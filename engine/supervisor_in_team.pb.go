// Code generated by protoc-gen-go. DO NOT EDIT.
// source: supervisor_in_team.proto

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

type UpdateSupervisorInTeamRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TeamId               int64    `protobuf:"varint,2,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	Agent                *Lookup  `protobuf:"bytes,3,opt,name=agent,proto3" json:"agent,omitempty"`
	DomainId             int64    `protobuf:"varint,4,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateSupervisorInTeamRequest) Reset()         { *m = UpdateSupervisorInTeamRequest{} }
func (m *UpdateSupervisorInTeamRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateSupervisorInTeamRequest) ProtoMessage()    {}
func (*UpdateSupervisorInTeamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_232fda82b39ce17f, []int{0}
}

func (m *UpdateSupervisorInTeamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateSupervisorInTeamRequest.Unmarshal(m, b)
}
func (m *UpdateSupervisorInTeamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateSupervisorInTeamRequest.Marshal(b, m, deterministic)
}
func (m *UpdateSupervisorInTeamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateSupervisorInTeamRequest.Merge(m, src)
}
func (m *UpdateSupervisorInTeamRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateSupervisorInTeamRequest.Size(m)
}
func (m *UpdateSupervisorInTeamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateSupervisorInTeamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateSupervisorInTeamRequest proto.InternalMessageInfo

func (m *UpdateSupervisorInTeamRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateSupervisorInTeamRequest) GetTeamId() int64 {
	if m != nil {
		return m.TeamId
	}
	return 0
}

func (m *UpdateSupervisorInTeamRequest) GetAgent() *Lookup {
	if m != nil {
		return m.Agent
	}
	return nil
}

func (m *UpdateSupervisorInTeamRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

type ReadSupervisorInTeamRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TeamId               int64    `protobuf:"varint,2,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	DomainId             int64    `protobuf:"varint,3,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadSupervisorInTeamRequest) Reset()         { *m = ReadSupervisorInTeamRequest{} }
func (m *ReadSupervisorInTeamRequest) String() string { return proto.CompactTextString(m) }
func (*ReadSupervisorInTeamRequest) ProtoMessage()    {}
func (*ReadSupervisorInTeamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_232fda82b39ce17f, []int{1}
}

func (m *ReadSupervisorInTeamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadSupervisorInTeamRequest.Unmarshal(m, b)
}
func (m *ReadSupervisorInTeamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadSupervisorInTeamRequest.Marshal(b, m, deterministic)
}
func (m *ReadSupervisorInTeamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadSupervisorInTeamRequest.Merge(m, src)
}
func (m *ReadSupervisorInTeamRequest) XXX_Size() int {
	return xxx_messageInfo_ReadSupervisorInTeamRequest.Size(m)
}
func (m *ReadSupervisorInTeamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadSupervisorInTeamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadSupervisorInTeamRequest proto.InternalMessageInfo

func (m *ReadSupervisorInTeamRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ReadSupervisorInTeamRequest) GetTeamId() int64 {
	if m != nil {
		return m.TeamId
	}
	return 0
}

func (m *ReadSupervisorInTeamRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

type SearchSupervisorInTeamRequest struct {
	TeamId               int64    `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	Page                 int32    `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Size                 int32    `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Q                    string   `protobuf:"bytes,4,opt,name=q,proto3" json:"q,omitempty"`
	DomainId             int64    `protobuf:"varint,5,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchSupervisorInTeamRequest) Reset()         { *m = SearchSupervisorInTeamRequest{} }
func (m *SearchSupervisorInTeamRequest) String() string { return proto.CompactTextString(m) }
func (*SearchSupervisorInTeamRequest) ProtoMessage()    {}
func (*SearchSupervisorInTeamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_232fda82b39ce17f, []int{2}
}

func (m *SearchSupervisorInTeamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchSupervisorInTeamRequest.Unmarshal(m, b)
}
func (m *SearchSupervisorInTeamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchSupervisorInTeamRequest.Marshal(b, m, deterministic)
}
func (m *SearchSupervisorInTeamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchSupervisorInTeamRequest.Merge(m, src)
}
func (m *SearchSupervisorInTeamRequest) XXX_Size() int {
	return xxx_messageInfo_SearchSupervisorInTeamRequest.Size(m)
}
func (m *SearchSupervisorInTeamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchSupervisorInTeamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchSupervisorInTeamRequest proto.InternalMessageInfo

func (m *SearchSupervisorInTeamRequest) GetTeamId() int64 {
	if m != nil {
		return m.TeamId
	}
	return 0
}

func (m *SearchSupervisorInTeamRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *SearchSupervisorInTeamRequest) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *SearchSupervisorInTeamRequest) GetQ() string {
	if m != nil {
		return m.Q
	}
	return ""
}

func (m *SearchSupervisorInTeamRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

type CreateSupervisorInTeamRequest struct {
	TeamId               int64    `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	Agent                *Lookup  `protobuf:"bytes,2,opt,name=agent,proto3" json:"agent,omitempty"`
	DomainId             int64    `protobuf:"varint,3,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateSupervisorInTeamRequest) Reset()         { *m = CreateSupervisorInTeamRequest{} }
func (m *CreateSupervisorInTeamRequest) String() string { return proto.CompactTextString(m) }
func (*CreateSupervisorInTeamRequest) ProtoMessage()    {}
func (*CreateSupervisorInTeamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_232fda82b39ce17f, []int{3}
}

func (m *CreateSupervisorInTeamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSupervisorInTeamRequest.Unmarshal(m, b)
}
func (m *CreateSupervisorInTeamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSupervisorInTeamRequest.Marshal(b, m, deterministic)
}
func (m *CreateSupervisorInTeamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSupervisorInTeamRequest.Merge(m, src)
}
func (m *CreateSupervisorInTeamRequest) XXX_Size() int {
	return xxx_messageInfo_CreateSupervisorInTeamRequest.Size(m)
}
func (m *CreateSupervisorInTeamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSupervisorInTeamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSupervisorInTeamRequest proto.InternalMessageInfo

func (m *CreateSupervisorInTeamRequest) GetTeamId() int64 {
	if m != nil {
		return m.TeamId
	}
	return 0
}

func (m *CreateSupervisorInTeamRequest) GetAgent() *Lookup {
	if m != nil {
		return m.Agent
	}
	return nil
}

func (m *CreateSupervisorInTeamRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

type DeleteSupervisorInTeamRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TeamId               int64    `protobuf:"varint,2,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	DomainId             int64    `protobuf:"varint,3,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteSupervisorInTeamRequest) Reset()         { *m = DeleteSupervisorInTeamRequest{} }
func (m *DeleteSupervisorInTeamRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteSupervisorInTeamRequest) ProtoMessage()    {}
func (*DeleteSupervisorInTeamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_232fda82b39ce17f, []int{4}
}

func (m *DeleteSupervisorInTeamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteSupervisorInTeamRequest.Unmarshal(m, b)
}
func (m *DeleteSupervisorInTeamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteSupervisorInTeamRequest.Marshal(b, m, deterministic)
}
func (m *DeleteSupervisorInTeamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteSupervisorInTeamRequest.Merge(m, src)
}
func (m *DeleteSupervisorInTeamRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteSupervisorInTeamRequest.Size(m)
}
func (m *DeleteSupervisorInTeamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteSupervisorInTeamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteSupervisorInTeamRequest proto.InternalMessageInfo

func (m *DeleteSupervisorInTeamRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DeleteSupervisorInTeamRequest) GetTeamId() int64 {
	if m != nil {
		return m.TeamId
	}
	return 0
}

func (m *DeleteSupervisorInTeamRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

type SupervisorInTeam struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TeamId               int64    `protobuf:"varint,2,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	Agent                *Lookup  `protobuf:"bytes,3,opt,name=agent,proto3" json:"agent,omitempty"`
	DomainId             int64    `protobuf:"varint,4,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SupervisorInTeam) Reset()         { *m = SupervisorInTeam{} }
func (m *SupervisorInTeam) String() string { return proto.CompactTextString(m) }
func (*SupervisorInTeam) ProtoMessage()    {}
func (*SupervisorInTeam) Descriptor() ([]byte, []int) {
	return fileDescriptor_232fda82b39ce17f, []int{5}
}

func (m *SupervisorInTeam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SupervisorInTeam.Unmarshal(m, b)
}
func (m *SupervisorInTeam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SupervisorInTeam.Marshal(b, m, deterministic)
}
func (m *SupervisorInTeam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SupervisorInTeam.Merge(m, src)
}
func (m *SupervisorInTeam) XXX_Size() int {
	return xxx_messageInfo_SupervisorInTeam.Size(m)
}
func (m *SupervisorInTeam) XXX_DiscardUnknown() {
	xxx_messageInfo_SupervisorInTeam.DiscardUnknown(m)
}

var xxx_messageInfo_SupervisorInTeam proto.InternalMessageInfo

func (m *SupervisorInTeam) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SupervisorInTeam) GetTeamId() int64 {
	if m != nil {
		return m.TeamId
	}
	return 0
}

func (m *SupervisorInTeam) GetAgent() *Lookup {
	if m != nil {
		return m.Agent
	}
	return nil
}

func (m *SupervisorInTeam) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

type ListSupervisorInTeam struct {
	Next                 bool                `protobuf:"varint,1,opt,name=next,proto3" json:"next,omitempty"`
	Items                []*SupervisorInTeam `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ListSupervisorInTeam) Reset()         { *m = ListSupervisorInTeam{} }
func (m *ListSupervisorInTeam) String() string { return proto.CompactTextString(m) }
func (*ListSupervisorInTeam) ProtoMessage()    {}
func (*ListSupervisorInTeam) Descriptor() ([]byte, []int) {
	return fileDescriptor_232fda82b39ce17f, []int{6}
}

func (m *ListSupervisorInTeam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListSupervisorInTeam.Unmarshal(m, b)
}
func (m *ListSupervisorInTeam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListSupervisorInTeam.Marshal(b, m, deterministic)
}
func (m *ListSupervisorInTeam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListSupervisorInTeam.Merge(m, src)
}
func (m *ListSupervisorInTeam) XXX_Size() int {
	return xxx_messageInfo_ListSupervisorInTeam.Size(m)
}
func (m *ListSupervisorInTeam) XXX_DiscardUnknown() {
	xxx_messageInfo_ListSupervisorInTeam.DiscardUnknown(m)
}

var xxx_messageInfo_ListSupervisorInTeam proto.InternalMessageInfo

func (m *ListSupervisorInTeam) GetNext() bool {
	if m != nil {
		return m.Next
	}
	return false
}

func (m *ListSupervisorInTeam) GetItems() []*SupervisorInTeam {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*UpdateSupervisorInTeamRequest)(nil), "engine.UpdateSupervisorInTeamRequest")
	proto.RegisterType((*ReadSupervisorInTeamRequest)(nil), "engine.ReadSupervisorInTeamRequest")
	proto.RegisterType((*SearchSupervisorInTeamRequest)(nil), "engine.SearchSupervisorInTeamRequest")
	proto.RegisterType((*CreateSupervisorInTeamRequest)(nil), "engine.CreateSupervisorInTeamRequest")
	proto.RegisterType((*DeleteSupervisorInTeamRequest)(nil), "engine.DeleteSupervisorInTeamRequest")
	proto.RegisterType((*SupervisorInTeam)(nil), "engine.SupervisorInTeam")
	proto.RegisterType((*ListSupervisorInTeam)(nil), "engine.ListSupervisorInTeam")
}

func init() { proto.RegisterFile("supervisor_in_team.proto", fileDescriptor_232fda82b39ce17f) }

var fileDescriptor_232fda82b39ce17f = []byte{
	// 542 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x95, 0xd1, 0x6a, 0x13, 0x4d,
	0x14, 0xc7, 0x99, 0x4d, 0xb7, 0x5f, 0x7b, 0xf2, 0x51, 0x64, 0x28, 0xe9, 0x92, 0x36, 0x10, 0x56,
	0x85, 0x10, 0x68, 0x46, 0x52, 0x04, 0xf1, 0x52, 0xbd, 0x09, 0xf4, 0x6a, 0xa3, 0x37, 0xbd, 0x09,
	0x93, 0xdd, 0xc3, 0x76, 0x30, 0x99, 0xd9, 0xec, 0x4c, 0x54, 0x0c, 0xf5, 0xc2, 0x5b, 0x45, 0x85,
	0x3e, 0x9a, 0xaf, 0xe0, 0x23, 0xf8, 0x00, 0xb2, 0xb3, 0x26, 0x31, 0xe9, 0x66, 0x4d, 0x51, 0xbc,
	0x9b, 0x39, 0xb3, 0x9c, 0xff, 0x6f, 0xfe, 0x67, 0xce, 0x59, 0xf0, 0xf4, 0x34, 0xc1, 0xf4, 0x95,
	0xd0, 0x2a, 0x1d, 0x08, 0x39, 0x30, 0xc8, 0xc7, 0x9d, 0x24, 0x55, 0x46, 0xd1, 0x5d, 0x94, 0xb1,
	0x90, 0x58, 0xaf, 0x86, 0x4a, 0x6a, 0x93, 0x07, 0xeb, 0x27, 0xb1, 0x52, 0xf1, 0x08, 0x19, 0x4f,
	0x04, 0xe3, 0x52, 0x2a, 0xc3, 0x8d, 0x50, 0x52, 0xe7, 0xa7, 0xfe, 0x07, 0x02, 0x8d, 0x17, 0x49,
	0xc4, 0x0d, 0xf6, 0x17, 0x59, 0x7b, 0xf2, 0x39, 0xf2, 0x71, 0x80, 0x93, 0x29, 0x6a, 0x43, 0x0f,
	0xc0, 0x11, 0x91, 0x47, 0x9a, 0xa4, 0x55, 0x09, 0x1c, 0x11, 0xd1, 0x23, 0xf8, 0x2f, 0x93, 0x1c,
	0x88, 0xc8, 0x73, 0x6c, 0x70, 0x37, 0xdb, 0xf6, 0x22, 0x7a, 0x0f, 0x5c, 0x1e, 0xa3, 0x34, 0x5e,
	0xa5, 0x49, 0x5a, 0xd5, 0xee, 0x41, 0x27, 0xa7, 0xe9, 0x9c, 0x2b, 0xf5, 0x72, 0x9a, 0x04, 0xf9,
	0x21, 0x3d, 0x86, 0xfd, 0x48, 0x8d, 0xb9, 0x90, 0x59, 0x82, 0x1d, 0x9b, 0x60, 0x2f, 0x0f, 0xf4,
	0x22, 0x3f, 0x84, 0xe3, 0x00, 0x79, 0xf4, 0xc7, 0x28, 0x2b, 0x22, 0x95, 0x35, 0x91, 0xec, 0xca,
	0x7d, 0xe4, 0x69, 0x78, 0xb9, 0x49, 0xe7, 0x97, 0xbc, 0x64, 0x25, 0x2f, 0x85, 0x9d, 0x84, 0xc7,
	0x68, 0xd5, 0xdc, 0xc0, 0xae, 0xb3, 0x98, 0x16, 0x6f, 0xd1, 0xca, 0xb8, 0x81, 0x5d, 0xd3, 0xff,
	0x81, 0x4c, 0xec, 0xe5, 0xf6, 0x03, 0x32, 0x59, 0xa5, 0x71, 0xd7, 0x68, 0x66, 0xd0, 0x78, 0x9a,
	0x62, 0x89, 0xff, 0x1b, 0x61, 0x16, 0x7e, 0x3b, 0x5b, 0xfb, 0xbd, 0x6e, 0x05, 0x42, 0xe3, 0x19,
	0x8e, 0xf0, 0x2f, 0x14, 0xbf, 0x54, 0xe6, 0x1d, 0xdc, 0x59, 0x17, 0xf8, 0xa7, 0xcf, 0xea, 0x02,
	0x0e, 0xcf, 0x85, 0x36, 0x37, 0x18, 0x28, 0xec, 0x48, 0x7c, 0x63, 0x2c, 0xc5, 0x5e, 0x60, 0xd7,
	0xb4, 0x03, 0xae, 0x30, 0x38, 0xd6, 0x9e, 0xd3, 0xac, 0xb4, 0xaa, 0x5d, 0x6f, 0x2e, 0x77, 0xc3,
	0xa1, 0xfc, 0xb3, 0xee, 0x77, 0x17, 0x8e, 0xd6, 0xcf, 0xfa, 0xd9, 0x2e, 0x44, 0xfa, 0x89, 0x40,
	0xad, 0xb8, 0xb8, 0xf4, 0xfe, 0x3c, 0x6f, 0x69, 0xf1, 0xeb, 0x1b, 0xe5, 0xfd, 0xb3, 0xf7, 0x5f,
	0xbf, 0x5d, 0x3b, 0xa7, 0x7e, 0x8b, 0x85, 0x7c, 0x34, 0x1a, 0x84, 0x28, 0x0d, 0xa6, 0x2c, 0xf3,
	0x4c, 0xb3, 0xd9, 0x4f, 0x27, 0xaf, 0xd8, 0x72, 0x54, 0xe8, 0xc7, 0xa4, 0x4d, 0x3f, 0x13, 0xa8,
	0x15, 0x3f, 0xfd, 0x25, 0x50, 0x69, 0x6b, 0xd4, 0x4f, 0x16, 0xf6, 0x17, 0x18, 0xea, 0x3f, 0xb0,
	0x50, 0x6d, 0xba, 0x35, 0x14, 0xfd, 0x48, 0xe0, 0xb0, 0xa8, 0xe5, 0xe9, 0xdd, 0xb9, 0x50, 0xc9,
	0x40, 0x28, 0xb1, 0xe7, 0xa1, 0x25, 0x61, 0xf4, 0x74, 0x5b, 0x12, 0x36, 0x13, 0xd1, 0x15, 0xbd,
	0x26, 0x50, 0x2b, 0x1e, 0x87, 0x4b, 0x83, 0x4a, 0xc7, 0x65, 0x09, 0xd2, 0x23, 0x8b, 0xd4, 0xad,
	0xdf, 0x0e, 0x29, 0x2b, 0xdb, 0x17, 0x02, 0xb5, 0xe2, 0x3e, 0x5d, 0x52, 0x95, 0xf6, 0xf1, 0xef,
	0x8d, 0x6a, 0xdf, 0x8e, 0xea, 0x89, 0x7f, 0xd1, 0x8c, 0x85, 0xb9, 0x9c, 0x0e, 0x3b, 0xa1, 0x1a,
	0xb3, 0xd7, 0x38, 0x14, 0x06, 0x47, 0xcc, 0xfe, 0x53, 0x34, 0xcb, 0xb5, 0x86, 0xbb, 0x76, 0x7b,
	0xf6, 0x23, 0x00, 0x00, 0xff, 0xff, 0xa3, 0xac, 0x07, 0x04, 0xb1, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SupervisorInTeamServiceClient is the client API for SupervisorInTeamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SupervisorInTeamServiceClient interface {
	// Create SupervisorInTeam
	CreateSupervisorInTeam(ctx context.Context, in *CreateSupervisorInTeamRequest, opts ...grpc.CallOption) (*SupervisorInTeam, error)
	// List of SupervisorInTeam
	SearchSupervisorInTeam(ctx context.Context, in *SearchSupervisorInTeamRequest, opts ...grpc.CallOption) (*ListSupervisorInTeam, error)
	// SupervisorInTeam item
	ReadSupervisorInTeam(ctx context.Context, in *ReadSupervisorInTeamRequest, opts ...grpc.CallOption) (*SupervisorInTeam, error)
	// Update SupervisorInTeam
	UpdateSupervisorInTeam(ctx context.Context, in *UpdateSupervisorInTeamRequest, opts ...grpc.CallOption) (*SupervisorInTeam, error)
	// Remove SupervisorInTeam
	DeleteSupervisorInTeam(ctx context.Context, in *DeleteSupervisorInTeamRequest, opts ...grpc.CallOption) (*SupervisorInTeam, error)
}

type supervisorInTeamServiceClient struct {
	cc *grpc.ClientConn
}

func NewSupervisorInTeamServiceClient(cc *grpc.ClientConn) SupervisorInTeamServiceClient {
	return &supervisorInTeamServiceClient{cc}
}

func (c *supervisorInTeamServiceClient) CreateSupervisorInTeam(ctx context.Context, in *CreateSupervisorInTeamRequest, opts ...grpc.CallOption) (*SupervisorInTeam, error) {
	out := new(SupervisorInTeam)
	err := c.cc.Invoke(ctx, "/engine.SupervisorInTeamService/CreateSupervisorInTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supervisorInTeamServiceClient) SearchSupervisorInTeam(ctx context.Context, in *SearchSupervisorInTeamRequest, opts ...grpc.CallOption) (*ListSupervisorInTeam, error) {
	out := new(ListSupervisorInTeam)
	err := c.cc.Invoke(ctx, "/engine.SupervisorInTeamService/SearchSupervisorInTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supervisorInTeamServiceClient) ReadSupervisorInTeam(ctx context.Context, in *ReadSupervisorInTeamRequest, opts ...grpc.CallOption) (*SupervisorInTeam, error) {
	out := new(SupervisorInTeam)
	err := c.cc.Invoke(ctx, "/engine.SupervisorInTeamService/ReadSupervisorInTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supervisorInTeamServiceClient) UpdateSupervisorInTeam(ctx context.Context, in *UpdateSupervisorInTeamRequest, opts ...grpc.CallOption) (*SupervisorInTeam, error) {
	out := new(SupervisorInTeam)
	err := c.cc.Invoke(ctx, "/engine.SupervisorInTeamService/UpdateSupervisorInTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supervisorInTeamServiceClient) DeleteSupervisorInTeam(ctx context.Context, in *DeleteSupervisorInTeamRequest, opts ...grpc.CallOption) (*SupervisorInTeam, error) {
	out := new(SupervisorInTeam)
	err := c.cc.Invoke(ctx, "/engine.SupervisorInTeamService/DeleteSupervisorInTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SupervisorInTeamServiceServer is the server API for SupervisorInTeamService service.
type SupervisorInTeamServiceServer interface {
	// Create SupervisorInTeam
	CreateSupervisorInTeam(context.Context, *CreateSupervisorInTeamRequest) (*SupervisorInTeam, error)
	// List of SupervisorInTeam
	SearchSupervisorInTeam(context.Context, *SearchSupervisorInTeamRequest) (*ListSupervisorInTeam, error)
	// SupervisorInTeam item
	ReadSupervisorInTeam(context.Context, *ReadSupervisorInTeamRequest) (*SupervisorInTeam, error)
	// Update SupervisorInTeam
	UpdateSupervisorInTeam(context.Context, *UpdateSupervisorInTeamRequest) (*SupervisorInTeam, error)
	// Remove SupervisorInTeam
	DeleteSupervisorInTeam(context.Context, *DeleteSupervisorInTeamRequest) (*SupervisorInTeam, error)
}

// UnimplementedSupervisorInTeamServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSupervisorInTeamServiceServer struct {
}

func (*UnimplementedSupervisorInTeamServiceServer) CreateSupervisorInTeam(ctx context.Context, req *CreateSupervisorInTeamRequest) (*SupervisorInTeam, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSupervisorInTeam not implemented")
}
func (*UnimplementedSupervisorInTeamServiceServer) SearchSupervisorInTeam(ctx context.Context, req *SearchSupervisorInTeamRequest) (*ListSupervisorInTeam, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchSupervisorInTeam not implemented")
}
func (*UnimplementedSupervisorInTeamServiceServer) ReadSupervisorInTeam(ctx context.Context, req *ReadSupervisorInTeamRequest) (*SupervisorInTeam, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadSupervisorInTeam not implemented")
}
func (*UnimplementedSupervisorInTeamServiceServer) UpdateSupervisorInTeam(ctx context.Context, req *UpdateSupervisorInTeamRequest) (*SupervisorInTeam, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSupervisorInTeam not implemented")
}
func (*UnimplementedSupervisorInTeamServiceServer) DeleteSupervisorInTeam(ctx context.Context, req *DeleteSupervisorInTeamRequest) (*SupervisorInTeam, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSupervisorInTeam not implemented")
}

func RegisterSupervisorInTeamServiceServer(s *grpc.Server, srv SupervisorInTeamServiceServer) {
	s.RegisterService(&_SupervisorInTeamService_serviceDesc, srv)
}

func _SupervisorInTeamService_CreateSupervisorInTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSupervisorInTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupervisorInTeamServiceServer).CreateSupervisorInTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.SupervisorInTeamService/CreateSupervisorInTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupervisorInTeamServiceServer).CreateSupervisorInTeam(ctx, req.(*CreateSupervisorInTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupervisorInTeamService_SearchSupervisorInTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchSupervisorInTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupervisorInTeamServiceServer).SearchSupervisorInTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.SupervisorInTeamService/SearchSupervisorInTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupervisorInTeamServiceServer).SearchSupervisorInTeam(ctx, req.(*SearchSupervisorInTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupervisorInTeamService_ReadSupervisorInTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadSupervisorInTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupervisorInTeamServiceServer).ReadSupervisorInTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.SupervisorInTeamService/ReadSupervisorInTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupervisorInTeamServiceServer).ReadSupervisorInTeam(ctx, req.(*ReadSupervisorInTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupervisorInTeamService_UpdateSupervisorInTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSupervisorInTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupervisorInTeamServiceServer).UpdateSupervisorInTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.SupervisorInTeamService/UpdateSupervisorInTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupervisorInTeamServiceServer).UpdateSupervisorInTeam(ctx, req.(*UpdateSupervisorInTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupervisorInTeamService_DeleteSupervisorInTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSupervisorInTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupervisorInTeamServiceServer).DeleteSupervisorInTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.SupervisorInTeamService/DeleteSupervisorInTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupervisorInTeamServiceServer).DeleteSupervisorInTeam(ctx, req.(*DeleteSupervisorInTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SupervisorInTeamService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "engine.SupervisorInTeamService",
	HandlerType: (*SupervisorInTeamServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSupervisorInTeam",
			Handler:    _SupervisorInTeamService_CreateSupervisorInTeam_Handler,
		},
		{
			MethodName: "SearchSupervisorInTeam",
			Handler:    _SupervisorInTeamService_SearchSupervisorInTeam_Handler,
		},
		{
			MethodName: "ReadSupervisorInTeam",
			Handler:    _SupervisorInTeamService_ReadSupervisorInTeam_Handler,
		},
		{
			MethodName: "UpdateSupervisorInTeam",
			Handler:    _SupervisorInTeamService_UpdateSupervisorInTeam_Handler,
		},
		{
			MethodName: "DeleteSupervisorInTeam",
			Handler:    _SupervisorInTeamService_DeleteSupervisorInTeam_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "supervisor_in_team.proto",
}