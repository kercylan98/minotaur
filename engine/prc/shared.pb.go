// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: shared.proto

package prc

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

type SharedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to MessageType:
	//
	//	*SharedMessage_Handshake
	//	*SharedMessage_Farewell
	//	*SharedMessage_DeliveryMessage
	//	*SharedMessage_BatchDeliveryMessage
	MessageType isSharedMessage_MessageType `protobuf_oneof:"message_type"`
}

func (x *SharedMessage) Reset() {
	*x = SharedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SharedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SharedMessage) ProtoMessage() {}

func (x *SharedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_shared_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SharedMessage.ProtoReflect.Descriptor instead.
func (*SharedMessage) Descriptor() ([]byte, []int) {
	return file_shared_proto_rawDescGZIP(), []int{0}
}

func (m *SharedMessage) GetMessageType() isSharedMessage_MessageType {
	if m != nil {
		return m.MessageType
	}
	return nil
}

func (x *SharedMessage) GetHandshake() *Handshake {
	if x, ok := x.GetMessageType().(*SharedMessage_Handshake); ok {
		return x.Handshake
	}
	return nil
}

func (x *SharedMessage) GetFarewell() *Farewell {
	if x, ok := x.GetMessageType().(*SharedMessage_Farewell); ok {
		return x.Farewell
	}
	return nil
}

func (x *SharedMessage) GetDeliveryMessage() *DeliveryMessage {
	if x, ok := x.GetMessageType().(*SharedMessage_DeliveryMessage); ok {
		return x.DeliveryMessage
	}
	return nil
}

func (x *SharedMessage) GetBatchDeliveryMessage() *BatchDeliveryMessage {
	if x, ok := x.GetMessageType().(*SharedMessage_BatchDeliveryMessage); ok {
		return x.BatchDeliveryMessage
	}
	return nil
}

type isSharedMessage_MessageType interface {
	isSharedMessage_MessageType()
}

type SharedMessage_Handshake struct {
	Handshake *Handshake `protobuf:"bytes,1,opt,name=handshake,proto3,oneof"` // 握手
}

type SharedMessage_Farewell struct {
	Farewell *Farewell `protobuf:"bytes,2,opt,name=farewell,proto3,oneof"` // 告别
}

type SharedMessage_DeliveryMessage struct {
	DeliveryMessage *DeliveryMessage `protobuf:"bytes,3,opt,name=delivery_message,json=deliveryMessage,proto3,oneof"` // 传递单条消息
}

type SharedMessage_BatchDeliveryMessage struct {
	BatchDeliveryMessage *BatchDeliveryMessage `protobuf:"bytes,4,opt,name=batch_delivery_message,json=batchDeliveryMessage,proto3,oneof"` // 传递多条消息
}

func (*SharedMessage_Handshake) isSharedMessage_MessageType() {}

func (*SharedMessage_Farewell) isSharedMessage_MessageType() {}

func (*SharedMessage_DeliveryMessage) isSharedMessage_MessageType() {}

func (*SharedMessage_BatchDeliveryMessage) isSharedMessage_MessageType() {}

type Handshake struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *Handshake) Reset() {
	*x = Handshake{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Handshake) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Handshake) ProtoMessage() {}

func (x *Handshake) ProtoReflect() protoreflect.Message {
	mi := &file_shared_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Handshake.ProtoReflect.Descriptor instead.
func (*Handshake) Descriptor() ([]byte, []int) {
	return file_shared_proto_rawDescGZIP(), []int{1}
}

func (x *Handshake) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type Farewell struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *Farewell) Reset() {
	*x = Farewell{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Farewell) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Farewell) ProtoMessage() {}

func (x *Farewell) ProtoReflect() protoreflect.Message {
	mi := &file_shared_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Farewell.ProtoReflect.Descriptor instead.
func (*Farewell) Descriptor() ([]byte, []int) {
	return file_shared_proto_rawDescGZIP(), []int{2}
}

func (x *Farewell) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type DeliveryMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender      *ProcessId `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`                              // 发送方
	Receiver    *ProcessId `protobuf:"bytes,2,opt,name=receiver,proto3" json:"receiver,omitempty"`                          // 接收方
	MessageType string     `protobuf:"bytes,3,opt,name=message_type,json=messageType,proto3" json:"message_type,omitempty"` // 消息类型名称
	MessageData []byte     `protobuf:"bytes,4,opt,name=message_data,json=messageData,proto3" json:"message_data,omitempty"` // 消息数据
	System      bool       `protobuf:"varint,5,opt,name=system,proto3" json:"system,omitempty"`                             // 是否是系统消息
}

func (x *DeliveryMessage) Reset() {
	*x = DeliveryMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliveryMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveryMessage) ProtoMessage() {}

func (x *DeliveryMessage) ProtoReflect() protoreflect.Message {
	mi := &file_shared_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveryMessage.ProtoReflect.Descriptor instead.
func (*DeliveryMessage) Descriptor() ([]byte, []int) {
	return file_shared_proto_rawDescGZIP(), []int{3}
}

func (x *DeliveryMessage) GetSender() *ProcessId {
	if x != nil {
		return x.Sender
	}
	return nil
}

func (x *DeliveryMessage) GetReceiver() *ProcessId {
	if x != nil {
		return x.Receiver
	}
	return nil
}

func (x *DeliveryMessage) GetMessageType() string {
	if x != nil {
		return x.MessageType
	}
	return ""
}

func (x *DeliveryMessage) GetMessageData() []byte {
	if x != nil {
		return x.MessageData
	}
	return nil
}

func (x *DeliveryMessage) GetSystem() bool {
	if x != nil {
		return x.System
	}
	return false
}

type BatchDeliveryMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []*DeliveryMessage `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *BatchDeliveryMessage) Reset() {
	*x = BatchDeliveryMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shared_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchDeliveryMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchDeliveryMessage) ProtoMessage() {}

func (x *BatchDeliveryMessage) ProtoReflect() protoreflect.Message {
	mi := &file_shared_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchDeliveryMessage.ProtoReflect.Descriptor instead.
func (*BatchDeliveryMessage) Descriptor() ([]byte, []int) {
	return file_shared_proto_rawDescGZIP(), []int{4}
}

func (x *BatchDeliveryMessage) GetMessages() []*DeliveryMessage {
	if x != nil {
		return x.Messages
	}
	return nil
}

var File_shared_proto protoreflect.FileDescriptor

var file_shared_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03,
	0x70, 0x72, 0x63, 0x1a, 0x10, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x02, 0x0a, 0x0d, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x09, 0x68, 0x61, 0x6e, 0x64, 0x73,
	0x68, 0x61, 0x6b, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x63,
	0x2e, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x48, 0x00, 0x52, 0x09, 0x68, 0x61,
	0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x66, 0x61, 0x72, 0x65, 0x77,
	0x65, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x63, 0x2e,
	0x46, 0x61, 0x72, 0x65, 0x77, 0x65, 0x6c, 0x6c, 0x48, 0x00, 0x52, 0x08, 0x66, 0x61, 0x72, 0x65,
	0x77, 0x65, 0x6c, 0x6c, 0x12, 0x41, 0x0a, 0x10, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x70, 0x72, 0x63, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x51, 0x0a, 0x16, 0x62, 0x61, 0x74, 0x63, 0x68,
	0x5f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x63, 0x2e, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x48, 0x00, 0x52, 0x14, 0x62, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x25, 0x0a, 0x09, 0x48, 0x61,
	0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x22, 0x24, 0x0a, 0x08, 0x46, 0x61, 0x72, 0x65, 0x77, 0x65, 0x6c, 0x6c, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0xc3, 0x01, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x73,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72,
	0x63, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x64, 0x52, 0x06, 0x73, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x63, 0x2e, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x49, 0x64, 0x52, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12,
	0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x22, 0x48, 0x0a,
	0x14, 0x42, 0x61, 0x74, 0x63, 0x68, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x63, 0x2e, 0x44, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x32, 0x47, 0x0a, 0x06, 0x53, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x12, 0x3d, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x48, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x72, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x63, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x63, 0x2e, 0x53, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01,
	0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b,
	0x65, 0x72, 0x63, 0x79, 0x6c, 0x61, 0x6e, 0x39, 0x38, 0x2f, 0x6d, 0x69, 0x6e, 0x6f, 0x74, 0x61,
	0x75, 0x72, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x76, 0x69, 0x76, 0x69, 0x64, 0x2f, 0x70, 0x72, 0x63,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shared_proto_rawDescOnce sync.Once
	file_shared_proto_rawDescData = file_shared_proto_rawDesc
)

func file_shared_proto_rawDescGZIP() []byte {
	file_shared_proto_rawDescOnce.Do(func() {
		file_shared_proto_rawDescData = protoimpl.X.CompressGZIP(file_shared_proto_rawDescData)
	})
	return file_shared_proto_rawDescData
}

var file_shared_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_shared_proto_goTypes = []interface{}{
	(*SharedMessage)(nil),        // 0: prc.SharedMessage
	(*Handshake)(nil),            // 1: prc.Handshake
	(*Farewell)(nil),             // 2: prc.Farewell
	(*DeliveryMessage)(nil),      // 3: prc.DeliveryMessage
	(*BatchDeliveryMessage)(nil), // 4: prc.BatchDeliveryMessage
	(*ProcessId)(nil),            // 5: prc.ProcessId
}
var file_shared_proto_depIdxs = []int32{
	1, // 0: prc.SharedMessage.handshake:type_name -> prc.Handshake
	2, // 1: prc.SharedMessage.farewell:type_name -> prc.Farewell
	3, // 2: prc.SharedMessage.delivery_message:type_name -> prc.DeliveryMessage
	4, // 3: prc.SharedMessage.batch_delivery_message:type_name -> prc.BatchDeliveryMessage
	5, // 4: prc.DeliveryMessage.sender:type_name -> prc.ProcessId
	5, // 5: prc.DeliveryMessage.receiver:type_name -> prc.ProcessId
	3, // 6: prc.BatchDeliveryMessage.messages:type_name -> prc.DeliveryMessage
	0, // 7: prc.Shared.StreamHandler:input_type -> prc.SharedMessage
	0, // 8: prc.Shared.StreamHandler:output_type -> prc.SharedMessage
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_shared_proto_init() }
func file_shared_proto_init() {
	if File_shared_proto != nil {
		return
	}
	file_process_id_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_shared_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SharedMessage); i {
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
		file_shared_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Handshake); i {
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
		file_shared_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Farewell); i {
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
		file_shared_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliveryMessage); i {
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
		file_shared_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchDeliveryMessage); i {
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
	file_shared_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SharedMessage_Handshake)(nil),
		(*SharedMessage_Farewell)(nil),
		(*SharedMessage_DeliveryMessage)(nil),
		(*SharedMessage_BatchDeliveryMessage)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_shared_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shared_proto_goTypes,
		DependencyIndexes: file_shared_proto_depIdxs,
		MessageInfos:      file_shared_proto_msgTypes,
	}.Build()
	File_shared_proto = out.File
	file_shared_proto_rawDesc = nil
	file_shared_proto_goTypes = nil
	file_shared_proto_depIdxs = nil
}