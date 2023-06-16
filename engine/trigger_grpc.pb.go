// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.6
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

const (
	TriggerService_CreateTrigger_FullMethodName    = "/engine.TriggerService/CreateTrigger"
	TriggerService_SearchTrigger_FullMethodName    = "/engine.TriggerService/SearchTrigger"
	TriggerService_ReadTrigger_FullMethodName      = "/engine.TriggerService/ReadTrigger"
	TriggerService_UpdateTrigger_FullMethodName    = "/engine.TriggerService/UpdateTrigger"
	TriggerService_PatchTrigger_FullMethodName     = "/engine.TriggerService/PatchTrigger"
	TriggerService_DeleteTrigger_FullMethodName    = "/engine.TriggerService/DeleteTrigger"
	TriggerService_CreateTriggerJob_FullMethodName = "/engine.TriggerService/CreateTriggerJob"
	TriggerService_SearchTriggerJob_FullMethodName = "/engine.TriggerService/SearchTriggerJob"
)

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
	CreateTriggerJob(ctx context.Context, in *CreateTriggerJobRequest, opts ...grpc.CallOption) (*TriggerJob, error)
	SearchTriggerJob(ctx context.Context, in *SearchTriggerJobRequest, opts ...grpc.CallOption) (*ListTriggerJob, error)
}

type triggerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTriggerServiceClient(cc grpc.ClientConnInterface) TriggerServiceClient {
	return &triggerServiceClient{cc}
}

func (c *triggerServiceClient) CreateTrigger(ctx context.Context, in *CreateTriggerRequest, opts ...grpc.CallOption) (*Trigger, error) {
	out := new(Trigger)
	err := c.cc.Invoke(ctx, TriggerService_CreateTrigger_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *triggerServiceClient) SearchTrigger(ctx context.Context, in *SearchTriggerRequest, opts ...grpc.CallOption) (*ListTrigger, error) {
	out := new(ListTrigger)
	err := c.cc.Invoke(ctx, TriggerService_SearchTrigger_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *triggerServiceClient) ReadTrigger(ctx context.Context, in *ReadTriggerRequest, opts ...grpc.CallOption) (*Trigger, error) {
	out := new(Trigger)
	err := c.cc.Invoke(ctx, TriggerService_ReadTrigger_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *triggerServiceClient) UpdateTrigger(ctx context.Context, in *UpdateTriggerRequest, opts ...grpc.CallOption) (*Trigger, error) {
	out := new(Trigger)
	err := c.cc.Invoke(ctx, TriggerService_UpdateTrigger_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *triggerServiceClient) PatchTrigger(ctx context.Context, in *PatchTriggerRequest, opts ...grpc.CallOption) (*Trigger, error) {
	out := new(Trigger)
	err := c.cc.Invoke(ctx, TriggerService_PatchTrigger_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *triggerServiceClient) DeleteTrigger(ctx context.Context, in *DeleteTriggerRequest, opts ...grpc.CallOption) (*Trigger, error) {
	out := new(Trigger)
	err := c.cc.Invoke(ctx, TriggerService_DeleteTrigger_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *triggerServiceClient) CreateTriggerJob(ctx context.Context, in *CreateTriggerJobRequest, opts ...grpc.CallOption) (*TriggerJob, error) {
	out := new(TriggerJob)
	err := c.cc.Invoke(ctx, TriggerService_CreateTriggerJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *triggerServiceClient) SearchTriggerJob(ctx context.Context, in *SearchTriggerJobRequest, opts ...grpc.CallOption) (*ListTriggerJob, error) {
	out := new(ListTriggerJob)
	err := c.cc.Invoke(ctx, TriggerService_SearchTriggerJob_FullMethodName, in, out, opts...)
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
	CreateTriggerJob(context.Context, *CreateTriggerJobRequest) (*TriggerJob, error)
	SearchTriggerJob(context.Context, *SearchTriggerJobRequest) (*ListTriggerJob, error)
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
func (UnimplementedTriggerServiceServer) CreateTriggerJob(context.Context, *CreateTriggerJobRequest) (*TriggerJob, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTriggerJob not implemented")
}
func (UnimplementedTriggerServiceServer) SearchTriggerJob(context.Context, *SearchTriggerJobRequest) (*ListTriggerJob, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchTriggerJob not implemented")
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
		FullMethod: TriggerService_CreateTrigger_FullMethodName,
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
		FullMethod: TriggerService_SearchTrigger_FullMethodName,
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
		FullMethod: TriggerService_ReadTrigger_FullMethodName,
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
		FullMethod: TriggerService_UpdateTrigger_FullMethodName,
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
		FullMethod: TriggerService_PatchTrigger_FullMethodName,
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
		FullMethod: TriggerService_DeleteTrigger_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TriggerServiceServer).DeleteTrigger(ctx, req.(*DeleteTriggerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TriggerService_CreateTriggerJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTriggerJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TriggerServiceServer).CreateTriggerJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TriggerService_CreateTriggerJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TriggerServiceServer).CreateTriggerJob(ctx, req.(*CreateTriggerJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TriggerService_SearchTriggerJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchTriggerJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TriggerServiceServer).SearchTriggerJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TriggerService_SearchTriggerJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TriggerServiceServer).SearchTriggerJob(ctx, req.(*SearchTriggerJobRequest))
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
		{
			MethodName: "CreateTriggerJob",
			Handler:    _TriggerService_CreateTriggerJob_Handler,
		},
		{
			MethodName: "SearchTriggerJob",
			Handler:    _TriggerService_SearchTriggerJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "trigger.proto",
}
