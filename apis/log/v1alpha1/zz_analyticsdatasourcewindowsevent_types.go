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

type AnalyticsDatasourceWindowsEventObservation struct {
}

type AnalyticsDatasourceWindowsEventParameters struct {

	// +kubebuilder:validation:Required
	EventLogName *string `json:"eventLogName" tf:"event_log_name,omitempty"`

	// +kubebuilder:validation:Required
	EventTypes []*string `json:"eventTypes" tf:"event_types,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Required
	WorkspaceName *string `json:"workspaceName" tf:"workspace_name,omitempty"`
}

// AnalyticsDatasourceWindowsEventSpec defines the desired state of AnalyticsDatasourceWindowsEvent
type AnalyticsDatasourceWindowsEventSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AnalyticsDatasourceWindowsEventParameters `json:"forProvider"`
}

// AnalyticsDatasourceWindowsEventStatus defines the observed state of AnalyticsDatasourceWindowsEvent.
type AnalyticsDatasourceWindowsEventStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AnalyticsDatasourceWindowsEventObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AnalyticsDatasourceWindowsEvent is the Schema for the AnalyticsDatasourceWindowsEvents API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type AnalyticsDatasourceWindowsEvent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AnalyticsDatasourceWindowsEventSpec   `json:"spec"`
	Status            AnalyticsDatasourceWindowsEventStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AnalyticsDatasourceWindowsEventList contains a list of AnalyticsDatasourceWindowsEvents
type AnalyticsDatasourceWindowsEventList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AnalyticsDatasourceWindowsEvent `json:"items"`
}

// Repository type metadata.
var (
	AnalyticsDatasourceWindowsEvent_Kind             = "AnalyticsDatasourceWindowsEvent"
	AnalyticsDatasourceWindowsEvent_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AnalyticsDatasourceWindowsEvent_Kind}.String()
	AnalyticsDatasourceWindowsEvent_KindAPIVersion   = AnalyticsDatasourceWindowsEvent_Kind + "." + CRDGroupVersion.String()
	AnalyticsDatasourceWindowsEvent_GroupVersionKind = CRDGroupVersion.WithKind(AnalyticsDatasourceWindowsEvent_Kind)
)

func init() {
	SchemeBuilder.Register(&AnalyticsDatasourceWindowsEvent{}, &AnalyticsDatasourceWindowsEventList{})
}
