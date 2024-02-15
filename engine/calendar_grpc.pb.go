// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: calendar.proto

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

// CalendarServiceClient is the client API for CalendarService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalendarServiceClient interface {
	// Create calendar
	CreateCalendar(ctx context.Context, in *CreateCalendarRequest, opts ...grpc.CallOption) (*Calendar, error)
	// List of calendar
	SearchCalendar(ctx context.Context, in *SearchCalendarRequest, opts ...grpc.CallOption) (*ListCalendar, error)
	// Calendar item
	ReadCalendar(ctx context.Context, in *ReadCalendarRequest, opts ...grpc.CallOption) (*Calendar, error)
	// Update calendar
	UpdateCalendar(ctx context.Context, in *UpdateCalendarRequest, opts ...grpc.CallOption) (*Calendar, error)
	// Remove calendar
	DeleteCalendar(ctx context.Context, in *DeleteCalendarRequest, opts ...grpc.CallOption) (*Calendar, error)
	// List timezones
	SearchTimezones(ctx context.Context, in *SearchTimezonesRequest, opts ...grpc.CallOption) (*ListTimezoneResponse, error)
}

type calendarServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalendarServiceClient(cc grpc.ClientConnInterface) CalendarServiceClient {
	return &calendarServiceClient{cc}
}

func (c *calendarServiceClient) CreateCalendar(ctx context.Context, in *CreateCalendarRequest, opts ...grpc.CallOption) (*Calendar, error) {
	out := new(Calendar)
	err := c.cc.Invoke(ctx, "/engine.CalendarService/CreateCalendar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarServiceClient) SearchCalendar(ctx context.Context, in *SearchCalendarRequest, opts ...grpc.CallOption) (*ListCalendar, error) {
	out := new(ListCalendar)
	err := c.cc.Invoke(ctx, "/engine.CalendarService/SearchCalendar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarServiceClient) ReadCalendar(ctx context.Context, in *ReadCalendarRequest, opts ...grpc.CallOption) (*Calendar, error) {
	out := new(Calendar)
	err := c.cc.Invoke(ctx, "/engine.CalendarService/ReadCalendar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarServiceClient) UpdateCalendar(ctx context.Context, in *UpdateCalendarRequest, opts ...grpc.CallOption) (*Calendar, error) {
	out := new(Calendar)
	err := c.cc.Invoke(ctx, "/engine.CalendarService/UpdateCalendar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarServiceClient) DeleteCalendar(ctx context.Context, in *DeleteCalendarRequest, opts ...grpc.CallOption) (*Calendar, error) {
	out := new(Calendar)
	err := c.cc.Invoke(ctx, "/engine.CalendarService/DeleteCalendar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarServiceClient) SearchTimezones(ctx context.Context, in *SearchTimezonesRequest, opts ...grpc.CallOption) (*ListTimezoneResponse, error) {
	out := new(ListTimezoneResponse)
	err := c.cc.Invoke(ctx, "/engine.CalendarService/SearchTimezones", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalendarServiceServer is the server API for CalendarService service.
// All implementations must embed UnimplementedCalendarServiceServer
// for forward compatibility
type CalendarServiceServer interface {
	// Create calendar
	CreateCalendar(context.Context, *CreateCalendarRequest) (*Calendar, error)
	// List of calendar
	SearchCalendar(context.Context, *SearchCalendarRequest) (*ListCalendar, error)
	// Calendar item
	ReadCalendar(context.Context, *ReadCalendarRequest) (*Calendar, error)
	// Update calendar
	UpdateCalendar(context.Context, *UpdateCalendarRequest) (*Calendar, error)
	// Remove calendar
	DeleteCalendar(context.Context, *DeleteCalendarRequest) (*Calendar, error)
	// List timezones
	SearchTimezones(context.Context, *SearchTimezonesRequest) (*ListTimezoneResponse, error)
	mustEmbedUnimplementedCalendarServiceServer()
}

// UnimplementedCalendarServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCalendarServiceServer struct {
}

func (UnimplementedCalendarServiceServer) CreateCalendar(context.Context, *CreateCalendarRequest) (*Calendar, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCalendar not implemented")
}
func (UnimplementedCalendarServiceServer) SearchCalendar(context.Context, *SearchCalendarRequest) (*ListCalendar, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchCalendar not implemented")
}
func (UnimplementedCalendarServiceServer) ReadCalendar(context.Context, *ReadCalendarRequest) (*Calendar, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadCalendar not implemented")
}
func (UnimplementedCalendarServiceServer) UpdateCalendar(context.Context, *UpdateCalendarRequest) (*Calendar, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCalendar not implemented")
}
func (UnimplementedCalendarServiceServer) DeleteCalendar(context.Context, *DeleteCalendarRequest) (*Calendar, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCalendar not implemented")
}
func (UnimplementedCalendarServiceServer) SearchTimezones(context.Context, *SearchTimezonesRequest) (*ListTimezoneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchTimezones not implemented")
}
func (UnimplementedCalendarServiceServer) mustEmbedUnimplementedCalendarServiceServer() {}

// UnsafeCalendarServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalendarServiceServer will
// result in compilation errors.
type UnsafeCalendarServiceServer interface {
	mustEmbedUnimplementedCalendarServiceServer()
}

func RegisterCalendarServiceServer(s grpc.ServiceRegistrar, srv CalendarServiceServer) {
	s.RegisterService(&CalendarService_ServiceDesc, srv)
}

func _CalendarService_CreateCalendar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCalendarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServiceServer).CreateCalendar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.CalendarService/CreateCalendar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServiceServer).CreateCalendar(ctx, req.(*CreateCalendarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarService_SearchCalendar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchCalendarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServiceServer).SearchCalendar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.CalendarService/SearchCalendar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServiceServer).SearchCalendar(ctx, req.(*SearchCalendarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarService_ReadCalendar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadCalendarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServiceServer).ReadCalendar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.CalendarService/ReadCalendar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServiceServer).ReadCalendar(ctx, req.(*ReadCalendarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarService_UpdateCalendar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCalendarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServiceServer).UpdateCalendar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.CalendarService/UpdateCalendar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServiceServer).UpdateCalendar(ctx, req.(*UpdateCalendarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarService_DeleteCalendar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCalendarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServiceServer).DeleteCalendar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.CalendarService/DeleteCalendar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServiceServer).DeleteCalendar(ctx, req.(*DeleteCalendarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalendarService_SearchTimezones_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchTimezonesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServiceServer).SearchTimezones(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.CalendarService/SearchTimezones",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServiceServer).SearchTimezones(ctx, req.(*SearchTimezonesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CalendarService_ServiceDesc is the grpc.ServiceDesc for CalendarService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalendarService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "engine.CalendarService",
	HandlerType: (*CalendarServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCalendar",
			Handler:    _CalendarService_CreateCalendar_Handler,
		},
		{
			MethodName: "SearchCalendar",
			Handler:    _CalendarService_SearchCalendar_Handler,
		},
		{
			MethodName: "ReadCalendar",
			Handler:    _CalendarService_ReadCalendar_Handler,
		},
		{
			MethodName: "UpdateCalendar",
			Handler:    _CalendarService_UpdateCalendar_Handler,
		},
		{
			MethodName: "DeleteCalendar",
			Handler:    _CalendarService_DeleteCalendar_Handler,
		},
		{
			MethodName: "SearchTimezones",
			Handler:    _CalendarService_SearchTimezones_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calendar.proto",
}
