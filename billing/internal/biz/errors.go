package biz

import (
	v1 "github.com/Skijetler/GoBillingService/billing/api/billing/v1/"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrAccountNotFound   = status.Errorf(codes.NotFound, "reason: %v", v1.ServiceErrors_ACCOUNT_NOT_FOUND.String())
	ErrAccountNotBelong  = status.Errorf(codes.Unauthenticated, "reason: %v", v1.ServiceErrors_ACCOUNT_NOT_BELONG.String())
	ErrInsufficientFunds = status.Errorf(codes.PermissionDenied, "reason: %v", v1.ServiceErrors_INSUFFICIENT_FUNDS.String())
)

func internalErr(reason error) error {
	return status.Errorf(codes.Internal, "reason: %v", reason.Error())
}
