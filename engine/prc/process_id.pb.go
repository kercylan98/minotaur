// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.28.0--rc1
// source: process_id.proto

package prc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	"sync/atomic"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProcessId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogicalAddress  string `protobuf:"bytes,1,opt,name=logical_address,json=logicalAddress,proto3" json:"logical_address,omitempty"`
	PhysicalAddress string `protobuf:"bytes,2,opt,name=physical_address,json=physicalAddress,proto3" json:"physical_address,omitempty"`
	cache           atomic.Pointer[Process]
}

func (x *ProcessId) Reset() {
	*x = ProcessId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_process_id_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessId) ProtoMessage() {}

func (x *ProcessId) ProtoReflect() protoreflect.Message {
	mi := &file_process_id_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessId.ProtoReflect.Descriptor instead.
func (*ProcessId) Descriptor() ([]byte, []int) {
	return file_process_id_proto_rawDescGZIP(), []int{0}
}

var File_process_id_proto protoreflect.FileDescriptor

var file_process_id_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x70, 0x72, 0x63, 0x22, 0x5f, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c,
	0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x29, 0x0a,
	0x10, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x68, 0x79, 0x73, 0x69, 0x63, 0x61,
	0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x65, 0x72, 0x63, 0x79, 0x6c, 0x61, 0x6e, 0x39,
	0x38, 0x2f, 0x6d, 0x69, 0x6e, 0x6f, 0x74, 0x61, 0x75, 0x72, 0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2f, 0x70, 0x72, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_process_id_proto_rawDescOnce sync.Once
	file_process_id_proto_rawDescData = file_process_id_proto_rawDesc
)

func file_process_id_proto_rawDescGZIP() []byte {
	file_process_id_proto_rawDescOnce.Do(func() {
		file_process_id_proto_rawDescData = protoimpl.X.CompressGZIP(file_process_id_proto_rawDescData)
	})
	return file_process_id_proto_rawDescData
}

var file_process_id_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_process_id_proto_goTypes = []interface{}{
	(*ProcessId)(nil), // 0: prc.ProcessId
}
var file_process_id_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_process_id_proto_init() }
func file_process_id_proto_init() {
	if File_process_id_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_process_id_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessId); i {
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
	type x struct{ cache atomic.Pointer[Process] }
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_process_id_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_process_id_proto_goTypes,
		DependencyIndexes: file_process_id_proto_depIdxs,
		MessageInfos:      file_process_id_proto_msgTypes,
	}.Build()
	File_process_id_proto = out.File
	file_process_id_proto_rawDesc = nil
	file_process_id_proto_goTypes = nil
	file_process_id_proto_depIdxs = nil
}
