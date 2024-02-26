// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: gateway/contacts/managers.proto

package contacts

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/genproto/googleapis/api/visibility"
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

// Api Endpoints for Managers service

func NewManagersEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		&api.Endpoint{
			Name:    "Managers.ListManagers",
			Path:    []string{"/contacts/{contact_id}/managers"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "Managers.MergeManagers",
			Path:    []string{"/contacts/{contact_id}/managers"},
			Method:  []string{"POST"},
			Body:    "input",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "Managers.ResetManagers",
			Path:    []string{"/contacts/{contact_id}/managers"},
			Method:  []string{"PUT"},
			Body:    "input",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "Managers.DeleteManagers",
			Path:    []string{"/contacts/{contact_id}/managers"},
			Method:  []string{"DELETE"},
			Body:    "",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "Managers.LocateManager",
			Path:    []string{"/contacts/{contact_id}/managers/{etag}"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "Managers.UpdateManager",
			Path:    []string{"/contacts/{contact_id}/managers/{input.etag}"},
			Method:  []string{"PUT"},
			Body:    "input",
			Handler: "rpc",
		},
		&api.Endpoint{
			Name:    "Managers.DeleteManager",
			Path:    []string{"/contacts/{contact_id}/managers/{etag}"},
			Method:  []string{"DELETE"},
			Body:    "",
			Handler: "rpc",
		},
	}
}

// Client API for Managers service

type ManagersService interface {
	// Search the Contact's Managers.
	ListManagers(ctx context.Context, in *ListManagersRequest, opts ...client.CallOption) (*ManagerList, error)
	// Associate new Managers to the Contact.
	MergeManagers(ctx context.Context, in *MergeManagersRequest, opts ...client.CallOption) (*ManagerList, error)
	// Reset Managers to fit the specified final set.
	ResetManagers(ctx context.Context, in *ResetManagersRequest, opts ...client.CallOption) (*ManagerList, error)
	// Remove Contact Managers associations.
	DeleteManagers(ctx context.Context, in *DeleteManagersRequest, opts ...client.CallOption) (*ManagerList, error)
	// Locate the manager address link.
	LocateManager(ctx context.Context, in *LocateManagerRequest, opts ...client.CallOption) (*Manager, error)
	// Update the contact's manager address link details
	UpdateManager(ctx context.Context, in *UpdateManagerRequest, opts ...client.CallOption) (*ManagerList, error)
	// Remove the contact's manager address link
	DeleteManager(ctx context.Context, in *DeleteManagerRequest, opts ...client.CallOption) (*Manager, error)
}

type managersService struct {
	c    client.Client
	name string
}

func NewManagersService(name string, c client.Client) ManagersService {
	return &managersService{
		c:    c,
		name: name,
	}
}

func (c *managersService) ListManagers(ctx context.Context, in *ListManagersRequest, opts ...client.CallOption) (*ManagerList, error) {
	req := c.c.NewRequest(c.name, "Managers.ListManagers", in)
	out := new(ManagerList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managersService) MergeManagers(ctx context.Context, in *MergeManagersRequest, opts ...client.CallOption) (*ManagerList, error) {
	req := c.c.NewRequest(c.name, "Managers.MergeManagers", in)
	out := new(ManagerList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managersService) ResetManagers(ctx context.Context, in *ResetManagersRequest, opts ...client.CallOption) (*ManagerList, error) {
	req := c.c.NewRequest(c.name, "Managers.ResetManagers", in)
	out := new(ManagerList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managersService) DeleteManagers(ctx context.Context, in *DeleteManagersRequest, opts ...client.CallOption) (*ManagerList, error) {
	req := c.c.NewRequest(c.name, "Managers.DeleteManagers", in)
	out := new(ManagerList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managersService) LocateManager(ctx context.Context, in *LocateManagerRequest, opts ...client.CallOption) (*Manager, error) {
	req := c.c.NewRequest(c.name, "Managers.LocateManager", in)
	out := new(Manager)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managersService) UpdateManager(ctx context.Context, in *UpdateManagerRequest, opts ...client.CallOption) (*ManagerList, error) {
	req := c.c.NewRequest(c.name, "Managers.UpdateManager", in)
	out := new(ManagerList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managersService) DeleteManager(ctx context.Context, in *DeleteManagerRequest, opts ...client.CallOption) (*Manager, error) {
	req := c.c.NewRequest(c.name, "Managers.DeleteManager", in)
	out := new(Manager)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Managers service

type ManagersHandler interface {
	// Search the Contact's Managers.
	ListManagers(context.Context, *ListManagersRequest, *ManagerList) error
	// Associate new Managers to the Contact.
	MergeManagers(context.Context, *MergeManagersRequest, *ManagerList) error
	// Reset Managers to fit the specified final set.
	ResetManagers(context.Context, *ResetManagersRequest, *ManagerList) error
	// Remove Contact Managers associations.
	DeleteManagers(context.Context, *DeleteManagersRequest, *ManagerList) error
	// Locate the manager address link.
	LocateManager(context.Context, *LocateManagerRequest, *Manager) error
	// Update the contact's manager address link details
	UpdateManager(context.Context, *UpdateManagerRequest, *ManagerList) error
	// Remove the contact's manager address link
	DeleteManager(context.Context, *DeleteManagerRequest, *Manager) error
}

func RegisterManagersHandler(s server.Server, hdlr ManagersHandler, opts ...server.HandlerOption) error {
	type managers interface {
		ListManagers(ctx context.Context, in *ListManagersRequest, out *ManagerList) error
		MergeManagers(ctx context.Context, in *MergeManagersRequest, out *ManagerList) error
		ResetManagers(ctx context.Context, in *ResetManagersRequest, out *ManagerList) error
		DeleteManagers(ctx context.Context, in *DeleteManagersRequest, out *ManagerList) error
		LocateManager(ctx context.Context, in *LocateManagerRequest, out *Manager) error
		UpdateManager(ctx context.Context, in *UpdateManagerRequest, out *ManagerList) error
		DeleteManager(ctx context.Context, in *DeleteManagerRequest, out *Manager) error
	}
	type Managers struct {
		managers
	}
	h := &managersHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Managers.ListManagers",
		Path:    []string{"/contacts/{contact_id}/managers"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Managers.MergeManagers",
		Path:    []string{"/contacts/{contact_id}/managers"},
		Method:  []string{"POST"},
		Body:    "input",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Managers.ResetManagers",
		Path:    []string{"/contacts/{contact_id}/managers"},
		Method:  []string{"PUT"},
		Body:    "input",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Managers.DeleteManagers",
		Path:    []string{"/contacts/{contact_id}/managers"},
		Method:  []string{"DELETE"},
		Body:    "",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Managers.LocateManager",
		Path:    []string{"/contacts/{contact_id}/managers/{etag}"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Managers.UpdateManager",
		Path:    []string{"/contacts/{contact_id}/managers/{input.etag}"},
		Method:  []string{"PUT"},
		Body:    "input",
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "Managers.DeleteManager",
		Path:    []string{"/contacts/{contact_id}/managers/{etag}"},
		Method:  []string{"DELETE"},
		Body:    "",
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&Managers{h}, opts...))
}

type managersHandler struct {
	ManagersHandler
}

func (h *managersHandler) ListManagers(ctx context.Context, in *ListManagersRequest, out *ManagerList) error {
	return h.ManagersHandler.ListManagers(ctx, in, out)
}

func (h *managersHandler) MergeManagers(ctx context.Context, in *MergeManagersRequest, out *ManagerList) error {
	return h.ManagersHandler.MergeManagers(ctx, in, out)
}

func (h *managersHandler) ResetManagers(ctx context.Context, in *ResetManagersRequest, out *ManagerList) error {
	return h.ManagersHandler.ResetManagers(ctx, in, out)
}

func (h *managersHandler) DeleteManagers(ctx context.Context, in *DeleteManagersRequest, out *ManagerList) error {
	return h.ManagersHandler.DeleteManagers(ctx, in, out)
}

func (h *managersHandler) LocateManager(ctx context.Context, in *LocateManagerRequest, out *Manager) error {
	return h.ManagersHandler.LocateManager(ctx, in, out)
}

func (h *managersHandler) UpdateManager(ctx context.Context, in *UpdateManagerRequest, out *ManagerList) error {
	return h.ManagersHandler.UpdateManager(ctx, in, out)
}

func (h *managersHandler) DeleteManager(ctx context.Context, in *DeleteManagerRequest, out *Manager) error {
	return h.ManagersHandler.DeleteManager(ctx, in, out)
}