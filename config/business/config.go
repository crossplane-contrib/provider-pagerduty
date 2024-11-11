package business

import (
	"github.com/crossplane-contrib/provider-pagerduty/config/fake"
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("pagerduty_business_service", func(r *config.Resource) {

		r.ShortGroup = "business"
		r.References = config.References{
			"team": {
				Type: "github.com/crossplane-contrib/provider-pagerduty/apis/team/v1alpha1.Team",
			},
		}
		r.ExternalName.GetExternalNameFn = fake.GetExternalName
		r.ExternalName.GetIDFn = fake.GetID
		// Deprecated
		if s, ok := r.TerraformResource.Schema["type"]; ok {
			s.Optional = false
			s.Computed = true
		}
	})

	p.AddResourceConfigurator("pagerduty_business_service_subscriber", func(r *config.Resource) {

		r.ShortGroup = "business"
		r.References = config.References{
			"business_service_id": {
				Type: "Service",
			},
		}

	})
}
