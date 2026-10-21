package config

import (
	"context"
	"sort"

	"github.com/PagerDuty/terraform-provider-pagerduty/pagerduty"
	pagerdutyplugin "github.com/PagerDuty/terraform-provider-pagerduty/pagerdutyplugin"
	fwprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	fwresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// newTerraformProviders returns fresh instances of the SDKv2 and Plugin
// Framework halves of the muxed PagerDuty Terraform provider. IsMuxed drops
// the SDKv2 resources that the Framework half re-implements, so every
// resource is served by exactly one of them.
func newTerraformProviders() (*schema.Provider, fwprovider.Provider) {
	sdk := pagerduty.Provider(pagerduty.IsMuxed)
	// The upstream provider shares package-level schema maps between
	// resources (e.g. eventOrchestrationPathConditionsSchema) and between
	// provider instances. upjet mutates resource schemas while configuring
	// them (server-side apply injected keys, overrides in config/*), so each
	// resource needs its own copy.
	for name, r := range sdk.ResourcesMap {
		sdk.ResourcesMap[name] = copyResource(r)
	}
	return sdk, pagerdutyplugin.New()
}

func copyResource(r *schema.Resource) *schema.Resource {
	c := *r
	if r.Schema != nil {
		c.Schema = make(map[string]*schema.Schema, len(r.Schema))
		for k, s := range r.Schema {
			c.Schema[k] = copySchema(s)
		}
	}
	return &c
}

func copySchema(s *schema.Schema) *schema.Schema {
	c := *s
	switch e := s.Elem.(type) {
	case *schema.Resource:
		c.Elem = copyResource(e)
	case *schema.Schema:
		c.Elem = copySchema(e)
	}
	return &c
}

// includeLists splits the external-name configured resources into the
// Terraform Plugin SDKv2 and Plugin Framework include lists, depending on
// which half of the muxed provider implements each resource.
func includeLists(sdk *schema.Provider, fw fwprovider.Provider) (sdkList, fwList []string) {
	fwNames := map[string]struct{}{}
	for _, f := range fw.Resources(context.Background()) {
		resp := &fwresource.MetadataResponse{}
		f().Metadata(context.Background(), fwresource.MetadataRequest{ProviderTypeName: resourcePrefix}, resp)
		fwNames[resp.TypeName] = struct{}{}
	}

	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		switch _, isFW := fwNames[name]; {
		case isFW:
			fwList = append(fwList, name+"$")
		case sdk.ResourcesMap[name] != nil:
			sdkList = append(sdkList, name+"$")
		}
	}
	sort.Strings(sdkList)
	sort.Strings(fwList)
	return sdkList, fwList
}
