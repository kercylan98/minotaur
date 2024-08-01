// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.28.0--rc1
// source: internal_messages.proto

package messages

import (
	prc "github.com/kercylan98/minotaur/engine/prc"
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

type GenerateRemoteActor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParentPid                      *prc.ProcessId   `protobuf:"bytes,1,opt,name=parent_pid,json=parentPid,proto3" json:"parent_pid,omitempty"`
	ProviderName                   string           `protobuf:"bytes,2,opt,name=provider_name,json=providerName,proto3" json:"provider_name,omitempty"`
	ActorName                      string           `protobuf:"bytes,100,opt,name=actor_name,json=actorName,proto3" json:"actor_name,omitempty"`
	ActorNamePrefix                string           `protobuf:"bytes,101,opt,name=actor_name_prefix,json=actorNamePrefix,proto3" json:"actor_name_prefix,omitempty"`
	MailboxProviderName            string           `protobuf:"bytes,102,opt,name=mailbox_provider_name,json=mailboxProviderName,proto3" json:"mailbox_provider_name,omitempty"`
	DispatcherProviderName         string           `protobuf:"bytes,103,opt,name=dispatcher_provider_name,json=dispatcherProviderName,proto3" json:"dispatcher_provider_name,omitempty"`
	SupervisionStrategyName        string           `protobuf:"bytes,104,opt,name=supervision_strategy_name,json=supervisionStrategyName,proto3" json:"supervision_strategy_name,omitempty"`
	ExpireDuration                 int64            `protobuf:"varint,105,opt,name=expire_duration,json=expireDuration,proto3" json:"expire_duration,omitempty"`
	IdleDeadline                   int64            `protobuf:"varint,106,opt,name=idle_deadline,json=idleDeadline,proto3" json:"idle_deadline,omitempty"`
	PersistenceStorageProviderName string           `protobuf:"bytes,107,opt,name=persistence_storage_provider_name,json=persistenceStorageProviderName,proto3" json:"persistence_storage_provider_name,omitempty"`
	PersistenceName                string           `protobuf:"bytes,108,opt,name=persistence_name,json=persistenceName,proto3" json:"persistence_name,omitempty"`
	PersistenceEventThreshold      int32            `protobuf:"varint,109,opt,name=persistence_event_threshold,json=persistenceEventThreshold,proto3" json:"persistence_event_threshold,omitempty"`
	SlowProcessDuration            int64            `protobuf:"varint,110,opt,name=slow_process_duration,json=slowProcessDuration,proto3" json:"slow_process_duration,omitempty"`
	SlowProcessReceivers           []*prc.ProcessId `protobuf:"bytes,111,rep,name=slow_process_receivers,json=slowProcessReceivers,proto3" json:"slow_process_receivers,omitempty"`
}

func (x *GenerateRemoteActor) Reset() {
	*x = GenerateRemoteActor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateRemoteActor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateRemoteActor) ProtoMessage() {}

func (x *GenerateRemoteActor) ProtoReflect() protoreflect.Message {
	mi := &file_internal_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateRemoteActor.ProtoReflect.Descriptor instead.
func (*GenerateRemoteActor) Descriptor() ([]byte, []int) {
	return file_internal_messages_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateRemoteActor) GetParentPid() *prc.ProcessId {
	if x != nil {
		return x.ParentPid
	}
	return nil
}

func (x *GenerateRemoteActor) GetProviderName() string {
	if x != nil {
		return x.ProviderName
	}
	return ""
}

func (x *GenerateRemoteActor) GetActorName() string {
	if x != nil {
		return x.ActorName
	}
	return ""
}

func (x *GenerateRemoteActor) GetActorNamePrefix() string {
	if x != nil {
		return x.ActorNamePrefix
	}
	return ""
}

func (x *GenerateRemoteActor) GetMailboxProviderName() string {
	if x != nil {
		return x.MailboxProviderName
	}
	return ""
}

func (x *GenerateRemoteActor) GetDispatcherProviderName() string {
	if x != nil {
		return x.DispatcherProviderName
	}
	return ""
}

func (x *GenerateRemoteActor) GetSupervisionStrategyName() string {
	if x != nil {
		return x.SupervisionStrategyName
	}
	return ""
}

func (x *GenerateRemoteActor) GetExpireDuration() int64 {
	if x != nil {
		return x.ExpireDuration
	}
	return 0
}

func (x *GenerateRemoteActor) GetIdleDeadline() int64 {
	if x != nil {
		return x.IdleDeadline
	}
	return 0
}

func (x *GenerateRemoteActor) GetPersistenceStorageProviderName() string {
	if x != nil {
		return x.PersistenceStorageProviderName
	}
	return ""
}

func (x *GenerateRemoteActor) GetPersistenceName() string {
	if x != nil {
		return x.PersistenceName
	}
	return ""
}

func (x *GenerateRemoteActor) GetPersistenceEventThreshold() int32 {
	if x != nil {
		return x.PersistenceEventThreshold
	}
	return 0
}

func (x *GenerateRemoteActor) GetSlowProcessDuration() int64 {
	if x != nil {
		return x.SlowProcessDuration
	}
	return 0
}

func (x *GenerateRemoteActor) GetSlowProcessReceivers() []*prc.ProcessId {
	if x != nil {
		return x.SlowProcessReceivers
	}
	return nil
}

type GenerateRemoteActorResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid *prc.ProcessId `protobuf:"bytes,1,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (x *GenerateRemoteActorResult) Reset() {
	*x = GenerateRemoteActorResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateRemoteActorResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateRemoteActorResult) ProtoMessage() {}

func (x *GenerateRemoteActorResult) ProtoReflect() protoreflect.Message {
	mi := &file_internal_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateRemoteActorResult.ProtoReflect.Descriptor instead.
func (*GenerateRemoteActorResult) Descriptor() ([]byte, []int) {
	return file_internal_messages_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateRemoteActorResult) GetPid() *prc.ProcessId {
	if x != nil {
		return x.Pid
	}
	return nil
}

// Terminated 当收到该消息时，说明 TerminatedActor 已经被终止，如果是自身，那么表示自身已被终止。
type Terminated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TerminatedProcess *prc.ProcessId `protobuf:"bytes,1,opt,name=terminated_process,json=terminatedProcess,proto3" json:"terminated_process,omitempty"`
}

func (x *Terminated) Reset() {
	*x = Terminated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Terminated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Terminated) ProtoMessage() {}

func (x *Terminated) ProtoReflect() protoreflect.Message {
	mi := &file_internal_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Terminated.ProtoReflect.Descriptor instead.
func (*Terminated) Descriptor() ([]byte, []int) {
	return file_internal_messages_proto_rawDescGZIP(), []int{2}
}

func (x *Terminated) GetTerminatedProcess() *prc.ProcessId {
	if x != nil {
		return x.TerminatedProcess
	}
	return nil
}

type Watch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Watch) Reset() {
	*x = Watch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Watch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Watch) ProtoMessage() {}

func (x *Watch) ProtoReflect() protoreflect.Message {
	mi := &file_internal_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Watch.ProtoReflect.Descriptor instead.
func (*Watch) Descriptor() ([]byte, []int) {
	return file_internal_messages_proto_rawDescGZIP(), []int{3}
}

type Unwatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Unwatch) Reset() {
	*x = Unwatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unwatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unwatch) ProtoMessage() {}

func (x *Unwatch) ProtoReflect() protoreflect.Message {
	mi := &file_internal_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unwatch.ProtoReflect.Descriptor instead.
func (*Unwatch) Descriptor() ([]byte, []int) {
	return file_internal_messages_proto_rawDescGZIP(), []int{4}
}

type SlowProcess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Duration int64          `protobuf:"varint,1,opt,name=duration,proto3" json:"duration,omitempty"` // 耗时
	Pid      *prc.ProcessId `protobuf:"bytes,2,opt,name=pid,proto3" json:"pid,omitempty"`            // 耗时进程
}

func (x *SlowProcess) Reset() {
	*x = SlowProcess{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlowProcess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlowProcess) ProtoMessage() {}

func (x *SlowProcess) ProtoReflect() protoreflect.Message {
	mi := &file_internal_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlowProcess.ProtoReflect.Descriptor instead.
func (*SlowProcess) Descriptor() ([]byte, []int) {
	return file_internal_messages_proto_rawDescGZIP(), []int{5}
}

func (x *SlowProcess) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *SlowProcess) GetPid() *prc.ProcessId {
	if x != nil {
		return x.Pid
	}
	return nil
}

var File_internal_messages_proto protoreflect.FileDescriptor

var file_internal_messages_proto_rawDesc = []byte{
	0x0a, 0x17, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x1a, 0x10, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdc, 0x05, 0x0a, 0x13, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x2d, 0x0a,
	0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x63, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49,
	0x64, 0x52, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x2a, 0x0a, 0x11, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x70,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x65, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x32, 0x0a, 0x15,
	0x6d, 0x61, 0x69, 0x6c, 0x62, 0x6f, 0x78, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x6d, 0x61, 0x69,
	0x6c, 0x62, 0x6f, 0x78, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x38, 0x0a, 0x18, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x67, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x16, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x19, 0x73, 0x75,
	0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x68, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x73,
	0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x69, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x23, 0x0a, 0x0d, 0x69, 0x64, 0x6c, 0x65, 0x5f, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65,
	0x18, 0x6a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x69, 0x64, 0x6c, 0x65, 0x44, 0x65, 0x61, 0x64,
	0x6c, 0x69, 0x6e, 0x65, 0x12, 0x49, 0x0a, 0x21, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65,
	0x6e, 0x63, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x6b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x1e, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x29, 0x0a, 0x10, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x6c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x65, 0x72, 0x73, 0x69,
	0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x1b, 0x70, 0x65,
	0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x6d, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x19, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x32, 0x0a, 0x15, 0x73, 0x6c,
	0x6f, 0x77, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x73, 0x6c, 0x6f, 0x77, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x44,
	0x0a, 0x16, 0x73, 0x6c, 0x6f, 0x77, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x73, 0x18, 0x6f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x70, 0x72, 0x63, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x64, 0x52, 0x14,
	0x73, 0x6c, 0x6f, 0x77, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x72, 0x73, 0x22, 0x3d, 0x0a, 0x19, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x20, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x70, 0x72, 0x63, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x64, 0x52, 0x03,
	0x70, 0x69, 0x64, 0x22, 0x4b, 0x0a, 0x0a, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x3d, 0x0a, 0x12, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x70, 0x72, 0x63, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x64, 0x52, 0x11, 0x74,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x22, 0x07, 0x0a, 0x05, 0x57, 0x61, 0x74, 0x63, 0x68, 0x22, 0x09, 0x0a, 0x07, 0x55, 0x6e, 0x77,
	0x61, 0x74, 0x63, 0x68, 0x22, 0x4b, 0x0a, 0x0b, 0x53, 0x6c, 0x6f, 0x77, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x20, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70,
	0x72, 0x63, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x64, 0x52, 0x03, 0x70, 0x69,
	0x64, 0x42, 0x40, 0x5a, 0x3e, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6b, 0x65, 0x72, 0x63, 0x79, 0x6c, 0x61, 0x6e, 0x39, 0x38, 0x2f, 0x6d, 0x69, 0x6e, 0x6f,
	0x74, 0x61, 0x75, 0x72, 0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x69, 0x76, 0x69,
	0x64, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_messages_proto_rawDescOnce sync.Once
	file_internal_messages_proto_rawDescData = file_internal_messages_proto_rawDesc
)

func file_internal_messages_proto_rawDescGZIP() []byte {
	file_internal_messages_proto_rawDescOnce.Do(func() {
		file_internal_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_messages_proto_rawDescData)
	})
	return file_internal_messages_proto_rawDescData
}

var file_internal_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_internal_messages_proto_goTypes = []interface{}{
	(*GenerateRemoteActor)(nil),       // 0: messages.GenerateRemoteActor
	(*GenerateRemoteActorResult)(nil), // 1: messages.GenerateRemoteActorResult
	(*Terminated)(nil),                // 2: messages.Terminated
	(*Watch)(nil),                     // 3: messages.Watch
	(*Unwatch)(nil),                   // 4: messages.Unwatch
	(*SlowProcess)(nil),               // 5: messages.SlowProcess
	(*prc.ProcessId)(nil),             // 6: prc.ProcessId
}
var file_internal_messages_proto_depIdxs = []int32{
	6, // 0: messages.GenerateRemoteActor.parent_pid:type_name -> prc.ProcessId
	6, // 1: messages.GenerateRemoteActor.slow_process_receivers:type_name -> prc.ProcessId
	6, // 2: messages.GenerateRemoteActorResult.pid:type_name -> prc.ProcessId
	6, // 3: messages.Terminated.terminated_process:type_name -> prc.ProcessId
	6, // 4: messages.SlowProcess.pid:type_name -> prc.ProcessId
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_internal_messages_proto_init() }
func file_internal_messages_proto_init() {
	if File_internal_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateRemoteActor); i {
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
		file_internal_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateRemoteActorResult); i {
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
		file_internal_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Terminated); i {
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
		file_internal_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Watch); i {
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
		file_internal_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unwatch); i {
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
		file_internal_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlowProcess); i {
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
			RawDescriptor: file_internal_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_messages_proto_goTypes,
		DependencyIndexes: file_internal_messages_proto_depIdxs,
		MessageInfos:      file_internal_messages_proto_msgTypes,
	}.Build()
	File_internal_messages_proto = out.File
	file_internal_messages_proto_rawDesc = nil
	file_internal_messages_proto_goTypes = nil
	file_internal_messages_proto_depIdxs = nil
}
