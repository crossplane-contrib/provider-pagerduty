package ruleset

import (
	"context"
	"fmt"

	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("pagerduty_ruleset", func(r *config.Resource) {

		r.ShortGroup = "ruleset"
		r.References = config.References{
			"team.id": {
				Type: "github.com/crossplane-contrib/provider-pagerduty/apis/team/v1alpha1.Team",
			},
		}
	})

	p.AddResourceConfigurator("pagerduty_ruleset_rule", func(r *config.Resource) {

		r.ShortGroup = "ruleset"
		r.References = config.References{
			"ruleset": {
				Type: "github.com/crossplane-contrib/provider-pagerduty/apis/ruleset/v1alpha1.Ruleset",
			},
		}
		r.ExternalName.GetIDFn = func(_ context.Context, externalName string, parameters map[string]any, _ map[string]any) (string, error) {
			return fmt.Sprintf("%s:%s", parameters["ruleset"].(string), externalName), nil
		}
	})
}
