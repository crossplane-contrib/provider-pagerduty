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

type TagObservation struct {

	// URL at which the entity is uniquely displayed in the Web app.
	HTMLURL *string `json:"htmlUrl,omitempty" tf:"html_url,omitempty"`

	// The ID of the tag.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A short-form, server-generated string that provides succinct, important information about an object suitable for primary labeling of an entity in a client. In many cases, this will be identical to name, though it is not intended to be an identifier.
	Summary *string `json:"summary,omitempty" tf:"summary,omitempty"`
}

type TagParameters struct {

	// The label of the tag.
	// +kubebuilder:validation:Required
	Label *string `json:"label" tf:"label,omitempty"`
}

// TagSpec defines the desired state of Tag
type TagSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TagParameters `json:"forProvider"`
}

// TagStatus defines the observed state of Tag.
type TagStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TagObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Tag is the Schema for the Tags API. Creates and manages a tag in PagerDuty.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,pagerduty}
type Tag struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TagSpec   `json:"spec"`
	Status            TagStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TagList contains a list of Tags
type TagList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Tag `json:"items"`
}

// Repository type metadata.
var (
	Tag_Kind             = "Tag"
	Tag_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Tag_Kind}.String()
	Tag_KindAPIVersion   = Tag_Kind + "." + CRDGroupVersion.String()
	Tag_GroupVersionKind = CRDGroupVersion.WithKind(Tag_Kind)
)

func init() {
	SchemeBuilder.Register(&Tag{}, &TagList{})
}