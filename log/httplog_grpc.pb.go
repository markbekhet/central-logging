// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: httplog.proto

package log

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

const (
	Logger_Send_FullMethodName = "/Logger/Send"
)

// LoggerClient is the client API for Logger service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoggerClient interface {
	Send(ctx context.Context, in *HTTPLog, opts ...grpc.CallOption) (*LogRes, error)
}

type loggerClient struct {
	cc grpc.ClientConnInterface
}

func NewLoggerClient(cc grpc.ClientConnInterface) LoggerClient {
	return &loggerClient{cc}
}

func (c *loggerClient) Send(ctx context.Context, in *HTTPLog, opts ...grpc.CallOption) (*LogRes, error) {
	out := new(LogRes)
	err := c.cc.Invoke(ctx, Logger_Send_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoggerServer is the server API for Logger service.
// All implementations must embed UnimplementedLoggerServer
// for forward compatibility
type LoggerServer interface {
	Send(context.Context, *HTTPLog) (*LogRes, error)
	mustEmbedUnimplementedLoggerServer()
}

// UnimplementedLoggerServer must be embedded to have forward compatible implementations.
type UnimplementedLoggerServer struct {
}

func (UnimplementedLoggerServer) Send(context.Context, *HTTPLog) (*LogRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedLoggerServer) mustEmbedUnimplementedLoggerServer() {}

// UnsafeLoggerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoggerServer will
// result in compilation errors.
type UnsafeLoggerServer interface {
	mustEmbedUnimplementedLoggerServer()
}

func RegisterLoggerServer(s grpc.ServiceRegistrar, srv LoggerServer) {
	s.RegisterService(&Logger_ServiceDesc, srv)
}

func _Logger_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HTTPLog)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggerServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Logger_Send_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggerServer).Send(ctx, req.(*HTTPLog))
	}
	return interceptor(ctx, in, info, handler)
}

// Logger_ServiceDesc is the grpc.ServiceDesc for Logger service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Logger_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Logger",
	HandlerType: (*LoggerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _Logger_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "httplog.proto",
}
