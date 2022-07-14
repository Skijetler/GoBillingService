package biz

import (
	v1 "github.com/Skijetler/GoBillingService/auth/api/auth/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrUserNotFound      = status.Errorf(codes.NotFound, "reason: %v", v1.ServiceErrors_USER_NOT_FOUND.String())
	ErrUserAlreadyExists = status.Errorf(codes.AlreadyExists, "reason: %v", v1.ServiceErrors_USER_ALREADY_EXISTS.String())
	ErrWrongPassword     = status.Errorf(codes.PermissionDenied, "reason: %v", v1.ServiceErrors_WRONG_PASSWORD.String())
	ErrInvalidToken      = status.Errorf(codes.Unauthenticated, "reason: %v", v1.ServiceErrors_INVALID_TOKEN.String())
)

func internalErr(reason error) error {
	return status.Errorf(codes.Internal, "reason: %v", reason.Error())
}
