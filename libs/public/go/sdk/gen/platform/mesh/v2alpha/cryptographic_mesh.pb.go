// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: platform/mesh/v2alpha/cryptographic_mesh.proto

package meshv2alphapb

import (
	_ "github.com/openecosystems/ecosystem/libs/protobuf/go/protobuf/gen/platform/options/v2"
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

type CryptographicMeshConfiguration struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MeshConfig    string                 `protobuf:"bytes,1,opt,name=mesh_config,json=meshConfig,proto3" json:"mesh_config,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CryptographicMeshConfiguration) Reset() {
	*x = CryptographicMeshConfiguration{}
	mi := &file_platform_mesh_v2alpha_cryptographic_mesh_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CryptographicMeshConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CryptographicMeshConfiguration) ProtoMessage() {}

func (x *CryptographicMeshConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_platform_mesh_v2alpha_cryptographic_mesh_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CryptographicMeshConfiguration.ProtoReflect.Descriptor instead.
func (*CryptographicMeshConfiguration) Descriptor() ([]byte, []int) {
	return file_platform_mesh_v2alpha_cryptographic_mesh_proto_rawDescGZIP(), []int{0}
}

func (x *CryptographicMeshConfiguration) GetMeshConfig() string {
	if x != nil {
		return x.MeshConfig
	}
	return ""
}

var File_platform_mesh_v2alpha_cryptographic_mesh_proto protoreflect.FileDescriptor

const file_platform_mesh_v2alpha_cryptographic_mesh_proto_rawDesc = "" +
	"\n" +
	".platform/mesh/v2alpha/cryptographic_mesh.proto\x12\x15platform.mesh.v2alpha\x1a%platform/options/v2/annotations.proto\"A\n" +
	"\x1eCryptographicMeshConfiguration\x12\x1f\n" +
	"\vmesh_config\x18\x01 \x01(\tR\n" +
	"meshConfigB\x83\x01\x82\xc4\x13\x02\b\x02\x82\xb5\x18\x06\b\x03\x10\x01\x18\x06\x92\xb5\x18\x03\n" +
	"\x01\x03\x9a\xb5\x18\x02\b\x01\xa2\xb5\x18\x02\b\x01Z^github.com/openecosystems/ecosystem/libs/public/go/sdk/gen/platform/mesh/v2alpha;meshv2alphapbb\x06proto3"

var (
	file_platform_mesh_v2alpha_cryptographic_mesh_proto_rawDescOnce sync.Once
	file_platform_mesh_v2alpha_cryptographic_mesh_proto_rawDescData []byte
)

func file_platform_mesh_v2alpha_cryptographic_mesh_proto_rawDescGZIP() []byte {
	file_platform_mesh_v2alpha_cryptographic_mesh_proto_rawDescOnce.Do(func() {
		file_platform_mesh_v2alpha_cryptographic_mesh_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_platform_mesh_v2alpha_cryptographic_mesh_proto_rawDesc), len(file_platform_mesh_v2alpha_cryptographic_mesh_proto_rawDesc)))
	})
	return file_platform_mesh_v2alpha_cryptographic_mesh_proto_rawDescData
}

var file_platform_mesh_v2alpha_cryptographic_mesh_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_platform_mesh_v2alpha_cryptographic_mesh_proto_goTypes = []any{
	(*CryptographicMeshConfiguration)(nil), // 0: platform.mesh.v2alpha.CryptographicMeshConfiguration
}
var file_platform_mesh_v2alpha_cryptographic_mesh_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_platform_mesh_v2alpha_cryptographic_mesh_proto_init() }
func file_platform_mesh_v2alpha_cryptographic_mesh_proto_init() {
	if File_platform_mesh_v2alpha_cryptographic_mesh_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_platform_mesh_v2alpha_cryptographic_mesh_proto_rawDesc), len(file_platform_mesh_v2alpha_cryptographic_mesh_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_platform_mesh_v2alpha_cryptographic_mesh_proto_goTypes,
		DependencyIndexes: file_platform_mesh_v2alpha_cryptographic_mesh_proto_depIdxs,
		MessageInfos:      file_platform_mesh_v2alpha_cryptographic_mesh_proto_msgTypes,
	}.Build()
	File_platform_mesh_v2alpha_cryptographic_mesh_proto = out.File
	file_platform_mesh_v2alpha_cryptographic_mesh_proto_goTypes = nil
	file_platform_mesh_v2alpha_cryptographic_mesh_proto_depIdxs = nil
}
