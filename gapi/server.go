package gapi

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
	"strings"

	"github.com/AhmetSBulbul/quarterback-server/helpers"
	"github.com/AhmetSBulbul/quarterback-server/pb/authpb"
	"github.com/AhmetSBulbul/quarterback-server/pb/chatpb"
	"github.com/AhmetSBulbul/quarterback-server/pb/courtpb"
	"github.com/AhmetSBulbul/quarterback-server/pb/filepb"
	"github.com/AhmetSBulbul/quarterback-server/pb/regionpb"
	"github.com/AhmetSBulbul/quarterback-server/pb/userpb"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var validate *validator.Validate

func gerr(c codes.Code, err error) error {
	if err != nil {
		log.Printf("%v", err)
		return status.Errorf(c, err.Error())
	}
	return status.Errorf(c, "unknown error")
}

type Sub_id string

func unaryInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	allowedMethods := []string{
		"^/auth.AuthService/.+$",
		"^/region.RegionService/.+$",
	}

	// Check for allowed methods without authentication
	for _, method := range allowedMethods {
		if match, _ := regexp.MatchString(method, info.FullMethod); match {
			return handler(ctx, req)
		}
	}

	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return nil, gerr(codes.InvalidArgument, nil)
	}

	// Check for jwt token
	id, ok := valid(md["authorization"])
	if !ok {
		return nil, gerr(codes.Unauthenticated, nil)
	}

	m, err := handler(context.WithValue(ctx, Sub_id("sub_id"), id), req)
	if err != nil {
		log.Printf("RPC failed with error: %v", err)
	}

	return m, err
}

type customServerStream struct {
	grpc.ServerStream
	ctx context.Context
}

func streamInterceptor(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	md, ok := metadata.FromIncomingContext(ss.Context())
	// TODO: doesn't work

	if !ok {
		return gerr(codes.InvalidArgument, fmt.Errorf("Stram Interceptor Context Error"))
	}

	id, ok := valid(md["authorization"])
	if !ok {
		return gerr(codes.Unauthenticated, fmt.Errorf("Error on valid check at stream interceptor"))
	}

	ctx := context.WithValue(ss.Context(), Sub_id("sub_id"), id)
	newStream := &customServerStream{
		ServerStream: ss,
		ctx:          ctx,
	}

	return handler(srv, newStream)

}

func valid(authorization []string) (int, bool) {
	if len(authorization) < 1 {
		return 0, false
	}
	token := strings.TrimPrefix(authorization[0], "Bearer ")
	jwt_token, err := helpers.ValidateJWT(token)

	if err != nil {
		fmt.Println(err)
		return 0, false
	}

	claims := jwt_token.Claims.(jwt.MapClaims)

	fsub_id, ok := claims["sub_id"].(float64)
	if !ok {
		return 0, false
	}
	sub_id := int(fsub_id)

	return sub_id, true
}

func NewServer(db *sql.DB) {
	validate = validator.New()
	s := grpc.NewServer(grpc.UnaryInterceptor(unaryInterceptor), grpc.StreamInterceptor(streamInterceptor))

	authpb.RegisterAuthServiceServer(s, &AuthService{db: db, validate: validate})
	regionpb.RegisterRegionServiceServer(s, &RegionService{db: db})
	userpb.RegisterUserServiceServer(s, &UserService{db: db})
	filepb.RegisterFileServiceServer(s, &FileService{db: db})
	courtpb.RegisterCourtServiceServer(s, &CourtService{db: db})
	chatpb.RegisterChatServiceServer(s, &ChatService{
		channel: make(map[int32]chan *chatpb.ChatMessage),
		db:      db,
	})

	lis, err := net.Listen("tcp", os.Getenv("LISTEN_ADDR"))
	if err != nil {
		panic(err)
	}

	log.Printf("listening on %s", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
