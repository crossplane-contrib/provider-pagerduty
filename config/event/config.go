package event

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	// The pagerduty_event_rule resource has been deprecated in favor of the
	// pagerduty_ruleset and pagerduty_ruleset_rule resources.
	// Please use the ruleset based resources for working with Event Rules.

	p.AddResourceConfigurator("pagerduty_event_orchestration_unrouted", func(r *config.Resource) {

		r.ShortGroup = "event"
		r.References = config.References{
			"event_orchestration": {
				Type: "Orchestration",
			},
		}
	})

	p.AddResourceConfigurator("pagerduty_event_orchestration_service", func(r *config.Resource) {

		r.ShortGroup = "event"
		r.References = config.References{
			"service": {
				Type: "github.com/crossplane-contrib/provider-pagerduty/apis/service/v1alpha1.Service",
			},
		}

	})

	p.AddResourceConfigurator("pagerduty_event_orchestration_router", func(r *config.Resource) {

		r.ShortGroup = "event"
		r.References = config.References{
			"event_orchestration": {
				Type: "Orchestration",
			},
		}
	})

	p.AddResourceConfigurator("pagerduty_event_orchestration", func(r *config.Resource) {

		r.ShortGroup = "event"
		r.References = config.References{
			"team": {
				Type: "github.com/crossplane-contrib/provider-pagerduty/apis/team/v1alpha1.Team",
			},
		}
	})

}
