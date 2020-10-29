// Code generated by protoc-gen-go. DO NOT EDIT.
// source: agent_skill.proto

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

type SearchLookupAgentNotExistsSkillRequest struct {
	AgentId              int64    `protobuf:"varint,1,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	Page                 int32    `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Size                 int32    `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Q                    string   `protobuf:"bytes,4,opt,name=q,proto3" json:"q,omitempty"`
	DomainId             int64    `protobuf:"varint,5,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchLookupAgentNotExistsSkillRequest) Reset() {
	*m = SearchLookupAgentNotExistsSkillRequest{}
}
func (m *SearchLookupAgentNotExistsSkillRequest) String() string { return proto.CompactTextString(m) }
func (*SearchLookupAgentNotExistsSkillRequest) ProtoMessage()    {}
func (*SearchLookupAgentNotExistsSkillRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8423e2df0c2237d, []int{0}
}

func (m *SearchLookupAgentNotExistsSkillRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchLookupAgentNotExistsSkillRequest.Unmarshal(m, b)
}
func (m *SearchLookupAgentNotExistsSkillRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchLookupAgentNotExistsSkillRequest.Marshal(b, m, deterministic)
}
func (m *SearchLookupAgentNotExistsSkillRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchLookupAgentNotExistsSkillRequest.Merge(m, src)
}
func (m *SearchLookupAgentNotExistsSkillRequest) XXX_Size() int {
	return xxx_messageInfo_SearchLookupAgentNotExistsSkillRequest.Size(m)
}
func (m *SearchLookupAgentNotExistsSkillRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchLookupAgentNotExistsSkillRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchLookupAgentNotExistsSkillRequest proto.InternalMessageInfo

func (m *SearchLookupAgentNotExistsSkillRequest) GetAgentId() int64 {
	if m != nil {
		return m.AgentId
	}
	return 0
}

func (m *SearchLookupAgentNotExistsSkillRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *SearchLookupAgentNotExistsSkillRequest) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *SearchLookupAgentNotExistsSkillRequest) GetQ() string {
	if m != nil {
		return m.Q
	}
	return ""
}

func (m *SearchLookupAgentNotExistsSkillRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

type UpdateAgentSkillRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AgentId              int64    `protobuf:"varint,2,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	Skill                *Lookup  `protobuf:"bytes,3,opt,name=skill,proto3" json:"skill,omitempty"`
	Capacity             int32    `protobuf:"varint,4,opt,name=capacity,proto3" json:"capacity,omitempty"`
	DomainId             int64    `protobuf:"varint,5,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateAgentSkillRequest) Reset()         { *m = UpdateAgentSkillRequest{} }
func (m *UpdateAgentSkillRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateAgentSkillRequest) ProtoMessage()    {}
func (*UpdateAgentSkillRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8423e2df0c2237d, []int{1}
}

func (m *UpdateAgentSkillRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateAgentSkillRequest.Unmarshal(m, b)
}
func (m *UpdateAgentSkillRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateAgentSkillRequest.Marshal(b, m, deterministic)
}
func (m *UpdateAgentSkillRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateAgentSkillRequest.Merge(m, src)
}
func (m *UpdateAgentSkillRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateAgentSkillRequest.Size(m)
}
func (m *UpdateAgentSkillRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateAgentSkillRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateAgentSkillRequest proto.InternalMessageInfo

func (m *UpdateAgentSkillRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateAgentSkillRequest) GetAgentId() int64 {
	if m != nil {
		return m.AgentId
	}
	return 0
}

func (m *UpdateAgentSkillRequest) GetSkill() *Lookup {
	if m != nil {
		return m.Skill
	}
	return nil
}

func (m *UpdateAgentSkillRequest) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *UpdateAgentSkillRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

type SearchAgentSkillRequest struct {
	AgentId              int64    `protobuf:"varint,1,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	Page                 int32    `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Size                 int32    `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Q                    string   `protobuf:"bytes,4,opt,name=q,proto3" json:"q,omitempty"`
	DomainId             int64    `protobuf:"varint,5,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchAgentSkillRequest) Reset()         { *m = SearchAgentSkillRequest{} }
func (m *SearchAgentSkillRequest) String() string { return proto.CompactTextString(m) }
func (*SearchAgentSkillRequest) ProtoMessage()    {}
func (*SearchAgentSkillRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8423e2df0c2237d, []int{2}
}

func (m *SearchAgentSkillRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchAgentSkillRequest.Unmarshal(m, b)
}
func (m *SearchAgentSkillRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchAgentSkillRequest.Marshal(b, m, deterministic)
}
func (m *SearchAgentSkillRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchAgentSkillRequest.Merge(m, src)
}
func (m *SearchAgentSkillRequest) XXX_Size() int {
	return xxx_messageInfo_SearchAgentSkillRequest.Size(m)
}
func (m *SearchAgentSkillRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchAgentSkillRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchAgentSkillRequest proto.InternalMessageInfo

func (m *SearchAgentSkillRequest) GetAgentId() int64 {
	if m != nil {
		return m.AgentId
	}
	return 0
}

func (m *SearchAgentSkillRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *SearchAgentSkillRequest) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *SearchAgentSkillRequest) GetQ() string {
	if m != nil {
		return m.Q
	}
	return ""
}

func (m *SearchAgentSkillRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

type CreateAgentSkillRequest struct {
	AgentId              int64    `protobuf:"varint,1,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	Skill                *Lookup  `protobuf:"bytes,2,opt,name=skill,proto3" json:"skill,omitempty"`
	Capacity             int32    `protobuf:"varint,3,opt,name=capacity,proto3" json:"capacity,omitempty"`
	DomainId             int64    `protobuf:"varint,4,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateAgentSkillRequest) Reset()         { *m = CreateAgentSkillRequest{} }
func (m *CreateAgentSkillRequest) String() string { return proto.CompactTextString(m) }
func (*CreateAgentSkillRequest) ProtoMessage()    {}
func (*CreateAgentSkillRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8423e2df0c2237d, []int{3}
}

func (m *CreateAgentSkillRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAgentSkillRequest.Unmarshal(m, b)
}
func (m *CreateAgentSkillRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAgentSkillRequest.Marshal(b, m, deterministic)
}
func (m *CreateAgentSkillRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAgentSkillRequest.Merge(m, src)
}
func (m *CreateAgentSkillRequest) XXX_Size() int {
	return xxx_messageInfo_CreateAgentSkillRequest.Size(m)
}
func (m *CreateAgentSkillRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAgentSkillRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAgentSkillRequest proto.InternalMessageInfo

func (m *CreateAgentSkillRequest) GetAgentId() int64 {
	if m != nil {
		return m.AgentId
	}
	return 0
}

func (m *CreateAgentSkillRequest) GetSkill() *Lookup {
	if m != nil {
		return m.Skill
	}
	return nil
}

func (m *CreateAgentSkillRequest) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *CreateAgentSkillRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

type AgentSkillItemRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AgentId              int64    `protobuf:"varint,2,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	DomainId             int64    `protobuf:"varint,3,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgentSkillItemRequest) Reset()         { *m = AgentSkillItemRequest{} }
func (m *AgentSkillItemRequest) String() string { return proto.CompactTextString(m) }
func (*AgentSkillItemRequest) ProtoMessage()    {}
func (*AgentSkillItemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8423e2df0c2237d, []int{4}
}

func (m *AgentSkillItemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgentSkillItemRequest.Unmarshal(m, b)
}
func (m *AgentSkillItemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgentSkillItemRequest.Marshal(b, m, deterministic)
}
func (m *AgentSkillItemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentSkillItemRequest.Merge(m, src)
}
func (m *AgentSkillItemRequest) XXX_Size() int {
	return xxx_messageInfo_AgentSkillItemRequest.Size(m)
}
func (m *AgentSkillItemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentSkillItemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AgentSkillItemRequest proto.InternalMessageInfo

func (m *AgentSkillItemRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AgentSkillItemRequest) GetAgentId() int64 {
	if m != nil {
		return m.AgentId
	}
	return 0
}

func (m *AgentSkillItemRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

type DeleteAgentSkillRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AgentId              int64    `protobuf:"varint,2,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	DomainId             int64    `protobuf:"varint,3,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteAgentSkillRequest) Reset()         { *m = DeleteAgentSkillRequest{} }
func (m *DeleteAgentSkillRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteAgentSkillRequest) ProtoMessage()    {}
func (*DeleteAgentSkillRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8423e2df0c2237d, []int{5}
}

func (m *DeleteAgentSkillRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteAgentSkillRequest.Unmarshal(m, b)
}
func (m *DeleteAgentSkillRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteAgentSkillRequest.Marshal(b, m, deterministic)
}
func (m *DeleteAgentSkillRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteAgentSkillRequest.Merge(m, src)
}
func (m *DeleteAgentSkillRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteAgentSkillRequest.Size(m)
}
func (m *DeleteAgentSkillRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteAgentSkillRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteAgentSkillRequest proto.InternalMessageInfo

func (m *DeleteAgentSkillRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DeleteAgentSkillRequest) GetAgentId() int64 {
	if m != nil {
		return m.AgentId
	}
	return 0
}

func (m *DeleteAgentSkillRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

type AgentSkillItem struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Skill                *Lookup  `protobuf:"bytes,2,opt,name=skill,proto3" json:"skill,omitempty"`
	Capacity             int32    `protobuf:"varint,3,opt,name=capacity,proto3" json:"capacity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgentSkillItem) Reset()         { *m = AgentSkillItem{} }
func (m *AgentSkillItem) String() string { return proto.CompactTextString(m) }
func (*AgentSkillItem) ProtoMessage()    {}
func (*AgentSkillItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8423e2df0c2237d, []int{6}
}

func (m *AgentSkillItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgentSkillItem.Unmarshal(m, b)
}
func (m *AgentSkillItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgentSkillItem.Marshal(b, m, deterministic)
}
func (m *AgentSkillItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentSkillItem.Merge(m, src)
}
func (m *AgentSkillItem) XXX_Size() int {
	return xxx_messageInfo_AgentSkillItem.Size(m)
}
func (m *AgentSkillItem) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentSkillItem.DiscardUnknown(m)
}

var xxx_messageInfo_AgentSkillItem proto.InternalMessageInfo

func (m *AgentSkillItem) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AgentSkillItem) GetSkill() *Lookup {
	if m != nil {
		return m.Skill
	}
	return nil
}

func (m *AgentSkillItem) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

type AgentSkill struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DomainId             int64    `protobuf:"varint,2,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	CreatedAt            int64    `protobuf:"varint,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	CreatedBy            *Lookup  `protobuf:"bytes,4,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	UpdatedAt            int64    `protobuf:"varint,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	UpdatedBy            *Lookup  `protobuf:"bytes,6,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
	Agent                *Lookup  `protobuf:"bytes,7,opt,name=agent,proto3" json:"agent,omitempty"`
	Skill                *Lookup  `protobuf:"bytes,8,opt,name=skill,proto3" json:"skill,omitempty"`
	Capacity             int32    `protobuf:"varint,9,opt,name=capacity,proto3" json:"capacity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgentSkill) Reset()         { *m = AgentSkill{} }
func (m *AgentSkill) String() string { return proto.CompactTextString(m) }
func (*AgentSkill) ProtoMessage()    {}
func (*AgentSkill) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8423e2df0c2237d, []int{7}
}

func (m *AgentSkill) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgentSkill.Unmarshal(m, b)
}
func (m *AgentSkill) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgentSkill.Marshal(b, m, deterministic)
}
func (m *AgentSkill) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentSkill.Merge(m, src)
}
func (m *AgentSkill) XXX_Size() int {
	return xxx_messageInfo_AgentSkill.Size(m)
}
func (m *AgentSkill) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentSkill.DiscardUnknown(m)
}

var xxx_messageInfo_AgentSkill proto.InternalMessageInfo

func (m *AgentSkill) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AgentSkill) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

func (m *AgentSkill) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *AgentSkill) GetCreatedBy() *Lookup {
	if m != nil {
		return m.CreatedBy
	}
	return nil
}

func (m *AgentSkill) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *AgentSkill) GetUpdatedBy() *Lookup {
	if m != nil {
		return m.UpdatedBy
	}
	return nil
}

func (m *AgentSkill) GetAgent() *Lookup {
	if m != nil {
		return m.Agent
	}
	return nil
}

func (m *AgentSkill) GetSkill() *Lookup {
	if m != nil {
		return m.Skill
	}
	return nil
}

func (m *AgentSkill) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

type ListAgentSkill struct {
	Next                 bool              `protobuf:"varint,1,opt,name=next,proto3" json:"next,omitempty"`
	Items                []*AgentSkillItem `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListAgentSkill) Reset()         { *m = ListAgentSkill{} }
func (m *ListAgentSkill) String() string { return proto.CompactTextString(m) }
func (*ListAgentSkill) ProtoMessage()    {}
func (*ListAgentSkill) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8423e2df0c2237d, []int{8}
}

func (m *ListAgentSkill) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAgentSkill.Unmarshal(m, b)
}
func (m *ListAgentSkill) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAgentSkill.Marshal(b, m, deterministic)
}
func (m *ListAgentSkill) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAgentSkill.Merge(m, src)
}
func (m *ListAgentSkill) XXX_Size() int {
	return xxx_messageInfo_ListAgentSkill.Size(m)
}
func (m *ListAgentSkill) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAgentSkill.DiscardUnknown(m)
}

var xxx_messageInfo_ListAgentSkill proto.InternalMessageInfo

func (m *ListAgentSkill) GetNext() bool {
	if m != nil {
		return m.Next
	}
	return false
}

func (m *ListAgentSkill) GetItems() []*AgentSkillItem {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*SearchLookupAgentNotExistsSkillRequest)(nil), "engine.SearchLookupAgentNotExistsSkillRequest")
	proto.RegisterType((*UpdateAgentSkillRequest)(nil), "engine.UpdateAgentSkillRequest")
	proto.RegisterType((*SearchAgentSkillRequest)(nil), "engine.SearchAgentSkillRequest")
	proto.RegisterType((*CreateAgentSkillRequest)(nil), "engine.CreateAgentSkillRequest")
	proto.RegisterType((*AgentSkillItemRequest)(nil), "engine.AgentSkillItemRequest")
	proto.RegisterType((*DeleteAgentSkillRequest)(nil), "engine.DeleteAgentSkillRequest")
	proto.RegisterType((*AgentSkillItem)(nil), "engine.AgentSkillItem")
	proto.RegisterType((*AgentSkill)(nil), "engine.AgentSkill")
	proto.RegisterType((*ListAgentSkill)(nil), "engine.ListAgentSkill")
}

func init() { proto.RegisterFile("agent_skill.proto", fileDescriptor_c8423e2df0c2237d) }

var fileDescriptor_c8423e2df0c2237d = []byte{
	// 696 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x96, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0x65, 0xe7, 0xa3, 0xc9, 0x04, 0x45, 0xe9, 0x4a, 0x10, 0x63, 0xa8, 0x1a, 0x59, 0x7c,
	0x44, 0x51, 0x1b, 0xa3, 0x20, 0x2e, 0xdc, 0x5a, 0xe0, 0x50, 0xa9, 0xe2, 0xe0, 0x8a, 0x0b, 0x97,
	0x68, 0x63, 0xaf, 0xd2, 0x55, 0x1d, 0x6f, 0x9a, 0xdd, 0xd0, 0xa6, 0x55, 0x39, 0x70, 0xe1, 0x01,
	0x2a, 0x71, 0x41, 0xe2, 0xc2, 0x73, 0xf0, 0x14, 0xbc, 0x02, 0x0f, 0x82, 0xbc, 0x1b, 0xc7, 0xb1,
	0x93, 0xb8, 0x29, 0x95, 0xb8, 0x79, 0x67, 0x56, 0xf3, 0xff, 0xcf, 0x2f, 0xb3, 0xbb, 0x81, 0x4d,
	0xdc, 0x27, 0x81, 0xe8, 0xf2, 0x13, 0xea, 0xfb, 0xed, 0xe1, 0x88, 0x09, 0x86, 0x8a, 0x24, 0xe8,
	0xd3, 0x80, 0x98, 0x15, 0x97, 0x05, 0x5c, 0xa8, 0xa0, 0x59, 0x99, 0xdb, 0x61, 0x3e, 0xee, 0x33,
	0xd6, 0xf7, 0x89, 0x8d, 0x87, 0xd4, 0xc6, 0x41, 0xc0, 0x04, 0x16, 0x94, 0x05, 0x5c, 0x65, 0xad,
	0x6f, 0x1a, 0x3c, 0x3b, 0x22, 0x78, 0xe4, 0x1e, 0x1f, 0x32, 0x76, 0x32, 0x1e, 0xee, 0x85, 0x0a,
	0xef, 0x99, 0x78, 0x77, 0x4e, 0xb9, 0xe0, 0x47, 0x61, 0x1d, 0x87, 0x9c, 0x8e, 0x09, 0x17, 0xe8,
	0x21, 0x94, 0x94, 0x3e, 0xf5, 0x0c, 0xad, 0xa1, 0x35, 0x73, 0xce, 0x86, 0x5c, 0x1f, 0x78, 0x08,
	0x41, 0x7e, 0x88, 0xfb, 0xc4, 0xd0, 0x1b, 0x5a, 0xb3, 0xe0, 0xc8, 0xef, 0x30, 0xc6, 0xe9, 0x05,
	0x31, 0x72, 0x2a, 0x16, 0x7e, 0xa3, 0x7b, 0xa0, 0x9d, 0x1a, 0xf9, 0x86, 0xd6, 0x2c, 0x3b, 0xda,
	0x29, 0x7a, 0x04, 0x65, 0x8f, 0x0d, 0x30, 0x0d, 0xc2, 0x8a, 0x05, 0x59, 0xb1, 0xa4, 0x02, 0x07,
	0x9e, 0xf5, 0x53, 0x83, 0xfa, 0x87, 0xa1, 0x87, 0x05, 0x91, 0x96, 0x12, 0x4e, 0xaa, 0xa0, 0xcf,
	0x3c, 0xe8, 0xd4, 0x4b, 0x38, 0xd3, 0x93, 0xce, 0x9e, 0x40, 0x41, 0xc2, 0x90, 0x36, 0x2a, 0x9d,
	0x6a, 0x5b, 0xf1, 0x6a, 0xab, 0x6e, 0x1d, 0x95, 0x44, 0x26, 0x94, 0x5c, 0x3c, 0xc4, 0x2e, 0x15,
	0x13, 0x69, 0xaf, 0xe0, 0xcc, 0xd6, 0xd9, 0x2e, 0xbf, 0x6a, 0x50, 0x57, 0xf8, 0x16, 0x5d, 0xfe,
	0x5f, 0x5e, 0xd7, 0x1a, 0xd4, 0xdf, 0x8c, 0xc8, 0x52, 0x5e, 0x19, 0x4e, 0x66, 0x7c, 0xf4, 0x75,
	0xf9, 0xe4, 0xb2, 0xf8, 0xe4, 0x53, 0xae, 0xba, 0x70, 0x3f, 0xb6, 0x73, 0x20, 0xc8, 0xe0, 0x1f,
	0x7e, 0xc2, 0x84, 0x40, 0x2e, 0x25, 0x80, 0xa1, 0xfe, 0x96, 0xf8, 0xe4, 0x8e, 0x53, 0x92, 0x29,
	0xd1, 0x83, 0x6a, 0xb2, 0x87, 0x85, 0xca, 0x77, 0x86, 0x68, 0xfd, 0xd2, 0x01, 0x62, 0x91, 0x05,
	0x81, 0x84, 0x3f, 0x3d, 0xe9, 0x0f, 0x6d, 0x01, 0xb8, 0xf2, 0x87, 0xf7, 0xba, 0x58, 0x4c, 0xdd,
	0x97, 0xa7, 0x91, 0x3d, 0x81, 0x76, 0xe3, 0x74, 0x4f, 0x4d, 0xf7, 0xa2, 0xc3, 0x68, 0xfb, 0xfe,
	0x24, 0xac, 0x36, 0x96, 0xc7, 0x4e, 0x56, 0x53, 0x53, 0x56, 0x9e, 0x46, 0x54, 0xb5, 0x28, 0xdd,
	0x9b, 0x18, 0xc5, 0xe5, 0xd5, 0xa6, 0x3b, 0xf6, 0x27, 0x21, 0x19, 0xc9, 0xd8, 0xd8, 0x58, 0x4e,
	0x46, 0x26, 0x63, 0x7e, 0xa5, 0x75, 0xf9, 0x95, 0x53, 0xfc, 0x1c, 0xa8, 0x1e, 0x52, 0x2e, 0xe6,
	0x10, 0x22, 0xc8, 0x07, 0xe4, 0x5c, 0x48, 0x88, 0x25, 0x47, 0x7e, 0xa3, 0x1d, 0x28, 0x50, 0x41,
	0x06, 0xdc, 0xd0, 0x1b, 0xb9, 0x66, 0xa5, 0xf3, 0x20, 0xd2, 0x49, 0x8d, 0xa8, 0xda, 0xd4, 0xf9,
	0x51, 0x84, 0xcd, 0x38, 0x73, 0x44, 0x46, 0x9f, 0xa8, 0x4b, 0xd0, 0x04, 0x6a, 0xe9, 0x63, 0x86,
	0xb6, 0xa3, 0x42, 0x2b, 0x0e, 0xa0, 0x89, 0x16, 0x95, 0xac, 0x17, 0x5f, 0x7e, 0xff, 0xb9, 0xd6,
	0x5b, 0xd6, 0x53, 0xdb, 0xc5, 0xbe, 0xdf, 0x75, 0x49, 0x20, 0xc8, 0xc8, 0x96, 0x40, 0xb8, 0x7d,
	0x19, 0x4d, 0xea, 0x95, 0x2d, 0x9b, 0xe7, 0xaf, 0xb5, 0x16, 0xba, 0x80, 0x5a, 0xfa, 0xae, 0x89,
	0xa5, 0x57, 0xdc, 0x42, 0xe6, 0xac, 0xc9, 0x24, 0x1f, 0x6b, 0x57, 0xca, 0x3f, 0x47, 0xeb, 0xc9,
	0xa3, 0x33, 0xa8, 0x3a, 0x04, 0x7b, 0x73, 0xca, 0x5b, 0x2b, 0xe8, 0x65, 0xb4, 0xdc, 0x91, 0x9a,
	0x3b, 0xa8, 0xb5, 0x96, 0xa6, 0x7d, 0x49, 0xbd, 0x2b, 0xf4, 0x19, 0x6a, 0xe9, 0x67, 0x20, 0x6e,
	0x7a, 0xc5, 0x03, 0xb1, 0x54, 0xfc, 0x95, 0x14, 0xb7, 0xcd, 0x5b, 0x88, 0x87, 0xd0, 0x2f, 0xa1,
	0x96, 0xbe, 0x60, 0x62, 0xfd, 0x15, 0x57, 0x4f, 0x56, 0xf3, 0xad, 0xdb, 0x34, 0xff, 0x5d, 0x83,
	0xed, 0x1b, 0x5e, 0x67, 0xd4, 0x4e, 0x4e, 0xc0, 0x4d, 0xcf, 0xb8, 0xb9, 0x39, 0x3f, 0x10, 0x09,
	0x34, 0x68, 0x37, 0x61, 0xcd, 0x97, 0x95, 0x78, 0x64, 0x31, 0xf2, 0x35, 0x73, 0xba, 0x6f, 0x7d,
	0x6c, 0xf4, 0xa9, 0x38, 0x1e, 0xf7, 0xda, 0x2e, 0x1b, 0xd8, 0x67, 0xa4, 0x47, 0x05, 0xf1, 0x6d,
	0xf9, 0xbf, 0x82, 0xdb, 0x4a, 0xa4, 0x57, 0x94, 0xcb, 0x97, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x5b, 0xab, 0xd4, 0x01, 0xbb, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AgentSkillServiceClient is the client API for AgentSkillService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AgentSkillServiceClient interface {
	// Create AgentSkill
	CreateAgentSkill(ctx context.Context, in *CreateAgentSkillRequest, opts ...grpc.CallOption) (*AgentSkill, error)
	// List of AgentSkill
	SearchAgentSkill(ctx context.Context, in *SearchAgentSkillRequest, opts ...grpc.CallOption) (*ListAgentSkill, error)
	// AgentSkill item
	ReadAgentSkill(ctx context.Context, in *AgentSkillItemRequest, opts ...grpc.CallOption) (*AgentSkill, error)
	// Update AgentSkill
	UpdateAgentSkill(ctx context.Context, in *UpdateAgentSkillRequest, opts ...grpc.CallOption) (*AgentSkill, error)
	// Remove AgentSkill
	DeleteAgentSkill(ctx context.Context, in *DeleteAgentSkillRequest, opts ...grpc.CallOption) (*AgentSkill, error)
	// SearchLookupAgentNotExistsSkill
	SearchLookupAgentNotExistsSkill(ctx context.Context, in *SearchLookupAgentNotExistsSkillRequest, opts ...grpc.CallOption) (*ListSkill, error)
}

type agentSkillServiceClient struct {
	cc *grpc.ClientConn
}

func NewAgentSkillServiceClient(cc *grpc.ClientConn) AgentSkillServiceClient {
	return &agentSkillServiceClient{cc}
}

func (c *agentSkillServiceClient) CreateAgentSkill(ctx context.Context, in *CreateAgentSkillRequest, opts ...grpc.CallOption) (*AgentSkill, error) {
	out := new(AgentSkill)
	err := c.cc.Invoke(ctx, "/engine.AgentSkillService/CreateAgentSkill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentSkillServiceClient) SearchAgentSkill(ctx context.Context, in *SearchAgentSkillRequest, opts ...grpc.CallOption) (*ListAgentSkill, error) {
	out := new(ListAgentSkill)
	err := c.cc.Invoke(ctx, "/engine.AgentSkillService/SearchAgentSkill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentSkillServiceClient) ReadAgentSkill(ctx context.Context, in *AgentSkillItemRequest, opts ...grpc.CallOption) (*AgentSkill, error) {
	out := new(AgentSkill)
	err := c.cc.Invoke(ctx, "/engine.AgentSkillService/ReadAgentSkill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentSkillServiceClient) UpdateAgentSkill(ctx context.Context, in *UpdateAgentSkillRequest, opts ...grpc.CallOption) (*AgentSkill, error) {
	out := new(AgentSkill)
	err := c.cc.Invoke(ctx, "/engine.AgentSkillService/UpdateAgentSkill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentSkillServiceClient) DeleteAgentSkill(ctx context.Context, in *DeleteAgentSkillRequest, opts ...grpc.CallOption) (*AgentSkill, error) {
	out := new(AgentSkill)
	err := c.cc.Invoke(ctx, "/engine.AgentSkillService/DeleteAgentSkill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentSkillServiceClient) SearchLookupAgentNotExistsSkill(ctx context.Context, in *SearchLookupAgentNotExistsSkillRequest, opts ...grpc.CallOption) (*ListSkill, error) {
	out := new(ListSkill)
	err := c.cc.Invoke(ctx, "/engine.AgentSkillService/SearchLookupAgentNotExistsSkill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentSkillServiceServer is the server API for AgentSkillService service.
type AgentSkillServiceServer interface {
	// Create AgentSkill
	CreateAgentSkill(context.Context, *CreateAgentSkillRequest) (*AgentSkill, error)
	// List of AgentSkill
	SearchAgentSkill(context.Context, *SearchAgentSkillRequest) (*ListAgentSkill, error)
	// AgentSkill item
	ReadAgentSkill(context.Context, *AgentSkillItemRequest) (*AgentSkill, error)
	// Update AgentSkill
	UpdateAgentSkill(context.Context, *UpdateAgentSkillRequest) (*AgentSkill, error)
	// Remove AgentSkill
	DeleteAgentSkill(context.Context, *DeleteAgentSkillRequest) (*AgentSkill, error)
	// SearchLookupAgentNotExistsSkill
	SearchLookupAgentNotExistsSkill(context.Context, *SearchLookupAgentNotExistsSkillRequest) (*ListSkill, error)
}

// UnimplementedAgentSkillServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAgentSkillServiceServer struct {
}

func (*UnimplementedAgentSkillServiceServer) CreateAgentSkill(ctx context.Context, req *CreateAgentSkillRequest) (*AgentSkill, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAgentSkill not implemented")
}
func (*UnimplementedAgentSkillServiceServer) SearchAgentSkill(ctx context.Context, req *SearchAgentSkillRequest) (*ListAgentSkill, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchAgentSkill not implemented")
}
func (*UnimplementedAgentSkillServiceServer) ReadAgentSkill(ctx context.Context, req *AgentSkillItemRequest) (*AgentSkill, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadAgentSkill not implemented")
}
func (*UnimplementedAgentSkillServiceServer) UpdateAgentSkill(ctx context.Context, req *UpdateAgentSkillRequest) (*AgentSkill, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAgentSkill not implemented")
}
func (*UnimplementedAgentSkillServiceServer) DeleteAgentSkill(ctx context.Context, req *DeleteAgentSkillRequest) (*AgentSkill, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAgentSkill not implemented")
}
func (*UnimplementedAgentSkillServiceServer) SearchLookupAgentNotExistsSkill(ctx context.Context, req *SearchLookupAgentNotExistsSkillRequest) (*ListSkill, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchLookupAgentNotExistsSkill not implemented")
}

func RegisterAgentSkillServiceServer(s *grpc.Server, srv AgentSkillServiceServer) {
	s.RegisterService(&_AgentSkillService_serviceDesc, srv)
}

func _AgentSkillService_CreateAgentSkill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAgentSkillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentSkillServiceServer).CreateAgentSkill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.AgentSkillService/CreateAgentSkill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentSkillServiceServer).CreateAgentSkill(ctx, req.(*CreateAgentSkillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentSkillService_SearchAgentSkill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchAgentSkillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentSkillServiceServer).SearchAgentSkill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.AgentSkillService/SearchAgentSkill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentSkillServiceServer).SearchAgentSkill(ctx, req.(*SearchAgentSkillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentSkillService_ReadAgentSkill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentSkillItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentSkillServiceServer).ReadAgentSkill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.AgentSkillService/ReadAgentSkill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentSkillServiceServer).ReadAgentSkill(ctx, req.(*AgentSkillItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentSkillService_UpdateAgentSkill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAgentSkillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentSkillServiceServer).UpdateAgentSkill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.AgentSkillService/UpdateAgentSkill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentSkillServiceServer).UpdateAgentSkill(ctx, req.(*UpdateAgentSkillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentSkillService_DeleteAgentSkill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAgentSkillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentSkillServiceServer).DeleteAgentSkill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.AgentSkillService/DeleteAgentSkill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentSkillServiceServer).DeleteAgentSkill(ctx, req.(*DeleteAgentSkillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentSkillService_SearchLookupAgentNotExistsSkill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchLookupAgentNotExistsSkillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentSkillServiceServer).SearchLookupAgentNotExistsSkill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.AgentSkillService/SearchLookupAgentNotExistsSkill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentSkillServiceServer).SearchLookupAgentNotExistsSkill(ctx, req.(*SearchLookupAgentNotExistsSkillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AgentSkillService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "engine.AgentSkillService",
	HandlerType: (*AgentSkillServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAgentSkill",
			Handler:    _AgentSkillService_CreateAgentSkill_Handler,
		},
		{
			MethodName: "SearchAgentSkill",
			Handler:    _AgentSkillService_SearchAgentSkill_Handler,
		},
		{
			MethodName: "ReadAgentSkill",
			Handler:    _AgentSkillService_ReadAgentSkill_Handler,
		},
		{
			MethodName: "UpdateAgentSkill",
			Handler:    _AgentSkillService_UpdateAgentSkill_Handler,
		},
		{
			MethodName: "DeleteAgentSkill",
			Handler:    _AgentSkillService_DeleteAgentSkill_Handler,
		},
		{
			MethodName: "SearchLookupAgentNotExistsSkill",
			Handler:    _AgentSkillService_SearchLookupAgentNotExistsSkill_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agent_skill.proto",
}