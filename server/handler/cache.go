package handler

import (
	"fmt"
	"ne_cache/server"
	"neko_server_go"
	"neko_server_go/utils"
)

func CacheAdd(c *neko_server_go.Context, w neko_server_go.ResWriter) {
	err := c.Request.ParseForm() //解析参数，默认是不会解析的
	if err != nil {
		utils.LogError(err)
		return
	}

	p := *c.PathParams
	cacheKey := p["cache_key"]
	expire := c.Request.Form.Get("expire")
	cacheContent := c.Request.Body

	s := server.NodeManager.SetNode(cacheKey)
	err = s.NodeSet(cacheKey, cacheContent, expire)
	if err != nil {
		utils.LogError(err)
		return
	}

	_, err = fmt.Fprintf(w, "") //这个写入到w的是输出到客户端的
	if err != nil {
		utils.LogError(err)
		return
	}
}

func CacheGet(c *neko_server_go.Context, w neko_server_go.ResWriter) {
	err := c.Request.ParseForm() //解析参数，默认是不会解析的
	if err != nil {
		utils.LogError(err)
		return
	}

	p := *c.PathParams
	cacheKey := p["cache_key"]

	s := server.NodeManager.GetNode(cacheKey)
	cache, err := s.NodeGet(cacheKey)
	if err != nil {
		utils.LogError(err)
		return
	}

	_, err := fmt.Fprintf(w, cache) //这个写入到w的是输出到客户端的
	if err != nil {
		utils.LogError(err)
		return
	}
}
