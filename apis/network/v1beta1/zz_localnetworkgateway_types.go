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

type BGPSettingsInitParameters struct {

	// The BGP speaker's ASN.
	Asn *float64 `json:"asn,omitempty" tf:"asn,omitempty"`

	// The BGP peering address and BGP identifier of this BGP speaker.
	BGPPeeringAddress *string `json:"bgpPeeringAddress,omitempty" tf:"bgp_peering_address,omitempty"`

	// The weight added to routes learned from this BGP speaker.
	PeerWeight *float64 `json:"peerWeight,omitempty" tf:"peer_weight,omitempty"`
}

type BGPSettingsObservation struct {

	// The BGP speaker's ASN.
	Asn *float64 `json:"asn,omitempty" tf:"asn,omitempty"`

	// The BGP peering address and BGP identifier of this BGP speaker.
	BGPPeeringAddress *string `json:"bgpPeeringAddress,omitempty" tf:"bgp_peering_address,omitempty"`

	// The weight added to routes learned from this BGP speaker.
	PeerWeight *float64 `json:"peerWeight,omitempty" tf:"peer_weight,omitempty"`
}

type BGPSettingsParameters struct {

	// The BGP speaker's ASN.
	// +kubebuilder:validation:Optional
	Asn *float64 `json:"asn" tf:"asn,omitempty"`

	// The BGP peering address and BGP identifier of this BGP speaker.
	// +kubebuilder:validation:Optional
	BGPPeeringAddress *string `json:"bgpPeeringAddress" tf:"bgp_peering_address,omitempty"`

	// The weight added to routes learned from this BGP speaker.
	// +kubebuilder:validation:Optional
	PeerWeight *float64 `json:"peerWeight,omitempty" tf:"peer_weight,omitempty"`
}

type LocalNetworkGatewayInitParameters struct {

	// The list of string CIDRs representing the address spaces the gateway exposes.
	AddressSpace []*string `json:"addressSpace,omitempty" tf:"address_space,omitempty"`

	// A bgp_settings block as defined below containing the Local Network Gateway's BGP speaker settings.
	BGPSettings []BGPSettingsInitParameters `json:"bgpSettings,omitempty" tf:"bgp_settings,omitempty"`

	// The gateway IP address to connect with.
	GatewayAddress *string `json:"gatewayAddress,omitempty" tf:"gateway_address,omitempty"`

	// The gateway FQDN to connect with.
	GatewayFqdn *string `json:"gatewayFqdn,omitempty" tf:"gateway_fqdn,omitempty"`

	// The location/region where the local network gateway is created. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type LocalNetworkGatewayObservation struct {

	// The list of string CIDRs representing the address spaces the gateway exposes.
	AddressSpace []*string `json:"addressSpace,omitempty" tf:"address_space,omitempty"`

	// A bgp_settings block as defined below containing the Local Network Gateway's BGP speaker settings.
	BGPSettings []BGPSettingsObservation `json:"bgpSettings,omitempty" tf:"bgp_settings,omitempty"`

	// The gateway IP address to connect with.
	GatewayAddress *string `json:"gatewayAddress,omitempty" tf:"gateway_address,omitempty"`

	// The gateway FQDN to connect with.
	GatewayFqdn *string `json:"gatewayFqdn,omitempty" tf:"gateway_fqdn,omitempty"`

	// The ID of the Local Network Gateway.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The location/region where the local network gateway is created. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the resource group in which to create the local network gateway. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// A mapping of tags to assign to the resource.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type LocalNetworkGatewayParameters struct {

	// The list of string CIDRs representing the address spaces the gateway exposes.
	// +kubebuilder:validation:Optional
	AddressSpace []*string `json:"addressSpace,omitempty" tf:"address_space,omitempty"`

	// A bgp_settings block as defined below containing the Local Network Gateway's BGP speaker settings.
	// +kubebuilder:validation:Optional
	BGPSettings []BGPSettingsParameters `json:"bgpSettings,omitempty" tf:"bgp_settings,omitempty"`

	// The gateway IP address to connect with.
	// +kubebuilder:validation:Optional
	GatewayAddress *string `json:"gatewayAddress,omitempty" tf:"gateway_address,omitempty"`

	// The gateway FQDN to connect with.
	// +kubebuilder:validation:Optional
	GatewayFqdn *string `json:"gatewayFqdn,omitempty" tf:"gateway_fqdn,omitempty"`

	// The location/region where the local network gateway is created. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the resource group in which to create the local network gateway. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// LocalNetworkGatewaySpec defines the desired state of LocalNetworkGateway
type LocalNetworkGatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LocalNetworkGatewayParameters `json:"forProvider"`
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
	InitProvider LocalNetworkGatewayInitParameters `json:"initProvider,omitempty"`
}

// LocalNetworkGatewayStatus defines the observed state of LocalNetworkGateway.
type LocalNetworkGatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LocalNetworkGatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LocalNetworkGateway is the Schema for the LocalNetworkGateways API. Manages a local network gateway connection over which specific connections can be configured.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LocalNetworkGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	Spec   LocalNetworkGatewaySpec   `json:"spec"`
	Status LocalNetworkGatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LocalNetworkGatewayList contains a list of LocalNetworkGateways
type LocalNetworkGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LocalNetworkGateway `json:"items"`
}

// Repository type metadata.
var (
	LocalNetworkGateway_Kind             = "LocalNetworkGateway"
	LocalNetworkGateway_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LocalNetworkGateway_Kind}.String()
	LocalNetworkGateway_KindAPIVersion   = LocalNetworkGateway_Kind + "." + CRDGroupVersion.String()
	LocalNetworkGateway_GroupVersionKind = CRDGroupVersion.WithKind(LocalNetworkGateway_Kind)
)

func init() {
	SchemeBuilder.Register(&LocalNetworkGateway{}, &LocalNetworkGatewayList{})
}
