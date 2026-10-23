/*
Copyright 2021 Upbound Inc.
*/

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/crossplane/upjet/v2/pkg/pipeline"

	"github.com/crossplane-contrib/provider-pagerduty/config"
	"github.com/crossplane-contrib/provider-pagerduty/config/versions"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "" {
		panic("root directory is required to be given as argument")
	}
	rootDir := os.Args[1]
	absRootDir, err := filepath.Abs(rootDir)
	if err != nil {
		panic(fmt.Sprintf("cannot calculate the absolute path with %s", rootDir))
	}
	pipeline.Run(config.GetProvider(), config.GetProviderNamespaced(), absRootDir)
	if err := versions.ConvertExamples(filepath.Join(absRootDir, "examples-generated")); err != nil {
		panic(err)
	}
}
