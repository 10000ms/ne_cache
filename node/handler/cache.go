package handler

import (
	"context"
	grpcService "ne_cache/grpc"
	"ne_cache/node/cache"
)

type NodeServer struct{}

func (h *NodeServer) SetValue(ctx context.Context, request *grpcService.SetValueRequest) (*grpcService.SetValueResponse, error) {
	s := cache.SingleCache{
		Key:    request.Key,
		Value:  request.Value,
		Expire: request.Expire,
	}

	cache.CacheManager.Add(request.Key, s)

	r := grpcService.SetValueResponse{
		Status: grpcService.SetValueResponse_OK,
	}
	return &r, nil
}

func (h *NodeServer) GetValue(ctx context.Context, request *grpcService.GetValueRequest) (*grpcService.GetValueResponse, error) {
	v, s := cache.CacheManager.Get(request.Key)

	r := grpcService.GetValueResponse{
		Value:  v,
		Status: s,
	}
	return &r, nil
}
