package server

import (
	"encoding/json"
	"io/ioutil"
	"neko_server_go"
	"neko_server_go/utils"
	"net/http"
	"time"
)

func main() {
	o := neko_server_go.Options{}
	neko_server_go.StartAPP(Settings, &Router, &o)

	// 定期获取node
	GetNodeTimer()
}

// 获取node节点
func GetNode() {
	serverAddr := Settings["node_manage_addr"].(string)
	resp, err := http.Get(serverAddr)
	if err != nil {
		utils.LogError(err)
		return
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			utils.LogError(err)
		}
	}()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			utils.LogError(err)
		}
		var node map[string]*SingleNode
		err = json.Unmarshal(bodyBytes, &node)
		if err != nil {
			utils.LogError(err)
		}
		NodeManager.UpdateNodeList(node)
		NodeManager.InitNodeManager()
	} else {
		utils.LogError("request node manager error")
	}
}

func GetNodeTimer() {
	go func() {
		ticker := time.NewTicker(time.Second * 30)
		for {
			GetNode()
			<-ticker.C
		}
	}()
}

