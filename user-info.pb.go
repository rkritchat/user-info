// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: user-info.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Firstname string `protobuf:"bytes,2,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname  string `protobuf:"bytes,3,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Email     string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_user_info_proto_msgTypes[0]
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
	return file_user_info_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *User) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type GetUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_user_info_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type GetUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Detail *User `protobuf:"bytes,1,opt,name=detail,proto3" json:"detail,omitempty"`
}

func (x *GetUserResponse) Reset() {
	*x = GetUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_info_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserResponse) ProtoMessage() {}

func (x *GetUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_info_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserResponse.ProtoReflect.Descriptor instead.
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return file_user_info_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserResponse) GetDetail() *User {
	if x != nil {
		return x.Detail
	}
	return nil
}

type GetAllUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllUserRequest) Reset() {
	*x = GetAllUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_info_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllUserRequest) ProtoMessage() {}

func (x *GetAllUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_info_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllUserRequest.ProtoReflect.Descriptor instead.
func (*GetAllUserRequest) Descriptor() ([]byte, []int) {
	return file_user_info_proto_rawDescGZIP(), []int{3}
}

type GetAllUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Details []*User `protobuf:"bytes,1,rep,name=details,proto3" json:"details,omitempty"`
}

func (x *GetAllUserResponse) Reset() {
	*x = GetAllUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_info_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllUserResponse) ProtoMessage() {}

func (x *GetAllUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_info_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllUserResponse.ProtoReflect.Descriptor instead.
func (*GetAllUserResponse) Descriptor() ([]byte, []int) {
	return file_user_info_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllUserResponse) GetDetails() []*User {
	if x != nil {
		return x.Details
	}
	return nil
}

var File_user_info_proto protoreflect.FileDescriptor

var file_user_info_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x75, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x66, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x22, 0x26, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x39, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26,
	0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x75, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x52, 0x06,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3e, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x28, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x75, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x32, 0x9e, 0x01, 0x0a, 0x0f,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x40, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x75, 0x65, 0x72,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x75, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x49, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x1b, 0x2e, 0x75, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x75,
	0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x04, 0x5a, 0x02,
	0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_info_proto_rawDescOnce sync.Once
	file_user_info_proto_rawDescData = file_user_info_proto_rawDesc
)

func file_user_info_proto_rawDescGZIP() []byte {
	file_user_info_proto_rawDescOnce.Do(func() {
		file_user_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_info_proto_rawDescData)
	})
	return file_user_info_proto_rawDescData
}

var file_user_info_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_user_info_proto_goTypes = []interface{}{
	(*User)(nil),               // 0: uer_info.user
	(*GetUserRequest)(nil),     // 1: uer_info.GetUserRequest
	(*GetUserResponse)(nil),    // 2: uer_info.GetUserResponse
	(*GetAllUserRequest)(nil),  // 3: uer_info.GetAllUserRequest
	(*GetAllUserResponse)(nil), // 4: uer_info.GetAllUserResponse
}
var file_user_info_proto_depIdxs = []int32{
	0, // 0: uer_info.GetUserResponse.detail:type_name -> uer_info.user
	0, // 1: uer_info.GetAllUserResponse.details:type_name -> uer_info.user
	1, // 2: uer_info.UserInfoService.GetUser:input_type -> uer_info.GetUserRequest
	3, // 3: uer_info.UserInfoService.GetAllUser:input_type -> uer_info.GetAllUserRequest
	2, // 4: uer_info.UserInfoService.GetUser:output_type -> uer_info.GetUserResponse
	4, // 5: uer_info.UserInfoService.GetAllUser:output_type -> uer_info.GetAllUserResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_user_info_proto_init() }
func file_user_info_proto_init() {
	if File_user_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_user_info_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRequest); i {
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
		file_user_info_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserResponse); i {
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
		file_user_info_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllUserRequest); i {
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
		file_user_info_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllUserResponse); i {
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
			RawDescriptor: file_user_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_info_proto_goTypes,
		DependencyIndexes: file_user_info_proto_depIdxs,
		MessageInfos:      file_user_info_proto_msgTypes,
	}.Build()
	File_user_info_proto = out.File
	file_user_info_proto_rawDesc = nil
	file_user_info_proto_goTypes = nil
	file_user_info_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserInfoServiceClient is the client API for UserInfoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserInfoServiceClient interface {
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	GetAllUser(ctx context.Context, in *GetAllUserRequest, opts ...grpc.CallOption) (*GetAllUserResponse, error)
}

type userInfoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserInfoServiceClient(cc grpc.ClientConnInterface) UserInfoServiceClient {
	return &userInfoServiceClient{cc}
}

func (c *userInfoServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/uer_info.UserInfoService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userInfoServiceClient) GetAllUser(ctx context.Context, in *GetAllUserRequest, opts ...grpc.CallOption) (*GetAllUserResponse, error) {
	out := new(GetAllUserResponse)
	err := c.cc.Invoke(ctx, "/uer_info.UserInfoService/GetAllUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserInfoServiceServer is the server API for UserInfoService service.
type UserInfoServiceServer interface {
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	GetAllUser(context.Context, *GetAllUserRequest) (*GetAllUserResponse, error)
}

// UnimplementedUserInfoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserInfoServiceServer struct {
}

func (*UnimplementedUserInfoServiceServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (*UnimplementedUserInfoServiceServer) GetAllUser(context.Context, *GetAllUserRequest) (*GetAllUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUser not implemented")
}

func RegisterUserInfoServiceServer(s *grpc.Server, srv UserInfoServiceServer) {
	s.RegisterService(&_UserInfoService_serviceDesc, srv)
}

func _UserInfoService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInfoServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/uer_info.UserInfoService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInfoServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserInfoService_GetAllUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInfoServiceServer).GetAllUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/uer_info.UserInfoService/GetAllUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInfoServiceServer).GetAllUser(ctx, req.(*GetAllUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserInfoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "uer_info.UserInfoService",
	HandlerType: (*UserInfoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _UserInfoService_GetUser_Handler,
		},
		{
			MethodName: "GetAllUser",
			Handler:    _UserInfoService_GetAllUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user-info.proto",
}
