// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/k8s/v1/maintenance.proto

package k8s

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	dayofweek "google.golang.org/genproto/googleapis/type/dayofweek"
	timeofday "google.golang.org/genproto/googleapis/type/timeofday"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MaintenanceWindow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Maintenance policy.
	//
	// Types that are assignable to Policy:
	//
	//	*MaintenanceWindow_Anytime
	//	*MaintenanceWindow_DailyMaintenanceWindow
	//	*MaintenanceWindow_WeeklyMaintenanceWindow
	Policy isMaintenanceWindow_Policy `protobuf_oneof:"policy"`
}

func (x *MaintenanceWindow) Reset() {
	*x = MaintenanceWindow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_k8s_v1_maintenance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaintenanceWindow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaintenanceWindow) ProtoMessage() {}

func (x *MaintenanceWindow) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_k8s_v1_maintenance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaintenanceWindow.ProtoReflect.Descriptor instead.
func (*MaintenanceWindow) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_k8s_v1_maintenance_proto_rawDescGZIP(), []int{0}
}

func (m *MaintenanceWindow) GetPolicy() isMaintenanceWindow_Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

func (x *MaintenanceWindow) GetAnytime() *AnytimeMaintenanceWindow {
	if x, ok := x.GetPolicy().(*MaintenanceWindow_Anytime); ok {
		return x.Anytime
	}
	return nil
}

func (x *MaintenanceWindow) GetDailyMaintenanceWindow() *DailyMaintenanceWindow {
	if x, ok := x.GetPolicy().(*MaintenanceWindow_DailyMaintenanceWindow); ok {
		return x.DailyMaintenanceWindow
	}
	return nil
}

func (x *MaintenanceWindow) GetWeeklyMaintenanceWindow() *WeeklyMaintenanceWindow {
	if x, ok := x.GetPolicy().(*MaintenanceWindow_WeeklyMaintenanceWindow); ok {
		return x.WeeklyMaintenanceWindow
	}
	return nil
}

type isMaintenanceWindow_Policy interface {
	isMaintenanceWindow_Policy()
}

type MaintenanceWindow_Anytime struct {
	// Updating the master at any time.
	Anytime *AnytimeMaintenanceWindow `protobuf:"bytes,1,opt,name=anytime,proto3,oneof"`
}

type MaintenanceWindow_DailyMaintenanceWindow struct {
	// Updating the master on any day during the specified time window.
	DailyMaintenanceWindow *DailyMaintenanceWindow `protobuf:"bytes,2,opt,name=daily_maintenance_window,json=dailyMaintenanceWindow,proto3,oneof"`
}

type MaintenanceWindow_WeeklyMaintenanceWindow struct {
	// Updating the master on selected days during the specified time window.
	WeeklyMaintenanceWindow *WeeklyMaintenanceWindow `protobuf:"bytes,3,opt,name=weekly_maintenance_window,json=weeklyMaintenanceWindow,proto3,oneof"`
}

func (*MaintenanceWindow_Anytime) isMaintenanceWindow_Policy() {}

func (*MaintenanceWindow_DailyMaintenanceWindow) isMaintenanceWindow_Policy() {}

func (*MaintenanceWindow_WeeklyMaintenanceWindow) isMaintenanceWindow_Policy() {}

type AnytimeMaintenanceWindow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AnytimeMaintenanceWindow) Reset() {
	*x = AnytimeMaintenanceWindow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_k8s_v1_maintenance_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnytimeMaintenanceWindow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnytimeMaintenanceWindow) ProtoMessage() {}

func (x *AnytimeMaintenanceWindow) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_k8s_v1_maintenance_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnytimeMaintenanceWindow.ProtoReflect.Descriptor instead.
func (*AnytimeMaintenanceWindow) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_k8s_v1_maintenance_proto_rawDescGZIP(), []int{1}
}

type DailyMaintenanceWindow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Window start time, in the UTC timezone.
	StartTime *timeofday.TimeOfDay `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// Window duration.
	Duration *durationpb.Duration `protobuf:"bytes,2,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *DailyMaintenanceWindow) Reset() {
	*x = DailyMaintenanceWindow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_k8s_v1_maintenance_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DailyMaintenanceWindow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DailyMaintenanceWindow) ProtoMessage() {}

func (x *DailyMaintenanceWindow) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_k8s_v1_maintenance_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DailyMaintenanceWindow.ProtoReflect.Descriptor instead.
func (*DailyMaintenanceWindow) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_k8s_v1_maintenance_proto_rawDescGZIP(), []int{2}
}

func (x *DailyMaintenanceWindow) GetStartTime() *timeofday.TimeOfDay {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *DailyMaintenanceWindow) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

type DaysOfWeekMaintenanceWindow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Days of the week when automatic updates are allowed.
	Days []dayofweek.DayOfWeek `protobuf:"varint,1,rep,packed,name=days,proto3,enum=google.type.DayOfWeek" json:"days,omitempty"`
	// Window start time, in the UTC timezone.
	StartTime *timeofday.TimeOfDay `protobuf:"bytes,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// Window duration.
	Duration *durationpb.Duration `protobuf:"bytes,3,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *DaysOfWeekMaintenanceWindow) Reset() {
	*x = DaysOfWeekMaintenanceWindow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_k8s_v1_maintenance_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DaysOfWeekMaintenanceWindow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DaysOfWeekMaintenanceWindow) ProtoMessage() {}

func (x *DaysOfWeekMaintenanceWindow) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_k8s_v1_maintenance_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DaysOfWeekMaintenanceWindow.ProtoReflect.Descriptor instead.
func (*DaysOfWeekMaintenanceWindow) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_k8s_v1_maintenance_proto_rawDescGZIP(), []int{3}
}

func (x *DaysOfWeekMaintenanceWindow) GetDays() []dayofweek.DayOfWeek {
	if x != nil {
		return x.Days
	}
	return nil
}

func (x *DaysOfWeekMaintenanceWindow) GetStartTime() *timeofday.TimeOfDay {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *DaysOfWeekMaintenanceWindow) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

type WeeklyMaintenanceWindow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Days of the week and the maintenance window for these days when automatic updates are allowed.
	DaysOfWeek []*DaysOfWeekMaintenanceWindow `protobuf:"bytes,1,rep,name=days_of_week,json=daysOfWeek,proto3" json:"days_of_week,omitempty"`
}

func (x *WeeklyMaintenanceWindow) Reset() {
	*x = WeeklyMaintenanceWindow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_k8s_v1_maintenance_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeeklyMaintenanceWindow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeeklyMaintenanceWindow) ProtoMessage() {}

func (x *WeeklyMaintenanceWindow) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_k8s_v1_maintenance_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeeklyMaintenanceWindow.ProtoReflect.Descriptor instead.
func (*WeeklyMaintenanceWindow) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_k8s_v1_maintenance_proto_rawDescGZIP(), []int{4}
}

func (x *WeeklyMaintenanceWindow) GetDaysOfWeek() []*DaysOfWeekMaintenanceWindow {
	if x != nil {
		return x.DaysOfWeek
	}
	return nil
}

var File_yandex_cloud_k8s_v1_maintenance_proto protoreflect.FileDescriptor

var file_yandex_cloud_k8s_v1_maintenance_proto_rawDesc = []byte{
	0x0a, 0x25, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6b,
	0x38, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x64, 0x61, 0x79, 0x6f, 0x66, 0x77,
	0x65, 0x65, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x66, 0x64, 0x61, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc3, 0x02, 0x0a, 0x11, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x63, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x12, 0x49, 0x0a, 0x07, 0x61,
	0x6e, 0x79, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6b, 0x38, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x6e, 0x79, 0x74, 0x69, 0x6d, 0x65, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x63, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x48, 0x00, 0x52, 0x07, 0x61,
	0x6e, 0x79, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x67, 0x0a, 0x18, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x5f,
	0x6d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x77, 0x69, 0x6e, 0x64,
	0x6f, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x61, 0x69, 0x6c, 0x79, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x57,
	0x69, 0x6e, 0x64, 0x6f, 0x77, 0x48, 0x00, 0x52, 0x16, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x4d, 0x61,
	0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x12,
	0x6a, 0x0a, 0x19, 0x77, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x5f, 0x6d, 0x61, 0x69, 0x6e, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x4d,
	0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77,
	0x48, 0x00, 0x52, 0x17, 0x77, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x63, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x42, 0x0e, 0x0a, 0x06, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x04, 0xc0, 0xc1, 0x31, 0x01, 0x22, 0x1a, 0x0a, 0x18, 0x41,
	0x6e, 0x79, 0x74, 0x69, 0x6d, 0x65, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63,
	0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x22, 0x98, 0x01, 0x0a, 0x16, 0x44, 0x61, 0x69, 0x6c,
	0x79, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x57, 0x69, 0x6e, 0x64,
	0x6f, 0x77, 0x12, 0x3b, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x4f, 0x66, 0x44, 0x61, 0x79, 0x42, 0x04,
	0xe8, 0xc7, 0x31, 0x01, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x41, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0a, 0xfa, 0xc7,
	0x31, 0x06, 0x31, 0x68, 0x2d, 0x32, 0x34, 0x68, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0xd2, 0x01, 0x0a, 0x1b, 0x44, 0x61, 0x79, 0x73, 0x4f, 0x66, 0x57, 0x65, 0x65,
	0x6b, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x57, 0x69, 0x6e, 0x64,
	0x6f, 0x77, 0x12, 0x33, 0x0a, 0x04, 0x64, 0x61, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e,
	0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44,
	0x61, 0x79, 0x4f, 0x66, 0x57, 0x65, 0x65, 0x6b, 0x42, 0x07, 0x82, 0xc8, 0x31, 0x03, 0x31, 0x2d,
	0x37, 0x52, 0x04, 0x64, 0x61, 0x79, 0x73, 0x12, 0x3b, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x4f, 0x66,
	0x44, 0x61, 0x79, 0x42, 0x04, 0xe8, 0xc7, 0x31, 0x01, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x0a, 0xfa, 0xc7, 0x31, 0x06, 0x31, 0x68, 0x2d, 0x32, 0x34, 0x68, 0x52, 0x08, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x76, 0x0a, 0x17, 0x57, 0x65, 0x65, 0x6b, 0x6c,
	0x79, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x57, 0x69, 0x6e, 0x64,
	0x6f, 0x77, 0x12, 0x5b, 0x0a, 0x0c, 0x64, 0x61, 0x79, 0x73, 0x5f, 0x6f, 0x66, 0x5f, 0x77, 0x65,
	0x65, 0x6b, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x61, 0x79, 0x73, 0x4f, 0x66, 0x57, 0x65, 0x65, 0x6b, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x63, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x42, 0x07, 0x82, 0xc8, 0x31, 0x03,
	0x31, 0x2d, 0x37, 0x52, 0x0a, 0x64, 0x61, 0x79, 0x73, 0x4f, 0x66, 0x57, 0x65, 0x65, 0x6b, 0x42,
	0x56, 0x0a, 0x17, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6b, 0x38, 0x73,
	0x2f, 0x76, 0x31, 0x3b, 0x6b, 0x38, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_k8s_v1_maintenance_proto_rawDescOnce sync.Once
	file_yandex_cloud_k8s_v1_maintenance_proto_rawDescData = file_yandex_cloud_k8s_v1_maintenance_proto_rawDesc
)

func file_yandex_cloud_k8s_v1_maintenance_proto_rawDescGZIP() []byte {
	file_yandex_cloud_k8s_v1_maintenance_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_k8s_v1_maintenance_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_k8s_v1_maintenance_proto_rawDescData)
	})
	return file_yandex_cloud_k8s_v1_maintenance_proto_rawDescData
}

var file_yandex_cloud_k8s_v1_maintenance_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_yandex_cloud_k8s_v1_maintenance_proto_goTypes = []any{
	(*MaintenanceWindow)(nil),           // 0: yandex.cloud.k8s.v1.MaintenanceWindow
	(*AnytimeMaintenanceWindow)(nil),    // 1: yandex.cloud.k8s.v1.AnytimeMaintenanceWindow
	(*DailyMaintenanceWindow)(nil),      // 2: yandex.cloud.k8s.v1.DailyMaintenanceWindow
	(*DaysOfWeekMaintenanceWindow)(nil), // 3: yandex.cloud.k8s.v1.DaysOfWeekMaintenanceWindow
	(*WeeklyMaintenanceWindow)(nil),     // 4: yandex.cloud.k8s.v1.WeeklyMaintenanceWindow
	(*timeofday.TimeOfDay)(nil),         // 5: google.type.TimeOfDay
	(*durationpb.Duration)(nil),         // 6: google.protobuf.Duration
	(dayofweek.DayOfWeek)(0),            // 7: google.type.DayOfWeek
}
var file_yandex_cloud_k8s_v1_maintenance_proto_depIdxs = []int32{
	1, // 0: yandex.cloud.k8s.v1.MaintenanceWindow.anytime:type_name -> yandex.cloud.k8s.v1.AnytimeMaintenanceWindow
	2, // 1: yandex.cloud.k8s.v1.MaintenanceWindow.daily_maintenance_window:type_name -> yandex.cloud.k8s.v1.DailyMaintenanceWindow
	4, // 2: yandex.cloud.k8s.v1.MaintenanceWindow.weekly_maintenance_window:type_name -> yandex.cloud.k8s.v1.WeeklyMaintenanceWindow
	5, // 3: yandex.cloud.k8s.v1.DailyMaintenanceWindow.start_time:type_name -> google.type.TimeOfDay
	6, // 4: yandex.cloud.k8s.v1.DailyMaintenanceWindow.duration:type_name -> google.protobuf.Duration
	7, // 5: yandex.cloud.k8s.v1.DaysOfWeekMaintenanceWindow.days:type_name -> google.type.DayOfWeek
	5, // 6: yandex.cloud.k8s.v1.DaysOfWeekMaintenanceWindow.start_time:type_name -> google.type.TimeOfDay
	6, // 7: yandex.cloud.k8s.v1.DaysOfWeekMaintenanceWindow.duration:type_name -> google.protobuf.Duration
	3, // 8: yandex.cloud.k8s.v1.WeeklyMaintenanceWindow.days_of_week:type_name -> yandex.cloud.k8s.v1.DaysOfWeekMaintenanceWindow
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_yandex_cloud_k8s_v1_maintenance_proto_init() }
func file_yandex_cloud_k8s_v1_maintenance_proto_init() {
	if File_yandex_cloud_k8s_v1_maintenance_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_k8s_v1_maintenance_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*MaintenanceWindow); i {
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
		file_yandex_cloud_k8s_v1_maintenance_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AnytimeMaintenanceWindow); i {
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
		file_yandex_cloud_k8s_v1_maintenance_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*DailyMaintenanceWindow); i {
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
		file_yandex_cloud_k8s_v1_maintenance_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*DaysOfWeekMaintenanceWindow); i {
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
		file_yandex_cloud_k8s_v1_maintenance_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*WeeklyMaintenanceWindow); i {
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
	file_yandex_cloud_k8s_v1_maintenance_proto_msgTypes[0].OneofWrappers = []any{
		(*MaintenanceWindow_Anytime)(nil),
		(*MaintenanceWindow_DailyMaintenanceWindow)(nil),
		(*MaintenanceWindow_WeeklyMaintenanceWindow)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_k8s_v1_maintenance_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_k8s_v1_maintenance_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_k8s_v1_maintenance_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_k8s_v1_maintenance_proto_msgTypes,
	}.Build()
	File_yandex_cloud_k8s_v1_maintenance_proto = out.File
	file_yandex_cloud_k8s_v1_maintenance_proto_rawDesc = nil
	file_yandex_cloud_k8s_v1_maintenance_proto_goTypes = nil
	file_yandex_cloud_k8s_v1_maintenance_proto_depIdxs = nil
}
