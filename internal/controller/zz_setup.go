/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	policy "github.com/crossplane-contrib/provider-pagerduty/internal/controller/escalation/policy"
	providerconfig "github.com/crossplane-contrib/provider-pagerduty/internal/controller/providerconfig"
	rule "github.com/crossplane-contrib/provider-pagerduty/internal/controller/ruleset/rule"
	ruleset "github.com/crossplane-contrib/provider-pagerduty/internal/controller/ruleset/ruleset"
	schedule "github.com/crossplane-contrib/provider-pagerduty/internal/controller/schedule/schedule"
	dependency "github.com/crossplane-contrib/provider-pagerduty/internal/controller/service/dependency"
	eventrule "github.com/crossplane-contrib/provider-pagerduty/internal/controller/service/eventrule"
	integration "github.com/crossplane-contrib/provider-pagerduty/internal/controller/service/integration"
	service "github.com/crossplane-contrib/provider-pagerduty/internal/controller/service/service"
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
		policy.Setup,
		providerconfig.Setup,
		rule.Setup,
		ruleset.Setup,
		schedule.Setup,
		dependency.Setup,
		eventrule.Setup,
		integration.Setup,
		service.Setup,
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
