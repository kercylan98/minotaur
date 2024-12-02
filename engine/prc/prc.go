package prc

import prcv1 "github.com/kercylan98/minotaur/engine/prc/v1"

// Message 消息是用于传递的数据
type Message = any

const (
	LocalhostPhysicalAddress = "localhost" // 无网络本地
)

type (
	// PhysicalAddress 物理地址是用于标识内容的网络地址
	PhysicalAddress = prcv1.PhysicalAddress
	// LogicalAddress 逻辑地址是用于标识内容的本地内部地址
	LogicalAddress = prcv1.LogicalAddress

	ProcessId = prcv1.ProcessId
)

var (
	sharedServiceDesc = prcv1.Shared_ServiceDesc
	newSharedClient   = prcv1.NewSharedClient
)

type (
	deliveryMessage                   = prcv1.DeliveryMessage
	batchDeliveryMessage              = prcv1.BatchDeliveryMessage
	sharedMessage                     = prcv1.SharedMessage
	sharedMessageHandshake            = prcv1.SharedMessage_Handshake
	sharedMessageFarewell             = prcv1.SharedMessage_Farewell
	sharedMessageDeliveryMessage      = prcv1.SharedMessage_DeliveryMessage
	sharedMessageBatchDeliveryMessage = prcv1.SharedMessage_BatchDeliveryMessage
	sharedErrorMessage                = prcv1.SharedErrorMessage
	handshake                         = prcv1.Handshake
	farewell                          = prcv1.Farewell
	sharedStreamHandlerServer         = prcv1.Shared_StreamHandlerServer
	sharedStreamHandlerClient         = prcv1.Shared_StreamHandlerClient
	unimplementedSharedServer         = prcv1.UnimplementedSharedServer
)
