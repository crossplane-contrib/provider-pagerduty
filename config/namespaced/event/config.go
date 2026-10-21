package event

import (
	c "github.com/crossplane-contrib/provider-pagerduty/config/namespaced/common"
	"github.com/crossplane/upjet/v2/pkg/config"
)

const (
	shortGroup = "event"

	eventOrchestrationResource = "pagerduty_event_orchestration"
	eventOrchestrationField    = "event_orchestration"
)

// injectedIndexMergeStrategy merges an object list by a synthetic "index"
// key, defaulted to "0", injected for server-side apply.
func injectedIndexMergeStrategy() config.MergeStrategy {
	return config.MergeStrategy{
		ListMergeStrategy: config.ListMergeStrategy{
			MergeStrategy: config.ListTypeMap,
			ListMapKeys: config.ListMapKeys{
				InjectedKey: config.InjectedKey{
					Key:          "index",
					DefaultValue: `"0"`,
				},
			},
		},
	}
}

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {

	// The pagerduty_event_rule resource has been deprecated in favor of the
	// pagerduty_ruleset and pagerduty_ruleset_rule resources.
	// Please use the ruleset based resources for working with Event Rules.

	p.AddResourceConfigurator("pagerduty_event_orchestration_unrouted", func(r *config.Resource) {

		r.ShortGroup = shortGroup
		r.References = config.References{
			eventOrchestrationField: {
				TerraformName: eventOrchestrationResource,
			},
		}

		r.ServerSideApplyMergeStrategies["set"] = injectedIndexMergeStrategy()
		r.ServerSideApplyMergeStrategies["catch_all"] = injectedIndexMergeStrategy()
	})

	p.AddResourceConfigurator("pagerduty_event_orchestration_service", func(r *config.Resource) {

		r.ShortGroup = shortGroup
		r.References = config.References{
			"service": {
				TerraformName: "pagerduty_service",
			},
		}

		r.ServerSideApplyMergeStrategies["set"] = injectedIndexMergeStrategy()
		r.ServerSideApplyMergeStrategies["catch_all"] = injectedIndexMergeStrategy()
	})

	p.AddResourceConfigurator("pagerduty_event_orchestration_router", func(r *config.Resource) {

		r.ShortGroup = shortGroup
		r.References = config.References{
			eventOrchestrationField: {
				TerraformName: eventOrchestrationResource,
			},
			"set.rule.actions.route_to": {
				TerraformName:     "pagerduty_service",
				RefFieldName:      "RouteToServiceRef",
				SelectorFieldName: "RouteToServiceSelector",
			},
			"catch_all.actions.route_to": {
				TerraformName:     "pagerduty_service",
				RefFieldName:      "RouteToServiceRef",
				SelectorFieldName: "RouteToServiceSelector",
			},
		}

		// set has a natural "id" map key; use map-merge so SSA patches that only
		// write back resolved reference values do not atomically remove this
		// required field, which would fail XValidation.
		r.ServerSideApplyMergeStrategies["set"] = injectedIndexMergeStrategy()

		// set[].rule has no natural key (id is Terraform-computed); inject one
		// so SSA can target individual rule items when patching back routeTo.
		r.ServerSideApplyMergeStrategies["set.rule"] = injectedIndexMergeStrategy()

		// set[].rule[].actions (routeTo lives here)
		r.ServerSideApplyMergeStrategies["set.rule.actions"] = injectedIndexMergeStrategy()

		// set[].rule[].condition uses expression as the natural map key.
		r.ServerSideApplyMergeStrategies["set.rule.condition"] = injectedIndexMergeStrategy()

		// set[].rule[].actions[].dynamic_route_to is effectively a singleton;
		// inject a synthetic key so SSA does not atomically replace it.
		r.ServerSideApplyMergeStrategies["set.rule.actions.dynamic_route_to"] = injectedIndexMergeStrategy()

		// catch_all has no natural key; inject a synthetic one.
		r.ServerSideApplyMergeStrategies["catch_all"] = injectedIndexMergeStrategy()

		// catch_all[].actions (routeTo also lives here)
		r.ServerSideApplyMergeStrategies["catch_all.actions"] = injectedIndexMergeStrategy()
	})

	p.AddResourceConfigurator(eventOrchestrationResource, func(r *config.Resource) {

		r.ShortGroup = shortGroup
		r.References = config.References{
			"team": {
				TerraformName: "pagerduty_team",
			},
		}
	})

	p.AddResourceConfigurator("pagerduty_event_orchestration_global_cache_variable", func(r *config.Resource) {

		r.ShortGroup = shortGroup
		r.ExternalName.GetIDFn = c.SplitIdFromExternalName(':', 1)
		r.ExternalName.GetExternalNameFn = c.GetExternalNameFromTfstate([]string{eventOrchestrationField, "id"}, ':')
		r.References = config.References{
			eventOrchestrationField: {
				TerraformName: eventOrchestrationResource,
			},
		}
	})

	p.AddResourceConfigurator("pagerduty_event_orchestration_global", func(r *config.Resource) {

		r.ShortGroup = shortGroup
		r.References = config.References{
			eventOrchestrationField: {
				TerraformName: eventOrchestrationResource,
			},
		}

		r.ServerSideApplyMergeStrategies["set"] = injectedIndexMergeStrategy()
		r.ServerSideApplyMergeStrategies["catch_all"] = injectedIndexMergeStrategy()
	})

	p.AddResourceConfigurator("pagerduty_event_orchestration_integration", func(r *config.Resource) {

		r.ShortGroup = shortGroup
		r.ExternalName.GetIDFn = c.SplitIdFromExternalName(':', 1)
		r.ExternalName.GetExternalNameFn = c.GetExternalNameFromTfstate([]string{eventOrchestrationField, "id"}, ':')
		r.References = config.References{
			eventOrchestrationField: {
				TerraformName: eventOrchestrationResource,
			},
		}
	})

	p.AddResourceConfigurator("pagerduty_event_orchestration_service_cache_variable", func(r *config.Resource) {

		r.ShortGroup = shortGroup
		r.ExternalName.GetIDFn = c.SplitIdFromExternalName(':', 1)
		r.ExternalName.GetExternalNameFn = c.GetExternalNameFromTfstate([]string{eventOrchestrationField, "id"}, ':')
		r.References = config.References{
			"service": {
				TerraformName: "pagerduty_event_orchestration_service",
			},
		}
	})

}
