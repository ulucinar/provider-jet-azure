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

type DNSConfigObservation struct {
}

type DNSConfigParameters struct {

	// +kubebuilder:validation:Required
	RelativeName *string `json:"relativeName" tf:"relative_name,omitempty"`

	// +kubebuilder:validation:Required
	TTL *int64 `json:"ttl" tf:"ttl,omitempty"`
}

type ManagerProfileObservation struct {
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`
}

type ManagerProfileParameters struct {

	// +kubebuilder:validation:Required
	DNSConfig []DNSConfigParameters `json:"dnsConfig" tf:"dns_config,omitempty"`

	// +kubebuilder:validation:Optional
	MaxReturn *int64 `json:"maxReturn,omitempty" tf:"max_return,omitempty"`

	// +kubebuilder:validation:Required
	MonitorConfig []MonitorConfigParameters `json:"monitorConfig" tf:"monitor_config,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	ProfileStatus *string `json:"profileStatus,omitempty" tf:"profile_status,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	TrafficRoutingMethod *string `json:"trafficRoutingMethod" tf:"traffic_routing_method,omitempty"`

	// +kubebuilder:validation:Optional
	TrafficViewEnabled *bool `json:"trafficViewEnabled,omitempty" tf:"traffic_view_enabled,omitempty"`
}

type MonitorConfigCustomHeaderObservation struct {
}

type MonitorConfigCustomHeaderParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type MonitorConfigObservation struct {
}

type MonitorConfigParameters struct {

	// +kubebuilder:validation:Optional
	CustomHeader []MonitorConfigCustomHeaderParameters `json:"customHeader,omitempty" tf:"custom_header,omitempty"`

	// +kubebuilder:validation:Optional
	ExpectedStatusCodeRanges []*string `json:"expectedStatusCodeRanges,omitempty" tf:"expected_status_code_ranges,omitempty"`

	// +kubebuilder:validation:Optional
	IntervalInSeconds *int64 `json:"intervalInSeconds,omitempty" tf:"interval_in_seconds,omitempty"`

	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// +kubebuilder:validation:Required
	Port *int64 `json:"port" tf:"port,omitempty"`

	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Optional
	TimeoutInSeconds *int64 `json:"timeoutInSeconds,omitempty" tf:"timeout_in_seconds,omitempty"`

	// +kubebuilder:validation:Optional
	ToleratedNumberOfFailures *int64 `json:"toleratedNumberOfFailures,omitempty" tf:"tolerated_number_of_failures,omitempty"`
}

// ManagerProfileSpec defines the desired state of ManagerProfile
type ManagerProfileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagerProfileParameters `json:"forProvider"`
}

// ManagerProfileStatus defines the observed state of ManagerProfile.
type ManagerProfileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagerProfileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ManagerProfile is the Schema for the ManagerProfiles API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type ManagerProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagerProfileSpec   `json:"spec"`
	Status            ManagerProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagerProfileList contains a list of ManagerProfiles
type ManagerProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagerProfile `json:"items"`
}

// Repository type metadata.
var (
	ManagerProfile_Kind             = "ManagerProfile"
	ManagerProfile_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ManagerProfile_Kind}.String()
	ManagerProfile_KindAPIVersion   = ManagerProfile_Kind + "." + CRDGroupVersion.String()
	ManagerProfile_GroupVersionKind = CRDGroupVersion.WithKind(ManagerProfile_Kind)
)

func init() {
	SchemeBuilder.Register(&ManagerProfile{}, &ManagerProfileList{})
}
