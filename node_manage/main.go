package main

import (
	"ne_cache/node_manage/node"
	"neko_server_go"
)

func main() {
	o := neko_server_go.Options{
		InitFunc: []func(){node.CheckTimer},  // 健康检查
	}
	neko_server_go.StartAPP(Settings, &Router, &o)
}
