package incident

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("pagerduty_incident_custom_field_option", func(r *config.Resource) {
		r.References = config.References{
			"field": {
				Type:              "CustomField",
				RefFieldName:      "FieldRefs",
				SelectorFieldName: "FieldSelector",
			},
		}
	})
	p.AddResourceConfigurator("pagerduty_incident_workflow", func(r *config.Resource) {
		r.References = config.References{
			"team": {
				Type:              "github.com/crossplane-contrib/provider-pagerduty/apis/team/v1alpha1.Team",
				RefFieldName:      "TeamRefs",
				SelectorFieldName: "TeamSelector",
			},
		}
	})
	p.AddResourceConfigurator("pagerduty_incident_workflow_trigger", func(r *config.Resource) {
		r.References = config.References{
			"permissions.team_id": {
				Type: "github.com/crossplane-contrib/provider-pagerduty/apis/team/v1alpha1.Team",
			},
			"service": {
				Type:              "github.com/crossplane-contrib/provider-pagerduty/apis/service/v1alpha1.Service",
				RefFieldName:      "ServiceRefs",
				SelectorFieldName: "ServiceSelector",
			},
			"workflow": {
				Type:              "Workflow",
				RefFieldName:      "WorkflowRefs",
				SelectorFieldName: "WorkflowSelector",
			},
		}
	})

}
