// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.6
// source: schema_variables.proto

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

type SchemaVariable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32           `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Encrypt bool            `protobuf:"varint,3,opt,name=encrypt,proto3" json:"encrypt,omitempty"`
	Value   *structpb.Value `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SchemaVariable) Reset() {
	*x = SchemaVariable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_variables_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchemaVariable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchemaVariable) ProtoMessage() {}

func (x *SchemaVariable) ProtoReflect() protoreflect.Message {
	mi := &file_schema_variables_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchemaVariable.ProtoReflect.Descriptor instead.
func (*SchemaVariable) Descriptor() ([]byte, []int) {
	return file_schema_variables_proto_rawDescGZIP(), []int{0}
}

func (x *SchemaVariable) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SchemaVariable) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SchemaVariable) GetEncrypt() bool {
	if x != nil {
		return x.Encrypt
	}
	return false
}

func (x *SchemaVariable) GetValue() *structpb.Value {
	if x != nil {
		return x.Value
	}
	return nil
}

type CreateSchemaVariableRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Encrypt bool            `protobuf:"varint,2,opt,name=encrypt,proto3" json:"encrypt,omitempty"`
	Value   *structpb.Value `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *CreateSchemaVariableRequest) Reset() {
	*x = CreateSchemaVariableRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_variables_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSchemaVariableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSchemaVariableRequest) ProtoMessage() {}

func (x *CreateSchemaVariableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schema_variables_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSchemaVariableRequest.ProtoReflect.Descriptor instead.
func (*CreateSchemaVariableRequest) Descriptor() ([]byte, []int) {
	return file_schema_variables_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSchemaVariableRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateSchemaVariableRequest) GetEncrypt() bool {
	if x != nil {
		return x.Encrypt
	}
	return false
}

func (x *CreateSchemaVariableRequest) GetValue() *structpb.Value {
	if x != nil {
		return x.Value
	}
	return nil
}

type SearchSchemaVariableRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page   int32    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size   int32    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Q      string   `protobuf:"bytes,3,opt,name=q,proto3" json:"q,omitempty"`
	Sort   string   `protobuf:"bytes,5,opt,name=sort,proto3" json:"sort,omitempty"`
	Fields []string `protobuf:"bytes,6,rep,name=fields,proto3" json:"fields,omitempty"`
}

func (x *SearchSchemaVariableRequest) Reset() {
	*x = SearchSchemaVariableRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_variables_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchSchemaVariableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchSchemaVariableRequest) ProtoMessage() {}

func (x *SearchSchemaVariableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schema_variables_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchSchemaVariableRequest.ProtoReflect.Descriptor instead.
func (*SearchSchemaVariableRequest) Descriptor() ([]byte, []int) {
	return file_schema_variables_proto_rawDescGZIP(), []int{2}
}

func (x *SearchSchemaVariableRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SearchSchemaVariableRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *SearchSchemaVariableRequest) GetQ() string {
	if x != nil {
		return x.Q
	}
	return ""
}

func (x *SearchSchemaVariableRequest) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

func (x *SearchSchemaVariableRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

type ListSchemaVariable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Next  bool              `protobuf:"varint,1,opt,name=next,proto3" json:"next,omitempty"`
	Items []*SchemaVariable `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListSchemaVariable) Reset() {
	*x = ListSchemaVariable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_variables_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSchemaVariable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSchemaVariable) ProtoMessage() {}

func (x *ListSchemaVariable) ProtoReflect() protoreflect.Message {
	mi := &file_schema_variables_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSchemaVariable.ProtoReflect.Descriptor instead.
func (*ListSchemaVariable) Descriptor() ([]byte, []int) {
	return file_schema_variables_proto_rawDescGZIP(), []int{3}
}

func (x *ListSchemaVariable) GetNext() bool {
	if x != nil {
		return x.Next
	}
	return false
}

func (x *ListSchemaVariable) GetItems() []*SchemaVariable {
	if x != nil {
		return x.Items
	}
	return nil
}

type ReadSchemaVariableRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReadSchemaVariableRequest) Reset() {
	*x = ReadSchemaVariableRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_variables_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadSchemaVariableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadSchemaVariableRequest) ProtoMessage() {}

func (x *ReadSchemaVariableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schema_variables_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadSchemaVariableRequest.ProtoReflect.Descriptor instead.
func (*ReadSchemaVariableRequest) Descriptor() ([]byte, []int) {
	return file_schema_variables_proto_rawDescGZIP(), []int{4}
}

func (x *ReadSchemaVariableRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateSchemaVariableRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int32           `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Value *structpb.Value `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *UpdateSchemaVariableRequest) Reset() {
	*x = UpdateSchemaVariableRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_variables_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSchemaVariableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSchemaVariableRequest) ProtoMessage() {}

func (x *UpdateSchemaVariableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schema_variables_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSchemaVariableRequest.ProtoReflect.Descriptor instead.
func (*UpdateSchemaVariableRequest) Descriptor() ([]byte, []int) {
	return file_schema_variables_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateSchemaVariableRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateSchemaVariableRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateSchemaVariableRequest) GetValue() *structpb.Value {
	if x != nil {
		return x.Value
	}
	return nil
}

type DeleteSchemaVariableRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteSchemaVariableRequest) Reset() {
	*x = DeleteSchemaVariableRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_variables_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSchemaVariableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSchemaVariableRequest) ProtoMessage() {}

func (x *DeleteSchemaVariableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schema_variables_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSchemaVariableRequest.ProtoReflect.Descriptor instead.
func (*DeleteSchemaVariableRequest) Descriptor() ([]byte, []int) {
	return file_schema_variables_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteSchemaVariableRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_schema_variables_proto protoreflect.FileDescriptor

var file_schema_variables_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7c, 0x0a, 0x0e,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x12, 0x2c, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x79, 0x0a, 0x1b, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x7f, 0x0a, 0x1b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x0c, 0x0a, 0x01,
	0x71, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f,
	0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x22, 0x56, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6e, 0x65, 0x78, 0x74,
	0x12, 0x2c, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x56,
	0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x2b,
	0x0a, 0x19, 0x52, 0x65, 0x61, 0x64, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x56, 0x61, 0x72, 0x69,
	0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x6f, 0x0a, 0x1b, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x56, 0x61, 0x72, 0x69, 0x61,
	0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2c,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x2d, 0x0a, 0x1b,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x56, 0x61, 0x72, 0x69,
	0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x32, 0x85, 0x05, 0x0a, 0x16,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x79, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x23,
	0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1e, 0x3a, 0x01, 0x2a, 0x22, 0x19, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67,
	0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65,
	0x73, 0x12, 0x7a, 0x0a, 0x14, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x23, 0x2e, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x56,
	0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1b, 0x12, 0x19, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x77, 0x0a,
	0x12, 0x52, 0x65, 0x61, 0x64, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x56, 0x61, 0x72, 0x69, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x21, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x52, 0x65, 0x61,
	0x64, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x26,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x12, 0x1e, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67,
	0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65,
	0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x7e, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x23,
	0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x29, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x23, 0x3a, 0x01, 0x2a, 0x1a, 0x1e, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67,
	0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65,
	0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x7b, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x23,
	0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x20, 0x2a, 0x1e, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x2f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x77, 0x65, 0x62, 0x69, 0x74, 0x65, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schema_variables_proto_rawDescOnce sync.Once
	file_schema_variables_proto_rawDescData = file_schema_variables_proto_rawDesc
)

func file_schema_variables_proto_rawDescGZIP() []byte {
	file_schema_variables_proto_rawDescOnce.Do(func() {
		file_schema_variables_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_variables_proto_rawDescData)
	})
	return file_schema_variables_proto_rawDescData
}

var file_schema_variables_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_schema_variables_proto_goTypes = []interface{}{
	(*SchemaVariable)(nil),              // 0: engine.SchemaVariable
	(*CreateSchemaVariableRequest)(nil), // 1: engine.CreateSchemaVariableRequest
	(*SearchSchemaVariableRequest)(nil), // 2: engine.SearchSchemaVariableRequest
	(*ListSchemaVariable)(nil),          // 3: engine.ListSchemaVariable
	(*ReadSchemaVariableRequest)(nil),   // 4: engine.ReadSchemaVariableRequest
	(*UpdateSchemaVariableRequest)(nil), // 5: engine.UpdateSchemaVariableRequest
	(*DeleteSchemaVariableRequest)(nil), // 6: engine.DeleteSchemaVariableRequest
	(*structpb.Value)(nil),              // 7: google.protobuf.Value
}
var file_schema_variables_proto_depIdxs = []int32{
	7, // 0: engine.SchemaVariable.value:type_name -> google.protobuf.Value
	7, // 1: engine.CreateSchemaVariableRequest.value:type_name -> google.protobuf.Value
	0, // 2: engine.ListSchemaVariable.items:type_name -> engine.SchemaVariable
	7, // 3: engine.UpdateSchemaVariableRequest.value:type_name -> google.protobuf.Value
	1, // 4: engine.SchemaVariablesService.CreateSchemaVariable:input_type -> engine.CreateSchemaVariableRequest
	2, // 5: engine.SchemaVariablesService.SearchSchemaVariable:input_type -> engine.SearchSchemaVariableRequest
	4, // 6: engine.SchemaVariablesService.ReadSchemaVariable:input_type -> engine.ReadSchemaVariableRequest
	5, // 7: engine.SchemaVariablesService.UpdateSchemaVariable:input_type -> engine.UpdateSchemaVariableRequest
	6, // 8: engine.SchemaVariablesService.DeleteSchemaVariable:input_type -> engine.DeleteSchemaVariableRequest
	0, // 9: engine.SchemaVariablesService.CreateSchemaVariable:output_type -> engine.SchemaVariable
	3, // 10: engine.SchemaVariablesService.SearchSchemaVariable:output_type -> engine.ListSchemaVariable
	0, // 11: engine.SchemaVariablesService.ReadSchemaVariable:output_type -> engine.SchemaVariable
	0, // 12: engine.SchemaVariablesService.UpdateSchemaVariable:output_type -> engine.SchemaVariable
	0, // 13: engine.SchemaVariablesService.DeleteSchemaVariable:output_type -> engine.SchemaVariable
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_schema_variables_proto_init() }
func file_schema_variables_proto_init() {
	if File_schema_variables_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_schema_variables_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SchemaVariable); i {
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
		file_schema_variables_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSchemaVariableRequest); i {
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
		file_schema_variables_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchSchemaVariableRequest); i {
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
		file_schema_variables_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSchemaVariable); i {
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
		file_schema_variables_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadSchemaVariableRequest); i {
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
		file_schema_variables_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSchemaVariableRequest); i {
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
		file_schema_variables_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSchemaVariableRequest); i {
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
			RawDescriptor: file_schema_variables_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_schema_variables_proto_goTypes,
		DependencyIndexes: file_schema_variables_proto_depIdxs,
		MessageInfos:      file_schema_variables_proto_msgTypes,
	}.Build()
	File_schema_variables_proto = out.File
	file_schema_variables_proto_rawDesc = nil
	file_schema_variables_proto_goTypes = nil
	file_schema_variables_proto_depIdxs = nil
}
