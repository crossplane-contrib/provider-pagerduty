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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DeliveryMethodObservation struct {
}

type DeliveryMethodParameters struct {

	// +kubebuilder:validation:Optional
	TemporarilyDisabled *bool `json:"temporarilyDisabled,omitempty" tf:"temporarily_disabled,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type FilterObservation struct {
}

type FilterParameters struct {

	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type SubscriptionObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SubscriptionParameters struct {

	// +kubebuilder:validation:Optional
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`

	// +kubebuilder:validation:Required
	DeliveryMethod []DeliveryMethodParameters `json:"deliveryMethod" tf:"delivery_method,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Events []*string `json:"events" tf:"events,omitempty"`

	// +kubebuilder:validation:Required
	Filter []FilterParameters `json:"filter" tf:"filter,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// SubscriptionSpec defines the desired state of Subscription
type SubscriptionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubscriptionParameters `json:"forProvider"`
}

// SubscriptionStatus defines the observed state of Subscription.
type SubscriptionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubscriptionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Subscription is the Schema for the Subscriptions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,pagerdutyjet}
type Subscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SubscriptionSpec   `json:"spec"`
	Status            SubscriptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubscriptionList contains a list of Subscriptions
type SubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Subscription `json:"items"`
}

// Repository type metadata.
var (
	Subscription_Kind             = "Subscription"
	Subscription_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Subscription_Kind}.String()
	Subscription_KindAPIVersion   = Subscription_Kind + "." + CRDGroupVersion.String()
	Subscription_GroupVersionKind = CRDGroupVersion.WithKind(Subscription_Kind)
)

func init() {
	SchemeBuilder.Register(&Subscription{}, &SubscriptionList{})
}