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

type ClusterSettingObservation struct {
}

type ClusterSettingParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type ServiceEnvironmentObservation struct {
	InternalIPAddress *string `json:"internalIpAddress,omitempty" tf:"internal_ip_address,omitempty"`

	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	OutboundIPAddresses []*string `json:"outboundIpAddresses,omitempty" tf:"outbound_ip_addresses,omitempty"`

	ServiceIPAddress *string `json:"serviceIpAddress,omitempty" tf:"service_ip_address,omitempty"`
}

type ServiceEnvironmentParameters struct {

	// +kubebuilder:validation:Optional
	AllowedUserIPCidrs []*string `json:"allowedUserIpCidrs,omitempty" tf:"allowed_user_ip_cidrs,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterSetting []ClusterSettingParameters `json:"clusterSetting,omitempty" tf:"cluster_setting,omitempty"`

	// +kubebuilder:validation:Optional
	FrontEndScaleFactor *int64 `json:"frontEndScaleFactor,omitempty" tf:"front_end_scale_factor,omitempty"`

	// +kubebuilder:validation:Optional
	InternalLoadBalancingMode *string `json:"internalLoadBalancingMode,omitempty" tf:"internal_load_balancing_mode,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	PricingTier *string `json:"pricingTier,omitempty" tf:"pricing_tier,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Required
	SubnetID *string `json:"subnetId" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	UserWhitelistedIPRanges []*string `json:"userWhitelistedIpRanges,omitempty" tf:"user_whitelisted_ip_ranges,omitempty"`
}

// ServiceEnvironmentSpec defines the desired state of ServiceEnvironment
type ServiceEnvironmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceEnvironmentParameters `json:"forProvider"`
}

// ServiceEnvironmentStatus defines the observed state of ServiceEnvironment.
type ServiceEnvironmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceEnvironmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceEnvironment is the Schema for the ServiceEnvironments API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type ServiceEnvironment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceEnvironmentSpec   `json:"spec"`
	Status            ServiceEnvironmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceEnvironmentList contains a list of ServiceEnvironments
type ServiceEnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceEnvironment `json:"items"`
}

// Repository type metadata.
var (
	ServiceEnvironment_Kind             = "ServiceEnvironment"
	ServiceEnvironment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServiceEnvironment_Kind}.String()
	ServiceEnvironment_KindAPIVersion   = ServiceEnvironment_Kind + "." + CRDGroupVersion.String()
	ServiceEnvironment_GroupVersionKind = CRDGroupVersion.WithKind(ServiceEnvironment_Kind)
)

func init() {
	SchemeBuilder.Register(&ServiceEnvironment{}, &ServiceEnvironmentList{})
}
