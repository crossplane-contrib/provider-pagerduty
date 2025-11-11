/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	actionsCluster "github.com/crossplane-contrib/provider-pagerduty/config/cluster/actions"
	addonCluster "github.com/crossplane-contrib/provider-pagerduty/config/cluster/addon"
	alertgroupingCluster "github.com/crossplane-contrib/provider-pagerduty/config/cluster/alertgrouping"
	businessCluster "github.com/crossplane-contrib/provider-pagerduty/config/cluster/business"
	escalationCluster "github.com/crossplane-contrib/provider-pagerduty/config/cluster/escalation"
	eventCluster "github.com/crossplane-contrib/provider-pagerduty/config/cluster/event"
	extensionsCluster "github.com/crossplane-contrib/provider-pagerduty/config/cluster/extensions"
	incidentCluster "github.com/crossplane-contrib/provider-pagerduty/config/cluster/incident"
	jiraCluster "github.com/crossplane-contrib/provider-pagerduty/config/cluster/jira"
	maintenanceCluster "github.com/crossplane-contrib/provider-pagerduty/config/cluster/maintenance"
	responseCluster "github.com/crossplane-contrib/provider-pagerduty/config/cluster/response"
	rulesetCluster "github.com/crossplane-contrib/provider-pagerduty/config/cluster/ruleset"
	scheduleCluster "github.com/crossplane-contrib/provider-pagerduty/config/cluster/schedule"
	serviceCluster "github.com/crossplane-contrib/provider-pagerduty/config/cluster/service"
	slackCluster "github.com/crossplane-contrib/provider-pagerduty/config/cluster/slack"
	tagCluster "github.com/crossplane-contrib/provider-pagerduty/config/cluster/tag"
	teamCluster "github.com/crossplane-contrib/provider-pagerduty/config/cluster/team"
	userCluster "github.com/crossplane-contrib/provider-pagerduty/config/cluster/user"
	webhookCluster "github.com/crossplane-contrib/provider-pagerduty/config/cluster/webhook"

	actionsNamespaced "github.com/crossplane-contrib/provider-pagerduty/config/namespaced/actions"
	addonNamespaced "github.com/crossplane-contrib/provider-pagerduty/config/namespaced/addon"
	alertgroupingNamespaced "github.com/crossplane-contrib/provider-pagerduty/config/namespaced/alertgrouping"
	businessNamespaced "github.com/crossplane-contrib/provider-pagerduty/config/namespaced/business"
	escalationNamespaced "github.com/crossplane-contrib/provider-pagerduty/config/namespaced/escalation"
	eventNamespaced "github.com/crossplane-contrib/provider-pagerduty/config/namespaced/event"
	extensionsNamespaced "github.com/crossplane-contrib/provider-pagerduty/config/namespaced/extensions"
	incidentNamespaced "github.com/crossplane-contrib/provider-pagerduty/config/namespaced/incident"
	jiraNamespaced "github.com/crossplane-contrib/provider-pagerduty/config/namespaced/jira"
	maintenanceNamespaced "github.com/crossplane-contrib/provider-pagerduty/config/namespaced/maintenance"
	responseNamespaced "github.com/crossplane-contrib/provider-pagerduty/config/namespaced/response"
	rulesetNamespaced "github.com/crossplane-contrib/provider-pagerduty/config/namespaced/ruleset"
	scheduleNamespaced "github.com/crossplane-contrib/provider-pagerduty/config/namespaced/schedule"
	serviceNamespaced "github.com/crossplane-contrib/provider-pagerduty/config/namespaced/service"
	slackNamespaced "github.com/crossplane-contrib/provider-pagerduty/config/namespaced/slack"
	tagNamespaced "github.com/crossplane-contrib/provider-pagerduty/config/namespaced/tag"
	teamNamespaced "github.com/crossplane-contrib/provider-pagerduty/config/namespaced/team"
	userNamespaced "github.com/crossplane-contrib/provider-pagerduty/config/namespaced/user"
	webhookNamespaced "github.com/crossplane-contrib/provider-pagerduty/config/namespaced/webhook"
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
		addonCluster.Configure,
		actionsCluster.Configure,
		alertgroupingCluster.Configure,
		businessCluster.Configure,
		escalationCluster.Configure,
		eventCluster.Configure,
		extensionsCluster.Configure,
		incidentCluster.Configure,
		jiraCluster.Configure,
		maintenanceCluster.Configure,
		responseCluster.Configure,
		rulesetCluster.Configure,
		scheduleCluster.Configure,
		serviceCluster.Configure,
		slackCluster.Configure,
		teamCluster.Configure,
		tagCluster.Configure,
		userCluster.Configure,
		webhookCluster.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// GetProviderNamespaced returns provider configuration
func GetProviderNamespaced() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithShortName("pagerduty"),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithRootGroup("pagerduty.m.crossplane.io"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		addonNamespaced.Configure,
		actionsNamespaced.Configure,
		alertgroupingNamespaced.Configure,
		businessNamespaced.Configure,
		escalationNamespaced.Configure,
		eventNamespaced.Configure,
		extensionsNamespaced.Configure,
		incidentNamespaced.Configure,
		jiraNamespaced.Configure,
		maintenanceNamespaced.Configure,
		responseNamespaced.Configure,
		rulesetNamespaced.Configure,
		scheduleNamespaced.Configure,
		serviceNamespaced.Configure,
		slackNamespaced.Configure,
		teamNamespaced.Configure,
		tagNamespaced.Configure,
		userNamespaced.Configure,
		webhookNamespaced.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
