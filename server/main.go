package main

import (
	"encoding/json"
	"io/ioutil"
	"ne_cache/server/node"
	"neko_server_go"
	"neko_server_go/utils"
	"net/http"
	"time"
)

func main() {
	o := neko_server_go.Options{
		InitFunc: []func(){GetNodeTimer},  // 定期获取node
	}
	neko_server_go.StartAPP(Settings, &Router, &o)
}

// 获取node节点
func GetNode() {
	serverAddr := "http://" + Settings["nodeManageAddr"].(string) + "/v1/node/info"
	resp, err := http.Get(serverAddr)
	if err != nil {
		utils.LogError(err)
		return
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
		var n map[string]*node.SingleNode
		err = json.Unmarshal(bodyBytes, &n)
		if err != nil {
			utils.LogError(err)
		}
		node.NodeManager.UpdateNodeList(n)
		node.NodeManager.InitNodeManager()
	} else {
		utils.LogError("request node manager error")
	}
}

func GetNodeTimer() {
	go func() {
		ticker := time.NewTicker(time.Second * 30)
		for {
			GetNode()
			<-ticker.C
		}
	}()
}
