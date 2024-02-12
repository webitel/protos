// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: acr_routing_out_call.proto

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

// RoutingOutboundCallServiceClient is the client API for RoutingOutboundCallService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoutingOutboundCallServiceClient interface {
	// Create RoutingOutboundCall
	CreateRoutingOutboundCall(ctx context.Context, in *CreateRoutingOutboundCallRequest, opts ...grpc.CallOption) (*RoutingOutboundCall, error)
	// List of RoutingOutboundCall
	SearchRoutingOutboundCall(ctx context.Context, in *SearchRoutingOutboundCallRequest, opts ...grpc.CallOption) (*ListRoutingOutboundCall, error)
	// RoutingOutboundCall item
	ReadRoutingOutboundCall(ctx context.Context, in *ReadRoutingOutboundCallRequest, opts ...grpc.CallOption) (*RoutingOutboundCall, error)
	// Update RoutingOutboundCall
	UpdateRoutingOutboundCall(ctx context.Context, in *UpdateRoutingOutboundCallRequest, opts ...grpc.CallOption) (*RoutingOutboundCall, error)
	// Patch RoutingOutboundCall
	PatchRoutingOutboundCall(ctx context.Context, in *PatchRoutingOutboundCallRequest, opts ...grpc.CallOption) (*RoutingOutboundCall, error)
	// Move RoutingOutboundCall
	MovePositionRoutingOutboundCall(ctx context.Context, in *MovePositionRoutingOutboundCallRequest, opts ...grpc.CallOption) (*MovePositionRoutingOutboundCallResponse, error)
	// Remove RoutingOutboundCall
	DeleteRoutingOutboundCall(ctx context.Context, in *DeleteRoutingOutboundCallRequest, opts ...grpc.CallOption) (*RoutingOutboundCall, error)
}

type routingOutboundCallServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRoutingOutboundCallServiceClient(cc grpc.ClientConnInterface) RoutingOutboundCallServiceClient {
	return &routingOutboundCallServiceClient{cc}
}

func (c *routingOutboundCallServiceClient) CreateRoutingOutboundCall(ctx context.Context, in *CreateRoutingOutboundCallRequest, opts ...grpc.CallOption) (*RoutingOutboundCall, error) {
	out := new(RoutingOutboundCall)
	err := c.cc.Invoke(ctx, "/engine.RoutingOutboundCallService/CreateRoutingOutboundCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routingOutboundCallServiceClient) SearchRoutingOutboundCall(ctx context.Context, in *SearchRoutingOutboundCallRequest, opts ...grpc.CallOption) (*ListRoutingOutboundCall, error) {
	out := new(ListRoutingOutboundCall)
	err := c.cc.Invoke(ctx, "/engine.RoutingOutboundCallService/SearchRoutingOutboundCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routingOutboundCallServiceClient) ReadRoutingOutboundCall(ctx context.Context, in *ReadRoutingOutboundCallRequest, opts ...grpc.CallOption) (*RoutingOutboundCall, error) {
	out := new(RoutingOutboundCall)
	err := c.cc.Invoke(ctx, "/engine.RoutingOutboundCallService/ReadRoutingOutboundCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routingOutboundCallServiceClient) UpdateRoutingOutboundCall(ctx context.Context, in *UpdateRoutingOutboundCallRequest, opts ...grpc.CallOption) (*RoutingOutboundCall, error) {
	out := new(RoutingOutboundCall)
	err := c.cc.Invoke(ctx, "/engine.RoutingOutboundCallService/UpdateRoutingOutboundCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routingOutboundCallServiceClient) PatchRoutingOutboundCall(ctx context.Context, in *PatchRoutingOutboundCallRequest, opts ...grpc.CallOption) (*RoutingOutboundCall, error) {
	out := new(RoutingOutboundCall)
	err := c.cc.Invoke(ctx, "/engine.RoutingOutboundCallService/PatchRoutingOutboundCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routingOutboundCallServiceClient) MovePositionRoutingOutboundCall(ctx context.Context, in *MovePositionRoutingOutboundCallRequest, opts ...grpc.CallOption) (*MovePositionRoutingOutboundCallResponse, error) {
	out := new(MovePositionRoutingOutboundCallResponse)
	err := c.cc.Invoke(ctx, "/engine.RoutingOutboundCallService/MovePositionRoutingOutboundCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routingOutboundCallServiceClient) DeleteRoutingOutboundCall(ctx context.Context, in *DeleteRoutingOutboundCallRequest, opts ...grpc.CallOption) (*RoutingOutboundCall, error) {
	out := new(RoutingOutboundCall)
	err := c.cc.Invoke(ctx, "/engine.RoutingOutboundCallService/DeleteRoutingOutboundCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoutingOutboundCallServiceServer is the server API for RoutingOutboundCallService service.
// All implementations must embed UnimplementedRoutingOutboundCallServiceServer
// for forward compatibility
type RoutingOutboundCallServiceServer interface {
	// Create RoutingOutboundCall
	CreateRoutingOutboundCall(context.Context, *CreateRoutingOutboundCallRequest) (*RoutingOutboundCall, error)
	// List of RoutingOutboundCall
	SearchRoutingOutboundCall(context.Context, *SearchRoutingOutboundCallRequest) (*ListRoutingOutboundCall, error)
	// RoutingOutboundCall item
	ReadRoutingOutboundCall(context.Context, *ReadRoutingOutboundCallRequest) (*RoutingOutboundCall, error)
	// Update RoutingOutboundCall
	UpdateRoutingOutboundCall(context.Context, *UpdateRoutingOutboundCallRequest) (*RoutingOutboundCall, error)
	// Patch RoutingOutboundCall
	PatchRoutingOutboundCall(context.Context, *PatchRoutingOutboundCallRequest) (*RoutingOutboundCall, error)
	// Move RoutingOutboundCall
	MovePositionRoutingOutboundCall(context.Context, *MovePositionRoutingOutboundCallRequest) (*MovePositionRoutingOutboundCallResponse, error)
	// Remove RoutingOutboundCall
	DeleteRoutingOutboundCall(context.Context, *DeleteRoutingOutboundCallRequest) (*RoutingOutboundCall, error)
	mustEmbedUnimplementedRoutingOutboundCallServiceServer()
}

// UnimplementedRoutingOutboundCallServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRoutingOutboundCallServiceServer struct {
}

func (UnimplementedRoutingOutboundCallServiceServer) CreateRoutingOutboundCall(context.Context, *CreateRoutingOutboundCallRequest) (*RoutingOutboundCall, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoutingOutboundCall not implemented")
}
func (UnimplementedRoutingOutboundCallServiceServer) SearchRoutingOutboundCall(context.Context, *SearchRoutingOutboundCallRequest) (*ListRoutingOutboundCall, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchRoutingOutboundCall not implemented")
}
func (UnimplementedRoutingOutboundCallServiceServer) ReadRoutingOutboundCall(context.Context, *ReadRoutingOutboundCallRequest) (*RoutingOutboundCall, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadRoutingOutboundCall not implemented")
}
func (UnimplementedRoutingOutboundCallServiceServer) UpdateRoutingOutboundCall(context.Context, *UpdateRoutingOutboundCallRequest) (*RoutingOutboundCall, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRoutingOutboundCall not implemented")
}
func (UnimplementedRoutingOutboundCallServiceServer) PatchRoutingOutboundCall(context.Context, *PatchRoutingOutboundCallRequest) (*RoutingOutboundCall, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchRoutingOutboundCall not implemented")
}
func (UnimplementedRoutingOutboundCallServiceServer) MovePositionRoutingOutboundCall(context.Context, *MovePositionRoutingOutboundCallRequest) (*MovePositionRoutingOutboundCallResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MovePositionRoutingOutboundCall not implemented")
}
func (UnimplementedRoutingOutboundCallServiceServer) DeleteRoutingOutboundCall(context.Context, *DeleteRoutingOutboundCallRequest) (*RoutingOutboundCall, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRoutingOutboundCall not implemented")
}
func (UnimplementedRoutingOutboundCallServiceServer) mustEmbedUnimplementedRoutingOutboundCallServiceServer() {
}

// UnsafeRoutingOutboundCallServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoutingOutboundCallServiceServer will
// result in compilation errors.
type UnsafeRoutingOutboundCallServiceServer interface {
	mustEmbedUnimplementedRoutingOutboundCallServiceServer()
}

func RegisterRoutingOutboundCallServiceServer(s grpc.ServiceRegistrar, srv RoutingOutboundCallServiceServer) {
	s.RegisterService(&RoutingOutboundCallService_ServiceDesc, srv)
}

func _RoutingOutboundCallService_CreateRoutingOutboundCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoutingOutboundCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutingOutboundCallServiceServer).CreateRoutingOutboundCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.RoutingOutboundCallService/CreateRoutingOutboundCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutingOutboundCallServiceServer).CreateRoutingOutboundCall(ctx, req.(*CreateRoutingOutboundCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoutingOutboundCallService_SearchRoutingOutboundCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRoutingOutboundCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutingOutboundCallServiceServer).SearchRoutingOutboundCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.RoutingOutboundCallService/SearchRoutingOutboundCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutingOutboundCallServiceServer).SearchRoutingOutboundCall(ctx, req.(*SearchRoutingOutboundCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoutingOutboundCallService_ReadRoutingOutboundCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRoutingOutboundCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutingOutboundCallServiceServer).ReadRoutingOutboundCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.RoutingOutboundCallService/ReadRoutingOutboundCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutingOutboundCallServiceServer).ReadRoutingOutboundCall(ctx, req.(*ReadRoutingOutboundCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoutingOutboundCallService_UpdateRoutingOutboundCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoutingOutboundCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutingOutboundCallServiceServer).UpdateRoutingOutboundCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.RoutingOutboundCallService/UpdateRoutingOutboundCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutingOutboundCallServiceServer).UpdateRoutingOutboundCall(ctx, req.(*UpdateRoutingOutboundCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoutingOutboundCallService_PatchRoutingOutboundCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchRoutingOutboundCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutingOutboundCallServiceServer).PatchRoutingOutboundCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.RoutingOutboundCallService/PatchRoutingOutboundCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutingOutboundCallServiceServer).PatchRoutingOutboundCall(ctx, req.(*PatchRoutingOutboundCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoutingOutboundCallService_MovePositionRoutingOutboundCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MovePositionRoutingOutboundCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutingOutboundCallServiceServer).MovePositionRoutingOutboundCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.RoutingOutboundCallService/MovePositionRoutingOutboundCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutingOutboundCallServiceServer).MovePositionRoutingOutboundCall(ctx, req.(*MovePositionRoutingOutboundCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoutingOutboundCallService_DeleteRoutingOutboundCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoutingOutboundCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutingOutboundCallServiceServer).DeleteRoutingOutboundCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.RoutingOutboundCallService/DeleteRoutingOutboundCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutingOutboundCallServiceServer).DeleteRoutingOutboundCall(ctx, req.(*DeleteRoutingOutboundCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RoutingOutboundCallService_ServiceDesc is the grpc.ServiceDesc for RoutingOutboundCallService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RoutingOutboundCallService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "engine.RoutingOutboundCallService",
	HandlerType: (*RoutingOutboundCallServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRoutingOutboundCall",
			Handler:    _RoutingOutboundCallService_CreateRoutingOutboundCall_Handler,
		},
		{
			MethodName: "SearchRoutingOutboundCall",
			Handler:    _RoutingOutboundCallService_SearchRoutingOutboundCall_Handler,
		},
		{
			MethodName: "ReadRoutingOutboundCall",
			Handler:    _RoutingOutboundCallService_ReadRoutingOutboundCall_Handler,
		},
		{
			MethodName: "UpdateRoutingOutboundCall",
			Handler:    _RoutingOutboundCallService_UpdateRoutingOutboundCall_Handler,
		},
		{
			MethodName: "PatchRoutingOutboundCall",
			Handler:    _RoutingOutboundCallService_PatchRoutingOutboundCall_Handler,
		},
		{
			MethodName: "MovePositionRoutingOutboundCall",
			Handler:    _RoutingOutboundCallService_MovePositionRoutingOutboundCall_Handler,
		},
		{
			MethodName: "DeleteRoutingOutboundCall",
			Handler:    _RoutingOutboundCallService_DeleteRoutingOutboundCall_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "acr_routing_out_call.proto",
}
