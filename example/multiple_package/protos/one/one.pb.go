// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v5.27.4
// source: protos/one/one.proto

package one_pb

import (
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

type GiroTestRequest1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GiroTestRequest1) Reset() {
	*x = GiroTestRequest1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_one_one_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GiroTestRequest1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GiroTestRequest1) ProtoMessage() {}

func (x *GiroTestRequest1) ProtoReflect() protoreflect.Message {
	mi := &file_protos_one_one_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GiroTestRequest1.ProtoReflect.Descriptor instead.
func (*GiroTestRequest1) Descriptor() ([]byte, []int) {
	return file_protos_one_one_proto_rawDescGZIP(), []int{0}
}

func (x *GiroTestRequest1) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GiroTestRequest2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GiroTestRequest2) Reset() {
	*x = GiroTestRequest2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_one_one_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GiroTestRequest2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GiroTestRequest2) ProtoMessage() {}

func (x *GiroTestRequest2) ProtoReflect() protoreflect.Message {
	mi := &file_protos_one_one_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GiroTestRequest2.ProtoReflect.Descriptor instead.
func (*GiroTestRequest2) Descriptor() ([]byte, []int) {
	return file_protos_one_one_proto_rawDescGZIP(), []int{1}
}

type GiroTestResponse1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GiroTestResponse1) Reset() {
	*x = GiroTestResponse1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_one_one_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GiroTestResponse1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GiroTestResponse1) ProtoMessage() {}

func (x *GiroTestResponse1) ProtoReflect() protoreflect.Message {
	mi := &file_protos_one_one_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GiroTestResponse1.ProtoReflect.Descriptor instead.
func (*GiroTestResponse1) Descriptor() ([]byte, []int) {
	return file_protos_one_one_proto_rawDescGZIP(), []int{2}
}

func (x *GiroTestResponse1) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GiroTestResponse2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GiroTestResponse2) Reset() {
	*x = GiroTestResponse2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_one_one_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GiroTestResponse2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GiroTestResponse2) ProtoMessage() {}

func (x *GiroTestResponse2) ProtoReflect() protoreflect.Message {
	mi := &file_protos_one_one_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GiroTestResponse2.ProtoReflect.Descriptor instead.
func (*GiroTestResponse2) Descriptor() ([]byte, []int) {
	return file_protos_one_one_proto_rawDescGZIP(), []int{3}
}

var File_protos_one_one_proto protoreflect.FileDescriptor

var file_protos_one_one_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x6e, 0x65, 0x2f, 0x6f, 0x6e, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x6e, 0x65, 0x22, 0x2c, 0x0a, 0x10, 0x47,
	0x69, 0x72, 0x6f, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x31, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x47, 0x69, 0x72,
	0x6f, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0x22, 0x2d, 0x0a,
	0x11, 0x47, 0x69, 0x72, 0x6f, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x31, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x13, 0x0a, 0x11,
	0x47, 0x69, 0x72, 0x6f, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x32, 0x32, 0x89, 0x02, 0x0a, 0x0b, 0x47, 0x69, 0x72, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x7c, 0x0a, 0x09, 0x47, 0x69, 0x72, 0x6f, 0x54, 0x65, 0x73, 0x74, 0x31, 0x12, 0x35,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c,
	0x65, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x6f, 0x6e, 0x65, 0x2e, 0x47, 0x69, 0x72, 0x6f, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x31, 0x1a, 0x36, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x6e, 0x65, 0x2e, 0x47, 0x69, 0x72, 0x6f,
	0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x31, 0x22, 0x00, 0x12,
	0x7c, 0x0a, 0x09, 0x47, 0x69, 0x72, 0x6f, 0x54, 0x65, 0x73, 0x74, 0x32, 0x12, 0x35, 0x2e, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x5f,
	0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f,
	0x6e, 0x65, 0x2e, 0x47, 0x69, 0x72, 0x6f, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x32, 0x1a, 0x36, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x6d, 0x75,
	0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x6e, 0x65, 0x2e, 0x47, 0x69, 0x72, 0x6f, 0x54, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x22, 0x00, 0x42, 0x43, 0x5a,
	0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x72, 0x6f,
	0x73, 0x74, 0x2f, 0x67, 0x69, 0x72, 0x6f, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f,
	0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x6e, 0x65, 0x3b, 0x6f, 0x6e, 0x65, 0x5f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_one_one_proto_rawDescOnce sync.Once
	file_protos_one_one_proto_rawDescData = file_protos_one_one_proto_rawDesc
)

func file_protos_one_one_proto_rawDescGZIP() []byte {
	file_protos_one_one_proto_rawDescOnce.Do(func() {
		file_protos_one_one_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_one_one_proto_rawDescData)
	})
	return file_protos_one_one_proto_rawDescData
}

var file_protos_one_one_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protos_one_one_proto_goTypes = []interface{}{
	(*GiroTestRequest1)(nil),  // 0: example.multiple_package.protos.one.GiroTestRequest1
	(*GiroTestRequest2)(nil),  // 1: example.multiple_package.protos.one.GiroTestRequest2
	(*GiroTestResponse1)(nil), // 2: example.multiple_package.protos.one.GiroTestResponse1
	(*GiroTestResponse2)(nil), // 3: example.multiple_package.protos.one.GiroTestResponse2
}
var file_protos_one_one_proto_depIdxs = []int32{
	0, // 0: example.multiple_package.protos.one.GiroService.GiroTest1:input_type -> example.multiple_package.protos.one.GiroTestRequest1
	1, // 1: example.multiple_package.protos.one.GiroService.GiroTest2:input_type -> example.multiple_package.protos.one.GiroTestRequest2
	2, // 2: example.multiple_package.protos.one.GiroService.GiroTest1:output_type -> example.multiple_package.protos.one.GiroTestResponse1
	3, // 3: example.multiple_package.protos.one.GiroService.GiroTest2:output_type -> example.multiple_package.protos.one.GiroTestResponse2
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_one_one_proto_init() }
func file_protos_one_one_proto_init() {
	if File_protos_one_one_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_one_one_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GiroTestRequest1); i {
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
		file_protos_one_one_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GiroTestRequest2); i {
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
		file_protos_one_one_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GiroTestResponse1); i {
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
		file_protos_one_one_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GiroTestResponse2); i {
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
			RawDescriptor: file_protos_one_one_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_one_one_proto_goTypes,
		DependencyIndexes: file_protos_one_one_proto_depIdxs,
		MessageInfos:      file_protos_one_one_proto_msgTypes,
	}.Build()
	File_protos_one_one_proto = out.File
	file_protos_one_one_proto_rawDesc = nil
	file_protos_one_one_proto_goTypes = nil
	file_protos_one_one_proto_depIdxs = nil
}
