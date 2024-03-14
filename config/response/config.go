package response

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("pagerduty_response_play", func(r *config.Resource) {
		r.ShortGroup = "response"
	})
}
