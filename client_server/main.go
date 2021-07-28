package main

import (
	"flag"
	"ne_cache/client_server/server"
	"ne_cache/common"
	"neko_server_go/utils"
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

	// 等待NodeManager启动完成
	for {
		if common.NodeManager.InitComplete == true {
			break
		}
		utils.LogDebug("等待NodeManager启动完成中...")
		time.Sleep(1 * time.Second)
	}

	// 启动tcp server
	server.StartServer(Settings)
}

func GetNodeTimer() {
	go func() {
		ticker := time.NewTicker(time.Second * 30)
		for {
			common.GetNode(*nodeManagerAddr)
			<-ticker.C
		}
	}()
}
