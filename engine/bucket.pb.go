// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bucket.proto

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

type DeleteBucketRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DomainId             int64    `protobuf:"varint,2,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteBucketRequest) Reset()         { *m = DeleteBucketRequest{} }
func (m *DeleteBucketRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteBucketRequest) ProtoMessage()    {}
func (*DeleteBucketRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e8f4be413d27239, []int{0}
}

func (m *DeleteBucketRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteBucketRequest.Unmarshal(m, b)
}
func (m *DeleteBucketRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteBucketRequest.Marshal(b, m, deterministic)
}
func (m *DeleteBucketRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteBucketRequest.Merge(m, src)
}
func (m *DeleteBucketRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteBucketRequest.Size(m)
}
func (m *DeleteBucketRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteBucketRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteBucketRequest proto.InternalMessageInfo

func (m *DeleteBucketRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DeleteBucketRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

type UpdateBucketRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	DomainId             int64    `protobuf:"varint,4,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateBucketRequest) Reset()         { *m = UpdateBucketRequest{} }
func (m *UpdateBucketRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateBucketRequest) ProtoMessage()    {}
func (*UpdateBucketRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e8f4be413d27239, []int{1}
}

func (m *UpdateBucketRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateBucketRequest.Unmarshal(m, b)
}
func (m *UpdateBucketRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateBucketRequest.Marshal(b, m, deterministic)
}
func (m *UpdateBucketRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateBucketRequest.Merge(m, src)
}
func (m *UpdateBucketRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateBucketRequest.Size(m)
}
func (m *UpdateBucketRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateBucketRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateBucketRequest proto.InternalMessageInfo

func (m *UpdateBucketRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateBucketRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateBucketRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *UpdateBucketRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

type SearchBucketRequest struct {
	Page                 int32    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size                 int32    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Q                    string   `protobuf:"bytes,3,opt,name=q,proto3" json:"q,omitempty"`
	Sort                 string   `protobuf:"bytes,4,opt,name=sort,proto3" json:"sort,omitempty"`
	Fields               []string `protobuf:"bytes,5,rep,name=fields,proto3" json:"fields,omitempty"`
	Id                   []uint32 `protobuf:"varint,6,rep,packed,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchBucketRequest) Reset()         { *m = SearchBucketRequest{} }
func (m *SearchBucketRequest) String() string { return proto.CompactTextString(m) }
func (*SearchBucketRequest) ProtoMessage()    {}
func (*SearchBucketRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e8f4be413d27239, []int{2}
}

func (m *SearchBucketRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchBucketRequest.Unmarshal(m, b)
}
func (m *SearchBucketRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchBucketRequest.Marshal(b, m, deterministic)
}
func (m *SearchBucketRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchBucketRequest.Merge(m, src)
}
func (m *SearchBucketRequest) XXX_Size() int {
	return xxx_messageInfo_SearchBucketRequest.Size(m)
}
func (m *SearchBucketRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchBucketRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchBucketRequest proto.InternalMessageInfo

func (m *SearchBucketRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *SearchBucketRequest) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *SearchBucketRequest) GetQ() string {
	if m != nil {
		return m.Q
	}
	return ""
}

func (m *SearchBucketRequest) GetSort() string {
	if m != nil {
		return m.Sort
	}
	return ""
}

func (m *SearchBucketRequest) GetFields() []string {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *SearchBucketRequest) GetId() []uint32 {
	if m != nil {
		return m.Id
	}
	return nil
}

type ListBucket struct {
	Next                 bool      `protobuf:"varint,1,opt,name=next,proto3" json:"next,omitempty"`
	Items                []*Bucket `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListBucket) Reset()         { *m = ListBucket{} }
func (m *ListBucket) String() string { return proto.CompactTextString(m) }
func (*ListBucket) ProtoMessage()    {}
func (*ListBucket) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e8f4be413d27239, []int{3}
}

func (m *ListBucket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBucket.Unmarshal(m, b)
}
func (m *ListBucket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBucket.Marshal(b, m, deterministic)
}
func (m *ListBucket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBucket.Merge(m, src)
}
func (m *ListBucket) XXX_Size() int {
	return xxx_messageInfo_ListBucket.Size(m)
}
func (m *ListBucket) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBucket.DiscardUnknown(m)
}

var xxx_messageInfo_ListBucket proto.InternalMessageInfo

func (m *ListBucket) GetNext() bool {
	if m != nil {
		return m.Next
	}
	return false
}

func (m *ListBucket) GetItems() []*Bucket {
	if m != nil {
		return m.Items
	}
	return nil
}

type ReadBucketRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DomainId             int64    `protobuf:"varint,2,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadBucketRequest) Reset()         { *m = ReadBucketRequest{} }
func (m *ReadBucketRequest) String() string { return proto.CompactTextString(m) }
func (*ReadBucketRequest) ProtoMessage()    {}
func (*ReadBucketRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e8f4be413d27239, []int{4}
}

func (m *ReadBucketRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadBucketRequest.Unmarshal(m, b)
}
func (m *ReadBucketRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadBucketRequest.Marshal(b, m, deterministic)
}
func (m *ReadBucketRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadBucketRequest.Merge(m, src)
}
func (m *ReadBucketRequest) XXX_Size() int {
	return xxx_messageInfo_ReadBucketRequest.Size(m)
}
func (m *ReadBucketRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadBucketRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadBucketRequest proto.InternalMessageInfo

func (m *ReadBucketRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ReadBucketRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

type CreateBucketRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	DomainId             int64    `protobuf:"varint,3,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBucketRequest) Reset()         { *m = CreateBucketRequest{} }
func (m *CreateBucketRequest) String() string { return proto.CompactTextString(m) }
func (*CreateBucketRequest) ProtoMessage()    {}
func (*CreateBucketRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e8f4be413d27239, []int{5}
}

func (m *CreateBucketRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBucketRequest.Unmarshal(m, b)
}
func (m *CreateBucketRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBucketRequest.Marshal(b, m, deterministic)
}
func (m *CreateBucketRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBucketRequest.Merge(m, src)
}
func (m *CreateBucketRequest) XXX_Size() int {
	return xxx_messageInfo_CreateBucketRequest.Size(m)
}
func (m *CreateBucketRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBucketRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBucketRequest proto.InternalMessageInfo

func (m *CreateBucketRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateBucketRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreateBucketRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

type Bucket struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bucket) Reset()         { *m = Bucket{} }
func (m *Bucket) String() string { return proto.CompactTextString(m) }
func (*Bucket) ProtoMessage()    {}
func (*Bucket) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e8f4be413d27239, []int{6}
}

func (m *Bucket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bucket.Unmarshal(m, b)
}
func (m *Bucket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bucket.Marshal(b, m, deterministic)
}
func (m *Bucket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bucket.Merge(m, src)
}
func (m *Bucket) XXX_Size() int {
	return xxx_messageInfo_Bucket.Size(m)
}
func (m *Bucket) XXX_DiscardUnknown() {
	xxx_messageInfo_Bucket.DiscardUnknown(m)
}

var xxx_messageInfo_Bucket proto.InternalMessageInfo

func (m *Bucket) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Bucket) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Bucket) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*DeleteBucketRequest)(nil), "engine.DeleteBucketRequest")
	proto.RegisterType((*UpdateBucketRequest)(nil), "engine.UpdateBucketRequest")
	proto.RegisterType((*SearchBucketRequest)(nil), "engine.SearchBucketRequest")
	proto.RegisterType((*ListBucket)(nil), "engine.ListBucket")
	proto.RegisterType((*ReadBucketRequest)(nil), "engine.ReadBucketRequest")
	proto.RegisterType((*CreateBucketRequest)(nil), "engine.CreateBucketRequest")
	proto.RegisterType((*Bucket)(nil), "engine.Bucket")
}

func init() { proto.RegisterFile("bucket.proto", fileDescriptor_8e8f4be413d27239) }

var fileDescriptor_8e8f4be413d27239 = []byte{
	// 509 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0xed, 0xc4, 0x6a, 0xa6, 0x6e, 0x25, 0x36, 0xa8, 0x72, 0x92, 0x0a, 0x8c, 0xd5, 0x43,
	0xd4, 0x43, 0x2c, 0x95, 0x1b, 0x27, 0x14, 0x10, 0x12, 0x12, 0xe2, 0xe0, 0x8a, 0x4b, 0x05, 0x54,
	0x1b, 0xef, 0xe0, 0xac, 0x70, 0x6c, 0xc7, 0xde, 0x40, 0x05, 0xe2, 0xc2, 0x85, 0x0f, 0xe0, 0xd3,
	0xf8, 0x05, 0x6e, 0xfc, 0x04, 0xca, 0xac, 0x23, 0xbb, 0x89, 0xa3, 0x1e, 0xe0, 0x36, 0x3b, 0xbb,
	0xfb, 0xde, 0xbc, 0x79, 0xbb, 0x03, 0xce, 0x6c, 0x15, 0x7d, 0x44, 0x35, 0xc9, 0x8b, 0x4c, 0x65,
	0xcc, 0xc6, 0x34, 0x96, 0x29, 0x0e, 0x4f, 0xe3, 0x2c, 0x8b, 0x13, 0x0c, 0x78, 0x2e, 0x03, 0x9e,
	0xa6, 0x99, 0xe2, 0x4a, 0x66, 0x69, 0xa9, 0x4f, 0xf9, 0x53, 0xe8, 0x3f, 0xc7, 0x04, 0x15, 0x4e,
	0xe9, 0x6e, 0x88, 0xcb, 0x15, 0x96, 0x8a, 0x1d, 0x83, 0x29, 0x85, 0x6b, 0x78, 0xc6, 0xd8, 0x0a,
	0x4d, 0x29, 0xd8, 0x08, 0x7a, 0x22, 0x5b, 0x70, 0x99, 0x5e, 0x4b, 0xe1, 0x9a, 0x94, 0x3e, 0xd0,
	0x89, 0x97, 0xc2, 0xbf, 0x81, 0xfe, 0x9b, 0x5c, 0xf0, 0xbb, 0x30, 0x18, 0x74, 0x52, 0xbe, 0x40,
	0xba, 0xde, 0x0b, 0x29, 0x66, 0x1e, 0x1c, 0x0a, 0x2c, 0xa3, 0x42, 0xe6, 0xeb, 0xa2, 0x5c, 0x8b,
	0xb6, 0x9a, 0xa9, 0xdb, 0xcc, 0x9d, 0x2d, 0xe6, 0x1f, 0x06, 0xf4, 0x2f, 0x91, 0x17, 0xd1, 0xfc,
	0x36, 0x35, 0x83, 0x4e, 0xce, 0x63, 0x24, 0xf2, 0x6e, 0x48, 0xf1, 0x3a, 0x57, 0xca, 0x2f, 0x9a,
	0xbe, 0x1b, 0x52, 0xcc, 0x1c, 0x30, 0x96, 0x15, 0xa9, 0xb1, 0xa4, 0x13, 0x59, 0xa1, 0x88, 0xa5,
	0x17, 0x52, 0xcc, 0x4e, 0xc0, 0xfe, 0x20, 0x31, 0x11, 0xa5, 0xdb, 0xf5, 0xac, 0x71, 0x2f, 0xac,
	0x56, 0x95, 0x38, 0xdb, 0xb3, 0xc6, 0x47, 0x6b, 0x71, 0xfe, 0x0b, 0x80, 0x57, 0xb2, 0x54, 0xba,
	0x0c, 0x92, 0x8a, 0x37, 0x8a, 0xf8, 0x0f, 0x42, 0x8a, 0xd9, 0x19, 0x74, 0xa5, 0xc2, 0x45, 0xe9,
	0x9a, 0x9e, 0x35, 0x3e, 0xbc, 0x38, 0x9e, 0x68, 0x7f, 0x26, 0x55, 0xe5, 0x7a, 0xd3, 0x7f, 0x0a,
	0xf7, 0x42, 0xe4, 0xe2, 0x1f, 0xdc, 0x98, 0x43, 0xff, 0x59, 0x81, 0x3b, 0x6e, 0x6c, 0xba, 0x6f,
	0xec, 0xef, 0xbe, 0x79, 0x47, 0xf7, 0xad, 0x2d, 0xa6, 0xd7, 0x60, 0x57, 0x7a, 0xff, 0x8b, 0xd5,
	0x17, 0x7f, 0x2c, 0x38, 0xd2, 0x80, 0x97, 0x58, 0x7c, 0x92, 0x11, 0xb2, 0xb7, 0xe0, 0x34, 0xb5,
	0xb0, 0xd1, 0xa6, 0x69, 0x2d, 0x0a, 0x87, 0x5b, 0x1d, 0xf5, 0x1f, 0x7e, 0xff, 0xf5, 0xfb, 0xa7,
	0x39, 0xf0, 0xef, 0x07, 0x11, 0x4f, 0x92, 0xeb, 0x08, 0x53, 0x85, 0x45, 0xa0, 0xff, 0x48, 0xf9,
	0xc4, 0x38, 0x67, 0xef, 0xc0, 0x69, 0x3e, 0x9e, 0x1a, 0xbd, 0xe5, 0x49, 0x0d, 0xd9, 0x66, 0xb3,
	0xb6, 0xd9, 0x3f, 0x25, 0x86, 0x13, 0xd6, 0xca, 0xc0, 0xae, 0x00, 0x6a, 0x2b, 0xd9, 0x60, 0x73,
	0x7f, 0xc7, 0xde, 0x9d, 0xc2, 0x1f, 0x11, 0xec, 0x88, 0x0d, 0xda, 0x60, 0x83, 0xaf, 0x52, 0x7c,
	0x63, 0x1c, 0x9c, 0xe6, 0x97, 0xab, 0x4b, 0x6f, 0xf9, 0x88, 0x3b, 0xf8, 0x67, 0x84, 0xff, 0x60,
	0xb8, 0x1f, 0x7f, 0xdd, 0x9d, 0xf7, 0xe0, 0x34, 0x27, 0x43, 0x4d, 0xd1, 0x32, 0x2f, 0xf6, 0x49,
	0x38, 0xdf, 0x4f, 0x31, 0xf5, 0xaf, 0xbc, 0x58, 0xaa, 0xf9, 0x6a, 0x36, 0x89, 0xb2, 0x45, 0xf0,
	0x19, 0x67, 0x52, 0x61, 0x12, 0xd0, 0x54, 0x2a, 0x03, 0x8d, 0x36, 0xb3, 0x69, 0xf9, 0xf8, 0x6f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xf3, 0xa9, 0x60, 0x5c, 0xda, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BucketServiceClient is the client API for BucketService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BucketServiceClient interface {
	// Create Bucket
	CreateBucket(ctx context.Context, in *CreateBucketRequest, opts ...grpc.CallOption) (*Bucket, error)
	// List of Bucket
	SearchBucket(ctx context.Context, in *SearchBucketRequest, opts ...grpc.CallOption) (*ListBucket, error)
	// Bucket item
	ReadBucket(ctx context.Context, in *ReadBucketRequest, opts ...grpc.CallOption) (*Bucket, error)
	// Update Bucket
	UpdateBucket(ctx context.Context, in *UpdateBucketRequest, opts ...grpc.CallOption) (*Bucket, error)
	// Remove Bucket
	DeleteBucket(ctx context.Context, in *DeleteBucketRequest, opts ...grpc.CallOption) (*Bucket, error)
}

type bucketServiceClient struct {
	cc *grpc.ClientConn
}

func NewBucketServiceClient(cc *grpc.ClientConn) BucketServiceClient {
	return &bucketServiceClient{cc}
}

func (c *bucketServiceClient) CreateBucket(ctx context.Context, in *CreateBucketRequest, opts ...grpc.CallOption) (*Bucket, error) {
	out := new(Bucket)
	err := c.cc.Invoke(ctx, "/engine.BucketService/CreateBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bucketServiceClient) SearchBucket(ctx context.Context, in *SearchBucketRequest, opts ...grpc.CallOption) (*ListBucket, error) {
	out := new(ListBucket)
	err := c.cc.Invoke(ctx, "/engine.BucketService/SearchBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bucketServiceClient) ReadBucket(ctx context.Context, in *ReadBucketRequest, opts ...grpc.CallOption) (*Bucket, error) {
	out := new(Bucket)
	err := c.cc.Invoke(ctx, "/engine.BucketService/ReadBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bucketServiceClient) UpdateBucket(ctx context.Context, in *UpdateBucketRequest, opts ...grpc.CallOption) (*Bucket, error) {
	out := new(Bucket)
	err := c.cc.Invoke(ctx, "/engine.BucketService/UpdateBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bucketServiceClient) DeleteBucket(ctx context.Context, in *DeleteBucketRequest, opts ...grpc.CallOption) (*Bucket, error) {
	out := new(Bucket)
	err := c.cc.Invoke(ctx, "/engine.BucketService/DeleteBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BucketServiceServer is the server API for BucketService service.
type BucketServiceServer interface {
	// Create Bucket
	CreateBucket(context.Context, *CreateBucketRequest) (*Bucket, error)
	// List of Bucket
	SearchBucket(context.Context, *SearchBucketRequest) (*ListBucket, error)
	// Bucket item
	ReadBucket(context.Context, *ReadBucketRequest) (*Bucket, error)
	// Update Bucket
	UpdateBucket(context.Context, *UpdateBucketRequest) (*Bucket, error)
	// Remove Bucket
	DeleteBucket(context.Context, *DeleteBucketRequest) (*Bucket, error)
}

// UnimplementedBucketServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBucketServiceServer struct {
}

func (*UnimplementedBucketServiceServer) CreateBucket(ctx context.Context, req *CreateBucketRequest) (*Bucket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBucket not implemented")
}
func (*UnimplementedBucketServiceServer) SearchBucket(ctx context.Context, req *SearchBucketRequest) (*ListBucket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchBucket not implemented")
}
func (*UnimplementedBucketServiceServer) ReadBucket(ctx context.Context, req *ReadBucketRequest) (*Bucket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadBucket not implemented")
}
func (*UnimplementedBucketServiceServer) UpdateBucket(ctx context.Context, req *UpdateBucketRequest) (*Bucket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBucket not implemented")
}
func (*UnimplementedBucketServiceServer) DeleteBucket(ctx context.Context, req *DeleteBucketRequest) (*Bucket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBucket not implemented")
}

func RegisterBucketServiceServer(s *grpc.Server, srv BucketServiceServer) {
	s.RegisterService(&_BucketService_serviceDesc, srv)
}

func _BucketService_CreateBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BucketServiceServer).CreateBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.BucketService/CreateBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BucketServiceServer).CreateBucket(ctx, req.(*CreateBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BucketService_SearchBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BucketServiceServer).SearchBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.BucketService/SearchBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BucketServiceServer).SearchBucket(ctx, req.(*SearchBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BucketService_ReadBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BucketServiceServer).ReadBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.BucketService/ReadBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BucketServiceServer).ReadBucket(ctx, req.(*ReadBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BucketService_UpdateBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BucketServiceServer).UpdateBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.BucketService/UpdateBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BucketServiceServer).UpdateBucket(ctx, req.(*UpdateBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BucketService_DeleteBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BucketServiceServer).DeleteBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.BucketService/DeleteBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BucketServiceServer).DeleteBucket(ctx, req.(*DeleteBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BucketService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "engine.BucketService",
	HandlerType: (*BucketServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBucket",
			Handler:    _BucketService_CreateBucket_Handler,
		},
		{
			MethodName: "SearchBucket",
			Handler:    _BucketService_SearchBucket_Handler,
		},
		{
			MethodName: "ReadBucket",
			Handler:    _BucketService_ReadBucket_Handler,
		},
		{
			MethodName: "UpdateBucket",
			Handler:    _BucketService_UpdateBucket_Handler,
		},
		{
			MethodName: "DeleteBucket",
			Handler:    _BucketService_DeleteBucket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bucket.proto",
}
