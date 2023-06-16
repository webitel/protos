// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.6
// source: queue_resources.proto

package engine

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

const (
	QueueResourcesService_CreateQueueResourceGroup_FullMethodName = "/engine.QueueResourcesService/CreateQueueResourceGroup"
	QueueResourcesService_SearchQueueResourceGroup_FullMethodName = "/engine.QueueResourcesService/SearchQueueResourceGroup"
	QueueResourcesService_ReadQueueResourceGroup_FullMethodName   = "/engine.QueueResourcesService/ReadQueueResourceGroup"
	QueueResourcesService_UpdateQueueResourceGroup_FullMethodName = "/engine.QueueResourcesService/UpdateQueueResourceGroup"
	QueueResourcesService_DeleteQueueResourceGroup_FullMethodName = "/engine.QueueResourcesService/DeleteQueueResourceGroup"
)

// QueueResourcesServiceClient is the client API for QueueResourcesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueueResourcesServiceClient interface {
	// CreateQueueResourceGroup
	CreateQueueResourceGroup(ctx context.Context, in *CreateQueueResourceGroupRequest, opts ...grpc.CallOption) (*QueueResourceGroup, error)
	// SearchQueueResourceGroup
	SearchQueueResourceGroup(ctx context.Context, in *SearchQueueResourceGroupRequest, opts ...grpc.CallOption) (*ListQueueResourceGroup, error)
	// ReadQueueResourceGroup
	ReadQueueResourceGroup(ctx context.Context, in *ReadQueueResourceGroupRequest, opts ...grpc.CallOption) (*QueueResourceGroup, error)
	// UpdateQueueResourceGroup
	UpdateQueueResourceGroup(ctx context.Context, in *UpdateQueueResourceGroupRequest, opts ...grpc.CallOption) (*QueueResourceGroup, error)
	// DeleteQueueResourceGroup
	DeleteQueueResourceGroup(ctx context.Context, in *DeleteQueueResourceGroupRequest, opts ...grpc.CallOption) (*QueueResourceGroup, error)
}

type queueResourcesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQueueResourcesServiceClient(cc grpc.ClientConnInterface) QueueResourcesServiceClient {
	return &queueResourcesServiceClient{cc}
}

func (c *queueResourcesServiceClient) CreateQueueResourceGroup(ctx context.Context, in *CreateQueueResourceGroupRequest, opts ...grpc.CallOption) (*QueueResourceGroup, error) {
	out := new(QueueResourceGroup)
	err := c.cc.Invoke(ctx, QueueResourcesService_CreateQueueResourceGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueResourcesServiceClient) SearchQueueResourceGroup(ctx context.Context, in *SearchQueueResourceGroupRequest, opts ...grpc.CallOption) (*ListQueueResourceGroup, error) {
	out := new(ListQueueResourceGroup)
	err := c.cc.Invoke(ctx, QueueResourcesService_SearchQueueResourceGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueResourcesServiceClient) ReadQueueResourceGroup(ctx context.Context, in *ReadQueueResourceGroupRequest, opts ...grpc.CallOption) (*QueueResourceGroup, error) {
	out := new(QueueResourceGroup)
	err := c.cc.Invoke(ctx, QueueResourcesService_ReadQueueResourceGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueResourcesServiceClient) UpdateQueueResourceGroup(ctx context.Context, in *UpdateQueueResourceGroupRequest, opts ...grpc.CallOption) (*QueueResourceGroup, error) {
	out := new(QueueResourceGroup)
	err := c.cc.Invoke(ctx, QueueResourcesService_UpdateQueueResourceGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueResourcesServiceClient) DeleteQueueResourceGroup(ctx context.Context, in *DeleteQueueResourceGroupRequest, opts ...grpc.CallOption) (*QueueResourceGroup, error) {
	out := new(QueueResourceGroup)
	err := c.cc.Invoke(ctx, QueueResourcesService_DeleteQueueResourceGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueueResourcesServiceServer is the server API for QueueResourcesService service.
// All implementations must embed UnimplementedQueueResourcesServiceServer
// for forward compatibility
type QueueResourcesServiceServer interface {
	// CreateQueueResourceGroup
	CreateQueueResourceGroup(context.Context, *CreateQueueResourceGroupRequest) (*QueueResourceGroup, error)
	// SearchQueueResourceGroup
	SearchQueueResourceGroup(context.Context, *SearchQueueResourceGroupRequest) (*ListQueueResourceGroup, error)
	// ReadQueueResourceGroup
	ReadQueueResourceGroup(context.Context, *ReadQueueResourceGroupRequest) (*QueueResourceGroup, error)
	// UpdateQueueResourceGroup
	UpdateQueueResourceGroup(context.Context, *UpdateQueueResourceGroupRequest) (*QueueResourceGroup, error)
	// DeleteQueueResourceGroup
	DeleteQueueResourceGroup(context.Context, *DeleteQueueResourceGroupRequest) (*QueueResourceGroup, error)
	mustEmbedUnimplementedQueueResourcesServiceServer()
}

// UnimplementedQueueResourcesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedQueueResourcesServiceServer struct {
}

func (UnimplementedQueueResourcesServiceServer) CreateQueueResourceGroup(context.Context, *CreateQueueResourceGroupRequest) (*QueueResourceGroup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateQueueResourceGroup not implemented")
}
func (UnimplementedQueueResourcesServiceServer) SearchQueueResourceGroup(context.Context, *SearchQueueResourceGroupRequest) (*ListQueueResourceGroup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchQueueResourceGroup not implemented")
}
func (UnimplementedQueueResourcesServiceServer) ReadQueueResourceGroup(context.Context, *ReadQueueResourceGroupRequest) (*QueueResourceGroup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadQueueResourceGroup not implemented")
}
func (UnimplementedQueueResourcesServiceServer) UpdateQueueResourceGroup(context.Context, *UpdateQueueResourceGroupRequest) (*QueueResourceGroup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateQueueResourceGroup not implemented")
}
func (UnimplementedQueueResourcesServiceServer) DeleteQueueResourceGroup(context.Context, *DeleteQueueResourceGroupRequest) (*QueueResourceGroup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQueueResourceGroup not implemented")
}
func (UnimplementedQueueResourcesServiceServer) mustEmbedUnimplementedQueueResourcesServiceServer() {}

// UnsafeQueueResourcesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueueResourcesServiceServer will
// result in compilation errors.
type UnsafeQueueResourcesServiceServer interface {
	mustEmbedUnimplementedQueueResourcesServiceServer()
}

func RegisterQueueResourcesServiceServer(s grpc.ServiceRegistrar, srv QueueResourcesServiceServer) {
	s.RegisterService(&QueueResourcesService_ServiceDesc, srv)
}

func _QueueResourcesService_CreateQueueResourceGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateQueueResourceGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueResourcesServiceServer).CreateQueueResourceGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QueueResourcesService_CreateQueueResourceGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueResourcesServiceServer).CreateQueueResourceGroup(ctx, req.(*CreateQueueResourceGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueueResourcesService_SearchQueueResourceGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchQueueResourceGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueResourcesServiceServer).SearchQueueResourceGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QueueResourcesService_SearchQueueResourceGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueResourcesServiceServer).SearchQueueResourceGroup(ctx, req.(*SearchQueueResourceGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueueResourcesService_ReadQueueResourceGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadQueueResourceGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueResourcesServiceServer).ReadQueueResourceGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QueueResourcesService_ReadQueueResourceGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueResourcesServiceServer).ReadQueueResourceGroup(ctx, req.(*ReadQueueResourceGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueueResourcesService_UpdateQueueResourceGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateQueueResourceGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueResourcesServiceServer).UpdateQueueResourceGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QueueResourcesService_UpdateQueueResourceGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueResourcesServiceServer).UpdateQueueResourceGroup(ctx, req.(*UpdateQueueResourceGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueueResourcesService_DeleteQueueResourceGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteQueueResourceGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueResourcesServiceServer).DeleteQueueResourceGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QueueResourcesService_DeleteQueueResourceGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueResourcesServiceServer).DeleteQueueResourceGroup(ctx, req.(*DeleteQueueResourceGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QueueResourcesService_ServiceDesc is the grpc.ServiceDesc for QueueResourcesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QueueResourcesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "engine.QueueResourcesService",
	HandlerType: (*QueueResourcesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateQueueResourceGroup",
			Handler:    _QueueResourcesService_CreateQueueResourceGroup_Handler,
		},
		{
			MethodName: "SearchQueueResourceGroup",
			Handler:    _QueueResourcesService_SearchQueueResourceGroup_Handler,
		},
		{
			MethodName: "ReadQueueResourceGroup",
			Handler:    _QueueResourcesService_ReadQueueResourceGroup_Handler,
		},
		{
			MethodName: "UpdateQueueResourceGroup",
			Handler:    _QueueResourcesService_UpdateQueueResourceGroup_Handler,
		},
		{
			MethodName: "DeleteQueueResourceGroup",
			Handler:    _QueueResourcesService_DeleteQueueResourceGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "queue_resources.proto",
}
