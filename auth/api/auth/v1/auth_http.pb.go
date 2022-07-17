// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.3.1

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationAuthSignUp = "/api.auth.v1.Auth/SignUp"
const OperationAuthSignIn = "/api.auth.v1.Auth/SignIn"
const OperationAuthRefreshToken = "/api.auth.v1.Auth/RefreshToken"

type AuthHTTPServer interface {
	RefreshToken(context.Context, *RefreshTokenRequest) (*RefreshTokenReply, error)
	SignIn(context.Context, *SignInRequest) (*SignInReply, error)
	SignUp(context.Context, *SignUpRequest) (*SignUpReply, error)
}

func RegisterAuthHTTPServer(s *http.Server, srv AuthHTTPServer) {
	r := s.Route("/")
	r.POST("/auth/sign-up", _Auth_SignUp0_HTTP_Handler(srv))
	r.POST("/auth/sign-in", _Auth_SignIn0_HTTP_Handler(srv))
	r.POST("/auth/refresh", _Auth_RefreshToken0_HTTP_Handler(srv))
}

func _Auth_SignUp0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SignUpRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthSignUp)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SignUp(ctx, req.(*SignUpRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SignUpReply)
		return ctx.Result(200, reply)
	}
}

func _Auth_SignIn0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SignInRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthSignIn)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SignIn(ctx, req.(*SignInRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SignInReply)
		return ctx.Result(200, reply)
	}
}

func _Auth_RefreshToken0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RefreshTokenRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAuthRefreshToken)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RefreshToken(ctx, req.(*RefreshTokenRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RefreshTokenReply)
		return ctx.Result(200, reply)
	}
}

type AuthHTTPClient interface {
	RefreshToken(ctx context.Context, req *RefreshTokenRequest, opts ...http.CallOption) (rsp *RefreshTokenReply, err error)
	SignIn(ctx context.Context, req *SignInRequest, opts ...http.CallOption) (rsp *SignInReply, err error)
	SignUp(ctx context.Context, req *SignUpRequest, opts ...http.CallOption) (rsp *SignUpReply, err error)
}

type AuthHTTPClientImpl struct {
	cc *http.Client
}

func NewAuthHTTPClient(client *http.Client) AuthHTTPClient {
	return &AuthHTTPClientImpl{client}
}

func (c *AuthHTTPClientImpl) RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...http.CallOption) (*RefreshTokenReply, error) {
	var out RefreshTokenReply
	pattern := "/auth/refresh"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAuthRefreshToken))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) SignIn(ctx context.Context, in *SignInRequest, opts ...http.CallOption) (*SignInReply, error) {
	var out SignInReply
	pattern := "/auth/sign-in"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAuthSignIn))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) SignUp(ctx context.Context, in *SignUpRequest, opts ...http.CallOption) (*SignUpReply, error) {
	var out SignUpReply
	pattern := "/auth/sign-up"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAuthSignUp))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}