// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.0
// source: outbound_resource.proto

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

// OutboundResourceServiceClient is the client API for OutboundResourceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OutboundResourceServiceClient interface {
	// Create OutboundResource
	CreateOutboundResource(ctx context.Context, in *CreateOutboundResourceRequest, opts ...grpc.CallOption) (*OutboundResource, error)
	// List of OutboundResource
	SearchOutboundResource(ctx context.Context, in *SearchOutboundResourceRequest, opts ...grpc.CallOption) (*ListOutboundResource, error)
	// OutboundResource item
	ReadOutboundResource(ctx context.Context, in *ReadOutboundResourceRequest, opts ...grpc.CallOption) (*OutboundResource, error)
	// Update OutboundResource
	UpdateOutboundResource(ctx context.Context, in *UpdateOutboundResourceRequest, opts ...grpc.CallOption) (*OutboundResource, error)
	// Patch OutboundResource
	PatchOutboundResource(ctx context.Context, in *PatchOutboundResourceRequest, opts ...grpc.CallOption) (*OutboundResource, error)
	// Remove OutboundResource
	DeleteOutboundResource(ctx context.Context, in *DeleteOutboundResourceRequest, opts ...grpc.CallOption) (*OutboundResource, error)
	// Create CreateOutboundResourceDisplay
	CreateOutboundResourceDisplay(ctx context.Context, in *CreateOutboundResourceDisplayRequest, opts ...grpc.CallOption) (*ResourceDisplay, error)
	// List of ResourceDisplay
	SearchOutboundResourceDisplay(ctx context.Context, in *SearchOutboundResourceDisplayRequest, opts ...grpc.CallOption) (*ListOutboundResourceDisplay, error)
	// ResourceDisplay item
	ReadOutboundResourceDisplay(ctx context.Context, in *ReadOutboundResourceDisplayRequest, opts ...grpc.CallOption) (*ResourceDisplay, error)
	// Update ResourceDisplay
	UpdateOutboundResourceDisplay(ctx context.Context, in *UpdateOutboundResourceDisplayRequest, opts ...grpc.CallOption) (*ResourceDisplay, error)
	// Remove ResourceDisplay
	DeleteOutboundResourceDisplay(ctx context.Context, in *DeleteOutboundResourceDisplayRequest, opts ...grpc.CallOption) (*ResourceDisplay, error)
}

type outboundResourceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOutboundResourceServiceClient(cc grpc.ClientConnInterface) OutboundResourceServiceClient {
	return &outboundResourceServiceClient{cc}
}

func (c *outboundResourceServiceClient) CreateOutboundResource(ctx context.Context, in *CreateOutboundResourceRequest, opts ...grpc.CallOption) (*OutboundResource, error) {
	out := new(OutboundResource)
	err := c.cc.Invoke(ctx, "/engine.OutboundResourceService/CreateOutboundResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *outboundResourceServiceClient) SearchOutboundResource(ctx context.Context, in *SearchOutboundResourceRequest, opts ...grpc.CallOption) (*ListOutboundResource, error) {
	out := new(ListOutboundResource)
	err := c.cc.Invoke(ctx, "/engine.OutboundResourceService/SearchOutboundResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *outboundResourceServiceClient) ReadOutboundResource(ctx context.Context, in *ReadOutboundResourceRequest, opts ...grpc.CallOption) (*OutboundResource, error) {
	out := new(OutboundResource)
	err := c.cc.Invoke(ctx, "/engine.OutboundResourceService/ReadOutboundResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *outboundResourceServiceClient) UpdateOutboundResource(ctx context.Context, in *UpdateOutboundResourceRequest, opts ...grpc.CallOption) (*OutboundResource, error) {
	out := new(OutboundResource)
	err := c.cc.Invoke(ctx, "/engine.OutboundResourceService/UpdateOutboundResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *outboundResourceServiceClient) PatchOutboundResource(ctx context.Context, in *PatchOutboundResourceRequest, opts ...grpc.CallOption) (*OutboundResource, error) {
	out := new(OutboundResource)
	err := c.cc.Invoke(ctx, "/engine.OutboundResourceService/PatchOutboundResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *outboundResourceServiceClient) DeleteOutboundResource(ctx context.Context, in *DeleteOutboundResourceRequest, opts ...grpc.CallOption) (*OutboundResource, error) {
	out := new(OutboundResource)
	err := c.cc.Invoke(ctx, "/engine.OutboundResourceService/DeleteOutboundResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *outboundResourceServiceClient) CreateOutboundResourceDisplay(ctx context.Context, in *CreateOutboundResourceDisplayRequest, opts ...grpc.CallOption) (*ResourceDisplay, error) {
	out := new(ResourceDisplay)
	err := c.cc.Invoke(ctx, "/engine.OutboundResourceService/CreateOutboundResourceDisplay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *outboundResourceServiceClient) SearchOutboundResourceDisplay(ctx context.Context, in *SearchOutboundResourceDisplayRequest, opts ...grpc.CallOption) (*ListOutboundResourceDisplay, error) {
	out := new(ListOutboundResourceDisplay)
	err := c.cc.Invoke(ctx, "/engine.OutboundResourceService/SearchOutboundResourceDisplay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *outboundResourceServiceClient) ReadOutboundResourceDisplay(ctx context.Context, in *ReadOutboundResourceDisplayRequest, opts ...grpc.CallOption) (*ResourceDisplay, error) {
	out := new(ResourceDisplay)
	err := c.cc.Invoke(ctx, "/engine.OutboundResourceService/ReadOutboundResourceDisplay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *outboundResourceServiceClient) UpdateOutboundResourceDisplay(ctx context.Context, in *UpdateOutboundResourceDisplayRequest, opts ...grpc.CallOption) (*ResourceDisplay, error) {
	out := new(ResourceDisplay)
	err := c.cc.Invoke(ctx, "/engine.OutboundResourceService/UpdateOutboundResourceDisplay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *outboundResourceServiceClient) DeleteOutboundResourceDisplay(ctx context.Context, in *DeleteOutboundResourceDisplayRequest, opts ...grpc.CallOption) (*ResourceDisplay, error) {
	out := new(ResourceDisplay)
	err := c.cc.Invoke(ctx, "/engine.OutboundResourceService/DeleteOutboundResourceDisplay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OutboundResourceServiceServer is the server API for OutboundResourceService service.
// All implementations must embed UnimplementedOutboundResourceServiceServer
// for forward compatibility
type OutboundResourceServiceServer interface {
	// Create OutboundResource
	CreateOutboundResource(context.Context, *CreateOutboundResourceRequest) (*OutboundResource, error)
	// List of OutboundResource
	SearchOutboundResource(context.Context, *SearchOutboundResourceRequest) (*ListOutboundResource, error)
	// OutboundResource item
	ReadOutboundResource(context.Context, *ReadOutboundResourceRequest) (*OutboundResource, error)
	// Update OutboundResource
	UpdateOutboundResource(context.Context, *UpdateOutboundResourceRequest) (*OutboundResource, error)
	// Patch OutboundResource
	PatchOutboundResource(context.Context, *PatchOutboundResourceRequest) (*OutboundResource, error)
	// Remove OutboundResource
	DeleteOutboundResource(context.Context, *DeleteOutboundResourceRequest) (*OutboundResource, error)
	// Create CreateOutboundResourceDisplay
	CreateOutboundResourceDisplay(context.Context, *CreateOutboundResourceDisplayRequest) (*ResourceDisplay, error)
	// List of ResourceDisplay
	SearchOutboundResourceDisplay(context.Context, *SearchOutboundResourceDisplayRequest) (*ListOutboundResourceDisplay, error)
	// ResourceDisplay item
	ReadOutboundResourceDisplay(context.Context, *ReadOutboundResourceDisplayRequest) (*ResourceDisplay, error)
	// Update ResourceDisplay
	UpdateOutboundResourceDisplay(context.Context, *UpdateOutboundResourceDisplayRequest) (*ResourceDisplay, error)
	// Remove ResourceDisplay
	DeleteOutboundResourceDisplay(context.Context, *DeleteOutboundResourceDisplayRequest) (*ResourceDisplay, error)
	mustEmbedUnimplementedOutboundResourceServiceServer()
}

// UnimplementedOutboundResourceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOutboundResourceServiceServer struct {
}

func (UnimplementedOutboundResourceServiceServer) CreateOutboundResource(context.Context, *CreateOutboundResourceRequest) (*OutboundResource, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOutboundResource not implemented")
}
func (UnimplementedOutboundResourceServiceServer) SearchOutboundResource(context.Context, *SearchOutboundResourceRequest) (*ListOutboundResource, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchOutboundResource not implemented")
}
func (UnimplementedOutboundResourceServiceServer) ReadOutboundResource(context.Context, *ReadOutboundResourceRequest) (*OutboundResource, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadOutboundResource not implemented")
}
func (UnimplementedOutboundResourceServiceServer) UpdateOutboundResource(context.Context, *UpdateOutboundResourceRequest) (*OutboundResource, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOutboundResource not implemented")
}
func (UnimplementedOutboundResourceServiceServer) PatchOutboundResource(context.Context, *PatchOutboundResourceRequest) (*OutboundResource, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchOutboundResource not implemented")
}
func (UnimplementedOutboundResourceServiceServer) DeleteOutboundResource(context.Context, *DeleteOutboundResourceRequest) (*OutboundResource, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOutboundResource not implemented")
}
func (UnimplementedOutboundResourceServiceServer) CreateOutboundResourceDisplay(context.Context, *CreateOutboundResourceDisplayRequest) (*ResourceDisplay, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOutboundResourceDisplay not implemented")
}
func (UnimplementedOutboundResourceServiceServer) SearchOutboundResourceDisplay(context.Context, *SearchOutboundResourceDisplayRequest) (*ListOutboundResourceDisplay, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchOutboundResourceDisplay not implemented")
}
func (UnimplementedOutboundResourceServiceServer) ReadOutboundResourceDisplay(context.Context, *ReadOutboundResourceDisplayRequest) (*ResourceDisplay, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadOutboundResourceDisplay not implemented")
}
func (UnimplementedOutboundResourceServiceServer) UpdateOutboundResourceDisplay(context.Context, *UpdateOutboundResourceDisplayRequest) (*ResourceDisplay, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOutboundResourceDisplay not implemented")
}
func (UnimplementedOutboundResourceServiceServer) DeleteOutboundResourceDisplay(context.Context, *DeleteOutboundResourceDisplayRequest) (*ResourceDisplay, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOutboundResourceDisplay not implemented")
}
func (UnimplementedOutboundResourceServiceServer) mustEmbedUnimplementedOutboundResourceServiceServer() {
}

// UnsafeOutboundResourceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OutboundResourceServiceServer will
// result in compilation errors.
type UnsafeOutboundResourceServiceServer interface {
	mustEmbedUnimplementedOutboundResourceServiceServer()
}

func RegisterOutboundResourceServiceServer(s grpc.ServiceRegistrar, srv OutboundResourceServiceServer) {
	s.RegisterService(&OutboundResourceService_ServiceDesc, srv)
}

func _OutboundResourceService_CreateOutboundResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOutboundResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OutboundResourceServiceServer).CreateOutboundResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.OutboundResourceService/CreateOutboundResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OutboundResourceServiceServer).CreateOutboundResource(ctx, req.(*CreateOutboundResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OutboundResourceService_SearchOutboundResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchOutboundResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OutboundResourceServiceServer).SearchOutboundResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.OutboundResourceService/SearchOutboundResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OutboundResourceServiceServer).SearchOutboundResource(ctx, req.(*SearchOutboundResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OutboundResourceService_ReadOutboundResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadOutboundResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OutboundResourceServiceServer).ReadOutboundResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.OutboundResourceService/ReadOutboundResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OutboundResourceServiceServer).ReadOutboundResource(ctx, req.(*ReadOutboundResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OutboundResourceService_UpdateOutboundResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOutboundResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OutboundResourceServiceServer).UpdateOutboundResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.OutboundResourceService/UpdateOutboundResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OutboundResourceServiceServer).UpdateOutboundResource(ctx, req.(*UpdateOutboundResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OutboundResourceService_PatchOutboundResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchOutboundResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OutboundResourceServiceServer).PatchOutboundResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.OutboundResourceService/PatchOutboundResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OutboundResourceServiceServer).PatchOutboundResource(ctx, req.(*PatchOutboundResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OutboundResourceService_DeleteOutboundResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOutboundResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OutboundResourceServiceServer).DeleteOutboundResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.OutboundResourceService/DeleteOutboundResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OutboundResourceServiceServer).DeleteOutboundResource(ctx, req.(*DeleteOutboundResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OutboundResourceService_CreateOutboundResourceDisplay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOutboundResourceDisplayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OutboundResourceServiceServer).CreateOutboundResourceDisplay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.OutboundResourceService/CreateOutboundResourceDisplay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OutboundResourceServiceServer).CreateOutboundResourceDisplay(ctx, req.(*CreateOutboundResourceDisplayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OutboundResourceService_SearchOutboundResourceDisplay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchOutboundResourceDisplayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OutboundResourceServiceServer).SearchOutboundResourceDisplay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.OutboundResourceService/SearchOutboundResourceDisplay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OutboundResourceServiceServer).SearchOutboundResourceDisplay(ctx, req.(*SearchOutboundResourceDisplayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OutboundResourceService_ReadOutboundResourceDisplay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadOutboundResourceDisplayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OutboundResourceServiceServer).ReadOutboundResourceDisplay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.OutboundResourceService/ReadOutboundResourceDisplay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OutboundResourceServiceServer).ReadOutboundResourceDisplay(ctx, req.(*ReadOutboundResourceDisplayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OutboundResourceService_UpdateOutboundResourceDisplay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOutboundResourceDisplayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OutboundResourceServiceServer).UpdateOutboundResourceDisplay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.OutboundResourceService/UpdateOutboundResourceDisplay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OutboundResourceServiceServer).UpdateOutboundResourceDisplay(ctx, req.(*UpdateOutboundResourceDisplayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OutboundResourceService_DeleteOutboundResourceDisplay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOutboundResourceDisplayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OutboundResourceServiceServer).DeleteOutboundResourceDisplay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.OutboundResourceService/DeleteOutboundResourceDisplay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OutboundResourceServiceServer).DeleteOutboundResourceDisplay(ctx, req.(*DeleteOutboundResourceDisplayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OutboundResourceService_ServiceDesc is the grpc.ServiceDesc for OutboundResourceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OutboundResourceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "engine.OutboundResourceService",
	HandlerType: (*OutboundResourceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOutboundResource",
			Handler:    _OutboundResourceService_CreateOutboundResource_Handler,
		},
		{
			MethodName: "SearchOutboundResource",
			Handler:    _OutboundResourceService_SearchOutboundResource_Handler,
		},
		{
			MethodName: "ReadOutboundResource",
			Handler:    _OutboundResourceService_ReadOutboundResource_Handler,
		},
		{
			MethodName: "UpdateOutboundResource",
			Handler:    _OutboundResourceService_UpdateOutboundResource_Handler,
		},
		{
			MethodName: "PatchOutboundResource",
			Handler:    _OutboundResourceService_PatchOutboundResource_Handler,
		},
		{
			MethodName: "DeleteOutboundResource",
			Handler:    _OutboundResourceService_DeleteOutboundResource_Handler,
		},
		{
			MethodName: "CreateOutboundResourceDisplay",
			Handler:    _OutboundResourceService_CreateOutboundResourceDisplay_Handler,
		},
		{
			MethodName: "SearchOutboundResourceDisplay",
			Handler:    _OutboundResourceService_SearchOutboundResourceDisplay_Handler,
		},
		{
			MethodName: "ReadOutboundResourceDisplay",
			Handler:    _OutboundResourceService_ReadOutboundResourceDisplay_Handler,
		},
		{
			MethodName: "UpdateOutboundResourceDisplay",
			Handler:    _OutboundResourceService_UpdateOutboundResourceDisplay_Handler,
		},
		{
			MethodName: "DeleteOutboundResourceDisplay",
			Handler:    _OutboundResourceService_DeleteOutboundResourceDisplay_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "outbound_resource.proto",
}