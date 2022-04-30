// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: followers_microservice/followers_microservice.proto

package followers_microservice

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FollowersServiceClient is the client API for FollowersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FollowersServiceClient interface {
	Follow(ctx context.Context, in *FollowRequest, opts ...grpc.CallOption) (*FollowResponse, error)
	ConfirmFollow(ctx context.Context, in *ConfirmFollowRequest, opts ...grpc.CallOption) (*ConfirmFollowResponse, error)
	GetFollowing(ctx context.Context, in *GetFollowingRequest, opts ...grpc.CallOption) (*GetFollowingResponse, error)
}

type followersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFollowersServiceClient(cc grpc.ClientConnInterface) FollowersServiceClient {
	return &followersServiceClient{cc}
}

func (c *followersServiceClient) Follow(ctx context.Context, in *FollowRequest, opts ...grpc.CallOption) (*FollowResponse, error) {
	out := new(FollowResponse)
	err := c.cc.Invoke(ctx, "/followers.FollowersService/Follow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followersServiceClient) ConfirmFollow(ctx context.Context, in *ConfirmFollowRequest, opts ...grpc.CallOption) (*ConfirmFollowResponse, error) {
	out := new(ConfirmFollowResponse)
	err := c.cc.Invoke(ctx, "/followers.FollowersService/ConfirmFollow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followersServiceClient) GetFollowing(ctx context.Context, in *GetFollowingRequest, opts ...grpc.CallOption) (*GetFollowingResponse, error) {
	out := new(GetFollowingResponse)
	err := c.cc.Invoke(ctx, "/followers.FollowersService/GetFollowing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FollowersServiceServer is the server API for FollowersService service.
// All implementations must embed UnimplementedFollowersServiceServer
// for forward compatibility
type FollowersServiceServer interface {
	Follow(context.Context, *FollowRequest) (*FollowResponse, error)
	ConfirmFollow(context.Context, *ConfirmFollowRequest) (*ConfirmFollowResponse, error)
	GetFollowing(context.Context, *GetFollowingRequest) (*GetFollowingResponse, error)
	mustEmbedUnimplementedFollowersServiceServer()
}

// UnimplementedFollowersServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFollowersServiceServer struct {
}

func (UnimplementedFollowersServiceServer) Follow(context.Context, *FollowRequest) (*FollowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Follow not implemented")
}
func (UnimplementedFollowersServiceServer) ConfirmFollow(context.Context, *ConfirmFollowRequest) (*ConfirmFollowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmFollow not implemented")
}
func (UnimplementedFollowersServiceServer) GetFollowing(context.Context, *GetFollowingRequest) (*GetFollowingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollowing not implemented")
}
func (UnimplementedFollowersServiceServer) mustEmbedUnimplementedFollowersServiceServer() {}

// UnsafeFollowersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FollowersServiceServer will
// result in compilation errors.
type UnsafeFollowersServiceServer interface {
	mustEmbedUnimplementedFollowersServiceServer()
}

func RegisterFollowersServiceServer(s grpc.ServiceRegistrar, srv FollowersServiceServer) {
	s.RegisterService(&FollowersService_ServiceDesc, srv)
}

func _FollowersService_Follow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowersServiceServer).Follow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/followers.FollowersService/Follow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowersServiceServer).Follow(ctx, req.(*FollowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FollowersService_ConfirmFollow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmFollowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowersServiceServer).ConfirmFollow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/followers.FollowersService/ConfirmFollow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowersServiceServer).ConfirmFollow(ctx, req.(*ConfirmFollowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FollowersService_GetFollowing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFollowingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowersServiceServer).GetFollowing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/followers.FollowersService/GetFollowing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowersServiceServer).GetFollowing(ctx, req.(*GetFollowingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FollowersService_ServiceDesc is the grpc.ServiceDesc for FollowersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FollowersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "followers.FollowersService",
	HandlerType: (*FollowersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Follow",
			Handler:    _FollowersService_Follow_Handler,
		},
		{
			MethodName: "ConfirmFollow",
			Handler:    _FollowersService_ConfirmFollow_Handler,
		},
		{
			MethodName: "GetFollowing",
			Handler:    _FollowersService_GetFollowing_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "followers_microservice/followers_microservice.proto",
}
