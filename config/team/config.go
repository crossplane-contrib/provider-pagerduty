package team

import (
	c "github.com/crossplane-contrib/provider-pagerduty/config/common"
	"github.com/crossplane/upjet/pkg/config"
)

const (
	shortGroup = "team"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("pagerduty_team", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ExternalName.GetExternalNameFn = c.GetExternalNameFromId
		r.ExternalName.GetIDFn = c.GetFakeID
		r.ShortGroup = shortGroup
	})

	p.AddResourceConfigurator("pagerduty_team_membership", func(r *config.Resource) {

		r.ShortGroup = shortGroup
		r.ExternalName.GetIDFn = c.GetIDFromParams([]string{"user_id", "team_id"}, ':')
		r.References = config.References{
			"user_id": {
				TerraformName: "pagerduty_user",
			},
			"team_id": {
				TerraformName: "pagerduty_team",
			},
		}
	})
}
