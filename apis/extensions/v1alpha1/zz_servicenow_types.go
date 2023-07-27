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

type ServicenowObservation struct {

	// This is the objects for which the extension applies (An array of service ids).
	ExtensionObjects []*string `json:"extensionObjects,omitempty" tf:"extension_objects,omitempty"`

	// This is the schema for this extension.
	ExtensionSchema *string `json:"extensionSchema,omitempty" tf:"extension_schema,omitempty"`

	// URL at which the entity is uniquely displayed in the Web app.
	HTMLURL *string `json:"htmlUrl,omitempty" tf:"html_url,omitempty"`

	// The ID of the extension.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the service extension.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ServiceNow referer.
	Referer *string `json:"referer,omitempty" tf:"referer,omitempty"`

	// The ServiceNow username.
	SnowUser *string `json:"snowUser,omitempty" tf:"snow_user,omitempty"`

	// A short-form, server-generated string that provides succinct, important information about an object suitable for primary labeling of an entity in a client. In many cases, this will be identical to name, though it is not intended to be an identifier.
	Summary *string `json:"summary,omitempty" tf:"summary,omitempty"`

	// The ServiceNow sync option.
	SyncOptions *string `json:"syncOptions,omitempty" tf:"sync_options,omitempty"`

	// Target Webhook URL.
	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	// The ServiceNow task type, typically incident.
	TaskType *string `json:"taskType,omitempty" tf:"task_type,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ServicenowParameters struct {

	// +kubebuilder:validation:Optional
	EndpointURLSecretRef *v1.SecretKeySelector `json:"endpointUrlSecretRef,omitempty" tf:"-"`

	// This is the objects for which the extension applies (An array of service ids).
	// +kubebuilder:validation:Optional
	ExtensionObjects []*string `json:"extensionObjects,omitempty" tf:"extension_objects,omitempty"`

	// This is the schema for this extension.
	// +kubebuilder:validation:Optional
	ExtensionSchema *string `json:"extensionSchema,omitempty" tf:"extension_schema,omitempty"`

	// The name of the service extension.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ServiceNow referer.
	// +kubebuilder:validation:Optional
	Referer *string `json:"referer,omitempty" tf:"referer,omitempty"`

	// The ServiceNow password.
	// +kubebuilder:validation:Optional
	SnowPasswordSecretRef v1.SecretKeySelector `json:"snowPasswordSecretRef" tf:"-"`

	// The ServiceNow username.
	// +kubebuilder:validation:Optional
	SnowUser *string `json:"snowUser,omitempty" tf:"snow_user,omitempty"`

	// A short-form, server-generated string that provides succinct, important information about an object suitable for primary labeling of an entity in a client. In many cases, this will be identical to name, though it is not intended to be an identifier.
	// +kubebuilder:validation:Optional
	Summary *string `json:"summary,omitempty" tf:"summary,omitempty"`

	// The ServiceNow sync option.
	// +kubebuilder:validation:Optional
	SyncOptions *string `json:"syncOptions,omitempty" tf:"sync_options,omitempty"`

	// Target Webhook URL.
	// +kubebuilder:validation:Optional
	Target *string `json:"target,omitempty" tf:"target,omitempty"`

	// The ServiceNow task type, typically incident.
	// +kubebuilder:validation:Optional
	TaskType *string `json:"taskType,omitempty" tf:"task_type,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// ServicenowSpec defines the desired state of Servicenow
type ServicenowSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServicenowParameters `json:"forProvider"`
}

// ServicenowStatus defines the observed state of Servicenow.
type ServicenowStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServicenowObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Servicenow is the Schema for the Servicenows API. Creates and manages a ServiceNow service extension in PagerDuty.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,pagerduty}
type Servicenow struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.extensionObjects)",message="extensionObjects is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.extensionSchema)",message="extensionSchema is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.referer)",message="referer is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.snowPasswordSecretRef)",message="snowPasswordSecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.snowUser)",message="snowUser is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.syncOptions)",message="syncOptions is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.target)",message="target is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.taskType)",message="taskType is a required parameter"
	Spec   ServicenowSpec   `json:"spec"`
	Status ServicenowStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServicenowList contains a list of Servicenows
type ServicenowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Servicenow `json:"items"`
}

// Repository type metadata.
var (
	Servicenow_Kind             = "Servicenow"
	Servicenow_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Servicenow_Kind}.String()
	Servicenow_KindAPIVersion   = Servicenow_Kind + "." + CRDGroupVersion.String()
	Servicenow_GroupVersionKind = CRDGroupVersion.WithKind(Servicenow_Kind)
)

func init() {
	SchemeBuilder.Register(&Servicenow{}, &ServicenowList{})
}
