package gapi

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/AhmetSBulbul/quarterback-server/pb/gamepb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GameService struct {
	db *sql.DB
	gamepb.UnimplementedGameServiceServer
}

func (s *GameService) CreateGame(ctx context.Context, in *gamepb.CreateGameRequest) (*gamepb.GameIdResponse, error) {
	userID := getUserIdFromCtx(ctx)
	query := `INSERT INTO game 
	(
		courtID,
		startedAt,
		homeScore,
		awayScore
	) VALUES (
		?,
		NOW(),
		0,
		0
	);`

	res, err := s.db.Exec(query, in.GetCourtId())
	if err != nil {
		fmt.Println(err)
		return nil, status.Errorf(codes.Internal, "Internal error")
	}

	gameID, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err)

		return nil, status.Errorf(codes.Internal, "Internal error")
	}

	query = `INSERT INTO game_player 
	  (gameID, playerID, isHomeSide)
	  VALUES (
		  ?,
		  ?,
		  1
		);`

	_, err = s.db.Exec(query, gameID, userID)

	if err != nil {
		fmt.Println(err)

		return nil, status.Errorf(codes.Internal, "Internal error")
	}

	return &gamepb.GameIdResponse{
		GameId: int32(gameID),
	}, nil
}

func (s *GameService) JoinGame(ctx context.Context, in *gamepb.JoinGameRequest) (*gamepb.GameIdResponse, error) {
	userID := getUserIdFromCtx(ctx)

	query := `SELECT endedAt FROM game WHERE id = ?;`

	var endedAt sql.NullTime
	s.db.QueryRow(query, in.GetGameId()).Scan(&endedAt)

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

	_, err := s.db.Exec(query, in.GetGameId(), userID, in.GetIsHomeSide())

	if err != nil {
		fmt.Println(err)

		return nil, status.Errorf(codes.Internal, "Internal error")
	}

	return &gamepb.GameIdResponse{
		GameId: in.GetGameId(),
	}, nil
}

func (s *GameService) StartGame(ctx context.Context, in *gamepb.StartGameRequest) (*gamepb.GameResponse, error) {
	query := `UPDATE game SET startedAt = NOW() WHERE id = ?;`

	_, err := s.db.Exec(query, in.GetGameId())
	if err != nil {
		fmt.Println(err)
		return nil, status.Errorf(codes.Internal, "Internal error")
	}

	query = `SELECT id, homeScore, awayScore, startedAt, endedAt FROM game WHERE id = ?;`
	row := s.db.QueryRow(query, in.GetGameId())

	var game gamepb.Game{}
	err = row.Scan(&game.Id, &game.HomeScore, &game.AwayScore, &game.StartedAt, &game.EndedAt)

	if err != nil {
		return nil, gerr(codes.Internal, err)
	}

	return nil, status.Errorf(codes.Unimplemented, "method StartGame not implemented")
}

func (s *GameService) EndGame(ctx context.Context, in *gamepb.EndGameRequest) (*gamepb.GameResponse, error) {

	query := `UPDATE game SET endedAt = NOW(), homeScore = ?, awayScore = ? WHERE id = ?;`

	_, err := s.db.Exec(query, in.GetHomeScore(), in.GetAwayScore(), in.GetGameId())

	if err != nil {
		fmt.Println(err)
		return nil, status.Errorf(codes.Internal, "Internal error")
	}

	var game gamepb.Game
	query = `SELECT id, homeScore, awayScore, startedAt, endedAt FROM game WHERE id = ?;`
	s.db.QueryRow(query, in.GetGameId()).Scan(&game.Id, &game.HomeScore, &game.AwayScore, &game.StartedAt, &game.EndedAt)

	return &gamepb.GameResponse{
		Game: &game,
	}, nil
}
