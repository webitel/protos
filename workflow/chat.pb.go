// Code generated by protoc-gen-go. DO NOT EDIT.
// source: workflow/chat.proto

package workflow

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
	return fileDescriptor_7b58d5ca10fd60ce, []int{0}
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
	Error                *Error   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BreakBridgeResponse) Reset()         { *m = BreakBridgeResponse{} }
func (m *BreakBridgeResponse) String() string { return proto.CompactTextString(m) }
func (*BreakBridgeResponse) ProtoMessage()    {}
func (*BreakBridgeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b58d5ca10fd60ce, []int{1}
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

func (m *BreakBridgeResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type Message struct {
	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*Message_Text
	//	*Message_File_
	Value                isMessage_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b58d5ca10fd60ce, []int{2}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Message) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type isMessage_Value interface {
	isMessage_Value()
}

type Message_Text struct {
	Text string `protobuf:"bytes,3,opt,name=text,proto3,oneof"`
}

type Message_File_ struct {
	File *Message_File `protobuf:"bytes,4,opt,name=file,proto3,oneof"`
}

func (*Message_Text) isMessage_Value() {}

func (*Message_File_) isMessage_Value() {}

func (m *Message) GetValue() isMessage_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Message) GetText() string {
	if x, ok := m.GetValue().(*Message_Text); ok {
		return x.Text
	}
	return ""
}

func (m *Message) GetFile() *Message_File {
	if x, ok := m.GetValue().(*Message_File_); ok {
		return x.File
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Message) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Message_Text)(nil),
		(*Message_File_)(nil),
	}
}

type Message_File struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	MimeType             string   `protobuf:"bytes,3,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message_File) Reset()         { *m = Message_File{} }
func (m *Message_File) String() string { return proto.CompactTextString(m) }
func (*Message_File) ProtoMessage()    {}
func (*Message_File) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b58d5ca10fd60ce, []int{2, 0}
}

func (m *Message_File) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message_File.Unmarshal(m, b)
}
func (m *Message_File) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message_File.Marshal(b, m, deterministic)
}
func (m *Message_File) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message_File.Merge(m, src)
}
func (m *Message_File) XXX_Size() int {
	return xxx_messageInfo_Message_File.Size(m)
}
func (m *Message_File) XXX_DiscardUnknown() {
	xxx_messageInfo_Message_File.DiscardUnknown(m)
}

var xxx_messageInfo_Message_File proto.InternalMessageInfo

func (m *Message_File) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Message_File) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Message_File) GetMimeType() string {
	if m != nil {
		return m.MimeType
	}
	return ""
}

type StartRequest struct {
	ConversationId       string            `protobuf:"bytes,1,opt,name=conversation_id,json=conversationId,proto3" json:"conversation_id,omitempty"`
	ProfileId            int64             `protobuf:"varint,2,opt,name=profile_id,json=profileId,proto3" json:"profile_id,omitempty"`
	DomainId             int64             `protobuf:"varint,3,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	Message              *Message          `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	Variables            map[string]string `protobuf:"bytes,5,rep,name=variables,proto3" json:"variables,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *StartRequest) Reset()         { *m = StartRequest{} }
func (m *StartRequest) String() string { return proto.CompactTextString(m) }
func (*StartRequest) ProtoMessage()    {}
func (*StartRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b58d5ca10fd60ce, []int{3}
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

func (m *StartRequest) GetMessage() *Message {
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

type Error struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b58d5ca10fd60ce, []int{4}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type StartResponse struct {
	Error                *Error   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartResponse) Reset()         { *m = StartResponse{} }
func (m *StartResponse) String() string { return proto.CompactTextString(m) }
func (*StartResponse) ProtoMessage()    {}
func (*StartResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b58d5ca10fd60ce, []int{5}
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

func (m *StartResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type BreakRequest struct {
	ConversationId       string   `protobuf:"bytes,1,opt,name=conversation_id,json=conversationId,proto3" json:"conversation_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BreakRequest) Reset()         { *m = BreakRequest{} }
func (m *BreakRequest) String() string { return proto.CompactTextString(m) }
func (*BreakRequest) ProtoMessage()    {}
func (*BreakRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b58d5ca10fd60ce, []int{6}
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

type BreakResponse struct {
	Error                *Error   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BreakResponse) Reset()         { *m = BreakResponse{} }
func (m *BreakResponse) String() string { return proto.CompactTextString(m) }
func (*BreakResponse) ProtoMessage()    {}
func (*BreakResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b58d5ca10fd60ce, []int{7}
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

func (m *BreakResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type ConfirmationMessageRequest struct {
	ConversationId       string     `protobuf:"bytes,1,opt,name=conversation_id,json=conversationId,proto3" json:"conversation_id,omitempty"`
	ConfirmationId       string     `protobuf:"bytes,2,opt,name=confirmation_id,json=confirmationId,proto3" json:"confirmation_id,omitempty"`
	Messages             []*Message `protobuf:"bytes,3,rep,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ConfirmationMessageRequest) Reset()         { *m = ConfirmationMessageRequest{} }
func (m *ConfirmationMessageRequest) String() string { return proto.CompactTextString(m) }
func (*ConfirmationMessageRequest) ProtoMessage()    {}
func (*ConfirmationMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b58d5ca10fd60ce, []int{8}
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

func (m *ConfirmationMessageRequest) GetMessages() []*Message {
	if m != nil {
		return m.Messages
	}
	return nil
}

type ConfirmationMessageResponse struct {
	Error                *Error   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfirmationMessageResponse) Reset()         { *m = ConfirmationMessageResponse{} }
func (m *ConfirmationMessageResponse) String() string { return proto.CompactTextString(m) }
func (*ConfirmationMessageResponse) ProtoMessage()    {}
func (*ConfirmationMessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b58d5ca10fd60ce, []int{9}
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

func (m *ConfirmationMessageResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*BreakBridgeRequest)(nil), "workflow.BreakBridgeRequest")
	proto.RegisterType((*BreakBridgeResponse)(nil), "workflow.BreakBridgeResponse")
	proto.RegisterType((*Message)(nil), "workflow.Message")
	proto.RegisterType((*Message_File)(nil), "workflow.Message.File")
	proto.RegisterType((*StartRequest)(nil), "workflow.StartRequest")
	proto.RegisterMapType((map[string]string)(nil), "workflow.StartRequest.VariablesEntry")
	proto.RegisterType((*Error)(nil), "workflow.Error")
	proto.RegisterType((*StartResponse)(nil), "workflow.StartResponse")
	proto.RegisterType((*BreakRequest)(nil), "workflow.BreakRequest")
	proto.RegisterType((*BreakResponse)(nil), "workflow.BreakResponse")
	proto.RegisterType((*ConfirmationMessageRequest)(nil), "workflow.ConfirmationMessageRequest")
	proto.RegisterType((*ConfirmationMessageResponse)(nil), "workflow.ConfirmationMessageResponse")
}

func init() { proto.RegisterFile("workflow/chat.proto", fileDescriptor_7b58d5ca10fd60ce) }

var fileDescriptor_7b58d5ca10fd60ce = []byte{
	// 591 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x5f, 0x6f, 0xd2, 0x50,
	0x14, 0x1f, 0x2d, 0x75, 0x70, 0x98, 0x4c, 0x2f, 0x73, 0x36, 0x9d, 0x4b, 0x96, 0x66, 0xc4, 0x25,
	0x6a, 0x89, 0x98, 0xa8, 0x59, 0xf6, 0x04, 0xb2, 0x48, 0xa2, 0x2f, 0xc5, 0xf8, 0xe0, 0xcb, 0x52,
	0xe8, 0x01, 0x6e, 0x68, 0xb9, 0x78, 0x7b, 0x01, 0xf9, 0x30, 0x3e, 0xf9, 0x49, 0x8c, 0x5f, 0xcc,
	0xdc, 0xdb, 0x16, 0x8a, 0xb0, 0x98, 0xbe, 0x34, 0x3d, 0xbf, 0xf3, 0x3b, 0xe7, 0x77, 0xfe, 0xe5,
	0x42, 0x6d, 0xc9, 0xf8, 0x64, 0x18, 0xb0, 0x65, 0x63, 0x30, 0xf6, 0x84, 0x33, 0xe3, 0x4c, 0x30,
	0x52, 0x4a, 0x41, 0xbb, 0x07, 0xa4, 0xc5, 0xd1, 0x9b, 0xb4, 0x38, 0xf5, 0x47, 0xe8, 0xe2, 0xf7,
	0x39, 0x46, 0x82, 0x3c, 0x87, 0xe3, 0x01, 0x9b, 0x2e, 0x90, 0x47, 0x9e, 0xa0, 0x6c, 0x7a, 0x47,
	0x7d, 0xb3, 0x70, 0x51, 0xb8, 0x2a, 0xbb, 0xd5, 0x2c, 0xdc, 0xf5, 0xc9, 0x09, 0x18, 0x03, 0x6f,
	0x1e, 0xa1, 0xa9, 0x29, 0x77, 0x6c, 0xd8, 0x37, 0x50, 0xdb, 0x4a, 0x1a, 0xcd, 0xd8, 0x34, 0x42,
	0x52, 0x07, 0x03, 0x39, 0x67, 0x5c, 0xe5, 0xaa, 0x34, 0x8f, 0x9d, 0xb4, 0x0a, 0xa7, 0x23, 0x61,
	0x37, 0xf6, 0xda, 0xbf, 0x0b, 0x70, 0xf8, 0x19, 0xa3, 0xc8, 0x1b, 0x21, 0xa9, 0x82, 0x96, 0x68,
	0xeb, 0xae, 0x46, 0x7d, 0x42, 0xa0, 0x28, 0x56, 0xb3, 0x54, 0x4e, 0xfd, 0x93, 0x13, 0x28, 0x0a,
	0xfc, 0x21, 0x4c, 0x5d, 0x62, 0x1f, 0x0f, 0x5c, 0x65, 0x91, 0x97, 0x50, 0x1c, 0xd2, 0x00, 0xcd,
	0xa2, 0xd2, 0x3a, 0xdd, 0x68, 0x25, 0xa9, 0x9d, 0x5b, 0x1a, 0xa0, 0x64, 0x4b, 0x96, 0xd5, 0x81,
	0xa2, 0xb4, 0x77, 0xf4, 0x1e, 0x81, 0x3e, 0xe7, 0x41, 0x22, 0x27, 0x7f, 0xc9, 0x19, 0x94, 0x43,
	0x1a, 0xe2, 0x9d, 0x2a, 0x43, 0x49, 0xba, 0x25, 0x09, 0x7c, 0x59, 0xcd, 0xb0, 0x75, 0x08, 0xc6,
	0xc2, 0x0b, 0xe6, 0x68, 0xff, 0xd2, 0xe0, 0xa8, 0x27, 0x3c, 0x2e, 0x72, 0x4f, 0xf4, 0x1c, 0x60,
	0xc6, 0x99, 0x2c, 0x4a, 0x72, 0x34, 0x55, 0x49, 0x39, 0x41, 0xba, 0xbe, 0x94, 0xf7, 0x59, 0xe8,
	0x51, 0x95, 0x41, 0x57, 0xde, 0x52, 0x0c, 0x74, 0x7d, 0xf2, 0x02, 0x0e, 0xc3, 0xb8, 0xbb, 0xa4,
	0xed, 0xc7, 0x3b, 0x6d, 0xbb, 0x29, 0x83, 0xb4, 0xa1, 0xbc, 0xf0, 0x38, 0xf5, 0xfa, 0x01, 0x46,
	0xa6, 0x71, 0xa1, 0x5f, 0x55, 0x9a, 0xf5, 0x0d, 0x3d, 0x5b, 0xbc, 0xf3, 0x35, 0xe5, 0x75, 0xa6,
	0x82, 0xaf, 0xdc, 0x4d, 0x9c, 0x75, 0x03, 0xd5, 0x6d, 0xa7, 0x9c, 0xd8, 0x04, 0x57, 0x49, 0x73,
	0xf2, 0x57, 0xde, 0x88, 0x1a, 0x4a, 0x7a, 0x23, 0xca, 0xb8, 0xd6, 0xde, 0x17, 0xec, 0xd7, 0x60,
	0xa8, 0xcd, 0x67, 0xc6, 0x5e, 0x56, 0x63, 0x37, 0x37, 0x8d, 0xc4, 0x41, 0xa9, 0x69, 0xbf, 0x85,
	0x87, 0x49, 0x69, 0xf9, 0x8e, 0xea, 0x1d, 0x1c, 0xa9, 0x93, 0xcc, 0xbb, 0x0f, 0x29, 0x98, 0x04,
	0xe6, 0x13, 0xfc, 0x59, 0x00, 0xab, 0xcd, 0xa6, 0x43, 0xca, 0x43, 0x95, 0x2a, 0x9d, 0x7f, 0xde,
	0x7b, 0x88, 0x89, 0xeb, 0x34, 0xe9, 0x51, 0xc4, 0xc4, 0x35, 0xdc, 0xf5, 0xc9, 0x2b, 0x28, 0x25,
	0x43, 0x8a, 0x4c, 0x5d, 0xad, 0x73, 0xcf, 0xf6, 0xd7, 0x14, 0xfb, 0x03, 0x9c, 0xed, 0x2d, 0x2f,
	0x57, 0x97, 0xcd, 0x3f, 0x1a, 0x3c, 0xb9, 0x0d, 0xd8, 0xb2, 0x3d, 0xf6, 0x44, 0x0f, 0xf9, 0x02,
	0xb9, 0xfc, 0xd2, 0x01, 0x92, 0x6b, 0x30, 0xd4, 0xa2, 0xc8, 0xe9, 0xfe, 0xa3, 0xb2, 0x9e, 0xee,
	0xe0, 0xb1, 0xb4, 0x7d, 0x20, 0x63, 0xd5, 0xcc, 0xb3, 0xb1, 0xd9, 0xed, 0x65, 0x63, 0xb7, 0x96,
	0x63, 0x1f, 0x90, 0x4f, 0x50, 0xc9, 0xbc, 0x3d, 0xe4, 0xd9, 0x3f, 0xcc, 0xad, 0x77, 0xce, 0x3a,
	0xbf, 0xc7, 0xbb, 0xce, 0xe6, 0x43, 0x6d, 0xcf, 0x94, 0xc8, 0xe5, 0x26, 0xee, 0xfe, 0x1d, 0x5b,
	0xf5, 0xff, 0xb0, 0x52, 0x95, 0xd6, 0xe5, 0x37, 0x7b, 0x44, 0xc5, 0x78, 0xde, 0x77, 0x06, 0x2c,
	0x6c, 0x2c, 0xb1, 0x4f, 0x05, 0x06, 0x0d, 0xf5, 0x54, 0x47, 0x8d, 0x34, 0x47, 0xff, 0x81, 0x02,
	0xde, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xba, 0xf1, 0x0e, 0xb7, 0xd2, 0x05, 0x00, 0x00,
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

// FlowChatServerServiceServer is the server API for FlowChatServerService service.
type FlowChatServerServiceServer interface {
	Start(context.Context, *StartRequest) (*StartResponse, error)
	Break(context.Context, *BreakRequest) (*BreakResponse, error)
	BreakBridge(context.Context, *BreakBridgeRequest) (*BreakBridgeResponse, error)
	ConfirmationMessage(context.Context, *ConfirmationMessageRequest) (*ConfirmationMessageResponse, error)
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "workflow/chat.proto",
}
