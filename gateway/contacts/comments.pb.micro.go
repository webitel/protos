// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: gateway/contacts/comments.proto

package contacts

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Api Endpoints for Comments service

func NewCommentsEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		&api.Endpoint{
			Name:    "Comments.PublishComment",
			Path:    []string{"/contacts/{contact_id}/comments"},
			Method:  []string{"POST"},
			Body:    "input",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "Comments.SearchComments",
			Path:    []string{"/contacts/{contact_id}/comments"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "Comments.UpdateComment",
			Path:    []string{"/contacts/{contact_id}/comments/{input.etag}"},
			Method:  []string{"PUT"},
			Body:    "input",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "Comments.DeleteComment",
			Path:    []string{"/contacts/{contact_id}/comments/{etag}"},
			Method:  []string{"DELETE"},
			Body:    "",
			Handler: "rpc",
		},
	}
}

// Client API for Comments service

type CommentsService interface {
	// Publish comment for a Contact.
	PublishComment(ctx context.Context, in *PublishCommentRequest, opts ...client.CallOption) (*Comment, error)
	// Search for Contact Comment(s) ...
	SearchComments(ctx context.Context, in *SearchCommentsRequest, opts ...client.CallOption) (*CommentList, error)
	// Update (edit) specific Comment text owned
	UpdateComment(ctx context.Context, in *UpdateCommentRequest, opts ...client.CallOption) (*Comment, error)
	// Delete Comment(s) for Contact ...
	DeleteComment(ctx context.Context, in *DeleteCommentsRequest, opts ...client.CallOption) (*CommentList, error)
}

type commentsService struct {
	c    client.Client
	name string
}

func NewCommentsService(name string, c client.Client) CommentsService {
	return &commentsService{
		c:    c,
		name: name,
	}
}

func (c *commentsService) PublishComment(ctx context.Context, in *PublishCommentRequest, opts ...client.CallOption) (*Comment, error) {
	req := c.c.NewRequest(c.name, "Comments.PublishComment", in)
	out := new(Comment)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentsService) SearchComments(ctx context.Context, in *SearchCommentsRequest, opts ...client.CallOption) (*CommentList, error) {
	req := c.c.NewRequest(c.name, "Comments.SearchComments", in)
	out := new(CommentList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentsService) UpdateComment(ctx context.Context, in *UpdateCommentRequest, opts ...client.CallOption) (*Comment, error) {
	req := c.c.NewRequest(c.name, "Comments.UpdateComment", in)
	out := new(Comment)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentsService) DeleteComment(ctx context.Context, in *DeleteCommentsRequest, opts ...client.CallOption) (*CommentList, error) {
	req := c.c.NewRequest(c.name, "Comments.DeleteComment", in)
	out := new(CommentList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Comments service

type CommentsHandler interface {
	// Publish comment for a Contact.
	PublishComment(context.Context, *PublishCommentRequest, *Comment) error
	// Search for Contact Comment(s) ...
	SearchComments(context.Context, *SearchCommentsRequest, *CommentList) error
	// Update (edit) specific Comment text owned
	UpdateComment(context.Context, *UpdateCommentRequest, *Comment) error
	// Delete Comment(s) for Contact ...
	DeleteComment(context.Context, *DeleteCommentsRequest, *CommentList) error
}

func RegisterCommentsHandler(s server.Server, hdlr CommentsHandler, opts ...server.HandlerOption) error {
	type comments interface {
		PublishComment(ctx context.Context, in *PublishCommentRequest, out *Comment) error
		SearchComments(ctx context.Context, in *SearchCommentsRequest, out *CommentList) error
		UpdateComment(ctx context.Context, in *UpdateCommentRequest, out *Comment) error
		DeleteComment(ctx context.Context, in *DeleteCommentsRequest, out *CommentList) error
	}
	type Comments struct {
		comments
	}
	h := &commentsHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Comments.PublishComment",
		Path:    []string{"/contacts/{contact_id}/comments"},
		Method:  []string{"POST"},
		Body:    "input",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Comments.SearchComments",
		Path:    []string{"/contacts/{contact_id}/comments"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Comments.UpdateComment",
		Path:    []string{"/contacts/{contact_id}/comments/{input.etag}"},
		Method:  []string{"PUT"},
		Body:    "input",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Comments.DeleteComment",
		Path:    []string{"/contacts/{contact_id}/comments/{etag}"},
		Method:  []string{"DELETE"},
		Body:    "",
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&Comments{h}, opts...))
}

type commentsHandler struct {
	CommentsHandler
}

func (h *commentsHandler) PublishComment(ctx context.Context, in *PublishCommentRequest, out *Comment) error {
	return h.CommentsHandler.PublishComment(ctx, in, out)
}

func (h *commentsHandler) SearchComments(ctx context.Context, in *SearchCommentsRequest, out *CommentList) error {
	return h.CommentsHandler.SearchComments(ctx, in, out)
}

func (h *commentsHandler) UpdateComment(ctx context.Context, in *UpdateCommentRequest, out *Comment) error {
	return h.CommentsHandler.UpdateComment(ctx, in, out)
}

func (h *commentsHandler) DeleteComment(ctx context.Context, in *DeleteCommentsRequest, out *CommentList) error {
	return h.CommentsHandler.DeleteComment(ctx, in, out)
}
