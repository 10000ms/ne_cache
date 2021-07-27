package server

import (
	"fmt"
	"ne_cache/common"
	"neko_server_go/utils"
	"net"
)

type Response struct {
	Conn net.Conn
}

func (r *Response) Send(content []byte) {
	utils.LogDebug("Response Send: ", content)
	_, _ = r.Conn.Write(content)
}

func (r *Response) SendBulkStrings(content []byte) {
	length := len(content)
	r.Send([]byte(fmt.Sprintf("$%d\r\n%s\r\n", length, content)))
}

func (r *Response) ParamsError() {
	r.Send([]byte("-ERR params error\r\n"))
}

func (r *Response) InternalError() {
	r.Send([]byte("-ERR internal error\r\n"))
}

func (r *Response) UnknownCommand(op common.RedisCommand) {
	r.Send([]byte("-ERR unknown command " + string(op) + "\r\n"))
}

func (r *Response) OKStatus() {
	r.Send([]byte("+OK\r\n"))
}

func (r *Response) WrongNumberOfArguments(op common.RedisCommand) {
	r.Send([]byte("-ERR wrong number of arguments for " + string(op) + " command\r\n"))
}

func (r *Response) NoSuchKey() {
	r.Send([]byte("-ERR no such key\r\n"))
}
