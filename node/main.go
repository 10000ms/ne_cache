package node

import (
	"flag"
	"google.golang.org/grpc"
	"ne_cache/node/cache"
	grpcService "ne_cache/node/grpc"
	"ne_cache/node/handler"
	"neko_server_go/utils"
	"net"
	"net/http"
	"net/url"
)

var nodeUUID = flag.String("uuid", "", "节点的uuid")

func main() {
	// grpc 服务器启动
	rpcServer := grpc.NewServer()
	grpcService.RegisterNodeHealthServer(rpcServer, &handler.HealthServer{})
	grpcService.RegisterNodeServiceServer(rpcServer, &handler.NodeServer{})
	listener, err := net.Listen("tcp", ":"+Settings["node_port"].(string))
	if err != nil {
		utils.LogError("服务监听端口失败", err)
	}
	_ = rpcServer.Serve(listener)

	// 向服务发现注册服务
	RegisterNode()

	// 检测
	cache.Checker()
}

// 注册node节点方法
func RegisterNode() {
	flag.Parse()
	var u string
	if *nodeUUID == "" {
		u = Settings["uuid"].(string)
	} else {
		u = *nodeUUID
	}
	serverAddr := Settings["server_addr"].(string)
	nodePort := Settings["node_port"].(string)
	_, err := http.PostForm(serverAddr, url.Values{
		"node_addr": {GetLocalIp() + nodePort},
		"uuid":      {u},
	})
	if err != nil {
		utils.LogError(err)
	}

}

// TODO 需要指定，docker内获取不到真实地址
//获取本机ip
func GetLocalIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		utils.LogError("get local ip failed")
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
