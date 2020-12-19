package node

import (
	"Demo/example"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
)

// main 方法实现对 gRPC 接口的请求
func main() {
	addr := Settings["server_addr"].(string)
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Can't connect: " + addr)
	}
	defer conn.Close()
	client := example.NewFormatDataClient(conn)
	resp, err := client.DoFormat(context.Background(), &example.Data{Text: "hello,world!"})
	if err != nil {
		log.Fatalln("Do Format error:" + err.Error())
	}
	log.Println(resp.Text)
}

func RegisterNode() {
	serverAddr := Settings["server_addr"].(string)
	nodePortTemp := Settings["node_port"].(string)
	//nodePort, err := strconv.ParseInt(nodePortTemp, 10, 64)
	body := "node_addr=" + GetLocalIp() + nodePortTemp
	resp, err := http.Post(serverAddr,
		"application/x-www-form-urlencoded",
		strings.NewReader(body))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
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
