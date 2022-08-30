// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.0
// source: trigger.proto

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

// TriggerServiceClient is the client API for TriggerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TriggerServiceClient interface {
	// Create Trigger
	CreateTrigger(ctx context.Context, in *CreateTriggerRequest, opts ...grpc.CallOption) (*Trigger, error)
	// List of Trigger
	SearchTrigger(ctx context.Context, in *SearchTriggerRequest, opts ...grpc.CallOption) (*ListTrigger, error)
	// Trigger item
	ReadTrigger(ctx context.Context, in *ReadTriggerRequest, opts ...grpc.CallOption) (*Trigger, error)
	// Update Trigger
	UpdateTrigger(ctx context.Context, in *UpdateTriggerRequest, opts ...grpc.CallOption) (*Trigger, error)
	PatchTrigger(ctx context.Context, in *PatchTriggerRequest, opts ...grpc.CallOption) (*Trigger, error)
	// Remove Trigger
	DeleteTrigger(ctx context.Context, in *DeleteTriggerRequest, opts ...grpc.CallOption) (*Trigger, error)
}

type triggerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTriggerServiceClient(cc grpc.ClientConnInterface) TriggerServiceClient {
	return &triggerServiceClient{cc}
}

func (c *triggerServiceClient) CreateTrigger(ctx context.Context, in *CreateTriggerRequest, opts ...grpc.CallOption) (*Trigger, error) {
	out := new(Trigger)
	err := c.cc.Invoke(ctx, "/engine.TriggerService/CreateTrigger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *triggerServiceClient) SearchTrigger(ctx context.Context, in *SearchTriggerRequest, opts ...grpc.CallOption) (*ListTrigger, error) {
	out := new(ListTrigger)
	err := c.cc.Invoke(ctx, "/engine.TriggerService/SearchTrigger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *triggerServiceClient) ReadTrigger(ctx context.Context, in *ReadTriggerRequest, opts ...grpc.CallOption) (*Trigger, error) {
	out := new(Trigger)
	err := c.cc.Invoke(ctx, "/engine.TriggerService/ReadTrigger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *triggerServiceClient) UpdateTrigger(ctx context.Context, in *UpdateTriggerRequest, opts ...grpc.CallOption) (*Trigger, error) {
	out := new(Trigger)
	err := c.cc.Invoke(ctx, "/engine.TriggerService/UpdateTrigger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *triggerServiceClient) PatchTrigger(ctx context.Context, in *PatchTriggerRequest, opts ...grpc.CallOption) (*Trigger, error) {
	out := new(Trigger)
	err := c.cc.Invoke(ctx, "/engine.TriggerService/PatchTrigger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *triggerServiceClient) DeleteTrigger(ctx context.Context, in *DeleteTriggerRequest, opts ...grpc.CallOption) (*Trigger, error) {
	out := new(Trigger)
	err := c.cc.Invoke(ctx, "/engine.TriggerService/DeleteTrigger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TriggerServiceServer is the server API for TriggerService service.
// All implementations must embed UnimplementedTriggerServiceServer
// for forward compatibility
type TriggerServiceServer interface {
	// Create Trigger
	CreateTrigger(context.Context, *CreateTriggerRequest) (*Trigger, error)
	// List of Trigger
	SearchTrigger(context.Context, *SearchTriggerRequest) (*ListTrigger, error)
	// Trigger item
	ReadTrigger(context.Context, *ReadTriggerRequest) (*Trigger, error)
	// Update Trigger
	UpdateTrigger(context.Context, *UpdateTriggerRequest) (*Trigger, error)
	PatchTrigger(context.Context, *PatchTriggerRequest) (*Trigger, error)
	// Remove Trigger
	DeleteTrigger(context.Context, *DeleteTriggerRequest) (*Trigger, error)
	mustEmbedUnimplementedTriggerServiceServer()
}

// UnimplementedTriggerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTriggerServiceServer struct {
}

func (UnimplementedTriggerServiceServer) CreateTrigger(context.Context, *CreateTriggerRequest) (*Trigger, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTrigger not implemented")
}
func (UnimplementedTriggerServiceServer) SearchTrigger(context.Context, *SearchTriggerRequest) (*ListTrigger, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchTrigger not implemented")
}
func (UnimplementedTriggerServiceServer) ReadTrigger(context.Context, *ReadTriggerRequest) (*Trigger, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadTrigger not implemented")
}
func (UnimplementedTriggerServiceServer) UpdateTrigger(context.Context, *UpdateTriggerRequest) (*Trigger, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTrigger not implemented")
}
func (UnimplementedTriggerServiceServer) PatchTrigger(context.Context, *PatchTriggerRequest) (*Trigger, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchTrigger not implemented")
}
func (UnimplementedTriggerServiceServer) DeleteTrigger(context.Context, *DeleteTriggerRequest) (*Trigger, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTrigger not implemented")
}
func (UnimplementedTriggerServiceServer) mustEmbedUnimplementedTriggerServiceServer() {}

// UnsafeTriggerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TriggerServiceServer will
// result in compilation errors.
type UnsafeTriggerServiceServer interface {
	mustEmbedUnimplementedTriggerServiceServer()
}

func RegisterTriggerServiceServer(s grpc.ServiceRegistrar, srv TriggerServiceServer) {
	s.RegisterService(&TriggerService_ServiceDesc, srv)
}

func _TriggerService_CreateTrigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTriggerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TriggerServiceServer).CreateTrigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.TriggerService/CreateTrigger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TriggerServiceServer).CreateTrigger(ctx, req.(*CreateTriggerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TriggerService_SearchTrigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchTriggerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TriggerServiceServer).SearchTrigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.TriggerService/SearchTrigger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TriggerServiceServer).SearchTrigger(ctx, req.(*SearchTriggerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TriggerService_ReadTrigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadTriggerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TriggerServiceServer).ReadTrigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.TriggerService/ReadTrigger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TriggerServiceServer).ReadTrigger(ctx, req.(*ReadTriggerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TriggerService_UpdateTrigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTriggerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TriggerServiceServer).UpdateTrigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.TriggerService/UpdateTrigger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TriggerServiceServer).UpdateTrigger(ctx, req.(*UpdateTriggerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TriggerService_PatchTrigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchTriggerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TriggerServiceServer).PatchTrigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.TriggerService/PatchTrigger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TriggerServiceServer).PatchTrigger(ctx, req.(*PatchTriggerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TriggerService_DeleteTrigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTriggerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TriggerServiceServer).DeleteTrigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.TriggerService/DeleteTrigger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TriggerServiceServer).DeleteTrigger(ctx, req.(*DeleteTriggerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TriggerService_ServiceDesc is the grpc.ServiceDesc for TriggerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TriggerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "engine.TriggerService",
	HandlerType: (*TriggerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTrigger",
			Handler:    _TriggerService_CreateTrigger_Handler,
		},
		{
			MethodName: "SearchTrigger",
			Handler:    _TriggerService_SearchTrigger_Handler,
		},
		{
			MethodName: "ReadTrigger",
			Handler:    _TriggerService_ReadTrigger_Handler,
		},
		{
			MethodName: "UpdateTrigger",
			Handler:    _TriggerService_UpdateTrigger_Handler,
		},
		{
			MethodName: "PatchTrigger",
			Handler:    _TriggerService_PatchTrigger_Handler,
		},
		{
			MethodName: "DeleteTrigger",
			Handler:    _TriggerService_DeleteTrigger_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "trigger.proto",
}
