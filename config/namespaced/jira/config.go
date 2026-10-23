package jira

import (
	c "github.com/crossplane-contrib/provider-pagerduty/config/namespaced/common"
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/crossplane-contrib/provider-pagerduty/config/versions"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("pagerduty_jira_cloud_account_mapping_rule", func(r *config.Resource) {
		versions.EmbedSingletonLists(r,
			"config", "config[*].jira", "config[*].jira[*].issueType", "config[*].jira[*].project",
			"config[*].jira[*].statusMapping", "config[*].jira[*].statusMapping[*].acknowledged",
			"config[*].jira[*].statusMapping[*].resolved", "config[*].jira[*].statusMapping[*].triggered")
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
