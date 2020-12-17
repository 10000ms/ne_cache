package handler

import (
	"fmt"
	"neko_server_go"
	"neko_server_go/utils"
)

func CacheAdd(c *neko_server_go.Context, w neko_server_go.ResWriter) {
	p := *c.PathParams
	cacheKey := p["cache_key"]

	//cacheContent := c.Request.Body

	_, err := fmt.Fprintf(w, cacheKey) //这个写入到w的是输出到客户端的
	if err != nil {
		utils.LogError(err)
		return
	}
}

func CacheGet(c *neko_server_go.Context, w neko_server_go.ResWriter) {
	p := *c.PathParams
	cacheKey := p["cache_key"]

	_, err := fmt.Fprintf(w, cacheKey) //这个写入到w的是输出到客户端的
	if err != nil {
		utils.LogError(err)
		return
	}
}
