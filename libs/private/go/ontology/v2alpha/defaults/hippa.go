package ontologydefaultsv2alphalib

import (
	ontologyv2alphapb "github.com/openecosystems/ecosystem/libs/public/go/sdk/gen/platform/ontology/v2alpha"

	optionv2pb "github.com/openecosystems/ecosystem/libs/protobuf/go/protobuf/gen/platform/options/v2"
)

// Hippa represents a predefined variable containing a comprehensive configuration of SpecDataCatalog for managing HIPAA-related data.
var Hippa = ontologyv2alphapb.SpecDataCatalog{
	Audit: &ontologyv2alphapb.Audit{
		AuditV2Alpha: &ontologyv2alphapb.AuditV2Alpha{
			Id:           optionv2pb.ClassificationType_CLASSIFICATION_TYPE_CONFIDENTIAL,
			CreatedAt:    0,
			UpdatedAt:    0,
			Entry:        0,
			Jurisdiction: 0,
		},
	},
	Cli: &ontologyv2alphapb.Cli{},
	Communication: &ontologyv2alphapb.Communication{
		PreferenceCenterV1Beta: &ontologyv2alphapb.PreferenceCenterV1Beta{
			Id:                  0,
			CreatedAt:           0,
			UpdatedAt:           0,
			AnonymousId:         0,
			Email:               optionv2pb.ClassificationType_CLASSIFICATION_TYPE_CONFIDENTIAL,
			ExternalId:          0,
			PhoneNumber:         0,
			FirstName:           optionv2pb.ClassificationType_CLASSIFICATION_TYPE_CONFIDENTIAL,
			LastName:            optionv2pb.ClassificationType_CLASSIFICATION_TYPE_CONFIDENTIAL,
			PostalCode:          0,
			City:                0,
			StateProvinceRegion: 0,
			Country:             0,
			ListIds:             0,
			SegmentIds:          0,
			EmailSubscription:   0,
		},
		PreferenceCenterV1Alpha: &ontologyv2alphapb.PreferenceCenterV1Alpha{
			Id:                  0,
			CreatedAt:           0,
			UpdatedAt:           0,
			AnonymousId:         0,
			Email:               optionv2pb.ClassificationType_CLASSIFICATION_TYPE_CONFIDENTIAL,
			ExternalId:          0,
			PhoneNumber:         optionv2pb.ClassificationType_CLASSIFICATION_TYPE_CONFIDENTIAL,
			FirstName:           optionv2pb.ClassificationType_CLASSIFICATION_TYPE_CONFIDENTIAL,
			LastName:            optionv2pb.ClassificationType_CLASSIFICATION_TYPE_CONFIDENTIAL,
			PostalCode:          0,
			City:                0,
			StateProvinceRegion: 0,
			Country:             0,
			ListIds:             0,
			SegmentIds:          0,
			EmailSubscription:   0,
		},
	},
	Configuration: &ontologyv2alphapb.Configuration{
		ConfigurationV2Alpha: &ontologyv2alphapb.ConfigurationV2Alpha{
			Id:                      0,
			OrganizationSlug:        0,
			WorkspaceSlug:           0,
			CreatedAt:               0,
			UpdatedAt:               0,
			SourceId:                0,
			Type:                    0,
			Status:                  0,
			StatusDetails:           0,
			ParentId:                0,
			DataCatalog:             0,
			ClinicalCatalog:         0,
			PlatformConfiguration:   0,
			PlatformConfigurations:  0,
			SolutionConfigurations:  0,
			ConnectorConfigurations: 0,
		},
	},
	Cryptography: &ontologyv2alphapb.Cryptography{},
	Dns:          nil,
	Edge:         &ontologyv2alphapb.Edge{},
	Iam:          &ontologyv2alphapb.Iam{},
	Mesh:         nil,
	Ontology:     &ontologyv2alphapb.Ontology{},
	Options:      &ontologyv2alphapb.Options{},
	Reference: &ontologyv2alphapb.Reference{
		ReferenceV2Alpha: &ontologyv2alphapb.ReferenceV2Alpha{
			Id:        0,
			CreatedAt: 0,
			UpdatedAt: 0,
		},
	},
	Spec: &ontologyv2alphapb.Spec{},
	System: &ontologyv2alphapb.System{
		SystemV2Alpha: &ontologyv2alphapb.SystemV2Alpha{
			Id:        0,
			EnabledAt: 0,
			CreatedAt: 0,
			UpdatedAt: 0,
		},
	},
	Type: &ontologyv2alphapb.Type{},
}
