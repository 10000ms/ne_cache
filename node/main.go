package node

import (
	"Demo/example"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

// main 方法实现对 gRPC 接口的请求
func main() {
	addr := Settings["Server_addr"].(string)
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Can't connect: " + addr)
	}
	defer conn.Close()
	client := example.NewFormatDataClient(conn)
	resp,err := client.DoFormat(context.Background(), &example.Data{Text:"hello,world!"})
	if err != nil {
		log.Fatalln("Do Format error:" + err.Error())
	}
	log.Println(resp.Text)
}
