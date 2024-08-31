//Generated by gRPC Go plugin
//If you make any local changes, they will be lost
//source: user

package myapp

import (
	context "context"
	flatbuffers "github.com/google/flatbuffers/go"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Client API for UserService service
type UserServiceClient interface {
	GetUser(ctx context.Context, in *flatbuffers.Builder,
		opts ...grpc.CallOption) (*UserResponse, error)
	GetAllUsers(ctx context.Context, in *flatbuffers.Builder,
		opts ...grpc.CallOption) (*UsersListResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUser(ctx context.Context, in *flatbuffers.Builder,
	opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/myapp.UserService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAllUsers(ctx context.Context, in *flatbuffers.Builder,
	opts ...grpc.CallOption) (*UsersListResponse, error) {
	out := new(UsersListResponse)
	err := c.cc.Invoke(ctx, "/myapp.UserService/GetAllUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service
type UserServiceServer interface {
	GetUser(context.Context, *UserRequest) (*flatbuffers.Builder, error)
	GetAllUsers(context.Context, *GetAllUsersRequest) (*flatbuffers.Builder, error)
	mustEmbedUnimplementedUserServiceServer()
}

type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) GetUser(context.Context, *UserRequest) (*flatbuffers.Builder, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}

func (UnimplementedUserServiceServer) GetAllUsers(context.Context, *GetAllUsersRequest) (*flatbuffers.Builder, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllUsers not implemented")
}

func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context,
	dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/myapp.UserService/GetUser",
	}

	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _UserService_GetAllUsers_Handler(srv interface{}, ctx context.Context,
	dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAllUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/myapp.UserService/GetAllUsers",
	}

	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAllUsers(ctx, req.(*GetAllUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "myapp.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "GetAllUsers",
			Handler:    _UserService_GetAllUsers_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
