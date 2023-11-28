// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.6
// source: user_helper.proto

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
	UserHelperService_DefaultDeviceConfig_FullMethodName     = "/engine.UserHelperService/DefaultDeviceConfig"
	UserHelperService_ActivityWorkspaceWidget_FullMethodName = "/engine.UserHelperService/ActivityWorkspaceWidget"
)

// UserHelperServiceClient is the client API for UserHelperService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserHelperServiceClient interface {
	DefaultDeviceConfig(ctx context.Context, in *DefaultDeviceConfigRequest, opts ...grpc.CallOption) (*DefaultDeviceConfigResponse, error)
	ActivityWorkspaceWidget(ctx context.Context, in *ActivityWorkspaceWidgetRequest, opts ...grpc.CallOption) (*ActivityWorkspaceWidgetResponse, error)
}

type userHelperServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserHelperServiceClient(cc grpc.ClientConnInterface) UserHelperServiceClient {
	return &userHelperServiceClient{cc}
}

func (c *userHelperServiceClient) DefaultDeviceConfig(ctx context.Context, in *DefaultDeviceConfigRequest, opts ...grpc.CallOption) (*DefaultDeviceConfigResponse, error) {
	out := new(DefaultDeviceConfigResponse)
	err := c.cc.Invoke(ctx, UserHelperService_DefaultDeviceConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userHelperServiceClient) ActivityWorkspaceWidget(ctx context.Context, in *ActivityWorkspaceWidgetRequest, opts ...grpc.CallOption) (*ActivityWorkspaceWidgetResponse, error) {
	out := new(ActivityWorkspaceWidgetResponse)
	err := c.cc.Invoke(ctx, UserHelperService_ActivityWorkspaceWidget_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserHelperServiceServer is the server API for UserHelperService service.
// All implementations must embed UnimplementedUserHelperServiceServer
// for forward compatibility
type UserHelperServiceServer interface {
	DefaultDeviceConfig(context.Context, *DefaultDeviceConfigRequest) (*DefaultDeviceConfigResponse, error)
	ActivityWorkspaceWidget(context.Context, *ActivityWorkspaceWidgetRequest) (*ActivityWorkspaceWidgetResponse, error)
	mustEmbedUnimplementedUserHelperServiceServer()
}

// UnimplementedUserHelperServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserHelperServiceServer struct {
}

func (UnimplementedUserHelperServiceServer) DefaultDeviceConfig(context.Context, *DefaultDeviceConfigRequest) (*DefaultDeviceConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DefaultDeviceConfig not implemented")
}
func (UnimplementedUserHelperServiceServer) ActivityWorkspaceWidget(context.Context, *ActivityWorkspaceWidgetRequest) (*ActivityWorkspaceWidgetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivityWorkspaceWidget not implemented")
}
func (UnimplementedUserHelperServiceServer) mustEmbedUnimplementedUserHelperServiceServer() {}

// UnsafeUserHelperServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserHelperServiceServer will
// result in compilation errors.
type UnsafeUserHelperServiceServer interface {
	mustEmbedUnimplementedUserHelperServiceServer()
}

func RegisterUserHelperServiceServer(s grpc.ServiceRegistrar, srv UserHelperServiceServer) {
	s.RegisterService(&UserHelperService_ServiceDesc, srv)
}

func _UserHelperService_DefaultDeviceConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DefaultDeviceConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserHelperServiceServer).DefaultDeviceConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserHelperService_DefaultDeviceConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserHelperServiceServer).DefaultDeviceConfig(ctx, req.(*DefaultDeviceConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserHelperService_ActivityWorkspaceWidget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivityWorkspaceWidgetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserHelperServiceServer).ActivityWorkspaceWidget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserHelperService_ActivityWorkspaceWidget_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserHelperServiceServer).ActivityWorkspaceWidget(ctx, req.(*ActivityWorkspaceWidgetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserHelperService_ServiceDesc is the grpc.ServiceDesc for UserHelperService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserHelperService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "engine.UserHelperService",
	HandlerType: (*UserHelperServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DefaultDeviceConfig",
			Handler:    _UserHelperService_DefaultDeviceConfig_Handler,
		},
		{
			MethodName: "ActivityWorkspaceWidget",
			Handler:    _UserHelperService_ActivityWorkspaceWidget_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_helper.proto",
}
