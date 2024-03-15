package tag

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("pagerduty_tag", func(r *config.Resource) {

		r.ShortGroup = "tag"
	})

	p.AddResourceConfigurator("pagerduty_tag_assignment", func(r *config.Resource) {

		r.ShortGroup = "tag"
	})
}
