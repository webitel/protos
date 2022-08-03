// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: chat.proto

package chat

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

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	// SendMessage [FROM] created channel_id (+auth_user_id) [TO] conversation_id chat-room
	SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error)
	// StartConversation starts bot's (.user.type:.user.connection) flow schema NEW routine
	StartConversation(ctx context.Context, in *StartConversationRequest, opts ...grpc.CallOption) (*StartConversationResponse, error)
	// CloseConversation stops and close chat-bot's schema routine with all it's recipient(s)
	CloseConversation(ctx context.Context, in *CloseConversationRequest, opts ...grpc.CallOption) (*CloseConversationResponse, error)
	// JoinConversation accepts user's invitation to chat conversation
	JoinConversation(ctx context.Context, in *JoinConversationRequest, opts ...grpc.CallOption) (*JoinConversationResponse, error)
	// LeaveConversation kicks requested user from chat conversation
	LeaveConversation(ctx context.Context, in *LeaveConversationRequest, opts ...grpc.CallOption) (*LeaveConversationResponse, error)
	// InviteToConversation publish NEW invitation for .user
	InviteToConversation(ctx context.Context, in *InviteToConversationRequest, opts ...grpc.CallOption) (*InviteToConversationResponse, error)
	// DeclineInvitation declines chat invitation FROM user
	DeclineInvitation(ctx context.Context, in *DeclineInvitationRequest, opts ...grpc.CallOption) (*DeclineInvitationResponse, error)
	// CheckSession returns internal chat channel for external chat user
	CheckSession(ctx context.Context, in *CheckSessionRequest, opts ...grpc.CallOption) (*CheckSessionResponse, error)
	WaitMessage(ctx context.Context, in *WaitMessageRequest, opts ...grpc.CallOption) (*WaitMessageResponse, error)
	UpdateChannel(ctx context.Context, in *UpdateChannelRequest, opts ...grpc.CallOption) (*UpdateChannelResponse, error)
	GetConversations(ctx context.Context, in *GetConversationsRequest, opts ...grpc.CallOption) (*GetConversationsResponse, error)
	GetConversationByID(ctx context.Context, in *GetConversationByIDRequest, opts ...grpc.CallOption) (*GetConversationByIDResponse, error)
	GetHistoryMessages(ctx context.Context, in *GetHistoryMessagesRequest, opts ...grpc.CallOption) (*GetHistoryMessagesResponse, error)
	// API /v1
	SetVariables(ctx context.Context, in *SetVariablesRequest, opts ...grpc.CallOption) (*ChatVariablesResponse, error)
	BlindTransfer(ctx context.Context, in *ChatTransferRequest, opts ...grpc.CallOption) (*ChatTransferResponse, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error) {
	out := new(SendMessageResponse)
	err := c.cc.Invoke(ctx, "/webitel.chat.server.ChatService/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) StartConversation(ctx context.Context, in *StartConversationRequest, opts ...grpc.CallOption) (*StartConversationResponse, error) {
	out := new(StartConversationResponse)
	err := c.cc.Invoke(ctx, "/webitel.chat.server.ChatService/StartConversation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) CloseConversation(ctx context.Context, in *CloseConversationRequest, opts ...grpc.CallOption) (*CloseConversationResponse, error) {
	out := new(CloseConversationResponse)
	err := c.cc.Invoke(ctx, "/webitel.chat.server.ChatService/CloseConversation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) JoinConversation(ctx context.Context, in *JoinConversationRequest, opts ...grpc.CallOption) (*JoinConversationResponse, error) {
	out := new(JoinConversationResponse)
	err := c.cc.Invoke(ctx, "/webitel.chat.server.ChatService/JoinConversation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) LeaveConversation(ctx context.Context, in *LeaveConversationRequest, opts ...grpc.CallOption) (*LeaveConversationResponse, error) {
	out := new(LeaveConversationResponse)
	err := c.cc.Invoke(ctx, "/webitel.chat.server.ChatService/LeaveConversation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) InviteToConversation(ctx context.Context, in *InviteToConversationRequest, opts ...grpc.CallOption) (*InviteToConversationResponse, error) {
	out := new(InviteToConversationResponse)
	err := c.cc.Invoke(ctx, "/webitel.chat.server.ChatService/InviteToConversation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) DeclineInvitation(ctx context.Context, in *DeclineInvitationRequest, opts ...grpc.CallOption) (*DeclineInvitationResponse, error) {
	out := new(DeclineInvitationResponse)
	err := c.cc.Invoke(ctx, "/webitel.chat.server.ChatService/DeclineInvitation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) CheckSession(ctx context.Context, in *CheckSessionRequest, opts ...grpc.CallOption) (*CheckSessionResponse, error) {
	out := new(CheckSessionResponse)
	err := c.cc.Invoke(ctx, "/webitel.chat.server.ChatService/CheckSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) WaitMessage(ctx context.Context, in *WaitMessageRequest, opts ...grpc.CallOption) (*WaitMessageResponse, error) {
	out := new(WaitMessageResponse)
	err := c.cc.Invoke(ctx, "/webitel.chat.server.ChatService/WaitMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) UpdateChannel(ctx context.Context, in *UpdateChannelRequest, opts ...grpc.CallOption) (*UpdateChannelResponse, error) {
	out := new(UpdateChannelResponse)
	err := c.cc.Invoke(ctx, "/webitel.chat.server.ChatService/UpdateChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetConversations(ctx context.Context, in *GetConversationsRequest, opts ...grpc.CallOption) (*GetConversationsResponse, error) {
	out := new(GetConversationsResponse)
	err := c.cc.Invoke(ctx, "/webitel.chat.server.ChatService/GetConversations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetConversationByID(ctx context.Context, in *GetConversationByIDRequest, opts ...grpc.CallOption) (*GetConversationByIDResponse, error) {
	out := new(GetConversationByIDResponse)
	err := c.cc.Invoke(ctx, "/webitel.chat.server.ChatService/GetConversationByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetHistoryMessages(ctx context.Context, in *GetHistoryMessagesRequest, opts ...grpc.CallOption) (*GetHistoryMessagesResponse, error) {
	out := new(GetHistoryMessagesResponse)
	err := c.cc.Invoke(ctx, "/webitel.chat.server.ChatService/GetHistoryMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) SetVariables(ctx context.Context, in *SetVariablesRequest, opts ...grpc.CallOption) (*ChatVariablesResponse, error) {
	out := new(ChatVariablesResponse)
	err := c.cc.Invoke(ctx, "/webitel.chat.server.ChatService/SetVariables", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) BlindTransfer(ctx context.Context, in *ChatTransferRequest, opts ...grpc.CallOption) (*ChatTransferResponse, error) {
	out := new(ChatTransferResponse)
	err := c.cc.Invoke(ctx, "/webitel.chat.server.ChatService/BlindTransfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations must embed UnimplementedChatServiceServer
// for forward compatibility
type ChatServiceServer interface {
	// SendMessage [FROM] created channel_id (+auth_user_id) [TO] conversation_id chat-room
	SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error)
	// StartConversation starts bot's (.user.type:.user.connection) flow schema NEW routine
	StartConversation(context.Context, *StartConversationRequest) (*StartConversationResponse, error)
	// CloseConversation stops and close chat-bot's schema routine with all it's recipient(s)
	CloseConversation(context.Context, *CloseConversationRequest) (*CloseConversationResponse, error)
	// JoinConversation accepts user's invitation to chat conversation
	JoinConversation(context.Context, *JoinConversationRequest) (*JoinConversationResponse, error)
	// LeaveConversation kicks requested user from chat conversation
	LeaveConversation(context.Context, *LeaveConversationRequest) (*LeaveConversationResponse, error)
	// InviteToConversation publish NEW invitation for .user
	InviteToConversation(context.Context, *InviteToConversationRequest) (*InviteToConversationResponse, error)
	// DeclineInvitation declines chat invitation FROM user
	DeclineInvitation(context.Context, *DeclineInvitationRequest) (*DeclineInvitationResponse, error)
	// CheckSession returns internal chat channel for external chat user
	CheckSession(context.Context, *CheckSessionRequest) (*CheckSessionResponse, error)
	WaitMessage(context.Context, *WaitMessageRequest) (*WaitMessageResponse, error)
	UpdateChannel(context.Context, *UpdateChannelRequest) (*UpdateChannelResponse, error)
	GetConversations(context.Context, *GetConversationsRequest) (*GetConversationsResponse, error)
	GetConversationByID(context.Context, *GetConversationByIDRequest) (*GetConversationByIDResponse, error)
	GetHistoryMessages(context.Context, *GetHistoryMessagesRequest) (*GetHistoryMessagesResponse, error)
	// API /v1
	SetVariables(context.Context, *SetVariablesRequest) (*ChatVariablesResponse, error)
	BlindTransfer(context.Context, *ChatTransferRequest) (*ChatTransferResponse, error)
	mustEmbedUnimplementedChatServiceServer()
}

// UnimplementedChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (UnimplementedChatServiceServer) SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedChatServiceServer) StartConversation(context.Context, *StartConversationRequest) (*StartConversationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartConversation not implemented")
}
func (UnimplementedChatServiceServer) CloseConversation(context.Context, *CloseConversationRequest) (*CloseConversationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseConversation not implemented")
}
func (UnimplementedChatServiceServer) JoinConversation(context.Context, *JoinConversationRequest) (*JoinConversationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinConversation not implemented")
}
func (UnimplementedChatServiceServer) LeaveConversation(context.Context, *LeaveConversationRequest) (*LeaveConversationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveConversation not implemented")
}
func (UnimplementedChatServiceServer) InviteToConversation(context.Context, *InviteToConversationRequest) (*InviteToConversationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InviteToConversation not implemented")
}
func (UnimplementedChatServiceServer) DeclineInvitation(context.Context, *DeclineInvitationRequest) (*DeclineInvitationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeclineInvitation not implemented")
}
func (UnimplementedChatServiceServer) CheckSession(context.Context, *CheckSessionRequest) (*CheckSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckSession not implemented")
}
func (UnimplementedChatServiceServer) WaitMessage(context.Context, *WaitMessageRequest) (*WaitMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WaitMessage not implemented")
}
func (UnimplementedChatServiceServer) UpdateChannel(context.Context, *UpdateChannelRequest) (*UpdateChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateChannel not implemented")
}
func (UnimplementedChatServiceServer) GetConversations(context.Context, *GetConversationsRequest) (*GetConversationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConversations not implemented")
}
func (UnimplementedChatServiceServer) GetConversationByID(context.Context, *GetConversationByIDRequest) (*GetConversationByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConversationByID not implemented")
}
func (UnimplementedChatServiceServer) GetHistoryMessages(context.Context, *GetHistoryMessagesRequest) (*GetHistoryMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHistoryMessages not implemented")
}
func (UnimplementedChatServiceServer) SetVariables(context.Context, *SetVariablesRequest) (*ChatVariablesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetVariables not implemented")
}
func (UnimplementedChatServiceServer) BlindTransfer(context.Context, *ChatTransferRequest) (*ChatTransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlindTransfer not implemented")
}
func (UnimplementedChatServiceServer) mustEmbedUnimplementedChatServiceServer() {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/webitel.chat.server.ChatService/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).SendMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_StartConversation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartConversationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).StartConversation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/webitel.chat.server.ChatService/StartConversation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).StartConversation(ctx, req.(*StartConversationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_CloseConversation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseConversationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).CloseConversation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/webitel.chat.server.ChatService/CloseConversation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).CloseConversation(ctx, req.(*CloseConversationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_JoinConversation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinConversationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).JoinConversation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/webitel.chat.server.ChatService/JoinConversation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).JoinConversation(ctx, req.(*JoinConversationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_LeaveConversation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveConversationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).LeaveConversation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/webitel.chat.server.ChatService/LeaveConversation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).LeaveConversation(ctx, req.(*LeaveConversationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_InviteToConversation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InviteToConversationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).InviteToConversation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/webitel.chat.server.ChatService/InviteToConversation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).InviteToConversation(ctx, req.(*InviteToConversationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_DeclineInvitation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeclineInvitationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).DeclineInvitation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/webitel.chat.server.ChatService/DeclineInvitation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).DeclineInvitation(ctx, req.(*DeclineInvitationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_CheckSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).CheckSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/webitel.chat.server.ChatService/CheckSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).CheckSession(ctx, req.(*CheckSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_WaitMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WaitMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).WaitMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/webitel.chat.server.ChatService/WaitMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).WaitMessage(ctx, req.(*WaitMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_UpdateChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).UpdateChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/webitel.chat.server.ChatService/UpdateChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).UpdateChannel(ctx, req.(*UpdateChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetConversations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConversationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetConversations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/webitel.chat.server.ChatService/GetConversations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetConversations(ctx, req.(*GetConversationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetConversationByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConversationByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetConversationByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/webitel.chat.server.ChatService/GetConversationByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetConversationByID(ctx, req.(*GetConversationByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetHistoryMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHistoryMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetHistoryMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/webitel.chat.server.ChatService/GetHistoryMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetHistoryMessages(ctx, req.(*GetHistoryMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_SetVariables_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetVariablesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).SetVariables(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/webitel.chat.server.ChatService/SetVariables",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).SetVariables(ctx, req.(*SetVariablesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_BlindTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).BlindTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/webitel.chat.server.ChatService/BlindTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).BlindTransfer(ctx, req.(*ChatTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "webitel.chat.server.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _ChatService_SendMessage_Handler,
		},
		{
			MethodName: "StartConversation",
			Handler:    _ChatService_StartConversation_Handler,
		},
		{
			MethodName: "CloseConversation",
			Handler:    _ChatService_CloseConversation_Handler,
		},
		{
			MethodName: "JoinConversation",
			Handler:    _ChatService_JoinConversation_Handler,
		},
		{
			MethodName: "LeaveConversation",
			Handler:    _ChatService_LeaveConversation_Handler,
		},
		{
			MethodName: "InviteToConversation",
			Handler:    _ChatService_InviteToConversation_Handler,
		},
		{
			MethodName: "DeclineInvitation",
			Handler:    _ChatService_DeclineInvitation_Handler,
		},
		{
			MethodName: "CheckSession",
			Handler:    _ChatService_CheckSession_Handler,
		},
		{
			MethodName: "WaitMessage",
			Handler:    _ChatService_WaitMessage_Handler,
		},
		{
			MethodName: "UpdateChannel",
			Handler:    _ChatService_UpdateChannel_Handler,
		},
		{
			MethodName: "GetConversations",
			Handler:    _ChatService_GetConversations_Handler,
		},
		{
			MethodName: "GetConversationByID",
			Handler:    _ChatService_GetConversationByID_Handler,
		},
		{
			MethodName: "GetHistoryMessages",
			Handler:    _ChatService_GetHistoryMessages_Handler,
		},
		{
			MethodName: "SetVariables",
			Handler:    _ChatService_SetVariables_Handler,
		},
		{
			MethodName: "BlindTransfer",
			Handler:    _ChatService_BlindTransfer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat.proto",
}
