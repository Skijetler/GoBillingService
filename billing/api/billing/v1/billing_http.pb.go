// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.3.1

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationBillingCreateAccount = "/api.billing.v1.Billing/CreateAccount"
const OperationBillingGetAccountInfo = "/api.billing.v1.Billing/GetAccountInfo"
const OperationBillingGetAllAccounts = "/api.billing.v1.Billing/GetAllAccounts"
const OperationBillingDeleteAccount = "/api.billing.v1.Billing/DeleteAccount"
const OperationBillingTransferFunds = "/api.billing.v1.Billing/TransferFunds"
const OperationBillingShowHistory = "/api.billing.v1.Billing/ShowHistory"

type BillingHTTPServer interface {
	CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountReply, error)
	DeleteAccount(context.Context, *DeleteAccountRequest) (*DeleteAccountReply, error)
	GetAccountInfo(context.Context, *GetAccountInfoRequest) (*GetAccountInfoReply, error)
	GetAllAccounts(context.Context, *emptypb.Empty) (*GetAllAccountsReply, error)
	ShowHistory(context.Context, *ShowHistoryRequest) (*ShowHistoryReply, error)
	TransferFunds(context.Context, *TransferFundsRequest) (*TransferFundsReply, error)
}

func RegisterBillingHTTPServer(s *http.Server, srv BillingHTTPServer) {
	r := s.Route("/")
	r.POST("/billing/account/create", _Billing_CreateAccount0_HTTP_Handler(srv))
	r.GET("/billing/account", _Billing_GetAccountInfo0_HTTP_Handler(srv))
	r.GET("/billing/account/all", _Billing_GetAllAccounts0_HTTP_Handler(srv))
	r.POST("/billing/account/delete", _Billing_DeleteAccount0_HTTP_Handler(srv))
	r.POST("/billing/transfer", _Billing_TransferFunds0_HTTP_Handler(srv))
	r.GET("/billing/history", _Billing_ShowHistory0_HTTP_Handler(srv))
}

func _Billing_CreateAccount0_HTTP_Handler(srv BillingHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateAccountRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBillingCreateAccount)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateAccount(ctx, req.(*CreateAccountRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateAccountReply)
		return ctx.Result(200, reply)
	}
}

func _Billing_GetAccountInfo0_HTTP_Handler(srv BillingHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetAccountInfoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBillingGetAccountInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetAccountInfo(ctx, req.(*GetAccountInfoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetAccountInfoReply)
		return ctx.Result(200, reply)
	}
}

func _Billing_GetAllAccounts0_HTTP_Handler(srv BillingHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBillingGetAllAccounts)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetAllAccounts(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetAllAccountsReply)
		return ctx.Result(200, reply)
	}
}

func _Billing_DeleteAccount0_HTTP_Handler(srv BillingHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteAccountRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBillingDeleteAccount)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteAccount(ctx, req.(*DeleteAccountRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteAccountReply)
		return ctx.Result(200, reply)
	}
}

func _Billing_TransferFunds0_HTTP_Handler(srv BillingHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TransferFundsRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBillingTransferFunds)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.TransferFunds(ctx, req.(*TransferFundsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TransferFundsReply)
		return ctx.Result(200, reply)
	}
}

func _Billing_ShowHistory0_HTTP_Handler(srv BillingHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ShowHistoryRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBillingShowHistory)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ShowHistory(ctx, req.(*ShowHistoryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ShowHistoryReply)
		return ctx.Result(200, reply)
	}
}

type BillingHTTPClient interface {
	CreateAccount(ctx context.Context, req *CreateAccountRequest, opts ...http.CallOption) (rsp *CreateAccountReply, err error)
	DeleteAccount(ctx context.Context, req *DeleteAccountRequest, opts ...http.CallOption) (rsp *DeleteAccountReply, err error)
	GetAccountInfo(ctx context.Context, req *GetAccountInfoRequest, opts ...http.CallOption) (rsp *GetAccountInfoReply, err error)
	GetAllAccounts(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *GetAllAccountsReply, err error)
	ShowHistory(ctx context.Context, req *ShowHistoryRequest, opts ...http.CallOption) (rsp *ShowHistoryReply, err error)
	TransferFunds(ctx context.Context, req *TransferFundsRequest, opts ...http.CallOption) (rsp *TransferFundsReply, err error)
}

type BillingHTTPClientImpl struct {
	cc *http.Client
}

func NewBillingHTTPClient(client *http.Client) BillingHTTPClient {
	return &BillingHTTPClientImpl{client}
}

func (c *BillingHTTPClientImpl) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...http.CallOption) (*CreateAccountReply, error) {
	var out CreateAccountReply
	pattern := "/billing/account/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBillingCreateAccount))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BillingHTTPClientImpl) DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...http.CallOption) (*DeleteAccountReply, error) {
	var out DeleteAccountReply
	pattern := "/billing/account/delete"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBillingDeleteAccount))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BillingHTTPClientImpl) GetAccountInfo(ctx context.Context, in *GetAccountInfoRequest, opts ...http.CallOption) (*GetAccountInfoReply, error) {
	var out GetAccountInfoReply
	pattern := "/billing/account"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBillingGetAccountInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BillingHTTPClientImpl) GetAllAccounts(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*GetAllAccountsReply, error) {
	var out GetAllAccountsReply
	pattern := "/billing/account/all"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBillingGetAllAccounts))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BillingHTTPClientImpl) ShowHistory(ctx context.Context, in *ShowHistoryRequest, opts ...http.CallOption) (*ShowHistoryReply, error) {
	var out ShowHistoryReply
	pattern := "/billing/history"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBillingShowHistory))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *BillingHTTPClientImpl) TransferFunds(ctx context.Context, in *TransferFundsRequest, opts ...http.CallOption) (*TransferFundsReply, error) {
	var out TransferFundsReply
	pattern := "/billing/transfer"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBillingTransferFunds))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
