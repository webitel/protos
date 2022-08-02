// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.0
// source: storage/cognitive_profile.proto

package storage

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

// CognitiveProfileServiceClient is the client API for CognitiveProfileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CognitiveProfileServiceClient interface {
	CreateCognitiveProfile(ctx context.Context, in *CreateCognitiveProfileRequest, opts ...grpc.CallOption) (*CognitiveProfile, error)
	SearchCognitiveProfile(ctx context.Context, in *SearchCognitiveProfileRequest, opts ...grpc.CallOption) (*ListCognitiveProfile, error)
	ReadCognitiveProfile(ctx context.Context, in *ReadCognitiveProfileRequest, opts ...grpc.CallOption) (*CognitiveProfile, error)
	UpdateCognitiveProfile(ctx context.Context, in *UpdateCognitiveProfileRequest, opts ...grpc.CallOption) (*CognitiveProfile, error)
	PatchCognitiveProfile(ctx context.Context, in *PatchCognitiveProfileRequest, opts ...grpc.CallOption) (*CognitiveProfile, error)
	// Remove BackendProfile
	DeleteCognitiveProfile(ctx context.Context, in *DeleteCognitiveProfileRequest, opts ...grpc.CallOption) (*CognitiveProfile, error)
}

type cognitiveProfileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCognitiveProfileServiceClient(cc grpc.ClientConnInterface) CognitiveProfileServiceClient {
	return &cognitiveProfileServiceClient{cc}
}

func (c *cognitiveProfileServiceClient) CreateCognitiveProfile(ctx context.Context, in *CreateCognitiveProfileRequest, opts ...grpc.CallOption) (*CognitiveProfile, error) {
	out := new(CognitiveProfile)
	err := c.cc.Invoke(ctx, "/storage.CognitiveProfileService/CreateCognitiveProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cognitiveProfileServiceClient) SearchCognitiveProfile(ctx context.Context, in *SearchCognitiveProfileRequest, opts ...grpc.CallOption) (*ListCognitiveProfile, error) {
	out := new(ListCognitiveProfile)
	err := c.cc.Invoke(ctx, "/storage.CognitiveProfileService/SearchCognitiveProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cognitiveProfileServiceClient) ReadCognitiveProfile(ctx context.Context, in *ReadCognitiveProfileRequest, opts ...grpc.CallOption) (*CognitiveProfile, error) {
	out := new(CognitiveProfile)
	err := c.cc.Invoke(ctx, "/storage.CognitiveProfileService/ReadCognitiveProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cognitiveProfileServiceClient) UpdateCognitiveProfile(ctx context.Context, in *UpdateCognitiveProfileRequest, opts ...grpc.CallOption) (*CognitiveProfile, error) {
	out := new(CognitiveProfile)
	err := c.cc.Invoke(ctx, "/storage.CognitiveProfileService/UpdateCognitiveProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cognitiveProfileServiceClient) PatchCognitiveProfile(ctx context.Context, in *PatchCognitiveProfileRequest, opts ...grpc.CallOption) (*CognitiveProfile, error) {
	out := new(CognitiveProfile)
	err := c.cc.Invoke(ctx, "/storage.CognitiveProfileService/PatchCognitiveProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cognitiveProfileServiceClient) DeleteCognitiveProfile(ctx context.Context, in *DeleteCognitiveProfileRequest, opts ...grpc.CallOption) (*CognitiveProfile, error) {
	out := new(CognitiveProfile)
	err := c.cc.Invoke(ctx, "/storage.CognitiveProfileService/DeleteCognitiveProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CognitiveProfileServiceServer is the server API for CognitiveProfileService service.
// All implementations must embed UnimplementedCognitiveProfileServiceServer
// for forward compatibility
type CognitiveProfileServiceServer interface {
	CreateCognitiveProfile(context.Context, *CreateCognitiveProfileRequest) (*CognitiveProfile, error)
	SearchCognitiveProfile(context.Context, *SearchCognitiveProfileRequest) (*ListCognitiveProfile, error)
	ReadCognitiveProfile(context.Context, *ReadCognitiveProfileRequest) (*CognitiveProfile, error)
	UpdateCognitiveProfile(context.Context, *UpdateCognitiveProfileRequest) (*CognitiveProfile, error)
	PatchCognitiveProfile(context.Context, *PatchCognitiveProfileRequest) (*CognitiveProfile, error)
	// Remove BackendProfile
	DeleteCognitiveProfile(context.Context, *DeleteCognitiveProfileRequest) (*CognitiveProfile, error)
	mustEmbedUnimplementedCognitiveProfileServiceServer()
}

// UnimplementedCognitiveProfileServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCognitiveProfileServiceServer struct {
}

func (UnimplementedCognitiveProfileServiceServer) CreateCognitiveProfile(context.Context, *CreateCognitiveProfileRequest) (*CognitiveProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCognitiveProfile not implemented")
}
func (UnimplementedCognitiveProfileServiceServer) SearchCognitiveProfile(context.Context, *SearchCognitiveProfileRequest) (*ListCognitiveProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchCognitiveProfile not implemented")
}
func (UnimplementedCognitiveProfileServiceServer) ReadCognitiveProfile(context.Context, *ReadCognitiveProfileRequest) (*CognitiveProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadCognitiveProfile not implemented")
}
func (UnimplementedCognitiveProfileServiceServer) UpdateCognitiveProfile(context.Context, *UpdateCognitiveProfileRequest) (*CognitiveProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCognitiveProfile not implemented")
}
func (UnimplementedCognitiveProfileServiceServer) PatchCognitiveProfile(context.Context, *PatchCognitiveProfileRequest) (*CognitiveProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchCognitiveProfile not implemented")
}
func (UnimplementedCognitiveProfileServiceServer) DeleteCognitiveProfile(context.Context, *DeleteCognitiveProfileRequest) (*CognitiveProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCognitiveProfile not implemented")
}
func (UnimplementedCognitiveProfileServiceServer) mustEmbedUnimplementedCognitiveProfileServiceServer() {
}

// UnsafeCognitiveProfileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CognitiveProfileServiceServer will
// result in compilation errors.
type UnsafeCognitiveProfileServiceServer interface {
	mustEmbedUnimplementedCognitiveProfileServiceServer()
}

func RegisterCognitiveProfileServiceServer(s grpc.ServiceRegistrar, srv CognitiveProfileServiceServer) {
	s.RegisterService(&CognitiveProfileService_ServiceDesc, srv)
}

func _CognitiveProfileService_CreateCognitiveProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCognitiveProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CognitiveProfileServiceServer).CreateCognitiveProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.CognitiveProfileService/CreateCognitiveProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CognitiveProfileServiceServer).CreateCognitiveProfile(ctx, req.(*CreateCognitiveProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CognitiveProfileService_SearchCognitiveProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchCognitiveProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CognitiveProfileServiceServer).SearchCognitiveProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.CognitiveProfileService/SearchCognitiveProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CognitiveProfileServiceServer).SearchCognitiveProfile(ctx, req.(*SearchCognitiveProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CognitiveProfileService_ReadCognitiveProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadCognitiveProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CognitiveProfileServiceServer).ReadCognitiveProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.CognitiveProfileService/ReadCognitiveProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CognitiveProfileServiceServer).ReadCognitiveProfile(ctx, req.(*ReadCognitiveProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CognitiveProfileService_UpdateCognitiveProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCognitiveProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CognitiveProfileServiceServer).UpdateCognitiveProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.CognitiveProfileService/UpdateCognitiveProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CognitiveProfileServiceServer).UpdateCognitiveProfile(ctx, req.(*UpdateCognitiveProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CognitiveProfileService_PatchCognitiveProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchCognitiveProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CognitiveProfileServiceServer).PatchCognitiveProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.CognitiveProfileService/PatchCognitiveProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CognitiveProfileServiceServer).PatchCognitiveProfile(ctx, req.(*PatchCognitiveProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CognitiveProfileService_DeleteCognitiveProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCognitiveProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CognitiveProfileServiceServer).DeleteCognitiveProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.CognitiveProfileService/DeleteCognitiveProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CognitiveProfileServiceServer).DeleteCognitiveProfile(ctx, req.(*DeleteCognitiveProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CognitiveProfileService_ServiceDesc is the grpc.ServiceDesc for CognitiveProfileService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CognitiveProfileService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "storage.CognitiveProfileService",
	HandlerType: (*CognitiveProfileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCognitiveProfile",
			Handler:    _CognitiveProfileService_CreateCognitiveProfile_Handler,
		},
		{
			MethodName: "SearchCognitiveProfile",
			Handler:    _CognitiveProfileService_SearchCognitiveProfile_Handler,
		},
		{
			MethodName: "ReadCognitiveProfile",
			Handler:    _CognitiveProfileService_ReadCognitiveProfile_Handler,
		},
		{
			MethodName: "UpdateCognitiveProfile",
			Handler:    _CognitiveProfileService_UpdateCognitiveProfile_Handler,
		},
		{
			MethodName: "PatchCognitiveProfile",
			Handler:    _CognitiveProfileService_PatchCognitiveProfile_Handler,
		},
		{
			MethodName: "DeleteCognitiveProfile",
			Handler:    _CognitiveProfileService_DeleteCognitiveProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "storage/cognitive_profile.proto",
}
