// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v5.27.3
// source: e2etest/dummyserver/echo.proto

package dummyserver

import (
	_ "github.com/rerost/giro/pb/hosts"
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

type MetadataValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []string `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
}

func (x *MetadataValue) Reset() {
	*x = MetadataValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_e2etest_dummyserver_echo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetadataValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataValue) ProtoMessage() {}

func (x *MetadataValue) ProtoReflect() protoreflect.Message {
	mi := &file_e2etest_dummyserver_echo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataValue.ProtoReflect.Descriptor instead.
func (*MetadataValue) Descriptor() ([]byte, []int) {
	return file_e2etest_dummyserver_echo_proto_rawDescGZIP(), []int{0}
}

func (x *MetadataValue) GetValue() []string {
	if x != nil {
		return x.Value
	}
	return nil
}

type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata map[string]*MetadataValue `protobuf:"bytes,1,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_e2etest_dummyserver_echo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_e2etest_dummyserver_echo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_e2etest_dummyserver_echo_proto_rawDescGZIP(), []int{1}
}

func (x *Metadata) GetMetadata() map[string]*MetadataValue {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type EchoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EchoRequest) Reset() {
	*x = EchoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_e2etest_dummyserver_echo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoRequest) ProtoMessage() {}

func (x *EchoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_e2etest_dummyserver_echo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoRequest.ProtoReflect.Descriptor instead.
func (*EchoRequest) Descriptor() ([]byte, []int) {
	return file_e2etest_dummyserver_echo_proto_rawDescGZIP(), []int{2}
}

func (x *EchoRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type EchoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	// https://godoc.org/google.golang.org/grpc/metadata#MD
	Metadata *Metadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *EchoResponse) Reset() {
	*x = EchoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_e2etest_dummyserver_echo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoResponse) ProtoMessage() {}

func (x *EchoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_e2etest_dummyserver_echo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoResponse.ProtoReflect.Descriptor instead.
func (*EchoResponse) Descriptor() ([]byte, []int) {
	return file_e2etest_dummyserver_echo_proto_rawDescGZIP(), []int{3}
}

func (x *EchoResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *EchoResponse) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_e2etest_dummyserver_echo_proto protoreflect.FileDescriptor

var file_e2etest_dummyserver_echo_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x65, 0x32, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0e, 0x72, 0x65, 0x72, 0x6f, 0x73, 0x74, 0x2e, 0x67, 0x69, 0x72, 0x6f, 0x2e, 0x76, 0x31,
	0x1a, 0x17, 0x72, 0x65, 0x72, 0x6f, 0x73, 0x74, 0x2f, 0x67, 0x69, 0x72, 0x6f, 0x2f, 0x68, 0x6f,
	0x73, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x0d, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0xaa, 0x01, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x42, 0x0a,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x72, 0x65, 0x72, 0x6f, 0x73, 0x74, 0x2e, 0x67, 0x69, 0x72, 0x6f, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x1a, 0x5a, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x33, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x72, 0x65, 0x72, 0x6f, 0x73, 0x74, 0x2e, 0x67, 0x69, 0x72,
	0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x27, 0x0a,
	0x0b, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x5e, 0x0a, 0x0c, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x34, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x72, 0x65, 0x72, 0x6f, 0x73, 0x74, 0x2e, 0x67, 0x69, 0x72, 0x6f,
	0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x32, 0x68, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x1b, 0x2e,
	0x72, 0x65, 0x72, 0x6f, 0x73, 0x74, 0x2e, 0x67, 0x69, 0x72, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x72, 0x65, 0x72,
	0x6f, 0x73, 0x74, 0x2e, 0x67, 0x69, 0x72, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x63, 0x68, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x1a, 0x14, 0x82, 0xb5, 0x18, 0x10,
	0x52, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73, 0x74, 0x3a, 0x35, 0x30, 0x30, 0x30,
	0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72,
	0x65, 0x72, 0x6f, 0x73, 0x74, 0x2f, 0x67, 0x69, 0x72, 0x6f, 0x2f, 0x65, 0x32, 0x65, 0x74, 0x65,
	0x73, 0x74, 0x2f, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x3b, 0x64,
	0x75, 0x6d, 0x6d, 0x79, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_e2etest_dummyserver_echo_proto_rawDescOnce sync.Once
	file_e2etest_dummyserver_echo_proto_rawDescData = file_e2etest_dummyserver_echo_proto_rawDesc
)

func file_e2etest_dummyserver_echo_proto_rawDescGZIP() []byte {
	file_e2etest_dummyserver_echo_proto_rawDescOnce.Do(func() {
		file_e2etest_dummyserver_echo_proto_rawDescData = protoimpl.X.CompressGZIP(file_e2etest_dummyserver_echo_proto_rawDescData)
	})
	return file_e2etest_dummyserver_echo_proto_rawDescData
}

var file_e2etest_dummyserver_echo_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_e2etest_dummyserver_echo_proto_goTypes = []interface{}{
	(*MetadataValue)(nil), // 0: rerost.giro.v1.MetadataValue
	(*Metadata)(nil),      // 1: rerost.giro.v1.Metadata
	(*EchoRequest)(nil),   // 2: rerost.giro.v1.EchoRequest
	(*EchoResponse)(nil),  // 3: rerost.giro.v1.EchoResponse
	nil,                   // 4: rerost.giro.v1.Metadata.MetadataEntry
}
var file_e2etest_dummyserver_echo_proto_depIdxs = []int32{
	4, // 0: rerost.giro.v1.Metadata.metadata:type_name -> rerost.giro.v1.Metadata.MetadataEntry
	1, // 1: rerost.giro.v1.EchoResponse.metadata:type_name -> rerost.giro.v1.Metadata
	0, // 2: rerost.giro.v1.Metadata.MetadataEntry.value:type_name -> rerost.giro.v1.MetadataValue
	2, // 3: rerost.giro.v1.TestService.Echo:input_type -> rerost.giro.v1.EchoRequest
	3, // 4: rerost.giro.v1.TestService.Echo:output_type -> rerost.giro.v1.EchoResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_e2etest_dummyserver_echo_proto_init() }
func file_e2etest_dummyserver_echo_proto_init() {
	if File_e2etest_dummyserver_echo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_e2etest_dummyserver_echo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetadataValue); i {
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
		file_e2etest_dummyserver_echo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
		file_e2etest_dummyserver_echo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoRequest); i {
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
		file_e2etest_dummyserver_echo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoResponse); i {
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
			RawDescriptor: file_e2etest_dummyserver_echo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_e2etest_dummyserver_echo_proto_goTypes,
		DependencyIndexes: file_e2etest_dummyserver_echo_proto_depIdxs,
		MessageInfos:      file_e2etest_dummyserver_echo_proto_msgTypes,
	}.Build()
	File_e2etest_dummyserver_echo_proto = out.File
	file_e2etest_dummyserver_echo_proto_rawDesc = nil
	file_e2etest_dummyserver_echo_proto_goTypes = nil
	file_e2etest_dummyserver_echo_proto_depIdxs = nil
}
