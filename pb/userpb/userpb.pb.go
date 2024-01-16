// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: userpb.proto

package userpb

import (
	commonpb "github.com/AhmetSBulbul/quarterback-server/pb/commonpb"
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

type UserSearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query      *commonpb.Query             `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Pagination *commonpb.PaginationRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *UserSearchRequest) Reset() {
	*x = UserSearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userpb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSearchRequest) ProtoMessage() {}

func (x *UserSearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userpb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSearchRequest.ProtoReflect.Descriptor instead.
func (*UserSearchRequest) Descriptor() ([]byte, []int) {
	return file_userpb_proto_rawDescGZIP(), []int{0}
}

func (x *UserSearchRequest) GetQuery() *commonpb.Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *UserSearchRequest) GetPagination() *commonpb.PaginationRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

// Be able to update only own user
type UserUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       *string `protobuf:"bytes,4,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Lastname   *string `protobuf:"bytes,5,opt,name=lastname,proto3,oneof" json:"lastname,omitempty"`
	DistrictID *int32  `protobuf:"varint,6,opt,name=districtID,proto3,oneof" json:"districtID,omitempty"`
}

func (x *UserUpdateRequest) Reset() {
	*x = UserUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userpb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserUpdateRequest) ProtoMessage() {}

func (x *UserUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userpb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserUpdateRequest.ProtoReflect.Descriptor instead.
func (*UserUpdateRequest) Descriptor() ([]byte, []int) {
	return file_userpb_proto_rawDescGZIP(), []int{1}
}

func (x *UserUpdateRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *UserUpdateRequest) GetLastname() string {
	if x != nil && x.Lastname != nil {
		return *x.Lastname
	}
	return ""
}

func (x *UserUpdateRequest) GetDistrictID() int32 {
	if x != nil && x.DistrictID != nil {
		return *x.DistrictID
	}
	return 0
}

type UpdateAvatarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvatarFileId int32 `protobuf:"varint,1,opt,name=avatarFileId,proto3" json:"avatarFileId,omitempty"`
}

func (x *UpdateAvatarRequest) Reset() {
	*x = UpdateAvatarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userpb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAvatarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAvatarRequest) ProtoMessage() {}

func (x *UpdateAvatarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userpb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAvatarRequest.ProtoReflect.Descriptor instead.
func (*UpdateAvatarRequest) Descriptor() ([]byte, []int) {
	return file_userpb_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateAvatarRequest) GetAvatarFileId() int32 {
	if x != nil {
		return x.AvatarFileId
	}
	return 0
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Email      string  `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Username   string  `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Name       string  `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Lastname   string  `protobuf:"bytes,5,opt,name=lastname,proto3" json:"lastname,omitempty"`
	AvatarPath string  `protobuf:"bytes,6,opt,name=avatarPath,proto3" json:"avatarPath,omitempty"`
	Password   *string `protobuf:"bytes,7,opt,name=password,proto3,oneof" json:"password,omitempty"`
	DistrictID int32   `protobuf:"varint,8,opt,name=districtID,proto3" json:"districtID,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userpb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_userpb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_userpb_proto_rawDescGZIP(), []int{3}
}

func (x *User) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *User) GetAvatarPath() string {
	if x != nil {
		return x.AvatarPath
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil && x.Password != nil {
		return *x.Password
	}
	return ""
}

func (x *User) GetDistrictID() int32 {
	if x != nil {
		return x.DistrictID
	}
	return 0
}

type UserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UserResponse) Reset() {
	*x = UserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userpb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponse) ProtoMessage() {}

func (x *UserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_userpb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponse.ProtoReflect.Descriptor instead.
func (*UserResponse) Descriptor() ([]byte, []int) {
	return file_userpb_proto_rawDescGZIP(), []int{4}
}

func (x *UserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type UserListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User       []*User                      `protobuf:"bytes,1,rep,name=user,proto3" json:"user,omitempty"`
	Pagination *commonpb.PaginationResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *UserListResponse) Reset() {
	*x = UserListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userpb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListResponse) ProtoMessage() {}

func (x *UserListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_userpb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserListResponse.ProtoReflect.Descriptor instead.
func (*UserListResponse) Descriptor() ([]byte, []int) {
	return file_userpb_proto_rawDescGZIP(), []int{5}
}

func (x *UserListResponse) GetUser() []*User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserListResponse) GetPagination() *commonpb.PaginationResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type FollowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User        *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"` // It can be user Id
	IsFollowing bool  `protobuf:"varint,2,opt,name=isFollowing,proto3" json:"isFollowing,omitempty"`
}

func (x *FollowResponse) Reset() {
	*x = FollowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userpb_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowResponse) ProtoMessage() {}

func (x *FollowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_userpb_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowResponse.ProtoReflect.Descriptor instead.
func (*FollowResponse) Descriptor() ([]byte, []int) {
	return file_userpb_proto_rawDescGZIP(), []int{6}
}

func (x *FollowResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *FollowResponse) GetIsFollowing() bool {
	if x != nil {
		return x.IsFollowing
	}
	return false
}

type BlockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User      *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"` // It can be user Id
	IsBlocked bool  `protobuf:"varint,2,opt,name=isBlocked,proto3" json:"isBlocked,omitempty"`
}

func (x *BlockResponse) Reset() {
	*x = BlockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userpb_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockResponse) ProtoMessage() {}

func (x *BlockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_userpb_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockResponse.ProtoReflect.Descriptor instead.
func (*BlockResponse) Descriptor() ([]byte, []int) {
	return file_userpb_proto_rawDescGZIP(), []int{7}
}

func (x *BlockResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *BlockResponse) GetIsBlocked() bool {
	if x != nil {
		return x.IsBlocked
	}
	return false
}

var File_userpb_proto protoreflect.FileDescriptor

var file_userpb_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x1a, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x62, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x73, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x05, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x39,
	0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0a, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xa9, 0x01, 0x0a, 0x11, 0x55, 0x73,
	0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x08, 0x6c, 0x61,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x64, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x63, 0x74, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x48, 0x02, 0x52,
	0x0a, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x49, 0x44, 0x88, 0x01, 0x01, 0x42, 0x07,
	0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6c, 0x61, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63,
	0x74, 0x49, 0x44, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x4a,
	0x04, 0x08, 0x03, 0x10, 0x04, 0x22, 0x39, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0c, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x64,
	0x22, 0xe6, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x50, 0x61, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1f, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1e, 0x0a, 0x0a,
	0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x49, 0x44, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x49, 0x44, 0x42, 0x0b, 0x0a, 0x09,
	0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x2e, 0x0a, 0x0c, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x6e, 0x0a, 0x10, 0x55, 0x73, 0x65,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x3a, 0x0a,
	0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x52, 0x0a, 0x0e, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x69,
	0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0b, 0x69, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x22, 0x4d, 0x0a,
	0x0d, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1c,
	0x0a, 0x09, 0x69, 0x73, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x69, 0x73, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x32, 0xe4, 0x03, 0x0a,
	0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x05,
	0x47, 0x65, 0x74, 0x4d, 0x65, 0x12, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3e, 0x0a, 0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x17,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x39, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x19, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x54, 0x6f, 0x67,
	0x67, 0x6c, 0x65, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x54, 0x6f, 0x67, 0x67, 0x6c,
	0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_userpb_proto_rawDescOnce sync.Once
	file_userpb_proto_rawDescData = file_userpb_proto_rawDesc
)

func file_userpb_proto_rawDescGZIP() []byte {
	file_userpb_proto_rawDescOnce.Do(func() {
		file_userpb_proto_rawDescData = protoimpl.X.CompressGZIP(file_userpb_proto_rawDescData)
	})
	return file_userpb_proto_rawDescData
}

var file_userpb_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_userpb_proto_goTypes = []interface{}{
	(*UserSearchRequest)(nil),           // 0: user.UserSearchRequest
	(*UserUpdateRequest)(nil),           // 1: user.UserUpdateRequest
	(*UpdateAvatarRequest)(nil),         // 2: user.UpdateAvatarRequest
	(*User)(nil),                        // 3: user.User
	(*UserResponse)(nil),                // 4: user.UserResponse
	(*UserListResponse)(nil),            // 5: user.UserListResponse
	(*FollowResponse)(nil),              // 6: user.FollowResponse
	(*BlockResponse)(nil),               // 7: user.BlockResponse
	(*commonpb.Query)(nil),              // 8: common.Query
	(*commonpb.PaginationRequest)(nil),  // 9: common.PaginationRequest
	(*commonpb.PaginationResponse)(nil), // 10: common.PaginationResponse
	(*commonpb.Empty)(nil),              // 11: common.Empty
	(*commonpb.GetByIdRequest)(nil),     // 12: common.GetByIdRequest
	(*commonpb.CommentRequest)(nil),     // 13: common.CommentRequest
	(*commonpb.CommentResponse)(nil),    // 14: common.CommentResponse
}
var file_userpb_proto_depIdxs = []int32{
	8,  // 0: user.UserSearchRequest.query:type_name -> common.Query
	9,  // 1: user.UserSearchRequest.pagination:type_name -> common.PaginationRequest
	3,  // 2: user.UserResponse.user:type_name -> user.User
	3,  // 3: user.UserListResponse.user:type_name -> user.User
	10, // 4: user.UserListResponse.pagination:type_name -> common.PaginationResponse
	3,  // 5: user.FollowResponse.user:type_name -> user.User
	3,  // 6: user.BlockResponse.user:type_name -> user.User
	11, // 7: user.UserService.GetMe:input_type -> common.Empty
	12, // 8: user.UserService.GetUser:input_type -> common.GetByIdRequest
	0,  // 9: user.UserService.SearchUsers:input_type -> user.UserSearchRequest
	1,  // 10: user.UserService.UpdateUser:input_type -> user.UserUpdateRequest
	2,  // 11: user.UserService.UploadAvatar:input_type -> user.UpdateAvatarRequest
	12, // 12: user.UserService.ToggleFollow:input_type -> common.GetByIdRequest
	12, // 13: user.UserService.ToggleBlock:input_type -> common.GetByIdRequest
	13, // 14: user.UserService.AddComment:input_type -> common.CommentRequest
	4,  // 15: user.UserService.GetMe:output_type -> user.UserResponse
	4,  // 16: user.UserService.GetUser:output_type -> user.UserResponse
	5,  // 17: user.UserService.SearchUsers:output_type -> user.UserListResponse
	4,  // 18: user.UserService.UpdateUser:output_type -> user.UserResponse
	4,  // 19: user.UserService.UploadAvatar:output_type -> user.UserResponse
	6,  // 20: user.UserService.ToggleFollow:output_type -> user.FollowResponse
	6,  // 21: user.UserService.ToggleBlock:output_type -> user.FollowResponse
	14, // 22: user.UserService.AddComment:output_type -> common.CommentResponse
	15, // [15:23] is the sub-list for method output_type
	7,  // [7:15] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_userpb_proto_init() }
func file_userpb_proto_init() {
	if File_userpb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_userpb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserSearchRequest); i {
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
		file_userpb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserUpdateRequest); i {
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
		file_userpb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAvatarRequest); i {
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
		file_userpb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_userpb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserResponse); i {
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
		file_userpb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserListResponse); i {
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
		file_userpb_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowResponse); i {
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
		file_userpb_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockResponse); i {
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
	file_userpb_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_userpb_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_userpb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_userpb_proto_goTypes,
		DependencyIndexes: file_userpb_proto_depIdxs,
		MessageInfos:      file_userpb_proto_msgTypes,
	}.Build()
	File_userpb_proto = out.File
	file_userpb_proto_rawDesc = nil
	file_userpb_proto_goTypes = nil
	file_userpb_proto_depIdxs = nil
}
