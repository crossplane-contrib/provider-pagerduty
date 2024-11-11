/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AcknowledgedInitParameters struct {

	// Unique identifier for the Jira status.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the Jira status.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type AcknowledgedObservation struct {

	// Unique identifier for the Jira status.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the Jira status.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type AcknowledgedParameters struct {

	// Unique identifier for the Jira status.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the Jira status.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type CloudAccountMappingRuleInitParameters struct {

	// [Updating can cause a resource replacement] The account mapping this rule belongs to.
	// +crossplane:generate:reference:refFieldName=FieldRefs
	// +crossplane:generate:reference:selectorFieldName=FieldSelector
	AccountMapping *string `json:"accountMapping,omitempty" tf:"account_mapping,omitempty"`

	// Configuration for bidirectional synchronization between Jira issues and PagerDuty incidents.
	Config []ConfigInitParameters `json:"config,omitempty" tf:"config,omitempty"`

	// Reference to a  to populate accountMapping.
	// +kubebuilder:validation:Optional
	FieldRefs *v1.Reference `json:"fieldRefs,omitempty" tf:"-"`

	// Selector for a  to populate accountMapping.
	// +kubebuilder:validation:Optional
	FieldSelector *v1.Selector `json:"fieldSelector,omitempty" tf:"-"`
}

type CloudAccountMappingRuleObservation struct {

	// [Updating can cause a resource replacement] The account mapping this rule belongs to.
	AccountMapping *string `json:"accountMapping,omitempty" tf:"account_mapping,omitempty"`

	// If auto-creation using JQL is disabled, this field provides the reason for the disablement.
	AutocreateJqlDisabledReason *string `json:"autocreateJqlDisabledReason,omitempty" tf:"autocreate_jql_disabled_reason,omitempty"`

	// The timestamp until which the auto-creation using JQL feature is disabled.
	AutocreateJqlDisabledUntil *string `json:"autocreateJqlDisabledUntil,omitempty" tf:"autocreate_jql_disabled_until,omitempty"`

	// Configuration for bidirectional synchronization between Jira issues and PagerDuty incidents.
	Config []ConfigObservation `json:"config,omitempty" tf:"config,omitempty"`

	// The ID of the service.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type CloudAccountMappingRuleParameters struct {

	// [Updating can cause a resource replacement] The account mapping this rule belongs to.
	// +crossplane:generate:reference:refFieldName=FieldRefs
	// +crossplane:generate:reference:selectorFieldName=FieldSelector
	// +kubebuilder:validation:Optional
	AccountMapping *string `json:"accountMapping,omitempty" tf:"account_mapping,omitempty"`

	// Configuration for bidirectional synchronization between Jira issues and PagerDuty incidents.
	// +kubebuilder:validation:Optional
	Config []ConfigParameters `json:"config,omitempty" tf:"config,omitempty"`

	// Reference to a  to populate accountMapping.
	// +kubebuilder:validation:Optional
	FieldRefs *v1.Reference `json:"fieldRefs,omitempty" tf:"-"`

	// Selector for a  to populate accountMapping.
	// +kubebuilder:validation:Optional
	FieldSelector *v1.Selector `json:"fieldSelector,omitempty" tf:"-"`
}

type ConfigInitParameters struct {

	// Synchronization settings.
	Jira []JiraInitParameters `json:"jira,omitempty" tf:"jira,omitempty"`

	// [Updating can cause a resource replacement] The ID of the linked PagerDuty service.
	Service *string `json:"service,omitempty" tf:"service,omitempty"`
}

type ConfigObservation struct {

	// Synchronization settings.
	Jira []JiraObservation `json:"jira,omitempty" tf:"jira,omitempty"`

	// [Updating can cause a resource replacement] The ID of the linked PagerDuty service.
	Service *string `json:"service,omitempty" tf:"service,omitempty"`
}

type ConfigParameters struct {

	// Synchronization settings.
	// +kubebuilder:validation:Optional
	Jira []JiraParameters `json:"jira" tf:"jira,omitempty"`

	// [Updating can cause a resource replacement] The ID of the linked PagerDuty service.
	// +kubebuilder:validation:Optional
	Service *string `json:"service" tf:"service,omitempty"`
}

type CustomFieldsInitParameters struct {

	// The PagerDuty incident field from which the value will be extracted (only applicable if type is attribute); one of incident_number, incident_title, incident_description, incident_status, incident_created_at, incident_service, incident_escalation_policy, incident_impacted_services, incident_html_url, incident_assignees, incident_acknowledgers, incident_last_status_change_at, incident_last_status_change_by, incident_urgency or incident_priority.
	SourceIncidentField *string `json:"sourceIncidentField,omitempty" tf:"source_incident_field,omitempty"`

	// The unique identifier key of the Jira field that will be set.
	TargetIssueField *string `json:"targetIssueField,omitempty" tf:"target_issue_field,omitempty"`

	// The human-readable name of the Jira field.
	TargetIssueFieldName *string `json:"targetIssueFieldName,omitempty" tf:"target_issue_field_name,omitempty"`

	// The type of the value that will be set; one of attribute, const or jira_value.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The value to be set for the Jira field (only applicable if type is const or jira_value). It must be set as a JSON string.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type CustomFieldsObservation struct {

	// The PagerDuty incident field from which the value will be extracted (only applicable if type is attribute); one of incident_number, incident_title, incident_description, incident_status, incident_created_at, incident_service, incident_escalation_policy, incident_impacted_services, incident_html_url, incident_assignees, incident_acknowledgers, incident_last_status_change_at, incident_last_status_change_by, incident_urgency or incident_priority.
	SourceIncidentField *string `json:"sourceIncidentField,omitempty" tf:"source_incident_field,omitempty"`

	// The unique identifier key of the Jira field that will be set.
	TargetIssueField *string `json:"targetIssueField,omitempty" tf:"target_issue_field,omitempty"`

	// The human-readable name of the Jira field.
	TargetIssueFieldName *string `json:"targetIssueFieldName,omitempty" tf:"target_issue_field_name,omitempty"`

	// The type of the value that will be set; one of attribute, const or jira_value.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The value to be set for the Jira field (only applicable if type is const or jira_value). It must be set as a JSON string.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type CustomFieldsParameters struct {

	// The PagerDuty incident field from which the value will be extracted (only applicable if type is attribute); one of incident_number, incident_title, incident_description, incident_status, incident_created_at, incident_service, incident_escalation_policy, incident_impacted_services, incident_html_url, incident_assignees, incident_acknowledgers, incident_last_status_change_at, incident_last_status_change_by, incident_urgency or incident_priority.
	// +kubebuilder:validation:Optional
	SourceIncidentField *string `json:"sourceIncidentField,omitempty" tf:"source_incident_field,omitempty"`

	// The unique identifier key of the Jira field that will be set.
	// +kubebuilder:validation:Optional
	TargetIssueField *string `json:"targetIssueField" tf:"target_issue_field,omitempty"`

	// The human-readable name of the Jira field.
	// +kubebuilder:validation:Optional
	TargetIssueFieldName *string `json:"targetIssueFieldName" tf:"target_issue_field_name,omitempty"`

	// The type of the value that will be set; one of attribute, const or jira_value.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`

	// The value to be set for the Jira field (only applicable if type is const or jira_value). It must be set as a JSON string.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type IssueTypeInitParameters struct {

	// Unique identifier for the Jira status.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the Jira status.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type IssueTypeObservation struct {

	// Unique identifier for the Jira status.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the Jira status.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type IssueTypeParameters struct {

	// Unique identifier for the Jira status.
	// +kubebuilder:validation:Optional
	ID *string `json:"id" tf:"id,omitempty"`

	// Name of the Jira status.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

type JiraInitParameters struct {

	// JQL query to automatically create PagerDuty incidents when matching Jira issues are created. Leave empty to disable this feature.
	AutocreateJql *string `json:"autocreateJql,omitempty" tf:"autocreate_jql,omitempty"`

	// When enabled, automatically creates a Jira issue whenever a PagerDuty incident is triggered.
	CreateIssueOnIncidentTrigger *bool `json:"createIssueOnIncidentTrigger,omitempty" tf:"create_issue_on_incident_trigger,omitempty"`

	// Defines how Jira fields are populated when a Jira Issue is created from a PagerDuty Incident.
	CustomFields []CustomFieldsInitParameters `json:"customFields,omitempty" tf:"custom_fields,omitempty"`

	// Specifies the Jira issue type to be created or synchronized with PagerDuty incidents.
	IssueType []IssueTypeInitParameters `json:"issueType,omitempty" tf:"issue_type,omitempty"`

	// Maps PagerDuty incident priorities to Jira issue priorities for synchronization.
	Priorities []PrioritiesInitParameters `json:"priorities,omitempty" tf:"priorities,omitempty"`

	// [Updating can cause a resource replacement] Defines the Jira project where issues will be created or synchronized.
	Project []ProjectInitParameters `json:"project,omitempty" tf:"project,omitempty"`

	// Maps PagerDuty incident statuses to corresponding Jira issue statuses for synchronization.
	StatusMapping []StatusMappingInitParameters `json:"statusMapping,omitempty" tf:"status_mapping,omitempty"`

	// ID of the PagerDuty user for syncing notes and comments between Jira issues and PagerDuty incidents. If not provided, note synchronization is disabled.
	SyncNotesUser *string `json:"syncNotesUser,omitempty" tf:"sync_notes_user,omitempty"`
}

type JiraObservation struct {

	// JQL query to automatically create PagerDuty incidents when matching Jira issues are created. Leave empty to disable this feature.
	AutocreateJql *string `json:"autocreateJql,omitempty" tf:"autocreate_jql,omitempty"`

	// When enabled, automatically creates a Jira issue whenever a PagerDuty incident is triggered.
	CreateIssueOnIncidentTrigger *bool `json:"createIssueOnIncidentTrigger,omitempty" tf:"create_issue_on_incident_trigger,omitempty"`

	// Defines how Jira fields are populated when a Jira Issue is created from a PagerDuty Incident.
	CustomFields []CustomFieldsObservation `json:"customFields,omitempty" tf:"custom_fields,omitempty"`

	// Specifies the Jira issue type to be created or synchronized with PagerDuty incidents.
	IssueType []IssueTypeObservation `json:"issueType,omitempty" tf:"issue_type,omitempty"`

	// Maps PagerDuty incident priorities to Jira issue priorities for synchronization.
	Priorities []PrioritiesObservation `json:"priorities,omitempty" tf:"priorities,omitempty"`

	// [Updating can cause a resource replacement] Defines the Jira project where issues will be created or synchronized.
	Project []ProjectObservation `json:"project,omitempty" tf:"project,omitempty"`

	// Maps PagerDuty incident statuses to corresponding Jira issue statuses for synchronization.
	StatusMapping []StatusMappingObservation `json:"statusMapping,omitempty" tf:"status_mapping,omitempty"`

	// ID of the PagerDuty user for syncing notes and comments between Jira issues and PagerDuty incidents. If not provided, note synchronization is disabled.
	SyncNotesUser *string `json:"syncNotesUser,omitempty" tf:"sync_notes_user,omitempty"`
}

type JiraParameters struct {

	// JQL query to automatically create PagerDuty incidents when matching Jira issues are created. Leave empty to disable this feature.
	// +kubebuilder:validation:Optional
	AutocreateJql *string `json:"autocreateJql,omitempty" tf:"autocreate_jql,omitempty"`

	// When enabled, automatically creates a Jira issue whenever a PagerDuty incident is triggered.
	// +kubebuilder:validation:Optional
	CreateIssueOnIncidentTrigger *bool `json:"createIssueOnIncidentTrigger,omitempty" tf:"create_issue_on_incident_trigger,omitempty"`

	// Defines how Jira fields are populated when a Jira Issue is created from a PagerDuty Incident.
	// +kubebuilder:validation:Optional
	CustomFields []CustomFieldsParameters `json:"customFields,omitempty" tf:"custom_fields,omitempty"`

	// Specifies the Jira issue type to be created or synchronized with PagerDuty incidents.
	// +kubebuilder:validation:Optional
	IssueType []IssueTypeParameters `json:"issueType" tf:"issue_type,omitempty"`

	// Maps PagerDuty incident priorities to Jira issue priorities for synchronization.
	// +kubebuilder:validation:Optional
	Priorities []PrioritiesParameters `json:"priorities,omitempty" tf:"priorities,omitempty"`

	// [Updating can cause a resource replacement] Defines the Jira project where issues will be created or synchronized.
	// +kubebuilder:validation:Optional
	Project []ProjectParameters `json:"project" tf:"project,omitempty"`

	// Maps PagerDuty incident statuses to corresponding Jira issue statuses for synchronization.
	// +kubebuilder:validation:Optional
	StatusMapping []StatusMappingParameters `json:"statusMapping" tf:"status_mapping,omitempty"`

	// ID of the PagerDuty user for syncing notes and comments between Jira issues and PagerDuty incidents. If not provided, note synchronization is disabled.
	// +kubebuilder:validation:Optional
	SyncNotesUser *string `json:"syncNotesUser,omitempty" tf:"sync_notes_user,omitempty"`
}

type PrioritiesInitParameters struct {

	// The ID of the Jira priority.
	JiraID *string `json:"jiraId,omitempty" tf:"jira_id,omitempty"`

	// The ID of the PagerDuty priority.
	PagerdutyID *string `json:"pagerdutyId,omitempty" tf:"pagerduty_id,omitempty"`
}

type PrioritiesObservation struct {

	// The ID of the Jira priority.
	JiraID *string `json:"jiraId,omitempty" tf:"jira_id,omitempty"`

	// The ID of the PagerDuty priority.
	PagerdutyID *string `json:"pagerdutyId,omitempty" tf:"pagerduty_id,omitempty"`
}

type PrioritiesParameters struct {

	// The ID of the Jira priority.
	// +kubebuilder:validation:Optional
	JiraID *string `json:"jiraId" tf:"jira_id,omitempty"`

	// The ID of the PagerDuty priority.
	// +kubebuilder:validation:Optional
	PagerdutyID *string `json:"pagerdutyId" tf:"pagerduty_id,omitempty"`
}

type ProjectInitParameters struct {

	// Unique identifier for the Jira status.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The short key name of the Jira project.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Name of the Jira status.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ProjectObservation struct {

	// Unique identifier for the Jira status.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The short key name of the Jira project.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Name of the Jira status.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ProjectParameters struct {

	// Unique identifier for the Jira status.
	// +kubebuilder:validation:Optional
	ID *string `json:"id" tf:"id,omitempty"`

	// The short key name of the Jira project.
	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// Name of the Jira status.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

type ResolvedInitParameters struct {

	// Unique identifier for the Jira status.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the Jira status.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ResolvedObservation struct {

	// Unique identifier for the Jira status.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the Jira status.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ResolvedParameters struct {

	// Unique identifier for the Jira status.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the Jira status.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type StatusMappingInitParameters struct {

	// Jira status that maps to the PagerDuty acknowledged status.
	Acknowledged []AcknowledgedInitParameters `json:"acknowledged,omitempty" tf:"acknowledged,omitempty"`

	// Jira status that maps to the PagerDuty resolved status.
	Resolved []ResolvedInitParameters `json:"resolved,omitempty" tf:"resolved,omitempty"`

	// Jira status that maps to the PagerDuty triggered status.
	Triggered []TriggeredInitParameters `json:"triggered,omitempty" tf:"triggered,omitempty"`
}

type StatusMappingObservation struct {

	// Jira status that maps to the PagerDuty acknowledged status.
	Acknowledged []AcknowledgedObservation `json:"acknowledged,omitempty" tf:"acknowledged,omitempty"`

	// Jira status that maps to the PagerDuty resolved status.
	Resolved []ResolvedObservation `json:"resolved,omitempty" tf:"resolved,omitempty"`

	// Jira status that maps to the PagerDuty triggered status.
	Triggered []TriggeredObservation `json:"triggered,omitempty" tf:"triggered,omitempty"`
}

type StatusMappingParameters struct {

	// Jira status that maps to the PagerDuty acknowledged status.
	// +kubebuilder:validation:Optional
	Acknowledged []AcknowledgedParameters `json:"acknowledged,omitempty" tf:"acknowledged,omitempty"`

	// Jira status that maps to the PagerDuty resolved status.
	// +kubebuilder:validation:Optional
	Resolved []ResolvedParameters `json:"resolved,omitempty" tf:"resolved,omitempty"`

	// Jira status that maps to the PagerDuty triggered status.
	// +kubebuilder:validation:Optional
	Triggered []TriggeredParameters `json:"triggered" tf:"triggered,omitempty"`
}

type TriggeredInitParameters struct {

	// Unique identifier for the Jira status.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the Jira status.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type TriggeredObservation struct {

	// Unique identifier for the Jira status.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the Jira status.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type TriggeredParameters struct {

	// Unique identifier for the Jira status.
	// +kubebuilder:validation:Optional
	ID *string `json:"id" tf:"id,omitempty"`

	// Name of the Jira status.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

// CloudAccountMappingRuleSpec defines the desired state of CloudAccountMappingRule
type CloudAccountMappingRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CloudAccountMappingRuleParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider CloudAccountMappingRuleInitParameters `json:"initProvider,omitempty"`
}

// CloudAccountMappingRuleStatus defines the observed state of CloudAccountMappingRule.
type CloudAccountMappingRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CloudAccountMappingRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// CloudAccountMappingRule is the Schema for the CloudAccountMappingRules API. Creates and manages a Jira Cloud account mapping Rule to integrate with PagerDuty.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,pagerduty}
type CloudAccountMappingRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.config) || (has(self.initProvider) && has(self.initProvider.config))",message="spec.forProvider.config is a required parameter"
	Spec   CloudAccountMappingRuleSpec   `json:"spec"`
	Status CloudAccountMappingRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudAccountMappingRuleList contains a list of CloudAccountMappingRules
type CloudAccountMappingRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudAccountMappingRule `json:"items"`
}

// Repository type metadata.
var (
	CloudAccountMappingRule_Kind             = "CloudAccountMappingRule"
	CloudAccountMappingRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CloudAccountMappingRule_Kind}.String()
	CloudAccountMappingRule_KindAPIVersion   = CloudAccountMappingRule_Kind + "." + CRDGroupVersion.String()
	CloudAccountMappingRule_GroupVersionKind = CRDGroupVersion.WithKind(CloudAccountMappingRule_Kind)
)

func init() {
	SchemeBuilder.Register(&CloudAccountMappingRule{}, &CloudAccountMappingRuleList{})
}
