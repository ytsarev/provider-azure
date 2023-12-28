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

type NetworkACLInitParameters struct {

	// The default action to control the network access when no other rule matches. Possible values are Allow and Deny.
	DefaultAction *string `json:"defaultAction,omitempty" tf:"default_action,omitempty"`

	// A private_endpoint block as defined below.
	PrivateEndpoint []PrivateEndpointInitParameters `json:"privateEndpoint,omitempty" tf:"private_endpoint,omitempty"`

	// A public_network block as defined below.
	PublicNetwork []PublicNetworkInitParameters `json:"publicNetwork,omitempty" tf:"public_network,omitempty"`

	// The ID of the SignalR service. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/signalrservice/v1beta1.Service
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	SignalrServiceID *string `json:"signalrServiceId,omitempty" tf:"signalr_service_id,omitempty"`

	// Reference to a Service in signalrservice to populate signalrServiceId.
	// +kubebuilder:validation:Optional
	SignalrServiceIDRef *v1.Reference `json:"signalrServiceIdRef,omitempty" tf:"-"`

	// Selector for a Service in signalrservice to populate signalrServiceId.
	// +kubebuilder:validation:Optional
	SignalrServiceIDSelector *v1.Selector `json:"signalrServiceIdSelector,omitempty" tf:"-"`
}

type NetworkACLObservation struct {

	// The default action to control the network access when no other rule matches. Possible values are Allow and Deny.
	DefaultAction *string `json:"defaultAction,omitempty" tf:"default_action,omitempty"`

	// The ID of the SignalR service.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A private_endpoint block as defined below.
	PrivateEndpoint []PrivateEndpointObservation `json:"privateEndpoint,omitempty" tf:"private_endpoint,omitempty"`

	// A public_network block as defined below.
	PublicNetwork []PublicNetworkObservation `json:"publicNetwork,omitempty" tf:"public_network,omitempty"`

	// The ID of the SignalR service. Changing this forces a new resource to be created.
	SignalrServiceID *string `json:"signalrServiceId,omitempty" tf:"signalr_service_id,omitempty"`
}

type NetworkACLParameters struct {

	// The default action to control the network access when no other rule matches. Possible values are Allow and Deny.
	// +kubebuilder:validation:Optional
	DefaultAction *string `json:"defaultAction,omitempty" tf:"default_action,omitempty"`

	// A private_endpoint block as defined below.
	// +kubebuilder:validation:Optional
	PrivateEndpoint []PrivateEndpointParameters `json:"privateEndpoint,omitempty" tf:"private_endpoint,omitempty"`

	// A public_network block as defined below.
	// +kubebuilder:validation:Optional
	PublicNetwork []PublicNetworkParameters `json:"publicNetwork,omitempty" tf:"public_network,omitempty"`

	// The ID of the SignalR service. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/signalrservice/v1beta1.Service
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SignalrServiceID *string `json:"signalrServiceId,omitempty" tf:"signalr_service_id,omitempty"`

	// Reference to a Service in signalrservice to populate signalrServiceId.
	// +kubebuilder:validation:Optional
	SignalrServiceIDRef *v1.Reference `json:"signalrServiceIdRef,omitempty" tf:"-"`

	// Selector for a Service in signalrservice to populate signalrServiceId.
	// +kubebuilder:validation:Optional
	SignalrServiceIDSelector *v1.Selector `json:"signalrServiceIdSelector,omitempty" tf:"-"`
}

type PrivateEndpointInitParameters struct {

	// The allowed request types for the Private Endpoint Connection. Possible values are ClientConnection, ServerConnection, RESTAPI and Trace.
	// +listType=set
	AllowedRequestTypes []*string `json:"allowedRequestTypes,omitempty" tf:"allowed_request_types,omitempty"`

	// The denied request types for the Private Endpoint Connection. Possible values are ClientConnection, ServerConnection, RESTAPI and Trace.
	// +listType=set
	DeniedRequestTypes []*string `json:"deniedRequestTypes,omitempty" tf:"denied_request_types,omitempty"`

	// The ID of the Private Endpoint which is based on the SignalR service.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.PrivateEndpoint
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Reference to a PrivateEndpoint in network to populate id.
	// +kubebuilder:validation:Optional
	IDRef *v1.Reference `json:"idRef,omitempty" tf:"-"`

	// Selector for a PrivateEndpoint in network to populate id.
	// +kubebuilder:validation:Optional
	IDSelector *v1.Selector `json:"idSelector,omitempty" tf:"-"`
}

type PrivateEndpointObservation struct {

	// The allowed request types for the Private Endpoint Connection. Possible values are ClientConnection, ServerConnection, RESTAPI and Trace.
	// +listType=set
	AllowedRequestTypes []*string `json:"allowedRequestTypes,omitempty" tf:"allowed_request_types,omitempty"`

	// The denied request types for the Private Endpoint Connection. Possible values are ClientConnection, ServerConnection, RESTAPI and Trace.
	// +listType=set
	DeniedRequestTypes []*string `json:"deniedRequestTypes,omitempty" tf:"denied_request_types,omitempty"`

	// The ID of the Private Endpoint which is based on the SignalR service.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PrivateEndpointParameters struct {

	// The allowed request types for the Private Endpoint Connection. Possible values are ClientConnection, ServerConnection, RESTAPI and Trace.
	// +kubebuilder:validation:Optional
	// +listType=set
	AllowedRequestTypes []*string `json:"allowedRequestTypes,omitempty" tf:"allowed_request_types,omitempty"`

	// The denied request types for the Private Endpoint Connection. Possible values are ClientConnection, ServerConnection, RESTAPI and Trace.
	// +kubebuilder:validation:Optional
	// +listType=set
	DeniedRequestTypes []*string `json:"deniedRequestTypes,omitempty" tf:"denied_request_types,omitempty"`

	// The ID of the Private Endpoint which is based on the SignalR service.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.PrivateEndpoint
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Reference to a PrivateEndpoint in network to populate id.
	// +kubebuilder:validation:Optional
	IDRef *v1.Reference `json:"idRef,omitempty" tf:"-"`

	// Selector for a PrivateEndpoint in network to populate id.
	// +kubebuilder:validation:Optional
	IDSelector *v1.Selector `json:"idSelector,omitempty" tf:"-"`
}

type PublicNetworkInitParameters struct {

	// The allowed request types for the public network. Possible values are ClientConnection, ServerConnection, RESTAPI and Trace.
	// +listType=set
	AllowedRequestTypes []*string `json:"allowedRequestTypes,omitempty" tf:"allowed_request_types,omitempty"`

	// The denied request types for the public network. Possible values are ClientConnection, ServerConnection, RESTAPI and Trace.
	// +listType=set
	DeniedRequestTypes []*string `json:"deniedRequestTypes,omitempty" tf:"denied_request_types,omitempty"`
}

type PublicNetworkObservation struct {

	// The allowed request types for the public network. Possible values are ClientConnection, ServerConnection, RESTAPI and Trace.
	// +listType=set
	AllowedRequestTypes []*string `json:"allowedRequestTypes,omitempty" tf:"allowed_request_types,omitempty"`

	// The denied request types for the public network. Possible values are ClientConnection, ServerConnection, RESTAPI and Trace.
	// +listType=set
	DeniedRequestTypes []*string `json:"deniedRequestTypes,omitempty" tf:"denied_request_types,omitempty"`
}

type PublicNetworkParameters struct {

	// The allowed request types for the public network. Possible values are ClientConnection, ServerConnection, RESTAPI and Trace.
	// +kubebuilder:validation:Optional
	// +listType=set
	AllowedRequestTypes []*string `json:"allowedRequestTypes,omitempty" tf:"allowed_request_types,omitempty"`

	// The denied request types for the public network. Possible values are ClientConnection, ServerConnection, RESTAPI and Trace.
	// +kubebuilder:validation:Optional
	// +listType=set
	DeniedRequestTypes []*string `json:"deniedRequestTypes,omitempty" tf:"denied_request_types,omitempty"`
}

// NetworkACLSpec defines the desired state of NetworkACL
type NetworkACLSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetworkACLParameters `json:"forProvider"`
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
	InitProvider NetworkACLInitParameters `json:"initProvider,omitempty"`
}

// NetworkACLStatus defines the observed state of NetworkACL.
type NetworkACLStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetworkACLObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkACL is the Schema for the NetworkACLs API. Manages the Network ACL for a SignalR service.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type NetworkACL struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.defaultAction) || (has(self.initProvider) && has(self.initProvider.defaultAction))",message="spec.forProvider.defaultAction is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.publicNetwork) || (has(self.initProvider) && has(self.initProvider.publicNetwork))",message="spec.forProvider.publicNetwork is a required parameter"
	Spec   NetworkACLSpec   `json:"spec"`
	Status NetworkACLStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkACLList contains a list of NetworkACLs
type NetworkACLList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkACL `json:"items"`
}

// Repository type metadata.
var (
	NetworkACL_Kind             = "NetworkACL"
	NetworkACL_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NetworkACL_Kind}.String()
	NetworkACL_KindAPIVersion   = NetworkACL_Kind + "." + CRDGroupVersion.String()
	NetworkACL_GroupVersionKind = CRDGroupVersion.WithKind(NetworkACL_Kind)
)

func init() {
	SchemeBuilder.Register(&NetworkACL{}, &NetworkACLList{})
}
