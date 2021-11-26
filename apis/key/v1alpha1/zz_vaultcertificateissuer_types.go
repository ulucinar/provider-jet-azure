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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AdminObservation struct {
}

type AdminParameters struct {

	// +kubebuilder:validation:Required
	EmailAddress *string `json:"emailAddress" tf:"email_address,omitempty"`

	// +kubebuilder:validation:Optional
	FirstName *string `json:"firstName,omitempty" tf:"first_name,omitempty"`

	// +kubebuilder:validation:Optional
	LastName *string `json:"lastName,omitempty" tf:"last_name,omitempty"`

	// +kubebuilder:validation:Optional
	Phone *string `json:"phone,omitempty" tf:"phone,omitempty"`
}

type VaultCertificateIssuerObservation struct {
}

type VaultCertificateIssuerParameters struct {

	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// +kubebuilder:validation:Optional
	Admin []AdminParameters `json:"admin,omitempty" tf:"admin,omitempty"`

	// +kubebuilder:validation:Required
	KeyVaultID *string `json:"keyVaultId" tf:"key_vault_id,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	OrgID *string `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	ProviderName *string `json:"providerName" tf:"provider_name,omitempty"`
}

// VaultCertificateIssuerSpec defines the desired state of VaultCertificateIssuer
type VaultCertificateIssuerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VaultCertificateIssuerParameters `json:"forProvider"`
}

// VaultCertificateIssuerStatus defines the observed state of VaultCertificateIssuer.
type VaultCertificateIssuerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VaultCertificateIssuerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VaultCertificateIssuer is the Schema for the VaultCertificateIssuers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type VaultCertificateIssuer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VaultCertificateIssuerSpec   `json:"spec"`
	Status            VaultCertificateIssuerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VaultCertificateIssuerList contains a list of VaultCertificateIssuers
type VaultCertificateIssuerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VaultCertificateIssuer `json:"items"`
}

// Repository type metadata.
var (
	VaultCertificateIssuer_Kind             = "VaultCertificateIssuer"
	VaultCertificateIssuer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VaultCertificateIssuer_Kind}.String()
	VaultCertificateIssuer_KindAPIVersion   = VaultCertificateIssuer_Kind + "." + CRDGroupVersion.String()
	VaultCertificateIssuer_GroupVersionKind = CRDGroupVersion.WithKind(VaultCertificateIssuer_Kind)
)

func init() {
	SchemeBuilder.Register(&VaultCertificateIssuer{}, &VaultCertificateIssuerList{})
}
