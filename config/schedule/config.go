package schedule

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("pagerduty_schedule", func(r *config.Resource) {

		r.ShortGroup = "schedule"
		r.References = config.References{
			"teams": {
				Type:              "github.com/crossplane-contrib/provider-pagerduty/apis/team/v1alpha1.Team",
				RefFieldName:      "TeamRefs",
				SelectorFieldName: "TeamSelector",
			},
			"layer.users": {
				Type:              "github.com/crossplane-contrib/provider-pagerduty/apis/user/v1alpha1.User",
				RefFieldName:      "UserRefs",
				SelectorFieldName: "UserSelector",
			},
		}
	})
}
