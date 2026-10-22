package config

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// An existing integration without email filters must not plan a replacement
// of the Optional+Computed+ForceNew email_filter block.
func TestNoOpComputedCollectionDiffDropped(t *testing.T) {
	r := GetProvider().Resources["pagerduty_service_integration"]
	state := &terraform.InstanceState{
		ID: "PFJKV51",
		Attributes: map[string]string{
			"id":      "PFJKV51",
			"name":    "Github",
			"service": "PDBF4FP",
			"vendor":  "PFZ9O9D",
		},
	}
	cfg := terraform.NewResourceConfigRaw(map[string]any{
		"name":    "Github",
		"service": "PDBF4FP",
		"vendor":  "PFZ9O9D",
	})

	d, err := schema.InternalMap(r.TerraformResource.Schema).Diff(context.Background(), state, cfg, r.TerraformResource.CustomizeDiff, nil, false)
	if err != nil {
		t.Fatal(err)
	}
	if d == nil || d.Attributes["email_filter.#"] == nil {
		t.Fatal("precondition: SDK no longer emits the email_filter.# no-op diff")
	}

	d, err = r.TerraformCustomDiff(d, state, cfg)
	if err != nil {
		t.Fatal(err)
	}
	for k, ad := range d.Attributes {
		if ad.RequiresNew {
			t.Errorf("%s still requires replacement: %#v", k, ad)
		}
	}

	// Create must keep the diff: nothing exists yet to keep.
	d, _ = schema.InternalMap(r.TerraformResource.Schema).Diff(context.Background(), nil, cfg, r.TerraformResource.CustomizeDiff, nil, false)
	if d, _ = r.TerraformCustomDiff(d, nil, cfg); d.Attributes["email_filter.#"] == nil {
		t.Error("create diff lost email_filter.#")
	}
}
