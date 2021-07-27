package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"ne_cache/client_server/server"
	"ne_cache/common"
	"neko_server_go/utils"
	"net/http"
	"time"
)

var nodeManagerAddr = flag.String("node_manager_addr", "127.0.0.1:8090", "node服务管理的地址")
var clientServerAddr = flag.String("addr", Settings.SettingsServerAddr, "client服务器监听地址")

func main() {
	flag.Parse()

	// 先启动获取节点的定时任务
	GetNodeTimer()

	// 处理一下settings
	if Settings.SettingsServerAddr != *clientServerAddr {
		Settings.SettingsServerAddr = *clientServerAddr
	}

	// 启动tcp server
	server.StartServer(Settings)
}

// GetNode 获取node节点
func GetNode() {
	serverAddr := "http://" + *nodeManagerAddr + "/v1/node/info"
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
		var n map[string]*common.ServerSingleNode
		err = json.Unmarshal(bodyBytes, &n)
		if err != nil {
			utils.LogError(err)
		}
		common.NodeManager.UpdateNodeList(n)
		common.NodeManager.InitNodeManager()
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
