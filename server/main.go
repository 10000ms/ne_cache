package main

import (
	"flag"
	"ne_cache/common"
	"neko_server_go"
	"time"
)

var nodeManagerAddr = flag.String("node_manager_addr", "127.0.0.1:8090", "node服务管理的地址")

func main() {
	flag.Parse()
	o := neko_server_go.Options{
		InitFunc: []func(){GetNodeTimer}, // 定期获取node
	}
	neko_server_go.StartAPP(Settings, &Router, &o)
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
