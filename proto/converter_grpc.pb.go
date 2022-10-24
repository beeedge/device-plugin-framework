// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: proto/converter.proto

package proto

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

// ConverterClient is the client API for Converter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConverterClient interface {
	ConvertReportMessage2Devices(ctx context.Context, in *GetDeviceReportRequest, opts ...grpc.CallOption) (*GetDeviceReportResponse, error)
	ConvertIssueMessage2Device(ctx context.Context, in *GetDeviceIssueRequest, opts ...grpc.CallOption) (*GetDeviceIssueResponse, error)
	ConvertDeviceMessages2MQFormat(ctx context.Context, in *GetMQFormatRequest, opts ...grpc.CallOption) (*GetMQFormatResponse, error)
}

type converterClient struct {
	cc grpc.ClientConnInterface
}

func NewConverterClient(cc grpc.ClientConnInterface) ConverterClient {
	return &converterClient{cc}
}

func (c *converterClient) ConvertReportMessage2Devices(ctx context.Context, in *GetDeviceReportRequest, opts ...grpc.CallOption) (*GetDeviceReportResponse, error) {
	out := new(GetDeviceReportResponse)
	err := c.cc.Invoke(ctx, "/proto.Converter/ConvertReportMessage2Devices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *converterClient) ConvertIssueMessage2Device(ctx context.Context, in *GetDeviceIssueRequest, opts ...grpc.CallOption) (*GetDeviceIssueResponse, error) {
	out := new(GetDeviceIssueResponse)
	err := c.cc.Invoke(ctx, "/proto.Converter/ConvertIssueMessage2Device", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *converterClient) ConvertDeviceMessages2MQFormat(ctx context.Context, in *GetMQFormatRequest, opts ...grpc.CallOption) (*GetMQFormatResponse, error) {
	out := new(GetMQFormatResponse)
	err := c.cc.Invoke(ctx, "/proto.Converter/ConvertDeviceMessages2MQFormat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConverterServer is the server API for Converter service.
// All implementations should embed UnimplementedConverterServer
// for forward compatibility
type ConverterServer interface {
	ConvertReportMessage2Devices(context.Context, *GetDeviceReportRequest) (*GetDeviceReportResponse, error)
	ConvertIssueMessage2Device(context.Context, *GetDeviceIssueRequest) (*GetDeviceIssueResponse, error)
	ConvertDeviceMessages2MQFormat(context.Context, *GetMQFormatRequest) (*GetMQFormatResponse, error)
}

// UnimplementedConverterServer should be embedded to have forward compatible implementations.
type UnimplementedConverterServer struct {
}

func (UnimplementedConverterServer) ConvertReportMessage2Devices(context.Context, *GetDeviceReportRequest) (*GetDeviceReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConvertReportMessage2Devices not implemented")
}
func (UnimplementedConverterServer) ConvertIssueMessage2Device(context.Context, *GetDeviceIssueRequest) (*GetDeviceIssueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConvertIssueMessage2Device not implemented")
}
func (UnimplementedConverterServer) ConvertDeviceMessages2MQFormat(context.Context, *GetMQFormatRequest) (*GetMQFormatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConvertDeviceMessages2MQFormat not implemented")
}

// UnsafeConverterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConverterServer will
// result in compilation errors.
type UnsafeConverterServer interface {
	mustEmbedUnimplementedConverterServer()
}

func RegisterConverterServer(s grpc.ServiceRegistrar, srv ConverterServer) {
	s.RegisterService(&Converter_ServiceDesc, srv)
}

func _Converter_ConvertReportMessage2Devices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConverterServer).ConvertReportMessage2Devices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Converter/ConvertReportMessage2Devices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConverterServer).ConvertReportMessage2Devices(ctx, req.(*GetDeviceReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Converter_ConvertIssueMessage2Device_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceIssueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConverterServer).ConvertIssueMessage2Device(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Converter/ConvertIssueMessage2Device",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConverterServer).ConvertIssueMessage2Device(ctx, req.(*GetDeviceIssueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Converter_ConvertDeviceMessages2MQFormat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMQFormatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConverterServer).ConvertDeviceMessages2MQFormat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Converter/ConvertDeviceMessages2MQFormat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConverterServer).ConvertDeviceMessages2MQFormat(ctx, req.(*GetMQFormatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Converter_ServiceDesc is the grpc.ServiceDesc for Converter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Converter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Converter",
	HandlerType: (*ConverterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConvertReportMessage2Devices",
			Handler:    _Converter_ConvertReportMessage2Devices_Handler,
		},
		{
			MethodName: "ConvertIssueMessage2Device",
			Handler:    _Converter_ConvertIssueMessage2Device_Handler,
		},
		{
			MethodName: "ConvertDeviceMessages2MQFormat",
			Handler:    _Converter_ConvertDeviceMessages2MQFormat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/converter.proto",
}
