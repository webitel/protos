// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0-devel
// 	protoc        v3.21.6
// source: queue_resources.proto

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

type DeleteQueueResourceGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	QueueId  int64 `protobuf:"varint,2,opt,name=queue_id,json=queueId,proto3" json:"queue_id,omitempty"`
	DomainId int64 `protobuf:"varint,3,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
}

func (x *DeleteQueueResourceGroupRequest) Reset() {
	*x = DeleteQueueResourceGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queue_resources_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteQueueResourceGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteQueueResourceGroupRequest) ProtoMessage() {}

func (x *DeleteQueueResourceGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_queue_resources_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteQueueResourceGroupRequest.ProtoReflect.Descriptor instead.
func (*DeleteQueueResourceGroupRequest) Descriptor() ([]byte, []int) {
	return file_queue_resources_proto_rawDescGZIP(), []int{0}
}

func (x *DeleteQueueResourceGroupRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeleteQueueResourceGroupRequest) GetQueueId() int64 {
	if x != nil {
		return x.QueueId
	}
	return 0
}

func (x *DeleteQueueResourceGroupRequest) GetDomainId() int64 {
	if x != nil {
		return x.DomainId
	}
	return 0
}

type UpdateQueueResourceGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	QueueId       int64   `protobuf:"varint,2,opt,name=queue_id,json=queueId,proto3" json:"queue_id,omitempty"`
	ResourceGroup *Lookup `protobuf:"bytes,3,opt,name=resource_group,json=resourceGroup,proto3" json:"resource_group,omitempty"`
	DomainId      int64   `protobuf:"varint,4,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
}

func (x *UpdateQueueResourceGroupRequest) Reset() {
	*x = UpdateQueueResourceGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queue_resources_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateQueueResourceGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateQueueResourceGroupRequest) ProtoMessage() {}

func (x *UpdateQueueResourceGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_queue_resources_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateQueueResourceGroupRequest.ProtoReflect.Descriptor instead.
func (*UpdateQueueResourceGroupRequest) Descriptor() ([]byte, []int) {
	return file_queue_resources_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateQueueResourceGroupRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateQueueResourceGroupRequest) GetQueueId() int64 {
	if x != nil {
		return x.QueueId
	}
	return 0
}

func (x *UpdateQueueResourceGroupRequest) GetResourceGroup() *Lookup {
	if x != nil {
		return x.ResourceGroup
	}
	return nil
}

func (x *UpdateQueueResourceGroupRequest) GetDomainId() int64 {
	if x != nil {
		return x.DomainId
	}
	return 0
}

type ReadQueueResourceGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	QueueId  int64 `protobuf:"varint,2,opt,name=queue_id,json=queueId,proto3" json:"queue_id,omitempty"`
	DomainId int64 `protobuf:"varint,3,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
}

func (x *ReadQueueResourceGroupRequest) Reset() {
	*x = ReadQueueResourceGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queue_resources_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadQueueResourceGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadQueueResourceGroupRequest) ProtoMessage() {}

func (x *ReadQueueResourceGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_queue_resources_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadQueueResourceGroupRequest.ProtoReflect.Descriptor instead.
func (*ReadQueueResourceGroupRequest) Descriptor() ([]byte, []int) {
	return file_queue_resources_proto_rawDescGZIP(), []int{2}
}

func (x *ReadQueueResourceGroupRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ReadQueueResourceGroupRequest) GetQueueId() int64 {
	if x != nil {
		return x.QueueId
	}
	return 0
}

func (x *ReadQueueResourceGroupRequest) GetDomainId() int64 {
	if x != nil {
		return x.DomainId
	}
	return 0
}

type QueueResourceGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ResourceGroup *Lookup `protobuf:"bytes,2,opt,name=resource_group,json=resourceGroup,proto3" json:"resource_group,omitempty"`
	Communication *Lookup `protobuf:"bytes,3,opt,name=communication,proto3" json:"communication,omitempty"`
}

func (x *QueueResourceGroup) Reset() {
	*x = QueueResourceGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queue_resources_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueueResourceGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueResourceGroup) ProtoMessage() {}

func (x *QueueResourceGroup) ProtoReflect() protoreflect.Message {
	mi := &file_queue_resources_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueueResourceGroup.ProtoReflect.Descriptor instead.
func (*QueueResourceGroup) Descriptor() ([]byte, []int) {
	return file_queue_resources_proto_rawDescGZIP(), []int{3}
}

func (x *QueueResourceGroup) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *QueueResourceGroup) GetResourceGroup() *Lookup {
	if x != nil {
		return x.ResourceGroup
	}
	return nil
}

func (x *QueueResourceGroup) GetCommunication() *Lookup {
	if x != nil {
		return x.Communication
	}
	return nil
}

type SearchQueueResourceGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueueId int64    `protobuf:"varint,1,opt,name=queue_id,json=queueId,proto3" json:"queue_id,omitempty"`
	Page    int32    `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Size    int32    `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Q       string   `protobuf:"bytes,4,opt,name=q,proto3" json:"q,omitempty"`
	Sort    string   `protobuf:"bytes,5,opt,name=sort,proto3" json:"sort,omitempty"`
	Fields  []string `protobuf:"bytes,6,rep,name=fields,proto3" json:"fields,omitempty"`
	Id      []uint32 `protobuf:"varint,7,rep,packed,name=id,proto3" json:"id,omitempty"`
}

func (x *SearchQueueResourceGroupRequest) Reset() {
	*x = SearchQueueResourceGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queue_resources_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchQueueResourceGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchQueueResourceGroupRequest) ProtoMessage() {}

func (x *SearchQueueResourceGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_queue_resources_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchQueueResourceGroupRequest.ProtoReflect.Descriptor instead.
func (*SearchQueueResourceGroupRequest) Descriptor() ([]byte, []int) {
	return file_queue_resources_proto_rawDescGZIP(), []int{4}
}

func (x *SearchQueueResourceGroupRequest) GetQueueId() int64 {
	if x != nil {
		return x.QueueId
	}
	return 0
}

func (x *SearchQueueResourceGroupRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SearchQueueResourceGroupRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *SearchQueueResourceGroupRequest) GetQ() string {
	if x != nil {
		return x.Q
	}
	return ""
}

func (x *SearchQueueResourceGroupRequest) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

func (x *SearchQueueResourceGroupRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *SearchQueueResourceGroupRequest) GetId() []uint32 {
	if x != nil {
		return x.Id
	}
	return nil
}

type ListQueueResourceGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Next  bool                  `protobuf:"varint,1,opt,name=next,proto3" json:"next,omitempty"`
	Items []*QueueResourceGroup `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListQueueResourceGroup) Reset() {
	*x = ListQueueResourceGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queue_resources_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListQueueResourceGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListQueueResourceGroup) ProtoMessage() {}

func (x *ListQueueResourceGroup) ProtoReflect() protoreflect.Message {
	mi := &file_queue_resources_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListQueueResourceGroup.ProtoReflect.Descriptor instead.
func (*ListQueueResourceGroup) Descriptor() ([]byte, []int) {
	return file_queue_resources_proto_rawDescGZIP(), []int{5}
}

func (x *ListQueueResourceGroup) GetNext() bool {
	if x != nil {
		return x.Next
	}
	return false
}

func (x *ListQueueResourceGroup) GetItems() []*QueueResourceGroup {
	if x != nil {
		return x.Items
	}
	return nil
}

type CreateQueueResourceGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueueId       int64   `protobuf:"varint,1,opt,name=queue_id,json=queueId,proto3" json:"queue_id,omitempty"`
	ResourceGroup *Lookup `protobuf:"bytes,2,opt,name=resource_group,json=resourceGroup,proto3" json:"resource_group,omitempty"`
	DomainId      int64   `protobuf:"varint,3,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
}

func (x *CreateQueueResourceGroupRequest) Reset() {
	*x = CreateQueueResourceGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_queue_resources_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateQueueResourceGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateQueueResourceGroupRequest) ProtoMessage() {}

func (x *CreateQueueResourceGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_queue_resources_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateQueueResourceGroupRequest.ProtoReflect.Descriptor instead.
func (*CreateQueueResourceGroupRequest) Descriptor() ([]byte, []int) {
	return file_queue_resources_proto_rawDescGZIP(), []int{6}
}

func (x *CreateQueueResourceGroupRequest) GetQueueId() int64 {
	if x != nil {
		return x.QueueId
	}
	return 0
}

func (x *CreateQueueResourceGroupRequest) GetResourceGroup() *Lookup {
	if x != nil {
		return x.ResourceGroup
	}
	return nil
}

func (x *CreateQueueResourceGroupRequest) GetDomainId() int64 {
	if x != nil {
		return x.DomainId
	}
	return 0
}

var File_queue_resources_proto protoreflect.FileDescriptor

var file_queue_resources_proto_rawDesc = []byte{
	0x0a, 0x15, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x1a,
	0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x1f, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x71, 0x75, 0x65, 0x75, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x49, 0x64, 0x22, 0xa0, 0x01, 0x0a, 0x1f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x0d, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x22, 0x67, 0x0a, 0x1d, 0x52, 0x65, 0x61, 0x64,
	0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49,
	0x64, 0x22, 0x91, 0x01, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x35, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70,
	0x52, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12,
	0x34, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e,
	0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xae, 0x01, 0x0a, 0x1f, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x0c, 0x0a, 0x01,
	0x71, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f,
	0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5e, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x75,
	0x65, 0x75, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04,
	0x6e, 0x65, 0x78, 0x74, 0x12, 0x30, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x51, 0x75, 0x65,
	0x75, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x90, 0x01, 0x0a, 0x1f, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x71, 0x75,
	0x65, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x71, 0x75,
	0x65, 0x75, 0x65, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x0d, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1b, 0x0a, 0x09,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x32, 0xae, 0x06, 0x0a, 0x15, 0x51, 0x75,
	0x65, 0x75, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x9a, 0x01, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x75,
	0x65, 0x75, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x27, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x39, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x33, 0x3a, 0x01, 0x2a,
	0x22, 0x2e, 0x2f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x71,
	0x75, 0x65, 0x75, 0x65, 0x73, 0x2f, 0x7b, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x7d,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x12, 0x9b, 0x01, 0x0a, 0x18, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x51, 0x75, 0x65, 0x75, 0x65,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x27, 0x2e,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x51, 0x75, 0x65,
	0x75, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x36, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30, 0x12, 0x2e,
	0x2f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x73, 0x2f, 0x7b, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x98,
	0x01, 0x0a, 0x16, 0x52, 0x65, 0x61, 0x64, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x25, 0x2e, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x3b, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x35, 0x12, 0x33, 0x2f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x73, 0x2f, 0x7b, 0x71, 0x75, 0x65, 0x75, 0x65,
	0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x9f, 0x01, 0x0a, 0x18, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x27, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x3e, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x38, 0x3a, 0x01, 0x2a, 0x1a, 0x33, 0x2f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x63, 0x65,
	0x6e, 0x74, 0x65, 0x72, 0x2f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x73, 0x2f, 0x7b, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x9c, 0x01, 0x0a, 0x18,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x27, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x3b, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x35, 0x2a, 0x33, 0x2f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x73, 0x2f, 0x7b, 0x71, 0x75, 0x65, 0x75,
	0x65, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x62, 0x69, 0x74, 0x65, 0x6c,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_queue_resources_proto_rawDescOnce sync.Once
	file_queue_resources_proto_rawDescData = file_queue_resources_proto_rawDesc
)

func file_queue_resources_proto_rawDescGZIP() []byte {
	file_queue_resources_proto_rawDescOnce.Do(func() {
		file_queue_resources_proto_rawDescData = protoimpl.X.CompressGZIP(file_queue_resources_proto_rawDescData)
	})
	return file_queue_resources_proto_rawDescData
}

var file_queue_resources_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_queue_resources_proto_goTypes = []interface{}{
	(*DeleteQueueResourceGroupRequest)(nil), // 0: engine.DeleteQueueResourceGroupRequest
	(*UpdateQueueResourceGroupRequest)(nil), // 1: engine.UpdateQueueResourceGroupRequest
	(*ReadQueueResourceGroupRequest)(nil),   // 2: engine.ReadQueueResourceGroupRequest
	(*QueueResourceGroup)(nil),              // 3: engine.QueueResourceGroup
	(*SearchQueueResourceGroupRequest)(nil), // 4: engine.SearchQueueResourceGroupRequest
	(*ListQueueResourceGroup)(nil),          // 5: engine.ListQueueResourceGroup
	(*CreateQueueResourceGroupRequest)(nil), // 6: engine.CreateQueueResourceGroupRequest
	(*Lookup)(nil),                          // 7: engine.Lookup
}
var file_queue_resources_proto_depIdxs = []int32{
	7,  // 0: engine.UpdateQueueResourceGroupRequest.resource_group:type_name -> engine.Lookup
	7,  // 1: engine.QueueResourceGroup.resource_group:type_name -> engine.Lookup
	7,  // 2: engine.QueueResourceGroup.communication:type_name -> engine.Lookup
	3,  // 3: engine.ListQueueResourceGroup.items:type_name -> engine.QueueResourceGroup
	7,  // 4: engine.CreateQueueResourceGroupRequest.resource_group:type_name -> engine.Lookup
	6,  // 5: engine.QueueResourcesService.CreateQueueResourceGroup:input_type -> engine.CreateQueueResourceGroupRequest
	4,  // 6: engine.QueueResourcesService.SearchQueueResourceGroup:input_type -> engine.SearchQueueResourceGroupRequest
	2,  // 7: engine.QueueResourcesService.ReadQueueResourceGroup:input_type -> engine.ReadQueueResourceGroupRequest
	1,  // 8: engine.QueueResourcesService.UpdateQueueResourceGroup:input_type -> engine.UpdateQueueResourceGroupRequest
	0,  // 9: engine.QueueResourcesService.DeleteQueueResourceGroup:input_type -> engine.DeleteQueueResourceGroupRequest
	3,  // 10: engine.QueueResourcesService.CreateQueueResourceGroup:output_type -> engine.QueueResourceGroup
	5,  // 11: engine.QueueResourcesService.SearchQueueResourceGroup:output_type -> engine.ListQueueResourceGroup
	3,  // 12: engine.QueueResourcesService.ReadQueueResourceGroup:output_type -> engine.QueueResourceGroup
	3,  // 13: engine.QueueResourcesService.UpdateQueueResourceGroup:output_type -> engine.QueueResourceGroup
	3,  // 14: engine.QueueResourcesService.DeleteQueueResourceGroup:output_type -> engine.QueueResourceGroup
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_queue_resources_proto_init() }
func file_queue_resources_proto_init() {
	if File_queue_resources_proto != nil {
		return
	}
	file_const_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_queue_resources_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteQueueResourceGroupRequest); i {
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
		file_queue_resources_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateQueueResourceGroupRequest); i {
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
		file_queue_resources_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadQueueResourceGroupRequest); i {
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
		file_queue_resources_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueueResourceGroup); i {
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
		file_queue_resources_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchQueueResourceGroupRequest); i {
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
		file_queue_resources_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListQueueResourceGroup); i {
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
		file_queue_resources_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateQueueResourceGroupRequest); i {
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
			RawDescriptor: file_queue_resources_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_queue_resources_proto_goTypes,
		DependencyIndexes: file_queue_resources_proto_depIdxs,
		MessageInfos:      file_queue_resources_proto_msgTypes,
	}.Build()
	File_queue_resources_proto = out.File
	file_queue_resources_proto_rawDesc = nil
	file_queue_resources_proto_goTypes = nil
	file_queue_resources_proto_depIdxs = nil
}
