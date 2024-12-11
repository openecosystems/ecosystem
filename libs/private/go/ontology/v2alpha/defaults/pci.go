package ontologydefaultsv2alphalib

import (
	optionv2pb "libs/protobuf/go/protobuf/gen/platform/options/v2"
	ontologyv2alphapb "libs/public/go/protobuf/gen/platform/ontology/v2alpha"
)

var PCI = &ontologyv2alphapb.SpecDataCatalog{
	Audit: &ontologyv2alphapb.Audit{
		AuditV2Alpha: &ontologyv2alphapb.AuditV2Alpha{
			Id:           optionv2pb.ClassificationType_CLASSIFICATION_TYPE_CONFIDENTIAL,
			CreatedAt:    0,
			UpdatedAt:    0,
			Entry:        0,
			Jurisdiction: 0,
		},
	},
	Communication: &ontologyv2alphapb.Communication{
		PreferenceCenterV1Beta: &ontologyv2alphapb.PreferenceCenterV1Beta{
			Id:                  0,
			CreatedAt:           0,
			UpdatedAt:           0,
			AnonymousId:         0,
			Email:               0,
			ExternalId:          0,
			PhoneNumber:         0,
			FirstName:           0,
			LastName:            0,
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
			Email:               0,
			ExternalId:          0,
			PhoneNumber:         0,
			FirstName:           0,
			LastName:            0,
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
	Cryptography: &ontologyv2alphapb.Cryptography{
		CertificateV2Alpha: &ontologyv2alphapb.CertificateV2Alpha{
			Id:        0,
			CreatedAt: 0,
			UpdatedAt: 0,
			Name:      0,
			Duration:  0,
		},
		CertificateAuthorityV2Alpha: &ontologyv2alphapb.CertificateAuthorityV2Alpha{
			Id:        0,
			CreatedAt: 0,
			UpdatedAt: 0,
			Name:      0,
			Curve:     0,
			Duration:  0,
			CaCert:    0,
			CaKey:     0,
			CaQrCode:  0,
		},
	},
	Edge: &ontologyv2alphapb.Edge{},
	Event: &ontologyv2alphapb.Event{
		EventSubscriptionV2Alpha: &ontologyv2alphapb.EventSubscriptionV2Alpha{
			Id:            0,
			CreatedAt:     0,
			UpdatedAt:     0,
			Status:        0,
			StatusDetails: 0,
			Type:          0,
			Data:          0,
		},
	},
	Iam:      &ontologyv2alphapb.Iam{},
	Ontology: &ontologyv2alphapb.Ontology{},
	Options:  &ontologyv2alphapb.Options{},
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
