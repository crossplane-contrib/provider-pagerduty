package event

import (
	c "github.com/crossplane-contrib/provider-pagerduty/config/cluster/common"
	"github.com/crossplane/upjet/v2/pkg/config"
)

const (
	shortGroup = "event"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	// The pagerduty_event_rule resource has been deprecated in favor of the
	// pagerduty_ruleset and pagerduty_ruleset_rule resources.
	// Please use the ruleset based resources for working with Event Rules.

	p.AddResourceConfigurator("pagerduty_event_orchestration_unrouted", func(r *config.Resource) {

		r.ShortGroup = shortGroup
		r.References = config.References{
			"event_orchestration": {
				TerraformName: "pagerduty_event_orchestration",
			},
		}
	})

	p.AddResourceConfigurator("pagerduty_event_orchestration_service", func(r *config.Resource) {

		r.ShortGroup = shortGroup
		r.References = config.References{
			"service": {
				TerraformName: "pagerduty_service",
			},
		}

	})

	p.AddResourceConfigurator("pagerduty_event_orchestration_router", func(r *config.Resource) {

		r.ShortGroup = shortGroup
		r.References = config.References{
			"event_orchestration": {
				TerraformName: "pagerduty_event_orchestration",
			},
		}
	})

	p.AddResourceConfigurator("pagerduty_event_orchestration", func(r *config.Resource) {

		r.ShortGroup = shortGroup
		r.References = config.References{
			"team": {
				TerraformName: "pagerduty_team",
			},
		}
	})

	p.AddResourceConfigurator("pagerduty_event_orchestration_global_cache_variable", func(r *config.Resource) {

		r.ShortGroup = shortGroup
		r.ExternalName.GetIDFn = c.SplitIdFromExternalName(':', 1)
		r.ExternalName.GetExternalNameFn = c.GetExternalNameFromTfstate([]string{"event_orchestration", "id"}, ':')
		r.References = config.References{
			"event_orchestration": {
				TerraformName: "pagerduty_event_orchestration",
			},
		}
	})

	p.AddResourceConfigurator("pagerduty_event_orchestration_global", func(r *config.Resource) {

		r.ShortGroup = shortGroup
		r.References = config.References{
			"event_orchestration": {
				TerraformName: "pagerduty_event_orchestration",
			},
		}
	})

	p.AddResourceConfigurator("pagerduty_event_orchestration_integration", func(r *config.Resource) {

		r.ShortGroup = shortGroup
		r.ExternalName.GetIDFn = c.SplitIdFromExternalName(':', 1)
		r.ExternalName.GetExternalNameFn = c.GetExternalNameFromTfstate([]string{"event_orchestration", "id"}, ':')
		r.References = config.References{
			"event_orchestration": {
				TerraformName: "pagerduty_event_orchestration",
			},
		}
	})

	p.AddResourceConfigurator("pagerduty_event_orchestration_service_cache_variable", func(r *config.Resource) {

		r.ShortGroup = shortGroup
		r.ExternalName.GetIDFn = c.SplitIdFromExternalName(':', 1)
		r.ExternalName.GetExternalNameFn = c.GetExternalNameFromTfstate([]string{"event_orchestration", "id"}, ':')
		r.References = config.References{
			"service": {
				TerraformName: "pagerduty_event_orchestration_service",
			},
		}
	})

}
