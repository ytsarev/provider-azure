/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	v1beta11 "github.com/upbound/provider-azure/apis/azure/v1beta1"
	v1beta1 "github.com/upbound/provider-azure/apis/network/v1beta1"
	rconfig "github.com/upbound/provider-azure/apis/rconfig"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this KubernetesCluster.
func (mg *KubernetesCluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.APIServerAccessProfile); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.APIServerAccessProfile[i3].SubnetID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.APIServerAccessProfile[i3].SubnetIDRef,
			Selector:     mg.Spec.ForProvider.APIServerAccessProfile[i3].SubnetIDSelector,
			To: reference.To{
				List:    &v1beta1.SubnetList{},
				Managed: &v1beta1.Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.APIServerAccessProfile[i3].SubnetID")
		}
		mg.Spec.ForProvider.APIServerAccessProfile[i3].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.APIServerAccessProfile[i3].SubnetIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.AciConnectorLinux); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AciConnectorLinux[i3].SubnetName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.AciConnectorLinux[i3].SubnetNameRef,
			Selector:     mg.Spec.ForProvider.AciConnectorLinux[i3].SubnetNameSelector,
			To: reference.To{
				List:    &v1beta1.SubnetList{},
				Managed: &v1beta1.Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.AciConnectorLinux[i3].SubnetName")
		}
		mg.Spec.ForProvider.AciConnectorLinux[i3].SubnetName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.AciConnectorLinux[i3].SubnetNameRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.DefaultNodePool); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DefaultNodePool[i3].PodSubnetID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.DefaultNodePool[i3].PodSubnetIDRef,
			Selector:     mg.Spec.ForProvider.DefaultNodePool[i3].PodSubnetIDSelector,
			To: reference.To{
				List:    &v1beta1.SubnetList{},
				Managed: &v1beta1.Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.DefaultNodePool[i3].PodSubnetID")
		}
		mg.Spec.ForProvider.DefaultNodePool[i3].PodSubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.DefaultNodePool[i3].PodSubnetIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.DefaultNodePool); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DefaultNodePool[i3].VnetSubnetID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.DefaultNodePool[i3].VnetSubnetIDRef,
			Selector:     mg.Spec.ForProvider.DefaultNodePool[i3].VnetSubnetIDSelector,
			To: reference.To{
				List:    &v1beta1.SubnetList{},
				Managed: &v1beta1.Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.DefaultNodePool[i3].VnetSubnetID")
		}
		mg.Spec.ForProvider.DefaultNodePool[i3].VnetSubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.DefaultNodePool[i3].VnetSubnetIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.IngressApplicationGateway); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.IngressApplicationGateway[i3].SubnetID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.IngressApplicationGateway[i3].SubnetIDRef,
			Selector:     mg.Spec.ForProvider.IngressApplicationGateway[i3].SubnetIDSelector,
			To: reference.To{
				List:    &v1beta1.SubnetList{},
				Managed: &v1beta1.Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.IngressApplicationGateway[i3].SubnetID")
		}
		mg.Spec.ForProvider.IngressApplicationGateway[i3].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.IngressApplicationGateway[i3].SubnetIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PrivateDNSZoneID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.PrivateDNSZoneIDRef,
		Selector:     mg.Spec.ForProvider.PrivateDNSZoneIDSelector,
		To: reference.To{
			List:    &v1beta1.PrivateDNSZoneList{},
			Managed: &v1beta1.PrivateDNSZone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PrivateDNSZoneID")
	}
	mg.Spec.ForProvider.PrivateDNSZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PrivateDNSZoneIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta11.ResourceGroupList{},
			Managed: &v1beta11.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.APIServerAccessProfile); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.APIServerAccessProfile[i3].SubnetID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.APIServerAccessProfile[i3].SubnetIDRef,
			Selector:     mg.Spec.InitProvider.APIServerAccessProfile[i3].SubnetIDSelector,
			To: reference.To{
				List:    &v1beta1.SubnetList{},
				Managed: &v1beta1.Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.APIServerAccessProfile[i3].SubnetID")
		}
		mg.Spec.InitProvider.APIServerAccessProfile[i3].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.APIServerAccessProfile[i3].SubnetIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.AciConnectorLinux); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.AciConnectorLinux[i3].SubnetName),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.AciConnectorLinux[i3].SubnetNameRef,
			Selector:     mg.Spec.InitProvider.AciConnectorLinux[i3].SubnetNameSelector,
			To: reference.To{
				List:    &v1beta1.SubnetList{},
				Managed: &v1beta1.Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.AciConnectorLinux[i3].SubnetName")
		}
		mg.Spec.InitProvider.AciConnectorLinux[i3].SubnetName = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.AciConnectorLinux[i3].SubnetNameRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.DefaultNodePool); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DefaultNodePool[i3].PodSubnetID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.DefaultNodePool[i3].PodSubnetIDRef,
			Selector:     mg.Spec.InitProvider.DefaultNodePool[i3].PodSubnetIDSelector,
			To: reference.To{
				List:    &v1beta1.SubnetList{},
				Managed: &v1beta1.Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.DefaultNodePool[i3].PodSubnetID")
		}
		mg.Spec.InitProvider.DefaultNodePool[i3].PodSubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.DefaultNodePool[i3].PodSubnetIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.DefaultNodePool); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DefaultNodePool[i3].VnetSubnetID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.DefaultNodePool[i3].VnetSubnetIDRef,
			Selector:     mg.Spec.InitProvider.DefaultNodePool[i3].VnetSubnetIDSelector,
			To: reference.To{
				List:    &v1beta1.SubnetList{},
				Managed: &v1beta1.Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.DefaultNodePool[i3].VnetSubnetID")
		}
		mg.Spec.InitProvider.DefaultNodePool[i3].VnetSubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.DefaultNodePool[i3].VnetSubnetIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.IngressApplicationGateway); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.IngressApplicationGateway[i3].SubnetID),
			Extract:      rconfig.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.IngressApplicationGateway[i3].SubnetIDRef,
			Selector:     mg.Spec.InitProvider.IngressApplicationGateway[i3].SubnetIDSelector,
			To: reference.To{
				List:    &v1beta1.SubnetList{},
				Managed: &v1beta1.Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.IngressApplicationGateway[i3].SubnetID")
		}
		mg.Spec.InitProvider.IngressApplicationGateway[i3].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.IngressApplicationGateway[i3].SubnetIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PrivateDNSZoneID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.PrivateDNSZoneIDRef,
		Selector:     mg.Spec.InitProvider.PrivateDNSZoneIDSelector,
		To: reference.To{
			List:    &v1beta1.PrivateDNSZoneList{},
			Managed: &v1beta1.PrivateDNSZone{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.PrivateDNSZoneID")
	}
	mg.Spec.InitProvider.PrivateDNSZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.PrivateDNSZoneIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this KubernetesClusterNodePool.
func (mg *KubernetesClusterNodePool) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KubernetesClusterID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.KubernetesClusterIDRef,
		Selector:     mg.Spec.ForProvider.KubernetesClusterIDSelector,
		To: reference.To{
			List:    &KubernetesClusterList{},
			Managed: &KubernetesCluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KubernetesClusterID")
	}
	mg.Spec.ForProvider.KubernetesClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KubernetesClusterIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PodSubnetID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.PodSubnetIDRef,
		Selector:     mg.Spec.ForProvider.PodSubnetIDSelector,
		To: reference.To{
			List:    &v1beta1.SubnetList{},
			Managed: &v1beta1.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PodSubnetID")
	}
	mg.Spec.ForProvider.PodSubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PodSubnetIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VnetSubnetID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.VnetSubnetIDRef,
		Selector:     mg.Spec.ForProvider.VnetSubnetIDSelector,
		To: reference.To{
			List:    &v1beta1.SubnetList{},
			Managed: &v1beta1.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VnetSubnetID")
	}
	mg.Spec.ForProvider.VnetSubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VnetSubnetIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.PodSubnetID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.PodSubnetIDRef,
		Selector:     mg.Spec.InitProvider.PodSubnetIDSelector,
		To: reference.To{
			List:    &v1beta1.SubnetList{},
			Managed: &v1beta1.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.PodSubnetID")
	}
	mg.Spec.InitProvider.PodSubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.PodSubnetIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.VnetSubnetID),
		Extract:      rconfig.ExtractResourceID(),
		Reference:    mg.Spec.InitProvider.VnetSubnetIDRef,
		Selector:     mg.Spec.InitProvider.VnetSubnetIDSelector,
		To: reference.To{
			List:    &v1beta1.SubnetList{},
			Managed: &v1beta1.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.VnetSubnetID")
	}
	mg.Spec.InitProvider.VnetSubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.VnetSubnetIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this KubernetesFleetManager.
func (mg *KubernetesFleetManager) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta11.ResourceGroupList{},
			Managed: &v1beta11.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}
