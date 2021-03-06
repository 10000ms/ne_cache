package handler

import (
	"bytes"
	"fmt"
	"ne_cache/server/node"
	"neko_server_go"
	"neko_server_go/utils"
	"strconv"
)

func CacheSet(c *neko_server_go.Context, w neko_server_go.ResWriter) {
	err := c.Request.ParseForm() //解析参数，默认是不会解析的
	if err != nil {
		utils.LogError(err)
		return
	}

	p := *c.PathParams
	cacheKey := p["cache_key"]
	expire := c.Request.Form.Get("expire")
	var expireInt64 int64
	if expire == "" {
		expireInt64 = 0
	} else {
		expireInt64, err = strconv.ParseInt(expire, 10, 64)
		if err != nil {
			utils.LogError(err)
			return
		}
	}
	cacheContent := c.Request.Body
	buf := new(bytes.Buffer)
	_ ,err = buf.ReadFrom(cacheContent)
	if err != nil {
		utils.LogError(err)
		return
	}

	s := node.NodeManager.GetNode(cacheKey)
	if s != nil {
		err = s.NodeSet(cacheKey, buf.Bytes(), expireInt64)
		if err != nil {
			utils.LogError(err)
			return
		}
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

	s := node.NodeManager.GetNode(cacheKey)

	var cache []byte
	if s != nil {
		cache, err = s.NodeGet(cacheKey)
		if err != nil {
			utils.LogError(err)
			return
		}
	}

	_, err = w.Write(cache) //这个写入到w的是输出到客户端的
	if err != nil {
		utils.LogError(err)
		return
	}
}
