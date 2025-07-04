// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: platform/cryptography/v2alpha/encryption.proto

package cryptographyv2alphapb

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

type EncryptionConfiguration struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EncryptionConfiguration) Reset() {
	*x = EncryptionConfiguration{}
	mi := &file_platform_cryptography_v2alpha_encryption_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EncryptionConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptionConfiguration) ProtoMessage() {}

func (x *EncryptionConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_platform_cryptography_v2alpha_encryption_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptionConfiguration.ProtoReflect.Descriptor instead.
func (*EncryptionConfiguration) Descriptor() ([]byte, []int) {
	return file_platform_cryptography_v2alpha_encryption_proto_rawDescGZIP(), []int{0}
}

type EncryptionContext struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// TODO: Revisit these types.
	User          string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Entity        string `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"`
	Principal     string `protobuf:"bytes,3,opt,name=principal,proto3" json:"principal,omitempty"`
	Intent        string `protobuf:"bytes,4,opt,name=intent,proto3" json:"intent,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EncryptionContext) Reset() {
	*x = EncryptionContext{}
	mi := &file_platform_cryptography_v2alpha_encryption_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EncryptionContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptionContext) ProtoMessage() {}

func (x *EncryptionContext) ProtoReflect() protoreflect.Message {
	mi := &file_platform_cryptography_v2alpha_encryption_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptionContext.ProtoReflect.Descriptor instead.
func (*EncryptionContext) Descriptor() ([]byte, []int) {
	return file_platform_cryptography_v2alpha_encryption_proto_rawDescGZIP(), []int{1}
}

func (x *EncryptionContext) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *EncryptionContext) GetEntity() string {
	if x != nil {
		return x.Entity
	}
	return ""
}

func (x *EncryptionContext) GetPrincipal() string {
	if x != nil {
		return x.Principal
	}
	return ""
}

func (x *EncryptionContext) GetIntent() string {
	if x != nil {
		return x.Intent
	}
	return ""
}

type CipherText struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	CipherText        []byte                 `protobuf:"bytes,1,opt,name=cipher_text,json=cipherText,proto3" json:"cipher_text,omitempty"`
	EncryptionContext *EncryptionContext     `protobuf:"bytes,2,opt,name=encryption_context,json=encryptionContext,proto3" json:"encryption_context,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *CipherText) Reset() {
	*x = CipherText{}
	mi := &file_platform_cryptography_v2alpha_encryption_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CipherText) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CipherText) ProtoMessage() {}

func (x *CipherText) ProtoReflect() protoreflect.Message {
	mi := &file_platform_cryptography_v2alpha_encryption_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CipherText.ProtoReflect.Descriptor instead.
func (*CipherText) Descriptor() ([]byte, []int) {
	return file_platform_cryptography_v2alpha_encryption_proto_rawDescGZIP(), []int{2}
}

func (x *CipherText) GetCipherText() []byte {
	if x != nil {
		return x.CipherText
	}
	return nil
}

func (x *CipherText) GetEncryptionContext() *EncryptionContext {
	if x != nil {
		return x.EncryptionContext
	}
	return nil
}

type EncryptRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	PlainText      []byte                 `protobuf:"bytes,1,opt,name=plain_text,json=plainText,proto3" json:"plain_text,omitempty"`
	AssociatedData []byte                 `protobuf:"bytes,2,opt,name=associated_data,json=associatedData,proto3" json:"associated_data,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *EncryptRequest) Reset() {
	*x = EncryptRequest{}
	mi := &file_platform_cryptography_v2alpha_encryption_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EncryptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptRequest) ProtoMessage() {}

func (x *EncryptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_platform_cryptography_v2alpha_encryption_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptRequest.ProtoReflect.Descriptor instead.
func (*EncryptRequest) Descriptor() ([]byte, []int) {
	return file_platform_cryptography_v2alpha_encryption_proto_rawDescGZIP(), []int{3}
}

func (x *EncryptRequest) GetPlainText() []byte {
	if x != nil {
		return x.PlainText
	}
	return nil
}

func (x *EncryptRequest) GetAssociatedData() []byte {
	if x != nil {
		return x.AssociatedData
	}
	return nil
}

type EncryptResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Result:
	//
	//	*EncryptResponse_CipherText
	//	*EncryptResponse_Err
	Result        isEncryptResponse_Result `protobuf_oneof:"result"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EncryptResponse) Reset() {
	*x = EncryptResponse{}
	mi := &file_platform_cryptography_v2alpha_encryption_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EncryptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptResponse) ProtoMessage() {}

func (x *EncryptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_platform_cryptography_v2alpha_encryption_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptResponse.ProtoReflect.Descriptor instead.
func (*EncryptResponse) Descriptor() ([]byte, []int) {
	return file_platform_cryptography_v2alpha_encryption_proto_rawDescGZIP(), []int{4}
}

func (x *EncryptResponse) GetResult() isEncryptResponse_Result {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *EncryptResponse) GetCipherText() *CipherText {
	if x != nil {
		if x, ok := x.Result.(*EncryptResponse_CipherText); ok {
			return x.CipherText
		}
	}
	return nil
}

func (x *EncryptResponse) GetErr() string {
	if x != nil {
		if x, ok := x.Result.(*EncryptResponse_Err); ok {
			return x.Err
		}
	}
	return ""
}

type isEncryptResponse_Result interface {
	isEncryptResponse_Result()
}

type EncryptResponse_CipherText struct {
	CipherText *CipherText `protobuf:"bytes,1,opt,name=cipher_text,json=cipherText,proto3,oneof"`
}

type EncryptResponse_Err struct {
	Err string `protobuf:"bytes,2,opt,name=err,proto3,oneof"`
}

func (*EncryptResponse_CipherText) isEncryptResponse_Result() {}

func (*EncryptResponse_Err) isEncryptResponse_Result() {}

type DecryptRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	CipherText     *CipherText            `protobuf:"bytes,1,opt,name=cipher_text,json=cipherText,proto3" json:"cipher_text,omitempty"`
	AssociatedData []byte                 `protobuf:"bytes,2,opt,name=associated_data,json=associatedData,proto3" json:"associated_data,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *DecryptRequest) Reset() {
	*x = DecryptRequest{}
	mi := &file_platform_cryptography_v2alpha_encryption_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DecryptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecryptRequest) ProtoMessage() {}

func (x *DecryptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_platform_cryptography_v2alpha_encryption_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecryptRequest.ProtoReflect.Descriptor instead.
func (*DecryptRequest) Descriptor() ([]byte, []int) {
	return file_platform_cryptography_v2alpha_encryption_proto_rawDescGZIP(), []int{5}
}

func (x *DecryptRequest) GetCipherText() *CipherText {
	if x != nil {
		return x.CipherText
	}
	return nil
}

func (x *DecryptRequest) GetAssociatedData() []byte {
	if x != nil {
		return x.AssociatedData
	}
	return nil
}

type DecryptResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Result:
	//
	//	*DecryptResponse_PlainText
	//	*DecryptResponse_Err
	Result        isDecryptResponse_Result `protobuf_oneof:"result"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DecryptResponse) Reset() {
	*x = DecryptResponse{}
	mi := &file_platform_cryptography_v2alpha_encryption_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DecryptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecryptResponse) ProtoMessage() {}

func (x *DecryptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_platform_cryptography_v2alpha_encryption_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecryptResponse.ProtoReflect.Descriptor instead.
func (*DecryptResponse) Descriptor() ([]byte, []int) {
	return file_platform_cryptography_v2alpha_encryption_proto_rawDescGZIP(), []int{6}
}

func (x *DecryptResponse) GetResult() isDecryptResponse_Result {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *DecryptResponse) GetPlainText() []byte {
	if x != nil {
		if x, ok := x.Result.(*DecryptResponse_PlainText); ok {
			return x.PlainText
		}
	}
	return nil
}

func (x *DecryptResponse) GetErr() string {
	if x != nil {
		if x, ok := x.Result.(*DecryptResponse_Err); ok {
			return x.Err
		}
	}
	return ""
}

type isDecryptResponse_Result interface {
	isDecryptResponse_Result()
}

type DecryptResponse_PlainText struct {
	PlainText []byte `protobuf:"bytes,1,opt,name=plain_text,json=plainText,proto3,oneof"`
}

type DecryptResponse_Err struct {
	Err string `protobuf:"bytes,2,opt,name=err,proto3,oneof"`
}

func (*DecryptResponse_PlainText) isDecryptResponse_Result() {}

func (*DecryptResponse_Err) isDecryptResponse_Result() {}

var File_platform_cryptography_v2alpha_encryption_proto protoreflect.FileDescriptor

const file_platform_cryptography_v2alpha_encryption_proto_rawDesc = "" +
	"\n" +
	".platform/cryptography/v2alpha/encryption.proto\x12\x1dplatform.cryptography.v2alpha\x1a%platform/options/v2/annotations.proto\"\x19\n" +
	"\x17EncryptionConfiguration\"u\n" +
	"\x11EncryptionContext\x12\x12\n" +
	"\x04user\x18\x01 \x01(\tR\x04user\x12\x16\n" +
	"\x06entity\x18\x02 \x01(\tR\x06entity\x12\x1c\n" +
	"\tprincipal\x18\x03 \x01(\tR\tprincipal\x12\x16\n" +
	"\x06intent\x18\x04 \x01(\tR\x06intent\"\x8e\x01\n" +
	"\n" +
	"CipherText\x12\x1f\n" +
	"\vcipher_text\x18\x01 \x01(\fR\n" +
	"cipherText\x12_\n" +
	"\x12encryption_context\x18\x02 \x01(\v20.platform.cryptography.v2alpha.EncryptionContextR\x11encryptionContext\"X\n" +
	"\x0eEncryptRequest\x12\x1d\n" +
	"\n" +
	"plain_text\x18\x01 \x01(\fR\tplainText\x12'\n" +
	"\x0fassociated_data\x18\x02 \x01(\fR\x0eassociatedData\"}\n" +
	"\x0fEncryptResponse\x12L\n" +
	"\vcipher_text\x18\x01 \x01(\v2).platform.cryptography.v2alpha.CipherTextH\x00R\n" +
	"cipherText\x12\x12\n" +
	"\x03err\x18\x02 \x01(\tH\x00R\x03errB\b\n" +
	"\x06result\"\x85\x01\n" +
	"\x0eDecryptRequest\x12J\n" +
	"\vcipher_text\x18\x01 \x01(\v2).platform.cryptography.v2alpha.CipherTextR\n" +
	"cipherText\x12'\n" +
	"\x0fassociated_data\x18\x02 \x01(\fR\x0eassociatedData\"P\n" +
	"\x0fDecryptResponse\x12\x1f\n" +
	"\n" +
	"plain_text\x18\x01 \x01(\fH\x00R\tplainText\x12\x12\n" +
	"\x03err\x18\x02 \x01(\tH\x00R\x03errB\b\n" +
	"\x06result2\xeb\x01\n" +
	"\x11EncryptionService\x12j\n" +
	"\aEncrypt\x12-.platform.cryptography.v2alpha.EncryptRequest\x1a..platform.cryptography.v2alpha.EncryptResponse\"\x00\x12j\n" +
	"\aDecrypt\x12-.platform.cryptography.v2alpha.DecryptRequest\x1a..platform.cryptography.v2alpha.DecryptResponse\"\x00B\x92\x01\x82\xc4\x13\x02\b\x01\x82\xb5\x18\x06\b\x03\x10\x01\x18\x02\x92\xb5\x18\x04\n" +
	"\x02\x03\x01\x9a\xb5\x18\x00\xa2\xb5\x18\x02\b\x01Zngithub.com/openecosystems/ecosystem/libs/public/go/sdk/gen/platform/cryptography/v2alpha;cryptographyv2alphapbb\x06proto3"

var (
	file_platform_cryptography_v2alpha_encryption_proto_rawDescOnce sync.Once
	file_platform_cryptography_v2alpha_encryption_proto_rawDescData []byte
)

func file_platform_cryptography_v2alpha_encryption_proto_rawDescGZIP() []byte {
	file_platform_cryptography_v2alpha_encryption_proto_rawDescOnce.Do(func() {
		file_platform_cryptography_v2alpha_encryption_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_platform_cryptography_v2alpha_encryption_proto_rawDesc), len(file_platform_cryptography_v2alpha_encryption_proto_rawDesc)))
	})
	return file_platform_cryptography_v2alpha_encryption_proto_rawDescData
}

var file_platform_cryptography_v2alpha_encryption_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_platform_cryptography_v2alpha_encryption_proto_goTypes = []any{
	(*EncryptionConfiguration)(nil), // 0: platform.cryptography.v2alpha.EncryptionConfiguration
	(*EncryptionContext)(nil),       // 1: platform.cryptography.v2alpha.EncryptionContext
	(*CipherText)(nil),              // 2: platform.cryptography.v2alpha.CipherText
	(*EncryptRequest)(nil),          // 3: platform.cryptography.v2alpha.EncryptRequest
	(*EncryptResponse)(nil),         // 4: platform.cryptography.v2alpha.EncryptResponse
	(*DecryptRequest)(nil),          // 5: platform.cryptography.v2alpha.DecryptRequest
	(*DecryptResponse)(nil),         // 6: platform.cryptography.v2alpha.DecryptResponse
}
var file_platform_cryptography_v2alpha_encryption_proto_depIdxs = []int32{
	1, // 0: platform.cryptography.v2alpha.CipherText.encryption_context:type_name -> platform.cryptography.v2alpha.EncryptionContext
	2, // 1: platform.cryptography.v2alpha.EncryptResponse.cipher_text:type_name -> platform.cryptography.v2alpha.CipherText
	2, // 2: platform.cryptography.v2alpha.DecryptRequest.cipher_text:type_name -> platform.cryptography.v2alpha.CipherText
	3, // 3: platform.cryptography.v2alpha.EncryptionService.Encrypt:input_type -> platform.cryptography.v2alpha.EncryptRequest
	5, // 4: platform.cryptography.v2alpha.EncryptionService.Decrypt:input_type -> platform.cryptography.v2alpha.DecryptRequest
	4, // 5: platform.cryptography.v2alpha.EncryptionService.Encrypt:output_type -> platform.cryptography.v2alpha.EncryptResponse
	6, // 6: platform.cryptography.v2alpha.EncryptionService.Decrypt:output_type -> platform.cryptography.v2alpha.DecryptResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_platform_cryptography_v2alpha_encryption_proto_init() }
func file_platform_cryptography_v2alpha_encryption_proto_init() {
	if File_platform_cryptography_v2alpha_encryption_proto != nil {
		return
	}
	file_platform_cryptography_v2alpha_encryption_proto_msgTypes[4].OneofWrappers = []any{
		(*EncryptResponse_CipherText)(nil),
		(*EncryptResponse_Err)(nil),
	}
	file_platform_cryptography_v2alpha_encryption_proto_msgTypes[6].OneofWrappers = []any{
		(*DecryptResponse_PlainText)(nil),
		(*DecryptResponse_Err)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_platform_cryptography_v2alpha_encryption_proto_rawDesc), len(file_platform_cryptography_v2alpha_encryption_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_platform_cryptography_v2alpha_encryption_proto_goTypes,
		DependencyIndexes: file_platform_cryptography_v2alpha_encryption_proto_depIdxs,
		MessageInfos:      file_platform_cryptography_v2alpha_encryption_proto_msgTypes,
	}.Build()
	File_platform_cryptography_v2alpha_encryption_proto = out.File
	file_platform_cryptography_v2alpha_encryption_proto_goTypes = nil
	file_platform_cryptography_v2alpha_encryption_proto_depIdxs = nil
}
