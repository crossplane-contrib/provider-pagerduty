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

type CustomFieldOptionInitParameters struct {

	// The datatype of the field option. Only string is allowed here at present.
	DataType *string `json:"dataType,omitempty" tf:"data_type,omitempty"`

	// The ID of the field.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-pagerduty/apis/incident/v1alpha1.CustomField
	Field *string `json:"field,omitempty" tf:"field,omitempty"`

	// Reference to a CustomField in incident to populate field.
	// +kubebuilder:validation:Optional
	FieldRef *v1.Reference `json:"fieldRef,omitempty" tf:"-"`

	// Selector for a CustomField in incident to populate field.
	// +kubebuilder:validation:Optional
	FieldSelector *v1.Selector `json:"fieldSelector,omitempty" tf:"-"`

	// The allowed value.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type CustomFieldOptionObservation struct {

	// The datatype of the field option. Only string is allowed here at present.
	DataType *string `json:"dataType,omitempty" tf:"data_type,omitempty"`

	// The ID of the field.
	Field *string `json:"field,omitempty" tf:"field,omitempty"`

	// The ID of the field option.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The allowed value.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type CustomFieldOptionParameters struct {

	// The datatype of the field option. Only string is allowed here at present.
	// +kubebuilder:validation:Optional
	DataType *string `json:"dataType,omitempty" tf:"data_type,omitempty"`

	// The ID of the field.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-pagerduty/apis/incident/v1alpha1.CustomField
	// +kubebuilder:validation:Optional
	Field *string `json:"field,omitempty" tf:"field,omitempty"`

	// Reference to a CustomField in incident to populate field.
	// +kubebuilder:validation:Optional
	FieldRef *v1.Reference `json:"fieldRef,omitempty" tf:"-"`

	// Selector for a CustomField in incident to populate field.
	// +kubebuilder:validation:Optional
	FieldSelector *v1.Selector `json:"fieldSelector,omitempty" tf:"-"`

	// The allowed value.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

// CustomFieldOptionSpec defines the desired state of CustomFieldOption
type CustomFieldOptionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CustomFieldOptionParameters `json:"forProvider"`
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
	InitProvider CustomFieldOptionInitParameters `json:"initProvider,omitempty"`
}

// CustomFieldOptionStatus defines the observed state of CustomFieldOption.
type CustomFieldOptionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CustomFieldOptionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// CustomFieldOption is the Schema for the CustomFieldOptions API. Creates and manages an field option for an Incident Custom Field in PagerDuty.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,pagerduty}
type CustomFieldOption struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.dataType) || (has(self.initProvider) && has(self.initProvider.dataType))",message="spec.forProvider.dataType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.value) || (has(self.initProvider) && has(self.initProvider.value))",message="spec.forProvider.value is a required parameter"
	Spec   CustomFieldOptionSpec   `json:"spec"`
	Status CustomFieldOptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CustomFieldOptionList contains a list of CustomFieldOptions
type CustomFieldOptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CustomFieldOption `json:"items"`
}

// Repository type metadata.
var (
	CustomFieldOption_Kind             = "CustomFieldOption"
	CustomFieldOption_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CustomFieldOption_Kind}.String()
	CustomFieldOption_KindAPIVersion   = CustomFieldOption_Kind + "." + CRDGroupVersion.String()
	CustomFieldOption_GroupVersionKind = CRDGroupVersion.WithKind(CustomFieldOption_Kind)
)

func init() {
	SchemeBuilder.Register(&CustomFieldOption{}, &CustomFieldOptionList{})
}
