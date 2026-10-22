package config

import (
	"strings"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// NoOpComputedCollectionDiff drops the count diff the SDK emits for an
// Optional+Computed list, set or map that is empty in both the state and the
// config (e.g. pagerduty_service_integration.email_filter): Old "", New "",
// NewComputed. Terraform core plans the prior value for such attributes, but
// upjet's no-fork client applies the raw diff and, for ForceNew attributes,
// refuses the update as a replacement.
func NoOpComputedCollectionDiff() ujconfig.ResourceOption {
	return func(r *ujconfig.Resource) {
		r.TerraformCustomDiff = func(diff *terraform.InstanceDiff, state *terraform.InstanceState, _ *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
			if diff == nil || state == nil || state.ID == "" {
				return diff, nil
			}
			for k, ad := range diff.Attributes {
				if isNoOpComputedCount(k, ad) {
					delete(diff.Attributes, k)
				}
			}
			return diff, nil
		}
	}
}

func isNoOpComputedCount(k string, ad *terraform.ResourceAttrDiff) bool {
	return ad != nil && ad.NewComputed && !ad.NewRemoved && ad.Old == "" && ad.New == "" &&
		(strings.HasSuffix(k, ".#") || strings.HasSuffix(k, ".%"))
}
