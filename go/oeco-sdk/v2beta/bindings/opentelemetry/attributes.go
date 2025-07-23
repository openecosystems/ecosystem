package opentelemetryv1

const (
	AttrServiceName    = "service.name"
	AttrServiceVersion = "service.version"
	AttrEnvironment    = "deployment.environment"

	AttrRPCSystem     = "rpc.system"
	AttrRPCService    = "rpc.service"
	AttrRPCMethod     = "rpc.method"
	AttrRPCStatusCode = "rpc.grpc.status_code"

	AttrHTTPMethod     = "http.method"
	AttrHTTPStatusCode = "http.status_code"
	AttrHTTPTarget     = "http.target"

	AttrNetPeerIP    = "net.peer.ip"
	AttrNetPeerPort  = "net.peer.port"
	AttrNetTransport = "net.transport"

	AttrExceptionType       = "exception.type"
	AttrExceptionMessage    = "exception.message"
	AttrExceptionStacktrace = "exception.stacktrace"

	AttrMessagingSystem      = "messaging.system"
	AttrMessagingOperation   = "messaging.operation"
	AttrMessagingDestination = "messaging.destination"

	AttrOrganizationID   = "organization.id"
	AttrEcosystemID      = "ecosystem.id"
	AttrUserID           = "user.id"
	AttrConnectorName    = "connector.name"
	AttrConnectorVersion = "connector.version"
	AttrAuthMethod       = "auth.method"
	AttrAuthSuccess      = "auth.success"
	AttrWorkflowName     = "workflow.name"
	AttrWorkflowRunID    = "workflow.run_id"
)
