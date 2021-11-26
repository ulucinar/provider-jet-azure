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

type ConnectionServicePrincipalObservation struct {
}

type ConnectionServicePrincipalParameters struct {

	// +kubebuilder:validation:Required
	ApplicationID *string `json:"applicationId" tf:"application_id,omitempty"`

	// +kubebuilder:validation:Required
	AutomationAccountName *string `json:"automationAccountName" tf:"automation_account_name,omitempty"`

	// +kubebuilder:validation:Required
	CertificateThumbprint *string `json:"certificateThumbprint" tf:"certificate_thumbprint,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Required
	SubscriptionID *string `json:"subscriptionId" tf:"subscription_id,omitempty"`

	// +kubebuilder:validation:Required
	TenantID *string `json:"tenantId" tf:"tenant_id,omitempty"`
}

// ConnectionServicePrincipalSpec defines the desired state of ConnectionServicePrincipal
type ConnectionServicePrincipalSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConnectionServicePrincipalParameters `json:"forProvider"`
}

// ConnectionServicePrincipalStatus defines the observed state of ConnectionServicePrincipal.
type ConnectionServicePrincipalStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConnectionServicePrincipalObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ConnectionServicePrincipal is the Schema for the ConnectionServicePrincipals API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type ConnectionServicePrincipal struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConnectionServicePrincipalSpec   `json:"spec"`
	Status            ConnectionServicePrincipalStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConnectionServicePrincipalList contains a list of ConnectionServicePrincipals
type ConnectionServicePrincipalList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConnectionServicePrincipal `json:"items"`
}

// Repository type metadata.
var (
	ConnectionServicePrincipal_Kind             = "ConnectionServicePrincipal"
	ConnectionServicePrincipal_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ConnectionServicePrincipal_Kind}.String()
	ConnectionServicePrincipal_KindAPIVersion   = ConnectionServicePrincipal_Kind + "." + CRDGroupVersion.String()
	ConnectionServicePrincipal_GroupVersionKind = CRDGroupVersion.WithKind(ConnectionServicePrincipal_Kind)
)

func init() {
	SchemeBuilder.Register(&ConnectionServicePrincipal{}, &ConnectionServicePrincipalList{})
}
