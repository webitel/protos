// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: gateway/contacts/variables.proto

package contacts

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/genproto/googleapis/api/visibility"
	_ "google.golang.org/protobuf/types/known/structpb"
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

// Api Endpoints for Variables service

func NewVariablesEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		&api.Endpoint{
			Name:    "Variables.ListVariables",
			Path:    []string{"/contacts/{contact_id}/variables"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "Variables.MergeVariables",
			Path:    []string{"/contacts/{contact_id}/variables"},
			Method:  []string{"POST"},
			Body:    "input",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "Variables.ResetVariables",
			Path:    []string{"/contacts/{contact_id}/variables"},
			Method:  []string{"PUT"},
			Body:    "input",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "Variables.DeleteVariables",
			Path:    []string{"/contacts/{contact_id}/variables"},
			Method:  []string{"DELETE"},
			Body:    "",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "Variables.UpdateVariable",
			Path:    []string{"/contacts/{contact_id}/variables/{input.etag}"},
			Method:  []string{"PUT"},
			Body:    "input",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "Variables.DeleteVariable",
			Path:    []string{"/contacts/{contact_id}/variables/{etag}"},
			Method:  []string{"DELETE"},
			Body:    "",
			Handler: "rpc",
		},
	}
}

// Client API for Variables service

type VariablesService interface {
	// List variables of the contact
	ListVariables(ctx context.Context, in *SearchVariablesRequest, opts ...client.CallOption) (*VariableList, error)
	// Update or append variables to the contact
	MergeVariables(ctx context.Context, in *MergeVariablesRequest, opts ...client.CallOption) (*VariableList, error)
	// Reset all variables of the contact
	ResetVariables(ctx context.Context, in *ResetVariablesRequest, opts ...client.CallOption) (*VariableList, error)
	// Remove variable(s) of the contact
	DeleteVariables(ctx context.Context, in *DeleteVariablesRequest, opts ...client.CallOption) (*VariableList, error)
	// Update contact variable
	UpdateVariable(ctx context.Context, in *UpdateVariableRequest, opts ...client.CallOption) (*VariableList, error)
	// Remove the contact's variable by etag
	DeleteVariable(ctx context.Context, in *DeleteVariableRequest, opts ...client.CallOption) (*Variable, error)
}

type variablesService struct {
	c    client.Client
	name string
}

func NewVariablesService(name string, c client.Client) VariablesService {
	return &variablesService{
		c:    c,
		name: name,
	}
}

func (c *variablesService) ListVariables(ctx context.Context, in *SearchVariablesRequest, opts ...client.CallOption) (*VariableList, error) {
	req := c.c.NewRequest(c.name, "Variables.ListVariables", in)
	out := new(VariableList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *variablesService) MergeVariables(ctx context.Context, in *MergeVariablesRequest, opts ...client.CallOption) (*VariableList, error) {
	req := c.c.NewRequest(c.name, "Variables.MergeVariables", in)
	out := new(VariableList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *variablesService) ResetVariables(ctx context.Context, in *ResetVariablesRequest, opts ...client.CallOption) (*VariableList, error) {
	req := c.c.NewRequest(c.name, "Variables.ResetVariables", in)
	out := new(VariableList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *variablesService) DeleteVariables(ctx context.Context, in *DeleteVariablesRequest, opts ...client.CallOption) (*VariableList, error) {
	req := c.c.NewRequest(c.name, "Variables.DeleteVariables", in)
	out := new(VariableList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *variablesService) UpdateVariable(ctx context.Context, in *UpdateVariableRequest, opts ...client.CallOption) (*VariableList, error) {
	req := c.c.NewRequest(c.name, "Variables.UpdateVariable", in)
	out := new(VariableList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *variablesService) DeleteVariable(ctx context.Context, in *DeleteVariableRequest, opts ...client.CallOption) (*Variable, error) {
	req := c.c.NewRequest(c.name, "Variables.DeleteVariable", in)
	out := new(Variable)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Variables service

type VariablesHandler interface {
	// List variables of the contact
	ListVariables(context.Context, *SearchVariablesRequest, *VariableList) error
	// Update or append variables to the contact
	MergeVariables(context.Context, *MergeVariablesRequest, *VariableList) error
	// Reset all variables of the contact
	ResetVariables(context.Context, *ResetVariablesRequest, *VariableList) error
	// Remove variable(s) of the contact
	DeleteVariables(context.Context, *DeleteVariablesRequest, *VariableList) error
	// Update contact variable
	UpdateVariable(context.Context, *UpdateVariableRequest, *VariableList) error
	// Remove the contact's variable by etag
	DeleteVariable(context.Context, *DeleteVariableRequest, *Variable) error
}

func RegisterVariablesHandler(s server.Server, hdlr VariablesHandler, opts ...server.HandlerOption) error {
	type variables interface {
		ListVariables(ctx context.Context, in *SearchVariablesRequest, out *VariableList) error
		MergeVariables(ctx context.Context, in *MergeVariablesRequest, out *VariableList) error
		ResetVariables(ctx context.Context, in *ResetVariablesRequest, out *VariableList) error
		DeleteVariables(ctx context.Context, in *DeleteVariablesRequest, out *VariableList) error
		UpdateVariable(ctx context.Context, in *UpdateVariableRequest, out *VariableList) error
		DeleteVariable(ctx context.Context, in *DeleteVariableRequest, out *Variable) error
	}
	type Variables struct {
		variables
	}
	h := &variablesHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Variables.ListVariables",
		Path:    []string{"/contacts/{contact_id}/variables"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Variables.MergeVariables",
		Path:    []string{"/contacts/{contact_id}/variables"},
		Method:  []string{"POST"},
		Body:    "input",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Variables.ResetVariables",
		Path:    []string{"/contacts/{contact_id}/variables"},
		Method:  []string{"PUT"},
		Body:    "input",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Variables.DeleteVariables",
		Path:    []string{"/contacts/{contact_id}/variables"},
		Method:  []string{"DELETE"},
		Body:    "",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Variables.UpdateVariable",
		Path:    []string{"/contacts/{contact_id}/variables/{input.etag}"},
		Method:  []string{"PUT"},
		Body:    "input",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Variables.DeleteVariable",
		Path:    []string{"/contacts/{contact_id}/variables/{etag}"},
		Method:  []string{"DELETE"},
		Body:    "",
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&Variables{h}, opts...))
}

type variablesHandler struct {
	VariablesHandler
}

func (h *variablesHandler) ListVariables(ctx context.Context, in *SearchVariablesRequest, out *VariableList) error {
	return h.VariablesHandler.ListVariables(ctx, in, out)
}

func (h *variablesHandler) MergeVariables(ctx context.Context, in *MergeVariablesRequest, out *VariableList) error {
	return h.VariablesHandler.MergeVariables(ctx, in, out)
}

func (h *variablesHandler) ResetVariables(ctx context.Context, in *ResetVariablesRequest, out *VariableList) error {
	return h.VariablesHandler.ResetVariables(ctx, in, out)
}

func (h *variablesHandler) DeleteVariables(ctx context.Context, in *DeleteVariablesRequest, out *VariableList) error {
	return h.VariablesHandler.DeleteVariables(ctx, in, out)
}

func (h *variablesHandler) UpdateVariable(ctx context.Context, in *UpdateVariableRequest, out *VariableList) error {
	return h.VariablesHandler.UpdateVariable(ctx, in, out)
}

func (h *variablesHandler) DeleteVariable(ctx context.Context, in *DeleteVariableRequest, out *Variable) error {
	return h.VariablesHandler.DeleteVariable(ctx, in, out)
}
