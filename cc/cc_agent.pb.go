// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cc_agent.proto

package cc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type WaitingChannelRequest struct {
	AgentId              int32    `protobuf:"varint,1,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	Channel              string   `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel,omitempty"`
	DomainId             int64    `protobuf:"varint,3,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WaitingChannelRequest) Reset()         { *m = WaitingChannelRequest{} }
func (m *WaitingChannelRequest) String() string { return proto.CompactTextString(m) }
func (*WaitingChannelRequest) ProtoMessage()    {}
func (*WaitingChannelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_faeb812e10656ee6, []int{0}
}

func (m *WaitingChannelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WaitingChannelRequest.Unmarshal(m, b)
}
func (m *WaitingChannelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WaitingChannelRequest.Marshal(b, m, deterministic)
}
func (m *WaitingChannelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WaitingChannelRequest.Merge(m, src)
}
func (m *WaitingChannelRequest) XXX_Size() int {
	return xxx_messageInfo_WaitingChannelRequest.Size(m)
}
func (m *WaitingChannelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WaitingChannelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WaitingChannelRequest proto.InternalMessageInfo

func (m *WaitingChannelRequest) GetAgentId() int32 {
	if m != nil {
		return m.AgentId
	}
	return 0
}

func (m *WaitingChannelRequest) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *WaitingChannelRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

type WaitingChannelResponse struct {
	Timestamp            int64    `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WaitingChannelResponse) Reset()         { *m = WaitingChannelResponse{} }
func (m *WaitingChannelResponse) String() string { return proto.CompactTextString(m) }
func (*WaitingChannelResponse) ProtoMessage()    {}
func (*WaitingChannelResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_faeb812e10656ee6, []int{1}
}

func (m *WaitingChannelResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WaitingChannelResponse.Unmarshal(m, b)
}
func (m *WaitingChannelResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WaitingChannelResponse.Marshal(b, m, deterministic)
}
func (m *WaitingChannelResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WaitingChannelResponse.Merge(m, src)
}
func (m *WaitingChannelResponse) XXX_Size() int {
	return xxx_messageInfo_WaitingChannelResponse.Size(m)
}
func (m *WaitingChannelResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WaitingChannelResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WaitingChannelResponse proto.InternalMessageInfo

func (m *WaitingChannelResponse) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type Channel struct {
	Channel              string   `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	State                string   `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	JoinedAt             int64    `protobuf:"varint,3,opt,name=joined_at,json=joinedAt,proto3" json:"joined_at,omitempty"`
	Enabled              bool     `protobuf:"varint,4,opt,name=enabled,proto3" json:"enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Channel) Reset()         { *m = Channel{} }
func (m *Channel) String() string { return proto.CompactTextString(m) }
func (*Channel) ProtoMessage()    {}
func (*Channel) Descriptor() ([]byte, []int) {
	return fileDescriptor_faeb812e10656ee6, []int{2}
}

func (m *Channel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Channel.Unmarshal(m, b)
}
func (m *Channel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Channel.Marshal(b, m, deterministic)
}
func (m *Channel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Channel.Merge(m, src)
}
func (m *Channel) XXX_Size() int {
	return xxx_messageInfo_Channel.Size(m)
}
func (m *Channel) XXX_DiscardUnknown() {
	xxx_messageInfo_Channel.DiscardUnknown(m)
}

var xxx_messageInfo_Channel proto.InternalMessageInfo

func (m *Channel) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *Channel) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Channel) GetJoinedAt() int64 {
	if m != nil {
		return m.JoinedAt
	}
	return 0
}

func (m *Channel) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

type PauseRequest struct {
	AgentId              int64    `protobuf:"varint,1,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	Payload              string   `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	Timeout              int32    `protobuf:"varint,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	DomainId             int64    `protobuf:"varint,4,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PauseRequest) Reset()         { *m = PauseRequest{} }
func (m *PauseRequest) String() string { return proto.CompactTextString(m) }
func (*PauseRequest) ProtoMessage()    {}
func (*PauseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_faeb812e10656ee6, []int{3}
}

func (m *PauseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PauseRequest.Unmarshal(m, b)
}
func (m *PauseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PauseRequest.Marshal(b, m, deterministic)
}
func (m *PauseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PauseRequest.Merge(m, src)
}
func (m *PauseRequest) XXX_Size() int {
	return xxx_messageInfo_PauseRequest.Size(m)
}
func (m *PauseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PauseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PauseRequest proto.InternalMessageInfo

func (m *PauseRequest) GetAgentId() int64 {
	if m != nil {
		return m.AgentId
	}
	return 0
}

func (m *PauseRequest) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func (m *PauseRequest) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *PauseRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

type PauseResponse struct {
	Timestamp            int64    `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PauseResponse) Reset()         { *m = PauseResponse{} }
func (m *PauseResponse) String() string { return proto.CompactTextString(m) }
func (*PauseResponse) ProtoMessage()    {}
func (*PauseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_faeb812e10656ee6, []int{4}
}

func (m *PauseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PauseResponse.Unmarshal(m, b)
}
func (m *PauseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PauseResponse.Marshal(b, m, deterministic)
}
func (m *PauseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PauseResponse.Merge(m, src)
}
func (m *PauseResponse) XXX_Size() int {
	return xxx_messageInfo_PauseResponse.Size(m)
}
func (m *PauseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PauseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PauseResponse proto.InternalMessageInfo

func (m *PauseResponse) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type OfflineRequest struct {
	AgentId              int64    `protobuf:"varint,1,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	DomainId             int64    `protobuf:"varint,2,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OfflineRequest) Reset()         { *m = OfflineRequest{} }
func (m *OfflineRequest) String() string { return proto.CompactTextString(m) }
func (*OfflineRequest) ProtoMessage()    {}
func (*OfflineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_faeb812e10656ee6, []int{5}
}

func (m *OfflineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OfflineRequest.Unmarshal(m, b)
}
func (m *OfflineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OfflineRequest.Marshal(b, m, deterministic)
}
func (m *OfflineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OfflineRequest.Merge(m, src)
}
func (m *OfflineRequest) XXX_Size() int {
	return xxx_messageInfo_OfflineRequest.Size(m)
}
func (m *OfflineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OfflineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OfflineRequest proto.InternalMessageInfo

func (m *OfflineRequest) GetAgentId() int64 {
	if m != nil {
		return m.AgentId
	}
	return 0
}

func (m *OfflineRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

type OfflineResponse struct {
	Timestamp            int64    `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OfflineResponse) Reset()         { *m = OfflineResponse{} }
func (m *OfflineResponse) String() string { return proto.CompactTextString(m) }
func (*OfflineResponse) ProtoMessage()    {}
func (*OfflineResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_faeb812e10656ee6, []int{6}
}

func (m *OfflineResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OfflineResponse.Unmarshal(m, b)
}
func (m *OfflineResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OfflineResponse.Marshal(b, m, deterministic)
}
func (m *OfflineResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OfflineResponse.Merge(m, src)
}
func (m *OfflineResponse) XXX_Size() int {
	return xxx_messageInfo_OfflineResponse.Size(m)
}
func (m *OfflineResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OfflineResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OfflineResponse proto.InternalMessageInfo

func (m *OfflineResponse) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type OnlineRequest struct {
	AgentId              int64    `protobuf:"varint,1,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	Channels             []string `protobuf:"bytes,3,rep,name=channels,proto3" json:"channels,omitempty"`
	OnDemand             bool     `protobuf:"varint,4,opt,name=on_demand,json=onDemand,proto3" json:"on_demand,omitempty"`
	DomainId             int64    `protobuf:"varint,5,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OnlineRequest) Reset()         { *m = OnlineRequest{} }
func (m *OnlineRequest) String() string { return proto.CompactTextString(m) }
func (*OnlineRequest) ProtoMessage()    {}
func (*OnlineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_faeb812e10656ee6, []int{7}
}

func (m *OnlineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OnlineRequest.Unmarshal(m, b)
}
func (m *OnlineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OnlineRequest.Marshal(b, m, deterministic)
}
func (m *OnlineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OnlineRequest.Merge(m, src)
}
func (m *OnlineRequest) XXX_Size() int {
	return xxx_messageInfo_OnlineRequest.Size(m)
}
func (m *OnlineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OnlineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OnlineRequest proto.InternalMessageInfo

func (m *OnlineRequest) GetAgentId() int64 {
	if m != nil {
		return m.AgentId
	}
	return 0
}

func (m *OnlineRequest) GetChannels() []string {
	if m != nil {
		return m.Channels
	}
	return nil
}

func (m *OnlineRequest) GetOnDemand() bool {
	if m != nil {
		return m.OnDemand
	}
	return false
}

func (m *OnlineRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

type OnlineResponse struct {
	Timestamp            int64      `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Channels             []*Channel `protobuf:"bytes,2,rep,name=channels,proto3" json:"channels,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *OnlineResponse) Reset()         { *m = OnlineResponse{} }
func (m *OnlineResponse) String() string { return proto.CompactTextString(m) }
func (*OnlineResponse) ProtoMessage()    {}
func (*OnlineResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_faeb812e10656ee6, []int{8}
}

func (m *OnlineResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OnlineResponse.Unmarshal(m, b)
}
func (m *OnlineResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OnlineResponse.Marshal(b, m, deterministic)
}
func (m *OnlineResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OnlineResponse.Merge(m, src)
}
func (m *OnlineResponse) XXX_Size() int {
	return xxx_messageInfo_OnlineResponse.Size(m)
}
func (m *OnlineResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OnlineResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OnlineResponse proto.InternalMessageInfo

func (m *OnlineResponse) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *OnlineResponse) GetChannels() []*Channel {
	if m != nil {
		return m.Channels
	}
	return nil
}

func init() {
	proto.RegisterType((*WaitingChannelRequest)(nil), "cc.WaitingChannelRequest")
	proto.RegisterType((*WaitingChannelResponse)(nil), "cc.WaitingChannelResponse")
	proto.RegisterType((*Channel)(nil), "cc.Channel")
	proto.RegisterType((*PauseRequest)(nil), "cc.PauseRequest")
	proto.RegisterType((*PauseResponse)(nil), "cc.PauseResponse")
	proto.RegisterType((*OfflineRequest)(nil), "cc.OfflineRequest")
	proto.RegisterType((*OfflineResponse)(nil), "cc.OfflineResponse")
	proto.RegisterType((*OnlineRequest)(nil), "cc.OnlineRequest")
	proto.RegisterType((*OnlineResponse)(nil), "cc.OnlineResponse")
}

func init() { proto.RegisterFile("cc_agent.proto", fileDescriptor_faeb812e10656ee6) }

var fileDescriptor_faeb812e10656ee6 = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x6f, 0xd4, 0x30,
	0x10, 0xc5, 0x9b, 0xa6, 0x9b, 0x9d, 0xb6, 0x0b, 0x35, 0x1f, 0x4a, 0x03, 0x87, 0x28, 0x17, 0x72,
	0x61, 0x11, 0x05, 0x71, 0xaf, 0xe0, 0xc0, 0x9e, 0x8a, 0xcc, 0xa1, 0xc7, 0x95, 0x6b, 0xbb, 0xc5,
	0x28, 0xb1, 0x43, 0xe3, 0x45, 0x82, 0x13, 0xbf, 0x99, 0x5f, 0x80, 0x1c, 0xc7, 0x21, 0x8e, 0x8a,
	0x94, 0xe3, 0x7b, 0x33, 0xe3, 0xf7, 0xf2, 0xc6, 0x0e, 0xac, 0x19, 0xdb, 0xd1, 0x5b, 0xa1, 0xcc,
	0xa6, 0xb9, 0xd3, 0x46, 0xe3, 0x05, 0x63, 0x85, 0x84, 0xa7, 0x57, 0x54, 0x1a, 0xa9, 0x6e, 0x3f,
	0x7c, 0xa5, 0x4a, 0x89, 0x8a, 0x88, 0xef, 0x7b, 0xd1, 0x1a, 0x7c, 0x06, 0x49, 0xd7, 0xbb, 0x93,
	0x3c, 0x45, 0x39, 0x2a, 0x63, 0xb2, 0xec, 0xf0, 0x96, 0xe3, 0x14, 0x96, 0xcc, 0x35, 0xa7, 0x8b,
	0x1c, 0x95, 0x2b, 0xe2, 0x21, 0x7e, 0x0e, 0x2b, 0xae, 0x6b, 0x2a, 0x95, 0x9d, 0x8a, 0x72, 0x54,
	0x46, 0x24, 0x71, 0xc4, 0x96, 0x17, 0xef, 0xe1, 0xd9, 0x54, 0xaa, 0x6d, 0xb4, 0x6a, 0x05, 0x7e,
	0x01, 0x2b, 0x23, 0x6b, 0xd1, 0x1a, 0x5a, 0x37, 0x9d, 0x58, 0x44, 0xfe, 0x11, 0x45, 0x03, 0xcb,
	0x7e, 0x60, 0xac, 0x8c, 0x42, 0xe5, 0x27, 0x10, 0xb7, 0x86, 0x1a, 0xd1, 0x3b, 0x72, 0xc0, 0xfa,
	0xf9, 0xa6, 0xa5, 0x12, 0x7c, 0x47, 0x8d, 0xf7, 0xe3, 0x88, 0x0b, 0x63, 0x0f, 0x13, 0x8a, 0x5e,
	0x57, 0x82, 0xa7, 0x07, 0x39, 0x2a, 0x13, 0xe2, 0x61, 0xf1, 0x0b, 0x8e, 0x3f, 0xd3, 0x7d, 0x2b,
	0xfe, 0x97, 0x45, 0x14, 0x64, 0xd1, 0xd0, 0x9f, 0x95, 0xa6, 0xdc, 0x67, 0xd1, 0x43, 0x5b, 0xb1,
	0xdf, 0xa0, 0xf7, 0x4e, 0x39, 0x26, 0x1e, 0x86, 0x29, 0x1d, 0x4c, 0x52, 0x7a, 0x05, 0x27, 0xbd,
	0xf6, 0xac, 0x70, 0x3e, 0xc1, 0xfa, 0xf2, 0xe6, 0xa6, 0x92, 0x6a, 0x8e, 0xd9, 0x40, 0x78, 0x31,
	0x11, 0x7e, 0x0d, 0x0f, 0x87, 0x93, 0x66, 0x49, 0xff, 0x46, 0x70, 0x72, 0xa9, 0x66, 0x4a, 0x67,
	0x90, 0xf4, 0xab, 0x6a, 0xd3, 0x28, 0x8f, 0xca, 0x15, 0x19, 0xb0, 0xb5, 0xa5, 0xd5, 0x8e, 0x8b,
	0x9a, 0x2a, 0xbf, 0x8a, 0x44, 0xab, 0x8f, 0x1d, 0x0e, 0x3d, 0xc7, 0x13, 0xcf, 0x57, 0xb0, 0xf6,
	0x0e, 0xe6, 0x58, 0xc6, 0x2f, 0x47, 0x2e, 0x16, 0x79, 0x54, 0x1e, 0x9d, 0x1f, 0x6d, 0x18, 0xdb,
	0xf8, 0xfb, 0x38, 0x14, 0xcf, 0xff, 0x20, 0x38, 0xbe, 0xb0, 0xd6, 0xbf, 0x88, 0xbb, 0x1f, 0x92,
	0x09, 0xfc, 0x06, 0x0e, 0x9d, 0x12, 0x3e, 0xb5, 0x13, 0xc1, 0x77, 0x67, 0x78, 0x4c, 0x39, 0x23,
	0xc5, 0x03, 0xfc, 0x0e, 0x96, 0x7d, 0xa0, 0xd8, 0x35, 0x04, 0x7b, 0xca, 0x1e, 0x07, 0xdc, 0x30,
	0xb5, 0x81, 0xb8, 0xdb, 0x3f, 0x7e, 0x64, 0xeb, 0xe3, 0x6b, 0x98, 0x9d, 0x8e, 0x98, 0xa1, 0x7f,
	0x0b, 0xeb, 0xf0, 0x55, 0xe1, 0x33, 0xdb, 0x76, 0xef, 0xa3, 0xce, 0xb2, 0xfb, 0x4a, 0xfe, 0xa8,
	0xeb, 0xc3, 0xee, 0xb7, 0xf0, 0xf6, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x06, 0xd4, 0x5e,
	0x28, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AgentServiceClient is the client API for AgentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AgentServiceClient interface {
	Online(ctx context.Context, in *OnlineRequest, opts ...grpc.CallOption) (*OnlineResponse, error)
	Offline(ctx context.Context, in *OfflineRequest, opts ...grpc.CallOption) (*OfflineResponse, error)
	Pause(ctx context.Context, in *PauseRequest, opts ...grpc.CallOption) (*PauseResponse, error)
	WaitingChannel(ctx context.Context, in *WaitingChannelRequest, opts ...grpc.CallOption) (*WaitingChannelResponse, error)
}

type agentServiceClient struct {
	cc *grpc.ClientConn
}

func NewAgentServiceClient(cc *grpc.ClientConn) AgentServiceClient {
	return &agentServiceClient{cc}
}

func (c *agentServiceClient) Online(ctx context.Context, in *OnlineRequest, opts ...grpc.CallOption) (*OnlineResponse, error) {
	out := new(OnlineResponse)
	err := c.cc.Invoke(ctx, "/cc.AgentService/Online", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) Offline(ctx context.Context, in *OfflineRequest, opts ...grpc.CallOption) (*OfflineResponse, error) {
	out := new(OfflineResponse)
	err := c.cc.Invoke(ctx, "/cc.AgentService/Offline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) Pause(ctx context.Context, in *PauseRequest, opts ...grpc.CallOption) (*PauseResponse, error) {
	out := new(PauseResponse)
	err := c.cc.Invoke(ctx, "/cc.AgentService/Pause", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) WaitingChannel(ctx context.Context, in *WaitingChannelRequest, opts ...grpc.CallOption) (*WaitingChannelResponse, error) {
	out := new(WaitingChannelResponse)
	err := c.cc.Invoke(ctx, "/cc.AgentService/WaitingChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServiceServer is the server API for AgentService service.
type AgentServiceServer interface {
	Online(context.Context, *OnlineRequest) (*OnlineResponse, error)
	Offline(context.Context, *OfflineRequest) (*OfflineResponse, error)
	Pause(context.Context, *PauseRequest) (*PauseResponse, error)
	WaitingChannel(context.Context, *WaitingChannelRequest) (*WaitingChannelResponse, error)
}

// UnimplementedAgentServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAgentServiceServer struct {
}

func (*UnimplementedAgentServiceServer) Online(ctx context.Context, req *OnlineRequest) (*OnlineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Online not implemented")
}
func (*UnimplementedAgentServiceServer) Offline(ctx context.Context, req *OfflineRequest) (*OfflineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Offline not implemented")
}
func (*UnimplementedAgentServiceServer) Pause(ctx context.Context, req *PauseRequest) (*PauseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pause not implemented")
}
func (*UnimplementedAgentServiceServer) WaitingChannel(ctx context.Context, req *WaitingChannelRequest) (*WaitingChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WaitingChannel not implemented")
}

func RegisterAgentServiceServer(s *grpc.Server, srv AgentServiceServer) {
	s.RegisterService(&_AgentService_serviceDesc, srv)
}

func _AgentService_Online_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OnlineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).Online(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cc.AgentService/Online",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).Online(ctx, req.(*OnlineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_Offline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OfflineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).Offline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cc.AgentService/Offline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).Offline(ctx, req.(*OfflineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_Pause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PauseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).Pause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cc.AgentService/Pause",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).Pause(ctx, req.(*PauseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_WaitingChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WaitingChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).WaitingChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cc.AgentService/WaitingChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).WaitingChannel(ctx, req.(*WaitingChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AgentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cc.AgentService",
	HandlerType: (*AgentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Online",
			Handler:    _AgentService_Online_Handler,
		},
		{
			MethodName: "Offline",
			Handler:    _AgentService_Offline_Handler,
		},
		{
			MethodName: "Pause",
			Handler:    _AgentService_Pause_Handler,
		},
		{
			MethodName: "WaitingChannel",
			Handler:    _AgentService_WaitingChannel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cc_agent.proto",
}