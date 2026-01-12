/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/v2/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	"pagerduty_addon":                                         idWithStub(),
	"pagerduty_alert_grouping_setting":                        idWithStub(),
	"pagerduty_automation_actions_action_service_association": idWithStub(),
	"pagerduty_automation_actions_action_team_association":    idWithStub(),
	"pagerduty_automation_actions_action":                     idWithStub(),
	"pagerduty_automation_actions_runner":                     idWithStub(),
	"pagerduty_automation_actions_runner_team_association":    idWithStub(),
	"pagerduty_business_service_subscriber":                   idWithStub(),
	"pagerduty_business_service":                              idWithStub(),
	"pagerduty_enablement":                                    idWithStub(),
	"pagerduty_escalation_policy":                             idWithStub(),
	"pagerduty_event_orchestration_global_cache_variable":     idWithStub(),
	"pagerduty_event_orchestration_global":                    idWithStub(),
	"pagerduty_event_orchestration_integration":               idWithStub(),
	"pagerduty_event_orchestration_router":                    idWithStub(),
	"pagerduty_event_orchestration_service_cache_variable":    idWithStub(),
	"pagerduty_event_orchestration_service":                   idWithStub(),
	"pagerduty_event_orchestration_unrouted":                  idWithStub(),
	"pagerduty_event_orchestration":                           idWithStub(),
	"pagerduty_event_rule":                                    idWithStub(),
	"pagerduty_extension_servicenow":                          idWithStub(),
	"pagerduty_extension":                                     idWithStub(),
	"pagerduty_incident_custom_field_option":                  idWithStub(),
	"pagerduty_incident_custom_field":                         idWithStub(),
	"pagerduty_incident_type":                                 idWithStub(),
	"pagerduty_incident_type_custom_field":                    idWithStub(),
	"pagerduty_incident_workflow_trigger":                     idWithStub(),
	"pagerduty_incident_workflow":                             idWithStub(),
	"pagerduty_jira_cloud_account_mapping_rule":               idWithStub(),
	"pagerduty_maintenance_window":                            idWithStub(),
	"pagerduty_response_play":                                 idWithStub(),
	"pagerduty_ruleset_rule":                                  idWithStub(),
	"pagerduty_ruleset":                                       idWithStub(),
	"pagerduty_schedule":                                      idWithStub(),
	"pagerduty_service_custom_field":                          idWithStub(),
	"pagerduty_service_custom_field_value":                    idWithStub(),
	"pagerduty_service_dependency":                            idWithStub(),
	"pagerduty_service_event_rule":                            idWithStub(),
	"pagerduty_service_integration":                           idWithStub(),
	"pagerduty_service":                                       idWithStub(),
	"pagerduty_slack_connection":                              idWithStub(),
	"pagerduty_tag_assignment":                                idWithStub(),
	"pagerduty_tag":                                           idWithStub(),
	"pagerduty_team_membership":                               idWithStub(),
	"pagerduty_team":                                          idWithStub(),
	"pagerduty_user_contact_method":                           idWithStub(),
	"pagerduty_user_handoff_notification_rule":                idWithStub(),
	"pagerduty_user_notification_rule":                        idWithStub(),
	"pagerduty_user":                                          idWithStub(),
	"pagerduty_webhook_subscription":                          idWithStub(),
}

func idWithStub() config.ExternalName {
	e := config.IdentifierFromProvider
	e.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
		en, _ := config.IDAsExternalName(tfstate)
		return en, nil
	}
	return e
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
