package commmon

import (
	"context"
	"strings"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/pkg/errors"
)

// Errors
const (
	ErrFmtNoAttribute    = `attribute not found: %s`
	ErrFmtUnexpectedType = `unexpected type for attribute %s: Expecting a string`
	fakeId               = `managed`
)

func getStringFromParams(parameters map[string]interface{}, keys []string, separator rune) (string, error) {
	fields := []string{}
	for _, p := range keys {
		p, ok := parameters[p]
		if !ok {
			return "", errors.Errorf(ErrFmtNoAttribute, p)
		}
		pStr, ok := p.(string)
		if !ok {
			return "", errors.Errorf(ErrFmtUnexpectedType, p)
		}
		fields = append(fields, pStr)
	}
	return strings.Join(fields, string(separator)), nil
}

func GetExternalNameFromId(tfstate map[string]any) (string, error) {
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
		return fakeId, nil
	}
	return externalName, nil
}

func SplitIdFromExternalName(separator rune, position int) config.GetIDFn {
	return func(_ context.Context, externalName string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
		return strings.Split(externalName, string(separator))[position], nil
	}
}

func GetIDFromParams(params []string, separator rune) config.GetIDFn {
	return func(_ context.Context, externalName string, parameters map[string]interface{}, _ map[string]interface{}) (string, error) {
		return getStringFromParams(parameters, params, separator)
	}
}

func GetExternalNameFromTfstate(fields []string, separator rune) config.GetExternalNameFn {
	return func(tfstate map[string]any) (string, error) {
		return getStringFromParams(tfstate, fields, separator)
	}
}
