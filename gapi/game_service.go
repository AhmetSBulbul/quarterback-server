package gapi

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/AhmetSBulbul/quarterback-server/pb/commonpb"
	"github.com/AhmetSBulbul/quarterback-server/pb/courtpb"
	"github.com/AhmetSBulbul/quarterback-server/pb/gamepb"
	"github.com/AhmetSBulbul/quarterback-server/pb/regionpb"
	"github.com/AhmetSBulbul/quarterback-server/pb/userpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GameService struct {
	db *sql.DB
	gamepb.UnimplementedGameServiceServer
}

func (s *GameService) CreateGame(ctx context.Context, in *gamepb.CreateGameRequest) (*gamepb.GameIdResponse, error) {
	sub_id := getUserIdFromCtx(ctx)
	query := "INSERT INTO game (courtID, homeScore, awayScore) VALUES (?, 0, 0);"

	res, err := s.db.Exec(query, in.GetCourtId())
	if err != nil {
		fmt.Println(err)
		return nil, gerr(codes.Internal, err)
	}

	gameID, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err)

		return nil, gerr(codes.Internal, err)
	}

	query = `INSERT INTO game_player 
	  (gameID, playerID, isHomeSide)
	  VALUES (
		  ?,
		  ?,
		  1
		);`

	_, err = s.db.Exec(query, gameID, sub_id)

	if err != nil {
		fmt.Println(err)

		return nil, gerr(codes.Internal, err)
	}

	return &gamepb.GameIdResponse{
		GameId: int32(gameID),
	}, nil
}

func (s *GameService) JoinGame(ctx context.Context, in *gamepb.JoinGameRequest) (*gamepb.GameIdResponse, error) {
	sub_id := getUserIdFromCtx(ctx)

	query := `SELECT startedAt, endedAt FROM game WHERE id = ?;`

	var endedAt sql.NullTime
	var startedAt sql.NullTime
	s.db.QueryRow(query, in.GetGameId()).Scan(&startedAt, &endedAt)

	if startedAt.Valid {
		return nil, status.Errorf(codes.InvalidArgument, "Game already started")
	}

	if endedAt.Valid {
		return nil, status.Errorf(codes.InvalidArgument, "Game has ended")
	}

	query = `INSERT INTO game_player 
	(gameID, playerID, isHomeSide)
	VALUES (
		?,
		?,
		?
	  );`

	_, err := s.db.Exec(query, in.GetGameId(), sub_id, in.GetIsHomeSide())

	if err != nil {
		return nil, gerr(codes.Internal, err)
	}

	return &gamepb.GameIdResponse{
		GameId: in.GetGameId(),
	}, nil
}

func (s *GameService) StartGame(ctx context.Context, in *gamepb.StartGameRequest) (*gamepb.GameIdResponse, error) {
	// TODO: check player counts and sub_id is player of game
	query := `SELECT startedAt, endedAt FROM game WHERE id = ?;`

	var endedAt sql.NullTime
	var startedAt sql.NullTime
	s.db.QueryRow(query, in.GetGameId()).Scan(&startedAt, &endedAt)

	if startedAt.Valid {
		return nil, status.Errorf(codes.InvalidArgument, "Game already started")
	}

	if endedAt.Valid {
		return nil, status.Errorf(codes.InvalidArgument, "Game has ended")
	}

	query = `UPDATE game SET startedAt = NOW() WHERE id = ?;`

	_, err := s.db.Exec(query, in.GetGameId())
	if err != nil {
		fmt.Println(err)
		return nil, status.Errorf(codes.Internal, "Internal error")
	}

	return &gamepb.GameIdResponse{
		GameId: in.GetGameId(),
	}, nil
}

func (s *GameService) EndGame(ctx context.Context, in *gamepb.EndGameRequest) (*gamepb.GameIdResponse, error) {
	query := `SELECT endedAt FROM game WHERE id = ?;`

	var endedAt sql.NullTime
	s.db.QueryRow(query, in.GetGameId()).Scan(&endedAt)

	if endedAt.Valid {
		return nil, status.Errorf(codes.InvalidArgument, "Game has ended")
	}

	query = `UPDATE game SET endedAt = NOW(), homeScore = ?, awayScore = ? WHERE id = ?;`

	_, err := s.db.Exec(query, in.GetHomeScore(), in.GetAwayScore(), in.GetGameId())

	if err != nil {
		fmt.Println(err)
		return nil, gerr(codes.Internal, err)
	}

	return &gamepb.GameIdResponse{
		GameId: in.GetGameId(),
	}, nil
}

func (s *GameService) CancelGame(ctx context.Context, in *gamepb.CancelGameRequest) (*gamepb.GameIdResponse, error) {
	query := `SELECT endedAt FROM game WHERE id = ?;`

	var endedAt sql.NullTime
	s.db.QueryRow(query, in.GetGameId()).Scan(&endedAt)

	if !endedAt.Valid {
		return nil, status.Errorf(codes.InvalidArgument, "Game has not ended")
	}

	query = "UPDATE game_player SET isCanceled = 1 WHERE gameID = ?;"
	_, err := s.db.ExecContext(ctx, query, in.GetGameId())
	if err != nil {
		fmt.Println(err)
		return nil, gerr(codes.Internal, err)
	}

	return &gamepb.GameIdResponse{
		GameId: in.GetGameId(),
	}, nil
}

func (s *GameService) GetGame(ctx context.Context, in *gamepb.GetGameRequest) (*gamepb.GameResponse, error) {
	game, err := s.getGameById(ctx, in.GetGameId())
	if err != nil {
		fmt.Println(err)
		return nil, gerr(codes.Internal, err)
	}

	return &gamepb.GameResponse{
		Game: game,
	}, nil
}

func (s *GameService) ListGamesByLocation(ctx context.Context, in *gamepb.ListGamesByLocationRequest) (*gamepb.ListGamesResponse, error) {
	// TODO: add pagination for later phases
	// Gets games by distance
	// Not started games will not be listed, they can only reachable by id
	// For now they will be listed due to testing purposes

	// TODO: production query
	// query := `SELECT
	// 	g.id,
	// 	ST_Distance_Sphere(
	//     	ST_GeomFromText(?),
	//     	c.coordinate
	// 	) as distance
	// FROM game g
	// INNER JOIN court c ON g.courtID = c.id
	// WHERE g.startedAt IS NOT NULL
	// ORDER BY distance;`

	query := `SELECT
		g.id,
		ST_Distance_Sphere(
        	ST_GeomFromText(?),
        	c.coordinate
		) as distance
	FROM game g
	INNER JOIN court c ON g.courtID = c.id
	ORDER BY distance;`

	lat := in.GetLocation().GetLatitude()
	lng := in.GetLocation().GetLongitude()

	point := fmt.Sprintf("POINT(%f %f)", lat, lng)

	rows, err := s.db.QueryContext(ctx, query, point)
	if err != nil {
		fmt.Println(err)
		return nil, gerr(codes.Internal, err)
	}

	defer rows.Close()

	var games []*gamepb.Game
	for rows.Next() {
		var id int32
		var distance float64

		err := rows.Scan(&id, &distance)
		if err != nil {
			fmt.Println(err)
			return nil, gerr(codes.Internal, err)
		}

		game, err := s.getGameById(ctx, id)
		if err != nil {
			fmt.Println(err)
			return nil, gerr(codes.Internal, err)
		}

		games = append(games, game)
	}

	return &gamepb.ListGamesResponse{
		Games: games,
		// TODO: add pagination
		Pagination: &commonpb.PaginationResponse{},
	}, nil
}

func (s *GameService) ListGamesByUser(ctx context.Context, in *gamepb.ListGamesByUserRequest) (*gamepb.ListGamesResponse, error) {
	query := `SELECT
		g.id
	FROM game_player gp
	INNER JOIN game g ON gp.gameID = g.id
	WHERE gp.playerID = ? AND g.startedAt IS NOT NULL
	ORDER BY g.startedAt DESC;`

	rows, err := s.db.QueryContext(ctx, query, in.GetUserId())
	if err != nil {
		fmt.Println(err)
		return nil, gerr(codes.Internal, err)
	}

	defer rows.Close()

	var games []*gamepb.Game
	for rows.Next() {
		var id int32

		err := rows.Scan(&id)
		if err != nil {
			fmt.Println(err)
			return nil, gerr(codes.Internal, err)
		}

		game, err := s.getGameById(ctx, id)
		if err != nil {
			fmt.Println(err)
			return nil, gerr(codes.Internal, err)
		}

		games = append(games, game)
	}

	return &gamepb.ListGamesResponse{
		Games:      games,
		Pagination: &commonpb.PaginationResponse{},
	}, nil

}

func (s *GameService) ListGamesByCourt(ctx context.Context, in *gamepb.ListGamesByCourtRequest) (*gamepb.ListGamesResponse, error) {
	query := `SELECT
		g.id,
	FROM game g
	INNER JOIN court c ON g.courtID = c.id
	WHERE c.id = ? AND g.startedAt IS NOT NULL
	ORDER BY g.startedAt DESC;`

	rows, err := s.db.QueryContext(ctx, query, in.GetCourtId())
	if err != nil {
		fmt.Println(err)
		return nil, gerr(codes.Internal, err)
	}

	defer rows.Close()

	var games []*gamepb.Game
	for rows.Next() {
		var id int32

		err := rows.Scan(&id)
		if err != nil {
			fmt.Println(err)
			return nil, gerr(codes.Internal, err)
		}

		game, err := s.getGameById(ctx, id)
		if err != nil {
			fmt.Println(err)
			return nil, gerr(codes.Internal, err)
		}

		games = append(games, game)
	}

	return &gamepb.ListGamesResponse{
		Games:      games,
		Pagination: &commonpb.PaginationResponse{},
	}, nil
}

// If everyone leaves game, game will be deleted
func (s *GameService) LeaveGame(ctx context.Context, in *gamepb.LeaveGameRequest) (*gamepb.GameIdResponse, error) {
	query := `SELECT startedAt FROM game WHERE id = ?;`

	var startedAt sql.NullTime
	s.db.QueryRow(query, in.GetGameId()).Scan(&startedAt)

	if startedAt.Valid {
		return nil, status.Errorf(codes.InvalidArgument, "Game has ended")
	}

	sub_id := getUserIdFromCtx(ctx)

	query = `DELETE FROM game_player WHERE gameID = ? AND playerID = ?;`

	_, err := s.db.Exec(query, in.GetGameId(), sub_id)
	if err != nil {
		fmt.Println(err)
		return nil, gerr(codes.Internal, err)
	}

	return &gamepb.GameIdResponse{
		GameId: in.GetGameId(),
	}, nil
}

func (s *GameService) ListGamesByTeam(_ context.Context, _ *gamepb.ListGamesByTeamRequest) (*gamepb.ListGamesResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (s *GameService) CreateGameWithTeam(_ context.Context, _ *gamepb.CreateGameWithTeamRequest) (*gamepb.GameIdResponse, error) {
	panic("not implemented") // TODO: Implement
}

// Add Comment
func (s *GameService) AddMedia(_ context.Context, _ *commonpb.File) (*commonpb.Media, error) {
	panic("not implemented") // TODO: Implement
}

func (s *GameService) mustEmbedUnimplementedGameServiceServer() {
	panic("not implemented") // TODO: Implement
}

// End of generated code

// Will be in repository
func (s *GameService) getCourt(ctx context.Context, id int32) (*courtpb.Court, error) {
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

	err := s.db.QueryRowContext(ctx, query, id).Scan(&court.Id, &court.Name, &district.Id, &district.Name, &district.CityId, &court.Address, &court.CommentCount, &court.CheckInCount)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "court not found")
	}

	court.District = district
	court.Media = make([]*commonpb.Media, 0)

	return court, nil
}

// TODO: optimize this mess
// it's not scalable but it's ok for mvp
func (s *GameService) getGameById(ctx context.Context, id int32) (*gamepb.Game, error) {
	query := `SELECT
		g.id,
		g.startedAt,
		g.endedAt,
		g.homeScore,
		g.awayScore,
		g.courtID
	FROM game g
	WHERE g.id = ?;`

	var game gamepb.Game
	var startedAt sql.NullTime
	var endedAt sql.NullTime
	var courtId int32
	var homeScore int32
	var awayScore int32

	err := s.db.QueryRow(query, id).Scan(&game.Id, &startedAt, &endedAt, &homeScore, &awayScore, &courtId)
	if err != nil {
		fmt.Println(err)
		return nil, gerr(codes.Internal, err)
	}

	// Court
	court, err := s.getCourt(ctx, courtId)
	if err != nil {
		fmt.Println(err)
		return nil, gerr(codes.Internal, err)
	}

	game.Court = court
	//

	// Status
	if startedAt.Valid {
		if endedAt.Valid {
			// Game Completed
			endedGame := gamepb.GameEnded{
				StartedAt: timestamppb.New(startedAt.Time),
				EndedAt:   timestamppb.New(endedAt.Time),
				HomeScore: homeScore,
				AwayScore: awayScore,
			}

			game.Status = &gamepb.Game_Ended{
				Ended: &endedGame,
			}
		} else {
			// Game Started
			startedGame := gamepb.GameStarted{
				StartedAt: timestamppb.New(startedAt.Time),
			}

			game.Status = &gamepb.Game_Started{
				Started: &startedGame,
			}
		}
	} else {
		// Game Created
		game.Status = &gamepb.Game_InProgress{
			// I'm not sure if it's necessary
			InProgress: &gamepb.GameInProgress{},
		}
	}
	//

	// TODO: Teams will be added in later phases
	// Teams
	// Check if game is team game
	//

	// Players
	query = `SELECT
		u.id,
		u.email,
		u.districtID,
		u.name,
		u.lastName,
		u.username,
		f.path,
		gp.isHomeSide,
		gp.isCanceled
	FROM game_player gp
	INNER JOIN user u ON gp.playerID = u.id
	LEFT JOIN file f ON u.avatarID = f.id
	WHERE gp.gameID = ?;`

	rows, err := s.db.QueryContext(ctx, query, id)
	if err != nil {
		fmt.Println(err)
		return nil, gerr(codes.Internal, err)
	}

	defer rows.Close()
	var homePlayers []*userpb.User
	var awayPlayers []*userpb.User
	var canceledUsers []*userpb.User

	for rows.Next() {
		var user userpb.User
		var avatarPath sql.NullString
		var isHomeSide bool
		var isCanceled bool

		err := rows.Scan(&user.Id, &user.Email, &user.DistrictID, &user.Name, &user.Lastname, &user.Username, &avatarPath, &isHomeSide, &isCanceled)
		if err != nil {
			fmt.Println(err)
			return nil, gerr(codes.Internal, err)
		}
		user.AvatarPath = avatarPath.String
		if isHomeSide {
			homePlayers = append(homePlayers, &user)
		} else {
			awayPlayers = append(awayPlayers, &user)
		}

		if isCanceled {
			canceledUsers = append(canceledUsers, &user)
		}
	}

	game.HomePlayers = homePlayers
	game.AwayPlayers = awayPlayers
	game.CanceledBy = canceledUsers

	return &game, nil
}
