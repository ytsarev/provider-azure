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

type BackupPolicyVMBackupInitParameters struct {

	// Sets the backup frequency. Possible values are Hourly, Daily and Weekly.
	Frequency *string `json:"frequency,omitempty" tf:"frequency,omitempty"`

	// Duration of the backup window in hours. Possible values are between 4 and 24 This is used when frequency is Hourly.
	HourDuration *float64 `json:"hourDuration,omitempty" tf:"hour_duration,omitempty"`

	// Interval in hour at which backup is triggered. Possible values are 4, 6, 8 and 12. This is used when frequency is Hourly.
	HourInterval *float64 `json:"hourInterval,omitempty" tf:"hour_interval,omitempty"`

	// The time of day to perform the backup in 24hour format.
	Time *string `json:"time,omitempty" tf:"time,omitempty"`

	// The days of the week to perform backups on. Must be one of Sunday, Monday, Tuesday, Wednesday, Thursday, Friday or Saturday. This is used when frequency is Weekly.
	// +listType=set
	Weekdays []*string `json:"weekdays,omitempty" tf:"weekdays,omitempty"`
}

type BackupPolicyVMBackupObservation struct {

	// Sets the backup frequency. Possible values are Hourly, Daily and Weekly.
	Frequency *string `json:"frequency,omitempty" tf:"frequency,omitempty"`

	// Duration of the backup window in hours. Possible values are between 4 and 24 This is used when frequency is Hourly.
	HourDuration *float64 `json:"hourDuration,omitempty" tf:"hour_duration,omitempty"`

	// Interval in hour at which backup is triggered. Possible values are 4, 6, 8 and 12. This is used when frequency is Hourly.
	HourInterval *float64 `json:"hourInterval,omitempty" tf:"hour_interval,omitempty"`

	// The time of day to perform the backup in 24hour format.
	Time *string `json:"time,omitempty" tf:"time,omitempty"`

	// The days of the week to perform backups on. Must be one of Sunday, Monday, Tuesday, Wednesday, Thursday, Friday or Saturday. This is used when frequency is Weekly.
	// +listType=set
	Weekdays []*string `json:"weekdays,omitempty" tf:"weekdays,omitempty"`
}

type BackupPolicyVMBackupParameters struct {

	// Sets the backup frequency. Possible values are Hourly, Daily and Weekly.
	// +kubebuilder:validation:Optional
	Frequency *string `json:"frequency" tf:"frequency,omitempty"`

	// Duration of the backup window in hours. Possible values are between 4 and 24 This is used when frequency is Hourly.
	// +kubebuilder:validation:Optional
	HourDuration *float64 `json:"hourDuration,omitempty" tf:"hour_duration,omitempty"`

	// Interval in hour at which backup is triggered. Possible values are 4, 6, 8 and 12. This is used when frequency is Hourly.
	// +kubebuilder:validation:Optional
	HourInterval *float64 `json:"hourInterval,omitempty" tf:"hour_interval,omitempty"`

	// The time of day to perform the backup in 24hour format.
	// +kubebuilder:validation:Optional
	Time *string `json:"time" tf:"time,omitempty"`

	// The days of the week to perform backups on. Must be one of Sunday, Monday, Tuesday, Wednesday, Thursday, Friday or Saturday. This is used when frequency is Weekly.
	// +kubebuilder:validation:Optional
	// +listType=set
	Weekdays []*string `json:"weekdays,omitempty" tf:"weekdays,omitempty"`
}

type BackupPolicyVMInitParameters struct {

	// Configures the Policy backup frequency, times & days as documented in the backup block below.
	Backup []BackupPolicyVMBackupInitParameters `json:"backup,omitempty" tf:"backup,omitempty"`

	// Specifies the instant restore resource group name as documented in the instant_restore_resource_group block below.
	InstantRestoreResourceGroup []InstantRestoreResourceGroupInitParameters `json:"instantRestoreResourceGroup,omitempty" tf:"instant_restore_resource_group,omitempty"`

	// Specifies the instant restore retention range in days. Possible values are between 1 and 5 when policy_type is V1, and 1 to 30 when policy_type is V2.
	InstantRestoreRetentionDays *float64 `json:"instantRestoreRetentionDays,omitempty" tf:"instant_restore_retention_days,omitempty"`

	// Type of the Backup Policy. Possible values are V1 and V2 where V2 stands for the Enhanced Policy. Defaults to V1. Changing this forces a new resource to be created.
	PolicyType *string `json:"policyType,omitempty" tf:"policy_type,omitempty"`

	// Configures the policy daily retention as documented in the retention_daily block below. Required when backup frequency is Daily.
	RetentionDaily []BackupPolicyVMRetentionDailyInitParameters `json:"retentionDaily,omitempty" tf:"retention_daily,omitempty"`

	// Configures the policy monthly retention as documented in the retention_monthly block below.
	RetentionMonthly []BackupPolicyVMRetentionMonthlyInitParameters `json:"retentionMonthly,omitempty" tf:"retention_monthly,omitempty"`

	// Configures the policy weekly retention as documented in the retention_weekly block below. Required when backup frequency is Weekly.
	RetentionWeekly []BackupPolicyVMRetentionWeeklyInitParameters `json:"retentionWeekly,omitempty" tf:"retention_weekly,omitempty"`

	// Configures the policy yearly retention as documented in the retention_yearly block below.
	RetentionYearly []BackupPolicyVMRetentionYearlyInitParameters `json:"retentionYearly,omitempty" tf:"retention_yearly,omitempty"`

	// Specifies the timezone. the possible values are defined here. Defaults to UTC
	Timezone *string `json:"timezone,omitempty" tf:"timezone,omitempty"`
}

type BackupPolicyVMObservation struct {

	// Configures the Policy backup frequency, times & days as documented in the backup block below.
	Backup []BackupPolicyVMBackupObservation `json:"backup,omitempty" tf:"backup,omitempty"`

	// The ID of the VM Backup Policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the instant restore resource group name as documented in the instant_restore_resource_group block below.
	InstantRestoreResourceGroup []InstantRestoreResourceGroupObservation `json:"instantRestoreResourceGroup,omitempty" tf:"instant_restore_resource_group,omitempty"`

	// Specifies the instant restore retention range in days. Possible values are between 1 and 5 when policy_type is V1, and 1 to 30 when policy_type is V2.
	InstantRestoreRetentionDays *float64 `json:"instantRestoreRetentionDays,omitempty" tf:"instant_restore_retention_days,omitempty"`

	// Type of the Backup Policy. Possible values are V1 and V2 where V2 stands for the Enhanced Policy. Defaults to V1. Changing this forces a new resource to be created.
	PolicyType *string `json:"policyType,omitempty" tf:"policy_type,omitempty"`

	// Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
	RecoveryVaultName *string `json:"recoveryVaultName,omitempty" tf:"recovery_vault_name,omitempty"`

	// The name of the resource group in which to create the policy. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Configures the policy daily retention as documented in the retention_daily block below. Required when backup frequency is Daily.
	RetentionDaily []BackupPolicyVMRetentionDailyObservation `json:"retentionDaily,omitempty" tf:"retention_daily,omitempty"`

	// Configures the policy monthly retention as documented in the retention_monthly block below.
	RetentionMonthly []BackupPolicyVMRetentionMonthlyObservation `json:"retentionMonthly,omitempty" tf:"retention_monthly,omitempty"`

	// Configures the policy weekly retention as documented in the retention_weekly block below. Required when backup frequency is Weekly.
	RetentionWeekly []BackupPolicyVMRetentionWeeklyObservation `json:"retentionWeekly,omitempty" tf:"retention_weekly,omitempty"`

	// Configures the policy yearly retention as documented in the retention_yearly block below.
	RetentionYearly []BackupPolicyVMRetentionYearlyObservation `json:"retentionYearly,omitempty" tf:"retention_yearly,omitempty"`

	// Specifies the timezone. the possible values are defined here. Defaults to UTC
	Timezone *string `json:"timezone,omitempty" tf:"timezone,omitempty"`
}

type BackupPolicyVMParameters struct {

	// Configures the Policy backup frequency, times & days as documented in the backup block below.
	// +kubebuilder:validation:Optional
	Backup []BackupPolicyVMBackupParameters `json:"backup,omitempty" tf:"backup,omitempty"`

	// Specifies the instant restore resource group name as documented in the instant_restore_resource_group block below.
	// +kubebuilder:validation:Optional
	InstantRestoreResourceGroup []InstantRestoreResourceGroupParameters `json:"instantRestoreResourceGroup,omitempty" tf:"instant_restore_resource_group,omitempty"`

	// Specifies the instant restore retention range in days. Possible values are between 1 and 5 when policy_type is V1, and 1 to 30 when policy_type is V2.
	// +kubebuilder:validation:Optional
	InstantRestoreRetentionDays *float64 `json:"instantRestoreRetentionDays,omitempty" tf:"instant_restore_retention_days,omitempty"`

	// Type of the Backup Policy. Possible values are V1 and V2 where V2 stands for the Enhanced Policy. Defaults to V1. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	PolicyType *string `json:"policyType,omitempty" tf:"policy_type,omitempty"`

	// Specifies the name of the Recovery Services Vault to use. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/recoveryservices/v1beta1.Vault
	// +kubebuilder:validation:Optional
	RecoveryVaultName *string `json:"recoveryVaultName,omitempty" tf:"recovery_vault_name,omitempty"`

	// Reference to a Vault in recoveryservices to populate recoveryVaultName.
	// +kubebuilder:validation:Optional
	RecoveryVaultNameRef *v1.Reference `json:"recoveryVaultNameRef,omitempty" tf:"-"`

	// Selector for a Vault in recoveryservices to populate recoveryVaultName.
	// +kubebuilder:validation:Optional
	RecoveryVaultNameSelector *v1.Selector `json:"recoveryVaultNameSelector,omitempty" tf:"-"`

	// The name of the resource group in which to create the policy. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Configures the policy daily retention as documented in the retention_daily block below. Required when backup frequency is Daily.
	// +kubebuilder:validation:Optional
	RetentionDaily []BackupPolicyVMRetentionDailyParameters `json:"retentionDaily,omitempty" tf:"retention_daily,omitempty"`

	// Configures the policy monthly retention as documented in the retention_monthly block below.
	// +kubebuilder:validation:Optional
	RetentionMonthly []BackupPolicyVMRetentionMonthlyParameters `json:"retentionMonthly,omitempty" tf:"retention_monthly,omitempty"`

	// Configures the policy weekly retention as documented in the retention_weekly block below. Required when backup frequency is Weekly.
	// +kubebuilder:validation:Optional
	RetentionWeekly []BackupPolicyVMRetentionWeeklyParameters `json:"retentionWeekly,omitempty" tf:"retention_weekly,omitempty"`

	// Configures the policy yearly retention as documented in the retention_yearly block below.
	// +kubebuilder:validation:Optional
	RetentionYearly []BackupPolicyVMRetentionYearlyParameters `json:"retentionYearly,omitempty" tf:"retention_yearly,omitempty"`

	// Specifies the timezone. the possible values are defined here. Defaults to UTC
	// +kubebuilder:validation:Optional
	Timezone *string `json:"timezone,omitempty" tf:"timezone,omitempty"`
}

type BackupPolicyVMRetentionDailyInitParameters struct {

	// The number of daily backups to keep. Must be between 7 and 9999.
	Count *float64 `json:"count,omitempty" tf:"count,omitempty"`
}

type BackupPolicyVMRetentionDailyObservation struct {

	// The number of daily backups to keep. Must be between 7 and 9999.
	Count *float64 `json:"count,omitempty" tf:"count,omitempty"`
}

type BackupPolicyVMRetentionDailyParameters struct {

	// The number of daily backups to keep. Must be between 7 and 9999.
	// +kubebuilder:validation:Optional
	Count *float64 `json:"count" tf:"count,omitempty"`
}

type BackupPolicyVMRetentionMonthlyInitParameters struct {

	// The number of monthly backups to keep. Must be between 1 and 9999
	Count *float64 `json:"count,omitempty" tf:"count,omitempty"`

	// The weekday backups to retain . Must be one of Sunday, Monday, Tuesday, Wednesday, Thursday, Friday or Saturday.
	// +listType=set
	Weekdays []*string `json:"weekdays,omitempty" tf:"weekdays,omitempty"`

	// The weeks of the month to retain backups of. Must be one of First, Second, Third, Fourth, Last.
	// +listType=set
	Weeks []*string `json:"weeks,omitempty" tf:"weeks,omitempty"`
}

type BackupPolicyVMRetentionMonthlyObservation struct {

	// The number of monthly backups to keep. Must be between 1 and 9999
	Count *float64 `json:"count,omitempty" tf:"count,omitempty"`

	// The weekday backups to retain . Must be one of Sunday, Monday, Tuesday, Wednesday, Thursday, Friday or Saturday.
	// +listType=set
	Weekdays []*string `json:"weekdays,omitempty" tf:"weekdays,omitempty"`

	// The weeks of the month to retain backups of. Must be one of First, Second, Third, Fourth, Last.
	// +listType=set
	Weeks []*string `json:"weeks,omitempty" tf:"weeks,omitempty"`
}

type BackupPolicyVMRetentionMonthlyParameters struct {

	// The number of monthly backups to keep. Must be between 1 and 9999
	// +kubebuilder:validation:Optional
	Count *float64 `json:"count" tf:"count,omitempty"`

	// The weekday backups to retain . Must be one of Sunday, Monday, Tuesday, Wednesday, Thursday, Friday or Saturday.
	// +kubebuilder:validation:Optional
	// +listType=set
	Weekdays []*string `json:"weekdays" tf:"weekdays,omitempty"`

	// The weeks of the month to retain backups of. Must be one of First, Second, Third, Fourth, Last.
	// +kubebuilder:validation:Optional
	// +listType=set
	Weeks []*string `json:"weeks" tf:"weeks,omitempty"`
}

type BackupPolicyVMRetentionWeeklyInitParameters struct {

	// The number of weekly backups to keep. Must be between 1 and 9999
	Count *float64 `json:"count,omitempty" tf:"count,omitempty"`

	// The weekday backups to retain. Must be one of Sunday, Monday, Tuesday, Wednesday, Thursday, Friday or Saturday.
	// +listType=set
	Weekdays []*string `json:"weekdays,omitempty" tf:"weekdays,omitempty"`
}

type BackupPolicyVMRetentionWeeklyObservation struct {

	// The number of weekly backups to keep. Must be between 1 and 9999
	Count *float64 `json:"count,omitempty" tf:"count,omitempty"`

	// The weekday backups to retain. Must be one of Sunday, Monday, Tuesday, Wednesday, Thursday, Friday or Saturday.
	// +listType=set
	Weekdays []*string `json:"weekdays,omitempty" tf:"weekdays,omitempty"`
}

type BackupPolicyVMRetentionWeeklyParameters struct {

	// The number of weekly backups to keep. Must be between 1 and 9999
	// +kubebuilder:validation:Optional
	Count *float64 `json:"count" tf:"count,omitempty"`

	// The weekday backups to retain. Must be one of Sunday, Monday, Tuesday, Wednesday, Thursday, Friday or Saturday.
	// +kubebuilder:validation:Optional
	// +listType=set
	Weekdays []*string `json:"weekdays" tf:"weekdays,omitempty"`
}

type BackupPolicyVMRetentionYearlyInitParameters struct {

	// The number of yearly backups to keep. Must be between 1 and 9999
	Count *float64 `json:"count,omitempty" tf:"count,omitempty"`

	// The months of the year to retain backups of. Must be one of January, February, March, April, May, June, July, August, September, October, November and December.
	// +listType=set
	Months []*string `json:"months,omitempty" tf:"months,omitempty"`

	// The weekday backups to retain . Must be one of Sunday, Monday, Tuesday, Wednesday, Thursday, Friday or Saturday.
	// +listType=set
	Weekdays []*string `json:"weekdays,omitempty" tf:"weekdays,omitempty"`

	// The weeks of the month to retain backups of. Must be one of First, Second, Third, Fourth, Last.
	// +listType=set
	Weeks []*string `json:"weeks,omitempty" tf:"weeks,omitempty"`
}

type BackupPolicyVMRetentionYearlyObservation struct {

	// The number of yearly backups to keep. Must be between 1 and 9999
	Count *float64 `json:"count,omitempty" tf:"count,omitempty"`

	// The months of the year to retain backups of. Must be one of January, February, March, April, May, June, July, August, September, October, November and December.
	// +listType=set
	Months []*string `json:"months,omitempty" tf:"months,omitempty"`

	// The weekday backups to retain . Must be one of Sunday, Monday, Tuesday, Wednesday, Thursday, Friday or Saturday.
	// +listType=set
	Weekdays []*string `json:"weekdays,omitempty" tf:"weekdays,omitempty"`

	// The weeks of the month to retain backups of. Must be one of First, Second, Third, Fourth, Last.
	// +listType=set
	Weeks []*string `json:"weeks,omitempty" tf:"weeks,omitempty"`
}

type BackupPolicyVMRetentionYearlyParameters struct {

	// The number of yearly backups to keep. Must be between 1 and 9999
	// +kubebuilder:validation:Optional
	Count *float64 `json:"count" tf:"count,omitempty"`

	// The months of the year to retain backups of. Must be one of January, February, March, April, May, June, July, August, September, October, November and December.
	// +kubebuilder:validation:Optional
	// +listType=set
	Months []*string `json:"months" tf:"months,omitempty"`

	// The weekday backups to retain . Must be one of Sunday, Monday, Tuesday, Wednesday, Thursday, Friday or Saturday.
	// +kubebuilder:validation:Optional
	// +listType=set
	Weekdays []*string `json:"weekdays" tf:"weekdays,omitempty"`

	// The weeks of the month to retain backups of. Must be one of First, Second, Third, Fourth, Last.
	// +kubebuilder:validation:Optional
	// +listType=set
	Weeks []*string `json:"weeks" tf:"weeks,omitempty"`
}

type InstantRestoreResourceGroupInitParameters struct {

	// The prefix for the instant_restore_resource_group name.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// The suffix for the instant_restore_resource_group name.
	Suffix *string `json:"suffix,omitempty" tf:"suffix,omitempty"`
}

type InstantRestoreResourceGroupObservation struct {

	// The prefix for the instant_restore_resource_group name.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// The suffix for the instant_restore_resource_group name.
	Suffix *string `json:"suffix,omitempty" tf:"suffix,omitempty"`
}

type InstantRestoreResourceGroupParameters struct {

	// The prefix for the instant_restore_resource_group name.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix" tf:"prefix,omitempty"`

	// The suffix for the instant_restore_resource_group name.
	// +kubebuilder:validation:Optional
	Suffix *string `json:"suffix,omitempty" tf:"suffix,omitempty"`
}

// BackupPolicyVMSpec defines the desired state of BackupPolicyVM
type BackupPolicyVMSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BackupPolicyVMParameters `json:"forProvider"`
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
	InitProvider BackupPolicyVMInitParameters `json:"initProvider,omitempty"`
}

// BackupPolicyVMStatus defines the observed state of BackupPolicyVM.
type BackupPolicyVMStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BackupPolicyVMObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BackupPolicyVM is the Schema for the BackupPolicyVMs API. Manages an Azure Backup VM Backup Policy.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type BackupPolicyVM struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.backup) || (has(self.initProvider) && has(self.initProvider.backup))",message="spec.forProvider.backup is a required parameter"
	Spec   BackupPolicyVMSpec   `json:"spec"`
	Status BackupPolicyVMStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupPolicyVMList contains a list of BackupPolicyVMs
type BackupPolicyVMList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupPolicyVM `json:"items"`
}

// Repository type metadata.
var (
	BackupPolicyVM_Kind             = "BackupPolicyVM"
	BackupPolicyVM_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BackupPolicyVM_Kind}.String()
	BackupPolicyVM_KindAPIVersion   = BackupPolicyVM_Kind + "." + CRDGroupVersion.String()
	BackupPolicyVM_GroupVersionKind = CRDGroupVersion.WithKind(BackupPolicyVM_Kind)
)

func init() {
	SchemeBuilder.Register(&BackupPolicyVM{}, &BackupPolicyVMList{})
}
