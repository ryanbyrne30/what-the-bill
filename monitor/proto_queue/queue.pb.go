// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: queue.proto

package proto_queue

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

type USBillUpdatedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PackageId   string `protobuf:"bytes,1,opt,name=package_id,json=packageId,proto3" json:"package_id,omitempty"`
	UpdatedAt   string `protobuf:"bytes,2,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	PackageLink string `protobuf:"bytes,3,opt,name=package_link,json=packageLink,proto3" json:"package_link,omitempty"`
}

func (x *USBillUpdatedEvent) Reset() {
	*x = USBillUpdatedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queue_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *USBillUpdatedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*USBillUpdatedEvent) ProtoMessage() {}

func (x *USBillUpdatedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_queue_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use USBillUpdatedEvent.ProtoReflect.Descriptor instead.
func (*USBillUpdatedEvent) Descriptor() ([]byte, []int) {
	return file_queue_proto_rawDescGZIP(), []int{0}
}

func (x *USBillUpdatedEvent) GetPackageId() string {
	if x != nil {
		return x.PackageId
	}
	return ""
}

func (x *USBillUpdatedEvent) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *USBillUpdatedEvent) GetPackageLink() string {
	if x != nil {
		return x.PackageLink
	}
	return ""
}

type USBillCreatedEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PackageId   string `protobuf:"bytes,1,opt,name=package_id,json=packageId,proto3" json:"package_id,omitempty"`
	UpdatedAt   string `protobuf:"bytes,2,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	PackageLink string `protobuf:"bytes,3,opt,name=package_link,json=packageLink,proto3" json:"package_link,omitempty"`
}

func (x *USBillCreatedEvent) Reset() {
	*x = USBillCreatedEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queue_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *USBillCreatedEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*USBillCreatedEvent) ProtoMessage() {}

func (x *USBillCreatedEvent) ProtoReflect() protoreflect.Message {
	mi := &file_queue_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use USBillCreatedEvent.ProtoReflect.Descriptor instead.
func (*USBillCreatedEvent) Descriptor() ([]byte, []int) {
	return file_queue_proto_rawDescGZIP(), []int{1}
}

func (x *USBillCreatedEvent) GetPackageId() string {
	if x != nil {
		return x.PackageId
	}
	return ""
}

func (x *USBillCreatedEvent) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *USBillCreatedEvent) GetPackageLink() string {
	if x != nil {
		return x.PackageLink
	}
	return ""
}

type EventPosted struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *EventPosted) Reset() {
	*x = EventPosted{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queue_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventPosted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventPosted) ProtoMessage() {}

func (x *EventPosted) ProtoReflect() protoreflect.Message {
	mi := &file_queue_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventPosted.ProtoReflect.Descriptor instead.
func (*EventPosted) Descriptor() ([]byte, []int) {
	return file_queue_proto_rawDescGZIP(), []int{2}
}

func (x *EventPosted) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_queue_proto protoreflect.FileDescriptor

var file_queue_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x22, 0x75, 0x0a, 0x12, 0x55, 0x53,
	0x42, 0x69, 0x6c, 0x6c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x21,
	0x0a, 0x0c, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x6e,
	0x6b, 0x22, 0x75, 0x0a, 0x12, 0x55, 0x53, 0x42, 0x69, 0x6c, 0x6c, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x22, 0x25, 0x0a, 0x0b, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32,
	0xb5, 0x01, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x75, 0x65, 0x12, 0x55, 0x0a, 0x16, 0x50, 0x6f, 0x73,
	0x74, 0x55, 0x53, 0x42, 0x69, 0x6c, 0x6c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x71, 0x75, 0x65, 0x75,
	0x65, 0x2e, 0x55, 0x53, 0x42, 0x69, 0x6c, 0x6c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x64, 0x22, 0x00,
	0x12, 0x55, 0x0a, 0x16, 0x50, 0x6f, 0x73, 0x74, 0x55, 0x53, 0x42, 0x69, 0x6c, 0x6c, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x55, 0x53, 0x42, 0x69, 0x6c, 0x6c, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50,
	0x6f, 0x73, 0x74, 0x65, 0x64, 0x22, 0x00, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x79, 0x61, 0x6e, 0x62, 0x79, 0x72, 0x6e, 0x65, 0x33,
	0x30, 0x2f, 0x77, 0x68, 0x61, 0x74, 0x2d, 0x74, 0x68, 0x65, 0x2d, 0x62, 0x69, 0x6c, 0x6c, 0x2f,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x71, 0x75,
	0x65, 0x75, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_queue_proto_rawDescOnce sync.Once
	file_queue_proto_rawDescData = file_queue_proto_rawDesc
)

func file_queue_proto_rawDescGZIP() []byte {
	file_queue_proto_rawDescOnce.Do(func() {
		file_queue_proto_rawDescData = protoimpl.X.CompressGZIP(file_queue_proto_rawDescData)
	})
	return file_queue_proto_rawDescData
}

var file_queue_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_queue_proto_goTypes = []interface{}{
	(*USBillUpdatedEvent)(nil), // 0: proto_queue.USBillUpdatedEvent
	(*USBillCreatedEvent)(nil), // 1: proto_queue.USBillCreatedEvent
	(*EventPosted)(nil),        // 2: proto_queue.EventPosted
}
var file_queue_proto_depIdxs = []int32{
	0, // 0: proto_queue.Queue.PostUSBillUpdatedEvent:input_type -> proto_queue.USBillUpdatedEvent
	1, // 1: proto_queue.Queue.PostUSBillCreatedEvent:input_type -> proto_queue.USBillCreatedEvent
	2, // 2: proto_queue.Queue.PostUSBillUpdatedEvent:output_type -> proto_queue.EventPosted
	2, // 3: proto_queue.Queue.PostUSBillCreatedEvent:output_type -> proto_queue.EventPosted
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_queue_proto_init() }
func file_queue_proto_init() {
	if File_queue_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_queue_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*USBillUpdatedEvent); i {
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
		file_queue_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*USBillCreatedEvent); i {
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
		file_queue_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventPosted); i {
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
			RawDescriptor: file_queue_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_queue_proto_goTypes,
		DependencyIndexes: file_queue_proto_depIdxs,
		MessageInfos:      file_queue_proto_msgTypes,
	}.Build()
	File_queue_proto = out.File
	file_queue_proto_rawDesc = nil
	file_queue_proto_goTypes = nil
	file_queue_proto_depIdxs = nil
}
