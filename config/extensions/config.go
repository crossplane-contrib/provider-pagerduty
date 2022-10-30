package extensions

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("pagerduty_extension", func(r *config.Resource) {
		r.ShortGroup = "extensions"
		r.References = config.References{
			"extension_objects": {
				Type: "github.com/crossplane-contrib/provider-pagerduty/apis/service/v1alpha1.Service",
			},
		}

	})

	p.AddResourceConfigurator("pagerduty_extension_servicenow", func(r *config.Resource) {
		r.ShortGroup = "extensions"
	})
}
