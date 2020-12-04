// Code generated by protoc-gen-go. DO NOT EDIT.
// source: storage/file.proto

package storage

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

type UploadStatusCode int32

const (
	UploadStatusCode_Unknown UploadStatusCode = 0
	UploadStatusCode_Ok      UploadStatusCode = 1
	UploadStatusCode_Failed  UploadStatusCode = 2
)

var UploadStatusCode_name = map[int32]string{
	0: "Unknown",
	1: "Ok",
	2: "Failed",
}

var UploadStatusCode_value = map[string]int32{
	"Unknown": 0,
	"Ok":      1,
	"Failed":  2,
}

func (x UploadStatusCode) String() string {
	return proto.EnumName(UploadStatusCode_name, int32(x))
}

func (UploadStatusCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0090dcf2f9a0ac7d, []int{0}
}

type UploadFileUrlRequest struct {
	DomainId             int64    `protobuf:"varint,1,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	Uuid                 string   `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Url                  string   `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadFileUrlRequest) Reset()         { *m = UploadFileUrlRequest{} }
func (m *UploadFileUrlRequest) String() string { return proto.CompactTextString(m) }
func (*UploadFileUrlRequest) ProtoMessage()    {}
func (*UploadFileUrlRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0090dcf2f9a0ac7d, []int{0}
}

func (m *UploadFileUrlRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadFileUrlRequest.Unmarshal(m, b)
}
func (m *UploadFileUrlRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadFileUrlRequest.Marshal(b, m, deterministic)
}
func (m *UploadFileUrlRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadFileUrlRequest.Merge(m, src)
}
func (m *UploadFileUrlRequest) XXX_Size() int {
	return xxx_messageInfo_UploadFileUrlRequest.Size(m)
}
func (m *UploadFileUrlRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadFileUrlRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadFileUrlRequest proto.InternalMessageInfo

func (m *UploadFileUrlRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

func (m *UploadFileUrlRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *UploadFileUrlRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UploadFileUrlRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type UploadFileUrlResponse struct {
	FileId               int64            `protobuf:"varint,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	FileUrl              string           `protobuf:"bytes,2,opt,name=file_url,json=fileUrl,proto3" json:"file_url,omitempty"`
	MimeType             string           `protobuf:"bytes,4,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	Size                 int64            `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
	Code                 UploadStatusCode `protobuf:"varint,6,opt,name=code,proto3,enum=storage.UploadStatusCode" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *UploadFileUrlResponse) Reset()         { *m = UploadFileUrlResponse{} }
func (m *UploadFileUrlResponse) String() string { return proto.CompactTextString(m) }
func (*UploadFileUrlResponse) ProtoMessage()    {}
func (*UploadFileUrlResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0090dcf2f9a0ac7d, []int{1}
}

func (m *UploadFileUrlResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadFileUrlResponse.Unmarshal(m, b)
}
func (m *UploadFileUrlResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadFileUrlResponse.Marshal(b, m, deterministic)
}
func (m *UploadFileUrlResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadFileUrlResponse.Merge(m, src)
}
func (m *UploadFileUrlResponse) XXX_Size() int {
	return xxx_messageInfo_UploadFileUrlResponse.Size(m)
}
func (m *UploadFileUrlResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadFileUrlResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadFileUrlResponse proto.InternalMessageInfo

func (m *UploadFileUrlResponse) GetFileId() int64 {
	if m != nil {
		return m.FileId
	}
	return 0
}

func (m *UploadFileUrlResponse) GetFileUrl() string {
	if m != nil {
		return m.FileUrl
	}
	return ""
}

func (m *UploadFileUrlResponse) GetMimeType() string {
	if m != nil {
		return m.MimeType
	}
	return ""
}

func (m *UploadFileUrlResponse) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *UploadFileUrlResponse) GetCode() UploadStatusCode {
	if m != nil {
		return m.Code
	}
	return UploadStatusCode_Unknown
}

type UploadFileRequest struct {
	// Types that are valid to be assigned to Data:
	//	*UploadFileRequest_Metadata_
	//	*UploadFileRequest_Chunk
	Data                 isUploadFileRequest_Data `protobuf_oneof:"data"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *UploadFileRequest) Reset()         { *m = UploadFileRequest{} }
func (m *UploadFileRequest) String() string { return proto.CompactTextString(m) }
func (*UploadFileRequest) ProtoMessage()    {}
func (*UploadFileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0090dcf2f9a0ac7d, []int{2}
}

func (m *UploadFileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadFileRequest.Unmarshal(m, b)
}
func (m *UploadFileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadFileRequest.Marshal(b, m, deterministic)
}
func (m *UploadFileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadFileRequest.Merge(m, src)
}
func (m *UploadFileRequest) XXX_Size() int {
	return xxx_messageInfo_UploadFileRequest.Size(m)
}
func (m *UploadFileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadFileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadFileRequest proto.InternalMessageInfo

type isUploadFileRequest_Data interface {
	isUploadFileRequest_Data()
}

type UploadFileRequest_Metadata_ struct {
	Metadata *UploadFileRequest_Metadata `protobuf:"bytes,1,opt,name=metadata,proto3,oneof"`
}

type UploadFileRequest_Chunk struct {
	Chunk []byte `protobuf:"bytes,2,opt,name=chunk,proto3,oneof"`
}

func (*UploadFileRequest_Metadata_) isUploadFileRequest_Data() {}

func (*UploadFileRequest_Chunk) isUploadFileRequest_Data() {}

func (m *UploadFileRequest) GetData() isUploadFileRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *UploadFileRequest) GetMetadata() *UploadFileRequest_Metadata {
	if x, ok := m.GetData().(*UploadFileRequest_Metadata_); ok {
		return x.Metadata
	}
	return nil
}

func (m *UploadFileRequest) GetChunk() []byte {
	if x, ok := m.GetData().(*UploadFileRequest_Chunk); ok {
		return x.Chunk
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*UploadFileRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*UploadFileRequest_Metadata_)(nil),
		(*UploadFileRequest_Chunk)(nil),
	}
}

type UploadFileRequest_Metadata struct {
	DomainId             int64    `protobuf:"varint,1,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	MimeType             string   `protobuf:"bytes,3,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	Uuid                 string   `protobuf:"bytes,4,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadFileRequest_Metadata) Reset()         { *m = UploadFileRequest_Metadata{} }
func (m *UploadFileRequest_Metadata) String() string { return proto.CompactTextString(m) }
func (*UploadFileRequest_Metadata) ProtoMessage()    {}
func (*UploadFileRequest_Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_0090dcf2f9a0ac7d, []int{2, 0}
}

func (m *UploadFileRequest_Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadFileRequest_Metadata.Unmarshal(m, b)
}
func (m *UploadFileRequest_Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadFileRequest_Metadata.Marshal(b, m, deterministic)
}
func (m *UploadFileRequest_Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadFileRequest_Metadata.Merge(m, src)
}
func (m *UploadFileRequest_Metadata) XXX_Size() int {
	return xxx_messageInfo_UploadFileRequest_Metadata.Size(m)
}
func (m *UploadFileRequest_Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadFileRequest_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_UploadFileRequest_Metadata proto.InternalMessageInfo

func (m *UploadFileRequest_Metadata) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

func (m *UploadFileRequest_Metadata) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UploadFileRequest_Metadata) GetMimeType() string {
	if m != nil {
		return m.MimeType
	}
	return ""
}

func (m *UploadFileRequest_Metadata) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

type UploadFileResponse struct {
	FileId               int64            `protobuf:"varint,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	FileUrl              string           `protobuf:"bytes,2,opt,name=file_url,json=fileUrl,proto3" json:"file_url,omitempty"`
	Size                 int64            `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Code                 UploadStatusCode `protobuf:"varint,4,opt,name=code,proto3,enum=storage.UploadStatusCode" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *UploadFileResponse) Reset()         { *m = UploadFileResponse{} }
func (m *UploadFileResponse) String() string { return proto.CompactTextString(m) }
func (*UploadFileResponse) ProtoMessage()    {}
func (*UploadFileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0090dcf2f9a0ac7d, []int{3}
}

func (m *UploadFileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadFileResponse.Unmarshal(m, b)
}
func (m *UploadFileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadFileResponse.Marshal(b, m, deterministic)
}
func (m *UploadFileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadFileResponse.Merge(m, src)
}
func (m *UploadFileResponse) XXX_Size() int {
	return xxx_messageInfo_UploadFileResponse.Size(m)
}
func (m *UploadFileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadFileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadFileResponse proto.InternalMessageInfo

func (m *UploadFileResponse) GetFileId() int64 {
	if m != nil {
		return m.FileId
	}
	return 0
}

func (m *UploadFileResponse) GetFileUrl() string {
	if m != nil {
		return m.FileUrl
	}
	return ""
}

func (m *UploadFileResponse) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *UploadFileResponse) GetCode() UploadStatusCode {
	if m != nil {
		return m.Code
	}
	return UploadStatusCode_Unknown
}

func init() {
	proto.RegisterEnum("storage.UploadStatusCode", UploadStatusCode_name, UploadStatusCode_value)
	proto.RegisterType((*UploadFileUrlRequest)(nil), "storage.UploadFileUrlRequest")
	proto.RegisterType((*UploadFileUrlResponse)(nil), "storage.UploadFileUrlResponse")
	proto.RegisterType((*UploadFileRequest)(nil), "storage.UploadFileRequest")
	proto.RegisterType((*UploadFileRequest_Metadata)(nil), "storage.UploadFileRequest.Metadata")
	proto.RegisterType((*UploadFileResponse)(nil), "storage.UploadFileResponse")
}

func init() { proto.RegisterFile("storage/file.proto", fileDescriptor_0090dcf2f9a0ac7d) }

var fileDescriptor_0090dcf2f9a0ac7d = []byte{
	// 477 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x5d, 0x6e, 0xd3, 0x40,
	0x10, 0xce, 0x26, 0xae, 0x93, 0x4e, 0x00, 0x85, 0x15, 0x3f, 0x6e, 0x0a, 0x28, 0xb8, 0x2f, 0x11,
	0x12, 0xb6, 0x94, 0x9e, 0x80, 0x22, 0x55, 0xcd, 0x03, 0x02, 0xb9, 0xe4, 0x85, 0x97, 0x6a, 0x93,
	0x1d, 0xd2, 0x55, 0xd6, 0xbb, 0xc6, 0x5e, 0x53, 0x95, 0x2b, 0x70, 0x13, 0xb8, 0x16, 0x07, 0x41,
	0xbb, 0x76, 0x70, 0x1a, 0x52, 0x81, 0xd4, 0xb7, 0x99, 0x6f, 0xbe, 0x99, 0xf9, 0xfc, 0x8d, 0x17,
	0x68, 0x61, 0x74, 0xce, 0x96, 0x18, 0x7f, 0x16, 0x12, 0xa3, 0x2c, 0xd7, 0x46, 0xd3, 0x6e, 0x8d,
	0x0d, 0x9f, 0x2d, 0xb5, 0x5e, 0x4a, 0x8c, 0x59, 0x26, 0x62, 0xa6, 0x94, 0x36, 0xcc, 0x08, 0xad,
	0x8a, 0x8a, 0x16, 0xa6, 0xf0, 0x68, 0x96, 0x49, 0xcd, 0xf8, 0xa9, 0x90, 0x38, 0xcb, 0x65, 0x82,
	0x5f, 0x4a, 0x2c, 0x0c, 0x3d, 0x84, 0x7d, 0xae, 0x53, 0x26, 0xd4, 0x85, 0xe0, 0x01, 0x19, 0x91,
	0x71, 0x27, 0xe9, 0x55, 0xc0, 0x94, 0x53, 0x0a, 0x5e, 0x59, 0x0a, 0x1e, 0xb4, 0x47, 0x64, 0xbc,
	0x9f, 0xb8, 0xd8, 0x62, 0x8a, 0xa5, 0x18, 0x74, 0x2a, 0xcc, 0xc6, 0x74, 0x00, 0x9d, 0x32, 0x97,
	0x81, 0xe7, 0x20, 0x1b, 0x86, 0x3f, 0x09, 0x3c, 0xde, 0xda, 0x57, 0x64, 0x5a, 0x15, 0x48, 0x9f,
	0x42, 0xd7, 0xaa, 0x6f, 0xd6, 0xf9, 0x36, 0x9d, 0x72, 0x7a, 0x00, 0x3d, 0x57, 0xb0, 0x93, 0xaa,
	0x85, 0x8e, 0x38, 0xcb, 0xa5, 0x15, 0x99, 0x8a, 0x14, 0x2f, 0xcc, 0x75, 0x86, 0xf5, 0x96, 0x9e,
	0x05, 0x3e, 0x5e, 0x67, 0x68, 0x05, 0x15, 0xe2, 0x1b, 0x06, 0x7b, 0x6e, 0x9a, 0x8b, 0xe9, 0x6b,
	0xf0, 0x16, 0x9a, 0x63, 0xe0, 0x8f, 0xc8, 0xf8, 0xc1, 0xe4, 0x20, 0xaa, 0x3d, 0x8a, 0x2a, 0x49,
	0xe7, 0x86, 0x99, 0xb2, 0x78, 0xab, 0x39, 0x26, 0x8e, 0x16, 0xfe, 0x22, 0xf0, 0xb0, 0x51, 0xbb,
	0xb6, 0xe6, 0x0d, 0xf4, 0x52, 0x34, 0x8c, 0x33, 0xc3, 0x9c, 0xd4, 0xfe, 0xe4, 0x68, 0x6b, 0xd0,
	0x06, 0x3b, 0x7a, 0x57, 0x53, 0xcf, 0x5a, 0xc9, 0x9f, 0x36, 0xfa, 0x04, 0xf6, 0x16, 0x97, 0xa5,
	0x5a, 0xb9, 0x0f, 0xba, 0x77, 0xd6, 0x4a, 0xaa, 0x74, 0x28, 0xa1, 0xb7, 0xe6, 0xff, 0xf3, 0x02,
	0xce, 0xed, 0xf6, 0x86, 0xdb, 0x37, 0xdc, 0xe8, 0xfc, 0xed, 0x86, 0x3b, 0x99, 0xd7, 0x9c, 0xec,
	0xc4, 0x07, 0xcf, 0x6e, 0x0a, 0xbf, 0x13, 0xa0, 0x9b, 0xc2, 0xef, 0x70, 0x91, 0xb5, 0xe9, 0x9d,
	0x1d, 0xa6, 0x7b, 0xff, 0x65, 0xfa, 0xab, 0x63, 0x18, 0x6c, 0x57, 0x68, 0x1f, 0xba, 0x33, 0xb5,
	0x52, 0xfa, 0x4a, 0x0d, 0x5a, 0xd4, 0x87, 0xf6, 0xfb, 0xd5, 0x80, 0x50, 0x00, 0xff, 0x94, 0x09,
	0x89, 0x7c, 0xd0, 0x9e, 0xfc, 0x20, 0xd0, 0xb7, 0xe2, 0xcf, 0x31, 0xff, 0x2a, 0x16, 0x48, 0xa7,
	0x00, 0xcd, 0x17, 0xd1, 0xe1, 0xed, 0xf7, 0x19, 0x1e, 0xee, 0xac, 0x55, 0x16, 0x84, 0xad, 0x31,
	0xa1, 0x1f, 0xe0, 0xfe, 0x8d, 0x3f, 0x96, 0x3e, 0xdf, 0xd1, 0xd1, 0xbc, 0x9c, 0xe1, 0x8b, 0xdb,
	0xca, 0xeb, 0x99, 0x27, 0x47, 0x9f, 0x5e, 0x2e, 0x85, 0xb9, 0x2c, 0xe7, 0xd1, 0x42, 0xa7, 0xf1,
	0x15, 0xce, 0x85, 0x41, 0x19, 0xbb, 0xf7, 0x58, 0xc4, 0x75, 0xf3, 0xdc, 0x77, 0xf9, 0xf1, 0xef,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x36, 0xdb, 0x63, 0x56, 0xdc, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FileServiceClient is the client API for FileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FileServiceClient interface {
	UploadFile(ctx context.Context, opts ...grpc.CallOption) (FileService_UploadFileClient, error)
	UploadFileUrl(ctx context.Context, in *UploadFileUrlRequest, opts ...grpc.CallOption) (*UploadFileUrlResponse, error)
}

type fileServiceClient struct {
	cc *grpc.ClientConn
}

func NewFileServiceClient(cc *grpc.ClientConn) FileServiceClient {
	return &fileServiceClient{cc}
}

func (c *fileServiceClient) UploadFile(ctx context.Context, opts ...grpc.CallOption) (FileService_UploadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FileService_serviceDesc.Streams[0], "/storage.FileService/UploadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileServiceUploadFileClient{stream}
	return x, nil
}

type FileService_UploadFileClient interface {
	Send(*UploadFileRequest) error
	CloseAndRecv() (*UploadFileResponse, error)
	grpc.ClientStream
}

type fileServiceUploadFileClient struct {
	grpc.ClientStream
}

func (x *fileServiceUploadFileClient) Send(m *UploadFileRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileServiceUploadFileClient) CloseAndRecv() (*UploadFileResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadFileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileServiceClient) UploadFileUrl(ctx context.Context, in *UploadFileUrlRequest, opts ...grpc.CallOption) (*UploadFileUrlResponse, error) {
	out := new(UploadFileUrlResponse)
	err := c.cc.Invoke(ctx, "/storage.FileService/UploadFileUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileServiceServer is the server API for FileService service.
type FileServiceServer interface {
	UploadFile(FileService_UploadFileServer) error
	UploadFileUrl(context.Context, *UploadFileUrlRequest) (*UploadFileUrlResponse, error)
}

// UnimplementedFileServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFileServiceServer struct {
}

func (*UnimplementedFileServiceServer) UploadFile(srv FileService_UploadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (*UnimplementedFileServiceServer) UploadFileUrl(ctx context.Context, req *UploadFileUrlRequest) (*UploadFileUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadFileUrl not implemented")
}

func RegisterFileServiceServer(s *grpc.Server, srv FileServiceServer) {
	s.RegisterService(&_FileService_serviceDesc, srv)
}

func _FileService_UploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileServiceServer).UploadFile(&fileServiceUploadFileServer{stream})
}

type FileService_UploadFileServer interface {
	SendAndClose(*UploadFileResponse) error
	Recv() (*UploadFileRequest, error)
	grpc.ServerStream
}

type fileServiceUploadFileServer struct {
	grpc.ServerStream
}

func (x *fileServiceUploadFileServer) SendAndClose(m *UploadFileResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileServiceUploadFileServer) Recv() (*UploadFileRequest, error) {
	m := new(UploadFileRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _FileService_UploadFileUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadFileUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).UploadFileUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.FileService/UploadFileUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).UploadFileUrl(ctx, req.(*UploadFileUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FileService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "storage.FileService",
	HandlerType: (*FileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadFileUrl",
			Handler:    _FileService_UploadFileUrl_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadFile",
			Handler:       _FileService_UploadFile_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "storage/file.proto",
}
