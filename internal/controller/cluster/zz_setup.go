/*
Copyright 2022 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	addon "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/addon/addon"
	groupingsetting "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/alert/groupingsetting"
	action "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/automation/action"
	actionserviceassociation "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/automation/actionserviceassociation"
	actionteamassociation "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/automation/actionteamassociation"
	runner "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/automation/runner"
	runnerteamassociation "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/automation/runnerteamassociation"
	service "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/business/service"
	servicesubscriber "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/business/servicesubscriber"
	policy "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/escalation/policy"
	orchestration "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/event/orchestration"
	orchestrationglobal "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/event/orchestrationglobal"
	orchestrationglobalcachevariable "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/event/orchestrationglobalcachevariable"
	orchestrationintegration "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/event/orchestrationintegration"
	orchestrationrouter "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/event/orchestrationrouter"
	orchestrationservice "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/event/orchestrationservice"
	orchestrationservicecachevariable "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/event/orchestrationservicecachevariable"
	orchestrationunrouted "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/event/orchestrationunrouted"
	rule "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/event/rule"
	extension "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/extensions/extension"
	servicenow "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/extensions/servicenow"
	customfield "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/incident/customfield"
	customfieldoption "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/incident/customfieldoption"
	incidenttype "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/incident/incidenttype"
	typecustomfield "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/incident/typecustomfield"
	workflow "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/incident/workflow"
	workflowtrigger "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/incident/workflowtrigger"
	cloudaccountmappingrule "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/jira/cloudaccountmappingrule"
	window "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/maintenance/window"
	providerconfig "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/providerconfig"
	play "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/response/play"
	ruleruleset "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/ruleset/rule"
	ruleset "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/ruleset/ruleset"
	schedule "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/schedule/schedule"
	dependency "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/service/dependency"
	eventrule "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/service/eventrule"
	integration "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/service/integration"
	serviceservice "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/service/service"
	connection "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/slack/connection"
	assignment "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/tag/assignment"
	tag "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/tag/tag"
	membership "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/team/membership"
	team "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/team/team"
	contactmethod "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/user/contactmethod"
	handoffnotificationrule "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/user/handoffnotificationrule"
	notificationrule "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/user/notificationrule"
	user "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/user/user"
	subscription "github.com/crossplane-contrib/provider-pagerduty/internal/controller/cluster/webhook/subscription"
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
