// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.6
// source: preset_query.proto

package engine

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PresetQuerySection int32

const (
	PresetQuerySection_section_default PresetQuerySection = 0
	PresetQuerySection_section_calls   PresetQuerySection = 1
)

// Enum value maps for PresetQuerySection.
var (
	PresetQuerySection_name = map[int32]string{
		0: "section_default",
		1: "section_calls",
	}
	PresetQuerySection_value = map[string]int32{
		"section_default": 0,
		"section_calls":   1,
	}
)

func (x PresetQuerySection) Enum() *PresetQuerySection {
	p := new(PresetQuerySection)
	*p = x
	return p
}

func (x PresetQuerySection) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PresetQuerySection) Descriptor() protoreflect.EnumDescriptor {
	return file_preset_query_proto_enumTypes[0].Descriptor()
}

func (PresetQuerySection) Type() protoreflect.EnumType {
	return &file_preset_query_proto_enumTypes[0]
}

func (x PresetQuerySection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PresetQuerySection.Descriptor instead.
func (PresetQuerySection) EnumDescriptor() ([]byte, []int) {
	return file_preset_query_proto_rawDescGZIP(), []int{0}
}

type DeletePresetQueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeletePresetQueryRequest) Reset() {
	*x = DeletePresetQueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_preset_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePresetQueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePresetQueryRequest) ProtoMessage() {}

func (x *DeletePresetQueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_preset_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePresetQueryRequest.ProtoReflect.Descriptor instead.
func (*DeletePresetQueryRequest) Descriptor() ([]byte, []int) {
	return file_preset_query_proto_rawDescGZIP(), []int{0}
}

func (x *DeletePresetQueryRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type PatchPresetQueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32              `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Fields      []string           `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty"`
	Name        string             `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string             `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Preset      *structpb.Value    `protobuf:"bytes,5,opt,name=preset,proto3" json:"preset,omitempty"`
	Section     PresetQuerySection `protobuf:"varint,6,opt,name=section,proto3,enum=engine.PresetQuerySection" json:"section,omitempty"`
}

func (x *PatchPresetQueryRequest) Reset() {
	*x = PatchPresetQueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_preset_query_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatchPresetQueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatchPresetQueryRequest) ProtoMessage() {}

func (x *PatchPresetQueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_preset_query_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatchPresetQueryRequest.ProtoReflect.Descriptor instead.
func (*PatchPresetQueryRequest) Descriptor() ([]byte, []int) {
	return file_preset_query_proto_rawDescGZIP(), []int{1}
}

func (x *PatchPresetQueryRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PatchPresetQueryRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *PatchPresetQueryRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PatchPresetQueryRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PatchPresetQueryRequest) GetPreset() *structpb.Value {
	if x != nil {
		return x.Preset
	}
	return nil
}

func (x *PatchPresetQueryRequest) GetSection() PresetQuerySection {
	if x != nil {
		return x.Section
	}
	return PresetQuerySection_section_default
}

type UpdatePresetQueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32              `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string             `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string             `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Preset      *structpb.Value    `protobuf:"bytes,4,opt,name=preset,proto3" json:"preset,omitempty"`
	Section     PresetQuerySection `protobuf:"varint,5,opt,name=section,proto3,enum=engine.PresetQuerySection" json:"section,omitempty"`
}

func (x *UpdatePresetQueryRequest) Reset() {
	*x = UpdatePresetQueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_preset_query_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePresetQueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePresetQueryRequest) ProtoMessage() {}

func (x *UpdatePresetQueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_preset_query_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePresetQueryRequest.ProtoReflect.Descriptor instead.
func (*UpdatePresetQueryRequest) Descriptor() ([]byte, []int) {
	return file_preset_query_proto_rawDescGZIP(), []int{2}
}

func (x *UpdatePresetQueryRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdatePresetQueryRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdatePresetQueryRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdatePresetQueryRequest) GetPreset() *structpb.Value {
	if x != nil {
		return x.Preset
	}
	return nil
}

func (x *UpdatePresetQueryRequest) GetSection() PresetQuerySection {
	if x != nil {
		return x.Section
	}
	return PresetQuerySection_section_default
}

type ReadPresetQueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReadPresetQueryRequest) Reset() {
	*x = ReadPresetQueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_preset_query_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadPresetQueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadPresetQueryRequest) ProtoMessage() {}

func (x *ReadPresetQueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_preset_query_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadPresetQueryRequest.ProtoReflect.Descriptor instead.
func (*ReadPresetQueryRequest) Descriptor() ([]byte, []int) {
	return file_preset_query_proto_rawDescGZIP(), []int{3}
}

func (x *ReadPresetQueryRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SearchPresetQueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page    int32                `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size    int32                `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Q       string               `protobuf:"bytes,3,opt,name=q,proto3" json:"q,omitempty"`
	Sort    string               `protobuf:"bytes,4,opt,name=sort,proto3" json:"sort,omitempty"`
	Fields  []string             `protobuf:"bytes,5,rep,name=fields,proto3" json:"fields,omitempty"`
	Id      []int32              `protobuf:"varint,6,rep,packed,name=id,proto3" json:"id,omitempty"`
	Section []PresetQuerySection `protobuf:"varint,7,rep,packed,name=section,proto3,enum=engine.PresetQuerySection" json:"section,omitempty"`
}

func (x *SearchPresetQueryRequest) Reset() {
	*x = SearchPresetQueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_preset_query_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchPresetQueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchPresetQueryRequest) ProtoMessage() {}

func (x *SearchPresetQueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_preset_query_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchPresetQueryRequest.ProtoReflect.Descriptor instead.
func (*SearchPresetQueryRequest) Descriptor() ([]byte, []int) {
	return file_preset_query_proto_rawDescGZIP(), []int{4}
}

func (x *SearchPresetQueryRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SearchPresetQueryRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *SearchPresetQueryRequest) GetQ() string {
	if x != nil {
		return x.Q
	}
	return ""
}

func (x *SearchPresetQueryRequest) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

func (x *SearchPresetQueryRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *SearchPresetQueryRequest) GetId() []int32 {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *SearchPresetQueryRequest) GetSection() []PresetQuerySection {
	if x != nil {
		return x.Section
	}
	return nil
}

type ListPresetQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Next  bool           `protobuf:"varint,1,opt,name=next,proto3" json:"next,omitempty"`
	Items []*PresetQuery `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListPresetQuery) Reset() {
	*x = ListPresetQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_preset_query_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPresetQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPresetQuery) ProtoMessage() {}

func (x *ListPresetQuery) ProtoReflect() protoreflect.Message {
	mi := &file_preset_query_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPresetQuery.ProtoReflect.Descriptor instead.
func (*ListPresetQuery) Descriptor() ([]byte, []int) {
	return file_preset_query_proto_rawDescGZIP(), []int{5}
}

func (x *ListPresetQuery) GetNext() bool {
	if x != nil {
		return x.Next
	}
	return false
}

func (x *ListPresetQuery) GetItems() []*PresetQuery {
	if x != nil {
		return x.Items
	}
	return nil
}

type CreatePresetQueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string             `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Preset      *structpb.Value    `protobuf:"bytes,3,opt,name=preset,proto3" json:"preset,omitempty"`
	Section     PresetQuerySection `protobuf:"varint,4,opt,name=section,proto3,enum=engine.PresetQuerySection" json:"section,omitempty"`
}

func (x *CreatePresetQueryRequest) Reset() {
	*x = CreatePresetQueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_preset_query_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePresetQueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePresetQueryRequest) ProtoMessage() {}

func (x *CreatePresetQueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_preset_query_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePresetQueryRequest.ProtoReflect.Descriptor instead.
func (*CreatePresetQueryRequest) Descriptor() ([]byte, []int) {
	return file_preset_query_proto_rawDescGZIP(), []int{6}
}

func (x *CreatePresetQueryRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreatePresetQueryRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreatePresetQueryRequest) GetPreset() *structpb.Value {
	if x != nil {
		return x.Preset
	}
	return nil
}

func (x *CreatePresetQueryRequest) GetSection() PresetQuerySection {
	if x != nil {
		return x.Section
	}
	return PresetQuerySection_section_default
}

type PresetQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32              `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string             `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string             `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Preset      *structpb.Value    `protobuf:"bytes,4,opt,name=preset,proto3" json:"preset,omitempty"`
	CreatedAt   int64              `protobuf:"varint,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   int64              `protobuf:"varint,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Section     PresetQuerySection `protobuf:"varint,7,opt,name=section,proto3,enum=engine.PresetQuerySection" json:"section,omitempty"`
}

func (x *PresetQuery) Reset() {
	*x = PresetQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_preset_query_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PresetQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PresetQuery) ProtoMessage() {}

func (x *PresetQuery) ProtoReflect() protoreflect.Message {
	mi := &file_preset_query_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PresetQuery.ProtoReflect.Descriptor instead.
func (*PresetQuery) Descriptor() ([]byte, []int) {
	return file_preset_query_proto_rawDescGZIP(), []int{7}
}

func (x *PresetQuery) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PresetQuery) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PresetQuery) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PresetQuery) GetPreset() *structpb.Value {
	if x != nil {
		return x.Preset
	}
	return nil
}

func (x *PresetQuery) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *PresetQuery) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *PresetQuery) GetSection() PresetQuerySection {
	if x != nil {
		return x.Section
	}
	return PresetQuerySection_section_default
}

var File_preset_query_proto protoreflect.FileDescriptor

var file_preset_query_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x22, 0xdd, 0x01, 0x0a, 0x17, 0x50, 0x61, 0x74, 0x63, 0x68, 0x50, 0x72,
	0x65, 0x73, 0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e,
	0x0a, 0x06, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x12, 0x34,
	0x0a, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1a, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0xc6, 0x01, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x65, 0x73, 0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x73, 0x65,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x06, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x12, 0x34, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2e, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x28, 0x0a,
	0x16, 0x52, 0x65, 0x61, 0x64, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0xc2, 0x01, 0x0a, 0x18, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x0c, 0x0a, 0x01,
	0x71, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f,
	0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x34, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x50, 0x0a, 0x0f,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6e,
	0x65, 0x78, 0x74, 0x12, 0x29, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x50, 0x72, 0x65, 0x73,
	0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xb6,
	0x01, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x2e, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x70, 0x72, 0x65, 0x73, 0x65,
	0x74, 0x12, 0x34, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x50, 0x72, 0x65, 0x73,
	0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07,
	0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xf7, 0x01, 0x0a, 0x0b, 0x50, 0x72, 0x65, 0x73,
	0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a,
	0x06, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x34, 0x0a, 0x07, 0x73,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2a, 0x3c, 0x0a, 0x12, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x13, 0x0a, 0x0f, 0x73, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d,
	0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x10, 0x01, 0x32,
	0xc9, 0x05, 0x0a, 0x12, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x70, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x20, 0x2e, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x65, 0x73, 0x65,
	0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01, 0x2a, 0x22, 0x19, 0x2f,
	0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x65, 0x73,
	0x65, 0x74, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x71, 0x0a, 0x11, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x20, 0x2e,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x65,
	0x73, 0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x65,
	0x73, 0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b,
	0x12, 0x19, 0x2f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x70,
	0x72, 0x65, 0x73, 0x65, 0x74, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x6e, 0x0a, 0x0f, 0x52,
	0x65, 0x61, 0x64, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1e,
	0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x50, 0x72, 0x65, 0x73,
	0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13,
	0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x12, 0x1e, 0x2f, 0x63, 0x61,
	0x6c, 0x6c, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74,
	0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x75, 0x0a, 0x11, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x20, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x13, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x50, 0x72, 0x65, 0x73,
	0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x3a,
	0x01, 0x2a, 0x1a, 0x1e, 0x2f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x2f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x12, 0x73, 0x0a, 0x10, 0x50, 0x61, 0x74, 0x63, 0x68, 0x50, 0x72, 0x65, 0x73, 0x65,
	0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1f, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e,
	0x50, 0x61, 0x74, 0x63, 0x68, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x22, 0x29, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x23, 0x3a, 0x01, 0x2a, 0x32, 0x1e, 0x2f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x2f, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x72, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x20, 0x2e, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x65, 0x73,
	0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13,
	0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x2a, 0x1e, 0x2f, 0x63, 0x61,
	0x6c, 0x6c, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74,
	0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x22, 0x5a, 0x20, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x65, 0x62, 0x69, 0x74, 0x65,
	0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_preset_query_proto_rawDescOnce sync.Once
	file_preset_query_proto_rawDescData = file_preset_query_proto_rawDesc
)

func file_preset_query_proto_rawDescGZIP() []byte {
	file_preset_query_proto_rawDescOnce.Do(func() {
		file_preset_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_preset_query_proto_rawDescData)
	})
	return file_preset_query_proto_rawDescData
}

var file_preset_query_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_preset_query_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_preset_query_proto_goTypes = []interface{}{
	(PresetQuerySection)(0),          // 0: engine.PresetQuerySection
	(*DeletePresetQueryRequest)(nil), // 1: engine.DeletePresetQueryRequest
	(*PatchPresetQueryRequest)(nil),  // 2: engine.PatchPresetQueryRequest
	(*UpdatePresetQueryRequest)(nil), // 3: engine.UpdatePresetQueryRequest
	(*ReadPresetQueryRequest)(nil),   // 4: engine.ReadPresetQueryRequest
	(*SearchPresetQueryRequest)(nil), // 5: engine.SearchPresetQueryRequest
	(*ListPresetQuery)(nil),          // 6: engine.ListPresetQuery
	(*CreatePresetQueryRequest)(nil), // 7: engine.CreatePresetQueryRequest
	(*PresetQuery)(nil),              // 8: engine.PresetQuery
	(*structpb.Value)(nil),           // 9: google.protobuf.Value
}
var file_preset_query_proto_depIdxs = []int32{
	9,  // 0: engine.PatchPresetQueryRequest.preset:type_name -> google.protobuf.Value
	0,  // 1: engine.PatchPresetQueryRequest.section:type_name -> engine.PresetQuerySection
	9,  // 2: engine.UpdatePresetQueryRequest.preset:type_name -> google.protobuf.Value
	0,  // 3: engine.UpdatePresetQueryRequest.section:type_name -> engine.PresetQuerySection
	0,  // 4: engine.SearchPresetQueryRequest.section:type_name -> engine.PresetQuerySection
	8,  // 5: engine.ListPresetQuery.items:type_name -> engine.PresetQuery
	9,  // 6: engine.CreatePresetQueryRequest.preset:type_name -> google.protobuf.Value
	0,  // 7: engine.CreatePresetQueryRequest.section:type_name -> engine.PresetQuerySection
	9,  // 8: engine.PresetQuery.preset:type_name -> google.protobuf.Value
	0,  // 9: engine.PresetQuery.section:type_name -> engine.PresetQuerySection
	7,  // 10: engine.PresetQueryService.CreatePresetQuery:input_type -> engine.CreatePresetQueryRequest
	5,  // 11: engine.PresetQueryService.SearchPresetQuery:input_type -> engine.SearchPresetQueryRequest
	4,  // 12: engine.PresetQueryService.ReadPresetQuery:input_type -> engine.ReadPresetQueryRequest
	3,  // 13: engine.PresetQueryService.UpdatePresetQuery:input_type -> engine.UpdatePresetQueryRequest
	2,  // 14: engine.PresetQueryService.PatchPresetQuery:input_type -> engine.PatchPresetQueryRequest
	1,  // 15: engine.PresetQueryService.DeletePresetQuery:input_type -> engine.DeletePresetQueryRequest
	8,  // 16: engine.PresetQueryService.CreatePresetQuery:output_type -> engine.PresetQuery
	6,  // 17: engine.PresetQueryService.SearchPresetQuery:output_type -> engine.ListPresetQuery
	8,  // 18: engine.PresetQueryService.ReadPresetQuery:output_type -> engine.PresetQuery
	8,  // 19: engine.PresetQueryService.UpdatePresetQuery:output_type -> engine.PresetQuery
	8,  // 20: engine.PresetQueryService.PatchPresetQuery:output_type -> engine.PresetQuery
	8,  // 21: engine.PresetQueryService.DeletePresetQuery:output_type -> engine.PresetQuery
	16, // [16:22] is the sub-list for method output_type
	10, // [10:16] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_preset_query_proto_init() }
func file_preset_query_proto_init() {
	if File_preset_query_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_preset_query_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePresetQueryRequest); i {
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
		file_preset_query_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatchPresetQueryRequest); i {
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
		file_preset_query_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePresetQueryRequest); i {
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
		file_preset_query_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadPresetQueryRequest); i {
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
		file_preset_query_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchPresetQueryRequest); i {
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
		file_preset_query_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPresetQuery); i {
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
		file_preset_query_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePresetQueryRequest); i {
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
		file_preset_query_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PresetQuery); i {
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
			RawDescriptor: file_preset_query_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_preset_query_proto_goTypes,
		DependencyIndexes: file_preset_query_proto_depIdxs,
		EnumInfos:         file_preset_query_proto_enumTypes,
		MessageInfos:      file_preset_query_proto_msgTypes,
	}.Build()
	File_preset_query_proto = out.File
	file_preset_query_proto_rawDesc = nil
	file_preset_query_proto_goTypes = nil
	file_preset_query_proto_depIdxs = nil
}
