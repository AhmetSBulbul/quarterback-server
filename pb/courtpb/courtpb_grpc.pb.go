// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: courtpb.proto

package courtpb

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

// CourtServiceClient is the client API for CourtService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CourtServiceClient interface {
	GetCourt(ctx context.Context, in *GetCourtRequest, opts ...grpc.CallOption) (*CourtResponse, error)
	ListCourtByLocation(ctx context.Context, in *ListCourtByLocationRequest, opts ...grpc.CallOption) (*ListCourtResponse, error)
	SearchCourt(ctx context.Context, in *SearchCourtRequest, opts ...grpc.CallOption) (*ListCourtResponse, error)
	CreateCourt(ctx context.Context, in *Court, opts ...grpc.CallOption) (*CourtResponse, error)
	CheckInCourt(ctx context.Context, in *commonpb.GetByIdRequest, opts ...grpc.CallOption) (*CheckInCourtResponse, error)
	AddComment(ctx context.Context, in *CourtCommentRequest, opts ...grpc.CallOption) (*CourtComment, error)
	ListComment(ctx context.Context, in *CourtCommentListRequest, opts ...grpc.CallOption) (*CourtCommentListResponse, error)
}

type courtServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCourtServiceClient(cc grpc.ClientConnInterface) CourtServiceClient {
	return &courtServiceClient{cc}
}

func (c *courtServiceClient) GetCourt(ctx context.Context, in *GetCourtRequest, opts ...grpc.CallOption) (*CourtResponse, error) {
	out := new(CourtResponse)
	err := c.cc.Invoke(ctx, "/court.CourtService/GetCourt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courtServiceClient) ListCourtByLocation(ctx context.Context, in *ListCourtByLocationRequest, opts ...grpc.CallOption) (*ListCourtResponse, error) {
	out := new(ListCourtResponse)
	err := c.cc.Invoke(ctx, "/court.CourtService/ListCourtByLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courtServiceClient) SearchCourt(ctx context.Context, in *SearchCourtRequest, opts ...grpc.CallOption) (*ListCourtResponse, error) {
	out := new(ListCourtResponse)
	err := c.cc.Invoke(ctx, "/court.CourtService/SearchCourt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courtServiceClient) CreateCourt(ctx context.Context, in *Court, opts ...grpc.CallOption) (*CourtResponse, error) {
	out := new(CourtResponse)
	err := c.cc.Invoke(ctx, "/court.CourtService/CreateCourt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courtServiceClient) CheckInCourt(ctx context.Context, in *commonpb.GetByIdRequest, opts ...grpc.CallOption) (*CheckInCourtResponse, error) {
	out := new(CheckInCourtResponse)
	err := c.cc.Invoke(ctx, "/court.CourtService/CheckInCourt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courtServiceClient) AddComment(ctx context.Context, in *CourtCommentRequest, opts ...grpc.CallOption) (*CourtComment, error) {
	out := new(CourtComment)
	err := c.cc.Invoke(ctx, "/court.CourtService/AddComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courtServiceClient) ListComment(ctx context.Context, in *CourtCommentListRequest, opts ...grpc.CallOption) (*CourtCommentListResponse, error) {
	out := new(CourtCommentListResponse)
	err := c.cc.Invoke(ctx, "/court.CourtService/ListComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CourtServiceServer is the server API for CourtService service.
// All implementations must embed UnimplementedCourtServiceServer
// for forward compatibility
type CourtServiceServer interface {
	GetCourt(context.Context, *GetCourtRequest) (*CourtResponse, error)
	ListCourtByLocation(context.Context, *ListCourtByLocationRequest) (*ListCourtResponse, error)
	SearchCourt(context.Context, *SearchCourtRequest) (*ListCourtResponse, error)
	CreateCourt(context.Context, *Court) (*CourtResponse, error)
	CheckInCourt(context.Context, *commonpb.GetByIdRequest) (*CheckInCourtResponse, error)
	AddComment(context.Context, *CourtCommentRequest) (*CourtComment, error)
	ListComment(context.Context, *CourtCommentListRequest) (*CourtCommentListResponse, error)
	mustEmbedUnimplementedCourtServiceServer()
}

// UnimplementedCourtServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCourtServiceServer struct {
}

func (UnimplementedCourtServiceServer) GetCourt(context.Context, *GetCourtRequest) (*CourtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourt not implemented")
}
func (UnimplementedCourtServiceServer) ListCourtByLocation(context.Context, *ListCourtByLocationRequest) (*ListCourtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCourtByLocation not implemented")
}
func (UnimplementedCourtServiceServer) SearchCourt(context.Context, *SearchCourtRequest) (*ListCourtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchCourt not implemented")
}
func (UnimplementedCourtServiceServer) CreateCourt(context.Context, *Court) (*CourtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCourt not implemented")
}
func (UnimplementedCourtServiceServer) CheckInCourt(context.Context, *commonpb.GetByIdRequest) (*CheckInCourtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckInCourt not implemented")
}
func (UnimplementedCourtServiceServer) AddComment(context.Context, *CourtCommentRequest) (*CourtComment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddComment not implemented")
}
func (UnimplementedCourtServiceServer) ListComment(context.Context, *CourtCommentListRequest) (*CourtCommentListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListComment not implemented")
}
func (UnimplementedCourtServiceServer) mustEmbedUnimplementedCourtServiceServer() {}

// UnsafeCourtServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CourtServiceServer will
// result in compilation errors.
type UnsafeCourtServiceServer interface {
	mustEmbedUnimplementedCourtServiceServer()
}

func RegisterCourtServiceServer(s grpc.ServiceRegistrar, srv CourtServiceServer) {
	s.RegisterService(&CourtService_ServiceDesc, srv)
}

func _CourtService_GetCourt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCourtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourtServiceServer).GetCourt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/court.CourtService/GetCourt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourtServiceServer).GetCourt(ctx, req.(*GetCourtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourtService_ListCourtByLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCourtByLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourtServiceServer).ListCourtByLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/court.CourtService/ListCourtByLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourtServiceServer).ListCourtByLocation(ctx, req.(*ListCourtByLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourtService_SearchCourt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchCourtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourtServiceServer).SearchCourt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/court.CourtService/SearchCourt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourtServiceServer).SearchCourt(ctx, req.(*SearchCourtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourtService_CreateCourt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Court)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourtServiceServer).CreateCourt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/court.CourtService/CreateCourt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourtServiceServer).CreateCourt(ctx, req.(*Court))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourtService_CheckInCourt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commonpb.GetByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourtServiceServer).CheckInCourt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/court.CourtService/CheckInCourt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourtServiceServer).CheckInCourt(ctx, req.(*commonpb.GetByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourtService_AddComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CourtCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourtServiceServer).AddComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/court.CourtService/AddComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourtServiceServer).AddComment(ctx, req.(*CourtCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourtService_ListComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CourtCommentListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourtServiceServer).ListComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/court.CourtService/ListComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourtServiceServer).ListComment(ctx, req.(*CourtCommentListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CourtService_ServiceDesc is the grpc.ServiceDesc for CourtService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CourtService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "court.CourtService",
	HandlerType: (*CourtServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCourt",
			Handler:    _CourtService_GetCourt_Handler,
		},
		{
			MethodName: "ListCourtByLocation",
			Handler:    _CourtService_ListCourtByLocation_Handler,
		},
		{
			MethodName: "SearchCourt",
			Handler:    _CourtService_SearchCourt_Handler,
		},
		{
			MethodName: "CreateCourt",
			Handler:    _CourtService_CreateCourt_Handler,
		},
		{
			MethodName: "CheckInCourt",
			Handler:    _CourtService_CheckInCourt_Handler,
		},
		{
			MethodName: "AddComment",
			Handler:    _CourtService_AddComment_Handler,
		},
		{
			MethodName: "ListComment",
			Handler:    _CourtService_ListComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "courtpb.proto",
}
