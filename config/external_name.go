/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/upbound/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	"pagerduty_user":                   config.IdentifierFromProvider,
	"pagerduty_user_contact_method":    config.IdentifierFromProvider,
	"pagerduty_user_notification_rule": config.IdentifierFromProvider,
	"pagerduty_team":                   config.IdentifierFromProvider,
	"pagerduty_team_membership":        config.IdentifierFromProvider,
	"pagerduty_tag":                    config.IdentifierFromProvider,
	"pagerduty_tag_assignment":         config.IdentifierFromProvider,
	"pagerduty_webhook_subscription":   config.IdentifierFromProvider,
	"pagerduty_slack_connection":       config.IdentifierFromProvider,
	"pagerduty_service":                config.IdentifierFromProvider,
	"pagerduty_service_dependency":     config.IdentifierFromProvider,
	"pagerduty_service_event_rule":     config.IdentifierFromProvider,
	"pagerduty_service_integration":    config.IdentifierFromProvider,
	"pagerduty_escalation_policy":      config.IdentifierFromProvider,
	"pagerduty_schedule":               config.IdentifierFromProvider,
	"pagerduty_ruleset":                config.IdentifierFromProvider,
	"pagerduty_ruleset_rule":           config.IdentifierFromProvider,
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
