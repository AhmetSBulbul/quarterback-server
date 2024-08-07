// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.3
// source: pointspb.proto

package pointspb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type QueryByDistrict struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DistrictID int32 `protobuf:"varint,1,opt,name=districtID,proto3" json:"districtID,omitempty"`
}

func (x *QueryByDistrict) Reset() {
	*x = QueryByDistrict{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pointspb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryByDistrict) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryByDistrict) ProtoMessage() {}

func (x *QueryByDistrict) ProtoReflect() protoreflect.Message {
	mi := &file_pointspb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryByDistrict.ProtoReflect.Descriptor instead.
func (*QueryByDistrict) Descriptor() ([]byte, []int) {
	return file_pointspb_proto_rawDescGZIP(), []int{0}
}

func (x *QueryByDistrict) GetDistrictID() int32 {
	if x != nil {
		return x.DistrictID
	}
	return 0
}

type QueryByUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerID int32 `protobuf:"varint,1,opt,name=playerID,proto3" json:"playerID,omitempty"`
}

func (x *QueryByUser) Reset() {
	*x = QueryByUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pointspb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryByUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryByUser) ProtoMessage() {}

func (x *QueryByUser) ProtoReflect() protoreflect.Message {
	mi := &file_pointspb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryByUser.ProtoReflect.Descriptor instead.
func (*QueryByUser) Descriptor() ([]byte, []int) {
	return file_pointspb_proto_rawDescGZIP(), []int{1}
}

func (x *QueryByUser) GetPlayerID() int32 {
	if x != nil {
		return x.PlayerID
	}
	return 0
}

type PointsByDistrict struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserPoints []*UserPoints `protobuf:"bytes,2,rep,name=user_points,json=userPoints,proto3" json:"user_points,omitempty"`
}

func (x *PointsByDistrict) Reset() {
	*x = PointsByDistrict{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pointspb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PointsByDistrict) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PointsByDistrict) ProtoMessage() {}

func (x *PointsByDistrict) ProtoReflect() protoreflect.Message {
	mi := &file_pointspb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PointsByDistrict.ProtoReflect.Descriptor instead.
func (*PointsByDistrict) Descriptor() ([]byte, []int) {
	return file_pointspb_proto_rawDescGZIP(), []int{2}
}

func (x *PointsByDistrict) GetUserPoints() []*UserPoints {
	if x != nil {
		return x.UserPoints
	}
	return nil
}

type UserPoints struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerID    int32  `protobuf:"varint,1,opt,name=playerID,proto3" json:"playerID,omitempty"`
	Username    string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Wins        int32  `protobuf:"varint,3,opt,name=wins,proto3" json:"wins,omitempty"`
	Losses      int32  `protobuf:"varint,4,opt,name=losses,proto3" json:"losses,omitempty"`
	Cancelled   int32  `protobuf:"varint,5,opt,name=cancelled,proto3" json:"cancelled,omitempty"`
	TotalPoints int32  `protobuf:"varint,6,opt,name=totalPoints,proto3" json:"totalPoints,omitempty"`
}

func (x *UserPoints) Reset() {
	*x = UserPoints{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pointspb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPoints) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPoints) ProtoMessage() {}

func (x *UserPoints) ProtoReflect() protoreflect.Message {
	mi := &file_pointspb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPoints.ProtoReflect.Descriptor instead.
func (*UserPoints) Descriptor() ([]byte, []int) {
	return file_pointspb_proto_rawDescGZIP(), []int{3}
}

func (x *UserPoints) GetPlayerID() int32 {
	if x != nil {
		return x.PlayerID
	}
	return 0
}

func (x *UserPoints) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserPoints) GetWins() int32 {
	if x != nil {
		return x.Wins
	}
	return 0
}

func (x *UserPoints) GetLosses() int32 {
	if x != nil {
		return x.Losses
	}
	return 0
}

func (x *UserPoints) GetCancelled() int32 {
	if x != nil {
		return x.Cancelled
	}
	return 0
}

func (x *UserPoints) GetTotalPoints() int32 {
	if x != nil {
		return x.TotalPoints
	}
	return 0
}

type UserStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalGames  int32 `protobuf:"varint,1,opt,name=totalGames,proto3" json:"totalGames,omitempty"`
	Wins        int32 `protobuf:"varint,2,opt,name=wins,proto3" json:"wins,omitempty"`
	Losses      int32 `protobuf:"varint,3,opt,name=losses,proto3" json:"losses,omitempty"`
	Cancelled   int32 `protobuf:"varint,4,opt,name=cancelled,proto3" json:"cancelled,omitempty"`
	TotalPoints int32 `protobuf:"varint,5,opt,name=totalPoints,proto3" json:"totalPoints,omitempty"`
	Rank        int32 `protobuf:"varint,6,opt,name=rank,proto3" json:"rank,omitempty"`
}

func (x *UserStats) Reset() {
	*x = UserStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pointspb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserStats) ProtoMessage() {}

func (x *UserStats) ProtoReflect() protoreflect.Message {
	mi := &file_pointspb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserStats.ProtoReflect.Descriptor instead.
func (*UserStats) Descriptor() ([]byte, []int) {
	return file_pointspb_proto_rawDescGZIP(), []int{4}
}

func (x *UserStats) GetTotalGames() int32 {
	if x != nil {
		return x.TotalGames
	}
	return 0
}

func (x *UserStats) GetWins() int32 {
	if x != nil {
		return x.Wins
	}
	return 0
}

func (x *UserStats) GetLosses() int32 {
	if x != nil {
		return x.Losses
	}
	return 0
}

func (x *UserStats) GetCancelled() int32 {
	if x != nil {
		return x.Cancelled
	}
	return 0
}

func (x *UserStats) GetTotalPoints() int32 {
	if x != nil {
		return x.TotalPoints
	}
	return 0
}

func (x *UserStats) GetRank() int32 {
	if x != nil {
		return x.Rank
	}
	return 0
}

var File_pointspb_proto protoreflect.FileDescriptor

var file_pointspb_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22, 0x31, 0x0a, 0x0f, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x42, 0x79, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x64,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x49, 0x44, 0x22, 0x29, 0x0a, 0x0b, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x22, 0x47, 0x0a, 0x10, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73,
	0x42, 0x79, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x12, 0x33, 0x0a, 0x0b, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x73, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22,
	0xb0, 0x01, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x69, 0x6e, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x77, 0x69, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f,
	0x73, 0x73, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6c, 0x6f, 0x73, 0x73,
	0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64,
	0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x22, 0xab, 0x01, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x47, 0x61, 0x6d, 0x65, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x77, 0x69, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x77, 0x69, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x73, 0x73, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6c, 0x6f, 0x73, 0x73, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x72, 0x61, 0x6e, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b,
	0x32, 0x91, 0x01, 0x0a, 0x0d, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x48, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x42,
	0x79, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x12, 0x17, 0x2e, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x79, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x63, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x42, 0x79, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x12, 0x36, 0x0a, 0x0c,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x13, 0x2e, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x79, 0x55, 0x73, 0x65,
	0x72, 0x1a, 0x11, 0x2e, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pointspb_proto_rawDescOnce sync.Once
	file_pointspb_proto_rawDescData = file_pointspb_proto_rawDesc
)

func file_pointspb_proto_rawDescGZIP() []byte {
	file_pointspb_proto_rawDescOnce.Do(func() {
		file_pointspb_proto_rawDescData = protoimpl.X.CompressGZIP(file_pointspb_proto_rawDescData)
	})
	return file_pointspb_proto_rawDescData
}

var file_pointspb_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pointspb_proto_goTypes = []interface{}{
	(*QueryByDistrict)(nil),  // 0: points.QueryByDistrict
	(*QueryByUser)(nil),      // 1: points.QueryByUser
	(*PointsByDistrict)(nil), // 2: points.PointsByDistrict
	(*UserPoints)(nil),       // 3: points.UserPoints
	(*UserStats)(nil),        // 4: points.UserStats
}
var file_pointspb_proto_depIdxs = []int32{
	3, // 0: points.PointsByDistrict.user_points:type_name -> points.UserPoints
	0, // 1: points.PointsService.GetPointsByDistrict:input_type -> points.QueryByDistrict
	1, // 2: points.PointsService.GetUserStats:input_type -> points.QueryByUser
	2, // 3: points.PointsService.GetPointsByDistrict:output_type -> points.PointsByDistrict
	4, // 4: points.PointsService.GetUserStats:output_type -> points.UserStats
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pointspb_proto_init() }
func file_pointspb_proto_init() {
	if File_pointspb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pointspb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryByDistrict); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pointspb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryByUser); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pointspb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PointsByDistrict); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pointspb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserPoints); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pointspb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserStats); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pointspb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pointspb_proto_goTypes,
		DependencyIndexes: file_pointspb_proto_depIdxs,
		MessageInfos:      file_pointspb_proto_msgTypes,
	}.Build()
	File_pointspb_proto = out.File
	file_pointspb_proto_rawDesc = nil
	file_pointspb_proto_goTypes = nil
	file_pointspb_proto_depIdxs = nil
}
