package jira

import (
	c "github.com/crossplane-contrib/provider-pagerduty/config/common"
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("pagerduty_jira_cloud_account_mapping_rule", func(r *config.Resource) {
		r.ExternalName.GetExternalNameFn = c.GetExternalNameFromTfstate([]string{"account_mapping_id", "rule_id"}, ':')
		r.References = config.References{
			"config.service": {
				TerraformName: "pagerduty_service",
			},
			"config.jira.sync_notes_user": {
				TerraformName: "pagerduty_user",
			},
		}
	})
}
