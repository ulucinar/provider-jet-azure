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

type AppIntegrationAccountBatchConfigurationObservation struct {
}

type AppIntegrationAccountBatchConfigurationParameters struct {

	// +kubebuilder:validation:Required
	BatchGroupName *string `json:"batchGroupName" tf:"batch_group_name,omitempty"`

	// +kubebuilder:validation:Required
	IntegrationAccountName *string `json:"integrationAccountName" tf:"integration_account_name,omitempty"`

	// +kubebuilder:validation:Optional
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ReleaseCriteria []ReleaseCriteriaParameters `json:"releaseCriteria" tf:"release_criteria,omitempty"`

	// +kubebuilder:validation:Required
	ResourceGroupName *string `json:"resourceGroupName" tf:"resource_group_name,omitempty"`
}

type MonthlyObservation struct {
}

type MonthlyParameters struct {

	// +kubebuilder:validation:Required
	Week *int64 `json:"week" tf:"week,omitempty"`

	// +kubebuilder:validation:Required
	Weekday *string `json:"weekday" tf:"weekday,omitempty"`
}

type RecurrenceObservation struct {
}

type RecurrenceParameters struct {

	// +kubebuilder:validation:Optional
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// +kubebuilder:validation:Required
	Frequency *string `json:"frequency" tf:"frequency,omitempty"`

	// +kubebuilder:validation:Required
	Interval *int64 `json:"interval" tf:"interval,omitempty"`

	// +kubebuilder:validation:Optional
	Schedule []ScheduleParameters `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// +kubebuilder:validation:Optional
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`

	// +kubebuilder:validation:Optional
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`
}

type ReleaseCriteriaObservation struct {
}

type ReleaseCriteriaParameters struct {

	// +kubebuilder:validation:Optional
	BatchSize *int64 `json:"batchSize,omitempty" tf:"batch_size,omitempty"`

	// +kubebuilder:validation:Optional
	MessageCount *int64 `json:"messageCount,omitempty" tf:"message_count,omitempty"`

	// +kubebuilder:validation:Optional
	Recurrence []RecurrenceParameters `json:"recurrence,omitempty" tf:"recurrence,omitempty"`
}

type ScheduleObservation struct {
}

type ScheduleParameters struct {

	// +kubebuilder:validation:Optional
	Hours []*int64 `json:"hours,omitempty" tf:"hours,omitempty"`

	// +kubebuilder:validation:Optional
	Minutes []*int64 `json:"minutes,omitempty" tf:"minutes,omitempty"`

	// +kubebuilder:validation:Optional
	MonthDays []*int64 `json:"monthDays,omitempty" tf:"month_days,omitempty"`

	// +kubebuilder:validation:Optional
	Monthly []MonthlyParameters `json:"monthly,omitempty" tf:"monthly,omitempty"`

	// +kubebuilder:validation:Optional
	WeekDays []*string `json:"weekDays,omitempty" tf:"week_days,omitempty"`
}

// AppIntegrationAccountBatchConfigurationSpec defines the desired state of AppIntegrationAccountBatchConfiguration
type AppIntegrationAccountBatchConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AppIntegrationAccountBatchConfigurationParameters `json:"forProvider"`
}

// AppIntegrationAccountBatchConfigurationStatus defines the observed state of AppIntegrationAccountBatchConfiguration.
type AppIntegrationAccountBatchConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AppIntegrationAccountBatchConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AppIntegrationAccountBatchConfiguration is the Schema for the AppIntegrationAccountBatchConfigurations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type AppIntegrationAccountBatchConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppIntegrationAccountBatchConfigurationSpec   `json:"spec"`
	Status            AppIntegrationAccountBatchConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppIntegrationAccountBatchConfigurationList contains a list of AppIntegrationAccountBatchConfigurations
type AppIntegrationAccountBatchConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppIntegrationAccountBatchConfiguration `json:"items"`
}

// Repository type metadata.
var (
	AppIntegrationAccountBatchConfiguration_Kind             = "AppIntegrationAccountBatchConfiguration"
	AppIntegrationAccountBatchConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AppIntegrationAccountBatchConfiguration_Kind}.String()
	AppIntegrationAccountBatchConfiguration_KindAPIVersion   = AppIntegrationAccountBatchConfiguration_Kind + "." + CRDGroupVersion.String()
	AppIntegrationAccountBatchConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(AppIntegrationAccountBatchConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&AppIntegrationAccountBatchConfiguration{}, &AppIntegrationAccountBatchConfigurationList{})
}
