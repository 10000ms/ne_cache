package server

import (
	"hash/crc32"
	grpcService "ne_cache/server/grpc"
	"sort"
	"strconv"
	"sync"
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
	NodeMultiple int
}

func (n *nodeManage) InitNodeManager() {
	n.NodeListLock.RLock()
	defer n.NodeListLock.RUnlock()
	for k, v := range n.RawNodeList {
		for i := 0; i < n.NodeMultiple; i++ {
			hash := int(crc32.ChecksumIEEE([]byte(strconv.Itoa(i) + k)))
			n.HashMap[hash] = v
		}
	}

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
	hash := int(crc32.ChecksumIEEE([]byte(key)))
	// TODO
	//idx := sort.Search(len(n.HashMap), func(i int) bool {
	//	return m.keys[i] >= hash
	//})
}

var NodeManager = nodeManage{
	NodeMultiple: 4,  // 4倍节点数
}
