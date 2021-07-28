package common

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func InitTestServerNodeManger() {
	l := map[string]*ServerSingleNode{
		"d5990db7-a7ca-48f8-99df-c33b90948906": {NodeAddr: "1"},
		"5787a75a-a8d9-4fb2-9bf8-e50551f9b1de": {NodeAddr: "2"},
		"50961f9f-a731-4766-a38e-d81e79860fce": {NodeAddr: "3"},
		"fe828e2a-e47a-45c2-83c7-ab00e1bb62c3": {NodeAddr: "4"},
		"842078b8-c59b-4557-870d-516fa27bca8a": {NodeAddr: "5"},
		"aa4413b2-3369-4dc5-beef-12a6c6525522": {NodeAddr: "6"},
		"8d1a1d1b-412e-4162-bd5d-891527bc9026": {NodeAddr: "7"},
		"cab231e0-0612-4db1-90cf-21a277f782a5": {NodeAddr: "8"},
		"697daf19-6134-4fdd-8f27-94aabad995b5": {NodeAddr: "9"},
		"69705f19-6134-4fdd-8f27-94aabad995b5": {NodeAddr: "10"},
		"69705f19-6134-4fdd-8f27-94a04ad995b5": {NodeAddr: "11"},
	}
	NodeManager.UpdateNodeList(l)
	NodeManager.InitNodeManager()
}

func MultiTestGetNode(t *testing.T, key string) {
	var nodeAddr string
	var RawNodeList map[string]*ServerSingleNode
	var HashMap map[int]*ServerSingleNode
	var NodeHash []int
	for i := 0; i < 100; i++ {
		// 每次重新设值保证设值步骤也是一致的
		InitTestServerNodeManger()
		n := NodeManager.GetNode(key)
		assert.NotEmpty(t, n, fmt.Sprintf("%s 拿不到node, 第 %v 次", key, i+1))
		if nodeAddr == "" {
			nodeAddr = n.NodeAddr
		}
		if RawNodeList == nil {
			RawNodeList = NodeManager.RawNodeList
		}
		if HashMap == nil {
			HashMap = NodeManager.HashMap
		}
		if NodeHash == nil {
			NodeHash = NodeManager.NodeHash
		}
		if nodeAddr != n.NodeAddr {
			assert.Equal(t, nodeAddr, n.NodeAddr, fmt.Sprintf("%s 拿不到的node不对，之前的node: %s, 这次的node: %s, 第 %v 次", key, nodeAddr, n.NodeAddr, i+1))

		}
		if RawNodeList != nil {
			assert.Equal(t, RawNodeList, NodeManager.RawNodeList)
		}
		if HashMap != nil {
			assert.Equal(t, HashMap, NodeManager.HashMap)
		}
		if NodeHash != nil {
			assert.Equal(t, NodeHash, NodeManager.NodeHash)
		}
	}
}

func TestGetNode(t *testing.T) {
	InitTestServerNodeManger()
	testKey1 := "aaa"
	MultiTestGetNode(t, testKey1)
	testKey2 := "bbb"
	MultiTestGetNode(t, testKey2)
	testKey3 := "ccc"
	MultiTestGetNode(t, testKey3)
	testKey4 := "ddd"
	MultiTestGetNode(t, testKey4)
	testKey5 := "eee"
	MultiTestGetNode(t, testKey5)
	testKey6 := "fff"
	MultiTestGetNode(t, testKey6)
}
