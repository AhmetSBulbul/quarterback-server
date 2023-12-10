// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.3
// source: court.proto

package court

import (
	common "github.com/AhmetSBulbul/quarterback-server/api/pb/common"
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

type Court struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	DistrictId int32            `protobuf:"varint,3,opt,name=districtId,proto3" json:"districtId,omitempty"`
	Address    string           `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	Location   *common.Location `protobuf:"bytes,5,opt,name=location,proto3" json:"location,omitempty"`
	Medias     []*common.Media  `protobuf:"bytes,6,rep,name=medias,proto3" json:"medias,omitempty"`
}

func (x *Court) Reset() {
	*x = Court{}
	if protoimpl.UnsafeEnabled {
		mi := &file_court_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Court) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Court) ProtoMessage() {}

func (x *Court) ProtoReflect() protoreflect.Message {
	mi := &file_court_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Court.ProtoReflect.Descriptor instead.
func (*Court) Descriptor() ([]byte, []int) {
	return file_court_proto_rawDescGZIP(), []int{0}
}

func (x *Court) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Court) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Court) GetDistrictId() int32 {
	if x != nil {
		return x.DistrictId
	}
	return 0
}

func (x *Court) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Court) GetLocation() *common.Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *Court) GetMedias() []*common.Media {
	if x != nil {
		return x.Medias
	}
	return nil
}

// Get Court
type GetCourtRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCourtRequest) Reset() {
	*x = GetCourtRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_court_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCourtRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCourtRequest) ProtoMessage() {}

func (x *GetCourtRequest) ProtoReflect() protoreflect.Message {
	mi := &file_court_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCourtRequest.ProtoReflect.Descriptor instead.
func (*GetCourtRequest) Descriptor() ([]byte, []int) {
	return file_court_proto_rawDescGZIP(), []int{1}
}

func (x *GetCourtRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CourtResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Court *Court `protobuf:"bytes,1,opt,name=court,proto3" json:"court,omitempty"`
}

func (x *CourtResponse) Reset() {
	*x = CourtResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_court_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourtResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourtResponse) ProtoMessage() {}

func (x *CourtResponse) ProtoReflect() protoreflect.Message {
	mi := &file_court_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourtResponse.ProtoReflect.Descriptor instead.
func (*CourtResponse) Descriptor() ([]byte, []int) {
	return file_court_proto_rawDescGZIP(), []int{2}
}

func (x *CourtResponse) GetCourt() *Court {
	if x != nil {
		return x.Court
	}
	return nil
}

type ListCourtByLocationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location   *common.Location          `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	Pagination *common.PaginationRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *ListCourtByLocationRequest) Reset() {
	*x = ListCourtByLocationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_court_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCourtByLocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCourtByLocationRequest) ProtoMessage() {}

func (x *ListCourtByLocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_court_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCourtByLocationRequest.ProtoReflect.Descriptor instead.
func (*ListCourtByLocationRequest) Descriptor() ([]byte, []int) {
	return file_court_proto_rawDescGZIP(), []int{3}
}

func (x *ListCourtByLocationRequest) GetLocation() *common.Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *ListCourtByLocationRequest) GetPagination() *common.PaginationRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type SearchCourtRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query      *common.Query             `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Pagination *common.PaginationRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *SearchCourtRequest) Reset() {
	*x = SearchCourtRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_court_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchCourtRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchCourtRequest) ProtoMessage() {}

func (x *SearchCourtRequest) ProtoReflect() protoreflect.Message {
	mi := &file_court_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchCourtRequest.ProtoReflect.Descriptor instead.
func (*SearchCourtRequest) Descriptor() ([]byte, []int) {
	return file_court_proto_rawDescGZIP(), []int{4}
}

func (x *SearchCourtRequest) GetQuery() *common.Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *SearchCourtRequest) GetPagination() *common.PaginationRequest {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type ListCourtResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Courts     []*Court                   `protobuf:"bytes,1,rep,name=courts,proto3" json:"courts,omitempty"`
	Pagination *common.PaginationResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *ListCourtResponse) Reset() {
	*x = ListCourtResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_court_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCourtResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCourtResponse) ProtoMessage() {}

func (x *ListCourtResponse) ProtoReflect() protoreflect.Message {
	mi := &file_court_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCourtResponse.ProtoReflect.Descriptor instead.
func (*ListCourtResponse) Descriptor() ([]byte, []int) {
	return file_court_proto_rawDescGZIP(), []int{5}
}

func (x *ListCourtResponse) GetCourts() []*Court {
	if x != nil {
		return x.Courts
	}
	return nil
}

func (x *ListCourtResponse) GetPagination() *common.PaginationResponse {
	if x != nil {
		return x.Pagination
	}
	return nil
}

var File_court_proto protoreflect.FileDescriptor

var file_court_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x63,
	0x6f, 0x75, 0x72, 0x74, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xba, 0x01, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2c, 0x0a, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x06, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x06, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x22,
	0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x33, 0x0a, 0x0d, 0x43, 0x6f, 0x75, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x74, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x74,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x72, 0x74, 0x22, 0x85, 0x01, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x6f, 0x75, 0x72, 0x74, 0x42, 0x79, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x74, 0x0a, 0x12, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x6f, 0x75, 0x72, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x39, 0x0a, 0x0a, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x75, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x06, 0x63, 0x6f,
	0x75, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x75,
	0x72, 0x74, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x74, 0x52, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x74, 0x73,
	0x12, 0x3a, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xe0, 0x01, 0x0a,
	0x0c, 0x43, 0x6f, 0x75, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x74, 0x12, 0x16, 0x2e, 0x63, 0x6f, 0x75, 0x72,
	0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x74, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x6f, 0x75, 0x72, 0x74, 0x42, 0x79, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21,
	0x2e, 0x63, 0x6f, 0x75, 0x72, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x74,
	0x42, 0x79, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f,
	0x75, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0b, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x6f, 0x75, 0x72, 0x74, 0x12, 0x19, 0x2e, 0x63, 0x6f, 0x75,
	0x72, 0x74, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x6f, 0x75, 0x72, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x74, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_court_proto_rawDescOnce sync.Once
	file_court_proto_rawDescData = file_court_proto_rawDesc
)

func file_court_proto_rawDescGZIP() []byte {
	file_court_proto_rawDescOnce.Do(func() {
		file_court_proto_rawDescData = protoimpl.X.CompressGZIP(file_court_proto_rawDescData)
	})
	return file_court_proto_rawDescData
}

var file_court_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_court_proto_goTypes = []interface{}{
	(*Court)(nil),                      // 0: court.Court
	(*GetCourtRequest)(nil),            // 1: court.GetCourtRequest
	(*CourtResponse)(nil),              // 2: court.CourtResponse
	(*ListCourtByLocationRequest)(nil), // 3: court.ListCourtByLocationRequest
	(*SearchCourtRequest)(nil),         // 4: court.SearchCourtRequest
	(*ListCourtResponse)(nil),          // 5: court.ListCourtResponse
	(*common.Location)(nil),            // 6: common.Location
	(*common.Media)(nil),               // 7: common.Media
	(*common.PaginationRequest)(nil),   // 8: common.PaginationRequest
	(*common.Query)(nil),               // 9: common.Query
	(*common.PaginationResponse)(nil),  // 10: common.PaginationResponse
}
var file_court_proto_depIdxs = []int32{
	6,  // 0: court.Court.location:type_name -> common.Location
	7,  // 1: court.Court.medias:type_name -> common.Media
	0,  // 2: court.CourtResponse.court:type_name -> court.Court
	6,  // 3: court.ListCourtByLocationRequest.location:type_name -> common.Location
	8,  // 4: court.ListCourtByLocationRequest.pagination:type_name -> common.PaginationRequest
	9,  // 5: court.SearchCourtRequest.query:type_name -> common.Query
	8,  // 6: court.SearchCourtRequest.pagination:type_name -> common.PaginationRequest
	0,  // 7: court.ListCourtResponse.courts:type_name -> court.Court
	10, // 8: court.ListCourtResponse.pagination:type_name -> common.PaginationResponse
	1,  // 9: court.CourtService.GetCourt:input_type -> court.GetCourtRequest
	3,  // 10: court.CourtService.ListCourtByLocation:input_type -> court.ListCourtByLocationRequest
	4,  // 11: court.CourtService.SearchCourt:input_type -> court.SearchCourtRequest
	2,  // 12: court.CourtService.GetCourt:output_type -> court.CourtResponse
	5,  // 13: court.CourtService.ListCourtByLocation:output_type -> court.ListCourtResponse
	5,  // 14: court.CourtService.SearchCourt:output_type -> court.ListCourtResponse
	12, // [12:15] is the sub-list for method output_type
	9,  // [9:12] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_court_proto_init() }
func file_court_proto_init() {
	if File_court_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_court_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Court); i {
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
		file_court_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCourtRequest); i {
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
		file_court_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CourtResponse); i {
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
		file_court_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCourtByLocationRequest); i {
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
		file_court_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchCourtRequest); i {
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
		file_court_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCourtResponse); i {
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
			RawDescriptor: file_court_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_court_proto_goTypes,
		DependencyIndexes: file_court_proto_depIdxs,
		MessageInfos:      file_court_proto_msgTypes,
	}.Build()
	File_court_proto = out.File
	file_court_proto_rawDesc = nil
	file_court_proto_goTypes = nil
	file_court_proto_depIdxs = nil
}
