package handler

import (
	"context"
	"ne_cache/common"
	grpcService "ne_cache/grpc"
	"ne_cache/node/cache"
	"neko_server_go/utils"
	"time"
)

type NodeServer struct{}

func (h *NodeServer) SetValue(ctx context.Context, request *grpcService.SetValueRequest) (*grpcService.SetValueResponse, error) {
	s := cache.SingleCache{
		Key:     request.Key,
		Value:   request.Value,
		Expire:  request.Expire,
		SetTime: time.Now().UnixNano(),
	}

	cache.CacheManager.Add(request.Key, s)

	utils.LogDebug("node set value: ", request.Key)

	r := grpcService.SetValueResponse{
		Status: grpcService.SetValueResponse_OK,
	}
	return &r, nil
}

func (h *NodeServer) GetValue(ctx context.Context, request *grpcService.GetValueRequest) (*grpcService.GetValueResponse, error) {
	v, s := cache.CacheManager.Get(request.Key)

	utils.LogDebug("node get key: ", request.Key, ", value: ", common.LogToJSON(v))

	r := grpcService.GetValueResponse{
		Value:  v,
		Status: s,
	}
	return &r, nil
}
