package ruleset

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("pagerduty_ruleset", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
		r.ShortGroup = "ruleset"
		r.References = config.References{
			"team.id": {
				Type: "github.com/crossplane-contrib/provider-pagerduty/apis/team/v1alpha1.Team",
			},
		}
	})

	p.AddResourceConfigurator("pagerduty_ruleset_rule", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
		r.ShortGroup = "ruleset"
		r.References = config.References{
			"ruleset": {
				Type: "github.com/crossplane-contrib/provider-pagerduty/apis/ruleset/v1alpha1.Ruleset",
			},
		}

	})
}
