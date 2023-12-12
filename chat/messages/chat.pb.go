// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.3
// source: chat/messages/chat.proto

package messages

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The Chat info.
// Alias: participant, subscriber, member, peer, leg.
type Chat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// [D]omain[C]omponent primary ID.
	Dc int64 `protobuf:"varint,1,opt,name=dc,proto3" json:"dc,omitempty"`
	// Unique identifier for this chat.
	// [FROM] Member / Channel ID.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// [FROM] VIA text gateway profile.
	Via *Peer `protobuf:"bytes,5,opt,name=via,proto3" json:"via,omitempty"`
	// [FROM]: User identity. Seed.
	Peer *Peer `protobuf:"bytes,6,opt,name=peer,proto3" json:"peer,omitempty"`
	// [TO]: Chat title.
	Title string `protobuf:"bytes,8,opt,name=title,proto3" json:"title,omitempty"`
	// OPTIONAL. A non-zero value indicates that
	// the participant has left the chat. OFFLINE(!)
	Left int64 `protobuf:"varint,10,opt,name=left,proto3" json:"left,omitempty"`
	// OPTIONAL. A non-zero value indicates that
	// the participant has joined the chat.
	Join int64 `protobuf:"varint,11,opt,name=join,proto3" json:"join,omitempty"`
	// OPTIONAL. Invite[d] BY member info.
	Invite *Chat_Invite `protobuf:"bytes,13,opt,name=invite,proto3" json:"invite,omitempty"`
	// Context. Variables.
	Context map[string]string `protobuf:"bytes,15,rep,name=context,proto3" json:"context,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Chat) Reset() {
	*x = Chat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_messages_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chat) ProtoMessage() {}

func (x *Chat) ProtoReflect() protoreflect.Message {
	mi := &file_chat_messages_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chat.ProtoReflect.Descriptor instead.
func (*Chat) Descriptor() ([]byte, []int) {
	return file_chat_messages_chat_proto_rawDescGZIP(), []int{0}
}

func (x *Chat) GetDc() int64 {
	if x != nil {
		return x.Dc
	}
	return 0
}

func (x *Chat) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Chat) GetVia() *Peer {
	if x != nil {
		return x.Via
	}
	return nil
}

func (x *Chat) GetPeer() *Peer {
	if x != nil {
		return x.Peer
	}
	return nil
}

func (x *Chat) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Chat) GetLeft() int64 {
	if x != nil {
		return x.Left
	}
	return 0
}

func (x *Chat) GetJoin() int64 {
	if x != nil {
		return x.Join
	}
	return 0
}

func (x *Chat) GetInvite() *Chat_Invite {
	if x != nil {
		return x.Invite
	}
	return nil
}

func (x *Chat) GetContext() map[string]string {
	if x != nil {
		return x.Context
	}
	return nil
}

// Timerange filter value.
type Timerange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Since epochtime (milli).
	// **Match**: greater than ..
	Since int64 `protobuf:"varint,1,opt,name=since,proto3" json:"since,omitempty"`
	// Until epochtime (milli).
	// **Match**: less or equal ..
	Until int64 `protobuf:"varint,2,opt,name=until,proto3" json:"until,omitempty"`
}

func (x *Timerange) Reset() {
	*x = Timerange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_messages_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Timerange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Timerange) ProtoMessage() {}

func (x *Timerange) ProtoReflect() protoreflect.Message {
	mi := &file_chat_messages_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Timerange.ProtoReflect.Descriptor instead.
func (*Timerange) Descriptor() ([]byte, []int) {
	return file_chat_messages_chat_proto_rawDescGZIP(), []int{1}
}

func (x *Timerange) GetSince() int64 {
	if x != nil {
		return x.Since
	}
	return 0
}

func (x *Timerange) GetUntil() int64 {
	if x != nil {
		return x.Until
	}
	return 0
}

type ChatMembersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Page number to return. **default**: 1.
	Page int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	// Page records limit. **default**: 16.
	Size int32 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	// Search term: peer(type;name)
	Q string `protobuf:"bytes,6,opt,name=q,proto3" json:"q,omitempty"`
	// Sort records by { fields } specification.
	Sort []string `protobuf:"bytes,3,rep,name=sort,proto3" json:"sort,omitempty"`
	// Fields [Q]uery to build result dataset record.
	Fields []string `protobuf:"bytes,4,rep,name=fields,proto3" json:"fields,omitempty"`
	// ID of the chat dialog.
	ChatId string `protobuf:"bytes,5,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	// Set of unique chat member ID.
	Id []string `protobuf:"bytes,7,rep,name=id,proto3" json:"id,omitempty"`
	// [VIA] Text gateway.
	Via *Peer `protobuf:"bytes,8,opt,name=via,proto3" json:"via,omitempty"`
	// [PEER] Member of ...
	Peer *Peer `protobuf:"bytes,9,opt,name=peer,proto3" json:"peer,omitempty"`
	// Date within timerange.
	Date *Timerange `protobuf:"bytes,10,opt,name=date,proto3" json:"date,omitempty"`
	// Participants ONLY who are currently [not] connected to the chat.
	// ( left: ( 0 ? online : offline ) )
	Online *wrapperspb.BoolValue `protobuf:"bytes,11,opt,name=online,proto3" json:"online,omitempty"`
	// Participants ONLY who have [not] been connected to the chat.
	// ( join: ( 0 ? [request|abandoned] : connected ) )
	Joined *wrapperspb.BoolValue `protobuf:"bytes,12,opt,name=joined,proto3" json:"joined,omitempty"`
}

func (x *ChatMembersRequest) Reset() {
	*x = ChatMembersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_messages_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatMembersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatMembersRequest) ProtoMessage() {}

func (x *ChatMembersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_messages_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatMembersRequest.ProtoReflect.Descriptor instead.
func (*ChatMembersRequest) Descriptor() ([]byte, []int) {
	return file_chat_messages_chat_proto_rawDescGZIP(), []int{2}
}

func (x *ChatMembersRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ChatMembersRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ChatMembersRequest) GetQ() string {
	if x != nil {
		return x.Q
	}
	return ""
}

func (x *ChatMembersRequest) GetSort() []string {
	if x != nil {
		return x.Sort
	}
	return nil
}

func (x *ChatMembersRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *ChatMembersRequest) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

func (x *ChatMembersRequest) GetId() []string {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ChatMembersRequest) GetVia() *Peer {
	if x != nil {
		return x.Via
	}
	return nil
}

func (x *ChatMembersRequest) GetPeer() *Peer {
	if x != nil {
		return x.Peer
	}
	return nil
}

func (x *ChatMembersRequest) GetDate() *Timerange {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *ChatMembersRequest) GetOnline() *wrapperspb.BoolValue {
	if x != nil {
		return x.Online
	}
	return nil
}

func (x *ChatMembersRequest) GetJoined() *wrapperspb.BoolValue {
	if x != nil {
		return x.Joined
	}
	return nil
}

// ChatMembers dataset
type ChatMembers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Page of the chat participants.
	Data []*Chat `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	// Page number of results.
	Page int32 `protobuf:"varint,5,opt,name=page,proto3" json:"page,omitempty"`
	// Next page available ?
	Next bool `protobuf:"varint,6,opt,name=next,proto3" json:"next,omitempty"`
}

func (x *ChatMembers) Reset() {
	*x = ChatMembers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_messages_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatMembers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatMembers) ProtoMessage() {}

func (x *ChatMembers) ProtoReflect() protoreflect.Message {
	mi := &file_chat_messages_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatMembers.ProtoReflect.Descriptor instead.
func (*ChatMembers) Descriptor() ([]byte, []int) {
	return file_chat_messages_chat_proto_rawDescGZIP(), []int{3}
}

func (x *ChatMembers) GetData() []*Chat {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ChatMembers) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ChatMembers) GetNext() bool {
	if x != nil {
		return x.Next
	}
	return false
}

type Chat_Invite struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Timestamp when the invitation to join the chat was sent
	Date int64 `protobuf:"varint,1,opt,name=date,proto3" json:"date,omitempty"`
	// Chat member ID who invited to join the chat
	From string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
}

func (x *Chat_Invite) Reset() {
	*x = Chat_Invite{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_messages_chat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chat_Invite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chat_Invite) ProtoMessage() {}

func (x *Chat_Invite) ProtoReflect() protoreflect.Message {
	mi := &file_chat_messages_chat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chat_Invite.ProtoReflect.Descriptor instead.
func (*Chat_Invite) Descriptor() ([]byte, []int) {
	return file_chat_messages_chat_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Chat_Invite) GetDate() int64 {
	if x != nil {
		return x.Date
	}
	return 0
}

func (x *Chat_Invite) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

var File_chat_messages_chat_proto protoreflect.FileDescriptor

var file_chat_messages_chat_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x77, 0x65, 0x62, 0x69,
	0x74, 0x65, 0x6c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x1a, 0x18, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x8e, 0x03, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x64,
	0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x64, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x03, 0x76,
	0x69, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x77, 0x65, 0x62, 0x69, 0x74,
	0x65, 0x6c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x03, 0x76, 0x69,
	0x61, 0x12, 0x26, 0x0a, 0x04, 0x70, 0x65, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x77, 0x65, 0x62, 0x69, 0x74, 0x65, 0x6c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x50,
	0x65, 0x65, 0x72, 0x52, 0x04, 0x70, 0x65, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6c, 0x65, 0x66, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x6c,
	0x65, 0x66, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6a, 0x6f, 0x69, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x6a, 0x6f, 0x69, 0x6e, 0x12, 0x31, 0x0a, 0x06, 0x69, 0x6e, 0x76, 0x69, 0x74,
	0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x77, 0x65, 0x62, 0x69, 0x74, 0x65,
	0x6c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x49, 0x6e, 0x76, 0x69,
	0x74, 0x65, 0x52, 0x06, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x12, 0x39, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x77, 0x65,
	0x62, 0x69, 0x74, 0x65, 0x6c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x1a, 0x30, 0x0a, 0x06, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x1a, 0x3a, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x37, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x61, 0x6e, 0x67, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x6e, 0x74, 0x69, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x75, 0x6e, 0x74, 0x69, 0x6c, 0x22, 0x82, 0x03, 0x0a,
	0x12, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x71,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72,
	0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x24,
	0x0a, 0x03, 0x76, 0x69, 0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x77, 0x65,
	0x62, 0x69, 0x74, 0x65, 0x6c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52,
	0x03, 0x76, 0x69, 0x61, 0x12, 0x26, 0x0a, 0x04, 0x70, 0x65, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x77, 0x65, 0x62, 0x69, 0x74, 0x65, 0x6c, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x04, 0x70, 0x65, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x77, 0x65, 0x62,
	0x69, 0x74, 0x65, 0x6c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x61,
	0x6e, 0x67, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x6f, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x32, 0x0a,
	0x06, 0x6a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x6a, 0x6f, 0x69, 0x6e, 0x65,
	0x64, 0x22, 0x5d, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73,
	0x12, 0x26, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x77, 0x65, 0x62, 0x69, 0x74, 0x65, 0x6c, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68,
	0x61, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x65, 0x78, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6e, 0x65, 0x78, 0x74,
	0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77,
	0x65, 0x62, 0x69, 0x74, 0x65, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x68,
	0x61, 0x74, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_chat_messages_chat_proto_rawDescOnce sync.Once
	file_chat_messages_chat_proto_rawDescData = file_chat_messages_chat_proto_rawDesc
)

func file_chat_messages_chat_proto_rawDescGZIP() []byte {
	file_chat_messages_chat_proto_rawDescOnce.Do(func() {
		file_chat_messages_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_messages_chat_proto_rawDescData)
	})
	return file_chat_messages_chat_proto_rawDescData
}

var file_chat_messages_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_chat_messages_chat_proto_goTypes = []interface{}{
	(*Chat)(nil),                 // 0: webitel.chat.Chat
	(*Timerange)(nil),            // 1: webitel.chat.Timerange
	(*ChatMembersRequest)(nil),   // 2: webitel.chat.ChatMembersRequest
	(*ChatMembers)(nil),          // 3: webitel.chat.ChatMembers
	(*Chat_Invite)(nil),          // 4: webitel.chat.Chat.Invite
	nil,                          // 5: webitel.chat.Chat.ContextEntry
	(*Peer)(nil),                 // 6: webitel.chat.Peer
	(*wrapperspb.BoolValue)(nil), // 7: google.protobuf.BoolValue
}
var file_chat_messages_chat_proto_depIdxs = []int32{
	6,  // 0: webitel.chat.Chat.via:type_name -> webitel.chat.Peer
	6,  // 1: webitel.chat.Chat.peer:type_name -> webitel.chat.Peer
	4,  // 2: webitel.chat.Chat.invite:type_name -> webitel.chat.Chat.Invite
	5,  // 3: webitel.chat.Chat.context:type_name -> webitel.chat.Chat.ContextEntry
	6,  // 4: webitel.chat.ChatMembersRequest.via:type_name -> webitel.chat.Peer
	6,  // 5: webitel.chat.ChatMembersRequest.peer:type_name -> webitel.chat.Peer
	1,  // 6: webitel.chat.ChatMembersRequest.date:type_name -> webitel.chat.Timerange
	7,  // 7: webitel.chat.ChatMembersRequest.online:type_name -> google.protobuf.BoolValue
	7,  // 8: webitel.chat.ChatMembersRequest.joined:type_name -> google.protobuf.BoolValue
	0,  // 9: webitel.chat.ChatMembers.data:type_name -> webitel.chat.Chat
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_chat_messages_chat_proto_init() }
func file_chat_messages_chat_proto_init() {
	if File_chat_messages_chat_proto != nil {
		return
	}
	file_chat_messages_peer_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_chat_messages_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chat); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chat_messages_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Timerange); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chat_messages_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatMembersRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chat_messages_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatMembers); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chat_messages_chat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chat_Invite); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chat_messages_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chat_messages_chat_proto_goTypes,
		DependencyIndexes: file_chat_messages_chat_proto_depIdxs,
		MessageInfos:      file_chat_messages_chat_proto_msgTypes,
	}.Build()
	File_chat_messages_chat_proto = out.File
	file_chat_messages_chat_proto_rawDesc = nil
	file_chat_messages_chat_proto_goTypes = nil
	file_chat_messages_chat_proto_depIdxs = nil
}
