// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        (unknown)
// source: message.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Account contact info to extend and replace legacy chat.Client message type
type Account struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Unique IDentifier
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Channel communication type
	Channel string `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel,omitempty"` // e.g.: bot, user, phone, telegram, facebook, viber, skype ...
	// Channel specific contact string
	Contact       string `protobuf:"bytes,3,opt,name=contact,proto3" json:"contact,omitempty"`                      // optional: channel specific contact string
	FirstName     string `protobuf:"bytes,4,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"` // optional
	LastName      string `protobuf:"bytes,5,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`    // optional
	Username      string `protobuf:"bytes,6,opt,name=username,proto3" json:"username,omitempty"`                    // required
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Account) Reset() {
	*x = Account{}
	mi := &file_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

func (x *Account) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Account) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *Account) GetContact() string {
	if x != nil {
		return x.Contact
	}
	return ""
}

func (x *Account) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Account) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Account) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type Message struct {
	state protoimpl.MessageState `protogen:"open.v1"`
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
	Variables map[string]string `protobuf:"bytes,5,rep,name=variables,proto3" json:"variables,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
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
	ReplyToVariables map[string]string `protobuf:"bytes,11,rep,name=reply_to_variables,json=replyToVariables,proto3" json:"reply_to_variables,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Unique identifier for the chat where the original message was sent
	ForwardFromChatId string `protobuf:"bytes,12,opt,name=forward_from_chat_id,json=forwardFromChatId,proto3" json:"forward_from_chat_id,omitempty"`
	// Message identifier in the chat specified in from_chat_id
	ForwardFromMessageId int64 `protobuf:"varint,13,opt,name=forward_from_message_id,json=forwardFromMessageId,proto3" json:"forward_from_message_id,omitempty"`
	// External message sent-binding to be able to identify corresponding internal message
	ForwardFromVariables map[string]string `protobuf:"bytes,14,rep,name=forward_from_variables,json=forwardFromVariables,proto3" json:"forward_from_variables,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Reply Markup Buttons SET
	Buttons []*Buttons `protobuf:"bytes,6,rep,name=buttons,proto3" json:"buttons,omitempty"`
	// Inline Keyboard Buttons SET
	Inline []*Buttons `protobuf:"bytes,15,rep,name=inline,proto3" json:"inline,omitempty"`
	// From sender user account
	From *Account `protobuf:"bytes,16,opt,name=from,proto3" json:"from,omitempty"`
	// Postback. Reply Button Click[ed].
	Postback *Postback `protobuf:"bytes,17,opt,name=postback,proto3" json:"postback,omitempty"`
	// An option used to block input to force the user
	// to respond with one of the `Buttons`.
	// Instructs client agents to disable input capabilities.
	// Can only be used with a set of `Buttons`.
	NoInput bool `protobuf:"varint,18,opt,name=no_input,json=noInput,proto3" json:"no_input,omitempty"`
	// NewChatMembers description for {"type":"joined"} notification
	NewChatMembers []*Account `protobuf:"bytes,20,rep,name=new_chat_members,json=newChatMembers,proto3" json:"new_chat_members,omitempty"`
	// LeftChatMember description for {"type":"left"} notification
	LeftChatMember *Account `protobuf:"bytes,21,opt,name=left_chat_member,json=leftChatMember,proto3" json:"left_chat_member,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *Message) Reset() {
	*x = Message{}
	mi := &file_message_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

func (x *Message) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Message) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Message) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Message) GetFile() *File {
	if x != nil {
		return x.File
	}
	return nil
}

func (x *Message) GetVariables() map[string]string {
	if x != nil {
		return x.Variables
	}
	return nil
}

func (x *Message) GetContact() *Account {
	if x != nil {
		return x.Contact
	}
	return nil
}

func (x *Message) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Message) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *Message) GetReplyToMessageId() int64 {
	if x != nil {
		return x.ReplyToMessageId
	}
	return 0
}

func (x *Message) GetReplyToVariables() map[string]string {
	if x != nil {
		return x.ReplyToVariables
	}
	return nil
}

func (x *Message) GetForwardFromChatId() string {
	if x != nil {
		return x.ForwardFromChatId
	}
	return ""
}

func (x *Message) GetForwardFromMessageId() int64 {
	if x != nil {
		return x.ForwardFromMessageId
	}
	return 0
}

func (x *Message) GetForwardFromVariables() map[string]string {
	if x != nil {
		return x.ForwardFromVariables
	}
	return nil
}

func (x *Message) GetButtons() []*Buttons {
	if x != nil {
		return x.Buttons
	}
	return nil
}

func (x *Message) GetInline() []*Buttons {
	if x != nil {
		return x.Inline
	}
	return nil
}

func (x *Message) GetFrom() *Account {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *Message) GetPostback() *Postback {
	if x != nil {
		return x.Postback
	}
	return nil
}

func (x *Message) GetNoInput() bool {
	if x != nil {
		return x.NoInput
	}
	return false
}

func (x *Message) GetNewChatMembers() []*Account {
	if x != nil {
		return x.NewChatMembers
	}
	return nil
}

func (x *Message) GetLeftChatMember() *Account {
	if x != nil {
		return x.LeftChatMember
	}
	return nil
}

type File struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Url           string                 `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Mime          string                 `protobuf:"bytes,3,opt,name=mime,proto3" json:"mime,omitempty"`
	Name          string                 `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Size          int64                  `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *File) Reset() {
	*x = File{}
	mi := &file_message_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2}
}

func (x *File) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *File) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *File) GetMime() string {
	if x != nil {
		return x.Mime
	}
	return ""
}

func (x *File) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *File) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type Buttons struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Button        []*Button              `protobuf:"bytes,1,rep,name=button,proto3" json:"button,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Buttons) Reset() {
	*x = Buttons{}
	mi := &file_message_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Buttons) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Buttons) ProtoMessage() {}

func (x *Buttons) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Buttons.ProtoReflect.Descriptor instead.
func (*Buttons) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{3}
}

func (x *Buttons) GetButton() []*Button {
	if x != nil {
		return x.Button
	}
	return nil
}

type Button struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Caption       string                 `protobuf:"bytes,1,opt,name=caption,proto3" json:"caption,omitempty"`
	Text          string                 `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Type          string                 `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Url           string                 `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	Code          string                 `protobuf:"bytes,5,opt,name=code,proto3" json:"code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Button) Reset() {
	*x = Button{}
	mi := &file_message_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Button) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Button) ProtoMessage() {}

func (x *Button) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Button.ProtoReflect.Descriptor instead.
func (*Button) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{4}
}

func (x *Button) GetCaption() string {
	if x != nil {
		return x.Caption
	}
	return ""
}

func (x *Button) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Button) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Button) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Button) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

// Postback. Reply Button Click[ed].
type Postback struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Message ID of the button.
	Mid int64 `protobuf:"varint,1,opt,name=mid,proto3" json:"mid,omitempty"`
	// Button's callback data associated.
	Code string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	// Button's caption. Text to display.
	Text          string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Postback) Reset() {
	*x = Postback{}
	mi := &file_message_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Postback) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Postback) ProtoMessage() {}

func (x *Postback) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Postback.ProtoReflect.Descriptor instead.
func (*Postback) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{5}
}

func (x *Postback) GetMid() int64 {
	if x != nil {
		return x.Mid
	}
	return 0
}

func (x *Postback) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Postback) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x13, 0x77, 0x65, 0x62, 0x69, 0x74, 0x65, 0x6c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x22, 0xa5, 0x01, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xea, 0x09, 0x0a,
	0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x12, 0x2d, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x77, 0x65, 0x62, 0x69, 0x74, 0x65, 0x6c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12,
	0x49, 0x0a, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x77, 0x65, 0x62, 0x69, 0x74, 0x65, 0x6c, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x77, 0x65,
	0x62, 0x69, 0x74, 0x65, 0x6c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x2d, 0x0a, 0x13, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x72,
	0x65, 0x70, 0x6c, 0x79, 0x54, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12,
	0x60, 0x0a, 0x12, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x74, 0x6f, 0x5f, 0x76, 0x61, 0x72, 0x69,
	0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x77, 0x65,
	0x62, 0x69, 0x74, 0x65, 0x6c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x54,
	0x6f, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x10, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x54, 0x6f, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65,
	0x73, 0x12, 0x2f, 0x0a, 0x14, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x66, 0x72, 0x6f,
	0x6d, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x11, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x68, 0x61, 0x74,
	0x49, 0x64, 0x12, 0x35, 0x0a, 0x17, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x66, 0x72,
	0x6f, 0x6d, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x14, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x46, 0x72, 0x6f, 0x6d,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x6c, 0x0a, 0x16, 0x66, 0x6f, 0x72,
	0x77, 0x61, 0x72, 0x64, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62,
	0x6c, 0x65, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x77, 0x65, 0x62, 0x69,
	0x74, 0x65, 0x6c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x46,
	0x72, 0x6f, 0x6d, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x14, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x56, 0x61,
	0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x07, 0x62, 0x75, 0x74, 0x74, 0x6f,
	0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x77, 0x65, 0x62, 0x69, 0x74,
	0x65, 0x6c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x42,
	0x75, 0x74, 0x74, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x73, 0x12,
	0x34, 0x0a, 0x06, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x77, 0x65, 0x62, 0x69, 0x74, 0x65, 0x6c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x73, 0x52, 0x06, 0x69,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x77, 0x65, 0x62, 0x69, 0x74, 0x65, 0x6c, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x39, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x62,
	0x61, 0x63, 0x6b, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x77, 0x65, 0x62, 0x69,
	0x74, 0x65, 0x6c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x50, 0x6f, 0x73, 0x74, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x62, 0x61,
	0x63, 0x6b, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x6f, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x12,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x6e, 0x6f, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x46, 0x0a,
	0x10, 0x6e, 0x65, 0x77, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x18, 0x14, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x77, 0x65, 0x62, 0x69, 0x74, 0x65,
	0x6c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0e, 0x6e, 0x65, 0x77, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x46, 0x0a, 0x10, 0x6c, 0x65, 0x66, 0x74, 0x5f, 0x63, 0x68,
	0x61, 0x74, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x77, 0x65, 0x62, 0x69, 0x74, 0x65, 0x6c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0e, 0x6c,
	0x65, 0x66, 0x74, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x1a, 0x3c, 0x0a,
	0x0e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x43, 0x0a, 0x15, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x54, 0x6f, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x1a, 0x47, 0x0a, 0x19, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x56,
	0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x64, 0x0a, 0x04, 0x46, 0x69, 0x6c,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6d, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22,
	0x3e, 0x0a, 0x07, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x73, 0x12, 0x33, 0x0a, 0x06, 0x62, 0x75,
	0x74, 0x74, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x77, 0x65, 0x62,
	0x69, 0x74, 0x65, 0x6c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x52, 0x06, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x22,
	0x70, 0x0a, 0x06, 0x42, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x61, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x22, 0x44, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x74, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6d, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x42, 0xba, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e,
	0x77, 0x65, 0x62, 0x69, 0x74, 0x65, 0x6c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x42, 0x0c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x77, 0x65, 0x62, 0x69, 0x74, 0x65, 0x6c, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0xa2, 0x02, 0x03, 0x57, 0x43, 0x53, 0xaa, 0x02,
	0x13, 0x57, 0x65, 0x62, 0x69, 0x74, 0x65, 0x6c, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0xca, 0x02, 0x13, 0x57, 0x65, 0x62, 0x69, 0x74, 0x65, 0x6c, 0x5c, 0x43,
	0x68, 0x61, 0x74, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0xe2, 0x02, 0x1f, 0x57, 0x65, 0x62,
	0x69, 0x74, 0x65, 0x6c, 0x5c, 0x43, 0x68, 0x61, 0x74, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x15, 0x57,
	0x65, 0x62, 0x69, 0x74, 0x65, 0x6c, 0x3a, 0x3a, 0x43, 0x68, 0x61, 0x74, 0x3a, 0x3a, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_message_proto_goTypes = []any{
	(*Account)(nil),  // 0: webitel.chat.server.Account
	(*Message)(nil),  // 1: webitel.chat.server.Message
	(*File)(nil),     // 2: webitel.chat.server.File
	(*Buttons)(nil),  // 3: webitel.chat.server.Buttons
	(*Button)(nil),   // 4: webitel.chat.server.Button
	(*Postback)(nil), // 5: webitel.chat.server.Postback
	nil,              // 6: webitel.chat.server.Message.VariablesEntry
	nil,              // 7: webitel.chat.server.Message.ReplyToVariablesEntry
	nil,              // 8: webitel.chat.server.Message.ForwardFromVariablesEntry
}
var file_message_proto_depIdxs = []int32{
	2,  // 0: webitel.chat.server.Message.file:type_name -> webitel.chat.server.File
	6,  // 1: webitel.chat.server.Message.variables:type_name -> webitel.chat.server.Message.VariablesEntry
	0,  // 2: webitel.chat.server.Message.contact:type_name -> webitel.chat.server.Account
	7,  // 3: webitel.chat.server.Message.reply_to_variables:type_name -> webitel.chat.server.Message.ReplyToVariablesEntry
	8,  // 4: webitel.chat.server.Message.forward_from_variables:type_name -> webitel.chat.server.Message.ForwardFromVariablesEntry
	3,  // 5: webitel.chat.server.Message.buttons:type_name -> webitel.chat.server.Buttons
	3,  // 6: webitel.chat.server.Message.inline:type_name -> webitel.chat.server.Buttons
	0,  // 7: webitel.chat.server.Message.from:type_name -> webitel.chat.server.Account
	5,  // 8: webitel.chat.server.Message.postback:type_name -> webitel.chat.server.Postback
	0,  // 9: webitel.chat.server.Message.new_chat_members:type_name -> webitel.chat.server.Account
	0,  // 10: webitel.chat.server.Message.left_chat_member:type_name -> webitel.chat.server.Account
	4,  // 11: webitel.chat.server.Buttons.button:type_name -> webitel.chat.server.Button
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}