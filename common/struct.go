package common

import (
	grpcService "ne_cache/grpc"
)

type MangeSingleNode struct {
	NodeAddr string                       `json:"node_addr"`
	Status   NodeStatus                   `json:"status"`
	Client   grpcService.NodeHealthClient `json:"-"`
}

type ServerSingleNode struct {
	NodeAddr string `json:"node_addr"`
	Client   grpcService.NodeServiceClient
}

type ClientSettingsBase struct {
	// SettingsServerAddr 服务监听的地址
	SettingsServerAddr string
	// SettingsBufferSize TCP每次读大小
	SettingsBufferSize int
}

type RedisCommand string
