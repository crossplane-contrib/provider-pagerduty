package team

import (
	c "github.com/crossplane-contrib/provider-pagerduty/config/common"
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("pagerduty_team", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ExternalName.GetExternalNameFn = c.GetExternalName
		r.ExternalName.GetIDFn = c.GetFakeID
		r.ShortGroup = "team"
	})

	p.AddResourceConfigurator("pagerduty_team_membership", func(r *config.Resource) {

		r.ShortGroup = "team"
		r.References = config.References{
			"user_id": {
				Type: "github.com/crossplane-contrib/provider-pagerduty/apis/user/v1alpha1.User",
			},
			"team_id": {
				Type: "Team",
			},
		}
	})
}
