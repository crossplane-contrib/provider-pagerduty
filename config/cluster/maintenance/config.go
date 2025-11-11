package maintenance

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("pagerduty_maintenance_window", func(r *config.Resource) {

		r.ShortGroup = "maintenance"
		r.References = config.References{
			"services": {
				TerraformName:     "pagerduty_service",
				RefFieldName:      "ServiceRefs",
				SelectorFieldName: "ServiceSelector",
			},
		}
	})
}
