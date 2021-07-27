package common

import (
	"context"
	"google.golang.org/grpc"
	"hash/crc32"
	grpcService "ne_cache/grpc"
	"neko_server_go/utils"
	"sort"
	"strconv"
	"sync"
	"time"
)

type MangerNodeManage struct {
	List     map[string]*MangeSingleNode // key 是uuid
	NodeLock sync.RWMutex
}

func (n *MangerNodeManage) LiveNodeList() map[string]*MangeSingleNode {
	n.NodeLock.RLock()
	defer n.NodeLock.RUnlock()
	r := make(map[string]*MangeSingleNode)
	for k, v := range n.List {
		if v.Status == NodeStatusServing {
			r[k] = v
		}
	}
	return r
}

func (n *MangerNodeManage) AddNode(uuid string, nodeAddr string) {
	n.NodeLock.Lock()
	defer n.NodeLock.Unlock()
	s := MangeSingleNode{
		NodeAddr: nodeAddr,
		Status:   NodeStatusUnKnow,
	}
	n.List[uuid] = &s
}

func SingleCheck(node *MangeSingleNode) error {
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
	NodeList.NodeLock.Lock()
	defer NodeList.NodeLock.Unlock()
	for k, v := range NodeList.List {
		if v.Client == nil {
			// 没建立连接先建立连接
			conn, err := grpc.Dial(v.NodeAddr, grpc.WithInsecure())
			if err != nil {
				utils.LogError("建立连接失败，准备移除节点，error: ", err)
				delete(NodeList.List, k)
			}
			v.Client = grpcService.NewNodeHealthClient(conn)
		}
		err := SingleCheck(v)
		if err != nil {
			utils.LogError("健康检查失败，准备移除节点，error: ", err)
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

type ServerNodeManger struct {
	RawNodeList  map[string]*ServerSingleNode // key 是uuid
	NodeListLock sync.RWMutex
	HashMap      map[int]*ServerSingleNode // key是hash位置
	HashMapLock  sync.RWMutex
	NodeMultiple int // 复制的node数量
	NodeHash     []int
}

func (s *ServerSingleNode) NodeClientCheck() error {
	if s.Client == nil {
		// 没建立连接先建立连接
		conn, err := grpc.Dial(s.NodeAddr, grpc.WithInsecure())
		if err != nil {
			utils.LogError(err)
			return err
		}
		s.Client = grpcService.NewNodeServiceClient(conn)
	}
	return nil
}

func (s *ServerSingleNode) NodeGet(key string) ([]byte, error) {
	err := s.NodeClientCheck()
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10)
	defer cancel()
	resp, err := s.Client.GetValue(ctx, &grpcService.GetValueRequest{
		Key: key,
	})
	if err != nil {
		return nil, err
	} else if resp.Status != grpcService.GetValueResponse_OK {
		return nil, nil
	} else {
		return resp.Value, nil
	}
}

func (s *ServerSingleNode) NodeSet(key string, value []byte, expire int64) error {
	err := s.NodeClientCheck()
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10)
	defer cancel()
	_, err = s.Client.SetValue(ctx, &grpcService.SetValueRequest{
		Key:    key,
		Value:  value,
		Expire: expire,
	})
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (n *ServerNodeManger) InitNodeManager() {
	n.NodeListLock.RLock()
	defer n.NodeListLock.RUnlock()
	for k, v := range n.RawNodeList {
		for i := 0; i < n.NodeMultiple; i++ {
			hash := int(crc32.ChecksumIEEE([]byte(strconv.Itoa(i) + k)))
			n.HashMap[hash] = v
			n.NodeHash = append(n.NodeHash, hash)
		}
	}
	sort.Ints(n.NodeHash)

	n.HashMapLock.Lock()
	defer n.HashMapLock.Unlock()
}

func (n *ServerNodeManger) UpdateNodeList(node map[string]*ServerSingleNode) {
	n.NodeListLock.Lock()
	defer n.NodeListLock.Unlock()
	n.RawNodeList = node
}

func (n *ServerNodeManger) GetNode(key string) *ServerSingleNode {
	n.HashMapLock.RLock()
	defer n.HashMapLock.RUnlock()
	if len(n.HashMap) == 0 {
		return nil
	}
	hash := int(crc32.ChecksumIEEE([]byte(key)))
	idx := sort.Search(len(n.HashMap), func(i int) bool {
		return n.NodeHash[i] >= hash
	})
	return n.HashMap[n.NodeHash[idx%len(n.NodeHash)]]
}
