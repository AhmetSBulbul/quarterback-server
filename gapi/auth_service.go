package gapi

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/AhmetSBulbul/quarterback-server/helpers"
	"github.com/AhmetSBulbul/quarterback-server/pb/authpb"
	"github.com/AhmetSBulbul/quarterback-server/pb/commonpb"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
)

type AuthService struct {
	db *sql.DB
	authpb.UnimplementedAuthServiceServer
}

func (a *AuthService) genTokens(sub string, sub_id string) (string, string, error) {
	accessToken, err := helpers.CreateJWT(jwt.MapClaims{
		"sub":    sub,
		"sub_id": sub_id,
		"typ":    "access",
	})

	if err != nil {
		// log.Fatalf("token error: %v", err.Error())
		return "", "", gerr(codes.Internal, err)
	}

	refresh_jti, _ := uuid.NewRandom()
	refresh_jti_str := refresh_jti.String()
	refreshTTL := time.Hour * 24 * 7
	refreshToken, err := helpers.CreateJWT(jwt.MapClaims{
		"sub":    sub,
		"sub_id": sub_id,
		"exp":    time.Now().Add(refreshTTL).Unix(),
		"typ":    "refresh",
		"jti":    refresh_jti_str,
	})
	if err != nil {
		// log.Fatalf("token error: %v", err.Error())
		return "", "", gerr(codes.Internal, err)
	}

	// Bu burada olmayacak, refresh yapinca eklenecek...
	// ctx := context.Background()
	// err = s.redis.Set(ctx, sub_id, refresh_jti_str, refreshTTL).Err()
	// if err != nil {
	// 	log.Printf("[genTokens] redis error: %v", err.Error())
	// }

	return accessToken, refreshToken, nil
}

func (a *AuthService) Login(ctx context.Context, in *authpb.LoginRequest) (*authpb.Credentials, error) {
	var (
		id       int
		username string
		password string
	)

	row := a.db.QueryRow("SELECT id, username, password FROM user WHERE username = ? OR email = ?", in.Username, in.Username)
	err := row.Scan(&id, &username, &password)
	if err != nil {
		// log.Printf("[Login] scan error: %v", err.Error())
		return nil, gerr(codes.Unauthenticated, err)
	}

	if !helpers.ComparePasswords(password, in.Password) {
		return nil, gerr(codes.Unauthenticated, nil)
	}

	accessToken, refreshToken, err := a.genTokens(username, fmt.Sprintf("%d", id))
	if err != nil {
		return nil, err
	}

	return &authpb.Credentials{
		Token:        accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (a *AuthService) Register(_ context.Context, _ *authpb.RegisterRequest) (*authpb.Credentials, error) {
	panic("not implemented") // TODO: Implement
}

func (a *AuthService) Refresh(_ context.Context, _ *authpb.RefreshTokenRequest) (*authpb.Credentials, error) {
	panic("not implemented") // TODO: Implement
}

// Send reset token to email
func (a *AuthService) SendOtp(_ context.Context, _ *authpb.SendOtpRequest) (*commonpb.SuccessResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (a *AuthService) ResetPasswordWithOtp(_ context.Context, _ *authpb.ResetPasswordWithOtpRequest) (*commonpb.SuccessResponse, error) {
	panic("not implemented") // TODO: Implement
}