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

type DailyRecurrenceObservation struct {
}

type DailyRecurrenceParameters struct {

	// The time each day when the schedule takes effect.
	// +kubebuilder:validation:Required
	Time *string `json:"time" tf:"time,omitempty"`
}

type HourlyRecurrenceObservation struct {
}

type HourlyRecurrenceParameters struct {

	// Minutes of the hour the schedule will run.
	// +kubebuilder:validation:Required
	Minute *float64 `json:"minute" tf:"minute,omitempty"`
}

type ScheduleNotificationSettingsObservation struct {
}

type ScheduleNotificationSettingsParameters struct {

	// The status of the notification. Possible values are Enabled and Disabled. Defaults to Disabled
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Time in minutes before event at which notification will be sent.
	// +kubebuilder:validation:Optional
	TimeInMinutes *float64 `json:"timeInMinutes,omitempty" tf:"time_in_minutes,omitempty"`

	// The webhook URL to which the notification will be sent.
	// +kubebuilder:validation:Optional
	WebhookURL *string `json:"webhookUrl,omitempty" tf:"webhook_url,omitempty"`
}

type ScheduleObservation struct {

	// The ID of the DevTest Schedule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ScheduleParameters struct {

	// The properties of a daily schedule. If the schedule occurs once each day of the week, specify the daily recurrence. A daily_recurrence block as defined below.
	// +kubebuilder:validation:Optional
	DailyRecurrence []DailyRecurrenceParameters `json:"dailyRecurrence,omitempty" tf:"daily_recurrence,omitempty"`

	// The properties of an hourly schedule. If the schedule occurs multiple times a day, specify the hourly recurrence. A hourly_recurrence block as defined below.
	// +kubebuilder:validation:Optional
	HourlyRecurrence []HourlyRecurrenceParameters `json:"hourlyRecurrence,omitempty" tf:"hourly_recurrence,omitempty"`

	// The name of the dev test lab. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/devtestlab/v1beta1.Lab
	// +kubebuilder:validation:Optional
	LabName *string `json:"labName,omitempty" tf:"lab_name,omitempty"`

	// Reference to a Lab in devtestlab to populate labName.
	// +kubebuilder:validation:Optional
	LabNameRef *v1.Reference `json:"labNameRef,omitempty" tf:"-"`

	// Selector for a Lab in devtestlab to populate labName.
	// +kubebuilder:validation:Optional
	LabNameSelector *v1.Selector `json:"labNameSelector,omitempty" tf:"-"`

	// The location where the schedule is created. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The notification setting of a schedule. A notification_settings as defined below.
	// +kubebuilder:validation:Required
	NotificationSettings []ScheduleNotificationSettingsParameters `json:"notificationSettings" tf:"notification_settings,omitempty"`

	// The name of the resource group in which to create the dev test lab schedule. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The status of this schedule. Possible values are Enabled and Disabled. Defaults to Disabled.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The task type of the schedule. Possible values include LabVmsShutdownTask and LabVmAutoStart.
	// +kubebuilder:validation:Required
	TaskType *string `json:"taskType" tf:"task_type,omitempty"`

	// The time zone ID (e.g. Pacific Standard time).
	// +kubebuilder:validation:Required
	TimeZoneID *string `json:"timeZoneId" tf:"time_zone_id,omitempty"`

	// The properties of a weekly schedule. If the schedule occurs only some days of the week, specify the weekly recurrence. A weekly_recurrence block as defined below.
	// +kubebuilder:validation:Optional
	WeeklyRecurrence []WeeklyRecurrenceParameters `json:"weeklyRecurrence,omitempty" tf:"weekly_recurrence,omitempty"`
}

type WeeklyRecurrenceObservation struct {
}

type WeeklyRecurrenceParameters struct {

	// The time when the schedule takes effect.
	// +kubebuilder:validation:Required
	Time *string `json:"time" tf:"time,omitempty"`

	// A list of days that this schedule takes effect . Possible values include Monday, Tuesday, Wednesday, Thursday, Friday, Saturday and Sunday.
	// +kubebuilder:validation:Optional
	WeekDays []*string `json:"weekDays,omitempty" tf:"week_days,omitempty"`
}

// ScheduleSpec defines the desired state of Schedule
type ScheduleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ScheduleParameters `json:"forProvider"`
}

// ScheduleStatus defines the observed state of Schedule.
type ScheduleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ScheduleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Schedule is the Schema for the Schedules API. Manages automated startup and shutdown schedules for Azure Dev Test Lab.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Schedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ScheduleSpec   `json:"spec"`
	Status            ScheduleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ScheduleList contains a list of Schedules
type ScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Schedule `json:"items"`
}

// Repository type metadata.
var (
	Schedule_Kind             = "Schedule"
	Schedule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Schedule_Kind}.String()
	Schedule_KindAPIVersion   = Schedule_Kind + "." + CRDGroupVersion.String()
	Schedule_GroupVersionKind = CRDGroupVersion.WithKind(Schedule_Kind)
)

func init() {
	SchemeBuilder.Register(&Schedule{}, &ScheduleList{})
}
