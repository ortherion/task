// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: analytics_messaging/message_service.proto

package analytics_messaging

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type EventType int32

const (
	EventType_UNKNOWN     EventType = 0
	EventType_CREATED     EventType = 1
	EventType_SENT_TO     EventType = 2
	EventType_APPROVED_BY EventType = 3
	EventType_REJECTED_BY EventType = 4
	EventType_SIGNED      EventType = 5
	EventType_SENT        EventType = 6
)

// Enum value maps for EventType.
var (
	EventType_name = map[int32]string{
		0: "UNKNOWN",
		1: "CREATED",
		2: "SENT_TO",
		3: "APPROVED_BY",
		4: "REJECTED_BY",
		5: "SIGNED",
		6: "SENT",
	}
	EventType_value = map[string]int32{
		"UNKNOWN":     0,
		"CREATED":     1,
		"SENT_TO":     2,
		"APPROVED_BY": 3,
		"REJECTED_BY": 4,
		"SIGNED":      5,
		"SENT":        6,
	}
)

func (x EventType) Enum() *EventType {
	p := new(EventType)
	*p = x
	return p
}

func (x EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_analytics_messaging_message_service_proto_enumTypes[0].Descriptor()
}

func (EventType) Type() protoreflect.EnumType {
	return &file_analytics_messaging_message_service_proto_enumTypes[0]
}

func (x EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventType.Descriptor instead.
func (EventType) EnumDescriptor() ([]byte, []int) {
	return file_analytics_messaging_message_service_proto_rawDescGZIP(), []int{0}
}

type EventMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskUuid  string                 `protobuf:"bytes,1,opt,name=task_uuid,json=taskUuid,proto3" json:"task_uuid,omitempty"`
	EventType EventType              `protobuf:"varint,2,opt,name=event_type,json=eventType,proto3,enum=analytics_messaging.EventType" json:"event_type,omitempty"`
	UserUuid  string                 `protobuf:"bytes,3,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *EventMessage) Reset() {
	*x = EventMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analytics_messaging_message_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventMessage) ProtoMessage() {}

func (x *EventMessage) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_messaging_message_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventMessage.ProtoReflect.Descriptor instead.
func (*EventMessage) Descriptor() ([]byte, []int) {
	return file_analytics_messaging_message_service_proto_rawDescGZIP(), []int{0}
}

func (x *EventMessage) GetTaskUuid() string {
	if x != nil {
		return x.TaskUuid
	}
	return ""
}

func (x *EventMessage) GetEventType() EventType {
	if x != nil {
		return x.EventType
	}
	return EventType_UNKNOWN
}

func (x *EventMessage) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

func (x *EventMessage) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analytics_messaging_message_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_messaging_message_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_analytics_messaging_message_service_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_analytics_messaging_message_service_proto protoreflect.FileDescriptor

var file_analytics_messaging_message_service_proto_rawDesc = []byte{
	0x0a, 0x29, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc1, 0x01, 0x0a, 0x0c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x55, 0x75, 0x69, 0x64, 0x12,
	0x3d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x24, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2a, 0x6a, 0x0a, 0x09, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44,
	0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x4f, 0x10, 0x02, 0x12,
	0x0f, 0x0a, 0x0b, 0x41, 0x50, 0x50, 0x52, 0x4f, 0x56, 0x45, 0x44, 0x5f, 0x42, 0x59, 0x10, 0x03,
	0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x5f, 0x42, 0x59, 0x10,
	0x04, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x49, 0x47, 0x4e, 0x45, 0x44, 0x10, 0x05, 0x12, 0x08, 0x0a,
	0x04, 0x53, 0x45, 0x4e, 0x54, 0x10, 0x06, 0x32, 0x68, 0x0a, 0x13, 0x41, 0x6e, 0x61, 0x6c, 0x79,
	0x74, 0x69, 0x63, 0x73, 0x4d, 0x73, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51,
	0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x2e,
	0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x69, 0x6e, 0x67, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x1a, 0x1d, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x36, 0x38, 0x33, 0x34, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x31, 0x37, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x69, 0x6e, 0x67, 0x3b, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_analytics_messaging_message_service_proto_rawDescOnce sync.Once
	file_analytics_messaging_message_service_proto_rawDescData = file_analytics_messaging_message_service_proto_rawDesc
)

func file_analytics_messaging_message_service_proto_rawDescGZIP() []byte {
	file_analytics_messaging_message_service_proto_rawDescOnce.Do(func() {
		file_analytics_messaging_message_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_analytics_messaging_message_service_proto_rawDescData)
	})
	return file_analytics_messaging_message_service_proto_rawDescData
}

var file_analytics_messaging_message_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_analytics_messaging_message_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_analytics_messaging_message_service_proto_goTypes = []interface{}{
	(EventType)(0),                // 0: analytics_messaging.EventType
	(*EventMessage)(nil),          // 1: analytics_messaging.EventMessage
	(*Response)(nil),              // 2: analytics_messaging.Response
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_analytics_messaging_message_service_proto_depIdxs = []int32{
	0, // 0: analytics_messaging.EventMessage.event_type:type_name -> analytics_messaging.EventType
	3, // 1: analytics_messaging.EventMessage.timestamp:type_name -> google.protobuf.Timestamp
	1, // 2: analytics_messaging.AnalyticsMsgService.SendMessage:input_type -> analytics_messaging.EventMessage
	2, // 3: analytics_messaging.AnalyticsMsgService.SendMessage:output_type -> analytics_messaging.Response
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_analytics_messaging_message_service_proto_init() }
func file_analytics_messaging_message_service_proto_init() {
	if File_analytics_messaging_message_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_analytics_messaging_message_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventMessage); i {
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
		file_analytics_messaging_message_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_analytics_messaging_message_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_analytics_messaging_message_service_proto_goTypes,
		DependencyIndexes: file_analytics_messaging_message_service_proto_depIdxs,
		EnumInfos:         file_analytics_messaging_message_service_proto_enumTypes,
		MessageInfos:      file_analytics_messaging_message_service_proto_msgTypes,
	}.Build()
	File_analytics_messaging_message_service_proto = out.File
	file_analytics_messaging_message_service_proto_rawDesc = nil
	file_analytics_messaging_message_service_proto_goTypes = nil
	file_analytics_messaging_message_service_proto_depIdxs = nil
}
