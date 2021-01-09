package main

import (
	"errors"
	"flag"
	"google.golang.org/grpc"
	"ne_cache/node/cache"
	grpcService "ne_cache/node/grpc"
	"ne_cache/node/handler"
	"neko_server_go/utils"
	"net"
	"net/http"
	"net/url"
	"strconv"
)

var nodeUUID = flag.String("uuid", "test", "节点的uuid")
var nodeAddr = flag.String("addr", "127.0.0.1", "节点的addr")
var nodePort = flag.String("port", "8080", "节点的port")
var nodeManagerAddr = flag.String("node_manager_addr", "127.0.0.1:8090", "node服务管理的地址")

func main() {
	utils.LogInfo("准备启动node")

	flag.Parse()
	// grpc 服务器启动
	rpcServer := grpc.NewServer()
	grpcService.RegisterNodeHealthServer(rpcServer, &handler.HealthServer{})
	grpcService.RegisterNodeServiceServer(rpcServer, &handler.NodeServer{})
	listener, err := net.Listen("tcp", ":"+*nodePort)
	if err != nil {
		utils.LogError("服务监听端口失败", err)
	}

	// 向服务发现注册服务
	RegisterNode()

	// 检测
	cache.Checker()

	// 启动rpc服务
	_ = rpcServer.Serve(listener)

}

// 注册node节点方法
func RegisterNode() {
	utils.LogInfo("准备注册node")
	var u, a string
	if *nodeUUID == "" {
		u = Settings["uuid"].(string)
	} else {
		u = *nodeUUID
	}
	if *nodeAddr == "" {
		a = GetLocalIp()
	} else {
		a = *nodeAddr
	}
	serverAddr := "http://" + *nodeManagerAddr + "/v1/node/add"
	resp, err := http.PostForm(serverAddr, url.Values{
		"node_addr": {a + ":" + *nodePort},
		"uuid":      {u},
	})
	if err != nil {
		utils.LogError(err)
		panic(err)
	}

	if resp.StatusCode != 200 {
		s := "注册失败，StatusCode: " + strconv.Itoa(resp.StatusCode)
		utils.LogError(err)
		panic(errors.New(s))
	}

	utils.LogInfo("注册node成功！")
}

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
