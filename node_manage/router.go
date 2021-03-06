package main

import (
	"ne_cache/node_manage/handler"
	"neko_server_go"
	"neko_server_go/core"
	"neko_server_go/enum"
)

var Router = neko_server_go.Router{
	"/v1/node/add":  core.CreateMethodsHandler(enum.HttpMethodsPost, handler.NodeAdd),
	"/v1/node/info": core.CreateMethodsHandler(enum.HttpMethodsGet, handler.AllNodeInfo),
}
