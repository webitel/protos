// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: queue_communications.proto

package api

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
	CommunicationTypeService_CreateCommunicationType_FullMethodName = "/engine.CommunicationTypeService/CreateCommunicationType"
	CommunicationTypeService_SearchCommunicationType_FullMethodName = "/engine.CommunicationTypeService/SearchCommunicationType"
	CommunicationTypeService_ReadCommunicationType_FullMethodName   = "/engine.CommunicationTypeService/ReadCommunicationType"
	CommunicationTypeService_UpdateCommunicationType_FullMethodName = "/engine.CommunicationTypeService/UpdateCommunicationType"
	CommunicationTypeService_PatchCommunicationType_FullMethodName  = "/engine.CommunicationTypeService/PatchCommunicationType"
	CommunicationTypeService_DeleteCommunicationType_FullMethodName = "/engine.CommunicationTypeService/DeleteCommunicationType"
)

// CommunicationTypeServiceClient is the client API for CommunicationTypeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommunicationTypeServiceClient interface {
	// Create CommunicationType
	CreateCommunicationType(ctx context.Context, in *CommunicationTypeRequest, opts ...grpc.CallOption) (*CommunicationType, error)
	// List of CommunicationType
	SearchCommunicationType(ctx context.Context, in *SearchCommunicationTypeRequest, opts ...grpc.CallOption) (*ListCommunicationType, error)
	// CommunicationType item
	ReadCommunicationType(ctx context.Context, in *ReadCommunicationTypeRequest, opts ...grpc.CallOption) (*CommunicationType, error)
	// Update CommunicationType
	UpdateCommunicationType(ctx context.Context, in *UpdateCommunicationTypeRequest, opts ...grpc.CallOption) (*CommunicationType, error)
	PatchCommunicationType(ctx context.Context, in *PatchCommunicationTypeRequest, opts ...grpc.CallOption) (*CommunicationType, error)
	// Remove CommunicationType
	DeleteCommunicationType(ctx context.Context, in *DeleteCommunicationTypeRequest, opts ...grpc.CallOption) (*CommunicationType, error)
}

type communicationTypeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommunicationTypeServiceClient(cc grpc.ClientConnInterface) CommunicationTypeServiceClient {
	return &communicationTypeServiceClient{cc}
}

func (c *communicationTypeServiceClient) CreateCommunicationType(ctx context.Context, in *CommunicationTypeRequest, opts ...grpc.CallOption) (*CommunicationType, error) {
	out := new(CommunicationType)
	err := c.cc.Invoke(ctx, CommunicationTypeService_CreateCommunicationType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicationTypeServiceClient) SearchCommunicationType(ctx context.Context, in *SearchCommunicationTypeRequest, opts ...grpc.CallOption) (*ListCommunicationType, error) {
	out := new(ListCommunicationType)
	err := c.cc.Invoke(ctx, CommunicationTypeService_SearchCommunicationType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicationTypeServiceClient) ReadCommunicationType(ctx context.Context, in *ReadCommunicationTypeRequest, opts ...grpc.CallOption) (*CommunicationType, error) {
	out := new(CommunicationType)
	err := c.cc.Invoke(ctx, CommunicationTypeService_ReadCommunicationType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicationTypeServiceClient) UpdateCommunicationType(ctx context.Context, in *UpdateCommunicationTypeRequest, opts ...grpc.CallOption) (*CommunicationType, error) {
	out := new(CommunicationType)
	err := c.cc.Invoke(ctx, CommunicationTypeService_UpdateCommunicationType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicationTypeServiceClient) PatchCommunicationType(ctx context.Context, in *PatchCommunicationTypeRequest, opts ...grpc.CallOption) (*CommunicationType, error) {
	out := new(CommunicationType)
	err := c.cc.Invoke(ctx, CommunicationTypeService_PatchCommunicationType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicationTypeServiceClient) DeleteCommunicationType(ctx context.Context, in *DeleteCommunicationTypeRequest, opts ...grpc.CallOption) (*CommunicationType, error) {
	out := new(CommunicationType)
	err := c.cc.Invoke(ctx, CommunicationTypeService_DeleteCommunicationType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommunicationTypeServiceServer is the server API for CommunicationTypeService service.
// All implementations must embed UnimplementedCommunicationTypeServiceServer
// for forward compatibility
type CommunicationTypeServiceServer interface {
	// Create CommunicationType
	CreateCommunicationType(context.Context, *CommunicationTypeRequest) (*CommunicationType, error)
	// List of CommunicationType
	SearchCommunicationType(context.Context, *SearchCommunicationTypeRequest) (*ListCommunicationType, error)
	// CommunicationType item
	ReadCommunicationType(context.Context, *ReadCommunicationTypeRequest) (*CommunicationType, error)
	// Update CommunicationType
	UpdateCommunicationType(context.Context, *UpdateCommunicationTypeRequest) (*CommunicationType, error)
	PatchCommunicationType(context.Context, *PatchCommunicationTypeRequest) (*CommunicationType, error)
	// Remove CommunicationType
	DeleteCommunicationType(context.Context, *DeleteCommunicationTypeRequest) (*CommunicationType, error)
	mustEmbedUnimplementedCommunicationTypeServiceServer()
}

// UnimplementedCommunicationTypeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCommunicationTypeServiceServer struct {
}

func (UnimplementedCommunicationTypeServiceServer) CreateCommunicationType(context.Context, *CommunicationTypeRequest) (*CommunicationType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCommunicationType not implemented")
}
func (UnimplementedCommunicationTypeServiceServer) SearchCommunicationType(context.Context, *SearchCommunicationTypeRequest) (*ListCommunicationType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchCommunicationType not implemented")
}
func (UnimplementedCommunicationTypeServiceServer) ReadCommunicationType(context.Context, *ReadCommunicationTypeRequest) (*CommunicationType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadCommunicationType not implemented")
}
func (UnimplementedCommunicationTypeServiceServer) UpdateCommunicationType(context.Context, *UpdateCommunicationTypeRequest) (*CommunicationType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCommunicationType not implemented")
}
func (UnimplementedCommunicationTypeServiceServer) PatchCommunicationType(context.Context, *PatchCommunicationTypeRequest) (*CommunicationType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchCommunicationType not implemented")
}
func (UnimplementedCommunicationTypeServiceServer) DeleteCommunicationType(context.Context, *DeleteCommunicationTypeRequest) (*CommunicationType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCommunicationType not implemented")
}
func (UnimplementedCommunicationTypeServiceServer) mustEmbedUnimplementedCommunicationTypeServiceServer() {
}

// UnsafeCommunicationTypeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommunicationTypeServiceServer will
// result in compilation errors.
type UnsafeCommunicationTypeServiceServer interface {
	mustEmbedUnimplementedCommunicationTypeServiceServer()
}

func RegisterCommunicationTypeServiceServer(s grpc.ServiceRegistrar, srv CommunicationTypeServiceServer) {
	s.RegisterService(&CommunicationTypeService_ServiceDesc, srv)
}

func _CommunicationTypeService_CreateCommunicationType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommunicationTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationTypeServiceServer).CreateCommunicationType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommunicationTypeService_CreateCommunicationType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationTypeServiceServer).CreateCommunicationType(ctx, req.(*CommunicationTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunicationTypeService_SearchCommunicationType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchCommunicationTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationTypeServiceServer).SearchCommunicationType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommunicationTypeService_SearchCommunicationType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationTypeServiceServer).SearchCommunicationType(ctx, req.(*SearchCommunicationTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunicationTypeService_ReadCommunicationType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadCommunicationTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationTypeServiceServer).ReadCommunicationType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommunicationTypeService_ReadCommunicationType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationTypeServiceServer).ReadCommunicationType(ctx, req.(*ReadCommunicationTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunicationTypeService_UpdateCommunicationType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCommunicationTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationTypeServiceServer).UpdateCommunicationType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommunicationTypeService_UpdateCommunicationType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationTypeServiceServer).UpdateCommunicationType(ctx, req.(*UpdateCommunicationTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunicationTypeService_PatchCommunicationType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchCommunicationTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationTypeServiceServer).PatchCommunicationType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommunicationTypeService_PatchCommunicationType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationTypeServiceServer).PatchCommunicationType(ctx, req.(*PatchCommunicationTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunicationTypeService_DeleteCommunicationType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCommunicationTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationTypeServiceServer).DeleteCommunicationType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommunicationTypeService_DeleteCommunicationType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationTypeServiceServer).DeleteCommunicationType(ctx, req.(*DeleteCommunicationTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CommunicationTypeService_ServiceDesc is the grpc.ServiceDesc for CommunicationTypeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommunicationTypeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "engine.CommunicationTypeService",
	HandlerType: (*CommunicationTypeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCommunicationType",
			Handler:    _CommunicationTypeService_CreateCommunicationType_Handler,
		},
		{
			MethodName: "SearchCommunicationType",
			Handler:    _CommunicationTypeService_SearchCommunicationType_Handler,
		},
		{
			MethodName: "ReadCommunicationType",
			Handler:    _CommunicationTypeService_ReadCommunicationType_Handler,
		},
		{
			MethodName: "UpdateCommunicationType",
			Handler:    _CommunicationTypeService_UpdateCommunicationType_Handler,
		},
		{
			MethodName: "PatchCommunicationType",
			Handler:    _CommunicationTypeService_PatchCommunicationType_Handler,
		},
		{
			MethodName: "DeleteCommunicationType",
			Handler:    _CommunicationTypeService_DeleteCommunicationType_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "queue_communications.proto",
}