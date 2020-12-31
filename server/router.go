package main

import (
	"ne_cache/server/handler"
	"neko_server_go"
	"neko_server_go/core"
	"neko_server_go/enum"
)

var Router = neko_server_go.Router{
	"/v1/cache/add/(?P<cache_key>.*)": core.CreateMethodsHandler(enum.HttpMethodsPost, handler.CacheAdd),
	"/v1/cache/get/(?P<cache_key>.*)": core.CreateMethodsHandler(enum.HttpMethodsGet, handler.CacheGet),
}
