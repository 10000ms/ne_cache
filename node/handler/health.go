package handler

import (
	"context"
	grpcService "ne_cache/grpc"
)

type HealthServer struct{}

func (h *HealthServer) NodeHealthCheck(ctx context.Context, request *grpcService.HealthCheckRequest) (*grpcService.HealthCheckResponse, error) {
	r := grpcService.HealthCheckResponse{
		Status: grpcService.HealthCheckResponse_SERVING,
	}
	return &r, nil
}
