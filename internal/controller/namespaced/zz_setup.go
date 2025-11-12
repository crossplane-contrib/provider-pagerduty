/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	addon "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/addon/addon"
	groupingsetting "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/alert/groupingsetting"
	action "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/automation/action"
	actionserviceassociation "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/automation/actionserviceassociation"
	actionteamassociation "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/automation/actionteamassociation"
	runner "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/automation/runner"
	runnerteamassociation "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/automation/runnerteamassociation"
	service "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/business/service"
	servicesubscriber "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/business/servicesubscriber"
	policy "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/escalation/policy"
	orchestration "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/event/orchestration"
	orchestrationglobal "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/event/orchestrationglobal"
	orchestrationglobalcachevariable "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/event/orchestrationglobalcachevariable"
	orchestrationintegration "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/event/orchestrationintegration"
	orchestrationrouter "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/event/orchestrationrouter"
	orchestrationservice "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/event/orchestrationservice"
	orchestrationservicecachevariable "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/event/orchestrationservicecachevariable"
	orchestrationunrouted "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/event/orchestrationunrouted"
	rule "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/event/rule"
	extension "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/extensions/extension"
	servicenow "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/extensions/servicenow"
	customfield "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/incident/customfield"
	customfieldoption "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/incident/customfieldoption"
	incidenttype "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/incident/incidenttype"
	typecustomfield "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/incident/typecustomfield"
	workflow "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/incident/workflow"
	workflowtrigger "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/incident/workflowtrigger"
	cloudaccountmappingrule "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/jira/cloudaccountmappingrule"
	window "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/maintenance/window"
	providerconfig "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/providerconfig"
	play "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/response/play"
	ruleruleset "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/ruleset/rule"
	ruleset "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/ruleset/ruleset"
	schedule "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/schedule/schedule"
	dependency "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/service/dependency"
	eventrule "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/service/eventrule"
	integration "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/service/integration"
	serviceservice "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/service/service"
	connection "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/slack/connection"
	assignment "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/tag/assignment"
	tag "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/tag/tag"
	membership "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/team/membership"
	team "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/team/team"
	contactmethod "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/user/contactmethod"
	handoffnotificationrule "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/user/handoffnotificationrule"
	notificationrule "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/user/notificationrule"
	user "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/user/user"
	subscription "github.com/crossplane-contrib/provider-pagerduty/internal/controller/namespaced/webhook/subscription"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		addon.Setup,
		groupingsetting.Setup,
		action.Setup,
		actionserviceassociation.Setup,
		actionteamassociation.Setup,
		runner.Setup,
		runnerteamassociation.Setup,
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
		incidenttype.Setup,
		typecustomfield.Setup,
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

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		addon.SetupGated,
		groupingsetting.SetupGated,
		action.SetupGated,
		actionserviceassociation.SetupGated,
		actionteamassociation.SetupGated,
		runner.SetupGated,
		runnerteamassociation.SetupGated,
		service.SetupGated,
		servicesubscriber.SetupGated,
		policy.SetupGated,
		orchestration.SetupGated,
		orchestrationglobal.SetupGated,
		orchestrationglobalcachevariable.SetupGated,
		orchestrationintegration.SetupGated,
		orchestrationrouter.SetupGated,
		orchestrationservice.SetupGated,
		orchestrationservicecachevariable.SetupGated,
		orchestrationunrouted.SetupGated,
		rule.SetupGated,
		extension.SetupGated,
		servicenow.SetupGated,
		customfield.SetupGated,
		customfieldoption.SetupGated,
		incidenttype.SetupGated,
		typecustomfield.SetupGated,
		workflow.SetupGated,
		workflowtrigger.SetupGated,
		cloudaccountmappingrule.SetupGated,
		window.SetupGated,
		providerconfig.SetupGated,
		play.SetupGated,
		ruleruleset.SetupGated,
		ruleset.SetupGated,
		schedule.SetupGated,
		dependency.SetupGated,
		eventrule.SetupGated,
		integration.SetupGated,
		serviceservice.SetupGated,
		connection.SetupGated,
		assignment.SetupGated,
		tag.SetupGated,
		membership.SetupGated,
		team.SetupGated,
		contactmethod.SetupGated,
		handoffnotificationrule.SetupGated,
		notificationrule.SetupGated,
		user.SetupGated,
		subscription.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
