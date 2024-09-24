// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/datasphere/v2/docker_image_service.proto

package datasphere

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud/api"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ActivateDockerImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	DockerId  string `protobuf:"bytes,2,opt,name=docker_id,json=dockerId,proto3" json:"docker_id,omitempty"`
}

func (x *ActivateDockerImageRequest) Reset() {
	*x = ActivateDockerImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_datasphere_v2_docker_image_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivateDockerImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivateDockerImageRequest) ProtoMessage() {}

func (x *ActivateDockerImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datasphere_v2_docker_image_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivateDockerImageRequest.ProtoReflect.Descriptor instead.
func (*ActivateDockerImageRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datasphere_v2_docker_image_service_proto_rawDescGZIP(), []int{0}
}

func (x *ActivateDockerImageRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *ActivateDockerImageRequest) GetDockerId() string {
	if x != nil {
		return x.DockerId
	}
	return ""
}

var File_yandex_cloud_datasphere_v2_docker_image_service_proto protoreflect.FileDescriptor

var file_yandex_cloud_datasphere_v2_docker_image_service_proto_rawDesc = []byte{
	0x0a, 0x35, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x64, 0x6f, 0x63,
	0x6b, 0x65, 0x72, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65,
	0x2e, 0x76, 0x32, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x74, 0x0a, 0x1a, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7,
	0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x09, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8,
	0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x08, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x49, 0x64,
	0x32, 0xc2, 0x01, 0x0a, 0x12, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xab, 0x01, 0x0a, 0x08, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x12, 0x36, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65, 0x2e, 0x76,
	0x32, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x44, 0xb2, 0xd2, 0x2a, 0x17, 0x12, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x23, 0x3a, 0x01, 0x2a, 0x22, 0x1e, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x70, 0x68, 0x65,
	0x72, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x2f, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x42, 0x6b, 0x0a, 0x1e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x70,
	0x68, 0x65, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x70,
	0x68, 0x65, 0x72, 0x65, 0x2f, 0x76, 0x32, 0x3b, 0x64, 0x61, 0x74, 0x61, 0x73, 0x70, 0x68, 0x65,
	0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_datasphere_v2_docker_image_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_datasphere_v2_docker_image_service_proto_rawDescData = file_yandex_cloud_datasphere_v2_docker_image_service_proto_rawDesc
)

func file_yandex_cloud_datasphere_v2_docker_image_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_datasphere_v2_docker_image_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_datasphere_v2_docker_image_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_datasphere_v2_docker_image_service_proto_rawDescData)
	})
	return file_yandex_cloud_datasphere_v2_docker_image_service_proto_rawDescData
}

var file_yandex_cloud_datasphere_v2_docker_image_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_yandex_cloud_datasphere_v2_docker_image_service_proto_goTypes = []any{
	(*ActivateDockerImageRequest)(nil), // 0: yandex.cloud.datasphere.v2.ActivateDockerImageRequest
	(*operation.Operation)(nil),        // 1: yandex.cloud.operation.Operation
}
var file_yandex_cloud_datasphere_v2_docker_image_service_proto_depIdxs = []int32{
	0, // 0: yandex.cloud.datasphere.v2.DockerImageService.Activate:input_type -> yandex.cloud.datasphere.v2.ActivateDockerImageRequest
	1, // 1: yandex.cloud.datasphere.v2.DockerImageService.Activate:output_type -> yandex.cloud.operation.Operation
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_yandex_cloud_datasphere_v2_docker_image_service_proto_init() }
func file_yandex_cloud_datasphere_v2_docker_image_service_proto_init() {
	if File_yandex_cloud_datasphere_v2_docker_image_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_datasphere_v2_docker_image_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ActivateDockerImageRequest); i {
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
			RawDescriptor: file_yandex_cloud_datasphere_v2_docker_image_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_datasphere_v2_docker_image_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_datasphere_v2_docker_image_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_datasphere_v2_docker_image_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_datasphere_v2_docker_image_service_proto = out.File
	file_yandex_cloud_datasphere_v2_docker_image_service_proto_rawDesc = nil
	file_yandex_cloud_datasphere_v2_docker_image_service_proto_goTypes = nil
	file_yandex_cloud_datasphere_v2_docker_image_service_proto_depIdxs = nil
}
