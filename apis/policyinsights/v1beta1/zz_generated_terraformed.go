/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/resource"
	"github.com/upbound/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this ResourcePolicyRemediation
func (mg *ResourcePolicyRemediation) GetTerraformResourceType() string {
	return "azurerm_resource_policy_remediation"
}

// GetConnectionDetailsMapping for this ResourcePolicyRemediation
func (tr *ResourcePolicyRemediation) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this ResourcePolicyRemediation
func (tr *ResourcePolicyRemediation) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ResourcePolicyRemediation
func (tr *ResourcePolicyRemediation) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ResourcePolicyRemediation
func (tr *ResourcePolicyRemediation) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ResourcePolicyRemediation
func (tr *ResourcePolicyRemediation) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ResourcePolicyRemediation
func (tr *ResourcePolicyRemediation) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this ResourcePolicyRemediation using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ResourcePolicyRemediation) LateInitialize(attrs []byte) (bool, error) {
	params := &ResourcePolicyRemediationParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ResourcePolicyRemediation) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SubscriptionPolicyRemediation
func (mg *SubscriptionPolicyRemediation) GetTerraformResourceType() string {
	return "azurerm_subscription_policy_remediation"
}

// GetConnectionDetailsMapping for this SubscriptionPolicyRemediation
func (tr *SubscriptionPolicyRemediation) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SubscriptionPolicyRemediation
func (tr *SubscriptionPolicyRemediation) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SubscriptionPolicyRemediation
func (tr *SubscriptionPolicyRemediation) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SubscriptionPolicyRemediation
func (tr *SubscriptionPolicyRemediation) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SubscriptionPolicyRemediation
func (tr *SubscriptionPolicyRemediation) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SubscriptionPolicyRemediation
func (tr *SubscriptionPolicyRemediation) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SubscriptionPolicyRemediation using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SubscriptionPolicyRemediation) LateInitialize(attrs []byte) (bool, error) {
	params := &SubscriptionPolicyRemediationParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SubscriptionPolicyRemediation) GetTerraformSchemaVersion() int {
	return 0
}
