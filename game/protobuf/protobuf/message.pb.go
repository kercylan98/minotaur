// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.1
// source: game/protobuf/message.proto

package protobuf

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

type MessageCode int32

const (
	MessageCode_SystemHeartbeat MessageCode = 0 // 心跳
)

// Enum value maps for MessageCode.
var (
	MessageCode_name = map[int32]string{
		0: "SystemHeartbeat",
	}
	MessageCode_value = map[string]int32{
		"SystemHeartbeat": 0,
	}
)

func (x MessageCode) Enum() *MessageCode {
	p := new(MessageCode)
	*p = x
	return p
}

func (x MessageCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageCode) Descriptor() protoreflect.EnumDescriptor {
	return file_game_protobuf_message_proto_enumTypes[0].Descriptor()
}

func (MessageCode) Type() protoreflect.EnumType {
	return &file_game_protobuf_message_proto_enumTypes[0]
}

func (x MessageCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageCode.Descriptor instead.
func (MessageCode) EnumDescriptor() ([]byte, []int) {
	return file_game_protobuf_message_proto_rawDescGZIP(), []int{0}
}

// 通用消息
type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_protobuf_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_game_protobuf_message_proto_msgTypes[0]
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
	return file_game_protobuf_message_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Message) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_game_protobuf_message_proto protoreflect.FileDescriptor

var file_game_protobuf_message_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x67, 0x61, 0x6d, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0x31, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x2a, 0x22, 0x0a, 0x0b, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x10, 0x00, 0x42, 0x0d,
	0x5a, 0x0b, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_game_protobuf_message_proto_rawDescOnce sync.Once
	file_game_protobuf_message_proto_rawDescData = file_game_protobuf_message_proto_rawDesc
)

func file_game_protobuf_message_proto_rawDescGZIP() []byte {
	file_game_protobuf_message_proto_rawDescOnce.Do(func() {
		file_game_protobuf_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_game_protobuf_message_proto_rawDescData)
	})
	return file_game_protobuf_message_proto_rawDescData
}

var file_game_protobuf_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_game_protobuf_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_game_protobuf_message_proto_goTypes = []interface{}{
	(MessageCode)(0), // 0: protobuf.MessageCode
	(*Message)(nil),  // 1: protobuf.Message
}
var file_game_protobuf_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_game_protobuf_message_proto_init() }
func file_game_protobuf_message_proto_init() {
	if File_game_protobuf_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_game_protobuf_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_game_protobuf_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_game_protobuf_message_proto_goTypes,
		DependencyIndexes: file_game_protobuf_message_proto_depIdxs,
		EnumInfos:         file_game_protobuf_message_proto_enumTypes,
		MessageInfos:      file_game_protobuf_message_proto_msgTypes,
	}.Build()
	File_game_protobuf_message_proto = out.File
	file_game_protobuf_message_proto_rawDesc = nil
	file_game_protobuf_message_proto_goTypes = nil
	file_game_protobuf_message_proto_depIdxs = nil
}
