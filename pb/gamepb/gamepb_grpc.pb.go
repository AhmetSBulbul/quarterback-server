// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: gamepb.proto

package gamepb

import (
	context "context"
	commonpb "github.com/AhmetSBulbul/quarterback-server/pb/commonpb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GameServiceClient is the client API for GameService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GameServiceClient interface {
	GetGame(ctx context.Context, in *GetGameRequest, opts ...grpc.CallOption) (*GameResponse, error)
	ListGamesByLocation(ctx context.Context, in *ListGamesByLocationRequest, opts ...grpc.CallOption) (*ListGamesResponse, error)
	ListGamesByUser(ctx context.Context, in *ListGamesByUserRequest, opts ...grpc.CallOption) (*ListGamesResponse, error)
	ListGamesByCourt(ctx context.Context, in *ListGamesByCourtRequest, opts ...grpc.CallOption) (*ListGamesResponse, error)
	ListGamesByTeam(ctx context.Context, in *ListGamesByTeamRequest, opts ...grpc.CallOption) (*ListGamesResponse, error)
	CreateGame(ctx context.Context, in *CreateGameRequest, opts ...grpc.CallOption) (*GameIdResponse, error)
	CreateGameWithTeam(ctx context.Context, in *CreateGameWithTeamRequest, opts ...grpc.CallOption) (*GameIdResponse, error)
	JoinGame(ctx context.Context, in *JoinGameRequest, opts ...grpc.CallOption) (*GameIdResponse, error)
	StartGame(ctx context.Context, in *StartGameRequest, opts ...grpc.CallOption) (*GameIdResponse, error)
	EndGame(ctx context.Context, in *EndGameRequest, opts ...grpc.CallOption) (*GameIdResponse, error)
	CancelGame(ctx context.Context, in *CancelGameRequest, opts ...grpc.CallOption) (*GameIdResponse, error)
	LeaveGame(ctx context.Context, in *LeaveGameRequest, opts ...grpc.CallOption) (*GameIdResponse, error)
	// Add Comment
	AddMedia(ctx context.Context, in *commonpb.File, opts ...grpc.CallOption) (*commonpb.Media, error)
}

type gameServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGameServiceClient(cc grpc.ClientConnInterface) GameServiceClient {
	return &gameServiceClient{cc}
}

func (c *gameServiceClient) GetGame(ctx context.Context, in *GetGameRequest, opts ...grpc.CallOption) (*GameResponse, error) {
	out := new(GameResponse)
	err := c.cc.Invoke(ctx, "/game.GameService/GetGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) ListGamesByLocation(ctx context.Context, in *ListGamesByLocationRequest, opts ...grpc.CallOption) (*ListGamesResponse, error) {
	out := new(ListGamesResponse)
	err := c.cc.Invoke(ctx, "/game.GameService/ListGamesByLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) ListGamesByUser(ctx context.Context, in *ListGamesByUserRequest, opts ...grpc.CallOption) (*ListGamesResponse, error) {
	out := new(ListGamesResponse)
	err := c.cc.Invoke(ctx, "/game.GameService/ListGamesByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) ListGamesByCourt(ctx context.Context, in *ListGamesByCourtRequest, opts ...grpc.CallOption) (*ListGamesResponse, error) {
	out := new(ListGamesResponse)
	err := c.cc.Invoke(ctx, "/game.GameService/ListGamesByCourt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) ListGamesByTeam(ctx context.Context, in *ListGamesByTeamRequest, opts ...grpc.CallOption) (*ListGamesResponse, error) {
	out := new(ListGamesResponse)
	err := c.cc.Invoke(ctx, "/game.GameService/ListGamesByTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) CreateGame(ctx context.Context, in *CreateGameRequest, opts ...grpc.CallOption) (*GameIdResponse, error) {
	out := new(GameIdResponse)
	err := c.cc.Invoke(ctx, "/game.GameService/CreateGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) CreateGameWithTeam(ctx context.Context, in *CreateGameWithTeamRequest, opts ...grpc.CallOption) (*GameIdResponse, error) {
	out := new(GameIdResponse)
	err := c.cc.Invoke(ctx, "/game.GameService/CreateGameWithTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) JoinGame(ctx context.Context, in *JoinGameRequest, opts ...grpc.CallOption) (*GameIdResponse, error) {
	out := new(GameIdResponse)
	err := c.cc.Invoke(ctx, "/game.GameService/JoinGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) StartGame(ctx context.Context, in *StartGameRequest, opts ...grpc.CallOption) (*GameIdResponse, error) {
	out := new(GameIdResponse)
	err := c.cc.Invoke(ctx, "/game.GameService/StartGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) EndGame(ctx context.Context, in *EndGameRequest, opts ...grpc.CallOption) (*GameIdResponse, error) {
	out := new(GameIdResponse)
	err := c.cc.Invoke(ctx, "/game.GameService/EndGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) CancelGame(ctx context.Context, in *CancelGameRequest, opts ...grpc.CallOption) (*GameIdResponse, error) {
	out := new(GameIdResponse)
	err := c.cc.Invoke(ctx, "/game.GameService/CancelGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) LeaveGame(ctx context.Context, in *LeaveGameRequest, opts ...grpc.CallOption) (*GameIdResponse, error) {
	out := new(GameIdResponse)
	err := c.cc.Invoke(ctx, "/game.GameService/LeaveGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) AddMedia(ctx context.Context, in *commonpb.File, opts ...grpc.CallOption) (*commonpb.Media, error) {
	out := new(commonpb.Media)
	err := c.cc.Invoke(ctx, "/game.GameService/AddMedia", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameServiceServer is the server API for GameService service.
// All implementations must embed UnimplementedGameServiceServer
// for forward compatibility
type GameServiceServer interface {
	GetGame(context.Context, *GetGameRequest) (*GameResponse, error)
	ListGamesByLocation(context.Context, *ListGamesByLocationRequest) (*ListGamesResponse, error)
	ListGamesByUser(context.Context, *ListGamesByUserRequest) (*ListGamesResponse, error)
	ListGamesByCourt(context.Context, *ListGamesByCourtRequest) (*ListGamesResponse, error)
	ListGamesByTeam(context.Context, *ListGamesByTeamRequest) (*ListGamesResponse, error)
	CreateGame(context.Context, *CreateGameRequest) (*GameIdResponse, error)
	CreateGameWithTeam(context.Context, *CreateGameWithTeamRequest) (*GameIdResponse, error)
	JoinGame(context.Context, *JoinGameRequest) (*GameIdResponse, error)
	StartGame(context.Context, *StartGameRequest) (*GameIdResponse, error)
	EndGame(context.Context, *EndGameRequest) (*GameIdResponse, error)
	CancelGame(context.Context, *CancelGameRequest) (*GameIdResponse, error)
	LeaveGame(context.Context, *LeaveGameRequest) (*GameIdResponse, error)
	// Add Comment
	AddMedia(context.Context, *commonpb.File) (*commonpb.Media, error)
	mustEmbedUnimplementedGameServiceServer()
}

// UnimplementedGameServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGameServiceServer struct {
}

func (UnimplementedGameServiceServer) GetGame(context.Context, *GetGameRequest) (*GameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGame not implemented")
}
func (UnimplementedGameServiceServer) ListGamesByLocation(context.Context, *ListGamesByLocationRequest) (*ListGamesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGamesByLocation not implemented")
}
func (UnimplementedGameServiceServer) ListGamesByUser(context.Context, *ListGamesByUserRequest) (*ListGamesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGamesByUser not implemented")
}
func (UnimplementedGameServiceServer) ListGamesByCourt(context.Context, *ListGamesByCourtRequest) (*ListGamesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGamesByCourt not implemented")
}
func (UnimplementedGameServiceServer) ListGamesByTeam(context.Context, *ListGamesByTeamRequest) (*ListGamesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGamesByTeam not implemented")
}
func (UnimplementedGameServiceServer) CreateGame(context.Context, *CreateGameRequest) (*GameIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGame not implemented")
}
func (UnimplementedGameServiceServer) CreateGameWithTeam(context.Context, *CreateGameWithTeamRequest) (*GameIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGameWithTeam not implemented")
}
func (UnimplementedGameServiceServer) JoinGame(context.Context, *JoinGameRequest) (*GameIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinGame not implemented")
}
func (UnimplementedGameServiceServer) StartGame(context.Context, *StartGameRequest) (*GameIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartGame not implemented")
}
func (UnimplementedGameServiceServer) EndGame(context.Context, *EndGameRequest) (*GameIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EndGame not implemented")
}
func (UnimplementedGameServiceServer) CancelGame(context.Context, *CancelGameRequest) (*GameIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelGame not implemented")
}
func (UnimplementedGameServiceServer) LeaveGame(context.Context, *LeaveGameRequest) (*GameIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveGame not implemented")
}
func (UnimplementedGameServiceServer) AddMedia(context.Context, *commonpb.File) (*commonpb.Media, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMedia not implemented")
}
func (UnimplementedGameServiceServer) mustEmbedUnimplementedGameServiceServer() {}

// UnsafeGameServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GameServiceServer will
// result in compilation errors.
type UnsafeGameServiceServer interface {
	mustEmbedUnimplementedGameServiceServer()
}

func RegisterGameServiceServer(s grpc.ServiceRegistrar, srv GameServiceServer) {
	s.RegisterService(&GameService_ServiceDesc, srv)
}

func _GameService_GetGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).GetGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game.GameService/GetGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).GetGame(ctx, req.(*GetGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_ListGamesByLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGamesByLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).ListGamesByLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game.GameService/ListGamesByLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).ListGamesByLocation(ctx, req.(*ListGamesByLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_ListGamesByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGamesByUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).ListGamesByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game.GameService/ListGamesByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).ListGamesByUser(ctx, req.(*ListGamesByUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_ListGamesByCourt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGamesByCourtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).ListGamesByCourt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game.GameService/ListGamesByCourt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).ListGamesByCourt(ctx, req.(*ListGamesByCourtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_ListGamesByTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGamesByTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).ListGamesByTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game.GameService/ListGamesByTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).ListGamesByTeam(ctx, req.(*ListGamesByTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_CreateGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).CreateGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game.GameService/CreateGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).CreateGame(ctx, req.(*CreateGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_CreateGameWithTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGameWithTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).CreateGameWithTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game.GameService/CreateGameWithTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).CreateGameWithTeam(ctx, req.(*CreateGameWithTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_JoinGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).JoinGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game.GameService/JoinGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).JoinGame(ctx, req.(*JoinGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_StartGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).StartGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game.GameService/StartGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).StartGame(ctx, req.(*StartGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_EndGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).EndGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game.GameService/EndGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).EndGame(ctx, req.(*EndGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_CancelGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).CancelGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game.GameService/CancelGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).CancelGame(ctx, req.(*CancelGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_LeaveGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).LeaveGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game.GameService/LeaveGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).LeaveGame(ctx, req.(*LeaveGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_AddMedia_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commonpb.File)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).AddMedia(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/game.GameService/AddMedia",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).AddMedia(ctx, req.(*commonpb.File))
	}
	return interceptor(ctx, in, info, handler)
}

// GameService_ServiceDesc is the grpc.ServiceDesc for GameService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GameService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "game.GameService",
	HandlerType: (*GameServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGame",
			Handler:    _GameService_GetGame_Handler,
		},
		{
			MethodName: "ListGamesByLocation",
			Handler:    _GameService_ListGamesByLocation_Handler,
		},
		{
			MethodName: "ListGamesByUser",
			Handler:    _GameService_ListGamesByUser_Handler,
		},
		{
			MethodName: "ListGamesByCourt",
			Handler:    _GameService_ListGamesByCourt_Handler,
		},
		{
			MethodName: "ListGamesByTeam",
			Handler:    _GameService_ListGamesByTeam_Handler,
		},
		{
			MethodName: "CreateGame",
			Handler:    _GameService_CreateGame_Handler,
		},
		{
			MethodName: "CreateGameWithTeam",
			Handler:    _GameService_CreateGameWithTeam_Handler,
		},
		{
			MethodName: "JoinGame",
			Handler:    _GameService_JoinGame_Handler,
		},
		{
			MethodName: "StartGame",
			Handler:    _GameService_StartGame_Handler,
		},
		{
			MethodName: "EndGame",
			Handler:    _GameService_EndGame_Handler,
		},
		{
			MethodName: "CancelGame",
			Handler:    _GameService_CancelGame_Handler,
		},
		{
			MethodName: "LeaveGame",
			Handler:    _GameService_LeaveGame_Handler,
		},
		{
			MethodName: "AddMedia",
			Handler:    _GameService_AddMedia_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gamepb.proto",
}
