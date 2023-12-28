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

type BackupContainerStorageAccountInitParameters struct {

	// The name of the vault where the storage account will be registered. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/recoveryservices/v1beta1.Vault
	RecoveryVaultName *string `json:"recoveryVaultName,omitempty" tf:"recovery_vault_name,omitempty"`

	// Reference to a Vault in recoveryservices to populate recoveryVaultName.
	// +kubebuilder:validation:Optional
	RecoveryVaultNameRef *v1.Reference `json:"recoveryVaultNameRef,omitempty" tf:"-"`

	// Selector for a Vault in recoveryservices to populate recoveryVaultName.
	// +kubebuilder:validation:Optional
	RecoveryVaultNameSelector *v1.Selector `json:"recoveryVaultNameSelector,omitempty" tf:"-"`

	// Name of the resource group where the vault is located. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The ID of the Storage Account to be registered Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Account
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`

	// Reference to a Account in storage to populate storageAccountId.
	// +kubebuilder:validation:Optional
	StorageAccountIDRef *v1.Reference `json:"storageAccountIdRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate storageAccountId.
	// +kubebuilder:validation:Optional
	StorageAccountIDSelector *v1.Selector `json:"storageAccountIdSelector,omitempty" tf:"-"`
}

type BackupContainerStorageAccountObservation struct {

	// The ID of the Backup Storage Account Container.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the vault where the storage account will be registered. Changing this forces a new resource to be created.
	RecoveryVaultName *string `json:"recoveryVaultName,omitempty" tf:"recovery_vault_name,omitempty"`

	// Name of the resource group where the vault is located. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The ID of the Storage Account to be registered Changing this forces a new resource to be created.
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`
}

type BackupContainerStorageAccountParameters struct {

	// The name of the vault where the storage account will be registered. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/recoveryservices/v1beta1.Vault
	// +kubebuilder:validation:Optional
	RecoveryVaultName *string `json:"recoveryVaultName,omitempty" tf:"recovery_vault_name,omitempty"`

	// Reference to a Vault in recoveryservices to populate recoveryVaultName.
	// +kubebuilder:validation:Optional
	RecoveryVaultNameRef *v1.Reference `json:"recoveryVaultNameRef,omitempty" tf:"-"`

	// Selector for a Vault in recoveryservices to populate recoveryVaultName.
	// +kubebuilder:validation:Optional
	RecoveryVaultNameSelector *v1.Selector `json:"recoveryVaultNameSelector,omitempty" tf:"-"`

	// Name of the resource group where the vault is located. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The ID of the Storage Account to be registered Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Account
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	StorageAccountID *string `json:"storageAccountId,omitempty" tf:"storage_account_id,omitempty"`

	// Reference to a Account in storage to populate storageAccountId.
	// +kubebuilder:validation:Optional
	StorageAccountIDRef *v1.Reference `json:"storageAccountIdRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate storageAccountId.
	// +kubebuilder:validation:Optional
	StorageAccountIDSelector *v1.Selector `json:"storageAccountIdSelector,omitempty" tf:"-"`
}

// BackupContainerStorageAccountSpec defines the desired state of BackupContainerStorageAccount
type BackupContainerStorageAccountSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BackupContainerStorageAccountParameters `json:"forProvider"`
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
	InitProvider BackupContainerStorageAccountInitParameters `json:"initProvider,omitempty"`
}

// BackupContainerStorageAccountStatus defines the observed state of BackupContainerStorageAccount.
type BackupContainerStorageAccountStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BackupContainerStorageAccountObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BackupContainerStorageAccount is the Schema for the BackupContainerStorageAccounts API. Manages a storage account container in an Azure Recovery Vault
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type BackupContainerStorageAccount struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BackupContainerStorageAccountSpec   `json:"spec"`
	Status            BackupContainerStorageAccountStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupContainerStorageAccountList contains a list of BackupContainerStorageAccounts
type BackupContainerStorageAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupContainerStorageAccount `json:"items"`
}

// Repository type metadata.
var (
	BackupContainerStorageAccount_Kind             = "BackupContainerStorageAccount"
	BackupContainerStorageAccount_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BackupContainerStorageAccount_Kind}.String()
	BackupContainerStorageAccount_KindAPIVersion   = BackupContainerStorageAccount_Kind + "." + CRDGroupVersion.String()
	BackupContainerStorageAccount_GroupVersionKind = CRDGroupVersion.WithKind(BackupContainerStorageAccount_Kind)
)

func init() {
	SchemeBuilder.Register(&BackupContainerStorageAccount{}, &BackupContainerStorageAccountList{})
}
