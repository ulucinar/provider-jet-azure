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

type EncryptionScopeObservation struct {
}

type EncryptionScopeParameters struct {

	// +kubebuilder:validation:Optional
	InfrastructureEncryptionRequired *bool `json:"infrastructureEncryptionRequired,omitempty" tf:"infrastructure_encryption_required,omitempty"`

	// +kubebuilder:validation:Optional
	KeyVaultKeyID *string `json:"keyVaultKeyId,omitempty" tf:"key_vault_key_id,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Source *string `json:"source" tf:"source,omitempty"`

	// +kubebuilder:validation:Required
	StorageAccountID *string `json:"storageAccountId" tf:"storage_account_id,omitempty"`
}

// EncryptionScopeSpec defines the desired state of EncryptionScope
type EncryptionScopeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EncryptionScopeParameters `json:"forProvider"`
}

// EncryptionScopeStatus defines the observed state of EncryptionScope.
type EncryptionScopeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EncryptionScopeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EncryptionScope is the Schema for the EncryptionScopes API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type EncryptionScope struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EncryptionScopeSpec   `json:"spec"`
	Status            EncryptionScopeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EncryptionScopeList contains a list of EncryptionScopes
type EncryptionScopeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EncryptionScope `json:"items"`
}

// Repository type metadata.
var (
	EncryptionScope_Kind             = "EncryptionScope"
	EncryptionScope_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EncryptionScope_Kind}.String()
	EncryptionScope_KindAPIVersion   = EncryptionScope_Kind + "." + CRDGroupVersion.String()
	EncryptionScope_GroupVersionKind = CRDGroupVersion.WithKind(EncryptionScope_Kind)
)

func init() {
	SchemeBuilder.Register(&EncryptionScope{}, &EncryptionScopeList{})
}
