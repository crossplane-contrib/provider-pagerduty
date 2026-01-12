package enablement

import (
	c "github.com/crossplane-contrib/provider-pagerduty/config/cluster/common"
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("pagerduty_enablement", func(r *config.Resource) {
		r.Kind = "Enablement"
		r.ExternalName.GetIDFn = c.GetIDFromParams([]string{"runner_id", "team_id"}, ':')
		r.References = config.References{
			"entity_id": {
				TerraformName:     "pagerduty_team",
				RefFieldName:      "TeamRefs",
				SelectorFieldName: "TeamSelector",
			},
		}
	})
}
