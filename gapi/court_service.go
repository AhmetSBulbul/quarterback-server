package gapi

import (
	"context"
	"database/sql"

	"github.com/AhmetSBulbul/quarterback-server/pb/commonpb"
	"github.com/AhmetSBulbul/quarterback-server/pb/courtpb"
	"github.com/AhmetSBulbul/quarterback-server/pb/regionpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CourtService struct {
	db *sql.DB
	courtpb.UnimplementedCourtServiceServer
}

func (s *CourtService) GetCourt(ctx context.Context, in *courtpb.GetCourtRequest) (*courtpb.CourtResponse, error) {
	query := `SELECT
		c.ID,
		c.name,
		d.ID,
		d.name,
		d.cityID,
		c.address,
		COUNT(com.courtID) as commentCount,
		COUNT(checkin.courtID) as checkinCount
	FROM court c
	INNER JOIN district d ON c.districtID = d.ID
	LEFT JOIN court_comment com on com.courtID = c.ID
	LEFT JOIN checkin on checkin.courtID = c.ID
	WHERE
		c.districtId = ?
	GROUP BY c.id;`
	court := &courtpb.Court{}
	district := &regionpb.District{}

	err := s.db.QueryRowContext(ctx, query, in.GetId()).Scan(&court.Id, &court.Name, &district.Id, &district.Name, &district.CityId, &court.Address, &court.CommentCount, &court.CheckInCount)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "court not found")
	}

	court.District = district
	court.Media = make([]*commonpb.Media, 0)

	return &courtpb.CourtResponse{Court: court}, nil
}

func (s *CourtService) ListCourtByLocation(ctx context.Context, in *courtpb.ListCourtByLocationRequest) (*courtpb.ListCourtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCourtByLocation not implemented")
}

func (s *CourtService) SearchCourt(ctx context.Context, in *courtpb.SearchCourtRequest) (*courtpb.ListCourtResponse, error) {
	sub_id := getUserIdFromCtx(ctx)
	// TODO: coordinate serialization'i ekler misin buraya
	query := `SELECT
		c.ID,
		c.name,
		d.ID,
		d.name,
		d.cityID,
		c.address,
		COUNT(com.courtID) as commentCount,
		COUNT(checkin.courtID) as checkinCount
	FROM court c
	INNER JOIN district d ON c.districtID = d.ID
	INNER JOIN user u ON u.id = ?
	LEFT JOIN court_comment com on com.courtID = c.ID
	LEFT JOIN checkin on checkin.courtID = c.ID
	WHERE
		u.districtID = c.districtID
	GROUP BY c.id;`

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
		var district regionpb.District
		err := rows.Scan(&court.Id, &court.Name, &district.Id, &district.Name, &district.CityId, &court.Address, &court.CommentCount, &court.CheckInCount)
		if err != nil {
			return nil, gerr(codes.Internal, err)
		}
		court.District = &district
		court.Media = make([]*commonpb.Media, 0)
		courts = append(courts, &court)
	}

	return &courtpb.ListCourtResponse{
		Courts:     courts,
		Pagination: &commonpb.PaginationResponse{},
	}, nil
}

func (s *CourtService) CreateCourt(ctx context.Context, in *courtpb.Court) (*courtpb.CourtResponse, error) {
	// query := "INSERT INTO court (name, districtID, coordinate, address) VALUES (?, ?, ?, ?)"

	// res, err := s.db.ExecContext(ctx, query, in.Name, in.DistrictId, in.Location, in.Address)
	// if err != nil {
	// 	return nil, status.Errorf(codes.Internal, "could not create court")
	// }

	// id, err := res.LastInsertId()
	// if err != nil {
	// 	return nil, status.Errorf(codes.Internal, "could not create court")
	// }

	// return &courtpb.CourtResponse{Court: &courtpb.Court{Id: int32(id), Name: in.Name, DistrictId: in.DistrictId, Location: in.Location, Address: in.Address}}, nil
	panic("not implemented")
}

func (s *CourtService) CheckInCourt(ctx context.Context, in *commonpb.GetByIdRequest) (*courtpb.CheckInCourtResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *CourtService) mustEmbedUnimplementedCourtServiceServer() {
	panic("not implemented") // TODO: Implement
}

// select * from checkin where createdAt > DATE_SUB(CURRENT_TIMESTAMP, INTERVAL 20 MINUTE);
// select u.id, u.email, u.districtID, u.name, u.lastName, u.username, f.path, com.content from court_comment as com inner join user u on u.id = com.senderID left join file f on u.avatarID = f.ID where courtID = 1;

// TODO:
// AddComment
// ListComment
// Checkins

// SELECT c.ID, c.name, d.ID, d.name, d.cityID, c.address, c.coordinate, COUNT(com.courtID) as commentCount, COUNT(checkin.courtID) as checkinCount FROM court as c LEFT JOIN district d ON c.districtID = d.ID LEFT JOIN court_comment com on com.courtID = c.ID LEFT JOIN checkin ON checkin.courtID = c.ID INNER JOIN user u ON u.id = 2 WHERE c.districtID = u.districtID group by c.id;
