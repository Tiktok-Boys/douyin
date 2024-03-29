// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/userService.proto

package userService

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for UserService service

func NewUserServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for UserService service

type UserService interface {
	Register(ctx context.Context, in *DouyinUserRegisterRequest, opts ...client.CallOption) (*DouyinUserRegisterResponse, error)
	Login(ctx context.Context, in *DouyinUserLoginRequest, opts ...client.CallOption) (*DouyinUserLoginResponse, error)
	UserInfo(ctx context.Context, in *DouyinUserRequest, opts ...client.CallOption) (*DouyinUserResponse, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) Register(ctx context.Context, in *DouyinUserRegisterRequest, opts ...client.CallOption) (*DouyinUserRegisterResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.Register", in)
	out := new(DouyinUserRegisterResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Login(ctx context.Context, in *DouyinUserLoginRequest, opts ...client.CallOption) (*DouyinUserLoginResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.Login", in)
	out := new(DouyinUserLoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UserInfo(ctx context.Context, in *DouyinUserRequest, opts ...client.CallOption) (*DouyinUserResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.UserInfo", in)
	out := new(DouyinUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	Register(context.Context, *DouyinUserRegisterRequest, *DouyinUserRegisterResponse) error
	Login(context.Context, *DouyinUserLoginRequest, *DouyinUserLoginResponse) error
	UserInfo(context.Context, *DouyinUserRequest, *DouyinUserResponse) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) error {
	type userService interface {
		Register(ctx context.Context, in *DouyinUserRegisterRequest, out *DouyinUserRegisterResponse) error
		Login(ctx context.Context, in *DouyinUserLoginRequest, out *DouyinUserLoginResponse) error
		UserInfo(ctx context.Context, in *DouyinUserRequest, out *DouyinUserResponse) error
	}
	type UserService struct {
		userService
	}
	h := &userServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&UserService{h}, opts...))
}

type userServiceHandler struct {
	UserServiceHandler
}

func (h *userServiceHandler) Register(ctx context.Context, in *DouyinUserRegisterRequest, out *DouyinUserRegisterResponse) error {
	return h.UserServiceHandler.Register(ctx, in, out)
}

func (h *userServiceHandler) Login(ctx context.Context, in *DouyinUserLoginRequest, out *DouyinUserLoginResponse) error {
	return h.UserServiceHandler.Login(ctx, in, out)
}

func (h *userServiceHandler) UserInfo(ctx context.Context, in *DouyinUserRequest, out *DouyinUserResponse) error {
	return h.UserServiceHandler.UserInfo(ctx, in, out)
}
