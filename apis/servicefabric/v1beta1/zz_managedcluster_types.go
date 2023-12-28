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

type ActiveDirectoryInitParameters struct {

	// The ID of the Client Application.
	ClientApplicationID *string `json:"clientApplicationId,omitempty" tf:"client_application_id,omitempty"`

	// The ID of the Cluster Application.
	ClusterApplicationID *string `json:"clusterApplicationId,omitempty" tf:"cluster_application_id,omitempty"`

	// The ID of the Tenant.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type ActiveDirectoryObservation struct {

	// The ID of the Client Application.
	ClientApplicationID *string `json:"clientApplicationId,omitempty" tf:"client_application_id,omitempty"`

	// The ID of the Cluster Application.
	ClusterApplicationID *string `json:"clusterApplicationId,omitempty" tf:"cluster_application_id,omitempty"`

	// The ID of the Tenant.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type ActiveDirectoryParameters struct {

	// The ID of the Client Application.
	// +kubebuilder:validation:Optional
	ClientApplicationID *string `json:"clientApplicationId" tf:"client_application_id,omitempty"`

	// The ID of the Cluster Application.
	// +kubebuilder:validation:Optional
	ClusterApplicationID *string `json:"clusterApplicationId" tf:"cluster_application_id,omitempty"`

	// The ID of the Tenant.
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId" tf:"tenant_id,omitempty"`
}

type AuthenticationCertificateInitParameters struct {

	// The certificate's CN.
	CommonName *string `json:"commonName,omitempty" tf:"common_name,omitempty"`

	// The thumbprint of the certificate.
	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint,omitempty"`

	// The type of the certificate. Can be AdminClient or ReadOnlyClient.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type AuthenticationCertificateObservation struct {

	// The certificate's CN.
	CommonName *string `json:"commonName,omitempty" tf:"common_name,omitempty"`

	// The thumbprint of the certificate.
	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint,omitempty"`

	// The type of the certificate. Can be AdminClient or ReadOnlyClient.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type AuthenticationCertificateParameters struct {

	// The certificate's CN.
	// +kubebuilder:validation:Optional
	CommonName *string `json:"commonName,omitempty" tf:"common_name,omitempty"`

	// The thumbprint of the certificate.
	// +kubebuilder:validation:Optional
	Thumbprint *string `json:"thumbprint" tf:"thumbprint,omitempty"`

	// The type of the certificate. Can be AdminClient or ReadOnlyClient.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type AuthenticationInitParameters struct {

	// A active_directory block as defined above.
	ActiveDirectory []ActiveDirectoryInitParameters `json:"activeDirectory,omitempty" tf:"active_directory,omitempty"`

	// One or more certificate blocks as defined below.
	Certificate []AuthenticationCertificateInitParameters `json:"certificate,omitempty" tf:"certificate,omitempty"`
}

type AuthenticationObservation struct {

	// A active_directory block as defined above.
	ActiveDirectory []ActiveDirectoryObservation `json:"activeDirectory,omitempty" tf:"active_directory,omitempty"`

	// One or more certificate blocks as defined below.
	Certificate []AuthenticationCertificateObservation `json:"certificate,omitempty" tf:"certificate,omitempty"`
}

type AuthenticationParameters struct {

	// A active_directory block as defined above.
	// +kubebuilder:validation:Optional
	ActiveDirectory []ActiveDirectoryParameters `json:"activeDirectory,omitempty" tf:"active_directory,omitempty"`

	// One or more certificate blocks as defined below.
	// +kubebuilder:validation:Optional
	Certificate []AuthenticationCertificateParameters `json:"certificate,omitempty" tf:"certificate,omitempty"`
}

type CertificatesInitParameters struct {

	// The certificate store on the Virtual Machine to which the certificate should be added.
	Store *string `json:"store,omitempty" tf:"store,omitempty"`

	// The URL of a certificate that has been uploaded to Key Vault as a secret
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type CertificatesObservation struct {

	// The certificate store on the Virtual Machine to which the certificate should be added.
	Store *string `json:"store,omitempty" tf:"store,omitempty"`

	// The URL of a certificate that has been uploaded to Key Vault as a secret
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type CertificatesParameters struct {

	// The certificate store on the Virtual Machine to which the certificate should be added.
	// +kubebuilder:validation:Optional
	Store *string `json:"store" tf:"store,omitempty"`

	// The URL of a certificate that has been uploaded to Key Vault as a secret
	// +kubebuilder:validation:Optional
	URL *string `json:"url" tf:"url,omitempty"`
}

type CustomFabricSettingInitParameters struct {

	// Parameter name.
	Parameter *string `json:"parameter,omitempty" tf:"parameter,omitempty"`

	// Section name.
	Section *string `json:"section,omitempty" tf:"section,omitempty"`

	// Parameter value.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type CustomFabricSettingObservation struct {

	// Parameter name.
	Parameter *string `json:"parameter,omitempty" tf:"parameter,omitempty"`

	// Section name.
	Section *string `json:"section,omitempty" tf:"section,omitempty"`

	// Parameter value.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type CustomFabricSettingParameters struct {

	// Parameter name.
	// +kubebuilder:validation:Optional
	Parameter *string `json:"parameter" tf:"parameter,omitempty"`

	// Section name.
	// +kubebuilder:validation:Optional
	Section *string `json:"section" tf:"section,omitempty"`

	// Parameter value.
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type LBRuleInitParameters struct {

	// LB Backend port.
	BackendPort *float64 `json:"backendPort,omitempty" tf:"backend_port,omitempty"`

	// LB Frontend port.
	FrontendPort *float64 `json:"frontendPort,omitempty" tf:"frontend_port,omitempty"`

	// Protocol for the probe. Can be one of tcp, udp, http, or https.
	ProbeProtocol *string `json:"probeProtocol,omitempty" tf:"probe_protocol,omitempty"`

	// Path for the probe to check, when probe protocol is set to http.
	ProbeRequestPath *string `json:"probeRequestPath,omitempty" tf:"probe_request_path,omitempty"`

	// The transport protocol used in this rule. Can be one of tcp or udp.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type LBRuleObservation struct {

	// LB Backend port.
	BackendPort *float64 `json:"backendPort,omitempty" tf:"backend_port,omitempty"`

	// LB Frontend port.
	FrontendPort *float64 `json:"frontendPort,omitempty" tf:"frontend_port,omitempty"`

	// Protocol for the probe. Can be one of tcp, udp, http, or https.
	ProbeProtocol *string `json:"probeProtocol,omitempty" tf:"probe_protocol,omitempty"`

	// Path for the probe to check, when probe protocol is set to http.
	ProbeRequestPath *string `json:"probeRequestPath,omitempty" tf:"probe_request_path,omitempty"`

	// The transport protocol used in this rule. Can be one of tcp or udp.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type LBRuleParameters struct {

	// LB Backend port.
	// +kubebuilder:validation:Optional
	BackendPort *float64 `json:"backendPort" tf:"backend_port,omitempty"`

	// LB Frontend port.
	// +kubebuilder:validation:Optional
	FrontendPort *float64 `json:"frontendPort" tf:"frontend_port,omitempty"`

	// Protocol for the probe. Can be one of tcp, udp, http, or https.
	// +kubebuilder:validation:Optional
	ProbeProtocol *string `json:"probeProtocol" tf:"probe_protocol,omitempty"`

	// Path for the probe to check, when probe protocol is set to http.
	// +kubebuilder:validation:Optional
	ProbeRequestPath *string `json:"probeRequestPath,omitempty" tf:"probe_request_path,omitempty"`

	// The transport protocol used in this rule. Can be one of tcp or udp.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`
}

type ManagedClusterInitParameters struct {

	// Controls how connections to the cluster are authenticated. A authentication block as defined below.
	Authentication []AuthenticationInitParameters `json:"authentication,omitempty" tf:"authentication,omitempty"`

	// If true, backup service is enabled.
	BackupServiceEnabled *bool `json:"backupServiceEnabled,omitempty" tf:"backup_service_enabled,omitempty"`

	// Port to use when connecting to the cluster.
	ClientConnectionPort *float64 `json:"clientConnectionPort,omitempty" tf:"client_connection_port,omitempty"`

	// One or more custom_fabric_setting blocks as defined below.
	CustomFabricSetting []CustomFabricSettingInitParameters `json:"customFabricSetting,omitempty" tf:"custom_fabric_setting,omitempty"`

	// Hostname for the cluster. If unset the cluster's name will be used..
	DNSName *string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`

	// If true, DNS service is enabled.
	DNSServiceEnabled *bool `json:"dnsServiceEnabled,omitempty" tf:"dns_service_enabled,omitempty"`

	// Port that should be used by the Service Fabric Explorer to visualize applications and cluster status.
	HTTPGatewayPort *float64 `json:"httpGatewayPort,omitempty" tf:"http_gateway_port,omitempty"`

	// One or more lb_rule blocks as defined below.
	LBRule []LBRuleInitParameters `json:"lbRule,omitempty" tf:"lb_rule,omitempty"`

	// The Azure Region where the Resource Group should exist. Changing this forces a new Resource Group to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// One or more node_type blocks as defined below.
	NodeType []ManagedClusterNodeTypeInitParameters `json:"nodeType,omitempty" tf:"node_type,omitempty"`

	// SKU for this cluster. Changing this forces a new resource to be created. Default is Basic, allowed values are either Basic or Standard.
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags which should be assigned to the Resource Group.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Upgrade wave for the fabric runtime. Default is Wave0, allowed value must be one of Wave0, Wave1, or Wave2.
	UpgradeWave *string `json:"upgradeWave,omitempty" tf:"upgrade_wave,omitempty"`

	// Administrator password for the VMs that will be created as part of this cluster.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type ManagedClusterNodeTypeInitParameters struct {

	// Sets the port range available for applications. Format is <from_port>-<to_port>, for example 10000-20000.
	ApplicationPortRange *string `json:"applicationPortRange,omitempty" tf:"application_port_range,omitempty"`

	// Specifies a list of key/value pairs used to set capacity tags for this node type.
	// +mapType=granular
	Capacities map[string]*string `json:"capacities,omitempty" tf:"capacities,omitempty"`

	// The size of the data disk in gigabytes..
	DataDiskSizeGb *float64 `json:"dataDiskSizeGb,omitempty" tf:"data_disk_size_gb,omitempty"`

	// The type of the disk to use for storing data. It can be one of Premium_LRS, Standard_LRS, or StandardSSD_LRS.
	DataDiskType *string `json:"dataDiskType,omitempty" tf:"data_disk_type,omitempty"`

	// Sets the port range available for the OS. Format is <from_port>-<to_port>, for example 10000-20000. There has to be at least 255 ports available and cannot overlap with application_port_range..
	EphemeralPortRange *string `json:"ephemeralPortRange,omitempty" tf:"ephemeral_port_range,omitempty"`

	// If set the node type can be composed of multiple placement groups.
	MultiplePlacementGroupsEnabled *bool `json:"multiplePlacementGroupsEnabled,omitempty" tf:"multiple_placement_groups_enabled,omitempty"`

	// The name which should be used for this node type.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies a list of placement tags that can be used to indicate where services should run..
	// +mapType=granular
	PlacementProperties map[string]*string `json:"placementProperties,omitempty" tf:"placement_properties,omitempty"`

	// If set to true, system services will run on this node type. Only one node type should be marked as primary. Primary node type cannot be deleted or changed once they're created.
	Primary *bool `json:"primary,omitempty" tf:"primary,omitempty"`

	// If set to true, only stateless workloads can run on this node type.
	Stateless *bool `json:"stateless,omitempty" tf:"stateless,omitempty"`

	// The offer type of the marketplace image cluster VMs will use.
	VMImageOffer *string `json:"vmImageOffer,omitempty" tf:"vm_image_offer,omitempty"`

	// The publisher of the marketplace image cluster VMs will use.
	VMImagePublisher *string `json:"vmImagePublisher,omitempty" tf:"vm_image_publisher,omitempty"`

	// The SKU of the marketplace image cluster VMs will use.
	VMImageSku *string `json:"vmImageSku,omitempty" tf:"vm_image_sku,omitempty"`

	// The version of the marketplace image cluster VMs will use.
	VMImageVersion *string `json:"vmImageVersion,omitempty" tf:"vm_image_version,omitempty"`

	// The number of instances this node type will launch.
	VMInstanceCount *float64 `json:"vmInstanceCount,omitempty" tf:"vm_instance_count,omitempty"`

	// One or more vm_secrets blocks as defined below.
	VMSecrets []VMSecretsInitParameters `json:"vmSecrets,omitempty" tf:"vm_secrets,omitempty"`

	// The size of the instances in this node type.
	VMSize *string `json:"vmSize,omitempty" tf:"vm_size,omitempty"`
}

type ManagedClusterNodeTypeObservation struct {

	// Sets the port range available for applications. Format is <from_port>-<to_port>, for example 10000-20000.
	ApplicationPortRange *string `json:"applicationPortRange,omitempty" tf:"application_port_range,omitempty"`

	// Specifies a list of key/value pairs used to set capacity tags for this node type.
	// +mapType=granular
	Capacities map[string]*string `json:"capacities,omitempty" tf:"capacities,omitempty"`

	// The size of the data disk in gigabytes..
	DataDiskSizeGb *float64 `json:"dataDiskSizeGb,omitempty" tf:"data_disk_size_gb,omitempty"`

	// The type of the disk to use for storing data. It can be one of Premium_LRS, Standard_LRS, or StandardSSD_LRS.
	DataDiskType *string `json:"dataDiskType,omitempty" tf:"data_disk_type,omitempty"`

	// Sets the port range available for the OS. Format is <from_port>-<to_port>, for example 10000-20000. There has to be at least 255 ports available and cannot overlap with application_port_range..
	EphemeralPortRange *string `json:"ephemeralPortRange,omitempty" tf:"ephemeral_port_range,omitempty"`

	// The ID of the Resource Group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// If set the node type can be composed of multiple placement groups.
	MultiplePlacementGroupsEnabled *bool `json:"multiplePlacementGroupsEnabled,omitempty" tf:"multiple_placement_groups_enabled,omitempty"`

	// The name which should be used for this node type.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies a list of placement tags that can be used to indicate where services should run..
	// +mapType=granular
	PlacementProperties map[string]*string `json:"placementProperties,omitempty" tf:"placement_properties,omitempty"`

	// If set to true, system services will run on this node type. Only one node type should be marked as primary. Primary node type cannot be deleted or changed once they're created.
	Primary *bool `json:"primary,omitempty" tf:"primary,omitempty"`

	// If set to true, only stateless workloads can run on this node type.
	Stateless *bool `json:"stateless,omitempty" tf:"stateless,omitempty"`

	// The offer type of the marketplace image cluster VMs will use.
	VMImageOffer *string `json:"vmImageOffer,omitempty" tf:"vm_image_offer,omitempty"`

	// The publisher of the marketplace image cluster VMs will use.
	VMImagePublisher *string `json:"vmImagePublisher,omitempty" tf:"vm_image_publisher,omitempty"`

	// The SKU of the marketplace image cluster VMs will use.
	VMImageSku *string `json:"vmImageSku,omitempty" tf:"vm_image_sku,omitempty"`

	// The version of the marketplace image cluster VMs will use.
	VMImageVersion *string `json:"vmImageVersion,omitempty" tf:"vm_image_version,omitempty"`

	// The number of instances this node type will launch.
	VMInstanceCount *float64 `json:"vmInstanceCount,omitempty" tf:"vm_instance_count,omitempty"`

	// One or more vm_secrets blocks as defined below.
	VMSecrets []VMSecretsObservation `json:"vmSecrets,omitempty" tf:"vm_secrets,omitempty"`

	// The size of the instances in this node type.
	VMSize *string `json:"vmSize,omitempty" tf:"vm_size,omitempty"`
}

type ManagedClusterNodeTypeParameters struct {

	// Sets the port range available for applications. Format is <from_port>-<to_port>, for example 10000-20000.
	// +kubebuilder:validation:Optional
	ApplicationPortRange *string `json:"applicationPortRange" tf:"application_port_range,omitempty"`

	// Specifies a list of key/value pairs used to set capacity tags for this node type.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Capacities map[string]*string `json:"capacities,omitempty" tf:"capacities,omitempty"`

	// The size of the data disk in gigabytes..
	// +kubebuilder:validation:Optional
	DataDiskSizeGb *float64 `json:"dataDiskSizeGb" tf:"data_disk_size_gb,omitempty"`

	// The type of the disk to use for storing data. It can be one of Premium_LRS, Standard_LRS, or StandardSSD_LRS.
	// +kubebuilder:validation:Optional
	DataDiskType *string `json:"dataDiskType,omitempty" tf:"data_disk_type,omitempty"`

	// Sets the port range available for the OS. Format is <from_port>-<to_port>, for example 10000-20000. There has to be at least 255 ports available and cannot overlap with application_port_range..
	// +kubebuilder:validation:Optional
	EphemeralPortRange *string `json:"ephemeralPortRange" tf:"ephemeral_port_range,omitempty"`

	// If set the node type can be composed of multiple placement groups.
	// +kubebuilder:validation:Optional
	MultiplePlacementGroupsEnabled *bool `json:"multiplePlacementGroupsEnabled,omitempty" tf:"multiple_placement_groups_enabled,omitempty"`

	// The name which should be used for this node type.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies a list of placement tags that can be used to indicate where services should run..
	// +kubebuilder:validation:Optional
	// +mapType=granular
	PlacementProperties map[string]*string `json:"placementProperties,omitempty" tf:"placement_properties,omitempty"`

	// If set to true, system services will run on this node type. Only one node type should be marked as primary. Primary node type cannot be deleted or changed once they're created.
	// +kubebuilder:validation:Optional
	Primary *bool `json:"primary,omitempty" tf:"primary,omitempty"`

	// If set to true, only stateless workloads can run on this node type.
	// +kubebuilder:validation:Optional
	Stateless *bool `json:"stateless,omitempty" tf:"stateless,omitempty"`

	// The offer type of the marketplace image cluster VMs will use.
	// +kubebuilder:validation:Optional
	VMImageOffer *string `json:"vmImageOffer" tf:"vm_image_offer,omitempty"`

	// The publisher of the marketplace image cluster VMs will use.
	// +kubebuilder:validation:Optional
	VMImagePublisher *string `json:"vmImagePublisher" tf:"vm_image_publisher,omitempty"`

	// The SKU of the marketplace image cluster VMs will use.
	// +kubebuilder:validation:Optional
	VMImageSku *string `json:"vmImageSku" tf:"vm_image_sku,omitempty"`

	// The version of the marketplace image cluster VMs will use.
	// +kubebuilder:validation:Optional
	VMImageVersion *string `json:"vmImageVersion" tf:"vm_image_version,omitempty"`

	// The number of instances this node type will launch.
	// +kubebuilder:validation:Optional
	VMInstanceCount *float64 `json:"vmInstanceCount" tf:"vm_instance_count,omitempty"`

	// One or more vm_secrets blocks as defined below.
	// +kubebuilder:validation:Optional
	VMSecrets []VMSecretsParameters `json:"vmSecrets,omitempty" tf:"vm_secrets,omitempty"`

	// The size of the instances in this node type.
	// +kubebuilder:validation:Optional
	VMSize *string `json:"vmSize" tf:"vm_size,omitempty"`
}

type ManagedClusterObservation struct {

	// Controls how connections to the cluster are authenticated. A authentication block as defined below.
	Authentication []AuthenticationObservation `json:"authentication,omitempty" tf:"authentication,omitempty"`

	// If true, backup service is enabled.
	BackupServiceEnabled *bool `json:"backupServiceEnabled,omitempty" tf:"backup_service_enabled,omitempty"`

	// Port to use when connecting to the cluster.
	ClientConnectionPort *float64 `json:"clientConnectionPort,omitempty" tf:"client_connection_port,omitempty"`

	// One or more custom_fabric_setting blocks as defined below.
	CustomFabricSetting []CustomFabricSettingObservation `json:"customFabricSetting,omitempty" tf:"custom_fabric_setting,omitempty"`

	// Hostname for the cluster. If unset the cluster's name will be used..
	DNSName *string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`

	// If true, DNS service is enabled.
	DNSServiceEnabled *bool `json:"dnsServiceEnabled,omitempty" tf:"dns_service_enabled,omitempty"`

	// Port that should be used by the Service Fabric Explorer to visualize applications and cluster status.
	HTTPGatewayPort *float64 `json:"httpGatewayPort,omitempty" tf:"http_gateway_port,omitempty"`

	// The ID of the Resource Group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// One or more lb_rule blocks as defined below.
	LBRule []LBRuleObservation `json:"lbRule,omitempty" tf:"lb_rule,omitempty"`

	// The Azure Region where the Resource Group should exist. Changing this forces a new Resource Group to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// One or more node_type blocks as defined below.
	NodeType []ManagedClusterNodeTypeObservation `json:"nodeType,omitempty" tf:"node_type,omitempty"`

	// The name of the Resource Group where the Resource Group should exist. Changing this forces a new Resource Group to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// SKU for this cluster. Changing this forces a new resource to be created. Default is Basic, allowed values are either Basic or Standard.
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags which should be assigned to the Resource Group.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Upgrade wave for the fabric runtime. Default is Wave0, allowed value must be one of Wave0, Wave1, or Wave2.
	UpgradeWave *string `json:"upgradeWave,omitempty" tf:"upgrade_wave,omitempty"`

	// Administrator password for the VMs that will be created as part of this cluster.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type ManagedClusterParameters struct {

	// Controls how connections to the cluster are authenticated. A authentication block as defined below.
	// +kubebuilder:validation:Optional
	Authentication []AuthenticationParameters `json:"authentication,omitempty" tf:"authentication,omitempty"`

	// If true, backup service is enabled.
	// +kubebuilder:validation:Optional
	BackupServiceEnabled *bool `json:"backupServiceEnabled,omitempty" tf:"backup_service_enabled,omitempty"`

	// Port to use when connecting to the cluster.
	// +kubebuilder:validation:Optional
	ClientConnectionPort *float64 `json:"clientConnectionPort,omitempty" tf:"client_connection_port,omitempty"`

	// One or more custom_fabric_setting blocks as defined below.
	// +kubebuilder:validation:Optional
	CustomFabricSetting []CustomFabricSettingParameters `json:"customFabricSetting,omitempty" tf:"custom_fabric_setting,omitempty"`

	// Hostname for the cluster. If unset the cluster's name will be used..
	// +kubebuilder:validation:Optional
	DNSName *string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`

	// If true, DNS service is enabled.
	// +kubebuilder:validation:Optional
	DNSServiceEnabled *bool `json:"dnsServiceEnabled,omitempty" tf:"dns_service_enabled,omitempty"`

	// Port that should be used by the Service Fabric Explorer to visualize applications and cluster status.
	// +kubebuilder:validation:Optional
	HTTPGatewayPort *float64 `json:"httpGatewayPort,omitempty" tf:"http_gateway_port,omitempty"`

	// One or more lb_rule blocks as defined below.
	// +kubebuilder:validation:Optional
	LBRule []LBRuleParameters `json:"lbRule,omitempty" tf:"lb_rule,omitempty"`

	// The Azure Region where the Resource Group should exist. Changing this forces a new Resource Group to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// One or more node_type blocks as defined below.
	// +kubebuilder:validation:Optional
	NodeType []ManagedClusterNodeTypeParameters `json:"nodeType,omitempty" tf:"node_type,omitempty"`

	// Administrator password for the VMs that will be created as part of this cluster.
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// The name of the Resource Group where the Resource Group should exist. Changing this forces a new Resource Group to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// SKU for this cluster. Changing this forces a new resource to be created. Default is Basic, allowed values are either Basic or Standard.
	// +kubebuilder:validation:Optional
	Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags which should be assigned to the Resource Group.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Upgrade wave for the fabric runtime. Default is Wave0, allowed value must be one of Wave0, Wave1, or Wave2.
	// +kubebuilder:validation:Optional
	UpgradeWave *string `json:"upgradeWave,omitempty" tf:"upgrade_wave,omitempty"`

	// Administrator password for the VMs that will be created as part of this cluster.
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type VMSecretsInitParameters struct {

	// One or more certificates blocks as defined above.
	Certificates []CertificatesInitParameters `json:"certificates,omitempty" tf:"certificates,omitempty"`

	// The ID of the Vault that contain the certificates.
	VaultID *string `json:"vaultId,omitempty" tf:"vault_id,omitempty"`
}

type VMSecretsObservation struct {

	// One or more certificates blocks as defined above.
	Certificates []CertificatesObservation `json:"certificates,omitempty" tf:"certificates,omitempty"`

	// The ID of the Vault that contain the certificates.
	VaultID *string `json:"vaultId,omitempty" tf:"vault_id,omitempty"`
}

type VMSecretsParameters struct {

	// One or more certificates blocks as defined above.
	// +kubebuilder:validation:Optional
	Certificates []CertificatesParameters `json:"certificates" tf:"certificates,omitempty"`

	// The ID of the Vault that contain the certificates.
	// +kubebuilder:validation:Optional
	VaultID *string `json:"vaultId" tf:"vault_id,omitempty"`
}

// ManagedClusterSpec defines the desired state of ManagedCluster
type ManagedClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ManagedClusterParameters `json:"forProvider"`
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
	InitProvider ManagedClusterInitParameters `json:"initProvider,omitempty"`
}

// ManagedClusterStatus defines the observed state of ManagedCluster.
type ManagedClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ManagedClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ManagedCluster is the Schema for the ManagedClusters API. Manages a Resource Group.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ManagedCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.clientConnectionPort) || (has(self.initProvider) && has(self.initProvider.clientConnectionPort))",message="spec.forProvider.clientConnectionPort is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.httpGatewayPort) || (has(self.initProvider) && has(self.initProvider.httpGatewayPort))",message="spec.forProvider.httpGatewayPort is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.lbRule) || (has(self.initProvider) && has(self.initProvider.lbRule))",message="spec.forProvider.lbRule is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	Spec   ManagedClusterSpec   `json:"spec"`
	Status ManagedClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ManagedClusterList contains a list of ManagedClusters
type ManagedClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ManagedCluster `json:"items"`
}

// Repository type metadata.
var (
	ManagedCluster_Kind             = "ManagedCluster"
	ManagedCluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ManagedCluster_Kind}.String()
	ManagedCluster_KindAPIVersion   = ManagedCluster_Kind + "." + CRDGroupVersion.String()
	ManagedCluster_GroupVersionKind = CRDGroupVersion.WithKind(ManagedCluster_Kind)
)

func init() {
	SchemeBuilder.Register(&ManagedCluster{}, &ManagedClusterList{})
}
