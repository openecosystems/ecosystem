// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: platform/type/v2/credential.proto

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

type CredentialType int32

const (
	CredentialType_CREDENTIAL_TYPE_UNSPECIFIED       CredentialType = 0 // Unspecified or unknown credential.
	CredentialType_CREDENTIAL_TYPE_ACCOUNT_AUTHORITY CredentialType = 1 // Credential for an account authority.
	CredentialType_CREDENTIAL_TYPE_MESH_ACCOUNT      CredentialType = 2 // Credential for a mesh service account.
)

// Enum value maps for CredentialType.
var (
	CredentialType_name = map[int32]string{
		0: "CREDENTIAL_TYPE_UNSPECIFIED",
		1: "CREDENTIAL_TYPE_ACCOUNT_AUTHORITY",
		2: "CREDENTIAL_TYPE_MESH_ACCOUNT",
	}
	CredentialType_value = map[string]int32{
		"CREDENTIAL_TYPE_UNSPECIFIED":       0,
		"CREDENTIAL_TYPE_ACCOUNT_AUTHORITY": 1,
		"CREDENTIAL_TYPE_MESH_ACCOUNT":      2,
	}
)

func (x CredentialType) Enum() *CredentialType {
	p := new(CredentialType)
	*p = x
	return p
}

func (x CredentialType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CredentialType) Descriptor() protoreflect.EnumDescriptor {
	return file_platform_type_v2_credential_proto_enumTypes[0].Descriptor()
}

func (CredentialType) Type() protoreflect.EnumType {
	return &file_platform_type_v2_credential_proto_enumTypes[0]
}

func (x CredentialType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CredentialType.Descriptor instead.
func (CredentialType) EnumDescriptor() ([]byte, []int) {
	return file_platform_type_v2_credential_proto_rawDescGZIP(), []int{0}
}

type Curve int32

const (
	Curve_CURVE_UNSPECIFIED Curve = 0
	Curve_CURVE_EDDSA       Curve = 1
	Curve_CURVE_ECDSA       Curve = 2
)

// Enum value maps for Curve.
var (
	Curve_name = map[int32]string{
		0: "CURVE_UNSPECIFIED",
		1: "CURVE_EDDSA",
		2: "CURVE_ECDSA",
	}
	Curve_value = map[string]int32{
		"CURVE_UNSPECIFIED": 0,
		"CURVE_EDDSA":       1,
		"CURVE_ECDSA":       2,
	}
)

func (x Curve) Enum() *Curve {
	p := new(Curve)
	*p = x
	return p
}

func (x Curve) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Curve) Descriptor() protoreflect.EnumDescriptor {
	return file_platform_type_v2_credential_proto_enumTypes[1].Descriptor()
}

func (Curve) Type() protoreflect.EnumType {
	return &file_platform_type_v2_credential_proto_enumTypes[1]
}

func (x Curve) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Curve.Descriptor instead.
func (Curve) EnumDescriptor() ([]byte, []int) {
	return file_platform_type_v2_credential_proto_rawDescGZIP(), []int{1}
}

type Credential struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	Type             CredentialType         `protobuf:"varint,1,opt,name=type,proto3,enum=platform.type.v2.CredentialType" json:"type,omitempty"`
	MeshAccountId    string                 `protobuf:"bytes,2,opt,name=mesh_account_id,json=meshAccountId,proto3" json:"mesh_account_id,omitempty"`
	EcosystemSlug    string                 `protobuf:"bytes,3,opt,name=ecosystem_slug,json=ecosystemSlug,proto3" json:"ecosystem_slug,omitempty"`
	MeshHostname     string                 `protobuf:"bytes,4,opt,name=mesh_hostname,json=meshHostname,proto3" json:"mesh_hostname,omitempty"`
	MeshIp           string                 `protobuf:"bytes,5,opt,name=mesh_ip,json=meshIp,proto3" json:"mesh_ip,omitempty"`
	Curve            Curve                  `protobuf:"varint,6,opt,name=curve,proto3,enum=platform.type.v2.Curve" json:"curve,omitempty"`
	AaCertX509       string                 `protobuf:"bytes,7,opt,name=aa_cert_x509,json=aaCertX509,proto3" json:"aa_cert_x509,omitempty"`
	AaCertX509QrCode string                 `protobuf:"bytes,8,opt,name=aa_cert_x509_qr_code,json=aaCertX509QrCode,proto3" json:"aa_cert_x509_qr_code,omitempty"`
	AaPrivateKey     string                 `protobuf:"bytes,9,opt,name=aa_private_key,json=aaPrivateKey,proto3" json:"aa_private_key,omitempty"`
	CertX509         string                 `protobuf:"bytes,10,opt,name=cert_x509,json=certX509,proto3" json:"cert_x509,omitempty"`
	CertX509QrCode   string                 `protobuf:"bytes,11,opt,name=cert_x509_qr_code,json=certX509QrCode,proto3" json:"cert_x509_qr_code,omitempty"`
	PrivateKey       string                 `protobuf:"bytes,12,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	NKey             string                 `protobuf:"bytes,13,opt,name=n_key,json=nKey,proto3" json:"n_key,omitempty"`
	Groups           []string               `protobuf:"bytes,14,rep,name=groups,proto3" json:"groups,omitempty"`
	Subnets          []string               `protobuf:"bytes,15,rep,name=subnets,proto3" json:"subnets,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *Credential) Reset() {
	*x = Credential{}
	mi := &file_platform_type_v2_credential_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Credential) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Credential) ProtoMessage() {}

func (x *Credential) ProtoReflect() protoreflect.Message {
	mi := &file_platform_type_v2_credential_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Credential.ProtoReflect.Descriptor instead.
func (*Credential) Descriptor() ([]byte, []int) {
	return file_platform_type_v2_credential_proto_rawDescGZIP(), []int{0}
}

func (x *Credential) GetType() CredentialType {
	if x != nil {
		return x.Type
	}
	return CredentialType_CREDENTIAL_TYPE_UNSPECIFIED
}

func (x *Credential) GetMeshAccountId() string {
	if x != nil {
		return x.MeshAccountId
	}
	return ""
}

func (x *Credential) GetEcosystemSlug() string {
	if x != nil {
		return x.EcosystemSlug
	}
	return ""
}

func (x *Credential) GetMeshHostname() string {
	if x != nil {
		return x.MeshHostname
	}
	return ""
}

func (x *Credential) GetMeshIp() string {
	if x != nil {
		return x.MeshIp
	}
	return ""
}

func (x *Credential) GetCurve() Curve {
	if x != nil {
		return x.Curve
	}
	return Curve_CURVE_UNSPECIFIED
}

func (x *Credential) GetAaCertX509() string {
	if x != nil {
		return x.AaCertX509
	}
	return ""
}

func (x *Credential) GetAaCertX509QrCode() string {
	if x != nil {
		return x.AaCertX509QrCode
	}
	return ""
}

func (x *Credential) GetAaPrivateKey() string {
	if x != nil {
		return x.AaPrivateKey
	}
	return ""
}

func (x *Credential) GetCertX509() string {
	if x != nil {
		return x.CertX509
	}
	return ""
}

func (x *Credential) GetCertX509QrCode() string {
	if x != nil {
		return x.CertX509QrCode
	}
	return ""
}

func (x *Credential) GetPrivateKey() string {
	if x != nil {
		return x.PrivateKey
	}
	return ""
}

func (x *Credential) GetNKey() string {
	if x != nil {
		return x.NKey
	}
	return ""
}

func (x *Credential) GetGroups() []string {
	if x != nil {
		return x.Groups
	}
	return nil
}

func (x *Credential) GetSubnets() []string {
	if x != nil {
		return x.Subnets
	}
	return nil
}

var File_platform_type_v2_credential_proto protoreflect.FileDescriptor

var file_platform_type_v2_credential_proto_rawDesc = string([]byte{
	0x0a, 0x21, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f,
	0x76, 0x32, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x10, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x76, 0x32, 0x22, 0xa6, 0x04, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x12, 0x34, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x20, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x6d, 0x65,
	0x73, 0x68, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x65, 0x73, 0x68, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f,
	0x73, 0x6c, 0x75, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x63, 0x6f, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x53, 0x6c, 0x75, 0x67, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x65, 0x73,
	0x68, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x6d, 0x65, 0x73, 0x68, 0x48, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x68, 0x5f, 0x69, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6d, 0x65, 0x73, 0x68, 0x49, 0x70, 0x12, 0x2d, 0x0a, 0x05, 0x63, 0x75, 0x72, 0x76, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x75, 0x72, 0x76, 0x65, 0x52,
	0x05, 0x63, 0x75, 0x72, 0x76, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x61, 0x61, 0x5f, 0x63, 0x65, 0x72,
	0x74, 0x5f, 0x78, 0x35, 0x30, 0x39, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x61,
	0x43, 0x65, 0x72, 0x74, 0x58, 0x35, 0x30, 0x39, 0x12, 0x2e, 0x0a, 0x14, 0x61, 0x61, 0x5f, 0x63,
	0x65, 0x72, 0x74, 0x5f, 0x78, 0x35, 0x30, 0x39, 0x5f, 0x71, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x61, 0x43, 0x65, 0x72, 0x74, 0x58, 0x35,
	0x30, 0x39, 0x51, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x61, 0x61, 0x5f, 0x70,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x61, 0x61, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x1b,
	0x0a, 0x09, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x78, 0x35, 0x30, 0x39, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x65, 0x72, 0x74, 0x58, 0x35, 0x30, 0x39, 0x12, 0x29, 0x0a, 0x11, 0x63,
	0x65, 0x72, 0x74, 0x5f, 0x78, 0x35, 0x30, 0x39, 0x5f, 0x71, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x65, 0x72, 0x74, 0x58, 0x35, 0x30, 0x39,
	0x51, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x13, 0x0a, 0x05, 0x6e, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x73, 0x18,
	0x0f, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x73, 0x2a, 0x7a,
	0x0a, 0x0e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1f, 0x0a, 0x1b, 0x43, 0x52, 0x45, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x41, 0x4c, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x25, 0x0a, 0x21, 0x43, 0x52, 0x45, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x41, 0x4c, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x41, 0x55, 0x54,
	0x48, 0x4f, 0x52, 0x49, 0x54, 0x59, 0x10, 0x01, 0x12, 0x20, 0x0a, 0x1c, 0x43, 0x52, 0x45, 0x44,
	0x45, 0x4e, 0x54, 0x49, 0x41, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x45, 0x53, 0x48,
	0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x02, 0x2a, 0x40, 0x0a, 0x05, 0x43, 0x75,
	0x72, 0x76, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x55, 0x52, 0x56, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x55,
	0x52, 0x56, 0x45, 0x5f, 0x45, 0x44, 0x44, 0x53, 0x41, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x43,
	0x55, 0x52, 0x56, 0x45, 0x5f, 0x45, 0x43, 0x44, 0x53, 0x41, 0x10, 0x02, 0x42, 0x5d, 0x5a, 0x5b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x65,
	0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x2f, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2f, 0x6c, 0x69, 0x62, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f,
	0x76, 0x32, 0x3b, 0x74, 0x79, 0x70, 0x65, 0x76, 0x32, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_platform_type_v2_credential_proto_rawDescOnce sync.Once
	file_platform_type_v2_credential_proto_rawDescData []byte
)

func file_platform_type_v2_credential_proto_rawDescGZIP() []byte {
	file_platform_type_v2_credential_proto_rawDescOnce.Do(func() {
		file_platform_type_v2_credential_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_platform_type_v2_credential_proto_rawDesc), len(file_platform_type_v2_credential_proto_rawDesc)))
	})
	return file_platform_type_v2_credential_proto_rawDescData
}

var file_platform_type_v2_credential_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_platform_type_v2_credential_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_platform_type_v2_credential_proto_goTypes = []any{
	(CredentialType)(0), // 0: platform.type.v2.CredentialType
	(Curve)(0),          // 1: platform.type.v2.Curve
	(*Credential)(nil),  // 2: platform.type.v2.Credential
}
var file_platform_type_v2_credential_proto_depIdxs = []int32{
	0, // 0: platform.type.v2.Credential.type:type_name -> platform.type.v2.CredentialType
	1, // 1: platform.type.v2.Credential.curve:type_name -> platform.type.v2.Curve
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_platform_type_v2_credential_proto_init() }
func file_platform_type_v2_credential_proto_init() {
	if File_platform_type_v2_credential_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_platform_type_v2_credential_proto_rawDesc), len(file_platform_type_v2_credential_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_platform_type_v2_credential_proto_goTypes,
		DependencyIndexes: file_platform_type_v2_credential_proto_depIdxs,
		EnumInfos:         file_platform_type_v2_credential_proto_enumTypes,
		MessageInfos:      file_platform_type_v2_credential_proto_msgTypes,
	}.Build()
	File_platform_type_v2_credential_proto = out.File
	file_platform_type_v2_credential_proto_goTypes = nil
	file_platform_type_v2_credential_proto_depIdxs = nil
}
