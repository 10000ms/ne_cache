package node

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	grpcService "ne_cache/node/grpc"
	"neko_server_go/utils"
	"net"
	"net/http"
	"net/url"
)


type HealthServer struct {

}

func(h *HealthServer) NodeHealthCheck(ctx context.Context, request *grpcService.HealthCheckRequest) (*grpcService.HealthCheckResponse, error) {

}

// main 方法实现对 gRPC 接口的请求
func main() {
	// 1. new一个grpc的server
	rpcServer := grpc.NewServer()

	// 2. 将刚刚我们新建的ProdService注册进去
	grpcService.RegisterNodeHealthServer(rpcServer, &HealthServer{})

	// 3. 新建一个listener，以tcp方式监听8082端口
	listener, err := net.Listen("tcp", ":" + Settings["node_port"].(string))
	if err != nil {
		log.Fatal("服务监听端口失败", err)
	}

	// 4. 运行rpcServer，传入listener
	_ = rpcServer.Serve(listener)
}

func RegisterNode() {
	serverAddr := Settings["server_addr"].(string)
	nodePort := Settings["node_port"].(string)
	_, err := http.PostForm(serverAddr, url.Values{
		"node_addr": {GetLocalIp() + nodePort},
		"uuid":      {Settings["uuid"].(string)},
	})
	if err != nil {
		utils.LogError(err)
	}

}

//获取本机ip
func GetLocalIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("get local ip failed")
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
