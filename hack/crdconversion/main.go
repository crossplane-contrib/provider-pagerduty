// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// crdconversion enables webhook conversion in the CRDs that serve more than
// one version. Crossplane fills in the service and the CA bundle of the
// conversion webhook when it installs the package.
package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"sigs.k8s.io/yaml"
)

const conversion = `  conversion:
    strategy: Webhook
    webhook:
      clientConfig:
        service:
          path: /convert
      conversionReviewVersions:
      - v1
`

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: crdconversion <crd directory>")
		os.Exit(1)
	}
	files, err := filepath.Glob(filepath.Join(os.Args[1], "*.yaml"))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	for _, f := range files {
		if err := enable(f); err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", f, err)
			os.Exit(1)
		}
	}
}

func enable(file string) error {
	b, err := os.ReadFile(filepath.Clean(file)) //nolint:gosec // generation only
	if err != nil {
		return err
	}
	var crd struct {
		Spec struct {
			Conversion any   `json:"conversion"`
			Versions   []any `json:"versions"`
		} `json:"spec"`
	}
	if err := yaml.Unmarshal(b, &crd); err != nil {
		return err
	}
	if len(crd.Spec.Versions) < 2 || crd.Spec.Conversion != nil {
		return nil
	}
	// controller-gen writes "spec:" once, as a top-level key.
	out := bytes.Replace(b, []byte("\nspec:\n"), []byte("\nspec:\n"+conversion), 1)
	if bytes.Equal(out, b) {
		return fmt.Errorf("cannot find the spec of the CRD")
	}
	return os.WriteFile(file, out, 0o600) //nolint:gosec // generation only
}
