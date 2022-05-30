// Code generated by protoc-gen-go. DO NOT EDIT.
// source: storage/file_transcript.proto

package storage

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	engine "github.com/webitel/protos/engine"
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

type TranscriptPhrase struct {
	StartSec             float32  `protobuf:"fixed32,1,opt,name=start_sec,json=startSec,proto3" json:"start_sec,omitempty"`
	EndSec               float32  `protobuf:"fixed32,2,opt,name=end_sec,json=endSec,proto3" json:"end_sec,omitempty"`
	Channel              uint32   `protobuf:"varint,3,opt,name=channel,proto3" json:"channel,omitempty"`
	Phrase               string   `protobuf:"bytes,4,opt,name=phrase,proto3" json:"phrase,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TranscriptPhrase) Reset()         { *m = TranscriptPhrase{} }
func (m *TranscriptPhrase) String() string { return proto.CompactTextString(m) }
func (*TranscriptPhrase) ProtoMessage()    {}
func (*TranscriptPhrase) Descriptor() ([]byte, []int) {
	return fileDescriptor_31885769db7a6a0a, []int{0}
}

func (m *TranscriptPhrase) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TranscriptPhrase.Unmarshal(m, b)
}
func (m *TranscriptPhrase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TranscriptPhrase.Marshal(b, m, deterministic)
}
func (m *TranscriptPhrase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TranscriptPhrase.Merge(m, src)
}
func (m *TranscriptPhrase) XXX_Size() int {
	return xxx_messageInfo_TranscriptPhrase.Size(m)
}
func (m *TranscriptPhrase) XXX_DiscardUnknown() {
	xxx_messageInfo_TranscriptPhrase.DiscardUnknown(m)
}

var xxx_messageInfo_TranscriptPhrase proto.InternalMessageInfo

func (m *TranscriptPhrase) GetStartSec() float32 {
	if m != nil {
		return m.StartSec
	}
	return 0
}

func (m *TranscriptPhrase) GetEndSec() float32 {
	if m != nil {
		return m.EndSec
	}
	return 0
}

func (m *TranscriptPhrase) GetChannel() uint32 {
	if m != nil {
		return m.Channel
	}
	return 0
}

func (m *TranscriptPhrase) GetPhrase() string {
	if m != nil {
		return m.Phrase
	}
	return ""
}

type GetFileTranscriptPhrasesRequest struct {
	Page                 int32    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size                 int32    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Id                   int64    `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFileTranscriptPhrasesRequest) Reset()         { *m = GetFileTranscriptPhrasesRequest{} }
func (m *GetFileTranscriptPhrasesRequest) String() string { return proto.CompactTextString(m) }
func (*GetFileTranscriptPhrasesRequest) ProtoMessage()    {}
func (*GetFileTranscriptPhrasesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_31885769db7a6a0a, []int{1}
}

func (m *GetFileTranscriptPhrasesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFileTranscriptPhrasesRequest.Unmarshal(m, b)
}
func (m *GetFileTranscriptPhrasesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFileTranscriptPhrasesRequest.Marshal(b, m, deterministic)
}
func (m *GetFileTranscriptPhrasesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFileTranscriptPhrasesRequest.Merge(m, src)
}
func (m *GetFileTranscriptPhrasesRequest) XXX_Size() int {
	return xxx_messageInfo_GetFileTranscriptPhrasesRequest.Size(m)
}
func (m *GetFileTranscriptPhrasesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFileTranscriptPhrasesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFileTranscriptPhrasesRequest proto.InternalMessageInfo

func (m *GetFileTranscriptPhrasesRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *GetFileTranscriptPhrasesRequest) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *GetFileTranscriptPhrasesRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ListPhrases struct {
	Next                 bool                `protobuf:"varint,1,opt,name=next,proto3" json:"next,omitempty"`
	Items                []*TranscriptPhrase `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ListPhrases) Reset()         { *m = ListPhrases{} }
func (m *ListPhrases) String() string { return proto.CompactTextString(m) }
func (*ListPhrases) ProtoMessage()    {}
func (*ListPhrases) Descriptor() ([]byte, []int) {
	return fileDescriptor_31885769db7a6a0a, []int{2}
}

func (m *ListPhrases) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPhrases.Unmarshal(m, b)
}
func (m *ListPhrases) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPhrases.Marshal(b, m, deterministic)
}
func (m *ListPhrases) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPhrases.Merge(m, src)
}
func (m *ListPhrases) XXX_Size() int {
	return xxx_messageInfo_ListPhrases.Size(m)
}
func (m *ListPhrases) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPhrases.DiscardUnknown(m)
}

var xxx_messageInfo_ListPhrases proto.InternalMessageInfo

func (m *ListPhrases) GetNext() bool {
	if m != nil {
		return m.Next
	}
	return false
}

func (m *ListPhrases) GetItems() []*TranscriptPhrase {
	if m != nil {
		return m.Items
	}
	return nil
}

type StartFileTranscriptRequest struct {
	FileId               []int64        `protobuf:"varint,1,rep,packed,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	Locale               string         `protobuf:"bytes,2,opt,name=locale,proto3" json:"locale,omitempty"`
	Profile              *engine.Lookup `protobuf:"bytes,3,opt,name=profile,proto3" json:"profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *StartFileTranscriptRequest) Reset()         { *m = StartFileTranscriptRequest{} }
func (m *StartFileTranscriptRequest) String() string { return proto.CompactTextString(m) }
func (*StartFileTranscriptRequest) ProtoMessage()    {}
func (*StartFileTranscriptRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_31885769db7a6a0a, []int{3}
}

func (m *StartFileTranscriptRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartFileTranscriptRequest.Unmarshal(m, b)
}
func (m *StartFileTranscriptRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartFileTranscriptRequest.Marshal(b, m, deterministic)
}
func (m *StartFileTranscriptRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartFileTranscriptRequest.Merge(m, src)
}
func (m *StartFileTranscriptRequest) XXX_Size() int {
	return xxx_messageInfo_StartFileTranscriptRequest.Size(m)
}
func (m *StartFileTranscriptRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StartFileTranscriptRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StartFileTranscriptRequest proto.InternalMessageInfo

func (m *StartFileTranscriptRequest) GetFileId() []int64 {
	if m != nil {
		return m.FileId
	}
	return nil
}

func (m *StartFileTranscriptRequest) GetLocale() string {
	if m != nil {
		return m.Locale
	}
	return ""
}

func (m *StartFileTranscriptRequest) GetProfile() *engine.Lookup {
	if m != nil {
		return m.Profile
	}
	return nil
}

type StartFileTranscriptResponse struct {
	Items                []*StartFileTranscriptResponse_TranscriptJob `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                     `json:"-"`
	XXX_unrecognized     []byte                                       `json:"-"`
	XXX_sizecache        int32                                        `json:"-"`
}

func (m *StartFileTranscriptResponse) Reset()         { *m = StartFileTranscriptResponse{} }
func (m *StartFileTranscriptResponse) String() string { return proto.CompactTextString(m) }
func (*StartFileTranscriptResponse) ProtoMessage()    {}
func (*StartFileTranscriptResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_31885769db7a6a0a, []int{4}
}

func (m *StartFileTranscriptResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartFileTranscriptResponse.Unmarshal(m, b)
}
func (m *StartFileTranscriptResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartFileTranscriptResponse.Marshal(b, m, deterministic)
}
func (m *StartFileTranscriptResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartFileTranscriptResponse.Merge(m, src)
}
func (m *StartFileTranscriptResponse) XXX_Size() int {
	return xxx_messageInfo_StartFileTranscriptResponse.Size(m)
}
func (m *StartFileTranscriptResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StartFileTranscriptResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StartFileTranscriptResponse proto.InternalMessageInfo

func (m *StartFileTranscriptResponse) GetItems() []*StartFileTranscriptResponse_TranscriptJob {
	if m != nil {
		return m.Items
	}
	return nil
}

type StartFileTranscriptResponse_TranscriptJob struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FileId               int64    `protobuf:"varint,2,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	CreatedAt            int64    `protobuf:"varint,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Action               string   `protobuf:"bytes,4,opt,name=action,proto3" json:"action,omitempty"`
	State                string   `protobuf:"bytes,5,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartFileTranscriptResponse_TranscriptJob) Reset() {
	*m = StartFileTranscriptResponse_TranscriptJob{}
}
func (m *StartFileTranscriptResponse_TranscriptJob) String() string { return proto.CompactTextString(m) }
func (*StartFileTranscriptResponse_TranscriptJob) ProtoMessage()    {}
func (*StartFileTranscriptResponse_TranscriptJob) Descriptor() ([]byte, []int) {
	return fileDescriptor_31885769db7a6a0a, []int{4, 0}
}

func (m *StartFileTranscriptResponse_TranscriptJob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartFileTranscriptResponse_TranscriptJob.Unmarshal(m, b)
}
func (m *StartFileTranscriptResponse_TranscriptJob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartFileTranscriptResponse_TranscriptJob.Marshal(b, m, deterministic)
}
func (m *StartFileTranscriptResponse_TranscriptJob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartFileTranscriptResponse_TranscriptJob.Merge(m, src)
}
func (m *StartFileTranscriptResponse_TranscriptJob) XXX_Size() int {
	return xxx_messageInfo_StartFileTranscriptResponse_TranscriptJob.Size(m)
}
func (m *StartFileTranscriptResponse_TranscriptJob) XXX_DiscardUnknown() {
	xxx_messageInfo_StartFileTranscriptResponse_TranscriptJob.DiscardUnknown(m)
}

var xxx_messageInfo_StartFileTranscriptResponse_TranscriptJob proto.InternalMessageInfo

func (m *StartFileTranscriptResponse_TranscriptJob) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *StartFileTranscriptResponse_TranscriptJob) GetFileId() int64 {
	if m != nil {
		return m.FileId
	}
	return 0
}

func (m *StartFileTranscriptResponse_TranscriptJob) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *StartFileTranscriptResponse_TranscriptJob) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *StartFileTranscriptResponse_TranscriptJob) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func init() {
	proto.RegisterType((*TranscriptPhrase)(nil), "storage.TranscriptPhrase")
	proto.RegisterType((*GetFileTranscriptPhrasesRequest)(nil), "storage.GetFileTranscriptPhrasesRequest")
	proto.RegisterType((*ListPhrases)(nil), "storage.ListPhrases")
	proto.RegisterType((*StartFileTranscriptRequest)(nil), "storage.StartFileTranscriptRequest")
	proto.RegisterType((*StartFileTranscriptResponse)(nil), "storage.StartFileTranscriptResponse")
	proto.RegisterType((*StartFileTranscriptResponse_TranscriptJob)(nil), "storage.StartFileTranscriptResponse.TranscriptJob")
}

func init() { proto.RegisterFile("storage/file_transcript.proto", fileDescriptor_31885769db7a6a0a) }

var fileDescriptor_31885769db7a6a0a = []byte{
	// 556 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcd, 0x8e, 0xd3, 0x3c,
	0x14, 0x95, 0x93, 0x69, 0x3b, 0x75, 0x35, 0xa3, 0x4f, 0x56, 0x3f, 0x26, 0x64, 0xa8, 0x28, 0x29,
	0x88, 0x08, 0x89, 0x44, 0x2a, 0x3b, 0x76, 0x80, 0xc4, 0x9f, 0x66, 0x81, 0x5c, 0x36, 0xb0, 0xa9,
	0xdc, 0xe4, 0x92, 0x5a, 0x64, 0xec, 0x10, 0xbb, 0x0c, 0x1a, 0xc4, 0x06, 0x09, 0x24, 0xd6, 0xbc,
	0x06, 0x6f, 0xc3, 0x2b, 0xb0, 0xe7, 0x15, 0x50, 0x6e, 0x92, 0xd2, 0x8e, 0x98, 0x81, 0x9d, 0xef,
	0x3d, 0xf6, 0xf1, 0x39, 0xc7, 0xd7, 0x74, 0x64, 0xac, 0x2e, 0x45, 0x06, 0xf1, 0x2b, 0x99, 0xc3,
	0xdc, 0x96, 0x42, 0x99, 0xa4, 0x94, 0x85, 0x8d, 0x8a, 0x52, 0x5b, 0xcd, 0x7a, 0x0d, 0xec, 0x33,
	0x50, 0x99, 0x54, 0x10, 0x27, 0x5a, 0x99, 0x06, 0xf4, 0xaf, 0x64, 0x5a, 0x67, 0x39, 0xc4, 0xa2,
	0x90, 0xb1, 0x50, 0x4a, 0x5b, 0x61, 0xa5, 0x56, 0xa6, 0x46, 0x83, 0x53, 0xfa, 0xdf, 0xf3, 0x35,
	0xdd, 0xb3, 0x65, 0x29, 0x0c, 0xb0, 0x43, 0xda, 0x37, 0x56, 0x94, 0x76, 0x6e, 0x20, 0xf1, 0xc8,
	0x98, 0x84, 0x0e, 0xdf, 0xc5, 0xc6, 0x0c, 0x12, 0x76, 0x40, 0x7b, 0xa0, 0x52, 0x84, 0x1c, 0x84,
	0xba, 0xa0, 0xd2, 0x0a, 0xf0, 0x68, 0x2f, 0x59, 0x0a, 0xa5, 0x20, 0xf7, 0xdc, 0x31, 0x09, 0xf7,
	0x78, 0x5b, 0xb2, 0x4b, 0xb4, 0x5b, 0x20, 0xb3, 0xb7, 0x33, 0x26, 0x61, 0x9f, 0x37, 0x55, 0xf0,
	0x82, 0x5e, 0x7d, 0x04, 0xf6, 0xa1, 0xcc, 0xe1, 0xac, 0x04, 0xc3, 0xe1, 0xcd, 0x0a, 0x8c, 0x65,
	0x8c, 0xee, 0x14, 0x22, 0x03, 0x54, 0xd1, 0xe1, 0xb8, 0xae, 0x7a, 0x46, 0x9e, 0x02, 0x5e, 0xdf,
	0xe1, 0xb8, 0x66, 0xfb, 0xd4, 0x91, 0x29, 0xd2, 0xbb, 0xdc, 0x91, 0x69, 0xc0, 0xe9, 0xe0, 0x48,
	0x9a, 0x96, 0xad, 0x3a, 0xa2, 0xe0, 0x9d, 0x45, 0x9a, 0x5d, 0x8e, 0x6b, 0x16, 0xd3, 0x8e, 0xb4,
	0x70, 0x6c, 0x3c, 0x67, 0xec, 0x86, 0x83, 0xe9, 0xe5, 0xa8, 0x09, 0x31, 0x3a, 0x2b, 0x86, 0xd7,
	0xfb, 0x82, 0x13, 0xea, 0xcf, 0xaa, 0x14, 0xb6, 0x05, 0xb7, 0x4a, 0x0f, 0x68, 0x0f, 0x1f, 0x47,
	0xa6, 0x1e, 0x19, 0xbb, 0xa1, 0xcb, 0xbb, 0x55, 0xf9, 0x24, 0xad, 0xdc, 0xe7, 0x3a, 0x11, 0x79,
	0x2d, 0xb8, 0xcf, 0x9b, 0x8a, 0x85, 0xb4, 0x57, 0x94, 0xba, 0xda, 0x84, 0x79, 0x0d, 0xa6, 0xfb,
	0x51, 0xfd, 0x7a, 0xd1, 0x91, 0xd6, 0xaf, 0x57, 0x05, 0x6f, 0xe1, 0xe0, 0x27, 0xa1, 0x87, 0x7f,
	0xbc, 0xd9, 0x14, 0x5a, 0x19, 0x60, 0x8f, 0x5b, 0x27, 0x04, 0x9d, 0x4c, 0xd7, 0x4e, 0x2e, 0x38,
	0xb4, 0xe1, 0xf2, 0xa9, 0x5e, 0x34, 0x16, 0xfd, 0x4f, 0x84, 0xee, 0x6d, 0x01, 0x4d, 0xb0, 0xa4,
	0x0d, 0x76, 0xd3, 0xa6, 0x83, 0xcd, 0xd6, 0xe6, 0x88, 0xd2, 0xa4, 0x04, 0x61, 0x21, 0x9d, 0x0b,
	0x8b, 0x8e, 0x5c, 0xde, 0x6f, 0x3a, 0xf7, 0x6c, 0x95, 0x82, 0x48, 0xaa, 0xc1, 0x6b, 0x67, 0xa0,
	0xae, 0xd8, 0x90, 0x76, 0x8c, 0x15, 0x16, 0xbc, 0x0e, 0xb6, 0xeb, 0x62, 0xfa, 0xcd, 0xa1, 0xff,
	0x6f, 0xeb, 0x9e, 0x41, 0xf9, 0x56, 0x26, 0xc0, 0x3e, 0x13, 0x3a, 0x7c, 0x80, 0xac, 0xdb, 0x38,
	0x9b, 0x5c, 0xec, 0x1a, 0x1f, 0xc9, 0xbf, 0xfe, 0x2f, 0xd1, 0x04, 0x93, 0x8f, 0xdf, 0x7f, 0x7c,
	0x75, 0x46, 0x81, 0x17, 0xb7, 0xdf, 0xee, 0xf7, 0x8f, 0x9b, 0x57, 0x76, 0xef, 0x92, 0x5b, 0xec,
	0x0b, 0xa1, 0xde, 0x79, 0xd3, 0xcb, 0xc2, 0xf5, 0x3d, 0x7f, 0x19, 0x70, 0x7f, 0xb8, 0xde, 0xb9,
	0x31, 0xaf, 0xc1, 0x6d, 0x54, 0x70, 0x93, 0xdd, 0x38, 0x4f, 0x41, 0xfc, 0x5e, 0xa6, 0x1f, 0xe2,
	0xfa, 0x1f, 0x99, 0xfb, 0x93, 0x97, 0xd7, 0x32, 0x69, 0x97, 0xab, 0x45, 0x94, 0xe8, 0xe3, 0xf8,
	0x04, 0x16, 0xd2, 0x42, 0x1e, 0xe3, 0x07, 0x37, 0x2d, 0xc3, 0xa2, 0x8b, 0xf5, 0x9d, 0x5f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x62, 0xeb, 0xe0, 0xad, 0x4c, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FileTranscriptServiceClient is the client API for FileTranscriptService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FileTranscriptServiceClient interface {
	CreateFileTranscript(ctx context.Context, in *StartFileTranscriptRequest, opts ...grpc.CallOption) (*StartFileTranscriptResponse, error)
	GetFileTranscriptPhrases(ctx context.Context, in *GetFileTranscriptPhrasesRequest, opts ...grpc.CallOption) (*ListPhrases, error)
}

type fileTranscriptServiceClient struct {
	cc *grpc.ClientConn
}

func NewFileTranscriptServiceClient(cc *grpc.ClientConn) FileTranscriptServiceClient {
	return &fileTranscriptServiceClient{cc}
}

func (c *fileTranscriptServiceClient) CreateFileTranscript(ctx context.Context, in *StartFileTranscriptRequest, opts ...grpc.CallOption) (*StartFileTranscriptResponse, error) {
	out := new(StartFileTranscriptResponse)
	err := c.cc.Invoke(ctx, "/storage.FileTranscriptService/CreateFileTranscript", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileTranscriptServiceClient) GetFileTranscriptPhrases(ctx context.Context, in *GetFileTranscriptPhrasesRequest, opts ...grpc.CallOption) (*ListPhrases, error) {
	out := new(ListPhrases)
	err := c.cc.Invoke(ctx, "/storage.FileTranscriptService/GetFileTranscriptPhrases", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileTranscriptServiceServer is the server API for FileTranscriptService service.
type FileTranscriptServiceServer interface {
	CreateFileTranscript(context.Context, *StartFileTranscriptRequest) (*StartFileTranscriptResponse, error)
	GetFileTranscriptPhrases(context.Context, *GetFileTranscriptPhrasesRequest) (*ListPhrases, error)
}

// UnimplementedFileTranscriptServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFileTranscriptServiceServer struct {
}

func (*UnimplementedFileTranscriptServiceServer) CreateFileTranscript(ctx context.Context, req *StartFileTranscriptRequest) (*StartFileTranscriptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFileTranscript not implemented")
}
func (*UnimplementedFileTranscriptServiceServer) GetFileTranscriptPhrases(ctx context.Context, req *GetFileTranscriptPhrasesRequest) (*ListPhrases, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFileTranscriptPhrases not implemented")
}

func RegisterFileTranscriptServiceServer(s *grpc.Server, srv FileTranscriptServiceServer) {
	s.RegisterService(&_FileTranscriptService_serviceDesc, srv)
}

func _FileTranscriptService_CreateFileTranscript_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartFileTranscriptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileTranscriptServiceServer).CreateFileTranscript(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.FileTranscriptService/CreateFileTranscript",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileTranscriptServiceServer).CreateFileTranscript(ctx, req.(*StartFileTranscriptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileTranscriptService_GetFileTranscriptPhrases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFileTranscriptPhrasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileTranscriptServiceServer).GetFileTranscriptPhrases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.FileTranscriptService/GetFileTranscriptPhrases",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileTranscriptServiceServer).GetFileTranscriptPhrases(ctx, req.(*GetFileTranscriptPhrasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FileTranscriptService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "storage.FileTranscriptService",
	HandlerType: (*FileTranscriptServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFileTranscript",
			Handler:    _FileTranscriptService_CreateFileTranscript_Handler,
		},
		{
			MethodName: "GetFileTranscriptPhrases",
			Handler:    _FileTranscriptService_GetFileTranscriptPhrases_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "storage/file_transcript.proto",
}
