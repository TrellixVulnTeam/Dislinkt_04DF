// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package authentification

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthentificationServiceClient is the client API for AuthentificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthentificationServiceClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	GetByUserId(ctx context.Context, in *GetByUserIdRequest, opts ...grpc.CallOption) (*GetByUserIdResponse, error)
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
}

type authentificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthentificationServiceClient(cc grpc.ClientConnInterface) AuthentificationServiceClient {
	return &authentificationServiceClient{cc}
}

func (c *authentificationServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/authentification.AuthentificationService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authentificationServiceClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/authentification.AuthentificationService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authentificationServiceClient) GetByUserId(ctx context.Context, in *GetByUserIdRequest, opts ...grpc.CallOption) (*GetByUserIdResponse, error) {
	out := new(GetByUserIdResponse)
	err := c.cc.Invoke(ctx, "/authentification.AuthentificationService/GetByUserId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authentificationServiceClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/authentification.AuthentificationService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthentificationServiceServer is the server API for AuthentificationService service.
// All implementations must embed UnimplementedAuthentificationServiceServer
// for forward compatibility
type AuthentificationServiceServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
	GetByUserId(context.Context, *GetByUserIdRequest) (*GetByUserIdResponse, error)
	Add(context.Context, *AddRequest) (*AddResponse, error)
	mustEmbedUnimplementedAuthentificationServiceServer()
}

// UnimplementedAuthentificationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthentificationServiceServer struct {
}

func (*UnimplementedAuthentificationServiceServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedAuthentificationServiceServer) GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (*UnimplementedAuthentificationServiceServer) GetByUserId(context.Context, *GetByUserIdRequest) (*GetByUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByUserId not implemented")
}
func (*UnimplementedAuthentificationServiceServer) Add(context.Context, *AddRequest) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedAuthentificationServiceServer) mustEmbedUnimplementedAuthentificationServiceServer() {
}

func RegisterAuthentificationServiceServer(s *grpc.Server, srv AuthentificationServiceServer) {
	s.RegisterService(&_AuthentificationService_serviceDesc, srv)
}

func _AuthentificationService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthentificationServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentification.AuthentificationService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthentificationServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthentificationService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthentificationServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentification.AuthentificationService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthentificationServiceServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthentificationService_GetByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthentificationServiceServer).GetByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentification.AuthentificationService/GetByUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthentificationServiceServer).GetByUserId(ctx, req.(*GetByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthentificationService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthentificationServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentification.AuthentificationService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthentificationServiceServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthentificationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "authentification.AuthentificationService",
	HandlerType: (*AuthentificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _AuthentificationService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _AuthentificationService_GetAll_Handler,
		},
		{
			MethodName: "GetByUserId",
			Handler:    _AuthentificationService_GetByUserId_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _AuthentificationService_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "common/proto/authentification_service/authentification_service.proto",
}
