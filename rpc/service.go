package rpc

import "context"

type (
	// UnaryCaller 该调用器将同步的发起一个仅由一个处理器响应的调用，被调用方的响应将存储在 Reader 中进行返回
	UnaryCaller func(ctx context.Context, params any) (Reader, error)
	// UnaryNotifyCaller 该调用器将同步的发起一个仅由一个处理器处理的调用，该调用不需要被调用方响应
	UnaryNotifyCaller func(ctx context.Context, params any) error
	// AsyncUnaryCaller 该调用器将异步的发起一个仅由一个处理器响应的调用，被调用方的响应将存储在 callback 的 Reader 中进行返回
	AsyncUnaryCaller func(ctx context.Context, params any, callback func(reader Reader, err error))
	// AsyncNotifyCaller 该调用器将异步的发起一个由多个处理器处理的调用
	AsyncNotifyCaller func(params any) error
)

type Service struct {
	Name       string            `json:"name"`        // 服务名称
	InstanceId string            `json:"instance_id"` // 全局唯一的服务实例标识符
	Host       string            `json:"host"`        // 服务主机
	Port       int               `json:"port"`        // 服务端口
	Metadata   map[string]string `json:"metadata"`    // 服务元数据
}

type CallableService interface {
	// GetServiceInfo 获取服务信息
	GetServiceInfo() Service

	UnaryCall(route ...Route) UnaryCaller

	UnaryNotifyCall(route ...Route) UnaryNotifyCaller

	AsyncUnaryCall(route ...Route) AsyncUnaryCaller

	AsyncNotifyCall(route ...Route) AsyncNotifyCaller
}
