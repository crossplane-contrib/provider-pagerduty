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

type ContactMethodInitParameters struct {

	// The "address" to deliver to: email, phone number, etc., depending on the type.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The 1-to-3 digit country calling code. Required when using phone_contact_method or sms_contact_method.
	CountryCode *float64 `json:"countryCode,omitempty" tf:"country_code,omitempty"`

	// The label (e.g., "Work", "Mobile", etc.).
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// Send an abbreviated email message instead of the standard email output.
	SendShortEmail *bool `json:"sendShortEmail,omitempty" tf:"send_short_email,omitempty"`

	// The contact method type. May be (email_contact_method, phone_contact_method, sms_contact_method, push_notification_contact_method).
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The ID of the user.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-pagerduty/apis/user/v1alpha1.User
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`

	// Reference to a User in user to populate userId.
	// +kubebuilder:validation:Optional
	UserIDRef *v1.Reference `json:"userIdRef,omitempty" tf:"-"`

	// Selector for a User in user to populate userId.
	// +kubebuilder:validation:Optional
	UserIDSelector *v1.Selector `json:"userIdSelector,omitempty" tf:"-"`
}

type ContactMethodObservation struct {

	// The "address" to deliver to: email, phone number, etc., depending on the type.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// If true, this phone has been blacklisted by PagerDuty and no messages will be sent to it.
	Blacklisted *bool `json:"blacklisted,omitempty" tf:"blacklisted,omitempty"`

	// The 1-to-3 digit country calling code. Required when using phone_contact_method or sms_contact_method.
	CountryCode *float64 `json:"countryCode,omitempty" tf:"country_code,omitempty"`

	// If true, this phone is capable of receiving SMS messages.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The ID of the contact method.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The label (e.g., "Work", "Mobile", etc.).
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// Send an abbreviated email message instead of the standard email output.
	SendShortEmail *bool `json:"sendShortEmail,omitempty" tf:"send_short_email,omitempty"`

	// The contact method type. May be (email_contact_method, phone_contact_method, sms_contact_method, push_notification_contact_method).
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The ID of the user.
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

type ContactMethodParameters struct {

	// The "address" to deliver to: email, phone number, etc., depending on the type.
	// +kubebuilder:validation:Optional
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The 1-to-3 digit country calling code. Required when using phone_contact_method or sms_contact_method.
	// +kubebuilder:validation:Optional
	CountryCode *float64 `json:"countryCode,omitempty" tf:"country_code,omitempty"`

	// The label (e.g., "Work", "Mobile", etc.).
	// +kubebuilder:validation:Optional
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// Send an abbreviated email message instead of the standard email output.
	// +kubebuilder:validation:Optional
	SendShortEmail *bool `json:"sendShortEmail,omitempty" tf:"send_short_email,omitempty"`

	// The contact method type. May be (email_contact_method, phone_contact_method, sms_contact_method, push_notification_contact_method).
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The ID of the user.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-pagerduty/apis/user/v1alpha1.User
	// +kubebuilder:validation:Optional
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`

	// Reference to a User in user to populate userId.
	// +kubebuilder:validation:Optional
	UserIDRef *v1.Reference `json:"userIdRef,omitempty" tf:"-"`

	// Selector for a User in user to populate userId.
	// +kubebuilder:validation:Optional
	UserIDSelector *v1.Selector `json:"userIdSelector,omitempty" tf:"-"`
}

// ContactMethodSpec defines the desired state of ContactMethod
type ContactMethodSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ContactMethodParameters `json:"forProvider"`
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
	InitProvider ContactMethodInitParameters `json:"initProvider,omitempty"`
}

// ContactMethodStatus defines the observed state of ContactMethod.
type ContactMethodStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ContactMethodObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ContactMethod is the Schema for the ContactMethods API. Creates and manages contact methods for a user in PagerDuty.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,pagerduty}
type ContactMethod struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.address) || (has(self.initProvider) && has(self.initProvider.address))",message="spec.forProvider.address is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.label) || (has(self.initProvider) && has(self.initProvider.label))",message="spec.forProvider.label is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   ContactMethodSpec   `json:"spec"`
	Status ContactMethodStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ContactMethodList contains a list of ContactMethods
type ContactMethodList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContactMethod `json:"items"`
}

// Repository type metadata.
var (
	ContactMethod_Kind             = "ContactMethod"
	ContactMethod_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ContactMethod_Kind}.String()
	ContactMethod_KindAPIVersion   = ContactMethod_Kind + "." + CRDGroupVersion.String()
	ContactMethod_GroupVersionKind = CRDGroupVersion.WithKind(ContactMethod_Kind)
)

func init() {
	SchemeBuilder.Register(&ContactMethod{}, &ContactMethodList{})
}
