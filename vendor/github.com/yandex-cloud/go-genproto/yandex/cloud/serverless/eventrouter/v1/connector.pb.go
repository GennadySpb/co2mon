// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/serverless/eventrouter/v1/connector.proto

package eventrouter

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Status of the connector.
type Connector_Status int32

const (
	Connector_STATUS_UNSPECIFIED Connector_Status = 0
	Connector_RUNNING            Connector_Status = 1
	// disabled by user
	Connector_STOPPED Connector_Status = 2
	// source does not exist
	Connector_RESOURCE_NOT_FOUND Connector_Status = 3
	// service account does not have read permission on source
	Connector_PERMISSION_DENIED Connector_Status = 4
	// service account not found
	Connector_SUBJECT_NOT_FOUND Connector_Status = 5
)

// Enum value maps for Connector_Status.
var (
	Connector_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "RUNNING",
		2: "STOPPED",
		3: "RESOURCE_NOT_FOUND",
		4: "PERMISSION_DENIED",
		5: "SUBJECT_NOT_FOUND",
	}
	Connector_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"RUNNING":            1,
		"STOPPED":            2,
		"RESOURCE_NOT_FOUND": 3,
		"PERMISSION_DENIED":  4,
		"SUBJECT_NOT_FOUND":  5,
	}
)

func (x Connector_Status) Enum() *Connector_Status {
	p := new(Connector_Status)
	*p = x
	return p
}

func (x Connector_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Connector_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_serverless_eventrouter_v1_connector_proto_enumTypes[0].Descriptor()
}

func (Connector_Status) Type() protoreflect.EnumType {
	return &file_yandex_cloud_serverless_eventrouter_v1_connector_proto_enumTypes[0]
}

func (x Connector_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Connector_Status.Descriptor instead.
func (Connector_Status) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_serverless_eventrouter_v1_connector_proto_rawDescGZIP(), []int{0, 0}
}

type Connector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the connector.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the bus that the connector belongs to.
	BusId string `protobuf:"bytes,2,opt,name=bus_id,json=busId,proto3" json:"bus_id,omitempty"`
	// ID of the folder that the connector resides in.
	FolderId string `protobuf:"bytes,3,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// ID of the cloud that the connector resides in.
	CloudId string `protobuf:"bytes,4,opt,name=cloud_id,json=cloudId,proto3" json:"cloud_id,omitempty"`
	// Creation timestamp.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Name of the connector.
	Name string `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the connector.
	Description string `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	// Resource labels as `key:value` pairs.
	Labels map[string]string `protobuf:"bytes,8,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Source of the connector.
	Source *Source `protobuf:"bytes,9,opt,name=source,proto3" json:"source,omitempty"`
	// Deletion protection.
	DeletionProtection bool `protobuf:"varint,10,opt,name=deletion_protection,json=deletionProtection,proto3" json:"deletion_protection,omitempty"`
	// Status of the connector.
	Status Connector_Status `protobuf:"varint,11,opt,name=status,proto3,enum=yandex.cloud.serverless.eventrouter.v1.Connector_Status" json:"status,omitempty"`
}

func (x *Connector) Reset() {
	*x = Connector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_serverless_eventrouter_v1_connector_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Connector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Connector) ProtoMessage() {}

func (x *Connector) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_serverless_eventrouter_v1_connector_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Connector.ProtoReflect.Descriptor instead.
func (*Connector) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_serverless_eventrouter_v1_connector_proto_rawDescGZIP(), []int{0}
}

func (x *Connector) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Connector) GetBusId() string {
	if x != nil {
		return x.BusId
	}
	return ""
}

func (x *Connector) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *Connector) GetCloudId() string {
	if x != nil {
		return x.CloudId
	}
	return ""
}

func (x *Connector) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Connector) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Connector) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Connector) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Connector) GetSource() *Source {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *Connector) GetDeletionProtection() bool {
	if x != nil {
		return x.DeletionProtection
	}
	return false
}

func (x *Connector) GetStatus() Connector_Status {
	if x != nil {
		return x.Status
	}
	return Connector_STATUS_UNSPECIFIED
}

type Source struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Source:
	//
	//	*Source_DataStream
	//	*Source_MessageQueue
	Source isSource_Source `protobuf_oneof:"source"`
}

func (x *Source) Reset() {
	*x = Source{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_serverless_eventrouter_v1_connector_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Source) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Source) ProtoMessage() {}

func (x *Source) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_serverless_eventrouter_v1_connector_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Source.ProtoReflect.Descriptor instead.
func (*Source) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_serverless_eventrouter_v1_connector_proto_rawDescGZIP(), []int{1}
}

func (m *Source) GetSource() isSource_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (x *Source) GetDataStream() *DataStream {
	if x, ok := x.GetSource().(*Source_DataStream); ok {
		return x.DataStream
	}
	return nil
}

func (x *Source) GetMessageQueue() *MessageQueue {
	if x, ok := x.GetSource().(*Source_MessageQueue); ok {
		return x.MessageQueue
	}
	return nil
}

type isSource_Source interface {
	isSource_Source()
}

type Source_DataStream struct {
	DataStream *DataStream `protobuf:"bytes,1,opt,name=data_stream,json=dataStream,proto3,oneof"`
}

type Source_MessageQueue struct {
	MessageQueue *MessageQueue `protobuf:"bytes,2,opt,name=message_queue,json=messageQueue,proto3,oneof"`
}

func (*Source_DataStream) isSource_Source() {}

func (*Source_MessageQueue) isSource_Source() {}

type DataStream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Stream database.
	// example: /ru-central1/aoegtvhtp8ob********/cc8004q4lbo6********
	Database string `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	// Stream name, absolute or relative.
	StreamName string `protobuf:"bytes,2,opt,name=stream_name,json=streamName,proto3" json:"stream_name,omitempty"`
	// Consumer name.
	Consumer string `protobuf:"bytes,3,opt,name=consumer,proto3" json:"consumer,omitempty"`
	// Service account which has read permission on the stream.
	ServiceAccountId string `protobuf:"bytes,4,opt,name=service_account_id,json=serviceAccountId,proto3" json:"service_account_id,omitempty"`
}

func (x *DataStream) Reset() {
	*x = DataStream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_serverless_eventrouter_v1_connector_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataStream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataStream) ProtoMessage() {}

func (x *DataStream) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_serverless_eventrouter_v1_connector_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataStream.ProtoReflect.Descriptor instead.
func (*DataStream) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_serverless_eventrouter_v1_connector_proto_rawDescGZIP(), []int{2}
}

func (x *DataStream) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *DataStream) GetStreamName() string {
	if x != nil {
		return x.StreamName
	}
	return ""
}

func (x *DataStream) GetConsumer() string {
	if x != nil {
		return x.Consumer
	}
	return ""
}

func (x *DataStream) GetServiceAccountId() string {
	if x != nil {
		return x.ServiceAccountId
	}
	return ""
}

type MessageQueue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Queue ARN.
	// Example: yrn:yc:ymq:ru-central1:aoe***:test
	QueueArn string `protobuf:"bytes,1,opt,name=queue_arn,json=queueArn,proto3" json:"queue_arn,omitempty"`
	// Service account which has read access to the queue.
	ServiceAccountId string `protobuf:"bytes,2,opt,name=service_account_id,json=serviceAccountId,proto3" json:"service_account_id,omitempty"`
	// Queue visibility timeout override.
	VisibilityTimeout *durationpb.Duration `protobuf:"bytes,3,opt,name=visibility_timeout,json=visibilityTimeout,proto3" json:"visibility_timeout,omitempty"`
	// Batch size for polling.
	BatchSize int64 `protobuf:"varint,4,opt,name=batch_size,json=batchSize,proto3" json:"batch_size,omitempty"`
	// Queue polling timeout.
	PollingTimeout *durationpb.Duration `protobuf:"bytes,5,opt,name=polling_timeout,json=pollingTimeout,proto3" json:"polling_timeout,omitempty"`
}

func (x *MessageQueue) Reset() {
	*x = MessageQueue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_serverless_eventrouter_v1_connector_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageQueue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageQueue) ProtoMessage() {}

func (x *MessageQueue) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_serverless_eventrouter_v1_connector_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageQueue.ProtoReflect.Descriptor instead.
func (*MessageQueue) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_serverless_eventrouter_v1_connector_proto_rawDescGZIP(), []int{3}
}

func (x *MessageQueue) GetQueueArn() string {
	if x != nil {
		return x.QueueArn
	}
	return ""
}

func (x *MessageQueue) GetServiceAccountId() string {
	if x != nil {
		return x.ServiceAccountId
	}
	return ""
}

func (x *MessageQueue) GetVisibilityTimeout() *durationpb.Duration {
	if x != nil {
		return x.VisibilityTimeout
	}
	return nil
}

func (x *MessageQueue) GetBatchSize() int64 {
	if x != nil {
		return x.BatchSize
	}
	return 0
}

func (x *MessageQueue) GetPollingTimeout() *durationpb.Duration {
	if x != nil {
		return x.PollingTimeout
	}
	return nil
}

var File_yandex_cloud_serverless_eventrouter_v1_connector_proto protoreflect.FileDescriptor

var file_yandex_cloud_serverless_eventrouter_v1_connector_proto_rawDesc = []byte{
	0x0a, 0x36, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73,
	0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xbb, 0x05, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x15,
	0x0a, 0x06, 0x62, 0x75, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x62, 0x75, 0x73, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x64, 0x12, 0x39, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x55,
	0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d,
	0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x46, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x2f, 0x0a,
	0x13, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x50,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x38,
	0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x80, 0x01, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x53,
	0x54, 0x4f, 0x50, 0x50, 0x45, 0x44, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x45, 0x53, 0x4f,
	0x55, 0x52, 0x43, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x03,
	0x12, 0x15, 0x0a, 0x11, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x44,
	0x45, 0x4e, 0x49, 0x45, 0x44, 0x10, 0x04, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x55, 0x42, 0x4a, 0x45,
	0x43, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x05, 0x22, 0xcc,
	0x01, 0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x0b, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32,
	0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x48, 0x00, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x12, 0x5b, 0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x71, 0x75, 0x65, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73,
	0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65, 0x48, 0x00, 0x52,
	0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65, 0x42, 0x0e, 0x0a,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x04, 0xc0, 0xc1, 0x31, 0x01, 0x22, 0xab, 0x01,
	0x0a, 0x0a, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x20, 0x0a, 0x08,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04,
	0xe8, 0xc7, 0x31, 0x01, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x25,
	0x0a, 0x0b, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x04, 0xe8, 0xc7, 0x31, 0x01, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe8, 0xc7, 0x31, 0x01, 0x52, 0x08, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x04, 0xe8, 0xc7, 0x31, 0x01, 0x52, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0xba, 0x02, 0x0a, 0x0c,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65, 0x12, 0x21, 0x0a, 0x09,
	0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x61, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x04, 0xe8, 0xc7, 0x31, 0x01, 0x52, 0x08, 0x71, 0x75, 0x65, 0x75, 0x65, 0x41, 0x72, 0x6e, 0x12,
	0x3a, 0x0a, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31,
	0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x53, 0x0a, 0x12, 0x76,
	0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x09, 0xfa, 0xc7, 0x31, 0x05, 0x3c, 0x3d, 0x31, 0x32, 0x68, 0x52, 0x11, 0x76,
	0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x12, 0x27, 0x0a, 0x0a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x08, 0xfa, 0xc7, 0x31, 0x04, 0x3c, 0x3d, 0x31, 0x30, 0x52, 0x09,
	0x62, 0x61, 0x74, 0x63, 0x68, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x4d, 0x0a, 0x0f, 0x70, 0x6f, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0xfa,
	0xc7, 0x31, 0x05, 0x3c, 0x3d, 0x32, 0x30, 0x73, 0x52, 0x0e, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x42, 0x8a, 0x01, 0x0a, 0x2a, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x04, 0x50, 0x45, 0x52, 0x43, 0x5a, 0x56, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_serverless_eventrouter_v1_connector_proto_rawDescOnce sync.Once
	file_yandex_cloud_serverless_eventrouter_v1_connector_proto_rawDescData = file_yandex_cloud_serverless_eventrouter_v1_connector_proto_rawDesc
)

func file_yandex_cloud_serverless_eventrouter_v1_connector_proto_rawDescGZIP() []byte {
	file_yandex_cloud_serverless_eventrouter_v1_connector_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_serverless_eventrouter_v1_connector_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_serverless_eventrouter_v1_connector_proto_rawDescData)
	})
	return file_yandex_cloud_serverless_eventrouter_v1_connector_proto_rawDescData
}

var file_yandex_cloud_serverless_eventrouter_v1_connector_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_serverless_eventrouter_v1_connector_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_yandex_cloud_serverless_eventrouter_v1_connector_proto_goTypes = []any{
	(Connector_Status)(0),         // 0: yandex.cloud.serverless.eventrouter.v1.Connector.Status
	(*Connector)(nil),             // 1: yandex.cloud.serverless.eventrouter.v1.Connector
	(*Source)(nil),                // 2: yandex.cloud.serverless.eventrouter.v1.Source
	(*DataStream)(nil),            // 3: yandex.cloud.serverless.eventrouter.v1.DataStream
	(*MessageQueue)(nil),          // 4: yandex.cloud.serverless.eventrouter.v1.MessageQueue
	nil,                           // 5: yandex.cloud.serverless.eventrouter.v1.Connector.LabelsEntry
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),   // 7: google.protobuf.Duration
}
var file_yandex_cloud_serverless_eventrouter_v1_connector_proto_depIdxs = []int32{
	6, // 0: yandex.cloud.serverless.eventrouter.v1.Connector.created_at:type_name -> google.protobuf.Timestamp
	5, // 1: yandex.cloud.serverless.eventrouter.v1.Connector.labels:type_name -> yandex.cloud.serverless.eventrouter.v1.Connector.LabelsEntry
	2, // 2: yandex.cloud.serverless.eventrouter.v1.Connector.source:type_name -> yandex.cloud.serverless.eventrouter.v1.Source
	0, // 3: yandex.cloud.serverless.eventrouter.v1.Connector.status:type_name -> yandex.cloud.serverless.eventrouter.v1.Connector.Status
	3, // 4: yandex.cloud.serverless.eventrouter.v1.Source.data_stream:type_name -> yandex.cloud.serverless.eventrouter.v1.DataStream
	4, // 5: yandex.cloud.serverless.eventrouter.v1.Source.message_queue:type_name -> yandex.cloud.serverless.eventrouter.v1.MessageQueue
	7, // 6: yandex.cloud.serverless.eventrouter.v1.MessageQueue.visibility_timeout:type_name -> google.protobuf.Duration
	7, // 7: yandex.cloud.serverless.eventrouter.v1.MessageQueue.polling_timeout:type_name -> google.protobuf.Duration
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_yandex_cloud_serverless_eventrouter_v1_connector_proto_init() }
func file_yandex_cloud_serverless_eventrouter_v1_connector_proto_init() {
	if File_yandex_cloud_serverless_eventrouter_v1_connector_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_serverless_eventrouter_v1_connector_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Connector); i {
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
		file_yandex_cloud_serverless_eventrouter_v1_connector_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Source); i {
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
		file_yandex_cloud_serverless_eventrouter_v1_connector_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*DataStream); i {
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
		file_yandex_cloud_serverless_eventrouter_v1_connector_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*MessageQueue); i {
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
	file_yandex_cloud_serverless_eventrouter_v1_connector_proto_msgTypes[1].OneofWrappers = []any{
		(*Source_DataStream)(nil),
		(*Source_MessageQueue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_serverless_eventrouter_v1_connector_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_serverless_eventrouter_v1_connector_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_serverless_eventrouter_v1_connector_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_serverless_eventrouter_v1_connector_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_serverless_eventrouter_v1_connector_proto_msgTypes,
	}.Build()
	File_yandex_cloud_serverless_eventrouter_v1_connector_proto = out.File
	file_yandex_cloud_serverless_eventrouter_v1_connector_proto_rawDesc = nil
	file_yandex_cloud_serverless_eventrouter_v1_connector_proto_goTypes = nil
	file_yandex_cloud_serverless_eventrouter_v1_connector_proto_depIdxs = nil
}
