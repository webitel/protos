// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.6
// source: queue_bucket.proto

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
	QueueBucketService_CreateQueueBucket_FullMethodName = "/engine.QueueBucketService/CreateQueueBucket"
	QueueBucketService_SearchQueueBucket_FullMethodName = "/engine.QueueBucketService/SearchQueueBucket"
	QueueBucketService_ReadQueueBucket_FullMethodName   = "/engine.QueueBucketService/ReadQueueBucket"
	QueueBucketService_UpdateQueueBucket_FullMethodName = "/engine.QueueBucketService/UpdateQueueBucket"
	QueueBucketService_PatchQueueBucket_FullMethodName  = "/engine.QueueBucketService/PatchQueueBucket"
	QueueBucketService_DeleteQueueBucket_FullMethodName = "/engine.QueueBucketService/DeleteQueueBucket"
)

// QueueBucketServiceClient is the client API for QueueBucketService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueueBucketServiceClient interface {
	// Create QueueBucket
	CreateQueueBucket(ctx context.Context, in *CreateQueueBucketRequest, opts ...grpc.CallOption) (*QueueBucket, error)
	// SearchQueueRouting
	SearchQueueBucket(ctx context.Context, in *SearchQueueBucketRequest, opts ...grpc.CallOption) (*ListQueueBucket, error)
	// ReadQueueRouting
	ReadQueueBucket(ctx context.Context, in *ReadQueueBucketRequest, opts ...grpc.CallOption) (*QueueBucket, error)
	// UpdateQueueBucket
	UpdateQueueBucket(ctx context.Context, in *UpdateQueueBucketRequest, opts ...grpc.CallOption) (*QueueBucket, error)
	PatchQueueBucket(ctx context.Context, in *PatchQueueBucketRequest, opts ...grpc.CallOption) (*QueueBucket, error)
	// DeleteQueueRouting
	DeleteQueueBucket(ctx context.Context, in *DeleteQueueBucketRequest, opts ...grpc.CallOption) (*QueueBucket, error)
}

type queueBucketServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQueueBucketServiceClient(cc grpc.ClientConnInterface) QueueBucketServiceClient {
	return &queueBucketServiceClient{cc}
}

func (c *queueBucketServiceClient) CreateQueueBucket(ctx context.Context, in *CreateQueueBucketRequest, opts ...grpc.CallOption) (*QueueBucket, error) {
	out := new(QueueBucket)
	err := c.cc.Invoke(ctx, QueueBucketService_CreateQueueBucket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueBucketServiceClient) SearchQueueBucket(ctx context.Context, in *SearchQueueBucketRequest, opts ...grpc.CallOption) (*ListQueueBucket, error) {
	out := new(ListQueueBucket)
	err := c.cc.Invoke(ctx, QueueBucketService_SearchQueueBucket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueBucketServiceClient) ReadQueueBucket(ctx context.Context, in *ReadQueueBucketRequest, opts ...grpc.CallOption) (*QueueBucket, error) {
	out := new(QueueBucket)
	err := c.cc.Invoke(ctx, QueueBucketService_ReadQueueBucket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueBucketServiceClient) UpdateQueueBucket(ctx context.Context, in *UpdateQueueBucketRequest, opts ...grpc.CallOption) (*QueueBucket, error) {
	out := new(QueueBucket)
	err := c.cc.Invoke(ctx, QueueBucketService_UpdateQueueBucket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueBucketServiceClient) PatchQueueBucket(ctx context.Context, in *PatchQueueBucketRequest, opts ...grpc.CallOption) (*QueueBucket, error) {
	out := new(QueueBucket)
	err := c.cc.Invoke(ctx, QueueBucketService_PatchQueueBucket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueBucketServiceClient) DeleteQueueBucket(ctx context.Context, in *DeleteQueueBucketRequest, opts ...grpc.CallOption) (*QueueBucket, error) {
	out := new(QueueBucket)
	err := c.cc.Invoke(ctx, QueueBucketService_DeleteQueueBucket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueueBucketServiceServer is the server API for QueueBucketService service.
// All implementations must embed UnimplementedQueueBucketServiceServer
// for forward compatibility
type QueueBucketServiceServer interface {
	// Create QueueBucket
	CreateQueueBucket(context.Context, *CreateQueueBucketRequest) (*QueueBucket, error)
	// SearchQueueRouting
	SearchQueueBucket(context.Context, *SearchQueueBucketRequest) (*ListQueueBucket, error)
	// ReadQueueRouting
	ReadQueueBucket(context.Context, *ReadQueueBucketRequest) (*QueueBucket, error)
	// UpdateQueueBucket
	UpdateQueueBucket(context.Context, *UpdateQueueBucketRequest) (*QueueBucket, error)
	PatchQueueBucket(context.Context, *PatchQueueBucketRequest) (*QueueBucket, error)
	// DeleteQueueRouting
	DeleteQueueBucket(context.Context, *DeleteQueueBucketRequest) (*QueueBucket, error)
	mustEmbedUnimplementedQueueBucketServiceServer()
}

// UnimplementedQueueBucketServiceServer must be embedded to have forward compatible implementations.
type UnimplementedQueueBucketServiceServer struct {
}

func (UnimplementedQueueBucketServiceServer) CreateQueueBucket(context.Context, *CreateQueueBucketRequest) (*QueueBucket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateQueueBucket not implemented")
}
func (UnimplementedQueueBucketServiceServer) SearchQueueBucket(context.Context, *SearchQueueBucketRequest) (*ListQueueBucket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchQueueBucket not implemented")
}
func (UnimplementedQueueBucketServiceServer) ReadQueueBucket(context.Context, *ReadQueueBucketRequest) (*QueueBucket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadQueueBucket not implemented")
}
func (UnimplementedQueueBucketServiceServer) UpdateQueueBucket(context.Context, *UpdateQueueBucketRequest) (*QueueBucket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateQueueBucket not implemented")
}
func (UnimplementedQueueBucketServiceServer) PatchQueueBucket(context.Context, *PatchQueueBucketRequest) (*QueueBucket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchQueueBucket not implemented")
}
func (UnimplementedQueueBucketServiceServer) DeleteQueueBucket(context.Context, *DeleteQueueBucketRequest) (*QueueBucket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQueueBucket not implemented")
}
func (UnimplementedQueueBucketServiceServer) mustEmbedUnimplementedQueueBucketServiceServer() {}

// UnsafeQueueBucketServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueueBucketServiceServer will
// result in compilation errors.
type UnsafeQueueBucketServiceServer interface {
	mustEmbedUnimplementedQueueBucketServiceServer()
}

func RegisterQueueBucketServiceServer(s grpc.ServiceRegistrar, srv QueueBucketServiceServer) {
	s.RegisterService(&QueueBucketService_ServiceDesc, srv)
}

func _QueueBucketService_CreateQueueBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateQueueBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueBucketServiceServer).CreateQueueBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QueueBucketService_CreateQueueBucket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueBucketServiceServer).CreateQueueBucket(ctx, req.(*CreateQueueBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueueBucketService_SearchQueueBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchQueueBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueBucketServiceServer).SearchQueueBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QueueBucketService_SearchQueueBucket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueBucketServiceServer).SearchQueueBucket(ctx, req.(*SearchQueueBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueueBucketService_ReadQueueBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadQueueBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueBucketServiceServer).ReadQueueBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QueueBucketService_ReadQueueBucket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueBucketServiceServer).ReadQueueBucket(ctx, req.(*ReadQueueBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueueBucketService_UpdateQueueBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateQueueBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueBucketServiceServer).UpdateQueueBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QueueBucketService_UpdateQueueBucket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueBucketServiceServer).UpdateQueueBucket(ctx, req.(*UpdateQueueBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueueBucketService_PatchQueueBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchQueueBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueBucketServiceServer).PatchQueueBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QueueBucketService_PatchQueueBucket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueBucketServiceServer).PatchQueueBucket(ctx, req.(*PatchQueueBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueueBucketService_DeleteQueueBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteQueueBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueBucketServiceServer).DeleteQueueBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QueueBucketService_DeleteQueueBucket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueBucketServiceServer).DeleteQueueBucket(ctx, req.(*DeleteQueueBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QueueBucketService_ServiceDesc is the grpc.ServiceDesc for QueueBucketService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QueueBucketService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "engine.QueueBucketService",
	HandlerType: (*QueueBucketServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateQueueBucket",
			Handler:    _QueueBucketService_CreateQueueBucket_Handler,
		},
		{
			MethodName: "SearchQueueBucket",
			Handler:    _QueueBucketService_SearchQueueBucket_Handler,
		},
		{
			MethodName: "ReadQueueBucket",
			Handler:    _QueueBucketService_ReadQueueBucket_Handler,
		},
		{
			MethodName: "UpdateQueueBucket",
			Handler:    _QueueBucketService_UpdateQueueBucket_Handler,
		},
		{
			MethodName: "PatchQueueBucket",
			Handler:    _QueueBucketService_PatchQueueBucket_Handler,
		},
		{
			MethodName: "DeleteQueueBucket",
			Handler:    _QueueBucketService_DeleteQueueBucket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "queue_bucket.proto",
}
