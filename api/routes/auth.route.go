package routes

import (
	"context"

	"github.com/AhmetSBulbul/quarterback-server/api/pb/authpb"
	"github.com/AhmetSBulbul/quarterback-server/api/pb/commonpb"
)

type AuthService struct{}

func (a *AuthService) Login(_ context.Context, _ *authpb.LoginRequest) (*authpb.Credentials, error) {
	panic("not implemented") // TODO: Implement
}

func (a *AuthService) Register(_ context.Context, _ *authpb.RegisterRequest) (*authpb.Credentials, error) {
	panic("not implemented") // TODO: Implement
}

func (a *AuthService) Refresh(_ context.Context, _ *authpb.RefreshTokenRequest) (*authpb.Credentials, error) {
	panic("not implemented") // TODO: Implement
}

func (a *AuthService) SendOtp(_ context.Context, _ *authpb.SendOtpRequest) (*commonpb.SuccessResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (a *AuthService) ResetPasswordWithOtp(_ context.Context, _ *authpb.ResetPasswordWithOtpRequest) (*commonpb.SuccessResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (a *AuthService) mustEmbedUnimplementedAuthServiceServer() {
	panic("not implemented") // TODO: Implement
}
