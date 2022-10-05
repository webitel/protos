// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.0
// source: pause_cause.proto

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

// AgentPauseCauseServiceClient is the client API for AgentPauseCauseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentPauseCauseServiceClient interface {
	CreateAgentPauseCause(ctx context.Context, in *CreateAgentPauseCauseRequest, opts ...grpc.CallOption) (*AgentPauseCause, error)
	SearchAgentPauseCause(ctx context.Context, in *SearchAgentPauseCauseRequest, opts ...grpc.CallOption) (*ListAgentPauseCause, error)
	ReadAgentPauseCause(ctx context.Context, in *ReadAgentPauseCauseRequest, opts ...grpc.CallOption) (*AgentPauseCause, error)
	PatchAgentPauseCause(ctx context.Context, in *PatchAgentPauseCauseRequest, opts ...grpc.CallOption) (*AgentPauseCause, error)
	UpdateAgentPauseCause(ctx context.Context, in *UpdateAgentPauseCauseRequest, opts ...grpc.CallOption) (*AgentPauseCause, error)
	DeleteAgentPauseCause(ctx context.Context, in *DeleteAgentPauseCauseRequest, opts ...grpc.CallOption) (*AgentPauseCause, error)
}

type agentPauseCauseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentPauseCauseServiceClient(cc grpc.ClientConnInterface) AgentPauseCauseServiceClient {
	return &agentPauseCauseServiceClient{cc}
}

func (c *agentPauseCauseServiceClient) CreateAgentPauseCause(ctx context.Context, in *CreateAgentPauseCauseRequest, opts ...grpc.CallOption) (*AgentPauseCause, error) {
	out := new(AgentPauseCause)
	err := c.cc.Invoke(ctx, "/engine.AgentPauseCauseService/CreateAgentPauseCause", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentPauseCauseServiceClient) SearchAgentPauseCause(ctx context.Context, in *SearchAgentPauseCauseRequest, opts ...grpc.CallOption) (*ListAgentPauseCause, error) {
	out := new(ListAgentPauseCause)
	err := c.cc.Invoke(ctx, "/engine.AgentPauseCauseService/SearchAgentPauseCause", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentPauseCauseServiceClient) ReadAgentPauseCause(ctx context.Context, in *ReadAgentPauseCauseRequest, opts ...grpc.CallOption) (*AgentPauseCause, error) {
	out := new(AgentPauseCause)
	err := c.cc.Invoke(ctx, "/engine.AgentPauseCauseService/ReadAgentPauseCause", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentPauseCauseServiceClient) PatchAgentPauseCause(ctx context.Context, in *PatchAgentPauseCauseRequest, opts ...grpc.CallOption) (*AgentPauseCause, error) {
	out := new(AgentPauseCause)
	err := c.cc.Invoke(ctx, "/engine.AgentPauseCauseService/PatchAgentPauseCause", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentPauseCauseServiceClient) UpdateAgentPauseCause(ctx context.Context, in *UpdateAgentPauseCauseRequest, opts ...grpc.CallOption) (*AgentPauseCause, error) {
	out := new(AgentPauseCause)
	err := c.cc.Invoke(ctx, "/engine.AgentPauseCauseService/UpdateAgentPauseCause", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentPauseCauseServiceClient) DeleteAgentPauseCause(ctx context.Context, in *DeleteAgentPauseCauseRequest, opts ...grpc.CallOption) (*AgentPauseCause, error) {
	out := new(AgentPauseCause)
	err := c.cc.Invoke(ctx, "/engine.AgentPauseCauseService/DeleteAgentPauseCause", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentPauseCauseServiceServer is the server API for AgentPauseCauseService service.
// All implementations must embed UnimplementedAgentPauseCauseServiceServer
// for forward compatibility
type AgentPauseCauseServiceServer interface {
	CreateAgentPauseCause(context.Context, *CreateAgentPauseCauseRequest) (*AgentPauseCause, error)
	SearchAgentPauseCause(context.Context, *SearchAgentPauseCauseRequest) (*ListAgentPauseCause, error)
	ReadAgentPauseCause(context.Context, *ReadAgentPauseCauseRequest) (*AgentPauseCause, error)
	PatchAgentPauseCause(context.Context, *PatchAgentPauseCauseRequest) (*AgentPauseCause, error)
	UpdateAgentPauseCause(context.Context, *UpdateAgentPauseCauseRequest) (*AgentPauseCause, error)
	DeleteAgentPauseCause(context.Context, *DeleteAgentPauseCauseRequest) (*AgentPauseCause, error)
	mustEmbedUnimplementedAgentPauseCauseServiceServer()
}

// UnimplementedAgentPauseCauseServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAgentPauseCauseServiceServer struct {
}

func (UnimplementedAgentPauseCauseServiceServer) CreateAgentPauseCause(context.Context, *CreateAgentPauseCauseRequest) (*AgentPauseCause, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAgentPauseCause not implemented")
}
func (UnimplementedAgentPauseCauseServiceServer) SearchAgentPauseCause(context.Context, *SearchAgentPauseCauseRequest) (*ListAgentPauseCause, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchAgentPauseCause not implemented")
}
func (UnimplementedAgentPauseCauseServiceServer) ReadAgentPauseCause(context.Context, *ReadAgentPauseCauseRequest) (*AgentPauseCause, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadAgentPauseCause not implemented")
}
func (UnimplementedAgentPauseCauseServiceServer) PatchAgentPauseCause(context.Context, *PatchAgentPauseCauseRequest) (*AgentPauseCause, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchAgentPauseCause not implemented")
}
func (UnimplementedAgentPauseCauseServiceServer) UpdateAgentPauseCause(context.Context, *UpdateAgentPauseCauseRequest) (*AgentPauseCause, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAgentPauseCause not implemented")
}
func (UnimplementedAgentPauseCauseServiceServer) DeleteAgentPauseCause(context.Context, *DeleteAgentPauseCauseRequest) (*AgentPauseCause, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAgentPauseCause not implemented")
}
func (UnimplementedAgentPauseCauseServiceServer) mustEmbedUnimplementedAgentPauseCauseServiceServer() {
}

// UnsafeAgentPauseCauseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentPauseCauseServiceServer will
// result in compilation errors.
type UnsafeAgentPauseCauseServiceServer interface {
	mustEmbedUnimplementedAgentPauseCauseServiceServer()
}

func RegisterAgentPauseCauseServiceServer(s grpc.ServiceRegistrar, srv AgentPauseCauseServiceServer) {
	s.RegisterService(&AgentPauseCauseService_ServiceDesc, srv)
}

func _AgentPauseCauseService_CreateAgentPauseCause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAgentPauseCauseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentPauseCauseServiceServer).CreateAgentPauseCause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.AgentPauseCauseService/CreateAgentPauseCause",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentPauseCauseServiceServer).CreateAgentPauseCause(ctx, req.(*CreateAgentPauseCauseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentPauseCauseService_SearchAgentPauseCause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchAgentPauseCauseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentPauseCauseServiceServer).SearchAgentPauseCause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.AgentPauseCauseService/SearchAgentPauseCause",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentPauseCauseServiceServer).SearchAgentPauseCause(ctx, req.(*SearchAgentPauseCauseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentPauseCauseService_ReadAgentPauseCause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadAgentPauseCauseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentPauseCauseServiceServer).ReadAgentPauseCause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.AgentPauseCauseService/ReadAgentPauseCause",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentPauseCauseServiceServer).ReadAgentPauseCause(ctx, req.(*ReadAgentPauseCauseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentPauseCauseService_PatchAgentPauseCause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchAgentPauseCauseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentPauseCauseServiceServer).PatchAgentPauseCause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.AgentPauseCauseService/PatchAgentPauseCause",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentPauseCauseServiceServer).PatchAgentPauseCause(ctx, req.(*PatchAgentPauseCauseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentPauseCauseService_UpdateAgentPauseCause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAgentPauseCauseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentPauseCauseServiceServer).UpdateAgentPauseCause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.AgentPauseCauseService/UpdateAgentPauseCause",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentPauseCauseServiceServer).UpdateAgentPauseCause(ctx, req.(*UpdateAgentPauseCauseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentPauseCauseService_DeleteAgentPauseCause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAgentPauseCauseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentPauseCauseServiceServer).DeleteAgentPauseCause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.AgentPauseCauseService/DeleteAgentPauseCause",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentPauseCauseServiceServer).DeleteAgentPauseCause(ctx, req.(*DeleteAgentPauseCauseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AgentPauseCauseService_ServiceDesc is the grpc.ServiceDesc for AgentPauseCauseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AgentPauseCauseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "engine.AgentPauseCauseService",
	HandlerType: (*AgentPauseCauseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAgentPauseCause",
			Handler:    _AgentPauseCauseService_CreateAgentPauseCause_Handler,
		},
		{
			MethodName: "SearchAgentPauseCause",
			Handler:    _AgentPauseCauseService_SearchAgentPauseCause_Handler,
		},
		{
			MethodName: "ReadAgentPauseCause",
			Handler:    _AgentPauseCauseService_ReadAgentPauseCause_Handler,
		},
		{
			MethodName: "PatchAgentPauseCause",
			Handler:    _AgentPauseCauseService_PatchAgentPauseCause_Handler,
		},
		{
			MethodName: "UpdateAgentPauseCause",
			Handler:    _AgentPauseCauseService_UpdateAgentPauseCause_Handler,
		},
		{
			MethodName: "DeleteAgentPauseCause",
			Handler:    _AgentPauseCauseService_DeleteAgentPauseCause_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pause_cause.proto",
}