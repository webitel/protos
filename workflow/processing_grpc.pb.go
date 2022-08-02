// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.0
// source: workflow/processing.proto

package workflow

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

// FlowProcessingServiceClient is the client API for FlowProcessingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FlowProcessingServiceClient interface {
	StartProcessing(ctx context.Context, in *StartProcessingRequest, opts ...grpc.CallOption) (*Form, error)
	FormAction(ctx context.Context, in *FormActionRequest, opts ...grpc.CallOption) (*Form, error)
	CancelProcessing(ctx context.Context, in *CancelProcessingRequest, opts ...grpc.CallOption) (*CancelProcessingResponse, error)
}

type flowProcessingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFlowProcessingServiceClient(cc grpc.ClientConnInterface) FlowProcessingServiceClient {
	return &flowProcessingServiceClient{cc}
}

func (c *flowProcessingServiceClient) StartProcessing(ctx context.Context, in *StartProcessingRequest, opts ...grpc.CallOption) (*Form, error) {
	out := new(Form)
	err := c.cc.Invoke(ctx, "/workflow.FlowProcessingService/StartProcessing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowProcessingServiceClient) FormAction(ctx context.Context, in *FormActionRequest, opts ...grpc.CallOption) (*Form, error) {
	out := new(Form)
	err := c.cc.Invoke(ctx, "/workflow.FlowProcessingService/FormAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowProcessingServiceClient) CancelProcessing(ctx context.Context, in *CancelProcessingRequest, opts ...grpc.CallOption) (*CancelProcessingResponse, error) {
	out := new(CancelProcessingResponse)
	err := c.cc.Invoke(ctx, "/workflow.FlowProcessingService/CancelProcessing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FlowProcessingServiceServer is the server API for FlowProcessingService service.
// All implementations must embed UnimplementedFlowProcessingServiceServer
// for forward compatibility
type FlowProcessingServiceServer interface {
	StartProcessing(context.Context, *StartProcessingRequest) (*Form, error)
	FormAction(context.Context, *FormActionRequest) (*Form, error)
	CancelProcessing(context.Context, *CancelProcessingRequest) (*CancelProcessingResponse, error)
	mustEmbedUnimplementedFlowProcessingServiceServer()
}

// UnimplementedFlowProcessingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFlowProcessingServiceServer struct {
}

func (UnimplementedFlowProcessingServiceServer) StartProcessing(context.Context, *StartProcessingRequest) (*Form, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartProcessing not implemented")
}
func (UnimplementedFlowProcessingServiceServer) FormAction(context.Context, *FormActionRequest) (*Form, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FormAction not implemented")
}
func (UnimplementedFlowProcessingServiceServer) CancelProcessing(context.Context, *CancelProcessingRequest) (*CancelProcessingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelProcessing not implemented")
}
func (UnimplementedFlowProcessingServiceServer) mustEmbedUnimplementedFlowProcessingServiceServer() {}

// UnsafeFlowProcessingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FlowProcessingServiceServer will
// result in compilation errors.
type UnsafeFlowProcessingServiceServer interface {
	mustEmbedUnimplementedFlowProcessingServiceServer()
}

func RegisterFlowProcessingServiceServer(s grpc.ServiceRegistrar, srv FlowProcessingServiceServer) {
	s.RegisterService(&FlowProcessingService_ServiceDesc, srv)
}

func _FlowProcessingService_StartProcessing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartProcessingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlowProcessingServiceServer).StartProcessing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.FlowProcessingService/StartProcessing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlowProcessingServiceServer).StartProcessing(ctx, req.(*StartProcessingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlowProcessingService_FormAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FormActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlowProcessingServiceServer).FormAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.FlowProcessingService/FormAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlowProcessingServiceServer).FormAction(ctx, req.(*FormActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlowProcessingService_CancelProcessing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelProcessingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlowProcessingServiceServer).CancelProcessing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.FlowProcessingService/CancelProcessing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlowProcessingServiceServer).CancelProcessing(ctx, req.(*CancelProcessingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FlowProcessingService_ServiceDesc is the grpc.ServiceDesc for FlowProcessingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FlowProcessingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "workflow.FlowProcessingService",
	HandlerType: (*FlowProcessingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartProcessing",
			Handler:    _FlowProcessingService_StartProcessing_Handler,
		},
		{
			MethodName: "FormAction",
			Handler:    _FlowProcessingService_FormAction_Handler,
		},
		{
			MethodName: "CancelProcessing",
			Handler:    _FlowProcessingService_CancelProcessing_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "workflow/processing.proto",
}
