// Code generated by protoc-gen-goext. DO NOT EDIT.

package apploadbalancer

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *TargetGroup) SetId(v string) {
	m.Id = v
}

func (m *TargetGroup) SetName(v string) {
	m.Name = v
}

func (m *TargetGroup) SetDescription(v string) {
	m.Description = v
}

func (m *TargetGroup) SetFolderId(v string) {
	m.FolderId = v
}

func (m *TargetGroup) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *TargetGroup) SetTargets(v []*Target) {
	m.Targets = v
}

func (m *TargetGroup) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

type Target_AddressType = isTarget_AddressType

func (m *Target) SetAddressType(v Target_AddressType) {
	m.AddressType = v
}

func (m *Target) SetIpAddress(v string) {
	m.AddressType = &Target_IpAddress{
		IpAddress: v,
	}
}

func (m *Target) SetSubnetId(v string) {
	m.SubnetId = v
}

func (m *Target) SetPrivateIpv4Address(v bool) {
	m.PrivateIpv4Address = v
}
