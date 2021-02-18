// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pause_cause.proto

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

type CreateAgentPauseCauseRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	LimitMin             uint32   `protobuf:"varint,2,opt,name=limit_min,json=limitMin,proto3" json:"limit_min,omitempty"`
	AllowSupervisor      bool     `protobuf:"varint,3,opt,name=allow_supervisor,json=allowSupervisor,proto3" json:"allow_supervisor,omitempty"`
	AllowAgent           bool     `protobuf:"varint,4,opt,name=allow_agent,json=allowAgent,proto3" json:"allow_agent,omitempty"`
	AllowAdmin           bool     `protobuf:"varint,5,opt,name=allow_admin,json=allowAdmin,proto3" json:"allow_admin,omitempty"`
	Description          string   `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateAgentPauseCauseRequest) Reset()         { *m = CreateAgentPauseCauseRequest{} }
func (m *CreateAgentPauseCauseRequest) String() string { return proto.CompactTextString(m) }
func (*CreateAgentPauseCauseRequest) ProtoMessage()    {}
func (*CreateAgentPauseCauseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9e5c0251eb1bbe3, []int{0}
}

func (m *CreateAgentPauseCauseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAgentPauseCauseRequest.Unmarshal(m, b)
}
func (m *CreateAgentPauseCauseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAgentPauseCauseRequest.Marshal(b, m, deterministic)
}
func (m *CreateAgentPauseCauseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAgentPauseCauseRequest.Merge(m, src)
}
func (m *CreateAgentPauseCauseRequest) XXX_Size() int {
	return xxx_messageInfo_CreateAgentPauseCauseRequest.Size(m)
}
func (m *CreateAgentPauseCauseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAgentPauseCauseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAgentPauseCauseRequest proto.InternalMessageInfo

func (m *CreateAgentPauseCauseRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateAgentPauseCauseRequest) GetLimitMin() uint32 {
	if m != nil {
		return m.LimitMin
	}
	return 0
}

func (m *CreateAgentPauseCauseRequest) GetAllowSupervisor() bool {
	if m != nil {
		return m.AllowSupervisor
	}
	return false
}

func (m *CreateAgentPauseCauseRequest) GetAllowAgent() bool {
	if m != nil {
		return m.AllowAgent
	}
	return false
}

func (m *CreateAgentPauseCauseRequest) GetAllowAdmin() bool {
	if m != nil {
		return m.AllowAdmin
	}
	return false
}

func (m *CreateAgentPauseCauseRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type AgentPauseCause struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            int64    `protobuf:"varint,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	CreatedBy            *Lookup  `protobuf:"bytes,3,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	UpdatedAt            int64    `protobuf:"varint,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	UpdatedBy            *Lookup  `protobuf:"bytes,5,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
	Name                 string   `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	LimitMin             uint32   `protobuf:"varint,7,opt,name=limit_min,json=limitMin,proto3" json:"limit_min,omitempty"`
	AllowSupervisor      bool     `protobuf:"varint,8,opt,name=allow_supervisor,json=allowSupervisor,proto3" json:"allow_supervisor,omitempty"`
	AllowAgent           bool     `protobuf:"varint,9,opt,name=allow_agent,json=allowAgent,proto3" json:"allow_agent,omitempty"`
	AllowAdmin           bool     `protobuf:"varint,10,opt,name=allow_admin,json=allowAdmin,proto3" json:"allow_admin,omitempty"`
	Description          string   `protobuf:"bytes,11,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgentPauseCause) Reset()         { *m = AgentPauseCause{} }
func (m *AgentPauseCause) String() string { return proto.CompactTextString(m) }
func (*AgentPauseCause) ProtoMessage()    {}
func (*AgentPauseCause) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9e5c0251eb1bbe3, []int{1}
}

func (m *AgentPauseCause) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgentPauseCause.Unmarshal(m, b)
}
func (m *AgentPauseCause) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgentPauseCause.Marshal(b, m, deterministic)
}
func (m *AgentPauseCause) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentPauseCause.Merge(m, src)
}
func (m *AgentPauseCause) XXX_Size() int {
	return xxx_messageInfo_AgentPauseCause.Size(m)
}
func (m *AgentPauseCause) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentPauseCause.DiscardUnknown(m)
}

var xxx_messageInfo_AgentPauseCause proto.InternalMessageInfo

func (m *AgentPauseCause) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AgentPauseCause) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *AgentPauseCause) GetCreatedBy() *Lookup {
	if m != nil {
		return m.CreatedBy
	}
	return nil
}

func (m *AgentPauseCause) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *AgentPauseCause) GetUpdatedBy() *Lookup {
	if m != nil {
		return m.UpdatedBy
	}
	return nil
}

func (m *AgentPauseCause) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AgentPauseCause) GetLimitMin() uint32 {
	if m != nil {
		return m.LimitMin
	}
	return 0
}

func (m *AgentPauseCause) GetAllowSupervisor() bool {
	if m != nil {
		return m.AllowSupervisor
	}
	return false
}

func (m *AgentPauseCause) GetAllowAgent() bool {
	if m != nil {
		return m.AllowAgent
	}
	return false
}

func (m *AgentPauseCause) GetAllowAdmin() bool {
	if m != nil {
		return m.AllowAdmin
	}
	return false
}

func (m *AgentPauseCause) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type SearchAgentPauseCauseRequest struct {
	Page                 int32    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size                 int32    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Q                    string   `protobuf:"bytes,3,opt,name=q,proto3" json:"q,omitempty"`
	Fields               []string `protobuf:"bytes,4,rep,name=fields,proto3" json:"fields,omitempty"`
	Id                   []uint32 `protobuf:"varint,5,rep,packed,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchAgentPauseCauseRequest) Reset()         { *m = SearchAgentPauseCauseRequest{} }
func (m *SearchAgentPauseCauseRequest) String() string { return proto.CompactTextString(m) }
func (*SearchAgentPauseCauseRequest) ProtoMessage()    {}
func (*SearchAgentPauseCauseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9e5c0251eb1bbe3, []int{2}
}

func (m *SearchAgentPauseCauseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchAgentPauseCauseRequest.Unmarshal(m, b)
}
func (m *SearchAgentPauseCauseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchAgentPauseCauseRequest.Marshal(b, m, deterministic)
}
func (m *SearchAgentPauseCauseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchAgentPauseCauseRequest.Merge(m, src)
}
func (m *SearchAgentPauseCauseRequest) XXX_Size() int {
	return xxx_messageInfo_SearchAgentPauseCauseRequest.Size(m)
}
func (m *SearchAgentPauseCauseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchAgentPauseCauseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchAgentPauseCauseRequest proto.InternalMessageInfo

func (m *SearchAgentPauseCauseRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *SearchAgentPauseCauseRequest) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *SearchAgentPauseCauseRequest) GetQ() string {
	if m != nil {
		return m.Q
	}
	return ""
}

func (m *SearchAgentPauseCauseRequest) GetFields() []string {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *SearchAgentPauseCauseRequest) GetId() []uint32 {
	if m != nil {
		return m.Id
	}
	return nil
}

type ListAgentPauseCause struct {
	Next                 bool               `protobuf:"varint,1,opt,name=next,proto3" json:"next,omitempty"`
	Items                []*AgentPauseCause `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ListAgentPauseCause) Reset()         { *m = ListAgentPauseCause{} }
func (m *ListAgentPauseCause) String() string { return proto.CompactTextString(m) }
func (*ListAgentPauseCause) ProtoMessage()    {}
func (*ListAgentPauseCause) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9e5c0251eb1bbe3, []int{3}
}

func (m *ListAgentPauseCause) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAgentPauseCause.Unmarshal(m, b)
}
func (m *ListAgentPauseCause) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAgentPauseCause.Marshal(b, m, deterministic)
}
func (m *ListAgentPauseCause) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAgentPauseCause.Merge(m, src)
}
func (m *ListAgentPauseCause) XXX_Size() int {
	return xxx_messageInfo_ListAgentPauseCause.Size(m)
}
func (m *ListAgentPauseCause) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAgentPauseCause.DiscardUnknown(m)
}

var xxx_messageInfo_ListAgentPauseCause proto.InternalMessageInfo

func (m *ListAgentPauseCause) GetNext() bool {
	if m != nil {
		return m.Next
	}
	return false
}

func (m *ListAgentPauseCause) GetItems() []*AgentPauseCause {
	if m != nil {
		return m.Items
	}
	return nil
}

type ReadAgentPauseCauseRequest struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadAgentPauseCauseRequest) Reset()         { *m = ReadAgentPauseCauseRequest{} }
func (m *ReadAgentPauseCauseRequest) String() string { return proto.CompactTextString(m) }
func (*ReadAgentPauseCauseRequest) ProtoMessage()    {}
func (*ReadAgentPauseCauseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9e5c0251eb1bbe3, []int{4}
}

func (m *ReadAgentPauseCauseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadAgentPauseCauseRequest.Unmarshal(m, b)
}
func (m *ReadAgentPauseCauseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadAgentPauseCauseRequest.Marshal(b, m, deterministic)
}
func (m *ReadAgentPauseCauseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadAgentPauseCauseRequest.Merge(m, src)
}
func (m *ReadAgentPauseCauseRequest) XXX_Size() int {
	return xxx_messageInfo_ReadAgentPauseCauseRequest.Size(m)
}
func (m *ReadAgentPauseCauseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadAgentPauseCauseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadAgentPauseCauseRequest proto.InternalMessageInfo

func (m *ReadAgentPauseCauseRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type PatchAgentPauseCauseRequest struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Fields               []string `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	LimitMin             uint32   `protobuf:"varint,4,opt,name=limit_min,json=limitMin,proto3" json:"limit_min,omitempty"`
	AllowSupervisor      bool     `protobuf:"varint,5,opt,name=allow_supervisor,json=allowSupervisor,proto3" json:"allow_supervisor,omitempty"`
	AllowAgent           bool     `protobuf:"varint,6,opt,name=allow_agent,json=allowAgent,proto3" json:"allow_agent,omitempty"`
	AllowAdmin           bool     `protobuf:"varint,7,opt,name=allow_admin,json=allowAdmin,proto3" json:"allow_admin,omitempty"`
	Description          string   `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PatchAgentPauseCauseRequest) Reset()         { *m = PatchAgentPauseCauseRequest{} }
func (m *PatchAgentPauseCauseRequest) String() string { return proto.CompactTextString(m) }
func (*PatchAgentPauseCauseRequest) ProtoMessage()    {}
func (*PatchAgentPauseCauseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9e5c0251eb1bbe3, []int{5}
}

func (m *PatchAgentPauseCauseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PatchAgentPauseCauseRequest.Unmarshal(m, b)
}
func (m *PatchAgentPauseCauseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PatchAgentPauseCauseRequest.Marshal(b, m, deterministic)
}
func (m *PatchAgentPauseCauseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PatchAgentPauseCauseRequest.Merge(m, src)
}
func (m *PatchAgentPauseCauseRequest) XXX_Size() int {
	return xxx_messageInfo_PatchAgentPauseCauseRequest.Size(m)
}
func (m *PatchAgentPauseCauseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PatchAgentPauseCauseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PatchAgentPauseCauseRequest proto.InternalMessageInfo

func (m *PatchAgentPauseCauseRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PatchAgentPauseCauseRequest) GetFields() []string {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *PatchAgentPauseCauseRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PatchAgentPauseCauseRequest) GetLimitMin() uint32 {
	if m != nil {
		return m.LimitMin
	}
	return 0
}

func (m *PatchAgentPauseCauseRequest) GetAllowSupervisor() bool {
	if m != nil {
		return m.AllowSupervisor
	}
	return false
}

func (m *PatchAgentPauseCauseRequest) GetAllowAgent() bool {
	if m != nil {
		return m.AllowAgent
	}
	return false
}

func (m *PatchAgentPauseCauseRequest) GetAllowAdmin() bool {
	if m != nil {
		return m.AllowAdmin
	}
	return false
}

func (m *PatchAgentPauseCauseRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type UpdateAgentPauseCauseRequest struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	LimitMin             uint32   `protobuf:"varint,3,opt,name=limit_min,json=limitMin,proto3" json:"limit_min,omitempty"`
	AllowSupervisor      bool     `protobuf:"varint,4,opt,name=allow_supervisor,json=allowSupervisor,proto3" json:"allow_supervisor,omitempty"`
	AllowAgent           bool     `protobuf:"varint,5,opt,name=allow_agent,json=allowAgent,proto3" json:"allow_agent,omitempty"`
	AllowAdmin           bool     `protobuf:"varint,6,opt,name=allow_admin,json=allowAdmin,proto3" json:"allow_admin,omitempty"`
	Description          string   `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateAgentPauseCauseRequest) Reset()         { *m = UpdateAgentPauseCauseRequest{} }
func (m *UpdateAgentPauseCauseRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateAgentPauseCauseRequest) ProtoMessage()    {}
func (*UpdateAgentPauseCauseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9e5c0251eb1bbe3, []int{6}
}

func (m *UpdateAgentPauseCauseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateAgentPauseCauseRequest.Unmarshal(m, b)
}
func (m *UpdateAgentPauseCauseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateAgentPauseCauseRequest.Marshal(b, m, deterministic)
}
func (m *UpdateAgentPauseCauseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateAgentPauseCauseRequest.Merge(m, src)
}
func (m *UpdateAgentPauseCauseRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateAgentPauseCauseRequest.Size(m)
}
func (m *UpdateAgentPauseCauseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateAgentPauseCauseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateAgentPauseCauseRequest proto.InternalMessageInfo

func (m *UpdateAgentPauseCauseRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateAgentPauseCauseRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateAgentPauseCauseRequest) GetLimitMin() uint32 {
	if m != nil {
		return m.LimitMin
	}
	return 0
}

func (m *UpdateAgentPauseCauseRequest) GetAllowSupervisor() bool {
	if m != nil {
		return m.AllowSupervisor
	}
	return false
}

func (m *UpdateAgentPauseCauseRequest) GetAllowAgent() bool {
	if m != nil {
		return m.AllowAgent
	}
	return false
}

func (m *UpdateAgentPauseCauseRequest) GetAllowAdmin() bool {
	if m != nil {
		return m.AllowAdmin
	}
	return false
}

func (m *UpdateAgentPauseCauseRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type DeleteAgentPauseCauseRequest struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteAgentPauseCauseRequest) Reset()         { *m = DeleteAgentPauseCauseRequest{} }
func (m *DeleteAgentPauseCauseRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteAgentPauseCauseRequest) ProtoMessage()    {}
func (*DeleteAgentPauseCauseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9e5c0251eb1bbe3, []int{7}
}

func (m *DeleteAgentPauseCauseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteAgentPauseCauseRequest.Unmarshal(m, b)
}
func (m *DeleteAgentPauseCauseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteAgentPauseCauseRequest.Marshal(b, m, deterministic)
}
func (m *DeleteAgentPauseCauseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteAgentPauseCauseRequest.Merge(m, src)
}
func (m *DeleteAgentPauseCauseRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteAgentPauseCauseRequest.Size(m)
}
func (m *DeleteAgentPauseCauseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteAgentPauseCauseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteAgentPauseCauseRequest proto.InternalMessageInfo

func (m *DeleteAgentPauseCauseRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateAgentPauseCauseRequest)(nil), "engine.CreateAgentPauseCauseRequest")
	proto.RegisterType((*AgentPauseCause)(nil), "engine.AgentPauseCause")
	proto.RegisterType((*SearchAgentPauseCauseRequest)(nil), "engine.SearchAgentPauseCauseRequest")
	proto.RegisterType((*ListAgentPauseCause)(nil), "engine.ListAgentPauseCause")
	proto.RegisterType((*ReadAgentPauseCauseRequest)(nil), "engine.ReadAgentPauseCauseRequest")
	proto.RegisterType((*PatchAgentPauseCauseRequest)(nil), "engine.PatchAgentPauseCauseRequest")
	proto.RegisterType((*UpdateAgentPauseCauseRequest)(nil), "engine.UpdateAgentPauseCauseRequest")
	proto.RegisterType((*DeleteAgentPauseCauseRequest)(nil), "engine.DeleteAgentPauseCauseRequest")
}

func init() { proto.RegisterFile("pause_cause.proto", fileDescriptor_e9e5c0251eb1bbe3) }

var fileDescriptor_e9e5c0251eb1bbe3 = []byte{
	// 728 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0x41, 0x6f, 0xd3, 0x4c,
	0x10, 0x95, 0x9d, 0xd8, 0x4d, 0x26, 0x5f, 0xdb, 0x8f, 0x2d, 0x29, 0x26, 0x09, 0xc5, 0xb8, 0x15,
	0x4a, 0x2b, 0x1a, 0x4b, 0xe1, 0xc6, 0xad, 0x2d, 0xc7, 0x22, 0x55, 0xae, 0x90, 0x10, 0x97, 0x68,
	0x63, 0x2f, 0xe9, 0x0a, 0xc7, 0x76, 0xbd, 0x1b, 0x4a, 0x4a, 0x0b, 0x02, 0x89, 0x5f, 0x80, 0xc4,
	0x1f, 0xe3, 0xce, 0x89, 0x33, 0xff, 0x00, 0x09, 0x79, 0x1d, 0x27, 0x21, 0xb1, 0x63, 0x5f, 0xac,
	0xf5, 0x78, 0x34, 0xf3, 0xde, 0x9b, 0x37, 0x9b, 0xc0, 0x9d, 0x00, 0x8f, 0x18, 0xe9, 0xd9, 0xd1,
	0xb3, 0x13, 0x84, 0x3e, 0xf7, 0x91, 0x4a, 0xbc, 0x01, 0xf5, 0x48, 0xa3, 0x35, 0xf0, 0xfd, 0x81,
	0x4b, 0x4c, 0x1c, 0x50, 0x13, 0x7b, 0x9e, 0xcf, 0x31, 0xa7, 0xbe, 0xc7, 0xe2, 0xac, 0x46, 0xcd,
	0xf6, 0x3d, 0xc6, 0xe3, 0x17, 0xe3, 0xa7, 0x04, 0xad, 0x93, 0x90, 0x60, 0x4e, 0x8e, 0x06, 0xc4,
	0xe3, 0x67, 0x51, 0xb5, 0x93, 0xe8, 0x61, 0x91, 0xcb, 0x11, 0x61, 0x1c, 0x21, 0x28, 0x7b, 0x78,
	0x48, 0x34, 0x49, 0x97, 0xda, 0x55, 0x4b, 0x9c, 0x51, 0x13, 0xaa, 0x2e, 0x1d, 0x52, 0xde, 0x1b,
	0x52, 0x4f, 0x93, 0x75, 0xa9, 0xbd, 0x6e, 0x55, 0x44, 0xe0, 0x05, 0xf5, 0xd0, 0x3e, 0xfc, 0x8f,
	0x5d, 0xd7, 0xbf, 0xea, 0xb1, 0x51, 0x40, 0xc2, 0x77, 0x94, 0xf9, 0xa1, 0x56, 0xd2, 0xa5, 0x76,
	0xc5, 0xda, 0x14, 0xf1, 0xf3, 0x69, 0x18, 0x3d, 0x84, 0x5a, 0x9c, 0x8a, 0xa3, 0xde, 0x5a, 0x59,
	0x64, 0x81, 0x08, 0x09, 0x34, 0x73, 0x09, 0x4e, 0xd4, 0x4a, 0x99, 0x4f, 0x88, 0x22, 0x48, 0x87,
	0x9a, 0x43, 0x98, 0x1d, 0xd2, 0x20, 0x62, 0xa8, 0xa9, 0x02, 0xe4, 0x7c, 0xc8, 0xf8, 0x23, 0xc3,
	0xe6, 0x02, 0x35, 0xb4, 0x01, 0x32, 0x75, 0x04, 0xa3, 0x75, 0x4b, 0xa6, 0x0e, 0x7a, 0x00, 0x60,
	0x0b, 0x0d, 0x9c, 0x1e, 0xe6, 0x82, 0x50, 0xc9, 0xaa, 0x4e, 0x22, 0x47, 0x1c, 0x1d, 0xce, 0x3e,
	0xf7, 0xc7, 0x82, 0x4b, 0xad, 0xbb, 0xd1, 0x89, 0xb5, 0xee, 0x9c, 0xfa, 0xfe, 0xdb, 0x51, 0x30,
	0x4d, 0x3f, 0x1e, 0x47, 0xd5, 0x46, 0x81, 0x93, 0x54, 0x2b, 0xc7, 0xd5, 0x26, 0x91, 0xb8, 0x5a,
	0xf2, 0xb9, 0x3f, 0x16, 0x94, 0x52, 0xaa, 0x4d, 0x32, 0x8e, 0xc7, 0x53, 0xfd, 0xd5, 0x2c, 0xfd,
	0xd7, 0x0a, 0xe8, 0x5f, 0x29, 0xa4, 0x7f, 0x35, 0x4f, 0x7f, 0xc8, 0xd3, 0xbf, 0xb6, 0xac, 0xff,
	0x0d, 0xb4, 0xce, 0x09, 0x0e, 0xed, 0x8b, 0x6c, 0x7f, 0x05, 0x78, 0x10, 0xfb, 0x4b, 0xb1, 0xc4,
	0x39, 0x8a, 0x31, 0x7a, 0x4d, 0xc4, 0x24, 0x14, 0x4b, 0x9c, 0xd1, 0x7f, 0x20, 0x5d, 0x0a, 0xed,
	0xab, 0x96, 0x74, 0x89, 0xb6, 0x41, 0x7d, 0x43, 0x89, 0xeb, 0x30, 0xad, 0xac, 0x97, 0xda, 0x55,
	0x6b, 0xf2, 0x36, 0x99, 0xac, 0xa2, 0x97, 0xe2, 0xc9, 0x1a, 0xaf, 0x60, 0xeb, 0x94, 0x32, 0xbe,
	0x68, 0x80, 0x48, 0x54, 0xf2, 0x9e, 0x8b, 0xa6, 0x15, 0x4b, 0x9c, 0xd1, 0x21, 0x28, 0x94, 0x93,
	0x21, 0xd3, 0x64, 0xbd, 0xd4, 0xae, 0x75, 0xef, 0x25, 0x23, 0x59, 0xc4, 0x1d, 0x67, 0x19, 0x4f,
	0xa0, 0x61, 0x11, 0xec, 0x64, 0xb0, 0x5a, 0x70, 0x98, 0xf1, 0x55, 0x86, 0xe6, 0x19, 0xe6, 0x99,
	0x2a, 0x2c, 0x3a, 0x72, 0xc6, 0x4f, 0xfe, 0x87, 0x5f, 0xe2, 0x86, 0x52, 0x96, 0x1b, 0xca, 0x05,
	0xdc, 0xa0, 0x14, 0x72, 0x83, 0x9a, 0xe7, 0x86, 0xb5, 0x3c, 0x37, 0x54, 0x96, 0xdd, 0xf0, 0x5b,
	0x82, 0xd6, 0x4b, 0xe1, 0xed, 0x82, 0x42, 0x24, 0x84, 0xe5, 0x2c, 0xc2, 0xa5, 0x02, 0x84, 0xcb,
	0x85, 0x08, 0x2b, 0x79, 0x84, 0xd5, 0x3c, 0xc2, 0x6b, 0xcb, 0x84, 0x3b, 0xd0, 0x7a, 0x4e, 0x5c,
	0x52, 0x94, 0x6f, 0xf7, 0xbb, 0x0a, 0xdb, 0x0b, 0xa9, 0xe7, 0x11, 0x5e, 0x9b, 0xa0, 0x1b, 0xa8,
	0xa7, 0xde, 0xd4, 0x68, 0x2f, 0xb1, 0xea, 0xaa, 0x8b, 0xbc, 0x91, 0x65, 0x68, 0x63, 0xef, 0xcb,
	0x8f, 0x5f, 0xdf, 0xe4, 0x1d, 0xe3, 0xbe, 0x69, 0x63, 0xd7, 0xed, 0xd9, 0xc4, 0xe3, 0x24, 0x34,
	0xe7, 0x7e, 0x5d, 0xd8, 0x33, 0xe9, 0x00, 0xdd, 0x42, 0x3d, 0x75, 0x8f, 0x67, 0xdd, 0x57, 0xad,
	0x79, 0xa3, 0x39, 0xbd, 0xe1, 0x96, 0xd7, 0xd1, 0x78, 0x24, 0x10, 0x34, 0x51, 0x36, 0x02, 0x74,
	0x0d, 0x5b, 0x29, 0xeb, 0x86, 0x8c, 0xa4, 0x6c, 0xf6, 0x2e, 0x66, 0x13, 0x7f, 0x2c, 0xda, 0xea,
	0x68, 0x27, 0xb3, 0xad, 0xf9, 0x81, 0x3a, 0xb7, 0xe8, 0x13, 0xdc, 0x4d, 0xdb, 0x5d, 0xb4, 0x9b,
	0x14, 0x5e, 0xb1, 0xd9, 0xd9, 0xdd, 0xf7, 0x45, 0xf7, 0xdd, 0x6e, 0x4e, 0xf7, 0x48, 0xfb, 0xcf,
	0x12, 0xd4, 0x53, 0xb7, 0x66, 0x26, 0xfe, 0xaa, 0xa5, 0xca, 0xc5, 0xd0, 0x28, 0x80, 0xe1, 0x23,
	0xd4, 0x53, 0x8d, 0x3c, 0x83, 0xb0, 0xca, 0xe7, 0xb9, 0x43, 0x38, 0xc8, 0x81, 0x70, 0x6c, 0xbc,
	0xd6, 0x07, 0x94, 0x5f, 0x8c, 0xfa, 0x1d, 0xdb, 0x1f, 0x9a, 0x57, 0xa4, 0x4f, 0x39, 0x71, 0x4d,
	0xf1, 0x27, 0x86, 0x99, 0x71, 0xed, 0xbe, 0x2a, 0x5e, 0x9f, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff,
	0x38, 0x04, 0x77, 0xfe, 0x1b, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AgentPauseCauseServiceClient is the client API for AgentPauseCauseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AgentPauseCauseServiceClient interface {
	CreateAgentPauseCause(ctx context.Context, in *CreateAgentPauseCauseRequest, opts ...grpc.CallOption) (*AgentPauseCause, error)
	SearchAgentPauseCause(ctx context.Context, in *SearchAgentPauseCauseRequest, opts ...grpc.CallOption) (*ListAgentPauseCause, error)
	ReadAgentPauseCause(ctx context.Context, in *ReadAgentPauseCauseRequest, opts ...grpc.CallOption) (*AgentPauseCause, error)
	PatchAgentPauseCause(ctx context.Context, in *PatchAgentPauseCauseRequest, opts ...grpc.CallOption) (*AgentPauseCause, error)
	UpdateAgentPauseCause(ctx context.Context, in *UpdateAgentPauseCauseRequest, opts ...grpc.CallOption) (*AgentPauseCause, error)
	DeleteAgentPauseCause(ctx context.Context, in *DeleteAgentPauseCauseRequest, opts ...grpc.CallOption) (*AgentPauseCause, error)
}

type agentPauseCauseServiceClient struct {
	cc *grpc.ClientConn
}

func NewAgentPauseCauseServiceClient(cc *grpc.ClientConn) AgentPauseCauseServiceClient {
	return &agentPauseCauseServiceClient{cc}
}

func (c *agentPauseCauseServiceClient) CreateAgentPauseCause(ctx context.Context, in *CreateAgentPauseCauseRequest, opts ...grpc.CallOption) (*AgentPauseCause, error) {
	out := new(AgentPauseCause)
	err := c.cc.Invoke(ctx, "/engine.AgentPauseCauseService/CreateAgentPauseCause", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentPauseCauseServiceClient) SearchAgentPauseCause(ctx context.Context, in *SearchAgentPauseCauseRequest, opts ...grpc.CallOption) (*ListAgentPauseCause, error) {
	out := new(ListAgentPauseCause)
	err := c.cc.Invoke(ctx, "/engine.AgentPauseCauseService/SearchAgentPauseCause", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentPauseCauseServiceClient) ReadAgentPauseCause(ctx context.Context, in *ReadAgentPauseCauseRequest, opts ...grpc.CallOption) (*AgentPauseCause, error) {
	out := new(AgentPauseCause)
	err := c.cc.Invoke(ctx, "/engine.AgentPauseCauseService/ReadAgentPauseCause", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentPauseCauseServiceClient) PatchAgentPauseCause(ctx context.Context, in *PatchAgentPauseCauseRequest, opts ...grpc.CallOption) (*AgentPauseCause, error) {
	out := new(AgentPauseCause)
	err := c.cc.Invoke(ctx, "/engine.AgentPauseCauseService/PatchAgentPauseCause", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentPauseCauseServiceClient) UpdateAgentPauseCause(ctx context.Context, in *UpdateAgentPauseCauseRequest, opts ...grpc.CallOption) (*AgentPauseCause, error) {
	out := new(AgentPauseCause)
	err := c.cc.Invoke(ctx, "/engine.AgentPauseCauseService/UpdateAgentPauseCause", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentPauseCauseServiceClient) DeleteAgentPauseCause(ctx context.Context, in *DeleteAgentPauseCauseRequest, opts ...grpc.CallOption) (*AgentPauseCause, error) {
	out := new(AgentPauseCause)
	err := c.cc.Invoke(ctx, "/engine.AgentPauseCauseService/DeleteAgentPauseCause", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentPauseCauseServiceServer is the server API for AgentPauseCauseService service.
type AgentPauseCauseServiceServer interface {
	CreateAgentPauseCause(context.Context, *CreateAgentPauseCauseRequest) (*AgentPauseCause, error)
	SearchAgentPauseCause(context.Context, *SearchAgentPauseCauseRequest) (*ListAgentPauseCause, error)
	ReadAgentPauseCause(context.Context, *ReadAgentPauseCauseRequest) (*AgentPauseCause, error)
	PatchAgentPauseCause(context.Context, *PatchAgentPauseCauseRequest) (*AgentPauseCause, error)
	UpdateAgentPauseCause(context.Context, *UpdateAgentPauseCauseRequest) (*AgentPauseCause, error)
	DeleteAgentPauseCause(context.Context, *DeleteAgentPauseCauseRequest) (*AgentPauseCause, error)
}

// UnimplementedAgentPauseCauseServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAgentPauseCauseServiceServer struct {
}

func (*UnimplementedAgentPauseCauseServiceServer) CreateAgentPauseCause(ctx context.Context, req *CreateAgentPauseCauseRequest) (*AgentPauseCause, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAgentPauseCause not implemented")
}
func (*UnimplementedAgentPauseCauseServiceServer) SearchAgentPauseCause(ctx context.Context, req *SearchAgentPauseCauseRequest) (*ListAgentPauseCause, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchAgentPauseCause not implemented")
}
func (*UnimplementedAgentPauseCauseServiceServer) ReadAgentPauseCause(ctx context.Context, req *ReadAgentPauseCauseRequest) (*AgentPauseCause, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadAgentPauseCause not implemented")
}
func (*UnimplementedAgentPauseCauseServiceServer) PatchAgentPauseCause(ctx context.Context, req *PatchAgentPauseCauseRequest) (*AgentPauseCause, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchAgentPauseCause not implemented")
}
func (*UnimplementedAgentPauseCauseServiceServer) UpdateAgentPauseCause(ctx context.Context, req *UpdateAgentPauseCauseRequest) (*AgentPauseCause, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAgentPauseCause not implemented")
}
func (*UnimplementedAgentPauseCauseServiceServer) DeleteAgentPauseCause(ctx context.Context, req *DeleteAgentPauseCauseRequest) (*AgentPauseCause, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAgentPauseCause not implemented")
}

func RegisterAgentPauseCauseServiceServer(s *grpc.Server, srv AgentPauseCauseServiceServer) {
	s.RegisterService(&_AgentPauseCauseService_serviceDesc, srv)
}

func _AgentPauseCauseService_CreateAgentPauseCause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAgentPauseCauseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentPauseCauseServiceServer).CreateAgentPauseCause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.AgentPauseCauseService/CreateAgentPauseCause",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentPauseCauseServiceServer).CreateAgentPauseCause(ctx, req.(*CreateAgentPauseCauseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentPauseCauseService_SearchAgentPauseCause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchAgentPauseCauseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentPauseCauseServiceServer).SearchAgentPauseCause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.AgentPauseCauseService/SearchAgentPauseCause",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentPauseCauseServiceServer).SearchAgentPauseCause(ctx, req.(*SearchAgentPauseCauseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentPauseCauseService_ReadAgentPauseCause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadAgentPauseCauseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentPauseCauseServiceServer).ReadAgentPauseCause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.AgentPauseCauseService/ReadAgentPauseCause",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentPauseCauseServiceServer).ReadAgentPauseCause(ctx, req.(*ReadAgentPauseCauseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentPauseCauseService_PatchAgentPauseCause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchAgentPauseCauseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentPauseCauseServiceServer).PatchAgentPauseCause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.AgentPauseCauseService/PatchAgentPauseCause",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentPauseCauseServiceServer).PatchAgentPauseCause(ctx, req.(*PatchAgentPauseCauseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentPauseCauseService_UpdateAgentPauseCause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAgentPauseCauseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentPauseCauseServiceServer).UpdateAgentPauseCause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.AgentPauseCauseService/UpdateAgentPauseCause",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentPauseCauseServiceServer).UpdateAgentPauseCause(ctx, req.(*UpdateAgentPauseCauseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentPauseCauseService_DeleteAgentPauseCause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAgentPauseCauseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentPauseCauseServiceServer).DeleteAgentPauseCause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.AgentPauseCauseService/DeleteAgentPauseCause",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentPauseCauseServiceServer).DeleteAgentPauseCause(ctx, req.(*DeleteAgentPauseCauseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AgentPauseCauseService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "engine.AgentPauseCauseService",
	HandlerType: (*AgentPauseCauseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAgentPauseCause",
			Handler:    _AgentPauseCauseService_CreateAgentPauseCause_Handler,
		},
		{
			MethodName: "SearchAgentPauseCause",
			Handler:    _AgentPauseCauseService_SearchAgentPauseCause_Handler,
		},
		{
			MethodName: "ReadAgentPauseCause",
			Handler:    _AgentPauseCauseService_ReadAgentPauseCause_Handler,
		},
		{
			MethodName: "PatchAgentPauseCause",
			Handler:    _AgentPauseCauseService_PatchAgentPauseCause_Handler,
		},
		{
			MethodName: "UpdateAgentPauseCause",
			Handler:    _AgentPauseCauseService_UpdateAgentPauseCause_Handler,
		},
		{
			MethodName: "DeleteAgentPauseCause",
			Handler:    _AgentPauseCauseService_DeleteAgentPauseCause_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pause_cause.proto",
}
