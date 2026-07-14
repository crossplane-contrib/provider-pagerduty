package schedule

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("pagerduty_schedule", func(r *config.Resource) {

		r.ShortGroup = "schedule"
		r.References = config.References{
			"teams": {
				TerraformName:     "pagerduty_team",
				RefFieldName:      "TeamRefs",
				SelectorFieldName: "TeamSelector",
			},
			"layer.users": {
				TerraformName:     "pagerduty_user",
				RefFieldName:      "UserRefs",
				SelectorFieldName: "UserSelector",
			},
		}

		r.ServerSideApplyMergeStrategies["layer"] = config.MergeStrategy{
			ListMergeStrategy: config.ListMergeStrategy{
				MergeStrategy: config.ListTypeAtomic,
			},
		}
		r.ServerSideApplyMergeStrategies["layer.users"] = config.MergeStrategy{
			ListMergeStrategy: config.ListMergeStrategy{
				MergeStrategy: config.ListTypeAtomic,
			},
		}
		r.ServerSideApplyMergeStrategies["layer.restriction"] = config.MergeStrategy{
			ListMergeStrategy: config.ListMergeStrategy{
				MergeStrategy: config.ListTypeAtomic,
			},
		}
		r.ServerSideApplyMergeStrategies["teams"] = config.MergeStrategy{
			ListMergeStrategy: config.ListMergeStrategy{
				MergeStrategy: config.ListTypeAtomic,
			},
		}
	})
}
