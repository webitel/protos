// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.6
// source: storage/file_transcript.proto

package storage

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	FileTranscriptService_CreateFileTranscript_FullMethodName     = "/storage.FileTranscriptService/CreateFileTranscript"
	FileTranscriptService_FileTranscriptSafe_FullMethodName       = "/storage.FileTranscriptService/FileTranscriptSafe"
	FileTranscriptService_GetFileTranscriptPhrases_FullMethodName = "/storage.FileTranscriptService/GetFileTranscriptPhrases"
	FileTranscriptService_DeleteFileTranscript_FullMethodName     = "/storage.FileTranscriptService/DeleteFileTranscript"
)

// FileTranscriptServiceClient is the client API for FileTranscriptService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileTranscriptServiceClient interface {
	CreateFileTranscript(ctx context.Context, in *StartFileTranscriptRequest, opts ...grpc.CallOption) (*StartFileTranscriptResponse, error)
	FileTranscriptSafe(ctx context.Context, in *FileTranscriptSafeRequest, opts ...grpc.CallOption) (*FileTranscriptSafeResponse, error)
	GetFileTranscriptPhrases(ctx context.Context, in *GetFileTranscriptPhrasesRequest, opts ...grpc.CallOption) (*ListPhrases, error)
	DeleteFileTranscript(ctx context.Context, in *DeleteFileTranscriptRequest, opts ...grpc.CallOption) (*DeleteFileTranscriptResponse, error)
}

type fileTranscriptServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileTranscriptServiceClient(cc grpc.ClientConnInterface) FileTranscriptServiceClient {
	return &fileTranscriptServiceClient{cc}
}

func (c *fileTranscriptServiceClient) CreateFileTranscript(ctx context.Context, in *StartFileTranscriptRequest, opts ...grpc.CallOption) (*StartFileTranscriptResponse, error) {
	out := new(StartFileTranscriptResponse)
	err := c.cc.Invoke(ctx, FileTranscriptService_CreateFileTranscript_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileTranscriptServiceClient) FileTranscriptSafe(ctx context.Context, in *FileTranscriptSafeRequest, opts ...grpc.CallOption) (*FileTranscriptSafeResponse, error) {
	out := new(FileTranscriptSafeResponse)
	err := c.cc.Invoke(ctx, FileTranscriptService_FileTranscriptSafe_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileTranscriptServiceClient) GetFileTranscriptPhrases(ctx context.Context, in *GetFileTranscriptPhrasesRequest, opts ...grpc.CallOption) (*ListPhrases, error) {
	out := new(ListPhrases)
	err := c.cc.Invoke(ctx, FileTranscriptService_GetFileTranscriptPhrases_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileTranscriptServiceClient) DeleteFileTranscript(ctx context.Context, in *DeleteFileTranscriptRequest, opts ...grpc.CallOption) (*DeleteFileTranscriptResponse, error) {
	out := new(DeleteFileTranscriptResponse)
	err := c.cc.Invoke(ctx, FileTranscriptService_DeleteFileTranscript_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileTranscriptServiceServer is the server API for FileTranscriptService service.
// All implementations must embed UnimplementedFileTranscriptServiceServer
// for forward compatibility
type FileTranscriptServiceServer interface {
	CreateFileTranscript(context.Context, *StartFileTranscriptRequest) (*StartFileTranscriptResponse, error)
	FileTranscriptSafe(context.Context, *FileTranscriptSafeRequest) (*FileTranscriptSafeResponse, error)
	GetFileTranscriptPhrases(context.Context, *GetFileTranscriptPhrasesRequest) (*ListPhrases, error)
	DeleteFileTranscript(context.Context, *DeleteFileTranscriptRequest) (*DeleteFileTranscriptResponse, error)
	mustEmbedUnimplementedFileTranscriptServiceServer()
}

// UnimplementedFileTranscriptServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFileTranscriptServiceServer struct {
}

func (UnimplementedFileTranscriptServiceServer) CreateFileTranscript(context.Context, *StartFileTranscriptRequest) (*StartFileTranscriptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFileTranscript not implemented")
}
func (UnimplementedFileTranscriptServiceServer) FileTranscriptSafe(context.Context, *FileTranscriptSafeRequest) (*FileTranscriptSafeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FileTranscriptSafe not implemented")
}
func (UnimplementedFileTranscriptServiceServer) GetFileTranscriptPhrases(context.Context, *GetFileTranscriptPhrasesRequest) (*ListPhrases, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFileTranscriptPhrases not implemented")
}
func (UnimplementedFileTranscriptServiceServer) DeleteFileTranscript(context.Context, *DeleteFileTranscriptRequest) (*DeleteFileTranscriptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFileTranscript not implemented")
}
func (UnimplementedFileTranscriptServiceServer) mustEmbedUnimplementedFileTranscriptServiceServer() {}

// UnsafeFileTranscriptServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileTranscriptServiceServer will
// result in compilation errors.
type UnsafeFileTranscriptServiceServer interface {
	mustEmbedUnimplementedFileTranscriptServiceServer()
}

func RegisterFileTranscriptServiceServer(s grpc.ServiceRegistrar, srv FileTranscriptServiceServer) {
	s.RegisterService(&FileTranscriptService_ServiceDesc, srv)
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
		FullMethod: FileTranscriptService_CreateFileTranscript_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileTranscriptServiceServer).CreateFileTranscript(ctx, req.(*StartFileTranscriptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileTranscriptService_FileTranscriptSafe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileTranscriptSafeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileTranscriptServiceServer).FileTranscriptSafe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileTranscriptService_FileTranscriptSafe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileTranscriptServiceServer).FileTranscriptSafe(ctx, req.(*FileTranscriptSafeRequest))
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
		FullMethod: FileTranscriptService_GetFileTranscriptPhrases_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileTranscriptServiceServer).GetFileTranscriptPhrases(ctx, req.(*GetFileTranscriptPhrasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileTranscriptService_DeleteFileTranscript_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFileTranscriptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileTranscriptServiceServer).DeleteFileTranscript(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileTranscriptService_DeleteFileTranscript_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileTranscriptServiceServer).DeleteFileTranscript(ctx, req.(*DeleteFileTranscriptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FileTranscriptService_ServiceDesc is the grpc.ServiceDesc for FileTranscriptService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileTranscriptService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "storage.FileTranscriptService",
	HandlerType: (*FileTranscriptServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFileTranscript",
			Handler:    _FileTranscriptService_CreateFileTranscript_Handler,
		},
		{
			MethodName: "FileTranscriptSafe",
			Handler:    _FileTranscriptService_FileTranscriptSafe_Handler,
		},
		{
			MethodName: "GetFileTranscriptPhrases",
			Handler:    _FileTranscriptService_GetFileTranscriptPhrases_Handler,
		},
		{
			MethodName: "DeleteFileTranscript",
			Handler:    _FileTranscriptService_DeleteFileTranscript_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "storage/file_transcript.proto",
}
