package config

import (
	"testing"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Every configured resource must be reconciled in-process by exactly one of
// the SDKv2 or Framework clients. A resource left on the CLI client would
// need a terraform binary that the provider image no longer ships.
func TestProvidersReconcileEveryResourceInProcess(t *testing.T) {
	for name, get := range map[string]func() *ujconfig.Provider{
		"cluster":    GetProvider,
		"namespaced": GetProviderNamespaced,
	} {
		t.Run(name, func(t *testing.T) {
			pc := get()
			if len(pc.Resources) != len(ExternalNameConfigs) {
				t.Errorf("got %d resources, want %d (one per ExternalNameConfigs entry)", len(pc.Resources), len(ExternalNameConfigs))
			}
			for n, r := range pc.Resources {
				sdk, fw := r.ShouldUseTerraformPluginSDKClient(), r.ShouldUseTerraformPluginFrameworkClient()
				if sdk == fw {
					t.Errorf("%s: SDKv2 client %t, Framework client %t; want exactly one", n, sdk, fw)
				}
			}
		})
	}
}

func TestIncludeListsSplitByImplementation(t *testing.T) {
	sdk, fw := newTerraformProviders()
	sdkList, fwList := includeLists(sdk, fw)

	if got, want := len(sdkList)+len(fwList), len(ExternalNameConfigs); got != want {
		t.Errorf("SDKv2 + Framework lists hold %d resources, want %d", got, want)
	}
	// pagerduty_team moved to the Framework half; muxing removes it from the
	// SDKv2 half, so it must not be routed to the SDKv2 client.
	for _, tc := range []struct {
		list []string
		name string
	}{
		{fwList, "pagerduty_team$"},
		{sdkList, "pagerduty_service$"},
	} {
		if !contains(tc.list, tc.name) {
			t.Errorf("%q missing from its include list %v", tc.name, tc.list)
		}
	}
}

// The upstream provider shares schema maps between resources and provider
// instances. upjet injects server-side apply keys into them, so a shared map
// makes the second resource fail with "already contains the argument key".
func TestTerraformProvidersDoNotShareSchemas(t *testing.T) {
	a, _ := newTerraformProviders()
	b, _ := newTerraformProviders()

	elem := func(p *schema.Provider, res string) map[string]*schema.Schema {
		s := p.ResourcesMap[res].Schema["set"].Elem.(*schema.Resource).Schema["rule"].Elem.(*schema.Resource).Schema["condition"]
		return s.Elem.(*schema.Resource).Schema
	}

	elem(a, "pagerduty_event_orchestration_router")["index"] = &schema.Schema{Type: schema.TypeString}

	for _, tc := range []struct {
		p   *schema.Provider
		res string
	}{
		{a, "pagerduty_event_orchestration_service"},
		{b, "pagerduty_event_orchestration_router"},
	} {
		if _, ok := elem(tc.p, tc.res)["index"]; ok {
			t.Errorf("%s: mutation of another resource's schema leaked into set.rule.condition", tc.res)
		}
	}
}

func contains(l []string, s string) bool {
	for _, v := range l {
		if v == s {
			return true
		}
	}
	return false
}
