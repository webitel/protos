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
	Sort                 string   `protobuf:"bytes,4,opt,name=sort,proto3" json:"sort,omitempty"`
	Fields               []string `protobuf:"bytes,5,rep,name=fields,proto3" json:"fields,omitempty"`
	Id                   []uint32 `protobuf:"varint,6,rep,packed,name=id,proto3" json:"id,omitempty"`
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

func (m *SearchAgentPauseCauseRequest) GetSort() string {
	if m != nil {
		return m.Sort
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
	// 737 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0x41, 0x6f, 0xd3, 0x4c,
	0x10, 0x95, 0xed, 0xd8, 0x8d, 0x27, 0x5f, 0xdb, 0x8f, 0x2d, 0x29, 0x26, 0x09, 0xc5, 0xb8, 0x15,
	0x4a, 0x2b, 0x1a, 0x4b, 0xe1, 0xc6, 0xad, 0x2d, 0xc7, 0x22, 0x55, 0xae, 0x90, 0x10, 0x97, 0x68,
	0x63, 0x2f, 0xe9, 0x0a, 0xc7, 0x76, 0xed, 0x0d, 0x25, 0x85, 0x82, 0x40, 0xe2, 0x8e, 0x84, 0xc4,
	0x1f, 0xe3, 0xce, 0x89, 0x33, 0xff, 0x00, 0x09, 0x79, 0x1d, 0x27, 0x21, 0xb1, 0x63, 0x5f, 0xac,
	0xf5, 0x78, 0x34, 0xf3, 0xde, 0x9b, 0x37, 0x9b, 0xc0, 0xad, 0x00, 0x8f, 0x22, 0xd2, 0xb3, 0xe3,
	0x67, 0x27, 0x08, 0x7d, 0xe6, 0x23, 0x85, 0x78, 0x03, 0xea, 0x91, 0x46, 0x6b, 0xe0, 0xfb, 0x03,
	0x97, 0x98, 0x38, 0xa0, 0x26, 0xf6, 0x3c, 0x9f, 0x61, 0x46, 0x7d, 0x2f, 0x4a, 0xb2, 0x1a, 0x35,
	0xdb, 0xf7, 0x22, 0x96, 0xbc, 0x18, 0x3f, 0x05, 0x68, 0x9d, 0x84, 0x04, 0x33, 0x72, 0x34, 0x20,
	0x1e, 0x3b, 0x8b, 0xab, 0x9d, 0xc4, 0x0f, 0x8b, 0x5c, 0x8e, 0x48, 0xc4, 0x10, 0x82, 0x8a, 0x87,
	0x87, 0x44, 0x13, 0x74, 0xa1, 0xad, 0x5a, 0xfc, 0x8c, 0x9a, 0xa0, 0xba, 0x74, 0x48, 0x59, 0x6f,
	0x48, 0x3d, 0x4d, 0xd4, 0x85, 0xf6, 0xba, 0x55, 0xe5, 0x81, 0x67, 0xd4, 0x43, 0xfb, 0xf0, 0x3f,
	0x76, 0x5d, 0xff, 0xaa, 0x17, 0x8d, 0x02, 0x12, 0xbe, 0xa1, 0x91, 0x1f, 0x6a, 0x92, 0x2e, 0xb4,
	0xab, 0xd6, 0x26, 0x8f, 0x9f, 0x4f, 0xc3, 0xe8, 0x3e, 0xd4, 0x92, 0x54, 0x1c, 0xf7, 0xd6, 0x2a,
	0x3c, 0x0b, 0x78, 0x88, 0xa3, 0x99, 0x4b, 0x70, 0xe2, 0x56, 0xf2, 0x7c, 0x42, 0x1c, 0x41, 0x3a,
	0xd4, 0x1c, 0x12, 0xd9, 0x21, 0x0d, 0x62, 0x86, 0x9a, 0xc2, 0x41, 0xce, 0x87, 0x8c, 0x3f, 0x22,
	0x6c, 0x2e, 0x50, 0x43, 0x1b, 0x20, 0x52, 0x87, 0x33, 0x5a, 0xb7, 0x44, 0xea, 0xa0, 0x7b, 0x00,
	0x36, 0xd7, 0xc0, 0xe9, 0x61, 0xc6, 0x09, 0x49, 0x96, 0x3a, 0x89, 0x1c, 0x31, 0x74, 0x38, 0xfb,
	0xdc, 0x1f, 0x73, 0x2e, 0xb5, 0xee, 0x46, 0x27, 0xd1, 0xba, 0x73, 0xea, 0xfb, 0xaf, 0x47, 0xc1,
	0x34, 0xfd, 0x78, 0x1c, 0x57, 0x1b, 0x05, 0x4e, 0x5a, 0xad, 0x92, 0x54, 0x9b, 0x44, 0x92, 0x6a,
	0xe9, 0xe7, 0xfe, 0x98, 0x53, 0xca, 0xa8, 0x36, 0xc9, 0x38, 0x1e, 0x4f, 0xf5, 0x57, 0xf2, 0xf4,
	0x5f, 0x2b, 0xa1, 0x7f, 0xb5, 0x94, 0xfe, 0x6a, 0x91, 0xfe, 0x50, 0xa4, 0x7f, 0x6d, 0x59, 0xff,
	0xaf, 0x02, 0xb4, 0xce, 0x09, 0x0e, 0xed, 0x8b, 0x7c, 0x83, 0x05, 0x78, 0x90, 0x18, 0x4c, 0xb6,
	0xf8, 0x39, 0x8e, 0x45, 0xf4, 0x9a, 0xf0, 0x51, 0xc8, 0x16, 0x3f, 0xa3, 0xff, 0x40, 0xb8, 0xe4,
	0xe2, 0xab, 0x96, 0x70, 0xc9, 0x33, 0xfc, 0x30, 0x91, 0x57, 0xb5, 0xf8, 0x19, 0x6d, 0x83, 0xf2,
	0x8a, 0x12, 0xd7, 0x89, 0x34, 0x59, 0x97, 0xda, 0xaa, 0x35, 0x79, 0x9b, 0x8c, 0x5b, 0xd1, 0xa5,
	0x64, 0xdc, 0xc6, 0x0b, 0xd8, 0x3a, 0xa5, 0x11, 0x5b, 0x74, 0x45, 0xac, 0x34, 0x79, 0xcb, 0x38,
	0x90, 0xaa, 0xc5, 0xcf, 0xe8, 0x10, 0x64, 0xca, 0xc8, 0x30, 0xd2, 0x44, 0x5d, 0x6a, 0xd7, 0xba,
	0x77, 0xd2, 0x39, 0x2d, 0x72, 0x49, 0xb2, 0x8c, 0x47, 0xd0, 0xb0, 0x08, 0x76, 0x72, 0x98, 0x2e,
	0xd8, 0xce, 0xf8, 0x22, 0x42, 0xf3, 0x0c, 0xb3, 0x5c, 0x65, 0x16, 0x6d, 0x3a, 0xe3, 0x27, 0xfe,
	0xc3, 0x2f, 0xb5, 0x88, 0x94, 0x67, 0x91, 0x4a, 0x09, 0x8b, 0xc8, 0xa5, 0x2c, 0xa2, 0x14, 0x59,
	0x64, 0xad, 0xc8, 0x22, 0xd5, 0x65, 0x8b, 0xfc, 0x16, 0xa0, 0xf5, 0x9c, 0x1b, 0xbe, 0xa4, 0x10,
	0x29, 0x61, 0x31, 0x8f, 0xb0, 0x54, 0x82, 0x70, 0xa5, 0x14, 0x61, 0xb9, 0x88, 0xb0, 0x52, 0x44,
	0x78, 0x6d, 0x99, 0x70, 0x07, 0x5a, 0x4f, 0x89, 0x4b, 0xca, 0xf2, 0xed, 0x7e, 0x57, 0x60, 0x7b,
	0x21, 0xf5, 0x3c, 0xc6, 0x6b, 0x13, 0xf4, 0x1e, 0xea, 0x99, 0xd7, 0x37, 0xda, 0x4b, 0xad, 0xba,
	0xea, 0x76, 0x6f, 0xe4, 0x19, 0xda, 0xd8, 0xfb, 0xfc, 0xe3, 0xd7, 0x37, 0x71, 0xc7, 0xb8, 0x6b,
	0xda, 0xd8, 0x75, 0x7b, 0x36, 0xf1, 0x18, 0x09, 0xcd, 0xb9, 0x9f, 0x9c, 0xe8, 0x89, 0x70, 0x80,
	0x6e, 0xa0, 0x9e, 0xb9, 0xdb, 0xb3, 0xee, 0xab, 0x56, 0xbf, 0xd1, 0x9c, 0x5e, 0x7b, 0xcb, 0xeb,
	0x68, 0x3c, 0xe0, 0x08, 0x9a, 0x28, 0x1f, 0x01, 0xba, 0x86, 0xad, 0x8c, 0x75, 0x43, 0x46, 0x5a,
	0x36, 0x7f, 0x17, 0xf3, 0x89, 0x3f, 0xe4, 0x6d, 0x75, 0xb4, 0x93, 0xdb, 0xd6, 0x7c, 0x47, 0x9d,
	0x1b, 0xf4, 0x11, 0x6e, 0x67, 0xed, 0x2e, 0xda, 0x4d, 0x0b, 0xaf, 0xd8, 0xec, 0xfc, 0xee, 0xfb,
	0xbc, 0xfb, 0x6e, 0xb7, 0xa0, 0x7b, 0xac, 0xfd, 0x27, 0x01, 0xea, 0x99, 0x5b, 0x33, 0x13, 0x7f,
	0xd5, 0x52, 0x15, 0x62, 0x68, 0x94, 0xc0, 0xf0, 0x01, 0xea, 0x99, 0x46, 0x9e, 0x41, 0x58, 0xe5,
	0xf3, 0xc2, 0x21, 0x1c, 0x14, 0x40, 0x38, 0x36, 0x5e, 0xea, 0x03, 0xca, 0x2e, 0x46, 0xfd, 0x8e,
	0xed, 0x0f, 0xcd, 0x2b, 0xd2, 0xa7, 0x8c, 0xb8, 0x26, 0xff, 0x67, 0x13, 0x99, 0x49, 0xed, 0xbe,
	0xc2, 0x5f, 0x1f, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x32, 0xfc, 0xd0, 0x49, 0x30, 0x09, 0x00,
	0x00,
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
