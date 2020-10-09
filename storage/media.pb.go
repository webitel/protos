// Code generated by protoc-gen-go. DO NOT EDIT.
// source: storage/media.proto

package storage

import (
	context "context"
	engine "engine"
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

type DeleteMediaFileRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DomainId             int64    `protobuf:"varint,2,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteMediaFileRequest) Reset()         { *m = DeleteMediaFileRequest{} }
func (m *DeleteMediaFileRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteMediaFileRequest) ProtoMessage()    {}
func (*DeleteMediaFileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc8e6441afd8faef, []int{0}
}

func (m *DeleteMediaFileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteMediaFileRequest.Unmarshal(m, b)
}
func (m *DeleteMediaFileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteMediaFileRequest.Marshal(b, m, deterministic)
}
func (m *DeleteMediaFileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteMediaFileRequest.Merge(m, src)
}
func (m *DeleteMediaFileRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteMediaFileRequest.Size(m)
}
func (m *DeleteMediaFileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteMediaFileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteMediaFileRequest proto.InternalMessageInfo

func (m *DeleteMediaFileRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DeleteMediaFileRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

type ReadMediaFileRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DomainId             int64    `protobuf:"varint,2,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadMediaFileRequest) Reset()         { *m = ReadMediaFileRequest{} }
func (m *ReadMediaFileRequest) String() string { return proto.CompactTextString(m) }
func (*ReadMediaFileRequest) ProtoMessage()    {}
func (*ReadMediaFileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc8e6441afd8faef, []int{1}
}

func (m *ReadMediaFileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadMediaFileRequest.Unmarshal(m, b)
}
func (m *ReadMediaFileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadMediaFileRequest.Marshal(b, m, deterministic)
}
func (m *ReadMediaFileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadMediaFileRequest.Merge(m, src)
}
func (m *ReadMediaFileRequest) XXX_Size() int {
	return xxx_messageInfo_ReadMediaFileRequest.Size(m)
}
func (m *ReadMediaFileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadMediaFileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadMediaFileRequest proto.InternalMessageInfo

func (m *ReadMediaFileRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ReadMediaFileRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

type MediaFile struct {
	Id                   int64          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            int64          `protobuf:"varint,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	CreatedBy            *engine.Lookup `protobuf:"bytes,3,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	UpdatedAt            int64          `protobuf:"varint,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	UpdatedBy            *engine.Lookup `protobuf:"bytes,5,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
	Name                 string         `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Size                 int64          `protobuf:"varint,7,opt,name=size,proto3" json:"size,omitempty"`
	MimeType             string         `protobuf:"bytes,8,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *MediaFile) Reset()         { *m = MediaFile{} }
func (m *MediaFile) String() string { return proto.CompactTextString(m) }
func (*MediaFile) ProtoMessage()    {}
func (*MediaFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc8e6441afd8faef, []int{2}
}

func (m *MediaFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MediaFile.Unmarshal(m, b)
}
func (m *MediaFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MediaFile.Marshal(b, m, deterministic)
}
func (m *MediaFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MediaFile.Merge(m, src)
}
func (m *MediaFile) XXX_Size() int {
	return xxx_messageInfo_MediaFile.Size(m)
}
func (m *MediaFile) XXX_DiscardUnknown() {
	xxx_messageInfo_MediaFile.DiscardUnknown(m)
}

var xxx_messageInfo_MediaFile proto.InternalMessageInfo

func (m *MediaFile) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MediaFile) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *MediaFile) GetCreatedBy() *engine.Lookup {
	if m != nil {
		return m.CreatedBy
	}
	return nil
}

func (m *MediaFile) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *MediaFile) GetUpdatedBy() *engine.Lookup {
	if m != nil {
		return m.UpdatedBy
	}
	return nil
}

func (m *MediaFile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MediaFile) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *MediaFile) GetMimeType() string {
	if m != nil {
		return m.MimeType
	}
	return ""
}

type SearchMediaFileRequest struct {
	Page                 int32    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size                 int32    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Q                    string   `protobuf:"bytes,3,opt,name=q,proto3" json:"q,omitempty"`
	DomainId             int64    `protobuf:"varint,4,opt,name=domain_id,json=domainId,proto3" json:"domain_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchMediaFileRequest) Reset()         { *m = SearchMediaFileRequest{} }
func (m *SearchMediaFileRequest) String() string { return proto.CompactTextString(m) }
func (*SearchMediaFileRequest) ProtoMessage()    {}
func (*SearchMediaFileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc8e6441afd8faef, []int{3}
}

func (m *SearchMediaFileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchMediaFileRequest.Unmarshal(m, b)
}
func (m *SearchMediaFileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchMediaFileRequest.Marshal(b, m, deterministic)
}
func (m *SearchMediaFileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchMediaFileRequest.Merge(m, src)
}
func (m *SearchMediaFileRequest) XXX_Size() int {
	return xxx_messageInfo_SearchMediaFileRequest.Size(m)
}
func (m *SearchMediaFileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchMediaFileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchMediaFileRequest proto.InternalMessageInfo

func (m *SearchMediaFileRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *SearchMediaFileRequest) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *SearchMediaFileRequest) GetQ() string {
	if m != nil {
		return m.Q
	}
	return ""
}

func (m *SearchMediaFileRequest) GetDomainId() int64 {
	if m != nil {
		return m.DomainId
	}
	return 0
}

type ListMedia struct {
	Next                 bool         `protobuf:"varint,1,opt,name=next,proto3" json:"next,omitempty"`
	Items                []*MediaFile `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ListMedia) Reset()         { *m = ListMedia{} }
func (m *ListMedia) String() string { return proto.CompactTextString(m) }
func (*ListMedia) ProtoMessage()    {}
func (*ListMedia) Descriptor() ([]byte, []int) {
	return fileDescriptor_dc8e6441afd8faef, []int{4}
}

func (m *ListMedia) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMedia.Unmarshal(m, b)
}
func (m *ListMedia) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMedia.Marshal(b, m, deterministic)
}
func (m *ListMedia) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMedia.Merge(m, src)
}
func (m *ListMedia) XXX_Size() int {
	return xxx_messageInfo_ListMedia.Size(m)
}
func (m *ListMedia) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMedia.DiscardUnknown(m)
}

var xxx_messageInfo_ListMedia proto.InternalMessageInfo

func (m *ListMedia) GetNext() bool {
	if m != nil {
		return m.Next
	}
	return false
}

func (m *ListMedia) GetItems() []*MediaFile {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*DeleteMediaFileRequest)(nil), "storage.DeleteMediaFileRequest")
	proto.RegisterType((*ReadMediaFileRequest)(nil), "storage.ReadMediaFileRequest")
	proto.RegisterType((*MediaFile)(nil), "storage.MediaFile")
	proto.RegisterType((*SearchMediaFileRequest)(nil), "storage.SearchMediaFileRequest")
	proto.RegisterType((*ListMedia)(nil), "storage.ListMedia")
}

func init() { proto.RegisterFile("storage/media.proto", fileDescriptor_dc8e6441afd8faef) }

var fileDescriptor_dc8e6441afd8faef = []byte{
	// 489 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0x4d, 0x6e, 0xd3, 0x50,
	0x10, 0xc7, 0x65, 0x37, 0x69, 0xe3, 0x29, 0xa4, 0xd5, 0x2b, 0x44, 0x56, 0x4a, 0x45, 0xe4, 0x95,
	0x05, 0xaa, 0x2d, 0x85, 0x13, 0x10, 0x3e, 0xa4, 0x4a, 0x65, 0xe3, 0xb2, 0x62, 0x81, 0xf5, 0xec,
	0x37, 0x72, 0x47, 0x8d, 0xfd, 0x1c, 0xfb, 0x05, 0x30, 0x88, 0x0d, 0x57, 0xe0, 0x32, 0xdc, 0x83,
	0x2b, 0x70, 0x0c, 0x16, 0xc8, 0xcf, 0xae, 0x43, 0x52, 0x8b, 0x0d, 0xbb, 0xd1, 0x7f, 0x66, 0x7e,
	0xf3, 0xa1, 0x19, 0x38, 0x29, 0x95, 0x2c, 0x78, 0x82, 0x7e, 0x8a, 0x82, 0xb8, 0x97, 0x17, 0x52,
	0x49, 0x76, 0xd0, 0x8a, 0x53, 0x86, 0x59, 0x42, 0x19, 0xfa, 0xb1, 0xcc, 0x4a, 0xd5, 0x38, 0xa7,
	0x8f, 0x12, 0x29, 0x93, 0x25, 0xfa, 0x3c, 0x27, 0x9f, 0x67, 0x99, 0x54, 0x5c, 0x91, 0xcc, 0xca,
	0xc6, 0xeb, 0xbc, 0x82, 0xc9, 0x4b, 0x5c, 0xa2, 0xc2, 0x37, 0x35, 0xef, 0x35, 0x2d, 0x31, 0xc0,
	0xd5, 0x1a, 0x4b, 0xc5, 0xc6, 0x60, 0x92, 0xb0, 0x8d, 0x99, 0xe1, 0xee, 0x05, 0x26, 0x09, 0x76,
	0x0a, 0x96, 0x90, 0x29, 0xa7, 0x2c, 0x24, 0x61, 0x9b, 0x5a, 0x1e, 0x35, 0xc2, 0x85, 0x70, 0x5e,
	0xc0, 0x83, 0x00, 0xb9, 0xf8, 0x3f, 0xc8, 0x6f, 0x03, 0xac, 0x8e, 0x70, 0x27, 0xf5, 0x0c, 0x20,
	0x2e, 0x90, 0x2b, 0x14, 0x21, 0x57, 0x6d, 0xae, 0xd5, 0x2a, 0xcf, 0x15, 0x3b, 0xdf, 0xb8, 0xa3,
	0xca, 0xde, 0x9b, 0x19, 0xee, 0xe1, 0x7c, 0xec, 0x35, 0xfb, 0xf0, 0x2e, 0xa5, 0xbc, 0x59, 0xe7,
	0x5d, 0xf8, 0xa2, 0xaa, 0x69, 0xeb, 0x5c, 0xdc, 0xd2, 0x06, 0x0d, 0xad, 0x55, 0x1a, 0xda, 0xad,
	0x3b, 0xaa, 0xec, 0x61, 0x3f, 0xad, 0x8d, 0x58, 0x54, 0x8c, 0xc1, 0x20, 0xe3, 0x29, 0xda, 0xfb,
	0x33, 0xc3, 0xb5, 0x02, 0x6d, 0xd7, 0x5a, 0x49, 0x9f, 0xd1, 0x3e, 0xd0, 0x6c, 0x6d, 0xd7, 0xe3,
	0xa7, 0x94, 0x62, 0xa8, 0xaa, 0x1c, 0xed, 0x91, 0x0e, 0x1e, 0xd5, 0xc2, 0xdb, 0x2a, 0x47, 0xe7,
	0x06, 0x26, 0x57, 0xc8, 0x8b, 0xf8, 0xfa, 0xce, 0x16, 0x19, 0x0c, 0x72, 0x9e, 0xa0, 0x5e, 0xc6,
	0x30, 0xd0, 0x76, 0x87, 0x37, 0x1b, 0x4d, 0xe3, 0xef, 0x81, 0xb1, 0xd2, 0xa3, 0x5b, 0x81, 0xb1,
	0xda, 0xde, 0xf5, 0x60, 0x67, 0xd7, 0x17, 0x60, 0x5d, 0x52, 0xa9, 0x74, 0x29, 0xdd, 0x3e, 0x7e,
	0x52, 0x9a, 0x3f, 0x0a, 0xb4, 0xcd, 0x5c, 0x18, 0x92, 0xc2, 0xb4, 0xb4, 0xcd, 0xd9, 0x9e, 0x7b,
	0x38, 0x67, 0x5e, 0x7b, 0x63, 0xde, 0xa6, 0xbb, 0x26, 0x60, 0xfe, 0xc3, 0x84, 0xe3, 0x4e, 0xbc,
	0xc2, 0xe2, 0x03, 0xc5, 0xc8, 0xde, 0xc3, 0xd1, 0xce, 0x30, 0xec, 0x71, 0x87, 0xe8, 0x1f, 0x73,
	0xba, 0xa9, 0xd1, 0xb5, 0xe6, 0x4c, 0xbe, 0xfd, 0xfc, 0xf5, 0xdd, 0x3c, 0x66, 0x63, 0x7f, 0xeb,
	0xf0, 0x59, 0x08, 0xf7, 0xb7, 0x0e, 0x8e, 0x9d, 0x75, 0xc9, 0x7d, 0x87, 0x38, 0xed, 0xe9, 0xdf,
	0x39, 0xd5, 0xec, 0x87, 0xec, 0x64, 0x9b, 0xed, 0x7f, 0x21, 0xf1, 0x95, 0xc5, 0x70, 0xb4, 0xf3,
	0x18, 0x7f, 0x0d, 0xd0, 0xff, 0x32, 0xff, 0x2a, 0xf2, 0xa4, 0xaf, 0xc8, 0xe2, 0xfc, 0xdd, 0xd3,
	0x84, 0xd4, 0xf5, 0x3a, 0xf2, 0x62, 0x99, 0xfa, 0x1f, 0x31, 0x22, 0x85, 0xcb, 0x2e, 0x30, 0x29,
	0xf2, 0x38, 0xac, 0xdf, 0xb6, 0x15, 0xa2, 0x7d, 0xfd, 0xb3, 0xcf, 0xfe, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x68, 0x1c, 0x32, 0x7f, 0x05, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MediaFileServiceClient is the client API for MediaFileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MediaFileServiceClient interface {
	// Search MediaFile
	SearchMediaFile(ctx context.Context, in *SearchMediaFileRequest, opts ...grpc.CallOption) (*ListMedia, error)
	// MediaFile item
	ReadMediaFile(ctx context.Context, in *ReadMediaFileRequest, opts ...grpc.CallOption) (*MediaFile, error)
	// Remove MediaFile
	DeleteMediaFile(ctx context.Context, in *DeleteMediaFileRequest, opts ...grpc.CallOption) (*MediaFile, error)
}

type mediaFileServiceClient struct {
	cc *grpc.ClientConn
}

func NewMediaFileServiceClient(cc *grpc.ClientConn) MediaFileServiceClient {
	return &mediaFileServiceClient{cc}
}

func (c *mediaFileServiceClient) SearchMediaFile(ctx context.Context, in *SearchMediaFileRequest, opts ...grpc.CallOption) (*ListMedia, error) {
	out := new(ListMedia)
	err := c.cc.Invoke(ctx, "/storage.MediaFileService/SearchMediaFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaFileServiceClient) ReadMediaFile(ctx context.Context, in *ReadMediaFileRequest, opts ...grpc.CallOption) (*MediaFile, error) {
	out := new(MediaFile)
	err := c.cc.Invoke(ctx, "/storage.MediaFileService/ReadMediaFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mediaFileServiceClient) DeleteMediaFile(ctx context.Context, in *DeleteMediaFileRequest, opts ...grpc.CallOption) (*MediaFile, error) {
	out := new(MediaFile)
	err := c.cc.Invoke(ctx, "/storage.MediaFileService/DeleteMediaFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MediaFileServiceServer is the server API for MediaFileService service.
type MediaFileServiceServer interface {
	// Search MediaFile
	SearchMediaFile(context.Context, *SearchMediaFileRequest) (*ListMedia, error)
	// MediaFile item
	ReadMediaFile(context.Context, *ReadMediaFileRequest) (*MediaFile, error)
	// Remove MediaFile
	DeleteMediaFile(context.Context, *DeleteMediaFileRequest) (*MediaFile, error)
}

// UnimplementedMediaFileServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMediaFileServiceServer struct {
}

func (*UnimplementedMediaFileServiceServer) SearchMediaFile(ctx context.Context, req *SearchMediaFileRequest) (*ListMedia, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchMediaFile not implemented")
}
func (*UnimplementedMediaFileServiceServer) ReadMediaFile(ctx context.Context, req *ReadMediaFileRequest) (*MediaFile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadMediaFile not implemented")
}
func (*UnimplementedMediaFileServiceServer) DeleteMediaFile(ctx context.Context, req *DeleteMediaFileRequest) (*MediaFile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMediaFile not implemented")
}

func RegisterMediaFileServiceServer(s *grpc.Server, srv MediaFileServiceServer) {
	s.RegisterService(&_MediaFileService_serviceDesc, srv)
}

func _MediaFileService_SearchMediaFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchMediaFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaFileServiceServer).SearchMediaFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.MediaFileService/SearchMediaFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaFileServiceServer).SearchMediaFile(ctx, req.(*SearchMediaFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaFileService_ReadMediaFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadMediaFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaFileServiceServer).ReadMediaFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.MediaFileService/ReadMediaFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaFileServiceServer).ReadMediaFile(ctx, req.(*ReadMediaFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MediaFileService_DeleteMediaFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMediaFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MediaFileServiceServer).DeleteMediaFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.MediaFileService/DeleteMediaFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MediaFileServiceServer).DeleteMediaFile(ctx, req.(*DeleteMediaFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MediaFileService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "storage.MediaFileService",
	HandlerType: (*MediaFileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchMediaFile",
			Handler:    _MediaFileService_SearchMediaFile_Handler,
		},
		{
			MethodName: "ReadMediaFile",
			Handler:    _MediaFileService_ReadMediaFile_Handler,
		},
		{
			MethodName: "DeleteMediaFile",
			Handler:    _MediaFileService_DeleteMediaFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "storage/media.proto",
}
