// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type KeyVaultKeyInitParameters struct {

	// Specifies the name of an existing Key Vault Data Factory Linked Service.
	LinkedServiceName *string `json:"linkedServiceName,omitempty" tf:"linked_service_name,omitempty"`

	// Specifies the secret name in Azure Key Vault that stores the system key of the Azure Function.
	SecretName *string `json:"secretName,omitempty" tf:"secret_name,omitempty"`
}

type KeyVaultKeyObservation struct {

	// Specifies the name of an existing Key Vault Data Factory Linked Service.
	LinkedServiceName *string `json:"linkedServiceName,omitempty" tf:"linked_service_name,omitempty"`

	// Specifies the secret name in Azure Key Vault that stores the system key of the Azure Function.
	SecretName *string `json:"secretName,omitempty" tf:"secret_name,omitempty"`
}

type KeyVaultKeyParameters struct {

	// Specifies the name of an existing Key Vault Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	LinkedServiceName *string `json:"linkedServiceName" tf:"linked_service_name,omitempty"`

	// Specifies the secret name in Azure Key Vault that stores the system key of the Azure Function.
	// +kubebuilder:validation:Optional
	SecretName *string `json:"secretName" tf:"secret_name,omitempty"`
}

type LinkedServiceAzureFunctionInitParameters struct {

	// A map of additional properties to associate with the Data Factory Linked Service.
	// +mapType=granular
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The description for the Data Factory Linked Service.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// A key_vault_key block as defined below. Use this Argument to store the system key of the Azure Function in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service. Exactly one of either key or key_vault_key is required.
	KeyVaultKey []KeyVaultKeyInitParameters `json:"keyVaultKey,omitempty" tf:"key_vault_key,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The url of the Azure Function.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type LinkedServiceAzureFunctionObservation struct {

	// A map of additional properties to associate with the Data Factory Linked Service.
	// +mapType=granular
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// The description for the Data Factory Linked Service.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the Data Factory Linked Service.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The integration runtime reference to associate with the Data Factory Linked Service.
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// A key_vault_key block as defined below. Use this Argument to store the system key of the Azure Function in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service. Exactly one of either key or key_vault_key is required.
	KeyVaultKey []KeyVaultKeyObservation `json:"keyVaultKey,omitempty" tf:"key_vault_key,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The url of the Azure Function.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type LinkedServiceAzureFunctionParameters struct {

	// A map of additional properties to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.Factory
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// Reference to a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDRef *v1.Reference `json:"dataFactoryIdRef,omitempty" tf:"-"`

	// Selector for a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDSelector *v1.Selector `json:"dataFactoryIdSelector,omitempty" tf:"-"`

	// The description for the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The integration runtime reference to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	IntegrationRuntimeName *string `json:"integrationRuntimeName,omitempty" tf:"integration_runtime_name,omitempty"`

	// The system key of the Azure Function. Exactly one of either key or key_vault_key is required
	// +kubebuilder:validation:Optional
	KeySecretRef *v1.SecretKeySelector `json:"keySecretRef,omitempty" tf:"-"`

	// A key_vault_key block as defined below. Use this Argument to store the system key of the Azure Function in an existing Key Vault. It needs an existing Key Vault Data Factory Linked Service. Exactly one of either key or key_vault_key is required.
	// +kubebuilder:validation:Optional
	KeyVaultKey []KeyVaultKeyParameters `json:"keyVaultKey,omitempty" tf:"key_vault_key,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The url of the Azure Function.
	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

// LinkedServiceAzureFunctionSpec defines the desired state of LinkedServiceAzureFunction
type LinkedServiceAzureFunctionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LinkedServiceAzureFunctionParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider LinkedServiceAzureFunctionInitParameters `json:"initProvider,omitempty"`
}

// LinkedServiceAzureFunctionStatus defines the observed state of LinkedServiceAzureFunction.
type LinkedServiceAzureFunctionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LinkedServiceAzureFunctionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LinkedServiceAzureFunction is the Schema for the LinkedServiceAzureFunctions API. Manages a Linked Service (connection) between an Azure Function Account and Azure Data Factory.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LinkedServiceAzureFunction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.url) || (has(self.initProvider) && has(self.initProvider.url))",message="spec.forProvider.url is a required parameter"
	Spec   LinkedServiceAzureFunctionSpec   `json:"spec"`
	Status LinkedServiceAzureFunctionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LinkedServiceAzureFunctionList contains a list of LinkedServiceAzureFunctions
type LinkedServiceAzureFunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LinkedServiceAzureFunction `json:"items"`
}

// Repository type metadata.
var (
	LinkedServiceAzureFunction_Kind             = "LinkedServiceAzureFunction"
	LinkedServiceAzureFunction_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LinkedServiceAzureFunction_Kind}.String()
	LinkedServiceAzureFunction_KindAPIVersion   = LinkedServiceAzureFunction_Kind + "." + CRDGroupVersion.String()
	LinkedServiceAzureFunction_GroupVersionKind = CRDGroupVersion.WithKind(LinkedServiceAzureFunction_Kind)
)

func init() {
	SchemeBuilder.Register(&LinkedServiceAzureFunction{}, &LinkedServiceAzureFunctionList{})
}
