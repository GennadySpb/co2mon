// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/reference/reference.proto

package reference

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Reference_Type int32

const (
	Reference_TYPE_UNSPECIFIED Reference_Type = 0
	Reference_MANAGED_BY       Reference_Type = 1
	Reference_USED_BY          Reference_Type = 2
)

// Enum value maps for Reference_Type.
var (
	Reference_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "MANAGED_BY",
		2: "USED_BY",
	}
	Reference_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"MANAGED_BY":       1,
		"USED_BY":          2,
	}
)

func (x Reference_Type) Enum() *Reference_Type {
	p := new(Reference_Type)
	*p = x
	return p
}

func (x Reference_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Reference_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_reference_reference_proto_enumTypes[0].Descriptor()
}

func (Reference_Type) Type() protoreflect.EnumType {
	return &file_yandex_cloud_reference_reference_proto_enumTypes[0]
}

func (x Reference_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Reference_Type.Descriptor instead.
func (Reference_Type) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_reference_reference_proto_rawDescGZIP(), []int{0, 0}
}

type Reference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Referrer *Referrer      `protobuf:"bytes,1,opt,name=referrer,proto3" json:"referrer,omitempty"`
	Type     Reference_Type `protobuf:"varint,2,opt,name=type,proto3,enum=yandex.cloud.reference.Reference_Type" json:"type,omitempty"`
}

func (x *Reference) Reset() {
	*x = Reference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_reference_reference_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reference) ProtoMessage() {}

func (x *Reference) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_reference_reference_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reference.ProtoReflect.Descriptor instead.
func (*Reference) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_reference_reference_proto_rawDescGZIP(), []int{0}
}

func (x *Reference) GetReferrer() *Referrer {
	if x != nil {
		return x.Referrer
	}
	return nil
}

func (x *Reference) GetType() Reference_Type {
	if x != nil {
		return x.Type
	}
	return Reference_TYPE_UNSPECIFIED
}

type Referrer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// * `type = compute.instance, id = <instance id>`
	// * `type = compute.instanceGroup, id = <instanceGroup id>`
	// * `type = loadbalancer.networkLoadBalancer, id = <networkLoadBalancer id>`
	// * `type = managed-kubernetes.cluster, id = <cluster id>`
	// * `type = managed-mysql.cluster, id = <cluster id>`
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Id   string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Referrer) Reset() {
	*x = Referrer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_reference_reference_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Referrer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Referrer) ProtoMessage() {}

func (x *Referrer) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_reference_reference_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Referrer.ProtoReflect.Descriptor instead.
func (*Referrer) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_reference_reference_proto_rawDescGZIP(), []int{1}
}

func (x *Referrer) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Referrer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_yandex_cloud_reference_reference_proto protoreflect.FileDescriptor

var file_yandex_cloud_reference_reference_proto_rawDesc = []byte{
	0x0a, 0x26, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x22, 0xc0, 0x01, 0x0a, 0x09, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x3c,
	0x0a, 0x08, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x72,
	0x65, 0x72, 0x52, 0x08, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x72, 0x12, 0x3a, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x39, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x4d, 0x41, 0x4e, 0x41, 0x47, 0x45,
	0x44, 0x5f, 0x42, 0x59, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x53, 0x45, 0x44, 0x5f, 0x42,
	0x59, 0x10, 0x02, 0x22, 0x2e, 0x0a, 0x08, 0x52, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x42, 0x62, 0x0a, 0x1a, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x3b, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_reference_reference_proto_rawDescOnce sync.Once
	file_yandex_cloud_reference_reference_proto_rawDescData = file_yandex_cloud_reference_reference_proto_rawDesc
)

func file_yandex_cloud_reference_reference_proto_rawDescGZIP() []byte {
	file_yandex_cloud_reference_reference_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_reference_reference_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_reference_reference_proto_rawDescData)
	})
	return file_yandex_cloud_reference_reference_proto_rawDescData
}

var file_yandex_cloud_reference_reference_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_reference_reference_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yandex_cloud_reference_reference_proto_goTypes = []any{
	(Reference_Type)(0), // 0: yandex.cloud.reference.Reference.Type
	(*Reference)(nil),   // 1: yandex.cloud.reference.Reference
	(*Referrer)(nil),    // 2: yandex.cloud.reference.Referrer
}
var file_yandex_cloud_reference_reference_proto_depIdxs = []int32{
	2, // 0: yandex.cloud.reference.Reference.referrer:type_name -> yandex.cloud.reference.Referrer
	0, // 1: yandex.cloud.reference.Reference.type:type_name -> yandex.cloud.reference.Reference.Type
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_yandex_cloud_reference_reference_proto_init() }
func file_yandex_cloud_reference_reference_proto_init() {
	if File_yandex_cloud_reference_reference_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_reference_reference_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Reference); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_reference_reference_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Referrer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_reference_reference_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_reference_reference_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_reference_reference_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_reference_reference_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_reference_reference_proto_msgTypes,
	}.Build()
	File_yandex_cloud_reference_reference_proto = out.File
	file_yandex_cloud_reference_reference_proto_rawDesc = nil
	file_yandex_cloud_reference_reference_proto_goTypes = nil
	file_yandex_cloud_reference_reference_proto_depIdxs = nil
}
