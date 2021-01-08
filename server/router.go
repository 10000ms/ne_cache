package main

import (
	"ne_cache/server/handler"
	"neko_server_go"
	"neko_server_go/core"
	"neko_server_go/enum"
)

var Router = neko_server_go.Router{
	"/v1/cache/set/(?P<cache_key>.*)": core.CreateMethodsHandler(enum.HttpMethodsPost, handler.CacheSet),
	"/v1/cache/get/(?P<cache_key>.*)": core.CreateMethodsHandler(enum.HttpMethodsGet, handler.CacheGet),
}
