/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/terrajet/pkg/controller"

	service "github.com/crossplane-contrib/provider-jet-pagerduty/internal/controller/business/service"
	servicesubscriber "github.com/crossplane-contrib/provider-jet-pagerduty/internal/controller/business/servicesubscriber"
	policy "github.com/crossplane-contrib/provider-jet-pagerduty/internal/controller/escalation/policy"
	rule "github.com/crossplane-contrib/provider-jet-pagerduty/internal/controller/event/rule"
	servicenow "github.com/crossplane-contrib/provider-jet-pagerduty/internal/controller/extension/servicenow"
	window "github.com/crossplane-contrib/provider-jet-pagerduty/internal/controller/maintenance/window"
	addon "github.com/crossplane-contrib/provider-jet-pagerduty/internal/controller/pagerduty/addon"
	extension "github.com/crossplane-contrib/provider-jet-pagerduty/internal/controller/pagerduty/extension"
	ruleset "github.com/crossplane-contrib/provider-jet-pagerduty/internal/controller/pagerduty/ruleset"
	schedule "github.com/crossplane-contrib/provider-jet-pagerduty/internal/controller/pagerduty/schedule"
	servicepagerduty "github.com/crossplane-contrib/provider-jet-pagerduty/internal/controller/pagerduty/service"
	tag "github.com/crossplane-contrib/provider-jet-pagerduty/internal/controller/pagerduty/tag"
	team "github.com/crossplane-contrib/provider-jet-pagerduty/internal/controller/pagerduty/team"
	user "github.com/crossplane-contrib/provider-jet-pagerduty/internal/controller/pagerduty/user"
	providerconfig "github.com/crossplane-contrib/provider-jet-pagerduty/internal/controller/providerconfig"
	play "github.com/crossplane-contrib/provider-jet-pagerduty/internal/controller/response/play"
	ruleruleset "github.com/crossplane-contrib/provider-jet-pagerduty/internal/controller/ruleset/rule"
	dependency "github.com/crossplane-contrib/provider-jet-pagerduty/internal/controller/service/dependency"
	eventrule "github.com/crossplane-contrib/provider-jet-pagerduty/internal/controller/service/eventrule"
	integration "github.com/crossplane-contrib/provider-jet-pagerduty/internal/controller/service/integration"
	connection "github.com/crossplane-contrib/provider-jet-pagerduty/internal/controller/slack/connection"
	assignment "github.com/crossplane-contrib/provider-jet-pagerduty/internal/controller/tag/assignment"
	membership "github.com/crossplane-contrib/provider-jet-pagerduty/internal/controller/team/membership"
	contactmethod "github.com/crossplane-contrib/provider-jet-pagerduty/internal/controller/user/contactmethod"
	notificationrule "github.com/crossplane-contrib/provider-jet-pagerduty/internal/controller/user/notificationrule"
	subscription "github.com/crossplane-contrib/provider-jet-pagerduty/internal/controller/webhook/subscription"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		service.Setup,
		servicesubscriber.Setup,
		policy.Setup,
		rule.Setup,
		servicenow.Setup,
		window.Setup,
		addon.Setup,
		extension.Setup,
		ruleset.Setup,
		schedule.Setup,
		servicepagerduty.Setup,
		tag.Setup,
		team.Setup,
		user.Setup,
		providerconfig.Setup,
		play.Setup,
		ruleruleset.Setup,
		dependency.Setup,
		eventrule.Setup,
		integration.Setup,
		connection.Setup,
		assignment.Setup,
		membership.Setup,
		contactmethod.Setup,
		notificationrule.Setup,
		subscription.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
