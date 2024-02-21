// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: email_profile.proto

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

// EmailProfileServiceClient is the client API for EmailProfileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmailProfileServiceClient interface {
	// Create EmailProfile
	CreateEmailProfile(ctx context.Context, in *CreateEmailProfileRequest, opts ...grpc.CallOption) (*EmailProfile, error)
	// Search EmailProfile
	SearchEmailProfile(ctx context.Context, in *SearchEmailProfileRequest, opts ...grpc.CallOption) (*ListEmailProfile, error)
	// EmailProfile check login
	TestEmailProfile(ctx context.Context, in *TestEmailProfileRequest, opts ...grpc.CallOption) (*TestEmailProfileResponse, error)
	LoginEmailProfile(ctx context.Context, in *LoginEmailProfileRequest, opts ...grpc.CallOption) (*LoginEmailProfileResponse, error)
	// EmailProfile item
	ReadEmailProfile(ctx context.Context, in *ReadEmailProfileRequest, opts ...grpc.CallOption) (*EmailProfile, error)
	PatchEmailProfile(ctx context.Context, in *PatchEmailProfileRequest, opts ...grpc.CallOption) (*EmailProfile, error)
	// Update EmailProfile
	UpdateEmailProfile(ctx context.Context, in *UpdateEmailProfileRequest, opts ...grpc.CallOption) (*EmailProfile, error)
	// Remove EmailProfile
	DeleteEmailProfile(ctx context.Context, in *DeleteEmailProfileRequest, opts ...grpc.CallOption) (*EmailProfile, error)
}

type emailProfileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEmailProfileServiceClient(cc grpc.ClientConnInterface) EmailProfileServiceClient {
	return &emailProfileServiceClient{cc}
}

func (c *emailProfileServiceClient) CreateEmailProfile(ctx context.Context, in *CreateEmailProfileRequest, opts ...grpc.CallOption) (*EmailProfile, error) {
	out := new(EmailProfile)
	err := c.cc.Invoke(ctx, "/engine.EmailProfileService/CreateEmailProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailProfileServiceClient) SearchEmailProfile(ctx context.Context, in *SearchEmailProfileRequest, opts ...grpc.CallOption) (*ListEmailProfile, error) {
	out := new(ListEmailProfile)
	err := c.cc.Invoke(ctx, "/engine.EmailProfileService/SearchEmailProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailProfileServiceClient) TestEmailProfile(ctx context.Context, in *TestEmailProfileRequest, opts ...grpc.CallOption) (*TestEmailProfileResponse, error) {
	out := new(TestEmailProfileResponse)
	err := c.cc.Invoke(ctx, "/engine.EmailProfileService/TestEmailProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailProfileServiceClient) LoginEmailProfile(ctx context.Context, in *LoginEmailProfileRequest, opts ...grpc.CallOption) (*LoginEmailProfileResponse, error) {
	out := new(LoginEmailProfileResponse)
	err := c.cc.Invoke(ctx, "/engine.EmailProfileService/LoginEmailProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailProfileServiceClient) ReadEmailProfile(ctx context.Context, in *ReadEmailProfileRequest, opts ...grpc.CallOption) (*EmailProfile, error) {
	out := new(EmailProfile)
	err := c.cc.Invoke(ctx, "/engine.EmailProfileService/ReadEmailProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailProfileServiceClient) PatchEmailProfile(ctx context.Context, in *PatchEmailProfileRequest, opts ...grpc.CallOption) (*EmailProfile, error) {
	out := new(EmailProfile)
	err := c.cc.Invoke(ctx, "/engine.EmailProfileService/PatchEmailProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailProfileServiceClient) UpdateEmailProfile(ctx context.Context, in *UpdateEmailProfileRequest, opts ...grpc.CallOption) (*EmailProfile, error) {
	out := new(EmailProfile)
	err := c.cc.Invoke(ctx, "/engine.EmailProfileService/UpdateEmailProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailProfileServiceClient) DeleteEmailProfile(ctx context.Context, in *DeleteEmailProfileRequest, opts ...grpc.CallOption) (*EmailProfile, error) {
	out := new(EmailProfile)
	err := c.cc.Invoke(ctx, "/engine.EmailProfileService/DeleteEmailProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmailProfileServiceServer is the server API for EmailProfileService service.
// All implementations must embed UnimplementedEmailProfileServiceServer
// for forward compatibility
type EmailProfileServiceServer interface {
	// Create EmailProfile
	CreateEmailProfile(context.Context, *CreateEmailProfileRequest) (*EmailProfile, error)
	// Search EmailProfile
	SearchEmailProfile(context.Context, *SearchEmailProfileRequest) (*ListEmailProfile, error)
	// EmailProfile check login
	TestEmailProfile(context.Context, *TestEmailProfileRequest) (*TestEmailProfileResponse, error)
	LoginEmailProfile(context.Context, *LoginEmailProfileRequest) (*LoginEmailProfileResponse, error)
	// EmailProfile item
	ReadEmailProfile(context.Context, *ReadEmailProfileRequest) (*EmailProfile, error)
	PatchEmailProfile(context.Context, *PatchEmailProfileRequest) (*EmailProfile, error)
	// Update EmailProfile
	UpdateEmailProfile(context.Context, *UpdateEmailProfileRequest) (*EmailProfile, error)
	// Remove EmailProfile
	DeleteEmailProfile(context.Context, *DeleteEmailProfileRequest) (*EmailProfile, error)
	mustEmbedUnimplementedEmailProfileServiceServer()
}

// UnimplementedEmailProfileServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEmailProfileServiceServer struct {
}

func (UnimplementedEmailProfileServiceServer) CreateEmailProfile(context.Context, *CreateEmailProfileRequest) (*EmailProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEmailProfile not implemented")
}
func (UnimplementedEmailProfileServiceServer) SearchEmailProfile(context.Context, *SearchEmailProfileRequest) (*ListEmailProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchEmailProfile not implemented")
}
func (UnimplementedEmailProfileServiceServer) TestEmailProfile(context.Context, *TestEmailProfileRequest) (*TestEmailProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestEmailProfile not implemented")
}
func (UnimplementedEmailProfileServiceServer) LoginEmailProfile(context.Context, *LoginEmailProfileRequest) (*LoginEmailProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginEmailProfile not implemented")
}
func (UnimplementedEmailProfileServiceServer) ReadEmailProfile(context.Context, *ReadEmailProfileRequest) (*EmailProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadEmailProfile not implemented")
}
func (UnimplementedEmailProfileServiceServer) PatchEmailProfile(context.Context, *PatchEmailProfileRequest) (*EmailProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchEmailProfile not implemented")
}
func (UnimplementedEmailProfileServiceServer) UpdateEmailProfile(context.Context, *UpdateEmailProfileRequest) (*EmailProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEmailProfile not implemented")
}
func (UnimplementedEmailProfileServiceServer) DeleteEmailProfile(context.Context, *DeleteEmailProfileRequest) (*EmailProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEmailProfile not implemented")
}
func (UnimplementedEmailProfileServiceServer) mustEmbedUnimplementedEmailProfileServiceServer() {}

// UnsafeEmailProfileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmailProfileServiceServer will
// result in compilation errors.
type UnsafeEmailProfileServiceServer interface {
	mustEmbedUnimplementedEmailProfileServiceServer()
}

func RegisterEmailProfileServiceServer(s grpc.ServiceRegistrar, srv EmailProfileServiceServer) {
	s.RegisterService(&EmailProfileService_ServiceDesc, srv)
}

func _EmailProfileService_CreateEmailProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEmailProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailProfileServiceServer).CreateEmailProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.EmailProfileService/CreateEmailProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailProfileServiceServer).CreateEmailProfile(ctx, req.(*CreateEmailProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailProfileService_SearchEmailProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchEmailProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailProfileServiceServer).SearchEmailProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.EmailProfileService/SearchEmailProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailProfileServiceServer).SearchEmailProfile(ctx, req.(*SearchEmailProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailProfileService_TestEmailProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestEmailProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailProfileServiceServer).TestEmailProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.EmailProfileService/TestEmailProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailProfileServiceServer).TestEmailProfile(ctx, req.(*TestEmailProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailProfileService_LoginEmailProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginEmailProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailProfileServiceServer).LoginEmailProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.EmailProfileService/LoginEmailProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailProfileServiceServer).LoginEmailProfile(ctx, req.(*LoginEmailProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailProfileService_ReadEmailProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadEmailProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailProfileServiceServer).ReadEmailProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.EmailProfileService/ReadEmailProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailProfileServiceServer).ReadEmailProfile(ctx, req.(*ReadEmailProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailProfileService_PatchEmailProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchEmailProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailProfileServiceServer).PatchEmailProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.EmailProfileService/PatchEmailProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailProfileServiceServer).PatchEmailProfile(ctx, req.(*PatchEmailProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailProfileService_UpdateEmailProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEmailProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailProfileServiceServer).UpdateEmailProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.EmailProfileService/UpdateEmailProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailProfileServiceServer).UpdateEmailProfile(ctx, req.(*UpdateEmailProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailProfileService_DeleteEmailProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEmailProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailProfileServiceServer).DeleteEmailProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.EmailProfileService/DeleteEmailProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailProfileServiceServer).DeleteEmailProfile(ctx, req.(*DeleteEmailProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EmailProfileService_ServiceDesc is the grpc.ServiceDesc for EmailProfileService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EmailProfileService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "engine.EmailProfileService",
	HandlerType: (*EmailProfileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEmailProfile",
			Handler:    _EmailProfileService_CreateEmailProfile_Handler,
		},
		{
			MethodName: "SearchEmailProfile",
			Handler:    _EmailProfileService_SearchEmailProfile_Handler,
		},
		{
			MethodName: "TestEmailProfile",
			Handler:    _EmailProfileService_TestEmailProfile_Handler,
		},
		{
			MethodName: "LoginEmailProfile",
			Handler:    _EmailProfileService_LoginEmailProfile_Handler,
		},
		{
			MethodName: "ReadEmailProfile",
			Handler:    _EmailProfileService_ReadEmailProfile_Handler,
		},
		{
			MethodName: "PatchEmailProfile",
			Handler:    _EmailProfileService_PatchEmailProfile_Handler,
		},
		{
			MethodName: "UpdateEmailProfile",
			Handler:    _EmailProfileService_UpdateEmailProfile_Handler,
		},
		{
			MethodName: "DeleteEmailProfile",
			Handler:    _EmailProfileService_DeleteEmailProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "email_profile.proto",
}
