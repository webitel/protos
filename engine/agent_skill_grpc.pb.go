// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: agent_skill.proto

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

// AgentSkillServiceClient is the client API for AgentSkillService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentSkillServiceClient interface {
	// Create AgentSkill
	CreateAgentSkill(ctx context.Context, in *CreateAgentSkillRequest, opts ...grpc.CallOption) (*AgentSkill, error)
	CreateAgentSkills(ctx context.Context, in *CreateAgentSkillsRequest, opts ...grpc.CallOption) (*CreateAgentSkillsResponse, error)
	// List of AgentSkill
	SearchAgentSkill(ctx context.Context, in *SearchAgentSkillRequest, opts ...grpc.CallOption) (*ListAgentSkill, error)
	// AgentSkill item
	ReadAgentSkill(ctx context.Context, in *AgentSkillItemRequest, opts ...grpc.CallOption) (*AgentSkill, error)
	// Update AgentSkill
	UpdateAgentSkill(ctx context.Context, in *UpdateAgentSkillRequest, opts ...grpc.CallOption) (*AgentSkill, error)
	PatchAgentSkill(ctx context.Context, in *PatchAgentSkillRequest, opts ...grpc.CallOption) (*AgentSkill, error)
	PatchAgentSkills(ctx context.Context, in *PatchAgentSkillsRequest, opts ...grpc.CallOption) (*ListAgentSkill, error)
	// Remove AgentSkill
	DeleteAgentSkill(ctx context.Context, in *DeleteAgentSkillRequest, opts ...grpc.CallOption) (*AgentSkill, error)
	DeleteAgentSkills(ctx context.Context, in *DeleteAgentSkillsRequest, opts ...grpc.CallOption) (*ListAgentSkill, error)
	// SearchLookupAgentNotExistsSkill
	SearchLookupAgentNotExistsSkill(ctx context.Context, in *SearchLookupAgentNotExistsSkillRequest, opts ...grpc.CallOption) (*ListSkill, error)
}

type agentSkillServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentSkillServiceClient(cc grpc.ClientConnInterface) AgentSkillServiceClient {
	return &agentSkillServiceClient{cc}
}

func (c *agentSkillServiceClient) CreateAgentSkill(ctx context.Context, in *CreateAgentSkillRequest, opts ...grpc.CallOption) (*AgentSkill, error) {
	out := new(AgentSkill)
	err := c.cc.Invoke(ctx, "/engine.AgentSkillService/CreateAgentSkill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentSkillServiceClient) CreateAgentSkills(ctx context.Context, in *CreateAgentSkillsRequest, opts ...grpc.CallOption) (*CreateAgentSkillsResponse, error) {
	out := new(CreateAgentSkillsResponse)
	err := c.cc.Invoke(ctx, "/engine.AgentSkillService/CreateAgentSkills", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentSkillServiceClient) SearchAgentSkill(ctx context.Context, in *SearchAgentSkillRequest, opts ...grpc.CallOption) (*ListAgentSkill, error) {
	out := new(ListAgentSkill)
	err := c.cc.Invoke(ctx, "/engine.AgentSkillService/SearchAgentSkill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentSkillServiceClient) ReadAgentSkill(ctx context.Context, in *AgentSkillItemRequest, opts ...grpc.CallOption) (*AgentSkill, error) {
	out := new(AgentSkill)
	err := c.cc.Invoke(ctx, "/engine.AgentSkillService/ReadAgentSkill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentSkillServiceClient) UpdateAgentSkill(ctx context.Context, in *UpdateAgentSkillRequest, opts ...grpc.CallOption) (*AgentSkill, error) {
	out := new(AgentSkill)
	err := c.cc.Invoke(ctx, "/engine.AgentSkillService/UpdateAgentSkill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentSkillServiceClient) PatchAgentSkill(ctx context.Context, in *PatchAgentSkillRequest, opts ...grpc.CallOption) (*AgentSkill, error) {
	out := new(AgentSkill)
	err := c.cc.Invoke(ctx, "/engine.AgentSkillService/PatchAgentSkill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentSkillServiceClient) PatchAgentSkills(ctx context.Context, in *PatchAgentSkillsRequest, opts ...grpc.CallOption) (*ListAgentSkill, error) {
	out := new(ListAgentSkill)
	err := c.cc.Invoke(ctx, "/engine.AgentSkillService/PatchAgentSkills", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentSkillServiceClient) DeleteAgentSkill(ctx context.Context, in *DeleteAgentSkillRequest, opts ...grpc.CallOption) (*AgentSkill, error) {
	out := new(AgentSkill)
	err := c.cc.Invoke(ctx, "/engine.AgentSkillService/DeleteAgentSkill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentSkillServiceClient) DeleteAgentSkills(ctx context.Context, in *DeleteAgentSkillsRequest, opts ...grpc.CallOption) (*ListAgentSkill, error) {
	out := new(ListAgentSkill)
	err := c.cc.Invoke(ctx, "/engine.AgentSkillService/DeleteAgentSkills", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentSkillServiceClient) SearchLookupAgentNotExistsSkill(ctx context.Context, in *SearchLookupAgentNotExistsSkillRequest, opts ...grpc.CallOption) (*ListSkill, error) {
	out := new(ListSkill)
	err := c.cc.Invoke(ctx, "/engine.AgentSkillService/SearchLookupAgentNotExistsSkill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentSkillServiceServer is the server API for AgentSkillService service.
// All implementations must embed UnimplementedAgentSkillServiceServer
// for forward compatibility
type AgentSkillServiceServer interface {
	// Create AgentSkill
	CreateAgentSkill(context.Context, *CreateAgentSkillRequest) (*AgentSkill, error)
	CreateAgentSkills(context.Context, *CreateAgentSkillsRequest) (*CreateAgentSkillsResponse, error)
	// List of AgentSkill
	SearchAgentSkill(context.Context, *SearchAgentSkillRequest) (*ListAgentSkill, error)
	// AgentSkill item
	ReadAgentSkill(context.Context, *AgentSkillItemRequest) (*AgentSkill, error)
	// Update AgentSkill
	UpdateAgentSkill(context.Context, *UpdateAgentSkillRequest) (*AgentSkill, error)
	PatchAgentSkill(context.Context, *PatchAgentSkillRequest) (*AgentSkill, error)
	PatchAgentSkills(context.Context, *PatchAgentSkillsRequest) (*ListAgentSkill, error)
	// Remove AgentSkill
	DeleteAgentSkill(context.Context, *DeleteAgentSkillRequest) (*AgentSkill, error)
	DeleteAgentSkills(context.Context, *DeleteAgentSkillsRequest) (*ListAgentSkill, error)
	// SearchLookupAgentNotExistsSkill
	SearchLookupAgentNotExistsSkill(context.Context, *SearchLookupAgentNotExistsSkillRequest) (*ListSkill, error)
	mustEmbedUnimplementedAgentSkillServiceServer()
}

// UnimplementedAgentSkillServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAgentSkillServiceServer struct {
}

func (UnimplementedAgentSkillServiceServer) CreateAgentSkill(context.Context, *CreateAgentSkillRequest) (*AgentSkill, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAgentSkill not implemented")
}
func (UnimplementedAgentSkillServiceServer) CreateAgentSkills(context.Context, *CreateAgentSkillsRequest) (*CreateAgentSkillsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAgentSkills not implemented")
}
func (UnimplementedAgentSkillServiceServer) SearchAgentSkill(context.Context, *SearchAgentSkillRequest) (*ListAgentSkill, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchAgentSkill not implemented")
}
func (UnimplementedAgentSkillServiceServer) ReadAgentSkill(context.Context, *AgentSkillItemRequest) (*AgentSkill, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadAgentSkill not implemented")
}
func (UnimplementedAgentSkillServiceServer) UpdateAgentSkill(context.Context, *UpdateAgentSkillRequest) (*AgentSkill, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAgentSkill not implemented")
}
func (UnimplementedAgentSkillServiceServer) PatchAgentSkill(context.Context, *PatchAgentSkillRequest) (*AgentSkill, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchAgentSkill not implemented")
}
func (UnimplementedAgentSkillServiceServer) PatchAgentSkills(context.Context, *PatchAgentSkillsRequest) (*ListAgentSkill, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchAgentSkills not implemented")
}
func (UnimplementedAgentSkillServiceServer) DeleteAgentSkill(context.Context, *DeleteAgentSkillRequest) (*AgentSkill, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAgentSkill not implemented")
}
func (UnimplementedAgentSkillServiceServer) DeleteAgentSkills(context.Context, *DeleteAgentSkillsRequest) (*ListAgentSkill, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAgentSkills not implemented")
}
func (UnimplementedAgentSkillServiceServer) SearchLookupAgentNotExistsSkill(context.Context, *SearchLookupAgentNotExistsSkillRequest) (*ListSkill, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchLookupAgentNotExistsSkill not implemented")
}
func (UnimplementedAgentSkillServiceServer) mustEmbedUnimplementedAgentSkillServiceServer() {}

// UnsafeAgentSkillServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentSkillServiceServer will
// result in compilation errors.
type UnsafeAgentSkillServiceServer interface {
	mustEmbedUnimplementedAgentSkillServiceServer()
}

func RegisterAgentSkillServiceServer(s grpc.ServiceRegistrar, srv AgentSkillServiceServer) {
	s.RegisterService(&AgentSkillService_ServiceDesc, srv)
}

func _AgentSkillService_CreateAgentSkill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAgentSkillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentSkillServiceServer).CreateAgentSkill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.AgentSkillService/CreateAgentSkill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentSkillServiceServer).CreateAgentSkill(ctx, req.(*CreateAgentSkillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentSkillService_CreateAgentSkills_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAgentSkillsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentSkillServiceServer).CreateAgentSkills(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.AgentSkillService/CreateAgentSkills",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentSkillServiceServer).CreateAgentSkills(ctx, req.(*CreateAgentSkillsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentSkillService_SearchAgentSkill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchAgentSkillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentSkillServiceServer).SearchAgentSkill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.AgentSkillService/SearchAgentSkill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentSkillServiceServer).SearchAgentSkill(ctx, req.(*SearchAgentSkillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentSkillService_ReadAgentSkill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentSkillItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentSkillServiceServer).ReadAgentSkill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.AgentSkillService/ReadAgentSkill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentSkillServiceServer).ReadAgentSkill(ctx, req.(*AgentSkillItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentSkillService_UpdateAgentSkill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAgentSkillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentSkillServiceServer).UpdateAgentSkill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.AgentSkillService/UpdateAgentSkill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentSkillServiceServer).UpdateAgentSkill(ctx, req.(*UpdateAgentSkillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentSkillService_PatchAgentSkill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchAgentSkillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentSkillServiceServer).PatchAgentSkill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.AgentSkillService/PatchAgentSkill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentSkillServiceServer).PatchAgentSkill(ctx, req.(*PatchAgentSkillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentSkillService_PatchAgentSkills_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchAgentSkillsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentSkillServiceServer).PatchAgentSkills(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.AgentSkillService/PatchAgentSkills",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentSkillServiceServer).PatchAgentSkills(ctx, req.(*PatchAgentSkillsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentSkillService_DeleteAgentSkill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAgentSkillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentSkillServiceServer).DeleteAgentSkill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.AgentSkillService/DeleteAgentSkill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentSkillServiceServer).DeleteAgentSkill(ctx, req.(*DeleteAgentSkillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentSkillService_DeleteAgentSkills_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAgentSkillsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentSkillServiceServer).DeleteAgentSkills(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.AgentSkillService/DeleteAgentSkills",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentSkillServiceServer).DeleteAgentSkills(ctx, req.(*DeleteAgentSkillsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentSkillService_SearchLookupAgentNotExistsSkill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchLookupAgentNotExistsSkillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentSkillServiceServer).SearchLookupAgentNotExistsSkill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.AgentSkillService/SearchLookupAgentNotExistsSkill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentSkillServiceServer).SearchLookupAgentNotExistsSkill(ctx, req.(*SearchLookupAgentNotExistsSkillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AgentSkillService_ServiceDesc is the grpc.ServiceDesc for AgentSkillService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AgentSkillService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "engine.AgentSkillService",
	HandlerType: (*AgentSkillServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAgentSkill",
			Handler:    _AgentSkillService_CreateAgentSkill_Handler,
		},
		{
			MethodName: "CreateAgentSkills",
			Handler:    _AgentSkillService_CreateAgentSkills_Handler,
		},
		{
			MethodName: "SearchAgentSkill",
			Handler:    _AgentSkillService_SearchAgentSkill_Handler,
		},
		{
			MethodName: "ReadAgentSkill",
			Handler:    _AgentSkillService_ReadAgentSkill_Handler,
		},
		{
			MethodName: "UpdateAgentSkill",
			Handler:    _AgentSkillService_UpdateAgentSkill_Handler,
		},
		{
			MethodName: "PatchAgentSkill",
			Handler:    _AgentSkillService_PatchAgentSkill_Handler,
		},
		{
			MethodName: "PatchAgentSkills",
			Handler:    _AgentSkillService_PatchAgentSkills_Handler,
		},
		{
			MethodName: "DeleteAgentSkill",
			Handler:    _AgentSkillService_DeleteAgentSkill_Handler,
		},
		{
			MethodName: "DeleteAgentSkills",
			Handler:    _AgentSkillService_DeleteAgentSkills_Handler,
		},
		{
			MethodName: "SearchLookupAgentNotExistsSkill",
			Handler:    _AgentSkillService_SearchLookupAgentNotExistsSkill_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agent_skill.proto",
}
