// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.0
// source: fs.proto

package fs

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

// ApiClient is the client API for Api service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiClient interface {
	Originate(ctx context.Context, in *OriginateRequest, opts ...grpc.CallOption) (*OriginateResponse, error)
	Execute(ctx context.Context, in *ExecuteRequest, opts ...grpc.CallOption) (*ExecuteResponse, error)
	SetVariables(ctx context.Context, in *SetVariablesRequest, opts ...grpc.CallOption) (*SetVariablesResponse, error)
	Bridge(ctx context.Context, in *BridgeRequest, opts ...grpc.CallOption) (*BridgeResponse, error)
	BridgeCall(ctx context.Context, in *BridgeCallRequest, opts ...grpc.CallOption) (*BridgeCallResponse, error)
	StopPlayback(ctx context.Context, in *StopPlaybackRequest, opts ...grpc.CallOption) (*StopPlaybackResponse, error)
	Hangup(ctx context.Context, in *HangupRequest, opts ...grpc.CallOption) (*HangupResponse, error)
	HangupMatchingVars(ctx context.Context, in *HangupMatchingVarsReqeust, opts ...grpc.CallOption) (*HangupMatchingVarsResponse, error)
	Queue(ctx context.Context, in *QueueRequest, opts ...grpc.CallOption) (*QueueResponse, error)
	HangupMany(ctx context.Context, in *HangupManyRequest, opts ...grpc.CallOption) (*HangupManyResponse, error)
	Hold(ctx context.Context, in *HoldRequest, opts ...grpc.CallOption) (*HoldResponse, error)
	UnHold(ctx context.Context, in *UnHoldRequest, opts ...grpc.CallOption) (*UnHoldResponse, error)
	SetProfileVar(ctx context.Context, in *SetProfileVarRequest, opts ...grpc.CallOption) (*SetProfileVarResponse, error)
	ConfirmPush(ctx context.Context, in *ConfirmPushRequest, opts ...grpc.CallOption) (*ConfirmPushResponse, error)
}

type apiClient struct {
	cc grpc.ClientConnInterface
}

func NewApiClient(cc grpc.ClientConnInterface) ApiClient {
	return &apiClient{cc}
}

func (c *apiClient) Originate(ctx context.Context, in *OriginateRequest, opts ...grpc.CallOption) (*OriginateResponse, error) {
	out := new(OriginateResponse)
	err := c.cc.Invoke(ctx, "/fs.Api/Originate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) Execute(ctx context.Context, in *ExecuteRequest, opts ...grpc.CallOption) (*ExecuteResponse, error) {
	out := new(ExecuteResponse)
	err := c.cc.Invoke(ctx, "/fs.Api/Execute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) SetVariables(ctx context.Context, in *SetVariablesRequest, opts ...grpc.CallOption) (*SetVariablesResponse, error) {
	out := new(SetVariablesResponse)
	err := c.cc.Invoke(ctx, "/fs.Api/SetVariables", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) Bridge(ctx context.Context, in *BridgeRequest, opts ...grpc.CallOption) (*BridgeResponse, error) {
	out := new(BridgeResponse)
	err := c.cc.Invoke(ctx, "/fs.Api/Bridge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) BridgeCall(ctx context.Context, in *BridgeCallRequest, opts ...grpc.CallOption) (*BridgeCallResponse, error) {
	out := new(BridgeCallResponse)
	err := c.cc.Invoke(ctx, "/fs.Api/BridgeCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) StopPlayback(ctx context.Context, in *StopPlaybackRequest, opts ...grpc.CallOption) (*StopPlaybackResponse, error) {
	out := new(StopPlaybackResponse)
	err := c.cc.Invoke(ctx, "/fs.Api/StopPlayback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) Hangup(ctx context.Context, in *HangupRequest, opts ...grpc.CallOption) (*HangupResponse, error) {
	out := new(HangupResponse)
	err := c.cc.Invoke(ctx, "/fs.Api/Hangup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) HangupMatchingVars(ctx context.Context, in *HangupMatchingVarsReqeust, opts ...grpc.CallOption) (*HangupMatchingVarsResponse, error) {
	out := new(HangupMatchingVarsResponse)
	err := c.cc.Invoke(ctx, "/fs.Api/HangupMatchingVars", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) Queue(ctx context.Context, in *QueueRequest, opts ...grpc.CallOption) (*QueueResponse, error) {
	out := new(QueueResponse)
	err := c.cc.Invoke(ctx, "/fs.Api/Queue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) HangupMany(ctx context.Context, in *HangupManyRequest, opts ...grpc.CallOption) (*HangupManyResponse, error) {
	out := new(HangupManyResponse)
	err := c.cc.Invoke(ctx, "/fs.Api/HangupMany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) Hold(ctx context.Context, in *HoldRequest, opts ...grpc.CallOption) (*HoldResponse, error) {
	out := new(HoldResponse)
	err := c.cc.Invoke(ctx, "/fs.Api/Hold", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) UnHold(ctx context.Context, in *UnHoldRequest, opts ...grpc.CallOption) (*UnHoldResponse, error) {
	out := new(UnHoldResponse)
	err := c.cc.Invoke(ctx, "/fs.Api/UnHold", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) SetProfileVar(ctx context.Context, in *SetProfileVarRequest, opts ...grpc.CallOption) (*SetProfileVarResponse, error) {
	out := new(SetProfileVarResponse)
	err := c.cc.Invoke(ctx, "/fs.Api/SetProfileVar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) ConfirmPush(ctx context.Context, in *ConfirmPushRequest, opts ...grpc.CallOption) (*ConfirmPushResponse, error) {
	out := new(ConfirmPushResponse)
	err := c.cc.Invoke(ctx, "/fs.Api/ConfirmPush", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiServer is the server API for Api service.
// All implementations must embed UnimplementedApiServer
// for forward compatibility
type ApiServer interface {
	Originate(context.Context, *OriginateRequest) (*OriginateResponse, error)
	Execute(context.Context, *ExecuteRequest) (*ExecuteResponse, error)
	SetVariables(context.Context, *SetVariablesRequest) (*SetVariablesResponse, error)
	Bridge(context.Context, *BridgeRequest) (*BridgeResponse, error)
	BridgeCall(context.Context, *BridgeCallRequest) (*BridgeCallResponse, error)
	StopPlayback(context.Context, *StopPlaybackRequest) (*StopPlaybackResponse, error)
	Hangup(context.Context, *HangupRequest) (*HangupResponse, error)
	HangupMatchingVars(context.Context, *HangupMatchingVarsReqeust) (*HangupMatchingVarsResponse, error)
	Queue(context.Context, *QueueRequest) (*QueueResponse, error)
	HangupMany(context.Context, *HangupManyRequest) (*HangupManyResponse, error)
	Hold(context.Context, *HoldRequest) (*HoldResponse, error)
	UnHold(context.Context, *UnHoldRequest) (*UnHoldResponse, error)
	SetProfileVar(context.Context, *SetProfileVarRequest) (*SetProfileVarResponse, error)
	ConfirmPush(context.Context, *ConfirmPushRequest) (*ConfirmPushResponse, error)
	mustEmbedUnimplementedApiServer()
}

// UnimplementedApiServer must be embedded to have forward compatible implementations.
type UnimplementedApiServer struct {
}

func (UnimplementedApiServer) Originate(context.Context, *OriginateRequest) (*OriginateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Originate not implemented")
}
func (UnimplementedApiServer) Execute(context.Context, *ExecuteRequest) (*ExecuteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Execute not implemented")
}
func (UnimplementedApiServer) SetVariables(context.Context, *SetVariablesRequest) (*SetVariablesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetVariables not implemented")
}
func (UnimplementedApiServer) Bridge(context.Context, *BridgeRequest) (*BridgeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bridge not implemented")
}
func (UnimplementedApiServer) BridgeCall(context.Context, *BridgeCallRequest) (*BridgeCallResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BridgeCall not implemented")
}
func (UnimplementedApiServer) StopPlayback(context.Context, *StopPlaybackRequest) (*StopPlaybackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopPlayback not implemented")
}
func (UnimplementedApiServer) Hangup(context.Context, *HangupRequest) (*HangupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hangup not implemented")
}
func (UnimplementedApiServer) HangupMatchingVars(context.Context, *HangupMatchingVarsReqeust) (*HangupMatchingVarsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HangupMatchingVars not implemented")
}
func (UnimplementedApiServer) Queue(context.Context, *QueueRequest) (*QueueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Queue not implemented")
}
func (UnimplementedApiServer) HangupMany(context.Context, *HangupManyRequest) (*HangupManyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HangupMany not implemented")
}
func (UnimplementedApiServer) Hold(context.Context, *HoldRequest) (*HoldResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hold not implemented")
}
func (UnimplementedApiServer) UnHold(context.Context, *UnHoldRequest) (*UnHoldResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnHold not implemented")
}
func (UnimplementedApiServer) SetProfileVar(context.Context, *SetProfileVarRequest) (*SetProfileVarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetProfileVar not implemented")
}
func (UnimplementedApiServer) ConfirmPush(context.Context, *ConfirmPushRequest) (*ConfirmPushResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmPush not implemented")
}
func (UnimplementedApiServer) mustEmbedUnimplementedApiServer() {}

// UnsafeApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiServer will
// result in compilation errors.
type UnsafeApiServer interface {
	mustEmbedUnimplementedApiServer()
}

func RegisterApiServer(s grpc.ServiceRegistrar, srv ApiServer) {
	s.RegisterService(&Api_ServiceDesc, srv)
}

func _Api_Originate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OriginateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).Originate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fs.Api/Originate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).Originate(ctx, req.(*OriginateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fs.Api/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).Execute(ctx, req.(*ExecuteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_SetVariables_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetVariablesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).SetVariables(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fs.Api/SetVariables",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).SetVariables(ctx, req.(*SetVariablesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_Bridge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BridgeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).Bridge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fs.Api/Bridge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).Bridge(ctx, req.(*BridgeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_BridgeCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BridgeCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).BridgeCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fs.Api/BridgeCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).BridgeCall(ctx, req.(*BridgeCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_StopPlayback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopPlaybackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).StopPlayback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fs.Api/StopPlayback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).StopPlayback(ctx, req.(*StopPlaybackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_Hangup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HangupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).Hangup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fs.Api/Hangup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).Hangup(ctx, req.(*HangupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_HangupMatchingVars_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HangupMatchingVarsReqeust)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).HangupMatchingVars(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fs.Api/HangupMatchingVars",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).HangupMatchingVars(ctx, req.(*HangupMatchingVarsReqeust))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_Queue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).Queue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fs.Api/Queue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).Queue(ctx, req.(*QueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_HangupMany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HangupManyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).HangupMany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fs.Api/HangupMany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).HangupMany(ctx, req.(*HangupManyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_Hold_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HoldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).Hold(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fs.Api/Hold",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).Hold(ctx, req.(*HoldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_UnHold_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnHoldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).UnHold(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fs.Api/UnHold",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).UnHold(ctx, req.(*UnHoldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_SetProfileVar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetProfileVarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).SetProfileVar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fs.Api/SetProfileVar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).SetProfileVar(ctx, req.(*SetProfileVarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_ConfirmPush_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmPushRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).ConfirmPush(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fs.Api/ConfirmPush",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).ConfirmPush(ctx, req.(*ConfirmPushRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Api_ServiceDesc is the grpc.ServiceDesc for Api service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Api_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fs.Api",
	HandlerType: (*ApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Originate",
			Handler:    _Api_Originate_Handler,
		},
		{
			MethodName: "Execute",
			Handler:    _Api_Execute_Handler,
		},
		{
			MethodName: "SetVariables",
			Handler:    _Api_SetVariables_Handler,
		},
		{
			MethodName: "Bridge",
			Handler:    _Api_Bridge_Handler,
		},
		{
			MethodName: "BridgeCall",
			Handler:    _Api_BridgeCall_Handler,
		},
		{
			MethodName: "StopPlayback",
			Handler:    _Api_StopPlayback_Handler,
		},
		{
			MethodName: "Hangup",
			Handler:    _Api_Hangup_Handler,
		},
		{
			MethodName: "HangupMatchingVars",
			Handler:    _Api_HangupMatchingVars_Handler,
		},
		{
			MethodName: "Queue",
			Handler:    _Api_Queue_Handler,
		},
		{
			MethodName: "HangupMany",
			Handler:    _Api_HangupMany_Handler,
		},
		{
			MethodName: "Hold",
			Handler:    _Api_Hold_Handler,
		},
		{
			MethodName: "UnHold",
			Handler:    _Api_UnHold_Handler,
		},
		{
			MethodName: "SetProfileVar",
			Handler:    _Api_SetProfileVar_Handler,
		},
		{
			MethodName: "ConfirmPush",
			Handler:    _Api_ConfirmPush_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fs.proto",
}