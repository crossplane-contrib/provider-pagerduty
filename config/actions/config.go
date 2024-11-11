package actions

import (
	"context"
	"fmt"

	"github.com/crossplane/upjet/pkg/config"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("pagerduty_automation_actions_action", func(r *config.Resource) {
		r.ShortGroup = "automation.actions"
		r.Kind = "Action"
	})
	p.AddResourceConfigurator("pagerduty_automation_actions_action_service_association", func(r *config.Resource) {
		r.ShortGroup = "automation.actions"
		r.Kind = "ServiceAssociation"
		r.ExternalName.GetIDFn = func(_ context.Context, externalName string, parameters map[string]any, _ map[string]any) (string, error) {
			return fmt.Sprintf("%s:%s", parameters["action_id"].(string), parameters["service_id"].(string)), nil
		}
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
		r.ShortGroup = "automation.actions"
		r.Kind = "TeamAssociation"
		r.ExternalName.GetIDFn = func(_ context.Context, externalName string, parameters map[string]any, _ map[string]any) (string, error) {
			return fmt.Sprintf("%s:%s", parameters["action_id"].(string), parameters["team_id"].(string)), nil
		}
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
		r.ShortGroup = "automation.actions"
		r.Kind = "Runner"
	})
	p.AddResourceConfigurator("pagerduty_automation_actions_runner_team_association", func(r *config.Resource) {
		r.ShortGroup = "automation.actions"
		r.Kind = "RunnerTeamAssociation"
		r.ExternalName.GetIDFn = func(_ context.Context, externalName string, parameters map[string]any, _ map[string]any) (string, error) {
			return fmt.Sprintf("%s:%s", parameters["runner_id"].(string), parameters["team_id"].(string)), nil
		}
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
