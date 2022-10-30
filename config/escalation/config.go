package slack

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("pagerduty_escalation_policy", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
		r.ShortGroup = "escalation"
		r.References = config.References{
			"teams": {
				Type:              "github.com/crossplane-contrib/provider-pagerduty/apis/team/v1alpha1.Team",
				RefFieldName:      "TeamRefs",
				SelectorFieldName: "TeamSelector",
			},
		}
	})
}
