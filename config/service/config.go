package slack

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("pagerduty_service", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
		r.ShortGroup = "service"
		r.References = config.References{
			"escalation_policy": {
				Type: "github.com/crossplane-contrib/provider-pagerduty/apis/escalation/v1alpha1.Policy",
			},
		}
	})

	p.AddResourceConfigurator("pagerduty_service_dependency", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
		r.ShortGroup = "service"
	})

	p.AddResourceConfigurator("pagerduty_service_event_rule", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
		r.ShortGroup = "service"
		r.References = config.References{
			"service": {
				Type: "Service",
			},
		}
	})

	p.AddResourceConfigurator("pagerduty_service_integration", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
		r.ShortGroup = "service"
		r.References = config.References{
			"service": {
				Type: "Service",
			},
		}
	})

}
