// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protobuf

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// RouteChatClient is the client API for RouteChat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RouteChatClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*GeneralResponse, error)
	ReadLatest(ctx context.Context, in *ReadLatestRequest, opts ...grpc.CallOption) (*ReadGeneralResponse, error)
	ReadMore(ctx context.Context, in *ReadMoreRequest, opts ...grpc.CallOption) (*ReadGeneralResponse, error)
	SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*GeneralResponse, error)
	RemoveMessage(ctx context.Context, in *RemoveMessageRequest, opts ...grpc.CallOption) (*GeneralResponse, error)
	UpdateMessage(ctx context.Context, in *UpdateMessageRequest, opts ...grpc.CallOption) (*GeneralResponse, error)
	Stream(ctx context.Context, opts ...grpc.CallOption) (RouteChat_StreamClient, error)
}

type routeChatClient struct {
	cc grpc.ClientConnInterface
}

func NewRouteChatClient(cc grpc.ClientConnInterface) RouteChatClient {
	return &routeChatClient{cc}
}

func (c *routeChatClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/chat.RouteChat/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeChatClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*GeneralResponse, error) {
	out := new(GeneralResponse)
	err := c.cc.Invoke(ctx, "/chat.RouteChat/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeChatClient) ReadLatest(ctx context.Context, in *ReadLatestRequest, opts ...grpc.CallOption) (*ReadGeneralResponse, error) {
	out := new(ReadGeneralResponse)
	err := c.cc.Invoke(ctx, "/chat.RouteChat/ReadLatest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeChatClient) ReadMore(ctx context.Context, in *ReadMoreRequest, opts ...grpc.CallOption) (*ReadGeneralResponse, error) {
	out := new(ReadGeneralResponse)
	err := c.cc.Invoke(ctx, "/chat.RouteChat/ReadMore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeChatClient) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*GeneralResponse, error) {
	out := new(GeneralResponse)
	err := c.cc.Invoke(ctx, "/chat.RouteChat/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeChatClient) RemoveMessage(ctx context.Context, in *RemoveMessageRequest, opts ...grpc.CallOption) (*GeneralResponse, error) {
	out := new(GeneralResponse)
	err := c.cc.Invoke(ctx, "/chat.RouteChat/RemoveMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeChatClient) UpdateMessage(ctx context.Context, in *UpdateMessageRequest, opts ...grpc.CallOption) (*GeneralResponse, error) {
	out := new(GeneralResponse)
	err := c.cc.Invoke(ctx, "/chat.RouteChat/UpdateMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeChatClient) Stream(ctx context.Context, opts ...grpc.CallOption) (RouteChat_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RouteChat_serviceDesc.Streams[0], "/chat.RouteChat/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeChatStreamClient{stream}
	return x, nil
}

type RouteChat_StreamClient interface {
	Send(*StreamRequest) error
	Recv() (*StreamResponse, error)
	grpc.ClientStream
}

type routeChatStreamClient struct {
	grpc.ClientStream
}

func (x *routeChatStreamClient) Send(m *StreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeChatStreamClient) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RouteChatServer is the server API for RouteChat service.
// All implementations must embed UnimplementedRouteChatServer
// for forward compatibility
type RouteChatServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Logout(context.Context, *LogoutRequest) (*GeneralResponse, error)
	ReadLatest(context.Context, *ReadLatestRequest) (*ReadGeneralResponse, error)
	ReadMore(context.Context, *ReadMoreRequest) (*ReadGeneralResponse, error)
	SendMessage(context.Context, *SendMessageRequest) (*GeneralResponse, error)
	RemoveMessage(context.Context, *RemoveMessageRequest) (*GeneralResponse, error)
	UpdateMessage(context.Context, *UpdateMessageRequest) (*GeneralResponse, error)
	Stream(RouteChat_StreamServer) error
	mustEmbedUnimplementedRouteChatServer()
}

// UnimplementedRouteChatServer must be embedded to have forward compatible implementations.
type UnimplementedRouteChatServer struct {
}

func (UnimplementedRouteChatServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedRouteChatServer) Logout(context.Context, *LogoutRequest) (*GeneralResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedRouteChatServer) ReadLatest(context.Context, *ReadLatestRequest) (*ReadGeneralResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadLatest not implemented")
}
func (UnimplementedRouteChatServer) ReadMore(context.Context, *ReadMoreRequest) (*ReadGeneralResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadMore not implemented")
}
func (UnimplementedRouteChatServer) SendMessage(context.Context, *SendMessageRequest) (*GeneralResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedRouteChatServer) RemoveMessage(context.Context, *RemoveMessageRequest) (*GeneralResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveMessage not implemented")
}
func (UnimplementedRouteChatServer) UpdateMessage(context.Context, *UpdateMessageRequest) (*GeneralResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMessage not implemented")
}
func (UnimplementedRouteChatServer) Stream(RouteChat_StreamServer) error {
	return status.Errorf(codes.Unimplemented, "method Stream not implemented")
}
func (UnimplementedRouteChatServer) mustEmbedUnimplementedRouteChatServer() {}

// UnsafeRouteChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RouteChatServer will
// result in compilation errors.
type UnsafeRouteChatServer interface {
	mustEmbedUnimplementedRouteChatServer()
}

func RegisterRouteChatServer(s grpc.ServiceRegistrar, srv RouteChatServer) {
	s.RegisterService(&_RouteChat_serviceDesc, srv)
}

func _RouteChat_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteChatServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.RouteChat/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteChatServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteChat_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteChatServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.RouteChat/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteChatServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteChat_ReadLatest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadLatestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteChatServer).ReadLatest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.RouteChat/ReadLatest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteChatServer).ReadLatest(ctx, req.(*ReadLatestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteChat_ReadMore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadMoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteChatServer).ReadMore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.RouteChat/ReadMore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteChatServer).ReadMore(ctx, req.(*ReadMoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteChat_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteChatServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.RouteChat/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteChatServer).SendMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteChat_RemoveMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteChatServer).RemoveMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.RouteChat/RemoveMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteChatServer).RemoveMessage(ctx, req.(*RemoveMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteChat_UpdateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteChatServer).UpdateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.RouteChat/UpdateMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteChatServer).UpdateMessage(ctx, req.(*UpdateMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RouteChat_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteChatServer).Stream(&routeChatStreamServer{stream})
}

type RouteChat_StreamServer interface {
	Send(*StreamResponse) error
	Recv() (*StreamRequest, error)
	grpc.ServerStream
}

type routeChatStreamServer struct {
	grpc.ServerStream
}

func (x *routeChatStreamServer) Send(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeChatStreamServer) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _RouteChat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chat.RouteChat",
	HandlerType: (*RouteChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _RouteChat_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _RouteChat_Logout_Handler,
		},
		{
			MethodName: "ReadLatest",
			Handler:    _RouteChat_ReadLatest_Handler,
		},
		{
			MethodName: "ReadMore",
			Handler:    _RouteChat_ReadMore_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _RouteChat_SendMessage_Handler,
		},
		{
			MethodName: "RemoveMessage",
			Handler:    _RouteChat_RemoveMessage_Handler,
		},
		{
			MethodName: "UpdateMessage",
			Handler:    _RouteChat_UpdateMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _RouteChat_Stream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "protobuf/gogoatalk.proto",
}
