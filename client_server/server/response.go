package server

import (
	"ne_cache/client_server/common"
	"neko_server_go/utils"
	"net"
)

type Response struct {
	Conn net.Conn
}

func (r *Response) send(content []byte) {
	utils.LogDebug("Response Send: ", content)
	_, _ = r.Conn.Write(content)
}

func (r *Response) UnknownCommand(op common.RedisCommand) {
	r.send([]byte("-ERR unknown command " + string(op) + "\r\n"))
}

func (r *Response) OKStatus() {
	r.send([]byte("+OK\r\n"))
}
