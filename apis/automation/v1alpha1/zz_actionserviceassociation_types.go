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

type ActionServiceAssociationInitParameters struct {

	// Id of the action.
	// +crossplane:generate:reference:type=Action
	ActionID *string `json:"actionId,omitempty" tf:"action_id,omitempty"`

	// Reference to a Action to populate actionId.
	// +kubebuilder:validation:Optional
	ActionIDRef *v1.Reference `json:"actionIdRef,omitempty" tf:"-"`

	// Selector for a Action to populate actionId.
	// +kubebuilder:validation:Optional
	ActionIDSelector *v1.Selector `json:"actionIdSelector,omitempty" tf:"-"`

	// Id of the service associated to the action.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-pagerduty/apis/service/v1alpha1.Service
	ServiceID *string `json:"serviceId,omitempty" tf:"service_id,omitempty"`

	// Reference to a Service in service to populate serviceId.
	// +kubebuilder:validation:Optional
	ServiceIDRef *v1.Reference `json:"serviceIdRef,omitempty" tf:"-"`

	// Selector for a Service in service to populate serviceId.
	// +kubebuilder:validation:Optional
	ServiceIDSelector *v1.Selector `json:"serviceIdSelector,omitempty" tf:"-"`
}

type ActionServiceAssociationObservation struct {

	// Id of the action.
	ActionID *string `json:"actionId,omitempty" tf:"action_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Id of the service associated to the action.
	ServiceID *string `json:"serviceId,omitempty" tf:"service_id,omitempty"`
}

type ActionServiceAssociationParameters struct {

	// Id of the action.
	// +crossplane:generate:reference:type=Action
	// +kubebuilder:validation:Optional
	ActionID *string `json:"actionId,omitempty" tf:"action_id,omitempty"`

	// Reference to a Action to populate actionId.
	// +kubebuilder:validation:Optional
	ActionIDRef *v1.Reference `json:"actionIdRef,omitempty" tf:"-"`

	// Selector for a Action to populate actionId.
	// +kubebuilder:validation:Optional
	ActionIDSelector *v1.Selector `json:"actionIdSelector,omitempty" tf:"-"`

	// Id of the service associated to the action.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-pagerduty/apis/service/v1alpha1.Service
	// +kubebuilder:validation:Optional
	ServiceID *string `json:"serviceId,omitempty" tf:"service_id,omitempty"`

	// Reference to a Service in service to populate serviceId.
	// +kubebuilder:validation:Optional
	ServiceIDRef *v1.Reference `json:"serviceIdRef,omitempty" tf:"-"`

	// Selector for a Service in service to populate serviceId.
	// +kubebuilder:validation:Optional
	ServiceIDSelector *v1.Selector `json:"serviceIdSelector,omitempty" tf:"-"`
}

// ActionServiceAssociationSpec defines the desired state of ActionServiceAssociation
type ActionServiceAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ActionServiceAssociationParameters `json:"forProvider"`
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
	InitProvider ActionServiceAssociationInitParameters `json:"initProvider,omitempty"`
}

// ActionServiceAssociationStatus defines the observed state of ActionServiceAssociation.
type ActionServiceAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ActionServiceAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ActionServiceAssociation is the Schema for the ActionServiceAssociations API. Creates and manages an Automation Actions action association with a Service in PagerDuty.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,pagerduty}
type ActionServiceAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ActionServiceAssociationSpec   `json:"spec"`
	Status            ActionServiceAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ActionServiceAssociationList contains a list of ActionServiceAssociations
type ActionServiceAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ActionServiceAssociation `json:"items"`
}

// Repository type metadata.
var (
	ActionServiceAssociation_Kind             = "ActionServiceAssociation"
	ActionServiceAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ActionServiceAssociation_Kind}.String()
	ActionServiceAssociation_KindAPIVersion   = ActionServiceAssociation_Kind + "." + CRDGroupVersion.String()
	ActionServiceAssociation_GroupVersionKind = CRDGroupVersion.WithKind(ActionServiceAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&ActionServiceAssociation{}, &ActionServiceAssociationList{})
}
