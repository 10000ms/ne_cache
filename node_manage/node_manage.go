package node_manage

import (
	"time"
)

type NodeStatus int64

const (
	NodeStatusUnKnow NodeStatus = iota
	NodeStatusServing NodeStatus = iota
	NodeStatusNotServing NodeStatus = iota
)

type SingleNode struct {
	NodeAddr string  `json:"node_addr"`
	Status NodeStatus  `json:"status"`
	Uuid string  `json:"uuid"`
}


var NodeList = make([]*SingleNode, 0)


func NodeHealthCheck() {
	for _, n := range NodeList {
		if n != nil {
			// 进行检查
		}
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
