// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: queue_hook.proto

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

type DeleteQueueHookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueueId uint32 `protobuf:"varint,1,opt,name=queue_id,json=queueId,proto3" json:"queue_id,omitempty"`
	Id      uint32 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteQueueHookRequest) Reset() {
	*x = DeleteQueueHookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queue_hook_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteQueueHookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteQueueHookRequest) ProtoMessage() {}

func (x *DeleteQueueHookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_queue_hook_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteQueueHookRequest.ProtoReflect.Descriptor instead.
func (*DeleteQueueHookRequest) Descriptor() ([]byte, []int) {
	return file_queue_hook_proto_rawDescGZIP(), []int{0}
}

func (x *DeleteQueueHookRequest) GetQueueId() uint32 {
	if x != nil {
		return x.QueueId
	}
	return 0
}

func (x *DeleteQueueHookRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type PatchQueueHookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields     []string `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty"`
	QueueId    uint32   `protobuf:"varint,2,opt,name=queue_id,json=queueId,proto3" json:"queue_id,omitempty"`
	Id         uint32   `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Schema     *Lookup  `protobuf:"bytes,4,opt,name=schema,proto3" json:"schema,omitempty"`
	Event      string   `protobuf:"bytes,5,opt,name=event,proto3" json:"event,omitempty"`
	Enabled    bool     `protobuf:"varint,6,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Properties []string `protobuf:"bytes,7,rep,name=properties,proto3" json:"properties,omitempty"`
}

func (x *PatchQueueHookRequest) Reset() {
	*x = PatchQueueHookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queue_hook_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatchQueueHookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatchQueueHookRequest) ProtoMessage() {}

func (x *PatchQueueHookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_queue_hook_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatchQueueHookRequest.ProtoReflect.Descriptor instead.
func (*PatchQueueHookRequest) Descriptor() ([]byte, []int) {
	return file_queue_hook_proto_rawDescGZIP(), []int{1}
}

func (x *PatchQueueHookRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *PatchQueueHookRequest) GetQueueId() uint32 {
	if x != nil {
		return x.QueueId
	}
	return 0
}

func (x *PatchQueueHookRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PatchQueueHookRequest) GetSchema() *Lookup {
	if x != nil {
		return x.Schema
	}
	return nil
}

func (x *PatchQueueHookRequest) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

func (x *PatchQueueHookRequest) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *PatchQueueHookRequest) GetProperties() []string {
	if x != nil {
		return x.Properties
	}
	return nil
}

type UpdateQueueHookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueueId    uint32   `protobuf:"varint,1,opt,name=queue_id,json=queueId,proto3" json:"queue_id,omitempty"`
	Id         uint32   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Schema     *Lookup  `protobuf:"bytes,3,opt,name=schema,proto3" json:"schema,omitempty"`
	Event      string   `protobuf:"bytes,4,opt,name=event,proto3" json:"event,omitempty"`
	Enabled    bool     `protobuf:"varint,5,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Properties []string `protobuf:"bytes,6,rep,name=properties,proto3" json:"properties,omitempty"`
}

func (x *UpdateQueueHookRequest) Reset() {
	*x = UpdateQueueHookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queue_hook_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateQueueHookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateQueueHookRequest) ProtoMessage() {}

func (x *UpdateQueueHookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_queue_hook_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateQueueHookRequest.ProtoReflect.Descriptor instead.
func (*UpdateQueueHookRequest) Descriptor() ([]byte, []int) {
	return file_queue_hook_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateQueueHookRequest) GetQueueId() uint32 {
	if x != nil {
		return x.QueueId
	}
	return 0
}

func (x *UpdateQueueHookRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateQueueHookRequest) GetSchema() *Lookup {
	if x != nil {
		return x.Schema
	}
	return nil
}

func (x *UpdateQueueHookRequest) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

func (x *UpdateQueueHookRequest) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *UpdateQueueHookRequest) GetProperties() []string {
	if x != nil {
		return x.Properties
	}
	return nil
}

type SearchQueueHookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueueId  uint32   `protobuf:"varint,1,opt,name=queue_id,json=queueId,proto3" json:"queue_id,omitempty"`
	Page     int32    `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Size     int32    `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Q        string   `protobuf:"bytes,4,opt,name=q,proto3" json:"q,omitempty"`
	Sort     string   `protobuf:"bytes,5,opt,name=sort,proto3" json:"sort,omitempty"`
	Fields   []string `protobuf:"bytes,6,rep,name=fields,proto3" json:"fields,omitempty"`
	Id       []uint32 `protobuf:"varint,7,rep,packed,name=id,proto3" json:"id,omitempty"`
	SchemaId []uint32 `protobuf:"varint,8,rep,packed,name=schema_id,json=schemaId,proto3" json:"schema_id,omitempty"`
	Event    []string `protobuf:"bytes,9,rep,name=event,proto3" json:"event,omitempty"`
}

func (x *SearchQueueHookRequest) Reset() {
	*x = SearchQueueHookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queue_hook_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchQueueHookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchQueueHookRequest) ProtoMessage() {}

func (x *SearchQueueHookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_queue_hook_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchQueueHookRequest.ProtoReflect.Descriptor instead.
func (*SearchQueueHookRequest) Descriptor() ([]byte, []int) {
	return file_queue_hook_proto_rawDescGZIP(), []int{3}
}

func (x *SearchQueueHookRequest) GetQueueId() uint32 {
	if x != nil {
		return x.QueueId
	}
	return 0
}

func (x *SearchQueueHookRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SearchQueueHookRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *SearchQueueHookRequest) GetQ() string {
	if x != nil {
		return x.Q
	}
	return ""
}

func (x *SearchQueueHookRequest) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

func (x *SearchQueueHookRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *SearchQueueHookRequest) GetId() []uint32 {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *SearchQueueHookRequest) GetSchemaId() []uint32 {
	if x != nil {
		return x.SchemaId
	}
	return nil
}

func (x *SearchQueueHookRequest) GetEvent() []string {
	if x != nil {
		return x.Event
	}
	return nil
}

type CreateQueueHookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueueId    uint32   `protobuf:"varint,1,opt,name=queue_id,json=queueId,proto3" json:"queue_id,omitempty"`
	Schema     *Lookup  `protobuf:"bytes,2,opt,name=schema,proto3" json:"schema,omitempty"`
	Event      string   `protobuf:"bytes,3,opt,name=event,proto3" json:"event,omitempty"`
	Enabled    bool     `protobuf:"varint,4,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Properties []string `protobuf:"bytes,5,rep,name=properties,proto3" json:"properties,omitempty"`
}

func (x *CreateQueueHookRequest) Reset() {
	*x = CreateQueueHookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queue_hook_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateQueueHookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateQueueHookRequest) ProtoMessage() {}

func (x *CreateQueueHookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_queue_hook_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateQueueHookRequest.ProtoReflect.Descriptor instead.
func (*CreateQueueHookRequest) Descriptor() ([]byte, []int) {
	return file_queue_hook_proto_rawDescGZIP(), []int{4}
}

func (x *CreateQueueHookRequest) GetQueueId() uint32 {
	if x != nil {
		return x.QueueId
	}
	return 0
}

func (x *CreateQueueHookRequest) GetSchema() *Lookup {
	if x != nil {
		return x.Schema
	}
	return nil
}

func (x *CreateQueueHookRequest) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

func (x *CreateQueueHookRequest) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *CreateQueueHookRequest) GetProperties() []string {
	if x != nil {
		return x.Properties
	}
	return nil
}

type ReadQueueHookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueueId uint32 `protobuf:"varint,1,opt,name=queue_id,json=queueId,proto3" json:"queue_id,omitempty"`
	Id      uint32 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReadQueueHookRequest) Reset() {
	*x = ReadQueueHookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queue_hook_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadQueueHookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadQueueHookRequest) ProtoMessage() {}

func (x *ReadQueueHookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_queue_hook_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadQueueHookRequest.ProtoReflect.Descriptor instead.
func (*ReadQueueHookRequest) Descriptor() ([]byte, []int) {
	return file_queue_hook_proto_rawDescGZIP(), []int{5}
}

func (x *ReadQueueHookRequest) GetQueueId() uint32 {
	if x != nil {
		return x.QueueId
	}
	return 0
}

func (x *ReadQueueHookRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type QueueHook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Schema     *Lookup  `protobuf:"bytes,2,opt,name=schema,proto3" json:"schema,omitempty"`
	Event      string   `protobuf:"bytes,3,opt,name=event,proto3" json:"event,omitempty"`
	Enabled    bool     `protobuf:"varint,4,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Properties []string `protobuf:"bytes,5,rep,name=properties,proto3" json:"properties,omitempty"`
}

func (x *QueueHook) Reset() {
	*x = QueueHook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queue_hook_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueueHook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueHook) ProtoMessage() {}

func (x *QueueHook) ProtoReflect() protoreflect.Message {
	mi := &file_queue_hook_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueueHook.ProtoReflect.Descriptor instead.
func (*QueueHook) Descriptor() ([]byte, []int) {
	return file_queue_hook_proto_rawDescGZIP(), []int{6}
}

func (x *QueueHook) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *QueueHook) GetSchema() *Lookup {
	if x != nil {
		return x.Schema
	}
	return nil
}

func (x *QueueHook) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

func (x *QueueHook) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *QueueHook) GetProperties() []string {
	if x != nil {
		return x.Properties
	}
	return nil
}

type ListQueueHook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Next  bool         `protobuf:"varint,1,opt,name=next,proto3" json:"next,omitempty"`
	Items []*QueueHook `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListQueueHook) Reset() {
	*x = ListQueueHook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queue_hook_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListQueueHook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListQueueHook) ProtoMessage() {}

func (x *ListQueueHook) ProtoReflect() protoreflect.Message {
	mi := &file_queue_hook_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListQueueHook.ProtoReflect.Descriptor instead.
func (*ListQueueHook) Descriptor() ([]byte, []int) {
	return file_queue_hook_proto_rawDescGZIP(), []int{7}
}

func (x *ListQueueHook) GetNext() bool {
	if x != nil {
		return x.Next
	}
	return false
}

func (x *ListQueueHook) GetItems() []*QueueHook {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_queue_hook_proto protoreflect.FileDescriptor

var file_queue_hook_proto_rawDesc = []byte{
	0x0a, 0x10, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51,
	0x75, 0x65, 0x75, 0x65, 0x48, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x71, 0x75, 0x65, 0x75, 0x65, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0xd2, 0x01, 0x0a, 0x15, 0x50,
	0x61, 0x74, 0x63, 0x68, 0x51, 0x75, 0x65, 0x75, 0x65, 0x48, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x19, 0x0a, 0x08,
	0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x71, 0x75, 0x65, 0x75, 0x65, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x22,
	0xbb, 0x01, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65, 0x48,
	0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x71, 0x75,
	0x65, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x71, 0x75,
	0x65, 0x75, 0x65, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x4c,
	0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x22, 0xd8, 0x01,
	0x0a, 0x16, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x51, 0x75, 0x65, 0x75, 0x65, 0x48, 0x6f, 0x6f,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x75,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x71, 0x75, 0x65, 0x75,
	0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x71,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f,
	0x69, 0x64, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x08, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0xab, 0x01, 0x0a, 0x16, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65, 0x48, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x71, 0x75, 0x65, 0x75, 0x65, 0x49, 0x64, 0x12, 0x26,
	0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x06,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x22, 0x41, 0x0a, 0x14, 0x52, 0x65, 0x61, 0x64, 0x51, 0x75,
	0x65, 0x75, 0x65, 0x48, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x71, 0x75, 0x65, 0x75, 0x65, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x93, 0x01, 0x0a, 0x09, 0x51, 0x75,
	0x65, 0x75, 0x65, 0x48, 0x6f, 0x6f, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x22,
	0x4c, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x75, 0x65, 0x75, 0x65, 0x48, 0x6f, 0x6f, 0x6b,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04,
	0x6e, 0x65, 0x78, 0x74, 0x12, 0x27, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x51, 0x75, 0x65,
	0x75, 0x65, 0x48, 0x6f, 0x6f, 0x6b, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0xe5, 0x05,
	0x0a, 0x10, 0x51, 0x75, 0x65, 0x75, 0x65, 0x48, 0x6f, 0x6f, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x75, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x75,
	0x65, 0x48, 0x6f, 0x6f, 0x6b, 0x12, 0x1e, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65, 0x48, 0x6f, 0x6f, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x51,
	0x75, 0x65, 0x75, 0x65, 0x48, 0x6f, 0x6f, 0x6b, 0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29,
	0x22, 0x24, 0x2f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x71,
	0x75, 0x65, 0x75, 0x65, 0x73, 0x2f, 0x7b, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x7d,
	0x2f, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x76, 0x0a, 0x0f, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x51, 0x75, 0x65, 0x75, 0x65, 0x48, 0x6f, 0x6f, 0x6b, 0x12, 0x1e, 0x2e, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x51, 0x75, 0x65, 0x75,
	0x65, 0x48, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x75, 0x65, 0x75, 0x65, 0x48,
	0x6f, 0x6f, 0x6b, 0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x12, 0x24, 0x2f, 0x63, 0x61,
	0x6c, 0x6c, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x73,
	0x2f, 0x7b, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x68, 0x6f, 0x6f, 0x6b,
	0x73, 0x12, 0x73, 0x0a, 0x0d, 0x52, 0x65, 0x61, 0x64, 0x51, 0x75, 0x65, 0x75, 0x65, 0x48, 0x6f,
	0x6f, 0x6b, 0x12, 0x1c, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x52, 0x65, 0x61, 0x64,
	0x51, 0x75, 0x65, 0x75, 0x65, 0x48, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x11, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x48,
	0x6f, 0x6f, 0x6b, 0x22, 0x31, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2b, 0x12, 0x29, 0x2f, 0x63, 0x61,
	0x6c, 0x6c, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x73,
	0x2f, 0x7b, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x68, 0x6f, 0x6f, 0x6b,
	0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x7a, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x51, 0x75, 0x65, 0x75, 0x65, 0x48, 0x6f, 0x6f, 0x6b, 0x12, 0x1e, 0x2e, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65, 0x48, 0x6f,
	0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x48, 0x6f, 0x6f, 0x6b, 0x22, 0x34, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x2e, 0x1a, 0x29, 0x2f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x73, 0x2f, 0x7b, 0x71, 0x75, 0x65, 0x75, 0x65,
	0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a,
	0x01, 0x2a, 0x12, 0x78, 0x0a, 0x0e, 0x50, 0x61, 0x74, 0x63, 0x68, 0x51, 0x75, 0x65, 0x75, 0x65,
	0x48, 0x6f, 0x6f, 0x6b, 0x12, 0x1d, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x50, 0x61,
	0x74, 0x63, 0x68, 0x51, 0x75, 0x65, 0x75, 0x65, 0x48, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x51, 0x75, 0x65,
	0x75, 0x65, 0x48, 0x6f, 0x6f, 0x6b, 0x22, 0x34, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2e, 0x32, 0x29,
	0x2f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x73, 0x2f, 0x7b, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x68,
	0x6f, 0x6f, 0x6b, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x77, 0x0a, 0x0f,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65, 0x48, 0x6f, 0x6f, 0x6b, 0x12,
	0x1e, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51,
	0x75, 0x65, 0x75, 0x65, 0x48, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x11, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x48, 0x6f,
	0x6f, 0x6b, 0x22, 0x31, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2b, 0x2a, 0x29, 0x2f, 0x63, 0x61, 0x6c,
	0x6c, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x73, 0x2f,
	0x7b, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x68, 0x6f, 0x6f, 0x6b, 0x73,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x62, 0x69, 0x74, 0x65, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_queue_hook_proto_rawDescOnce sync.Once
	file_queue_hook_proto_rawDescData = file_queue_hook_proto_rawDesc
)

func file_queue_hook_proto_rawDescGZIP() []byte {
	file_queue_hook_proto_rawDescOnce.Do(func() {
		file_queue_hook_proto_rawDescData = protoimpl.X.CompressGZIP(file_queue_hook_proto_rawDescData)
	})
	return file_queue_hook_proto_rawDescData
}

var file_queue_hook_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_queue_hook_proto_goTypes = []interface{}{
	(*DeleteQueueHookRequest)(nil), // 0: engine.DeleteQueueHookRequest
	(*PatchQueueHookRequest)(nil),  // 1: engine.PatchQueueHookRequest
	(*UpdateQueueHookRequest)(nil), // 2: engine.UpdateQueueHookRequest
	(*SearchQueueHookRequest)(nil), // 3: engine.SearchQueueHookRequest
	(*CreateQueueHookRequest)(nil), // 4: engine.CreateQueueHookRequest
	(*ReadQueueHookRequest)(nil),   // 5: engine.ReadQueueHookRequest
	(*QueueHook)(nil),              // 6: engine.QueueHook
	(*ListQueueHook)(nil),          // 7: engine.ListQueueHook
	(*Lookup)(nil),                 // 8: engine.Lookup
}
var file_queue_hook_proto_depIdxs = []int32{
	8,  // 0: engine.PatchQueueHookRequest.schema:type_name -> engine.Lookup
	8,  // 1: engine.UpdateQueueHookRequest.schema:type_name -> engine.Lookup
	8,  // 2: engine.CreateQueueHookRequest.schema:type_name -> engine.Lookup
	8,  // 3: engine.QueueHook.schema:type_name -> engine.Lookup
	6,  // 4: engine.ListQueueHook.items:type_name -> engine.QueueHook
	4,  // 5: engine.QueueHookService.CreateQueueHook:input_type -> engine.CreateQueueHookRequest
	3,  // 6: engine.QueueHookService.SearchQueueHook:input_type -> engine.SearchQueueHookRequest
	5,  // 7: engine.QueueHookService.ReadQueueHook:input_type -> engine.ReadQueueHookRequest
	2,  // 8: engine.QueueHookService.UpdateQueueHook:input_type -> engine.UpdateQueueHookRequest
	1,  // 9: engine.QueueHookService.PatchQueueHook:input_type -> engine.PatchQueueHookRequest
	0,  // 10: engine.QueueHookService.DeleteQueueHook:input_type -> engine.DeleteQueueHookRequest
	6,  // 11: engine.QueueHookService.CreateQueueHook:output_type -> engine.QueueHook
	7,  // 12: engine.QueueHookService.SearchQueueHook:output_type -> engine.ListQueueHook
	6,  // 13: engine.QueueHookService.ReadQueueHook:output_type -> engine.QueueHook
	6,  // 14: engine.QueueHookService.UpdateQueueHook:output_type -> engine.QueueHook
	6,  // 15: engine.QueueHookService.PatchQueueHook:output_type -> engine.QueueHook
	6,  // 16: engine.QueueHookService.DeleteQueueHook:output_type -> engine.QueueHook
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_queue_hook_proto_init() }
func file_queue_hook_proto_init() {
	if File_queue_hook_proto != nil {
		return
	}
	file_const_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_queue_hook_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteQueueHookRequest); i {
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
		file_queue_hook_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatchQueueHookRequest); i {
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
		file_queue_hook_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateQueueHookRequest); i {
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
		file_queue_hook_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchQueueHookRequest); i {
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
		file_queue_hook_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateQueueHookRequest); i {
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
		file_queue_hook_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadQueueHookRequest); i {
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
		file_queue_hook_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueueHook); i {
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
		file_queue_hook_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListQueueHook); i {
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
			RawDescriptor: file_queue_hook_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_queue_hook_proto_goTypes,
		DependencyIndexes: file_queue_hook_proto_depIdxs,
		MessageInfos:      file_queue_hook_proto_msgTypes,
	}.Build()
	File_queue_hook_proto = out.File
	file_queue_hook_proto_rawDesc = nil
	file_queue_hook_proto_goTypes = nil
	file_queue_hook_proto_depIdxs = nil
}
