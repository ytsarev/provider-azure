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

type AccountInitParameters struct {

	// If kind is TextAnalytics this specifies the ID of the Search service.
	CustomQuestionAnsweringSearchServiceID *string `json:"customQuestionAnsweringSearchServiceId,omitempty" tf:"custom_question_answering_search_service_id,omitempty"`

	// The subdomain name used for token-based authentication. Changing this forces a new resource to be created.
	CustomSubdomainName *string `json:"customSubdomainName,omitempty" tf:"custom_subdomain_name,omitempty"`

	// A customer_managed_key block as documented below.
	CustomerManagedKey []CustomerManagedKeyInitParameters `json:"customerManagedKey,omitempty" tf:"customer_managed_key,omitempty"`

	// Whether to enable the dynamic throttling for this Cognitive Service Account.
	DynamicThrottlingEnabled *bool `json:"dynamicThrottlingEnabled,omitempty" tf:"dynamic_throttling_enabled,omitempty"`

	// List of FQDNs allowed for the Cognitive Account.
	Fqdns []*string `json:"fqdns,omitempty" tf:"fqdns,omitempty"`

	// An identity block as defined below.
	Identity []IdentityInitParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// Specifies the type of Cognitive Service Account that should be created. Possible values are Academic, AnomalyDetector, Bing.Autosuggest, Bing.Autosuggest.v7, Bing.CustomSearch, Bing.Search, Bing.Search.v7, Bing.Speech, Bing.SpellCheck, Bing.SpellCheck.v7, CognitiveServices, ComputerVision, ContentModerator, CustomSpeech, CustomVision.Prediction, CustomVision.Training, Emotion, Face, FormRecognizer, ImmersiveReader, LUIS, LUIS.Authoring, MetricsAdvisor, OpenAI, Personalizer, QnAMaker, Recommendations, SpeakerRecognition, Speech, SpeechServices, SpeechTranslation, TextAnalytics, TextTranslation and WebLM. Changing this forces a new resource to be created.
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// Whether local authentication methods is enabled for the Cognitive Account. Defaults to true.
	LocalAuthEnabled *bool `json:"localAuthEnabled,omitempty" tf:"local_auth_enabled,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The Azure AD Client ID (Application ID). This attribute is only set when kind is MetricsAdvisor. Changing this forces a new resource to be created.
	MetricsAdvisorAADClientID *string `json:"metricsAdvisorAadClientId,omitempty" tf:"metrics_advisor_aad_client_id,omitempty"`

	// The Azure AD Tenant ID. This attribute is only set when kind is MetricsAdvisor. Changing this forces a new resource to be created.
	MetricsAdvisorAADTenantID *string `json:"metricsAdvisorAadTenantId,omitempty" tf:"metrics_advisor_aad_tenant_id,omitempty"`

	// The super user of Metrics Advisor. This attribute is only set when kind is MetricsAdvisor. Changing this forces a new resource to be created.
	MetricsAdvisorSuperUserName *string `json:"metricsAdvisorSuperUserName,omitempty" tf:"metrics_advisor_super_user_name,omitempty"`

	// The website name of Metrics Advisor. This attribute is only set when kind is MetricsAdvisor. Changing this forces a new resource to be created.
	MetricsAdvisorWebsiteName *string `json:"metricsAdvisorWebsiteName,omitempty" tf:"metrics_advisor_website_name,omitempty"`

	// A network_acls block as defined below.
	NetworkAcls []NetworkAclsInitParameters `json:"networkAcls,omitempty" tf:"network_acls,omitempty"`

	// Whether outbound network access is restricted for the Cognitive Account. Defaults to false.
	OutboundNetworkAccessRestricted *bool `json:"outboundNetworkAccessRestricted,omitempty" tf:"outbound_network_access_restricted,omitempty"`

	// Whether public network access is allowed for the Cognitive Account. Defaults to true.
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// A URL to link a QnAMaker cognitive account to a QnA runtime.
	QnaRuntimeEndpoint *string `json:"qnaRuntimeEndpoint,omitempty" tf:"qna_runtime_endpoint,omitempty"`

	// Specifies the SKU Name for this Cognitive Service Account. Possible values are F0, F1, S0, S, S1, S2, S3, S4, S5, S6, P0, P1, P2, E0 and DC0.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// A storage block as defined below.
	Storage []StorageInitParameters `json:"storage,omitempty" tf:"storage,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type AccountObservation struct {

	// If kind is TextAnalytics this specifies the ID of the Search service.
	CustomQuestionAnsweringSearchServiceID *string `json:"customQuestionAnsweringSearchServiceId,omitempty" tf:"custom_question_answering_search_service_id,omitempty"`

	// The subdomain name used for token-based authentication. Changing this forces a new resource to be created.
	CustomSubdomainName *string `json:"customSubdomainName,omitempty" tf:"custom_subdomain_name,omitempty"`

	// A customer_managed_key block as documented below.
	CustomerManagedKey []CustomerManagedKeyObservation `json:"customerManagedKey,omitempty" tf:"customer_managed_key,omitempty"`

	// Whether to enable the dynamic throttling for this Cognitive Service Account.
	DynamicThrottlingEnabled *bool `json:"dynamicThrottlingEnabled,omitempty" tf:"dynamic_throttling_enabled,omitempty"`

	// The endpoint used to connect to the Cognitive Service Account.
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// List of FQDNs allowed for the Cognitive Account.
	Fqdns []*string `json:"fqdns,omitempty" tf:"fqdns,omitempty"`

	// The ID of the Cognitive Service Account.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity block as defined below.
	Identity []IdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// Specifies the type of Cognitive Service Account that should be created. Possible values are Academic, AnomalyDetector, Bing.Autosuggest, Bing.Autosuggest.v7, Bing.CustomSearch, Bing.Search, Bing.Search.v7, Bing.Speech, Bing.SpellCheck, Bing.SpellCheck.v7, CognitiveServices, ComputerVision, ContentModerator, CustomSpeech, CustomVision.Prediction, CustomVision.Training, Emotion, Face, FormRecognizer, ImmersiveReader, LUIS, LUIS.Authoring, MetricsAdvisor, OpenAI, Personalizer, QnAMaker, Recommendations, SpeakerRecognition, Speech, SpeechServices, SpeechTranslation, TextAnalytics, TextTranslation and WebLM. Changing this forces a new resource to be created.
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// Whether local authentication methods is enabled for the Cognitive Account. Defaults to true.
	LocalAuthEnabled *bool `json:"localAuthEnabled,omitempty" tf:"local_auth_enabled,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The Azure AD Client ID (Application ID). This attribute is only set when kind is MetricsAdvisor. Changing this forces a new resource to be created.
	MetricsAdvisorAADClientID *string `json:"metricsAdvisorAadClientId,omitempty" tf:"metrics_advisor_aad_client_id,omitempty"`

	// The Azure AD Tenant ID. This attribute is only set when kind is MetricsAdvisor. Changing this forces a new resource to be created.
	MetricsAdvisorAADTenantID *string `json:"metricsAdvisorAadTenantId,omitempty" tf:"metrics_advisor_aad_tenant_id,omitempty"`

	// The super user of Metrics Advisor. This attribute is only set when kind is MetricsAdvisor. Changing this forces a new resource to be created.
	MetricsAdvisorSuperUserName *string `json:"metricsAdvisorSuperUserName,omitempty" tf:"metrics_advisor_super_user_name,omitempty"`

	// The website name of Metrics Advisor. This attribute is only set when kind is MetricsAdvisor. Changing this forces a new resource to be created.
	MetricsAdvisorWebsiteName *string `json:"metricsAdvisorWebsiteName,omitempty" tf:"metrics_advisor_website_name,omitempty"`

	// A network_acls block as defined below.
	NetworkAcls []NetworkAclsObservation `json:"networkAcls,omitempty" tf:"network_acls,omitempty"`

	// Whether outbound network access is restricted for the Cognitive Account. Defaults to false.
	OutboundNetworkAccessRestricted *bool `json:"outboundNetworkAccessRestricted,omitempty" tf:"outbound_network_access_restricted,omitempty"`

	// Whether public network access is allowed for the Cognitive Account. Defaults to true.
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// A URL to link a QnAMaker cognitive account to a QnA runtime.
	QnaRuntimeEndpoint *string `json:"qnaRuntimeEndpoint,omitempty" tf:"qna_runtime_endpoint,omitempty"`

	// The name of the resource group in which the Cognitive Service Account is created. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Specifies the SKU Name for this Cognitive Service Account. Possible values are F0, F1, S0, S, S1, S2, S3, S4, S5, S6, P0, P1, P2, E0 and DC0.
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// A storage block as defined below.
	Storage []StorageObservation `json:"storage,omitempty" tf:"storage,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type AccountParameters struct {

	// If kind is TextAnalytics this specifies the ID of the Search service.
	// +kubebuilder:validation:Optional
	CustomQuestionAnsweringSearchServiceID *string `json:"customQuestionAnsweringSearchServiceId,omitempty" tf:"custom_question_answering_search_service_id,omitempty"`

	// If kind is TextAnalytics this specifies the key of the Search service.
	// +kubebuilder:validation:Optional
	CustomQuestionAnsweringSearchServiceKeySecretRef *v1.SecretKeySelector `json:"customQuestionAnsweringSearchServiceKeySecretRef,omitempty" tf:"-"`

	// The subdomain name used for token-based authentication. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	CustomSubdomainName *string `json:"customSubdomainName,omitempty" tf:"custom_subdomain_name,omitempty"`

	// A customer_managed_key block as documented below.
	// +kubebuilder:validation:Optional
	CustomerManagedKey []CustomerManagedKeyParameters `json:"customerManagedKey,omitempty" tf:"customer_managed_key,omitempty"`

	// Whether to enable the dynamic throttling for this Cognitive Service Account.
	// +kubebuilder:validation:Optional
	DynamicThrottlingEnabled *bool `json:"dynamicThrottlingEnabled,omitempty" tf:"dynamic_throttling_enabled,omitempty"`

	// List of FQDNs allowed for the Cognitive Account.
	// +kubebuilder:validation:Optional
	Fqdns []*string `json:"fqdns,omitempty" tf:"fqdns,omitempty"`

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity []IdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// Specifies the type of Cognitive Service Account that should be created. Possible values are Academic, AnomalyDetector, Bing.Autosuggest, Bing.Autosuggest.v7, Bing.CustomSearch, Bing.Search, Bing.Search.v7, Bing.Speech, Bing.SpellCheck, Bing.SpellCheck.v7, CognitiveServices, ComputerVision, ContentModerator, CustomSpeech, CustomVision.Prediction, CustomVision.Training, Emotion, Face, FormRecognizer, ImmersiveReader, LUIS, LUIS.Authoring, MetricsAdvisor, OpenAI, Personalizer, QnAMaker, Recommendations, SpeakerRecognition, Speech, SpeechServices, SpeechTranslation, TextAnalytics, TextTranslation and WebLM. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// Whether local authentication methods is enabled for the Cognitive Account. Defaults to true.
	// +kubebuilder:validation:Optional
	LocalAuthEnabled *bool `json:"localAuthEnabled,omitempty" tf:"local_auth_enabled,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The Azure AD Client ID (Application ID). This attribute is only set when kind is MetricsAdvisor. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	MetricsAdvisorAADClientID *string `json:"metricsAdvisorAadClientId,omitempty" tf:"metrics_advisor_aad_client_id,omitempty"`

	// The Azure AD Tenant ID. This attribute is only set when kind is MetricsAdvisor. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	MetricsAdvisorAADTenantID *string `json:"metricsAdvisorAadTenantId,omitempty" tf:"metrics_advisor_aad_tenant_id,omitempty"`

	// The super user of Metrics Advisor. This attribute is only set when kind is MetricsAdvisor. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	MetricsAdvisorSuperUserName *string `json:"metricsAdvisorSuperUserName,omitempty" tf:"metrics_advisor_super_user_name,omitempty"`

	// The website name of Metrics Advisor. This attribute is only set when kind is MetricsAdvisor. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	MetricsAdvisorWebsiteName *string `json:"metricsAdvisorWebsiteName,omitempty" tf:"metrics_advisor_website_name,omitempty"`

	// A network_acls block as defined below.
	// +kubebuilder:validation:Optional
	NetworkAcls []NetworkAclsParameters `json:"networkAcls,omitempty" tf:"network_acls,omitempty"`

	// Whether outbound network access is restricted for the Cognitive Account. Defaults to false.
	// +kubebuilder:validation:Optional
	OutboundNetworkAccessRestricted *bool `json:"outboundNetworkAccessRestricted,omitempty" tf:"outbound_network_access_restricted,omitempty"`

	// Whether public network access is allowed for the Cognitive Account. Defaults to true.
	// +kubebuilder:validation:Optional
	PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

	// A URL to link a QnAMaker cognitive account to a QnA runtime.
	// +kubebuilder:validation:Optional
	QnaRuntimeEndpoint *string `json:"qnaRuntimeEndpoint,omitempty" tf:"qna_runtime_endpoint,omitempty"`

	// The name of the resource group in which the Cognitive Service Account is created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Specifies the SKU Name for this Cognitive Service Account. Possible values are F0, F1, S0, S, S1, S2, S3, S4, S5, S6, P0, P1, P2, E0 and DC0.
	// +kubebuilder:validation:Optional
	SkuName *string `json:"skuName,omitempty" tf:"sku_name,omitempty"`

	// A storage block as defined below.
	// +kubebuilder:validation:Optional
	Storage []StorageParameters `json:"storage,omitempty" tf:"storage,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type CustomerManagedKeyInitParameters struct {

	// The Client ID of the User Assigned Identity that has access to the key. This property only needs to be specified when there're multiple identities attached to the Cognitive Account.
	IdentityClientID *string `json:"identityClientId,omitempty" tf:"identity_client_id,omitempty"`

	// The ID of the Key Vault Key which should be used to Encrypt the data in this Cognitive Account.
	KeyVaultKeyID *string `json:"keyVaultKeyId,omitempty" tf:"key_vault_key_id,omitempty"`
}

type CustomerManagedKeyObservation struct {

	// The Client ID of the User Assigned Identity that has access to the key. This property only needs to be specified when there're multiple identities attached to the Cognitive Account.
	IdentityClientID *string `json:"identityClientId,omitempty" tf:"identity_client_id,omitempty"`

	// The ID of the Key Vault Key which should be used to Encrypt the data in this Cognitive Account.
	KeyVaultKeyID *string `json:"keyVaultKeyId,omitempty" tf:"key_vault_key_id,omitempty"`
}

type CustomerManagedKeyParameters struct {

	// The Client ID of the User Assigned Identity that has access to the key. This property only needs to be specified when there're multiple identities attached to the Cognitive Account.
	// +kubebuilder:validation:Optional
	IdentityClientID *string `json:"identityClientId,omitempty" tf:"identity_client_id,omitempty"`

	// The ID of the Key Vault Key which should be used to Encrypt the data in this Cognitive Account.
	// +kubebuilder:validation:Optional
	KeyVaultKeyID *string `json:"keyVaultKeyId" tf:"key_vault_key_id,omitempty"`
}

type IdentityInitParameters struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Cognitive Account.
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this Cognitive Account. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both).
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityObservation struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Cognitive Account.
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// The Principal ID associated with this Managed Service Identity.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The Tenant ID associated with this Managed Service Identity.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this Cognitive Account. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both).
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IdentityParameters struct {

	// Specifies a list of User Assigned Managed Identity IDs to be assigned to this Cognitive Account.
	// +kubebuilder:validation:Optional
	// +listType=set
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the type of Managed Service Identity that should be configured on this Cognitive Account. Possible values are SystemAssigned, UserAssigned, SystemAssigned, UserAssigned (to enable both).
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type NetworkAclsInitParameters struct {

	// The Default Action to use when no rules match from ip_rules / virtual_network_rules. Possible values are Allow and Deny.
	DefaultAction *string `json:"defaultAction,omitempty" tf:"default_action,omitempty"`

	// One or more IP Addresses, or CIDR Blocks which should be able to access the Cognitive Account.
	// +listType=set
	IPRules []*string `json:"ipRules,omitempty" tf:"ip_rules,omitempty"`

	// A virtual_network_rules block as defined below.
	VirtualNetworkRules []VirtualNetworkRulesInitParameters `json:"virtualNetworkRules,omitempty" tf:"virtual_network_rules,omitempty"`
}

type NetworkAclsObservation struct {

	// The Default Action to use when no rules match from ip_rules / virtual_network_rules. Possible values are Allow and Deny.
	DefaultAction *string `json:"defaultAction,omitempty" tf:"default_action,omitempty"`

	// One or more IP Addresses, or CIDR Blocks which should be able to access the Cognitive Account.
	// +listType=set
	IPRules []*string `json:"ipRules,omitempty" tf:"ip_rules,omitempty"`

	// A virtual_network_rules block as defined below.
	VirtualNetworkRules []VirtualNetworkRulesObservation `json:"virtualNetworkRules,omitempty" tf:"virtual_network_rules,omitempty"`
}

type NetworkAclsParameters struct {

	// The Default Action to use when no rules match from ip_rules / virtual_network_rules. Possible values are Allow and Deny.
	// +kubebuilder:validation:Optional
	DefaultAction *string `json:"defaultAction" tf:"default_action,omitempty"`

	// One or more IP Addresses, or CIDR Blocks which should be able to access the Cognitive Account.
	// +kubebuilder:validation:Optional
	// +listType=set
	IPRules []*string `json:"ipRules,omitempty" tf:"ip_rules,omitempty"`

	// A virtual_network_rules block as defined below.
	// +kubebuilder:validation:Optional
	VirtualNetworkRules []VirtualNetworkRulesParameters `json:"virtualNetworkRules,omitempty" tf:"virtual_network_rules,omitempty"`
}

type StorageInitParameters struct {

	// The client ID of the managed identity associated with the storage resource.
	IdentityClientID *string `json:"identityClientId,omitempty" tf:"identity_client_id,omitempty"`

	// Full resource id of a Microsoft.Storage resource.
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`
}

type StorageObservation struct {

	// The client ID of the managed identity associated with the storage resource.
	IdentityClientID *string `json:"identityClientId,omitempty" tf:"identity_client_id,omitempty"`

	// Full resource id of a Microsoft.Storage resource.
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`
}

type StorageParameters struct {

	// The client ID of the managed identity associated with the storage resource.
	// +kubebuilder:validation:Optional
	IdentityClientID *string `json:"identityClientId,omitempty" tf:"identity_client_id,omitempty"`

	// Full resource id of a Microsoft.Storage resource.
	// +kubebuilder:validation:Optional
	StorageAccountID *string `json:"storageAccountId" tf:"storage_account_id,omitempty"`
}

type VirtualNetworkRulesInitParameters struct {

	// Whether ignore missing vnet service endpoint or not. Default to false.
	IgnoreMissingVnetServiceEndpoint *bool `json:"ignoreMissingVnetServiceEndpoint,omitempty" tf:"ignore_missing_vnet_service_endpoint,omitempty"`

	// The ID of the subnet which should be able to access this Cognitive Account.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

type VirtualNetworkRulesObservation struct {

	// Whether ignore missing vnet service endpoint or not. Default to false.
	IgnoreMissingVnetServiceEndpoint *bool `json:"ignoreMissingVnetServiceEndpoint,omitempty" tf:"ignore_missing_vnet_service_endpoint,omitempty"`

	// The ID of the subnet which should be able to access this Cognitive Account.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type VirtualNetworkRulesParameters struct {

	// Whether ignore missing vnet service endpoint or not. Default to false.
	// +kubebuilder:validation:Optional
	IgnoreMissingVnetServiceEndpoint *bool `json:"ignoreMissingVnetServiceEndpoint,omitempty" tf:"ignore_missing_vnet_service_endpoint,omitempty"`

	// The ID of the subnet which should be able to access this Cognitive Account.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

// AccountSpec defines the desired state of Account
type AccountSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AccountParameters `json:"forProvider"`
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
	InitProvider AccountInitParameters `json:"initProvider,omitempty"`
}

// AccountStatus defines the observed state of Account.
type AccountStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AccountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Account is the Schema for the Accounts API. Manages a Cognitive Services Account.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Account struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.kind) || (has(self.initProvider) && has(self.initProvider.kind))",message="spec.forProvider.kind is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.skuName) || (has(self.initProvider) && has(self.initProvider.skuName))",message="spec.forProvider.skuName is a required parameter"
	Spec   AccountSpec   `json:"spec"`
	Status AccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccountList contains a list of Accounts
type AccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Account `json:"items"`
}

// Repository type metadata.
var (
	Account_Kind             = "Account"
	Account_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Account_Kind}.String()
	Account_KindAPIVersion   = Account_Kind + "." + CRDGroupVersion.String()
	Account_GroupVersionKind = CRDGroupVersion.WithKind(Account_Kind)
)

func init() {
	SchemeBuilder.Register(&Account{}, &AccountList{})
}
