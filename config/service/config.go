package service

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("pagerduty_service", func(r *config.Resource) {

		r.LateInitializer = config.LateInitializer{
			// alert_grouping_parameters and alert_grouping_timeout are mutually exclusive
			IgnoredFields: []string{"alert_grouping_parameters", "alert_grouping_timeout"},
		}
		r.ShortGroup = "service"
		r.References = config.References{
			"escalation_policy": {
				Type: "github.com/crossplane-contrib/provider-pagerduty/apis/escalation/v1alpha1.Policy",
			},
		}
	})

	p.AddResourceConfigurator("pagerduty_service_dependency", func(r *config.Resource) {

		r.ShortGroup = "service"
	})

	p.AddResourceConfigurator("pagerduty_service_event_rule", func(r *config.Resource) {

		r.ShortGroup = "service"
		r.References = config.References{
			"service": {
				Type: "Service",
			},
		}
	})

	p.AddResourceConfigurator("pagerduty_service_integration", func(r *config.Resource) {

		r.LateInitializer = config.LateInitializer{
			// type and vendor are mutually-exclusive, so these cannot be late-initialized
			IgnoredFields: []string{"type", "vendor"},
		}
		r.ShortGroup = "service"
		r.References = config.References{
			"service": {
				Type: "Service",
			},
		}

		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if a, ok := attr["integration_key"].(string); ok {
				conn["integration_key"] = []byte(a)
			}
			return conn, nil
		}
	})

}
