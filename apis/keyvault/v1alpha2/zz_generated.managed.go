/*
Copyright 2020 The Crossplane Authors.

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
// Code generated by angryjet. DO NOT EDIT.

package v1alpha2

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this AccessPolicy.
func (mg *AccessPolicy) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this AccessPolicy.
func (mg *AccessPolicy) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this AccessPolicy.
func (mg *AccessPolicy) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this AccessPolicy.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *AccessPolicy) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this AccessPolicy.
func (mg *AccessPolicy) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this AccessPolicy.
func (mg *AccessPolicy) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this AccessPolicy.
func (mg *AccessPolicy) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this AccessPolicy.
func (mg *AccessPolicy) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this AccessPolicy.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *AccessPolicy) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this AccessPolicy.
func (mg *AccessPolicy) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this Certificate.
func (mg *Certificate) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this Certificate.
func (mg *Certificate) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this Certificate.
func (mg *Certificate) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this Certificate.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *Certificate) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this Certificate.
func (mg *Certificate) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this Certificate.
func (mg *Certificate) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this Certificate.
func (mg *Certificate) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this Certificate.
func (mg *Certificate) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this Certificate.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *Certificate) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this Certificate.
func (mg *Certificate) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this CertificateIssuer.
func (mg *CertificateIssuer) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this CertificateIssuer.
func (mg *CertificateIssuer) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this CertificateIssuer.
func (mg *CertificateIssuer) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this CertificateIssuer.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *CertificateIssuer) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this CertificateIssuer.
func (mg *CertificateIssuer) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this CertificateIssuer.
func (mg *CertificateIssuer) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this CertificateIssuer.
func (mg *CertificateIssuer) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this CertificateIssuer.
func (mg *CertificateIssuer) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this CertificateIssuer.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *CertificateIssuer) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this CertificateIssuer.
func (mg *CertificateIssuer) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this Key.
func (mg *Key) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this Key.
func (mg *Key) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this Key.
func (mg *Key) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this Key.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *Key) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this Key.
func (mg *Key) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this Key.
func (mg *Key) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this Key.
func (mg *Key) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this Key.
func (mg *Key) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this Key.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *Key) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this Key.
func (mg *Key) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this ManagedHardwareSecurityModule.
func (mg *ManagedHardwareSecurityModule) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this ManagedHardwareSecurityModule.
func (mg *ManagedHardwareSecurityModule) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this ManagedHardwareSecurityModule.
func (mg *ManagedHardwareSecurityModule) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this ManagedHardwareSecurityModule.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *ManagedHardwareSecurityModule) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this ManagedHardwareSecurityModule.
func (mg *ManagedHardwareSecurityModule) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this ManagedHardwareSecurityModule.
func (mg *ManagedHardwareSecurityModule) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this ManagedHardwareSecurityModule.
func (mg *ManagedHardwareSecurityModule) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this ManagedHardwareSecurityModule.
func (mg *ManagedHardwareSecurityModule) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this ManagedHardwareSecurityModule.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *ManagedHardwareSecurityModule) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this ManagedHardwareSecurityModule.
func (mg *ManagedHardwareSecurityModule) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this Secret.
func (mg *Secret) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this Secret.
func (mg *Secret) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this Secret.
func (mg *Secret) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this Secret.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *Secret) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this Secret.
func (mg *Secret) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this Secret.
func (mg *Secret) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this Secret.
func (mg *Secret) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this Secret.
func (mg *Secret) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this Secret.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *Secret) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this Secret.
func (mg *Secret) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this Vault.
func (mg *Vault) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this Vault.
func (mg *Vault) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this Vault.
func (mg *Vault) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this Vault.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *Vault) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this Vault.
func (mg *Vault) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this Vault.
func (mg *Vault) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this Vault.
func (mg *Vault) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this Vault.
func (mg *Vault) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this Vault.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *Vault) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this Vault.
func (mg *Vault) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
