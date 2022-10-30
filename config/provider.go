/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"

	addon "github.com/crossplane-contrib/provider-pagerduty/config/addon"
	business "github.com/crossplane-contrib/provider-pagerduty/config/business"
	escalation "github.com/crossplane-contrib/provider-pagerduty/config/escalation"
	extensions "github.com/crossplane-contrib/provider-pagerduty/config/extensions"
	maintenance "github.com/crossplane-contrib/provider-pagerduty/config/maintenance"
	response "github.com/crossplane-contrib/provider-pagerduty/config/response"
	ruleset "github.com/crossplane-contrib/provider-pagerduty/config/ruleset"
	schedule "github.com/crossplane-contrib/provider-pagerduty/config/schedule"
	service "github.com/crossplane-contrib/provider-pagerduty/config/service"
	slack "github.com/crossplane-contrib/provider-pagerduty/config/slack"
	tag "github.com/crossplane-contrib/provider-pagerduty/config/tag"
	team "github.com/crossplane-contrib/provider-pagerduty/config/team"
	user "github.com/crossplane-contrib/provider-pagerduty/config/user"
	webhook "github.com/crossplane-contrib/provider-pagerduty/config/webhook"
)

const (
	resourcePrefix = "pagerduty"
	modulePath     = "github.com/crossplane-contrib/provider-pagerduty"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithShortName("pagerduty"),
		ujconfig.WithRootGroup("pagerduty.crossplane.io"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		addon.Configure,
		business.Configure,
		escalation.Configure,
		extensions.Configure,
		maintenance.Configure,
		response.Configure,
		ruleset.Configure,
		schedule.Configure,
		service.Configure,
		slack.Configure,
		team.Configure,
		tag.Configure,
		user.Configure,
		webhook.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
