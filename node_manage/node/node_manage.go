package node

import (
	"context"
	"google.golang.org/grpc"
	grpcService "ne_cache/node_manage/grpc"
	"neko_server_go/utils"
	"sync"
	"time"
)

type NodeStatus int64

const (
	NodeStatusUnKnow     NodeStatus = iota
	NodeStatusServing    NodeStatus = iota
	NodeStatusNotServing NodeStatus = iota
)

type SingleNode struct {
	NodeAddr string                       `json:"node_addr"`
	Status   NodeStatus                   `json:"status"`
	Client   grpcService.NodeHealthClient `json:"-"`
}

/*
key 是uuid
*/
var NodeList = make(map[string]*SingleNode)
var NodeLock sync.RWMutex

func LiveNodeList() map[string]*SingleNode {
	NodeLock.RLock()
	defer NodeLock.RUnlock()
	r := make(map[string]*SingleNode)
	for k, v := range NodeList {
		if v.Status == NodeStatusServing {
			r[k] = v
		}
	}
	return r
}

func SingleCheck(node *SingleNode) error {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*1)
	defer cancel()
	_, err := node.Client.NodeHealthCheck(ctx, &grpcService.HealthCheckRequest{})
	if err != nil {
		return err
	} else {
		return nil
	}
}

func NodeHealthCheck() {
	NodeLock.Lock()
	defer NodeLock.Unlock()
	for k, v := range NodeList {
		if v.Client != nil {
			// 没建立连接先建立连接
			conn, err := grpc.Dial(v.NodeAddr, grpc.WithInsecure())
			if err != nil {
				utils.LogError(err)
				delete(NodeList, k)
			}
			v.Client = grpcService.NewNodeHealthClient(conn)
		}
		err := SingleCheck(v)
		if err != nil {
			utils.LogError(err)
			delete(NodeList, k)
		}
		v.Status = NodeStatusServing
	}
}

func CheckTimer() {
	go func() {
		ticker := time.NewTicker(time.Second * 30)
		for {
			NodeHealthCheck()
			<-ticker.C
		}
	}()
}
