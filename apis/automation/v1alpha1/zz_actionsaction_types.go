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

type ActionDataReferenceInitParameters struct {

	// The command to execute the script with.
	InvocationCommand *string `json:"invocationCommand,omitempty" tf:"invocation_command,omitempty"`

	// The arguments to pass to the Process Automation job execution.
	ProcessAutomationJobArguments *string `json:"processAutomationJobArguments,omitempty" tf:"process_automation_job_arguments,omitempty"`

	// The ID of the Process Automation job to execute.
	ProcessAutomationJobID *string `json:"processAutomationJobId,omitempty" tf:"process_automation_job_id,omitempty"`

	// The expression that filters on which nodes a Process Automation Job executes Learn more.
	ProcessAutomationNodeFilter *string `json:"processAutomationNodeFilter,omitempty" tf:"process_automation_node_filter,omitempty"`

	// Body of the script to be executed on the Runner. Max length is 16777215 characters.
	Script *string `json:"script,omitempty" tf:"script,omitempty"`
}

type ActionDataReferenceObservation struct {

	// The command to execute the script with.
	InvocationCommand *string `json:"invocationCommand,omitempty" tf:"invocation_command,omitempty"`

	// The arguments to pass to the Process Automation job execution.
	ProcessAutomationJobArguments *string `json:"processAutomationJobArguments,omitempty" tf:"process_automation_job_arguments,omitempty"`

	// The ID of the Process Automation job to execute.
	ProcessAutomationJobID *string `json:"processAutomationJobId,omitempty" tf:"process_automation_job_id,omitempty"`

	// The expression that filters on which nodes a Process Automation Job executes Learn more.
	ProcessAutomationNodeFilter *string `json:"processAutomationNodeFilter,omitempty" tf:"process_automation_node_filter,omitempty"`

	// Body of the script to be executed on the Runner. Max length is 16777215 characters.
	Script *string `json:"script,omitempty" tf:"script,omitempty"`
}

type ActionDataReferenceParameters struct {

	// The command to execute the script with.
	// +kubebuilder:validation:Optional
	InvocationCommand *string `json:"invocationCommand,omitempty" tf:"invocation_command,omitempty"`

	// The arguments to pass to the Process Automation job execution.
	// +kubebuilder:validation:Optional
	ProcessAutomationJobArguments *string `json:"processAutomationJobArguments,omitempty" tf:"process_automation_job_arguments,omitempty"`

	// The ID of the Process Automation job to execute.
	// +kubebuilder:validation:Optional
	ProcessAutomationJobID *string `json:"processAutomationJobId,omitempty" tf:"process_automation_job_id,omitempty"`

	// The expression that filters on which nodes a Process Automation Job executes Learn more.
	// +kubebuilder:validation:Optional
	ProcessAutomationNodeFilter *string `json:"processAutomationNodeFilter,omitempty" tf:"process_automation_node_filter,omitempty"`

	// Body of the script to be executed on the Runner. Max length is 16777215 characters.
	// +kubebuilder:validation:Optional
	Script *string `json:"script,omitempty" tf:"script,omitempty"`
}

type ActionsActionInitParameters struct {

	// The category of the action. The only allowed values are diagnostic and remediation.
	ActionClassification *string `json:"actionClassification,omitempty" tf:"action_classification,omitempty"`

	// Action Data block. Action Data is documented below.
	ActionDataReference []ActionDataReferenceInitParameters `json:"actionDataReference,omitempty" tf:"action_data_reference,omitempty"`

	// The type of the action. The only allowed values are process_automation and script. Cannot be changed once set.
	ActionType *string `json:"actionType,omitempty" tf:"action_type,omitempty"`

	// The time action was created. Represented as an ISO 8601 timestamp.
	CreationTime *string `json:"creationTime,omitempty" tf:"creation_time,omitempty"`

	// The description of the action. Max length is 1024 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The last time action has been modified. Represented as an ISO 8601 timestamp.
	ModifyTime *string `json:"modifyTime,omitempty" tf:"modify_time,omitempty"`

	// The name of the action. Max length is 255 characters.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The Process Automation Actions runner to associate the action with. Cannot be changed for the process_automation action type once set.
	RunnerID *string `json:"runnerId,omitempty" tf:"runner_id,omitempty"`

	// The type of the runner associated with the action.
	RunnerType *string `json:"runnerType,omitempty" tf:"runner_type,omitempty"`

	// The type of object. The value returned will be action.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ActionsActionObservation struct {

	// The category of the action. The only allowed values are diagnostic and remediation.
	ActionClassification *string `json:"actionClassification,omitempty" tf:"action_classification,omitempty"`

	// Action Data block. Action Data is documented below.
	ActionDataReference []ActionDataReferenceObservation `json:"actionDataReference,omitempty" tf:"action_data_reference,omitempty"`

	// The type of the action. The only allowed values are process_automation and script. Cannot be changed once set.
	ActionType *string `json:"actionType,omitempty" tf:"action_type,omitempty"`

	// The time action was created. Represented as an ISO 8601 timestamp.
	CreationTime *string `json:"creationTime,omitempty" tf:"creation_time,omitempty"`

	// The description of the action. Max length is 1024 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the action.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The last time action has been modified. Represented as an ISO 8601 timestamp.
	ModifyTime *string `json:"modifyTime,omitempty" tf:"modify_time,omitempty"`

	// The name of the action. Max length is 255 characters.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The Process Automation Actions runner to associate the action with. Cannot be changed for the process_automation action type once set.
	RunnerID *string `json:"runnerId,omitempty" tf:"runner_id,omitempty"`

	// The type of the runner associated with the action.
	RunnerType *string `json:"runnerType,omitempty" tf:"runner_type,omitempty"`

	// The type of object. The value returned will be action.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ActionsActionParameters struct {

	// The category of the action. The only allowed values are diagnostic and remediation.
	// +kubebuilder:validation:Optional
	ActionClassification *string `json:"actionClassification,omitempty" tf:"action_classification,omitempty"`

	// Action Data block. Action Data is documented below.
	// +kubebuilder:validation:Optional
	ActionDataReference []ActionDataReferenceParameters `json:"actionDataReference,omitempty" tf:"action_data_reference,omitempty"`

	// The type of the action. The only allowed values are process_automation and script. Cannot be changed once set.
	// +kubebuilder:validation:Optional
	ActionType *string `json:"actionType,omitempty" tf:"action_type,omitempty"`

	// The time action was created. Represented as an ISO 8601 timestamp.
	// +kubebuilder:validation:Optional
	CreationTime *string `json:"creationTime,omitempty" tf:"creation_time,omitempty"`

	// The description of the action. Max length is 1024 characters.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The last time action has been modified. Represented as an ISO 8601 timestamp.
	// +kubebuilder:validation:Optional
	ModifyTime *string `json:"modifyTime,omitempty" tf:"modify_time,omitempty"`

	// The name of the action. Max length is 255 characters.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The Process Automation Actions runner to associate the action with. Cannot be changed for the process_automation action type once set.
	// +kubebuilder:validation:Optional
	RunnerID *string `json:"runnerId,omitempty" tf:"runner_id,omitempty"`

	// The type of the runner associated with the action.
	// +kubebuilder:validation:Optional
	RunnerType *string `json:"runnerType,omitempty" tf:"runner_type,omitempty"`

	// The type of object. The value returned will be action.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// ActionsActionSpec defines the desired state of ActionsAction
type ActionsActionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ActionsActionParameters `json:"forProvider"`
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
	InitProvider ActionsActionInitParameters `json:"initProvider,omitempty"`
}

// ActionsActionStatus defines the observed state of ActionsAction.
type ActionsActionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ActionsActionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ActionsAction is the Schema for the ActionsActions API. Creates and manages an Automation Actions action in PagerDuty.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,pagerduty}
type ActionsAction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.actionDataReference) || (has(self.initProvider) && has(self.initProvider.actionDataReference))",message="spec.forProvider.actionDataReference is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.actionType) || (has(self.initProvider) && has(self.initProvider.actionType))",message="spec.forProvider.actionType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   ActionsActionSpec   `json:"spec"`
	Status ActionsActionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ActionsActionList contains a list of ActionsActions
type ActionsActionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ActionsAction `json:"items"`
}

// Repository type metadata.
var (
	ActionsAction_Kind             = "ActionsAction"
	ActionsAction_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ActionsAction_Kind}.String()
	ActionsAction_KindAPIVersion   = ActionsAction_Kind + "." + CRDGroupVersion.String()
	ActionsAction_GroupVersionKind = CRDGroupVersion.WithKind(ActionsAction_Kind)
)

func init() {
	SchemeBuilder.Register(&ActionsAction{}, &ActionsActionList{})
}