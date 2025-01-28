package configurationdefaultsv2alphalib

import (
	referencev2alphapb "libs/poc/go/protobuf/gen/platform/reference/v2alpha"
	auditv2alphapb "libs/private/go/protobuf/gen/platform/audit/v2alpha"
	edgev2alphapb "libs/private/go/protobuf/gen/platform/edge/v2alpha"
	communicationv1alphapb "libs/public/go/protobuf/gen/platform/communication/v1alpha"
	v1beta "libs/public/go/protobuf/gen/platform/communication/v1beta"
	configurationv2alphapb "libs/public/go/protobuf/gen/platform/configuration/v2alpha"
	cryptographyv2alphapb "libs/public/go/protobuf/gen/platform/cryptography/v2alpha"
	systemv2alphapb "libs/public/go/protobuf/gen/platform/system/v2alpha"
)

// DefaultEnterpriseConfiguration defines the default enterprise platform configuration with multiple sub-configurations.
var DefaultEnterpriseConfiguration = configurationv2alphapb.SpecPlatformConfiguration{
	PreferenceCenterConfigurationV1Beta:  &v1beta.PreferenceCenterConfiguration{},
	PreferenceCenterConfigurationV1Alpha: &communicationv1alphapb.PreferenceCenterConfiguration{},
	AuditConfigurationV2Alpha: &auditv2alphapb.AuditConfiguration{
		Id: "",
	},
	EncryptionConfigurationV2Alpha: &cryptographyv2alphapb.EncryptionConfiguration{},
	EdgeRouterConfigurationV2Alpha: &edgev2alphapb.EdgeRouterConfiguration{
		EdgeRouterConfig: "",
	},
	ReferenceConfigurationV2Alpha: &referencev2alphapb.ReferenceConfiguration{},
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
}
