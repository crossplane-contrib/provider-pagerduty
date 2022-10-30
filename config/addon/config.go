package addon

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("pagerduty_addon", func(r *config.Resource) {

		r.ShortGroup = "addon"
	})
}
