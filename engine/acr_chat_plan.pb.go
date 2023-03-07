// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: acr_chat_plan.proto

package engine

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ChatPlan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Schema      *Lookup `protobuf:"bytes,4,opt,name=schema,proto3" json:"schema,omitempty"`
	Enabled     bool    `protobuf:"varint,5,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *ChatPlan) Reset() {
	*x = ChatPlan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acr_chat_plan_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatPlan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatPlan) ProtoMessage() {}

func (x *ChatPlan) ProtoReflect() protoreflect.Message {
	mi := &file_acr_chat_plan_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatPlan.ProtoReflect.Descriptor instead.
func (*ChatPlan) Descriptor() ([]byte, []int) {
	return file_acr_chat_plan_proto_rawDescGZIP(), []int{0}
}

func (x *ChatPlan) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ChatPlan) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ChatPlan) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ChatPlan) GetSchema() *Lookup {
	if x != nil {
		return x.Schema
	}
	return nil
}

func (x *ChatPlan) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

type CreateChatPlanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string  `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Schema      *Lookup `protobuf:"bytes,3,opt,name=schema,proto3" json:"schema,omitempty"`
	Enabled     bool    `protobuf:"varint,4,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *CreateChatPlanRequest) Reset() {
	*x = CreateChatPlanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acr_chat_plan_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateChatPlanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChatPlanRequest) ProtoMessage() {}

func (x *CreateChatPlanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_acr_chat_plan_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateChatPlanRequest.ProtoReflect.Descriptor instead.
func (*CreateChatPlanRequest) Descriptor() ([]byte, []int) {
	return file_acr_chat_plan_proto_rawDescGZIP(), []int{1}
}

func (x *CreateChatPlanRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateChatPlanRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateChatPlanRequest) GetSchema() *Lookup {
	if x != nil {
		return x.Schema
	}
	return nil
}

func (x *CreateChatPlanRequest) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

type SearchChatPlanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page    int32    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size    int32    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Q       string   `protobuf:"bytes,3,opt,name=q,proto3" json:"q,omitempty"`
	Sort    string   `protobuf:"bytes,4,opt,name=sort,proto3" json:"sort,omitempty"`
	Fields  []string `protobuf:"bytes,5,rep,name=fields,proto3" json:"fields,omitempty"`
	Id      []int32  `protobuf:"varint,6,rep,packed,name=id,proto3" json:"id,omitempty"`
	Name    string   `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	Enabled bool     `protobuf:"varint,8,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *SearchChatPlanRequest) Reset() {
	*x = SearchChatPlanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acr_chat_plan_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchChatPlanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchChatPlanRequest) ProtoMessage() {}

func (x *SearchChatPlanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_acr_chat_plan_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchChatPlanRequest.ProtoReflect.Descriptor instead.
func (*SearchChatPlanRequest) Descriptor() ([]byte, []int) {
	return file_acr_chat_plan_proto_rawDescGZIP(), []int{2}
}

func (x *SearchChatPlanRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SearchChatPlanRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *SearchChatPlanRequest) GetQ() string {
	if x != nil {
		return x.Q
	}
	return ""
}

func (x *SearchChatPlanRequest) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

func (x *SearchChatPlanRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *SearchChatPlanRequest) GetId() []int32 {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *SearchChatPlanRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SearchChatPlanRequest) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

type ListChatPlan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Next  bool        `protobuf:"varint,1,opt,name=next,proto3" json:"next,omitempty"`
	Items []*ChatPlan `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListChatPlan) Reset() {
	*x = ListChatPlan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acr_chat_plan_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListChatPlan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListChatPlan) ProtoMessage() {}

func (x *ListChatPlan) ProtoReflect() protoreflect.Message {
	mi := &file_acr_chat_plan_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListChatPlan.ProtoReflect.Descriptor instead.
func (*ListChatPlan) Descriptor() ([]byte, []int) {
	return file_acr_chat_plan_proto_rawDescGZIP(), []int{3}
}

func (x *ListChatPlan) GetNext() bool {
	if x != nil {
		return x.Next
	}
	return false
}

func (x *ListChatPlan) GetItems() []*ChatPlan {
	if x != nil {
		return x.Items
	}
	return nil
}

type ReadChatPlanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReadChatPlanRequest) Reset() {
	*x = ReadChatPlanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acr_chat_plan_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadChatPlanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadChatPlanRequest) ProtoMessage() {}

func (x *ReadChatPlanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_acr_chat_plan_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadChatPlanRequest.ProtoReflect.Descriptor instead.
func (*ReadChatPlanRequest) Descriptor() ([]byte, []int) {
	return file_acr_chat_plan_proto_rawDescGZIP(), []int{4}
}

func (x *ReadChatPlanRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateChatPlanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Schema      *Lookup `protobuf:"bytes,4,opt,name=schema,proto3" json:"schema,omitempty"`
	Enabled     bool    `protobuf:"varint,5,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *UpdateChatPlanRequest) Reset() {
	*x = UpdateChatPlanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acr_chat_plan_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateChatPlanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateChatPlanRequest) ProtoMessage() {}

func (x *UpdateChatPlanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_acr_chat_plan_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateChatPlanRequest.ProtoReflect.Descriptor instead.
func (*UpdateChatPlanRequest) Descriptor() ([]byte, []int) {
	return file_acr_chat_plan_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateChatPlanRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateChatPlanRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateChatPlanRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateChatPlanRequest) GetSchema() *Lookup {
	if x != nil {
		return x.Schema
	}
	return nil
}

func (x *UpdateChatPlanRequest) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

type PatchChatPlanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Schema      *Lookup  `protobuf:"bytes,4,opt,name=schema,proto3" json:"schema,omitempty"`
	Enabled     bool     `protobuf:"varint,5,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Fields      []string `protobuf:"bytes,6,rep,name=fields,proto3" json:"fields,omitempty"`
}

func (x *PatchChatPlanRequest) Reset() {
	*x = PatchChatPlanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acr_chat_plan_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatchChatPlanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatchChatPlanRequest) ProtoMessage() {}

func (x *PatchChatPlanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_acr_chat_plan_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatchChatPlanRequest.ProtoReflect.Descriptor instead.
func (*PatchChatPlanRequest) Descriptor() ([]byte, []int) {
	return file_acr_chat_plan_proto_rawDescGZIP(), []int{6}
}

func (x *PatchChatPlanRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PatchChatPlanRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PatchChatPlanRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PatchChatPlanRequest) GetSchema() *Lookup {
	if x != nil {
		return x.Schema
	}
	return nil
}

func (x *PatchChatPlanRequest) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *PatchChatPlanRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

type DeleteChatPlanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteChatPlanRequest) Reset() {
	*x = DeleteChatPlanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_acr_chat_plan_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteChatPlanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteChatPlanRequest) ProtoMessage() {}

func (x *DeleteChatPlanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_acr_chat_plan_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteChatPlanRequest.ProtoReflect.Descriptor instead.
func (*DeleteChatPlanRequest) Descriptor() ([]byte, []int) {
	return file_acr_chat_plan_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteChatPlanRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_acr_chat_plan_proto protoreflect.FileDescriptor

var file_acr_chat_plan_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x63, 0x72, 0x5f, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x1a, 0x0b, 0x63,
	0x6f, 0x6e, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x01, 0x0a, 0x08, 0x43, 0x68, 0x61,
	0x74, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x06, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x06, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0x8f, 0x01,
	0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x50, 0x6c, 0x61, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a,
	0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x06, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22,
	0xb7, 0x01, 0x0a, 0x15, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x68, 0x61, 0x74, 0x50, 0x6c,
	0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x12, 0x0c, 0x0a, 0x01, 0x71, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x71, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73,
	0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x06, 0x20, 0x03, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0x4a, 0x0a, 0x0c, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x68, 0x61, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x65, 0x78,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x12, 0x26, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x25, 0x0a, 0x13, 0x52, 0x65, 0x61, 0x64, 0x43, 0x68, 0x61,
	0x74, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x9f, 0x01, 0x0a,
	0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x06,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x06, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0xb6,
	0x01, 0x0a, 0x14, 0x50, 0x61, 0x74, 0x63, 0x68, 0x43, 0x68, 0x61, 0x74, 0x50, 0x6c, 0x61, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a,
	0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x06, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x22, 0x27, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x68, 0x61, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x32, 0x85, 0x05, 0x0a, 0x16, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x43, 0x68, 0x61, 0x74,
	0x50, 0x6c, 0x61, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x64, 0x0a, 0x0e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x1d, 0x2e,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61,
	0x74, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x22, 0x21,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x22, 0x16, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67,
	0x2f, 0x6f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x3a, 0x01,
	0x2a, 0x12, 0x65, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x68, 0x61, 0x74, 0x50,
	0x6c, 0x61, 0x6e, 0x12, 0x1d, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x43, 0x68, 0x61, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x68, 0x61, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18,
	0x12, 0x16, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x6f, 0x75, 0x74, 0x62, 0x6f,
	0x75, 0x6e, 0x64, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x12, 0x62, 0x0a, 0x0c, 0x52, 0x65, 0x61, 0x64,
	0x43, 0x68, 0x61, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x1b, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x43, 0x68, 0x61, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x43,
	0x68, 0x61, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12,
	0x1b, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x6f, 0x75, 0x74, 0x62, 0x6f, 0x75,
	0x6e, 0x64, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x69, 0x0a, 0x0e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x1d,
	0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x68,
	0x61, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x22,
	0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x1a, 0x1b, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e,
	0x67, 0x2f, 0x6f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x67, 0x0a, 0x0d, 0x50, 0x61, 0x74, 0x63, 0x68,
	0x43, 0x68, 0x61, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x1c, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2e, 0x50, 0x61, 0x74, 0x63, 0x68, 0x43, 0x68, 0x61, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e,
	0x43, 0x68, 0x61, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20,
	0x32, 0x1b, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x6f, 0x75, 0x74, 0x62, 0x6f,
	0x75, 0x6e, 0x64, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a,
	0x12, 0x66, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x50, 0x6c,
	0x61, 0x6e, 0x12, 0x1d, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x10, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x50,
	0x6c, 0x61, 0x6e, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x2a, 0x1b, 0x2f, 0x72, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x6f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x2f, 0x63,
	0x68, 0x61, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x62, 0x69, 0x74, 0x65, 0x6c, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_acr_chat_plan_proto_rawDescOnce sync.Once
	file_acr_chat_plan_proto_rawDescData = file_acr_chat_plan_proto_rawDesc
)

func file_acr_chat_plan_proto_rawDescGZIP() []byte {
	file_acr_chat_plan_proto_rawDescOnce.Do(func() {
		file_acr_chat_plan_proto_rawDescData = protoimpl.X.CompressGZIP(file_acr_chat_plan_proto_rawDescData)
	})
	return file_acr_chat_plan_proto_rawDescData
}

var file_acr_chat_plan_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_acr_chat_plan_proto_goTypes = []interface{}{
	(*ChatPlan)(nil),              // 0: engine.ChatPlan
	(*CreateChatPlanRequest)(nil), // 1: engine.CreateChatPlanRequest
	(*SearchChatPlanRequest)(nil), // 2: engine.SearchChatPlanRequest
	(*ListChatPlan)(nil),          // 3: engine.ListChatPlan
	(*ReadChatPlanRequest)(nil),   // 4: engine.ReadChatPlanRequest
	(*UpdateChatPlanRequest)(nil), // 5: engine.UpdateChatPlanRequest
	(*PatchChatPlanRequest)(nil),  // 6: engine.PatchChatPlanRequest
	(*DeleteChatPlanRequest)(nil), // 7: engine.DeleteChatPlanRequest
	(*Lookup)(nil),                // 8: engine.Lookup
}
var file_acr_chat_plan_proto_depIdxs = []int32{
	8,  // 0: engine.ChatPlan.schema:type_name -> engine.Lookup
	8,  // 1: engine.CreateChatPlanRequest.schema:type_name -> engine.Lookup
	0,  // 2: engine.ListChatPlan.items:type_name -> engine.ChatPlan
	8,  // 3: engine.UpdateChatPlanRequest.schema:type_name -> engine.Lookup
	8,  // 4: engine.PatchChatPlanRequest.schema:type_name -> engine.Lookup
	1,  // 5: engine.RoutingChatPlanService.CreateChatPlan:input_type -> engine.CreateChatPlanRequest
	2,  // 6: engine.RoutingChatPlanService.SearchChatPlan:input_type -> engine.SearchChatPlanRequest
	4,  // 7: engine.RoutingChatPlanService.ReadChatPlan:input_type -> engine.ReadChatPlanRequest
	5,  // 8: engine.RoutingChatPlanService.UpdateChatPlan:input_type -> engine.UpdateChatPlanRequest
	6,  // 9: engine.RoutingChatPlanService.PatchChatPlan:input_type -> engine.PatchChatPlanRequest
	7,  // 10: engine.RoutingChatPlanService.DeleteChatPlan:input_type -> engine.DeleteChatPlanRequest
	0,  // 11: engine.RoutingChatPlanService.CreateChatPlan:output_type -> engine.ChatPlan
	3,  // 12: engine.RoutingChatPlanService.SearchChatPlan:output_type -> engine.ListChatPlan
	0,  // 13: engine.RoutingChatPlanService.ReadChatPlan:output_type -> engine.ChatPlan
	0,  // 14: engine.RoutingChatPlanService.UpdateChatPlan:output_type -> engine.ChatPlan
	0,  // 15: engine.RoutingChatPlanService.PatchChatPlan:output_type -> engine.ChatPlan
	0,  // 16: engine.RoutingChatPlanService.DeleteChatPlan:output_type -> engine.ChatPlan
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_acr_chat_plan_proto_init() }
func file_acr_chat_plan_proto_init() {
	if File_acr_chat_plan_proto != nil {
		return
	}
	file_const_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_acr_chat_plan_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatPlan); i {
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
		file_acr_chat_plan_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateChatPlanRequest); i {
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
		file_acr_chat_plan_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchChatPlanRequest); i {
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
		file_acr_chat_plan_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListChatPlan); i {
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
		file_acr_chat_plan_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadChatPlanRequest); i {
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
		file_acr_chat_plan_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateChatPlanRequest); i {
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
		file_acr_chat_plan_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatchChatPlanRequest); i {
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
		file_acr_chat_plan_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteChatPlanRequest); i {
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
			RawDescriptor: file_acr_chat_plan_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_acr_chat_plan_proto_goTypes,
		DependencyIndexes: file_acr_chat_plan_proto_depIdxs,
		MessageInfos:      file_acr_chat_plan_proto_msgTypes,
	}.Build()
	File_acr_chat_plan_proto = out.File
	file_acr_chat_plan_proto_rawDesc = nil
	file_acr_chat_plan_proto_goTypes = nil
	file_acr_chat_plan_proto_depIdxs = nil
}
