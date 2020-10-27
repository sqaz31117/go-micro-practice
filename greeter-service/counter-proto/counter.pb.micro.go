// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/counter.proto

package counter

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Api Endpoints for Counter service

func NewCounterEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Counter service

type CounterService interface {
	// 定義 API Interface，Count 為此 API 的 Name，
	// 代表 給 Count API Request 當參數，並返回 Response
	Count(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type counterService struct {
	c    client.Client
	name string
}

func NewCounterService(name string, c client.Client) CounterService {
	return &counterService{
		c:    c,
		name: name,
	}
}

func (c *counterService) Count(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Counter.Count", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Counter service

type CounterHandler interface {
	// 定義 API Interface，Count 為此 API 的 Name，
	// 代表 給 Count API Request 當參數，並返回 Response
	Count(context.Context, *Request, *Response) error
}

func RegisterCounterHandler(s server.Server, hdlr CounterHandler, opts ...server.HandlerOption) error {
	type counter interface {
		Count(ctx context.Context, in *Request, out *Response) error
	}
	type Counter struct {
		counter
	}
	h := &counterHandler{hdlr}
	return s.Handle(s.NewHandler(&Counter{h}, opts...))
}

type counterHandler struct {
	CounterHandler
}

func (h *counterHandler) Count(ctx context.Context, in *Request, out *Response) error {
	return h.CounterHandler.Count(ctx, in, out)
}
