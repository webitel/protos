// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        (unknown)
// source: pause_cause.proto

package api

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateAgentPauseCauseRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Name            string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	LimitMin        uint32                 `protobuf:"varint,2,opt,name=limit_min,json=limitMin,proto3" json:"limit_min,omitempty"`
	AllowSupervisor bool                   `protobuf:"varint,3,opt,name=allow_supervisor,json=allowSupervisor,proto3" json:"allow_supervisor,omitempty"`
	AllowAgent      bool                   `protobuf:"varint,4,opt,name=allow_agent,json=allowAgent,proto3" json:"allow_agent,omitempty"`
	AllowAdmin      bool                   `protobuf:"varint,5,opt,name=allow_admin,json=allowAdmin,proto3" json:"allow_admin,omitempty"`
	Description     string                 `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *CreateAgentPauseCauseRequest) Reset() {
	*x = CreateAgentPauseCauseRequest{}
	mi := &file_pause_cause_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAgentPauseCauseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAgentPauseCauseRequest) ProtoMessage() {}

func (x *CreateAgentPauseCauseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pause_cause_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAgentPauseCauseRequest.ProtoReflect.Descriptor instead.
func (*CreateAgentPauseCauseRequest) Descriptor() ([]byte, []int) {
	return file_pause_cause_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAgentPauseCauseRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateAgentPauseCauseRequest) GetLimitMin() uint32 {
	if x != nil {
		return x.LimitMin
	}
	return 0
}

func (x *CreateAgentPauseCauseRequest) GetAllowSupervisor() bool {
	if x != nil {
		return x.AllowSupervisor
	}
	return false
}

func (x *CreateAgentPauseCauseRequest) GetAllowAgent() bool {
	if x != nil {
		return x.AllowAgent
	}
	return false
}

func (x *CreateAgentPauseCauseRequest) GetAllowAdmin() bool {
	if x != nil {
		return x.AllowAdmin
	}
	return false
}

func (x *CreateAgentPauseCauseRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type AgentPauseCause struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Id              uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt       int64                  `protobuf:"varint,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	CreatedBy       *Lookup                `protobuf:"bytes,3,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	UpdatedAt       int64                  `protobuf:"varint,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	UpdatedBy       *Lookup                `protobuf:"bytes,5,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
	Name            string                 `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	LimitMin        uint32                 `protobuf:"varint,7,opt,name=limit_min,json=limitMin,proto3" json:"limit_min,omitempty"`
	AllowSupervisor bool                   `protobuf:"varint,8,opt,name=allow_supervisor,json=allowSupervisor,proto3" json:"allow_supervisor,omitempty"`
	AllowAgent      bool                   `protobuf:"varint,9,opt,name=allow_agent,json=allowAgent,proto3" json:"allow_agent,omitempty"`
	AllowAdmin      bool                   `protobuf:"varint,10,opt,name=allow_admin,json=allowAdmin,proto3" json:"allow_admin,omitempty"`
	Description     string                 `protobuf:"bytes,11,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *AgentPauseCause) Reset() {
	*x = AgentPauseCause{}
	mi := &file_pause_cause_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AgentPauseCause) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentPauseCause) ProtoMessage() {}

func (x *AgentPauseCause) ProtoReflect() protoreflect.Message {
	mi := &file_pause_cause_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentPauseCause.ProtoReflect.Descriptor instead.
func (*AgentPauseCause) Descriptor() ([]byte, []int) {
	return file_pause_cause_proto_rawDescGZIP(), []int{1}
}

func (x *AgentPauseCause) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AgentPauseCause) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *AgentPauseCause) GetCreatedBy() *Lookup {
	if x != nil {
		return x.CreatedBy
	}
	return nil
}

func (x *AgentPauseCause) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *AgentPauseCause) GetUpdatedBy() *Lookup {
	if x != nil {
		return x.UpdatedBy
	}
	return nil
}

func (x *AgentPauseCause) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AgentPauseCause) GetLimitMin() uint32 {
	if x != nil {
		return x.LimitMin
	}
	return 0
}

func (x *AgentPauseCause) GetAllowSupervisor() bool {
	if x != nil {
		return x.AllowSupervisor
	}
	return false
}

func (x *AgentPauseCause) GetAllowAgent() bool {
	if x != nil {
		return x.AllowAgent
	}
	return false
}

func (x *AgentPauseCause) GetAllowAdmin() bool {
	if x != nil {
		return x.AllowAdmin
	}
	return false
}

func (x *AgentPauseCause) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type SearchAgentPauseCauseRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Page          int32                  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size          int32                  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Q             string                 `protobuf:"bytes,3,opt,name=q,proto3" json:"q,omitempty"`
	Sort          string                 `protobuf:"bytes,4,opt,name=sort,proto3" json:"sort,omitempty"`
	Fields        []string               `protobuf:"bytes,5,rep,name=fields,proto3" json:"fields,omitempty"`
	Id            []uint32               `protobuf:"varint,6,rep,packed,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SearchAgentPauseCauseRequest) Reset() {
	*x = SearchAgentPauseCauseRequest{}
	mi := &file_pause_cause_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchAgentPauseCauseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchAgentPauseCauseRequest) ProtoMessage() {}

func (x *SearchAgentPauseCauseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pause_cause_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchAgentPauseCauseRequest.ProtoReflect.Descriptor instead.
func (*SearchAgentPauseCauseRequest) Descriptor() ([]byte, []int) {
	return file_pause_cause_proto_rawDescGZIP(), []int{2}
}

func (x *SearchAgentPauseCauseRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SearchAgentPauseCauseRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *SearchAgentPauseCauseRequest) GetQ() string {
	if x != nil {
		return x.Q
	}
	return ""
}

func (x *SearchAgentPauseCauseRequest) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

func (x *SearchAgentPauseCauseRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *SearchAgentPauseCauseRequest) GetId() []uint32 {
	if x != nil {
		return x.Id
	}
	return nil
}

type ListAgentPauseCause struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Next          bool                   `protobuf:"varint,1,opt,name=next,proto3" json:"next,omitempty"`
	Items         []*AgentPauseCause     `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAgentPauseCause) Reset() {
	*x = ListAgentPauseCause{}
	mi := &file_pause_cause_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAgentPauseCause) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAgentPauseCause) ProtoMessage() {}

func (x *ListAgentPauseCause) ProtoReflect() protoreflect.Message {
	mi := &file_pause_cause_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAgentPauseCause.ProtoReflect.Descriptor instead.
func (*ListAgentPauseCause) Descriptor() ([]byte, []int) {
	return file_pause_cause_proto_rawDescGZIP(), []int{3}
}

func (x *ListAgentPauseCause) GetNext() bool {
	if x != nil {
		return x.Next
	}
	return false
}

func (x *ListAgentPauseCause) GetItems() []*AgentPauseCause {
	if x != nil {
		return x.Items
	}
	return nil
}

type ReadAgentPauseCauseRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReadAgentPauseCauseRequest) Reset() {
	*x = ReadAgentPauseCauseRequest{}
	mi := &file_pause_cause_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReadAgentPauseCauseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadAgentPauseCauseRequest) ProtoMessage() {}

func (x *ReadAgentPauseCauseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pause_cause_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadAgentPauseCauseRequest.ProtoReflect.Descriptor instead.
func (*ReadAgentPauseCauseRequest) Descriptor() ([]byte, []int) {
	return file_pause_cause_proto_rawDescGZIP(), []int{4}
}

func (x *ReadAgentPauseCauseRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type PatchAgentPauseCauseRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Id              uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Fields          []string               `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty"`
	Name            string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	LimitMin        uint32                 `protobuf:"varint,4,opt,name=limit_min,json=limitMin,proto3" json:"limit_min,omitempty"`
	AllowSupervisor bool                   `protobuf:"varint,5,opt,name=allow_supervisor,json=allowSupervisor,proto3" json:"allow_supervisor,omitempty"`
	AllowAgent      bool                   `protobuf:"varint,6,opt,name=allow_agent,json=allowAgent,proto3" json:"allow_agent,omitempty"`
	AllowAdmin      bool                   `protobuf:"varint,7,opt,name=allow_admin,json=allowAdmin,proto3" json:"allow_admin,omitempty"`
	Description     string                 `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *PatchAgentPauseCauseRequest) Reset() {
	*x = PatchAgentPauseCauseRequest{}
	mi := &file_pause_cause_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PatchAgentPauseCauseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatchAgentPauseCauseRequest) ProtoMessage() {}

func (x *PatchAgentPauseCauseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pause_cause_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatchAgentPauseCauseRequest.ProtoReflect.Descriptor instead.
func (*PatchAgentPauseCauseRequest) Descriptor() ([]byte, []int) {
	return file_pause_cause_proto_rawDescGZIP(), []int{5}
}

func (x *PatchAgentPauseCauseRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PatchAgentPauseCauseRequest) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *PatchAgentPauseCauseRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PatchAgentPauseCauseRequest) GetLimitMin() uint32 {
	if x != nil {
		return x.LimitMin
	}
	return 0
}

func (x *PatchAgentPauseCauseRequest) GetAllowSupervisor() bool {
	if x != nil {
		return x.AllowSupervisor
	}
	return false
}

func (x *PatchAgentPauseCauseRequest) GetAllowAgent() bool {
	if x != nil {
		return x.AllowAgent
	}
	return false
}

func (x *PatchAgentPauseCauseRequest) GetAllowAdmin() bool {
	if x != nil {
		return x.AllowAdmin
	}
	return false
}

func (x *PatchAgentPauseCauseRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type UpdateAgentPauseCauseRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Id              uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name            string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	LimitMin        uint32                 `protobuf:"varint,3,opt,name=limit_min,json=limitMin,proto3" json:"limit_min,omitempty"`
	AllowSupervisor bool                   `protobuf:"varint,4,opt,name=allow_supervisor,json=allowSupervisor,proto3" json:"allow_supervisor,omitempty"`
	AllowAgent      bool                   `protobuf:"varint,5,opt,name=allow_agent,json=allowAgent,proto3" json:"allow_agent,omitempty"`
	AllowAdmin      bool                   `protobuf:"varint,6,opt,name=allow_admin,json=allowAdmin,proto3" json:"allow_admin,omitempty"`
	Description     string                 `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *UpdateAgentPauseCauseRequest) Reset() {
	*x = UpdateAgentPauseCauseRequest{}
	mi := &file_pause_cause_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateAgentPauseCauseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAgentPauseCauseRequest) ProtoMessage() {}

func (x *UpdateAgentPauseCauseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pause_cause_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAgentPauseCauseRequest.ProtoReflect.Descriptor instead.
func (*UpdateAgentPauseCauseRequest) Descriptor() ([]byte, []int) {
	return file_pause_cause_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateAgentPauseCauseRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateAgentPauseCauseRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateAgentPauseCauseRequest) GetLimitMin() uint32 {
	if x != nil {
		return x.LimitMin
	}
	return 0
}

func (x *UpdateAgentPauseCauseRequest) GetAllowSupervisor() bool {
	if x != nil {
		return x.AllowSupervisor
	}
	return false
}

func (x *UpdateAgentPauseCauseRequest) GetAllowAgent() bool {
	if x != nil {
		return x.AllowAgent
	}
	return false
}

func (x *UpdateAgentPauseCauseRequest) GetAllowAdmin() bool {
	if x != nil {
		return x.AllowAdmin
	}
	return false
}

func (x *UpdateAgentPauseCauseRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type DeleteAgentPauseCauseRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteAgentPauseCauseRequest) Reset() {
	*x = DeleteAgentPauseCauseRequest{}
	mi := &file_pause_cause_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAgentPauseCauseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAgentPauseCauseRequest) ProtoMessage() {}

func (x *DeleteAgentPauseCauseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pause_cause_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAgentPauseCauseRequest.ProtoReflect.Descriptor instead.
func (*DeleteAgentPauseCauseRequest) Descriptor() ([]byte, []int) {
	return file_pause_cause_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteAgentPauseCauseRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_pause_cause_proto protoreflect.FileDescriptor

var file_pause_cause_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x61, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x75, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xde, 0x01, 0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x75, 0x73, 0x65, 0x43, 0x61, 0x75, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x5f, 0x6d, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x4d, 0x69, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x5f, 0x73, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x53, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69,
	0x73, 0x6f, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xfd, 0x02, 0x0a, 0x0f, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x50, 0x61, 0x75, 0x73, 0x65, 0x43, 0x61, 0x75, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x5f, 0x6d, 0x69, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x4d, 0x69, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x5f, 0x73, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x53, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69,
	0x73, 0x6f, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x90, 0x01, 0x0a, 0x1c, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x75, 0x73, 0x65, 0x43, 0x61, 0x75, 0x73,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x12, 0x0c, 0x0a, 0x01, 0x71, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x71, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6f,
	0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x58, 0x0a, 0x13, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x75, 0x73, 0x65, 0x43, 0x61, 0x75, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x04, 0x6e, 0x65, 0x78, 0x74, 0x12, 0x2d, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x50, 0x61, 0x75, 0x73, 0x65, 0x43, 0x61, 0x75, 0x73, 0x65, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x22, 0x2c, 0x0a, 0x1a, 0x52, 0x65, 0x61, 0x64, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x50, 0x61, 0x75, 0x73, 0x65, 0x43, 0x61, 0x75, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x85, 0x02, 0x0a, 0x1b, 0x50, 0x61, 0x74, 0x63, 0x68, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x50, 0x61, 0x75, 0x73, 0x65, 0x43, 0x61, 0x75, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x6d, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x4d, 0x69, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x53, 0x75, 0x70, 0x65,
	0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x61, 0x6c, 0x6c,
	0x6f, 0x77, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xee, 0x01, 0x0a, 0x1c, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x75, 0x73, 0x65, 0x43,
	0x61, 0x75, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x6d, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x4d, 0x69, 0x6e, 0x12, 0x29, 0x0a, 0x10,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x53, 0x75, 0x70,
	0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x6c, 0x6c, 0x6f, 0x77,
	0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2e, 0x0a, 0x1c, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x75, 0x73, 0x65, 0x43,
	0x61, 0x75, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x32, 0x96, 0x06, 0x0a, 0x16,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x75, 0x73, 0x65, 0x43, 0x61, 0x75, 0x73, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7c, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x75, 0x73, 0x65, 0x43, 0x61, 0x75, 0x73, 0x65, 0x12,
	0x24, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x75, 0x73, 0x65, 0x43, 0x61, 0x75, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x75, 0x73, 0x65, 0x43, 0x61, 0x75, 0x73, 0x65, 0x22, 0x24,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01, 0x2a, 0x22, 0x19, 0x2f, 0x63, 0x61, 0x6c, 0x6c,
	0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x61, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61,
	0x75, 0x73, 0x65, 0x73, 0x12, 0x7d, 0x0a, 0x15, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x50, 0x61, 0x75, 0x73, 0x65, 0x43, 0x61, 0x75, 0x73, 0x65, 0x12, 0x24, 0x2e,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x50, 0x61, 0x75, 0x73, 0x65, 0x43, 0x61, 0x75, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x75, 0x73, 0x65, 0x43, 0x61, 0x75, 0x73, 0x65,
	0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x63, 0x61, 0x6c, 0x6c, 0x5f,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x61, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x75,
	0x73, 0x65, 0x73, 0x12, 0x7a, 0x0a, 0x13, 0x52, 0x65, 0x61, 0x64, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x50, 0x61, 0x75, 0x73, 0x65, 0x43, 0x61, 0x75, 0x73, 0x65, 0x12, 0x22, 0x2e, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x75,
	0x73, 0x65, 0x43, 0x61, 0x75, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x75,
	0x73, 0x65, 0x43, 0x61, 0x75, 0x73, 0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x12,
	0x1e, 0x2f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x61,
	0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x75, 0x73, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12,
	0x7f, 0x0a, 0x14, 0x50, 0x61, 0x74, 0x63, 0x68, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x75,
	0x73, 0x65, 0x43, 0x61, 0x75, 0x73, 0x65, 0x12, 0x23, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x50, 0x61, 0x74, 0x63, 0x68, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x75, 0x73, 0x65,
	0x43, 0x61, 0x75, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x75, 0x73, 0x65,
	0x43, 0x61, 0x75, 0x73, 0x65, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x3a, 0x01, 0x2a,
	0x32, 0x1e, 0x2f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x70,
	0x61, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x75, 0x73, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x12, 0x81, 0x01, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x50, 0x61, 0x75, 0x73, 0x65, 0x43, 0x61, 0x75, 0x73, 0x65, 0x12, 0x24, 0x2e, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x50,
	0x61, 0x75, 0x73, 0x65, 0x43, 0x61, 0x75, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x50,
	0x61, 0x75, 0x73, 0x65, 0x43, 0x61, 0x75, 0x73, 0x65, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x23, 0x3a, 0x01, 0x2a, 0x1a, 0x1e, 0x2f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2f, 0x70, 0x61, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x75, 0x73, 0x65, 0x73, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x12, 0x7e, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x50, 0x61, 0x75, 0x73, 0x65, 0x43, 0x61, 0x75, 0x73, 0x65, 0x12, 0x24, 0x2e,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x50, 0x61, 0x75, 0x73, 0x65, 0x43, 0x61, 0x75, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x50, 0x61, 0x75, 0x73, 0x65, 0x43, 0x61, 0x75, 0x73, 0x65, 0x22, 0x26, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x20, 0x2a, 0x1e, 0x2f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2f, 0x70, 0x61, 0x75, 0x73, 0x65, 0x5f, 0x63, 0x61, 0x75, 0x73, 0x65, 0x73, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x42, 0x74, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x42, 0x0f, 0x50, 0x61, 0x75, 0x73, 0x65, 0x43, 0x61, 0x75, 0x73, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x77, 0x65, 0x62, 0x69, 0x74, 0x65, 0x6c, 0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0xa2, 0x02, 0x03, 0x45, 0x58, 0x58, 0xaa, 0x02, 0x06, 0x45, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0xca, 0x02, 0x06, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0xe2, 0x02, 0x12, 0x45,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x06, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pause_cause_proto_rawDescOnce sync.Once
	file_pause_cause_proto_rawDescData = file_pause_cause_proto_rawDesc
)

func file_pause_cause_proto_rawDescGZIP() []byte {
	file_pause_cause_proto_rawDescOnce.Do(func() {
		file_pause_cause_proto_rawDescData = protoimpl.X.CompressGZIP(file_pause_cause_proto_rawDescData)
	})
	return file_pause_cause_proto_rawDescData
}

var file_pause_cause_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pause_cause_proto_goTypes = []any{
	(*CreateAgentPauseCauseRequest)(nil), // 0: engine.CreateAgentPauseCauseRequest
	(*AgentPauseCause)(nil),              // 1: engine.AgentPauseCause
	(*SearchAgentPauseCauseRequest)(nil), // 2: engine.SearchAgentPauseCauseRequest
	(*ListAgentPauseCause)(nil),          // 3: engine.ListAgentPauseCause
	(*ReadAgentPauseCauseRequest)(nil),   // 4: engine.ReadAgentPauseCauseRequest
	(*PatchAgentPauseCauseRequest)(nil),  // 5: engine.PatchAgentPauseCauseRequest
	(*UpdateAgentPauseCauseRequest)(nil), // 6: engine.UpdateAgentPauseCauseRequest
	(*DeleteAgentPauseCauseRequest)(nil), // 7: engine.DeleteAgentPauseCauseRequest
	(*Lookup)(nil),                       // 8: engine.Lookup
}
var file_pause_cause_proto_depIdxs = []int32{
	8, // 0: engine.AgentPauseCause.created_by:type_name -> engine.Lookup
	8, // 1: engine.AgentPauseCause.updated_by:type_name -> engine.Lookup
	1, // 2: engine.ListAgentPauseCause.items:type_name -> engine.AgentPauseCause
	0, // 3: engine.AgentPauseCauseService.CreateAgentPauseCause:input_type -> engine.CreateAgentPauseCauseRequest
	2, // 4: engine.AgentPauseCauseService.SearchAgentPauseCause:input_type -> engine.SearchAgentPauseCauseRequest
	4, // 5: engine.AgentPauseCauseService.ReadAgentPauseCause:input_type -> engine.ReadAgentPauseCauseRequest
	5, // 6: engine.AgentPauseCauseService.PatchAgentPauseCause:input_type -> engine.PatchAgentPauseCauseRequest
	6, // 7: engine.AgentPauseCauseService.UpdateAgentPauseCause:input_type -> engine.UpdateAgentPauseCauseRequest
	7, // 8: engine.AgentPauseCauseService.DeleteAgentPauseCause:input_type -> engine.DeleteAgentPauseCauseRequest
	1, // 9: engine.AgentPauseCauseService.CreateAgentPauseCause:output_type -> engine.AgentPauseCause
	3, // 10: engine.AgentPauseCauseService.SearchAgentPauseCause:output_type -> engine.ListAgentPauseCause
	1, // 11: engine.AgentPauseCauseService.ReadAgentPauseCause:output_type -> engine.AgentPauseCause
	1, // 12: engine.AgentPauseCauseService.PatchAgentPauseCause:output_type -> engine.AgentPauseCause
	1, // 13: engine.AgentPauseCauseService.UpdateAgentPauseCause:output_type -> engine.AgentPauseCause
	1, // 14: engine.AgentPauseCauseService.DeleteAgentPauseCause:output_type -> engine.AgentPauseCause
	9, // [9:15] is the sub-list for method output_type
	3, // [3:9] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pause_cause_proto_init() }
func file_pause_cause_proto_init() {
	if File_pause_cause_proto != nil {
		return
	}
	file_const_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pause_cause_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pause_cause_proto_goTypes,
		DependencyIndexes: file_pause_cause_proto_depIdxs,
		MessageInfos:      file_pause_cause_proto_msgTypes,
	}.Build()
	File_pause_cause_proto = out.File
	file_pause_cause_proto_rawDesc = nil
	file_pause_cause_proto_goTypes = nil
	file_pause_cause_proto_depIdxs = nil
}