// Code generated by protoc-gen-platform protobuf/configuration. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: platform/configuration/v2alpha/spec_configuration.proto

package configurationv2alphapb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	v2alpha9 "libs/poc/go/protobuf/gen/platform/reference/v2alpha"
	v2alpha "libs/private/go/protobuf/gen/platform/audit/v2alpha"
	v2alpha4 "libs/private/go/protobuf/gen/platform/edge/v2alpha"
	v2alpha5 "libs/private/go/protobuf/gen/platform/event/v2alpha"
	_ "libs/protobuf/go/protobuf/gen/platform/options/v2"
	v2alpha1 "libs/public/go/protobuf/gen/platform/cli/v2alpha"
	v1alpha "libs/public/go/protobuf/gen/platform/communication/v1alpha"
	v1beta "libs/public/go/protobuf/gen/platform/communication/v1beta"
	v2alpha2 "libs/public/go/protobuf/gen/platform/cryptography/v2alpha"
	v2alpha3 "libs/public/go/protobuf/gen/platform/dns/v2alpha"
	v2alpha6 "libs/public/go/protobuf/gen/platform/event/v2alpha"
	v2alpha7 "libs/public/go/protobuf/gen/platform/iam/v2alpha"
	v2alpha8 "libs/public/go/protobuf/gen/platform/mesh/v2alpha"
	v2alpha10 "libs/public/go/protobuf/gen/platform/system/v2alpha"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SpecPlatformConfiguration struct {
	state                                    protoimpl.MessageState                      `protogen:"open.v1"`
	AuditConfigurationV2Alpha                *v2alpha.AuditConfiguration                 `protobuf:"bytes,2,opt,name=audit_configuration_v2alpha,json=auditConfigurationV2alpha,proto3" json:"audit_configuration_v2alpha,omitempty"`
	OecoConfigurationV2Alpha                 *v2alpha1.OecoConfiguration                 `protobuf:"bytes,3,opt,name=oeco_configuration_v2alpha,json=oecoConfigurationV2alpha,proto3" json:"oeco_configuration_v2alpha,omitempty"`
	CertificateConfigurationV2Alpha          *v2alpha2.CertificateConfiguration          `protobuf:"bytes,4,opt,name=certificate_configuration_v2alpha,json=certificateConfigurationV2alpha,proto3" json:"certificate_configuration_v2alpha,omitempty"`
	CertificateAuthorityConfigurationV2Alpha *v2alpha2.CertificateAuthorityConfiguration `protobuf:"bytes,5,opt,name=certificate_authority_configuration_v2alpha,json=certificateAuthorityConfigurationV2alpha,proto3" json:"certificate_authority_configuration_v2alpha,omitempty"`
	EncryptionConfigurationV2Alpha           *v2alpha2.EncryptionConfiguration           `protobuf:"bytes,6,opt,name=encryption_configuration_v2alpha,json=encryptionConfigurationV2alpha,proto3" json:"encryption_configuration_v2alpha,omitempty"`
	DynamicDnsConfigurationV2Alpha           *v2alpha3.DynamicDnsConfiguration           `protobuf:"bytes,7,opt,name=dynamic_dns_configuration_v2alpha,json=dynamicDnsConfigurationV2alpha,proto3" json:"dynamic_dns_configuration_v2alpha,omitempty"`
	EdgeRouterConfigurationV2Alpha           *v2alpha4.EdgeRouterConfiguration           `protobuf:"bytes,8,opt,name=edge_router_configuration_v2alpha,json=edgeRouterConfigurationV2alpha,proto3" json:"edge_router_configuration_v2alpha,omitempty"`
	EventMultiplexerConfigurationV2Alpha     *v2alpha5.EventMultiplexerConfiguration     `protobuf:"bytes,9,opt,name=event_multiplexer_configuration_v2alpha,json=eventMultiplexerConfigurationV2alpha,proto3" json:"event_multiplexer_configuration_v2alpha,omitempty"`
	EventSubscriptionConfigurationV2Alpha    *v2alpha6.EventSubscriptionConfiguration    `protobuf:"bytes,10,opt,name=event_subscription_configuration_v2alpha,json=eventSubscriptionConfigurationV2alpha,proto3" json:"event_subscription_configuration_v2alpha,omitempty"`
	IamApiKeyConfigurationV2Alpha            *v2alpha7.IamApiKeyConfiguration            `protobuf:"bytes,11,opt,name=iam_api_key_configuration_v2alpha,json=iamApiKeyConfigurationV2alpha,proto3" json:"iam_api_key_configuration_v2alpha,omitempty"`
	IamAuthenticationConfigurationV2Alpha    *v2alpha7.IamAuthenticationConfiguration    `protobuf:"bytes,12,opt,name=iam_authentication_configuration_v2alpha,json=iamAuthenticationConfigurationV2alpha,proto3" json:"iam_authentication_configuration_v2alpha,omitempty"`
	CryptographicMeshConfigurationV2Alpha    *v2alpha8.CryptographicMeshConfiguration    `protobuf:"bytes,13,opt,name=cryptographic_mesh_configuration_v2alpha,json=cryptographicMeshConfigurationV2alpha,proto3" json:"cryptographic_mesh_configuration_v2alpha,omitempty"`
	ReferenceConfigurationV2Alpha            *v2alpha9.ReferenceConfiguration            `protobuf:"bytes,14,opt,name=reference_configuration_v2alpha,json=referenceConfigurationV2alpha,proto3" json:"reference_configuration_v2alpha,omitempty"`
	SystemConfigurationV2Alpha               *v2alpha10.SystemConfiguration              `protobuf:"bytes,15,opt,name=system_configuration_v2alpha,json=systemConfigurationV2alpha,proto3" json:"system_configuration_v2alpha,omitempty"`
	PreferenceCenterConfigurationV1Alpha     *v1alpha.PreferenceCenterConfiguration      `protobuf:"bytes,16,opt,name=preference_center_configuration_v1alpha,json=preferenceCenterConfigurationV1alpha,proto3" json:"preference_center_configuration_v1alpha,omitempty"`
	PreferenceCenterConfigurationV1Beta      *v1beta.PreferenceCenterConfiguration       `protobuf:"bytes,17,opt,name=preference_center_configuration_v1beta,json=preferenceCenterConfigurationV1beta,proto3" json:"preference_center_configuration_v1beta,omitempty"`
	unknownFields                            protoimpl.UnknownFields
	sizeCache                                protoimpl.SizeCache
}

func (x *SpecPlatformConfiguration) Reset() {
	*x = SpecPlatformConfiguration{}
	mi := &file_platform_configuration_v2alpha_spec_configuration_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SpecPlatformConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpecPlatformConfiguration) ProtoMessage() {}

func (x *SpecPlatformConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_platform_configuration_v2alpha_spec_configuration_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpecPlatformConfiguration.ProtoReflect.Descriptor instead.
func (*SpecPlatformConfiguration) Descriptor() ([]byte, []int) {
	return file_platform_configuration_v2alpha_spec_configuration_proto_rawDescGZIP(), []int{0}
}

func (x *SpecPlatformConfiguration) GetAuditConfigurationV2Alpha() *v2alpha.AuditConfiguration {
	if x != nil {
		return x.AuditConfigurationV2Alpha
	}
	return nil
}

func (x *SpecPlatformConfiguration) GetOecoConfigurationV2Alpha() *v2alpha1.OecoConfiguration {
	if x != nil {
		return x.OecoConfigurationV2Alpha
	}
	return nil
}

func (x *SpecPlatformConfiguration) GetCertificateConfigurationV2Alpha() *v2alpha2.CertificateConfiguration {
	if x != nil {
		return x.CertificateConfigurationV2Alpha
	}
	return nil
}

func (x *SpecPlatformConfiguration) GetCertificateAuthorityConfigurationV2Alpha() *v2alpha2.CertificateAuthorityConfiguration {
	if x != nil {
		return x.CertificateAuthorityConfigurationV2Alpha
	}
	return nil
}

func (x *SpecPlatformConfiguration) GetEncryptionConfigurationV2Alpha() *v2alpha2.EncryptionConfiguration {
	if x != nil {
		return x.EncryptionConfigurationV2Alpha
	}
	return nil
}

func (x *SpecPlatformConfiguration) GetDynamicDnsConfigurationV2Alpha() *v2alpha3.DynamicDnsConfiguration {
	if x != nil {
		return x.DynamicDnsConfigurationV2Alpha
	}
	return nil
}

func (x *SpecPlatformConfiguration) GetEdgeRouterConfigurationV2Alpha() *v2alpha4.EdgeRouterConfiguration {
	if x != nil {
		return x.EdgeRouterConfigurationV2Alpha
	}
	return nil
}

func (x *SpecPlatformConfiguration) GetEventMultiplexerConfigurationV2Alpha() *v2alpha5.EventMultiplexerConfiguration {
	if x != nil {
		return x.EventMultiplexerConfigurationV2Alpha
	}
	return nil
}

func (x *SpecPlatformConfiguration) GetEventSubscriptionConfigurationV2Alpha() *v2alpha6.EventSubscriptionConfiguration {
	if x != nil {
		return x.EventSubscriptionConfigurationV2Alpha
	}
	return nil
}

func (x *SpecPlatformConfiguration) GetIamApiKeyConfigurationV2Alpha() *v2alpha7.IamApiKeyConfiguration {
	if x != nil {
		return x.IamApiKeyConfigurationV2Alpha
	}
	return nil
}

func (x *SpecPlatformConfiguration) GetIamAuthenticationConfigurationV2Alpha() *v2alpha7.IamAuthenticationConfiguration {
	if x != nil {
		return x.IamAuthenticationConfigurationV2Alpha
	}
	return nil
}

func (x *SpecPlatformConfiguration) GetCryptographicMeshConfigurationV2Alpha() *v2alpha8.CryptographicMeshConfiguration {
	if x != nil {
		return x.CryptographicMeshConfigurationV2Alpha
	}
	return nil
}

func (x *SpecPlatformConfiguration) GetReferenceConfigurationV2Alpha() *v2alpha9.ReferenceConfiguration {
	if x != nil {
		return x.ReferenceConfigurationV2Alpha
	}
	return nil
}

func (x *SpecPlatformConfiguration) GetSystemConfigurationV2Alpha() *v2alpha10.SystemConfiguration {
	if x != nil {
		return x.SystemConfigurationV2Alpha
	}
	return nil
}

func (x *SpecPlatformConfiguration) GetPreferenceCenterConfigurationV1Alpha() *v1alpha.PreferenceCenterConfiguration {
	if x != nil {
		return x.PreferenceCenterConfigurationV1Alpha
	}
	return nil
}

func (x *SpecPlatformConfiguration) GetPreferenceCenterConfigurationV1Beta() *v1beta.PreferenceCenterConfiguration {
	if x != nil {
		return x.PreferenceCenterConfigurationV1Beta
	}
	return nil
}

var File_platform_configuration_v2alpha_spec_configuration_proto protoreflect.FileDescriptor

var file_platform_configuration_v2alpha_spec_configuration_proto_rawDesc = []byte{
	0x0a, 0x37, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2f, 0x73, 0x70, 0x65, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x25, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x22, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x63,
	0x6c, 0x69, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x6f, 0x65, 0x63, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x2f, 0x76, 0x32, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x39, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x2f, 0x76, 0x32,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x26, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x64, 0x6e, 0x73, 0x2f,
	0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f,
	0x64, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2f, 0x65, 0x64, 0x67, 0x65, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x78, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x69, 0x61,
	0x6d, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x69, 0x61, 0x6d, 0x5f, 0x61, 0x70,
	0x69, 0x5f, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2f, 0x69, 0x61, 0x6d, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x5f,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x2f, 0x76,
	0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x70, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x63,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdf, 0x10, 0x0a, 0x19,
	0x53, 0x70, 0x65, 0x63, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x6a, 0x0a, 0x1b, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a,
	0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e,
	0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x19, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x32,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x12, 0x65, 0x0a, 0x1a, 0x6f, 0x65, 0x63, 0x6f, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x32, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x63, 0x6c, 0x69, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x4f, 0x65, 0x63, 0x6f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x18, 0x6f, 0x65, 0x63, 0x6f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x12, 0x83, 0x01, 0x0a,
	0x21, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x32, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79,
	0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x1f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x32, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x12, 0x9f, 0x01, 0x0a, 0x2b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x32, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79,
	0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x28, 0x63, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x32, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x12, 0x80, 0x01, 0x0a, 0x20, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x36, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x1e, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x56, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x12, 0x78, 0x0a, 0x21, 0x64, 0x79, 0x6e, 0x61, 0x6d,
	0x69, 0x63, 0x5f, 0x64, 0x6e, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x64, 0x6e,
	0x73, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69,
	0x63, 0x44, 0x6e, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x1e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x44, 0x6e, 0x73, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x32, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x12, 0x79, 0x0a, 0x21, 0x65, 0x64, 0x67, 0x65, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76,
	0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x76, 0x32, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2e, 0x45, 0x64, 0x67, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x1e, 0x65, 0x64,
	0x67, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x12, 0x8c, 0x01, 0x0a,
	0x27, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x65, 0x78,
	0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35,
	0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x75, 0x6c,
	0x74, 0x69, 0x70, 0x6c, 0x65, 0x78, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x24, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x75, 0x6c, 0x74,
	0x69, 0x70, 0x6c, 0x65, 0x78, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x56, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x12, 0x8f, 0x01, 0x0a, 0x28,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36,
	0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x25, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x12, 0x76, 0x0a,
	0x21, 0x69, 0x61, 0x6d, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x32, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x49, 0x61, 0x6d, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x1d, 0x69, 0x61, 0x6d, 0x41, 0x70, 0x69, 0x4b, 0x65,
	0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x32,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x12, 0x8d, 0x01, 0x0a, 0x28, 0x69, 0x61, 0x6d, 0x5f, 0x61, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x32, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x49, 0x61, 0x6d, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x25,
	0x69, 0x61, 0x6d, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x32,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x12, 0x8e, 0x01, 0x0a, 0x28, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x68, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x32, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x4d, 0x65,
	0x73, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x25, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x4d, 0x65,
	0x73, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56,
	0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x12, 0x7a, 0x0a, 0x1f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x32, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x52, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x1d, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x32, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x12, 0x6e, 0x0a, 0x1c, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x32, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x1a, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x32, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x12, 0x94, 0x01, 0x0a, 0x27, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x24, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x43,
	0x65, 0x6e, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x12, 0x91, 0x01, 0x0a, 0x26, 0x70, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x50, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x23, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x42, 0x63, 0x82,
	0xc4, 0x13, 0x02, 0x08, 0x02, 0x82, 0xb5, 0x18, 0x06, 0x08, 0x03, 0x10, 0x01, 0x18, 0x02, 0x5a,
	0x51, 0x6c, 0x69, 0x62, 0x73, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x67, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3b, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_platform_configuration_v2alpha_spec_configuration_proto_rawDescOnce sync.Once
	file_platform_configuration_v2alpha_spec_configuration_proto_rawDescData = file_platform_configuration_v2alpha_spec_configuration_proto_rawDesc
)

func file_platform_configuration_v2alpha_spec_configuration_proto_rawDescGZIP() []byte {
	file_platform_configuration_v2alpha_spec_configuration_proto_rawDescOnce.Do(func() {
		file_platform_configuration_v2alpha_spec_configuration_proto_rawDescData = protoimpl.X.CompressGZIP(file_platform_configuration_v2alpha_spec_configuration_proto_rawDescData)
	})
	return file_platform_configuration_v2alpha_spec_configuration_proto_rawDescData
}

var file_platform_configuration_v2alpha_spec_configuration_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_platform_configuration_v2alpha_spec_configuration_proto_goTypes = []any{
	(*SpecPlatformConfiguration)(nil),                  // 0: platform.configuration.v2alpha.SpecPlatformConfiguration
	(*v2alpha.AuditConfiguration)(nil),                 // 1: platform.audit.v2alpha.AuditConfiguration
	(*v2alpha1.OecoConfiguration)(nil),                 // 2: platform.cli.v2alpha.OecoConfiguration
	(*v2alpha2.CertificateConfiguration)(nil),          // 3: platform.cryptography.v2alpha.CertificateConfiguration
	(*v2alpha2.CertificateAuthorityConfiguration)(nil), // 4: platform.cryptography.v2alpha.CertificateAuthorityConfiguration
	(*v2alpha2.EncryptionConfiguration)(nil),           // 5: platform.cryptography.v2alpha.EncryptionConfiguration
	(*v2alpha3.DynamicDnsConfiguration)(nil),           // 6: platform.dns.v2alpha.DynamicDnsConfiguration
	(*v2alpha4.EdgeRouterConfiguration)(nil),           // 7: platform.edge.v2alpha.EdgeRouterConfiguration
	(*v2alpha5.EventMultiplexerConfiguration)(nil),     // 8: platform.event.v2alpha.EventMultiplexerConfiguration
	(*v2alpha6.EventSubscriptionConfiguration)(nil),    // 9: platform.event.v2alpha.EventSubscriptionConfiguration
	(*v2alpha7.IamApiKeyConfiguration)(nil),            // 10: platform.iam.v2alpha.IamApiKeyConfiguration
	(*v2alpha7.IamAuthenticationConfiguration)(nil),    // 11: platform.iam.v2alpha.IamAuthenticationConfiguration
	(*v2alpha8.CryptographicMeshConfiguration)(nil),    // 12: platform.mesh.v2alpha.CryptographicMeshConfiguration
	(*v2alpha9.ReferenceConfiguration)(nil),            // 13: platform.reference.v2alpha.ReferenceConfiguration
	(*v2alpha10.SystemConfiguration)(nil),              // 14: platform.system.v2alpha.SystemConfiguration
	(*v1alpha.PreferenceCenterConfiguration)(nil),      // 15: platform.communication.v1alpha.PreferenceCenterConfiguration
	(*v1beta.PreferenceCenterConfiguration)(nil),       // 16: platform.communication.v1beta.PreferenceCenterConfiguration
}
var file_platform_configuration_v2alpha_spec_configuration_proto_depIdxs = []int32{
	1,  // 0: platform.configuration.v2alpha.SpecPlatformConfiguration.audit_configuration_v2alpha:type_name -> platform.audit.v2alpha.AuditConfiguration
	2,  // 1: platform.configuration.v2alpha.SpecPlatformConfiguration.oeco_configuration_v2alpha:type_name -> platform.cli.v2alpha.OecoConfiguration
	3,  // 2: platform.configuration.v2alpha.SpecPlatformConfiguration.certificate_configuration_v2alpha:type_name -> platform.cryptography.v2alpha.CertificateConfiguration
	4,  // 3: platform.configuration.v2alpha.SpecPlatformConfiguration.certificate_authority_configuration_v2alpha:type_name -> platform.cryptography.v2alpha.CertificateAuthorityConfiguration
	5,  // 4: platform.configuration.v2alpha.SpecPlatformConfiguration.encryption_configuration_v2alpha:type_name -> platform.cryptography.v2alpha.EncryptionConfiguration
	6,  // 5: platform.configuration.v2alpha.SpecPlatformConfiguration.dynamic_dns_configuration_v2alpha:type_name -> platform.dns.v2alpha.DynamicDnsConfiguration
	7,  // 6: platform.configuration.v2alpha.SpecPlatformConfiguration.edge_router_configuration_v2alpha:type_name -> platform.edge.v2alpha.EdgeRouterConfiguration
	8,  // 7: platform.configuration.v2alpha.SpecPlatformConfiguration.event_multiplexer_configuration_v2alpha:type_name -> platform.event.v2alpha.EventMultiplexerConfiguration
	9,  // 8: platform.configuration.v2alpha.SpecPlatformConfiguration.event_subscription_configuration_v2alpha:type_name -> platform.event.v2alpha.EventSubscriptionConfiguration
	10, // 9: platform.configuration.v2alpha.SpecPlatformConfiguration.iam_api_key_configuration_v2alpha:type_name -> platform.iam.v2alpha.IamApiKeyConfiguration
	11, // 10: platform.configuration.v2alpha.SpecPlatformConfiguration.iam_authentication_configuration_v2alpha:type_name -> platform.iam.v2alpha.IamAuthenticationConfiguration
	12, // 11: platform.configuration.v2alpha.SpecPlatformConfiguration.cryptographic_mesh_configuration_v2alpha:type_name -> platform.mesh.v2alpha.CryptographicMeshConfiguration
	13, // 12: platform.configuration.v2alpha.SpecPlatformConfiguration.reference_configuration_v2alpha:type_name -> platform.reference.v2alpha.ReferenceConfiguration
	14, // 13: platform.configuration.v2alpha.SpecPlatformConfiguration.system_configuration_v2alpha:type_name -> platform.system.v2alpha.SystemConfiguration
	15, // 14: platform.configuration.v2alpha.SpecPlatformConfiguration.preference_center_configuration_v1alpha:type_name -> platform.communication.v1alpha.PreferenceCenterConfiguration
	16, // 15: platform.configuration.v2alpha.SpecPlatformConfiguration.preference_center_configuration_v1beta:type_name -> platform.communication.v1beta.PreferenceCenterConfiguration
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_platform_configuration_v2alpha_spec_configuration_proto_init() }
func file_platform_configuration_v2alpha_spec_configuration_proto_init() {
	if File_platform_configuration_v2alpha_spec_configuration_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_platform_configuration_v2alpha_spec_configuration_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_platform_configuration_v2alpha_spec_configuration_proto_goTypes,
		DependencyIndexes: file_platform_configuration_v2alpha_spec_configuration_proto_depIdxs,
		MessageInfos:      file_platform_configuration_v2alpha_spec_configuration_proto_msgTypes,
	}.Build()
	File_platform_configuration_v2alpha_spec_configuration_proto = out.File
	file_platform_configuration_v2alpha_spec_configuration_proto_rawDesc = nil
	file_platform_configuration_v2alpha_spec_configuration_proto_goTypes = nil
	file_platform_configuration_v2alpha_spec_configuration_proto_depIdxs = nil
}
