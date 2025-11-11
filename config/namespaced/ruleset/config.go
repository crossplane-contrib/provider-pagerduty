package ruleset

import (
	c "github.com/crossplane-contrib/provider-pagerduty/config/namespaced/common"
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("pagerduty_ruleset", func(r *config.Resource) {

		r.ShortGroup = "ruleset"
		r.References = config.References{
			"team.id": {
				TerraformName: "pagerduty_team",
			},
		}
	})

	p.AddResourceConfigurator("pagerduty_ruleset_rule", func(r *config.Resource) {

		r.ShortGroup = "ruleset"
		r.References = config.References{
			"ruleset": {
				TerraformName: "pagerduty_ruleset",
			},
		}
		r.ExternalName.GetIDFn = c.SplitIdFromExternalName(':', 1)
		r.ExternalName.GetExternalNameFn = c.GetExternalNameFromTfstate([]string{"ruleset", "id"}, '.')
	})
}
