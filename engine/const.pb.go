// Code generated by protoc-gen-go. DO NOT EDIT.
// source: const.proto

package engine

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Lookup struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Lookup) Reset()         { *m = Lookup{} }
func (m *Lookup) String() string { return proto.CompactTextString(m) }
func (*Lookup) ProtoMessage()    {}
func (*Lookup) Descriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{0}
}

func (m *Lookup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Lookup.Unmarshal(m, b)
}
func (m *Lookup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Lookup.Marshal(b, m, deterministic)
}
func (m *Lookup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Lookup.Merge(m, src)
}
func (m *Lookup) XXX_Size() int {
	return xxx_messageInfo_Lookup.Size(m)
}
func (m *Lookup) XXX_DiscardUnknown() {
	xxx_messageInfo_Lookup.DiscardUnknown(m)
}

var xxx_messageInfo_Lookup proto.InternalMessageInfo

func (m *Lookup) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Lookup) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type FilterBetween struct {
	From                 int64    `protobuf:"varint,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   int64    `protobuf:"varint,2,opt,name=to,proto3" json:"to,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FilterBetween) Reset()         { *m = FilterBetween{} }
func (m *FilterBetween) String() string { return proto.CompactTextString(m) }
func (*FilterBetween) ProtoMessage()    {}
func (*FilterBetween) Descriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{1}
}

func (m *FilterBetween) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterBetween.Unmarshal(m, b)
}
func (m *FilterBetween) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterBetween.Marshal(b, m, deterministic)
}
func (m *FilterBetween) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterBetween.Merge(m, src)
}
func (m *FilterBetween) XXX_Size() int {
	return xxx_messageInfo_FilterBetween.Size(m)
}
func (m *FilterBetween) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterBetween.DiscardUnknown(m)
}

var xxx_messageInfo_FilterBetween proto.InternalMessageInfo

func (m *FilterBetween) GetFrom() int64 {
	if m != nil {
		return m.From
	}
	return 0
}

func (m *FilterBetween) GetTo() int64 {
	if m != nil {
		return m.To
	}
	return 0
}

type ListRequest struct {
	DomainId             int64    `protobuf:"varint,1,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	Size                 int32    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Page                 int32    `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{2}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

func (m *ListRequest) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *ListRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

type DomainRecord struct {
	DomainId             int64    `protobuf:"varint,1,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	CreatedAt            int64    `protobuf:"varint,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	CreatedBy            int64    `protobuf:"varint,3,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	UpdatedAt            int64    `protobuf:"varint,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	UpdatedBy            int64    `protobuf:"varint,5,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DomainRecord) Reset()         { *m = DomainRecord{} }
func (m *DomainRecord) String() string { return proto.CompactTextString(m) }
func (*DomainRecord) ProtoMessage()    {}
func (*DomainRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{3}
}

func (m *DomainRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DomainRecord.Unmarshal(m, b)
}
func (m *DomainRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DomainRecord.Marshal(b, m, deterministic)
}
func (m *DomainRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DomainRecord.Merge(m, src)
}
func (m *DomainRecord) XXX_Size() int {
	return xxx_messageInfo_DomainRecord.Size(m)
}
func (m *DomainRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_DomainRecord.DiscardUnknown(m)
}

var xxx_messageInfo_DomainRecord proto.InternalMessageInfo

func (m *DomainRecord) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

func (m *DomainRecord) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *DomainRecord) GetCreatedBy() int64 {
	if m != nil {
		return m.CreatedBy
	}
	return 0
}

func (m *DomainRecord) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *DomainRecord) GetUpdatedBy() int64 {
	if m != nil {
		return m.UpdatedBy
	}
	return 0
}

type ListForItemRequest struct {
	DomainId             int64    `protobuf:"varint,1,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	ItemId               int64    `protobuf:"varint,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	Size                 int32    `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Page                 int32    `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListForItemRequest) Reset()         { *m = ListForItemRequest{} }
func (m *ListForItemRequest) String() string { return proto.CompactTextString(m) }
func (*ListForItemRequest) ProtoMessage()    {}
func (*ListForItemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{4}
}

func (m *ListForItemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListForItemRequest.Unmarshal(m, b)
}
func (m *ListForItemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListForItemRequest.Marshal(b, m, deterministic)
}
func (m *ListForItemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListForItemRequest.Merge(m, src)
}
func (m *ListForItemRequest) XXX_Size() int {
	return xxx_messageInfo_ListForItemRequest.Size(m)
}
func (m *ListForItemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListForItemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListForItemRequest proto.InternalMessageInfo

func (m *ListForItemRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

func (m *ListForItemRequest) GetItemId() int64 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func (m *ListForItemRequest) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *ListForItemRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

type ItemRequest struct {
	DomainId             int64    `protobuf:"varint,1,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ItemRequest) Reset()         { *m = ItemRequest{} }
func (m *ItemRequest) String() string { return proto.CompactTextString(m) }
func (*ItemRequest) ProtoMessage()    {}
func (*ItemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{5}
}

func (m *ItemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemRequest.Unmarshal(m, b)
}
func (m *ItemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemRequest.Marshal(b, m, deterministic)
}
func (m *ItemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemRequest.Merge(m, src)
}
func (m *ItemRequest) XXX_Size() int {
	return xxx_messageInfo_ItemRequest.Size(m)
}
func (m *ItemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ItemRequest proto.InternalMessageInfo

func (m *ItemRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

func (m *ItemRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Response struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_5adb9555099c2688, []int{6}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*Lookup)(nil), "engine.Lookup")
	proto.RegisterType((*FilterBetween)(nil), "engine.FilterBetween")
	proto.RegisterType((*ListRequest)(nil), "engine.ListRequest")
	proto.RegisterType((*DomainRecord)(nil), "engine.DomainRecord")
	proto.RegisterType((*ListForItemRequest)(nil), "engine.ListForItemRequest")
	proto.RegisterType((*ItemRequest)(nil), "engine.ItemRequest")
	proto.RegisterType((*Response)(nil), "engine.Response")
}

func init() { proto.RegisterFile("const.proto", fileDescriptor_5adb9555099c2688) }

var fileDescriptor_5adb9555099c2688 = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x6b, 0xe3, 0x30,
	0x14, 0xc4, 0x76, 0xe2, 0x8d, 0x5f, 0x76, 0xf7, 0xe0, 0xc3, 0xae, 0xa1, 0x14, 0x82, 0x4e, 0x39,
	0x94, 0xe4, 0x90, 0x5b, 0x6f, 0x35, 0x25, 0x10, 0xc8, 0x49, 0xc7, 0x5e, 0x82, 0x3f, 0x5e, 0x53,
	0xd1, 0x58, 0x72, 0xa5, 0x67, 0x82, 0xfb, 0x87, 0xfa, 0x37, 0x8b, 0x24, 0xa7, 0x4d, 0xa1, 0x94,
	0xdc, 0xe6, 0xcd, 0x78, 0x34, 0xf2, 0xe8, 0xc1, 0xb4, 0x52, 0xd2, 0xd0, 0xa2, 0xd5, 0x8a, 0x54,
	0x1a, 0xa3, 0xdc, 0x0b, 0x89, 0xec, 0x06, 0xe2, 0xad, 0x52, 0xcf, 0x5d, 0x9b, 0xfe, 0x85, 0x50,
	0xd4, 0x59, 0x30, 0x0b, 0xe6, 0x11, 0x0f, 0x45, 0x9d, 0xa6, 0x30, 0x92, 0x45, 0x83, 0x59, 0x38,
	0x0b, 0xe6, 0x09, 0x77, 0x98, 0xad, 0xe0, 0xcf, 0x5a, 0x1c, 0x08, 0x75, 0x8e, 0x74, 0x44, 0x94,
	0xf6, 0xa3, 0x47, 0xad, 0x9a, 0xc1, 0xe6, 0xb0, 0x3d, 0x88, 0x94, 0xb3, 0x45, 0x3c, 0x24, 0xc5,
	0x38, 0x4c, 0xb7, 0xc2, 0x10, 0xc7, 0x97, 0x0e, 0x0d, 0xa5, 0x57, 0x90, 0xd4, 0xaa, 0x29, 0x84,
	0xdc, 0x7d, 0xc4, 0x4d, 0x3c, 0xb1, 0x71, 0xa1, 0x46, 0xbc, 0xfa, 0xd0, 0x31, 0x77, 0xd8, 0x72,
	0x6d, 0xb1, 0xc7, 0x2c, 0xf2, 0x9c, 0xc5, 0xec, 0x2d, 0x80, 0xdf, 0xf7, 0xce, 0xc4, 0xb1, 0x52,
	0xba, 0xfe, 0xf9, 0xd4, 0x6b, 0x80, 0x4a, 0x63, 0x41, 0x58, 0xef, 0x0a, 0x1a, 0x6e, 0x96, 0x0c,
	0xcc, 0x1d, 0x9d, 0xcb, 0x65, 0xef, 0x62, 0x3e, 0xe5, 0xbc, 0xb7, 0x72, 0xd7, 0xd6, 0x27, 0xf7,
	0xc8, 0xcb, 0x03, 0xe3, 0xdd, 0x27, 0xb9, 0xec, 0xb3, 0xf1, 0x17, 0x39, 0xef, 0x99, 0x86, 0xd4,
	0xfe, 0xfd, 0x5a, 0xe9, 0x0d, 0x61, 0x73, 0x51, 0x09, 0xff, 0xe1, 0x97, 0x20, 0x6c, 0xac, 0xe4,
	0xef, 0x1a, 0xdb, 0xf1, 0xac, 0x9d, 0xe8, 0x9b, 0x76, 0x46, 0x67, 0xed, 0xdc, 0xc2, 0xf4, 0xe2,
	0x30, 0xff, 0xec, 0xe1, 0xe9, 0xd9, 0x19, 0x83, 0x09, 0x47, 0xd3, 0x2a, 0x69, 0x30, 0xfd, 0x07,
	0xb1, 0xa1, 0x82, 0x3a, 0xe3, 0x5c, 0x09, 0x1f, 0xa6, 0x9c, 0x3d, 0xcc, 0xf6, 0x82, 0x9e, 0xba,
	0x72, 0x51, 0xa9, 0x66, 0x79, 0xc4, 0x52, 0x10, 0x1e, 0x96, 0x6e, 0xb1, 0xcc, 0xd2, 0x2f, 0x56,
	0x19, 0xbb, 0x71, 0xf5, 0x1e, 0x00, 0x00, 0xff, 0xff, 0x07, 0x95, 0xe9, 0x1f, 0x76, 0x02, 0x00,
	0x00,
}
