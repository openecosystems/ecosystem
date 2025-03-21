// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: platform/spec/v2/spec_settings.proto

package specv2pb

import (
	_ "github.com/openecosystems/ecosystem/libs/protobuf/go/protobuf/gen/platform/options/v2"
	v2 "github.com/openecosystems/ecosystem/libs/protobuf/go/protobuf/gen/platform/type/v2"
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

type SpecSettings struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
	Version       string                 `protobuf:"bytes,11,opt,name=version,proto3" json:"version,omitempty"`
	Description   string                 `protobuf:"bytes,12,opt,name=description,proto3" json:"description,omitempty"`
	App           *App                   `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
	Platform      *Platform              `protobuf:"bytes,4,opt,name=platform,proto3" json:"platform,omitempty"`
	Context       *Context               `protobuf:"bytes,5,opt,name=context,proto3" json:"context,omitempty"`
	Systems       []*SpecSystem          `protobuf:"bytes,6,rep,name=systems,proto3" json:"systems,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SpecSettings) Reset() {
	*x = SpecSettings{}
	mi := &file_platform_spec_v2_spec_settings_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SpecSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpecSettings) ProtoMessage() {}

func (x *SpecSettings) ProtoReflect() protoreflect.Message {
	mi := &file_platform_spec_v2_spec_settings_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpecSettings.ProtoReflect.Descriptor instead.
func (*SpecSettings) Descriptor() ([]byte, []int) {
	return file_platform_spec_v2_spec_settings_proto_rawDescGZIP(), []int{0}
}

func (x *SpecSettings) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SpecSettings) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *SpecSettings) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SpecSettings) GetApp() *App {
	if x != nil {
		return x.App
	}
	return nil
}

func (x *SpecSettings) GetPlatform() *Platform {
	if x != nil {
		return x.Platform
	}
	return nil
}

func (x *SpecSettings) GetContext() *Context {
	if x != nil {
		return x.Context
	}
	return nil
}

func (x *SpecSettings) GetSystems() []*SpecSystem {
	if x != nil {
		return x.Systems
	}
	return nil
}

type App struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Name            string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version         string                 `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Description     string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	EnvironmentName string                 `protobuf:"bytes,4,opt,name=environment_name,json=environmentName,proto3" json:"environment_name,omitempty"`
	EnvironmentType string                 `protobuf:"bytes,5,opt,name=environment_type,json=environmentType,proto3" json:"environment_type,omitempty"`
	Debug           bool                   `protobuf:"varint,6,opt,name=debug,proto3" json:"debug,omitempty"`
	Verbose         bool                   `protobuf:"varint,7,opt,name=verbose,proto3" json:"verbose,omitempty"`
	Quiet           bool                   `protobuf:"varint,8,opt,name=quiet,proto3" json:"quiet,omitempty"`
	LogToFile       bool                   `protobuf:"varint,9,opt,name=log_to_file,json=logToFile,proto3" json:"log_to_file,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *App) Reset() {
	*x = App{}
	mi := &file_platform_spec_v2_spec_settings_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *App) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*App) ProtoMessage() {}

func (x *App) ProtoReflect() protoreflect.Message {
	mi := &file_platform_spec_v2_spec_settings_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use App.ProtoReflect.Descriptor instead.
func (*App) Descriptor() ([]byte, []int) {
	return file_platform_spec_v2_spec_settings_proto_rawDescGZIP(), []int{1}
}

func (x *App) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *App) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *App) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *App) GetEnvironmentName() string {
	if x != nil {
		return x.EnvironmentName
	}
	return ""
}

func (x *App) GetEnvironmentType() string {
	if x != nil {
		return x.EnvironmentType
	}
	return ""
}

func (x *App) GetDebug() bool {
	if x != nil {
		return x.Debug
	}
	return false
}

func (x *App) GetVerbose() bool {
	if x != nil {
		return x.Verbose
	}
	return false
}

func (x *App) GetQuiet() bool {
	if x != nil {
		return x.Quiet
	}
	return false
}

func (x *App) GetLogToFile() bool {
	if x != nil {
		return x.LogToFile
	}
	return false
}

type Platform struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	Endpoint            string                 `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Insecure            bool                   `protobuf:"varint,2,opt,name=insecure,proto3" json:"insecure,omitempty"`
	DnsEndpoints        []string               `protobuf:"bytes,3,rep,name=dns_endpoints,json=dnsEndpoints,proto3" json:"dns_endpoints,omitempty"`
	DynamicConfigReload bool                   `protobuf:"varint,4,opt,name=dynamic_config_reload,json=dynamicConfigReload,proto3" json:"dynamic_config_reload,omitempty"`
	Mesh                *Mesh                  `protobuf:"bytes,5,opt,name=mesh,proto3" json:"mesh,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *Platform) Reset() {
	*x = Platform{}
	mi := &file_platform_spec_v2_spec_settings_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Platform) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Platform) ProtoMessage() {}

func (x *Platform) ProtoReflect() protoreflect.Message {
	mi := &file_platform_spec_v2_spec_settings_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Platform.ProtoReflect.Descriptor instead.
func (*Platform) Descriptor() ([]byte, []int) {
	return file_platform_spec_v2_spec_settings_proto_rawDescGZIP(), []int{2}
}

func (x *Platform) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *Platform) GetInsecure() bool {
	if x != nil {
		return x.Insecure
	}
	return false
}

func (x *Platform) GetDnsEndpoints() []string {
	if x != nil {
		return x.DnsEndpoints
	}
	return nil
}

func (x *Platform) GetDynamicConfigReload() bool {
	if x != nil {
		return x.DynamicConfigReload
	}
	return false
}

func (x *Platform) GetMesh() *Mesh {
	if x != nil {
		return x.Mesh
	}
	return nil
}

type Mesh struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Enabled        bool                   `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Endpoint       string                 `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Insecure       bool                   `protobuf:"varint,3,opt,name=insecure,proto3" json:"insecure,omitempty"`
	DnsEndpoint    string                 `protobuf:"bytes,4,opt,name=dns_endpoint,json=dnsEndpoint,proto3" json:"dns_endpoint,omitempty"`
	UdpEndpoint    string                 `protobuf:"bytes,5,opt,name=udp_endpoint,json=udpEndpoint,proto3" json:"udp_endpoint,omitempty"`
	Punchy         bool                   `protobuf:"varint,6,opt,name=punchy,proto3" json:"punchy,omitempty"`
	CredentialPath string                 `protobuf:"bytes,7,opt,name=credential_path,json=credentialPath,proto3" json:"credential_path,omitempty"`
	DnsServer      bool                   `protobuf:"varint,8,opt,name=dns_server,json=dnsServer,proto3" json:"dns_server,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *Mesh) Reset() {
	*x = Mesh{}
	mi := &file_platform_spec_v2_spec_settings_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Mesh) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mesh) ProtoMessage() {}

func (x *Mesh) ProtoReflect() protoreflect.Message {
	mi := &file_platform_spec_v2_spec_settings_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mesh.ProtoReflect.Descriptor instead.
func (*Mesh) Descriptor() ([]byte, []int) {
	return file_platform_spec_v2_spec_settings_proto_rawDescGZIP(), []int{3}
}

func (x *Mesh) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Mesh) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *Mesh) GetInsecure() bool {
	if x != nil {
		return x.Insecure
	}
	return false
}

func (x *Mesh) GetDnsEndpoint() string {
	if x != nil {
		return x.DnsEndpoint
	}
	return ""
}

func (x *Mesh) GetUdpEndpoint() string {
	if x != nil {
		return x.UdpEndpoint
	}
	return ""
}

func (x *Mesh) GetPunchy() bool {
	if x != nil {
		return x.Punchy
	}
	return false
}

func (x *Mesh) GetCredentialPath() string {
	if x != nil {
		return x.CredentialPath
	}
	return ""
}

func (x *Mesh) GetDnsServer() bool {
	if x != nil {
		return x.DnsServer
	}
	return false
}

type Context struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Headers       []*v2.Header           `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Context) Reset() {
	*x = Context{}
	mi := &file_platform_spec_v2_spec_settings_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Context) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Context) ProtoMessage() {}

func (x *Context) ProtoReflect() protoreflect.Message {
	mi := &file_platform_spec_v2_spec_settings_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Context.ProtoReflect.Descriptor instead.
func (*Context) Descriptor() ([]byte, []int) {
	return file_platform_spec_v2_spec_settings_proto_rawDescGZIP(), []int{4}
}

func (x *Context) GetHeaders() []*v2.Header {
	if x != nil {
		return x.Headers
	}
	return nil
}

type DependencyRegistry struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Git           *v2.GitRepository      `protobuf:"bytes,1,opt,name=git,proto3" json:"git,omitempty"`
	Path          string                 `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Registry      string                 `protobuf:"bytes,3,opt,name=registry,proto3" json:"registry,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DependencyRegistry) Reset() {
	*x = DependencyRegistry{}
	mi := &file_platform_spec_v2_spec_settings_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DependencyRegistry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DependencyRegistry) ProtoMessage() {}

func (x *DependencyRegistry) ProtoReflect() protoreflect.Message {
	mi := &file_platform_spec_v2_spec_settings_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DependencyRegistry.ProtoReflect.Descriptor instead.
func (*DependencyRegistry) Descriptor() ([]byte, []int) {
	return file_platform_spec_v2_spec_settings_proto_rawDescGZIP(), []int{5}
}

func (x *DependencyRegistry) GetGit() *v2.GitRepository {
	if x != nil {
		return x.Git
	}
	return nil
}

func (x *DependencyRegistry) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *DependencyRegistry) GetRegistry() string {
	if x != nil {
		return x.Registry
	}
	return ""
}

type SpecSystem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version       string                 `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Protocols     []v2.Protocol          `protobuf:"varint,3,rep,packed,name=protocols,proto3,enum=platform.type.v2.Protocol" json:"protocols,omitempty"`
	Registry      *DependencyRegistry    `protobuf:"bytes,4,opt,name=registry,proto3" json:"registry,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SpecSystem) Reset() {
	*x = SpecSystem{}
	mi := &file_platform_spec_v2_spec_settings_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SpecSystem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpecSystem) ProtoMessage() {}

func (x *SpecSystem) ProtoReflect() protoreflect.Message {
	mi := &file_platform_spec_v2_spec_settings_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpecSystem.ProtoReflect.Descriptor instead.
func (*SpecSystem) Descriptor() ([]byte, []int) {
	return file_platform_spec_v2_spec_settings_proto_rawDescGZIP(), []int{6}
}

func (x *SpecSystem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SpecSystem) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *SpecSystem) GetProtocols() []v2.Protocol {
	if x != nil {
		return x.Protocols
	}
	return nil
}

func (x *SpecSystem) GetRegistry() *DependencyRegistry {
	if x != nil {
		return x.Registry
	}
	return nil
}

var File_platform_spec_v2_spec_settings_proto protoreflect.FileDescriptor

var file_platform_spec_v2_spec_settings_proto_rawDesc = string([]byte{
	0x0a, 0x24, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x2f,
	0x76, 0x32, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2e, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x76, 0x32, 0x1a, 0x25, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76,
	0x32, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x2f, 0x76, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x25, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x2f, 0x76, 0x32, 0x2f, 0x67, 0x69, 0x74, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x02, 0x0a, 0x0c, 0x53, 0x70, 0x65,
	0x63, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x03, 0x61, 0x70, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x03, 0x61,
	0x70, 0x70, 0x12, 0x36, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x73, 0x70, 0x65, 0x63, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x33, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x76, 0x32, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12,
	0x36, 0x0a, 0x07, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x73, 0x70, 0x65, 0x63,
	0x2e, 0x76, 0x32, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x07,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x3a, 0x06, 0xfa, 0xb6, 0x18, 0x02, 0x08, 0x01, 0x22,
	0x91, 0x02, 0x0a, 0x03, 0x41, 0x70, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x65, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x64, 0x65, 0x62, 0x75, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x64, 0x65,
	0x62, 0x75, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x76, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x71, 0x75, 0x69, 0x65, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x71, 0x75,
	0x69, 0x65, 0x74, 0x12, 0x1e, 0x0a, 0x0b, 0x6c, 0x6f, 0x67, 0x5f, 0x74, 0x6f, 0x5f, 0x66, 0x69,
	0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x54, 0x6f, 0x46,
	0x69, 0x6c, 0x65, 0x22, 0xcf, 0x01, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x6e, 0x73, 0x5f,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0c, 0x64, 0x6e, 0x73, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x32, 0x0a,
	0x15, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f,
	0x72, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x64, 0x79,
	0x6e, 0x61, 0x6d, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x2a, 0x0a, 0x04, 0x6d, 0x65, 0x73, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x2e,
	0x76, 0x32, 0x2e, 0x4d, 0x65, 0x73, 0x68, 0x52, 0x04, 0x6d, 0x65, 0x73, 0x68, 0x3a, 0x06, 0xfa,
	0xb6, 0x18, 0x02, 0x08, 0x01, 0x22, 0x86, 0x02, 0x0a, 0x04, 0x4d, 0x65, 0x73, 0x68, 0x12, 0x18,
	0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x64, 0x6e, 0x73, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x6e, 0x73, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x64, 0x70, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x64, 0x70, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x6e, 0x63, 0x68, 0x79,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x70, 0x75, 0x6e, 0x63, 0x68, 0x79, 0x12, 0x27,
	0x0a, 0x0f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x6e, 0x73, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x64, 0x6e, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x3a, 0x06, 0xfa, 0xb6, 0x18, 0x02, 0x08, 0x01, 0x22, 0x45,
	0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x32, 0x0a, 0x07, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x3a, 0x06, 0xfa,
	0xb6, 0x18, 0x02, 0x08, 0x01, 0x22, 0x77, 0x0a, 0x12, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65,
	0x6e, 0x63, 0x79, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x31, 0x0a, 0x03, 0x67,
	0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x69, 0x74, 0x52,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x03, 0x67, 0x69, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x22, 0xb6,
	0x01, 0x0a, 0x0a, 0x53, 0x70, 0x65, 0x63, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x09, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1a,
	0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76,
	0x32, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x12, 0x40, 0x0a, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x65, 0x70, 0x65, 0x6e,
	0x64, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x52, 0x08, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x42, 0x5d, 0x5a, 0x5b, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x73, 0x2f, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x6c,
	0x69, 0x62, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x76, 0x32, 0x3b, 0x73, 0x70,
	0x65, 0x63, 0x76, 0x32, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_platform_spec_v2_spec_settings_proto_rawDescOnce sync.Once
	file_platform_spec_v2_spec_settings_proto_rawDescData []byte
)

func file_platform_spec_v2_spec_settings_proto_rawDescGZIP() []byte {
	file_platform_spec_v2_spec_settings_proto_rawDescOnce.Do(func() {
		file_platform_spec_v2_spec_settings_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_platform_spec_v2_spec_settings_proto_rawDesc), len(file_platform_spec_v2_spec_settings_proto_rawDesc)))
	})
	return file_platform_spec_v2_spec_settings_proto_rawDescData
}

var file_platform_spec_v2_spec_settings_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_platform_spec_v2_spec_settings_proto_goTypes = []any{
	(*SpecSettings)(nil),       // 0: platform.spec.v2.SpecSettings
	(*App)(nil),                // 1: platform.spec.v2.App
	(*Platform)(nil),           // 2: platform.spec.v2.Platform
	(*Mesh)(nil),               // 3: platform.spec.v2.Mesh
	(*Context)(nil),            // 4: platform.spec.v2.Context
	(*DependencyRegistry)(nil), // 5: platform.spec.v2.DependencyRegistry
	(*SpecSystem)(nil),         // 6: platform.spec.v2.SpecSystem
	(*v2.Header)(nil),          // 7: platform.type.v2.Header
	(*v2.GitRepository)(nil),   // 8: platform.type.v2.GitRepository
	(v2.Protocol)(0),           // 9: platform.type.v2.Protocol
}
var file_platform_spec_v2_spec_settings_proto_depIdxs = []int32{
	1, // 0: platform.spec.v2.SpecSettings.app:type_name -> platform.spec.v2.App
	2, // 1: platform.spec.v2.SpecSettings.platform:type_name -> platform.spec.v2.Platform
	4, // 2: platform.spec.v2.SpecSettings.context:type_name -> platform.spec.v2.Context
	6, // 3: platform.spec.v2.SpecSettings.systems:type_name -> platform.spec.v2.SpecSystem
	3, // 4: platform.spec.v2.Platform.mesh:type_name -> platform.spec.v2.Mesh
	7, // 5: platform.spec.v2.Context.headers:type_name -> platform.type.v2.Header
	8, // 6: platform.spec.v2.DependencyRegistry.git:type_name -> platform.type.v2.GitRepository
	9, // 7: platform.spec.v2.SpecSystem.protocols:type_name -> platform.type.v2.Protocol
	5, // 8: platform.spec.v2.SpecSystem.registry:type_name -> platform.spec.v2.DependencyRegistry
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_platform_spec_v2_spec_settings_proto_init() }
func file_platform_spec_v2_spec_settings_proto_init() {
	if File_platform_spec_v2_spec_settings_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_platform_spec_v2_spec_settings_proto_rawDesc), len(file_platform_spec_v2_spec_settings_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_platform_spec_v2_spec_settings_proto_goTypes,
		DependencyIndexes: file_platform_spec_v2_spec_settings_proto_depIdxs,
		MessageInfos:      file_platform_spec_v2_spec_settings_proto_msgTypes,
	}.Build()
	File_platform_spec_v2_spec_settings_proto = out.File
	file_platform_spec_v2_spec_settings_proto_goTypes = nil
	file_platform_spec_v2_spec_settings_proto_depIdxs = nil
}
