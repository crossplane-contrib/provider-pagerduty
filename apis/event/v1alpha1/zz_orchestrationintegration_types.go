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

type OrchestrationIntegrationInitParameters struct {

	// ID of the Event Orchestration to which this Integration belongs to. If value is changed, current Integration is associated with a newly provided ID.
	// +crossplane:generate:reference:type=Orchestration
	EventOrchestration *string `json:"eventOrchestration,omitempty" tf:"event_orchestration,omitempty"`

	// Reference to a Orchestration to populate eventOrchestration.
	// +kubebuilder:validation:Optional
	EventOrchestrationRef *v1.Reference `json:"eventOrchestrationRef,omitempty" tf:"-"`

	// Selector for a Orchestration to populate eventOrchestration.
	// +kubebuilder:validation:Optional
	EventOrchestrationSelector *v1.Selector `json:"eventOrchestrationSelector,omitempty" tf:"-"`

	// Name/description of the Integration.
	Label *string `json:"label,omitempty" tf:"label,omitempty"`
}

type OrchestrationIntegrationObservation struct {

	// ID of the Event Orchestration to which this Integration belongs to. If value is changed, current Integration is associated with a newly provided ID.
	EventOrchestration *string `json:"eventOrchestration,omitempty" tf:"event_orchestration,omitempty"`

	// ID of this Integration.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name/description of the Integration.
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	Parameters []OrchestrationIntegrationParametersObservation `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type OrchestrationIntegrationParameters struct {

	// ID of the Event Orchestration to which this Integration belongs to. If value is changed, current Integration is associated with a newly provided ID.
	// +crossplane:generate:reference:type=Orchestration
	// +kubebuilder:validation:Optional
	EventOrchestration *string `json:"eventOrchestration,omitempty" tf:"event_orchestration,omitempty"`

	// Reference to a Orchestration to populate eventOrchestration.
	// +kubebuilder:validation:Optional
	EventOrchestrationRef *v1.Reference `json:"eventOrchestrationRef,omitempty" tf:"-"`

	// Selector for a Orchestration to populate eventOrchestration.
	// +kubebuilder:validation:Optional
	EventOrchestrationSelector *v1.Selector `json:"eventOrchestrationSelector,omitempty" tf:"-"`

	// Name/description of the Integration.
	// +kubebuilder:validation:Optional
	Label *string `json:"label,omitempty" tf:"label,omitempty"`
}

type OrchestrationIntegrationParametersInitParameters struct {
}

type OrchestrationIntegrationParametersObservation struct {

	// Routing key that routes to this Orchestration.
	RoutingKey *string `json:"routingKey,omitempty" tf:"routing_key,omitempty"`

	// Type of the routing key. global is the default type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type OrchestrationIntegrationParametersParameters struct {
}

// OrchestrationIntegrationSpec defines the desired state of OrchestrationIntegration
type OrchestrationIntegrationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OrchestrationIntegrationParameters `json:"forProvider"`
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
	InitProvider OrchestrationIntegrationInitParameters `json:"initProvider,omitempty"`
}

// OrchestrationIntegrationStatus defines the observed state of OrchestrationIntegration.
type OrchestrationIntegrationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OrchestrationIntegrationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// OrchestrationIntegration is the Schema for the OrchestrationIntegrations API. Creates and manages an Integration for an Event Orchestration.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,pagerduty}
type OrchestrationIntegration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.label) || (has(self.initProvider) && has(self.initProvider.label))",message="spec.forProvider.label is a required parameter"
	Spec   OrchestrationIntegrationSpec   `json:"spec"`
	Status OrchestrationIntegrationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OrchestrationIntegrationList contains a list of OrchestrationIntegrations
type OrchestrationIntegrationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OrchestrationIntegration `json:"items"`
}

// Repository type metadata.
var (
	OrchestrationIntegration_Kind             = "OrchestrationIntegration"
	OrchestrationIntegration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OrchestrationIntegration_Kind}.String()
	OrchestrationIntegration_KindAPIVersion   = OrchestrationIntegration_Kind + "." + CRDGroupVersion.String()
	OrchestrationIntegration_GroupVersionKind = CRDGroupVersion.WithKind(OrchestrationIntegration_Kind)
)

func init() {
	SchemeBuilder.Register(&OrchestrationIntegration{}, &OrchestrationIntegrationList{})
}
