/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	addon "github.com/crossplane-contrib/provider-pagerduty/internal/controller/addon/addon"
	service "github.com/crossplane-contrib/provider-pagerduty/internal/controller/business/service"
	servicesubscriber "github.com/crossplane-contrib/provider-pagerduty/internal/controller/business/servicesubscriber"
	policy "github.com/crossplane-contrib/provider-pagerduty/internal/controller/escalation/policy"
	orchestration "github.com/crossplane-contrib/provider-pagerduty/internal/controller/event/orchestration"
	orchestrationrouter "github.com/crossplane-contrib/provider-pagerduty/internal/controller/event/orchestrationrouter"
	orchestrationservice "github.com/crossplane-contrib/provider-pagerduty/internal/controller/event/orchestrationservice"
	orchestrationunrouted "github.com/crossplane-contrib/provider-pagerduty/internal/controller/event/orchestrationunrouted"
	extension "github.com/crossplane-contrib/provider-pagerduty/internal/controller/extensions/extension"
	servicenow "github.com/crossplane-contrib/provider-pagerduty/internal/controller/extensions/servicenow"
	window "github.com/crossplane-contrib/provider-pagerduty/internal/controller/maintenance/window"
	providerconfig "github.com/crossplane-contrib/provider-pagerduty/internal/controller/providerconfig"
	play "github.com/crossplane-contrib/provider-pagerduty/internal/controller/response/play"
	rule "github.com/crossplane-contrib/provider-pagerduty/internal/controller/ruleset/rule"
	ruleset "github.com/crossplane-contrib/provider-pagerduty/internal/controller/ruleset/ruleset"
	schedule "github.com/crossplane-contrib/provider-pagerduty/internal/controller/schedule/schedule"
	dependency "github.com/crossplane-contrib/provider-pagerduty/internal/controller/service/dependency"
	eventrule "github.com/crossplane-contrib/provider-pagerduty/internal/controller/service/eventrule"
	integration "github.com/crossplane-contrib/provider-pagerduty/internal/controller/service/integration"
	serviceservice "github.com/crossplane-contrib/provider-pagerduty/internal/controller/service/service"
	connection "github.com/crossplane-contrib/provider-pagerduty/internal/controller/slack/connection"
	assignment "github.com/crossplane-contrib/provider-pagerduty/internal/controller/tag/assignment"
	tag "github.com/crossplane-contrib/provider-pagerduty/internal/controller/tag/tag"
	membership "github.com/crossplane-contrib/provider-pagerduty/internal/controller/team/membership"
	team "github.com/crossplane-contrib/provider-pagerduty/internal/controller/team/team"
	contactmethod "github.com/crossplane-contrib/provider-pagerduty/internal/controller/user/contactmethod"
	notificationrule "github.com/crossplane-contrib/provider-pagerduty/internal/controller/user/notificationrule"
	user "github.com/crossplane-contrib/provider-pagerduty/internal/controller/user/user"
	subscription "github.com/crossplane-contrib/provider-pagerduty/internal/controller/webhook/subscription"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		addon.Setup,
		service.Setup,
		servicesubscriber.Setup,
		policy.Setup,
		orchestration.Setup,
		orchestrationrouter.Setup,
		orchestrationservice.Setup,
		orchestrationunrouted.Setup,
		extension.Setup,
		servicenow.Setup,
		window.Setup,
		providerconfig.Setup,
		play.Setup,
		rule.Setup,
		ruleset.Setup,
		schedule.Setup,
		dependency.Setup,
		eventrule.Setup,
		integration.Setup,
		serviceservice.Setup,
		connection.Setup,
		assignment.Setup,
		tag.Setup,
		membership.Setup,
		team.Setup,
		contactmethod.Setup,
		notificationrule.Setup,
		user.Setup,
		subscription.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
