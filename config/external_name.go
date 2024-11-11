/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	"pagerduty_addon":                                         config.IdentifierFromProvider,
	"pagerduty_alert_grouping_setting":                        config.IdentifierFromProvider,
	"pagerduty_automation_actions_action_service_association": config.IdentifierFromProvider,
	"pagerduty_automation_actions_action_team_association":    config.IdentifierFromProvider,
	"pagerduty_automation_actions_action":                     config.IdentifierFromProvider,
	"pagerduty_automation_actions_runner":                     config.IdentifierFromProvider,
	"pagerduty_automation_actions_runner_team_association":    config.IdentifierFromProvider,
	"pagerduty_business_service_subscriber":                   config.IdentifierFromProvider,
	"pagerduty_business_service":                              config.IdentifierFromProvider,
	"pagerduty_escalation_policy":                             config.IdentifierFromProvider,
	"pagerduty_event_orchestration_global_cache_variable":     config.IdentifierFromProvider,
	"pagerduty_event_orchestration_global":                    config.IdentifierFromProvider,
	"pagerduty_event_orchestration_integration":               config.IdentifierFromProvider,
	"pagerduty_event_orchestration_router":                    config.IdentifierFromProvider,
	"pagerduty_event_orchestration_service_cache_variable":    config.IdentifierFromProvider,
	"pagerduty_event_orchestration_service":                   config.IdentifierFromProvider,
	"pagerduty_event_orchestration_unrouted":                  config.IdentifierFromProvider,
	"pagerduty_event_orchestration":                           config.IdentifierFromProvider,
	"pagerduty_event_rule":                                    config.IdentifierFromProvider,
	"pagerduty_extension_servicenow":                          config.IdentifierFromProvider,
	"pagerduty_extension":                                     config.IdentifierFromProvider,
	"pagerduty_incident_custom_field_option":                  config.IdentifierFromProvider,
	"pagerduty_incident_custom_field":                         config.IdentifierFromProvider,
	"pagerduty_incident_workflow_trigger":                     config.IdentifierFromProvider,
	"pagerduty_incident_workflow":                             config.IdentifierFromProvider,
	"pagerduty_maintenance_window":                            config.IdentifierFromProvider,
	"pagerduty_response_play":                                 config.IdentifierFromProvider,
	"pagerduty_ruleset_rule":                                  config.IdentifierFromProvider,
	"pagerduty_ruleset":                                       config.IdentifierFromProvider,
	"pagerduty_schedule":                                      config.IdentifierFromProvider,
	"pagerduty_service_dependency":                            config.IdentifierFromProvider,
	"pagerduty_service_event_rule":                            config.IdentifierFromProvider,
	"pagerduty_service_integration":                           config.IdentifierFromProvider,
	"pagerduty_service":                                       config.IdentifierFromProvider,
	"pagerduty_slack_connection":                              config.IdentifierFromProvider,
	"pagerduty_tag_assignment":                                config.IdentifierFromProvider,
	"pagerduty_tag":                                           config.IdentifierFromProvider,
	"pagerduty_team_membership":                               config.IdentifierFromProvider,
	"pagerduty_team":                                          config.IdentifierFromProvider,
	"pagerduty_user_contact_method":                           config.IdentifierFromProvider,
	"pagerduty_user_handoff_notification_rule":                config.IdentifierFromProvider,
	"pagerduty_user_notification_rule":                        config.IdentifierFromProvider,
	"pagerduty_user":                                          config.IdentifierFromProvider,
	"pagerduty_webhook_subscription":                          config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
