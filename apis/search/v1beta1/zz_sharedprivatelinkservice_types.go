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

type SharedPrivateLinkServiceInitParameters struct {

	// Specify the request message for requesting approval of the Shared Private Link Enabled Remote Resource.
	RequestMessage *string `json:"requestMessage,omitempty" tf:"request_message,omitempty"`

	// Specify the sub resource name which the Azure Search Private Endpoint is able to connect to. Changing this forces a new resource to be created.
	SubresourceName *string `json:"subresourceName,omitempty" tf:"subresource_name,omitempty"`

	// Specify the ID of the Shared Private Link Enabled Remote Resource which this Azure Search Private Endpoint should be connected to. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Account
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	TargetResourceID *string `json:"targetResourceId,omitempty" tf:"target_resource_id,omitempty"`

	// Reference to a Account in storage to populate targetResourceId.
	// +kubebuilder:validation:Optional
	TargetResourceIDRef *v1.Reference `json:"targetResourceIdRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate targetResourceId.
	// +kubebuilder:validation:Optional
	TargetResourceIDSelector *v1.Selector `json:"targetResourceIdSelector,omitempty" tf:"-"`
}

type SharedPrivateLinkServiceObservation struct {

	// The ID of the Azure Search Shared Private Link resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specify the request message for requesting approval of the Shared Private Link Enabled Remote Resource.
	RequestMessage *string `json:"requestMessage,omitempty" tf:"request_message,omitempty"`

	// Specify the id of the Azure Search Service. Changing this forces a new resource to be created.
	SearchServiceID *string `json:"searchServiceId,omitempty" tf:"search_service_id,omitempty"`

	// The status of a private endpoint connection. Possible values are Pending, Approved, Rejected or Disconnected.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Specify the sub resource name which the Azure Search Private Endpoint is able to connect to. Changing this forces a new resource to be created.
	SubresourceName *string `json:"subresourceName,omitempty" tf:"subresource_name,omitempty"`

	// Specify the ID of the Shared Private Link Enabled Remote Resource which this Azure Search Private Endpoint should be connected to. Changing this forces a new resource to be created.
	TargetResourceID *string `json:"targetResourceId,omitempty" tf:"target_resource_id,omitempty"`
}

type SharedPrivateLinkServiceParameters struct {

	// Specify the request message for requesting approval of the Shared Private Link Enabled Remote Resource.
	// +kubebuilder:validation:Optional
	RequestMessage *string `json:"requestMessage,omitempty" tf:"request_message,omitempty"`

	// Specify the id of the Azure Search Service. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/search/v1beta1.Service
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SearchServiceID *string `json:"searchServiceId,omitempty" tf:"search_service_id,omitempty"`

	// Reference to a Service in search to populate searchServiceId.
	// +kubebuilder:validation:Optional
	SearchServiceIDRef *v1.Reference `json:"searchServiceIdRef,omitempty" tf:"-"`

	// Selector for a Service in search to populate searchServiceId.
	// +kubebuilder:validation:Optional
	SearchServiceIDSelector *v1.Selector `json:"searchServiceIdSelector,omitempty" tf:"-"`

	// Specify the sub resource name which the Azure Search Private Endpoint is able to connect to. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SubresourceName *string `json:"subresourceName,omitempty" tf:"subresource_name,omitempty"`

	// Specify the ID of the Shared Private Link Enabled Remote Resource which this Azure Search Private Endpoint should be connected to. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Account
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	TargetResourceID *string `json:"targetResourceId,omitempty" tf:"target_resource_id,omitempty"`

	// Reference to a Account in storage to populate targetResourceId.
	// +kubebuilder:validation:Optional
	TargetResourceIDRef *v1.Reference `json:"targetResourceIdRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate targetResourceId.
	// +kubebuilder:validation:Optional
	TargetResourceIDSelector *v1.Selector `json:"targetResourceIdSelector,omitempty" tf:"-"`
}

// SharedPrivateLinkServiceSpec defines the desired state of SharedPrivateLinkService
type SharedPrivateLinkServiceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SharedPrivateLinkServiceParameters `json:"forProvider"`
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
	InitProvider SharedPrivateLinkServiceInitParameters `json:"initProvider,omitempty"`
}

// SharedPrivateLinkServiceStatus defines the observed state of SharedPrivateLinkService.
type SharedPrivateLinkServiceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SharedPrivateLinkServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SharedPrivateLinkService is the Schema for the SharedPrivateLinkServices API. Manages the Shared Private Link Service for an Azure Search Service.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SharedPrivateLinkService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.subresourceName) || (has(self.initProvider) && has(self.initProvider.subresourceName))",message="spec.forProvider.subresourceName is a required parameter"
	Spec   SharedPrivateLinkServiceSpec   `json:"spec"`
	Status SharedPrivateLinkServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SharedPrivateLinkServiceList contains a list of SharedPrivateLinkServices
type SharedPrivateLinkServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SharedPrivateLinkService `json:"items"`
}

// Repository type metadata.
var (
	SharedPrivateLinkService_Kind             = "SharedPrivateLinkService"
	SharedPrivateLinkService_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SharedPrivateLinkService_Kind}.String()
	SharedPrivateLinkService_KindAPIVersion   = SharedPrivateLinkService_Kind + "." + CRDGroupVersion.String()
	SharedPrivateLinkService_GroupVersionKind = CRDGroupVersion.WithKind(SharedPrivateLinkService_Kind)
)

func init() {
	SchemeBuilder.Register(&SharedPrivateLinkService{}, &SharedPrivateLinkServiceList{})
}
