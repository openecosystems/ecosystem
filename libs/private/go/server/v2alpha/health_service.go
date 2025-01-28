package serverv2alpha

import (
	"context"

	"google.golang.org/grpc/codes"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
)

// Service represents a structure that implements health check functionalities for gRPC services.
type Service struct{}

// Check performs a health check for the service and returns the current health status in the response.
func (s *Service) Check(_ context.Context, _ *healthpb.HealthCheckRequest) (*healthpb.HealthCheckResponse, error) {
	return &healthpb.HealthCheckResponse{Status: healthpb.HealthCheckResponse_SERVING}, nil
}

// Watch streams health check responses to the client for the requested service. Not implemented in this implementation.
func (s *Service) Watch(_ *healthpb.HealthCheckRequest, _ healthpb.Health_WatchServer) error {
	return status.Error(codes.Unimplemented, "Watch is not implemented")
}
