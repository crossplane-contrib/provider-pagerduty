package jira

import (
	"context"
	"fmt"

	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("pagerduty_jira_cloud_account_mapping_rule", func(r *config.Resource) {
		r.ExternalName.GetIDFn = func(_ context.Context, externalName string, parameters map[string]any, _ map[string]any) (string, error) {
			return fmt.Sprintf("%s:%s", parameters["action_mapping"].(string), externalName), nil
		}
		r.References = config.References{
			"account_mapping": {
				Type:              "",
				RefFieldName:      "FieldRefs",
				SelectorFieldName: "FieldSelector",
			},
		}
	})
}
