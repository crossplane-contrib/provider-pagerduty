package slack

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("pagerduty_slack_connection", func(r *config.Resource) {

		r.ShortGroup = "slack"
	})
}
