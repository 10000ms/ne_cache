package node

import (
	"context"
	"google.golang.org/grpc"
	"hash/crc32"
	grpcService "ne_cache/server/grpc"
	"neko_server_go/utils"
	"sort"
	"strconv"
	"sync"
	"time"
)

type SingleNode struct {
	NodeAddr string `json:"node_addr"`
	Client   grpcService.NodeServiceClient
}

type nodeManage struct {
	RawNodeList  map[string]*SingleNode // key 是uuid
	NodeListLock sync.RWMutex
	HashMap      map[int]*SingleNode // key是hash位置
	HashMapLock  sync.RWMutex
	NodeMultiple int // 复制的node数量
	NodeHash     []int
}

func (s *SingleNode) NodeClientCheck() error {
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

func (s *SingleNode) NodeGet(key string) ([]byte, error) {
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

func (s *SingleNode) NodeSet(key string, value []byte, expire int64) error {
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

func (n *nodeManage) InitNodeManager() {
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

func (n *nodeManage) UpdateNodeList(node map[string]*SingleNode) {
	n.NodeListLock.Lock()
	defer n.NodeListLock.Unlock()
	n.RawNodeList = node
}

func (n *nodeManage) GetNode(key string) *SingleNode {
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

var NodeManager = nodeManage{
	RawNodeList:  make(map[string]*SingleNode),
	HashMap:      make(map[int]*SingleNode),
	NodeMultiple: 4, // 4倍节点数
	NodeHash:     make([]int, 0),
}
