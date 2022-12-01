// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: grpc_proto/grpc.proto

package grpc_api

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

// EmpCRUDClient is the client API for EmpCRUD service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmpCRUDClient interface {
	GetEmpsDetails(ctx context.Context, in *Empty, opts ...grpc.CallOption) (EmpCRUD_GetEmpsDetailsClient, error)
	GetEmpDetails(ctx context.Context, in *Id, opts ...grpc.CallOption) (*EmpInfo, error)
	CreateEmpDetails(ctx context.Context, in *EmpInfo, opts ...grpc.CallOption) (*Id, error)
	UpdateEmpDetails(ctx context.Context, in *EmpInfo, opts ...grpc.CallOption) (*Status, error)
	DeleteEmpDetails(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Status, error)
}

type empCRUDClient struct {
	cc grpc.ClientConnInterface
}

func NewEmpCRUDClient(cc grpc.ClientConnInterface) EmpCRUDClient {
	return &empCRUDClient{cc}
}

func (c *empCRUDClient) GetEmpsDetails(ctx context.Context, in *Empty, opts ...grpc.CallOption) (EmpCRUD_GetEmpsDetailsClient, error) {
	stream, err := c.cc.NewStream(ctx, &EmpCRUD_ServiceDesc.Streams[0], "/empmgmt.EmpCRUD/GetEmpsDetails", opts...)
	if err != nil {
		return nil, err
	}
	x := &empCRUDGetEmpsDetailsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EmpCRUD_GetEmpsDetailsClient interface {
	Recv() (*EmpInfo, error)
	grpc.ClientStream
}

type empCRUDGetEmpsDetailsClient struct {
	grpc.ClientStream
}

func (x *empCRUDGetEmpsDetailsClient) Recv() (*EmpInfo, error) {
	m := new(EmpInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *empCRUDClient) GetEmpDetails(ctx context.Context, in *Id, opts ...grpc.CallOption) (*EmpInfo, error) {
	out := new(EmpInfo)
	err := c.cc.Invoke(ctx, "/empmgmt.EmpCRUD/GetEmpDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *empCRUDClient) CreateEmpDetails(ctx context.Context, in *EmpInfo, opts ...grpc.CallOption) (*Id, error) {
	out := new(Id)
	err := c.cc.Invoke(ctx, "/empmgmt.EmpCRUD/CreateEmpDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *empCRUDClient) UpdateEmpDetails(ctx context.Context, in *EmpInfo, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/empmgmt.EmpCRUD/UpdateEmpDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *empCRUDClient) DeleteEmpDetails(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/empmgmt.EmpCRUD/DeleteEmpDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmpCRUDServer is the server API for EmpCRUD service.
// All implementations must embed UnimplementedEmpCRUDServer
// for forward compatibility
type EmpCRUDServer interface {
	GetEmpsDetails(*Empty, EmpCRUD_GetEmpsDetailsServer) error
	GetEmpDetails(context.Context, *Id) (*EmpInfo, error)
	CreateEmpDetails(context.Context, *EmpInfo) (*Id, error)
	UpdateEmpDetails(context.Context, *EmpInfo) (*Status, error)
	DeleteEmpDetails(context.Context, *Id) (*Status, error)
	mustEmbedUnimplementedEmpCRUDServer()
}

// UnimplementedEmpCRUDServer must be embedded to have forward compatible implementations.
type UnimplementedEmpCRUDServer struct {
}

func (UnimplementedEmpCRUDServer) GetEmpsDetails(*Empty, EmpCRUD_GetEmpsDetailsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetEmpsDetails not implemented")
}
func (UnimplementedEmpCRUDServer) GetEmpDetails(context.Context, *Id) (*EmpInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmpDetails not implemented")
}
func (UnimplementedEmpCRUDServer) CreateEmpDetails(context.Context, *EmpInfo) (*Id, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEmpDetails not implemented")
}
func (UnimplementedEmpCRUDServer) UpdateEmpDetails(context.Context, *EmpInfo) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEmpDetails not implemented")
}
func (UnimplementedEmpCRUDServer) DeleteEmpDetails(context.Context, *Id) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEmpDetails not implemented")
}
func (UnimplementedEmpCRUDServer) mustEmbedUnimplementedEmpCRUDServer() {}

// UnsafeEmpCRUDServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmpCRUDServer will
// result in compilation errors.
type UnsafeEmpCRUDServer interface {
	mustEmbedUnimplementedEmpCRUDServer()
}

func RegisterEmpCRUDServer(s grpc.ServiceRegistrar, srv EmpCRUDServer) {
	s.RegisterService(&EmpCRUD_ServiceDesc, srv)
}

func _EmpCRUD_GetEmpsDetails_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EmpCRUDServer).GetEmpsDetails(m, &empCRUDGetEmpsDetailsServer{stream})
}

type EmpCRUD_GetEmpsDetailsServer interface {
	Send(*EmpInfo) error
	grpc.ServerStream
}

type empCRUDGetEmpsDetailsServer struct {
	grpc.ServerStream
}

func (x *empCRUDGetEmpsDetailsServer) Send(m *EmpInfo) error {
	return x.ServerStream.SendMsg(m)
}

func _EmpCRUD_GetEmpDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmpCRUDServer).GetEmpDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/empmgmt.EmpCRUD/GetEmpDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmpCRUDServer).GetEmpDetails(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmpCRUD_CreateEmpDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmpInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmpCRUDServer).CreateEmpDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/empmgmt.EmpCRUD/CreateEmpDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmpCRUDServer).CreateEmpDetails(ctx, req.(*EmpInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmpCRUD_UpdateEmpDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmpInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmpCRUDServer).UpdateEmpDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/empmgmt.EmpCRUD/UpdateEmpDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmpCRUDServer).UpdateEmpDetails(ctx, req.(*EmpInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmpCRUD_DeleteEmpDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmpCRUDServer).DeleteEmpDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/empmgmt.EmpCRUD/DeleteEmpDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmpCRUDServer).DeleteEmpDetails(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

// EmpCRUD_ServiceDesc is the grpc.ServiceDesc for EmpCRUD service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EmpCRUD_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "empmgmt.EmpCRUD",
	HandlerType: (*EmpCRUDServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEmpDetails",
			Handler:    _EmpCRUD_GetEmpDetails_Handler,
		},
		{
			MethodName: "CreateEmpDetails",
			Handler:    _EmpCRUD_CreateEmpDetails_Handler,
		},
		{
			MethodName: "UpdateEmpDetails",
			Handler:    _EmpCRUD_UpdateEmpDetails_Handler,
		},
		{
			MethodName: "DeleteEmpDetails",
			Handler:    _EmpCRUD_DeleteEmpDetails_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetEmpsDetails",
			Handler:       _EmpCRUD_GetEmpsDetails_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "grpc_proto/grpc.proto",
}
