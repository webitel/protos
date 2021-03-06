// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package chat

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Account contact info to extend and replace legacy chat.Client message type
type Account struct {
	// Unique IDentifier
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Channel communication type
	Channel string `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel,omitempty"`
	// Channel specific contact string
	Contact              string   `protobuf:"bytes,3,opt,name=contact,proto3" json:"contact,omitempty"`
	FirstName            string   `protobuf:"bytes,4,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,5,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Username             string   `protobuf:"bytes,6,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Account) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *Account) GetContact() string {
	if m != nil {
		return m.Contact
	}
	return ""
}

func (m *Account) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Account) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Account) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type Message struct {
	// Unique message identifier inside this chat
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Type of this Message to notify e.g.: text, file, read, status etc.
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// Text of the message to be sent, 1-4096
	Text string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	// File to send. Pass an HTTP .URL as a string for Webitel to get a file from the Internet,
	// or upload a new one using storage.FileService.UploadFile()
	File *File `protobuf:"bytes,4,opt,name=file,proto3" json:"file,omitempty"`
	// Optional. This Message extra properties
	Variables map[string]string `protobuf:"bytes,5,rep,name=variables,proto3" json:"variables,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// End-User extra contact info
	Contact *Account `protobuf:"bytes,7,opt,name=contact,proto3" json:"contact,omitempty"`
	// Optional. Send message date (epochtime ms) Generates by the service.
	CreatedAt int64 `protobuf:"varint,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Optional. Edit message date (epochtime ms) Generates by the service.
	// For "read" messages, you can specify the date the last READ message was created_at
	UpdatedAt int64 `protobuf:"varint,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// Optional. If the message is a reply,
	// this is the ID of the original message
	ReplyToMessageId int64 `protobuf:"varint,10,opt,name=reply_to_message_id,json=replyToMessageId,proto3" json:"reply_to_message_id,omitempty"`
	// External message sent-binding to be able to identify corresponding internal message
	ReplyToVariables map[string]string `protobuf:"bytes,11,rep,name=reply_to_variables,json=replyToVariables,proto3" json:"reply_to_variables,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Unique identifier for the chat where the original message was sent
	ForwardFromChatId string `protobuf:"bytes,12,opt,name=forward_from_chat_id,json=forwardFromChatId,proto3" json:"forward_from_chat_id,omitempty"`
	// Message identifier in the chat specified in from_chat_id
	ForwardFromMessageId int64 `protobuf:"varint,13,opt,name=forward_from_message_id,json=forwardFromMessageId,proto3" json:"forward_from_message_id,omitempty"`
	// External message sent-binding to be able to identify corresponding internal message
	ForwardFromVariables map[string]string `protobuf:"bytes,14,rep,name=forward_from_variables,json=forwardFromVariables,proto3" json:"forward_from_variables,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Reply Markup Buttons SET
	Buttons []*Buttons `protobuf:"bytes,6,rep,name=buttons,proto3" json:"buttons,omitempty"`
	// Inline Keyboard Buttons SET
	Inline               []*Buttons `protobuf:"bytes,15,rep,name=inline,proto3" json:"inline,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
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

func (m *Message) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Message) GetFile() *File {
	if m != nil {
		return m.File
	}
	return nil
}

func (m *Message) GetVariables() map[string]string {
	if m != nil {
		return m.Variables
	}
	return nil
}

func (m *Message) GetContact() *Account {
	if m != nil {
		return m.Contact
	}
	return nil
}

func (m *Message) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Message) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *Message) GetReplyToMessageId() int64 {
	if m != nil {
		return m.ReplyToMessageId
	}
	return 0
}

func (m *Message) GetReplyToVariables() map[string]string {
	if m != nil {
		return m.ReplyToVariables
	}
	return nil
}

func (m *Message) GetForwardFromChatId() string {
	if m != nil {
		return m.ForwardFromChatId
	}
	return ""
}

func (m *Message) GetForwardFromMessageId() int64 {
	if m != nil {
		return m.ForwardFromMessageId
	}
	return 0
}

func (m *Message) GetForwardFromVariables() map[string]string {
	if m != nil {
		return m.ForwardFromVariables
	}
	return nil
}

func (m *Message) GetButtons() []*Buttons {
	if m != nil {
		return m.Buttons
	}
	return nil
}

func (m *Message) GetInline() []*Buttons {
	if m != nil {
		return m.Inline
	}
	return nil
}

type File struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Mime                 string   `protobuf:"bytes,3,opt,name=mime,proto3" json:"mime,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Size                 int64    `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *File) Reset()         { *m = File{} }
func (m *File) String() string { return proto.CompactTextString(m) }
func (*File) ProtoMessage()    {}
func (*File) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{2}
}

func (m *File) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_File.Unmarshal(m, b)
}
func (m *File) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_File.Marshal(b, m, deterministic)
}
func (m *File) XXX_Merge(src proto.Message) {
	xxx_messageInfo_File.Merge(m, src)
}
func (m *File) XXX_Size() int {
	return xxx_messageInfo_File.Size(m)
}
func (m *File) XXX_DiscardUnknown() {
	xxx_messageInfo_File.DiscardUnknown(m)
}

var xxx_messageInfo_File proto.InternalMessageInfo

func (m *File) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *File) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *File) GetMime() string {
	if m != nil {
		return m.Mime
	}
	return ""
}

func (m *File) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *File) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

type Buttons struct {
	Button               []*Button `protobuf:"bytes,1,rep,name=button,proto3" json:"button,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Buttons) Reset()         { *m = Buttons{} }
func (m *Buttons) String() string { return proto.CompactTextString(m) }
func (*Buttons) ProtoMessage()    {}
func (*Buttons) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{3}
}

func (m *Buttons) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Buttons.Unmarshal(m, b)
}
func (m *Buttons) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Buttons.Marshal(b, m, deterministic)
}
func (m *Buttons) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Buttons.Merge(m, src)
}
func (m *Buttons) XXX_Size() int {
	return xxx_messageInfo_Buttons.Size(m)
}
func (m *Buttons) XXX_DiscardUnknown() {
	xxx_messageInfo_Buttons.DiscardUnknown(m)
}

var xxx_messageInfo_Buttons proto.InternalMessageInfo

func (m *Buttons) GetButton() []*Button {
	if m != nil {
		return m.Button
	}
	return nil
}

type Button struct {
	Caption              string   `protobuf:"bytes,1,opt,name=caption,proto3" json:"caption,omitempty"`
	Text                 string   `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Url                  string   `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	Code                 string   `protobuf:"bytes,5,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Button) Reset()         { *m = Button{} }
func (m *Button) String() string { return proto.CompactTextString(m) }
func (*Button) ProtoMessage()    {}
func (*Button) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{4}
}

func (m *Button) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Button.Unmarshal(m, b)
}
func (m *Button) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Button.Marshal(b, m, deterministic)
}
func (m *Button) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Button.Merge(m, src)
}
func (m *Button) XXX_Size() int {
	return xxx_messageInfo_Button.Size(m)
}
func (m *Button) XXX_DiscardUnknown() {
	xxx_messageInfo_Button.DiscardUnknown(m)
}

var xxx_messageInfo_Button proto.InternalMessageInfo

func (m *Button) GetCaption() string {
	if m != nil {
		return m.Caption
	}
	return ""
}

func (m *Button) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Button) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Button) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Button) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func init() {
	proto.RegisterType((*Account)(nil), "chat.Account")
	proto.RegisterType((*Message)(nil), "chat.Message")
	proto.RegisterMapType((map[string]string)(nil), "chat.Message.ForwardFromVariablesEntry")
	proto.RegisterMapType((map[string]string)(nil), "chat.Message.ReplyToVariablesEntry")
	proto.RegisterMapType((map[string]string)(nil), "chat.Message.VariablesEntry")
	proto.RegisterType((*File)(nil), "chat.File")
	proto.RegisterType((*Buttons)(nil), "chat.Buttons")
	proto.RegisterType((*Button)(nil), "chat.Button")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 617 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x6e, 0xd4, 0x3c,
	0x10, 0xd5, 0x6e, 0xb6, 0xd9, 0x66, 0xfa, 0xf3, 0xf5, 0x33, 0x05, 0x42, 0xf9, 0x51, 0xb5, 0x50,
	0xb5, 0x37, 0xec, 0x4a, 0x45, 0x48, 0xa8, 0xe2, 0xa6, 0xad, 0x28, 0xea, 0x05, 0x48, 0x44, 0xa8,
	0x17, 0x48, 0x28, 0xf2, 0x26, 0xde, 0xae, 0x45, 0x12, 0x47, 0x8e, 0xd3, 0xb2, 0xbc, 0x0f, 0x0f,
	0xc7, 0x5b, 0x20, 0x8f, 0xed, 0x4d, 0xd2, 0x96, 0x8b, 0xde, 0x8d, 0xcf, 0xf1, 0x4c, 0xce, 0x1c,
	0x67, 0x06, 0x36, 0x72, 0x56, 0x55, 0xf4, 0x92, 0x8d, 0x4b, 0x29, 0x94, 0x20, 0x83, 0x64, 0x4e,
	0xd5, 0xe8, 0x77, 0x0f, 0x86, 0xc7, 0x49, 0x22, 0xea, 0x42, 0x91, 0x4d, 0xe8, 0xf3, 0x34, 0xec,
	0xed, 0xf6, 0x0e, 0xbc, 0xa8, 0xcf, 0x53, 0x12, 0xc2, 0x30, 0x99, 0xd3, 0xa2, 0x60, 0x59, 0xd8,
	0xdf, 0xed, 0x1d, 0x04, 0x91, 0x3b, 0x22, 0x23, 0x0a, 0x45, 0x13, 0x15, 0x7a, 0x96, 0x31, 0x47,
	0xf2, 0x1c, 0x60, 0xc6, 0x65, 0xa5, 0xe2, 0x82, 0xe6, 0x2c, 0x1c, 0x20, 0x19, 0x20, 0xf2, 0x99,
	0xe6, 0x8c, 0x3c, 0x85, 0x20, 0xa3, 0x8e, 0x5d, 0x41, 0x76, 0x55, 0x03, 0x48, 0xee, 0xc0, 0x6a,
	0x5d, 0x31, 0x89, 0x9c, 0x6f, 0x38, 0x77, 0x1e, 0xfd, 0xf1, 0x61, 0xf8, 0xc9, 0xe8, 0xbf, 0xa5,
	0x93, 0xc0, 0x40, 0x2d, 0x4a, 0x66, 0x45, 0x62, 0x8c, 0x18, 0xfb, 0xe9, 0xe4, 0x61, 0x4c, 0x5e,
	0xc0, 0x60, 0xc6, 0x33, 0xa3, 0x6a, 0xed, 0x10, 0xc6, 0xda, 0x80, 0xf1, 0x19, 0xcf, 0x58, 0x84,
	0x38, 0x39, 0x82, 0xe0, 0x8a, 0x4a, 0x4e, 0xa7, 0x19, 0xab, 0xc2, 0x95, 0x5d, 0xef, 0x60, 0xed,
	0xf0, 0x99, 0xb9, 0x64, 0xbf, 0x3c, 0xbe, 0x70, 0xf4, 0x87, 0x42, 0xc9, 0x45, 0xd4, 0x5c, 0x27,
	0xfb, 0x8d, 0x23, 0x43, 0x2c, 0xbf, 0x61, 0x32, 0xad, 0xb7, 0x1d, 0x83, 0x12, 0xc9, 0xa8, 0x62,
	0x69, 0x4c, 0x55, 0xb8, 0x8a, 0x4d, 0x04, 0x16, 0x39, 0x46, 0xba, 0x2e, 0x53, 0x47, 0x07, 0x86,
	0xb6, 0xc8, 0xb1, 0x22, 0xaf, 0xe1, 0x81, 0x64, 0x65, 0xb6, 0x88, 0x95, 0x88, 0xed, 0x73, 0xc6,
	0x3c, 0x0d, 0x01, 0xef, 0x6d, 0x21, 0xf5, 0x55, 0x58, 0xb5, 0xe7, 0x29, 0xf9, 0x02, 0x64, 0x79,
	0xbd, 0x69, 0x6d, 0x0d, 0x5b, 0x7b, 0xd9, 0x6d, 0x2d, 0x32, 0xb9, 0x37, 0x3a, 0x74, 0x25, 0x97,
	0x30, 0x99, 0xc0, 0xf6, 0x4c, 0xc8, 0x6b, 0x2a, 0xd3, 0x78, 0x26, 0x45, 0x1e, 0xeb, 0x22, 0x5a,
	0xc2, 0x3a, 0x1a, 0xfd, 0xbf, 0xe5, 0xce, 0xa4, 0xc8, 0x4f, 0xe7, 0x54, 0x9d, 0xa7, 0xe4, 0x2d,
	0x3c, 0xee, 0x24, 0xb4, 0x64, 0x6f, 0xa0, 0xec, 0xed, 0x56, 0x4e, 0x23, 0xfd, 0x3b, 0x3c, 0xea,
	0xa4, 0x35, 0xf2, 0x37, 0x51, 0xfe, 0x7e, 0x57, 0xfe, 0x59, 0x53, 0xe3, 0x46, 0x0b, 0xed, 0xf2,
	0x17, 0xed, 0xf7, 0x9a, 0xd6, 0x4a, 0x89, 0xa2, 0x0a, 0x7d, 0xac, 0x67, 0xdf, 0xeb, 0xc4, 0x80,
	0x91, 0x63, 0xc9, 0x1e, 0xf8, 0xbc, 0xc8, 0x78, 0xc1, 0xc2, 0xff, 0xee, 0xba, 0x67, 0xc9, 0x9d,
	0xf7, 0xb0, 0xd9, 0xfd, 0x2e, 0xd9, 0x02, 0xef, 0x07, 0x5b, 0xe0, 0x6f, 0x1a, 0x44, 0x3a, 0x24,
	0xdb, 0xb0, 0x72, 0x45, 0xb3, 0xda, 0xfd, 0xa8, 0xe6, 0x70, 0xd4, 0x7f, 0xd7, 0xdb, 0x39, 0x85,
	0x87, 0x77, 0xfa, 0x7f, 0xaf, 0x22, 0x1f, 0xe1, 0xc9, 0x3f, 0x5d, 0xb8, 0x4f, 0xa1, 0x51, 0x0a,
	0x03, 0x3d, 0x15, 0xb7, 0xe6, 0x6c, 0x0b, 0xbc, 0x5a, 0xba, 0x5d, 0xa0, 0x43, 0x3d, 0x65, 0x39,
	0xcf, 0x99, 0x9b, 0x32, 0x1d, 0x6b, 0xac, 0x35, 0xfb, 0x18, 0x6b, 0xac, 0xe2, 0xbf, 0xcc, 0xc4,
	0x7b, 0x11, 0xc6, 0xa3, 0x09, 0x0c, 0xad, 0x89, 0xe4, 0x15, 0xf8, 0xc6, 0xee, 0xb0, 0x87, 0x1e,
	0xaf, 0xb7, 0x3d, 0x8e, 0x2c, 0x37, 0x2a, 0xc1, 0x37, 0x08, 0xae, 0x1f, 0x5a, 0x2a, 0x8e, 0x09,
	0x66, 0xfd, 0x98, 0xe3, 0x72, 0xec, 0xfb, 0xad, 0xb1, 0x77, 0xeb, 0xc1, 0x6b, 0xad, 0x07, 0xdb,
	0xca, 0xa0, 0xd3, 0x4a, 0x22, 0x52, 0xb7, 0x94, 0x30, 0x3e, 0xd9, 0xff, 0xb6, 0x77, 0xc9, 0xd5,
	0xbc, 0x9e, 0x8e, 0x13, 0x91, 0x4f, 0xae, 0xd9, 0x94, 0x2b, 0x96, 0x4d, 0x70, 0x7d, 0x56, 0x13,
	0x56, 0x5c, 0xf2, 0x82, 0x4d, 0xb4, 0xd2, 0xa9, 0x8f, 0xd8, 0x9b, 0xbf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x34, 0xaf, 0xc9, 0x91, 0x63, 0x05, 0x00, 0x00,
}
