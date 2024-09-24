// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/ai/foundation_models/v1/text_common.proto

package foundation_models

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Enum representing the generation status of the alternative.
type Alternative_AlternativeStatus int32

const (
	// Unspecified generation status.
	Alternative_ALTERNATIVE_STATUS_UNSPECIFIED Alternative_AlternativeStatus = 0
	// Partially generated alternative.
	Alternative_ALTERNATIVE_STATUS_PARTIAL Alternative_AlternativeStatus = 1
	// Incomplete final alternative resulting from reaching the maximum allowed number of tokens.
	Alternative_ALTERNATIVE_STATUS_TRUNCATED_FINAL Alternative_AlternativeStatus = 2
	// Final alternative generated without running into any limits.
	Alternative_ALTERNATIVE_STATUS_FINAL Alternative_AlternativeStatus = 3
	// Generation was stopped due to the discovery of potentially sensitive content in the prompt or generated response.
	// To fix, modify the prompt and restart generation.
	Alternative_ALTERNATIVE_STATUS_CONTENT_FILTER Alternative_AlternativeStatus = 4
)

// Enum value maps for Alternative_AlternativeStatus.
var (
	Alternative_AlternativeStatus_name = map[int32]string{
		0: "ALTERNATIVE_STATUS_UNSPECIFIED",
		1: "ALTERNATIVE_STATUS_PARTIAL",
		2: "ALTERNATIVE_STATUS_TRUNCATED_FINAL",
		3: "ALTERNATIVE_STATUS_FINAL",
		4: "ALTERNATIVE_STATUS_CONTENT_FILTER",
	}
	Alternative_AlternativeStatus_value = map[string]int32{
		"ALTERNATIVE_STATUS_UNSPECIFIED":     0,
		"ALTERNATIVE_STATUS_PARTIAL":         1,
		"ALTERNATIVE_STATUS_TRUNCATED_FINAL": 2,
		"ALTERNATIVE_STATUS_FINAL":           3,
		"ALTERNATIVE_STATUS_CONTENT_FILTER":  4,
	}
)

func (x Alternative_AlternativeStatus) Enum() *Alternative_AlternativeStatus {
	p := new(Alternative_AlternativeStatus)
	*p = x
	return p
}

func (x Alternative_AlternativeStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Alternative_AlternativeStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_ai_foundation_models_v1_text_common_proto_enumTypes[0].Descriptor()
}

func (Alternative_AlternativeStatus) Type() protoreflect.EnumType {
	return &file_yandex_cloud_ai_foundation_models_v1_text_common_proto_enumTypes[0]
}

func (x Alternative_AlternativeStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Alternative_AlternativeStatus.Descriptor instead.
func (Alternative_AlternativeStatus) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_foundation_models_v1_text_common_proto_rawDescGZIP(), []int{3, 0}
}

// Defines the options for completion generation.
type CompletionOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Enables streaming of partially generated text.
	Stream bool `protobuf:"varint,1,opt,name=stream,proto3" json:"stream,omitempty"`
	// Affects creativity and randomness of responses. Should be a double number between 0 (inclusive) and 1 (inclusive).
	// Lower values produce more straightforward responses while higher values lead to increased creativity and randomness.
	// Default temperature: 0.3
	Temperature *wrapperspb.DoubleValue `protobuf:"bytes,2,opt,name=temperature,proto3" json:"temperature,omitempty"`
	// The limit on the number of tokens used for single completion generation.
	// Must be greater than zero. This maximum allowed parameter value may depend on the model being used.
	MaxTokens *wrapperspb.Int64Value `protobuf:"bytes,3,opt,name=max_tokens,json=maxTokens,proto3" json:"max_tokens,omitempty"`
}

func (x *CompletionOptions) Reset() {
	*x = CompletionOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ai_foundation_models_v1_text_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompletionOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompletionOptions) ProtoMessage() {}

func (x *CompletionOptions) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_foundation_models_v1_text_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompletionOptions.ProtoReflect.Descriptor instead.
func (*CompletionOptions) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_foundation_models_v1_text_common_proto_rawDescGZIP(), []int{0}
}

func (x *CompletionOptions) GetStream() bool {
	if x != nil {
		return x.Stream
	}
	return false
}

func (x *CompletionOptions) GetTemperature() *wrapperspb.DoubleValue {
	if x != nil {
		return x.Temperature
	}
	return nil
}

func (x *CompletionOptions) GetMaxTokens() *wrapperspb.Int64Value {
	if x != nil {
		return x.MaxTokens
	}
	return nil
}

// A message object representing a wrapper over the inputs and outputs of the completion model.
type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the message sender. Supported roles:
	// * `system`: Special role used to define the behaviour of the completion model.
	// * `assistant`: A role used by the model to generate responses.
	// * `user`: A role used by the user to describe requests to the model.
	Role string `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	// Message content.
	//
	// Types that are assignable to Content:
	//
	//	*Message_Text
	Content isMessage_Content `protobuf_oneof:"Content"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ai_foundation_models_v1_text_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_foundation_models_v1_text_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_foundation_models_v1_text_common_proto_rawDescGZIP(), []int{1}
}

func (x *Message) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (m *Message) GetContent() isMessage_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *Message) GetText() string {
	if x, ok := x.GetContent().(*Message_Text); ok {
		return x.Text
	}
	return ""
}

type isMessage_Content interface {
	isMessage_Content()
}

type Message_Text struct {
	// Textual content of the message.
	Text string `protobuf:"bytes,2,opt,name=text,proto3,oneof"`
}

func (*Message_Text) isMessage_Content() {}

// An object representing the number of content [tokens](/docs/foundation-models/concepts/yandexgpt/tokens) used by the completion model.
type ContentUsage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The number of tokens in the textual part of the model input.
	InputTextTokens int64 `protobuf:"varint,1,opt,name=input_text_tokens,json=inputTextTokens,proto3" json:"input_text_tokens,omitempty"`
	// The total number of tokens in the generated completions.
	CompletionTokens int64 `protobuf:"varint,2,opt,name=completion_tokens,json=completionTokens,proto3" json:"completion_tokens,omitempty"`
	// The total number of tokens, including all input tokens and all generated tokens.
	TotalTokens int64 `protobuf:"varint,3,opt,name=total_tokens,json=totalTokens,proto3" json:"total_tokens,omitempty"`
}

func (x *ContentUsage) Reset() {
	*x = ContentUsage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ai_foundation_models_v1_text_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentUsage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentUsage) ProtoMessage() {}

func (x *ContentUsage) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_foundation_models_v1_text_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentUsage.ProtoReflect.Descriptor instead.
func (*ContentUsage) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_foundation_models_v1_text_common_proto_rawDescGZIP(), []int{2}
}

func (x *ContentUsage) GetInputTextTokens() int64 {
	if x != nil {
		return x.InputTextTokens
	}
	return 0
}

func (x *ContentUsage) GetCompletionTokens() int64 {
	if x != nil {
		return x.CompletionTokens
	}
	return 0
}

func (x *ContentUsage) GetTotalTokens() int64 {
	if x != nil {
		return x.TotalTokens
	}
	return 0
}

// Represents a generated completion alternative, including its content and generation status.
type Alternative struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A message containing the content of the alternative.
	Message *Message `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	// The generation status of the alternative
	Status Alternative_AlternativeStatus `protobuf:"varint,2,opt,name=status,proto3,enum=yandex.cloud.ai.foundation_models.v1.Alternative_AlternativeStatus" json:"status,omitempty"`
}

func (x *Alternative) Reset() {
	*x = Alternative{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ai_foundation_models_v1_text_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Alternative) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Alternative) ProtoMessage() {}

func (x *Alternative) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_foundation_models_v1_text_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Alternative.ProtoReflect.Descriptor instead.
func (*Alternative) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_foundation_models_v1_text_common_proto_rawDescGZIP(), []int{3}
}

func (x *Alternative) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *Alternative) GetStatus() Alternative_AlternativeStatus {
	if x != nil {
		return x.Status
	}
	return Alternative_ALTERNATIVE_STATUS_UNSPECIFIED
}

// Represents a token, the basic unit of content, used by the foundation model.
type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An internal token identifier.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// The textual representation of the token.
	Text string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	// Indicates whether the token is special or not. Special tokens may define the model's behavior and are not visible to users.
	Special bool `protobuf:"varint,3,opt,name=special,proto3" json:"special,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ai_foundation_models_v1_text_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_foundation_models_v1_text_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_foundation_models_v1_text_common_proto_rawDescGZIP(), []int{4}
}

func (x *Token) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Token) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Token) GetSpecial() bool {
	if x != nil {
		return x.Special
	}
	return false
}

var File_yandex_cloud_ai_foundation_models_v1_text_common_proto protoreflect.FileDescriptor

var file_yandex_cloud_ai_foundation_models_v1_text_common_proto_rawDesc = []byte{
	0x0a, 0x36, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x2f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa7,
	0x01, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x3e, 0x0a, 0x0b,
	0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x3a, 0x0a, 0x0a,
	0x6d, 0x61, 0x78, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x6d,
	0x61, 0x78, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x22, 0x3e, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x42, 0x09, 0x0a,
	0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x8a, 0x01, 0x0a, 0x0c, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x54, 0x65, 0x78, 0x74, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x22, 0xfa, 0x02, 0x0a, 0x0b, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x74, 0x69, 0x76, 0x65, 0x12, 0x47, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x5b,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x43,
	0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69,
	0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x69, 0x76,
	0x65, 0x2e, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xc4, 0x01, 0x0a, 0x11,
	0x41, 0x6c, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x22, 0x0a, 0x1e, 0x41, 0x4c, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x54, 0x49, 0x56, 0x45,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x41, 0x4c, 0x54, 0x45, 0x52, 0x4e, 0x41,
	0x54, 0x49, 0x56, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x41, 0x52, 0x54,
	0x49, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x26, 0x0a, 0x22, 0x41, 0x4c, 0x54, 0x45, 0x52, 0x4e, 0x41,
	0x54, 0x49, 0x56, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x54, 0x52, 0x55, 0x4e,
	0x43, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x10, 0x02, 0x12, 0x1c, 0x0a,
	0x18, 0x41, 0x4c, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x10, 0x03, 0x12, 0x25, 0x0a, 0x21, 0x41,
	0x4c, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x5f, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52,
	0x10, 0x04, 0x22, 0x45, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x42, 0x86, 0x01, 0x0a, 0x28, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x69, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x5a, 0x5a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x69, 0x2f, 0x66, 0x6f, 0x75, 0x6e,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x76, 0x31,
	0x3b, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_ai_foundation_models_v1_text_common_proto_rawDescOnce sync.Once
	file_yandex_cloud_ai_foundation_models_v1_text_common_proto_rawDescData = file_yandex_cloud_ai_foundation_models_v1_text_common_proto_rawDesc
)

func file_yandex_cloud_ai_foundation_models_v1_text_common_proto_rawDescGZIP() []byte {
	file_yandex_cloud_ai_foundation_models_v1_text_common_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_ai_foundation_models_v1_text_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_ai_foundation_models_v1_text_common_proto_rawDescData)
	})
	return file_yandex_cloud_ai_foundation_models_v1_text_common_proto_rawDescData
}

var file_yandex_cloud_ai_foundation_models_v1_text_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_ai_foundation_models_v1_text_common_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_yandex_cloud_ai_foundation_models_v1_text_common_proto_goTypes = []any{
	(Alternative_AlternativeStatus)(0), // 0: yandex.cloud.ai.foundation_models.v1.Alternative.AlternativeStatus
	(*CompletionOptions)(nil),          // 1: yandex.cloud.ai.foundation_models.v1.CompletionOptions
	(*Message)(nil),                    // 2: yandex.cloud.ai.foundation_models.v1.Message
	(*ContentUsage)(nil),               // 3: yandex.cloud.ai.foundation_models.v1.ContentUsage
	(*Alternative)(nil),                // 4: yandex.cloud.ai.foundation_models.v1.Alternative
	(*Token)(nil),                      // 5: yandex.cloud.ai.foundation_models.v1.Token
	(*wrapperspb.DoubleValue)(nil),     // 6: google.protobuf.DoubleValue
	(*wrapperspb.Int64Value)(nil),      // 7: google.protobuf.Int64Value
}
var file_yandex_cloud_ai_foundation_models_v1_text_common_proto_depIdxs = []int32{
	6, // 0: yandex.cloud.ai.foundation_models.v1.CompletionOptions.temperature:type_name -> google.protobuf.DoubleValue
	7, // 1: yandex.cloud.ai.foundation_models.v1.CompletionOptions.max_tokens:type_name -> google.protobuf.Int64Value
	2, // 2: yandex.cloud.ai.foundation_models.v1.Alternative.message:type_name -> yandex.cloud.ai.foundation_models.v1.Message
	0, // 3: yandex.cloud.ai.foundation_models.v1.Alternative.status:type_name -> yandex.cloud.ai.foundation_models.v1.Alternative.AlternativeStatus
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_yandex_cloud_ai_foundation_models_v1_text_common_proto_init() }
func file_yandex_cloud_ai_foundation_models_v1_text_common_proto_init() {
	if File_yandex_cloud_ai_foundation_models_v1_text_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_ai_foundation_models_v1_text_common_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CompletionOptions); i {
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
		file_yandex_cloud_ai_foundation_models_v1_text_common_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Message); i {
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
		file_yandex_cloud_ai_foundation_models_v1_text_common_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ContentUsage); i {
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
		file_yandex_cloud_ai_foundation_models_v1_text_common_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Alternative); i {
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
		file_yandex_cloud_ai_foundation_models_v1_text_common_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Token); i {
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
	file_yandex_cloud_ai_foundation_models_v1_text_common_proto_msgTypes[1].OneofWrappers = []any{
		(*Message_Text)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_ai_foundation_models_v1_text_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_ai_foundation_models_v1_text_common_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_ai_foundation_models_v1_text_common_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_ai_foundation_models_v1_text_common_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_ai_foundation_models_v1_text_common_proto_msgTypes,
	}.Build()
	File_yandex_cloud_ai_foundation_models_v1_text_common_proto = out.File
	file_yandex_cloud_ai_foundation_models_v1_text_common_proto_rawDesc = nil
	file_yandex_cloud_ai_foundation_models_v1_text_common_proto_goTypes = nil
	file_yandex_cloud_ai_foundation_models_v1_text_common_proto_depIdxs = nil
}
