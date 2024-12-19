package actions

import (
	c "github.com/crossplane-contrib/provider-pagerduty/config/common"
	"github.com/crossplane/upjet/pkg/config"
)

const (
	shortGroup = "automation.actions"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("pagerduty_automation_actions_action", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Action"
	})
	p.AddResourceConfigurator("pagerduty_automation_actions_action_service_association", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ActionServiceAssociation"
		r.ExternalName.GetIDFn = c.GetIDFromParams([]string{"action_id", "service_id"}, ':')
		r.References = config.References{
			"action_id": {
				Type: "Action", // TerraformName fails to resolve with shortGroup including dots
			},
			"service_id": {
				TerraformName: "pagerduty_service",
			},
		}
	})
	p.AddResourceConfigurator("pagerduty_automation_actions_action_team_association", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ActionTeamAssociation"
		r.ExternalName.GetIDFn = c.GetIDFromParams([]string{"action_id", "team_id"}, ':')
		r.References = config.References{
			"action_id": {
				Type: "Action", // TerraformName fails to resolve with shortGroup including dots
			},
			"team_id": {
				TerraformName: "pagerduty_team",
			},
		}
	})
	p.AddResourceConfigurator("pagerduty_automation_actions_runner", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Runner"
	})
	p.AddResourceConfigurator("pagerduty_automation_actions_runner_team_association", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "RunnerTeamAssociation"
		r.ExternalName.GetIDFn = c.GetIDFromParams([]string{"runner_id", "team_id"}, ':')
		r.References = config.References{
			"action_id": {
				Type: "Runner", // TerraformName fails to resolve with shortGroup including dots
			},
			"team_id": {
				TerraformName: "pagerduty_team",
			},
		}
	})
}
