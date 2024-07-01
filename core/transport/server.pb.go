// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: server.proto

package transport

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

type DistributedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to MessageType:
	//
	//	*DistributedMessage_ConnectionOpen
	//	*DistributedMessage_ConnectionOpened
	//	*DistributedMessage_ConnectionClosed
	//	*DistributedMessage_ConnectionMessageBatch
	MessageType isDistributedMessage_MessageType `protobuf_oneof:"message_type"`
}

func (x *DistributedMessage) Reset() {
	*x = DistributedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DistributedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DistributedMessage) ProtoMessage() {}

func (x *DistributedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DistributedMessage.ProtoReflect.Descriptor instead.
func (*DistributedMessage) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{0}
}

func (m *DistributedMessage) GetMessageType() isDistributedMessage_MessageType {
	if m != nil {
		return m.MessageType
	}
	return nil
}

func (x *DistributedMessage) GetConnectionOpen() *ConnectionOpen {
	if x, ok := x.GetMessageType().(*DistributedMessage_ConnectionOpen); ok {
		return x.ConnectionOpen
	}
	return nil
}

func (x *DistributedMessage) GetConnectionOpened() *ConnectionOpened {
	if x, ok := x.GetMessageType().(*DistributedMessage_ConnectionOpened); ok {
		return x.ConnectionOpened
	}
	return nil
}

func (x *DistributedMessage) GetConnectionClosed() *ConnectionClosed {
	if x, ok := x.GetMessageType().(*DistributedMessage_ConnectionClosed); ok {
		return x.ConnectionClosed
	}
	return nil
}

func (x *DistributedMessage) GetConnectionMessageBatch() *ConnectionMessageBatch {
	if x, ok := x.GetMessageType().(*DistributedMessage_ConnectionMessageBatch); ok {
		return x.ConnectionMessageBatch
	}
	return nil
}

type isDistributedMessage_MessageType interface {
	isDistributedMessage_MessageType()
}

type DistributedMessage_ConnectionOpen struct {
	ConnectionOpen *ConnectionOpen `protobuf:"bytes,1,opt,name=connection_open,json=connectionOpen,proto3,oneof"`
}

type DistributedMessage_ConnectionOpened struct {
	ConnectionOpened *ConnectionOpened `protobuf:"bytes,2,opt,name=connection_opened,json=connectionOpened,proto3,oneof"`
}

type DistributedMessage_ConnectionClosed struct {
	ConnectionClosed *ConnectionClosed `protobuf:"bytes,3,opt,name=connection_closed,json=connectionClosed,proto3,oneof"`
}

type DistributedMessage_ConnectionMessageBatch struct {
	ConnectionMessageBatch *ConnectionMessageBatch `protobuf:"bytes,5,opt,name=connection_message_batch,json=connectionMessageBatch,proto3,oneof"`
}

func (*DistributedMessage_ConnectionOpen) isDistributedMessage_MessageType() {}

func (*DistributedMessage_ConnectionOpened) isDistributedMessage_MessageType() {}

func (*DistributedMessage_ConnectionClosed) isDistributedMessage_MessageType() {}

func (*DistributedMessage_ConnectionMessageBatch) isDistributedMessage_MessageType() {}

type ConnectionOpen struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=Address,proto3" json:"Address,omitempty"`
}

func (x *ConnectionOpen) Reset() {
	*x = ConnectionOpen{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionOpen) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionOpen) ProtoMessage() {}

func (x *ConnectionOpen) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionOpen.ProtoReflect.Descriptor instead.
func (*ConnectionOpen) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{1}
}

func (x *ConnectionOpen) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type ConnectionOpened struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConnectionOpened) Reset() {
	*x = ConnectionOpened{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionOpened) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionOpened) ProtoMessage() {}

func (x *ConnectionOpened) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionOpened.ProtoReflect.Descriptor instead.
func (*ConnectionOpened) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{2}
}

type ConnectionClosed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConnectionClosed) Reset() {
	*x = ConnectionClosed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionClosed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionClosed) ProtoMessage() {}

func (x *ConnectionClosed) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionClosed.ProtoReflect.Descriptor instead.
func (*ConnectionClosed) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{3}
}

type ConnectionMessageBatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderAddress                  [][]byte `protobuf:"bytes,1,rep,name=sender_address,json=senderAddress,proto3" json:"sender_address,omitempty"`
	ReceiverAddress                [][]byte `protobuf:"bytes,2,rep,name=receiver_address,json=receiverAddress,proto3" json:"receiver_address,omitempty"`
	RegulatoryMessageSenderAddress [][]byte `protobuf:"bytes,3,rep,name=regulatory_message_sender_address,json=regulatoryMessageSenderAddress,proto3" json:"regulatory_message_sender_address,omitempty"`
	TypeName                       []string `protobuf:"bytes,4,rep,name=type_name,json=typeName,proto3" json:"type_name,omitempty"`
	MessageData                    [][]byte `protobuf:"bytes,5,rep,name=message_data,json=messageData,proto3" json:"message_data,omitempty"`
	Bad                            []bool   `protobuf:"varint,6,rep,packed,name=bad,proto3" json:"bad,omitempty"`
	System                         []bool   `protobuf:"varint,7,rep,packed,name=system,proto3" json:"system,omitempty"`
}

func (x *ConnectionMessageBatch) Reset() {
	*x = ConnectionMessageBatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionMessageBatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionMessageBatch) ProtoMessage() {}

func (x *ConnectionMessageBatch) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionMessageBatch.ProtoReflect.Descriptor instead.
func (*ConnectionMessageBatch) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{4}
}

func (x *ConnectionMessageBatch) GetSenderAddress() [][]byte {
	if x != nil {
		return x.SenderAddress
	}
	return nil
}

func (x *ConnectionMessageBatch) GetReceiverAddress() [][]byte {
	if x != nil {
		return x.ReceiverAddress
	}
	return nil
}

func (x *ConnectionMessageBatch) GetRegulatoryMessageSenderAddress() [][]byte {
	if x != nil {
		return x.RegulatoryMessageSenderAddress
	}
	return nil
}

func (x *ConnectionMessageBatch) GetTypeName() []string {
	if x != nil {
		return x.TypeName
	}
	return nil
}

func (x *ConnectionMessageBatch) GetMessageData() [][]byte {
	if x != nil {
		return x.MessageData
	}
	return nil
}

func (x *ConnectionMessageBatch) GetBad() []bool {
	if x != nil {
		return x.Bad
	}
	return nil
}

func (x *ConnectionMessageBatch) GetSystem() []bool {
	if x != nil {
		return x.System
	}
	return nil
}

var File_server_proto protoreflect.FileDescriptor

var file_server_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x22, 0xd5, 0x02, 0x0a, 0x12, 0x44, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x41, 0x0a,
	0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x70, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x6e, 0x48, 0x00,
	0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x6e,
	0x12, 0x47, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6f,
	0x70, 0x65, 0x6e, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x72, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f,
	0x70, 0x65, 0x6e, 0x65, 0x64, 0x48, 0x00, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x12, 0x47, 0x0a, 0x11, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x48, 0x00,
	0x52, 0x10, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6c, 0x6f, 0x73,
	0x65, 0x64, 0x12, 0x5a, 0x0a, 0x18, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x48, 0x00, 0x52, 0x16, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x42, 0x0e,
	0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x2a,
	0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x6e,
	0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x12, 0x0a, 0x10, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x22, 0x12,
	0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6c, 0x6f, 0x73,
	0x65, 0x64, 0x22, 0x9f, 0x02, 0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x25, 0x0a,
	0x0e, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0d, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0f,
	0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x49, 0x0a, 0x21, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x1e, 0x72, 0x65, 0x67, 0x75,
	0x6c, 0x61, 0x74, 0x6f, 0x72, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x79,
	0x70, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x74,
	0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0b, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x61,
	0x64, 0x18, 0x06, 0x20, 0x03, 0x28, 0x08, 0x52, 0x03, 0x62, 0x61, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x07, 0x20, 0x03, 0x28, 0x08, 0x52, 0x06, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x32, 0x69, 0x0a, 0x18, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x4d, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x72, 0x12, 0x1a, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1a, 0x2e,
	0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42,
	0x30, 0x5a, 0x2e, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b,
	0x65, 0x72, 0x63, 0x79, 0x6c, 0x61, 0x6e, 0x39, 0x38, 0x2f, 0x6d, 0x69, 0x6e, 0x6f, 0x74, 0x61,
	0x75, 0x72, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_proto_rawDescOnce sync.Once
	file_server_proto_rawDescData = file_server_proto_rawDesc
)

func file_server_proto_rawDescGZIP() []byte {
	file_server_proto_rawDescOnce.Do(func() {
		file_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_proto_rawDescData)
	})
	return file_server_proto_rawDescData
}

var file_server_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_server_proto_goTypes = []interface{}{
	(*DistributedMessage)(nil),     // 0: remote.DistributedMessage
	(*ConnectionOpen)(nil),         // 1: remote.ConnectionOpen
	(*ConnectionOpened)(nil),       // 2: remote.ConnectionOpened
	(*ConnectionClosed)(nil),       // 3: remote.ConnectionClosed
	(*ConnectionMessageBatch)(nil), // 4: remote.ConnectionMessageBatch
}
var file_server_proto_depIdxs = []int32{
	1, // 0: remote.DistributedMessage.connection_open:type_name -> remote.ConnectionOpen
	2, // 1: remote.DistributedMessage.connection_opened:type_name -> remote.ConnectionOpened
	3, // 2: remote.DistributedMessage.connection_closed:type_name -> remote.ConnectionClosed
	4, // 3: remote.DistributedMessage.connection_message_batch:type_name -> remote.ConnectionMessageBatch
	0, // 4: remote.ActorSystemCommunication.StreamHandler:input_type -> remote.DistributedMessage
	0, // 5: remote.ActorSystemCommunication.StreamHandler:output_type -> remote.DistributedMessage
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_server_proto_init() }
func file_server_proto_init() {
	if File_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DistributedMessage); i {
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
		file_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionOpen); i {
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
		file_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionOpened); i {
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
		file_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionClosed); i {
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
		file_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionMessageBatch); i {
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
	file_server_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*DistributedMessage_ConnectionOpen)(nil),
		(*DistributedMessage_ConnectionOpened)(nil),
		(*DistributedMessage_ConnectionClosed)(nil),
		(*DistributedMessage_ConnectionMessageBatch)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_proto_goTypes,
		DependencyIndexes: file_server_proto_depIdxs,
		MessageInfos:      file_server_proto_msgTypes,
	}.Build()
	File_server_proto = out.File
	file_server_proto_rawDesc = nil
	file_server_proto_goTypes = nil
	file_server_proto_depIdxs = nil
}
