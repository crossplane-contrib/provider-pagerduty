package config

import (
	"context"
	"encoding/json"
	"testing"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	fwresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
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

// GroupingSetting keeps config as a singleton list in its CRD while the
// Framework resource defines it as a single nested block, so the params
// must be converted before upjet encodes them with the Framework schema.
func TestGroupingSettingConfigEncodesAsFrameworkObject(t *testing.T) {
	ctx := context.Background()
	for name, get := range map[string]func() *ujconfig.Provider{
		"cluster":    GetProvider,
		"namespaced": GetProviderNamespaced,
	} {
		t.Run(name, func(t *testing.T) {
			r := get().Resources["pagerduty_alert_grouping_setting"]
			resp := &fwresource.SchemaResponse{}
			r.TerraformPluginFrameworkResource.Schema(ctx, fwresource.SchemaRequest{}, resp)
			tfType := resp.Schema.Type().TerraformType(ctx)

			params, err := r.ApplyTFConversions(map[string]any{
				"name":     "n",
				"type":     "time",
				"services": []any{"P1"},
				"config":   []any{map[string]any{"timeout": 5}},
			}, ujconfig.ToTerraform)
			if err != nil {
				t.Fatalf("ApplyTFConversions(ToTerraform): %v", err)
			}
			b, err := json.Marshal(params)
			if err != nil {
				t.Fatal(err)
			}
			if _, err := tftypes.ValueFromJSONWithOpts(b, tfType, tftypes.ValueFromJSONOpts{IgnoreUndefinedAttributes: true}); err != nil {
				t.Errorf("params do not match the Framework schema: %v", err)
			}

			state, err := r.ApplyTFConversions(map[string]any{"config": map[string]any{"timeout": 5}}, ujconfig.FromTerraform)
			if err != nil {
				t.Fatalf("ApplyTFConversions(FromTerraform): %v", err)
			}
			if _, ok := state["config"].([]any); !ok {
				t.Errorf("state config = %T, want a singleton list", state["config"])
			}
		})
	}
}
