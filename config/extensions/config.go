package extensions

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("pagerduty_extension", func(r *config.Resource) {
		r.ShortGroup = "extensions"
		r.References = config.References{
			"extension_objects": {
				TerraformName: "pagerduty_service",
			},
		}

	})

	p.AddResourceConfigurator("pagerduty_extension_servicenow", func(r *config.Resource) {
		r.ShortGroup = "extensions"
	})
}
