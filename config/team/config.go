package team

import (
	"context"

	"github.com/crossplane/upjet/pkg/config"
	"github.com/pkg/errors"
)

const (
	ErrFmtNoAttribute    = `attribute not found: %s`
	ErrFmtUnexpectedType = `unexpected type for attribute %s: Expecting a string`
)

func getExternalName(tfstate map[string]any) (string, error) {
	id, ok := tfstate["id"]
	if !ok {
		return "", errors.Errorf(ErrFmtNoAttribute, "id")
	}
	idStr, ok := id.(string)
	if !ok {
		return "", errors.Errorf(ErrFmtUnexpectedType, "id")
	}
	return idStr, nil
}

func getID(ctx context.Context, externalName string, parameters map[string]any, providerConfig map[string]any) (string, error) {
	if externalName == "" {
		return "PENDING", nil
	}
	return externalName, nil
}

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("pagerduty_team", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ExternalName.GetExternalNameFn = getExternalName
		r.ExternalName.GetIDFn = getID
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
