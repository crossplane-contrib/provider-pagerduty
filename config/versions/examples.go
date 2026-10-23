package versions

import (
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/config/conversion"
	"github.com/pkg/errors"
	"sigs.k8s.io/yaml"
)

// exampleSeparator separates the manifests in the example files that upjet
// generates.
const exampleSeparator = "\n---\n\n"

var (
	mu       sync.Mutex
	embedded = map[*config.Resource][]string{}
)

func recordEmbedded(r *config.Resource, crdPaths []string) {
	mu.Lock()
	defer mu.Unlock()
	embedded[r] = crdPaths
}

// ConvertExamples converts the singleton lists of the resources configured
// with EmbedSingletonLists to embedded objects in the example manifests under
// dir. upjet generates the examples from the Terraform registry docs, where
// single nested blocks are lists.
func ConvertExamples(dir string) error {
	return filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() || filepath.Ext(path) != ".yaml" {
			return err
		}
		return convertExampleFile(path)
	})
}

func convertExampleFile(path string) error {
	b, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return errors.Wrapf(err, "cannot read example %s", path)
	}
	docs := strings.Split(string(b), exampleSeparator)
	changed := false
	for i, doc := range docs {
		var m map[string]any
		if err := yaml.Unmarshal([]byte(doc), &m); err != nil {
			return errors.Wrapf(err, "cannot parse example %s", path)
		}
		paths := embeddedPaths(m)
		if paths == nil {
			continue
		}
		spec, _ := m["spec"].(map[string]any)
		fp, ok := spec["forProvider"].(map[string]any)
		if !ok {
			continue
		}
		if _, err := conversion.Convert(fp, paths, conversion.ToEmbeddedObject, nil); err != nil {
			return errors.Wrapf(err, "cannot convert example %s", path)
		}
		out, err := yaml.Marshal(m)
		if err != nil {
			return errors.Wrapf(err, "cannot marshal example %s", path)
		}
		docs[i] = string(out)
		changed = true
	}
	if !changed {
		return nil
	}
	return errors.Wrapf(os.WriteFile(path, []byte(strings.Join(docs, exampleSeparator)), 0o600), "cannot write example %s", path) //nolint:gosec // generation only
}

// embeddedPaths returns the CRD paths to embed for the given manifest, or nil
// if it is not the current version of a resource configured with
// EmbedSingletonLists.
func embeddedPaths(m map[string]any) []string {
	apiVersion, _ := m["apiVersion"].(string)
	kind, _ := m["kind"].(string)
	group, version, ok := strings.Cut(apiVersion, "/")
	if !ok {
		return nil
	}
	mu.Lock()
	defer mu.Unlock()
	for r, paths := range embedded {
		if r.Kind == kind && r.Version == version && strings.HasPrefix(group, strings.ToLower(r.ShortGroup)+".") {
			return paths
		}
	}
	return nil
}
