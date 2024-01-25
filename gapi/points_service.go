package gapi

import (
	"context"
	"database/sql"

	"github.com/AhmetSBulbul/quarterback-server/pb/pointspb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PointsService struct {
	db *sql.DB
	pointspb.UnimplementedPointsServiceServer
}

func (s *PointsService) GetPointsByDistrict(ctx context.Context, in *pointspb.QueryByDistrict) (*pointspb.PointsByDistrict, error) {
	query := `SELECT
	playerID,
	username,
	sum(isWin) as wins,
	sum(isLoss) as losses,
	sum(isCanceled) as cancelled,
	sum(isWin) * 2 + sum(isLoss) as totalPoints
	 from districtPoints
	 where districtID = ?
	 GROUP BY username
	 order by date desc
	`

	rows, err := s.db.Query(query, in.DistrictID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error while querying database: %v", err)
	}
	defer rows.Close()

	var pointsByDistrict pointspb.PointsByDistrict
	var userPointsArray []*pointspb.UserPoints
	for rows.Next() {
		var userPoints pointspb.UserPoints
		if err := rows.Scan(&userPoints.PlayerID,
			&userPoints.Username, &userPoints.Wins,
			&userPoints.Losses, &userPoints.Cancelled, &userPoints.TotalPoints); err != nil {
			return nil, status.Errorf(codes.Internal, "Error while scanning rows: %v", err)
		}
		userPointsArray = append(userPointsArray, &userPoints)
	}

	pointsByDistrict.UserPoints = userPointsArray
	return &pointsByDistrict, nil
}

func (s *PointsService) GetUserStats(ctx context.Context, in *pointspb.QueryByUser) (*pointspb.UserStats, error) {
	query := `
	SELECT 
	count(*) as games,
	sum(isWin) as wins,
	sum(isLoss) as losses,
	sum(isWin) * 2 + sum(isLoss) as totalPoints,
	0 as rank
	from districtPoints
	where playerID = ?
	group by playerID
	order by date desc`

	var userStats pointspb.UserStats
	if err := s.db.QueryRow(query, in.PlayerID).Scan(&userStats.TotalGames, &userStats.Wins, &userStats.Losses, &userStats.TotalPoints, &userStats.Rank); err != nil {
		return nil, status.Errorf(codes.Internal, "Error while scanning rows: %v", err)
	}

	return &userStats, nil
}
