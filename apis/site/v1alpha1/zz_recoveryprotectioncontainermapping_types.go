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

type RecoveryProtectionContainerMappingObservation struct {
}

type RecoveryProtectionContainerMappingParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	RecoveryFabricName *string `json:"recoveryFabricName" tf:"recovery_fabric_name,omitempty"`

	// +kubebuilder:validation:Required
	RecoveryReplicationPolicyID *string `json:"recoveryReplicationPolicyId" tf:"recovery_replication_policy_id,omitempty"`

	// +kubebuilder:validation:Required
	RecoverySourceProtectionContainerName *string `json:"recoverySourceProtectionContainerName" tf:"recovery_source_protection_container_name,omitempty"`

	// +kubebuilder:validation:Required
	RecoveryTargetProtectionContainerID *string `json:"recoveryTargetProtectionContainerId" tf:"recovery_target_protection_container_id,omitempty"`

	// +kubebuilder:validation:Required
	RecoveryVaultName *string `json:"recoveryVaultName" tf:"recovery_vault_name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`
}

// RecoveryProtectionContainerMappingSpec defines the desired state of RecoveryProtectionContainerMapping
type RecoveryProtectionContainerMappingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RecoveryProtectionContainerMappingParameters `json:"forProvider"`
}

// RecoveryProtectionContainerMappingStatus defines the observed state of RecoveryProtectionContainerMapping.
type RecoveryProtectionContainerMappingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RecoveryProtectionContainerMappingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RecoveryProtectionContainerMapping is the Schema for the RecoveryProtectionContainerMappings API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type RecoveryProtectionContainerMapping struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RecoveryProtectionContainerMappingSpec   `json:"spec"`
	Status            RecoveryProtectionContainerMappingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RecoveryProtectionContainerMappingList contains a list of RecoveryProtectionContainerMappings
type RecoveryProtectionContainerMappingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RecoveryProtectionContainerMapping `json:"items"`
}

// Repository type metadata.
var (
	RecoveryProtectionContainerMapping_Kind             = "RecoveryProtectionContainerMapping"
	RecoveryProtectionContainerMapping_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RecoveryProtectionContainerMapping_Kind}.String()
	RecoveryProtectionContainerMapping_KindAPIVersion   = RecoveryProtectionContainerMapping_Kind + "." + CRDGroupVersion.String()
	RecoveryProtectionContainerMapping_GroupVersionKind = CRDGroupVersion.WithKind(RecoveryProtectionContainerMapping_Kind)
)

func init() {
	SchemeBuilder.Register(&RecoveryProtectionContainerMapping{}, &RecoveryProtectionContainerMappingList{})
}
