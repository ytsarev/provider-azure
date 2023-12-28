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

type DataSetJSONAzureBlobStorageLocationInitParameters struct {

	// The container on the Azure Blob Storage Account hosting the file.
	Container *string `json:"container,omitempty" tf:"container,omitempty"`

	// Is the container using dynamic expression, function or system variables? Defaults to false.
	DynamicContainerEnabled *bool `json:"dynamicContainerEnabled,omitempty" tf:"dynamic_container_enabled,omitempty"`

	// Is the filename using dynamic expression, function or system variables? Defaults to false.
	DynamicFilenameEnabled *bool `json:"dynamicFilenameEnabled,omitempty" tf:"dynamic_filename_enabled,omitempty"`

	// Is the path using dynamic expression, function or system variables? Defaults to false.
	DynamicPathEnabled *bool `json:"dynamicPathEnabled,omitempty" tf:"dynamic_path_enabled,omitempty"`

	// The filename of the file on the web server.
	Filename *string `json:"filename,omitempty" tf:"filename,omitempty"`

	// The folder path to the file on the web server.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type DataSetJSONAzureBlobStorageLocationObservation struct {

	// The container on the Azure Blob Storage Account hosting the file.
	Container *string `json:"container,omitempty" tf:"container,omitempty"`

	// Is the container using dynamic expression, function or system variables? Defaults to false.
	DynamicContainerEnabled *bool `json:"dynamicContainerEnabled,omitempty" tf:"dynamic_container_enabled,omitempty"`

	// Is the filename using dynamic expression, function or system variables? Defaults to false.
	DynamicFilenameEnabled *bool `json:"dynamicFilenameEnabled,omitempty" tf:"dynamic_filename_enabled,omitempty"`

	// Is the path using dynamic expression, function or system variables? Defaults to false.
	DynamicPathEnabled *bool `json:"dynamicPathEnabled,omitempty" tf:"dynamic_path_enabled,omitempty"`

	// The filename of the file on the web server.
	Filename *string `json:"filename,omitempty" tf:"filename,omitempty"`

	// The folder path to the file on the web server.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type DataSetJSONAzureBlobStorageLocationParameters struct {

	// The container on the Azure Blob Storage Account hosting the file.
	// +kubebuilder:validation:Optional
	Container *string `json:"container" tf:"container,omitempty"`

	// Is the container using dynamic expression, function or system variables? Defaults to false.
	// +kubebuilder:validation:Optional
	DynamicContainerEnabled *bool `json:"dynamicContainerEnabled,omitempty" tf:"dynamic_container_enabled,omitempty"`

	// Is the filename using dynamic expression, function or system variables? Defaults to false.
	// +kubebuilder:validation:Optional
	DynamicFilenameEnabled *bool `json:"dynamicFilenameEnabled,omitempty" tf:"dynamic_filename_enabled,omitempty"`

	// Is the path using dynamic expression, function or system variables? Defaults to false.
	// +kubebuilder:validation:Optional
	DynamicPathEnabled *bool `json:"dynamicPathEnabled,omitempty" tf:"dynamic_path_enabled,omitempty"`

	// The filename of the file on the web server.
	// +kubebuilder:validation:Optional
	Filename *string `json:"filename" tf:"filename,omitempty"`

	// The folder path to the file on the web server.
	// +kubebuilder:validation:Optional
	Path *string `json:"path" tf:"path,omitempty"`
}

type DataSetJSONHTTPServerLocationInitParameters struct {

	// Is the filename using dynamic expression, function or system variables? Defaults to false.
	DynamicFilenameEnabled *bool `json:"dynamicFilenameEnabled,omitempty" tf:"dynamic_filename_enabled,omitempty"`

	// Is the path using dynamic expression, function or system variables? Defaults to false.
	DynamicPathEnabled *bool `json:"dynamicPathEnabled,omitempty" tf:"dynamic_path_enabled,omitempty"`

	// The filename of the file on the web server.
	Filename *string `json:"filename,omitempty" tf:"filename,omitempty"`

	// The folder path to the file on the web server.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// The base URL to the web server hosting the file.
	RelativeURL *string `json:"relativeUrl,omitempty" tf:"relative_url,omitempty"`
}

type DataSetJSONHTTPServerLocationObservation struct {

	// Is the filename using dynamic expression, function or system variables? Defaults to false.
	DynamicFilenameEnabled *bool `json:"dynamicFilenameEnabled,omitempty" tf:"dynamic_filename_enabled,omitempty"`

	// Is the path using dynamic expression, function or system variables? Defaults to false.
	DynamicPathEnabled *bool `json:"dynamicPathEnabled,omitempty" tf:"dynamic_path_enabled,omitempty"`

	// The filename of the file on the web server.
	Filename *string `json:"filename,omitempty" tf:"filename,omitempty"`

	// The folder path to the file on the web server.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// The base URL to the web server hosting the file.
	RelativeURL *string `json:"relativeUrl,omitempty" tf:"relative_url,omitempty"`
}

type DataSetJSONHTTPServerLocationParameters struct {

	// Is the filename using dynamic expression, function or system variables? Defaults to false.
	// +kubebuilder:validation:Optional
	DynamicFilenameEnabled *bool `json:"dynamicFilenameEnabled,omitempty" tf:"dynamic_filename_enabled,omitempty"`

	// Is the path using dynamic expression, function or system variables? Defaults to false.
	// +kubebuilder:validation:Optional
	DynamicPathEnabled *bool `json:"dynamicPathEnabled,omitempty" tf:"dynamic_path_enabled,omitempty"`

	// The filename of the file on the web server.
	// +kubebuilder:validation:Optional
	Filename *string `json:"filename" tf:"filename,omitempty"`

	// The folder path to the file on the web server.
	// +kubebuilder:validation:Optional
	Path *string `json:"path" tf:"path,omitempty"`

	// The base URL to the web server hosting the file.
	// +kubebuilder:validation:Optional
	RelativeURL *string `json:"relativeUrl" tf:"relative_url,omitempty"`
}

type DataSetJSONInitParameters struct {

	// A map of additional properties to associate with the Data Factory Dataset.
	// +mapType=granular
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Dataset.
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// A azure_blob_storage_location block as defined below.
	AzureBlobStorageLocation []DataSetJSONAzureBlobStorageLocationInitParameters `json:"azureBlobStorageLocation,omitempty" tf:"azure_blob_storage_location,omitempty"`

	// The description for the Data Factory Dataset.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The encoding format for the file.
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder *string `json:"folder,omitempty" tf:"folder,omitempty"`

	// A http_server_location block as defined below.
	HTTPServerLocation []DataSetJSONHTTPServerLocationInitParameters `json:"httpServerLocation,omitempty" tf:"http_server_location,omitempty"`

	// The Data Factory Linked Service name in which to associate the Dataset with.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.LinkedServiceWeb
	LinkedServiceName *string `json:"linkedServiceName,omitempty" tf:"linked_service_name,omitempty"`

	// Reference to a LinkedServiceWeb in datafactory to populate linkedServiceName.
	// +kubebuilder:validation:Optional
	LinkedServiceNameRef *v1.Reference `json:"linkedServiceNameRef,omitempty" tf:"-"`

	// Selector for a LinkedServiceWeb in datafactory to populate linkedServiceName.
	// +kubebuilder:validation:Optional
	LinkedServiceNameSelector *v1.Selector `json:"linkedServiceNameSelector,omitempty" tf:"-"`

	// A map of parameters to associate with the Data Factory Dataset.
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// A schema_column block as defined below.
	SchemaColumn []DataSetJSONSchemaColumnInitParameters `json:"schemaColumn,omitempty" tf:"schema_column,omitempty"`
}

type DataSetJSONObservation struct {

	// A map of additional properties to associate with the Data Factory Dataset.
	// +mapType=granular
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Dataset.
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// A azure_blob_storage_location block as defined below.
	AzureBlobStorageLocation []DataSetJSONAzureBlobStorageLocationObservation `json:"azureBlobStorageLocation,omitempty" tf:"azure_blob_storage_location,omitempty"`

	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// The description for the Data Factory Dataset.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The encoding format for the file.
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder *string `json:"folder,omitempty" tf:"folder,omitempty"`

	// A http_server_location block as defined below.
	HTTPServerLocation []DataSetJSONHTTPServerLocationObservation `json:"httpServerLocation,omitempty" tf:"http_server_location,omitempty"`

	// The ID of the Data Factory Dataset.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Data Factory Linked Service name in which to associate the Dataset with.
	LinkedServiceName *string `json:"linkedServiceName,omitempty" tf:"linked_service_name,omitempty"`

	// A map of parameters to associate with the Data Factory Dataset.
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// A schema_column block as defined below.
	SchemaColumn []DataSetJSONSchemaColumnObservation `json:"schemaColumn,omitempty" tf:"schema_column,omitempty"`
}

type DataSetJSONParameters struct {

	// A map of additional properties to associate with the Data Factory Dataset.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Dataset.
	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// A azure_blob_storage_location block as defined below.
	// +kubebuilder:validation:Optional
	AzureBlobStorageLocation []DataSetJSONAzureBlobStorageLocationParameters `json:"azureBlobStorageLocation,omitempty" tf:"azure_blob_storage_location,omitempty"`

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

	// The description for the Data Factory Dataset.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The encoding format for the file.
	// +kubebuilder:validation:Optional
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	// +kubebuilder:validation:Optional
	Folder *string `json:"folder,omitempty" tf:"folder,omitempty"`

	// A http_server_location block as defined below.
	// +kubebuilder:validation:Optional
	HTTPServerLocation []DataSetJSONHTTPServerLocationParameters `json:"httpServerLocation,omitempty" tf:"http_server_location,omitempty"`

	// The Data Factory Linked Service name in which to associate the Dataset with.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.LinkedServiceWeb
	// +kubebuilder:validation:Optional
	LinkedServiceName *string `json:"linkedServiceName,omitempty" tf:"linked_service_name,omitempty"`

	// Reference to a LinkedServiceWeb in datafactory to populate linkedServiceName.
	// +kubebuilder:validation:Optional
	LinkedServiceNameRef *v1.Reference `json:"linkedServiceNameRef,omitempty" tf:"-"`

	// Selector for a LinkedServiceWeb in datafactory to populate linkedServiceName.
	// +kubebuilder:validation:Optional
	LinkedServiceNameSelector *v1.Selector `json:"linkedServiceNameSelector,omitempty" tf:"-"`

	// A map of parameters to associate with the Data Factory Dataset.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// A schema_column block as defined below.
	// +kubebuilder:validation:Optional
	SchemaColumn []DataSetJSONSchemaColumnParameters `json:"schemaColumn,omitempty" tf:"schema_column,omitempty"`
}

type DataSetJSONSchemaColumnInitParameters struct {

	// The description of the column.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the column.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Type of the column. Valid values are Byte, Byte[], Boolean, Date, DateTime,DateTimeOffset, Decimal, Double, Guid, Int16, Int32, Int64, Single, String, TimeSpan. Please note these values are case sensitive.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DataSetJSONSchemaColumnObservation struct {

	// The description of the column.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the column.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Type of the column. Valid values are Byte, Byte[], Boolean, Date, DateTime,DateTimeOffset, Decimal, Double, Guid, Int16, Int32, Int64, Single, String, TimeSpan. Please note these values are case sensitive.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type DataSetJSONSchemaColumnParameters struct {

	// The description of the column.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the column.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// Type of the column. Valid values are Byte, Byte[], Boolean, Date, DateTime,DateTimeOffset, Decimal, Double, Guid, Int16, Int32, Int64, Single, String, TimeSpan. Please note these values are case sensitive.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// DataSetJSONSpec defines the desired state of DataSetJSON
type DataSetJSONSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DataSetJSONParameters `json:"forProvider"`
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
	InitProvider DataSetJSONInitParameters `json:"initProvider,omitempty"`
}

// DataSetJSONStatus defines the observed state of DataSetJSON.
type DataSetJSONStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DataSetJSONObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DataSetJSON is the Schema for the DataSetJSONs API. Manages an Azure JSON Dataset inside an Azure Data Factory.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DataSetJSON struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataSetJSONSpec   `json:"spec"`
	Status            DataSetJSONStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataSetJSONList contains a list of DataSetJSONs
type DataSetJSONList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataSetJSON `json:"items"`
}

// Repository type metadata.
var (
	DataSetJSON_Kind             = "DataSetJSON"
	DataSetJSON_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DataSetJSON_Kind}.String()
	DataSetJSON_KindAPIVersion   = DataSetJSON_Kind + "." + CRDGroupVersion.String()
	DataSetJSON_GroupVersionKind = CRDGroupVersion.WithKind(DataSetJSON_Kind)
)

func init() {
	SchemeBuilder.Register(&DataSetJSON{}, &DataSetJSONList{})
}
