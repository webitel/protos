// Code generated by protoc-gen-go. DO NOT EDIT.
// source: queue_communications.proto

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

type DeleteCommunicationTypeRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DomainId             int64    `protobuf:"varint,2,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteCommunicationTypeRequest) Reset()         { *m = DeleteCommunicationTypeRequest{} }
func (m *DeleteCommunicationTypeRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteCommunicationTypeRequest) ProtoMessage()    {}
func (*DeleteCommunicationTypeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_090a32393dcacaf2, []int{0}
}

func (m *DeleteCommunicationTypeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteCommunicationTypeRequest.Unmarshal(m, b)
}
func (m *DeleteCommunicationTypeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteCommunicationTypeRequest.Marshal(b, m, deterministic)
}
func (m *DeleteCommunicationTypeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteCommunicationTypeRequest.Merge(m, src)
}
func (m *DeleteCommunicationTypeRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteCommunicationTypeRequest.Size(m)
}
func (m *DeleteCommunicationTypeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteCommunicationTypeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteCommunicationTypeRequest proto.InternalMessageInfo

func (m *DeleteCommunicationTypeRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DeleteCommunicationTypeRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

type UpdateCommunicationTypeRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Code                 string   `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Type                 string   `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Description          string   `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	DomainId             int64    `protobuf:"varint,6,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateCommunicationTypeRequest) Reset()         { *m = UpdateCommunicationTypeRequest{} }
func (m *UpdateCommunicationTypeRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateCommunicationTypeRequest) ProtoMessage()    {}
func (*UpdateCommunicationTypeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_090a32393dcacaf2, []int{1}
}

func (m *UpdateCommunicationTypeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCommunicationTypeRequest.Unmarshal(m, b)
}
func (m *UpdateCommunicationTypeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCommunicationTypeRequest.Marshal(b, m, deterministic)
}
func (m *UpdateCommunicationTypeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCommunicationTypeRequest.Merge(m, src)
}
func (m *UpdateCommunicationTypeRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateCommunicationTypeRequest.Size(m)
}
func (m *UpdateCommunicationTypeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCommunicationTypeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCommunicationTypeRequest proto.InternalMessageInfo

func (m *UpdateCommunicationTypeRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateCommunicationTypeRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateCommunicationTypeRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *UpdateCommunicationTypeRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *UpdateCommunicationTypeRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *UpdateCommunicationTypeRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

type ListCommunicationType struct {
	Next                 bool                 `protobuf:"varint,1,opt,name=next,proto3" json:"next,omitempty"`
	Items                []*CommunicationType `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ListCommunicationType) Reset()         { *m = ListCommunicationType{} }
func (m *ListCommunicationType) String() string { return proto.CompactTextString(m) }
func (*ListCommunicationType) ProtoMessage()    {}
func (*ListCommunicationType) Descriptor() ([]byte, []int) {
	return fileDescriptor_090a32393dcacaf2, []int{2}
}

func (m *ListCommunicationType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCommunicationType.Unmarshal(m, b)
}
func (m *ListCommunicationType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCommunicationType.Marshal(b, m, deterministic)
}
func (m *ListCommunicationType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCommunicationType.Merge(m, src)
}
func (m *ListCommunicationType) XXX_Size() int {
	return xxx_messageInfo_ListCommunicationType.Size(m)
}
func (m *ListCommunicationType) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCommunicationType.DiscardUnknown(m)
}

var xxx_messageInfo_ListCommunicationType proto.InternalMessageInfo

func (m *ListCommunicationType) GetNext() bool {
	if m != nil {
		return m.Next
	}
	return false
}

func (m *ListCommunicationType) GetItems() []*CommunicationType {
	if m != nil {
		return m.Items
	}
	return nil
}

type SearchCommunicationTypeRequest struct {
	Page                 int32    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size                 int32    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Q                    string   `protobuf:"bytes,3,opt,name=q,proto3" json:"q,omitempty"`
	DomainId             int64    `protobuf:"varint,4,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchCommunicationTypeRequest) Reset()         { *m = SearchCommunicationTypeRequest{} }
func (m *SearchCommunicationTypeRequest) String() string { return proto.CompactTextString(m) }
func (*SearchCommunicationTypeRequest) ProtoMessage()    {}
func (*SearchCommunicationTypeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_090a32393dcacaf2, []int{3}
}

func (m *SearchCommunicationTypeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchCommunicationTypeRequest.Unmarshal(m, b)
}
func (m *SearchCommunicationTypeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchCommunicationTypeRequest.Marshal(b, m, deterministic)
}
func (m *SearchCommunicationTypeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchCommunicationTypeRequest.Merge(m, src)
}
func (m *SearchCommunicationTypeRequest) XXX_Size() int {
	return xxx_messageInfo_SearchCommunicationTypeRequest.Size(m)
}
func (m *SearchCommunicationTypeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchCommunicationTypeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchCommunicationTypeRequest proto.InternalMessageInfo

func (m *SearchCommunicationTypeRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *SearchCommunicationTypeRequest) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *SearchCommunicationTypeRequest) GetQ() string {
	if m != nil {
		return m.Q
	}
	return ""
}

func (m *SearchCommunicationTypeRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

type ReadCommunicationTypeRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DomainId             int64    `protobuf:"varint,2,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadCommunicationTypeRequest) Reset()         { *m = ReadCommunicationTypeRequest{} }
func (m *ReadCommunicationTypeRequest) String() string { return proto.CompactTextString(m) }
func (*ReadCommunicationTypeRequest) ProtoMessage()    {}
func (*ReadCommunicationTypeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_090a32393dcacaf2, []int{4}
}

func (m *ReadCommunicationTypeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadCommunicationTypeRequest.Unmarshal(m, b)
}
func (m *ReadCommunicationTypeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadCommunicationTypeRequest.Marshal(b, m, deterministic)
}
func (m *ReadCommunicationTypeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadCommunicationTypeRequest.Merge(m, src)
}
func (m *ReadCommunicationTypeRequest) XXX_Size() int {
	return xxx_messageInfo_ReadCommunicationTypeRequest.Size(m)
}
func (m *ReadCommunicationTypeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadCommunicationTypeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadCommunicationTypeRequest proto.InternalMessageInfo

func (m *ReadCommunicationTypeRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ReadCommunicationTypeRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

type CommunicationTypeRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	DomainId             int64    `protobuf:"varint,5,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommunicationTypeRequest) Reset()         { *m = CommunicationTypeRequest{} }
func (m *CommunicationTypeRequest) String() string { return proto.CompactTextString(m) }
func (*CommunicationTypeRequest) ProtoMessage()    {}
func (*CommunicationTypeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_090a32393dcacaf2, []int{5}
}

func (m *CommunicationTypeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommunicationTypeRequest.Unmarshal(m, b)
}
func (m *CommunicationTypeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommunicationTypeRequest.Marshal(b, m, deterministic)
}
func (m *CommunicationTypeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommunicationTypeRequest.Merge(m, src)
}
func (m *CommunicationTypeRequest) XXX_Size() int {
	return xxx_messageInfo_CommunicationTypeRequest.Size(m)
}
func (m *CommunicationTypeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CommunicationTypeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CommunicationTypeRequest proto.InternalMessageInfo

func (m *CommunicationTypeRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CommunicationTypeRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *CommunicationTypeRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *CommunicationTypeRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CommunicationTypeRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

type CommunicationType struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DomainId             int64    `protobuf:"varint,2,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Code                 string   `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	Type                 string   `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	Description          string   `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommunicationType) Reset()         { *m = CommunicationType{} }
func (m *CommunicationType) String() string { return proto.CompactTextString(m) }
func (*CommunicationType) ProtoMessage()    {}
func (*CommunicationType) Descriptor() ([]byte, []int) {
	return fileDescriptor_090a32393dcacaf2, []int{6}
}

func (m *CommunicationType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommunicationType.Unmarshal(m, b)
}
func (m *CommunicationType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommunicationType.Marshal(b, m, deterministic)
}
func (m *CommunicationType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommunicationType.Merge(m, src)
}
func (m *CommunicationType) XXX_Size() int {
	return xxx_messageInfo_CommunicationType.Size(m)
}
func (m *CommunicationType) XXX_DiscardUnknown() {
	xxx_messageInfo_CommunicationType.DiscardUnknown(m)
}

var xxx_messageInfo_CommunicationType proto.InternalMessageInfo

func (m *CommunicationType) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CommunicationType) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

func (m *CommunicationType) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CommunicationType) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *CommunicationType) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *CommunicationType) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*DeleteCommunicationTypeRequest)(nil), "engine.DeleteCommunicationTypeRequest")
	proto.RegisterType((*UpdateCommunicationTypeRequest)(nil), "engine.UpdateCommunicationTypeRequest")
	proto.RegisterType((*ListCommunicationType)(nil), "engine.ListCommunicationType")
	proto.RegisterType((*SearchCommunicationTypeRequest)(nil), "engine.SearchCommunicationTypeRequest")
	proto.RegisterType((*ReadCommunicationTypeRequest)(nil), "engine.ReadCommunicationTypeRequest")
	proto.RegisterType((*CommunicationTypeRequest)(nil), "engine.CommunicationTypeRequest")
	proto.RegisterType((*CommunicationType)(nil), "engine.CommunicationType")
}

func init() { proto.RegisterFile("queue_communications.proto", fileDescriptor_090a32393dcacaf2) }

var fileDescriptor_090a32393dcacaf2 = []byte{
	// 526 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0xcf, 0x8a, 0x13, 0x41,
	0x10, 0xc6, 0xe9, 0xf9, 0x13, 0x76, 0x6b, 0x45, 0xb0, 0x61, 0xc9, 0x38, 0xae, 0x31, 0x0e, 0x61,
	0x5d, 0x83, 0x64, 0x60, 0xbd, 0x79, 0x5d, 0x2f, 0xa2, 0x5e, 0x66, 0xf5, 0x26, 0x84, 0x76, 0xba,
	0x88, 0x0d, 0x99, 0xee, 0xc9, 0x4c, 0x47, 0x5c, 0xc5, 0xcb, 0x1e, 0x04, 0x05, 0x4f, 0xe2, 0xd5,
	0x17, 0xf0, 0x69, 0xc4, 0x57, 0xf0, 0x41, 0x64, 0x7a, 0x8c, 0x9b, 0x64, 0xfe, 0x24, 0x0b, 0x7b,
	0x2b, 0xbe, 0x34, 0xdd, 0xbf, 0xfa, 0xbe, 0x9a, 0x0a, 0xf8, 0xb3, 0x39, 0xce, 0x71, 0x1c, 0xab,
	0x24, 0x99, 0x4b, 0x11, 0x33, 0x2d, 0x94, 0xcc, 0x47, 0x69, 0xa6, 0xb4, 0xa2, 0x1d, 0x94, 0x13,
	0x21, 0xd1, 0x3f, 0x98, 0x28, 0x35, 0x99, 0x62, 0xc8, 0x52, 0x11, 0x32, 0x29, 0x95, 0x5e, 0x3e,
	0x15, 0x3c, 0x87, 0xde, 0x63, 0x9c, 0xa2, 0xc6, 0x93, 0xe5, 0x3b, 0x5e, 0x9c, 0xa5, 0x18, 0xe1,
	0x6c, 0x8e, 0xb9, 0xa6, 0xd7, 0xc1, 0x12, 0xdc, 0x23, 0x7d, 0x72, 0x64, 0x47, 0x96, 0xe0, 0xf4,
	0x16, 0xec, 0x72, 0x95, 0x30, 0x21, 0xc7, 0x82, 0x7b, 0x96, 0x91, 0x77, 0x4a, 0xe1, 0x09, 0x0f,
	0x7e, 0x12, 0xe8, 0xbd, 0x4c, 0x39, 0xbb, 0xc4, 0x7d, 0x14, 0x1c, 0xc9, 0x12, 0x34, 0x57, 0xed,
	0x46, 0xa6, 0x2e, 0xb4, 0x58, 0x71, 0xf4, 0xec, 0x52, 0x2b, 0xea, 0x42, 0xd3, 0x67, 0x29, 0x7a,
	0x4e, 0xa9, 0x15, 0x35, 0xed, 0xc3, 0x1e, 0xc7, 0x3c, 0xce, 0x44, 0x5a, 0xbc, 0xe2, 0xb9, 0xe6,
	0xa7, 0x65, 0x69, 0x95, 0xb6, 0xb3, 0x46, 0xfb, 0x0a, 0xf6, 0x9f, 0x89, 0x5c, 0x57, 0x50, 0x0d,
	0x13, 0xbe, 0xd3, 0x86, 0x72, 0x27, 0x32, 0x35, 0x0d, 0xc1, 0x15, 0x1a, 0x93, 0xdc, 0xb3, 0xfa,
	0xf6, 0xd1, 0xde, 0xf1, 0xcd, 0x51, 0xe9, 0xef, 0xa8, 0xda, 0x68, 0x79, 0x2e, 0xc8, 0xa1, 0x77,
	0x8a, 0x2c, 0x8b, 0xdf, 0x34, 0x5a, 0x41, 0xc1, 0x49, 0xd9, 0x04, 0xcd, 0x33, 0x6e, 0x64, 0xea,
	0x42, 0xcb, 0xc5, 0xfb, 0xd2, 0x0e, 0x37, 0x32, 0x35, 0xbd, 0x06, 0x64, 0xf6, 0xcf, 0x0b, 0x32,
	0x5b, 0x6d, 0xc9, 0x59, 0x6b, 0xe9, 0x29, 0x1c, 0x44, 0xc8, 0xf8, 0xd5, 0xa4, 0xf9, 0x9d, 0x80,
	0xd7, 0x06, 0x6f, 0x72, 0x23, 0x35, 0xb9, 0x59, 0x35, 0xb9, 0xd9, 0xcd, 0xb9, 0x39, 0x1b, 0x72,
	0x73, 0xd7, 0xb8, 0x7e, 0x10, 0xb8, 0x51, 0x0d, 0xed, 0x32, 0xad, 0xfd, 0xa7, 0xb7, 0x6b, 0xe8,
	0x9d, 0x1a, 0x7a, 0xb7, 0x99, 0xbe, 0x53, 0xa1, 0x3f, 0xfe, 0xe5, 0xd6, 0x18, 0x77, 0x8a, 0xd9,
	0x5b, 0x11, 0x23, 0x3d, 0x27, 0xd0, 0x3d, 0xc9, 0xb0, 0xee, 0x1b, 0xa1, 0xfd, 0xe6, 0xa9, 0x2a,
	0x6d, 0xf7, 0x9b, 0xe7, 0x2e, 0x18, 0x9e, 0xff, 0xfe, 0xf3, 0xcd, 0x1a, 0x04, 0x77, 0xc2, 0x98,
	0x4d, 0xa7, 0xe3, 0x18, 0xa5, 0xc6, 0x2c, 0x5c, 0x59, 0x0e, 0xe3, 0xa2, 0x81, 0x47, 0x64, 0x48,
	0x3f, 0x13, 0xe8, 0x36, 0x4c, 0x27, 0x3d, 0x5c, 0x3c, 0xd1, 0x3e, 0xbe, 0xfe, 0xed, 0xc5, 0xb9,
	0xda, 0x8f, 0x28, 0xb8, 0x67, 0x70, 0xee, 0xd2, 0x4d, 0x38, 0xf4, 0x13, 0x81, 0xfd, 0xda, 0xa1,
	0xa5, 0x83, 0xc5, 0x0b, 0x6d, 0x33, 0xdd, 0x66, 0xc9, 0x03, 0xc3, 0x70, 0x48, 0x07, 0x1b, 0x18,
	0xc2, 0x0f, 0x82, 0x7f, 0xa4, 0x5f, 0x09, 0x74, 0x1b, 0xb6, 0xd7, 0x85, 0x29, 0xed, 0xeb, 0xad,
	0x0d, 0x26, 0x34, 0x30, 0xf7, 0xfd, 0xad, 0x60, 0x8a, 0x90, 0xbe, 0x10, 0xe8, 0x36, 0x6c, 0xe7,
	0x0b, 0x9e, 0xf6, 0xf5, 0xbd, 0x85, 0x39, 0xc3, 0xad, 0x78, 0x5e, 0x77, 0xcc, 0x1f, 0xc6, 0xc3,
	0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xad, 0x31, 0xb6, 0xf0, 0x74, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CommunicationTypeServiceClient is the client API for CommunicationTypeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommunicationTypeServiceClient interface {
	// Create CommunicationType
	CreateCommunicationType(ctx context.Context, in *CommunicationTypeRequest, opts ...grpc.CallOption) (*CommunicationType, error)
	// List of CommunicationType
	SearchCommunicationType(ctx context.Context, in *SearchCommunicationTypeRequest, opts ...grpc.CallOption) (*ListCommunicationType, error)
	// CommunicationType item
	ReadCommunicationType(ctx context.Context, in *ReadCommunicationTypeRequest, opts ...grpc.CallOption) (*CommunicationType, error)
	// Update CommunicationType
	UpdateCommunicationType(ctx context.Context, in *UpdateCommunicationTypeRequest, opts ...grpc.CallOption) (*CommunicationType, error)
	// Remove CommunicationType
	DeleteCommunicationType(ctx context.Context, in *DeleteCommunicationTypeRequest, opts ...grpc.CallOption) (*CommunicationType, error)
}

type communicationTypeServiceClient struct {
	cc *grpc.ClientConn
}

func NewCommunicationTypeServiceClient(cc *grpc.ClientConn) CommunicationTypeServiceClient {
	return &communicationTypeServiceClient{cc}
}

func (c *communicationTypeServiceClient) CreateCommunicationType(ctx context.Context, in *CommunicationTypeRequest, opts ...grpc.CallOption) (*CommunicationType, error) {
	out := new(CommunicationType)
	err := c.cc.Invoke(ctx, "/engine.CommunicationTypeService/CreateCommunicationType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicationTypeServiceClient) SearchCommunicationType(ctx context.Context, in *SearchCommunicationTypeRequest, opts ...grpc.CallOption) (*ListCommunicationType, error) {
	out := new(ListCommunicationType)
	err := c.cc.Invoke(ctx, "/engine.CommunicationTypeService/SearchCommunicationType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicationTypeServiceClient) ReadCommunicationType(ctx context.Context, in *ReadCommunicationTypeRequest, opts ...grpc.CallOption) (*CommunicationType, error) {
	out := new(CommunicationType)
	err := c.cc.Invoke(ctx, "/engine.CommunicationTypeService/ReadCommunicationType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicationTypeServiceClient) UpdateCommunicationType(ctx context.Context, in *UpdateCommunicationTypeRequest, opts ...grpc.CallOption) (*CommunicationType, error) {
	out := new(CommunicationType)
	err := c.cc.Invoke(ctx, "/engine.CommunicationTypeService/UpdateCommunicationType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicationTypeServiceClient) DeleteCommunicationType(ctx context.Context, in *DeleteCommunicationTypeRequest, opts ...grpc.CallOption) (*CommunicationType, error) {
	out := new(CommunicationType)
	err := c.cc.Invoke(ctx, "/engine.CommunicationTypeService/DeleteCommunicationType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommunicationTypeServiceServer is the server API for CommunicationTypeService service.
type CommunicationTypeServiceServer interface {
	// Create CommunicationType
	CreateCommunicationType(context.Context, *CommunicationTypeRequest) (*CommunicationType, error)
	// List of CommunicationType
	SearchCommunicationType(context.Context, *SearchCommunicationTypeRequest) (*ListCommunicationType, error)
	// CommunicationType item
	ReadCommunicationType(context.Context, *ReadCommunicationTypeRequest) (*CommunicationType, error)
	// Update CommunicationType
	UpdateCommunicationType(context.Context, *UpdateCommunicationTypeRequest) (*CommunicationType, error)
	// Remove CommunicationType
	DeleteCommunicationType(context.Context, *DeleteCommunicationTypeRequest) (*CommunicationType, error)
}

// UnimplementedCommunicationTypeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCommunicationTypeServiceServer struct {
}

func (*UnimplementedCommunicationTypeServiceServer) CreateCommunicationType(ctx context.Context, req *CommunicationTypeRequest) (*CommunicationType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCommunicationType not implemented")
}
func (*UnimplementedCommunicationTypeServiceServer) SearchCommunicationType(ctx context.Context, req *SearchCommunicationTypeRequest) (*ListCommunicationType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchCommunicationType not implemented")
}
func (*UnimplementedCommunicationTypeServiceServer) ReadCommunicationType(ctx context.Context, req *ReadCommunicationTypeRequest) (*CommunicationType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadCommunicationType not implemented")
}
func (*UnimplementedCommunicationTypeServiceServer) UpdateCommunicationType(ctx context.Context, req *UpdateCommunicationTypeRequest) (*CommunicationType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCommunicationType not implemented")
}
func (*UnimplementedCommunicationTypeServiceServer) DeleteCommunicationType(ctx context.Context, req *DeleteCommunicationTypeRequest) (*CommunicationType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCommunicationType not implemented")
}

func RegisterCommunicationTypeServiceServer(s *grpc.Server, srv CommunicationTypeServiceServer) {
	s.RegisterService(&_CommunicationTypeService_serviceDesc, srv)
}

func _CommunicationTypeService_CreateCommunicationType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommunicationTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationTypeServiceServer).CreateCommunicationType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.CommunicationTypeService/CreateCommunicationType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationTypeServiceServer).CreateCommunicationType(ctx, req.(*CommunicationTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunicationTypeService_SearchCommunicationType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchCommunicationTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationTypeServiceServer).SearchCommunicationType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.CommunicationTypeService/SearchCommunicationType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationTypeServiceServer).SearchCommunicationType(ctx, req.(*SearchCommunicationTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunicationTypeService_ReadCommunicationType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadCommunicationTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationTypeServiceServer).ReadCommunicationType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.CommunicationTypeService/ReadCommunicationType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationTypeServiceServer).ReadCommunicationType(ctx, req.(*ReadCommunicationTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunicationTypeService_UpdateCommunicationType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCommunicationTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationTypeServiceServer).UpdateCommunicationType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.CommunicationTypeService/UpdateCommunicationType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationTypeServiceServer).UpdateCommunicationType(ctx, req.(*UpdateCommunicationTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommunicationTypeService_DeleteCommunicationType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCommunicationTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationTypeServiceServer).DeleteCommunicationType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.CommunicationTypeService/DeleteCommunicationType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationTypeServiceServer).DeleteCommunicationType(ctx, req.(*DeleteCommunicationTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CommunicationTypeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "engine.CommunicationTypeService",
	HandlerType: (*CommunicationTypeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCommunicationType",
			Handler:    _CommunicationTypeService_CreateCommunicationType_Handler,
		},
		{
			MethodName: "SearchCommunicationType",
			Handler:    _CommunicationTypeService_SearchCommunicationType_Handler,
		},
		{
			MethodName: "ReadCommunicationType",
			Handler:    _CommunicationTypeService_ReadCommunicationType_Handler,
		},
		{
			MethodName: "UpdateCommunicationType",
			Handler:    _CommunicationTypeService_UpdateCommunicationType_Handler,
		},
		{
			MethodName: "DeleteCommunicationType",
			Handler:    _CommunicationTypeService_DeleteCommunicationType_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "queue_communications.proto",
}
