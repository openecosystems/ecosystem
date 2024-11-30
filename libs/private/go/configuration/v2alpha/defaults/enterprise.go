package configurationdefaultsv2alphalib

import (
	"libs/poc/go/protobuf/gen/platform/reference/v2alpha"
	"libs/private/go/protobuf/gen/platform/audit/v2alpha"
	v2alpha1 "libs/private/go/protobuf/gen/platform/edge/v2alpha"
	"libs/private/go/protobuf/gen/platform/encryption/v2alpha"
	v2alpha3 "libs/private/go/protobuf/gen/platform/event/v2alpha"
	"libs/public/go/protobuf/gen/platform/communication/v1alpha"
	v1beta "libs/public/go/protobuf/gen/platform/communication/v1beta"
	"libs/public/go/protobuf/gen/platform/configuration/v2alpha"
	"libs/public/go/protobuf/gen/platform/event/v2alpha"
	v2alpha "libs/public/go/protobuf/gen/platform/iam/v2alpha"
	"libs/public/go/protobuf/gen/platform/system/v2alpha"
)

var DefaultEnterpriseConfiguration = configurationv2alphapb.SpecPlatformConfiguration{
	AuditConfigurationV2Alpha: &auditv2alphapb.AuditConfiguration{
		Id: "",
	},
	EdgeRouterConfigurationV2Alpha: &v2alpha1.EdgeRouterConfiguration{
		EdgeRouterConfig: "",
	},
	EncryptionConfigurationV2Alpha:       &encryptionv2alphapb.EncryptionConfiguration{},
	EventMultiplexerConfigurationV2Alpha: &v2alpha3.EventMultiplexerConfiguration{},
	EventSubscriptionConfigurationV2Alpha: &eventv2alphapb.EventSubscriptionConfiguration{
		ConfigOne:   "",
		ConfigTwo:   false,
		ConfigThree: 0,
	},
	IamApiKeyConfigurationV2Alpha:         &v2alpha.IamApiKeyConfiguration{},
	IamAuthenticationConfigurationV2Alpha: &v2alpha.IamAuthenticationConfiguration{},
	ReferenceConfigurationV2Alpha:         &referencev2alphapb.ReferenceConfiguration{},
	SystemConfigurationV2Alpha: &systemv2alphapb.SystemConfiguration{
		PublicSystems: &systemv2alphapb.PublicSystems{
			Communication: false,
			Configuration: false,
			Event:         false,
			Iam:           false,
			Ontology:      false,
		},
		PrivateSystems: &systemv2alphapb.PrivateSystems{
			Audit:      false,
			Edge:       false,
			Encryption: false,
			Event:      false,
		},
		PocSystems: &systemv2alphapb.PocSystems{
			Reference: false,
		},
	},
	PreferenceCenterConfigurationV1Beta:  &v1beta.PreferenceCenterConfiguration{},
	PreferenceCenterConfigurationV1Alpha: &communicationv1alphapb.PreferenceCenterConfiguration{},
}
