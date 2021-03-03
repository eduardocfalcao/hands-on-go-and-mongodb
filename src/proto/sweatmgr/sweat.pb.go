// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.3
// source: src/proto/sweatmgr/sweat.proto

package sweatmgr

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Sweat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreatedAt       *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Glucose         float32                `protobuf:"fixed32,2,opt,name=glucose,proto3" json:"glucose,omitempty"`
	Chloride        float32                `protobuf:"fixed32,3,opt,name=chloride,proto3" json:"chloride,omitempty"`
	Sodium          float32                `protobuf:"fixed32,4,opt,name=sodium,proto3" json:"sodium,omitempty"`
	Potassium       float32                `protobuf:"fixed32,5,opt,name=potassium,proto3" json:"potassium,omitempty"`
	Magnesium       float32                `protobuf:"fixed32,6,opt,name=magnesium,proto3" json:"magnesium,omitempty"`
	Calcium         float32                `protobuf:"fixed32,7,opt,name=calcium,proto3" json:"calcium,omitempty"`
	Humidity        float32                `protobuf:"fixed32,8,opt,name=humidity,proto3" json:"humidity,omitempty"`
	RoomTemperature float32                `protobuf:"fixed32,9,opt,name=room_temperature,json=roomTemperature,proto3" json:"room_temperature,omitempty"`
	BodyTemperature float32                `protobuf:"fixed32,10,opt,name=body_temperature,json=bodyTemperature,proto3" json:"body_temperature,omitempty"`
	Heartbeat       int32                  `protobuf:"varint,11,opt,name=heartbeat,proto3" json:"heartbeat,omitempty"`
}

func (x *Sweat) Reset() {
	*x = Sweat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_sweatmgr_sweat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sweat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sweat) ProtoMessage() {}

func (x *Sweat) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_sweatmgr_sweat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sweat.ProtoReflect.Descriptor instead.
func (*Sweat) Descriptor() ([]byte, []int) {
	return file_src_proto_sweatmgr_sweat_proto_rawDescGZIP(), []int{0}
}

func (x *Sweat) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Sweat) GetGlucose() float32 {
	if x != nil {
		return x.Glucose
	}
	return 0
}

func (x *Sweat) GetChloride() float32 {
	if x != nil {
		return x.Chloride
	}
	return 0
}

func (x *Sweat) GetSodium() float32 {
	if x != nil {
		return x.Sodium
	}
	return 0
}

func (x *Sweat) GetPotassium() float32 {
	if x != nil {
		return x.Potassium
	}
	return 0
}

func (x *Sweat) GetMagnesium() float32 {
	if x != nil {
		return x.Magnesium
	}
	return 0
}

func (x *Sweat) GetCalcium() float32 {
	if x != nil {
		return x.Calcium
	}
	return 0
}

func (x *Sweat) GetHumidity() float32 {
	if x != nil {
		return x.Humidity
	}
	return 0
}

func (x *Sweat) GetRoomTemperature() float32 {
	if x != nil {
		return x.RoomTemperature
	}
	return 0
}

func (x *Sweat) GetBodyTemperature() float32 {
	if x != nil {
		return x.BodyTemperature
	}
	return 0
}

func (x *Sweat) GetHeartbeat() int32 {
	if x != nil {
		return x.Heartbeat
	}
	return 0
}

type SweatStatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userid string `protobuf:"bytes,1,opt,name=userid,proto3" json:"userid,omitempty"`
}

func (x *SweatStatsRequest) Reset() {
	*x = SweatStatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_sweatmgr_sweat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SweatStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SweatStatsRequest) ProtoMessage() {}

func (x *SweatStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_sweatmgr_sweat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SweatStatsRequest.ProtoReflect.Descriptor instead.
func (*SweatStatsRequest) Descriptor() ([]byte, []int) {
	return file_src_proto_sweatmgr_sweat_proto_rawDescGZIP(), []int{1}
}

func (x *SweatStatsRequest) GetUserid() string {
	if x != nil {
		return x.Userid
	}
	return ""
}

type SweatStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Userid string   `protobuf:"bytes,1,opt,name=userid,proto3" json:"userid,omitempty"`
	Sweat  []*Sweat `protobuf:"bytes,2,rep,name=sweat,proto3" json:"sweat,omitempty"`
}

func (x *SweatStatsResponse) Reset() {
	*x = SweatStatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_proto_sweatmgr_sweat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SweatStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SweatStatsResponse) ProtoMessage() {}

func (x *SweatStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_proto_sweatmgr_sweat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SweatStatsResponse.ProtoReflect.Descriptor instead.
func (*SweatStatsResponse) Descriptor() ([]byte, []int) {
	return file_src_proto_sweatmgr_sweat_proto_rawDescGZIP(), []int{2}
}

func (x *SweatStatsResponse) GetUserid() string {
	if x != nil {
		return x.Userid
	}
	return ""
}

func (x *SweatStatsResponse) GetSweat() []*Sweat {
	if x != nil {
		return x.Sweat
	}
	return nil
}

var File_src_proto_sweatmgr_sweat_proto protoreflect.FileDescriptor

var file_src_proto_sweatmgr_sweat_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x77, 0x65, 0x61,
	0x74, 0x6d, 0x67, 0x72, 0x2f, 0x73, 0x77, 0x65, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x73, 0x77, 0x65, 0x61, 0x74, 0x6d, 0x67, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf6, 0x02, 0x0a, 0x05,
	0x53, 0x77, 0x65, 0x61, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x67, 0x6c, 0x75, 0x63, 0x6f, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x07, 0x67, 0x6c, 0x75, 0x63, 0x6f, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68,
	0x6c, 0x6f, 0x72, 0x69, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x63, 0x68,
	0x6c, 0x6f, 0x72, 0x69, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x64, 0x69, 0x75, 0x6d,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x73, 0x6f, 0x64, 0x69, 0x75, 0x6d, 0x12, 0x1c,
	0x0a, 0x09, 0x70, 0x6f, 0x74, 0x61, 0x73, 0x73, 0x69, 0x75, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x09, 0x70, 0x6f, 0x74, 0x61, 0x73, 0x73, 0x69, 0x75, 0x6d, 0x12, 0x1c, 0x0a, 0x09,
	0x6d, 0x61, 0x67, 0x6e, 0x65, 0x73, 0x69, 0x75, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x09, 0x6d, 0x61, 0x67, 0x6e, 0x65, 0x73, 0x69, 0x75, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61,
	0x6c, 0x63, 0x69, 0x75, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x63, 0x61, 0x6c,
	0x63, 0x69, 0x75, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x75, 0x6d, 0x69, 0x64, 0x69, 0x74, 0x79,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x68, 0x75, 0x6d, 0x69, 0x64, 0x69, 0x74, 0x79,
	0x12, 0x29, 0x0a, 0x10, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0f, 0x72, 0x6f, 0x6f, 0x6d,
	0x54, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x62,
	0x6f, 0x64, 0x79, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0f, 0x62, 0x6f, 0x64, 0x79, 0x54, 0x65, 0x6d, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62,
	0x65, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x68, 0x65, 0x61, 0x72, 0x74,
	0x62, 0x65, 0x61, 0x74, 0x22, 0x2b, 0x0a, 0x11, 0x53, 0x77, 0x65, 0x61, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69,
	0x64, 0x22, 0x53, 0x0a, 0x12, 0x53, 0x77, 0x65, 0x61, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x12,
	0x25, 0x0a, 0x05, 0x73, 0x77, 0x65, 0x61, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x73, 0x77, 0x65, 0x61, 0x74, 0x6d, 0x67, 0x72, 0x2e, 0x53, 0x77, 0x65, 0x61, 0x74, 0x52,
	0x05, 0x73, 0x77, 0x65, 0x61, 0x74, 0x32, 0x53, 0x0a, 0x03, 0x41, 0x70, 0x69, 0x12, 0x4c, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x53, 0x77, 0x65, 0x61, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x1b,
	0x2e, 0x73, 0x77, 0x65, 0x61, 0x74, 0x6d, 0x67, 0x72, 0x2e, 0x53, 0x77, 0x65, 0x61, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x77,
	0x65, 0x61, 0x74, 0x6d, 0x67, 0x72, 0x2e, 0x53, 0x77, 0x65, 0x61, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_src_proto_sweatmgr_sweat_proto_rawDescOnce sync.Once
	file_src_proto_sweatmgr_sweat_proto_rawDescData = file_src_proto_sweatmgr_sweat_proto_rawDesc
)

func file_src_proto_sweatmgr_sweat_proto_rawDescGZIP() []byte {
	file_src_proto_sweatmgr_sweat_proto_rawDescOnce.Do(func() {
		file_src_proto_sweatmgr_sweat_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_proto_sweatmgr_sweat_proto_rawDescData)
	})
	return file_src_proto_sweatmgr_sweat_proto_rawDescData
}

var file_src_proto_sweatmgr_sweat_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_src_proto_sweatmgr_sweat_proto_goTypes = []interface{}{
	(*Sweat)(nil),                 // 0: sweatmgr.Sweat
	(*SweatStatsRequest)(nil),     // 1: sweatmgr.SweatStatsRequest
	(*SweatStatsResponse)(nil),    // 2: sweatmgr.SweatStatsResponse
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_src_proto_sweatmgr_sweat_proto_depIdxs = []int32{
	3, // 0: sweatmgr.Sweat.created_at:type_name -> google.protobuf.Timestamp
	0, // 1: sweatmgr.SweatStatsResponse.sweat:type_name -> sweatmgr.Sweat
	1, // 2: sweatmgr.Api.GetSweatStats:input_type -> sweatmgr.SweatStatsRequest
	2, // 3: sweatmgr.Api.GetSweatStats:output_type -> sweatmgr.SweatStatsResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_src_proto_sweatmgr_sweat_proto_init() }
func file_src_proto_sweatmgr_sweat_proto_init() {
	if File_src_proto_sweatmgr_sweat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_proto_sweatmgr_sweat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sweat); i {
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
		file_src_proto_sweatmgr_sweat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SweatStatsRequest); i {
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
		file_src_proto_sweatmgr_sweat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SweatStatsResponse); i {
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
			RawDescriptor: file_src_proto_sweatmgr_sweat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_src_proto_sweatmgr_sweat_proto_goTypes,
		DependencyIndexes: file_src_proto_sweatmgr_sweat_proto_depIdxs,
		MessageInfos:      file_src_proto_sweatmgr_sweat_proto_msgTypes,
	}.Build()
	File_src_proto_sweatmgr_sweat_proto = out.File
	file_src_proto_sweatmgr_sweat_proto_rawDesc = nil
	file_src_proto_sweatmgr_sweat_proto_goTypes = nil
	file_src_proto_sweatmgr_sweat_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ApiClient is the client API for Api service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ApiClient interface {
	GetSweatStats(ctx context.Context, in *SweatStatsRequest, opts ...grpc.CallOption) (*SweatStatsResponse, error)
}

type apiClient struct {
	cc grpc.ClientConnInterface
}

func NewApiClient(cc grpc.ClientConnInterface) ApiClient {
	return &apiClient{cc}
}

func (c *apiClient) GetSweatStats(ctx context.Context, in *SweatStatsRequest, opts ...grpc.CallOption) (*SweatStatsResponse, error) {
	out := new(SweatStatsResponse)
	err := c.cc.Invoke(ctx, "/sweatmgr.Api/GetSweatStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiServer is the server API for Api service.
type ApiServer interface {
	GetSweatStats(context.Context, *SweatStatsRequest) (*SweatStatsResponse, error)
}

// UnimplementedApiServer can be embedded to have forward compatible implementations.
type UnimplementedApiServer struct {
}

func (*UnimplementedApiServer) GetSweatStats(context.Context, *SweatStatsRequest) (*SweatStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSweatStats not implemented")
}

func RegisterApiServer(s *grpc.Server, srv ApiServer) {
	s.RegisterService(&_Api_serviceDesc, srv)
}

func _Api_GetSweatStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SweatStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).GetSweatStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sweatmgr.Api/GetSweatStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).GetSweatStats(ctx, req.(*SweatStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Api_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sweatmgr.Api",
	HandlerType: (*ApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSweatStats",
			Handler:    _Api_GetSweatStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/proto/sweatmgr/sweat.proto",
}
