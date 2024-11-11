package user

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("pagerduty_user", func(r *config.Resource) {

		r.ShortGroup = "user"
		// Managed by pagerduty_team_membership resource.
		if s, ok := r.TerraformResource.Schema["teams"]; ok {
			s.Optional = false
			s.Computed = true
		}
	})

	p.AddResourceConfigurator("pagerduty_user_notification_rule", func(r *config.Resource) {

		r.ShortGroup = "user"
		r.References = config.References{
			"user_id": {
				Type: "User",
			},
		}
	})

	p.AddResourceConfigurator("pagerduty_user_contact_method", func(r *config.Resource) {

		r.ShortGroup = "user"
		r.References = config.References{
			"user_id": {
				Type: "User",
			},
		}
	})

	p.AddResourceConfigurator("pagerduty_user_handoff_notification_rule", func(r *config.Resource) {

		r.ShortGroup = "user"
		r.References = config.References{
			"user_id": {
				Type: "User",
			},
		}
	})

}
