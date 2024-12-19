package incident

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("pagerduty_incident_custom_field_option", func(r *config.Resource) {
		r.References = config.References{
			"field": {
				TerraformName: "pagerduty_incident_custom_field",
			},
		}
	})
	p.AddResourceConfigurator("pagerduty_incident_workflow", func(r *config.Resource) {
		r.References = config.References{
			"team": {
				TerraformName: "pagerduty_team",
			},
		}
	})
	p.AddResourceConfigurator("pagerduty_incident_workflow_trigger", func(r *config.Resource) {
		r.References = config.References{
			"permissions.team_id": {
				TerraformName: "pagerduty_team",
			},
			"service": {
				TerraformName: "pagerduty_service",
			},
			"workflow": {
				TerraformName: "pagerduty_incident_workflow",
			},
		}
	})

}
