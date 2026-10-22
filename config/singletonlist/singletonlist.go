// Package singletonlist keeps Terraform Plugin Framework single nested blocks
// as singleton lists in the generated CRDs.
package singletonlist

import (
	"fmt"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/config/conversion"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Keep generates the top-level single nested blocks at the given Terraform
// paths as singleton lists, the shape upjet generated before it mapped single
// nested blocks to embedded objects, so the CRD API does not change. At
// runtime the lists are converted to the objects the Plugin Framework
// resource expects, and back.
func Keep(r *config.Resource, paths ...string) {
	for _, p := range paths {
		s := config.GetSchema(r.TerraformResource, p)
		if s == nil {
			panic(fmt.Sprintf("singletonlist: %s has no Terraform schema at path %q", r.Name, p))
		}
		s.Type, s.MaxItems = schema.TypeList, 1
	}
	r.TerraformConversions = append(r.TerraformConversions, listToObject(paths))
}

// listToObject converts between the CRD singleton lists and the Terraform
// objects at the given paths: the reverse of upjet's NewTFSingletonConversion.
type listToObject []string

func (paths listToObject) Convert(params map[string]any, _ *config.Resource, mode config.Mode) (map[string]any, error) {
	if mode == config.ToTerraform {
		return conversion.Convert(params, paths, conversion.ToEmbeddedObject, nil)
	}
	return conversion.Convert(params, paths, conversion.ToSingletonList, nil)
}
