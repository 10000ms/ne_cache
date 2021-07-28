package handler

import (
	"encoding/json"
	"fmt"
	"ne_cache/common"
	"neko_server_go"
	"neko_server_go/utils"
)

func NodeAdd(c *neko_server_go.Context, w neko_server_go.ResWriter) {
	err := c.Request.ParseForm() //解析参数，默认是不会解析的
	if err != nil {
		utils.LogError(err)
		return
	}

	nodeAddr := c.Request.PostForm.Get("node_addr")
	uuid := c.Request.PostForm.Get("uuid")

	common.NodeList.AddNode(uuid, nodeAddr)

	utils.LogInfo(c.LogRequestID(), "注册node成功，node_addr: ", nodeAddr, " uuid: ", uuid)

	_, err = fmt.Fprintf(w, "") //这个写入到w的是输出到客户端的
	if err != nil {
		utils.LogError(err)
		return
	}
}

func AllNodeInfo(c *neko_server_go.Context, w neko_server_go.ResWriter) {
	l, LastUpdateTime := common.NodeList.LiveNodeList()
	r := map[string]interface{}{
		common.StringKeyLiveNodeList:   l,
		common.StringKeyLastUpdateTime: LastUpdateTime,
	}

	body, err := json.Marshal(r)
	if err != nil {
		utils.LogError(c.LogRequestID(), err, "Json Marshal失败， c: %v", r)
	}

	utils.LogDebug("response body: ", body)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	_, err = w.Write(body) //这个写入到w的是输出到客户端的
	if err != nil {
		utils.LogError(err)
		return
	}
}
