// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: queue.proto

package proto_queue

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

// QueueClient is the client API for Queue service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueueClient interface {
	PostUSBillUpdatedEvent(ctx context.Context, in *USBillUpdatedEvent, opts ...grpc.CallOption) (*EventPosted, error)
	PostUSBillCreatedEvent(ctx context.Context, in *USBillCreatedEvent, opts ...grpc.CallOption) (*EventPosted, error)
}

type queueClient struct {
	cc grpc.ClientConnInterface
}

func NewQueueClient(cc grpc.ClientConnInterface) QueueClient {
	return &queueClient{cc}
}

func (c *queueClient) PostUSBillUpdatedEvent(ctx context.Context, in *USBillUpdatedEvent, opts ...grpc.CallOption) (*EventPosted, error) {
	out := new(EventPosted)
	err := c.cc.Invoke(ctx, "/proto_queue.Queue/PostUSBillUpdatedEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueClient) PostUSBillCreatedEvent(ctx context.Context, in *USBillCreatedEvent, opts ...grpc.CallOption) (*EventPosted, error) {
	out := new(EventPosted)
	err := c.cc.Invoke(ctx, "/proto_queue.Queue/PostUSBillCreatedEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueueServer is the server API for Queue service.
// All implementations must embed UnimplementedQueueServer
// for forward compatibility
type QueueServer interface {
	PostUSBillUpdatedEvent(context.Context, *USBillUpdatedEvent) (*EventPosted, error)
	PostUSBillCreatedEvent(context.Context, *USBillCreatedEvent) (*EventPosted, error)
	mustEmbedUnimplementedQueueServer()
}

// UnimplementedQueueServer must be embedded to have forward compatible implementations.
type UnimplementedQueueServer struct {
}

func (UnimplementedQueueServer) PostUSBillUpdatedEvent(context.Context, *USBillUpdatedEvent) (*EventPosted, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostUSBillUpdatedEvent not implemented")
}
func (UnimplementedQueueServer) PostUSBillCreatedEvent(context.Context, *USBillCreatedEvent) (*EventPosted, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostUSBillCreatedEvent not implemented")
}
func (UnimplementedQueueServer) mustEmbedUnimplementedQueueServer() {}

// UnsafeQueueServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueueServer will
// result in compilation errors.
type UnsafeQueueServer interface {
	mustEmbedUnimplementedQueueServer()
}

func RegisterQueueServer(s grpc.ServiceRegistrar, srv QueueServer) {
	s.RegisterService(&Queue_ServiceDesc, srv)
}

func _Queue_PostUSBillUpdatedEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(USBillUpdatedEvent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueServer).PostUSBillUpdatedEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_queue.Queue/PostUSBillUpdatedEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueServer).PostUSBillUpdatedEvent(ctx, req.(*USBillUpdatedEvent))
	}
	return interceptor(ctx, in, info, handler)
}

func _Queue_PostUSBillCreatedEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(USBillCreatedEvent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueServer).PostUSBillCreatedEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_queue.Queue/PostUSBillCreatedEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueServer).PostUSBillCreatedEvent(ctx, req.(*USBillCreatedEvent))
	}
	return interceptor(ctx, in, info, handler)
}

// Queue_ServiceDesc is the grpc.ServiceDesc for Queue service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Queue_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto_queue.Queue",
	HandlerType: (*QueueServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostUSBillUpdatedEvent",
			Handler:    _Queue_PostUSBillUpdatedEvent_Handler,
		},
		{
			MethodName: "PostUSBillCreatedEvent",
			Handler:    _Queue_PostUSBillCreatedEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "queue.proto",
}
