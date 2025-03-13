//nolint:revive
package sdkv2alphalib

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	specv2pb "libs/protobuf/go/protobuf/gen/platform/spec/v2"
	typev2pb "libs/protobuf/go/protobuf/gen/platform/type/v2"

	"connectrpc.com/connect"

	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"github.com/segmentio/ksuid"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// DefaultSpecVersion defines the default specification version to be used.
// DefaultConnectionId specifies the default connection identifier.
const (
	DefaultSpecVersion  = "v2"
	DefaultConnectionId = "corporate" //nolint:revive
)

// NatsMsgId is used for message deduplication in NATS.
// ApiKey represents the API key required to access the platform.
// SentAtKey specifies the timestamp sent by the client (not sanitized).
// AnonymousIdKey refers to the anonymous identifier in the context principal.
// PrincipalIdKey is the sanitized principal ID obtained from the authorization workload proxy.
// PrincipalEmailKey is the sanitized email of the principal.
// PrincipalTypeKey denotes the type of the principal (sanitized).
// ConnectionIdKey represents the connection identifier in the context.
// RequestIdKey is the request identifier sent by the client (not sanitized).
// B3ContextHeader represents the B3 propagation context header.
// B3DebugFlagKey indicates the B3 debug flag.
// B3TraceIDKey contains the B3 trace identifier.
// B3SpanIDKey holds the B3 span identifier.
// B3SampledKey indicates if the trace is sampled in the B3 context.
// B3ParentSpanIDKey represents the B3 parent span identifier.
// EcosystemSlug refers to the sanitized ecosystem slug from the edge cache.
// OrganizationSlug refers to the sanitized organization slug from the edge cache.
// WorkspaceSlug refers to the sanitized workspace slug from the edge cache.
// WorkspaceJurisdictionAreaNetworkKey indicates the workspace jurisdiction area network.
// IpKey specifies the IP information in the context.
// LocaleKey contains the locale information in the context.
// TimezoneKey specifies the timezone in the context.
// UserAgentKey denotes the user agent string.
// ValidateOnlyKey indicates if the request is for validation only.
// ChannelNameKey represents the name of the channel in the context.
// ChannelVersionKey represents the version of the channel in the context.
// DeviceIdKey contains the device identifier sent by the client (not sanitized).
// DeviceAdvertisingIdKey refers to the advertising identifier of the device.
// DeviceManufacturerKey specifies the manufacturer of the device.
// DeviceModelKey denotes the model of the device.
// DeviceNameKey represents the name of the device.
// DeviceTypeKey specifies the type of the device.
// DeviceTokenKey holds the device token.
// CityKey specifies the city information (sanitized, from edge cache).
// CountryKey contains the country information (sanitized, from edge cache).
// LatitudeKey provides the latitude information (sanitized, from edge cache).
// LongitudeKey provides the longitude information (sanitized, from edge cache).
// SpeedKey indicates the speed information (sanitized, from edge cache).
// BluetoothKey indicates Bluetooth usage information (not sanitized).
// CellularKey indicates cellular network information (not sanitized).
// WifiKey indicates Wi-Fi network usage information (not sanitized).
// CarrierKey specifies the carrier information in the network context.
// OsNameKey provides the name of the operating system (not sanitized).
// OsVersionKey specifies the version of the operating system (not sanitized).
// FieldMask contains the field mask data sent in the spec (not sanitized).
const (

	// NatsMsgId Message Deduplication
	NatsMsgId = "Nats-Msg-Id"

	// ApiKey API Key to access the platform
	ApiKey = "x-spec-apikey" //nolint:gosec
	// SentAtKey Spec
	// Not sanitized and allowed from the client
	SentAtKey = "x-spec-sent-at"

	// AnonymousIdKey Spec.Context.Principal
	AnonymousIdKey = "x-spec-anonymous-id"
	// PrincipalIdKey Sanitized comes from authorization workload proxy
	PrincipalIdKey    = "x-spec-principal-id"
	PrincipalEmailKey = "x-spec-principal-email"
	PrincipalTypeKey  = "x-spec-principal-type"
	ConnectionIdKey   = "x-spec-connection-id"

	// RequestIdKey Spec.SpanContext
	// Not sanitized and allowed from the client
	RequestIdKey      = "x-request-id"
	B3ContextHeader   = "b3"
	B3DebugFlagKey    = "x-b3-flags"
	B3TraceIDKey      = "x-b3-traceid"
	B3SpanIDKey       = "x-b3-spanid"
	B3SampledKey      = "x-b3-sampled"
	B3ParentSpanIDKey = "x-b3-parentspanid"

	// EcosystemSlug Spec.Context
	// OrganizationSlug Spec.Context
	// Sanitized comes from edge cache
	EcosystemSlug                       = "x-spec-ecosystem-slug"
	OrganizationSlug                    = "x-spec-organization-slug"
	WorkspaceSlug                       = "x-spec-workspace-slug"
	WorkspaceJurisdictionAreaNetworkKey = "x-spec-workspace-jan"
	IpKey                               = "x-spec-ip"
	LocaleKey                           = "x-spec-locale"
	TimezoneKey                         = "x-spec-timezone"
	UserAgentKey                        = "user-agent"

	// ValidateOnlyKey Spec.Context.Validation
	ValidateOnlyKey = "x-spec-validate-only"

	// ChannelNameKey Spec.Context.Channel
	ChannelNameKey    = "channel-name"
	ChannelVersionKey = "channel-version"

	// DeviceIdKey Spec.Context.Device
	// Not sanitized and allowed from the client
	DeviceIdKey            = "x-spec-device-id"
	DeviceAdvertisingIdKey = "x-spec-device-adv-id"
	DeviceManufacturerKey  = "x-spec-device-manufacturer"
	DeviceModelKey         = "x-spec-device-model"
	DeviceNameKey          = "x-spec-device-name"
	DeviceTypeKey          = "x-spec-device-type"
	DeviceTokenKey         = "x-spec-device-token" //nolint:gosec

	// CityKey Spec.Context.Location
	// Sanitized and comes from edge cache
	CityKey      = "x-spec-city"
	CountryKey   = "x-spec-country"
	LatitudeKey  = "x-spec-lat"
	LongitudeKey = "x-spec-long"
	SpeedKey     = "x-spec-speed"

	// BluetoothKey Spec.Context.Network
	// Not sanitized and allowed from the client
	BluetoothKey = "x-spec-bluetooth"
	CellularKey  = "x-spec-cellular"
	WifiKey      = "x-spec-wifi"
	CarrierKey   = "x-spec-carrier"

	// OsNameKey Spec.Context.OS
	// Not sanitized and allowed from the client
	OsNameKey    = "x-spec-os-name"
	OsVersionKey = "x-spec-os-version"

	// FieldMask Spec.SpecData.fields
	// Not sanitized and allowed from the client
	FieldMask = "x-spec-fieldmask"
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
func NewFactory(req connect.AnyRequest) Factory {
	h := req.Header()
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

	sentAt := &timestamppb.Timestamp{Seconds: 0, Nanos: 0}
	if _sentAt, ok := headers[SentAtKey]; ok {
		t, err := time.Parse(time.RFC3339, _sentAt)
		if err != nil {
			sentAt = timestamppb.New(t)
		}
	}

	receivedAt := timestamppb.Now()

	// Completed at is provided upstream by the implementing service/turbine
	completedAt := &timestamppb.Timestamp{Seconds: 0, Nanos: 0}

	fmt.Println(req.Spec().Procedure)
	specType := ""

	// Spec.SpecPrincipal
	// ===============================
	anonymousId := h.Get(AnonymousIdKey)
	principalId := h.Get(PrincipalIdKey)

	var principalType specv2pb.SpecPrincipalType
	_principalType := h.Get(PrincipalTypeKey)
	switch _principalType {
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

	// Span.Context
	// ===============================
	traceId := h.Get(B3TraceIDKey)
	spanId := h.Get(B3SpanIDKey)
	parentSpanId := h.Get(B3ParentSpanIDKey)
	traceFlags := h.Get(B3DebugFlagKey)

	// Spec.Context
	// ===============================
	ecosystemSlug := h.Get(EcosystemSlug)
	organizationSlug := h.Get(OrganizationSlug)
	workspaceSlug := h.Get(WorkspaceSlug)

	var workspaceJAN typev2pb.Jurisdiction
	_workspaceJAN := h.Get(WorkspaceJurisdictionAreaNetworkKey)
	switch _workspaceJAN {
	case "JURISDICTION_NA_US_1":
		workspaceJAN = typev2pb.Jurisdiction_JURISDICTION_NA_US_1
	case "JURISDICTION_GOV_US_1":
		workspaceJAN = typev2pb.Jurisdiction_JURISDICTION_GOV_US_1
	case "JURISDICTION_EU_DE_1":
		workspaceJAN = typev2pb.Jurisdiction_JURISDICTION_EU_DE_1
	case "JURISDICTION_GOV_EU_1  ":
		workspaceJAN = typev2pb.Jurisdiction_JURISDICTION_GOV_EU_1
	case "JURISDICTION_AS_CN_1":
		workspaceJAN = typev2pb.Jurisdiction_JURISDICTION_AS_CN_1
	case "JURISDICTION_SA_BR_1":
		workspaceJAN = typev2pb.Jurisdiction_JURISDICTION_SA_BR_1
	default:
		workspaceJAN = typev2pb.Jurisdiction_JURISDICTION_UNSPECIFIED
	}

	ip := h.Get(IpKey)
	locale := h.Get(LocaleKey)
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
	city := h.Get(CityKey)
	country := h.Get(CountryKey)
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
		},
		SpanContext: &specv2pb.SpanContext{
			TraceId:      traceId,
			SpanId:       spanId,
			ParentSpanId: parentSpanId,
			TraceFlags:   traceFlags,
		},
		Context: &specv2pb.SpecContext{
			EcosystemSlug:    ecosystemSlug,
			OrganizationSlug: organizationSlug,
			WorkspaceSlug:    workspaceSlug,
			WorkspaceJan:     workspaceJAN,
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
				City:      city,
				Country:   country,
				Latitude:  latitude,
				Longitude: longitude,
				Speed:     speed,
			},
			Network: &specv2pb.SpecNetwork{
				Bluetooth: bluetooth,
				Cellular:  cellular,
				Wifi:      wifi,
				Carrier:   carrier,
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
