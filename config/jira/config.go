package jira

import (
	c "github.com/crossplane-contrib/provider-pagerduty/config/common"
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("pagerduty_jira_cloud_account_mapping_rule", func(r *config.Resource) {
		r.ExternalName.GetExternalNameFn = c.GetExternalNameFromTfstate([]string{"account_mapping", "id"}, ':')
		r.ExternalName.GetIDFn = c.SplitIdFromExternalName(':', 1)
		r.References = config.References{
			"config.service": {
				TerraformName:     "pagerduty_service",
				RefFieldName:      "ServiceRefs",
				SelectorFieldName: "ServiceSelector",
			},
			"config.jira.sync_notes_user": {
				TerraformName:     "pagerduty_user",
				RefFieldName:      "UserRefs",
				SelectorFieldName: "UserSelector",
			},
		}
	})
}
