// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.6
// source: gateway/contacts/managers.proto

package contacts

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
	Managers_ListManagers_FullMethodName   = "/webitel.contacts.Managers/ListManagers"
	Managers_MergeManagers_FullMethodName  = "/webitel.contacts.Managers/MergeManagers"
	Managers_ResetManagers_FullMethodName  = "/webitel.contacts.Managers/ResetManagers"
	Managers_DeleteManagers_FullMethodName = "/webitel.contacts.Managers/DeleteManagers"
	Managers_LocateManager_FullMethodName  = "/webitel.contacts.Managers/LocateManager"
	Managers_UpdateManager_FullMethodName  = "/webitel.contacts.Managers/UpdateManager"
	Managers_DeleteManager_FullMethodName  = "/webitel.contacts.Managers/DeleteManager"
)

// ManagersClient is the client API for Managers service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManagersClient interface {
	// Search the Contact's Managers.
	ListManagers(ctx context.Context, in *ListManagersRequest, opts ...grpc.CallOption) (*ManagerList, error)
	// Associate new Managers to the Contact.
	MergeManagers(ctx context.Context, in *MergeManagersRequest, opts ...grpc.CallOption) (*ManagerList, error)
	// Reset Managers to fit the specified final set.
	ResetManagers(ctx context.Context, in *ResetManagersRequest, opts ...grpc.CallOption) (*ManagerList, error)
	// Remove Contact Managers associations.
	DeleteManagers(ctx context.Context, in *DeleteManagersRequest, opts ...grpc.CallOption) (*ManagerList, error)
	// Locate the manager address link.
	LocateManager(ctx context.Context, in *LocateManagerRequest, opts ...grpc.CallOption) (*Manager, error)
	// Update the contact's manager address link details
	UpdateManager(ctx context.Context, in *UpdateManagerRequest, opts ...grpc.CallOption) (*ManagerList, error)
	// Remove the contact's manager address link
	DeleteManager(ctx context.Context, in *DeleteManagerRequest, opts ...grpc.CallOption) (*Manager, error)
}

type managersClient struct {
	cc grpc.ClientConnInterface
}

func NewManagersClient(cc grpc.ClientConnInterface) ManagersClient {
	return &managersClient{cc}
}

func (c *managersClient) ListManagers(ctx context.Context, in *ListManagersRequest, opts ...grpc.CallOption) (*ManagerList, error) {
	out := new(ManagerList)
	err := c.cc.Invoke(ctx, Managers_ListManagers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managersClient) MergeManagers(ctx context.Context, in *MergeManagersRequest, opts ...grpc.CallOption) (*ManagerList, error) {
	out := new(ManagerList)
	err := c.cc.Invoke(ctx, Managers_MergeManagers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managersClient) ResetManagers(ctx context.Context, in *ResetManagersRequest, opts ...grpc.CallOption) (*ManagerList, error) {
	out := new(ManagerList)
	err := c.cc.Invoke(ctx, Managers_ResetManagers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managersClient) DeleteManagers(ctx context.Context, in *DeleteManagersRequest, opts ...grpc.CallOption) (*ManagerList, error) {
	out := new(ManagerList)
	err := c.cc.Invoke(ctx, Managers_DeleteManagers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managersClient) LocateManager(ctx context.Context, in *LocateManagerRequest, opts ...grpc.CallOption) (*Manager, error) {
	out := new(Manager)
	err := c.cc.Invoke(ctx, Managers_LocateManager_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managersClient) UpdateManager(ctx context.Context, in *UpdateManagerRequest, opts ...grpc.CallOption) (*ManagerList, error) {
	out := new(ManagerList)
	err := c.cc.Invoke(ctx, Managers_UpdateManager_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managersClient) DeleteManager(ctx context.Context, in *DeleteManagerRequest, opts ...grpc.CallOption) (*Manager, error) {
	out := new(Manager)
	err := c.cc.Invoke(ctx, Managers_DeleteManager_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagersServer is the server API for Managers service.
// All implementations must embed UnimplementedManagersServer
// for forward compatibility
type ManagersServer interface {
	// Search the Contact's Managers.
	ListManagers(context.Context, *ListManagersRequest) (*ManagerList, error)
	// Associate new Managers to the Contact.
	MergeManagers(context.Context, *MergeManagersRequest) (*ManagerList, error)
	// Reset Managers to fit the specified final set.
	ResetManagers(context.Context, *ResetManagersRequest) (*ManagerList, error)
	// Remove Contact Managers associations.
	DeleteManagers(context.Context, *DeleteManagersRequest) (*ManagerList, error)
	// Locate the manager address link.
	LocateManager(context.Context, *LocateManagerRequest) (*Manager, error)
	// Update the contact's manager address link details
	UpdateManager(context.Context, *UpdateManagerRequest) (*ManagerList, error)
	// Remove the contact's manager address link
	DeleteManager(context.Context, *DeleteManagerRequest) (*Manager, error)
	mustEmbedUnimplementedManagersServer()
}

// UnimplementedManagersServer must be embedded to have forward compatible implementations.
type UnimplementedManagersServer struct {
}

func (UnimplementedManagersServer) ListManagers(context.Context, *ListManagersRequest) (*ManagerList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListManagers not implemented")
}
func (UnimplementedManagersServer) MergeManagers(context.Context, *MergeManagersRequest) (*ManagerList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MergeManagers not implemented")
}
func (UnimplementedManagersServer) ResetManagers(context.Context, *ResetManagersRequest) (*ManagerList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetManagers not implemented")
}
func (UnimplementedManagersServer) DeleteManagers(context.Context, *DeleteManagersRequest) (*ManagerList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteManagers not implemented")
}
func (UnimplementedManagersServer) LocateManager(context.Context, *LocateManagerRequest) (*Manager, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LocateManager not implemented")
}
func (UnimplementedManagersServer) UpdateManager(context.Context, *UpdateManagerRequest) (*ManagerList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateManager not implemented")
}
func (UnimplementedManagersServer) DeleteManager(context.Context, *DeleteManagerRequest) (*Manager, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteManager not implemented")
}
func (UnimplementedManagersServer) mustEmbedUnimplementedManagersServer() {}

// UnsafeManagersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManagersServer will
// result in compilation errors.
type UnsafeManagersServer interface {
	mustEmbedUnimplementedManagersServer()
}

func RegisterManagersServer(s grpc.ServiceRegistrar, srv ManagersServer) {
	s.RegisterService(&Managers_ServiceDesc, srv)
}

func _Managers_ListManagers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListManagersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagersServer).ListManagers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Managers_ListManagers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagersServer).ListManagers(ctx, req.(*ListManagersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Managers_MergeManagers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MergeManagersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagersServer).MergeManagers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Managers_MergeManagers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagersServer).MergeManagers(ctx, req.(*MergeManagersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Managers_ResetManagers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetManagersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagersServer).ResetManagers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Managers_ResetManagers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagersServer).ResetManagers(ctx, req.(*ResetManagersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Managers_DeleteManagers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteManagersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagersServer).DeleteManagers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Managers_DeleteManagers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagersServer).DeleteManagers(ctx, req.(*DeleteManagersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Managers_LocateManager_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LocateManagerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagersServer).LocateManager(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Managers_LocateManager_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagersServer).LocateManager(ctx, req.(*LocateManagerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Managers_UpdateManager_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateManagerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagersServer).UpdateManager(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Managers_UpdateManager_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagersServer).UpdateManager(ctx, req.(*UpdateManagerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Managers_DeleteManager_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteManagerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagersServer).DeleteManager(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Managers_DeleteManager_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagersServer).DeleteManager(ctx, req.(*DeleteManagerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Managers_ServiceDesc is the grpc.ServiceDesc for Managers service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Managers_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "webitel.contacts.Managers",
	HandlerType: (*ManagersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListManagers",
			Handler:    _Managers_ListManagers_Handler,
		},
		{
			MethodName: "MergeManagers",
			Handler:    _Managers_MergeManagers_Handler,
		},
		{
			MethodName: "ResetManagers",
			Handler:    _Managers_ResetManagers_Handler,
		},
		{
			MethodName: "DeleteManagers",
			Handler:    _Managers_DeleteManagers_Handler,
		},
		{
			MethodName: "LocateManager",
			Handler:    _Managers_LocateManager_Handler,
		},
		{
			MethodName: "UpdateManager",
			Handler:    _Managers_UpdateManager_Handler,
		},
		{
			MethodName: "DeleteManager",
			Handler:    _Managers_DeleteManager_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gateway/contacts/managers.proto",
}
