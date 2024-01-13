package gapi

import (
	"database/sql"
	"log"
	"net"
	"os"

	"github.com/AhmetSBulbul/quarterback-server/pb/authpb"
	"github.com/AhmetSBulbul/quarterback-server/pb/regionpb"
	"github.com/go-playground/validator/v10"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
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

func NewServer(db *sql.DB) {
	validate = validator.New()
	s := grpc.NewServer()

	authpb.RegisterAuthServiceServer(s, &AuthService{db: db, validate: validate})
	regionpb.RegisterRegionServiceServer(s, &RegionService{db: db})

	lis, err := net.Listen("tcp", os.Getenv("LISTEN_ADDR"))
	if err != nil {
		panic(err)
	}

	log.Printf("listening on %s", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
