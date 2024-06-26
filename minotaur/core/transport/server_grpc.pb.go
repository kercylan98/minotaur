// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package transport

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ActorSystemCommunicationClient is the client API for ActorSystemCommunication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActorSystemCommunicationClient interface {
	StreamHandler(ctx context.Context, opts ...grpc.CallOption) (ActorSystemCommunication_StreamHandlerClient, error)
}

type actorSystemCommunicationClient struct {
	cc grpc.ClientConnInterface
}

func NewActorSystemCommunicationClient(cc grpc.ClientConnInterface) ActorSystemCommunicationClient {
	return &actorSystemCommunicationClient{cc}
}

func (c *actorSystemCommunicationClient) StreamHandler(ctx context.Context, opts ...grpc.CallOption) (ActorSystemCommunication_StreamHandlerClient, error) {
	stream, err := c.cc.NewStream(ctx, &ActorSystemCommunication_ServiceDesc.Streams[0], "/remote.ActorSystemCommunication/StreamHandler", opts...)
	if err != nil {
		return nil, err
	}
	x := &actorSystemCommunicationStreamHandlerClient{stream}
	return x, nil
}

type ActorSystemCommunication_StreamHandlerClient interface {
	Send(*DistributedMessage) error
	Recv() (*DistributedMessage, error)
	grpc.ClientStream
}

type actorSystemCommunicationStreamHandlerClient struct {
	grpc.ClientStream
}

func (x *actorSystemCommunicationStreamHandlerClient) Send(m *DistributedMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *actorSystemCommunicationStreamHandlerClient) Recv() (*DistributedMessage, error) {
	m := new(DistributedMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ActorSystemCommunicationServer is the server API for ActorSystemCommunication service.
// All implementations must embed UnimplementedActorSystemCommunicationServer
// for forward compatibility
type ActorSystemCommunicationServer interface {
	StreamHandler(ActorSystemCommunication_StreamHandlerServer) error
	mustEmbedUnimplementedActorSystemCommunicationServer()
}

// UnimplementedActorSystemCommunicationServer must be embedded to have forward compatible implementations.
type UnimplementedActorSystemCommunicationServer struct {
}

func (UnimplementedActorSystemCommunicationServer) StreamHandler(ActorSystemCommunication_StreamHandlerServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamHandler not implemented")
}
func (UnimplementedActorSystemCommunicationServer) mustEmbedUnimplementedActorSystemCommunicationServer() {
}

// UnsafeActorSystemCommunicationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActorSystemCommunicationServer will
// result in compilation errors.
type UnsafeActorSystemCommunicationServer interface {
	mustEmbedUnimplementedActorSystemCommunicationServer()
}

func RegisterActorSystemCommunicationServer(s grpc.ServiceRegistrar, srv ActorSystemCommunicationServer) {
	s.RegisterService(&ActorSystemCommunication_ServiceDesc, srv)
}

func _ActorSystemCommunication_StreamHandler_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ActorSystemCommunicationServer).StreamHandler(&actorSystemCommunicationStreamHandlerServer{stream})
}

type ActorSystemCommunication_StreamHandlerServer interface {
	Send(*DistributedMessage) error
	Recv() (*DistributedMessage, error)
	grpc.ServerStream
}

type actorSystemCommunicationStreamHandlerServer struct {
	grpc.ServerStream
}

func (x *actorSystemCommunicationStreamHandlerServer) Send(m *DistributedMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *actorSystemCommunicationStreamHandlerServer) Recv() (*DistributedMessage, error) {
	m := new(DistributedMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ActorSystemCommunication_ServiceDesc is the grpc.ServiceDesc for ActorSystemCommunication service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ActorSystemCommunication_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "remote.ActorSystemCommunication",
	HandlerType: (*ActorSystemCommunicationServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamHandler",
			Handler:       _ActorSystemCommunication_StreamHandler_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "server.proto",
}
