package main

import (
	"ne_cache/node_manage/node"
	"neko_server_go"
)

func main() {
	o := neko_server_go.Options{}
	neko_server_go.StartAPP(Settings, &Router, &o)

	node.CheckTimer()
}
