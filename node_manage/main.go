package main

import (
	"ne_cache/common"
	"neko_server_go"
)

func main() {
	o := neko_server_go.Options{
		InitFunc: []func(){common.CheckTimer}, // 健康检查
	}
	neko_server_go.StartAPP(Settings, &Router, &o)
}
