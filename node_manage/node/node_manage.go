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


type NodeManage struct{
	List     map[string]*SingleNode // key 是uuid
	NodeLock sync.RWMutex
}

func (n *NodeManage) LiveNodeList() map[string]*SingleNode {
	n.NodeLock.RLock()
	defer n.NodeLock.RUnlock()
	r := make(map[string]*SingleNode)
	for k, v := range n.List {
		if v.Status == NodeStatusServing {
			r[k] = v
		}
	}
	return r
}

func (n *NodeManage) AddNode(uuid string, nodeAddr string) {
	n.NodeLock.Lock()
	defer n.NodeLock.Unlock()
	s := SingleNode{
		NodeAddr: nodeAddr,
		Status:   NodeStatusUnKnow,
	}
	n.List[uuid] = &s
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

var NodeList = NodeManage{
	List: make(map[string]*SingleNode),
}

func NodeHealthCheck() {
	utils.LogInfo("进行节点健康检查")
	NodeList.NodeLock.Lock()
	defer NodeList.NodeLock.Unlock()
	for k, v := range NodeList.List {
		if v.Client != nil {
			// 没建立连接先建立连接
			conn, err := grpc.Dial(v.NodeAddr, grpc.WithInsecure())
			if err != nil {
				utils.LogError(err)
				delete(NodeList.List, k)
			}
			v.Client = grpcService.NewNodeHealthClient(conn)
		}
		err := SingleCheck(v)
		if err != nil {
			utils.LogError(err)
			delete(NodeList.List, k)
		}
		v.Status = NodeStatusServing
	}
}

func CheckTimer() {
	go func() {
		ticker := time.NewTicker(time.Second * 3)
		for {
			NodeHealthCheck()
			<-ticker.C
		}
	}()
}
