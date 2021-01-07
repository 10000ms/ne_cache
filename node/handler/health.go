package handler

import (
	"context"
	grpcService "ne_cache/node/grpc"
	"neko_server_go/utils"
)

type HealthServer struct{}

func (h *HealthServer) NodeHealthCheck(ctx context.Context, request *grpcService.HealthCheckRequest) (*grpcService.HealthCheckResponse, error) {
	utils.LogInfo("节点健康检查")
	r := grpcService.HealthCheckResponse{
		Status: grpcService.HealthCheckResponse_SERVING,
	}
	return &r, nil
}
