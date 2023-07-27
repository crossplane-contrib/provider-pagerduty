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

type TeamObservation struct {

	// A human-friendly description of the team.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// URL at which the entity is uniquely displayed in the Web app
	HTMLURL *string `json:"htmlUrl,omitempty" tf:"html_url,omitempty"`

	// The ID of the team.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of the parent team. This is available to accounts with the Team Hierarchy feature enabled. Please contact your account manager for more information.
	Parent *string `json:"parent,omitempty" tf:"parent,omitempty"`
}

type TeamParameters struct {

	// A human-friendly description of the team.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the group.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of the parent team. This is available to accounts with the Team Hierarchy feature enabled. Please contact your account manager for more information.
	// +kubebuilder:validation:Optional
	Parent *string `json:"parent,omitempty" tf:"parent,omitempty"`
}

// TeamSpec defines the desired state of Team
type TeamSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TeamParameters `json:"forProvider"`
}

// TeamStatus defines the observed state of Team.
type TeamStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TeamObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Team is the Schema for the Teams API. Creates and manages a team in PagerDuty.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,pagerduty}
type Team struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	Spec   TeamSpec   `json:"spec"`
	Status TeamStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TeamList contains a list of Teams
type TeamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Team `json:"items"`
}

// Repository type metadata.
var (
	Team_Kind             = "Team"
	Team_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Team_Kind}.String()
	Team_KindAPIVersion   = Team_Kind + "." + CRDGroupVersion.String()
	Team_GroupVersionKind = CRDGroupVersion.WithKind(Team_Kind)
)

func init() {
	SchemeBuilder.Register(&Team{}, &TeamList{})
}
