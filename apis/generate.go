//go:build generate
// +build generate

/*
Copyright 2021 Upbound Inc.
*/

// NOTE: See the below link for details on what is happening here.
// https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module

// Remove existing CRDs
//go:generate rm -rf ../package/crds

// Remove generated files. The types and Terraformed files of previous API
// versions (config.Resource.PreviousVersions) are kept: upjet loads them to
// generate the conversion spokes and never regenerates them. So is the group
// version info of previous versions that no longer have a current resource.
//go:generate bash -c "find . -iname 'zz_*' ! -iname 'zz_generated.managed*.go' ! -path '*/alert/v1alpha1/zz_groupingsetting_*' ! -path '*/jira/v1alpha1/zz_cloudaccountmappingrule_*' ! -path '*/user/v1alpha1/zz_notificationrule_*' ! -path '*/alert/v1alpha1/zz_groupversion_info.go' ! -path '*/jira/v1alpha1/zz_groupversion_info.go' -delete"
//go:generate bash -c "find . -type d -empty -delete"
//go:generate bash -c "find ../internal/controller -iname 'zz_*' -delete"
//go:generate bash -c "find ../internal/controller -type d -empty -delete"
//go:generate rm -rf ../examples-generated

// Generate documentation from Terraform docs.
//go:generate go run github.com/crossplane/upjet/v2/cmd/scraper -n ${TERRAFORM_PROVIDER_SOURCE} -r ../.work/${TERRAFORM_PROVIDER_SOURCE}/${TERRAFORM_DOCS_PATH} -o ../config/provider-metadata.yaml

// Run Upjet generator
//go:generate go run ../cmd/generator/main.go ..

// Generate deepcopy methodsets and CRD manifests
//go:generate go run -tags generate sigs.k8s.io/controller-tools/cmd/controller-gen object:headerFile=../hack/boilerplate.go.txt paths=./... crd:allowDangerousTypes=true,crdVersions=v1 output:artifacts:config=../package/crds

// Enable webhook conversion in the CRDs that serve more than one version
//go:generate go run ../hack/crdconversion ../package/crds

// Generate crossplane-runtime methodsets (resource.Claim, etc)
//go:generate go run -tags generate github.com/crossplane/crossplane-tools/cmd/angryjet generate-methodsets --header-file=../hack/boilerplate.go.txt ./...

package apis

import (
	_ "sigs.k8s.io/controller-tools/cmd/controller-gen" //nolint:typecheck

	_ "github.com/crossplane/crossplane-tools/cmd/angryjet" //nolint:typecheck
)
