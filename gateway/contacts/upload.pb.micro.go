// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: gateway/contacts/upload.proto

package contacts

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/rpc/status"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Upload service

func NewUploadEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Upload service

type UploadService interface {
	// Upload an image or photo
	UploadMedia(ctx context.Context, opts ...client.CallOption) (Upload_UploadMediaService, error)
}

type uploadService struct {
	c    client.Client
	name string
}

func NewUploadService(name string, c client.Client) UploadService {
	return &uploadService{
		c:    c,
		name: name,
	}
}

func (c *uploadService) UploadMedia(ctx context.Context, opts ...client.CallOption) (Upload_UploadMediaService, error) {
	req := c.c.NewRequest(c.name, "Upload.UploadMedia", &UploadMediaRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &uploadServiceUploadMedia{stream}, nil
}

type Upload_UploadMediaService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*UploadMediaRequest) error
	Recv() (*UploadMediaResponse, error)
}

type uploadServiceUploadMedia struct {
	stream client.Stream
}

func (x *uploadServiceUploadMedia) Close() error {
	return x.stream.Close()
}

func (x *uploadServiceUploadMedia) Context() context.Context {
	return x.stream.Context()
}

func (x *uploadServiceUploadMedia) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *uploadServiceUploadMedia) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *uploadServiceUploadMedia) Send(m *UploadMediaRequest) error {
	return x.stream.Send(m)
}

func (x *uploadServiceUploadMedia) Recv() (*UploadMediaResponse, error) {
	m := new(UploadMediaResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Upload service

type UploadHandler interface {
	// Upload an image or photo
	UploadMedia(context.Context, Upload_UploadMediaStream) error
}

func RegisterUploadHandler(s server.Server, hdlr UploadHandler, opts ...server.HandlerOption) error {
	type upload interface {
		UploadMedia(ctx context.Context, stream server.Stream) error
	}
	type Upload struct {
		upload
	}
	h := &uploadHandler{hdlr}
	return s.Handle(s.NewHandler(&Upload{h}, opts...))
}

type uploadHandler struct {
	UploadHandler
}

func (h *uploadHandler) UploadMedia(ctx context.Context, stream server.Stream) error {
	return h.UploadHandler.UploadMedia(ctx, &uploadUploadMediaStream{stream})
}

type Upload_UploadMediaStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*UploadMediaResponse) error
	Recv() (*UploadMediaRequest, error)
}

type uploadUploadMediaStream struct {
	stream server.Stream
}

func (x *uploadUploadMediaStream) Close() error {
	return x.stream.Close()
}

func (x *uploadUploadMediaStream) Context() context.Context {
	return x.stream.Context()
}

func (x *uploadUploadMediaStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *uploadUploadMediaStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *uploadUploadMediaStream) Send(m *UploadMediaResponse) error {
	return x.stream.Send(m)
}

func (x *uploadUploadMediaStream) Recv() (*UploadMediaRequest, error) {
	m := new(UploadMediaRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}