// Code generated by protoc-gen-go. DO NOT EDIT.
// source: agent_team.proto

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

type DeleteAgentTeamRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DomainId             int64    `protobuf:"varint,2,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteAgentTeamRequest) Reset()         { *m = DeleteAgentTeamRequest{} }
func (m *DeleteAgentTeamRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteAgentTeamRequest) ProtoMessage()    {}
func (*DeleteAgentTeamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3083856eea12172, []int{0}
}

func (m *DeleteAgentTeamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteAgentTeamRequest.Unmarshal(m, b)
}
func (m *DeleteAgentTeamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteAgentTeamRequest.Marshal(b, m, deterministic)
}
func (m *DeleteAgentTeamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteAgentTeamRequest.Merge(m, src)
}
func (m *DeleteAgentTeamRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteAgentTeamRequest.Size(m)
}
func (m *DeleteAgentTeamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteAgentTeamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteAgentTeamRequest proto.InternalMessageInfo

func (m *DeleteAgentTeamRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DeleteAgentTeamRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

type UpdateAgentTeamRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Strategy             string   `protobuf:"bytes,4,opt,name=strategy,proto3" json:"strategy,omitempty"`
	MaxNoAnswer          int32    `protobuf:"varint,5,opt,name=max_no_answer,json=maxNoAnswer,proto3" json:"max_no_answer,omitempty"`
	NoAnswerDelayTime    int32    `protobuf:"varint,6,opt,name=no_answer_delay_time,json=noAnswerDelayTime,proto3" json:"no_answer_delay_time,omitempty"`
	WrapUpTime           int32    `protobuf:"varint,7,opt,name=wrap_up_time,json=wrapUpTime,proto3" json:"wrap_up_time,omitempty"`
	CallTimeout          int32    `protobuf:"varint,8,opt,name=call_timeout,json=callTimeout,proto3" json:"call_timeout,omitempty"`
	Admin                *Lookup  `protobuf:"bytes,9,opt,name=admin,proto3" json:"admin,omitempty"`
	DomainId             int64    `protobuf:"varint,10,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateAgentTeamRequest) Reset()         { *m = UpdateAgentTeamRequest{} }
func (m *UpdateAgentTeamRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateAgentTeamRequest) ProtoMessage()    {}
func (*UpdateAgentTeamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3083856eea12172, []int{1}
}

func (m *UpdateAgentTeamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateAgentTeamRequest.Unmarshal(m, b)
}
func (m *UpdateAgentTeamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateAgentTeamRequest.Marshal(b, m, deterministic)
}
func (m *UpdateAgentTeamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateAgentTeamRequest.Merge(m, src)
}
func (m *UpdateAgentTeamRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateAgentTeamRequest.Size(m)
}
func (m *UpdateAgentTeamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateAgentTeamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateAgentTeamRequest proto.InternalMessageInfo

func (m *UpdateAgentTeamRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateAgentTeamRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateAgentTeamRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *UpdateAgentTeamRequest) GetStrategy() string {
	if m != nil {
		return m.Strategy
	}
	return ""
}

func (m *UpdateAgentTeamRequest) GetMaxNoAnswer() int32 {
	if m != nil {
		return m.MaxNoAnswer
	}
	return 0
}

func (m *UpdateAgentTeamRequest) GetNoAnswerDelayTime() int32 {
	if m != nil {
		return m.NoAnswerDelayTime
	}
	return 0
}

func (m *UpdateAgentTeamRequest) GetWrapUpTime() int32 {
	if m != nil {
		return m.WrapUpTime
	}
	return 0
}

func (m *UpdateAgentTeamRequest) GetCallTimeout() int32 {
	if m != nil {
		return m.CallTimeout
	}
	return 0
}

func (m *UpdateAgentTeamRequest) GetAdmin() *Lookup {
	if m != nil {
		return m.Admin
	}
	return nil
}

func (m *UpdateAgentTeamRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

type ReadAgentTeamRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DomainId             int64    `protobuf:"varint,2,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadAgentTeamRequest) Reset()         { *m = ReadAgentTeamRequest{} }
func (m *ReadAgentTeamRequest) String() string { return proto.CompactTextString(m) }
func (*ReadAgentTeamRequest) ProtoMessage()    {}
func (*ReadAgentTeamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3083856eea12172, []int{2}
}

func (m *ReadAgentTeamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadAgentTeamRequest.Unmarshal(m, b)
}
func (m *ReadAgentTeamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadAgentTeamRequest.Marshal(b, m, deterministic)
}
func (m *ReadAgentTeamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadAgentTeamRequest.Merge(m, src)
}
func (m *ReadAgentTeamRequest) XXX_Size() int {
	return xxx_messageInfo_ReadAgentTeamRequest.Size(m)
}
func (m *ReadAgentTeamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadAgentTeamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadAgentTeamRequest proto.InternalMessageInfo

func (m *ReadAgentTeamRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ReadAgentTeamRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

type SearchAgentTeamRequest struct {
	Page                 int32    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size                 int32    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Q                    string   `protobuf:"bytes,3,opt,name=q,proto3" json:"q,omitempty"`
	DomainId             int64    `protobuf:"varint,4,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	Fields               []string `protobuf:"bytes,5,rep,name=fields,proto3" json:"fields,omitempty"`
	Sort                 string   `protobuf:"bytes,6,opt,name=sort,proto3" json:"sort,omitempty"`
	Id                   []uint32 `protobuf:"varint,7,rep,packed,name=id,proto3" json:"id,omitempty"`
	Strategy             []string `protobuf:"bytes,8,rep,name=strategy,proto3" json:"strategy,omitempty"`
	AdminId              []uint32 `protobuf:"varint,9,rep,packed,name=admin_id,json=adminId,proto3" json:"admin_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchAgentTeamRequest) Reset()         { *m = SearchAgentTeamRequest{} }
func (m *SearchAgentTeamRequest) String() string { return proto.CompactTextString(m) }
func (*SearchAgentTeamRequest) ProtoMessage()    {}
func (*SearchAgentTeamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3083856eea12172, []int{3}
}

func (m *SearchAgentTeamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchAgentTeamRequest.Unmarshal(m, b)
}
func (m *SearchAgentTeamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchAgentTeamRequest.Marshal(b, m, deterministic)
}
func (m *SearchAgentTeamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchAgentTeamRequest.Merge(m, src)
}
func (m *SearchAgentTeamRequest) XXX_Size() int {
	return xxx_messageInfo_SearchAgentTeamRequest.Size(m)
}
func (m *SearchAgentTeamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchAgentTeamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchAgentTeamRequest proto.InternalMessageInfo

func (m *SearchAgentTeamRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *SearchAgentTeamRequest) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *SearchAgentTeamRequest) GetQ() string {
	if m != nil {
		return m.Q
	}
	return ""
}

func (m *SearchAgentTeamRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

func (m *SearchAgentTeamRequest) GetFields() []string {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *SearchAgentTeamRequest) GetSort() string {
	if m != nil {
		return m.Sort
	}
	return ""
}

func (m *SearchAgentTeamRequest) GetId() []uint32 {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *SearchAgentTeamRequest) GetStrategy() []string {
	if m != nil {
		return m.Strategy
	}
	return nil
}

func (m *SearchAgentTeamRequest) GetAdminId() []uint32 {
	if m != nil {
		return m.AdminId
	}
	return nil
}

type CreateAgentTeamRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Strategy             string   `protobuf:"bytes,3,opt,name=strategy,proto3" json:"strategy,omitempty"`
	MaxNoAnswer          int32    `protobuf:"varint,4,opt,name=max_no_answer,json=maxNoAnswer,proto3" json:"max_no_answer,omitempty"`
	NoAnswerDelayTime    int32    `protobuf:"varint,5,opt,name=no_answer_delay_time,json=noAnswerDelayTime,proto3" json:"no_answer_delay_time,omitempty"`
	WrapUpTime           int32    `protobuf:"varint,6,opt,name=wrap_up_time,json=wrapUpTime,proto3" json:"wrap_up_time,omitempty"`
	CallTimeout          int32    `protobuf:"varint,7,opt,name=call_timeout,json=callTimeout,proto3" json:"call_timeout,omitempty"`
	Admin                *Lookup  `protobuf:"bytes,8,opt,name=admin,proto3" json:"admin,omitempty"`
	DomainId             int64    `protobuf:"varint,9,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateAgentTeamRequest) Reset()         { *m = CreateAgentTeamRequest{} }
func (m *CreateAgentTeamRequest) String() string { return proto.CompactTextString(m) }
func (*CreateAgentTeamRequest) ProtoMessage()    {}
func (*CreateAgentTeamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3083856eea12172, []int{4}
}

func (m *CreateAgentTeamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAgentTeamRequest.Unmarshal(m, b)
}
func (m *CreateAgentTeamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAgentTeamRequest.Marshal(b, m, deterministic)
}
func (m *CreateAgentTeamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAgentTeamRequest.Merge(m, src)
}
func (m *CreateAgentTeamRequest) XXX_Size() int {
	return xxx_messageInfo_CreateAgentTeamRequest.Size(m)
}
func (m *CreateAgentTeamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAgentTeamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAgentTeamRequest proto.InternalMessageInfo

func (m *CreateAgentTeamRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateAgentTeamRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreateAgentTeamRequest) GetStrategy() string {
	if m != nil {
		return m.Strategy
	}
	return ""
}

func (m *CreateAgentTeamRequest) GetMaxNoAnswer() int32 {
	if m != nil {
		return m.MaxNoAnswer
	}
	return 0
}

func (m *CreateAgentTeamRequest) GetNoAnswerDelayTime() int32 {
	if m != nil {
		return m.NoAnswerDelayTime
	}
	return 0
}

func (m *CreateAgentTeamRequest) GetWrapUpTime() int32 {
	if m != nil {
		return m.WrapUpTime
	}
	return 0
}

func (m *CreateAgentTeamRequest) GetCallTimeout() int32 {
	if m != nil {
		return m.CallTimeout
	}
	return 0
}

func (m *CreateAgentTeamRequest) GetAdmin() *Lookup {
	if m != nil {
		return m.Admin
	}
	return nil
}

func (m *CreateAgentTeamRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

type AgentTeam struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DomainId             int64    `protobuf:"varint,2,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Strategy             string   `protobuf:"bytes,5,opt,name=strategy,proto3" json:"strategy,omitempty"`
	MaxNoAnswer          int32    `protobuf:"varint,6,opt,name=max_no_answer,json=maxNoAnswer,proto3" json:"max_no_answer,omitempty"`
	NoAnswerDelayTime    int32    `protobuf:"varint,7,opt,name=no_answer_delay_time,json=noAnswerDelayTime,proto3" json:"no_answer_delay_time,omitempty"`
	WrapUpTime           int32    `protobuf:"varint,8,opt,name=wrap_up_time,json=wrapUpTime,proto3" json:"wrap_up_time,omitempty"`
	CallTimeout          int32    `protobuf:"varint,9,opt,name=call_timeout,json=callTimeout,proto3" json:"call_timeout,omitempty"`
	UpdatedAt            int64    `protobuf:"varint,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Admin                *Lookup  `protobuf:"bytes,11,opt,name=admin,proto3" json:"admin,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgentTeam) Reset()         { *m = AgentTeam{} }
func (m *AgentTeam) String() string { return proto.CompactTextString(m) }
func (*AgentTeam) ProtoMessage()    {}
func (*AgentTeam) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3083856eea12172, []int{5}
}

func (m *AgentTeam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgentTeam.Unmarshal(m, b)
}
func (m *AgentTeam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgentTeam.Marshal(b, m, deterministic)
}
func (m *AgentTeam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentTeam.Merge(m, src)
}
func (m *AgentTeam) XXX_Size() int {
	return xxx_messageInfo_AgentTeam.Size(m)
}
func (m *AgentTeam) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentTeam.DiscardUnknown(m)
}

var xxx_messageInfo_AgentTeam proto.InternalMessageInfo

func (m *AgentTeam) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AgentTeam) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

func (m *AgentTeam) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AgentTeam) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *AgentTeam) GetStrategy() string {
	if m != nil {
		return m.Strategy
	}
	return ""
}

func (m *AgentTeam) GetMaxNoAnswer() int32 {
	if m != nil {
		return m.MaxNoAnswer
	}
	return 0
}

func (m *AgentTeam) GetNoAnswerDelayTime() int32 {
	if m != nil {
		return m.NoAnswerDelayTime
	}
	return 0
}

func (m *AgentTeam) GetWrapUpTime() int32 {
	if m != nil {
		return m.WrapUpTime
	}
	return 0
}

func (m *AgentTeam) GetCallTimeout() int32 {
	if m != nil {
		return m.CallTimeout
	}
	return 0
}

func (m *AgentTeam) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *AgentTeam) GetAdmin() *Lookup {
	if m != nil {
		return m.Admin
	}
	return nil
}

type ListAgentTeam struct {
	Next                 bool         `protobuf:"varint,1,opt,name=next,proto3" json:"next,omitempty"`
	Items                []*AgentTeam `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ListAgentTeam) Reset()         { *m = ListAgentTeam{} }
func (m *ListAgentTeam) String() string { return proto.CompactTextString(m) }
func (*ListAgentTeam) ProtoMessage()    {}
func (*ListAgentTeam) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3083856eea12172, []int{6}
}

func (m *ListAgentTeam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAgentTeam.Unmarshal(m, b)
}
func (m *ListAgentTeam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAgentTeam.Marshal(b, m, deterministic)
}
func (m *ListAgentTeam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAgentTeam.Merge(m, src)
}
func (m *ListAgentTeam) XXX_Size() int {
	return xxx_messageInfo_ListAgentTeam.Size(m)
}
func (m *ListAgentTeam) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAgentTeam.DiscardUnknown(m)
}

var xxx_messageInfo_ListAgentTeam proto.InternalMessageInfo

func (m *ListAgentTeam) GetNext() bool {
	if m != nil {
		return m.Next
	}
	return false
}

func (m *ListAgentTeam) GetItems() []*AgentTeam {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*DeleteAgentTeamRequest)(nil), "engine.DeleteAgentTeamRequest")
	proto.RegisterType((*UpdateAgentTeamRequest)(nil), "engine.UpdateAgentTeamRequest")
	proto.RegisterType((*ReadAgentTeamRequest)(nil), "engine.ReadAgentTeamRequest")
	proto.RegisterType((*SearchAgentTeamRequest)(nil), "engine.SearchAgentTeamRequest")
	proto.RegisterType((*CreateAgentTeamRequest)(nil), "engine.CreateAgentTeamRequest")
	proto.RegisterType((*AgentTeam)(nil), "engine.AgentTeam")
	proto.RegisterType((*ListAgentTeam)(nil), "engine.ListAgentTeam")
}

func init() { proto.RegisterFile("agent_team.proto", fileDescriptor_e3083856eea12172) }

var fileDescriptor_e3083856eea12172 = []byte{
	// 747 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xbd, 0x6e, 0xe3, 0x46,
	0x10, 0x06, 0x29, 0x51, 0x22, 0x47, 0x96, 0x65, 0x2f, 0x1c, 0x99, 0x56, 0xec, 0x84, 0x21, 0x02,
	0x44, 0x70, 0x21, 0x02, 0x4e, 0x97, 0xce, 0xb1, 0x53, 0x18, 0x30, 0x52, 0xd0, 0x76, 0x93, 0x86,
	0x58, 0x93, 0x13, 0x79, 0x11, 0xf1, 0xc7, 0xe4, 0x2a, 0xb6, 0x13, 0x5c, 0x73, 0x6f, 0x70, 0xb8,
	0x87, 0xba, 0x07, 0xb8, 0xe2, 0xba, 0xab, 0xae, 0xb9, 0xb7, 0x38, 0x70, 0xa8, 0x7f, 0x51, 0x3f,
	0x87, 0xeb, 0x76, 0x67, 0x86, 0xdf, 0xc7, 0xfd, 0xe6, 0xdb, 0x1d, 0xd8, 0xe3, 0x7d, 0x8c, 0xa4,
	0x27, 0x91, 0x87, 0xbd, 0x24, 0x8d, 0x65, 0xcc, 0x6a, 0x18, 0xf5, 0x45, 0x84, 0x9d, 0xe3, 0x7e,
	0x1c, 0xf7, 0x07, 0xe8, 0xf0, 0x44, 0x38, 0x3c, 0x8a, 0x62, 0xc9, 0xa5, 0x88, 0xa3, 0xac, 0xa8,
	0xea, 0x34, 0xfc, 0x38, 0xca, 0x64, 0xb1, 0xb1, 0xff, 0x80, 0xf6, 0x25, 0x0e, 0x50, 0xe2, 0x79,
	0x0e, 0x76, 0x8b, 0x3c, 0x74, 0xf1, 0x71, 0x88, 0x99, 0x64, 0xbb, 0xa0, 0x8a, 0xc0, 0x54, 0x2c,
	0xa5, 0x5b, 0x71, 0x55, 0x11, 0xb0, 0xef, 0xc1, 0x08, 0xe2, 0x90, 0x8b, 0xc8, 0x13, 0x81, 0xa9,
	0x52, 0x58, 0x2f, 0x02, 0x57, 0x81, 0xfd, 0x41, 0x85, 0xf6, 0x5d, 0x12, 0xf0, 0x2d, 0x70, 0x18,
	0x54, 0x23, 0x1e, 0x22, 0x41, 0x18, 0x2e, 0xad, 0x99, 0x05, 0x8d, 0x00, 0x33, 0x3f, 0x15, 0x49,
	0xfe, 0xa3, 0x66, 0x85, 0x52, 0xb3, 0x21, 0xd6, 0x01, 0x3d, 0x93, 0x29, 0x97, 0xd8, 0x7f, 0x31,
	0xab, 0x94, 0x9e, 0xec, 0x99, 0x0d, 0xcd, 0x90, 0x3f, 0x7b, 0x51, 0xec, 0xf1, 0x28, 0x7b, 0xc2,
	0xd4, 0xd4, 0x2c, 0xa5, 0xab, 0xb9, 0x8d, 0x90, 0x3f, 0xff, 0x19, 0x9f, 0x53, 0x88, 0x39, 0x70,
	0x30, 0xc9, 0x7b, 0x01, 0x0e, 0xf8, 0x8b, 0x27, 0x45, 0x88, 0x66, 0x8d, 0x4a, 0xf7, 0xa3, 0x51,
	0xdd, 0x65, 0x9e, 0xb9, 0x15, 0xf4, 0x4b, 0x3b, 0x4f, 0x29, 0x4f, 0xbc, 0x61, 0x52, 0x14, 0xd6,
	0xa9, 0x10, 0xf2, 0xd8, 0x5d, 0x42, 0x15, 0x3f, 0xc1, 0x8e, 0xcf, 0x07, 0x03, 0x4a, 0xc7, 0x43,
	0x69, 0xea, 0x05, 0x6b, 0x1e, 0xbb, 0x2d, 0x42, 0xec, 0x67, 0xd0, 0x78, 0x10, 0x8a, 0xc8, 0x34,
	0x2c, 0xa5, 0xdb, 0x38, 0xdb, 0xed, 0x15, 0x0d, 0xea, 0x5d, 0xc7, 0xf1, 0x3f, 0xc3, 0xc4, 0x2d,
	0x92, 0xf3, 0xca, 0xc2, 0x82, 0xb2, 0x17, 0x70, 0xe0, 0x22, 0x0f, 0xbe, 0xad, 0x3d, 0x1f, 0x15,
	0x68, 0xdf, 0x20, 0x4f, 0xfd, 0x87, 0x25, 0x1c, 0x06, 0xd5, 0x84, 0xf7, 0x91, 0x90, 0x34, 0x97,
	0xd6, 0x79, 0x2c, 0x13, 0xff, 0x15, 0x2d, 0xd2, 0x5c, 0x5a, 0xb3, 0x1d, 0x50, 0x1e, 0x47, 0x8d,
	0x51, 0x1e, 0xe7, 0xd9, 0xaa, 0xf3, 0x6c, 0xac, 0x0d, 0xb5, 0xbf, 0x05, 0x0e, 0x82, 0xcc, 0xd4,
	0xac, 0x4a, 0xd7, 0x70, 0x47, 0x3b, 0x82, 0x8d, 0x53, 0x49, 0x9a, 0x1b, 0x2e, 0xad, 0x47, 0xc7,
	0xa8, 0x5b, 0x95, 0x6e, 0x93, 0x8e, 0x31, 0xdb, 0x67, 0x9d, 0xbe, 0x9e, 0xf6, 0xf9, 0x08, 0x74,
	0x12, 0x2c, 0xe7, 0x34, 0xe8, 0x8b, 0x3a, 0xed, 0xaf, 0x02, 0xfb, 0x9d, 0x0a, 0xed, 0x8b, 0x14,
	0xcb, 0xfc, 0x37, 0xf6, 0x9b, 0xb2, 0xda, 0x6f, 0xea, 0x7a, 0xbf, 0x55, 0x36, 0xf9, 0xad, 0xba,
	0xbd, 0xdf, 0xb4, 0x6d, 0xfd, 0x56, 0xdb, 0xe8, 0xb7, 0xfa, 0x1a, 0xbf, 0xe9, 0x5b, 0xfb, 0xcd,
	0x58, 0xb0, 0xca, 0x67, 0x15, 0x8c, 0x89, 0x86, 0x5f, 0xe5, 0xb2, 0x89, 0xd2, 0x95, 0xd5, 0x4a,
	0x57, 0xd7, 0x2b, 0xad, 0x6d, 0x52, 0xba, 0xb6, 0xbd, 0xd2, 0xf5, 0x6d, 0x95, 0xd6, 0x37, 0x2a,
	0x6d, 0x2c, 0x2b, 0x7d, 0x02, 0x30, 0xa4, 0xf7, 0x2e, 0xf0, 0xb8, 0x1c, 0x5d, 0x5a, 0x63, 0x14,
	0x39, 0x9f, 0x69, 0x44, 0x63, 0x4d, 0x23, 0xec, 0x6b, 0x68, 0x5e, 0x8b, 0x4c, 0x4e, 0xe5, 0xce,
	0x15, 0xc4, 0x67, 0x49, 0x82, 0xeb, 0x2e, 0xad, 0xd9, 0x2f, 0xa0, 0x09, 0x89, 0x61, 0x66, 0xaa,
	0x56, 0xa5, 0xdb, 0x38, 0xdb, 0x1f, 0x43, 0x4d, 0x8d, 0x5e, 0xe4, 0xcf, 0xde, 0x54, 0x61, 0x6f,
	0x12, 0xbc, 0xc1, 0xf4, 0x5f, 0xe1, 0x23, 0xf3, 0xa1, 0xb5, 0x70, 0x2f, 0xd8, 0x0f, 0x63, 0x84,
	0xf2, 0x0b, 0xd3, 0x59, 0x66, 0xb0, 0x4f, 0x5e, 0xbf, 0xff, 0xf4, 0x56, 0x3d, 0xb4, 0x99, 0x43,
	0xba, 0xf8, 0x18, 0x49, 0x4c, 0x9d, 0x7c, 0xf2, 0x64, 0xbf, 0x29, 0xa7, 0x2c, 0x80, 0xd6, 0xc2,
	0xeb, 0x32, 0x25, 0x29, 0x7f, 0x76, 0x3a, 0xdf, 0x4d, 0x14, 0x99, 0x15, 0xc0, 0xee, 0x10, 0xd1,
	0x01, 0x2b, 0x21, 0x62, 0x1c, 0x9a, 0x73, 0x2f, 0x21, 0x3b, 0x1e, 0x63, 0x94, 0x3d, 0x90, 0x65,
	0xc7, 0xf8, 0x91, 0xd0, 0x8f, 0xd8, 0xe1, 0x32, 0xba, 0xf3, 0xbf, 0x08, 0x5e, 0xb1, 0x07, 0x68,
	0x2d, 0x4c, 0xb1, 0xe9, 0x41, 0xca, 0xc7, 0x5b, 0x19, 0x8d, 0x4d, 0x34, 0xc7, 0x9d, 0x55, 0x34,
	0xb9, 0x64, 0x08, 0xad, 0x85, 0xb9, 0x3b, 0x65, 0x2a, 0x1f, 0xc8, 0x6b, 0x0e, 0x74, 0xba, 0x8a,
	0xe9, 0x77, 0xfb, 0x2f, 0xab, 0x2f, 0xe4, 0xc3, 0xf0, 0xbe, 0xe7, 0xc7, 0xa1, 0xf3, 0x84, 0xf7,
	0x42, 0xe2, 0xc0, 0xa1, 0xd1, 0x9f, 0x39, 0x05, 0xdc, 0x7d, 0x8d, 0xb6, 0xbf, 0x7e, 0x09, 0x00,
	0x00, 0xff, 0xff, 0xa2, 0x93, 0xac, 0xfd, 0x50, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AgentTeamServiceClient is the client API for AgentTeamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AgentTeamServiceClient interface {
	// Create AgentTeam
	CreateAgentTeam(ctx context.Context, in *CreateAgentTeamRequest, opts ...grpc.CallOption) (*AgentTeam, error)
	// List of AgentTeam
	SearchAgentTeam(ctx context.Context, in *SearchAgentTeamRequest, opts ...grpc.CallOption) (*ListAgentTeam, error)
	// AgentTeam item
	ReadAgentTeam(ctx context.Context, in *ReadAgentTeamRequest, opts ...grpc.CallOption) (*AgentTeam, error)
	// Update AgentTeam
	UpdateAgentTeam(ctx context.Context, in *UpdateAgentTeamRequest, opts ...grpc.CallOption) (*AgentTeam, error)
	// Remove AgentTeam
	DeleteAgentTeam(ctx context.Context, in *DeleteAgentTeamRequest, opts ...grpc.CallOption) (*AgentTeam, error)
}

type agentTeamServiceClient struct {
	cc *grpc.ClientConn
}

func NewAgentTeamServiceClient(cc *grpc.ClientConn) AgentTeamServiceClient {
	return &agentTeamServiceClient{cc}
}

func (c *agentTeamServiceClient) CreateAgentTeam(ctx context.Context, in *CreateAgentTeamRequest, opts ...grpc.CallOption) (*AgentTeam, error) {
	out := new(AgentTeam)
	err := c.cc.Invoke(ctx, "/engine.AgentTeamService/CreateAgentTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentTeamServiceClient) SearchAgentTeam(ctx context.Context, in *SearchAgentTeamRequest, opts ...grpc.CallOption) (*ListAgentTeam, error) {
	out := new(ListAgentTeam)
	err := c.cc.Invoke(ctx, "/engine.AgentTeamService/SearchAgentTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentTeamServiceClient) ReadAgentTeam(ctx context.Context, in *ReadAgentTeamRequest, opts ...grpc.CallOption) (*AgentTeam, error) {
	out := new(AgentTeam)
	err := c.cc.Invoke(ctx, "/engine.AgentTeamService/ReadAgentTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentTeamServiceClient) UpdateAgentTeam(ctx context.Context, in *UpdateAgentTeamRequest, opts ...grpc.CallOption) (*AgentTeam, error) {
	out := new(AgentTeam)
	err := c.cc.Invoke(ctx, "/engine.AgentTeamService/UpdateAgentTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentTeamServiceClient) DeleteAgentTeam(ctx context.Context, in *DeleteAgentTeamRequest, opts ...grpc.CallOption) (*AgentTeam, error) {
	out := new(AgentTeam)
	err := c.cc.Invoke(ctx, "/engine.AgentTeamService/DeleteAgentTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentTeamServiceServer is the server API for AgentTeamService service.
type AgentTeamServiceServer interface {
	// Create AgentTeam
	CreateAgentTeam(context.Context, *CreateAgentTeamRequest) (*AgentTeam, error)
	// List of AgentTeam
	SearchAgentTeam(context.Context, *SearchAgentTeamRequest) (*ListAgentTeam, error)
	// AgentTeam item
	ReadAgentTeam(context.Context, *ReadAgentTeamRequest) (*AgentTeam, error)
	// Update AgentTeam
	UpdateAgentTeam(context.Context, *UpdateAgentTeamRequest) (*AgentTeam, error)
	// Remove AgentTeam
	DeleteAgentTeam(context.Context, *DeleteAgentTeamRequest) (*AgentTeam, error)
}

// UnimplementedAgentTeamServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAgentTeamServiceServer struct {
}

func (*UnimplementedAgentTeamServiceServer) CreateAgentTeam(ctx context.Context, req *CreateAgentTeamRequest) (*AgentTeam, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAgentTeam not implemented")
}
func (*UnimplementedAgentTeamServiceServer) SearchAgentTeam(ctx context.Context, req *SearchAgentTeamRequest) (*ListAgentTeam, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchAgentTeam not implemented")
}
func (*UnimplementedAgentTeamServiceServer) ReadAgentTeam(ctx context.Context, req *ReadAgentTeamRequest) (*AgentTeam, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadAgentTeam not implemented")
}
func (*UnimplementedAgentTeamServiceServer) UpdateAgentTeam(ctx context.Context, req *UpdateAgentTeamRequest) (*AgentTeam, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAgentTeam not implemented")
}
func (*UnimplementedAgentTeamServiceServer) DeleteAgentTeam(ctx context.Context, req *DeleteAgentTeamRequest) (*AgentTeam, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAgentTeam not implemented")
}

func RegisterAgentTeamServiceServer(s *grpc.Server, srv AgentTeamServiceServer) {
	s.RegisterService(&_AgentTeamService_serviceDesc, srv)
}

func _AgentTeamService_CreateAgentTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAgentTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentTeamServiceServer).CreateAgentTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.AgentTeamService/CreateAgentTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentTeamServiceServer).CreateAgentTeam(ctx, req.(*CreateAgentTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentTeamService_SearchAgentTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchAgentTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentTeamServiceServer).SearchAgentTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.AgentTeamService/SearchAgentTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentTeamServiceServer).SearchAgentTeam(ctx, req.(*SearchAgentTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentTeamService_ReadAgentTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadAgentTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentTeamServiceServer).ReadAgentTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.AgentTeamService/ReadAgentTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentTeamServiceServer).ReadAgentTeam(ctx, req.(*ReadAgentTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentTeamService_UpdateAgentTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAgentTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentTeamServiceServer).UpdateAgentTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.AgentTeamService/UpdateAgentTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentTeamServiceServer).UpdateAgentTeam(ctx, req.(*UpdateAgentTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentTeamService_DeleteAgentTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAgentTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentTeamServiceServer).DeleteAgentTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.AgentTeamService/DeleteAgentTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentTeamServiceServer).DeleteAgentTeam(ctx, req.(*DeleteAgentTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AgentTeamService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "engine.AgentTeamService",
	HandlerType: (*AgentTeamServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAgentTeam",
			Handler:    _AgentTeamService_CreateAgentTeam_Handler,
		},
		{
			MethodName: "SearchAgentTeam",
			Handler:    _AgentTeamService_SearchAgentTeam_Handler,
		},
		{
			MethodName: "ReadAgentTeam",
			Handler:    _AgentTeamService_ReadAgentTeam_Handler,
		},
		{
			MethodName: "UpdateAgentTeam",
			Handler:    _AgentTeamService_UpdateAgentTeam_Handler,
		},
		{
			MethodName: "DeleteAgentTeam",
			Handler:    _AgentTeamService_DeleteAgentTeam_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agent_team.proto",
}
