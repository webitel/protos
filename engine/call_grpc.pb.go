// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: call.proto

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

// CallServiceClient is the client API for CallService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CallServiceClient interface {
	// List of call
	SearchHistoryCall(ctx context.Context, in *SearchHistoryCallRequest, opts ...grpc.CallOption) (*ListHistoryCall, error)
	// List of call
	SearchHistoryCallPost(ctx context.Context, in *SearchHistoryCallRequest, opts ...grpc.CallOption) (*ListHistoryCall, error)
	PatchHistoryCall(ctx context.Context, in *PatchHistoryCallRequest, opts ...grpc.CallOption) (*HistoryCall, error)
	AggregateHistoryCall(ctx context.Context, in *AggregateHistoryCallRequest, opts ...grpc.CallOption) (*ListAggregate, error)
	SearchActiveCall(ctx context.Context, in *SearchCallRequest, opts ...grpc.CallOption) (*ListCall, error)
	// Call item
	ReadCall(ctx context.Context, in *ReadCallRequest, opts ...grpc.CallOption) (*ActiveCall, error)
	CreateCall(ctx context.Context, in *CreateCallRequest, opts ...grpc.CallOption) (*CreateCallResponse, error)
	HangupCall(ctx context.Context, in *HangupCallRequest, opts ...grpc.CallOption) (*HangupCallResponse, error)
	HoldCall(ctx context.Context, in *UserCallRequest, opts ...grpc.CallOption) (*HoldCallResponse, error)
	UnHoldCall(ctx context.Context, in *UserCallRequest, opts ...grpc.CallOption) (*HoldCallResponse, error)
	DtmfCall(ctx context.Context, in *DtmfCallRequest, opts ...grpc.CallOption) (*DtmfCallResponse, error)
	BlindTransferCall(ctx context.Context, in *BlindTransferCallRequest, opts ...grpc.CallOption) (*BlindTransferCallResponse, error)
	EavesdropCall(ctx context.Context, in *EavesdropCallRequest, opts ...grpc.CallOption) (*CreateCallResponse, error)
	// Call item
	ConfirmPush(ctx context.Context, in *ConfirmPushRequest, opts ...grpc.CallOption) (*ConfirmPushResponse, error)
	SetVariablesCall(ctx context.Context, in *SetVariablesCallRequest, opts ...grpc.CallOption) (*SetVariablesCallResponse, error)
	CreateCallAnnotation(ctx context.Context, in *CreateCallAnnotationRequest, opts ...grpc.CallOption) (*CallAnnotation, error)
	UpdateCallAnnotation(ctx context.Context, in *UpdateCallAnnotationRequest, opts ...grpc.CallOption) (*CallAnnotation, error)
	DeleteCallAnnotation(ctx context.Context, in *DeleteCallAnnotationRequest, opts ...grpc.CallOption) (*CallAnnotation, error)
	RedialCall(ctx context.Context, in *RedialCallRequest, opts ...grpc.CallOption) (*CreateCallResponse, error)
}

type callServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCallServiceClient(cc grpc.ClientConnInterface) CallServiceClient {
	return &callServiceClient{cc}
}

func (c *callServiceClient) SearchHistoryCall(ctx context.Context, in *SearchHistoryCallRequest, opts ...grpc.CallOption) (*ListHistoryCall, error) {
	out := new(ListHistoryCall)
	err := c.cc.Invoke(ctx, "/engine.CallService/SearchHistoryCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *callServiceClient) SearchHistoryCallPost(ctx context.Context, in *SearchHistoryCallRequest, opts ...grpc.CallOption) (*ListHistoryCall, error) {
	out := new(ListHistoryCall)
	err := c.cc.Invoke(ctx, "/engine.CallService/SearchHistoryCallPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *callServiceClient) PatchHistoryCall(ctx context.Context, in *PatchHistoryCallRequest, opts ...grpc.CallOption) (*HistoryCall, error) {
	out := new(HistoryCall)
	err := c.cc.Invoke(ctx, "/engine.CallService/PatchHistoryCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *callServiceClient) AggregateHistoryCall(ctx context.Context, in *AggregateHistoryCallRequest, opts ...grpc.CallOption) (*ListAggregate, error) {
	out := new(ListAggregate)
	err := c.cc.Invoke(ctx, "/engine.CallService/AggregateHistoryCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *callServiceClient) SearchActiveCall(ctx context.Context, in *SearchCallRequest, opts ...grpc.CallOption) (*ListCall, error) {
	out := new(ListCall)
	err := c.cc.Invoke(ctx, "/engine.CallService/SearchActiveCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *callServiceClient) ReadCall(ctx context.Context, in *ReadCallRequest, opts ...grpc.CallOption) (*ActiveCall, error) {
	out := new(ActiveCall)
	err := c.cc.Invoke(ctx, "/engine.CallService/ReadCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *callServiceClient) CreateCall(ctx context.Context, in *CreateCallRequest, opts ...grpc.CallOption) (*CreateCallResponse, error) {
	out := new(CreateCallResponse)
	err := c.cc.Invoke(ctx, "/engine.CallService/CreateCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *callServiceClient) HangupCall(ctx context.Context, in *HangupCallRequest, opts ...grpc.CallOption) (*HangupCallResponse, error) {
	out := new(HangupCallResponse)
	err := c.cc.Invoke(ctx, "/engine.CallService/HangupCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *callServiceClient) HoldCall(ctx context.Context, in *UserCallRequest, opts ...grpc.CallOption) (*HoldCallResponse, error) {
	out := new(HoldCallResponse)
	err := c.cc.Invoke(ctx, "/engine.CallService/HoldCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *callServiceClient) UnHoldCall(ctx context.Context, in *UserCallRequest, opts ...grpc.CallOption) (*HoldCallResponse, error) {
	out := new(HoldCallResponse)
	err := c.cc.Invoke(ctx, "/engine.CallService/UnHoldCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *callServiceClient) DtmfCall(ctx context.Context, in *DtmfCallRequest, opts ...grpc.CallOption) (*DtmfCallResponse, error) {
	out := new(DtmfCallResponse)
	err := c.cc.Invoke(ctx, "/engine.CallService/DtmfCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *callServiceClient) BlindTransferCall(ctx context.Context, in *BlindTransferCallRequest, opts ...grpc.CallOption) (*BlindTransferCallResponse, error) {
	out := new(BlindTransferCallResponse)
	err := c.cc.Invoke(ctx, "/engine.CallService/BlindTransferCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *callServiceClient) EavesdropCall(ctx context.Context, in *EavesdropCallRequest, opts ...grpc.CallOption) (*CreateCallResponse, error) {
	out := new(CreateCallResponse)
	err := c.cc.Invoke(ctx, "/engine.CallService/EavesdropCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *callServiceClient) ConfirmPush(ctx context.Context, in *ConfirmPushRequest, opts ...grpc.CallOption) (*ConfirmPushResponse, error) {
	out := new(ConfirmPushResponse)
	err := c.cc.Invoke(ctx, "/engine.CallService/ConfirmPush", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *callServiceClient) SetVariablesCall(ctx context.Context, in *SetVariablesCallRequest, opts ...grpc.CallOption) (*SetVariablesCallResponse, error) {
	out := new(SetVariablesCallResponse)
	err := c.cc.Invoke(ctx, "/engine.CallService/SetVariablesCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *callServiceClient) CreateCallAnnotation(ctx context.Context, in *CreateCallAnnotationRequest, opts ...grpc.CallOption) (*CallAnnotation, error) {
	out := new(CallAnnotation)
	err := c.cc.Invoke(ctx, "/engine.CallService/CreateCallAnnotation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *callServiceClient) UpdateCallAnnotation(ctx context.Context, in *UpdateCallAnnotationRequest, opts ...grpc.CallOption) (*CallAnnotation, error) {
	out := new(CallAnnotation)
	err := c.cc.Invoke(ctx, "/engine.CallService/UpdateCallAnnotation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *callServiceClient) DeleteCallAnnotation(ctx context.Context, in *DeleteCallAnnotationRequest, opts ...grpc.CallOption) (*CallAnnotation, error) {
	out := new(CallAnnotation)
	err := c.cc.Invoke(ctx, "/engine.CallService/DeleteCallAnnotation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *callServiceClient) RedialCall(ctx context.Context, in *RedialCallRequest, opts ...grpc.CallOption) (*CreateCallResponse, error) {
	out := new(CreateCallResponse)
	err := c.cc.Invoke(ctx, "/engine.CallService/RedialCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CallServiceServer is the server API for CallService service.
// All implementations must embed UnimplementedCallServiceServer
// for forward compatibility
type CallServiceServer interface {
	// List of call
	SearchHistoryCall(context.Context, *SearchHistoryCallRequest) (*ListHistoryCall, error)
	// List of call
	SearchHistoryCallPost(context.Context, *SearchHistoryCallRequest) (*ListHistoryCall, error)
	PatchHistoryCall(context.Context, *PatchHistoryCallRequest) (*HistoryCall, error)
	AggregateHistoryCall(context.Context, *AggregateHistoryCallRequest) (*ListAggregate, error)
	SearchActiveCall(context.Context, *SearchCallRequest) (*ListCall, error)
	// Call item
	ReadCall(context.Context, *ReadCallRequest) (*ActiveCall, error)
	CreateCall(context.Context, *CreateCallRequest) (*CreateCallResponse, error)
	HangupCall(context.Context, *HangupCallRequest) (*HangupCallResponse, error)
	HoldCall(context.Context, *UserCallRequest) (*HoldCallResponse, error)
	UnHoldCall(context.Context, *UserCallRequest) (*HoldCallResponse, error)
	DtmfCall(context.Context, *DtmfCallRequest) (*DtmfCallResponse, error)
	BlindTransferCall(context.Context, *BlindTransferCallRequest) (*BlindTransferCallResponse, error)
	EavesdropCall(context.Context, *EavesdropCallRequest) (*CreateCallResponse, error)
	// Call item
	ConfirmPush(context.Context, *ConfirmPushRequest) (*ConfirmPushResponse, error)
	SetVariablesCall(context.Context, *SetVariablesCallRequest) (*SetVariablesCallResponse, error)
	CreateCallAnnotation(context.Context, *CreateCallAnnotationRequest) (*CallAnnotation, error)
	UpdateCallAnnotation(context.Context, *UpdateCallAnnotationRequest) (*CallAnnotation, error)
	DeleteCallAnnotation(context.Context, *DeleteCallAnnotationRequest) (*CallAnnotation, error)
	RedialCall(context.Context, *RedialCallRequest) (*CreateCallResponse, error)
	mustEmbedUnimplementedCallServiceServer()
}

// UnimplementedCallServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCallServiceServer struct {
}

func (UnimplementedCallServiceServer) SearchHistoryCall(context.Context, *SearchHistoryCallRequest) (*ListHistoryCall, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchHistoryCall not implemented")
}
func (UnimplementedCallServiceServer) SearchHistoryCallPost(context.Context, *SearchHistoryCallRequest) (*ListHistoryCall, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchHistoryCallPost not implemented")
}
func (UnimplementedCallServiceServer) PatchHistoryCall(context.Context, *PatchHistoryCallRequest) (*HistoryCall, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchHistoryCall not implemented")
}
func (UnimplementedCallServiceServer) AggregateHistoryCall(context.Context, *AggregateHistoryCallRequest) (*ListAggregate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AggregateHistoryCall not implemented")
}
func (UnimplementedCallServiceServer) SearchActiveCall(context.Context, *SearchCallRequest) (*ListCall, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchActiveCall not implemented")
}
func (UnimplementedCallServiceServer) ReadCall(context.Context, *ReadCallRequest) (*ActiveCall, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadCall not implemented")
}
func (UnimplementedCallServiceServer) CreateCall(context.Context, *CreateCallRequest) (*CreateCallResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCall not implemented")
}
func (UnimplementedCallServiceServer) HangupCall(context.Context, *HangupCallRequest) (*HangupCallResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HangupCall not implemented")
}
func (UnimplementedCallServiceServer) HoldCall(context.Context, *UserCallRequest) (*HoldCallResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HoldCall not implemented")
}
func (UnimplementedCallServiceServer) UnHoldCall(context.Context, *UserCallRequest) (*HoldCallResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnHoldCall not implemented")
}
func (UnimplementedCallServiceServer) DtmfCall(context.Context, *DtmfCallRequest) (*DtmfCallResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DtmfCall not implemented")
}
func (UnimplementedCallServiceServer) BlindTransferCall(context.Context, *BlindTransferCallRequest) (*BlindTransferCallResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlindTransferCall not implemented")
}
func (UnimplementedCallServiceServer) EavesdropCall(context.Context, *EavesdropCallRequest) (*CreateCallResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EavesdropCall not implemented")
}
func (UnimplementedCallServiceServer) ConfirmPush(context.Context, *ConfirmPushRequest) (*ConfirmPushResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmPush not implemented")
}
func (UnimplementedCallServiceServer) SetVariablesCall(context.Context, *SetVariablesCallRequest) (*SetVariablesCallResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetVariablesCall not implemented")
}
func (UnimplementedCallServiceServer) CreateCallAnnotation(context.Context, *CreateCallAnnotationRequest) (*CallAnnotation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCallAnnotation not implemented")
}
func (UnimplementedCallServiceServer) UpdateCallAnnotation(context.Context, *UpdateCallAnnotationRequest) (*CallAnnotation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCallAnnotation not implemented")
}
func (UnimplementedCallServiceServer) DeleteCallAnnotation(context.Context, *DeleteCallAnnotationRequest) (*CallAnnotation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCallAnnotation not implemented")
}
func (UnimplementedCallServiceServer) RedialCall(context.Context, *RedialCallRequest) (*CreateCallResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RedialCall not implemented")
}
func (UnimplementedCallServiceServer) mustEmbedUnimplementedCallServiceServer() {}

// UnsafeCallServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CallServiceServer will
// result in compilation errors.
type UnsafeCallServiceServer interface {
	mustEmbedUnimplementedCallServiceServer()
}

func RegisterCallServiceServer(s grpc.ServiceRegistrar, srv CallServiceServer) {
	s.RegisterService(&CallService_ServiceDesc, srv)
}

func _CallService_SearchHistoryCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchHistoryCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallServiceServer).SearchHistoryCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.CallService/SearchHistoryCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallServiceServer).SearchHistoryCall(ctx, req.(*SearchHistoryCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CallService_SearchHistoryCallPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchHistoryCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallServiceServer).SearchHistoryCallPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.CallService/SearchHistoryCallPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallServiceServer).SearchHistoryCallPost(ctx, req.(*SearchHistoryCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CallService_PatchHistoryCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchHistoryCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallServiceServer).PatchHistoryCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.CallService/PatchHistoryCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallServiceServer).PatchHistoryCall(ctx, req.(*PatchHistoryCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CallService_AggregateHistoryCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AggregateHistoryCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallServiceServer).AggregateHistoryCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.CallService/AggregateHistoryCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallServiceServer).AggregateHistoryCall(ctx, req.(*AggregateHistoryCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CallService_SearchActiveCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallServiceServer).SearchActiveCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.CallService/SearchActiveCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallServiceServer).SearchActiveCall(ctx, req.(*SearchCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CallService_ReadCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallServiceServer).ReadCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.CallService/ReadCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallServiceServer).ReadCall(ctx, req.(*ReadCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CallService_CreateCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallServiceServer).CreateCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.CallService/CreateCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallServiceServer).CreateCall(ctx, req.(*CreateCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CallService_HangupCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HangupCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallServiceServer).HangupCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.CallService/HangupCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallServiceServer).HangupCall(ctx, req.(*HangupCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CallService_HoldCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallServiceServer).HoldCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.CallService/HoldCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallServiceServer).HoldCall(ctx, req.(*UserCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CallService_UnHoldCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallServiceServer).UnHoldCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.CallService/UnHoldCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallServiceServer).UnHoldCall(ctx, req.(*UserCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CallService_DtmfCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DtmfCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallServiceServer).DtmfCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.CallService/DtmfCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallServiceServer).DtmfCall(ctx, req.(*DtmfCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CallService_BlindTransferCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlindTransferCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallServiceServer).BlindTransferCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.CallService/BlindTransferCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallServiceServer).BlindTransferCall(ctx, req.(*BlindTransferCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CallService_EavesdropCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EavesdropCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallServiceServer).EavesdropCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.CallService/EavesdropCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallServiceServer).EavesdropCall(ctx, req.(*EavesdropCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CallService_ConfirmPush_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmPushRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallServiceServer).ConfirmPush(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.CallService/ConfirmPush",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallServiceServer).ConfirmPush(ctx, req.(*ConfirmPushRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CallService_SetVariablesCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetVariablesCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallServiceServer).SetVariablesCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.CallService/SetVariablesCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallServiceServer).SetVariablesCall(ctx, req.(*SetVariablesCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CallService_CreateCallAnnotation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCallAnnotationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallServiceServer).CreateCallAnnotation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.CallService/CreateCallAnnotation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallServiceServer).CreateCallAnnotation(ctx, req.(*CreateCallAnnotationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CallService_UpdateCallAnnotation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCallAnnotationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallServiceServer).UpdateCallAnnotation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.CallService/UpdateCallAnnotation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallServiceServer).UpdateCallAnnotation(ctx, req.(*UpdateCallAnnotationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CallService_DeleteCallAnnotation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCallAnnotationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallServiceServer).DeleteCallAnnotation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.CallService/DeleteCallAnnotation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallServiceServer).DeleteCallAnnotation(ctx, req.(*DeleteCallAnnotationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CallService_RedialCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RedialCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallServiceServer).RedialCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.CallService/RedialCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallServiceServer).RedialCall(ctx, req.(*RedialCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CallService_ServiceDesc is the grpc.ServiceDesc for CallService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CallService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "engine.CallService",
	HandlerType: (*CallServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchHistoryCall",
			Handler:    _CallService_SearchHistoryCall_Handler,
		},
		{
			MethodName: "SearchHistoryCallPost",
			Handler:    _CallService_SearchHistoryCallPost_Handler,
		},
		{
			MethodName: "PatchHistoryCall",
			Handler:    _CallService_PatchHistoryCall_Handler,
		},
		{
			MethodName: "AggregateHistoryCall",
			Handler:    _CallService_AggregateHistoryCall_Handler,
		},
		{
			MethodName: "SearchActiveCall",
			Handler:    _CallService_SearchActiveCall_Handler,
		},
		{
			MethodName: "ReadCall",
			Handler:    _CallService_ReadCall_Handler,
		},
		{
			MethodName: "CreateCall",
			Handler:    _CallService_CreateCall_Handler,
		},
		{
			MethodName: "HangupCall",
			Handler:    _CallService_HangupCall_Handler,
		},
		{
			MethodName: "HoldCall",
			Handler:    _CallService_HoldCall_Handler,
		},
		{
			MethodName: "UnHoldCall",
			Handler:    _CallService_UnHoldCall_Handler,
		},
		{
			MethodName: "DtmfCall",
			Handler:    _CallService_DtmfCall_Handler,
		},
		{
			MethodName: "BlindTransferCall",
			Handler:    _CallService_BlindTransferCall_Handler,
		},
		{
			MethodName: "EavesdropCall",
			Handler:    _CallService_EavesdropCall_Handler,
		},
		{
			MethodName: "ConfirmPush",
			Handler:    _CallService_ConfirmPush_Handler,
		},
		{
			MethodName: "SetVariablesCall",
			Handler:    _CallService_SetVariablesCall_Handler,
		},
		{
			MethodName: "CreateCallAnnotation",
			Handler:    _CallService_CreateCallAnnotation_Handler,
		},
		{
			MethodName: "UpdateCallAnnotation",
			Handler:    _CallService_UpdateCallAnnotation_Handler,
		},
		{
			MethodName: "DeleteCallAnnotation",
			Handler:    _CallService_DeleteCallAnnotation_Handler,
		},
		{
			MethodName: "RedialCall",
			Handler:    _CallService_RedialCall_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "call.proto",
}
