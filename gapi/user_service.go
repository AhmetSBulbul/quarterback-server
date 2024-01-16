package gapi

import (
	"context"
	"database/sql"

	"github.com/AhmetSBulbul/quarterback-server/pb/commonpb"
	"github.com/AhmetSBulbul/quarterback-server/pb/userpb"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	db *sql.DB
	userpb.UnimplementedUserServiceServer
	validate *validator.Validate
}

func (s *UserService) getUserByID(ctx context.Context, userid int) (*userpb.User, error) {
	var user userpb.User
	var avatarID sql.NullString

	query := "SELECT id, email, districtID, name, lastName, username, avatarID FROM user WHERE id = ?"
	row := s.db.QueryRowContext(ctx, query, userid)
	err := row.Scan(&user.Id, &user.Email, &user.DistrictID, &user.Name, &user.Lastname, &user.Username, &avatarID)

	user.AvatarPath = avatarID.String

	if err != nil {
		return nil, gerr(codes.Internal, err)
	}

	return &user, nil
}

func (s *UserService) getUserBySearch(ctx context.Context, query string) ([]*userpb.User, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT id, email, districtID, name, lastName, username, avatarID FROM user WHERE username LIKE %?% OR name LIKE %?% OR lastName LIKE %?%", query, query, query)
	if err != nil {
		return nil, gerr(codes.Internal, err)
	}
	defer rows.Close()

	var users []*userpb.User
	for rows.Next() {
		var user userpb.User
		var avatarID sql.NullString
		err := rows.Scan(&user.Id, &user.Email, &user.DistrictID, &user.Name, &user.Lastname, &user.Username, &avatarID)
		if err != nil {
			return nil, gerr(codes.Internal, err)
		}
		user.AvatarPath = avatarID.String
		users = append(users, &user)
	}

	return users, nil
}

func (s *UserService) GetMe(ctx context.Context, in *commonpb.Empty) (*userpb.UserResponse, error) {
	var sub_id = ctx.Value(Sub_id("sub_id")).(int)
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
	if in.Query.String() != "" {
		users, err := s.getUserBySearch(ctx, in.Query.String())
		if err != nil {
			return nil, err
		}
		return &userpb.UserListResponse{
			User: users,
		}, nil
	}
	// TODO: implement query by district
	var users []*userpb.User
	return &userpb.UserListResponse{
		User: users,
	}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, in *userpb.UserUpdateRequest) (*userpb.UserResponse, error) {
	sub_id := ctx.Value(Sub_id("sub_id")).(int)

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

func (s *UserService) UploadAvatar(ctx context.Context, in *commonpb.File) (*userpb.UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadAvatar not implemented")
}

func (s *UserService) ToggleFollow(ctx context.Context, in *commonpb.GetByIdRequest) (*userpb.FollowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ToggleFollow not implemented")
}
func (s *UserService) ToggleBlock(ctx context.Context, in *commonpb.GetByIdRequest) (*userpb.FollowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ToggleBlock not implemented")
}
func (s *UserService) AddComment(ctx context.Context, in *commonpb.CommentRequest) (*commonpb.CommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddComment not implemented")
}
