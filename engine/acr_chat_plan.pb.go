// Code generated by protoc-gen-go. DO NOT EDIT.
// source: acr_chat_plan.proto

package engine

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type ChatPlan struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Schema               *Lookup  `protobuf:"bytes,4,opt,name=schema,proto3" json:"schema,omitempty"`
	Enabled              bool     `protobuf:"varint,5,opt,name=enabled,proto3" json:"enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChatPlan) Reset()         { *m = ChatPlan{} }
func (m *ChatPlan) String() string { return proto.CompactTextString(m) }
func (*ChatPlan) ProtoMessage()    {}
func (*ChatPlan) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5ca83868761bb59, []int{0}
}

func (m *ChatPlan) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatPlan.Unmarshal(m, b)
}
func (m *ChatPlan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatPlan.Marshal(b, m, deterministic)
}
func (m *ChatPlan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatPlan.Merge(m, src)
}
func (m *ChatPlan) XXX_Size() int {
	return xxx_messageInfo_ChatPlan.Size(m)
}
func (m *ChatPlan) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatPlan.DiscardUnknown(m)
}

var xxx_messageInfo_ChatPlan proto.InternalMessageInfo

func (m *ChatPlan) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ChatPlan) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ChatPlan) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ChatPlan) GetSchema() *Lookup {
	if m != nil {
		return m.Schema
	}
	return nil
}

func (m *ChatPlan) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

type CreateChatPlanRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Schema               *Lookup  `protobuf:"bytes,3,opt,name=schema,proto3" json:"schema,omitempty"`
	Enabled              bool     `protobuf:"varint,4,opt,name=enabled,proto3" json:"enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateChatPlanRequest) Reset()         { *m = CreateChatPlanRequest{} }
func (m *CreateChatPlanRequest) String() string { return proto.CompactTextString(m) }
func (*CreateChatPlanRequest) ProtoMessage()    {}
func (*CreateChatPlanRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5ca83868761bb59, []int{1}
}

func (m *CreateChatPlanRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateChatPlanRequest.Unmarshal(m, b)
}
func (m *CreateChatPlanRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateChatPlanRequest.Marshal(b, m, deterministic)
}
func (m *CreateChatPlanRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateChatPlanRequest.Merge(m, src)
}
func (m *CreateChatPlanRequest) XXX_Size() int {
	return xxx_messageInfo_CreateChatPlanRequest.Size(m)
}
func (m *CreateChatPlanRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateChatPlanRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateChatPlanRequest proto.InternalMessageInfo

func (m *CreateChatPlanRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateChatPlanRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreateChatPlanRequest) GetSchema() *Lookup {
	if m != nil {
		return m.Schema
	}
	return nil
}

func (m *CreateChatPlanRequest) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

type SearchChatPlanRequest struct {
	Page                 int32    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size                 int32    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Q                    string   `protobuf:"bytes,3,opt,name=q,proto3" json:"q,omitempty"`
	Sort                 string   `protobuf:"bytes,4,opt,name=sort,proto3" json:"sort,omitempty"`
	Fields               []string `protobuf:"bytes,5,rep,name=fields,proto3" json:"fields,omitempty"`
	Id                   []int32  `protobuf:"varint,6,rep,packed,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	Enabled              bool     `protobuf:"varint,8,opt,name=enabled,proto3" json:"enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchChatPlanRequest) Reset()         { *m = SearchChatPlanRequest{} }
func (m *SearchChatPlanRequest) String() string { return proto.CompactTextString(m) }
func (*SearchChatPlanRequest) ProtoMessage()    {}
func (*SearchChatPlanRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5ca83868761bb59, []int{2}
}

func (m *SearchChatPlanRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchChatPlanRequest.Unmarshal(m, b)
}
func (m *SearchChatPlanRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchChatPlanRequest.Marshal(b, m, deterministic)
}
func (m *SearchChatPlanRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchChatPlanRequest.Merge(m, src)
}
func (m *SearchChatPlanRequest) XXX_Size() int {
	return xxx_messageInfo_SearchChatPlanRequest.Size(m)
}
func (m *SearchChatPlanRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchChatPlanRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchChatPlanRequest proto.InternalMessageInfo

func (m *SearchChatPlanRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *SearchChatPlanRequest) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *SearchChatPlanRequest) GetQ() string {
	if m != nil {
		return m.Q
	}
	return ""
}

func (m *SearchChatPlanRequest) GetSort() string {
	if m != nil {
		return m.Sort
	}
	return ""
}

func (m *SearchChatPlanRequest) GetFields() []string {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *SearchChatPlanRequest) GetId() []int32 {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *SearchChatPlanRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SearchChatPlanRequest) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

type ListChatPlan struct {
	Next                 bool        `protobuf:"varint,1,opt,name=next,proto3" json:"next,omitempty"`
	Items                []*ChatPlan `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ListChatPlan) Reset()         { *m = ListChatPlan{} }
func (m *ListChatPlan) String() string { return proto.CompactTextString(m) }
func (*ListChatPlan) ProtoMessage()    {}
func (*ListChatPlan) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5ca83868761bb59, []int{3}
}

func (m *ListChatPlan) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListChatPlan.Unmarshal(m, b)
}
func (m *ListChatPlan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListChatPlan.Marshal(b, m, deterministic)
}
func (m *ListChatPlan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListChatPlan.Merge(m, src)
}
func (m *ListChatPlan) XXX_Size() int {
	return xxx_messageInfo_ListChatPlan.Size(m)
}
func (m *ListChatPlan) XXX_DiscardUnknown() {
	xxx_messageInfo_ListChatPlan.DiscardUnknown(m)
}

var xxx_messageInfo_ListChatPlan proto.InternalMessageInfo

func (m *ListChatPlan) GetNext() bool {
	if m != nil {
		return m.Next
	}
	return false
}

func (m *ListChatPlan) GetItems() []*ChatPlan {
	if m != nil {
		return m.Items
	}
	return nil
}

type ReadChatPlanRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadChatPlanRequest) Reset()         { *m = ReadChatPlanRequest{} }
func (m *ReadChatPlanRequest) String() string { return proto.CompactTextString(m) }
func (*ReadChatPlanRequest) ProtoMessage()    {}
func (*ReadChatPlanRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5ca83868761bb59, []int{4}
}

func (m *ReadChatPlanRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadChatPlanRequest.Unmarshal(m, b)
}
func (m *ReadChatPlanRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadChatPlanRequest.Marshal(b, m, deterministic)
}
func (m *ReadChatPlanRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadChatPlanRequest.Merge(m, src)
}
func (m *ReadChatPlanRequest) XXX_Size() int {
	return xxx_messageInfo_ReadChatPlanRequest.Size(m)
}
func (m *ReadChatPlanRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadChatPlanRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadChatPlanRequest proto.InternalMessageInfo

func (m *ReadChatPlanRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type UpdateChatPlanRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Schema               *Lookup  `protobuf:"bytes,4,opt,name=schema,proto3" json:"schema,omitempty"`
	Enabled              bool     `protobuf:"varint,5,opt,name=enabled,proto3" json:"enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateChatPlanRequest) Reset()         { *m = UpdateChatPlanRequest{} }
func (m *UpdateChatPlanRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateChatPlanRequest) ProtoMessage()    {}
func (*UpdateChatPlanRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5ca83868761bb59, []int{5}
}

func (m *UpdateChatPlanRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateChatPlanRequest.Unmarshal(m, b)
}
func (m *UpdateChatPlanRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateChatPlanRequest.Marshal(b, m, deterministic)
}
func (m *UpdateChatPlanRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateChatPlanRequest.Merge(m, src)
}
func (m *UpdateChatPlanRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateChatPlanRequest.Size(m)
}
func (m *UpdateChatPlanRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateChatPlanRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateChatPlanRequest proto.InternalMessageInfo

func (m *UpdateChatPlanRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateChatPlanRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateChatPlanRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *UpdateChatPlanRequest) GetSchema() *Lookup {
	if m != nil {
		return m.Schema
	}
	return nil
}

func (m *UpdateChatPlanRequest) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

type PatchChatPlanRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Schema               *Lookup  `protobuf:"bytes,4,opt,name=schema,proto3" json:"schema,omitempty"`
	Enabled              bool     `protobuf:"varint,5,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Fields               []string `protobuf:"bytes,6,rep,name=fields,proto3" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PatchChatPlanRequest) Reset()         { *m = PatchChatPlanRequest{} }
func (m *PatchChatPlanRequest) String() string { return proto.CompactTextString(m) }
func (*PatchChatPlanRequest) ProtoMessage()    {}
func (*PatchChatPlanRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5ca83868761bb59, []int{6}
}

func (m *PatchChatPlanRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PatchChatPlanRequest.Unmarshal(m, b)
}
func (m *PatchChatPlanRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PatchChatPlanRequest.Marshal(b, m, deterministic)
}
func (m *PatchChatPlanRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PatchChatPlanRequest.Merge(m, src)
}
func (m *PatchChatPlanRequest) XXX_Size() int {
	return xxx_messageInfo_PatchChatPlanRequest.Size(m)
}
func (m *PatchChatPlanRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PatchChatPlanRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PatchChatPlanRequest proto.InternalMessageInfo

func (m *PatchChatPlanRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PatchChatPlanRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PatchChatPlanRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *PatchChatPlanRequest) GetSchema() *Lookup {
	if m != nil {
		return m.Schema
	}
	return nil
}

func (m *PatchChatPlanRequest) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *PatchChatPlanRequest) GetFields() []string {
	if m != nil {
		return m.Fields
	}
	return nil
}

type DeleteChatPlanRequest struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteChatPlanRequest) Reset()         { *m = DeleteChatPlanRequest{} }
func (m *DeleteChatPlanRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteChatPlanRequest) ProtoMessage()    {}
func (*DeleteChatPlanRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5ca83868761bb59, []int{7}
}

func (m *DeleteChatPlanRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteChatPlanRequest.Unmarshal(m, b)
}
func (m *DeleteChatPlanRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteChatPlanRequest.Marshal(b, m, deterministic)
}
func (m *DeleteChatPlanRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteChatPlanRequest.Merge(m, src)
}
func (m *DeleteChatPlanRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteChatPlanRequest.Size(m)
}
func (m *DeleteChatPlanRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteChatPlanRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteChatPlanRequest proto.InternalMessageInfo

func (m *DeleteChatPlanRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*ChatPlan)(nil), "engine.ChatPlan")
	proto.RegisterType((*CreateChatPlanRequest)(nil), "engine.CreateChatPlanRequest")
	proto.RegisterType((*SearchChatPlanRequest)(nil), "engine.SearchChatPlanRequest")
	proto.RegisterType((*ListChatPlan)(nil), "engine.ListChatPlan")
	proto.RegisterType((*ReadChatPlanRequest)(nil), "engine.ReadChatPlanRequest")
	proto.RegisterType((*UpdateChatPlanRequest)(nil), "engine.UpdateChatPlanRequest")
	proto.RegisterType((*PatchChatPlanRequest)(nil), "engine.PatchChatPlanRequest")
	proto.RegisterType((*DeleteChatPlanRequest)(nil), "engine.DeleteChatPlanRequest")
}

func init() { proto.RegisterFile("acr_chat_plan.proto", fileDescriptor_f5ca83868761bb59) }

var fileDescriptor_f5ca83868761bb59 = []byte{
	// 608 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xd6, 0x3a, 0xb5, 0x9b, 0x6e, 0x8a, 0x85, 0xb6, 0x4d, 0x65, 0xa5, 0x3f, 0x32, 0x06, 0x8a,
	0xd5, 0x43, 0x2c, 0x85, 0x1b, 0x47, 0xca, 0x09, 0xf5, 0x50, 0xb9, 0xe2, 0xc2, 0xa5, 0x5a, 0xdb,
	0x83, 0xb3, 0xc2, 0xd9, 0x75, 0xed, 0x35, 0x20, 0x10, 0x17, 0x0e, 0x5c, 0x91, 0xe0, 0x01, 0x78,
	0x0b, 0x78, 0x10, 0x5e, 0x81, 0x07, 0x41, 0x59, 0xc7, 0xa9, 0xdd, 0x6e, 0x5a, 0x8e, 0xbd, 0xcd,
	0xee, 0x6c, 0xe6, 0xfb, 0xe6, 0x9b, 0x6f, 0x62, 0xbc, 0x45, 0xe3, 0xe2, 0x3c, 0x9e, 0x52, 0x79,
	0x9e, 0x67, 0x94, 0x8f, 0xf3, 0x42, 0x48, 0x41, 0x2c, 0xe0, 0x29, 0xe3, 0x30, 0x1a, 0xc4, 0x82,
	0x97, 0xb2, 0xbe, 0x1c, 0xed, 0xa5, 0x42, 0xa4, 0x19, 0x04, 0x34, 0x67, 0x01, 0xe5, 0x5c, 0x48,
	0x2a, 0x99, 0xe0, 0x65, 0x9d, 0xf5, 0xbe, 0x23, 0xdc, 0x3f, 0x9e, 0x52, 0x79, 0x9a, 0x51, 0x4e,
	0x6c, 0x6c, 0xb0, 0xc4, 0x41, 0x2e, 0xf2, 0xcd, 0xd0, 0x60, 0x09, 0x21, 0x78, 0x8d, 0xd3, 0x19,
	0x38, 0x86, 0x8b, 0xfc, 0x8d, 0x50, 0xc5, 0xc4, 0xc5, 0x83, 0x04, 0xca, 0xb8, 0x60, 0xf9, 0xbc,
	0x8c, 0xd3, 0x53, 0xa9, 0xf6, 0x15, 0x39, 0xc4, 0x56, 0x19, 0x4f, 0x61, 0x46, 0x9d, 0x35, 0x17,
	0xf9, 0x83, 0x89, 0x3d, 0xae, 0x69, 0x8d, 0x4f, 0x84, 0x78, 0x5b, 0xe5, 0xe1, 0x22, 0x4b, 0x1c,
	0xbc, 0x0e, 0x9c, 0x46, 0x19, 0x24, 0x8e, 0xe9, 0x22, 0xbf, 0x1f, 0x36, 0x47, 0xef, 0x1b, 0xc2,
	0xc3, 0xe3, 0x02, 0xa8, 0x84, 0x86, 0x5a, 0x08, 0x17, 0x15, 0x94, 0x72, 0xc9, 0x08, 0xad, 0x66,
	0x64, 0xdc, 0xc4, 0xa8, 0xf7, 0xbf, 0x8c, 0xd6, 0xba, 0x8c, 0x7e, 0x23, 0x3c, 0x3c, 0x03, 0x5a,
	0xc4, 0x53, 0x0d, 0xa3, 0x9c, 0xa6, 0xb0, 0x50, 0x4d, 0xc5, 0xf3, 0xbb, 0x92, 0x7d, 0xac, 0x75,
	0x33, 0x43, 0x15, 0x93, 0x4d, 0x8c, 0x2e, 0x16, 0x6a, 0xa1, 0x0b, 0xf5, 0x42, 0x14, 0x52, 0xc1,
	0x6c, 0x84, 0x2a, 0x26, 0x3b, 0xd8, 0x7a, 0xc3, 0x20, 0x4b, 0x4a, 0xc7, 0x74, 0x7b, 0xfe, 0x46,
	0xb8, 0x38, 0x2d, 0xa6, 0x62, 0xb9, 0xbd, 0x2b, 0x53, 0x59, 0x6f, 0x69, 0xd0, 0x62, 0xde, 0xef,
	0x32, 0x7f, 0x89, 0x37, 0x4f, 0x58, 0x29, 0x97, 0x33, 0x9e, 0xff, 0x1a, 0x3e, 0x48, 0xc5, 0xb7,
	0x1f, 0xaa, 0x98, 0x1c, 0x62, 0x93, 0x49, 0x98, 0x95, 0x8e, 0xe1, 0xf6, 0xfc, 0xc1, 0xe4, 0x7e,
	0x23, 0xcf, 0xb2, 0xd7, 0x3a, 0xed, 0x3d, 0xc6, 0x5b, 0x21, 0xd0, 0xe4, 0xaa, 0x04, 0x57, 0x6c,
	0xe3, 0xfd, 0x44, 0x78, 0xf8, 0x2a, 0x4f, 0x34, 0xe3, 0xbb, 0x2b, 0x06, 0xfb, 0x85, 0xf0, 0xf6,
	0x29, 0x95, 0xd7, 0xa7, 0x79, 0x47, 0x08, 0xb6, 0xbc, 0x60, 0xb5, 0xbd, 0xe0, 0x3d, 0xc1, 0xc3,
	0x17, 0x90, 0xc1, 0xad, 0xca, 0x4e, 0xbe, 0x9a, 0x78, 0x27, 0x14, 0x95, 0x64, 0x3c, 0x6d, 0x9e,
	0x9e, 0x41, 0xf1, 0x8e, 0xc5, 0x40, 0x12, 0x6c, 0x77, 0x97, 0x8b, 0xec, 0x2f, 0x07, 0xae, 0x5b,
	0xba, 0xd1, 0x35, 0x3f, 0x78, 0x0f, 0xbe, 0xfc, 0xf9, 0xfb, 0xc3, 0xd8, 0xf5, 0x76, 0x82, 0xa2,
	0xc6, 0x08, 0x44, 0x25, 0x23, 0x51, 0xf1, 0x24, 0x98, 0xff, 0x27, 0x3d, 0x43, 0x47, 0x04, 0xb0,
	0xdd, 0x5d, 0x98, 0x4b, 0x14, 0xed, 0x22, 0x8d, 0xb6, 0x97, 0x22, 0xb5, 0xec, 0xea, 0x1d, 0x28,
	0x24, 0x87, 0xac, 0x40, 0x22, 0x11, 0xde, 0x6c, 0x5b, 0x92, 0xec, 0x36, 0x55, 0x34, 0x46, 0xd5,
	0x34, 0xf2, 0x50, 0x95, 0xdf, 0x27, 0xbb, 0xfa, 0xf2, 0xc1, 0x27, 0x96, 0x7c, 0x26, 0x0c, 0xdb,
	0x5d, 0x3b, 0x5f, 0xb6, 0xa2, 0xb5, 0xb9, 0x06, 0xe7, 0x50, 0xe1, 0xb8, 0xa3, 0x9b, 0x70, 0xe6,
	0xaa, 0xa5, 0xf8, 0x5e, 0xc7, 0x97, 0x64, 0xaf, 0x29, 0xa5, 0xb3, 0xeb, 0x6a, 0xa0, 0xc9, 0xed,
	0x40, 0x76, 0xd7, 0x48, 0x97, 0x3d, 0x69, 0x0d, 0xa6, 0x81, 0x7a, 0xa4, 0xa0, 0x0e, 0x8e, 0xf6,
	0x34, 0x50, 0x34, 0xcb, 0x4a, 0x85, 0xf5, 0xdc, 0x7b, 0xed, 0xa6, 0x4c, 0x4e, 0xab, 0x68, 0x1c,
	0x8b, 0x59, 0xf0, 0x1e, 0x22, 0x26, 0x21, 0x0b, 0xd4, 0xc7, 0xa7, 0x0c, 0xea, 0x92, 0x91, 0xa5,
	0x8e, 0x4f, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x3f, 0x36, 0x2c, 0xf8, 0xd5, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RoutingChatPlanServiceClient is the client API for RoutingChatPlanService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RoutingChatPlanServiceClient interface {
	CreateChatPlan(ctx context.Context, in *CreateChatPlanRequest, opts ...grpc.CallOption) (*ChatPlan, error)
	SearchChatPlan(ctx context.Context, in *SearchChatPlanRequest, opts ...grpc.CallOption) (*ListChatPlan, error)
	ReadChatPlan(ctx context.Context, in *ReadChatPlanRequest, opts ...grpc.CallOption) (*ChatPlan, error)
	UpdateChatPlan(ctx context.Context, in *UpdateChatPlanRequest, opts ...grpc.CallOption) (*ChatPlan, error)
	PatchChatPlan(ctx context.Context, in *PatchChatPlanRequest, opts ...grpc.CallOption) (*ChatPlan, error)
	DeleteChatPlan(ctx context.Context, in *DeleteChatPlanRequest, opts ...grpc.CallOption) (*ChatPlan, error)
}

type routingChatPlanServiceClient struct {
	cc *grpc.ClientConn
}

func NewRoutingChatPlanServiceClient(cc *grpc.ClientConn) RoutingChatPlanServiceClient {
	return &routingChatPlanServiceClient{cc}
}

func (c *routingChatPlanServiceClient) CreateChatPlan(ctx context.Context, in *CreateChatPlanRequest, opts ...grpc.CallOption) (*ChatPlan, error) {
	out := new(ChatPlan)
	err := c.cc.Invoke(ctx, "/engine.RoutingChatPlanService/CreateChatPlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routingChatPlanServiceClient) SearchChatPlan(ctx context.Context, in *SearchChatPlanRequest, opts ...grpc.CallOption) (*ListChatPlan, error) {
	out := new(ListChatPlan)
	err := c.cc.Invoke(ctx, "/engine.RoutingChatPlanService/SearchChatPlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routingChatPlanServiceClient) ReadChatPlan(ctx context.Context, in *ReadChatPlanRequest, opts ...grpc.CallOption) (*ChatPlan, error) {
	out := new(ChatPlan)
	err := c.cc.Invoke(ctx, "/engine.RoutingChatPlanService/ReadChatPlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routingChatPlanServiceClient) UpdateChatPlan(ctx context.Context, in *UpdateChatPlanRequest, opts ...grpc.CallOption) (*ChatPlan, error) {
	out := new(ChatPlan)
	err := c.cc.Invoke(ctx, "/engine.RoutingChatPlanService/UpdateChatPlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routingChatPlanServiceClient) PatchChatPlan(ctx context.Context, in *PatchChatPlanRequest, opts ...grpc.CallOption) (*ChatPlan, error) {
	out := new(ChatPlan)
	err := c.cc.Invoke(ctx, "/engine.RoutingChatPlanService/PatchChatPlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routingChatPlanServiceClient) DeleteChatPlan(ctx context.Context, in *DeleteChatPlanRequest, opts ...grpc.CallOption) (*ChatPlan, error) {
	out := new(ChatPlan)
	err := c.cc.Invoke(ctx, "/engine.RoutingChatPlanService/DeleteChatPlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoutingChatPlanServiceServer is the server API for RoutingChatPlanService service.
type RoutingChatPlanServiceServer interface {
	CreateChatPlan(context.Context, *CreateChatPlanRequest) (*ChatPlan, error)
	SearchChatPlan(context.Context, *SearchChatPlanRequest) (*ListChatPlan, error)
	ReadChatPlan(context.Context, *ReadChatPlanRequest) (*ChatPlan, error)
	UpdateChatPlan(context.Context, *UpdateChatPlanRequest) (*ChatPlan, error)
	PatchChatPlan(context.Context, *PatchChatPlanRequest) (*ChatPlan, error)
	DeleteChatPlan(context.Context, *DeleteChatPlanRequest) (*ChatPlan, error)
}

// UnimplementedRoutingChatPlanServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRoutingChatPlanServiceServer struct {
}

func (*UnimplementedRoutingChatPlanServiceServer) CreateChatPlan(ctx context.Context, req *CreateChatPlanRequest) (*ChatPlan, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChatPlan not implemented")
}
func (*UnimplementedRoutingChatPlanServiceServer) SearchChatPlan(ctx context.Context, req *SearchChatPlanRequest) (*ListChatPlan, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchChatPlan not implemented")
}
func (*UnimplementedRoutingChatPlanServiceServer) ReadChatPlan(ctx context.Context, req *ReadChatPlanRequest) (*ChatPlan, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadChatPlan not implemented")
}
func (*UnimplementedRoutingChatPlanServiceServer) UpdateChatPlan(ctx context.Context, req *UpdateChatPlanRequest) (*ChatPlan, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateChatPlan not implemented")
}
func (*UnimplementedRoutingChatPlanServiceServer) PatchChatPlan(ctx context.Context, req *PatchChatPlanRequest) (*ChatPlan, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchChatPlan not implemented")
}
func (*UnimplementedRoutingChatPlanServiceServer) DeleteChatPlan(ctx context.Context, req *DeleteChatPlanRequest) (*ChatPlan, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteChatPlan not implemented")
}

func RegisterRoutingChatPlanServiceServer(s *grpc.Server, srv RoutingChatPlanServiceServer) {
	s.RegisterService(&_RoutingChatPlanService_serviceDesc, srv)
}

func _RoutingChatPlanService_CreateChatPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChatPlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutingChatPlanServiceServer).CreateChatPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.RoutingChatPlanService/CreateChatPlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutingChatPlanServiceServer).CreateChatPlan(ctx, req.(*CreateChatPlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoutingChatPlanService_SearchChatPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchChatPlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutingChatPlanServiceServer).SearchChatPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.RoutingChatPlanService/SearchChatPlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutingChatPlanServiceServer).SearchChatPlan(ctx, req.(*SearchChatPlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoutingChatPlanService_ReadChatPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadChatPlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutingChatPlanServiceServer).ReadChatPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.RoutingChatPlanService/ReadChatPlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutingChatPlanServiceServer).ReadChatPlan(ctx, req.(*ReadChatPlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoutingChatPlanService_UpdateChatPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateChatPlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutingChatPlanServiceServer).UpdateChatPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.RoutingChatPlanService/UpdateChatPlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutingChatPlanServiceServer).UpdateChatPlan(ctx, req.(*UpdateChatPlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoutingChatPlanService_PatchChatPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchChatPlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutingChatPlanServiceServer).PatchChatPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.RoutingChatPlanService/PatchChatPlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutingChatPlanServiceServer).PatchChatPlan(ctx, req.(*PatchChatPlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoutingChatPlanService_DeleteChatPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteChatPlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutingChatPlanServiceServer).DeleteChatPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.RoutingChatPlanService/DeleteChatPlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutingChatPlanServiceServer).DeleteChatPlan(ctx, req.(*DeleteChatPlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RoutingChatPlanService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "engine.RoutingChatPlanService",
	HandlerType: (*RoutingChatPlanServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateChatPlan",
			Handler:    _RoutingChatPlanService_CreateChatPlan_Handler,
		},
		{
			MethodName: "SearchChatPlan",
			Handler:    _RoutingChatPlanService_SearchChatPlan_Handler,
		},
		{
			MethodName: "ReadChatPlan",
			Handler:    _RoutingChatPlanService_ReadChatPlan_Handler,
		},
		{
			MethodName: "UpdateChatPlan",
			Handler:    _RoutingChatPlanService_UpdateChatPlan_Handler,
		},
		{
			MethodName: "PatchChatPlan",
			Handler:    _RoutingChatPlanService_PatchChatPlan_Handler,
		},
		{
			MethodName: "DeleteChatPlan",
			Handler:    _RoutingChatPlanService_DeleteChatPlan_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "acr_chat_plan.proto",
}
