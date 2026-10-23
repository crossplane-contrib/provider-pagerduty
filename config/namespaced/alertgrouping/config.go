package alertgrouping

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/crossplane-contrib/provider-pagerduty/config/versions"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("pagerduty_alert_grouping_setting", func(r *config.Resource) {
		versions.EmbedSingletonLists(r, "config")
		r.References = config.References{
			"services": {
				TerraformName:     "pagerduty_service",
				RefFieldName:      "ServiceRefs",
				SelectorFieldName: "ServiceSelector",
			},
		}
	})
}
