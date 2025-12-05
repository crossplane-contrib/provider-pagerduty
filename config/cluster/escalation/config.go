package escalation

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("pagerduty_escalation_policy", func(r *config.Resource) {

		r.ShortGroup = "escalation"
		r.References = config.References{
			"teams": {
				TerraformName:     "pagerduty_team",
				RefFieldName:      "TeamRefs",
				SelectorFieldName: "TeamSelector",
			},
		}
	})
}
