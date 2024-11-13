package commmon

import (
	"context"
	"strings"

	"github.com/crossplane/upjet/pkg/config"
	"github.com/pkg/errors"
)

// Errors
const (
	ErrFmtNoAttribute    = `attribute not found: %s`
	ErrFmtUnexpectedType = `unexpected type for attribute %s: Expecting a string`
	Managed              = `managed`
)

func GetExternalName(tfstate map[string]any) (string, error) {
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

func GetFakeID(_ context.Context, externalName string, _ map[string]any, _ map[string]any) (string, error) {
	if externalName == "" {
		return Managed, nil
	}
	return externalName, nil
}

func SplitExternalNameFromId() config.ExternalName {
	e := config.IdentifierFromProvider
	e.GetExternalNameFn = func(tfstate map[string]interface{}) (string, error) {
		id, ok := tfstate["id"]
		if !ok {
			return "", errors.New("id in tfstate cannot be empty")
		}
		w := strings.Split(id.(string), ":")
		return w[len(w)-1], nil
	}
	return e
}

func ExternalNameFromParams(params []string) config.ExternalName {
	e := config.IdentifierFromProvider
	e.GetExternalNameFn = func(tfstate map[string]interface{}) (string, error) {
		return strings.Join(params, ":"), nil
	}
	return e
}
