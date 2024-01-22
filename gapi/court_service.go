package gapi

import (
	"context"
	"database/sql"

	"github.com/AhmetSBulbul/quarterback-server/pb/courtpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CourtService struct {
	db *sql.DB
	courtpb.UnimplementedCourtServiceServer
}

func (s *CourtService) GetCourt(ctx context.Context, in *courtpb.GetCourtRequest) (*courtpb.CourtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourt not implemented")
}
func (s *CourtService) ListCourtByLocation(ctx context.Context, in *courtpb.ListCourtByLocationRequest) (*courtpb.ListCourtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCourtByLocation not implemented")
}
func (s *CourtService) SearchCourt(ctx context.Context, in *courtpb.SearchCourtRequest) (*courtpb.ListCourtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchCourt not implemented")
}
func (s *CourtService) CreateCourt(ctx context.Context, in *courtpb.Court) (*courtpb.CourtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCourt not implemented")
}
