//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventRelayNamespace) DeepCopyInto(out *EventRelayNamespace) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventRelayNamespace.
func (in *EventRelayNamespace) DeepCopy() *EventRelayNamespace {
	if in == nil {
		return nil
	}
	out := new(EventRelayNamespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventRelayNamespace) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventRelayNamespaceInitParameters) DeepCopyInto(out *EventRelayNamespaceInitParameters) {
	*out = *in
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.SkuName != nil {
		in, out := &in.SkuName, &out.SkuName
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventRelayNamespaceInitParameters.
func (in *EventRelayNamespaceInitParameters) DeepCopy() *EventRelayNamespaceInitParameters {
	if in == nil {
		return nil
	}
	out := new(EventRelayNamespaceInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventRelayNamespaceList) DeepCopyInto(out *EventRelayNamespaceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EventRelayNamespace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventRelayNamespaceList.
func (in *EventRelayNamespaceList) DeepCopy() *EventRelayNamespaceList {
	if in == nil {
		return nil
	}
	out := new(EventRelayNamespaceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventRelayNamespaceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventRelayNamespaceObservation) DeepCopyInto(out *EventRelayNamespaceObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MetricID != nil {
		in, out := &in.MetricID, &out.MetricID
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.SkuName != nil {
		in, out := &in.SkuName, &out.SkuName
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventRelayNamespaceObservation.
func (in *EventRelayNamespaceObservation) DeepCopy() *EventRelayNamespaceObservation {
	if in == nil {
		return nil
	}
	out := new(EventRelayNamespaceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventRelayNamespaceParameters) DeepCopyInto(out *EventRelayNamespaceParameters) {
	*out = *in
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.SkuName != nil {
		in, out := &in.SkuName, &out.SkuName
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventRelayNamespaceParameters.
func (in *EventRelayNamespaceParameters) DeepCopy() *EventRelayNamespaceParameters {
	if in == nil {
		return nil
	}
	out := new(EventRelayNamespaceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventRelayNamespaceSpec) DeepCopyInto(out *EventRelayNamespaceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventRelayNamespaceSpec.
func (in *EventRelayNamespaceSpec) DeepCopy() *EventRelayNamespaceSpec {
	if in == nil {
		return nil
	}
	out := new(EventRelayNamespaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventRelayNamespaceStatus) DeepCopyInto(out *EventRelayNamespaceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventRelayNamespaceStatus.
func (in *EventRelayNamespaceStatus) DeepCopy() *EventRelayNamespaceStatus {
	if in == nil {
		return nil
	}
	out := new(EventRelayNamespaceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HybridConnection) DeepCopyInto(out *HybridConnection) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HybridConnection.
func (in *HybridConnection) DeepCopy() *HybridConnection {
	if in == nil {
		return nil
	}
	out := new(HybridConnection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HybridConnection) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HybridConnectionAuthorizationRule) DeepCopyInto(out *HybridConnectionAuthorizationRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HybridConnectionAuthorizationRule.
func (in *HybridConnectionAuthorizationRule) DeepCopy() *HybridConnectionAuthorizationRule {
	if in == nil {
		return nil
	}
	out := new(HybridConnectionAuthorizationRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HybridConnectionAuthorizationRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HybridConnectionAuthorizationRuleInitParameters) DeepCopyInto(out *HybridConnectionAuthorizationRuleInitParameters) {
	*out = *in
	if in.Listen != nil {
		in, out := &in.Listen, &out.Listen
		*out = new(bool)
		**out = **in
	}
	if in.Manage != nil {
		in, out := &in.Manage, &out.Manage
		*out = new(bool)
		**out = **in
	}
	if in.NamespaceName != nil {
		in, out := &in.NamespaceName, &out.NamespaceName
		*out = new(string)
		**out = **in
	}
	if in.NamespaceNameRef != nil {
		in, out := &in.NamespaceNameRef, &out.NamespaceNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.NamespaceNameSelector != nil {
		in, out := &in.NamespaceNameSelector, &out.NamespaceNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Send != nil {
		in, out := &in.Send, &out.Send
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HybridConnectionAuthorizationRuleInitParameters.
func (in *HybridConnectionAuthorizationRuleInitParameters) DeepCopy() *HybridConnectionAuthorizationRuleInitParameters {
	if in == nil {
		return nil
	}
	out := new(HybridConnectionAuthorizationRuleInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HybridConnectionAuthorizationRuleList) DeepCopyInto(out *HybridConnectionAuthorizationRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HybridConnectionAuthorizationRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HybridConnectionAuthorizationRuleList.
func (in *HybridConnectionAuthorizationRuleList) DeepCopy() *HybridConnectionAuthorizationRuleList {
	if in == nil {
		return nil
	}
	out := new(HybridConnectionAuthorizationRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HybridConnectionAuthorizationRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HybridConnectionAuthorizationRuleObservation) DeepCopyInto(out *HybridConnectionAuthorizationRuleObservation) {
	*out = *in
	if in.HybridConnectionName != nil {
		in, out := &in.HybridConnectionName, &out.HybridConnectionName
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Listen != nil {
		in, out := &in.Listen, &out.Listen
		*out = new(bool)
		**out = **in
	}
	if in.Manage != nil {
		in, out := &in.Manage, &out.Manage
		*out = new(bool)
		**out = **in
	}
	if in.NamespaceName != nil {
		in, out := &in.NamespaceName, &out.NamespaceName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.Send != nil {
		in, out := &in.Send, &out.Send
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HybridConnectionAuthorizationRuleObservation.
func (in *HybridConnectionAuthorizationRuleObservation) DeepCopy() *HybridConnectionAuthorizationRuleObservation {
	if in == nil {
		return nil
	}
	out := new(HybridConnectionAuthorizationRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HybridConnectionAuthorizationRuleParameters) DeepCopyInto(out *HybridConnectionAuthorizationRuleParameters) {
	*out = *in
	if in.HybridConnectionName != nil {
		in, out := &in.HybridConnectionName, &out.HybridConnectionName
		*out = new(string)
		**out = **in
	}
	if in.HybridConnectionNameRef != nil {
		in, out := &in.HybridConnectionNameRef, &out.HybridConnectionNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.HybridConnectionNameSelector != nil {
		in, out := &in.HybridConnectionNameSelector, &out.HybridConnectionNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Listen != nil {
		in, out := &in.Listen, &out.Listen
		*out = new(bool)
		**out = **in
	}
	if in.Manage != nil {
		in, out := &in.Manage, &out.Manage
		*out = new(bool)
		**out = **in
	}
	if in.NamespaceName != nil {
		in, out := &in.NamespaceName, &out.NamespaceName
		*out = new(string)
		**out = **in
	}
	if in.NamespaceNameRef != nil {
		in, out := &in.NamespaceNameRef, &out.NamespaceNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.NamespaceNameSelector != nil {
		in, out := &in.NamespaceNameSelector, &out.NamespaceNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Send != nil {
		in, out := &in.Send, &out.Send
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HybridConnectionAuthorizationRuleParameters.
func (in *HybridConnectionAuthorizationRuleParameters) DeepCopy() *HybridConnectionAuthorizationRuleParameters {
	if in == nil {
		return nil
	}
	out := new(HybridConnectionAuthorizationRuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HybridConnectionAuthorizationRuleSpec) DeepCopyInto(out *HybridConnectionAuthorizationRuleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HybridConnectionAuthorizationRuleSpec.
func (in *HybridConnectionAuthorizationRuleSpec) DeepCopy() *HybridConnectionAuthorizationRuleSpec {
	if in == nil {
		return nil
	}
	out := new(HybridConnectionAuthorizationRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HybridConnectionAuthorizationRuleStatus) DeepCopyInto(out *HybridConnectionAuthorizationRuleStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HybridConnectionAuthorizationRuleStatus.
func (in *HybridConnectionAuthorizationRuleStatus) DeepCopy() *HybridConnectionAuthorizationRuleStatus {
	if in == nil {
		return nil
	}
	out := new(HybridConnectionAuthorizationRuleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HybridConnectionInitParameters) DeepCopyInto(out *HybridConnectionInitParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RelayNamespaceName != nil {
		in, out := &in.RelayNamespaceName, &out.RelayNamespaceName
		*out = new(string)
		**out = **in
	}
	if in.RelayNamespaceNameRef != nil {
		in, out := &in.RelayNamespaceNameRef, &out.RelayNamespaceNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.RelayNamespaceNameSelector != nil {
		in, out := &in.RelayNamespaceNameSelector, &out.RelayNamespaceNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.RequiresClientAuthorization != nil {
		in, out := &in.RequiresClientAuthorization, &out.RequiresClientAuthorization
		*out = new(bool)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.UserMetadata != nil {
		in, out := &in.UserMetadata, &out.UserMetadata
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HybridConnectionInitParameters.
func (in *HybridConnectionInitParameters) DeepCopy() *HybridConnectionInitParameters {
	if in == nil {
		return nil
	}
	out := new(HybridConnectionInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HybridConnectionList) DeepCopyInto(out *HybridConnectionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HybridConnection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HybridConnectionList.
func (in *HybridConnectionList) DeepCopy() *HybridConnectionList {
	if in == nil {
		return nil
	}
	out := new(HybridConnectionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HybridConnectionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HybridConnectionObservation) DeepCopyInto(out *HybridConnectionObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RelayNamespaceName != nil {
		in, out := &in.RelayNamespaceName, &out.RelayNamespaceName
		*out = new(string)
		**out = **in
	}
	if in.RequiresClientAuthorization != nil {
		in, out := &in.RequiresClientAuthorization, &out.RequiresClientAuthorization
		*out = new(bool)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.UserMetadata != nil {
		in, out := &in.UserMetadata, &out.UserMetadata
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HybridConnectionObservation.
func (in *HybridConnectionObservation) DeepCopy() *HybridConnectionObservation {
	if in == nil {
		return nil
	}
	out := new(HybridConnectionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HybridConnectionParameters) DeepCopyInto(out *HybridConnectionParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RelayNamespaceName != nil {
		in, out := &in.RelayNamespaceName, &out.RelayNamespaceName
		*out = new(string)
		**out = **in
	}
	if in.RelayNamespaceNameRef != nil {
		in, out := &in.RelayNamespaceNameRef, &out.RelayNamespaceNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.RelayNamespaceNameSelector != nil {
		in, out := &in.RelayNamespaceNameSelector, &out.RelayNamespaceNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.RequiresClientAuthorization != nil {
		in, out := &in.RequiresClientAuthorization, &out.RequiresClientAuthorization
		*out = new(bool)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.UserMetadata != nil {
		in, out := &in.UserMetadata, &out.UserMetadata
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HybridConnectionParameters.
func (in *HybridConnectionParameters) DeepCopy() *HybridConnectionParameters {
	if in == nil {
		return nil
	}
	out := new(HybridConnectionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HybridConnectionSpec) DeepCopyInto(out *HybridConnectionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HybridConnectionSpec.
func (in *HybridConnectionSpec) DeepCopy() *HybridConnectionSpec {
	if in == nil {
		return nil
	}
	out := new(HybridConnectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HybridConnectionStatus) DeepCopyInto(out *HybridConnectionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HybridConnectionStatus.
func (in *HybridConnectionStatus) DeepCopy() *HybridConnectionStatus {
	if in == nil {
		return nil
	}
	out := new(HybridConnectionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceAuthorizationRule) DeepCopyInto(out *NamespaceAuthorizationRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceAuthorizationRule.
func (in *NamespaceAuthorizationRule) DeepCopy() *NamespaceAuthorizationRule {
	if in == nil {
		return nil
	}
	out := new(NamespaceAuthorizationRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NamespaceAuthorizationRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceAuthorizationRuleInitParameters) DeepCopyInto(out *NamespaceAuthorizationRuleInitParameters) {
	*out = *in
	if in.Listen != nil {
		in, out := &in.Listen, &out.Listen
		*out = new(bool)
		**out = **in
	}
	if in.Manage != nil {
		in, out := &in.Manage, &out.Manage
		*out = new(bool)
		**out = **in
	}
	if in.Send != nil {
		in, out := &in.Send, &out.Send
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceAuthorizationRuleInitParameters.
func (in *NamespaceAuthorizationRuleInitParameters) DeepCopy() *NamespaceAuthorizationRuleInitParameters {
	if in == nil {
		return nil
	}
	out := new(NamespaceAuthorizationRuleInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceAuthorizationRuleList) DeepCopyInto(out *NamespaceAuthorizationRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NamespaceAuthorizationRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceAuthorizationRuleList.
func (in *NamespaceAuthorizationRuleList) DeepCopy() *NamespaceAuthorizationRuleList {
	if in == nil {
		return nil
	}
	out := new(NamespaceAuthorizationRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NamespaceAuthorizationRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceAuthorizationRuleObservation) DeepCopyInto(out *NamespaceAuthorizationRuleObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Listen != nil {
		in, out := &in.Listen, &out.Listen
		*out = new(bool)
		**out = **in
	}
	if in.Manage != nil {
		in, out := &in.Manage, &out.Manage
		*out = new(bool)
		**out = **in
	}
	if in.NamespaceName != nil {
		in, out := &in.NamespaceName, &out.NamespaceName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.Send != nil {
		in, out := &in.Send, &out.Send
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceAuthorizationRuleObservation.
func (in *NamespaceAuthorizationRuleObservation) DeepCopy() *NamespaceAuthorizationRuleObservation {
	if in == nil {
		return nil
	}
	out := new(NamespaceAuthorizationRuleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceAuthorizationRuleParameters) DeepCopyInto(out *NamespaceAuthorizationRuleParameters) {
	*out = *in
	if in.Listen != nil {
		in, out := &in.Listen, &out.Listen
		*out = new(bool)
		**out = **in
	}
	if in.Manage != nil {
		in, out := &in.Manage, &out.Manage
		*out = new(bool)
		**out = **in
	}
	if in.NamespaceName != nil {
		in, out := &in.NamespaceName, &out.NamespaceName
		*out = new(string)
		**out = **in
	}
	if in.NamespaceNameRef != nil {
		in, out := &in.NamespaceNameRef, &out.NamespaceNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.NamespaceNameSelector != nil {
		in, out := &in.NamespaceNameSelector, &out.NamespaceNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceGroupName != nil {
		in, out := &in.ResourceGroupName, &out.ResourceGroupName
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroupNameRef != nil {
		in, out := &in.ResourceGroupNameRef, &out.ResourceGroupNameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceGroupNameSelector != nil {
		in, out := &in.ResourceGroupNameSelector, &out.ResourceGroupNameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Send != nil {
		in, out := &in.Send, &out.Send
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceAuthorizationRuleParameters.
func (in *NamespaceAuthorizationRuleParameters) DeepCopy() *NamespaceAuthorizationRuleParameters {
	if in == nil {
		return nil
	}
	out := new(NamespaceAuthorizationRuleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceAuthorizationRuleSpec) DeepCopyInto(out *NamespaceAuthorizationRuleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceAuthorizationRuleSpec.
func (in *NamespaceAuthorizationRuleSpec) DeepCopy() *NamespaceAuthorizationRuleSpec {
	if in == nil {
		return nil
	}
	out := new(NamespaceAuthorizationRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceAuthorizationRuleStatus) DeepCopyInto(out *NamespaceAuthorizationRuleStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceAuthorizationRuleStatus.
func (in *NamespaceAuthorizationRuleStatus) DeepCopy() *NamespaceAuthorizationRuleStatus {
	if in == nil {
		return nil
	}
	out := new(NamespaceAuthorizationRuleStatus)
	in.DeepCopyInto(out)
	return out
}
