// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: srv/appcron/proto/appcron.proto

package appcron

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	proto1 "global/global_proto/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Appcron service

type AppcronService interface {
	GetCronTask(ctx context.Context, in *proto1.Empty, opts ...client.CallOption) (*ListCronTask, error)
	RegisterCronTask(ctx context.Context, in *CronTask, opts ...client.CallOption) (*proto1.Response, error)
	UpdateCronTask(ctx context.Context, in *CronTask, opts ...client.CallOption) (*proto1.Response, error)
}

type appcronService struct {
	c    client.Client
	name string
}

func NewAppcronService(name string, c client.Client) AppcronService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "appcron"
	}
	return &appcronService{
		c:    c,
		name: name,
	}
}

func (c *appcronService) GetCronTask(ctx context.Context, in *proto1.Empty, opts ...client.CallOption) (*ListCronTask, error) {
	req := c.c.NewRequest(c.name, "Appcron.GetCronTask", in)
	out := new(ListCronTask)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appcronService) RegisterCronTask(ctx context.Context, in *CronTask, opts ...client.CallOption) (*proto1.Response, error) {
	req := c.c.NewRequest(c.name, "Appcron.RegisterCronTask", in)
	out := new(proto1.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appcronService) UpdateCronTask(ctx context.Context, in *CronTask, opts ...client.CallOption) (*proto1.Response, error) {
	req := c.c.NewRequest(c.name, "Appcron.UpdateCronTask", in)
	out := new(proto1.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Appcron service

type AppcronHandler interface {
	GetCronTask(context.Context, *proto1.Empty, *ListCronTask) error
	RegisterCronTask(context.Context, *CronTask, *proto1.Response) error
	UpdateCronTask(context.Context, *CronTask, *proto1.Response) error
}

func RegisterAppcronHandler(s server.Server, hdlr AppcronHandler, opts ...server.HandlerOption) error {
	type appcron interface {
		GetCronTask(ctx context.Context, in *proto1.Empty, out *ListCronTask) error
		RegisterCronTask(ctx context.Context, in *CronTask, out *proto1.Response) error
		UpdateCronTask(ctx context.Context, in *CronTask, out *proto1.Response) error
	}
	type Appcron struct {
		appcron
	}
	h := &appcronHandler{hdlr}
	return s.Handle(s.NewHandler(&Appcron{h}, opts...))
}

type appcronHandler struct {
	AppcronHandler
}

func (h *appcronHandler) GetCronTask(ctx context.Context, in *proto1.Empty, out *ListCronTask) error {
	return h.AppcronHandler.GetCronTask(ctx, in, out)
}

func (h *appcronHandler) RegisterCronTask(ctx context.Context, in *CronTask, out *proto1.Response) error {
	return h.AppcronHandler.RegisterCronTask(ctx, in, out)
}

func (h *appcronHandler) UpdateCronTask(ctx context.Context, in *CronTask, out *proto1.Response) error {
	return h.AppcronHandler.UpdateCronTask(ctx, in, out)
}
