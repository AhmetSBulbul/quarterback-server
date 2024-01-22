package gapi

import (
	"context"
	"database/sql"

	"github.com/AhmetSBulbul/quarterback-server/pb/commonpb"
	"github.com/AhmetSBulbul/quarterback-server/pb/courtpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CourtService struct {
	db *sql.DB
	courtpb.UnimplementedCourtServiceServer
}

func (s *CourtService) GetCourt(ctx context.Context, in *courtpb.GetCourtRequest) (*courtpb.CourtResponse, error) {
	query := "SELECT ID, name, districtID, coordinate, address FROM court WHERE id = ?"
	court := &courtpb.Court{}

	err := s.db.QueryRowContext(ctx, query, in.GetId()).Scan(&court.Id, &court.Name, &court.DistrictId, &court.Location, &court.Address)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "court not found")
	}

	return &courtpb.CourtResponse{Court: court}, nil
}

func (s *CourtService) ListCourtByLocation(ctx context.Context, in *courtpb.ListCourtByLocationRequest) (*courtpb.ListCourtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCourtByLocation not implemented")
}

func (s *CourtService) SearchCourt(ctx context.Context, in *courtpb.SearchCourtRequest) (*courtpb.ListCourtResponse, error) {
	sub_id := getUserIdFromCtx(ctx)
	query := `SELECT
		c.ID,
		c.name,
		c.districtID,
		c.coordinate,
		c.address
	FROM court c
	INNER JOIN user u ON u.id = ?
	WHERE
		u.districtID = c.districtID`

	rows, err := s.db.QueryContext(
		ctx, query, sub_id,
	)

	if err != nil {
		return nil, gerr(codes.Internal, err)
	}

	defer rows.Close()

	var courts []*courtpb.Court

	for rows.Next() {
		var court courtpb.Court
		err := rows.Scan(&court.Id, &court.Name, &court.DistrictId, &court.Location, &court.Address)
		if err != nil {
			return nil, gerr(codes.Internal, err)
		}
		court.Media = make([]*commonpb.Media, 0)
		courts = append(courts, &court)
	}

	return &courtpb.ListCourtResponse{
		Courts:     courts,
		Pagination: &commonpb.PaginationResponse{},
	}, nil
}

func (s *CourtService) CreateCourt(ctx context.Context, in *courtpb.Court) (*courtpb.CourtResponse, error) {
	query := "INSERT INTO court (name, districtID, coordinate, address) VALUES (?, ?, ?, ?)"

	res, err := s.db.ExecContext(ctx, query, in.Name, in.DistrictId, in.Location, in.Address)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not create court")
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not create court")
	}

	return &courtpb.CourtResponse{Court: &courtpb.Court{Id: int32(id), Name: in.Name, DistrictId: in.DistrictId, Location: in.Location, Address: in.Address}}, nil
}
