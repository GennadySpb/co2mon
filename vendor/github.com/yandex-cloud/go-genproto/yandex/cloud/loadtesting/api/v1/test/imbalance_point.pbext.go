// Code generated by protoc-gen-goext. DO NOT EDIT.

package test

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *ImbalancePoint) SetAt(v *timestamppb.Timestamp) {
	m.At = v
}

func (m *ImbalancePoint) SetRps(v int64) {
	m.Rps = v
}

func (m *ImbalancePoint) SetComment(v string) {
	m.Comment = v
}
