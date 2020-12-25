package handler

import (
	"context"
	"ne_cache/node"
	grpcService "ne_cache/node/grpc"
)

type NodeServer struct{}

func (h *NodeServer) SetValue(ctx context.Context, request *grpcService.SetValueRequest) (*grpcService.SetValueResponse, error) {
	s := node.SingleCache{
		Value:  request.Value,
		Expire: request.Expire,
	}

	node.CacheManager.Add(request.Key, s)

	r := grpcService.SetValueResponse{
		Status: grpcService.SetValueResponse_OK,
	}
	return &r, nil
}

func (h *NodeServer) GetValue(ctx context.Context, request *grpcService.GetValueRequest) (*grpcService.GetValueResponse, error) {
	v, s := node.CacheManager.Get(request.Key)

	r := grpcService.GetValueResponse{
		Value:  v,
		Status: s,
	}
	return &r, nil
}
