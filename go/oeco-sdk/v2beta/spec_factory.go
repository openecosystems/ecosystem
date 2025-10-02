//nolint:revive
package sdkv2betalib

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	specv2pb "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta/gen/platform/spec/v2"
	typev2pb "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta/gen/platform/type/v2"

	"google.golang.org/protobuf/types/known/fieldmaskpb"

	apexlog "github.com/apex/log"
	optionv2pb "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta/gen/platform/options/v2"
	"github.com/segmentio/ksuid"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// DefaultSpecVersion defines the default specification version to be used.
// DefaultConnectionId specifies the default connection identifier.
const (
	DefaultSpecVersion  = "v2"
	DefaultConnectionId = "corporate" //nolint:revive
)

// Factory represents an entity responsible for creating and initializing resources or objects.
// Spec is a pointer to a Spec structure that holds specification details.
// Headers is a map containing key-value pairs for custom headers.
type Factory struct {
	Spec    *specv2pb.Spec
	Headers map[string]string
}

// NewFactory creates and initializes a new Factory instance using the provided `connect.AnyRequest`.
// It extracts headers, processes key metadata, and constructs a structured `specv2pb.Spec` object.
// Returns a Factory containing the built `specv2pb.Spec` and a map of parsed headers.
func NewFactory(ctx context.Context, h http.Header, procedure string) Factory {
	headers := make(map[string]string, len(h))
	for k, v := range h {
		if len(v) > 0 {
			values := strings.Join(v, "; ")
			headers[k] = values
		}
	}

	// Spec
	// ===============================
	specVersion := DefaultSpecVersion

	messageId := ksuid.New().String()

	requestId := h.Get(RequestIdKey)
	if requestId == "" {
		requestId = ksuid.New().String()
	}

	sentAt := timestamppb.Now()
	if _sentAt, ok := headers[SentAtKey]; ok {
		t, err := time.Parse(time.RFC3339, _sentAt)
		if err != nil {
			sentAt = timestamppb.New(t)
		}
	}

	receivedAt := timestamppb.Now()

	// Completed at is provided upstream by the implementing service/turbine
	completedAt := &timestamppb.Timestamp{Seconds: 0, Nanos: 0}

	specType := ""

	// Spec.SpecPrincipal
	// ===============================
	anonymousId := h.Get(AnonymousIdKey)
	principalId := h.Get(PrincipalIdKey)

	var principalType specv2pb.SpecPrincipalType
	_principalType := h.Get(PrincipalTypeKey)
	switch _principalType {
	case "SPEC_PRINCIPAL_TYPE_ANONYMOUS":
		principalType = specv2pb.SpecPrincipalType_SPEC_PRINCIPAL_TYPE_ANONYMOUS
	case "SPEC_PRINCIPAL_TYPE_USER":
		principalType = specv2pb.SpecPrincipalType_SPEC_PRINCIPAL_TYPE_USER
	case "SPEC_PRINCIPAL_TYPE_SERVICE_ACCOUNT":
		principalType = specv2pb.SpecPrincipalType_SPEC_PRINCIPAL_TYPE_SERVICE_ACCOUNT
	case "SPEC_PRINCIPAL_TYPE_GROUP":
		principalType = specv2pb.SpecPrincipalType_SPEC_PRINCIPAL_TYPE_GROUP
	case "SPEC_PRINCIPAL_TYPE_DOMAIN":
		principalType = specv2pb.SpecPrincipalType_SPEC_PRINCIPAL_TYPE_DOMAIN
	default:
		principalType = specv2pb.SpecPrincipalType_SPEC_PRINCIPAL_TYPE_UNSPECIFIED
	}

	principalEmail := h.Get(PrincipalEmailKey)
	connectionId := h.Get(ConnectionIdKey)

	_roles := h.Values(RolesKey)
	var roles []optionv2pb.AuthRole
	for _, role := range _roles {
		var r optionv2pb.AuthRole
		switch role {
		case "AUTH_ROLE_ANONYMOUS":
			r = optionv2pb.AuthRole_AUTH_ROLE_ANONYMOUS
		case "AUTH_ROLE_PLATFORM_SUPER_ADMIN":
			r = optionv2pb.AuthRole_AUTH_ROLE_PLATFORM_SUPER_ADMIN
		case "AUTH_ROLE_PLATFORM_ADMIN":
			r = optionv2pb.AuthRole_AUTH_ROLE_PLATFORM_ADMIN
		case "AUTH_ROLE_ORGANIZATION_ADMIN":
			r = optionv2pb.AuthRole_AUTH_ROLE_ORGANIZATION_ADMIN
		case "AUTH_ROLE_ORGANIZATION_USER":
			r = optionv2pb.AuthRole_AUTH_ROLE_ORGANIZATION_USER
		default:
		}

		roles = append(roles, r)
	}

	// Span.Context
	// ===============================
	traceId := h.Get(XB3Traceid)
	spanId := h.Get(XB3Spanid)
	parentSpanId := h.Get(XB3Parentspanid)
	traceFlags := h.Get(XB3Flags)
	spanCtx := trace.SpanContextFromContext(ctx)
	if spanCtx.HasTraceID() {
		traceId = spanCtx.TraceID().String()
	}
	if spanCtx.HasSpanID() {
		spanId = spanCtx.SpanID().String()
	}

	if traceId == "" {
		traceId = GenerateTraceID()
	}
	if spanId == "" {
		spanId = GenerateSpanID()
	}
	if traceFlags == "" {
		traceFlags = fmt.Sprintf("%02x", trace.TraceFlags(0x00))
	}

	// Spec.Context
	// ===============================
	ecosystemID := h.Get(EcosystemID)
	ecosystemSlug := h.Get(EcosystemSlug)
	organizationID := h.Get(OrganizationID)
	organizationSlug := h.Get(OrganizationSlug)

	var jan typev2pb.Jurisdiction
	_jan := h.Get(JurisdictionAreaNetworkKey)
	switch _jan {
	case "JURISDICTION_NA_US_1":
		jan = typev2pb.Jurisdiction_JURISDICTION_NA_US_1
	case "JURISDICTION_GOV_US_1":
		jan = typev2pb.Jurisdiction_JURISDICTION_GOV_US_1
	case "JURISDICTION_EU_DE_1":
		jan = typev2pb.Jurisdiction_JURISDICTION_EU_DE_1
	case "JURISDICTION_GOV_EU_1  ":
		jan = typev2pb.Jurisdiction_JURISDICTION_GOV_EU_1
	case "JURISDICTION_AS_CN_1":
		jan = typev2pb.Jurisdiction_JURISDICTION_AS_CN_1
	case "JURISDICTION_SA_BR_1":
		jan = typev2pb.Jurisdiction_JURISDICTION_SA_BR_1
	default:
		jan = typev2pb.Jurisdiction_JURISDICTION_UNSPECIFIED
	}

	ip := h.Get(IpKey)
	if h.Get(XForwardedFor) != "" {
		ip = h.Get(XForwardedFor)
	}

	locale := h.Get(AcceptLanguage)
	timezone := h.Get(TimezoneKey)
	userAgent := h.Get(UserAgentKey)

	// Spec.Context.Validation
	// ===============================
	validateOnly := false
	_validateOnly := h.Get(ValidateOnlyKey)
	vo, err := strconv.ParseBool(_validateOnly)
	if err == nil {
		validateOnly = vo
	}

	// Spec.Context.Device
	// ===============================
	deviceId := h.Get(DeviceIdKey)
	deviceAdvertisingId := h.Get(DeviceAdvertisingIdKey)
	deviceManufacturer := h.Get(DeviceManufacturerKey)
	deviceModel := h.Get(DeviceModelKey)
	deviceName := h.Get(DeviceNameKey)
	deviceType := h.Get(DeviceTypeKey)
	deviceToken := h.Get(DeviceTokenKey)

	// Spec.Context.Location
	// ===============================
	continent := h.Get(ContinentKey)
	country := h.Get(CountryKey)

	isEUCountry := false
	_isEUCountry := h.Get(IsEUCountryKey)
	eu, err := strconv.ParseBool(_isEUCountry)
	if err == nil {
		isEUCountry = eu
	}

	city := h.Get(CityKey)
	region := h.Get(RegionKey)
	regionCode := h.Get(RegionCodeKey)

	latitude := 0.0
	_latitude := h.Get(LatitudeKey)
	l, err := strconv.ParseFloat(_latitude, 64)
	if err == nil {
		latitude = l
	}

	longitude := 0.0
	_longitude := h.Get(LongitudeKey)
	la, err := strconv.ParseFloat(_longitude, 64)
	if err == nil {
		longitude = la
	}

	postalCode := h.Get(PostalCodeKey)
	metroCode := h.Get(MetroCodeKey)
	speed := h.Get(SpeedKey)

	// Spec.Context.Network
	// ===============================
	bluetooth := false
	_bluetooth := h.Get(BluetoothKey)
	b, err := strconv.ParseBool(_bluetooth)
	if err != nil {
		bluetooth = b
	}

	cellular := false
	_cellular := h.Get(CellularKey)
	c, err := strconv.ParseBool(_cellular)
	if err != nil {
		cellular = c
	}

	wifi := false
	_wifi := h.Get(WifiKey)
	w, err := strconv.ParseBool(_wifi)
	if err != nil {
		wifi = w
	}

	carrier := h.Get(CarrierKey)
	asn := h.Get(AsnKey)
	asnOrganization := h.Get(AsnOrganizationKey)

	// Spec.Context.OS
	// ===============================
	osName := h.Get(OsNameKey)
	osVersion := h.Get(OsVersionKey)

	// Spec.SpecData.fields
	// ===============================
	var fieldMask []string
	if h.Get(FieldMask) != "" {
		fieldMask = strings.Split(h.Get(FieldMask), ",")
	}

	// Supporting
	// ===============================
	// contentType := headers["content-type"]
	// acceptEncoding := headers["accept-encoding"]
	// grpcAcceptEncoding := headers["grpc-accept-encoding"]

	s := specv2pb.Spec{
		SpecVersion: specVersion,
		MessageId:   messageId,
		RequestId:   requestId,
		SentAt:      sentAt,
		ReceivedAt:  receivedAt,
		CompletedAt: completedAt,
		SpecType:    specType,
		Principal: &specv2pb.SpecPrincipal{
			AnonymousId:    anonymousId,
			Type:           principalType,
			PrincipalId:    principalId,
			PrincipalEmail: principalEmail,
			ConnectionId:   connectionId,
			AuthRoles:      roles,
		},
		SpanContext: &specv2pb.SpanContext{
			TraceId:      traceId,
			SpanId:       spanId,
			ParentSpanId: parentSpanId,
			TraceFlags:   traceFlags,
		},
		Context: &specv2pb.SpecContext{
			OrganizationId:   organizationID,
			OrganizationSlug: organizationSlug,
			EcosystemId:      ecosystemID,
			EcosystemSlug:    ecosystemSlug,
			Jan:              jan,
			Ip:               ip,
			Locale:           locale,
			Timezone:         timezone,
			UserAgent:        userAgent,
			Validation: &specv2pb.SpecValidation{
				ValidateOnly: validateOnly,
			},
			Device: &specv2pb.SpecDevice{
				Id:            deviceId,
				Type:          deviceType,
				AdvertisingId: deviceAdvertisingId,
				Manufacturer:  deviceManufacturer,
				Model:         deviceModel,
				Name:          deviceName,
				Token:         deviceToken,
			},
			Location: &specv2pb.SpecLocation{
				Continent:   continent,
				Country:     country,
				IsEuCountry: isEUCountry,
				City:        city,
				Region:      region,
				RegionCode:  regionCode,
				Latitude:    latitude,
				Longitude:   longitude,
				PostalCode:  postalCode,
				MetroCode:   metroCode,
				Speed:       speed,
			},
			Network: &specv2pb.SpecNetwork{
				Bluetooth:       bluetooth,
				Cellular:        cellular,
				Wifi:            wifi,
				Carrier:         carrier,
				Asn:             asn,
				AsnOrganization: asnOrganization,
			},
			Os: &specv2pb.SpecOS{
				Name:    osName,
				Version: osVersion,
			},
		},
		Data: &anypb.Any{
			TypeUrl: "",
			Value:   nil,
		},
		SpecData: &specv2pb.SpecData{
			Configuration: nil,
			Data: &anypb.Any{
				TypeUrl: "",
				Value:   nil,
			},
			FieldMask: &fieldmaskpb.FieldMask{
				Paths: fieldMask,
			},
		},
	}

	return Factory{
		Spec:    &s,
		Headers: headers,
	}
}

// GenerateTraceID generate a trace id
func GenerateTraceID() string {
	var tid [16]byte
	if _, err := rand.Read(tid[:]); err != nil {
		apexlog.Error("failed to generate trace ID: " + err.Error())
	}
	return hex.EncodeToString(tid[:])
}

// GenerateSpanID generate a span ID
func GenerateSpanID() string {
	var sid [8]byte
	if _, err := rand.Read(sid[:]); err != nil {
		apexlog.Error("failed to generate span ID: " + err.Error())
	}
	return hex.EncodeToString(sid[:])
}
