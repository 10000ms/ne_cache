package client_server

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"ne_cache/client_server/node"
	"neko_server_go/utils"
	"net"
	"net/http"
	"time"
)

var nodeManagerAddr = flag.String("node_manager_addr", "127.0.0.1:8090", "node服务管理的地址")
var clientServerAddr = flag.String("port", SettingsServerAddr, "client服务器监听地址")


func main() {
	flag.Parse()
	// 先启动获取节点的定时任务
	GetNodeTimer()
	// 启动tcp server
	startServer()
}


//func handleConn(c net.Conn) {
//	defer func() {
//		_ = c.Close()
//	}()
//	remoteAddr := c.RemoteAddr()
//	fmt.Println(remoteAddr, " connect success")
//	// 接收数据
//	buf := make([]byte, 1024)
//	for {
//		n, err := c.Read(buf)
//		if err != nil {
//			fmt.Println("err:", err)
//			return
//		}
//		// windows会发送\r\n
//		if "exit" == string(buf[:n-2]) {
//			fmt.Println(remoteAddr, " 已断开")
//			return
//		}
//		fmt.Printf("from %s data:%s\n", remoteAddr, string(buf[:n])) // 发送数据
//		to := strings.ToUpper(string(buf[:n]))
//		_, _ = c.Write([]byte(to))
//	}
//}


func ProcessConn(c net.Conn) {
	// TODO 连接池管理功能，避免keep alive的连接太多
	defer func() {
		_ = c.Close()
	}()
	remoteAddr := c.RemoteAddr()
	utils.LogInfo(remoteAddr, " connect success")
	r := RequestHandler{
		Conn: c,
		EndConn: false,
	}
	r.Process()
}


func startServer() {
	l, err := net.Listen("tcp", *clientServerAddr)
	if err != nil {
		utils.LogError("listen error:", err)
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			utils.LogError("accept error:", err)
			break
		}
		// start a new goroutine to handle
		// the new connection.
		go ProcessConn(c)
	}
}

// GetNode 获取node节点
func GetNode() {
	serverAddr := "http://" + *nodeManagerAddr + "/v1/node/info"
	resp, err := http.Get(serverAddr)
	if err != nil {
		utils.LogError(err)
		panic(err)
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
		var n map[string]*node.SingleNode
		err = json.Unmarshal(bodyBytes, &n)
		if err != nil {
			utils.LogError(err)
		}
		node.NodeManager.UpdateNodeList(n)
		node.NodeManager.InitNodeManager()
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


