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

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ActionObservation struct {
}

type ActionParameters struct {

	// +kubebuilder:validation:Required
	ActionType *string `json:"actionType" tf:"action_type,omitempty"`
}

type CertificateAttributeObservation struct {
	Created *string `json:"created,omitempty" tf:"created,omitempty"`

	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	Expires *string `json:"expires,omitempty" tf:"expires,omitempty"`

	NotBefore *string `json:"notBefore,omitempty" tf:"not_before,omitempty"`

	RecoveryLevel *string `json:"recoveryLevel,omitempty" tf:"recovery_level,omitempty"`

	Updated *string `json:"updated,omitempty" tf:"updated,omitempty"`
}

type CertificateAttributeParameters struct {
}

type CertificateCertificateObservation struct {
}

type CertificateCertificateParameters struct {

	// +kubebuilder:validation:Required
	ContentsSecretRef v1.SecretKeySelector `json:"contentsSecretRef" tf:"-"`

	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`
}

type CertificateObservation struct {
	CertificateAttribute []CertificateAttributeObservation `json:"certificateAttribute,omitempty" tf:"certificate_attribute,omitempty"`

	CertificateData *string `json:"certificateData,omitempty" tf:"certificate_data,omitempty"`

	CertificateDataBase64 *string `json:"certificateDataBase64,omitempty" tf:"certificate_data_base64,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	SecretID *string `json:"secretId,omitempty" tf:"secret_id,omitempty"`

	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint,omitempty"`

	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type CertificateParameters struct {

	// +kubebuilder:validation:Optional
	Certificate []CertificateCertificateParameters `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// +kubebuilder:validation:Required
	CertificatePolicy []CertificatePolicyParameters `json:"certificatePolicy" tf:"certificate_policy,omitempty"`

	// +crossplane:generate:reference:type=Vault
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-jet-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	KeyVaultID *string `json:"keyVaultId,omitempty" tf:"key_vault_id,omitempty"`

	// +kubebuilder:validation:Optional
	KeyVaultIDRef *v1.Reference `json:"keyVaultIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	KeyVaultIDSelector *v1.Selector `json:"keyVaultIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type CertificatePolicyObservation struct {
}

type CertificatePolicyParameters struct {

	// +kubebuilder:validation:Required
	IssuerParameters []IssuerParametersParameters `json:"issuerParameters" tf:"issuer_parameters,omitempty"`

	// +kubebuilder:validation:Required
	KeyProperties []KeyPropertiesParameters `json:"keyProperties" tf:"key_properties,omitempty"`

	// +kubebuilder:validation:Optional
	LifetimeAction []LifetimeActionParameters `json:"lifetimeAction,omitempty" tf:"lifetime_action,omitempty"`

	// +kubebuilder:validation:Required
	SecretProperties []SecretPropertiesParameters `json:"secretProperties" tf:"secret_properties,omitempty"`

	// +kubebuilder:validation:Optional
	X509CertificateProperties []X509CertificatePropertiesParameters `json:"x509CertificateProperties,omitempty" tf:"x509_certificate_properties,omitempty"`
}

type IssuerParametersObservation struct {
}

type IssuerParametersParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type KeyPropertiesObservation struct {
}

type KeyPropertiesParameters struct {

	// +kubebuilder:validation:Optional
	Curve *string `json:"curve,omitempty" tf:"curve,omitempty"`

	// +kubebuilder:validation:Required
	Exportable *bool `json:"exportable" tf:"exportable,omitempty"`

	// +kubebuilder:validation:Optional
	KeySize *int64 `json:"keySize,omitempty" tf:"key_size,omitempty"`

	// +kubebuilder:validation:Required
	KeyType *string `json:"keyType" tf:"key_type,omitempty"`

	// +kubebuilder:validation:Required
	ReuseKey *bool `json:"reuseKey" tf:"reuse_key,omitempty"`
}

type LifetimeActionObservation struct {
}

type LifetimeActionParameters struct {

	// +kubebuilder:validation:Required
	Action []ActionParameters `json:"action" tf:"action,omitempty"`

	// +kubebuilder:validation:Required
	Trigger []TriggerParameters `json:"trigger" tf:"trigger,omitempty"`
}

type SecretPropertiesObservation struct {
}

type SecretPropertiesParameters struct {

	// +kubebuilder:validation:Required
	ContentType *string `json:"contentType" tf:"content_type,omitempty"`
}

type SubjectAlternativeNamesObservation struct {
}

type SubjectAlternativeNamesParameters struct {

	// +kubebuilder:validation:Optional
	DNSNames []*string `json:"dnsNames,omitempty" tf:"dns_names,omitempty"`

	// +kubebuilder:validation:Optional
	Emails []*string `json:"emails,omitempty" tf:"emails,omitempty"`

	// +kubebuilder:validation:Optional
	Upns []*string `json:"upns,omitempty" tf:"upns,omitempty"`
}

type TriggerObservation struct {
}

type TriggerParameters struct {

	// +kubebuilder:validation:Optional
	DaysBeforeExpiry *int64 `json:"daysBeforeExpiry,omitempty" tf:"days_before_expiry,omitempty"`

	// +kubebuilder:validation:Optional
	LifetimePercentage *int64 `json:"lifetimePercentage,omitempty" tf:"lifetime_percentage,omitempty"`
}

type X509CertificatePropertiesObservation struct {
}

type X509CertificatePropertiesParameters struct {

	// +kubebuilder:validation:Optional
	ExtendedKeyUsage []*string `json:"extendedKeyUsage,omitempty" tf:"extended_key_usage,omitempty"`

	// +kubebuilder:validation:Required
	KeyUsage []*string `json:"keyUsage" tf:"key_usage,omitempty"`

	// +kubebuilder:validation:Required
	Subject *string `json:"subject" tf:"subject,omitempty"`

	// +kubebuilder:validation:Optional
	SubjectAlternativeNames []SubjectAlternativeNamesParameters `json:"subjectAlternativeNames,omitempty" tf:"subject_alternative_names,omitempty"`

	// +kubebuilder:validation:Required
	ValidityInMonths *int64 `json:"validityInMonths" tf:"validity_in_months,omitempty"`
}

// CertificateSpec defines the desired state of Certificate
type CertificateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CertificateParameters `json:"forProvider"`
}

// CertificateStatus defines the observed state of Certificate.
type CertificateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CertificateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Certificate is the Schema for the Certificates API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azurejet}
type Certificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CertificateSpec   `json:"spec"`
	Status            CertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CertificateList contains a list of Certificates
type CertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Certificate `json:"items"`
}

// Repository type metadata.
var (
	Certificate_Kind             = "Certificate"
	Certificate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Certificate_Kind}.String()
	Certificate_KindAPIVersion   = Certificate_Kind + "." + CRDGroupVersion.String()
	Certificate_GroupVersionKind = CRDGroupVersion.WithKind(Certificate_Kind)
)

func init() {
	SchemeBuilder.Register(&Certificate{}, &CertificateList{})
}
