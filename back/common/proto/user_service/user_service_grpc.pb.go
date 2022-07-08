// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package user

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	GetByUsername(ctx context.Context, in *GetByUsernameRequest, opts ...grpc.CallOption) (*GetByUsernameResponse, error)
	GetPublicUserByUsername(ctx context.Context, in *GetPublicUserByUsernameRequest, opts ...grpc.CallOption) (*GetPublicUserByUsernameResponse, error)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	GetAllPublicUsers(ctx context.Context, in *GetAllPublicUsersRequest, opts ...grpc.CallOption) (*GetAllPublicUsersResponse, error)
	GetAllPublicUsersByUsername(ctx context.Context, in *GetAllPublicUsersByUsernameRequest, opts ...grpc.CallOption) (*GetAllPublicUsersByUsernameResponse, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	UpdateBasicInfo(ctx context.Context, in *UpdateBasicInfoRequest, opts ...grpc.CallOption) (*UpdateBasicInfoResponse, error)
	UpdateAdvancedInfo(ctx context.Context, in *UpdateAdvancedInfoRequest, opts ...grpc.CallOption) (*UpdateAdvancedInfoResponse, error)
	UpdatePersonalInfo(ctx context.Context, in *UpdatePersonalInfoRequest, opts ...grpc.CallOption) (*UpdatePersonalInfoResponse, error)
	UpdateAllInfo(ctx context.Context, in *UpdateAllInfoRequest, opts ...grpc.CallOption) (*UpdateAllInfoResponse, error)
	FollowPublicProfile(ctx context.Context, in *FollowPublicProfileRequest, opts ...grpc.CallOption) (*FollowPublicProfileResponse, error)
	AcceptFollowingRequest(ctx context.Context, in *AcceptFollowingRequestRequest, opts ...grpc.CallOption) (*AcceptFollowingRequestResponse, error)
	RejectFollowingRequest(ctx context.Context, in *RejectFollowingRequestRequest, opts ...grpc.CallOption) (*RejectFollowingRequestResponse, error)
	BlockUser(ctx context.Context, in *BlockUserRequest, opts ...grpc.CallOption) (*BlockUserResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetByUsername(ctx context.Context, in *GetByUsernameRequest, opts ...grpc.CallOption) (*GetByUsernameResponse, error) {
	out := new(GetByUsernameResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/GetByUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetPublicUserByUsername(ctx context.Context, in *GetPublicUserByUsernameRequest, opts ...grpc.CallOption) (*GetPublicUserByUsernameResponse, error) {
	out := new(GetPublicUserByUsernameResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/GetPublicUserByUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAllPublicUsers(ctx context.Context, in *GetAllPublicUsersRequest, opts ...grpc.CallOption) (*GetAllPublicUsersResponse, error) {
	out := new(GetAllPublicUsersResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/GetAllPublicUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAllPublicUsersByUsername(ctx context.Context, in *GetAllPublicUsersByUsernameRequest, opts ...grpc.CallOption) (*GetAllPublicUsersByUsernameResponse, error) {
	out := new(GetAllPublicUsersByUsernameResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/GetAllPublicUsersByUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateBasicInfo(ctx context.Context, in *UpdateBasicInfoRequest, opts ...grpc.CallOption) (*UpdateBasicInfoResponse, error) {
	out := new(UpdateBasicInfoResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/UpdateBasicInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateAdvancedInfo(ctx context.Context, in *UpdateAdvancedInfoRequest, opts ...grpc.CallOption) (*UpdateAdvancedInfoResponse, error) {
	out := new(UpdateAdvancedInfoResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/UpdateAdvancedInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdatePersonalInfo(ctx context.Context, in *UpdatePersonalInfoRequest, opts ...grpc.CallOption) (*UpdatePersonalInfoResponse, error) {
	out := new(UpdatePersonalInfoResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/UpdatePersonalInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateAllInfo(ctx context.Context, in *UpdateAllInfoRequest, opts ...grpc.CallOption) (*UpdateAllInfoResponse, error) {
	out := new(UpdateAllInfoResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/UpdateAllInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) FollowPublicProfile(ctx context.Context, in *FollowPublicProfileRequest, opts ...grpc.CallOption) (*FollowPublicProfileResponse, error) {
	out := new(FollowPublicProfileResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/FollowPublicProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AcceptFollowingRequest(ctx context.Context, in *AcceptFollowingRequestRequest, opts ...grpc.CallOption) (*AcceptFollowingRequestResponse, error) {
	out := new(AcceptFollowingRequestResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/AcceptFollowingRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) RejectFollowingRequest(ctx context.Context, in *RejectFollowingRequestRequest, opts ...grpc.CallOption) (*RejectFollowingRequestResponse, error) {
	out := new(RejectFollowingRequestResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/RejectFollowingRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) BlockUser(ctx context.Context, in *BlockUserRequest, opts ...grpc.CallOption) (*BlockUserResponse, error) {
	out := new(BlockUserResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/BlockUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	GetByUsername(context.Context, *GetByUsernameRequest) (*GetByUsernameResponse, error)
	GetPublicUserByUsername(context.Context, *GetPublicUserByUsernameRequest) (*GetPublicUserByUsernameResponse, error)
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
	GetAllPublicUsers(context.Context, *GetAllPublicUsersRequest) (*GetAllPublicUsersResponse, error)
	GetAllPublicUsersByUsername(context.Context, *GetAllPublicUsersByUsernameRequest) (*GetAllPublicUsersByUsernameResponse, error)
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	UpdateBasicInfo(context.Context, *UpdateBasicInfoRequest) (*UpdateBasicInfoResponse, error)
	UpdateAdvancedInfo(context.Context, *UpdateAdvancedInfoRequest) (*UpdateAdvancedInfoResponse, error)
	UpdatePersonalInfo(context.Context, *UpdatePersonalInfoRequest) (*UpdatePersonalInfoResponse, error)
	UpdateAllInfo(context.Context, *UpdateAllInfoRequest) (*UpdateAllInfoResponse, error)
	FollowPublicProfile(context.Context, *FollowPublicProfileRequest) (*FollowPublicProfileResponse, error)
	AcceptFollowingRequest(context.Context, *AcceptFollowingRequestRequest) (*AcceptFollowingRequestResponse, error)
	RejectFollowingRequest(context.Context, *RejectFollowingRequestRequest) (*RejectFollowingRequestResponse, error)
	BlockUser(context.Context, *BlockUserRequest) (*BlockUserResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedUserServiceServer) GetByUsername(context.Context, *GetByUsernameRequest) (*GetByUsernameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByUsername not implemented")
}
func (*UnimplementedUserServiceServer) GetPublicUserByUsername(context.Context, *GetPublicUserByUsernameRequest) (*GetPublicUserByUsernameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublicUserByUsername not implemented")
}
func (*UnimplementedUserServiceServer) GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (*UnimplementedUserServiceServer) GetAllPublicUsers(context.Context, *GetAllPublicUsersRequest) (*GetAllPublicUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPublicUsers not implemented")
}
func (*UnimplementedUserServiceServer) GetAllPublicUsersByUsername(context.Context, *GetAllPublicUsersByUsernameRequest) (*GetAllPublicUsersByUsernameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPublicUsersByUsername not implemented")
}
func (*UnimplementedUserServiceServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedUserServiceServer) UpdateBasicInfo(context.Context, *UpdateBasicInfoRequest) (*UpdateBasicInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBasicInfo not implemented")
}
func (*UnimplementedUserServiceServer) UpdateAdvancedInfo(context.Context, *UpdateAdvancedInfoRequest) (*UpdateAdvancedInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAdvancedInfo not implemented")
}
func (*UnimplementedUserServiceServer) UpdatePersonalInfo(context.Context, *UpdatePersonalInfoRequest) (*UpdatePersonalInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePersonalInfo not implemented")
}
func (*UnimplementedUserServiceServer) UpdateAllInfo(context.Context, *UpdateAllInfoRequest) (*UpdateAllInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAllInfo not implemented")
}
func (*UnimplementedUserServiceServer) FollowPublicProfile(context.Context, *FollowPublicProfileRequest) (*FollowPublicProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowPublicProfile not implemented")
}
func (*UnimplementedUserServiceServer) AcceptFollowingRequest(context.Context, *AcceptFollowingRequestRequest) (*AcceptFollowingRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptFollowingRequest not implemented")
}
func (*UnimplementedUserServiceServer) RejectFollowingRequest(context.Context, *RejectFollowingRequestRequest) (*RejectFollowingRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RejectFollowingRequest not implemented")
}
func (*UnimplementedUserServiceServer) BlockUser(context.Context, *BlockUserRequest) (*BlockUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockUser not implemented")
}
func (*UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetByUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetByUsername(ctx, req.(*GetByUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetPublicUserByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPublicUserByUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetPublicUserByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetPublicUserByUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetPublicUserByUsername(ctx, req.(*GetPublicUserByUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetAllPublicUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllPublicUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAllPublicUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetAllPublicUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAllPublicUsers(ctx, req.(*GetAllPublicUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetAllPublicUsersByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllPublicUsersByUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAllPublicUsersByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetAllPublicUsersByUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAllPublicUsersByUsername(ctx, req.(*GetAllPublicUsersByUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateBasicInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBasicInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateBasicInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UpdateBasicInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateBasicInfo(ctx, req.(*UpdateBasicInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateAdvancedInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAdvancedInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateAdvancedInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UpdateAdvancedInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateAdvancedInfo(ctx, req.(*UpdateAdvancedInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdatePersonalInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePersonalInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdatePersonalInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UpdatePersonalInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdatePersonalInfo(ctx, req.(*UpdatePersonalInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateAllInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAllInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateAllInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UpdateAllInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateAllInfo(ctx, req.(*UpdateAllInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_FollowPublicProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowPublicProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).FollowPublicProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/FollowPublicProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).FollowPublicProfile(ctx, req.(*FollowPublicProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AcceptFollowingRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptFollowingRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AcceptFollowingRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/AcceptFollowingRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AcceptFollowingRequest(ctx, req.(*AcceptFollowingRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_RejectFollowingRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RejectFollowingRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).RejectFollowingRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/RejectFollowingRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).RejectFollowingRequest(ctx, req.(*RejectFollowingRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_BlockUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).BlockUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/BlockUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).BlockUser(ctx, req.(*BlockUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _UserService_Get_Handler,
		},
		{
			MethodName: "GetByUsername",
			Handler:    _UserService_GetByUsername_Handler,
		},
		{
			MethodName: "GetPublicUserByUsername",
			Handler:    _UserService_GetPublicUserByUsername_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _UserService_GetAll_Handler,
		},
		{
			MethodName: "GetAllPublicUsers",
			Handler:    _UserService_GetAllPublicUsers_Handler,
		},
		{
			MethodName: "GetAllPublicUsersByUsername",
			Handler:    _UserService_GetAllPublicUsersByUsername_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _UserService_Register_Handler,
		},
		{
			MethodName: "UpdateBasicInfo",
			Handler:    _UserService_UpdateBasicInfo_Handler,
		},
		{
			MethodName: "UpdateAdvancedInfo",
			Handler:    _UserService_UpdateAdvancedInfo_Handler,
		},
		{
			MethodName: "UpdatePersonalInfo",
			Handler:    _UserService_UpdatePersonalInfo_Handler,
		},
		{
			MethodName: "UpdateAllInfo",
			Handler:    _UserService_UpdateAllInfo_Handler,
		},
		{
			MethodName: "FollowPublicProfile",
			Handler:    _UserService_FollowPublicProfile_Handler,
		},
		{
			MethodName: "AcceptFollowingRequest",
			Handler:    _UserService_AcceptFollowingRequest_Handler,
		},
		{
			MethodName: "RejectFollowingRequest",
			Handler:    _UserService_RejectFollowingRequest_Handler,
		},
		{
			MethodName: "BlockUser",
			Handler:    _UserService_BlockUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "common/proto/user_service/user_service.proto",
}
