// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: platform/type/v2/transport.proto

package typev2pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Change this to only be TCP, UDP, QUIC. Remove others
type Transport int32

const (
	Transport_TRANSPORT_UNSPECIFIED Transport = 0
	Transport_TRANSPORT_GRPC        Transport = 1
	Transport_TRANSPORT_REST        Transport = 2
	Transport_TRANSPORT_GRAPHQL     Transport = 3
	Transport_TRANSPORT_TCP         Transport = 4
	Transport_TRANSPORT_UDP         Transport = 5
	Transport_TRANSPORT_QUIC        Transport = 6
)

// Enum value maps for Transport.
var (
	Transport_name = map[int32]string{
		0: "TRANSPORT_UNSPECIFIED",
		1: "TRANSPORT_GRPC",
		2: "TRANSPORT_REST",
		3: "TRANSPORT_GRAPHQL",
		4: "TRANSPORT_TCP",
		5: "TRANSPORT_UDP",
		6: "TRANSPORT_QUIC",
	}
	Transport_value = map[string]int32{
		"TRANSPORT_UNSPECIFIED": 0,
		"TRANSPORT_GRPC":        1,
		"TRANSPORT_REST":        2,
		"TRANSPORT_GRAPHQL":     3,
		"TRANSPORT_TCP":         4,
		"TRANSPORT_UDP":         5,
		"TRANSPORT_QUIC":        6,
	}
)

func (x Transport) Enum() *Transport {
	p := new(Transport)
	*p = x
	return p
}

func (x Transport) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Transport) Descriptor() protoreflect.EnumDescriptor {
	return file_platform_type_v2_transport_proto_enumTypes[0].Descriptor()
}

func (Transport) Type() protoreflect.EnumType {
	return &file_platform_type_v2_transport_proto_enumTypes[0]
}

func (x Transport) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Transport.Descriptor instead.
func (Transport) EnumDescriptor() ([]byte, []int) {
	return file_platform_type_v2_transport_proto_rawDescGZIP(), []int{0}
}

var File_platform_type_v2_transport_proto protoreflect.FileDescriptor

var file_platform_type_v2_transport_proto_rawDesc = string([]byte{
	0x0a, 0x20, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f,
	0x76, 0x32, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x10, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x76, 0x32, 0x2a, 0x9f, 0x01, 0x0a, 0x09, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x12, 0x19, 0x0a, 0x15, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x50, 0x4f, 0x52, 0x54, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a,
	0x0e, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x47, 0x52, 0x50, 0x43, 0x10,
	0x01, 0x12, 0x12, 0x0a, 0x0e, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x52,
	0x45, 0x53, 0x54, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x50, 0x4f,
	0x52, 0x54, 0x5f, 0x47, 0x52, 0x41, 0x50, 0x48, 0x51, 0x4c, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d,
	0x54, 0x52, 0x41, 0x4e, 0x53, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x54, 0x43, 0x50, 0x10, 0x04, 0x12,
	0x11, 0x0a, 0x0d, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x55, 0x44, 0x50,
	0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x50, 0x4f, 0x52, 0x54, 0x5f,
	0x51, 0x55, 0x49, 0x43, 0x10, 0x06, 0x42, 0x5d, 0x5a, 0x5b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x73, 0x2f, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x6c, 0x69,
	0x62, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x6f, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x32, 0x3b, 0x74, 0x79, 0x70,
	0x65, 0x76, 0x32, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_platform_type_v2_transport_proto_rawDescOnce sync.Once
	file_platform_type_v2_transport_proto_rawDescData []byte
)

func file_platform_type_v2_transport_proto_rawDescGZIP() []byte {
	file_platform_type_v2_transport_proto_rawDescOnce.Do(func() {
		file_platform_type_v2_transport_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_platform_type_v2_transport_proto_rawDesc), len(file_platform_type_v2_transport_proto_rawDesc)))
	})
	return file_platform_type_v2_transport_proto_rawDescData
}

var file_platform_type_v2_transport_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_platform_type_v2_transport_proto_goTypes = []any{
	(Transport)(0), // 0: platform.type.v2.Transport
}
var file_platform_type_v2_transport_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_platform_type_v2_transport_proto_init() }
func file_platform_type_v2_transport_proto_init() {
	if File_platform_type_v2_transport_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_platform_type_v2_transport_proto_rawDesc), len(file_platform_type_v2_transport_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_platform_type_v2_transport_proto_goTypes,
		DependencyIndexes: file_platform_type_v2_transport_proto_depIdxs,
		EnumInfos:         file_platform_type_v2_transport_proto_enumTypes,
	}.Build()
	File_platform_type_v2_transport_proto = out.File
	file_platform_type_v2_transport_proto_goTypes = nil
	file_platform_type_v2_transport_proto_depIdxs = nil
}
