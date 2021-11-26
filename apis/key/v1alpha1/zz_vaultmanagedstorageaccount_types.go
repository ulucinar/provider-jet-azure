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

type VaultManagedStorageAccountObservation struct {
}

type VaultManagedStorageAccountParameters struct {

	// +kubebuilder:validation:Required
	KeyVaultID *string `json:"keyVaultId" tf:"key_vault_id,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	RegenerateKeyAutomatically *bool `json:"regenerateKeyAutomatically,omitempty" tf:"regenerate_key_automatically,omitempty"`

	// +kubebuilder:validation:Optional
	RegenerationPeriod *string `json:"regenerationPeriod,omitempty" tf:"regeneration_period,omitempty"`

	// +kubebuilder:validation:Required
	StorageAccountID *string `json:"storageAccountId" tf:"storage_account_id,omitempty"`

	// +kubebuilder:validation:Required
	StorageAccountKey *string `json:"storageAccountKey" tf:"storage_account_key,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// VaultManagedStorageAccountSpec defines the desired state of VaultManagedStorageAccount
type VaultManagedStorageAccountSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VaultManagedStorageAccountParameters `json:"forProvider"`
}

// VaultManagedStorageAccountStatus defines the observed state of VaultManagedStorageAccount.
type VaultManagedStorageAccountStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VaultManagedStorageAccountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VaultManagedStorageAccount is the Schema for the VaultManagedStorageAccounts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type VaultManagedStorageAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VaultManagedStorageAccountSpec   `json:"spec"`
	Status            VaultManagedStorageAccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VaultManagedStorageAccountList contains a list of VaultManagedStorageAccounts
type VaultManagedStorageAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VaultManagedStorageAccount `json:"items"`
}

// Repository type metadata.
var (
	VaultManagedStorageAccount_Kind             = "VaultManagedStorageAccount"
	VaultManagedStorageAccount_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VaultManagedStorageAccount_Kind}.String()
	VaultManagedStorageAccount_KindAPIVersion   = VaultManagedStorageAccount_Kind + "." + CRDGroupVersion.String()
	VaultManagedStorageAccount_GroupVersionKind = CRDGroupVersion.WithKind(VaultManagedStorageAccount_Kind)
)

func init() {
	SchemeBuilder.Register(&VaultManagedStorageAccount{}, &VaultManagedStorageAccountList{})
}
