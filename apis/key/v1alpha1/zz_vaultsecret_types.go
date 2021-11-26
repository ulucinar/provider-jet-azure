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

type VaultSecretObservation struct {
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	VersionlessID *string `json:"versionlessId,omitempty" tf:"versionless_id,omitempty"`
}

type VaultSecretParameters struct {

	// +kubebuilder:validation:Optional
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// +kubebuilder:validation:Optional
	ExpirationDate *string `json:"expirationDate,omitempty" tf:"expiration_date,omitempty"`

	// +kubebuilder:validation:Required
	KeyVaultID *string `json:"keyVaultId" tf:"key_vault_id,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	NotBeforeDate *string `json:"notBeforeDate,omitempty" tf:"not_before_date,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	ValueSecretRef v1.SecretKeySelector `json:"valueSecretRef" tf:"-"`
}

// VaultSecretSpec defines the desired state of VaultSecret
type VaultSecretSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VaultSecretParameters `json:"forProvider"`
}

// VaultSecretStatus defines the observed state of VaultSecret.
type VaultSecretStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VaultSecretObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VaultSecret is the Schema for the VaultSecrets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type VaultSecret struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VaultSecretSpec   `json:"spec"`
	Status            VaultSecretStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VaultSecretList contains a list of VaultSecrets
type VaultSecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VaultSecret `json:"items"`
}

// Repository type metadata.
var (
	VaultSecret_Kind             = "VaultSecret"
	VaultSecret_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VaultSecret_Kind}.String()
	VaultSecret_KindAPIVersion   = VaultSecret_Kind + "." + CRDGroupVersion.String()
	VaultSecret_GroupVersionKind = CRDGroupVersion.WithKind(VaultSecret_Kind)
)

func init() {
	SchemeBuilder.Register(&VaultSecret{}, &VaultSecretList{})
}
