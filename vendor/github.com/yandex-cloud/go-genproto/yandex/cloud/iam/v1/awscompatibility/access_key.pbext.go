// Code generated by protoc-gen-goext. DO NOT EDIT.

package awscompatibility

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *AccessKey) SetId(v string) {
	m.Id = v
}

func (m *AccessKey) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *AccessKey) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *AccessKey) SetDescription(v string) {
	m.Description = v
}

func (m *AccessKey) SetKeyId(v string) {
	m.KeyId = v
}

func (m *AccessKey) SetLastUsedAt(v *timestamppb.Timestamp) {
	m.LastUsedAt = v
}
