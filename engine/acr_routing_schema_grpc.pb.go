// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.0
// source: acr_routing_schema.proto

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

// RoutingSchemaServiceClient is the client API for RoutingSchemaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoutingSchemaServiceClient interface {
	// Create RoutingSchema
	CreateRoutingSchema(ctx context.Context, in *CreateRoutingSchemaRequest, opts ...grpc.CallOption) (*RoutingSchema, error)
	// List RoutingSchema
	SearchRoutingSchema(ctx context.Context, in *SearchRoutingSchemaRequest, opts ...grpc.CallOption) (*ListRoutingSchema, error)
	// RoutingSchema item
	ReadRoutingSchema(ctx context.Context, in *ReadRoutingSchemaRequest, opts ...grpc.CallOption) (*RoutingSchema, error)
	// Update RoutingSchema
	UpdateRoutingSchema(ctx context.Context, in *UpdateRoutingSchemaRequest, opts ...grpc.CallOption) (*RoutingSchema, error)
	// Patch RoutingSchema
	PatchRoutingSchema(ctx context.Context, in *PatchRoutingSchemaRequest, opts ...grpc.CallOption) (*RoutingSchema, error)
	// Remove RoutingSchema
	DeleteRoutingSchema(ctx context.Context, in *DeleteRoutingSchemaRequest, opts ...grpc.CallOption) (*RoutingSchema, error)
	// List RoutingSchemaTags
	SearchRoutingSchemaTags(ctx context.Context, in *SearchRoutingSchemaTagsRequest, opts ...grpc.CallOption) (*ListRoutingSchemaTags, error)
}

type routingSchemaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRoutingSchemaServiceClient(cc grpc.ClientConnInterface) RoutingSchemaServiceClient {
	return &routingSchemaServiceClient{cc}
}

func (c *routingSchemaServiceClient) CreateRoutingSchema(ctx context.Context, in *CreateRoutingSchemaRequest, opts ...grpc.CallOption) (*RoutingSchema, error) {
	out := new(RoutingSchema)
	err := c.cc.Invoke(ctx, "/engine.RoutingSchemaService/CreateRoutingSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routingSchemaServiceClient) SearchRoutingSchema(ctx context.Context, in *SearchRoutingSchemaRequest, opts ...grpc.CallOption) (*ListRoutingSchema, error) {
	out := new(ListRoutingSchema)
	err := c.cc.Invoke(ctx, "/engine.RoutingSchemaService/SearchRoutingSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routingSchemaServiceClient) ReadRoutingSchema(ctx context.Context, in *ReadRoutingSchemaRequest, opts ...grpc.CallOption) (*RoutingSchema, error) {
	out := new(RoutingSchema)
	err := c.cc.Invoke(ctx, "/engine.RoutingSchemaService/ReadRoutingSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routingSchemaServiceClient) UpdateRoutingSchema(ctx context.Context, in *UpdateRoutingSchemaRequest, opts ...grpc.CallOption) (*RoutingSchema, error) {
	out := new(RoutingSchema)
	err := c.cc.Invoke(ctx, "/engine.RoutingSchemaService/UpdateRoutingSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routingSchemaServiceClient) PatchRoutingSchema(ctx context.Context, in *PatchRoutingSchemaRequest, opts ...grpc.CallOption) (*RoutingSchema, error) {
	out := new(RoutingSchema)
	err := c.cc.Invoke(ctx, "/engine.RoutingSchemaService/PatchRoutingSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routingSchemaServiceClient) DeleteRoutingSchema(ctx context.Context, in *DeleteRoutingSchemaRequest, opts ...grpc.CallOption) (*RoutingSchema, error) {
	out := new(RoutingSchema)
	err := c.cc.Invoke(ctx, "/engine.RoutingSchemaService/DeleteRoutingSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routingSchemaServiceClient) SearchRoutingSchemaTags(ctx context.Context, in *SearchRoutingSchemaTagsRequest, opts ...grpc.CallOption) (*ListRoutingSchemaTags, error) {
	out := new(ListRoutingSchemaTags)
	err := c.cc.Invoke(ctx, "/engine.RoutingSchemaService/SearchRoutingSchemaTags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoutingSchemaServiceServer is the server API for RoutingSchemaService service.
// All implementations must embed UnimplementedRoutingSchemaServiceServer
// for forward compatibility
type RoutingSchemaServiceServer interface {
	// Create RoutingSchema
	CreateRoutingSchema(context.Context, *CreateRoutingSchemaRequest) (*RoutingSchema, error)
	// List RoutingSchema
	SearchRoutingSchema(context.Context, *SearchRoutingSchemaRequest) (*ListRoutingSchema, error)
	// RoutingSchema item
	ReadRoutingSchema(context.Context, *ReadRoutingSchemaRequest) (*RoutingSchema, error)
	// Update RoutingSchema
	UpdateRoutingSchema(context.Context, *UpdateRoutingSchemaRequest) (*RoutingSchema, error)
	// Patch RoutingSchema
	PatchRoutingSchema(context.Context, *PatchRoutingSchemaRequest) (*RoutingSchema, error)
	// Remove RoutingSchema
	DeleteRoutingSchema(context.Context, *DeleteRoutingSchemaRequest) (*RoutingSchema, error)
	// List RoutingSchemaTags
	SearchRoutingSchemaTags(context.Context, *SearchRoutingSchemaTagsRequest) (*ListRoutingSchemaTags, error)
	mustEmbedUnimplementedRoutingSchemaServiceServer()
}

// UnimplementedRoutingSchemaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRoutingSchemaServiceServer struct {
}

func (UnimplementedRoutingSchemaServiceServer) CreateRoutingSchema(context.Context, *CreateRoutingSchemaRequest) (*RoutingSchema, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoutingSchema not implemented")
}
func (UnimplementedRoutingSchemaServiceServer) SearchRoutingSchema(context.Context, *SearchRoutingSchemaRequest) (*ListRoutingSchema, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchRoutingSchema not implemented")
}
func (UnimplementedRoutingSchemaServiceServer) ReadRoutingSchema(context.Context, *ReadRoutingSchemaRequest) (*RoutingSchema, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadRoutingSchema not implemented")
}
func (UnimplementedRoutingSchemaServiceServer) UpdateRoutingSchema(context.Context, *UpdateRoutingSchemaRequest) (*RoutingSchema, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRoutingSchema not implemented")
}
func (UnimplementedRoutingSchemaServiceServer) PatchRoutingSchema(context.Context, *PatchRoutingSchemaRequest) (*RoutingSchema, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchRoutingSchema not implemented")
}
func (UnimplementedRoutingSchemaServiceServer) DeleteRoutingSchema(context.Context, *DeleteRoutingSchemaRequest) (*RoutingSchema, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRoutingSchema not implemented")
}
func (UnimplementedRoutingSchemaServiceServer) SearchRoutingSchemaTags(context.Context, *SearchRoutingSchemaTagsRequest) (*ListRoutingSchemaTags, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchRoutingSchemaTags not implemented")
}
func (UnimplementedRoutingSchemaServiceServer) mustEmbedUnimplementedRoutingSchemaServiceServer() {}

// UnsafeRoutingSchemaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoutingSchemaServiceServer will
// result in compilation errors.
type UnsafeRoutingSchemaServiceServer interface {
	mustEmbedUnimplementedRoutingSchemaServiceServer()
}

func RegisterRoutingSchemaServiceServer(s grpc.ServiceRegistrar, srv RoutingSchemaServiceServer) {
	s.RegisterService(&RoutingSchemaService_ServiceDesc, srv)
}

func _RoutingSchemaService_CreateRoutingSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoutingSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutingSchemaServiceServer).CreateRoutingSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.RoutingSchemaService/CreateRoutingSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutingSchemaServiceServer).CreateRoutingSchema(ctx, req.(*CreateRoutingSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoutingSchemaService_SearchRoutingSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRoutingSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutingSchemaServiceServer).SearchRoutingSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.RoutingSchemaService/SearchRoutingSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutingSchemaServiceServer).SearchRoutingSchema(ctx, req.(*SearchRoutingSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoutingSchemaService_ReadRoutingSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRoutingSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutingSchemaServiceServer).ReadRoutingSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.RoutingSchemaService/ReadRoutingSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutingSchemaServiceServer).ReadRoutingSchema(ctx, req.(*ReadRoutingSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoutingSchemaService_UpdateRoutingSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoutingSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutingSchemaServiceServer).UpdateRoutingSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.RoutingSchemaService/UpdateRoutingSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutingSchemaServiceServer).UpdateRoutingSchema(ctx, req.(*UpdateRoutingSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoutingSchemaService_PatchRoutingSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchRoutingSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutingSchemaServiceServer).PatchRoutingSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.RoutingSchemaService/PatchRoutingSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutingSchemaServiceServer).PatchRoutingSchema(ctx, req.(*PatchRoutingSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoutingSchemaService_DeleteRoutingSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoutingSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutingSchemaServiceServer).DeleteRoutingSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.RoutingSchemaService/DeleteRoutingSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutingSchemaServiceServer).DeleteRoutingSchema(ctx, req.(*DeleteRoutingSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoutingSchemaService_SearchRoutingSchemaTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRoutingSchemaTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutingSchemaServiceServer).SearchRoutingSchemaTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.RoutingSchemaService/SearchRoutingSchemaTags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutingSchemaServiceServer).SearchRoutingSchemaTags(ctx, req.(*SearchRoutingSchemaTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RoutingSchemaService_ServiceDesc is the grpc.ServiceDesc for RoutingSchemaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RoutingSchemaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "engine.RoutingSchemaService",
	HandlerType: (*RoutingSchemaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRoutingSchema",
			Handler:    _RoutingSchemaService_CreateRoutingSchema_Handler,
		},
		{
			MethodName: "SearchRoutingSchema",
			Handler:    _RoutingSchemaService_SearchRoutingSchema_Handler,
		},
		{
			MethodName: "ReadRoutingSchema",
			Handler:    _RoutingSchemaService_ReadRoutingSchema_Handler,
		},
		{
			MethodName: "UpdateRoutingSchema",
			Handler:    _RoutingSchemaService_UpdateRoutingSchema_Handler,
		},
		{
			MethodName: "PatchRoutingSchema",
			Handler:    _RoutingSchemaService_PatchRoutingSchema_Handler,
		},
		{
			MethodName: "DeleteRoutingSchema",
			Handler:    _RoutingSchemaService_DeleteRoutingSchema_Handler,
		},
		{
			MethodName: "SearchRoutingSchemaTags",
			Handler:    _RoutingSchemaService_SearchRoutingSchemaTags_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "acr_routing_schema.proto",
}
