// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: pkg/auth_svc/infrastructure/pb/auth.proto

package pb

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

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	UserSignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error)
	UserOTPVerication(ctx context.Context, in *RequestOtpVefification, opts ...grpc.CallOption) (*ResponseOtpVerification, error)
	UserLogin(ctx context.Context, in *RequestUserLogin, opts ...grpc.CallOption) (*ResponseUserLogin, error)
	ForgotPasswordRequest(ctx context.Context, in *RequestForgotPass, opts ...grpc.CallOption) (*ResponseForgotPass, error)
	ResetPassword(ctx context.Context, in *RequestResetPass, opts ...grpc.CallOption) (*ResponseErrorMessage, error)
	VerifyAccessToken(ctx context.Context, in *RequestVerifyAccess, opts ...grpc.CallOption) (*ResponseVerifyAccess, error)
	AccessRegenerator(ctx context.Context, in *RequestAccessGenerator, opts ...grpc.CallOption) (*ResponseAccessGenerator, error)
	GetUserProfile(ctx context.Context, in *RequestGetUserProfile, opts ...grpc.CallOption) (*ResponseUserProfile, error)
	EditUserProfile(ctx context.Context, in *RequestEditUserProfile, opts ...grpc.CallOption) (*ResponseErrorMessage, error)
	GetFollowersDetails(ctx context.Context, in *RequestUserId, opts ...grpc.CallOption) (*ResponseGetUsersDetails, error)
	GetFollowingsDetails(ctx context.Context, in *RequestUserId, opts ...grpc.CallOption) (*ResponseGetUsersDetails, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) UserSignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error) {
	out := new(SignUpResponse)
	err := c.cc.Invoke(ctx, "/auth_proto.AuthService/UserSignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UserOTPVerication(ctx context.Context, in *RequestOtpVefification, opts ...grpc.CallOption) (*ResponseOtpVerification, error) {
	out := new(ResponseOtpVerification)
	err := c.cc.Invoke(ctx, "/auth_proto.AuthService/UserOTPVerication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) UserLogin(ctx context.Context, in *RequestUserLogin, opts ...grpc.CallOption) (*ResponseUserLogin, error) {
	out := new(ResponseUserLogin)
	err := c.cc.Invoke(ctx, "/auth_proto.AuthService/UserLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ForgotPasswordRequest(ctx context.Context, in *RequestForgotPass, opts ...grpc.CallOption) (*ResponseForgotPass, error) {
	out := new(ResponseForgotPass)
	err := c.cc.Invoke(ctx, "/auth_proto.AuthService/ForgotPasswordRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ResetPassword(ctx context.Context, in *RequestResetPass, opts ...grpc.CallOption) (*ResponseErrorMessage, error) {
	out := new(ResponseErrorMessage)
	err := c.cc.Invoke(ctx, "/auth_proto.AuthService/ResetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) VerifyAccessToken(ctx context.Context, in *RequestVerifyAccess, opts ...grpc.CallOption) (*ResponseVerifyAccess, error) {
	out := new(ResponseVerifyAccess)
	err := c.cc.Invoke(ctx, "/auth_proto.AuthService/VerifyAccessToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) AccessRegenerator(ctx context.Context, in *RequestAccessGenerator, opts ...grpc.CallOption) (*ResponseAccessGenerator, error) {
	out := new(ResponseAccessGenerator)
	err := c.cc.Invoke(ctx, "/auth_proto.AuthService/AccessRegenerator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetUserProfile(ctx context.Context, in *RequestGetUserProfile, opts ...grpc.CallOption) (*ResponseUserProfile, error) {
	out := new(ResponseUserProfile)
	err := c.cc.Invoke(ctx, "/auth_proto.AuthService/GetUserProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) EditUserProfile(ctx context.Context, in *RequestEditUserProfile, opts ...grpc.CallOption) (*ResponseErrorMessage, error) {
	out := new(ResponseErrorMessage)
	err := c.cc.Invoke(ctx, "/auth_proto.AuthService/EditUserProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetFollowersDetails(ctx context.Context, in *RequestUserId, opts ...grpc.CallOption) (*ResponseGetUsersDetails, error) {
	out := new(ResponseGetUsersDetails)
	err := c.cc.Invoke(ctx, "/auth_proto.AuthService/GetFollowersDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetFollowingsDetails(ctx context.Context, in *RequestUserId, opts ...grpc.CallOption) (*ResponseGetUsersDetails, error) {
	out := new(ResponseGetUsersDetails)
	err := c.cc.Invoke(ctx, "/auth_proto.AuthService/GetFollowingsDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	UserSignUp(context.Context, *SignUpRequest) (*SignUpResponse, error)
	UserOTPVerication(context.Context, *RequestOtpVefification) (*ResponseOtpVerification, error)
	UserLogin(context.Context, *RequestUserLogin) (*ResponseUserLogin, error)
	ForgotPasswordRequest(context.Context, *RequestForgotPass) (*ResponseForgotPass, error)
	ResetPassword(context.Context, *RequestResetPass) (*ResponseErrorMessage, error)
	VerifyAccessToken(context.Context, *RequestVerifyAccess) (*ResponseVerifyAccess, error)
	AccessRegenerator(context.Context, *RequestAccessGenerator) (*ResponseAccessGenerator, error)
	GetUserProfile(context.Context, *RequestGetUserProfile) (*ResponseUserProfile, error)
	EditUserProfile(context.Context, *RequestEditUserProfile) (*ResponseErrorMessage, error)
	GetFollowersDetails(context.Context, *RequestUserId) (*ResponseGetUsersDetails, error)
	GetFollowingsDetails(context.Context, *RequestUserId) (*ResponseGetUsersDetails, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) UserSignUp(context.Context, *SignUpRequest) (*SignUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserSignUp not implemented")
}
func (UnimplementedAuthServiceServer) UserOTPVerication(context.Context, *RequestOtpVefification) (*ResponseOtpVerification, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserOTPVerication not implemented")
}
func (UnimplementedAuthServiceServer) UserLogin(context.Context, *RequestUserLogin) (*ResponseUserLogin, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (UnimplementedAuthServiceServer) ForgotPasswordRequest(context.Context, *RequestForgotPass) (*ResponseForgotPass, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForgotPasswordRequest not implemented")
}
func (UnimplementedAuthServiceServer) ResetPassword(context.Context, *RequestResetPass) (*ResponseErrorMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (UnimplementedAuthServiceServer) VerifyAccessToken(context.Context, *RequestVerifyAccess) (*ResponseVerifyAccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyAccessToken not implemented")
}
func (UnimplementedAuthServiceServer) AccessRegenerator(context.Context, *RequestAccessGenerator) (*ResponseAccessGenerator, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccessRegenerator not implemented")
}
func (UnimplementedAuthServiceServer) GetUserProfile(context.Context, *RequestGetUserProfile) (*ResponseUserProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProfile not implemented")
}
func (UnimplementedAuthServiceServer) EditUserProfile(context.Context, *RequestEditUserProfile) (*ResponseErrorMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditUserProfile not implemented")
}
func (UnimplementedAuthServiceServer) GetFollowersDetails(context.Context, *RequestUserId) (*ResponseGetUsersDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollowersDetails not implemented")
}
func (UnimplementedAuthServiceServer) GetFollowingsDetails(context.Context, *RequestUserId) (*ResponseGetUsersDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollowingsDetails not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_UserSignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UserSignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_proto.AuthService/UserSignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UserSignUp(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UserOTPVerication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestOtpVefification)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UserOTPVerication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_proto.AuthService/UserOTPVerication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UserOTPVerication(ctx, req.(*RequestOtpVefification))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestUserLogin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_proto.AuthService/UserLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).UserLogin(ctx, req.(*RequestUserLogin))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ForgotPasswordRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestForgotPass)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ForgotPasswordRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_proto.AuthService/ForgotPasswordRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ForgotPasswordRequest(ctx, req.(*RequestForgotPass))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestResetPass)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_proto.AuthService/ResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ResetPassword(ctx, req.(*RequestResetPass))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_VerifyAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestVerifyAccess)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).VerifyAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_proto.AuthService/VerifyAccessToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).VerifyAccessToken(ctx, req.(*RequestVerifyAccess))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_AccessRegenerator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestAccessGenerator)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AccessRegenerator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_proto.AuthService/AccessRegenerator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AccessRegenerator(ctx, req.(*RequestAccessGenerator))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestGetUserProfile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_proto.AuthService/GetUserProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetUserProfile(ctx, req.(*RequestGetUserProfile))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_EditUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestEditUserProfile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).EditUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_proto.AuthService/EditUserProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).EditUserProfile(ctx, req.(*RequestEditUserProfile))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetFollowersDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestUserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetFollowersDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_proto.AuthService/GetFollowersDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetFollowersDetails(ctx, req.(*RequestUserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetFollowingsDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestUserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetFollowingsDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_proto.AuthService/GetFollowingsDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetFollowingsDetails(ctx, req.(*RequestUserId))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth_proto.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserSignUp",
			Handler:    _AuthService_UserSignUp_Handler,
		},
		{
			MethodName: "UserOTPVerication",
			Handler:    _AuthService_UserOTPVerication_Handler,
		},
		{
			MethodName: "UserLogin",
			Handler:    _AuthService_UserLogin_Handler,
		},
		{
			MethodName: "ForgotPasswordRequest",
			Handler:    _AuthService_ForgotPasswordRequest_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _AuthService_ResetPassword_Handler,
		},
		{
			MethodName: "VerifyAccessToken",
			Handler:    _AuthService_VerifyAccessToken_Handler,
		},
		{
			MethodName: "AccessRegenerator",
			Handler:    _AuthService_AccessRegenerator_Handler,
		},
		{
			MethodName: "GetUserProfile",
			Handler:    _AuthService_GetUserProfile_Handler,
		},
		{
			MethodName: "EditUserProfile",
			Handler:    _AuthService_EditUserProfile_Handler,
		},
		{
			MethodName: "GetFollowersDetails",
			Handler:    _AuthService_GetFollowersDetails_Handler,
		},
		{
			MethodName: "GetFollowingsDetails",
			Handler:    _AuthService_GetFollowingsDetails_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/auth_svc/infrastructure/pb/auth.proto",
}
