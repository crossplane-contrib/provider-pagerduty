package webhook

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("pagerduty_webhook_subscription", func(r *config.Resource) {

		r.ShortGroup = "webhook"
	})
}
