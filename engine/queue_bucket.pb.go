// Code generated by protoc-gen-go. DO NOT EDIT.
// source: queue_bucket.proto

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

type DeleteQueueBucketRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	QueueId              int64    `protobuf:"varint,2,opt,name=queue_id,json=queueId,proto3" json:"queue_id,omitempty"`
	DomainId             int64    `protobuf:"varint,3,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteQueueBucketRequest) Reset()         { *m = DeleteQueueBucketRequest{} }
func (m *DeleteQueueBucketRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteQueueBucketRequest) ProtoMessage()    {}
func (*DeleteQueueBucketRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4c81630b355ff82, []int{0}
}

func (m *DeleteQueueBucketRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteQueueBucketRequest.Unmarshal(m, b)
}
func (m *DeleteQueueBucketRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteQueueBucketRequest.Marshal(b, m, deterministic)
}
func (m *DeleteQueueBucketRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteQueueBucketRequest.Merge(m, src)
}
func (m *DeleteQueueBucketRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteQueueBucketRequest.Size(m)
}
func (m *DeleteQueueBucketRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteQueueBucketRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteQueueBucketRequest proto.InternalMessageInfo

func (m *DeleteQueueBucketRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DeleteQueueBucketRequest) GetQueueId() int64 {
	if m != nil {
		return m.QueueId
	}
	return 0
}

func (m *DeleteQueueBucketRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

type UpdateQueueBucketRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	QueueId              int64    `protobuf:"varint,2,opt,name=queue_id,json=queueId,proto3" json:"queue_id,omitempty"`
	Ratio                int32    `protobuf:"varint,3,opt,name=ratio,proto3" json:"ratio,omitempty"`
	Bucket               *Lookup  `protobuf:"bytes,4,opt,name=bucket,proto3" json:"bucket,omitempty"`
	DomainId             int64    `protobuf:"varint,5,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateQueueBucketRequest) Reset()         { *m = UpdateQueueBucketRequest{} }
func (m *UpdateQueueBucketRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateQueueBucketRequest) ProtoMessage()    {}
func (*UpdateQueueBucketRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4c81630b355ff82, []int{1}
}

func (m *UpdateQueueBucketRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateQueueBucketRequest.Unmarshal(m, b)
}
func (m *UpdateQueueBucketRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateQueueBucketRequest.Marshal(b, m, deterministic)
}
func (m *UpdateQueueBucketRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateQueueBucketRequest.Merge(m, src)
}
func (m *UpdateQueueBucketRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateQueueBucketRequest.Size(m)
}
func (m *UpdateQueueBucketRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateQueueBucketRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateQueueBucketRequest proto.InternalMessageInfo

func (m *UpdateQueueBucketRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateQueueBucketRequest) GetQueueId() int64 {
	if m != nil {
		return m.QueueId
	}
	return 0
}

func (m *UpdateQueueBucketRequest) GetRatio() int32 {
	if m != nil {
		return m.Ratio
	}
	return 0
}

func (m *UpdateQueueBucketRequest) GetBucket() *Lookup {
	if m != nil {
		return m.Bucket
	}
	return nil
}

func (m *UpdateQueueBucketRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

type SearchQueueBucketRequest struct {
	QueueId              int64    `protobuf:"varint,1,opt,name=queue_id,json=queueId,proto3" json:"queue_id,omitempty"`
	Page                 int32    `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Size                 int32    `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Q                    string   `protobuf:"bytes,4,opt,name=q,proto3" json:"q,omitempty"`
	Sort                 string   `protobuf:"bytes,5,opt,name=sort,proto3" json:"sort,omitempty"`
	Fields               []string `protobuf:"bytes,6,rep,name=fields,proto3" json:"fields,omitempty"`
	Id                   []uint32 `protobuf:"varint,7,rep,packed,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchQueueBucketRequest) Reset()         { *m = SearchQueueBucketRequest{} }
func (m *SearchQueueBucketRequest) String() string { return proto.CompactTextString(m) }
func (*SearchQueueBucketRequest) ProtoMessage()    {}
func (*SearchQueueBucketRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4c81630b355ff82, []int{2}
}

func (m *SearchQueueBucketRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchQueueBucketRequest.Unmarshal(m, b)
}
func (m *SearchQueueBucketRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchQueueBucketRequest.Marshal(b, m, deterministic)
}
func (m *SearchQueueBucketRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchQueueBucketRequest.Merge(m, src)
}
func (m *SearchQueueBucketRequest) XXX_Size() int {
	return xxx_messageInfo_SearchQueueBucketRequest.Size(m)
}
func (m *SearchQueueBucketRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchQueueBucketRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchQueueBucketRequest proto.InternalMessageInfo

func (m *SearchQueueBucketRequest) GetQueueId() int64 {
	if m != nil {
		return m.QueueId
	}
	return 0
}

func (m *SearchQueueBucketRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *SearchQueueBucketRequest) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *SearchQueueBucketRequest) GetQ() string {
	if m != nil {
		return m.Q
	}
	return ""
}

func (m *SearchQueueBucketRequest) GetSort() string {
	if m != nil {
		return m.Sort
	}
	return ""
}

func (m *SearchQueueBucketRequest) GetFields() []string {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *SearchQueueBucketRequest) GetId() []uint32 {
	if m != nil {
		return m.Id
	}
	return nil
}

type ListQueueBucket struct {
	Next                 bool           `protobuf:"varint,1,opt,name=next,proto3" json:"next,omitempty"`
	Items                []*QueueBucket `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ListQueueBucket) Reset()         { *m = ListQueueBucket{} }
func (m *ListQueueBucket) String() string { return proto.CompactTextString(m) }
func (*ListQueueBucket) ProtoMessage()    {}
func (*ListQueueBucket) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4c81630b355ff82, []int{3}
}

func (m *ListQueueBucket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListQueueBucket.Unmarshal(m, b)
}
func (m *ListQueueBucket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListQueueBucket.Marshal(b, m, deterministic)
}
func (m *ListQueueBucket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListQueueBucket.Merge(m, src)
}
func (m *ListQueueBucket) XXX_Size() int {
	return xxx_messageInfo_ListQueueBucket.Size(m)
}
func (m *ListQueueBucket) XXX_DiscardUnknown() {
	xxx_messageInfo_ListQueueBucket.DiscardUnknown(m)
}

var xxx_messageInfo_ListQueueBucket proto.InternalMessageInfo

func (m *ListQueueBucket) GetNext() bool {
	if m != nil {
		return m.Next
	}
	return false
}

func (m *ListQueueBucket) GetItems() []*QueueBucket {
	if m != nil {
		return m.Items
	}
	return nil
}

type ReadQueueBucketRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	QueueId              int64    `protobuf:"varint,2,opt,name=queue_id,json=queueId,proto3" json:"queue_id,omitempty"`
	DomainId             int64    `protobuf:"varint,3,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadQueueBucketRequest) Reset()         { *m = ReadQueueBucketRequest{} }
func (m *ReadQueueBucketRequest) String() string { return proto.CompactTextString(m) }
func (*ReadQueueBucketRequest) ProtoMessage()    {}
func (*ReadQueueBucketRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4c81630b355ff82, []int{4}
}

func (m *ReadQueueBucketRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadQueueBucketRequest.Unmarshal(m, b)
}
func (m *ReadQueueBucketRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadQueueBucketRequest.Marshal(b, m, deterministic)
}
func (m *ReadQueueBucketRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadQueueBucketRequest.Merge(m, src)
}
func (m *ReadQueueBucketRequest) XXX_Size() int {
	return xxx_messageInfo_ReadQueueBucketRequest.Size(m)
}
func (m *ReadQueueBucketRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadQueueBucketRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadQueueBucketRequest proto.InternalMessageInfo

func (m *ReadQueueBucketRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ReadQueueBucketRequest) GetQueueId() int64 {
	if m != nil {
		return m.QueueId
	}
	return 0
}

func (m *ReadQueueBucketRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

type CreateQueueBucketRequest struct {
	QueueId              int64    `protobuf:"varint,1,opt,name=queue_id,json=queueId,proto3" json:"queue_id,omitempty"`
	Ratio                int32    `protobuf:"varint,2,opt,name=ratio,proto3" json:"ratio,omitempty"`
	Bucket               *Lookup  `protobuf:"bytes,3,opt,name=bucket,proto3" json:"bucket,omitempty"`
	DomainId             int64    `protobuf:"varint,4,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateQueueBucketRequest) Reset()         { *m = CreateQueueBucketRequest{} }
func (m *CreateQueueBucketRequest) String() string { return proto.CompactTextString(m) }
func (*CreateQueueBucketRequest) ProtoMessage()    {}
func (*CreateQueueBucketRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4c81630b355ff82, []int{5}
}

func (m *CreateQueueBucketRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateQueueBucketRequest.Unmarshal(m, b)
}
func (m *CreateQueueBucketRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateQueueBucketRequest.Marshal(b, m, deterministic)
}
func (m *CreateQueueBucketRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateQueueBucketRequest.Merge(m, src)
}
func (m *CreateQueueBucketRequest) XXX_Size() int {
	return xxx_messageInfo_CreateQueueBucketRequest.Size(m)
}
func (m *CreateQueueBucketRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateQueueBucketRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateQueueBucketRequest proto.InternalMessageInfo

func (m *CreateQueueBucketRequest) GetQueueId() int64 {
	if m != nil {
		return m.QueueId
	}
	return 0
}

func (m *CreateQueueBucketRequest) GetRatio() int32 {
	if m != nil {
		return m.Ratio
	}
	return 0
}

func (m *CreateQueueBucketRequest) GetBucket() *Lookup {
	if m != nil {
		return m.Bucket
	}
	return nil
}

func (m *CreateQueueBucketRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

type QueueBucket struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Ratio                int32    `protobuf:"varint,2,opt,name=ratio,proto3" json:"ratio,omitempty"`
	Bucket               *Lookup  `protobuf:"bytes,3,opt,name=bucket,proto3" json:"bucket,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueueBucket) Reset()         { *m = QueueBucket{} }
func (m *QueueBucket) String() string { return proto.CompactTextString(m) }
func (*QueueBucket) ProtoMessage()    {}
func (*QueueBucket) Descriptor() ([]byte, []int) {
	return fileDescriptor_e4c81630b355ff82, []int{6}
}

func (m *QueueBucket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueueBucket.Unmarshal(m, b)
}
func (m *QueueBucket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueueBucket.Marshal(b, m, deterministic)
}
func (m *QueueBucket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueueBucket.Merge(m, src)
}
func (m *QueueBucket) XXX_Size() int {
	return xxx_messageInfo_QueueBucket.Size(m)
}
func (m *QueueBucket) XXX_DiscardUnknown() {
	xxx_messageInfo_QueueBucket.DiscardUnknown(m)
}

var xxx_messageInfo_QueueBucket proto.InternalMessageInfo

func (m *QueueBucket) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *QueueBucket) GetRatio() int32 {
	if m != nil {
		return m.Ratio
	}
	return 0
}

func (m *QueueBucket) GetBucket() *Lookup {
	if m != nil {
		return m.Bucket
	}
	return nil
}

func init() {
	proto.RegisterType((*DeleteQueueBucketRequest)(nil), "engine.DeleteQueueBucketRequest")
	proto.RegisterType((*UpdateQueueBucketRequest)(nil), "engine.UpdateQueueBucketRequest")
	proto.RegisterType((*SearchQueueBucketRequest)(nil), "engine.SearchQueueBucketRequest")
	proto.RegisterType((*ListQueueBucket)(nil), "engine.ListQueueBucket")
	proto.RegisterType((*ReadQueueBucketRequest)(nil), "engine.ReadQueueBucketRequest")
	proto.RegisterType((*CreateQueueBucketRequest)(nil), "engine.CreateQueueBucketRequest")
	proto.RegisterType((*QueueBucket)(nil), "engine.QueueBucket")
}

func init() { proto.RegisterFile("queue_bucket.proto", fileDescriptor_e4c81630b355ff82) }

var fileDescriptor_e4c81630b355ff82 = []byte{
	// 575 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x95, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xb5, 0x76, 0x9c, 0x26, 0x13, 0x68, 0x95, 0x05, 0x15, 0x13, 0x10, 0xb2, 0x7c, 0xa8,
	0x4c, 0x2b, 0x62, 0x48, 0x25, 0x0e, 0x1c, 0x0b, 0x97, 0x4a, 0x3d, 0x80, 0x2b, 0x2e, 0x70, 0x28,
	0x8e, 0x3d, 0xa4, 0xab, 0x3a, 0xde, 0xc4, 0xbb, 0x06, 0x44, 0x14, 0x90, 0xfa, 0x04, 0x48, 0x48,
	0x9c, 0xb9, 0xf1, 0x40, 0xbc, 0x02, 0x0f, 0x82, 0xbc, 0x4e, 0x1a, 0x27, 0x71, 0xaa, 0x54, 0x55,
	0x6f, 0xbb, 0x3b, 0xd6, 0xfc, 0xdf, 0xfc, 0x33, 0x1a, 0x03, 0x1d, 0xa6, 0x98, 0xe2, 0x49, 0x37,
	0x0d, 0xce, 0x50, 0xb6, 0x07, 0x09, 0x97, 0x9c, 0x56, 0x31, 0xee, 0xb1, 0x18, 0x5b, 0x8d, 0x80,
	0xc7, 0x62, 0xf2, 0xd8, 0x7a, 0xd8, 0xe3, 0xbc, 0x17, 0xa1, 0xeb, 0x0f, 0x98, 0xeb, 0xc7, 0x31,
	0x97, 0xbe, 0x64, 0x3c, 0x16, 0x79, 0xd4, 0xee, 0x82, 0xf9, 0x0a, 0x23, 0x94, 0xf8, 0x26, 0x4b,
	0x77, 0xa0, 0xb2, 0x79, 0x38, 0x4c, 0x51, 0x48, 0xba, 0x09, 0x1a, 0x0b, 0x4d, 0x62, 0x11, 0x47,
	0xf7, 0x34, 0x16, 0xd2, 0xfb, 0x50, 0xcb, 0x45, 0x59, 0x68, 0x6a, 0xea, 0x75, 0x43, 0xdd, 0x0f,
	0x43, 0xfa, 0x00, 0xea, 0x21, 0xef, 0xfb, 0x2c, 0xce, 0x62, 0xba, 0x8a, 0xd5, 0xf2, 0x87, 0xc3,
	0xd0, 0xfe, 0x4d, 0xc0, 0x7c, 0x3b, 0x08, 0xfd, 0xeb, 0x8a, 0xdc, 0x05, 0x23, 0xc9, 0xe0, 0x95,
	0x80, 0xe1, 0xe5, 0x17, 0xba, 0x03, 0xd5, 0xdc, 0x04, 0xb3, 0x62, 0x11, 0xa7, 0xd1, 0xd9, 0x6c,
	0xe7, 0x2e, 0xb4, 0x8f, 0x38, 0x3f, 0x4b, 0x07, 0xde, 0x24, 0x3a, 0x8f, 0x68, 0x2c, 0x20, 0xfe,
	0x21, 0x60, 0x1e, 0xa3, 0x9f, 0x04, 0xa7, 0x25, 0x88, 0x45, 0x24, 0x32, 0x8f, 0x44, 0xa1, 0x32,
	0xf0, 0x7b, 0xa8, 0x48, 0x0d, 0x4f, 0x9d, 0xb3, 0x37, 0xc1, 0xbe, 0xe2, 0x84, 0x52, 0x9d, 0xe9,
	0x2d, 0x20, 0x43, 0xc5, 0x57, 0xf7, 0xc8, 0x50, 0x7d, 0xc1, 0x13, 0xa9, 0x28, 0xea, 0x9e, 0x3a,
	0xd3, 0x6d, 0xa8, 0x7e, 0x64, 0x18, 0x85, 0xc2, 0xac, 0x5a, 0xba, 0x53, 0xf7, 0x26, 0xb7, 0x89,
	0x3f, 0x1b, 0x96, 0xee, 0xdc, 0xce, 0xfc, 0xb1, 0x5f, 0xc3, 0xd6, 0x11, 0x13, 0xb2, 0x80, 0x99,
	0xa5, 0x8b, 0xf1, 0x8b, 0x54, 0x6c, 0x35, 0x4f, 0x9d, 0xe9, 0x63, 0x30, 0x98, 0xc4, 0xbe, 0x30,
	0x35, 0x4b, 0x77, 0x1a, 0x9d, 0x3b, 0x53, 0x53, 0x8a, 0xe5, 0xe5, 0x5f, 0xd8, 0x1f, 0x60, 0xdb,
	0x43, 0x3f, 0xbc, 0xc1, 0x01, 0xf8, 0x41, 0xc0, 0x7c, 0x99, 0x60, 0xf9, 0x00, 0x5c, 0xe2, 0xee,
	0x45, 0xc3, 0xb5, 0xf2, 0x86, 0xeb, 0xeb, 0x37, 0xbc, 0xb2, 0x80, 0xf4, 0x1e, 0x1a, 0x45, 0x0b,
	0x17, 0x2b, 0xbd, 0x96, 0x72, 0xe7, 0x97, 0x01, 0xb4, 0x90, 0xfd, 0x18, 0x93, 0x4f, 0x2c, 0x40,
	0x3a, 0x86, 0xe6, 0x92, 0x0b, 0xd4, 0x9a, 0xe6, 0x58, 0x65, 0x50, 0xab, 0xac, 0x77, 0xf6, 0xb3,
	0xf3, 0xbf, 0xff, 0x7e, 0x6a, 0x7b, 0xf6, 0x8e, 0x1b, 0xf8, 0x51, 0x74, 0x12, 0x60, 0x2c, 0x31,
	0x71, 0x95, 0x71, 0xc2, 0x1d, 0x4d, 0x0d, 0x1d, 0xbb, 0x39, 0x91, 0x78, 0x41, 0x76, 0xe9, 0x37,
	0x68, 0x2e, 0x8d, 0xf8, 0x4c, 0x7e, 0xd5, 0xf4, 0xb7, 0xee, 0x5d, 0x14, 0x39, 0x3f, 0x76, 0x76,
	0x5b, 0x21, 0x38, 0x74, 0x4d, 0x04, 0x3a, 0x82, 0xad, 0x85, 0x39, 0xa3, 0x8f, 0xa6, 0xb9, 0xcb,
	0x07, 0xb0, 0xbc, 0xf4, 0x7d, 0xa5, 0xfb, 0x84, 0xee, 0xad, 0xa7, 0xeb, 0x8e, 0x58, 0x38, 0xa6,
	0xe7, 0x04, 0x9a, 0x4b, 0x3b, 0x68, 0x56, 0xfd, 0xaa, 0xf5, 0x54, 0x4e, 0xf0, 0x5c, 0x11, 0x3c,
	0x6d, 0x5d, 0x85, 0x20, 0xeb, 0xc0, 0x77, 0x68, 0x2e, 0x2d, 0xdb, 0x19, 0xc3, 0xaa, 0x3d, 0x7c,
	0xa9, 0x0b, 0xbb, 0x57, 0x61, 0x38, 0xb0, 0xdf, 0x59, 0x3d, 0x26, 0x4f, 0xd3, 0x6e, 0x3b, 0xe0,
	0x7d, 0xf7, 0x33, 0x76, 0x99, 0xc4, 0xc8, 0x55, 0x7f, 0x02, 0xe1, 0xe6, 0x22, 0xdd, 0xaa, 0xba,
	0xee, 0xff, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x56, 0x9e, 0x95, 0xb4, 0x61, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueueBucketServiceClient is the client API for QueueBucketService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueueBucketServiceClient interface {
	// Create QueueBucket
	CreateQueueBucket(ctx context.Context, in *CreateQueueBucketRequest, opts ...grpc.CallOption) (*QueueBucket, error)
	// SearchQueueRouting
	SearchQueueBucket(ctx context.Context, in *SearchQueueBucketRequest, opts ...grpc.CallOption) (*ListQueueBucket, error)
	// ReadQueueRouting
	ReadQueueBucket(ctx context.Context, in *ReadQueueBucketRequest, opts ...grpc.CallOption) (*QueueBucket, error)
	// UpdateQueueBucket
	UpdateQueueBucket(ctx context.Context, in *UpdateQueueBucketRequest, opts ...grpc.CallOption) (*QueueBucket, error)
	// DeleteQueueRouting
	DeleteQueueBucket(ctx context.Context, in *DeleteQueueBucketRequest, opts ...grpc.CallOption) (*QueueBucket, error)
}

type queueBucketServiceClient struct {
	cc *grpc.ClientConn
}

func NewQueueBucketServiceClient(cc *grpc.ClientConn) QueueBucketServiceClient {
	return &queueBucketServiceClient{cc}
}

func (c *queueBucketServiceClient) CreateQueueBucket(ctx context.Context, in *CreateQueueBucketRequest, opts ...grpc.CallOption) (*QueueBucket, error) {
	out := new(QueueBucket)
	err := c.cc.Invoke(ctx, "/engine.QueueBucketService/CreateQueueBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueBucketServiceClient) SearchQueueBucket(ctx context.Context, in *SearchQueueBucketRequest, opts ...grpc.CallOption) (*ListQueueBucket, error) {
	out := new(ListQueueBucket)
	err := c.cc.Invoke(ctx, "/engine.QueueBucketService/SearchQueueBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueBucketServiceClient) ReadQueueBucket(ctx context.Context, in *ReadQueueBucketRequest, opts ...grpc.CallOption) (*QueueBucket, error) {
	out := new(QueueBucket)
	err := c.cc.Invoke(ctx, "/engine.QueueBucketService/ReadQueueBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueBucketServiceClient) UpdateQueueBucket(ctx context.Context, in *UpdateQueueBucketRequest, opts ...grpc.CallOption) (*QueueBucket, error) {
	out := new(QueueBucket)
	err := c.cc.Invoke(ctx, "/engine.QueueBucketService/UpdateQueueBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueBucketServiceClient) DeleteQueueBucket(ctx context.Context, in *DeleteQueueBucketRequest, opts ...grpc.CallOption) (*QueueBucket, error) {
	out := new(QueueBucket)
	err := c.cc.Invoke(ctx, "/engine.QueueBucketService/DeleteQueueBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueueBucketServiceServer is the server API for QueueBucketService service.
type QueueBucketServiceServer interface {
	// Create QueueBucket
	CreateQueueBucket(context.Context, *CreateQueueBucketRequest) (*QueueBucket, error)
	// SearchQueueRouting
	SearchQueueBucket(context.Context, *SearchQueueBucketRequest) (*ListQueueBucket, error)
	// ReadQueueRouting
	ReadQueueBucket(context.Context, *ReadQueueBucketRequest) (*QueueBucket, error)
	// UpdateQueueBucket
	UpdateQueueBucket(context.Context, *UpdateQueueBucketRequest) (*QueueBucket, error)
	// DeleteQueueRouting
	DeleteQueueBucket(context.Context, *DeleteQueueBucketRequest) (*QueueBucket, error)
}

// UnimplementedQueueBucketServiceServer can be embedded to have forward compatible implementations.
type UnimplementedQueueBucketServiceServer struct {
}

func (*UnimplementedQueueBucketServiceServer) CreateQueueBucket(ctx context.Context, req *CreateQueueBucketRequest) (*QueueBucket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateQueueBucket not implemented")
}
func (*UnimplementedQueueBucketServiceServer) SearchQueueBucket(ctx context.Context, req *SearchQueueBucketRequest) (*ListQueueBucket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchQueueBucket not implemented")
}
func (*UnimplementedQueueBucketServiceServer) ReadQueueBucket(ctx context.Context, req *ReadQueueBucketRequest) (*QueueBucket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadQueueBucket not implemented")
}
func (*UnimplementedQueueBucketServiceServer) UpdateQueueBucket(ctx context.Context, req *UpdateQueueBucketRequest) (*QueueBucket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateQueueBucket not implemented")
}
func (*UnimplementedQueueBucketServiceServer) DeleteQueueBucket(ctx context.Context, req *DeleteQueueBucketRequest) (*QueueBucket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQueueBucket not implemented")
}

func RegisterQueueBucketServiceServer(s *grpc.Server, srv QueueBucketServiceServer) {
	s.RegisterService(&_QueueBucketService_serviceDesc, srv)
}

func _QueueBucketService_CreateQueueBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateQueueBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueBucketServiceServer).CreateQueueBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.QueueBucketService/CreateQueueBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueBucketServiceServer).CreateQueueBucket(ctx, req.(*CreateQueueBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueueBucketService_SearchQueueBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchQueueBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueBucketServiceServer).SearchQueueBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.QueueBucketService/SearchQueueBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueBucketServiceServer).SearchQueueBucket(ctx, req.(*SearchQueueBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueueBucketService_ReadQueueBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadQueueBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueBucketServiceServer).ReadQueueBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.QueueBucketService/ReadQueueBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueBucketServiceServer).ReadQueueBucket(ctx, req.(*ReadQueueBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueueBucketService_UpdateQueueBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateQueueBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueBucketServiceServer).UpdateQueueBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.QueueBucketService/UpdateQueueBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueBucketServiceServer).UpdateQueueBucket(ctx, req.(*UpdateQueueBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueueBucketService_DeleteQueueBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteQueueBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueBucketServiceServer).DeleteQueueBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.QueueBucketService/DeleteQueueBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueBucketServiceServer).DeleteQueueBucket(ctx, req.(*DeleteQueueBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QueueBucketService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "engine.QueueBucketService",
	HandlerType: (*QueueBucketServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateQueueBucket",
			Handler:    _QueueBucketService_CreateQueueBucket_Handler,
		},
		{
			MethodName: "SearchQueueBucket",
			Handler:    _QueueBucketService_SearchQueueBucket_Handler,
		},
		{
			MethodName: "ReadQueueBucket",
			Handler:    _QueueBucketService_ReadQueueBucket_Handler,
		},
		{
			MethodName: "UpdateQueueBucket",
			Handler:    _QueueBucketService_UpdateQueueBucket_Handler,
		},
		{
			MethodName: "DeleteQueueBucket",
			Handler:    _QueueBucketService_DeleteQueueBucket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "queue_bucket.proto",
}
