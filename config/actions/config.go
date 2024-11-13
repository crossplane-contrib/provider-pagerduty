package actions

import (
	c "github.com/crossplane-contrib/provider-pagerduty/config/common"
	"github.com/crossplane/upjet/pkg/config"
)

const (
	ShortGroup = "automation.actions"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("pagerduty_automation_actions_action", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Action"
	})
	p.AddResourceConfigurator("pagerduty_automation_actions_action_service_association", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "ActionServiceAssociation"
		r.ExternalName = c.ExternalNameFromParams([]string{"action_id", "service_id"})
		r.References = config.References{
			"action_id": {
				Type:              "Action",
				RefFieldName:      "ActionRefs",
				SelectorFieldName: "ActionSelector",
			},
			"service_id": {
				Type:              "github.com/crossplane-contrib/provider-pagerduty/apis/service/v1alpha1.Service",
				RefFieldName:      "ServiceRefs",
				SelectorFieldName: "ServiceSelector",
			},
		}
	})
	p.AddResourceConfigurator("pagerduty_automation_actions_action_team_association", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "ActionTeamAssociation"
		r.ExternalName = c.ExternalNameFromParams([]string{"action_id", "team_id"})
		r.References = config.References{
			"action_id": {
				Type:              "Action",
				RefFieldName:      "ActionRefs",
				SelectorFieldName: "ActionSelector",
			},
			"team_id": {
				Type:              "github.com/crossplane-contrib/provider-pagerduty/apis/team/v1alpha1.Team",
				RefFieldName:      "TeamRefs",
				SelectorFieldName: "TeamSelector",
			},
		}
	})
	p.AddResourceConfigurator("pagerduty_automation_actions_runner", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "Runner"
	})
	p.AddResourceConfigurator("pagerduty_automation_actions_runner_team_association", func(r *config.Resource) {
		r.ShortGroup = ShortGroup
		r.Kind = "RunnerTeamAssociation"
		r.ExternalName = c.ExternalNameFromParams([]string{"runner_id", "team_id"})
		r.References = config.References{
			"runner_id": {
				Type:              "github.com/crossplane-contrib/provider-pagerduty/apis/automation/v1alpha1.Runner",
				RefFieldName:      "RunnerRefs",
				SelectorFieldName: "RunnerSelector",
			},
			"team_id": {
				Type:              "github.com/crossplane-contrib/provider-pagerduty/apis/team/v1alpha1.Team",
				RefFieldName:      "TeamRefs",
				SelectorFieldName: "TeamSelector",
			},
		}
	})
}
