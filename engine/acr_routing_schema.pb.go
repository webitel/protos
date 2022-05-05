// Code generated by protoc-gen-go. DO NOT EDIT.
// source: acr_routing_schema.proto

package engine

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type PatchRoutingSchemaRequest struct {
	Id                   int64          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string         `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Type                 string         `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Schema               *_struct.Value `protobuf:"bytes,5,opt,name=schema,proto3" json:"schema,omitempty"`
	Payload              *_struct.Value `protobuf:"bytes,6,opt,name=payload,proto3" json:"payload,omitempty"`
	Debug                bool           `protobuf:"varint,7,opt,name=debug,proto3" json:"debug,omitempty"`
	Fields               []string       `protobuf:"bytes,8,rep,name=fields,proto3" json:"fields,omitempty"`
	Editor               bool           `protobuf:"varint,9,opt,name=editor,proto3" json:"editor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *PatchRoutingSchemaRequest) Reset()         { *m = PatchRoutingSchemaRequest{} }
func (m *PatchRoutingSchemaRequest) String() string { return proto.CompactTextString(m) }
func (*PatchRoutingSchemaRequest) ProtoMessage()    {}
func (*PatchRoutingSchemaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e8e6387a1066a92, []int{0}
}

func (m *PatchRoutingSchemaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PatchRoutingSchemaRequest.Unmarshal(m, b)
}
func (m *PatchRoutingSchemaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PatchRoutingSchemaRequest.Marshal(b, m, deterministic)
}
func (m *PatchRoutingSchemaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PatchRoutingSchemaRequest.Merge(m, src)
}
func (m *PatchRoutingSchemaRequest) XXX_Size() int {
	return xxx_messageInfo_PatchRoutingSchemaRequest.Size(m)
}
func (m *PatchRoutingSchemaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PatchRoutingSchemaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PatchRoutingSchemaRequest proto.InternalMessageInfo

func (m *PatchRoutingSchemaRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PatchRoutingSchemaRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PatchRoutingSchemaRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *PatchRoutingSchemaRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *PatchRoutingSchemaRequest) GetSchema() *_struct.Value {
	if m != nil {
		return m.Schema
	}
	return nil
}

func (m *PatchRoutingSchemaRequest) GetPayload() *_struct.Value {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *PatchRoutingSchemaRequest) GetDebug() bool {
	if m != nil {
		return m.Debug
	}
	return false
}

func (m *PatchRoutingSchemaRequest) GetFields() []string {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *PatchRoutingSchemaRequest) GetEditor() bool {
	if m != nil {
		return m.Editor
	}
	return false
}

type DeleteRoutingSchemaRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DomainId             int64    `protobuf:"varint,2,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRoutingSchemaRequest) Reset()         { *m = DeleteRoutingSchemaRequest{} }
func (m *DeleteRoutingSchemaRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRoutingSchemaRequest) ProtoMessage()    {}
func (*DeleteRoutingSchemaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e8e6387a1066a92, []int{1}
}

func (m *DeleteRoutingSchemaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRoutingSchemaRequest.Unmarshal(m, b)
}
func (m *DeleteRoutingSchemaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRoutingSchemaRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRoutingSchemaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRoutingSchemaRequest.Merge(m, src)
}
func (m *DeleteRoutingSchemaRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRoutingSchemaRequest.Size(m)
}
func (m *DeleteRoutingSchemaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRoutingSchemaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRoutingSchemaRequest proto.InternalMessageInfo

func (m *DeleteRoutingSchemaRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DeleteRoutingSchemaRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

type UpdateRoutingSchemaRequest struct {
	Id                   int64          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string         `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string         `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Type                 string         `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Schema               *_struct.Value `protobuf:"bytes,5,opt,name=schema,proto3" json:"schema,omitempty"`
	Payload              *_struct.Value `protobuf:"bytes,6,opt,name=payload,proto3" json:"payload,omitempty"`
	Debug                bool           `protobuf:"varint,7,opt,name=debug,proto3" json:"debug,omitempty"`
	Editor               bool           `protobuf:"varint,8,opt,name=editor,proto3" json:"editor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UpdateRoutingSchemaRequest) Reset()         { *m = UpdateRoutingSchemaRequest{} }
func (m *UpdateRoutingSchemaRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRoutingSchemaRequest) ProtoMessage()    {}
func (*UpdateRoutingSchemaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e8e6387a1066a92, []int{2}
}

func (m *UpdateRoutingSchemaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRoutingSchemaRequest.Unmarshal(m, b)
}
func (m *UpdateRoutingSchemaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRoutingSchemaRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRoutingSchemaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRoutingSchemaRequest.Merge(m, src)
}
func (m *UpdateRoutingSchemaRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRoutingSchemaRequest.Size(m)
}
func (m *UpdateRoutingSchemaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRoutingSchemaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRoutingSchemaRequest proto.InternalMessageInfo

func (m *UpdateRoutingSchemaRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateRoutingSchemaRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateRoutingSchemaRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *UpdateRoutingSchemaRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *UpdateRoutingSchemaRequest) GetSchema() *_struct.Value {
	if m != nil {
		return m.Schema
	}
	return nil
}

func (m *UpdateRoutingSchemaRequest) GetPayload() *_struct.Value {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *UpdateRoutingSchemaRequest) GetDebug() bool {
	if m != nil {
		return m.Debug
	}
	return false
}

func (m *UpdateRoutingSchemaRequest) GetEditor() bool {
	if m != nil {
		return m.Editor
	}
	return false
}

type ReadRoutingSchemaRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DomainId             int64    `protobuf:"varint,2,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRoutingSchemaRequest) Reset()         { *m = ReadRoutingSchemaRequest{} }
func (m *ReadRoutingSchemaRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRoutingSchemaRequest) ProtoMessage()    {}
func (*ReadRoutingSchemaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e8e6387a1066a92, []int{3}
}

func (m *ReadRoutingSchemaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRoutingSchemaRequest.Unmarshal(m, b)
}
func (m *ReadRoutingSchemaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRoutingSchemaRequest.Marshal(b, m, deterministic)
}
func (m *ReadRoutingSchemaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRoutingSchemaRequest.Merge(m, src)
}
func (m *ReadRoutingSchemaRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRoutingSchemaRequest.Size(m)
}
func (m *ReadRoutingSchemaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRoutingSchemaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRoutingSchemaRequest proto.InternalMessageInfo

func (m *ReadRoutingSchemaRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ReadRoutingSchemaRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

type SearchRoutingSchemaRequest struct {
	Page                 int32    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size                 int32    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Q                    string   `protobuf:"bytes,3,opt,name=q,proto3" json:"q,omitempty"`
	Sort                 string   `protobuf:"bytes,4,opt,name=sort,proto3" json:"sort,omitempty"`
	Fields               []string `protobuf:"bytes,5,rep,name=fields,proto3" json:"fields,omitempty"`
	Id                   []uint32 `protobuf:"varint,6,rep,packed,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	Type                 []string `protobuf:"bytes,8,rep,name=type,proto3" json:"type,omitempty"`
	Editor               bool     `protobuf:"varint,9,opt,name=editor,proto3" json:"editor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchRoutingSchemaRequest) Reset()         { *m = SearchRoutingSchemaRequest{} }
func (m *SearchRoutingSchemaRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRoutingSchemaRequest) ProtoMessage()    {}
func (*SearchRoutingSchemaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e8e6387a1066a92, []int{4}
}

func (m *SearchRoutingSchemaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRoutingSchemaRequest.Unmarshal(m, b)
}
func (m *SearchRoutingSchemaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRoutingSchemaRequest.Marshal(b, m, deterministic)
}
func (m *SearchRoutingSchemaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRoutingSchemaRequest.Merge(m, src)
}
func (m *SearchRoutingSchemaRequest) XXX_Size() int {
	return xxx_messageInfo_SearchRoutingSchemaRequest.Size(m)
}
func (m *SearchRoutingSchemaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRoutingSchemaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRoutingSchemaRequest proto.InternalMessageInfo

func (m *SearchRoutingSchemaRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *SearchRoutingSchemaRequest) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *SearchRoutingSchemaRequest) GetQ() string {
	if m != nil {
		return m.Q
	}
	return ""
}

func (m *SearchRoutingSchemaRequest) GetSort() string {
	if m != nil {
		return m.Sort
	}
	return ""
}

func (m *SearchRoutingSchemaRequest) GetFields() []string {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *SearchRoutingSchemaRequest) GetId() []uint32 {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *SearchRoutingSchemaRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SearchRoutingSchemaRequest) GetType() []string {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *SearchRoutingSchemaRequest) GetEditor() bool {
	if m != nil {
		return m.Editor
	}
	return false
}

type CreateRoutingSchemaRequest struct {
	Name                 string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string         `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Type                 string         `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Schema               *_struct.Value `protobuf:"bytes,4,opt,name=schema,proto3" json:"schema,omitempty"`
	Payload              *_struct.Value `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
	Debug                bool           `protobuf:"varint,6,opt,name=debug,proto3" json:"debug,omitempty"`
	Editor               bool           `protobuf:"varint,7,opt,name=editor,proto3" json:"editor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreateRoutingSchemaRequest) Reset()         { *m = CreateRoutingSchemaRequest{} }
func (m *CreateRoutingSchemaRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRoutingSchemaRequest) ProtoMessage()    {}
func (*CreateRoutingSchemaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e8e6387a1066a92, []int{5}
}

func (m *CreateRoutingSchemaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRoutingSchemaRequest.Unmarshal(m, b)
}
func (m *CreateRoutingSchemaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRoutingSchemaRequest.Marshal(b, m, deterministic)
}
func (m *CreateRoutingSchemaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRoutingSchemaRequest.Merge(m, src)
}
func (m *CreateRoutingSchemaRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRoutingSchemaRequest.Size(m)
}
func (m *CreateRoutingSchemaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRoutingSchemaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRoutingSchemaRequest proto.InternalMessageInfo

func (m *CreateRoutingSchemaRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRoutingSchemaRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreateRoutingSchemaRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *CreateRoutingSchemaRequest) GetSchema() *_struct.Value {
	if m != nil {
		return m.Schema
	}
	return nil
}

func (m *CreateRoutingSchemaRequest) GetPayload() *_struct.Value {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *CreateRoutingSchemaRequest) GetDebug() bool {
	if m != nil {
		return m.Debug
	}
	return false
}

func (m *CreateRoutingSchemaRequest) GetEditor() bool {
	if m != nil {
		return m.Editor
	}
	return false
}

type RoutingSchema struct {
	Id                   int64          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            int64          `protobuf:"varint,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	CreatedBy            *Lookup        `protobuf:"bytes,3,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	UpdatedAt            int64          `protobuf:"varint,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	UpdatedBy            *Lookup        `protobuf:"bytes,5,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
	Name                 string         `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Description          string         `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	Type                 string         `protobuf:"bytes,8,opt,name=type,proto3" json:"type,omitempty"`
	Schema               *_struct.Value `protobuf:"bytes,9,opt,name=schema,proto3" json:"schema,omitempty"`
	Payload              *_struct.Value `protobuf:"bytes,10,opt,name=payload,proto3" json:"payload,omitempty"`
	Debug                bool           `protobuf:"varint,11,opt,name=debug,proto3" json:"debug,omitempty"`
	Editor               bool           `protobuf:"varint,12,opt,name=editor,proto3" json:"editor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *RoutingSchema) Reset()         { *m = RoutingSchema{} }
func (m *RoutingSchema) String() string { return proto.CompactTextString(m) }
func (*RoutingSchema) ProtoMessage()    {}
func (*RoutingSchema) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e8e6387a1066a92, []int{6}
}

func (m *RoutingSchema) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoutingSchema.Unmarshal(m, b)
}
func (m *RoutingSchema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoutingSchema.Marshal(b, m, deterministic)
}
func (m *RoutingSchema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoutingSchema.Merge(m, src)
}
func (m *RoutingSchema) XXX_Size() int {
	return xxx_messageInfo_RoutingSchema.Size(m)
}
func (m *RoutingSchema) XXX_DiscardUnknown() {
	xxx_messageInfo_RoutingSchema.DiscardUnknown(m)
}

var xxx_messageInfo_RoutingSchema proto.InternalMessageInfo

func (m *RoutingSchema) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RoutingSchema) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *RoutingSchema) GetCreatedBy() *Lookup {
	if m != nil {
		return m.CreatedBy
	}
	return nil
}

func (m *RoutingSchema) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *RoutingSchema) GetUpdatedBy() *Lookup {
	if m != nil {
		return m.UpdatedBy
	}
	return nil
}

func (m *RoutingSchema) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RoutingSchema) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *RoutingSchema) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *RoutingSchema) GetSchema() *_struct.Value {
	if m != nil {
		return m.Schema
	}
	return nil
}

func (m *RoutingSchema) GetPayload() *_struct.Value {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *RoutingSchema) GetDebug() bool {
	if m != nil {
		return m.Debug
	}
	return false
}

func (m *RoutingSchema) GetEditor() bool {
	if m != nil {
		return m.Editor
	}
	return false
}

type ListRoutingSchema struct {
	Next                 bool             `protobuf:"varint,1,opt,name=next,proto3" json:"next,omitempty"`
	Items                []*RoutingSchema `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ListRoutingSchema) Reset()         { *m = ListRoutingSchema{} }
func (m *ListRoutingSchema) String() string { return proto.CompactTextString(m) }
func (*ListRoutingSchema) ProtoMessage()    {}
func (*ListRoutingSchema) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e8e6387a1066a92, []int{7}
}

func (m *ListRoutingSchema) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRoutingSchema.Unmarshal(m, b)
}
func (m *ListRoutingSchema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRoutingSchema.Marshal(b, m, deterministic)
}
func (m *ListRoutingSchema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRoutingSchema.Merge(m, src)
}
func (m *ListRoutingSchema) XXX_Size() int {
	return xxx_messageInfo_ListRoutingSchema.Size(m)
}
func (m *ListRoutingSchema) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRoutingSchema.DiscardUnknown(m)
}

var xxx_messageInfo_ListRoutingSchema proto.InternalMessageInfo

func (m *ListRoutingSchema) GetNext() bool {
	if m != nil {
		return m.Next
	}
	return false
}

func (m *ListRoutingSchema) GetItems() []*RoutingSchema {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*PatchRoutingSchemaRequest)(nil), "engine.PatchRoutingSchemaRequest")
	proto.RegisterType((*DeleteRoutingSchemaRequest)(nil), "engine.DeleteRoutingSchemaRequest")
	proto.RegisterType((*UpdateRoutingSchemaRequest)(nil), "engine.UpdateRoutingSchemaRequest")
	proto.RegisterType((*ReadRoutingSchemaRequest)(nil), "engine.ReadRoutingSchemaRequest")
	proto.RegisterType((*SearchRoutingSchemaRequest)(nil), "engine.SearchRoutingSchemaRequest")
	proto.RegisterType((*CreateRoutingSchemaRequest)(nil), "engine.CreateRoutingSchemaRequest")
	proto.RegisterType((*RoutingSchema)(nil), "engine.RoutingSchema")
	proto.RegisterType((*ListRoutingSchema)(nil), "engine.ListRoutingSchema")
}

func init() { proto.RegisterFile("acr_routing_schema.proto", fileDescriptor_7e8e6387a1066a92) }

var fileDescriptor_7e8e6387a1066a92 = []byte{
	// 766 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x56, 0xc1, 0x6e, 0xd3, 0x4a,
	0x14, 0x95, 0xed, 0xd8, 0x49, 0x26, 0x6d, 0x9f, 0x3a, 0xc9, 0x2b, 0xae, 0x29, 0xc2, 0x78, 0x15,
	0x15, 0x61, 0xa3, 0xb0, 0x63, 0x47, 0x41, 0x42, 0x95, 0xba, 0x40, 0x2e, 0xb0, 0x60, 0x13, 0x4d,
	0xec, 0x69, 0x3a, 0x90, 0x78, 0x1c, 0x7b, 0x0c, 0x04, 0xc4, 0x86, 0x05, 0x3f, 0x80, 0x58, 0xf2,
	0x15, 0xfc, 0x07, 0x1b, 0x7e, 0x81, 0x6f, 0x60, 0x8d, 0x3c, 0x63, 0x27, 0x75, 0xe2, 0xa1, 0x8d,
	0xd8, 0xb1, 0x9b, 0xb9, 0xf7, 0xfa, 0x1e, 0xdf, 0x73, 0x8e, 0x67, 0x0c, 0x4c, 0x14, 0x24, 0xc3,
	0x84, 0x66, 0x8c, 0x44, 0xe3, 0x61, 0x1a, 0x9c, 0xe3, 0x29, 0x72, 0xe3, 0x84, 0x32, 0x0a, 0x0d,
	0x1c, 0x8d, 0x49, 0x84, 0xad, 0x4e, 0x40, 0xa3, 0x94, 0x89, 0xa0, 0x75, 0x30, 0xa6, 0x74, 0x3c,
	0xc1, 0x1e, 0xdf, 0x8d, 0xb2, 0x33, 0x2f, 0x65, 0x49, 0x16, 0xac, 0x66, 0x51, 0x4c, 0x3c, 0x14,
	0x45, 0x94, 0x21, 0x46, 0x68, 0x94, 0x8a, 0xac, 0xf3, 0x55, 0x05, 0xfb, 0x4f, 0x10, 0x0b, 0xce,
	0x7d, 0x01, 0x77, 0xca, 0xd1, 0x7c, 0x3c, 0xcb, 0x70, 0xca, 0xe0, 0x0e, 0x50, 0x49, 0x68, 0x2a,
	0xb6, 0xd2, 0xd7, 0x7c, 0x95, 0x84, 0x10, 0x82, 0x46, 0x84, 0xa6, 0xd8, 0x54, 0x6d, 0xa5, 0xdf,
	0xf6, 0xf9, 0x1a, 0xda, 0xa0, 0x13, 0xe2, 0x34, 0x48, 0x48, 0x9c, 0xf7, 0x35, 0x35, 0x9e, 0xba,
	0x18, 0xca, 0x9f, 0x62, 0xf3, 0x18, 0x9b, 0x0d, 0xf1, 0x54, 0xbe, 0x86, 0x2e, 0x30, 0xc4, 0x60,
	0xa6, 0x6e, 0x2b, 0xfd, 0xce, 0x60, 0xcf, 0x15, 0xaf, 0xe9, 0x96, 0x43, 0xb8, 0xcf, 0xd1, 0x24,
	0xc3, 0x7e, 0x51, 0x05, 0xef, 0x82, 0x66, 0x8c, 0xe6, 0x13, 0x8a, 0x42, 0xd3, 0xf8, 0xe3, 0x03,
	0x65, 0x19, 0xec, 0x01, 0x3d, 0xc4, 0xa3, 0x6c, 0x6c, 0x36, 0x6d, 0xa5, 0xdf, 0xf2, 0xc5, 0x06,
	0xee, 0x01, 0xe3, 0x8c, 0xe0, 0x49, 0x98, 0x9a, 0x2d, 0x5b, 0xeb, 0xb7, 0xfd, 0x62, 0x97, 0xc7,
	0x71, 0x48, 0x18, 0x4d, 0xcc, 0x36, 0x2f, 0x2f, 0x76, 0xce, 0x31, 0xb0, 0x1e, 0xe1, 0x09, 0x66,
	0xf8, 0x4a, 0xfc, 0x5c, 0x07, 0xed, 0x90, 0x4e, 0x11, 0x89, 0x86, 0x24, 0xe4, 0x24, 0x69, 0x7e,
	0x4b, 0x04, 0x8e, 0x43, 0xe7, 0x93, 0x0a, 0xac, 0x67, 0x71, 0x88, 0xae, 0xd8, 0xeb, 0x9f, 0xe0,
	0xba, 0xe0, 0xb4, 0x55, 0xe1, 0xf4, 0x31, 0x30, 0x7d, 0x8c, 0xc2, 0xbf, 0x67, 0xf4, 0xbb, 0x02,
	0xac, 0x53, 0x8c, 0x12, 0x89, 0x7b, 0x21, 0x68, 0xc4, 0x68, 0x8c, 0x79, 0x37, 0xdd, 0xe7, 0xeb,
	0x3c, 0x96, 0x92, 0x77, 0x82, 0x55, 0xdd, 0xe7, 0x6b, 0xb8, 0x05, 0x94, 0x59, 0xc1, 0xa5, 0x32,
	0xe3, 0x15, 0x34, 0x61, 0x25, 0x83, 0xf9, 0xfa, 0x82, 0x6b, 0xf4, 0x8a, 0x6b, 0xc4, 0xdb, 0x1a,
	0xb6, 0xd6, 0xdf, 0xae, 0x68, 0xd6, 0xbc, 0xa0, 0x59, 0xa9, 0x88, 0xf0, 0x9b, 0x50, 0x44, 0xe6,
	0xb6, 0x5f, 0x0a, 0xb0, 0x1e, 0x26, 0x58, 0x66, 0x91, 0xb2, 0xbd, 0x22, 0xb7, 0x84, 0x2a, 0xb7,
	0x84, 0x56, 0x6b, 0x89, 0xc6, 0xa6, 0x96, 0xd0, 0x37, 0xb4, 0x84, 0x51, 0x6f, 0x89, 0x66, 0x65,
	0xf0, 0x2f, 0x1a, 0xd8, 0xae, 0x8c, 0xbc, 0x66, 0x84, 0x1b, 0x00, 0x04, 0x9c, 0x99, 0x70, 0x88,
	0x58, 0xe1, 0x84, 0x76, 0x11, 0x79, 0xc0, 0xe0, 0x9d, 0x65, 0x7a, 0x34, 0xe7, 0xa3, 0x76, 0x06,
	0x3b, 0xae, 0x38, 0x2d, 0xdd, 0x13, 0x4a, 0x5f, 0x65, 0xf1, 0xa2, 0xfc, 0x68, 0x9e, 0x77, 0xcb,
	0xf8, 0xa7, 0xc8, 0xbb, 0x35, 0x44, 0xb7, 0x22, 0x22, 0xba, 0x95, 0xe9, 0xd1, 0xbc, 0x98, 0x78,
	0xad, 0x5b, 0x51, 0x71, 0x34, 0x5f, 0xe8, 0x62, 0xc8, 0x75, 0x69, 0xca, 0x75, 0x69, 0xd5, 0xea,
	0xd2, 0xde, 0x54, 0x17, 0xb0, 0xa1, 0x2e, 0x9d, 0x7a, 0x5d, 0xb6, 0x2a, 0xba, 0x3c, 0x05, 0xbb,
	0x27, 0x24, 0x65, 0x55, 0x69, 0xf2, 0x71, 0xf1, 0x5b, 0xc6, 0xc5, 0x69, 0xf9, 0x7c, 0x0d, 0x6f,
	0x03, 0x9d, 0x30, 0x3c, 0x4d, 0x4d, 0xd5, 0xd6, 0xfa, 0x9d, 0xc1, 0xff, 0x25, 0x59, 0x55, 0x1f,
	0x8b, 0x9a, 0xc1, 0x37, 0x1d, 0xf4, 0x2a, 0x89, 0x53, 0x9c, 0xbc, 0x26, 0x01, 0x86, 0x13, 0xd0,
	0xad, 0xb1, 0x3f, 0x74, 0xca, 0x6e, 0xf2, 0x6f, 0xc3, 0xaa, 0x47, 0x74, 0xac, 0x8f, 0x3f, 0x7e,
	0x7e, 0x56, 0x7b, 0xce, 0x7f, 0x5e, 0x71, 0x9f, 0x7a, 0x82, 0xb9, 0xfb, 0xca, 0x21, 0x9c, 0x82,
	0x6e, 0xcd, 0xe9, 0xb1, 0x44, 0x93, 0x1f, 0x2d, 0xd6, 0xfe, 0xc2, 0x0c, 0xab, 0xec, 0x38, 0xd7,
	0x38, 0xe2, 0x2e, 0x5c, 0x45, 0x84, 0x2f, 0xc1, 0xee, 0xda, 0xb1, 0x07, 0xed, 0xc5, 0x6b, 0x4b,
	0x4e, 0x44, 0xd9, 0x60, 0x07, 0x1c, 0x66, 0x0f, 0xf6, 0x56, 0x60, 0xbc, 0xf7, 0x24, 0xfc, 0x00,
	0x67, 0xa0, 0x5b, 0x73, 0xd5, 0x2c, 0x47, 0x93, 0xdf, 0x43, 0x32, 0xbc, 0x9b, 0x1c, 0x6f, 0xdf,
	0xaa, 0xc5, 0xcb, 0xd9, 0xa4, 0x00, 0xae, 0xff, 0x48, 0xc0, 0x5b, 0x65, 0x37, 0xe9, 0x4f, 0xc6,
	0x25, 0x80, 0x03, 0x29, 0x60, 0x04, 0xba, 0x35, 0x57, 0xf3, 0x72, 0x46, 0xf9, 0xbd, 0x7d, 0x09,
	0xa7, 0x87, 0xb5, 0x90, 0x47, 0xce, 0x0b, 0x7b, 0x4c, 0xd8, 0x79, 0x36, 0x72, 0x03, 0x3a, 0xf5,
	0xde, 0xe0, 0x11, 0x61, 0x78, 0x22, 0x7e, 0xba, 0x52, 0x4f, 0xf4, 0x1b, 0x19, 0x7c, 0x7b, 0xef,
	0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2e, 0x21, 0xc0, 0xaf, 0xc2, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RoutingSchemaServiceClient is the client API for RoutingSchemaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RoutingSchemaServiceClient interface {
	// Create RoutingSchema
	CreateRoutingSchema(ctx context.Context, in *CreateRoutingSchemaRequest, opts ...grpc.CallOption) (*RoutingSchema, error)
	// List RoutingSchema
	SearchRoutingSchema(ctx context.Context, in *SearchRoutingSchemaRequest, opts ...grpc.CallOption) (*ListRoutingSchema, error)
	// RoutingSchema item
	ReadRoutingSchema(ctx context.Context, in *ReadRoutingSchemaRequest, opts ...grpc.CallOption) (*RoutingSchema, error)
	// Update RoutingSchema
	UpdateRoutingSchema(ctx context.Context, in *UpdateRoutingSchemaRequest, opts ...grpc.CallOption) (*RoutingSchema, error)
	// Patch RoutingSchema
	PatchRoutingSchema(ctx context.Context, in *PatchRoutingSchemaRequest, opts ...grpc.CallOption) (*RoutingSchema, error)
	// Remove RoutingSchema
	DeleteRoutingSchema(ctx context.Context, in *DeleteRoutingSchemaRequest, opts ...grpc.CallOption) (*RoutingSchema, error)
}

type routingSchemaServiceClient struct {
	cc *grpc.ClientConn
}

func NewRoutingSchemaServiceClient(cc *grpc.ClientConn) RoutingSchemaServiceClient {
	return &routingSchemaServiceClient{cc}
}

func (c *routingSchemaServiceClient) CreateRoutingSchema(ctx context.Context, in *CreateRoutingSchemaRequest, opts ...grpc.CallOption) (*RoutingSchema, error) {
	out := new(RoutingSchema)
	err := c.cc.Invoke(ctx, "/engine.RoutingSchemaService/CreateRoutingSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routingSchemaServiceClient) SearchRoutingSchema(ctx context.Context, in *SearchRoutingSchemaRequest, opts ...grpc.CallOption) (*ListRoutingSchema, error) {
	out := new(ListRoutingSchema)
	err := c.cc.Invoke(ctx, "/engine.RoutingSchemaService/SearchRoutingSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routingSchemaServiceClient) ReadRoutingSchema(ctx context.Context, in *ReadRoutingSchemaRequest, opts ...grpc.CallOption) (*RoutingSchema, error) {
	out := new(RoutingSchema)
	err := c.cc.Invoke(ctx, "/engine.RoutingSchemaService/ReadRoutingSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routingSchemaServiceClient) UpdateRoutingSchema(ctx context.Context, in *UpdateRoutingSchemaRequest, opts ...grpc.CallOption) (*RoutingSchema, error) {
	out := new(RoutingSchema)
	err := c.cc.Invoke(ctx, "/engine.RoutingSchemaService/UpdateRoutingSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routingSchemaServiceClient) PatchRoutingSchema(ctx context.Context, in *PatchRoutingSchemaRequest, opts ...grpc.CallOption) (*RoutingSchema, error) {
	out := new(RoutingSchema)
	err := c.cc.Invoke(ctx, "/engine.RoutingSchemaService/PatchRoutingSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routingSchemaServiceClient) DeleteRoutingSchema(ctx context.Context, in *DeleteRoutingSchemaRequest, opts ...grpc.CallOption) (*RoutingSchema, error) {
	out := new(RoutingSchema)
	err := c.cc.Invoke(ctx, "/engine.RoutingSchemaService/DeleteRoutingSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoutingSchemaServiceServer is the server API for RoutingSchemaService service.
type RoutingSchemaServiceServer interface {
	// Create RoutingSchema
	CreateRoutingSchema(context.Context, *CreateRoutingSchemaRequest) (*RoutingSchema, error)
	// List RoutingSchema
	SearchRoutingSchema(context.Context, *SearchRoutingSchemaRequest) (*ListRoutingSchema, error)
	// RoutingSchema item
	ReadRoutingSchema(context.Context, *ReadRoutingSchemaRequest) (*RoutingSchema, error)
	// Update RoutingSchema
	UpdateRoutingSchema(context.Context, *UpdateRoutingSchemaRequest) (*RoutingSchema, error)
	// Patch RoutingSchema
	PatchRoutingSchema(context.Context, *PatchRoutingSchemaRequest) (*RoutingSchema, error)
	// Remove RoutingSchema
	DeleteRoutingSchema(context.Context, *DeleteRoutingSchemaRequest) (*RoutingSchema, error)
}

// UnimplementedRoutingSchemaServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRoutingSchemaServiceServer struct {
}

func (*UnimplementedRoutingSchemaServiceServer) CreateRoutingSchema(ctx context.Context, req *CreateRoutingSchemaRequest) (*RoutingSchema, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoutingSchema not implemented")
}
func (*UnimplementedRoutingSchemaServiceServer) SearchRoutingSchema(ctx context.Context, req *SearchRoutingSchemaRequest) (*ListRoutingSchema, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchRoutingSchema not implemented")
}
func (*UnimplementedRoutingSchemaServiceServer) ReadRoutingSchema(ctx context.Context, req *ReadRoutingSchemaRequest) (*RoutingSchema, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadRoutingSchema not implemented")
}
func (*UnimplementedRoutingSchemaServiceServer) UpdateRoutingSchema(ctx context.Context, req *UpdateRoutingSchemaRequest) (*RoutingSchema, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRoutingSchema not implemented")
}
func (*UnimplementedRoutingSchemaServiceServer) PatchRoutingSchema(ctx context.Context, req *PatchRoutingSchemaRequest) (*RoutingSchema, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchRoutingSchema not implemented")
}
func (*UnimplementedRoutingSchemaServiceServer) DeleteRoutingSchema(ctx context.Context, req *DeleteRoutingSchemaRequest) (*RoutingSchema, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRoutingSchema not implemented")
}

func RegisterRoutingSchemaServiceServer(s *grpc.Server, srv RoutingSchemaServiceServer) {
	s.RegisterService(&_RoutingSchemaService_serviceDesc, srv)
}

func _RoutingSchemaService_CreateRoutingSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoutingSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutingSchemaServiceServer).CreateRoutingSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.RoutingSchemaService/CreateRoutingSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutingSchemaServiceServer).CreateRoutingSchema(ctx, req.(*CreateRoutingSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoutingSchemaService_SearchRoutingSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRoutingSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutingSchemaServiceServer).SearchRoutingSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.RoutingSchemaService/SearchRoutingSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutingSchemaServiceServer).SearchRoutingSchema(ctx, req.(*SearchRoutingSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoutingSchemaService_ReadRoutingSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRoutingSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutingSchemaServiceServer).ReadRoutingSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.RoutingSchemaService/ReadRoutingSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutingSchemaServiceServer).ReadRoutingSchema(ctx, req.(*ReadRoutingSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoutingSchemaService_UpdateRoutingSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoutingSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutingSchemaServiceServer).UpdateRoutingSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.RoutingSchemaService/UpdateRoutingSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutingSchemaServiceServer).UpdateRoutingSchema(ctx, req.(*UpdateRoutingSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoutingSchemaService_PatchRoutingSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchRoutingSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutingSchemaServiceServer).PatchRoutingSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.RoutingSchemaService/PatchRoutingSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutingSchemaServiceServer).PatchRoutingSchema(ctx, req.(*PatchRoutingSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoutingSchemaService_DeleteRoutingSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoutingSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutingSchemaServiceServer).DeleteRoutingSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.RoutingSchemaService/DeleteRoutingSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutingSchemaServiceServer).DeleteRoutingSchema(ctx, req.(*DeleteRoutingSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RoutingSchemaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "engine.RoutingSchemaService",
	HandlerType: (*RoutingSchemaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRoutingSchema",
			Handler:    _RoutingSchemaService_CreateRoutingSchema_Handler,
		},
		{
			MethodName: "SearchRoutingSchema",
			Handler:    _RoutingSchemaService_SearchRoutingSchema_Handler,
		},
		{
			MethodName: "ReadRoutingSchema",
			Handler:    _RoutingSchemaService_ReadRoutingSchema_Handler,
		},
		{
			MethodName: "UpdateRoutingSchema",
			Handler:    _RoutingSchemaService_UpdateRoutingSchema_Handler,
		},
		{
			MethodName: "PatchRoutingSchema",
			Handler:    _RoutingSchemaService_PatchRoutingSchema_Handler,
		},
		{
			MethodName: "DeleteRoutingSchema",
			Handler:    _RoutingSchemaService_DeleteRoutingSchema_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "acr_routing_schema.proto",
}
