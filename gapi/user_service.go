package gapi

import (
	"context"
	"database/sql"
	"math"

	"github.com/AhmetSBulbul/quarterback-server/pb/commonpb"
	"github.com/AhmetSBulbul/quarterback-server/pb/userpb"
	"google.golang.org/grpc/codes"
)

type UserService struct {
	db *sql.DB
	userpb.UnimplementedUserServiceServer
}

func getUserIdFromCtx(ctx context.Context) int {
	return ctx.Value(Sub_id("sub_id")).(int)
}

func (s *UserService) getUserByID(ctx context.Context, userid int) (*userpb.User, error) {
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

func (s *UserService) getUsersBySearch(ctx context.Context, value string, cursor int, limit int) ([]*userpb.User, int, error) {
	// TODO: Filter blocked users?
	sub_id := getUserIdFromCtx(ctx)
	query := `SELECT 
		u.id,
		u.email,
		u.districtID,
		u.name,
		u.lastName,
		u.username,
		f.path,
		COUNT(*) OVER() AS total
	FROM user u
	LEFT JOIN file f ON f.id = u.avatarID
	WHERE 
		u.id > ? AND
		u.id != ? AND
		(u.username LIKE ? OR
		u.name LIKE ? OR
		u.lastName LIKE ?)
	ORDER BY u.id ASC
	LIMIT ?`

	_like := "%%"

	if value != "" {
		_like = "%" + value + "%"
	}

	rows, err := s.db.QueryContext(
		ctx,
		query,
		cursor,
		sub_id,
		_like,
		_like,
		_like,
		limit,
	)
	if err != nil {
		return nil, 0, gerr(codes.Internal, err)
	}
	defer rows.Close()
	var total int
	var users []*userpb.User

	for rows.Next() {
		var user userpb.User
		var avatarID sql.NullString
		err := rows.Scan(&user.Id, &user.Email, &user.DistrictID, &user.Name, &user.Lastname, &user.Username, &avatarID, &total)
		if err != nil {
			return nil, 0, gerr(codes.Internal, err)
		}
		user.AvatarPath = avatarID.String
		users = append(users, &user)
	}

	return users, total, nil
}

func (s *UserService) GetMe(ctx context.Context, in *commonpb.Empty) (*userpb.UserResponse, error) {
	var sub_id = getUserIdFromCtx(ctx)
	user, err := s.getUserByID(ctx, sub_id)

	if err != nil {
		return nil, err
	}

	return &userpb.UserResponse{
		User: user,
	}, nil
}

func (s *UserService) GetUser(ctx context.Context, in *commonpb.GetByIdRequest) (*userpb.UserResponse, error) {
	user, err := s.getUserByID(ctx, int(in.Id))
	if err != nil {
		return nil, err
	}
	return &userpb.UserResponse{
		User: user,
	}, nil
}

func (s *UserService) SearchUsers(ctx context.Context, in *userpb.UserSearchRequest) (*userpb.UserListResponse, error) {
	cursor := in.Pagination.GetCursorId()
	limit := int(math.Min(float64(in.Pagination.GetLimit()), 200))

	if limit == 0 {
		limit = 10
	}

	users, total, err := s.getUsersBySearch(ctx, in.Query.GetQuery(), int(cursor), limit)
	if err != nil {
		return nil, err
	}
	var cursorId int32 = 0
	if len(users) > 0 {
		cursorId = users[len(users)-1].Id
	}
	return &userpb.UserListResponse{
		User: users,
		Pagination: &commonpb.PaginationResponse{
			CursorId: cursorId,
			Total:    int32(total),
		},
	}, nil
	// TODO: implement query by district
	// var users []*userpb.User
	// return &userpb.UserListResponse{
	// 	User: users,
	// 	Pagination: &commonpb.PaginationResponse{
	// 		CursorId: 0,
	// 		Total:    0,
	// 	},
	// }, nil
}

func (s *UserService) UpdateUser(ctx context.Context, in *userpb.UserUpdateRequest) (*userpb.UserResponse, error) {
	sub_id := getUserIdFromCtx(ctx)

	var user userpb.User
	row := s.db.QueryRowContext(ctx, "SELECT id, name, lastName, districtID FROM user WHERE id = ?", sub_id)
	err := row.Scan(&user.Id, &user.Name, &user.Lastname, &user.DistrictID)

	if err != nil {
		return nil, gerr(codes.Internal, err)
	}

	if in.GetName() != "" {
		user.Name = in.GetName()
	}

	if in.GetLastname() != "" {
		user.Lastname = in.GetLastname()
	}

	if in.GetDistrictID() != 0 {
		user.DistrictID = in.GetDistrictID()
	}

	query := "UPDATE user SET name = ?, lastName = ?, districtID = ? WHERE id = ?"
	_, err = s.db.ExecContext(ctx, query, user.Name, user.Lastname, user.DistrictID, user.Id)

	if err != nil {
		return nil, gerr(codes.Internal, err)
	}

	updatedUser, err := s.getUserByID(ctx, sub_id)

	if err != nil {
		return nil, gerr(codes.Internal, err)
	}

	return &userpb.UserResponse{
		User: updatedUser,
	}, nil
}

func (s *UserService) UploadAvatar(ctx context.Context, in *userpb.UpdateAvatarRequest) (*userpb.UserResponse, error) {
	sub_id := getUserIdFromCtx(ctx)
	fileId := in.GetAvatarFileId()

	query := "UPDATE user SET avatarID = ? WHERE id = ?"
	_, err := s.db.ExecContext(ctx, query, fileId, sub_id)

	if err != nil {
		return nil, gerr(codes.Internal, err)
	}

	updatedUser, err := s.getUserByID(ctx, sub_id)

	if err != nil {
		return nil, gerr(codes.Internal, err)
	}

	return &userpb.UserResponse{
		User: updatedUser,
	}, nil
}

func (s *UserService) ToggleFollow(ctx context.Context, in *commonpb.GetByIdRequest) (*userpb.FollowResponse, error) {
	sub_id := getUserIdFromCtx(ctx)

	rows := s.db.QueryRowContext(ctx, "SELECT followingID FROM follower WHERE followerID = ? AND followingID = ?", sub_id, in.Id)

	var followingID int
	err := rows.Scan(&followingID)

	if err != nil && err != sql.ErrNoRows {
		return nil, gerr(codes.Internal, err)
	}

	if err == sql.ErrNoRows {
		_, err := s.db.ExecContext(ctx, "INSERT INTO follower (followerID, followingID) VALUES (?, ?)", sub_id, in.Id)
		if err != nil {
			return nil, gerr(codes.Internal, err)
		}
		return &userpb.FollowResponse{
			IsFollowing:   true,
			SubjectUserID: in.Id,
		}, nil
	}

	_, err = s.db.ExecContext(ctx, "DELETE FROM follower WHERE followerID = ? AND followingID = ?", sub_id, in.Id)
	if err != nil {
		return nil, gerr(codes.Internal, err)
	}

	return &userpb.FollowResponse{
		IsFollowing:   false,
		SubjectUserID: in.Id,
	}, nil
}

func (s *UserService) ToggleBlock(ctx context.Context, in *commonpb.GetByIdRequest) (*userpb.BlockResponse, error) {
	sub_id := getUserIdFromCtx(ctx)

	rows := s.db.QueryRowContext(ctx, "SELECT blockedID FROM block WHERE blockerID = ? AND blockedID = ?", sub_id, in.Id)

	var blockedID int
	err := rows.Scan(&blockedID)

	if err != nil && err != sql.ErrNoRows {
		return nil, gerr(codes.Internal, err)
	}

	if err == sql.ErrNoRows {
		_, err := s.db.ExecContext(ctx, "INSERT INTO block (blockerID, blockedID) VALUES (?, ?)", sub_id, in.Id)
		if err != nil {
			return nil, gerr(codes.Internal, err)
		}
		return &userpb.BlockResponse{
			IsBlocked:     true,
			SubjectUserID: in.Id,
		}, nil
	}

	_, err = s.db.ExecContext(ctx, "DELETE FROM block WHERE blockerID = ? AND blockedID = ?", sub_id, in.Id)
	if err != nil {
		return nil, gerr(codes.Internal, err)
	}

	return &userpb.BlockResponse{
		IsBlocked:     false,
		SubjectUserID: in.Id,
	}, nil
}

func (s *UserService) AddComment(ctx context.Context, in *commonpb.CommentRequest) (*commonpb.CommentResponse, error) {
	sub_id := getUserIdFromCtx(ctx)

	result, err := s.db.ExecContext(ctx, "INSERT INTO comment (senderID, receiverID, content) VALUES (?, ?, ?)", sub_id, in.TargetId, in.Content)

	if err != nil {
		return nil, gerr(codes.Internal, err)
	}

	commentId, err := result.LastInsertId()
	if err != nil {
		return nil, gerr(codes.Internal, err)
	}

	comment := commonpb.Comment{}
	row := s.db.QueryRowContext(ctx, "SELECT id, senderID, receiverID, content, isHidden FROM comment WHERE id = ?", commentId)

	err = row.Scan(&comment.Id, &comment.CommenterId, &comment.TargetId, &comment.Content, &comment.IsHidden)

	if err != nil {
		return nil, gerr(codes.Internal, err)
	}

	return &commonpb.CommentResponse{
		Comment: &comment,
	}, nil
}

func (s *UserService) GetFollowers(ctx context.Context, in *commonpb.GetByIdPaginatedRequest) (*userpb.UserListResponse, error) {
	// isFollower and isFollowing will be checked on client side. So due of that we don't use pagination here.
	// But we will!
	query := `SELECT
		u.id,
		u.email,
		u.districtId,
		u.name,
		u.lastName,
		u.username,
		f.path
	FROM follower fol
	INNER JOIN user u ON u.id = fol.followerID
	LEFT JOIN file f ON f.id = u.avatarID
	WHERE fol.followingID = ?`

	rows, err := s.db.QueryContext(ctx, query, in.Id)
	if err != nil {
		return nil, gerr(codes.Internal, err)
	}

	defer rows.Close()

	var users []*userpb.User

	for rows.Next() {
		var user userpb.User
		var avatarPath sql.NullString
		err := rows.Scan(&user.Id, &user.Email, &user.DistrictID, &user.Name, &user.Lastname, &user.Username, &avatarPath)
		if err != nil {
			return nil, gerr(codes.Internal, err)
		}
		user.AvatarPath = avatarPath.String
		users = append(users, &user)
	}

	return &userpb.UserListResponse{
		User:       users,
		Pagination: &commonpb.PaginationResponse{},
	}, nil

}

func (s *UserService) GetFollowing(ctx context.Context, in *commonpb.GetByIdPaginatedRequest) (*userpb.UserListResponse, error) {
	// Copy - Paste from follower
	query := `SELECT
		u.id,
		u.email,
		u.districtId,
		u.name,
		u.lastName,
		u.username,
		f.path
	FROM follower fol
	INNER JOIN user u ON u.id = fol.followingID
	LEFT JOIN file f ON f.id = u.avatarID
	WHERE fol.followerID = ?`

	rows, err := s.db.QueryContext(ctx, query, in.Id)
	if err != nil {
		return nil, gerr(codes.Internal, err)
	}

	defer rows.Close()

	var users []*userpb.User

	for rows.Next() {
		var user userpb.User
		var avatarPath sql.NullString
		err := rows.Scan(&user.Id, &user.Email, &user.DistrictID, &user.Name, &user.Lastname, &user.Username, &avatarPath)
		if err != nil {
			return nil, gerr(codes.Internal, err)
		}
		user.AvatarPath = avatarPath.String
		users = append(users, &user)
	}

	return &userpb.UserListResponse{
		User:       users,
		Pagination: &commonpb.PaginationResponse{},
	}, nil
}

func (s *UserService) GetBlocked(_ context.Context, _ *commonpb.GetByIdPaginatedRequest) (*userpb.UserListResponse, error) {
	return nil, gerr(codes.Unimplemented, nil) // TODO: Implement
}

func (s *UserService) GetComments(_ context.Context, _ *commonpb.GetByIdPaginatedRequest) (*commonpb.CommentsResponse, error) {
	return nil, gerr(codes.Unimplemented, nil) // TODO: Implement
}
