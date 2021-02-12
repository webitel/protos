// Code generated by protoc-gen-go. DO NOT EDIT.
// source: queue_skill.proto

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

type CreateQueueSkillRequest struct {
	QueueId              uint32    `protobuf:"varint,1,opt,name=queue_id,json=queueId,proto3" json:"queue_id,omitempty"`
	Skill                *Lookup   `protobuf:"bytes,2,opt,name=skill,proto3" json:"skill,omitempty"`
	Buckets              []*Lookup `protobuf:"bytes,3,rep,name=buckets,proto3" json:"buckets,omitempty"`
	Lvl                  int32     `protobuf:"varint,4,opt,name=lvl,proto3" json:"lvl,omitempty"`
	MinCapacity          int32     `protobuf:"varint,5,opt,name=min_capacity,json=minCapacity,proto3" json:"min_capacity,omitempty"`
	MaxCapacity          int32     `protobuf:"varint,6,opt,name=max_capacity,json=maxCapacity,proto3" json:"max_capacity,omitempty"`
	Disabled             bool      `protobuf:"varint,7,opt,name=disabled,proto3" json:"disabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CreateQueueSkillRequest) Reset()         { *m = CreateQueueSkillRequest{} }
func (m *CreateQueueSkillRequest) String() string { return proto.CompactTextString(m) }
func (*CreateQueueSkillRequest) ProtoMessage()    {}
func (*CreateQueueSkillRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ad92fc63d7cf65f, []int{0}
}

func (m *CreateQueueSkillRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateQueueSkillRequest.Unmarshal(m, b)
}
func (m *CreateQueueSkillRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateQueueSkillRequest.Marshal(b, m, deterministic)
}
func (m *CreateQueueSkillRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateQueueSkillRequest.Merge(m, src)
}
func (m *CreateQueueSkillRequest) XXX_Size() int {
	return xxx_messageInfo_CreateQueueSkillRequest.Size(m)
}
func (m *CreateQueueSkillRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateQueueSkillRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateQueueSkillRequest proto.InternalMessageInfo

func (m *CreateQueueSkillRequest) GetQueueId() uint32 {
	if m != nil {
		return m.QueueId
	}
	return 0
}

func (m *CreateQueueSkillRequest) GetSkill() *Lookup {
	if m != nil {
		return m.Skill
	}
	return nil
}

func (m *CreateQueueSkillRequest) GetBuckets() []*Lookup {
	if m != nil {
		return m.Buckets
	}
	return nil
}

func (m *CreateQueueSkillRequest) GetLvl() int32 {
	if m != nil {
		return m.Lvl
	}
	return 0
}

func (m *CreateQueueSkillRequest) GetMinCapacity() int32 {
	if m != nil {
		return m.MinCapacity
	}
	return 0
}

func (m *CreateQueueSkillRequest) GetMaxCapacity() int32 {
	if m != nil {
		return m.MaxCapacity
	}
	return 0
}

func (m *CreateQueueSkillRequest) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

type QueueSkill struct {
	Id                   uint32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Skill                *Lookup   `protobuf:"bytes,2,opt,name=skill,proto3" json:"skill,omitempty"`
	Buckets              []*Lookup `protobuf:"bytes,3,rep,name=buckets,proto3" json:"buckets,omitempty"`
	Lvl                  int32     `protobuf:"varint,4,opt,name=lvl,proto3" json:"lvl,omitempty"`
	MinCapacity          int32     `protobuf:"varint,5,opt,name=min_capacity,json=minCapacity,proto3" json:"min_capacity,omitempty"`
	MaxCapacity          int32     `protobuf:"varint,6,opt,name=max_capacity,json=maxCapacity,proto3" json:"max_capacity,omitempty"`
	Disabled             bool      `protobuf:"varint,7,opt,name=disabled,proto3" json:"disabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *QueueSkill) Reset()         { *m = QueueSkill{} }
func (m *QueueSkill) String() string { return proto.CompactTextString(m) }
func (*QueueSkill) ProtoMessage()    {}
func (*QueueSkill) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ad92fc63d7cf65f, []int{1}
}

func (m *QueueSkill) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueueSkill.Unmarshal(m, b)
}
func (m *QueueSkill) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueueSkill.Marshal(b, m, deterministic)
}
func (m *QueueSkill) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueueSkill.Merge(m, src)
}
func (m *QueueSkill) XXX_Size() int {
	return xxx_messageInfo_QueueSkill.Size(m)
}
func (m *QueueSkill) XXX_DiscardUnknown() {
	xxx_messageInfo_QueueSkill.DiscardUnknown(m)
}

var xxx_messageInfo_QueueSkill proto.InternalMessageInfo

func (m *QueueSkill) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *QueueSkill) GetSkill() *Lookup {
	if m != nil {
		return m.Skill
	}
	return nil
}

func (m *QueueSkill) GetBuckets() []*Lookup {
	if m != nil {
		return m.Buckets
	}
	return nil
}

func (m *QueueSkill) GetLvl() int32 {
	if m != nil {
		return m.Lvl
	}
	return 0
}

func (m *QueueSkill) GetMinCapacity() int32 {
	if m != nil {
		return m.MinCapacity
	}
	return 0
}

func (m *QueueSkill) GetMaxCapacity() int32 {
	if m != nil {
		return m.MaxCapacity
	}
	return 0
}

func (m *QueueSkill) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

type UpdateQueueSkillRequest struct {
	QueueId              uint32    `protobuf:"varint,1,opt,name=queue_id,json=queueId,proto3" json:"queue_id,omitempty"`
	Id                   uint32    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Skill                *Lookup   `protobuf:"bytes,3,opt,name=skill,proto3" json:"skill,omitempty"`
	Buckets              []*Lookup `protobuf:"bytes,4,rep,name=buckets,proto3" json:"buckets,omitempty"`
	Lvl                  int32     `protobuf:"varint,5,opt,name=lvl,proto3" json:"lvl,omitempty"`
	MinCapacity          int32     `protobuf:"varint,6,opt,name=min_capacity,json=minCapacity,proto3" json:"min_capacity,omitempty"`
	MaxCapacity          int32     `protobuf:"varint,7,opt,name=max_capacity,json=maxCapacity,proto3" json:"max_capacity,omitempty"`
	Disabled             bool      `protobuf:"varint,8,opt,name=disabled,proto3" json:"disabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UpdateQueueSkillRequest) Reset()         { *m = UpdateQueueSkillRequest{} }
func (m *UpdateQueueSkillRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateQueueSkillRequest) ProtoMessage()    {}
func (*UpdateQueueSkillRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ad92fc63d7cf65f, []int{2}
}

func (m *UpdateQueueSkillRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateQueueSkillRequest.Unmarshal(m, b)
}
func (m *UpdateQueueSkillRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateQueueSkillRequest.Marshal(b, m, deterministic)
}
func (m *UpdateQueueSkillRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateQueueSkillRequest.Merge(m, src)
}
func (m *UpdateQueueSkillRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateQueueSkillRequest.Size(m)
}
func (m *UpdateQueueSkillRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateQueueSkillRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateQueueSkillRequest proto.InternalMessageInfo

func (m *UpdateQueueSkillRequest) GetQueueId() uint32 {
	if m != nil {
		return m.QueueId
	}
	return 0
}

func (m *UpdateQueueSkillRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateQueueSkillRequest) GetSkill() *Lookup {
	if m != nil {
		return m.Skill
	}
	return nil
}

func (m *UpdateQueueSkillRequest) GetBuckets() []*Lookup {
	if m != nil {
		return m.Buckets
	}
	return nil
}

func (m *UpdateQueueSkillRequest) GetLvl() int32 {
	if m != nil {
		return m.Lvl
	}
	return 0
}

func (m *UpdateQueueSkillRequest) GetMinCapacity() int32 {
	if m != nil {
		return m.MinCapacity
	}
	return 0
}

func (m *UpdateQueueSkillRequest) GetMaxCapacity() int32 {
	if m != nil {
		return m.MaxCapacity
	}
	return 0
}

func (m *UpdateQueueSkillRequest) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

type PatchQueueSkillRequest struct {
	QueueId              uint32    `protobuf:"varint,1,opt,name=queue_id,json=queueId,proto3" json:"queue_id,omitempty"`
	Id                   uint32    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Skill                *Lookup   `protobuf:"bytes,3,opt,name=skill,proto3" json:"skill,omitempty"`
	Buckets              []*Lookup `protobuf:"bytes,4,rep,name=buckets,proto3" json:"buckets,omitempty"`
	Lvl                  int32     `protobuf:"varint,5,opt,name=lvl,proto3" json:"lvl,omitempty"`
	MinCapacity          int32     `protobuf:"varint,6,opt,name=min_capacity,json=minCapacity,proto3" json:"min_capacity,omitempty"`
	MaxCapacity          int32     `protobuf:"varint,7,opt,name=max_capacity,json=maxCapacity,proto3" json:"max_capacity,omitempty"`
	Disabled             bool      `protobuf:"varint,8,opt,name=disabled,proto3" json:"disabled,omitempty"`
	Fields               []string  `protobuf:"bytes,9,rep,name=fields,proto3" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PatchQueueSkillRequest) Reset()         { *m = PatchQueueSkillRequest{} }
func (m *PatchQueueSkillRequest) String() string { return proto.CompactTextString(m) }
func (*PatchQueueSkillRequest) ProtoMessage()    {}
func (*PatchQueueSkillRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ad92fc63d7cf65f, []int{3}
}

func (m *PatchQueueSkillRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PatchQueueSkillRequest.Unmarshal(m, b)
}
func (m *PatchQueueSkillRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PatchQueueSkillRequest.Marshal(b, m, deterministic)
}
func (m *PatchQueueSkillRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PatchQueueSkillRequest.Merge(m, src)
}
func (m *PatchQueueSkillRequest) XXX_Size() int {
	return xxx_messageInfo_PatchQueueSkillRequest.Size(m)
}
func (m *PatchQueueSkillRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PatchQueueSkillRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PatchQueueSkillRequest proto.InternalMessageInfo

func (m *PatchQueueSkillRequest) GetQueueId() uint32 {
	if m != nil {
		return m.QueueId
	}
	return 0
}

func (m *PatchQueueSkillRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PatchQueueSkillRequest) GetSkill() *Lookup {
	if m != nil {
		return m.Skill
	}
	return nil
}

func (m *PatchQueueSkillRequest) GetBuckets() []*Lookup {
	if m != nil {
		return m.Buckets
	}
	return nil
}

func (m *PatchQueueSkillRequest) GetLvl() int32 {
	if m != nil {
		return m.Lvl
	}
	return 0
}

func (m *PatchQueueSkillRequest) GetMinCapacity() int32 {
	if m != nil {
		return m.MinCapacity
	}
	return 0
}

func (m *PatchQueueSkillRequest) GetMaxCapacity() int32 {
	if m != nil {
		return m.MaxCapacity
	}
	return 0
}

func (m *PatchQueueSkillRequest) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func (m *PatchQueueSkillRequest) GetFields() []string {
	if m != nil {
		return m.Fields
	}
	return nil
}

type SearchQueueSkillRequest struct {
	QueueId              uint32   `protobuf:"varint,1,opt,name=queue_id,json=queueId,proto3" json:"queue_id,omitempty"`
	Page                 int32    `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Size                 int32    `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Q                    string   `protobuf:"bytes,4,opt,name=q,proto3" json:"q,omitempty"`
	Id                   []uint32 `protobuf:"varint,6,rep,packed,name=id,proto3" json:"id,omitempty"`
	SkillId              []uint32 `protobuf:"varint,7,rep,packed,name=skill_id,json=skillId,proto3" json:"skill_id,omitempty"`
	BucketId             []uint32 `protobuf:"varint,8,rep,packed,name=bucket_id,json=bucketId,proto3" json:"bucket_id,omitempty"`
	Lvl                  []int32  `protobuf:"varint,9,rep,packed,name=lvl,proto3" json:"lvl,omitempty"`
	MinCapacity          []int32  `protobuf:"varint,10,rep,packed,name=min_capacity,json=minCapacity,proto3" json:"min_capacity,omitempty"`
	MaxCapacity          []int32  `protobuf:"varint,11,rep,packed,name=max_capacity,json=maxCapacity,proto3" json:"max_capacity,omitempty"`
	Disabled             bool     `protobuf:"varint,12,opt,name=disabled,proto3" json:"disabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchQueueSkillRequest) Reset()         { *m = SearchQueueSkillRequest{} }
func (m *SearchQueueSkillRequest) String() string { return proto.CompactTextString(m) }
func (*SearchQueueSkillRequest) ProtoMessage()    {}
func (*SearchQueueSkillRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ad92fc63d7cf65f, []int{4}
}

func (m *SearchQueueSkillRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchQueueSkillRequest.Unmarshal(m, b)
}
func (m *SearchQueueSkillRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchQueueSkillRequest.Marshal(b, m, deterministic)
}
func (m *SearchQueueSkillRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchQueueSkillRequest.Merge(m, src)
}
func (m *SearchQueueSkillRequest) XXX_Size() int {
	return xxx_messageInfo_SearchQueueSkillRequest.Size(m)
}
func (m *SearchQueueSkillRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchQueueSkillRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchQueueSkillRequest proto.InternalMessageInfo

func (m *SearchQueueSkillRequest) GetQueueId() uint32 {
	if m != nil {
		return m.QueueId
	}
	return 0
}

func (m *SearchQueueSkillRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *SearchQueueSkillRequest) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *SearchQueueSkillRequest) GetQ() string {
	if m != nil {
		return m.Q
	}
	return ""
}

func (m *SearchQueueSkillRequest) GetId() []uint32 {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *SearchQueueSkillRequest) GetSkillId() []uint32 {
	if m != nil {
		return m.SkillId
	}
	return nil
}

func (m *SearchQueueSkillRequest) GetBucketId() []uint32 {
	if m != nil {
		return m.BucketId
	}
	return nil
}

func (m *SearchQueueSkillRequest) GetLvl() []int32 {
	if m != nil {
		return m.Lvl
	}
	return nil
}

func (m *SearchQueueSkillRequest) GetMinCapacity() []int32 {
	if m != nil {
		return m.MinCapacity
	}
	return nil
}

func (m *SearchQueueSkillRequest) GetMaxCapacity() []int32 {
	if m != nil {
		return m.MaxCapacity
	}
	return nil
}

func (m *SearchQueueSkillRequest) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

type DeleteQueueSkillRequest struct {
	QueueId              uint32   `protobuf:"varint,1,opt,name=queue_id,json=queueId,proto3" json:"queue_id,omitempty"`
	Id                   uint32   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteQueueSkillRequest) Reset()         { *m = DeleteQueueSkillRequest{} }
func (m *DeleteQueueSkillRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteQueueSkillRequest) ProtoMessage()    {}
func (*DeleteQueueSkillRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ad92fc63d7cf65f, []int{5}
}

func (m *DeleteQueueSkillRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteQueueSkillRequest.Unmarshal(m, b)
}
func (m *DeleteQueueSkillRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteQueueSkillRequest.Marshal(b, m, deterministic)
}
func (m *DeleteQueueSkillRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteQueueSkillRequest.Merge(m, src)
}
func (m *DeleteQueueSkillRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteQueueSkillRequest.Size(m)
}
func (m *DeleteQueueSkillRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteQueueSkillRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteQueueSkillRequest proto.InternalMessageInfo

func (m *DeleteQueueSkillRequest) GetQueueId() uint32 {
	if m != nil {
		return m.QueueId
	}
	return 0
}

func (m *DeleteQueueSkillRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ListQueueSkill struct {
	Next                 bool          `protobuf:"varint,1,opt,name=next,proto3" json:"next,omitempty"`
	Items                []*QueueSkill `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ListQueueSkill) Reset()         { *m = ListQueueSkill{} }
func (m *ListQueueSkill) String() string { return proto.CompactTextString(m) }
func (*ListQueueSkill) ProtoMessage()    {}
func (*ListQueueSkill) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ad92fc63d7cf65f, []int{6}
}

func (m *ListQueueSkill) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListQueueSkill.Unmarshal(m, b)
}
func (m *ListQueueSkill) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListQueueSkill.Marshal(b, m, deterministic)
}
func (m *ListQueueSkill) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListQueueSkill.Merge(m, src)
}
func (m *ListQueueSkill) XXX_Size() int {
	return xxx_messageInfo_ListQueueSkill.Size(m)
}
func (m *ListQueueSkill) XXX_DiscardUnknown() {
	xxx_messageInfo_ListQueueSkill.DiscardUnknown(m)
}

var xxx_messageInfo_ListQueueSkill proto.InternalMessageInfo

func (m *ListQueueSkill) GetNext() bool {
	if m != nil {
		return m.Next
	}
	return false
}

func (m *ListQueueSkill) GetItems() []*QueueSkill {
	if m != nil {
		return m.Items
	}
	return nil
}

type ReadQueueSkillRequest struct {
	QueueId              uint32   `protobuf:"varint,1,opt,name=queue_id,json=queueId,proto3" json:"queue_id,omitempty"`
	Id                   uint32   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadQueueSkillRequest) Reset()         { *m = ReadQueueSkillRequest{} }
func (m *ReadQueueSkillRequest) String() string { return proto.CompactTextString(m) }
func (*ReadQueueSkillRequest) ProtoMessage()    {}
func (*ReadQueueSkillRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ad92fc63d7cf65f, []int{7}
}

func (m *ReadQueueSkillRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadQueueSkillRequest.Unmarshal(m, b)
}
func (m *ReadQueueSkillRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadQueueSkillRequest.Marshal(b, m, deterministic)
}
func (m *ReadQueueSkillRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadQueueSkillRequest.Merge(m, src)
}
func (m *ReadQueueSkillRequest) XXX_Size() int {
	return xxx_messageInfo_ReadQueueSkillRequest.Size(m)
}
func (m *ReadQueueSkillRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadQueueSkillRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadQueueSkillRequest proto.InternalMessageInfo

func (m *ReadQueueSkillRequest) GetQueueId() uint32 {
	if m != nil {
		return m.QueueId
	}
	return 0
}

func (m *ReadQueueSkillRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateQueueSkillRequest)(nil), "engine.CreateQueueSkillRequest")
	proto.RegisterType((*QueueSkill)(nil), "engine.QueueSkill")
	proto.RegisterType((*UpdateQueueSkillRequest)(nil), "engine.UpdateQueueSkillRequest")
	proto.RegisterType((*PatchQueueSkillRequest)(nil), "engine.PatchQueueSkillRequest")
	proto.RegisterType((*SearchQueueSkillRequest)(nil), "engine.SearchQueueSkillRequest")
	proto.RegisterType((*DeleteQueueSkillRequest)(nil), "engine.DeleteQueueSkillRequest")
	proto.RegisterType((*ListQueueSkill)(nil), "engine.ListQueueSkill")
	proto.RegisterType((*ReadQueueSkillRequest)(nil), "engine.ReadQueueSkillRequest")
}

func init() { proto.RegisterFile("queue_skill.proto", fileDescriptor_9ad92fc63d7cf65f) }

var fileDescriptor_9ad92fc63d7cf65f = []byte{
	// 689 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x96, 0xcf, 0x6f, 0xd3, 0x4a,
	0x10, 0xc7, 0x65, 0x27, 0x4e, 0x9c, 0x49, 0x5f, 0x5e, 0xba, 0xd2, 0x6b, 0xdc, 0xbc, 0xf7, 0xc0,
	0x58, 0x20, 0xac, 0x08, 0x62, 0x14, 0xc4, 0x85, 0x63, 0xdb, 0x4b, 0xa5, 0x0a, 0x81, 0x2b, 0x2e,
	0x5c, 0xaa, 0x8d, 0x3d, 0xa4, 0xab, 0x3a, 0x76, 0x92, 0xdd, 0xf4, 0x27, 0xe5, 0xc0, 0x89, 0x3b,
	0x7f, 0x00, 0x37, 0xfe, 0x21, 0xfe, 0x01, 0x0e, 0xfc, 0x11, 0x9c, 0x10, 0xf2, 0x3a, 0x4e, 0x5a,
	0x13, 0x87, 0x44, 0xbd, 0xc1, 0x6d, 0x67, 0x67, 0xb2, 0xdf, 0x99, 0x8f, 0x67, 0x67, 0x03, 0xeb,
	0xc3, 0x31, 0x8e, 0xf1, 0x80, 0x1f, 0xb1, 0x20, 0x68, 0x0f, 0x46, 0x91, 0x88, 0x48, 0x09, 0xc3,
	0x1e, 0x0b, 0xb1, 0xf9, 0x5f, 0x2f, 0x8a, 0x7a, 0x01, 0x3a, 0x74, 0xc0, 0x1c, 0x1a, 0x86, 0x91,
	0xa0, 0x82, 0x45, 0x21, 0x4f, 0xa2, 0x9a, 0x55, 0x2f, 0x0a, 0xb9, 0x48, 0x0c, 0xeb, 0x9b, 0x02,
	0x8d, 0xed, 0x11, 0x52, 0x81, 0x2f, 0xe2, 0xe3, 0xf6, 0xe3, 0xd3, 0x5c, 0x1c, 0x8e, 0x91, 0x0b,
	0xb2, 0x09, 0x7a, 0xa2, 0xc1, 0x7c, 0x43, 0x31, 0x15, 0xfb, 0x2f, 0xb7, 0x2c, 0xed, 0x5d, 0x9f,
	0xdc, 0x05, 0x4d, 0x0a, 0x1b, 0xaa, 0xa9, 0xd8, 0xd5, 0x4e, 0xad, 0x9d, 0x28, 0xb7, 0xf7, 0xa2,
	0xe8, 0x68, 0x3c, 0x70, 0x13, 0x27, 0xb1, 0xa1, 0xdc, 0x1d, 0x7b, 0x47, 0x28, 0xb8, 0x51, 0x30,
	0x0b, 0x73, 0xe2, 0x52, 0x37, 0xa9, 0x43, 0x21, 0x38, 0x0e, 0x8c, 0xa2, 0xa9, 0xd8, 0x9a, 0x1b,
	0x2f, 0xc9, 0x1d, 0x58, 0xeb, 0xb3, 0xf0, 0xc0, 0xa3, 0x03, 0xea, 0x31, 0x71, 0x66, 0x68, 0xd2,
	0x55, 0xed, 0xb3, 0x70, 0x7b, 0xb2, 0x25, 0x43, 0xe8, 0xe9, 0x2c, 0xa4, 0x34, 0x09, 0xa1, 0xa7,
	0xd3, 0x90, 0x26, 0xe8, 0x3e, 0xe3, 0xb4, 0x1b, 0xa0, 0x6f, 0x94, 0x4d, 0xc5, 0xd6, 0xdd, 0xa9,
	0x6d, 0x7d, 0x51, 0x00, 0x66, 0x45, 0x93, 0x1a, 0xa8, 0xd3, 0x3a, 0x55, 0xf6, 0x7b, 0x94, 0xf8,
	0x5e, 0x85, 0xc6, 0xcb, 0x81, 0xbf, 0xea, 0xd7, 0x4d, 0x50, 0xa8, 0x3f, 0xa3, 0x28, 0x2c, 0x89,
	0xa2, 0xb8, 0x14, 0x0a, 0x2d, 0x1f, 0x45, 0xe9, 0xd7, 0x28, 0xca, 0x8b, 0x51, 0xe8, 0x19, 0x14,
	0x1f, 0x55, 0xd8, 0x78, 0x4e, 0x85, 0x77, 0xf8, 0xa7, 0x93, 0x20, 0x1b, 0x50, 0x7a, 0xcd, 0x30,
	0xf0, 0xb9, 0x51, 0x31, 0x0b, 0x76, 0xc5, 0x9d, 0x58, 0xd6, 0x27, 0x15, 0x1a, 0xfb, 0x48, 0x47,
	0x2b, 0x22, 0x22, 0x50, 0x1c, 0xd0, 0x1e, 0x4a, 0x48, 0x9a, 0x2b, 0xd7, 0xf1, 0x1e, 0x67, 0xe7,
	0x28, 0x29, 0x69, 0xae, 0x5c, 0x93, 0x35, 0x50, 0x86, 0xb2, 0xfb, 0x2b, 0xae, 0x32, 0x9c, 0x80,
	0x2d, 0x99, 0x85, 0x09, 0xd8, 0x4d, 0xd0, 0x25, 0xbb, 0x58, 0xa0, 0x2c, 0x77, 0xcb, 0xd2, 0xde,
	0xf5, 0xc9, 0xbf, 0x50, 0x49, 0x70, 0xc5, 0x3e, 0x5d, 0xfa, 0xf4, 0x64, 0x63, 0xd7, 0x4f, 0x01,
	0xc6, 0x95, 0xe4, 0x00, 0x04, 0xe9, 0x5a, 0x08, 0xb0, 0x3a, 0x09, 0xc9, 0x01, 0xb8, 0x96, 0x69,
	0xa5, 0x1d, 0x68, 0xec, 0x60, 0x80, 0x37, 0xbb, 0x54, 0xd6, 0x33, 0xa8, 0xed, 0x31, 0x2e, 0xae,
	0x4c, 0x20, 0x02, 0xc5, 0x10, 0x4f, 0x85, 0xfc, 0xa1, 0xee, 0xca, 0x35, 0xb1, 0x41, 0x63, 0x02,
	0xfb, 0xdc, 0x50, 0x65, 0x23, 0x91, 0xb4, 0x91, 0xae, 0x48, 0x27, 0x01, 0xd6, 0x16, 0xfc, 0xe3,
	0x22, 0xf5, 0x6f, 0x92, 0x53, 0xe7, 0xbb, 0x06, 0xeb, 0xb3, 0x03, 0xf6, 0x71, 0x74, 0xcc, 0x3c,
	0x24, 0x67, 0x50, 0xcf, 0x3e, 0x11, 0xe4, 0x76, 0x9a, 0x48, 0xce, 0xe3, 0xd1, 0x9c, 0x93, 0xa9,
	0xf5, 0xe8, 0xdd, 0xe7, 0xaf, 0x1f, 0xd4, 0x96, 0x75, 0xcf, 0xf1, 0x68, 0x10, 0x1c, 0x78, 0x18,
	0x0a, 0x1c, 0x39, 0x32, 0x19, 0xee, 0x5c, 0xa4, 0x49, 0x5e, 0x3a, 0xf2, 0xcb, 0xf3, 0xa7, 0x4a,
	0x8b, 0x9c, 0x43, 0x3d, 0xdb, 0x92, 0x33, 0xe9, 0x9c, 0x66, 0x6d, 0x6e, 0x4c, 0x6f, 0xdb, 0x35,
	0xbe, 0xd6, 0x43, 0x29, 0x7f, 0x9f, 0x2c, 0x27, 0x4f, 0x4e, 0xa0, 0x76, 0x1d, 0x28, 0xf9, 0x3f,
	0x3d, 0x78, 0x2e, 0xe8, 0xb9, 0x25, 0x77, 0xa4, 0xe6, 0x03, 0xd2, 0x5a, 0x4a, 0xd3, 0xb9, 0x60,
	0xfe, 0x25, 0x79, 0x0b, 0xf5, 0xec, 0xd0, 0x9e, 0x15, 0x9d, 0x33, 0xce, 0xe7, 0x8a, 0x3f, 0x91,
	0xe2, 0x4e, 0x73, 0x05, 0xf1, 0x18, 0xfa, 0x1b, 0xf8, 0x3b, 0x33, 0x29, 0xc9, 0xad, 0xf4, 0xf4,
	0xf9, 0x23, 0x74, 0x91, 0x7a, 0x67, 0x45, 0xf5, 0x0b, 0xa8, 0x67, 0x6f, 0xd7, 0xac, 0xfa, 0x9c,
	0x7b, 0xb7, 0x08, 0x7d, 0x6b, 0x05, 0xfd, 0x2d, 0xeb, 0x95, 0xd9, 0x63, 0xe2, 0x70, 0xdc, 0x6d,
	0x7b, 0x51, 0xdf, 0x39, 0xc1, 0x2e, 0x13, 0x18, 0x38, 0xf2, 0xaf, 0x12, 0x77, 0x12, 0x89, 0x6e,
	0x49, 0x9a, 0x8f, 0x7f, 0x04, 0x00, 0x00, 0xff, 0xff, 0x11, 0x5f, 0x31, 0xbd, 0x81, 0x09, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueueSkillServiceClient is the client API for QueueSkillService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueueSkillServiceClient interface {
	CreateQueueSkill(ctx context.Context, in *CreateQueueSkillRequest, opts ...grpc.CallOption) (*QueueSkill, error)
	SearchQueueSkill(ctx context.Context, in *SearchQueueSkillRequest, opts ...grpc.CallOption) (*ListQueueSkill, error)
	ReadQueueSkill(ctx context.Context, in *ReadQueueSkillRequest, opts ...grpc.CallOption) (*QueueSkill, error)
	UpdateQueueSkill(ctx context.Context, in *UpdateQueueSkillRequest, opts ...grpc.CallOption) (*QueueSkill, error)
	PatchQueueSkill(ctx context.Context, in *PatchQueueSkillRequest, opts ...grpc.CallOption) (*QueueSkill, error)
	DeleteQueueSkill(ctx context.Context, in *DeleteQueueSkillRequest, opts ...grpc.CallOption) (*QueueSkill, error)
}

type queueSkillServiceClient struct {
	cc *grpc.ClientConn
}

func NewQueueSkillServiceClient(cc *grpc.ClientConn) QueueSkillServiceClient {
	return &queueSkillServiceClient{cc}
}

func (c *queueSkillServiceClient) CreateQueueSkill(ctx context.Context, in *CreateQueueSkillRequest, opts ...grpc.CallOption) (*QueueSkill, error) {
	out := new(QueueSkill)
	err := c.cc.Invoke(ctx, "/engine.QueueSkillService/CreateQueueSkill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueSkillServiceClient) SearchQueueSkill(ctx context.Context, in *SearchQueueSkillRequest, opts ...grpc.CallOption) (*ListQueueSkill, error) {
	out := new(ListQueueSkill)
	err := c.cc.Invoke(ctx, "/engine.QueueSkillService/SearchQueueSkill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueSkillServiceClient) ReadQueueSkill(ctx context.Context, in *ReadQueueSkillRequest, opts ...grpc.CallOption) (*QueueSkill, error) {
	out := new(QueueSkill)
	err := c.cc.Invoke(ctx, "/engine.QueueSkillService/ReadQueueSkill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueSkillServiceClient) UpdateQueueSkill(ctx context.Context, in *UpdateQueueSkillRequest, opts ...grpc.CallOption) (*QueueSkill, error) {
	out := new(QueueSkill)
	err := c.cc.Invoke(ctx, "/engine.QueueSkillService/UpdateQueueSkill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueSkillServiceClient) PatchQueueSkill(ctx context.Context, in *PatchQueueSkillRequest, opts ...grpc.CallOption) (*QueueSkill, error) {
	out := new(QueueSkill)
	err := c.cc.Invoke(ctx, "/engine.QueueSkillService/PatchQueueSkill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueSkillServiceClient) DeleteQueueSkill(ctx context.Context, in *DeleteQueueSkillRequest, opts ...grpc.CallOption) (*QueueSkill, error) {
	out := new(QueueSkill)
	err := c.cc.Invoke(ctx, "/engine.QueueSkillService/DeleteQueueSkill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueueSkillServiceServer is the server API for QueueSkillService service.
type QueueSkillServiceServer interface {
	CreateQueueSkill(context.Context, *CreateQueueSkillRequest) (*QueueSkill, error)
	SearchQueueSkill(context.Context, *SearchQueueSkillRequest) (*ListQueueSkill, error)
	ReadQueueSkill(context.Context, *ReadQueueSkillRequest) (*QueueSkill, error)
	UpdateQueueSkill(context.Context, *UpdateQueueSkillRequest) (*QueueSkill, error)
	PatchQueueSkill(context.Context, *PatchQueueSkillRequest) (*QueueSkill, error)
	DeleteQueueSkill(context.Context, *DeleteQueueSkillRequest) (*QueueSkill, error)
}

// UnimplementedQueueSkillServiceServer can be embedded to have forward compatible implementations.
type UnimplementedQueueSkillServiceServer struct {
}

func (*UnimplementedQueueSkillServiceServer) CreateQueueSkill(ctx context.Context, req *CreateQueueSkillRequest) (*QueueSkill, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateQueueSkill not implemented")
}
func (*UnimplementedQueueSkillServiceServer) SearchQueueSkill(ctx context.Context, req *SearchQueueSkillRequest) (*ListQueueSkill, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchQueueSkill not implemented")
}
func (*UnimplementedQueueSkillServiceServer) ReadQueueSkill(ctx context.Context, req *ReadQueueSkillRequest) (*QueueSkill, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadQueueSkill not implemented")
}
func (*UnimplementedQueueSkillServiceServer) UpdateQueueSkill(ctx context.Context, req *UpdateQueueSkillRequest) (*QueueSkill, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateQueueSkill not implemented")
}
func (*UnimplementedQueueSkillServiceServer) PatchQueueSkill(ctx context.Context, req *PatchQueueSkillRequest) (*QueueSkill, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchQueueSkill not implemented")
}
func (*UnimplementedQueueSkillServiceServer) DeleteQueueSkill(ctx context.Context, req *DeleteQueueSkillRequest) (*QueueSkill, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQueueSkill not implemented")
}

func RegisterQueueSkillServiceServer(s *grpc.Server, srv QueueSkillServiceServer) {
	s.RegisterService(&_QueueSkillService_serviceDesc, srv)
}

func _QueueSkillService_CreateQueueSkill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateQueueSkillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueSkillServiceServer).CreateQueueSkill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.QueueSkillService/CreateQueueSkill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueSkillServiceServer).CreateQueueSkill(ctx, req.(*CreateQueueSkillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueueSkillService_SearchQueueSkill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchQueueSkillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueSkillServiceServer).SearchQueueSkill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.QueueSkillService/SearchQueueSkill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueSkillServiceServer).SearchQueueSkill(ctx, req.(*SearchQueueSkillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueueSkillService_ReadQueueSkill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadQueueSkillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueSkillServiceServer).ReadQueueSkill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.QueueSkillService/ReadQueueSkill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueSkillServiceServer).ReadQueueSkill(ctx, req.(*ReadQueueSkillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueueSkillService_UpdateQueueSkill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateQueueSkillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueSkillServiceServer).UpdateQueueSkill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.QueueSkillService/UpdateQueueSkill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueSkillServiceServer).UpdateQueueSkill(ctx, req.(*UpdateQueueSkillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueueSkillService_PatchQueueSkill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchQueueSkillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueSkillServiceServer).PatchQueueSkill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.QueueSkillService/PatchQueueSkill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueSkillServiceServer).PatchQueueSkill(ctx, req.(*PatchQueueSkillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueueSkillService_DeleteQueueSkill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteQueueSkillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueSkillServiceServer).DeleteQueueSkill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.QueueSkillService/DeleteQueueSkill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueSkillServiceServer).DeleteQueueSkill(ctx, req.(*DeleteQueueSkillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QueueSkillService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "engine.QueueSkillService",
	HandlerType: (*QueueSkillServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateQueueSkill",
			Handler:    _QueueSkillService_CreateQueueSkill_Handler,
		},
		{
			MethodName: "SearchQueueSkill",
			Handler:    _QueueSkillService_SearchQueueSkill_Handler,
		},
		{
			MethodName: "ReadQueueSkill",
			Handler:    _QueueSkillService_ReadQueueSkill_Handler,
		},
		{
			MethodName: "UpdateQueueSkill",
			Handler:    _QueueSkillService_UpdateQueueSkill_Handler,
		},
		{
			MethodName: "PatchQueueSkill",
			Handler:    _QueueSkillService_PatchQueueSkill_Handler,
		},
		{
			MethodName: "DeleteQueueSkill",
			Handler:    _QueueSkillService_DeleteQueueSkill_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "queue_skill.proto",
}
