package gapi

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/AhmetSBulbul/quarterback-server/pb/commonpb"
	"github.com/AhmetSBulbul/quarterback-server/pb/courtpb"
	"github.com/AhmetSBulbul/quarterback-server/pb/regionpb"
	"github.com/AhmetSBulbul/quarterback-server/pb/userpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CourtService struct {
	db *sql.DB
	courtpb.UnimplementedCourtServiceServer
}

// Repository olusturmaya usendim .s
func (s *CourtService) getUserByID(ctx context.Context, userid int) (*userpb.User, error) {
	var user userpb.User
	var avatarID sql.NullString

	query := "SELECT u.id, u.email, u.districtID, u.name, u.lastName, u.username, f.path FROM user as u LEFT JOIN file f ON f.id = u.avatarID WHERE u.id = ?"
	row := s.db.QueryRowContext(ctx, query, userid)
	err := row.Scan(&user.Id, &user.Email, &user.DistrictID, &user.Name, &user.Lastname, &user.Username, &avatarID)

	user.AvatarPath = avatarID.String

	if err != nil {
		return nil, gerr(codes.Internal, err)
	}

	return &user, nil
}

func (s *CourtService) GetCourt(ctx context.Context, in *courtpb.GetCourtRequest) (*courtpb.CourtResponse, error) {
	// TODO: get checkins at different endpoint
	// This code is used on game service
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
		c.ID = ?
	GROUP BY c.id;`
	// TODO: no need to group by!
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

func (s *CourtService) ListCourtByLocation(ctx context.Context, in *courtpb.ListCourtByLocationRequest) (*courtpb.CourtsWithDistance, error) {
	query := `select 
    c.id,
	ST_Y(c.coordinate) as latitude,
	ST_X(c.coordinate) as longitude,
    c.name,
    c.address,
    d.name as district,
	d.id as districtID,
	d.cityID as cityID,
    ST_Distance_Sphere(
        ST_GeomFromText(?),
        c.coordinate
	) as distance
	from court c 
	inner join district d on d.id = c.districtID 
	ORDER BY distance;`

	lat := in.GetLocation().GetLatitude()
	lng := in.GetLocation().GetLongitude()

	point := fmt.Sprintf("POINT(%f %f)", lat, lng)

	rows, err := s.db.QueryContext(ctx, query, point)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not list courts")
	}
	defer rows.Close()

	courts := []*courtpb.CourtWithDistance{}
	for rows.Next() {
		cwd := &courtpb.CourtWithDistance{}
		court := &courtpb.Court{}
		dst := &regionpb.District{}
		loc := &commonpb.Location{}
		err := rows.Scan(&court.Id, &loc.Latitude, &loc.Longitude, &court.Name, &court.Address, &dst.Name, &dst.Id, &dst.CityId, &cwd.Distance)
		court.District = dst
		court.Location = loc
		cwd.Court = court
		if err != nil {
			return nil, status.Errorf(codes.Internal, "could not list courts")
		}
		courts = append(courts, cwd)
	}

	return &courtpb.CourtsWithDistance{Courts: courts}, nil
}

func (s *CourtService) SearchCourt(ctx context.Context, in *courtpb.SearchCourtRequest) (*courtpb.ListCourtResponse, error) {
	sub_id := getUserIdFromCtx(ctx)
	query := `SELECT
		c.ID,
		ST_Y(c.coordinate) as latitude,
		ST_X(c.coordinate) as longitude,
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
	LEFT JOIN checkin on checkin.courtID = c.ID AND checkin.createdAt > DATE_SUB(CURRENT_TIMESTAMP, INTERVAL 20 MINUTE)
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
		loc := &commonpb.Location{}
		err := rows.Scan(&court.Id, &loc.Latitude, &loc.Longitude, &court.Name, &district.Id, &district.Name, &district.CityId, &court.Address, &court.CommentCount, &court.CheckInCount)
		if err != nil {
			return nil, gerr(codes.Internal, err)
		}
		court.District = &district
		court.Location = loc
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
	// TODO: check if sender has any other checkin in last 30 min
	sub_id := getUserIdFromCtx(ctx)
	query := "INSERT INTO checkin (userID, courtID) VALUES (?, ?)"

	result, err := s.db.ExecContext(ctx, query, sub_id, in.Id)
	if err != nil {
		return nil, gerr(codes.Internal, err)
	}

	checkInId, err := result.LastInsertId()
	if err != nil {
		return nil, gerr(codes.Internal, err)
	}

	return &courtpb.CheckInCourtResponse{
		Id:        int32(checkInId),
		CheckedIn: true,
	}, nil
}

func (s *CourtService) AddComment(ctx context.Context, req *courtpb.CourtCommentRequest) (*courtpb.CourtComment, error) {
	sub_id := getUserIdFromCtx(ctx)

	query := "INSERT INTO court_comment (senderID, courtID, content) VALUES (?, ?, ?)"
	result, err := s.db.ExecContext(ctx, query, sub_id, req.CourtId, req.Content)
	if err != nil {
		return nil, gerr(codes.Internal, err)
	}

	commentId, err := result.LastInsertId()
	if err != nil {
		return nil, gerr(codes.Internal, err)
	}

	sender, err := s.getUserByID(ctx, sub_id)
	if err != nil {
		return nil, gerr(codes.Internal, err)
	}

	var comment courtpb.CourtComment

	query = "SELECT com.id, com.courtID, com.content FROM court_comment as com WHERE com.ID = ?"
	row := s.db.QueryRowContext(ctx, query, commentId)

	err = row.Scan(&comment.Id, &comment.CourtId, &comment.Content)
	if err != nil {
		return nil, gerr(codes.Internal, err)
	}

	comment.Sender = sender

	return &comment, nil

}

func (s *CourtService) ListComment(ctx context.Context, in *courtpb.CourtCommentListRequest) (*courtpb.CourtCommentListResponse, error) {
	query := "SELECT com.id, com.courtID, com.senderID, com.content FROM court_comment as com WHERE com.courtID = ?"
	rows, err := s.db.QueryContext(ctx, query, in.CourtId)
	if err != nil {
		return nil, gerr(codes.Internal, err)
	}

	defer rows.Close()

	var comments []*courtpb.CourtComment

	for rows.Next() {
		var senderId int
		var comment courtpb.CourtComment

		err := rows.Scan(&comment.Id, &comment.CourtId, &senderId, &comment.Content)
		if err != nil {
			return nil, gerr(codes.Internal, err)
		}

		sender, err := s.getUserByID(ctx, senderId)
		if err != nil {
			return nil, gerr(codes.Internal, err)
		}
		comment.Sender = sender
		comments = append(comments, &comment)
	}
	return &courtpb.CourtCommentListResponse{
		Comments:   comments,
		Pagination: &commonpb.PaginationResponse{},
	}, nil
}

// select * from checkin where createdAt > DATE_SUB(CURRENT_TIMESTAMP, INTERVAL 20 MINUTE);
// select u.id, u.email, u.districtID, u.name, u.lastName, u.username, f.path, com.content from court_comment as com inner join user u on u.id = com.senderID left join file f on u.avatarID = f.ID where courtID = 1;

// TODO:
// AddComment
// ListComment
// Checkins

// SELECT c.ID, c.name, d.ID, d.name, d.cityID, c.address, c.coordinate, COUNT(com.courtID) as commentCount, COUNT(checkin.courtID) as checkinCount FROM court as c LEFT JOIN district d ON c.districtID = d.ID LEFT JOIN court_comment com on com.courtID = c.ID LEFT JOIN checkin ON checkin.courtID = c.ID INNER JOIN user u ON u.id = 2 WHERE c.districtID = u.districtID group by c.id;
