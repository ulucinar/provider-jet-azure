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

type AppIntegrationAccountSchemaObservation struct {
}

type AppIntegrationAccountSchemaParameters struct {

	// +kubebuilder:validation:Required
	Content *string `json:"content" tf:"content,omitempty"`

	// +kubebuilder:validation:Optional
	FileName *string `json:"fileName,omitempty" tf:"file_name,omitempty"`

	// +kubebuilder:validation:Required
	IntegrationAccountName *string `json:"integrationAccountName" tf:"integration_account_name,omitempty"`

	// +kubebuilder:validation:Optional
	Metadata *string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`
}

// AppIntegrationAccountSchemaSpec defines the desired state of AppIntegrationAccountSchema
type AppIntegrationAccountSchemaSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AppIntegrationAccountSchemaParameters `json:"forProvider"`
}

// AppIntegrationAccountSchemaStatus defines the observed state of AppIntegrationAccountSchema.
type AppIntegrationAccountSchemaStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AppIntegrationAccountSchemaObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AppIntegrationAccountSchema is the Schema for the AppIntegrationAccountSchemas API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type AppIntegrationAccountSchema struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppIntegrationAccountSchemaSpec   `json:"spec"`
	Status            AppIntegrationAccountSchemaStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppIntegrationAccountSchemaList contains a list of AppIntegrationAccountSchemas
type AppIntegrationAccountSchemaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppIntegrationAccountSchema `json:"items"`
}

// Repository type metadata.
var (
	AppIntegrationAccountSchema_Kind             = "AppIntegrationAccountSchema"
	AppIntegrationAccountSchema_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AppIntegrationAccountSchema_Kind}.String()
	AppIntegrationAccountSchema_KindAPIVersion   = AppIntegrationAccountSchema_Kind + "." + CRDGroupVersion.String()
	AppIntegrationAccountSchema_GroupVersionKind = CRDGroupVersion.WithKind(AppIntegrationAccountSchema_Kind)
)

func init() {
	SchemeBuilder.Register(&AppIntegrationAccountSchema{}, &AppIntegrationAccountSchemaList{})
}
