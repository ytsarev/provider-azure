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

type BotWebAppInitParameters struct {

	// The Application Insights Application ID to associate with the Web App Bot.
	DeveloperAppInsightsApplicationID *string `json:"developerAppInsightsApplicationId,omitempty" tf:"developer_app_insights_application_id,omitempty"`

	// The Application Insights Key to associate with the Web App Bot.
	DeveloperAppInsightsKey *string `json:"developerAppInsightsKey,omitempty" tf:"developer_app_insights_key,omitempty"`

	// The name of the Web App Bot will be displayed as. This defaults to name if not specified.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The Web App Bot endpoint.
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A list of LUIS App IDs to associate with the Web App Bot.
	LuisAppIds []*string `json:"luisAppIds,omitempty" tf:"luis_app_ids,omitempty"`

	// The Microsoft Application ID for the Web App Bot. Changing this forces a new resource to be created.
	MicrosoftAppID *string `json:"microsoftAppId,omitempty" tf:"microsoft_app_id,omitempty"`

	// The SKU of the Web App Bot. Valid values include F0 or S1. Changing this forces a new resource to be created.
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type BotWebAppObservation struct {

	// The Application Insights Application ID to associate with the Web App Bot.
	DeveloperAppInsightsApplicationID *string `json:"developerAppInsightsApplicationId,omitempty" tf:"developer_app_insights_application_id,omitempty"`

	// The Application Insights Key to associate with the Web App Bot.
	DeveloperAppInsightsKey *string `json:"developerAppInsightsKey,omitempty" tf:"developer_app_insights_key,omitempty"`

	// The name of the Web App Bot will be displayed as. This defaults to name if not specified.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The Web App Bot endpoint.
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// The ID of the Bot Web App.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A list of LUIS App IDs to associate with the Web App Bot.
	LuisAppIds []*string `json:"luisAppIds,omitempty" tf:"luis_app_ids,omitempty"`

	// The Microsoft Application ID for the Web App Bot. Changing this forces a new resource to be created.
	MicrosoftAppID *string `json:"microsoftAppId,omitempty" tf:"microsoft_app_id,omitempty"`

	// The name of the resource group in which to create the Web App Bot. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The SKU of the Web App Bot. Valid values include F0 or S1. Changing this forces a new resource to be created.
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type BotWebAppParameters struct {

	// The Application Insights API Key to associate with the Web App Bot.
	// +kubebuilder:validation:Optional
	DeveloperAppInsightsAPIKeySecretRef *v1.SecretKeySelector `json:"developerAppInsightsApiKeySecretRef,omitempty" tf:"-"`

	// The Application Insights Application ID to associate with the Web App Bot.
	// +kubebuilder:validation:Optional
	DeveloperAppInsightsApplicationID *string `json:"developerAppInsightsApplicationId,omitempty" tf:"developer_app_insights_application_id,omitempty"`

	// The Application Insights Key to associate with the Web App Bot.
	// +kubebuilder:validation:Optional
	DeveloperAppInsightsKey *string `json:"developerAppInsightsKey,omitempty" tf:"developer_app_insights_key,omitempty"`

	// The name of the Web App Bot will be displayed as. This defaults to name if not specified.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The Web App Bot endpoint.
	// +kubebuilder:validation:Optional
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// The supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A list of LUIS App IDs to associate with the Web App Bot.
	// +kubebuilder:validation:Optional
	LuisAppIds []*string `json:"luisAppIds,omitempty" tf:"luis_app_ids,omitempty"`

	// The LUIS key to associate with the Web App Bot.
	// +kubebuilder:validation:Optional
	LuisKeySecretRef *v1.SecretKeySelector `json:"luisKeySecretRef,omitempty" tf:"-"`

	// The Microsoft Application ID for the Web App Bot. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	MicrosoftAppID *string `json:"microsoftAppId,omitempty" tf:"microsoft_app_id,omitempty"`

	// The name of the resource group in which to create the Web App Bot. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The SKU of the Web App Bot. Valid values include F0 or S1. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// BotWebAppSpec defines the desired state of BotWebApp
type BotWebAppSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BotWebAppParameters `json:"forProvider"`
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
	InitProvider BotWebAppInitParameters `json:"initProvider,omitempty"`
}

// BotWebAppStatus defines the observed state of BotWebApp.
type BotWebAppStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BotWebAppObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BotWebApp is the Schema for the BotWebApps API. Manages a Web App Bot.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type BotWebApp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.microsoftAppId) || (has(self.initProvider) && has(self.initProvider.microsoftAppId))",message="spec.forProvider.microsoftAppId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.sku) || (has(self.initProvider) && has(self.initProvider.sku))",message="spec.forProvider.sku is a required parameter"
	Spec   BotWebAppSpec   `json:"spec"`
	Status BotWebAppStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BotWebAppList contains a list of BotWebApps
type BotWebAppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BotWebApp `json:"items"`
}

// Repository type metadata.
var (
	BotWebApp_Kind             = "BotWebApp"
	BotWebApp_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BotWebApp_Kind}.String()
	BotWebApp_KindAPIVersion   = BotWebApp_Kind + "." + CRDGroupVersion.String()
	BotWebApp_GroupVersionKind = CRDGroupVersion.WithKind(BotWebApp_Kind)
)

func init() {
	SchemeBuilder.Register(&BotWebApp{}, &BotWebAppList{})
}
