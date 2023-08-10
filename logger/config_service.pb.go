// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.6
// source: config_service.proto

package proto

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

type Configs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  int32     `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Next  bool      `protobuf:"varint,2,opt,name=next,proto3" json:"next,omitempty"`
	Items []*Config `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *Configs) Reset() {
	*x = Configs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Configs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Configs) ProtoMessage() {}

func (x *Configs) ProtoReflect() protoreflect.Message {
	mi := &file_config_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Configs.ProtoReflect.Descriptor instead.
func (*Configs) Descriptor() ([]byte, []int) {
	return file_config_service_proto_rawDescGZIP(), []int{0}
}

func (x *Configs) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *Configs) GetNext() bool {
	if x != nil {
		return x.Next
	}
	return false
}

func (x *Configs) GetItems() []*Config {
	if x != nil {
		return x.Items
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_config_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_config_service_proto_rawDescGZIP(), []int{1}
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Object      *Lookup `protobuf:"bytes,2,opt,name=object,proto3" json:"object,omitempty"`
	Enabled     bool    `protobuf:"varint,3,opt,name=enabled,proto3" json:"enabled,omitempty"`
	DaysToStore int32   `protobuf:"varint,4,opt,name=days_to_store,json=daysToStore,proto3" json:"days_to_store,omitempty"`
	Period      string  `protobuf:"bytes,5,opt,name=period,proto3" json:"period,omitempty"`
	Storage     *Lookup `protobuf:"bytes,6,opt,name=storage,proto3" json:"storage,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_config_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_config_service_proto_rawDescGZIP(), []int{2}
}

func (x *Config) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Config) GetObject() *Lookup {
	if x != nil {
		return x.Object
	}
	return nil
}

func (x *Config) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Config) GetDaysToStore() int32 {
	if x != nil {
		return x.DaysToStore
	}
	return 0
}

func (x *Config) GetPeriod() string {
	if x != nil {
		return x.Period
	}
	return ""
}

func (x *Config) GetStorage() *Lookup {
	if x != nil {
		return x.Storage
	}
	return nil
}

type DeleteConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteConfigRequest) Reset() {
	*x = DeleteConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteConfigRequest) ProtoMessage() {}

func (x *DeleteConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteConfigRequest.ProtoReflect.Descriptor instead.
func (*DeleteConfigRequest) Descriptor() ([]byte, []int) {
	return file_config_service_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteConfigRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteConfigsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []int32 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *DeleteConfigsRequest) Reset() {
	*x = DeleteConfigsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteConfigsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteConfigsRequest) ProtoMessage() {}

func (x *DeleteConfigsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteConfigsRequest.ProtoReflect.Descriptor instead.
func (*DeleteConfigsRequest) Descriptor() ([]byte, []int) {
	return file_config_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteConfigsRequest) GetIds() []int32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type GetConfigByObjectIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectId int32 `protobuf:"varint,1,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	DomainId int32 `protobuf:"varint,2,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
}

func (x *GetConfigByObjectIdRequest) Reset() {
	*x = GetConfigByObjectIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigByObjectIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigByObjectIdRequest) ProtoMessage() {}

func (x *GetConfigByObjectIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigByObjectIdRequest.ProtoReflect.Descriptor instead.
func (*GetConfigByObjectIdRequest) Descriptor() ([]byte, []int) {
	return file_config_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetConfigByObjectIdRequest) GetObjectId() int32 {
	if x != nil {
		return x.ObjectId
	}
	return 0
}

func (x *GetConfigByObjectIdRequest) GetDomainId() int32 {
	if x != nil {
		return x.DomainId
	}
	return 0
}

type GetConfigByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` //  int32 domainId = 8;
}

func (x *GetConfigByIdRequest) Reset() {
	*x = GetConfigByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigByIdRequest) ProtoMessage() {}

func (x *GetConfigByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigByIdRequest.ProtoReflect.Descriptor instead.
func (*GetConfigByIdRequest) Descriptor() ([]byte, []int) {
	return file_config_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetConfigByIdRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetAllConfigsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page   int32    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size   int32    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Q      string   `protobuf:"bytes,3,opt,name=q,proto3" json:"q,omitempty"`
	Sort   string   `protobuf:"bytes,4,opt,name=sort,proto3" json:"sort,omitempty"`
	Fields []string `protobuf:"bytes,5,rep,name=fields,proto3" json:"fields,omitempty"`
}

func (x *GetAllConfigsRequest) Reset() {
	*x = GetAllConfigsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllConfigsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllConfigsRequest) ProtoMessage() {}

func (x *GetAllConfigsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllConfigsRequest.ProtoReflect.Descriptor instead.
func (*GetAllConfigsRequest) Descriptor() ([]byte, []int) {
	return file_config_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetAllConfigsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetAllConfigsRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *GetAllConfigsRequest) GetQ() string {
	if x != nil {
		return x.Q
	}
	return ""
}

func (x *GetAllConfigsRequest) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

func (x *GetAllConfigsRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

type UpdateConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConfigId    int32  `protobuf:"varint,1,opt,name=config_id,json=configId,proto3" json:"config_id,omitempty"`
	Enabled     bool   `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	DaysToStore int32  `protobuf:"varint,3,opt,name=days_to_store,json=daysToStore,proto3" json:"days_to_store,omitempty"`
	Period      string `protobuf:"bytes,4,opt,name=period,proto3" json:"period,omitempty"`
	StorageId   int32  `protobuf:"varint,5,opt,name=storage_id,json=storageId,proto3" json:"storage_id,omitempty"`
}

func (x *UpdateConfigRequest) Reset() {
	*x = UpdateConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateConfigRequest) ProtoMessage() {}

func (x *UpdateConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateConfigRequest.ProtoReflect.Descriptor instead.
func (*UpdateConfigRequest) Descriptor() ([]byte, []int) {
	return file_config_service_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateConfigRequest) GetConfigId() int32 {
	if x != nil {
		return x.ConfigId
	}
	return 0
}

func (x *UpdateConfigRequest) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *UpdateConfigRequest) GetDaysToStore() int32 {
	if x != nil {
		return x.DaysToStore
	}
	return 0
}

func (x *UpdateConfigRequest) GetPeriod() string {
	if x != nil {
		return x.Period
	}
	return ""
}

func (x *UpdateConfigRequest) GetStorageId() int32 {
	if x != nil {
		return x.StorageId
	}
	return 0
}

type PatchUpdateConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConfigId    int32    `protobuf:"varint,1,opt,name=config_id,json=configId,proto3" json:"config_id,omitempty"`
	Enabled     bool     `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	DaysToStore int32    `protobuf:"varint,3,opt,name=days_to_store,json=daysToStore,proto3" json:"days_to_store,omitempty"`
	Period      string   `protobuf:"bytes,4,opt,name=period,proto3" json:"period,omitempty"`
	StorageId   int32    `protobuf:"varint,5,opt,name=storage_id,json=storageId,proto3" json:"storage_id,omitempty"`
	Fields      []string `protobuf:"bytes,6,rep,name=fields,proto3" json:"fields,omitempty"`
}

func (x *PatchUpdateConfigRequest) Reset() {
	*x = PatchUpdateConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatchUpdateConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatchUpdateConfigRequest) ProtoMessage() {}

func (x *PatchUpdateConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatchUpdateConfigRequest.ProtoReflect.Descriptor instead.
func (*PatchUpdateConfigRequest) Descriptor() ([]byte, []int) {
	return file_config_service_proto_rawDescGZIP(), []int{9}
}

func (x *PatchUpdateConfigRequest) GetConfigId() int32 {
	if x != nil {
		return x.ConfigId
	}
	return 0
}

func (x *PatchUpdateConfigRequest) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *PatchUpdateConfigRequest) GetDaysToStore() int32 {
	if x != nil {
		return x.DaysToStore
	}
	return 0
}

func (x *PatchUpdateConfigRequest) GetPeriod() string {
	if x != nil {
		return x.Period
	}
	return ""
}

func (x *PatchUpdateConfigRequest) GetStorageId() int32 {
	if x != nil {
		return x.StorageId
	}
	return 0
}

func (x *PatchUpdateConfigRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

type InsertConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectId    int32  `protobuf:"varint,1,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	Enabled     bool   `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	DaysToStore int32  `protobuf:"varint,3,opt,name=days_to_store,json=daysToStore,proto3" json:"days_to_store,omitempty"`
	Period      string `protobuf:"bytes,4,opt,name=period,proto3" json:"period,omitempty"`
	StorageId   int32  `protobuf:"varint,5,opt,name=storage_id,json=storageId,proto3" json:"storage_id,omitempty"`
}

func (x *InsertConfigRequest) Reset() {
	*x = InsertConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertConfigRequest) ProtoMessage() {}

func (x *InsertConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertConfigRequest.ProtoReflect.Descriptor instead.
func (*InsertConfigRequest) Descriptor() ([]byte, []int) {
	return file_config_service_proto_rawDescGZIP(), []int{10}
}

func (x *InsertConfigRequest) GetObjectId() int32 {
	if x != nil {
		return x.ObjectId
	}
	return 0
}

func (x *InsertConfigRequest) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *InsertConfigRequest) GetDaysToStore() int32 {
	if x != nil {
		return x.DaysToStore
	}
	return 0
}

func (x *InsertConfigRequest) GetPeriod() string {
	if x != nil {
		return x.Period
	}
	return ""
}

func (x *InsertConfigRequest) GetStorageId() int32 {
	if x != nil {
		return x.StorageId
	}
	return 0
}

var File_config_service_proto protoreflect.FileDescriptor

var file_config_service_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x1a, 0x14,
	0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x57, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x04, 0x6e, 0x65, 0x78, 0x74, 0x12, 0x24, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x07, 0x0a, 0x05, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0xc0, 0x01, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x26, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52,
	0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x12, 0x22, 0x0a, 0x0d, 0x64, 0x61, 0x79, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x64, 0x61, 0x79, 0x73, 0x54, 0x6f,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x28, 0x0a,
	0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x07,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x22, 0x25, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x28,
	0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x05, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x56, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x79, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x64,
	0x22, 0x26, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x78, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x71, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x01, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x22, 0xa7, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x12, 0x22, 0x0a, 0x0d, 0x64, 0x61, 0x79, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x64, 0x61, 0x79, 0x73, 0x54, 0x6f,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0xc4, 0x01, 0x0a,
	0x18, 0x50, 0x61, 0x74, 0x63, 0x68, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x12, 0x22, 0x0a, 0x0d, 0x64, 0x61, 0x79, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x64, 0x61, 0x79, 0x73, 0x54, 0x6f, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x22, 0xa7, 0x01, 0x0a, 0x13, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x64, 0x61, 0x79, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x64, 0x61, 0x79, 0x73, 0x54,
	0x6f, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x49, 0x64, 0x32, 0xec, 0x05,
	0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x62, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x1b, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x6c,
	0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x25, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x1a, 0x1a, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x7b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f,
	0x69, 0x64, 0x7d, 0x12, 0x6c, 0x0a, 0x11, 0x50, 0x61, 0x74, 0x63, 0x68, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x20, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65,
	0x72, 0x2e, 0x50, 0x61, 0x74, 0x63, 0x68, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x6c, 0x6f, 0x67,
	0x67, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x32, 0x1a, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x7b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69, 0x64,
	0x7d, 0x12, 0x56, 0x0a, 0x0c, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x1b, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e,
	0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x19,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x22, 0x0e, 0x2f, 0x6c, 0x6f, 0x67, 0x67,
	0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x57, 0x0a, 0x0c, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1b, 0x2e, 0x6c, 0x6f, 0x67, 0x67,
	0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x2a, 0x13, 0x2f,
	0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x12, 0x57, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x73, 0x12, 0x1c, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0d, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x2a, 0x0e, 0x2f, 0x6c, 0x6f,
	0x67, 0x67, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x4b, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x79, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x49, 0x64, 0x12, 0x22, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x79, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1c, 0x2e, 0x6c, 0x6f, 0x67, 0x67,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x79, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72,
	0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12,
	0x13, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x12, 0x56, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x1c, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x73, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x6c,
	0x6f, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x16, 0x5a, 0x14,
	0x77, 0x65, 0x62, 0x69, 0x74, 0x65, 0x6c, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_service_proto_rawDescOnce sync.Once
	file_config_service_proto_rawDescData = file_config_service_proto_rawDesc
)

func file_config_service_proto_rawDescGZIP() []byte {
	file_config_service_proto_rawDescOnce.Do(func() {
		file_config_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_service_proto_rawDescData)
	})
	return file_config_service_proto_rawDescData
}

var file_config_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_config_service_proto_goTypes = []interface{}{
	(*Configs)(nil),                    // 0: logger.Configs
	(*Empty)(nil),                      // 1: logger.Empty
	(*Config)(nil),                     // 2: logger.Config
	(*DeleteConfigRequest)(nil),        // 3: logger.DeleteConfigRequest
	(*DeleteConfigsRequest)(nil),       // 4: logger.DeleteConfigsRequest
	(*GetConfigByObjectIdRequest)(nil), // 5: logger.GetConfigByObjectIdRequest
	(*GetConfigByIdRequest)(nil),       // 6: logger.GetConfigByIdRequest
	(*GetAllConfigsRequest)(nil),       // 7: logger.GetAllConfigsRequest
	(*UpdateConfigRequest)(nil),        // 8: logger.UpdateConfigRequest
	(*PatchUpdateConfigRequest)(nil),   // 9: logger.PatchUpdateConfigRequest
	(*InsertConfigRequest)(nil),        // 10: logger.InsertConfigRequest
	(*Lookup)(nil),                     // 11: logger.Lookup
}
var file_config_service_proto_depIdxs = []int32{
	2,  // 0: logger.Configs.items:type_name -> logger.Config
	11, // 1: logger.Config.object:type_name -> logger.Lookup
	11, // 2: logger.Config.storage:type_name -> logger.Lookup
	8,  // 3: logger.ConfigService.UpdateConfig:input_type -> logger.UpdateConfigRequest
	9,  // 4: logger.ConfigService.PatchUpdateConfig:input_type -> logger.PatchUpdateConfigRequest
	10, // 5: logger.ConfigService.InsertConfig:input_type -> logger.InsertConfigRequest
	3,  // 6: logger.ConfigService.DeleteConfig:input_type -> logger.DeleteConfigRequest
	4,  // 7: logger.ConfigService.DeleteConfigs:input_type -> logger.DeleteConfigsRequest
	5,  // 8: logger.ConfigService.GetConfigByObjectId:input_type -> logger.GetConfigByObjectIdRequest
	6,  // 9: logger.ConfigService.GetConfigById:input_type -> logger.GetConfigByIdRequest
	7,  // 10: logger.ConfigService.GetAllConfigs:input_type -> logger.GetAllConfigsRequest
	2,  // 11: logger.ConfigService.UpdateConfig:output_type -> logger.Config
	2,  // 12: logger.ConfigService.PatchUpdateConfig:output_type -> logger.Config
	2,  // 13: logger.ConfigService.InsertConfig:output_type -> logger.Config
	1,  // 14: logger.ConfigService.DeleteConfig:output_type -> logger.Empty
	1,  // 15: logger.ConfigService.DeleteConfigs:output_type -> logger.Empty
	2,  // 16: logger.ConfigService.GetConfigByObjectId:output_type -> logger.Config
	2,  // 17: logger.ConfigService.GetConfigById:output_type -> logger.Config
	0,  // 18: logger.ConfigService.GetAllConfigs:output_type -> logger.Configs
	11, // [11:19] is the sub-list for method output_type
	3,  // [3:11] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_config_service_proto_init() }
func file_config_service_proto_init() {
	if File_config_service_proto != nil {
		return
	}
	file_logger_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_config_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Configs); i {
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
		file_config_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_config_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_config_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteConfigRequest); i {
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
		file_config_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteConfigsRequest); i {
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
		file_config_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfigByObjectIdRequest); i {
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
		file_config_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfigByIdRequest); i {
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
		file_config_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllConfigsRequest); i {
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
		file_config_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateConfigRequest); i {
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
		file_config_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatchUpdateConfigRequest); i {
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
		file_config_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertConfigRequest); i {
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
			RawDescriptor: file_config_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_config_service_proto_goTypes,
		DependencyIndexes: file_config_service_proto_depIdxs,
		MessageInfos:      file_config_service_proto_msgTypes,
	}.Build()
	File_config_service_proto = out.File
	file_config_service_proto_rawDesc = nil
	file_config_service_proto_goTypes = nil
	file_config_service_proto_depIdxs = nil
}
