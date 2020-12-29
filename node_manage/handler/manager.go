package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"neko_server_go"
	"neko_server_go/utils"

	"ne_cache/node_manage"
)

func NodeAdd(c *neko_server_go.Context, w neko_server_go.ResWriter) {
	err := c.Request.ParseForm() //解析参数，默认是不会解析的
	if err != nil {
		utils.LogError(err)
		return
	}

	nodeAddr := c.Request.PostForm.Get("node_addr")
	uuid := c.Request.PostForm.Get("uuid")

	n := node_manage.SingleNode{
		NodeAddr: nodeAddr,
		Status:   node_manage.NodeStatusUnKnow,
	}

	node_manage.NodeLock.Lock()
	defer node_manage.NodeLock.Unlock()
	node_manage.NodeList[uuid] = &n

	_, err = fmt.Fprintf(w, "") //这个写入到w的是输出到客户端的
	if err != nil {
		utils.LogError(err)
		return
	}
}

func AllNodeInfo(c *neko_server_go.Context, w neko_server_go.ResWriter) {
	buf := &bytes.Buffer{}
	encoder := json.NewEncoder(buf)
	l := node_manage.LiveNodeList()
	err := encoder.Encode(l)
	if err != nil {
		utils.LogError(err, "Json unMarshal失败， c:%v", l)
	}
	body := buf.Bytes()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	_, err = w.Write(body) //这个写入到w的是输出到客户端的
	if err != nil {
		utils.LogError(err)
		return
	}
}
