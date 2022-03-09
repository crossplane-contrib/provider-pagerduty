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

type LayerObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type LayerParameters struct {

	// +kubebuilder:validation:Optional
	End *string `json:"end,omitempty" tf:"end,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Restriction []RestrictionParameters `json:"restriction,omitempty" tf:"restriction,omitempty"`

	// +kubebuilder:validation:Required
	RotationTurnLengthSeconds *float64 `json:"rotationTurnLengthSeconds" tf:"rotation_turn_length_seconds,omitempty"`

	// +kubebuilder:validation:Required
	RotationVirtualStart *string `json:"rotationVirtualStart" tf:"rotation_virtual_start,omitempty"`

	// +kubebuilder:validation:Required
	Start *string `json:"start" tf:"start,omitempty"`

	// +kubebuilder:validation:Required
	Users []*string `json:"users" tf:"users,omitempty"`
}

type RestrictionObservation struct {
}

type RestrictionParameters struct {

	// +kubebuilder:validation:Required
	DurationSeconds *float64 `json:"durationSeconds" tf:"duration_seconds,omitempty"`

	// +kubebuilder:validation:Optional
	StartDayOfWeek *float64 `json:"startDayOfWeek,omitempty" tf:"start_day_of_week,omitempty"`

	// +kubebuilder:validation:Required
	StartTimeOfDay *string `json:"startTimeOfDay" tf:"start_time_of_day,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type ScheduleObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ScheduleParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Layer []LayerParameters `json:"layer" tf:"layer,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Overflow *bool `json:"overflow,omitempty" tf:"overflow,omitempty"`

	// +kubebuilder:validation:Optional
	Teams []*string `json:"teams,omitempty" tf:"teams,omitempty"`

	// +kubebuilder:validation:Required
	TimeZone *string `json:"timeZone" tf:"time_zone,omitempty"`
}

// ScheduleSpec defines the desired state of Schedule
type ScheduleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ScheduleParameters `json:"forProvider"`
}

// ScheduleStatus defines the observed state of Schedule.
type ScheduleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ScheduleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Schedule is the Schema for the Schedules API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,pagerdutyjet}
type Schedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ScheduleSpec   `json:"spec"`
	Status            ScheduleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ScheduleList contains a list of Schedules
type ScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Schedule `json:"items"`
}

// Repository type metadata.
var (
	Schedule_Kind             = "Schedule"
	Schedule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Schedule_Kind}.String()
	Schedule_KindAPIVersion   = Schedule_Kind + "." + CRDGroupVersion.String()
	Schedule_GroupVersionKind = CRDGroupVersion.WithKind(Schedule_Kind)
)

func init() {
	SchemeBuilder.Register(&Schedule{}, &ScheduleList{})
}