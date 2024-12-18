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

type ActionsActionServiceAssociationInitParameters struct {

	// Id of the action.
	ActionID *string `json:"actionId,omitempty" tf:"action_id,omitempty"`

	// Id of the service associated to the action.
	ServiceID *string `json:"serviceId,omitempty" tf:"service_id,omitempty"`
}

type ActionsActionServiceAssociationObservation struct {

	// Id of the action.
	ActionID *string `json:"actionId,omitempty" tf:"action_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Id of the service associated to the action.
	ServiceID *string `json:"serviceId,omitempty" tf:"service_id,omitempty"`
}

type ActionsActionServiceAssociationParameters struct {

	// Id of the action.
	// +kubebuilder:validation:Optional
	ActionID *string `json:"actionId,omitempty" tf:"action_id,omitempty"`

	// Id of the service associated to the action.
	// +kubebuilder:validation:Optional
	ServiceID *string `json:"serviceId,omitempty" tf:"service_id,omitempty"`
}

// ActionsActionServiceAssociationSpec defines the desired state of ActionsActionServiceAssociation
type ActionsActionServiceAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ActionsActionServiceAssociationParameters `json:"forProvider"`
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
	InitProvider ActionsActionServiceAssociationInitParameters `json:"initProvider,omitempty"`
}

// ActionsActionServiceAssociationStatus defines the observed state of ActionsActionServiceAssociation.
type ActionsActionServiceAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ActionsActionServiceAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ActionsActionServiceAssociation is the Schema for the ActionsActionServiceAssociations API. Creates and manages an Automation Actions action association with a Service in PagerDuty.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,pagerduty}
type ActionsActionServiceAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.actionId) || (has(self.initProvider) && has(self.initProvider.actionId))",message="spec.forProvider.actionId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceId) || (has(self.initProvider) && has(self.initProvider.serviceId))",message="spec.forProvider.serviceId is a required parameter"
	Spec   ActionsActionServiceAssociationSpec   `json:"spec"`
	Status ActionsActionServiceAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ActionsActionServiceAssociationList contains a list of ActionsActionServiceAssociations
type ActionsActionServiceAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ActionsActionServiceAssociation `json:"items"`
}

// Repository type metadata.
var (
	ActionsActionServiceAssociation_Kind             = "ActionsActionServiceAssociation"
	ActionsActionServiceAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ActionsActionServiceAssociation_Kind}.String()
	ActionsActionServiceAssociation_KindAPIVersion   = ActionsActionServiceAssociation_Kind + "." + CRDGroupVersion.String()
	ActionsActionServiceAssociation_GroupVersionKind = CRDGroupVersion.WithKind(ActionsActionServiceAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&ActionsActionServiceAssociation{}, &ActionsActionServiceAssociationList{})
}
