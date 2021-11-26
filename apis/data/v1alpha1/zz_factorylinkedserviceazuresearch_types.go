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

type FactoryLinkedServiceAzureSearchObservation struct {
	EncryptedCredential *string `json:"encryptedCredential,omitempty" tf:"encrypted_credential,omitempty"`
}

type FactoryLinkedServiceAzureSearchParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// +kubebuilder:validation:Required
	DataFactoryID *string `json:"dataFactoryId" tf:"data_factory_id,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// +kubebuilder:validation:Required
	SearchServiceKey *string `json:"searchServiceKey" tf:"search_service_key,omitempty"`

	// +kubebuilder:validation:Required
	URL *string `json:"url" tf:"url,omitempty"`
}

// FactoryLinkedServiceAzureSearchSpec defines the desired state of FactoryLinkedServiceAzureSearch
type FactoryLinkedServiceAzureSearchSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FactoryLinkedServiceAzureSearchParameters `json:"forProvider"`
}

// FactoryLinkedServiceAzureSearchStatus defines the observed state of FactoryLinkedServiceAzureSearch.
type FactoryLinkedServiceAzureSearchStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FactoryLinkedServiceAzureSearchObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FactoryLinkedServiceAzureSearch is the Schema for the FactoryLinkedServiceAzureSearchs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type FactoryLinkedServiceAzureSearch struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FactoryLinkedServiceAzureSearchSpec   `json:"spec"`
	Status            FactoryLinkedServiceAzureSearchStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FactoryLinkedServiceAzureSearchList contains a list of FactoryLinkedServiceAzureSearchs
type FactoryLinkedServiceAzureSearchList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FactoryLinkedServiceAzureSearch `json:"items"`
}

// Repository type metadata.
var (
	FactoryLinkedServiceAzureSearch_Kind             = "FactoryLinkedServiceAzureSearch"
	FactoryLinkedServiceAzureSearch_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FactoryLinkedServiceAzureSearch_Kind}.String()
	FactoryLinkedServiceAzureSearch_KindAPIVersion   = FactoryLinkedServiceAzureSearch_Kind + "." + CRDGroupVersion.String()
	FactoryLinkedServiceAzureSearch_GroupVersionKind = CRDGroupVersion.WithKind(FactoryLinkedServiceAzureSearch_Kind)
)

func init() {
	SchemeBuilder.Register(&FactoryLinkedServiceAzureSearch{}, &FactoryLinkedServiceAzureSearchList{})
}
