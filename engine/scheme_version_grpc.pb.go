// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.6
// source: scheme_version.proto

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
	SchemeVersionService_Search_FullMethodName = "/engine.SchemeVersionService/Search"
)

// SchemeVersionServiceClient is the client API for SchemeVersionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SchemeVersionServiceClient interface {
	Search(ctx context.Context, in *SearchSchemeVersionRequest, opts ...grpc.CallOption) (*SearchSchemeVersionResponse, error)
}

type schemeVersionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSchemeVersionServiceClient(cc grpc.ClientConnInterface) SchemeVersionServiceClient {
	return &schemeVersionServiceClient{cc}
}

func (c *schemeVersionServiceClient) Search(ctx context.Context, in *SearchSchemeVersionRequest, opts ...grpc.CallOption) (*SearchSchemeVersionResponse, error) {
	out := new(SearchSchemeVersionResponse)
	err := c.cc.Invoke(ctx, SchemeVersionService_Search_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SchemeVersionServiceServer is the server API for SchemeVersionService service.
// All implementations must embed UnimplementedSchemeVersionServiceServer
// for forward compatibility
type SchemeVersionServiceServer interface {
	Search(context.Context, *SearchSchemeVersionRequest) (*SearchSchemeVersionResponse, error)
	mustEmbedUnimplementedSchemeVersionServiceServer()
}

// UnimplementedSchemeVersionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSchemeVersionServiceServer struct {
}

func (UnimplementedSchemeVersionServiceServer) Search(context.Context, *SearchSchemeVersionRequest) (*SearchSchemeVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedSchemeVersionServiceServer) mustEmbedUnimplementedSchemeVersionServiceServer() {}

// UnsafeSchemeVersionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SchemeVersionServiceServer will
// result in compilation errors.
type UnsafeSchemeVersionServiceServer interface {
	mustEmbedUnimplementedSchemeVersionServiceServer()
}

func RegisterSchemeVersionServiceServer(s grpc.ServiceRegistrar, srv SchemeVersionServiceServer) {
	s.RegisterService(&SchemeVersionService_ServiceDesc, srv)
}

func _SchemeVersionService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchSchemeVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemeVersionServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SchemeVersionService_Search_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemeVersionServiceServer).Search(ctx, req.(*SearchSchemeVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SchemeVersionService_ServiceDesc is the grpc.ServiceDesc for SchemeVersionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SchemeVersionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "engine.SchemeVersionService",
	HandlerType: (*SchemeVersionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _SchemeVersionService_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scheme_version.proto",
}
