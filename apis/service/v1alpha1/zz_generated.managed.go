/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this Dependency.
func (mg *Dependency) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this Dependency.
func (mg *Dependency) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicy of this Dependency.
func (mg *Dependency) GetManagementPolicy() xpv1.ManagementPolicy {
	return mg.Spec.ManagementPolicy
}

// GetProviderConfigReference of this Dependency.
func (mg *Dependency) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this Dependency.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *Dependency) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this Dependency.
func (mg *Dependency) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this Dependency.
func (mg *Dependency) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this Dependency.
func (mg *Dependency) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this Dependency.
func (mg *Dependency) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicy of this Dependency.
func (mg *Dependency) SetManagementPolicy(r xpv1.ManagementPolicy) {
	mg.Spec.ManagementPolicy = r
}

// SetProviderConfigReference of this Dependency.
func (mg *Dependency) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this Dependency.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *Dependency) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this Dependency.
func (mg *Dependency) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this Dependency.
func (mg *Dependency) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this EventRule.
func (mg *EventRule) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this EventRule.
func (mg *EventRule) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicy of this EventRule.
func (mg *EventRule) GetManagementPolicy() xpv1.ManagementPolicy {
	return mg.Spec.ManagementPolicy
}

// GetProviderConfigReference of this EventRule.
func (mg *EventRule) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this EventRule.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *EventRule) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this EventRule.
func (mg *EventRule) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this EventRule.
func (mg *EventRule) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this EventRule.
func (mg *EventRule) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this EventRule.
func (mg *EventRule) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicy of this EventRule.
func (mg *EventRule) SetManagementPolicy(r xpv1.ManagementPolicy) {
	mg.Spec.ManagementPolicy = r
}

// SetProviderConfigReference of this EventRule.
func (mg *EventRule) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this EventRule.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *EventRule) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this EventRule.
func (mg *EventRule) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this EventRule.
func (mg *EventRule) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this Integration.
func (mg *Integration) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this Integration.
func (mg *Integration) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicy of this Integration.
func (mg *Integration) GetManagementPolicy() xpv1.ManagementPolicy {
	return mg.Spec.ManagementPolicy
}

// GetProviderConfigReference of this Integration.
func (mg *Integration) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this Integration.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *Integration) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this Integration.
func (mg *Integration) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this Integration.
func (mg *Integration) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this Integration.
func (mg *Integration) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this Integration.
func (mg *Integration) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicy of this Integration.
func (mg *Integration) SetManagementPolicy(r xpv1.ManagementPolicy) {
	mg.Spec.ManagementPolicy = r
}

// SetProviderConfigReference of this Integration.
func (mg *Integration) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this Integration.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *Integration) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this Integration.
func (mg *Integration) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this Integration.
func (mg *Integration) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this Service.
func (mg *Service) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this Service.
func (mg *Service) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicy of this Service.
func (mg *Service) GetManagementPolicy() xpv1.ManagementPolicy {
	return mg.Spec.ManagementPolicy
}

// GetProviderConfigReference of this Service.
func (mg *Service) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this Service.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *Service) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this Service.
func (mg *Service) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this Service.
func (mg *Service) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this Service.
func (mg *Service) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this Service.
func (mg *Service) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicy of this Service.
func (mg *Service) SetManagementPolicy(r xpv1.ManagementPolicy) {
	mg.Spec.ManagementPolicy = r
}

// SetProviderConfigReference of this Service.
func (mg *Service) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this Service.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *Service) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this Service.
func (mg *Service) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this Service.
func (mg *Service) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
