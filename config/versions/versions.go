// Package versions configures the API versions of resources that changed
// shape between releases.
package versions

import (
	"slices"
	"strings"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/config/conversion"
)

const (
	v1alpha1 = "v1alpha1"
	v1alpha2 = "v1alpha2"
)

// EmbedSingletonLists serves the singleton-list shape of the given CRD paths
// as v1alpha1 and the embedded-object shape, which upjet generates for Plugin
// Framework single nested blocks, as v1alpha2.
//
// v1alpha2 is the hub and is reconciled by the controller. v1alpha1 stays the
// storage version for now, so that rolling back to a release that only serves
// v1alpha1 keeps working. The v1alpha1 types are not regenerated (see
// apis/generate.go).
//
// crdPaths are camelCase CRD field paths relative to spec.forProvider,
// spec.initProvider and status.atProvider, written in their singleton-list
// form, with all ancestors of nested paths listed, e.g. "config" and
// "config[*].jira".
func EmbedSingletonLists(r *config.Resource, crdPaths ...string) {
	recordEmbedded(r, crdPaths)
	r.Version = v1alpha2
	r.PreviousVersions = []string{v1alpha1}
	r.SetCRDStorageVersion(v1alpha1)
	prefixes := conversion.DefaultPathPrefixes()
	r.Conversions = []conversion.Conversion{
		conversion.NewIdentityConversionExpandPaths(conversion.AllVersions, conversion.AllVersions, prefixes, crdPaths...),
		conversion.NewSingletonListConversion(v1alpha1, v1alpha2, prefixes, crdPaths, conversion.ToEmbeddedObject),
		conversion.NewSingletonListConversion(v1alpha2, v1alpha1, prefixes, crdPaths, conversion.ToSingletonList),
	}
	if err := r.SetDeprecatedVersion(v1alpha1, config.VersionDeprecation{
		Warning:            r.Kind + " " + v1alpha1 + " is deprecated, use " + v1alpha2 + ": its singleton lists are embedded objects.",
		DeprecationRelease: "v1.2.2",
	}); err != nil {
		panic(err)
	}
}

// RegisterPreviousVersions registers the API packages of previous versions
// that no resource is generated in anymore, e.g. alert/v1alpha1 once
// GroupingSetting moved to v1alpha2. upjet only registers the packages of
// current versions, and the conversion webhook needs the previous ones in the
// scheme. It must run after the resources are configured.
func RegisterPreviousVersions(pc *config.Provider) {
	current := map[string]bool{}
	for _, r := range pc.Resources {
		current[apiPackage(r.ShortGroup, r.Version)] = true
	}
	var previous []string
	for _, r := range pc.Resources {
		for _, v := range r.PreviousVersions {
			p := apiPackage(r.ShortGroup, v)
			if !current[p] && !slices.Contains(previous, p) && !slices.Contains(pc.BasePackages.APIVersion, p) {
				previous = append(previous, p)
			}
		}
	}
	slices.Sort(previous)
	// Copy, as the slice is shared with config.DefaultBasePackages.
	pc.BasePackages.APIVersion = append(slices.Clone(pc.BasePackages.APIVersion), previous...)
}

func apiPackage(shortGroup, version string) string {
	return strings.ToLower(shortGroup) + "/" + version
}
