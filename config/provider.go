/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/crossplane-contrib/provider-pagerduty/config/actions"
	"github.com/crossplane-contrib/provider-pagerduty/config/addon"
	"github.com/crossplane-contrib/provider-pagerduty/config/alertgrouping"
	"github.com/crossplane-contrib/provider-pagerduty/config/business"
	"github.com/crossplane-contrib/provider-pagerduty/config/escalation"
	"github.com/crossplane-contrib/provider-pagerduty/config/event"
	"github.com/crossplane-contrib/provider-pagerduty/config/extensions"
	"github.com/crossplane-contrib/provider-pagerduty/config/incident"
	"github.com/crossplane-contrib/provider-pagerduty/config/jira"
	"github.com/crossplane-contrib/provider-pagerduty/config/maintenance"
	"github.com/crossplane-contrib/provider-pagerduty/config/response"
	"github.com/crossplane-contrib/provider-pagerduty/config/ruleset"
	"github.com/crossplane-contrib/provider-pagerduty/config/schedule"
	"github.com/crossplane-contrib/provider-pagerduty/config/service"
	"github.com/crossplane-contrib/provider-pagerduty/config/slack"
	"github.com/crossplane-contrib/provider-pagerduty/config/tag"
	"github.com/crossplane-contrib/provider-pagerduty/config/team"
	"github.com/crossplane-contrib/provider-pagerduty/config/user"
	"github.com/crossplane-contrib/provider-pagerduty/config/webhook"
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
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithRootGroup("pagerduty.crossplane.io"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		addon.Configure,
		actions.Configure,
		alertgrouping.Configure,
		business.Configure,
		escalation.Configure,
		event.Configure,
		extensions.Configure,
		incident.Configure,
		jira.Configure,
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
