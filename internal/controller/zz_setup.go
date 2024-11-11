/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	addon "github.com/crossplane-contrib/provider-pagerduty/internal/controller/addon/addon"
	groupingsetting "github.com/crossplane-contrib/provider-pagerduty/internal/controller/alert/groupingsetting"
	action "github.com/crossplane-contrib/provider-pagerduty/internal/controller/automation/action"
	runner "github.com/crossplane-contrib/provider-pagerduty/internal/controller/automation/runner"
	runnerteamassociation "github.com/crossplane-contrib/provider-pagerduty/internal/controller/automation/runnerteamassociation"
	serviceassociation "github.com/crossplane-contrib/provider-pagerduty/internal/controller/automation/serviceassociation"
	teamassociation "github.com/crossplane-contrib/provider-pagerduty/internal/controller/automation/teamassociation"
	service "github.com/crossplane-contrib/provider-pagerduty/internal/controller/business/service"
	servicesubscriber "github.com/crossplane-contrib/provider-pagerduty/internal/controller/business/servicesubscriber"
	policy "github.com/crossplane-contrib/provider-pagerduty/internal/controller/escalation/policy"
	orchestration "github.com/crossplane-contrib/provider-pagerduty/internal/controller/event/orchestration"
	orchestrationglobal "github.com/crossplane-contrib/provider-pagerduty/internal/controller/event/orchestrationglobal"
	orchestrationglobalcachevariable "github.com/crossplane-contrib/provider-pagerduty/internal/controller/event/orchestrationglobalcachevariable"
	orchestrationintegration "github.com/crossplane-contrib/provider-pagerduty/internal/controller/event/orchestrationintegration"
	orchestrationrouter "github.com/crossplane-contrib/provider-pagerduty/internal/controller/event/orchestrationrouter"
	orchestrationservice "github.com/crossplane-contrib/provider-pagerduty/internal/controller/event/orchestrationservice"
	orchestrationservicecachevariable "github.com/crossplane-contrib/provider-pagerduty/internal/controller/event/orchestrationservicecachevariable"
	orchestrationunrouted "github.com/crossplane-contrib/provider-pagerduty/internal/controller/event/orchestrationunrouted"
	rule "github.com/crossplane-contrib/provider-pagerduty/internal/controller/event/rule"
	extension "github.com/crossplane-contrib/provider-pagerduty/internal/controller/extensions/extension"
	servicenow "github.com/crossplane-contrib/provider-pagerduty/internal/controller/extensions/servicenow"
	customfield "github.com/crossplane-contrib/provider-pagerduty/internal/controller/incident/customfield"
	customfieldoption "github.com/crossplane-contrib/provider-pagerduty/internal/controller/incident/customfieldoption"
	workflow "github.com/crossplane-contrib/provider-pagerduty/internal/controller/incident/workflow"
	workflowtrigger "github.com/crossplane-contrib/provider-pagerduty/internal/controller/incident/workflowtrigger"
	cloudaccountmappingrule "github.com/crossplane-contrib/provider-pagerduty/internal/controller/jira/cloudaccountmappingrule"
	window "github.com/crossplane-contrib/provider-pagerduty/internal/controller/maintenance/window"
	providerconfig "github.com/crossplane-contrib/provider-pagerduty/internal/controller/providerconfig"
	play "github.com/crossplane-contrib/provider-pagerduty/internal/controller/response/play"
	ruleruleset "github.com/crossplane-contrib/provider-pagerduty/internal/controller/ruleset/rule"
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
	handoffnotificationrule "github.com/crossplane-contrib/provider-pagerduty/internal/controller/user/handoffnotificationrule"
	notificationrule "github.com/crossplane-contrib/provider-pagerduty/internal/controller/user/notificationrule"
	user "github.com/crossplane-contrib/provider-pagerduty/internal/controller/user/user"
	subscription "github.com/crossplane-contrib/provider-pagerduty/internal/controller/webhook/subscription"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		addon.Setup,
		groupingsetting.Setup,
		action.Setup,
		runner.Setup,
		runnerteamassociation.Setup,
		serviceassociation.Setup,
		teamassociation.Setup,
		service.Setup,
		servicesubscriber.Setup,
		policy.Setup,
		orchestration.Setup,
		orchestrationglobal.Setup,
		orchestrationglobalcachevariable.Setup,
		orchestrationintegration.Setup,
		orchestrationrouter.Setup,
		orchestrationservice.Setup,
		orchestrationservicecachevariable.Setup,
		orchestrationunrouted.Setup,
		rule.Setup,
		extension.Setup,
		servicenow.Setup,
		customfield.Setup,
		customfieldoption.Setup,
		workflow.Setup,
		workflowtrigger.Setup,
		cloudaccountmappingrule.Setup,
		window.Setup,
		providerconfig.Setup,
		play.Setup,
		ruleruleset.Setup,
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
		handoffnotificationrule.Setup,
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
