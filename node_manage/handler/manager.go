package handler

import (
	"fmt"
	"neko_server_go"
	"neko_server_go/utils"

	"ne_cache/node_manage"
)

func NodeAdd(c *neko_server_go.Context, w neko_server_go.ResWriter) {

	nodeAddr := c.Request.PostForm.Get("node_addr")

	n := node_manage.SingleNode{
		NodeAddr: nodeAddr,
		Status: node_manage.NodeStatusUnKnow,
	}

	node_manage.NodeList = append(node_manage.NodeList, &n)

	_, err := fmt.Fprintf(w, "") //这个写入到w的是输出到客户端的
	if err != nil {
		utils.LogError(err)
		return
	}
}
