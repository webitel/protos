// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0-devel
// 	protoc        v3.21.6
// source: web_hook.proto

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

type DeleteWebHookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteWebHookRequest) Reset() {
	*x = DeleteWebHookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_hook_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteWebHookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWebHookRequest) ProtoMessage() {}

func (x *DeleteWebHookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_hook_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWebHookRequest.ProtoReflect.Descriptor instead.
func (*DeleteWebHookRequest) Descriptor() ([]byte, []int) {
	return file_web_hook_proto_rawDescGZIP(), []int{0}
}

func (x *DeleteWebHookRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateWebHookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Origin        []string `protobuf:"bytes,4,rep,name=origin,proto3" json:"origin,omitempty"`
	Schema        *Lookup  `protobuf:"bytes,5,opt,name=schema,proto3" json:"schema,omitempty"`
	Enabled       bool     `protobuf:"varint,6,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Authorization string   `protobuf:"bytes,7,opt,name=authorization,proto3" json:"authorization,omitempty"`
}

func (x *UpdateWebHookRequest) Reset() {
	*x = UpdateWebHookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_hook_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateWebHookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateWebHookRequest) ProtoMessage() {}

func (x *UpdateWebHookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_hook_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateWebHookRequest.ProtoReflect.Descriptor instead.
func (*UpdateWebHookRequest) Descriptor() ([]byte, []int) {
	return file_web_hook_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateWebHookRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateWebHookRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateWebHookRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateWebHookRequest) GetOrigin() []string {
	if x != nil {
		return x.Origin
	}
	return nil
}

func (x *UpdateWebHookRequest) GetSchema() *Lookup {
	if x != nil {
		return x.Schema
	}
	return nil
}

func (x *UpdateWebHookRequest) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *UpdateWebHookRequest) GetAuthorization() string {
	if x != nil {
		return x.Authorization
	}
	return ""
}

type PatchWebHookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Fields        []string `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty"`
	Name          string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description   string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Origin        []string `protobuf:"bytes,5,rep,name=origin,proto3" json:"origin,omitempty"`
	Schema        *Lookup  `protobuf:"bytes,6,opt,name=schema,proto3" json:"schema,omitempty"`
	Enabled       bool     `protobuf:"varint,7,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Authorization string   `protobuf:"bytes,8,opt,name=authorization,proto3" json:"authorization,omitempty"`
}

func (x *PatchWebHookRequest) Reset() {
	*x = PatchWebHookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_hook_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatchWebHookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatchWebHookRequest) ProtoMessage() {}

func (x *PatchWebHookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_hook_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatchWebHookRequest.ProtoReflect.Descriptor instead.
func (*PatchWebHookRequest) Descriptor() ([]byte, []int) {
	return file_web_hook_proto_rawDescGZIP(), []int{2}
}

func (x *PatchWebHookRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PatchWebHookRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *PatchWebHookRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PatchWebHookRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PatchWebHookRequest) GetOrigin() []string {
	if x != nil {
		return x.Origin
	}
	return nil
}

func (x *PatchWebHookRequest) GetSchema() *Lookup {
	if x != nil {
		return x.Schema
	}
	return nil
}

func (x *PatchWebHookRequest) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *PatchWebHookRequest) GetAuthorization() string {
	if x != nil {
		return x.Authorization
	}
	return ""
}

type ReadWebHookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReadWebHookRequest) Reset() {
	*x = ReadWebHookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_hook_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadWebHookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadWebHookRequest) ProtoMessage() {}

func (x *ReadWebHookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_hook_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadWebHookRequest.ProtoReflect.Descriptor instead.
func (*ReadWebHookRequest) Descriptor() ([]byte, []int) {
	return file_web_hook_proto_rawDescGZIP(), []int{3}
}

func (x *ReadWebHookRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SearchWebHookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page   int32    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size   int32    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Q      string   `protobuf:"bytes,3,opt,name=q,proto3" json:"q,omitempty"`
	Sort   string   `protobuf:"bytes,4,opt,name=sort,proto3" json:"sort,omitempty"`
	Fields []string `protobuf:"bytes,5,rep,name=fields,proto3" json:"fields,omitempty"`
	Id     []int32  `protobuf:"varint,6,rep,packed,name=id,proto3" json:"id,omitempty"`
}

func (x *SearchWebHookRequest) Reset() {
	*x = SearchWebHookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_hook_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchWebHookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchWebHookRequest) ProtoMessage() {}

func (x *SearchWebHookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_hook_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchWebHookRequest.ProtoReflect.Descriptor instead.
func (*SearchWebHookRequest) Descriptor() ([]byte, []int) {
	return file_web_hook_proto_rawDescGZIP(), []int{4}
}

func (x *SearchWebHookRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SearchWebHookRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *SearchWebHookRequest) GetQ() string {
	if x != nil {
		return x.Q
	}
	return ""
}

func (x *SearchWebHookRequest) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

func (x *SearchWebHookRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *SearchWebHookRequest) GetId() []int32 {
	if x != nil {
		return x.Id
	}
	return nil
}

type ListWebHook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Next  bool       `protobuf:"varint,1,opt,name=next,proto3" json:"next,omitempty"`
	Items []*WebHook `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListWebHook) Reset() {
	*x = ListWebHook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_hook_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWebHook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWebHook) ProtoMessage() {}

func (x *ListWebHook) ProtoReflect() protoreflect.Message {
	mi := &file_web_hook_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWebHook.ProtoReflect.Descriptor instead.
func (*ListWebHook) Descriptor() ([]byte, []int) {
	return file_web_hook_proto_rawDescGZIP(), []int{5}
}

func (x *ListWebHook) GetNext() bool {
	if x != nil {
		return x.Next
	}
	return false
}

func (x *ListWebHook) GetItems() []*WebHook {
	if x != nil {
		return x.Items
	}
	return nil
}

type CreateWebHookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description   string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Origin        []string `protobuf:"bytes,3,rep,name=origin,proto3" json:"origin,omitempty"`
	Schema        *Lookup  `protobuf:"bytes,4,opt,name=schema,proto3" json:"schema,omitempty"`
	Enabled       bool     `protobuf:"varint,5,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Authorization string   `protobuf:"bytes,6,opt,name=authorization,proto3" json:"authorization,omitempty"`
}

func (x *CreateWebHookRequest) Reset() {
	*x = CreateWebHookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_hook_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWebHookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWebHookRequest) ProtoMessage() {}

func (x *CreateWebHookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_web_hook_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWebHookRequest.ProtoReflect.Descriptor instead.
func (*CreateWebHookRequest) Descriptor() ([]byte, []int) {
	return file_web_hook_proto_rawDescGZIP(), []int{6}
}

func (x *CreateWebHookRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateWebHookRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateWebHookRequest) GetOrigin() []string {
	if x != nil {
		return x.Origin
	}
	return nil
}

func (x *CreateWebHookRequest) GetSchema() *Lookup {
	if x != nil {
		return x.Schema
	}
	return nil
}

func (x *CreateWebHookRequest) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *CreateWebHookRequest) GetAuthorization() string {
	if x != nil {
		return x.Authorization
	}
	return ""
}

type WebHook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Key           string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	CreatedAt     int64    `protobuf:"varint,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	CreatedBy     *Lookup  `protobuf:"bytes,4,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	UpdatedAt     int64    `protobuf:"varint,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	UpdatedBy     *Lookup  `protobuf:"bytes,6,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
	Name          string   `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	Description   string   `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	Origin        []string `protobuf:"bytes,9,rep,name=origin,proto3" json:"origin,omitempty"`
	Schema        *Lookup  `protobuf:"bytes,10,opt,name=schema,proto3" json:"schema,omitempty"`
	Enabled       bool     `protobuf:"varint,11,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Authorization string   `protobuf:"bytes,12,opt,name=authorization,proto3" json:"authorization,omitempty"`
}

func (x *WebHook) Reset() {
	*x = WebHook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_hook_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebHook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebHook) ProtoMessage() {}

func (x *WebHook) ProtoReflect() protoreflect.Message {
	mi := &file_web_hook_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebHook.ProtoReflect.Descriptor instead.
func (*WebHook) Descriptor() ([]byte, []int) {
	return file_web_hook_proto_rawDescGZIP(), []int{7}
}

func (x *WebHook) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *WebHook) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *WebHook) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *WebHook) GetCreatedBy() *Lookup {
	if x != nil {
		return x.CreatedBy
	}
	return nil
}

func (x *WebHook) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *WebHook) GetUpdatedBy() *Lookup {
	if x != nil {
		return x.UpdatedBy
	}
	return nil
}

func (x *WebHook) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WebHook) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *WebHook) GetOrigin() []string {
	if x != nil {
		return x.Origin
	}
	return nil
}

func (x *WebHook) GetSchema() *Lookup {
	if x != nil {
		return x.Schema
	}
	return nil
}

func (x *WebHook) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *WebHook) GetAuthorization() string {
	if x != nil {
		return x.Authorization
	}
	return ""
}

var File_web_hook_proto protoreflect.FileDescriptor

var file_web_hook_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x77, 0x65, 0x62, 0x5f, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x06, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x1a, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x65, 0x62,
	0x48, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0xdc, 0x01, 0x0a, 0x14,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x65, 0x62, 0x48, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b,
	0x75, 0x70, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xf3, 0x01, 0x0a, 0x13, 0x50,
	0x61, 0x74, 0x63, 0x68, 0x57, 0x65, 0x62, 0x48, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x24, 0x0a, 0x12, 0x52, 0x65, 0x61, 0x64, 0x57, 0x65, 0x62, 0x48, 0x6f, 0x6f, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x88, 0x01, 0x0a, 0x14, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x57, 0x65, 0x62, 0x48, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x71, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x01, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x06, 0x20, 0x03, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x48, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x65, 0x62, 0x48, 0x6f, 0x6f, 0x6b,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04,
	0x6e, 0x65, 0x78, 0x74, 0x12, 0x25, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x57, 0x65, 0x62,
	0x48, 0x6f, 0x6f, 0x6b, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xcc, 0x01, 0x0a, 0x14,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x65, 0x62, 0x48, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b,
	0x75, 0x70, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xfd, 0x02, 0x0a, 0x07, 0x57,
	0x65, 0x62, 0x48, 0x6f, 0x6f, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x62, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x42, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b,
	0x75, 0x70, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x85, 0x04, 0x0a, 0x0e, 0x57,
	0x65, 0x62, 0x48, 0x6f, 0x6f, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a,
	0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x65, 0x62, 0x48, 0x6f, 0x6f, 0x6b, 0x12, 0x1c,
	0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x65,
	0x62, 0x48, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x57, 0x65, 0x62, 0x48, 0x6f, 0x6f, 0x6b, 0x22, 0x10, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x3a, 0x01, 0x2a, 0x22, 0x05, 0x2f, 0x68, 0x6f, 0x6f, 0x6b, 0x12,
	0x51, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x57, 0x65, 0x62, 0x48, 0x6f, 0x6f, 0x6b,
	0x12, 0x1c, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x57, 0x65, 0x62, 0x48, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13,
	0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x65, 0x62, 0x48,
	0x6f, 0x6f, 0x6b, 0x22, 0x0d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x07, 0x12, 0x05, 0x2f, 0x68, 0x6f,
	0x6f, 0x6b, 0x12, 0x4e, 0x0a, 0x0b, 0x52, 0x65, 0x61, 0x64, 0x57, 0x65, 0x62, 0x48, 0x6f, 0x6f,
	0x6b, 0x12, 0x1a, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x57,
	0x65, 0x62, 0x48, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x57, 0x65, 0x62, 0x48, 0x6f, 0x6f, 0x6b, 0x22, 0x12,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x68, 0x6f, 0x6f, 0x6b, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x12, 0x53, 0x0a, 0x0c, 0x50, 0x61, 0x74, 0x63, 0x68, 0x57, 0x65, 0x62, 0x48, 0x6f,
	0x6f, 0x6b, 0x12, 0x1b, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x50, 0x61, 0x74, 0x63,
	0x68, 0x57, 0x65, 0x62, 0x48, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0f, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x57, 0x65, 0x62, 0x48, 0x6f, 0x6f, 0x6b,
	0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x3a, 0x01, 0x2a, 0x32, 0x0a, 0x2f, 0x68, 0x6f,
	0x6f, 0x6b, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x55, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x57, 0x65, 0x62, 0x48, 0x6f, 0x6f, 0x6b, 0x12, 0x1c, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x65, 0x62, 0x48, 0x6f, 0x6f, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e,
	0x57, 0x65, 0x62, 0x48, 0x6f, 0x6f, 0x6b, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x3a,
	0x01, 0x2a, 0x1a, 0x0a, 0x2f, 0x68, 0x6f, 0x6f, 0x6b, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x52,
	0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x65, 0x62, 0x48, 0x6f, 0x6f, 0x6b, 0x12,
	0x1c, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57,
	0x65, 0x62, 0x48, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x57, 0x65, 0x62, 0x48, 0x6f, 0x6f, 0x6b, 0x22, 0x12,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x2a, 0x0a, 0x2f, 0x68, 0x6f, 0x6f, 0x6b, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x77, 0x65, 0x62, 0x69, 0x74, 0x65, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_web_hook_proto_rawDescOnce sync.Once
	file_web_hook_proto_rawDescData = file_web_hook_proto_rawDesc
)

func file_web_hook_proto_rawDescGZIP() []byte {
	file_web_hook_proto_rawDescOnce.Do(func() {
		file_web_hook_proto_rawDescData = protoimpl.X.CompressGZIP(file_web_hook_proto_rawDescData)
	})
	return file_web_hook_proto_rawDescData
}

var file_web_hook_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_web_hook_proto_goTypes = []interface{}{
	(*DeleteWebHookRequest)(nil), // 0: engine.DeleteWebHookRequest
	(*UpdateWebHookRequest)(nil), // 1: engine.UpdateWebHookRequest
	(*PatchWebHookRequest)(nil),  // 2: engine.PatchWebHookRequest
	(*ReadWebHookRequest)(nil),   // 3: engine.ReadWebHookRequest
	(*SearchWebHookRequest)(nil), // 4: engine.SearchWebHookRequest
	(*ListWebHook)(nil),          // 5: engine.ListWebHook
	(*CreateWebHookRequest)(nil), // 6: engine.CreateWebHookRequest
	(*WebHook)(nil),              // 7: engine.WebHook
	(*Lookup)(nil),               // 8: engine.Lookup
}
var file_web_hook_proto_depIdxs = []int32{
	8,  // 0: engine.UpdateWebHookRequest.schema:type_name -> engine.Lookup
	8,  // 1: engine.PatchWebHookRequest.schema:type_name -> engine.Lookup
	7,  // 2: engine.ListWebHook.items:type_name -> engine.WebHook
	8,  // 3: engine.CreateWebHookRequest.schema:type_name -> engine.Lookup
	8,  // 4: engine.WebHook.created_by:type_name -> engine.Lookup
	8,  // 5: engine.WebHook.updated_by:type_name -> engine.Lookup
	8,  // 6: engine.WebHook.schema:type_name -> engine.Lookup
	6,  // 7: engine.WebHookService.CreateWebHook:input_type -> engine.CreateWebHookRequest
	4,  // 8: engine.WebHookService.SearchWebHook:input_type -> engine.SearchWebHookRequest
	3,  // 9: engine.WebHookService.ReadWebHook:input_type -> engine.ReadWebHookRequest
	2,  // 10: engine.WebHookService.PatchWebHook:input_type -> engine.PatchWebHookRequest
	1,  // 11: engine.WebHookService.UpdateWebHook:input_type -> engine.UpdateWebHookRequest
	0,  // 12: engine.WebHookService.DeleteWebHook:input_type -> engine.DeleteWebHookRequest
	7,  // 13: engine.WebHookService.CreateWebHook:output_type -> engine.WebHook
	5,  // 14: engine.WebHookService.SearchWebHook:output_type -> engine.ListWebHook
	7,  // 15: engine.WebHookService.ReadWebHook:output_type -> engine.WebHook
	7,  // 16: engine.WebHookService.PatchWebHook:output_type -> engine.WebHook
	7,  // 17: engine.WebHookService.UpdateWebHook:output_type -> engine.WebHook
	7,  // 18: engine.WebHookService.DeleteWebHook:output_type -> engine.WebHook
	13, // [13:19] is the sub-list for method output_type
	7,  // [7:13] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_web_hook_proto_init() }
func file_web_hook_proto_init() {
	if File_web_hook_proto != nil {
		return
	}
	file_const_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_web_hook_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteWebHookRequest); i {
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
		file_web_hook_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateWebHookRequest); i {
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
		file_web_hook_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatchWebHookRequest); i {
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
		file_web_hook_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadWebHookRequest); i {
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
		file_web_hook_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchWebHookRequest); i {
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
		file_web_hook_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWebHook); i {
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
		file_web_hook_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWebHookRequest); i {
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
		file_web_hook_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebHook); i {
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
			RawDescriptor: file_web_hook_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_web_hook_proto_goTypes,
		DependencyIndexes: file_web_hook_proto_depIdxs,
		MessageInfos:      file_web_hook_proto_msgTypes,
	}.Build()
	File_web_hook_proto = out.File
	file_web_hook_proto_rawDesc = nil
	file_web_hook_proto_goTypes = nil
	file_web_hook_proto_depIdxs = nil
}
