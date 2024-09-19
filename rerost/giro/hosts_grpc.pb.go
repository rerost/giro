// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.5
// source: rerost/giro/hosts.proto

package hosts_pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	HostService_GetHost_FullMethodName = "/rerost.giro.v1.HostService/GetHost"
)

// HostServiceClient is the client API for HostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HostServiceClient interface {
	GetHost(ctx context.Context, in *GetHostRequest, opts ...grpc.CallOption) (*GetHostResponse, error)
}

type hostServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHostServiceClient(cc grpc.ClientConnInterface) HostServiceClient {
	return &hostServiceClient{cc}
}

func (c *hostServiceClient) GetHost(ctx context.Context, in *GetHostRequest, opts ...grpc.CallOption) (*GetHostResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetHostResponse)
	err := c.cc.Invoke(ctx, HostService_GetHost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HostServiceServer is the server API for HostService service.
// All implementations must embed UnimplementedHostServiceServer
// for forward compatibility.
type HostServiceServer interface {
	GetHost(context.Context, *GetHostRequest) (*GetHostResponse, error)
	mustEmbedUnimplementedHostServiceServer()
}

// UnimplementedHostServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedHostServiceServer struct{}

func (UnimplementedHostServiceServer) GetHost(context.Context, *GetHostRequest) (*GetHostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHost not implemented")
}
func (UnimplementedHostServiceServer) mustEmbedUnimplementedHostServiceServer() {}
func (UnimplementedHostServiceServer) testEmbeddedByValue()                     {}

// UnsafeHostServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HostServiceServer will
// result in compilation errors.
type UnsafeHostServiceServer interface {
	mustEmbedUnimplementedHostServiceServer()
}

func RegisterHostServiceServer(s grpc.ServiceRegistrar, srv HostServiceServer) {
	// If the following call pancis, it indicates UnimplementedHostServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&HostService_ServiceDesc, srv)
}

func _HostService_GetHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostServiceServer).GetHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HostService_GetHost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostServiceServer).GetHost(ctx, req.(*GetHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HostService_ServiceDesc is the grpc.ServiceDesc for HostService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HostService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rerost.giro.v1.HostService",
	HandlerType: (*HostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHost",
			Handler:    _HostService_GetHost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rerost/giro/hosts.proto",
}
