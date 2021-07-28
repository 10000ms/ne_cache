package common

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"hash/crc32"
	"io/ioutil"
	grpcService "ne_cache/grpc"
	"neko_server_go/utils"
	"net/http"
	"sort"
	"strconv"
	"sync"
	"time"
)

type MangerNodeManage struct {
	List           map[string]*MangeSingleNode // key 是uuid
	LastUpdateTime int64
	NodeLock       sync.RWMutex
}

func (n *MangerNodeManage) LiveNodeList() (map[string]*MangeSingleNode, int64) {
	n.NodeLock.RLock()
	defer n.NodeLock.RUnlock()
	r := make(map[string]*MangeSingleNode)
	for k, v := range n.List {
		if v.Status == NodeStatusServing {
			r[k] = v
		}
	}
	return r, n.LastUpdateTime
}

func (n *MangerNodeManage) AddNode(uuid string, nodeAddr string) {
	n.NodeLock.Lock()
	defer n.NodeLock.Unlock()
	s := MangeSingleNode{
		NodeAddr: nodeAddr,
		Status:   NodeStatusUnKnow,
	}
	n.List[uuid] = &s
	n.LastUpdateTime = time.Now().UnixNano()
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
				NodeList.LastUpdateTime = time.Now().UnixNano()
			}
			v.Client = grpcService.NewNodeHealthClient(conn)
		}
		err := SingleCheck(v)
		if err != nil {
			utils.LogError("健康检查失败，准备移除节点，error: ", err)
			delete(NodeList.List, k)
			NodeList.LastUpdateTime = time.Now().UnixNano()
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
	RawNodeList    map[string]*ServerSingleNode // key 是uuid
	NodeListLock   sync.RWMutex
	HashMap        map[int]*ServerSingleNode // key是hash位置
	HashMapLock    sync.RWMutex
	NodeMultiple   int // 复制的node数量
	NodeHash       []int
	LastUpdateTime int64
	InitComplete   bool
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

	n.HashMapLock.Lock()
	defer n.HashMapLock.Unlock()

	// 清空
	n.NodeHash = make([]int, 0)
	n.HashMap = make(map[int]*ServerSingleNode)

	for k, v := range n.RawNodeList {
		for i := 0; i < n.NodeMultiple; i++ {
			hash := int(crc32.ChecksumIEEE([]byte(strconv.Itoa(i) + k)))
			n.HashMap[hash] = v
			n.NodeHash = append(n.NodeHash, hash)
		}
	}
	sort.Ints(n.NodeHash)

	if len(n.HashMap) > 0 {
		n.InitComplete = true
	}
}

func (n *ServerNodeManger) UpdateNodeList(nodeList map[string]*ServerSingleNode) {
	n.NodeListLock.Lock()
	defer n.NodeListLock.Unlock()

	for key, singleNode := range nodeList {
		if _, ok := n.RawNodeList[key]; !ok {
			n.RawNodeList[key] = singleNode
		}
	}

	for key, _ := range n.RawNodeList {
		if _, ok := nodeList[key]; !ok {
			delete(n.RawNodeList, key)
		}
	}
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

// GetNode 获取node节点
func GetNode(nodeManagerAddr string) {
	serverAddr := "http://" + nodeManagerAddr + "/v1/node/info"
	resp, err := http.Get(serverAddr)
	if err != nil {
		utils.LogError(err)
		panic(err)
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			utils.LogError(err)
		}
	}()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			utils.LogError(err)
		}
		var n map[string]interface{}
		err = json.Unmarshal(bodyBytes, &n)
		utils.LogDebug(LogToJSON(n))
		if err != nil {
			utils.LogError(err)
		}

		var (
			nodeList           map[string]*ServerSingleNode
			lastUpdateTime     int64
			ok                 bool
			tempNodeList       interface{}
			tempNodeListByte   []byte
			tempLastUpdateTime interface{}
		)
		tempNodeList, ok = n[StringKeyLiveNodeList]
		if !ok {
			utils.LogError("request node manager error")
			return
		}
		tempLastUpdateTime, ok = n[StringKeyLastUpdateTime]
		if !ok {
			utils.LogError("request node manager error")
			return
		}
		lastUpdateTime = int64(tempLastUpdateTime.(float64))

		tempNodeListByte, err = json.Marshal(tempNodeList)
		if err != nil {
			utils.LogError("json.Marshal(tempNodeList) error ", err)
			return
		}
		err = json.Unmarshal(tempNodeListByte, &nodeList)
		if err != nil {
			utils.LogError("json.Unmarshal(tempNodeListByte, &nodeList) error ", err)
			return
		}

		if lastUpdateTime != NodeManager.LastUpdateTime {
			utils.LogDebug(fmt.Sprintf("node list有更新，当前更新时间: %v, 服务器更新时间: %v, 进行更新", NodeManager.LastUpdateTime, lastUpdateTime))
			utils.LogDebug("新的node list: ", LogToJSON(nodeList), " ", LogToJSON(tempNodeList))
			NodeManager.UpdateNodeList(nodeList)
			NodeManager.InitNodeManager()
			NodeManager.LastUpdateTime = lastUpdateTime
		}

	} else {
		utils.LogError("request node manager error")
	}
}
