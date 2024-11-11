package fake

import (
	"context"

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

func GetID(_ context.Context, externalName string, _ map[string]any, _ map[string]any) (string, error) {
	if externalName == "" {
		return Managed, nil
	}
	return externalName, nil
}
