// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: schema_version_service.proto

package proto

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
	SchemaVersionsService_Search_FullMethodName = "/logger.SchemaVersionsService/Search"
)

// SchemaVersionsServiceClient is the client API for SchemaVersionsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SchemaVersionsServiceClient interface {
	Search(ctx context.Context, in *SearchSchemaVersionRequest, opts ...grpc.CallOption) (*SchemaVersions, error)
}

type schemaVersionsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSchemaVersionsServiceClient(cc grpc.ClientConnInterface) SchemaVersionsServiceClient {
	return &schemaVersionsServiceClient{cc}
}

func (c *schemaVersionsServiceClient) Search(ctx context.Context, in *SearchSchemaVersionRequest, opts ...grpc.CallOption) (*SchemaVersions, error) {
	out := new(SchemaVersions)
	err := c.cc.Invoke(ctx, SchemaVersionsService_Search_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SchemaVersionsServiceServer is the server API for SchemaVersionsService service.
// All implementations must embed UnimplementedSchemaVersionsServiceServer
// for forward compatibility
type SchemaVersionsServiceServer interface {
	Search(context.Context, *SearchSchemaVersionRequest) (*SchemaVersions, error)
	mustEmbedUnimplementedSchemaVersionsServiceServer()
}

// UnimplementedSchemaVersionsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSchemaVersionsServiceServer struct {
}

func (UnimplementedSchemaVersionsServiceServer) Search(context.Context, *SearchSchemaVersionRequest) (*SchemaVersions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedSchemaVersionsServiceServer) mustEmbedUnimplementedSchemaVersionsServiceServer() {}

// UnsafeSchemaVersionsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SchemaVersionsServiceServer will
// result in compilation errors.
type UnsafeSchemaVersionsServiceServer interface {
	mustEmbedUnimplementedSchemaVersionsServiceServer()
}

func RegisterSchemaVersionsServiceServer(s grpc.ServiceRegistrar, srv SchemaVersionsServiceServer) {
	s.RegisterService(&SchemaVersionsService_ServiceDesc, srv)
}

func _SchemaVersionsService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchSchemaVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaVersionsServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SchemaVersionsService_Search_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaVersionsServiceServer).Search(ctx, req.(*SearchSchemaVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SchemaVersionsService_ServiceDesc is the grpc.ServiceDesc for SchemaVersionsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SchemaVersionsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "logger.SchemaVersionsService",
	HandlerType: (*SchemaVersionsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _SchemaVersionsService_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "schema_version_service.proto",
}
