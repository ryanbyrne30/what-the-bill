// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: consumer_us_bills.proto

package proto_consumer_us_bills

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

// ConsumerUSBillsClient is the client API for ConsumerUSBills service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConsumerUSBillsClient interface {
	SendUSBillCreated(ctx context.Context, in *USBillCreatedEvent, opts ...grpc.CallOption) (*EventPosted, error)
	SendUSBillUpdated(ctx context.Context, in *USBillUpdatedEvent, opts ...grpc.CallOption) (*EventPosted, error)
}

type consumerUSBillsClient struct {
	cc grpc.ClientConnInterface
}

func NewConsumerUSBillsClient(cc grpc.ClientConnInterface) ConsumerUSBillsClient {
	return &consumerUSBillsClient{cc}
}

func (c *consumerUSBillsClient) SendUSBillCreated(ctx context.Context, in *USBillCreatedEvent, opts ...grpc.CallOption) (*EventPosted, error) {
	out := new(EventPosted)
	err := c.cc.Invoke(ctx, "/proto_consumer_us_bills.ConsumerUSBills/SendUSBillCreated", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consumerUSBillsClient) SendUSBillUpdated(ctx context.Context, in *USBillUpdatedEvent, opts ...grpc.CallOption) (*EventPosted, error) {
	out := new(EventPosted)
	err := c.cc.Invoke(ctx, "/proto_consumer_us_bills.ConsumerUSBills/SendUSBillUpdated", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConsumerUSBillsServer is the server API for ConsumerUSBills service.
// All implementations must embed UnimplementedConsumerUSBillsServer
// for forward compatibility
type ConsumerUSBillsServer interface {
	SendUSBillCreated(context.Context, *USBillCreatedEvent) (*EventPosted, error)
	SendUSBillUpdated(context.Context, *USBillUpdatedEvent) (*EventPosted, error)
	mustEmbedUnimplementedConsumerUSBillsServer()
}

// UnimplementedConsumerUSBillsServer must be embedded to have forward compatible implementations.
type UnimplementedConsumerUSBillsServer struct {
}

func (UnimplementedConsumerUSBillsServer) SendUSBillCreated(context.Context, *USBillCreatedEvent) (*EventPosted, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendUSBillCreated not implemented")
}
func (UnimplementedConsumerUSBillsServer) SendUSBillUpdated(context.Context, *USBillUpdatedEvent) (*EventPosted, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendUSBillUpdated not implemented")
}
func (UnimplementedConsumerUSBillsServer) mustEmbedUnimplementedConsumerUSBillsServer() {}

// UnsafeConsumerUSBillsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConsumerUSBillsServer will
// result in compilation errors.
type UnsafeConsumerUSBillsServer interface {
	mustEmbedUnimplementedConsumerUSBillsServer()
}

func RegisterConsumerUSBillsServer(s grpc.ServiceRegistrar, srv ConsumerUSBillsServer) {
	s.RegisterService(&ConsumerUSBills_ServiceDesc, srv)
}

func _ConsumerUSBills_SendUSBillCreated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(USBillCreatedEvent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerUSBillsServer).SendUSBillCreated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_consumer_us_bills.ConsumerUSBills/SendUSBillCreated",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerUSBillsServer).SendUSBillCreated(ctx, req.(*USBillCreatedEvent))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsumerUSBills_SendUSBillUpdated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(USBillUpdatedEvent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsumerUSBillsServer).SendUSBillUpdated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_consumer_us_bills.ConsumerUSBills/SendUSBillUpdated",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsumerUSBillsServer).SendUSBillUpdated(ctx, req.(*USBillUpdatedEvent))
	}
	return interceptor(ctx, in, info, handler)
}

// ConsumerUSBills_ServiceDesc is the grpc.ServiceDesc for ConsumerUSBills service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConsumerUSBills_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto_consumer_us_bills.ConsumerUSBills",
	HandlerType: (*ConsumerUSBillsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendUSBillCreated",
			Handler:    _ConsumerUSBills_SendUSBillCreated_Handler,
		},
		{
			MethodName: "SendUSBillUpdated",
			Handler:    _ConsumerUSBills_SendUSBillUpdated_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "consumer_us_bills.proto",
}