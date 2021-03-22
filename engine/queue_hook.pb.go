// Code generated by protoc-gen-go. DO NOT EDIT.
// source: queue_hook.proto

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

type DeleteQueueHookRequest struct {
	QueueId              uint32   `protobuf:"varint,1,opt,name=queue_id,json=queueId,proto3" json:"queue_id,omitempty"`
	Id                   uint32   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteQueueHookRequest) Reset()         { *m = DeleteQueueHookRequest{} }
func (m *DeleteQueueHookRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteQueueHookRequest) ProtoMessage()    {}
func (*DeleteQueueHookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d7d7c317d7f4f51, []int{0}
}

func (m *DeleteQueueHookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteQueueHookRequest.Unmarshal(m, b)
}
func (m *DeleteQueueHookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteQueueHookRequest.Marshal(b, m, deterministic)
}
func (m *DeleteQueueHookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteQueueHookRequest.Merge(m, src)
}
func (m *DeleteQueueHookRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteQueueHookRequest.Size(m)
}
func (m *DeleteQueueHookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteQueueHookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteQueueHookRequest proto.InternalMessageInfo

func (m *DeleteQueueHookRequest) GetQueueId() uint32 {
	if m != nil {
		return m.QueueId
	}
	return 0
}

func (m *DeleteQueueHookRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type PatchQueueHookRequest struct {
	Fields               []string `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty"`
	QueueId              uint32   `protobuf:"varint,2,opt,name=queue_id,json=queueId,proto3" json:"queue_id,omitempty"`
	Id                   uint32   `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Schema               *Lookup  `protobuf:"bytes,4,opt,name=schema,proto3" json:"schema,omitempty"`
	Event                string   `protobuf:"bytes,5,opt,name=event,proto3" json:"event,omitempty"`
	Enabled              bool     `protobuf:"varint,6,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Properties           []string `protobuf:"bytes,7,rep,name=properties,proto3" json:"properties,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PatchQueueHookRequest) Reset()         { *m = PatchQueueHookRequest{} }
func (m *PatchQueueHookRequest) String() string { return proto.CompactTextString(m) }
func (*PatchQueueHookRequest) ProtoMessage()    {}
func (*PatchQueueHookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d7d7c317d7f4f51, []int{1}
}

func (m *PatchQueueHookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PatchQueueHookRequest.Unmarshal(m, b)
}
func (m *PatchQueueHookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PatchQueueHookRequest.Marshal(b, m, deterministic)
}
func (m *PatchQueueHookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PatchQueueHookRequest.Merge(m, src)
}
func (m *PatchQueueHookRequest) XXX_Size() int {
	return xxx_messageInfo_PatchQueueHookRequest.Size(m)
}
func (m *PatchQueueHookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PatchQueueHookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PatchQueueHookRequest proto.InternalMessageInfo

func (m *PatchQueueHookRequest) GetFields() []string {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *PatchQueueHookRequest) GetQueueId() uint32 {
	if m != nil {
		return m.QueueId
	}
	return 0
}

func (m *PatchQueueHookRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PatchQueueHookRequest) GetSchema() *Lookup {
	if m != nil {
		return m.Schema
	}
	return nil
}

func (m *PatchQueueHookRequest) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func (m *PatchQueueHookRequest) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *PatchQueueHookRequest) GetProperties() []string {
	if m != nil {
		return m.Properties
	}
	return nil
}

type UpdateQueueHookRequest struct {
	QueueId              uint32   `protobuf:"varint,1,opt,name=queue_id,json=queueId,proto3" json:"queue_id,omitempty"`
	Id                   uint32   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Schema               *Lookup  `protobuf:"bytes,3,opt,name=schema,proto3" json:"schema,omitempty"`
	Event                string   `protobuf:"bytes,4,opt,name=event,proto3" json:"event,omitempty"`
	Enabled              bool     `protobuf:"varint,5,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Properties           []string `protobuf:"bytes,6,rep,name=properties,proto3" json:"properties,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateQueueHookRequest) Reset()         { *m = UpdateQueueHookRequest{} }
func (m *UpdateQueueHookRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateQueueHookRequest) ProtoMessage()    {}
func (*UpdateQueueHookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d7d7c317d7f4f51, []int{2}
}

func (m *UpdateQueueHookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateQueueHookRequest.Unmarshal(m, b)
}
func (m *UpdateQueueHookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateQueueHookRequest.Marshal(b, m, deterministic)
}
func (m *UpdateQueueHookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateQueueHookRequest.Merge(m, src)
}
func (m *UpdateQueueHookRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateQueueHookRequest.Size(m)
}
func (m *UpdateQueueHookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateQueueHookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateQueueHookRequest proto.InternalMessageInfo

func (m *UpdateQueueHookRequest) GetQueueId() uint32 {
	if m != nil {
		return m.QueueId
	}
	return 0
}

func (m *UpdateQueueHookRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateQueueHookRequest) GetSchema() *Lookup {
	if m != nil {
		return m.Schema
	}
	return nil
}

func (m *UpdateQueueHookRequest) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func (m *UpdateQueueHookRequest) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *UpdateQueueHookRequest) GetProperties() []string {
	if m != nil {
		return m.Properties
	}
	return nil
}

type SearchQueueHookRequest struct {
	QueueId              uint32   `protobuf:"varint,1,opt,name=queue_id,json=queueId,proto3" json:"queue_id,omitempty"`
	Page                 int32    `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Size                 int32    `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Q                    string   `protobuf:"bytes,4,opt,name=q,proto3" json:"q,omitempty"`
	Sort                 string   `protobuf:"bytes,5,opt,name=sort,proto3" json:"sort,omitempty"`
	Fields               []string `protobuf:"bytes,6,rep,name=fields,proto3" json:"fields,omitempty"`
	Id                   []uint32 `protobuf:"varint,7,rep,packed,name=id,proto3" json:"id,omitempty"`
	SchemaId             []uint32 `protobuf:"varint,8,rep,packed,name=schema_id,json=schemaId,proto3" json:"schema_id,omitempty"`
	Event                []string `protobuf:"bytes,9,rep,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchQueueHookRequest) Reset()         { *m = SearchQueueHookRequest{} }
func (m *SearchQueueHookRequest) String() string { return proto.CompactTextString(m) }
func (*SearchQueueHookRequest) ProtoMessage()    {}
func (*SearchQueueHookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d7d7c317d7f4f51, []int{3}
}

func (m *SearchQueueHookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchQueueHookRequest.Unmarshal(m, b)
}
func (m *SearchQueueHookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchQueueHookRequest.Marshal(b, m, deterministic)
}
func (m *SearchQueueHookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchQueueHookRequest.Merge(m, src)
}
func (m *SearchQueueHookRequest) XXX_Size() int {
	return xxx_messageInfo_SearchQueueHookRequest.Size(m)
}
func (m *SearchQueueHookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchQueueHookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchQueueHookRequest proto.InternalMessageInfo

func (m *SearchQueueHookRequest) GetQueueId() uint32 {
	if m != nil {
		return m.QueueId
	}
	return 0
}

func (m *SearchQueueHookRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *SearchQueueHookRequest) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *SearchQueueHookRequest) GetQ() string {
	if m != nil {
		return m.Q
	}
	return ""
}

func (m *SearchQueueHookRequest) GetSort() string {
	if m != nil {
		return m.Sort
	}
	return ""
}

func (m *SearchQueueHookRequest) GetFields() []string {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *SearchQueueHookRequest) GetId() []uint32 {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *SearchQueueHookRequest) GetSchemaId() []uint32 {
	if m != nil {
		return m.SchemaId
	}
	return nil
}

func (m *SearchQueueHookRequest) GetEvent() []string {
	if m != nil {
		return m.Event
	}
	return nil
}

type CreateQueueHookRequest struct {
	QueueId              uint32   `protobuf:"varint,1,opt,name=queue_id,json=queueId,proto3" json:"queue_id,omitempty"`
	Schema               *Lookup  `protobuf:"bytes,2,opt,name=schema,proto3" json:"schema,omitempty"`
	Event                string   `protobuf:"bytes,3,opt,name=event,proto3" json:"event,omitempty"`
	Enabled              bool     `protobuf:"varint,4,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Properties           []string `protobuf:"bytes,5,rep,name=properties,proto3" json:"properties,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateQueueHookRequest) Reset()         { *m = CreateQueueHookRequest{} }
func (m *CreateQueueHookRequest) String() string { return proto.CompactTextString(m) }
func (*CreateQueueHookRequest) ProtoMessage()    {}
func (*CreateQueueHookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d7d7c317d7f4f51, []int{4}
}

func (m *CreateQueueHookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateQueueHookRequest.Unmarshal(m, b)
}
func (m *CreateQueueHookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateQueueHookRequest.Marshal(b, m, deterministic)
}
func (m *CreateQueueHookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateQueueHookRequest.Merge(m, src)
}
func (m *CreateQueueHookRequest) XXX_Size() int {
	return xxx_messageInfo_CreateQueueHookRequest.Size(m)
}
func (m *CreateQueueHookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateQueueHookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateQueueHookRequest proto.InternalMessageInfo

func (m *CreateQueueHookRequest) GetQueueId() uint32 {
	if m != nil {
		return m.QueueId
	}
	return 0
}

func (m *CreateQueueHookRequest) GetSchema() *Lookup {
	if m != nil {
		return m.Schema
	}
	return nil
}

func (m *CreateQueueHookRequest) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func (m *CreateQueueHookRequest) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *CreateQueueHookRequest) GetProperties() []string {
	if m != nil {
		return m.Properties
	}
	return nil
}

type ReadQueueHookRequest struct {
	QueueId              uint32   `protobuf:"varint,1,opt,name=queue_id,json=queueId,proto3" json:"queue_id,omitempty"`
	Id                   uint32   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadQueueHookRequest) Reset()         { *m = ReadQueueHookRequest{} }
func (m *ReadQueueHookRequest) String() string { return proto.CompactTextString(m) }
func (*ReadQueueHookRequest) ProtoMessage()    {}
func (*ReadQueueHookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d7d7c317d7f4f51, []int{5}
}

func (m *ReadQueueHookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadQueueHookRequest.Unmarshal(m, b)
}
func (m *ReadQueueHookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadQueueHookRequest.Marshal(b, m, deterministic)
}
func (m *ReadQueueHookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadQueueHookRequest.Merge(m, src)
}
func (m *ReadQueueHookRequest) XXX_Size() int {
	return xxx_messageInfo_ReadQueueHookRequest.Size(m)
}
func (m *ReadQueueHookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadQueueHookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadQueueHookRequest proto.InternalMessageInfo

func (m *ReadQueueHookRequest) GetQueueId() uint32 {
	if m != nil {
		return m.QueueId
	}
	return 0
}

func (m *ReadQueueHookRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type QueueHook struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Schema               *Lookup  `protobuf:"bytes,2,opt,name=schema,proto3" json:"schema,omitempty"`
	Event                string   `protobuf:"bytes,3,opt,name=event,proto3" json:"event,omitempty"`
	Enabled              bool     `protobuf:"varint,4,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Properties           []string `protobuf:"bytes,5,rep,name=properties,proto3" json:"properties,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueueHook) Reset()         { *m = QueueHook{} }
func (m *QueueHook) String() string { return proto.CompactTextString(m) }
func (*QueueHook) ProtoMessage()    {}
func (*QueueHook) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d7d7c317d7f4f51, []int{6}
}

func (m *QueueHook) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueueHook.Unmarshal(m, b)
}
func (m *QueueHook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueueHook.Marshal(b, m, deterministic)
}
func (m *QueueHook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueueHook.Merge(m, src)
}
func (m *QueueHook) XXX_Size() int {
	return xxx_messageInfo_QueueHook.Size(m)
}
func (m *QueueHook) XXX_DiscardUnknown() {
	xxx_messageInfo_QueueHook.DiscardUnknown(m)
}

var xxx_messageInfo_QueueHook proto.InternalMessageInfo

func (m *QueueHook) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *QueueHook) GetSchema() *Lookup {
	if m != nil {
		return m.Schema
	}
	return nil
}

func (m *QueueHook) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func (m *QueueHook) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *QueueHook) GetProperties() []string {
	if m != nil {
		return m.Properties
	}
	return nil
}

type ListQueueHook struct {
	Next                 bool         `protobuf:"varint,1,opt,name=next,proto3" json:"next,omitempty"`
	Items                []*QueueHook `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ListQueueHook) Reset()         { *m = ListQueueHook{} }
func (m *ListQueueHook) String() string { return proto.CompactTextString(m) }
func (*ListQueueHook) ProtoMessage()    {}
func (*ListQueueHook) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d7d7c317d7f4f51, []int{7}
}

func (m *ListQueueHook) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListQueueHook.Unmarshal(m, b)
}
func (m *ListQueueHook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListQueueHook.Marshal(b, m, deterministic)
}
func (m *ListQueueHook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListQueueHook.Merge(m, src)
}
func (m *ListQueueHook) XXX_Size() int {
	return xxx_messageInfo_ListQueueHook.Size(m)
}
func (m *ListQueueHook) XXX_DiscardUnknown() {
	xxx_messageInfo_ListQueueHook.DiscardUnknown(m)
}

var xxx_messageInfo_ListQueueHook proto.InternalMessageInfo

func (m *ListQueueHook) GetNext() bool {
	if m != nil {
		return m.Next
	}
	return false
}

func (m *ListQueueHook) GetItems() []*QueueHook {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*DeleteQueueHookRequest)(nil), "engine.DeleteQueueHookRequest")
	proto.RegisterType((*PatchQueueHookRequest)(nil), "engine.PatchQueueHookRequest")
	proto.RegisterType((*UpdateQueueHookRequest)(nil), "engine.UpdateQueueHookRequest")
	proto.RegisterType((*SearchQueueHookRequest)(nil), "engine.SearchQueueHookRequest")
	proto.RegisterType((*CreateQueueHookRequest)(nil), "engine.CreateQueueHookRequest")
	proto.RegisterType((*ReadQueueHookRequest)(nil), "engine.ReadQueueHookRequest")
	proto.RegisterType((*QueueHook)(nil), "engine.QueueHook")
	proto.RegisterType((*ListQueueHook)(nil), "engine.ListQueueHook")
}

func init() { proto.RegisterFile("queue_hook.proto", fileDescriptor_5d7d7c317d7f4f51) }

var fileDescriptor_5d7d7c317d7f4f51 = []byte{
	// 660 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xd6, 0x3a, 0xb1, 0x9b, 0x4c, 0x69, 0x53, 0x56, 0x6d, 0x64, 0x42, 0xa9, 0x2c, 0x0b, 0x95,
	0xb4, 0xa0, 0x58, 0x04, 0x4e, 0xdc, 0xa0, 0x1c, 0xa8, 0xd4, 0x03, 0xb8, 0xe2, 0xc2, 0xa5, 0x72,
	0xec, 0x21, 0x59, 0xd5, 0xf5, 0x3a, 0xde, 0x4d, 0x5b, 0xb5, 0xea, 0x85, 0x57, 0xe8, 0x63, 0xf0,
	0x08, 0xbc, 0x05, 0x5c, 0x78, 0x00, 0x78, 0x0f, 0x94, 0x75, 0x5c, 0xe7, 0xc7, 0x8d, 0x52, 0x55,
	0xe2, 0xb6, 0x33, 0x6b, 0xcf, 0xf7, 0xcd, 0xf7, 0xed, 0xee, 0xc0, 0x5a, 0x7f, 0x80, 0x03, 0x3c,
	0xea, 0x71, 0x7e, 0xdc, 0x8a, 0x13, 0x2e, 0x39, 0x35, 0x30, 0xea, 0xb2, 0x08, 0x1b, 0x9b, 0x5d,
	0xce, 0xbb, 0x21, 0x3a, 0x5e, 0xcc, 0x1c, 0x2f, 0x8a, 0xb8, 0xf4, 0x24, 0xe3, 0x91, 0x48, 0xbf,
	0x6a, 0x2c, 0xfb, 0x3c, 0x12, 0x32, 0x0d, 0xec, 0x3d, 0xa8, 0xbf, 0xc7, 0x10, 0x25, 0x7e, 0x1a,
	0x16, 0xfb, 0xc0, 0xf9, 0xb1, 0x8b, 0xfd, 0x01, 0x0a, 0x49, 0x1f, 0x41, 0x25, 0x05, 0x60, 0x81,
	0x49, 0x2c, 0xd2, 0x5c, 0x71, 0x97, 0x54, 0xbc, 0x1f, 0xd0, 0x55, 0xd0, 0x58, 0x60, 0x6a, 0x2a,
	0xa9, 0xb1, 0xc0, 0xfe, 0x49, 0x60, 0xe3, 0xa3, 0x27, 0xfd, 0xde, 0x4c, 0x91, 0x3a, 0x18, 0x5f,
	0x19, 0x86, 0x81, 0x30, 0x89, 0x55, 0x6a, 0x56, 0xdd, 0x51, 0x34, 0x51, 0x5c, 0x2b, 0x2a, 0x5e,
	0xca, 0x8a, 0xd3, 0x6d, 0x30, 0x84, 0xdf, 0xc3, 0x13, 0xcf, 0x2c, 0x5b, 0xa4, 0xb9, 0xdc, 0x5e,
	0x6d, 0xa5, 0x5d, 0xb6, 0x0e, 0x38, 0x3f, 0x1e, 0xc4, 0xee, 0x68, 0x97, 0xae, 0x83, 0x8e, 0xa7,
	0x18, 0x49, 0x53, 0xb7, 0x48, 0xb3, 0xea, 0xa6, 0x01, 0x35, 0x61, 0x09, 0x23, 0xaf, 0x13, 0x62,
	0x60, 0x1a, 0x16, 0x69, 0x56, 0xdc, 0x2c, 0xa4, 0x5b, 0x00, 0x71, 0xc2, 0x63, 0x4c, 0x24, 0x43,
	0x61, 0x2e, 0x29, 0x7a, 0x63, 0x19, 0xfb, 0x07, 0x81, 0xfa, 0xe7, 0x38, 0xf0, 0xee, 0x25, 0xcd,
	0x18, 0xfb, 0xd2, 0x62, 0xec, 0xcb, 0xb7, 0xb0, 0xd7, 0xe7, 0xb1, 0x37, 0x66, 0xd8, 0xff, 0x26,
	0x50, 0x3f, 0x44, 0x2f, 0x29, 0xf0, 0x64, 0x0e, 0x7b, 0x0a, 0xe5, 0xd8, 0xeb, 0xa2, 0xe2, 0xaf,
	0xbb, 0x6a, 0x3d, 0xcc, 0x09, 0x76, 0x81, 0x8a, 0xbf, 0xee, 0xaa, 0x35, 0x7d, 0x00, 0xa4, 0x3f,
	0x62, 0x4a, 0xfa, 0xea, 0x0b, 0x9e, 0x64, 0xc2, 0xab, 0xf5, 0x98, 0xf1, 0xc6, 0x84, 0xf1, 0xa9,
	0x3e, 0x43, 0xb5, 0x53, 0x7d, 0x1e, 0x43, 0x35, 0x55, 0x60, 0xc8, 0xa6, 0xa2, 0xd2, 0x95, 0x34,
	0xb1, 0x1f, 0xe4, 0xa2, 0x54, 0x55, 0x8d, 0x34, 0xb0, 0xbf, 0x13, 0xa8, 0xef, 0x25, 0x78, 0x47,
	0x63, 0x72, 0x23, 0xb4, 0xc5, 0x8c, 0x28, 0xdd, 0x62, 0x44, 0x79, 0x9e, 0x11, 0xfa, 0x8c, 0x11,
	0x6f, 0x61, 0xdd, 0x45, 0x2f, 0xb8, 0xcf, 0xf5, 0xba, 0x26, 0x50, 0xbd, 0xf9, 0x7f, 0xb4, 0x4b,
	0x0a, 0x4e, 0xd8, 0xff, 0x6d, 0xec, 0x00, 0x56, 0x0e, 0x98, 0x90, 0x39, 0x31, 0x0a, 0xe5, 0x08,
	0xcf, 0xa5, 0xa2, 0x56, 0x71, 0xd5, 0x9a, 0x3e, 0x03, 0x9d, 0x49, 0x3c, 0x11, 0xa6, 0x66, 0x95,
	0x9a, 0xcb, 0xed, 0x87, 0x19, 0xb7, 0x5c, 0x8e, 0x74, 0xbf, 0xfd, 0x57, 0x87, 0xb5, 0x9b, 0xe4,
	0x21, 0x26, 0xa7, 0xcc, 0x47, 0x3a, 0x80, 0xda, 0x94, 0xd1, 0x74, 0x2b, 0xab, 0x50, 0x7c, 0x02,
	0x1a, 0xb3, 0x08, 0xb6, 0xf3, 0xed, 0xd7, 0x9f, 0x6b, 0x6d, 0xc7, 0x7e, 0xea, 0xf8, 0x5e, 0x18,
	0x1e, 0xf9, 0x18, 0x49, 0x4c, 0x1c, 0x25, 0xb6, 0x70, 0x2e, 0x33, 0x13, 0xae, 0x9c, 0xe1, 0x43,
	0x2a, 0xde, 0x90, 0x5d, 0x7a, 0x0a, 0xb5, 0xa9, 0xab, 0x93, 0xc3, 0x16, 0xdf, 0xa9, 0xc6, 0xc6,
	0x8d, 0xe8, 0xe3, 0x92, 0xd8, 0x2f, 0x14, 0xf4, 0x36, 0x5d, 0x08, 0x9a, 0x0a, 0x58, 0x99, 0x38,
	0x2a, 0x74, 0x33, 0xab, 0x5a, 0x74, 0x82, 0x8a, 0x5a, 0x7d, 0xa9, 0xf0, 0x9e, 0xd3, 0x9d, 0x45,
	0xf0, 0x9c, 0x4b, 0x16, 0x5c, 0xd1, 0x0b, 0xa8, 0x4d, 0xbd, 0x72, 0x79, 0xb3, 0xc5, 0xcf, 0x5f,
	0x11, 0xf0, 0x6b, 0x05, 0xdc, 0x6a, 0x2c, 0x0e, 0x3c, 0x14, 0xfa, 0x1c, 0x56, 0x27, 0xc7, 0x06,
	0x7d, 0x92, 0x95, 0x2e, 0x1c, 0x27, 0x73, 0x90, 0xdb, 0x77, 0x43, 0x3e, 0x83, 0xda, 0xd4, 0xd8,
	0xcb, 0xbb, 0x2e, 0x9e, 0x87, 0x73, 0xe4, 0xde, 0x5d, 0x1c, 0xfb, 0x9d, 0xfd, 0xc5, 0xea, 0x32,
	0xd9, 0x1b, 0x74, 0x5a, 0x3e, 0x3f, 0x71, 0xce, 0xb0, 0xc3, 0x24, 0x86, 0x8e, 0x9a, 0xc5, 0xc2,
	0x49, 0x01, 0x3a, 0x86, 0x0a, 0x5f, 0xfd, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xaa, 0x2c, 0xc2, 0x3c,
	0xe1, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueueHookServiceClient is the client API for QueueHookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueueHookServiceClient interface {
	CreateQueueHook(ctx context.Context, in *CreateQueueHookRequest, opts ...grpc.CallOption) (*QueueHook, error)
	SearchQueueHook(ctx context.Context, in *SearchQueueHookRequest, opts ...grpc.CallOption) (*ListQueueHook, error)
	ReadQueueHook(ctx context.Context, in *ReadQueueHookRequest, opts ...grpc.CallOption) (*QueueHook, error)
	UpdateQueueHook(ctx context.Context, in *UpdateQueueHookRequest, opts ...grpc.CallOption) (*QueueHook, error)
	PatchQueueHook(ctx context.Context, in *PatchQueueHookRequest, opts ...grpc.CallOption) (*QueueHook, error)
	DeleteQueueHook(ctx context.Context, in *DeleteQueueHookRequest, opts ...grpc.CallOption) (*QueueHook, error)
}

type queueHookServiceClient struct {
	cc *grpc.ClientConn
}

func NewQueueHookServiceClient(cc *grpc.ClientConn) QueueHookServiceClient {
	return &queueHookServiceClient{cc}
}

func (c *queueHookServiceClient) CreateQueueHook(ctx context.Context, in *CreateQueueHookRequest, opts ...grpc.CallOption) (*QueueHook, error) {
	out := new(QueueHook)
	err := c.cc.Invoke(ctx, "/engine.QueueHookService/CreateQueueHook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueHookServiceClient) SearchQueueHook(ctx context.Context, in *SearchQueueHookRequest, opts ...grpc.CallOption) (*ListQueueHook, error) {
	out := new(ListQueueHook)
	err := c.cc.Invoke(ctx, "/engine.QueueHookService/SearchQueueHook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueHookServiceClient) ReadQueueHook(ctx context.Context, in *ReadQueueHookRequest, opts ...grpc.CallOption) (*QueueHook, error) {
	out := new(QueueHook)
	err := c.cc.Invoke(ctx, "/engine.QueueHookService/ReadQueueHook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueHookServiceClient) UpdateQueueHook(ctx context.Context, in *UpdateQueueHookRequest, opts ...grpc.CallOption) (*QueueHook, error) {
	out := new(QueueHook)
	err := c.cc.Invoke(ctx, "/engine.QueueHookService/UpdateQueueHook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueHookServiceClient) PatchQueueHook(ctx context.Context, in *PatchQueueHookRequest, opts ...grpc.CallOption) (*QueueHook, error) {
	out := new(QueueHook)
	err := c.cc.Invoke(ctx, "/engine.QueueHookService/PatchQueueHook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueHookServiceClient) DeleteQueueHook(ctx context.Context, in *DeleteQueueHookRequest, opts ...grpc.CallOption) (*QueueHook, error) {
	out := new(QueueHook)
	err := c.cc.Invoke(ctx, "/engine.QueueHookService/DeleteQueueHook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueueHookServiceServer is the server API for QueueHookService service.
type QueueHookServiceServer interface {
	CreateQueueHook(context.Context, *CreateQueueHookRequest) (*QueueHook, error)
	SearchQueueHook(context.Context, *SearchQueueHookRequest) (*ListQueueHook, error)
	ReadQueueHook(context.Context, *ReadQueueHookRequest) (*QueueHook, error)
	UpdateQueueHook(context.Context, *UpdateQueueHookRequest) (*QueueHook, error)
	PatchQueueHook(context.Context, *PatchQueueHookRequest) (*QueueHook, error)
	DeleteQueueHook(context.Context, *DeleteQueueHookRequest) (*QueueHook, error)
}

// UnimplementedQueueHookServiceServer can be embedded to have forward compatible implementations.
type UnimplementedQueueHookServiceServer struct {
}

func (*UnimplementedQueueHookServiceServer) CreateQueueHook(ctx context.Context, req *CreateQueueHookRequest) (*QueueHook, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateQueueHook not implemented")
}
func (*UnimplementedQueueHookServiceServer) SearchQueueHook(ctx context.Context, req *SearchQueueHookRequest) (*ListQueueHook, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchQueueHook not implemented")
}
func (*UnimplementedQueueHookServiceServer) ReadQueueHook(ctx context.Context, req *ReadQueueHookRequest) (*QueueHook, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadQueueHook not implemented")
}
func (*UnimplementedQueueHookServiceServer) UpdateQueueHook(ctx context.Context, req *UpdateQueueHookRequest) (*QueueHook, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateQueueHook not implemented")
}
func (*UnimplementedQueueHookServiceServer) PatchQueueHook(ctx context.Context, req *PatchQueueHookRequest) (*QueueHook, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchQueueHook not implemented")
}
func (*UnimplementedQueueHookServiceServer) DeleteQueueHook(ctx context.Context, req *DeleteQueueHookRequest) (*QueueHook, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQueueHook not implemented")
}

func RegisterQueueHookServiceServer(s *grpc.Server, srv QueueHookServiceServer) {
	s.RegisterService(&_QueueHookService_serviceDesc, srv)
}

func _QueueHookService_CreateQueueHook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateQueueHookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueHookServiceServer).CreateQueueHook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.QueueHookService/CreateQueueHook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueHookServiceServer).CreateQueueHook(ctx, req.(*CreateQueueHookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueueHookService_SearchQueueHook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchQueueHookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueHookServiceServer).SearchQueueHook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.QueueHookService/SearchQueueHook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueHookServiceServer).SearchQueueHook(ctx, req.(*SearchQueueHookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueueHookService_ReadQueueHook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadQueueHookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueHookServiceServer).ReadQueueHook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.QueueHookService/ReadQueueHook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueHookServiceServer).ReadQueueHook(ctx, req.(*ReadQueueHookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueueHookService_UpdateQueueHook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateQueueHookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueHookServiceServer).UpdateQueueHook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.QueueHookService/UpdateQueueHook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueHookServiceServer).UpdateQueueHook(ctx, req.(*UpdateQueueHookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueueHookService_PatchQueueHook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchQueueHookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueHookServiceServer).PatchQueueHook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.QueueHookService/PatchQueueHook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueHookServiceServer).PatchQueueHook(ctx, req.(*PatchQueueHookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueueHookService_DeleteQueueHook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteQueueHookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueHookServiceServer).DeleteQueueHook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.QueueHookService/DeleteQueueHook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueHookServiceServer).DeleteQueueHook(ctx, req.(*DeleteQueueHookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QueueHookService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "engine.QueueHookService",
	HandlerType: (*QueueHookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateQueueHook",
			Handler:    _QueueHookService_CreateQueueHook_Handler,
		},
		{
			MethodName: "SearchQueueHook",
			Handler:    _QueueHookService_SearchQueueHook_Handler,
		},
		{
			MethodName: "ReadQueueHook",
			Handler:    _QueueHookService_ReadQueueHook_Handler,
		},
		{
			MethodName: "UpdateQueueHook",
			Handler:    _QueueHookService_UpdateQueueHook_Handler,
		},
		{
			MethodName: "PatchQueueHook",
			Handler:    _QueueHookService_PatchQueueHook_Handler,
		},
		{
			MethodName: "DeleteQueueHook",
			Handler:    _QueueHookService_DeleteQueueHook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "queue_hook.proto",
}