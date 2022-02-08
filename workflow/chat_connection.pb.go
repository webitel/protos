// Code generated by protoc-gen-go. DO NOT EDIT.
// source: workflow/chat_connection.proto

package workflow

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	chat "github.com/webitel/protos/engine/chat"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type BreakBridgeRequest struct {
	ConversationId       string   `protobuf:"bytes,1,opt,name=conversation_id,json=conversationId,proto3" json:"conversation_id,omitempty"`
	Cause                string   `protobuf:"bytes,2,opt,name=cause,proto3" json:"cause,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BreakBridgeRequest) Reset()         { *m = BreakBridgeRequest{} }
func (m *BreakBridgeRequest) String() string { return proto.CompactTextString(m) }
func (*BreakBridgeRequest) ProtoMessage()    {}
func (*BreakBridgeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5783028bbcf301a4, []int{0}
}

func (m *BreakBridgeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BreakBridgeRequest.Unmarshal(m, b)
}
func (m *BreakBridgeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BreakBridgeRequest.Marshal(b, m, deterministic)
}
func (m *BreakBridgeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BreakBridgeRequest.Merge(m, src)
}
func (m *BreakBridgeRequest) XXX_Size() int {
	return xxx_messageInfo_BreakBridgeRequest.Size(m)
}
func (m *BreakBridgeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BreakBridgeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BreakBridgeRequest proto.InternalMessageInfo

func (m *BreakBridgeRequest) GetConversationId() string {
	if m != nil {
		return m.ConversationId
	}
	return ""
}

func (m *BreakBridgeRequest) GetCause() string {
	if m != nil {
		return m.Cause
	}
	return ""
}

type BreakBridgeResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BreakBridgeResponse) Reset()         { *m = BreakBridgeResponse{} }
func (m *BreakBridgeResponse) String() string { return proto.CompactTextString(m) }
func (*BreakBridgeResponse) ProtoMessage()    {}
func (*BreakBridgeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5783028bbcf301a4, []int{1}
}

func (m *BreakBridgeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BreakBridgeResponse.Unmarshal(m, b)
}
func (m *BreakBridgeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BreakBridgeResponse.Marshal(b, m, deterministic)
}
func (m *BreakBridgeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BreakBridgeResponse.Merge(m, src)
}
func (m *BreakBridgeResponse) XXX_Size() int {
	return xxx_messageInfo_BreakBridgeResponse.Size(m)
}
func (m *BreakBridgeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BreakBridgeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BreakBridgeResponse proto.InternalMessageInfo

type TransferChatPlanRequest struct {
	ConversationId       string   `protobuf:"bytes,1,opt,name=conversation_id,json=conversationId,proto3" json:"conversation_id,omitempty"`
	PlanId               int32    `protobuf:"varint,2,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransferChatPlanRequest) Reset()         { *m = TransferChatPlanRequest{} }
func (m *TransferChatPlanRequest) String() string { return proto.CompactTextString(m) }
func (*TransferChatPlanRequest) ProtoMessage()    {}
func (*TransferChatPlanRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5783028bbcf301a4, []int{2}
}

func (m *TransferChatPlanRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferChatPlanRequest.Unmarshal(m, b)
}
func (m *TransferChatPlanRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferChatPlanRequest.Marshal(b, m, deterministic)
}
func (m *TransferChatPlanRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferChatPlanRequest.Merge(m, src)
}
func (m *TransferChatPlanRequest) XXX_Size() int {
	return xxx_messageInfo_TransferChatPlanRequest.Size(m)
}
func (m *TransferChatPlanRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferChatPlanRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransferChatPlanRequest proto.InternalMessageInfo

func (m *TransferChatPlanRequest) GetConversationId() string {
	if m != nil {
		return m.ConversationId
	}
	return ""
}

func (m *TransferChatPlanRequest) GetPlanId() int32 {
	if m != nil {
		return m.PlanId
	}
	return 0
}

type TransferChatPlanResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransferChatPlanResponse) Reset()         { *m = TransferChatPlanResponse{} }
func (m *TransferChatPlanResponse) String() string { return proto.CompactTextString(m) }
func (*TransferChatPlanResponse) ProtoMessage()    {}
func (*TransferChatPlanResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5783028bbcf301a4, []int{3}
}

func (m *TransferChatPlanResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferChatPlanResponse.Unmarshal(m, b)
}
func (m *TransferChatPlanResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferChatPlanResponse.Marshal(b, m, deterministic)
}
func (m *TransferChatPlanResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferChatPlanResponse.Merge(m, src)
}
func (m *TransferChatPlanResponse) XXX_Size() int {
	return xxx_messageInfo_TransferChatPlanResponse.Size(m)
}
func (m *TransferChatPlanResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferChatPlanResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransferChatPlanResponse proto.InternalMessageInfo

type StartRequest struct {
	ConversationId       string            `protobuf:"bytes,1,opt,name=conversation_id,json=conversationId,proto3" json:"conversation_id,omitempty"`
	ProfileId            int64             `protobuf:"varint,2,opt,name=profile_id,json=profileId,proto3" json:"profile_id,omitempty"`
	DomainId             int64             `protobuf:"varint,3,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	Message              *chat.Message     `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	Variables            map[string]string `protobuf:"bytes,5,rep,name=variables,proto3" json:"variables,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	SchemaId             int32             `protobuf:"varint,6,opt,name=schema_id,json=schemaId,proto3" json:"schema_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *StartRequest) Reset()         { *m = StartRequest{} }
func (m *StartRequest) String() string { return proto.CompactTextString(m) }
func (*StartRequest) ProtoMessage()    {}
func (*StartRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5783028bbcf301a4, []int{4}
}

func (m *StartRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartRequest.Unmarshal(m, b)
}
func (m *StartRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartRequest.Marshal(b, m, deterministic)
}
func (m *StartRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartRequest.Merge(m, src)
}
func (m *StartRequest) XXX_Size() int {
	return xxx_messageInfo_StartRequest.Size(m)
}
func (m *StartRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StartRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StartRequest proto.InternalMessageInfo

func (m *StartRequest) GetConversationId() string {
	if m != nil {
		return m.ConversationId
	}
	return ""
}

func (m *StartRequest) GetProfileId() int64 {
	if m != nil {
		return m.ProfileId
	}
	return 0
}

func (m *StartRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

func (m *StartRequest) GetMessage() *chat.Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *StartRequest) GetVariables() map[string]string {
	if m != nil {
		return m.Variables
	}
	return nil
}

func (m *StartRequest) GetSchemaId() int32 {
	if m != nil {
		return m.SchemaId
	}
	return 0
}

type StartResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartResponse) Reset()         { *m = StartResponse{} }
func (m *StartResponse) String() string { return proto.CompactTextString(m) }
func (*StartResponse) ProtoMessage()    {}
func (*StartResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5783028bbcf301a4, []int{5}
}

func (m *StartResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartResponse.Unmarshal(m, b)
}
func (m *StartResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartResponse.Marshal(b, m, deterministic)
}
func (m *StartResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartResponse.Merge(m, src)
}
func (m *StartResponse) XXX_Size() int {
	return xxx_messageInfo_StartResponse.Size(m)
}
func (m *StartResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StartResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StartResponse proto.InternalMessageInfo

type BreakRequest struct {
	ConversationId       string   `protobuf:"bytes,1,opt,name=conversation_id,json=conversationId,proto3" json:"conversation_id,omitempty"`
	Cause                string   `protobuf:"bytes,2,opt,name=cause,proto3" json:"cause,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BreakRequest) Reset()         { *m = BreakRequest{} }
func (m *BreakRequest) String() string { return proto.CompactTextString(m) }
func (*BreakRequest) ProtoMessage()    {}
func (*BreakRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5783028bbcf301a4, []int{6}
}

func (m *BreakRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BreakRequest.Unmarshal(m, b)
}
func (m *BreakRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BreakRequest.Marshal(b, m, deterministic)
}
func (m *BreakRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BreakRequest.Merge(m, src)
}
func (m *BreakRequest) XXX_Size() int {
	return xxx_messageInfo_BreakRequest.Size(m)
}
func (m *BreakRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BreakRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BreakRequest proto.InternalMessageInfo

func (m *BreakRequest) GetConversationId() string {
	if m != nil {
		return m.ConversationId
	}
	return ""
}

func (m *BreakRequest) GetCause() string {
	if m != nil {
		return m.Cause
	}
	return ""
}

type BreakResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BreakResponse) Reset()         { *m = BreakResponse{} }
func (m *BreakResponse) String() string { return proto.CompactTextString(m) }
func (*BreakResponse) ProtoMessage()    {}
func (*BreakResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5783028bbcf301a4, []int{7}
}

func (m *BreakResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BreakResponse.Unmarshal(m, b)
}
func (m *BreakResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BreakResponse.Marshal(b, m, deterministic)
}
func (m *BreakResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BreakResponse.Merge(m, src)
}
func (m *BreakResponse) XXX_Size() int {
	return xxx_messageInfo_BreakResponse.Size(m)
}
func (m *BreakResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BreakResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BreakResponse proto.InternalMessageInfo

type ConfirmationMessageRequest struct {
	ConversationId       string          `protobuf:"bytes,1,opt,name=conversation_id,json=conversationId,proto3" json:"conversation_id,omitempty"`
	ConfirmationId       string          `protobuf:"bytes,2,opt,name=confirmation_id,json=confirmationId,proto3" json:"confirmation_id,omitempty"`
	Messages             []*chat.Message `protobuf:"bytes,3,rep,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ConfirmationMessageRequest) Reset()         { *m = ConfirmationMessageRequest{} }
func (m *ConfirmationMessageRequest) String() string { return proto.CompactTextString(m) }
func (*ConfirmationMessageRequest) ProtoMessage()    {}
func (*ConfirmationMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5783028bbcf301a4, []int{8}
}

func (m *ConfirmationMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfirmationMessageRequest.Unmarshal(m, b)
}
func (m *ConfirmationMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfirmationMessageRequest.Marshal(b, m, deterministic)
}
func (m *ConfirmationMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfirmationMessageRequest.Merge(m, src)
}
func (m *ConfirmationMessageRequest) XXX_Size() int {
	return xxx_messageInfo_ConfirmationMessageRequest.Size(m)
}
func (m *ConfirmationMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfirmationMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfirmationMessageRequest proto.InternalMessageInfo

func (m *ConfirmationMessageRequest) GetConversationId() string {
	if m != nil {
		return m.ConversationId
	}
	return ""
}

func (m *ConfirmationMessageRequest) GetConfirmationId() string {
	if m != nil {
		return m.ConfirmationId
	}
	return ""
}

func (m *ConfirmationMessageRequest) GetMessages() []*chat.Message {
	if m != nil {
		return m.Messages
	}
	return nil
}

type ConfirmationMessageResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfirmationMessageResponse) Reset()         { *m = ConfirmationMessageResponse{} }
func (m *ConfirmationMessageResponse) String() string { return proto.CompactTextString(m) }
func (*ConfirmationMessageResponse) ProtoMessage()    {}
func (*ConfirmationMessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5783028bbcf301a4, []int{9}
}

func (m *ConfirmationMessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfirmationMessageResponse.Unmarshal(m, b)
}
func (m *ConfirmationMessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfirmationMessageResponse.Marshal(b, m, deterministic)
}
func (m *ConfirmationMessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfirmationMessageResponse.Merge(m, src)
}
func (m *ConfirmationMessageResponse) XXX_Size() int {
	return xxx_messageInfo_ConfirmationMessageResponse.Size(m)
}
func (m *ConfirmationMessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfirmationMessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConfirmationMessageResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*BreakBridgeRequest)(nil), "workflow.BreakBridgeRequest")
	proto.RegisterType((*BreakBridgeResponse)(nil), "workflow.BreakBridgeResponse")
	proto.RegisterType((*TransferChatPlanRequest)(nil), "workflow.TransferChatPlanRequest")
	proto.RegisterType((*TransferChatPlanResponse)(nil), "workflow.TransferChatPlanResponse")
	proto.RegisterType((*StartRequest)(nil), "workflow.StartRequest")
	proto.RegisterMapType((map[string]string)(nil), "workflow.StartRequest.VariablesEntry")
	proto.RegisterType((*StartResponse)(nil), "workflow.StartResponse")
	proto.RegisterType((*BreakRequest)(nil), "workflow.BreakRequest")
	proto.RegisterType((*BreakResponse)(nil), "workflow.BreakResponse")
	proto.RegisterType((*ConfirmationMessageRequest)(nil), "workflow.ConfirmationMessageRequest")
	proto.RegisterType((*ConfirmationMessageResponse)(nil), "workflow.ConfirmationMessageResponse")
}

func init() { proto.RegisterFile("workflow/chat_connection.proto", fileDescriptor_5783028bbcf301a4) }

var fileDescriptor_5783028bbcf301a4 = []byte{
	// 550 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xc1, 0x6e, 0xda, 0x40,
	0x10, 0x0d, 0xb8, 0x10, 0x18, 0x42, 0x88, 0x36, 0x4d, 0xb1, 0x9c, 0x52, 0x51, 0x2b, 0x51, 0xe8,
	0xc5, 0x48, 0xf4, 0x52, 0x45, 0x3d, 0x81, 0x5a, 0xc9, 0x52, 0x23, 0x55, 0xa6, 0xea, 0xa1, 0x39,
	0x44, 0x8b, 0xbd, 0x80, 0x85, 0xed, 0xa5, 0xbb, 0x06, 0x94, 0x4f, 0xe9, 0xad, 0x1f, 0xd2, 0x8f,
	0xab, 0x76, 0xbd, 0x36, 0x86, 0x40, 0x2b, 0xa4, 0x5c, 0x10, 0xfb, 0x66, 0xde, 0x9b, 0xf1, 0x7b,
	0x5e, 0xc3, 0x9b, 0x15, 0x65, 0xb3, 0x71, 0x40, 0x57, 0x5d, 0x77, 0x8a, 0xe3, 0x07, 0x97, 0x46,
	0x11, 0x71, 0x63, 0x9f, 0x46, 0xd6, 0x9c, 0xd1, 0x98, 0xa2, 0x4a, 0x5a, 0x37, 0xea, 0x21, 0xe1,
	0x1c, 0x4f, 0x48, 0x52, 0x30, 0x87, 0x80, 0xfa, 0x8c, 0xe0, 0x59, 0x9f, 0xf9, 0xde, 0x84, 0x38,
	0xe4, 0xe7, 0x82, 0xf0, 0x18, 0xdd, 0x40, 0xc3, 0xa5, 0xd1, 0x92, 0x30, 0x8e, 0x85, 0xc8, 0x83,
	0xef, 0xe9, 0x85, 0x76, 0xa1, 0x53, 0x75, 0x4e, 0xf3, 0xb0, 0xed, 0xa1, 0x97, 0x50, 0x72, 0xf1,
	0x82, 0x13, 0xbd, 0x28, 0xcb, 0xc9, 0xc1, 0xbc, 0x80, 0xf3, 0x0d, 0x51, 0x3e, 0xa7, 0x11, 0x27,
	0xe6, 0x3d, 0x34, 0xbf, 0x31, 0x1c, 0xf1, 0x31, 0x61, 0x83, 0x29, 0x8e, 0xbf, 0x06, 0x38, 0x3a,
	0x78, 0x60, 0x13, 0x8e, 0xe7, 0x01, 0x96, 0x0d, 0x62, 0x64, 0xc9, 0x29, 0x8b, 0xa3, 0xed, 0x99,
	0x06, 0xe8, 0x4f, 0xc5, 0xd5, 0xe0, 0x3f, 0x45, 0x38, 0x19, 0xc6, 0x98, 0xc5, 0x07, 0x8f, 0x6b,
	0x01, 0xcc, 0x19, 0x1d, 0xfb, 0x01, 0x49, 0x27, 0x6a, 0x4e, 0x55, 0x21, 0xb6, 0x87, 0x2e, 0xa1,
	0xea, 0xd1, 0x10, 0xfb, 0x52, 0x41, 0x93, 0xd5, 0x4a, 0x02, 0xd8, 0x1e, 0xba, 0x81, 0x63, 0xe5,
	0xb5, 0xfe, 0xa2, 0x5d, 0xe8, 0xd4, 0x7a, 0x75, 0x4b, 0x84, 0x63, 0xdd, 0x25, 0xa0, 0x93, 0x56,
	0xd1, 0x00, 0xaa, 0x4b, 0xcc, 0x7c, 0x3c, 0x0a, 0x08, 0xd7, 0x4b, 0x6d, 0xad, 0x53, 0xeb, 0x5d,
	0x5b, 0x69, 0x60, 0x56, 0x7e, 0x71, 0xeb, 0x7b, 0xda, 0xf7, 0x29, 0x8a, 0xd9, 0xa3, 0xb3, 0xe6,
	0x89, 0x55, 0xb8, 0x3b, 0x25, 0x21, 0x16, 0xab, 0x94, 0xa5, 0x35, 0x95, 0x04, 0xb0, 0x3d, 0xe3,
	0x23, 0x9c, 0x6e, 0x32, 0xd1, 0x19, 0x68, 0x33, 0xf2, 0xa8, 0x9e, 0x5a, 0xfc, 0x15, 0x51, 0x2e,
	0x71, 0xb0, 0xc8, 0xa2, 0x94, 0x87, 0xdb, 0xe2, 0x87, 0x82, 0xd9, 0x80, 0xba, 0x5a, 0x42, 0xf9,
	0x79, 0x07, 0x27, 0x32, 0xdf, 0x67, 0x7a, 0x5d, 0x1a, 0x50, 0x57, 0x72, 0x4a, 0xff, 0x57, 0x01,
	0x8c, 0x01, 0x8d, 0xc6, 0x3e, 0x0b, 0x25, 0x33, 0x75, 0xec, 0xd0, 0x71, 0x49, 0x63, 0x26, 0x93,
	0x46, 0x98, 0x34, 0x66, 0xb0, 0xed, 0xa1, 0x77, 0x50, 0x51, 0x61, 0x70, 0x5d, 0x93, 0x01, 0x6c,
	0x65, 0x95, 0x95, 0xcd, 0x16, 0x5c, 0xee, 0x5c, 0x2d, 0x59, 0xbd, 0xf7, 0x5b, 0x83, 0x8b, 0xcf,
	0x01, 0x5d, 0x89, 0x77, 0x70, 0x48, 0xd8, 0x92, 0x30, 0xf1, 0xeb, 0xbb, 0x04, 0xdd, 0x42, 0x49,
	0xba, 0x88, 0x5e, 0xed, 0xce, 0xd6, 0x68, 0x3e, 0xc1, 0x95, 0x1d, 0x47, 0x82, 0x2b, 0x1d, 0xca,
	0x73, 0xf3, 0x09, 0xe4, 0xb9, 0x9b, 0x56, 0x1e, 0xa1, 0x2f, 0x50, 0xcb, 0x5d, 0x46, 0xf4, 0x7a,
	0xab, 0x73, 0xe3, 0xe2, 0x1b, 0xad, 0x3d, 0xd5, 0x4c, 0xcd, 0x83, 0xf3, 0x1d, 0x8f, 0x8f, 0xae,
	0xd6, 0xbc, 0xfd, 0xc1, 0x19, 0xd7, 0xff, 0xe9, 0xca, 0xa6, 0xdc, 0xc3, 0xd9, 0xf6, 0x65, 0x46,
	0x6f, 0xd7, 0xe4, 0x3d, 0x5f, 0x11, 0xc3, 0xfc, 0x57, 0x4b, 0x2a, 0xde, 0xbf, 0xfa, 0x61, 0x4e,
	0xfc, 0x78, 0xba, 0x18, 0x59, 0x2e, 0x0d, 0xbb, 0x2b, 0x32, 0xf2, 0x63, 0x12, 0x74, 0xe5, 0xe7,
	0x90, 0x77, 0x53, 0x81, 0x51, 0x59, 0x02, 0xef, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x85, 0x17,
	0xb6, 0xc1, 0x5a, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FlowChatServerServiceClient is the client API for FlowChatServerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FlowChatServerServiceClient interface {
	Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error)
	Break(ctx context.Context, in *BreakRequest, opts ...grpc.CallOption) (*BreakResponse, error)
	BreakBridge(ctx context.Context, in *BreakBridgeRequest, opts ...grpc.CallOption) (*BreakBridgeResponse, error)
	ConfirmationMessage(ctx context.Context, in *ConfirmationMessageRequest, opts ...grpc.CallOption) (*ConfirmationMessageResponse, error)
	TransferChatPlan(ctx context.Context, in *TransferChatPlanRequest, opts ...grpc.CallOption) (*TransferChatPlanResponse, error)
}

type flowChatServerServiceClient struct {
	cc *grpc.ClientConn
}

func NewFlowChatServerServiceClient(cc *grpc.ClientConn) FlowChatServerServiceClient {
	return &flowChatServerServiceClient{cc}
}

func (c *flowChatServerServiceClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error) {
	out := new(StartResponse)
	err := c.cc.Invoke(ctx, "/workflow.FlowChatServerService/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowChatServerServiceClient) Break(ctx context.Context, in *BreakRequest, opts ...grpc.CallOption) (*BreakResponse, error) {
	out := new(BreakResponse)
	err := c.cc.Invoke(ctx, "/workflow.FlowChatServerService/Break", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowChatServerServiceClient) BreakBridge(ctx context.Context, in *BreakBridgeRequest, opts ...grpc.CallOption) (*BreakBridgeResponse, error) {
	out := new(BreakBridgeResponse)
	err := c.cc.Invoke(ctx, "/workflow.FlowChatServerService/BreakBridge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowChatServerServiceClient) ConfirmationMessage(ctx context.Context, in *ConfirmationMessageRequest, opts ...grpc.CallOption) (*ConfirmationMessageResponse, error) {
	out := new(ConfirmationMessageResponse)
	err := c.cc.Invoke(ctx, "/workflow.FlowChatServerService/ConfirmationMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowChatServerServiceClient) TransferChatPlan(ctx context.Context, in *TransferChatPlanRequest, opts ...grpc.CallOption) (*TransferChatPlanResponse, error) {
	out := new(TransferChatPlanResponse)
	err := c.cc.Invoke(ctx, "/workflow.FlowChatServerService/TransferChatPlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FlowChatServerServiceServer is the server API for FlowChatServerService service.
type FlowChatServerServiceServer interface {
	Start(context.Context, *StartRequest) (*StartResponse, error)
	Break(context.Context, *BreakRequest) (*BreakResponse, error)
	BreakBridge(context.Context, *BreakBridgeRequest) (*BreakBridgeResponse, error)
	ConfirmationMessage(context.Context, *ConfirmationMessageRequest) (*ConfirmationMessageResponse, error)
	TransferChatPlan(context.Context, *TransferChatPlanRequest) (*TransferChatPlanResponse, error)
}

// UnimplementedFlowChatServerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFlowChatServerServiceServer struct {
}

func (*UnimplementedFlowChatServerServiceServer) Start(ctx context.Context, req *StartRequest) (*StartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (*UnimplementedFlowChatServerServiceServer) Break(ctx context.Context, req *BreakRequest) (*BreakResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Break not implemented")
}
func (*UnimplementedFlowChatServerServiceServer) BreakBridge(ctx context.Context, req *BreakBridgeRequest) (*BreakBridgeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BreakBridge not implemented")
}
func (*UnimplementedFlowChatServerServiceServer) ConfirmationMessage(ctx context.Context, req *ConfirmationMessageRequest) (*ConfirmationMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmationMessage not implemented")
}
func (*UnimplementedFlowChatServerServiceServer) TransferChatPlan(ctx context.Context, req *TransferChatPlanRequest) (*TransferChatPlanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferChatPlan not implemented")
}

func RegisterFlowChatServerServiceServer(s *grpc.Server, srv FlowChatServerServiceServer) {
	s.RegisterService(&_FlowChatServerService_serviceDesc, srv)
}

func _FlowChatServerService_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlowChatServerServiceServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.FlowChatServerService/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlowChatServerServiceServer).Start(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlowChatServerService_Break_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BreakRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlowChatServerServiceServer).Break(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.FlowChatServerService/Break",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlowChatServerServiceServer).Break(ctx, req.(*BreakRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlowChatServerService_BreakBridge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BreakBridgeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlowChatServerServiceServer).BreakBridge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.FlowChatServerService/BreakBridge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlowChatServerServiceServer).BreakBridge(ctx, req.(*BreakBridgeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlowChatServerService_ConfirmationMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmationMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlowChatServerServiceServer).ConfirmationMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.FlowChatServerService/ConfirmationMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlowChatServerServiceServer).ConfirmationMessage(ctx, req.(*ConfirmationMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlowChatServerService_TransferChatPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferChatPlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlowChatServerServiceServer).TransferChatPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/workflow.FlowChatServerService/TransferChatPlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlowChatServerServiceServer).TransferChatPlan(ctx, req.(*TransferChatPlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FlowChatServerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "workflow.FlowChatServerService",
	HandlerType: (*FlowChatServerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _FlowChatServerService_Start_Handler,
		},
		{
			MethodName: "Break",
			Handler:    _FlowChatServerService_Break_Handler,
		},
		{
			MethodName: "BreakBridge",
			Handler:    _FlowChatServerService_BreakBridge_Handler,
		},
		{
			MethodName: "ConfirmationMessage",
			Handler:    _FlowChatServerService_ConfirmationMessage_Handler,
		},
		{
			MethodName: "TransferChatPlan",
			Handler:    _FlowChatServerService_TransferChatPlan_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "workflow/chat_connection.proto",
}
