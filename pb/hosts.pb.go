// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: rerost/giro/hosts.proto

package hosts_pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type HostOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host string `protobuf:"bytes,10,opt,name=host,proto3" json:"host,omitempty"`
}

func (x *HostOptions) Reset() {
	*x = HostOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rerost_giro_hosts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostOptions) ProtoMessage() {}

func (x *HostOptions) ProtoReflect() protoreflect.Message {
	mi := &file_rerost_giro_hosts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostOptions.ProtoReflect.Descriptor instead.
func (*HostOptions) Descriptor() ([]byte, []int) {
	return file_rerost_giro_hosts_proto_rawDescGZIP(), []int{0}
}

func (x *HostOptions) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

type ListHostsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
}

func (x *ListHostsRequest) Reset() {
	*x = ListHostsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rerost_giro_hosts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListHostsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHostsRequest) ProtoMessage() {}

func (x *ListHostsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rerost_giro_hosts_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHostsRequest.ProtoReflect.Descriptor instead.
func (*ListHostsRequest) Descriptor() ([]byte, []int) {
	return file_rerost_giro_hosts_proto_rawDescGZIP(), []int{1}
}

func (x *ListHostsRequest) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

type ListHostsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
}

func (x *ListHostsResponse) Reset() {
	*x = ListHostsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rerost_giro_hosts_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListHostsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHostsResponse) ProtoMessage() {}

func (x *ListHostsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rerost_giro_hosts_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHostsResponse.ProtoReflect.Descriptor instead.
func (*ListHostsResponse) Descriptor() ([]byte, []int) {
	return file_rerost_giro_hosts_proto_rawDescGZIP(), []int{2}
}

func (x *ListHostsResponse) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

var file_rerost_giro_hosts_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.ServiceOptions)(nil),
		ExtensionType: (*HostOptions)(nil),
		Field:         50000,
		Name:          "rerost.giro.v1.host_option",
		Tag:           "bytes,50000,opt,name=host_option",
		Filename:      "rerost/giro/hosts.proto",
	},
}

// Extension fields to descriptor.ServiceOptions.
var (
	// optional rerost.giro.v1.HostOptions host_option = 50000;
	E_HostOption = &file_rerost_giro_hosts_proto_extTypes[0]
)

var File_rerost_giro_hosts_proto protoreflect.FileDescriptor

var file_rerost_giro_hosts_proto_rawDesc = []byte{
	0x0a, 0x17, 0x72, 0x65, 0x72, 0x6f, 0x73, 0x74, 0x2f, 0x67, 0x69, 0x72, 0x6f, 0x2f, 0x68, 0x6f,
	0x73, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x72, 0x65, 0x72, 0x6f, 0x73,
	0x74, 0x2e, 0x67, 0x69, 0x72, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21, 0x0a, 0x0b, 0x48,
	0x6f, 0x73, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f,
	0x73, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x22, 0x35,
	0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x27, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x6f, 0x73,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x32, 0x61,
	0x0a, 0x0b, 0x48, 0x6f, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x52, 0x0a,
	0x09, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x20, 0x2e, 0x72, 0x65, 0x72,
	0x6f, 0x73, 0x74, 0x2e, 0x67, 0x69, 0x72, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x48, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x72,
	0x65, 0x72, 0x6f, 0x73, 0x74, 0x2e, 0x67, 0x69, 0x72, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x3a, 0x5f, 0x0a, 0x0b, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0xd0, 0x86, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x72, 0x65, 0x72, 0x6f,
	0x73, 0x74, 0x2e, 0x67, 0x69, 0x72, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0a, 0x68, 0x6f, 0x73, 0x74, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x72, 0x65, 0x72, 0x6f, 0x73, 0x74, 0x2f, 0x67, 0x69, 0x72, 0x6f, 0x2f, 0x70, 0x62, 0x3b,
	0x68, 0x6f, 0x73, 0x74, 0x73, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rerost_giro_hosts_proto_rawDescOnce sync.Once
	file_rerost_giro_hosts_proto_rawDescData = file_rerost_giro_hosts_proto_rawDesc
)

func file_rerost_giro_hosts_proto_rawDescGZIP() []byte {
	file_rerost_giro_hosts_proto_rawDescOnce.Do(func() {
		file_rerost_giro_hosts_proto_rawDescData = protoimpl.X.CompressGZIP(file_rerost_giro_hosts_proto_rawDescData)
	})
	return file_rerost_giro_hosts_proto_rawDescData
}

var file_rerost_giro_hosts_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_rerost_giro_hosts_proto_goTypes = []interface{}{
	(*HostOptions)(nil),               // 0: rerost.giro.v1.HostOptions
	(*ListHostsRequest)(nil),          // 1: rerost.giro.v1.ListHostsRequest
	(*ListHostsResponse)(nil),         // 2: rerost.giro.v1.ListHostsResponse
	(*descriptor.ServiceOptions)(nil), // 3: google.protobuf.ServiceOptions
}
var file_rerost_giro_hosts_proto_depIdxs = []int32{
	3, // 0: rerost.giro.v1.host_option:extendee -> google.protobuf.ServiceOptions
	0, // 1: rerost.giro.v1.host_option:type_name -> rerost.giro.v1.HostOptions
	1, // 2: rerost.giro.v1.HostService.ListHosts:input_type -> rerost.giro.v1.ListHostsRequest
	2, // 3: rerost.giro.v1.HostService.ListHosts:output_type -> rerost.giro.v1.ListHostsResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	1, // [1:2] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rerost_giro_hosts_proto_init() }
func file_rerost_giro_hosts_proto_init() {
	if File_rerost_giro_hosts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rerost_giro_hosts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rerost_giro_hosts_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListHostsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rerost_giro_hosts_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListHostsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rerost_giro_hosts_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 1,
			NumServices:   1,
		},
		GoTypes:           file_rerost_giro_hosts_proto_goTypes,
		DependencyIndexes: file_rerost_giro_hosts_proto_depIdxs,
		MessageInfos:      file_rerost_giro_hosts_proto_msgTypes,
		ExtensionInfos:    file_rerost_giro_hosts_proto_extTypes,
	}.Build()
	File_rerost_giro_hosts_proto = out.File
	file_rerost_giro_hosts_proto_rawDesc = nil
	file_rerost_giro_hosts_proto_goTypes = nil
	file_rerost_giro_hosts_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HostServiceClient is the client API for HostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HostServiceClient interface {
	ListHosts(ctx context.Context, in *ListHostsRequest, opts ...grpc.CallOption) (*ListHostsResponse, error)
}

type hostServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHostServiceClient(cc grpc.ClientConnInterface) HostServiceClient {
	return &hostServiceClient{cc}
}

func (c *hostServiceClient) ListHosts(ctx context.Context, in *ListHostsRequest, opts ...grpc.CallOption) (*ListHostsResponse, error) {
	out := new(ListHostsResponse)
	err := c.cc.Invoke(ctx, "/rerost.giro.v1.HostService/ListHosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HostServiceServer is the server API for HostService service.
type HostServiceServer interface {
	ListHosts(context.Context, *ListHostsRequest) (*ListHostsResponse, error)
}

// UnimplementedHostServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHostServiceServer struct {
}

func (*UnimplementedHostServiceServer) ListHosts(context.Context, *ListHostsRequest) (*ListHostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHosts not implemented")
}

func RegisterHostServiceServer(s *grpc.Server, srv HostServiceServer) {
	s.RegisterService(&_HostService_serviceDesc, srv)
}

func _HostService_ListHosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListHostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostServiceServer).ListHosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rerost.giro.v1.HostService/ListHosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostServiceServer).ListHosts(ctx, req.(*ListHostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HostService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rerost.giro.v1.HostService",
	HandlerType: (*HostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListHosts",
			Handler:    _HostService_ListHosts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rerost/giro/hosts.proto",
}
