// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.6
// source: preset_query.proto

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
	PresetQueryService_CreatePresetQuery_FullMethodName = "/engine.PresetQueryService/CreatePresetQuery"
	PresetQueryService_SearchPresetQuery_FullMethodName = "/engine.PresetQueryService/SearchPresetQuery"
	PresetQueryService_ReadPresetQuery_FullMethodName   = "/engine.PresetQueryService/ReadPresetQuery"
	PresetQueryService_UpdatePresetQuery_FullMethodName = "/engine.PresetQueryService/UpdatePresetQuery"
	PresetQueryService_PatchPresetQuery_FullMethodName  = "/engine.PresetQueryService/PatchPresetQuery"
	PresetQueryService_DeletePresetQuery_FullMethodName = "/engine.PresetQueryService/DeletePresetQuery"
)

// PresetQueryServiceClient is the client API for PresetQueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PresetQueryServiceClient interface {
	CreatePresetQuery(ctx context.Context, in *CreatePresetQueryRequest, opts ...grpc.CallOption) (*PresetQuery, error)
	SearchPresetQuery(ctx context.Context, in *SearchPresetQueryRequest, opts ...grpc.CallOption) (*ListPresetQuery, error)
	ReadPresetQuery(ctx context.Context, in *ReadPresetQueryRequest, opts ...grpc.CallOption) (*PresetQuery, error)
	UpdatePresetQuery(ctx context.Context, in *UpdatePresetQueryRequest, opts ...grpc.CallOption) (*PresetQuery, error)
	PatchPresetQuery(ctx context.Context, in *PatchPresetQueryRequest, opts ...grpc.CallOption) (*PresetQuery, error)
	DeletePresetQuery(ctx context.Context, in *DeletePresetQueryRequest, opts ...grpc.CallOption) (*PresetQuery, error)
}

type presetQueryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPresetQueryServiceClient(cc grpc.ClientConnInterface) PresetQueryServiceClient {
	return &presetQueryServiceClient{cc}
}

func (c *presetQueryServiceClient) CreatePresetQuery(ctx context.Context, in *CreatePresetQueryRequest, opts ...grpc.CallOption) (*PresetQuery, error) {
	out := new(PresetQuery)
	err := c.cc.Invoke(ctx, PresetQueryService_CreatePresetQuery_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presetQueryServiceClient) SearchPresetQuery(ctx context.Context, in *SearchPresetQueryRequest, opts ...grpc.CallOption) (*ListPresetQuery, error) {
	out := new(ListPresetQuery)
	err := c.cc.Invoke(ctx, PresetQueryService_SearchPresetQuery_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presetQueryServiceClient) ReadPresetQuery(ctx context.Context, in *ReadPresetQueryRequest, opts ...grpc.CallOption) (*PresetQuery, error) {
	out := new(PresetQuery)
	err := c.cc.Invoke(ctx, PresetQueryService_ReadPresetQuery_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presetQueryServiceClient) UpdatePresetQuery(ctx context.Context, in *UpdatePresetQueryRequest, opts ...grpc.CallOption) (*PresetQuery, error) {
	out := new(PresetQuery)
	err := c.cc.Invoke(ctx, PresetQueryService_UpdatePresetQuery_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presetQueryServiceClient) PatchPresetQuery(ctx context.Context, in *PatchPresetQueryRequest, opts ...grpc.CallOption) (*PresetQuery, error) {
	out := new(PresetQuery)
	err := c.cc.Invoke(ctx, PresetQueryService_PatchPresetQuery_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presetQueryServiceClient) DeletePresetQuery(ctx context.Context, in *DeletePresetQueryRequest, opts ...grpc.CallOption) (*PresetQuery, error) {
	out := new(PresetQuery)
	err := c.cc.Invoke(ctx, PresetQueryService_DeletePresetQuery_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PresetQueryServiceServer is the server API for PresetQueryService service.
// All implementations must embed UnimplementedPresetQueryServiceServer
// for forward compatibility
type PresetQueryServiceServer interface {
	CreatePresetQuery(context.Context, *CreatePresetQueryRequest) (*PresetQuery, error)
	SearchPresetQuery(context.Context, *SearchPresetQueryRequest) (*ListPresetQuery, error)
	ReadPresetQuery(context.Context, *ReadPresetQueryRequest) (*PresetQuery, error)
	UpdatePresetQuery(context.Context, *UpdatePresetQueryRequest) (*PresetQuery, error)
	PatchPresetQuery(context.Context, *PatchPresetQueryRequest) (*PresetQuery, error)
	DeletePresetQuery(context.Context, *DeletePresetQueryRequest) (*PresetQuery, error)
	mustEmbedUnimplementedPresetQueryServiceServer()
}

// UnimplementedPresetQueryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPresetQueryServiceServer struct {
}

func (UnimplementedPresetQueryServiceServer) CreatePresetQuery(context.Context, *CreatePresetQueryRequest) (*PresetQuery, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePresetQuery not implemented")
}
func (UnimplementedPresetQueryServiceServer) SearchPresetQuery(context.Context, *SearchPresetQueryRequest) (*ListPresetQuery, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchPresetQuery not implemented")
}
func (UnimplementedPresetQueryServiceServer) ReadPresetQuery(context.Context, *ReadPresetQueryRequest) (*PresetQuery, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadPresetQuery not implemented")
}
func (UnimplementedPresetQueryServiceServer) UpdatePresetQuery(context.Context, *UpdatePresetQueryRequest) (*PresetQuery, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePresetQuery not implemented")
}
func (UnimplementedPresetQueryServiceServer) PatchPresetQuery(context.Context, *PatchPresetQueryRequest) (*PresetQuery, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchPresetQuery not implemented")
}
func (UnimplementedPresetQueryServiceServer) DeletePresetQuery(context.Context, *DeletePresetQueryRequest) (*PresetQuery, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePresetQuery not implemented")
}
func (UnimplementedPresetQueryServiceServer) mustEmbedUnimplementedPresetQueryServiceServer() {}

// UnsafePresetQueryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PresetQueryServiceServer will
// result in compilation errors.
type UnsafePresetQueryServiceServer interface {
	mustEmbedUnimplementedPresetQueryServiceServer()
}

func RegisterPresetQueryServiceServer(s grpc.ServiceRegistrar, srv PresetQueryServiceServer) {
	s.RegisterService(&PresetQueryService_ServiceDesc, srv)
}

func _PresetQueryService_CreatePresetQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePresetQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresetQueryServiceServer).CreatePresetQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PresetQueryService_CreatePresetQuery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresetQueryServiceServer).CreatePresetQuery(ctx, req.(*CreatePresetQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PresetQueryService_SearchPresetQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchPresetQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresetQueryServiceServer).SearchPresetQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PresetQueryService_SearchPresetQuery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresetQueryServiceServer).SearchPresetQuery(ctx, req.(*SearchPresetQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PresetQueryService_ReadPresetQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadPresetQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresetQueryServiceServer).ReadPresetQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PresetQueryService_ReadPresetQuery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresetQueryServiceServer).ReadPresetQuery(ctx, req.(*ReadPresetQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PresetQueryService_UpdatePresetQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePresetQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresetQueryServiceServer).UpdatePresetQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PresetQueryService_UpdatePresetQuery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresetQueryServiceServer).UpdatePresetQuery(ctx, req.(*UpdatePresetQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PresetQueryService_PatchPresetQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchPresetQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresetQueryServiceServer).PatchPresetQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PresetQueryService_PatchPresetQuery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresetQueryServiceServer).PatchPresetQuery(ctx, req.(*PatchPresetQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PresetQueryService_DeletePresetQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePresetQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresetQueryServiceServer).DeletePresetQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PresetQueryService_DeletePresetQuery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresetQueryServiceServer).DeletePresetQuery(ctx, req.(*DeletePresetQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PresetQueryService_ServiceDesc is the grpc.ServiceDesc for PresetQueryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PresetQueryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "engine.PresetQueryService",
	HandlerType: (*PresetQueryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePresetQuery",
			Handler:    _PresetQueryService_CreatePresetQuery_Handler,
		},
		{
			MethodName: "SearchPresetQuery",
			Handler:    _PresetQueryService_SearchPresetQuery_Handler,
		},
		{
			MethodName: "ReadPresetQuery",
			Handler:    _PresetQueryService_ReadPresetQuery_Handler,
		},
		{
			MethodName: "UpdatePresetQuery",
			Handler:    _PresetQueryService_UpdatePresetQuery_Handler,
		},
		{
			MethodName: "PatchPresetQuery",
			Handler:    _PresetQueryService_PatchPresetQuery_Handler,
		},
		{
			MethodName: "DeletePresetQuery",
			Handler:    _PresetQueryService_DeletePresetQuery_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "preset_query.proto",
}
