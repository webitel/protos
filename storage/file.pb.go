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

type UploadFileResponse_UploadStatusCode int32

const (
	UploadFileResponse_Unknown UploadFileResponse_UploadStatusCode = 0
	UploadFileResponse_Ok      UploadFileResponse_UploadStatusCode = 1
	UploadFileResponse_Failed  UploadFileResponse_UploadStatusCode = 2
)

var UploadFileResponse_UploadStatusCode_name = map[int32]string{
	0: "Unknown",
	1: "Ok",
	2: "Failed",
}

var UploadFileResponse_UploadStatusCode_value = map[string]int32{
	"Unknown": 0,
	"Ok":      1,
	"Failed":  2,
}

func (x UploadFileResponse_UploadStatusCode) String() string {
	return proto.EnumName(UploadFileResponse_UploadStatusCode_name, int32(x))
}

func (UploadFileResponse_UploadStatusCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0090dcf2f9a0ac7d, []int{1, 0}
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
	return fileDescriptor_0090dcf2f9a0ac7d, []int{0}
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
	return fileDescriptor_0090dcf2f9a0ac7d, []int{0, 0}
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
	FileId               int64                               `protobuf:"varint,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	Code                 UploadFileResponse_UploadStatusCode `protobuf:"varint,2,opt,name=code,proto3,enum=storage.UploadFileResponse_UploadStatusCode" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *UploadFileResponse) Reset()         { *m = UploadFileResponse{} }
func (m *UploadFileResponse) String() string { return proto.CompactTextString(m) }
func (*UploadFileResponse) ProtoMessage()    {}
func (*UploadFileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0090dcf2f9a0ac7d, []int{1}
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

func (m *UploadFileResponse) GetCode() UploadFileResponse_UploadStatusCode {
	if m != nil {
		return m.Code
	}
	return UploadFileResponse_Unknown
}

func init() {
	proto.RegisterEnum("storage.UploadFileResponse_UploadStatusCode", UploadFileResponse_UploadStatusCode_name, UploadFileResponse_UploadStatusCode_value)
	proto.RegisterType((*UploadFileRequest)(nil), "storage.UploadFileRequest")
	proto.RegisterType((*UploadFileRequest_Metadata)(nil), "storage.UploadFileRequest.Metadata")
	proto.RegisterType((*UploadFileResponse)(nil), "storage.UploadFileResponse")
}

func init() { proto.RegisterFile("storage/file.proto", fileDescriptor_0090dcf2f9a0ac7d) }

var fileDescriptor_0090dcf2f9a0ac7d = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x4d, 0x8e, 0xda, 0x40,
	0x10, 0x85, 0x6d, 0x70, 0x8c, 0x29, 0xa2, 0xc8, 0xe9, 0x45, 0x62, 0x99, 0x2c, 0x88, 0xd9, 0xb0,
	0x88, 0x6c, 0x09, 0x2e, 0x90, 0x10, 0x09, 0xc1, 0x22, 0x8a, 0x64, 0x82, 0x14, 0x65, 0x83, 0x1a,
	0x77, 0xc5, 0xb4, 0xb0, 0xbb, 0x3d, 0xb8, 0x3d, 0x88, 0xfb, 0xcc, 0xd1, 0xe6, 0x20, 0xa3, 0xb6,
	0x3d, 0xc3, 0x68, 0xfe, 0x76, 0x55, 0xaf, 0xde, 0x57, 0xaf, 0xfc, 0x03, 0xa4, 0x54, 0xf2, 0x48,
	0x53, 0x8c, 0xfe, 0xf3, 0x0c, 0xc3, 0xe2, 0x28, 0x95, 0x24, 0xbd, 0x56, 0xf3, 0xbf, 0xa4, 0x52,
	0xa6, 0x19, 0x46, 0xb4, 0xe0, 0x11, 0x15, 0x42, 0x2a, 0xaa, 0xb8, 0x14, 0x65, 0x63, 0x0b, 0x6e,
	0x4d, 0xf8, 0xb8, 0x29, 0x32, 0x49, 0xd9, 0x82, 0x67, 0x18, 0xe3, 0x55, 0x85, 0xa5, 0x22, 0x3f,
	0xc0, 0xc9, 0x51, 0x51, 0x46, 0x15, 0xf5, 0xcc, 0x91, 0x39, 0x19, 0x4c, 0xc7, 0x61, 0xbb, 0x2f,
	0x7c, 0xe6, 0x0e, 0x7f, 0xb5, 0xd6, 0xa5, 0x11, 0x3f, 0x60, 0xe4, 0x13, 0xbc, 0x4b, 0xf6, 0x95,
	0x38, 0x78, 0x9d, 0x91, 0x39, 0x79, 0xbf, 0x34, 0xe2, 0xa6, 0xf5, 0x33, 0x70, 0xee, 0xfd, 0x64,
	0x08, 0x7d, 0x26, 0x73, 0xca, 0xc5, 0x96, 0xb3, 0x3a, 0xa7, 0x1b, 0x3b, 0x8d, 0xb0, 0x62, 0x84,
	0x80, 0x25, 0x68, 0x8e, 0x35, 0xdf, 0x8f, 0xeb, 0x5a, 0x03, 0x39, 0xcf, 0x71, 0xab, 0xce, 0x05,
	0x7a, 0xdd, 0x7a, 0xe0, 0x68, 0xe1, 0xcf, 0xb9, 0x40, 0x0d, 0x54, 0x15, 0x67, 0x9e, 0xd5, 0x00,
	0xba, 0x9e, 0xdb, 0x60, 0xe9, 0xa4, 0xe0, 0xc6, 0x04, 0xf2, 0xf8, 0xf0, 0xb2, 0x90, 0xa2, 0x44,
	0xf2, 0x19, 0x7a, 0xfa, 0x95, 0x5d, 0xe2, 0x6d, 0xdd, 0xae, 0x18, 0xf9, 0x0e, 0x56, 0x22, 0x59,
	0x13, 0xfe, 0x61, 0xfa, 0xed, 0xc5, 0x87, 0x6f, 0x76, 0xb4, 0xd2, 0x5a, 0x51, 0x55, 0x95, 0x3f,
	0x25, 0xc3, 0xb8, 0x26, 0x83, 0x19, 0xb8, 0x4f, 0x27, 0x64, 0x00, 0xbd, 0x8d, 0x38, 0x08, 0x79,
	0x12, 0xae, 0x41, 0x6c, 0xe8, 0xfc, 0x3e, 0xb8, 0x26, 0x01, 0xb0, 0x17, 0x94, 0x67, 0xc8, 0xdc,
	0xce, 0xf4, 0x2f, 0x0c, 0xf4, 0xee, 0x35, 0x1e, 0xaf, 0x79, 0x82, 0x64, 0x05, 0x70, 0x09, 0x24,
	0xfe, 0xeb, 0x9f, 0xc0, 0x1f, 0xbe, 0x71, 0x61, 0x60, 0x4c, 0xcc, 0xf9, 0xf8, 0xdf, 0xd7, 0x94,
	0xab, 0x7d, 0xb5, 0x0b, 0x13, 0x99, 0x47, 0x27, 0xdc, 0x71, 0x85, 0x59, 0x54, 0xff, 0x03, 0x65,
	0xd4, 0xb2, 0x3b, 0xbb, 0xee, 0x67, 0x77, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x1c, 0x2a, 0xb7,
	0x50, 0x02, 0x00, 0x00,
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

// FileServiceServer is the server API for FileService service.
type FileServiceServer interface {
	UploadFile(FileService_UploadFileServer) error
}

// UnimplementedFileServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFileServiceServer struct {
}

func (*UnimplementedFileServiceServer) UploadFile(srv FileService_UploadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
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

var _FileService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "storage.FileService",
	HandlerType: (*FileServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadFile",
			Handler:       _FileService_UploadFile_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "storage/file.proto",
}
