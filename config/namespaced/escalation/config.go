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

		// rule is a required list with no user-provided id (Terraform-computed);
		// inject a key so SSA patches resolving teams do not atomically remove it.
		r.ServerSideApplyMergeStrategies["rule"] = config.MergeStrategy{
			ListMergeStrategy: config.ListMergeStrategy{
				MergeStrategy: config.ListTypeMap,
				ListMapKeys: config.ListMapKeys{
					InjectedKey: config.InjectedKey{
						Key:          "index",
						DefaultValue: `"0"`,
					},
				},
			},
		}
	})
}
