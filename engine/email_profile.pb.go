// Code generated by protoc-gen-go. DO NOT EDIT.
// source: email_profile.proto

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

type ReadEmailProfileRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DomainId             int64    `protobuf:"varint,2,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadEmailProfileRequest) Reset()         { *m = ReadEmailProfileRequest{} }
func (m *ReadEmailProfileRequest) String() string { return proto.CompactTextString(m) }
func (*ReadEmailProfileRequest) ProtoMessage()    {}
func (*ReadEmailProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4c83f7b530e853a, []int{0}
}

func (m *ReadEmailProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadEmailProfileRequest.Unmarshal(m, b)
}
func (m *ReadEmailProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadEmailProfileRequest.Marshal(b, m, deterministic)
}
func (m *ReadEmailProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadEmailProfileRequest.Merge(m, src)
}
func (m *ReadEmailProfileRequest) XXX_Size() int {
	return xxx_messageInfo_ReadEmailProfileRequest.Size(m)
}
func (m *ReadEmailProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadEmailProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadEmailProfileRequest proto.InternalMessageInfo

func (m *ReadEmailProfileRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ReadEmailProfileRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

type DeleteEmailProfileRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DomainId             int64    `protobuf:"varint,2,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteEmailProfileRequest) Reset()         { *m = DeleteEmailProfileRequest{} }
func (m *DeleteEmailProfileRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteEmailProfileRequest) ProtoMessage()    {}
func (*DeleteEmailProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4c83f7b530e853a, []int{1}
}

func (m *DeleteEmailProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteEmailProfileRequest.Unmarshal(m, b)
}
func (m *DeleteEmailProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteEmailProfileRequest.Marshal(b, m, deterministic)
}
func (m *DeleteEmailProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteEmailProfileRequest.Merge(m, src)
}
func (m *DeleteEmailProfileRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteEmailProfileRequest.Size(m)
}
func (m *DeleteEmailProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteEmailProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteEmailProfileRequest proto.InternalMessageInfo

func (m *DeleteEmailProfileRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DeleteEmailProfileRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

type UpdateEmailProfileRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Schema               *Lookup  `protobuf:"bytes,4,opt,name=schema,proto3" json:"schema,omitempty"`
	Enabled              bool     `protobuf:"varint,5,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Host                 string   `protobuf:"bytes,6,opt,name=host,proto3" json:"host,omitempty"`
	Login                string   `protobuf:"bytes,7,opt,name=login,proto3" json:"login,omitempty"`
	Mailbox              string   `protobuf:"bytes,8,opt,name=mailbox,proto3" json:"mailbox,omitempty"`
	SmtpPort             int32    `protobuf:"varint,9,opt,name=smtp_port,json=smtpPort,proto3" json:"smtp_port,omitempty"`
	ImapPort             int32    `protobuf:"varint,10,opt,name=imap_port,json=imapPort,proto3" json:"imap_port,omitempty"`
	Password             string   `protobuf:"bytes,11,opt,name=password,proto3" json:"password,omitempty"`
	DomainId             int64    `protobuf:"varint,12,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateEmailProfileRequest) Reset()         { *m = UpdateEmailProfileRequest{} }
func (m *UpdateEmailProfileRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateEmailProfileRequest) ProtoMessage()    {}
func (*UpdateEmailProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4c83f7b530e853a, []int{2}
}

func (m *UpdateEmailProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateEmailProfileRequest.Unmarshal(m, b)
}
func (m *UpdateEmailProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateEmailProfileRequest.Marshal(b, m, deterministic)
}
func (m *UpdateEmailProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateEmailProfileRequest.Merge(m, src)
}
func (m *UpdateEmailProfileRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateEmailProfileRequest.Size(m)
}
func (m *UpdateEmailProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateEmailProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateEmailProfileRequest proto.InternalMessageInfo

func (m *UpdateEmailProfileRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateEmailProfileRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateEmailProfileRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *UpdateEmailProfileRequest) GetSchema() *Lookup {
	if m != nil {
		return m.Schema
	}
	return nil
}

func (m *UpdateEmailProfileRequest) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *UpdateEmailProfileRequest) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *UpdateEmailProfileRequest) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *UpdateEmailProfileRequest) GetMailbox() string {
	if m != nil {
		return m.Mailbox
	}
	return ""
}

func (m *UpdateEmailProfileRequest) GetSmtpPort() int32 {
	if m != nil {
		return m.SmtpPort
	}
	return 0
}

func (m *UpdateEmailProfileRequest) GetImapPort() int32 {
	if m != nil {
		return m.ImapPort
	}
	return 0
}

func (m *UpdateEmailProfileRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UpdateEmailProfileRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

type SearchEmailProfileRequest struct {
	Page                 int32    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size                 int32    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Q                    string   `protobuf:"bytes,3,opt,name=q,proto3" json:"q,omitempty"`
	Sort                 string   `protobuf:"bytes,4,opt,name=sort,proto3" json:"sort,omitempty"`
	Fields               []string `protobuf:"bytes,5,rep,name=fields,proto3" json:"fields,omitempty"`
	DomainId             int64    `protobuf:"varint,6,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchEmailProfileRequest) Reset()         { *m = SearchEmailProfileRequest{} }
func (m *SearchEmailProfileRequest) String() string { return proto.CompactTextString(m) }
func (*SearchEmailProfileRequest) ProtoMessage()    {}
func (*SearchEmailProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4c83f7b530e853a, []int{3}
}

func (m *SearchEmailProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchEmailProfileRequest.Unmarshal(m, b)
}
func (m *SearchEmailProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchEmailProfileRequest.Marshal(b, m, deterministic)
}
func (m *SearchEmailProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchEmailProfileRequest.Merge(m, src)
}
func (m *SearchEmailProfileRequest) XXX_Size() int {
	return xxx_messageInfo_SearchEmailProfileRequest.Size(m)
}
func (m *SearchEmailProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchEmailProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchEmailProfileRequest proto.InternalMessageInfo

func (m *SearchEmailProfileRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *SearchEmailProfileRequest) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *SearchEmailProfileRequest) GetQ() string {
	if m != nil {
		return m.Q
	}
	return ""
}

func (m *SearchEmailProfileRequest) GetSort() string {
	if m != nil {
		return m.Sort
	}
	return ""
}

func (m *SearchEmailProfileRequest) GetFields() []string {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *SearchEmailProfileRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

type ListEmailProfile struct {
	Next                 bool            `protobuf:"varint,1,opt,name=next,proto3" json:"next,omitempty"`
	Items                []*EmailProfile `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ListEmailProfile) Reset()         { *m = ListEmailProfile{} }
func (m *ListEmailProfile) String() string { return proto.CompactTextString(m) }
func (*ListEmailProfile) ProtoMessage()    {}
func (*ListEmailProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4c83f7b530e853a, []int{4}
}

func (m *ListEmailProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListEmailProfile.Unmarshal(m, b)
}
func (m *ListEmailProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListEmailProfile.Marshal(b, m, deterministic)
}
func (m *ListEmailProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListEmailProfile.Merge(m, src)
}
func (m *ListEmailProfile) XXX_Size() int {
	return xxx_messageInfo_ListEmailProfile.Size(m)
}
func (m *ListEmailProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_ListEmailProfile.DiscardUnknown(m)
}

var xxx_messageInfo_ListEmailProfile proto.InternalMessageInfo

func (m *ListEmailProfile) GetNext() bool {
	if m != nil {
		return m.Next
	}
	return false
}

func (m *ListEmailProfile) GetItems() []*EmailProfile {
	if m != nil {
		return m.Items
	}
	return nil
}

type EmailProfile struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DomainId             int64    `protobuf:"varint,2,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	CreatedAt            int64    `protobuf:"varint,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	CreatedBy            *Lookup  `protobuf:"bytes,4,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	UpdatedAt            int64    `protobuf:"varint,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	UpdatedBy            *Lookup  `protobuf:"bytes,6,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
	Name                 string   `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	Schema               *Lookup  `protobuf:"bytes,9,opt,name=schema,proto3" json:"schema,omitempty"`
	Enabled              bool     `protobuf:"varint,10,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Host                 string   `protobuf:"bytes,11,opt,name=host,proto3" json:"host,omitempty"`
	Login                string   `protobuf:"bytes,12,opt,name=login,proto3" json:"login,omitempty"`
	Mailbox              string   `protobuf:"bytes,13,opt,name=mailbox,proto3" json:"mailbox,omitempty"`
	SmtpPort             int32    `protobuf:"varint,14,opt,name=smtp_port,json=smtpPort,proto3" json:"smtp_port,omitempty"`
	ImapPort             int32    `protobuf:"varint,15,opt,name=imap_port,json=imapPort,proto3" json:"imap_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmailProfile) Reset()         { *m = EmailProfile{} }
func (m *EmailProfile) String() string { return proto.CompactTextString(m) }
func (*EmailProfile) ProtoMessage()    {}
func (*EmailProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4c83f7b530e853a, []int{5}
}

func (m *EmailProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmailProfile.Unmarshal(m, b)
}
func (m *EmailProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmailProfile.Marshal(b, m, deterministic)
}
func (m *EmailProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmailProfile.Merge(m, src)
}
func (m *EmailProfile) XXX_Size() int {
	return xxx_messageInfo_EmailProfile.Size(m)
}
func (m *EmailProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_EmailProfile.DiscardUnknown(m)
}

var xxx_messageInfo_EmailProfile proto.InternalMessageInfo

func (m *EmailProfile) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *EmailProfile) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

func (m *EmailProfile) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *EmailProfile) GetCreatedBy() *Lookup {
	if m != nil {
		return m.CreatedBy
	}
	return nil
}

func (m *EmailProfile) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *EmailProfile) GetUpdatedBy() *Lookup {
	if m != nil {
		return m.UpdatedBy
	}
	return nil
}

func (m *EmailProfile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EmailProfile) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *EmailProfile) GetSchema() *Lookup {
	if m != nil {
		return m.Schema
	}
	return nil
}

func (m *EmailProfile) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *EmailProfile) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *EmailProfile) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *EmailProfile) GetMailbox() string {
	if m != nil {
		return m.Mailbox
	}
	return ""
}

func (m *EmailProfile) GetSmtpPort() int32 {
	if m != nil {
		return m.SmtpPort
	}
	return 0
}

func (m *EmailProfile) GetImapPort() int32 {
	if m != nil {
		return m.ImapPort
	}
	return 0
}

type CreateEmailProfileRequest struct {
	DomainId             int64    `protobuf:"varint,1,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Schema               *Lookup  `protobuf:"bytes,4,opt,name=schema,proto3" json:"schema,omitempty"`
	Enabled              bool     `protobuf:"varint,5,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Host                 string   `protobuf:"bytes,6,opt,name=host,proto3" json:"host,omitempty"`
	Login                string   `protobuf:"bytes,7,opt,name=login,proto3" json:"login,omitempty"`
	Password             string   `protobuf:"bytes,8,opt,name=password,proto3" json:"password,omitempty"`
	Mailbox              string   `protobuf:"bytes,9,opt,name=mailbox,proto3" json:"mailbox,omitempty"`
	SmtpPort             int32    `protobuf:"varint,10,opt,name=smtp_port,json=smtpPort,proto3" json:"smtp_port,omitempty"`
	ImapPort             int32    `protobuf:"varint,11,opt,name=imap_port,json=imapPort,proto3" json:"imap_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateEmailProfileRequest) Reset()         { *m = CreateEmailProfileRequest{} }
func (m *CreateEmailProfileRequest) String() string { return proto.CompactTextString(m) }
func (*CreateEmailProfileRequest) ProtoMessage()    {}
func (*CreateEmailProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4c83f7b530e853a, []int{6}
}

func (m *CreateEmailProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEmailProfileRequest.Unmarshal(m, b)
}
func (m *CreateEmailProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEmailProfileRequest.Marshal(b, m, deterministic)
}
func (m *CreateEmailProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEmailProfileRequest.Merge(m, src)
}
func (m *CreateEmailProfileRequest) XXX_Size() int {
	return xxx_messageInfo_CreateEmailProfileRequest.Size(m)
}
func (m *CreateEmailProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEmailProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEmailProfileRequest proto.InternalMessageInfo

func (m *CreateEmailProfileRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

func (m *CreateEmailProfileRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateEmailProfileRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreateEmailProfileRequest) GetSchema() *Lookup {
	if m != nil {
		return m.Schema
	}
	return nil
}

func (m *CreateEmailProfileRequest) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *CreateEmailProfileRequest) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *CreateEmailProfileRequest) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *CreateEmailProfileRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CreateEmailProfileRequest) GetMailbox() string {
	if m != nil {
		return m.Mailbox
	}
	return ""
}

func (m *CreateEmailProfileRequest) GetSmtpPort() int32 {
	if m != nil {
		return m.SmtpPort
	}
	return 0
}

func (m *CreateEmailProfileRequest) GetImapPort() int32 {
	if m != nil {
		return m.ImapPort
	}
	return 0
}

func init() {
	proto.RegisterType((*ReadEmailProfileRequest)(nil), "engine.ReadEmailProfileRequest")
	proto.RegisterType((*DeleteEmailProfileRequest)(nil), "engine.DeleteEmailProfileRequest")
	proto.RegisterType((*UpdateEmailProfileRequest)(nil), "engine.UpdateEmailProfileRequest")
	proto.RegisterType((*SearchEmailProfileRequest)(nil), "engine.SearchEmailProfileRequest")
	proto.RegisterType((*ListEmailProfile)(nil), "engine.ListEmailProfile")
	proto.RegisterType((*EmailProfile)(nil), "engine.EmailProfile")
	proto.RegisterType((*CreateEmailProfileRequest)(nil), "engine.CreateEmailProfileRequest")
}

func init() { proto.RegisterFile("email_profile.proto", fileDescriptor_f4c83f7b530e853a) }

var fileDescriptor_f4c83f7b530e853a = []byte{
	// 746 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0xcd, 0x6e, 0x13, 0x3b,
	0x14, 0xd6, 0x24, 0x99, 0x74, 0x72, 0x92, 0x9b, 0x5b, 0x39, 0xbd, 0xbd, 0x4e, 0x7a, 0x2f, 0x0c,
	0xb3, 0x40, 0x51, 0x25, 0x12, 0xa9, 0xec, 0xd8, 0x51, 0x7e, 0x04, 0x52, 0x17, 0xd5, 0x54, 0x6c,
	0xd8, 0x44, 0x4e, 0xc6, 0x4d, 0x0c, 0x99, 0xf1, 0x74, 0xec, 0xd0, 0x06, 0xc4, 0x86, 0x57, 0x60,
	0xc7, 0x3b, 0xf0, 0x22, 0x48, 0xac, 0xd8, 0xb1, 0xe6, 0x41, 0x90, 0xed, 0x49, 0x93, 0x49, 0x66,
	0xa2, 0x22, 0x36, 0xec, 0x7c, 0xce, 0xb1, 0xbe, 0xcf, 0xe7, 0xf3, 0x67, 0x1f, 0x68, 0xd1, 0x90,
	0xb0, 0xe9, 0x20, 0x4e, 0xf8, 0x39, 0x9b, 0xd2, 0x5e, 0x9c, 0x70, 0xc9, 0x51, 0x95, 0x46, 0x63,
	0x16, 0xd1, 0x4e, 0x7d, 0xc4, 0x23, 0x21, 0x4d, 0xb2, 0xf3, 0xdf, 0x98, 0xf3, 0xf1, 0x94, 0xf6,
	0x49, 0xcc, 0xfa, 0x24, 0x8a, 0xb8, 0x24, 0x92, 0xf1, 0x48, 0x98, 0xaa, 0xf7, 0x14, 0xfe, 0xf5,
	0x29, 0x09, 0x9e, 0x28, 0xb4, 0x53, 0x03, 0xe6, 0xd3, 0x8b, 0x19, 0x15, 0x12, 0x35, 0xa1, 0xc4,
	0x02, 0x6c, 0xb9, 0x56, 0xb7, 0xec, 0x97, 0x58, 0x80, 0x0e, 0xa0, 0x16, 0xf0, 0x90, 0xb0, 0x68,
	0xc0, 0x02, 0x5c, 0xd2, 0x69, 0xc7, 0x24, 0x9e, 0x07, 0xde, 0x33, 0x68, 0x3f, 0xa6, 0x53, 0x2a,
	0xe9, 0x6f, 0x23, 0x7d, 0x2f, 0x41, 0xfb, 0x45, 0x1c, 0x90, 0x9b, 0x41, 0x21, 0xa8, 0x44, 0x24,
	0xa4, 0x1a, 0xa5, 0xe6, 0xeb, 0x35, 0x72, 0xa1, 0x1e, 0x50, 0x31, 0x4a, 0x58, 0xac, 0x3a, 0xc5,
	0x65, 0x5d, 0x5a, 0x4d, 0xa1, 0xbb, 0x50, 0x15, 0xa3, 0x09, 0x0d, 0x09, 0xae, 0xb8, 0x56, 0xb7,
	0x7e, 0xd4, 0xec, 0x19, 0xe5, 0x7a, 0x27, 0x9c, 0xbf, 0x9e, 0xc5, 0x7e, 0x5a, 0x45, 0x18, 0x76,
	0x68, 0x44, 0x86, 0x53, 0x1a, 0x60, 0xdb, 0xb5, 0xba, 0x8e, 0xbf, 0x08, 0x15, 0xef, 0x84, 0x0b,
	0x89, 0xab, 0x86, 0x57, 0xad, 0xd1, 0x1e, 0xd8, 0x53, 0x3e, 0x66, 0x11, 0xde, 0xd1, 0x49, 0x13,
	0x28, 0x0c, 0xd5, 0xc7, 0x90, 0x5f, 0x61, 0x47, 0xe7, 0x17, 0xa1, 0x92, 0x41, 0x84, 0x32, 0x1e,
	0xc4, 0x3c, 0x91, 0xb8, 0xe6, 0x5a, 0x5d, 0xdb, 0x77, 0x54, 0xe2, 0x94, 0x27, 0x52, 0x15, 0x59,
	0x48, 0xd2, 0x22, 0x98, 0xa2, 0x4a, 0xe8, 0x62, 0x07, 0x9c, 0x98, 0x08, 0x71, 0xc9, 0x93, 0x00,
	0xd7, 0x35, 0xe8, 0x75, 0x9c, 0x15, 0xb7, 0xb1, 0x26, 0xee, 0x27, 0x0b, 0xda, 0x67, 0x94, 0x24,
	0xa3, 0x49, 0x9e, 0xb8, 0x08, 0x2a, 0x31, 0x19, 0x53, 0x2d, 0xaf, 0xed, 0xeb, 0xb5, 0xca, 0x09,
	0xf6, 0xd6, 0x08, 0x6c, 0xfb, 0x7a, 0x8d, 0x1a, 0x60, 0x5d, 0xa4, 0xb2, 0x5a, 0x17, 0x7a, 0x87,
	0x3a, 0x64, 0xc5, 0x48, 0xa1, 0xd6, 0x68, 0x1f, 0xaa, 0xe7, 0x8c, 0x4e, 0x03, 0x81, 0x6d, 0xb7,
	0xdc, 0xad, 0xf9, 0x69, 0x94, 0x3d, 0x5c, 0x75, 0xed, 0x70, 0x3e, 0xec, 0x9e, 0x30, 0x21, 0x57,
	0x4f, 0xa6, 0xef, 0x97, 0x5e, 0x49, 0x7d, 0x24, 0xc7, 0xd7, 0x6b, 0x74, 0x08, 0x36, 0x93, 0x34,
	0x14, 0xb8, 0xe4, 0x96, 0xbb, 0xf5, 0xa3, 0xbd, 0xc5, 0xe5, 0x65, 0x5a, 0x32, 0x5b, 0xbc, 0xaf,
	0x65, 0x68, 0x64, 0x00, 0x7f, 0xc5, 0x8b, 0xe8, 0x7f, 0x80, 0x51, 0x42, 0x89, 0xa4, 0xc1, 0x80,
	0x48, 0xdd, 0x71, 0xd9, 0xaf, 0xa5, 0x99, 0x87, 0x12, 0xdd, 0x5b, 0x96, 0x87, 0xf3, 0x02, 0x2b,
	0x2d, 0xb6, 0x1f, 0xcf, 0x15, 0xda, 0x4c, 0x1b, 0x5b, 0xa3, 0xd9, 0x06, 0x2d, 0xcd, 0x18, 0xb4,
	0x45, 0x79, 0x38, 0xd7, 0xe2, 0xe4, 0xa0, 0xa5, 0x3b, 0x8e, 0xe7, 0xd7, 0xce, 0xdf, 0x29, 0x76,
	0xbe, 0xb3, 0xcd, 0xf9, 0xb5, 0x9b, 0x3a, 0x1f, 0xf2, 0x9d, 0x5f, 0xcf, 0x73, 0x7e, 0xa3, 0xc0,
	0xf9, 0x7f, 0x6d, 0x71, 0x7e, 0x73, 0x9b, 0xf3, 0xff, 0xce, 0x3a, 0xdf, 0xfb, 0x52, 0x82, 0xf6,
	0x23, 0xad, 0x68, 0x9e, 0x81, 0x33, 0x97, 0x69, 0xad, 0x5d, 0xe6, 0x9f, 0xff, 0x55, 0xac, 0x3e,
	0x6b, 0x67, 0xed, 0x59, 0xaf, 0x88, 0x59, 0xdb, 0x22, 0x26, 0x6c, 0x13, 0xb3, 0x9e, 0x15, 0xf3,
	0xe8, 0x73, 0x05, 0x5a, 0xab, 0x32, 0x9e, 0xd1, 0xe4, 0x0d, 0x1b, 0x51, 0x34, 0x01, 0xb4, 0xa9,
	0x31, 0xba, 0xb3, 0xe8, 0xbc, 0x50, 0xff, 0x4e, 0xee, 0x53, 0xf4, 0xda, 0x1f, 0xbe, 0xfd, 0xf8,
	0x58, 0x6a, 0x79, 0xcd, 0xbe, 0x1e, 0x5a, 0xfd, 0x74, 0x68, 0x3d, 0xb0, 0x0e, 0x11, 0x03, 0xb4,
	0xf9, 0x1d, 0x2d, 0x99, 0x0a, 0xbf, 0xaa, 0x0e, 0xbe, 0xbe, 0x86, 0xb5, 0x1f, 0xc3, 0xdb, 0xd7,
	0x6c, 0xbb, 0x68, 0x8d, 0x0d, 0x9d, 0xc3, 0xee, 0xfa, 0xa4, 0x43, 0xb7, 0x17, 0x28, 0x05, 0x33,
	0xb0, 0xa0, 0xa1, 0x03, 0x4d, 0xf1, 0x0f, 0x6a, 0x65, 0x29, 0xfa, 0xef, 0x58, 0xf0, 0x1e, 0x85,
	0x80, 0x36, 0xc7, 0xd7, 0xb2, 0xa5, 0xc2, 0xd1, 0x56, 0xc0, 0x75, 0x4b, 0x73, 0xe1, 0x4e, 0x1e,
	0x97, 0x52, 0xf0, 0x15, 0xa0, 0xcd, 0xc1, 0xbb, 0xa4, 0x2b, 0x1c, 0xca, 0xdb, 0x5b, 0x3b, 0xcc,
	0xa3, 0x3b, 0xf6, 0x5e, 0xba, 0x63, 0x26, 0x27, 0xb3, 0x61, 0x6f, 0xc4, 0xc3, 0xfe, 0x25, 0x1d,
	0x32, 0x49, 0xf5, 0x16, 0xc9, 0x45, 0xdf, 0xa0, 0x0d, 0xab, 0x3a, 0xbc, 0xff, 0x33, 0x00, 0x00,
	0xff, 0xff, 0x0e, 0xb6, 0x49, 0x8a, 0xa1, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EmailProfileServiceClient is the client API for EmailProfileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EmailProfileServiceClient interface {
	// Create EmailProfile
	CreateEmailProfile(ctx context.Context, in *CreateEmailProfileRequest, opts ...grpc.CallOption) (*EmailProfile, error)
	// Search EmailProfile
	SearchEmailProfile(ctx context.Context, in *SearchEmailProfileRequest, opts ...grpc.CallOption) (*ListEmailProfile, error)
	// EmailProfile item
	ReadEmailProfile(ctx context.Context, in *ReadEmailProfileRequest, opts ...grpc.CallOption) (*EmailProfile, error)
	// Update EmailProfile
	UpdateEmailProfile(ctx context.Context, in *UpdateEmailProfileRequest, opts ...grpc.CallOption) (*EmailProfile, error)
	// Remove EmailProfile
	DeleteEmailProfile(ctx context.Context, in *DeleteEmailProfileRequest, opts ...grpc.CallOption) (*EmailProfile, error)
}

type emailProfileServiceClient struct {
	cc *grpc.ClientConn
}

func NewEmailProfileServiceClient(cc *grpc.ClientConn) EmailProfileServiceClient {
	return &emailProfileServiceClient{cc}
}

func (c *emailProfileServiceClient) CreateEmailProfile(ctx context.Context, in *CreateEmailProfileRequest, opts ...grpc.CallOption) (*EmailProfile, error) {
	out := new(EmailProfile)
	err := c.cc.Invoke(ctx, "/engine.EmailProfileService/CreateEmailProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailProfileServiceClient) SearchEmailProfile(ctx context.Context, in *SearchEmailProfileRequest, opts ...grpc.CallOption) (*ListEmailProfile, error) {
	out := new(ListEmailProfile)
	err := c.cc.Invoke(ctx, "/engine.EmailProfileService/SearchEmailProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailProfileServiceClient) ReadEmailProfile(ctx context.Context, in *ReadEmailProfileRequest, opts ...grpc.CallOption) (*EmailProfile, error) {
	out := new(EmailProfile)
	err := c.cc.Invoke(ctx, "/engine.EmailProfileService/ReadEmailProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailProfileServiceClient) UpdateEmailProfile(ctx context.Context, in *UpdateEmailProfileRequest, opts ...grpc.CallOption) (*EmailProfile, error) {
	out := new(EmailProfile)
	err := c.cc.Invoke(ctx, "/engine.EmailProfileService/UpdateEmailProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailProfileServiceClient) DeleteEmailProfile(ctx context.Context, in *DeleteEmailProfileRequest, opts ...grpc.CallOption) (*EmailProfile, error) {
	out := new(EmailProfile)
	err := c.cc.Invoke(ctx, "/engine.EmailProfileService/DeleteEmailProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmailProfileServiceServer is the server API for EmailProfileService service.
type EmailProfileServiceServer interface {
	// Create EmailProfile
	CreateEmailProfile(context.Context, *CreateEmailProfileRequest) (*EmailProfile, error)
	// Search EmailProfile
	SearchEmailProfile(context.Context, *SearchEmailProfileRequest) (*ListEmailProfile, error)
	// EmailProfile item
	ReadEmailProfile(context.Context, *ReadEmailProfileRequest) (*EmailProfile, error)
	// Update EmailProfile
	UpdateEmailProfile(context.Context, *UpdateEmailProfileRequest) (*EmailProfile, error)
	// Remove EmailProfile
	DeleteEmailProfile(context.Context, *DeleteEmailProfileRequest) (*EmailProfile, error)
}

// UnimplementedEmailProfileServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEmailProfileServiceServer struct {
}

func (*UnimplementedEmailProfileServiceServer) CreateEmailProfile(ctx context.Context, req *CreateEmailProfileRequest) (*EmailProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEmailProfile not implemented")
}
func (*UnimplementedEmailProfileServiceServer) SearchEmailProfile(ctx context.Context, req *SearchEmailProfileRequest) (*ListEmailProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchEmailProfile not implemented")
}
func (*UnimplementedEmailProfileServiceServer) ReadEmailProfile(ctx context.Context, req *ReadEmailProfileRequest) (*EmailProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadEmailProfile not implemented")
}
func (*UnimplementedEmailProfileServiceServer) UpdateEmailProfile(ctx context.Context, req *UpdateEmailProfileRequest) (*EmailProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEmailProfile not implemented")
}
func (*UnimplementedEmailProfileServiceServer) DeleteEmailProfile(ctx context.Context, req *DeleteEmailProfileRequest) (*EmailProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEmailProfile not implemented")
}

func RegisterEmailProfileServiceServer(s *grpc.Server, srv EmailProfileServiceServer) {
	s.RegisterService(&_EmailProfileService_serviceDesc, srv)
}

func _EmailProfileService_CreateEmailProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEmailProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailProfileServiceServer).CreateEmailProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.EmailProfileService/CreateEmailProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailProfileServiceServer).CreateEmailProfile(ctx, req.(*CreateEmailProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailProfileService_SearchEmailProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchEmailProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailProfileServiceServer).SearchEmailProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.EmailProfileService/SearchEmailProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailProfileServiceServer).SearchEmailProfile(ctx, req.(*SearchEmailProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailProfileService_ReadEmailProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadEmailProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailProfileServiceServer).ReadEmailProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.EmailProfileService/ReadEmailProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailProfileServiceServer).ReadEmailProfile(ctx, req.(*ReadEmailProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailProfileService_UpdateEmailProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEmailProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailProfileServiceServer).UpdateEmailProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.EmailProfileService/UpdateEmailProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailProfileServiceServer).UpdateEmailProfile(ctx, req.(*UpdateEmailProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailProfileService_DeleteEmailProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEmailProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailProfileServiceServer).DeleteEmailProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.EmailProfileService/DeleteEmailProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailProfileServiceServer).DeleteEmailProfile(ctx, req.(*DeleteEmailProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EmailProfileService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "engine.EmailProfileService",
	HandlerType: (*EmailProfileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEmailProfile",
			Handler:    _EmailProfileService_CreateEmailProfile_Handler,
		},
		{
			MethodName: "SearchEmailProfile",
			Handler:    _EmailProfileService_SearchEmailProfile_Handler,
		},
		{
			MethodName: "ReadEmailProfile",
			Handler:    _EmailProfileService_ReadEmailProfile_Handler,
		},
		{
			MethodName: "UpdateEmailProfile",
			Handler:    _EmailProfileService_UpdateEmailProfile_Handler,
		},
		{
			MethodName: "DeleteEmailProfile",
			Handler:    _EmailProfileService_DeleteEmailProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "email_profile.proto",
}
