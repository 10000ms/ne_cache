package server

import (
	"ne_cache/common"
	"net"
)

func initRouter() map[common.RedisCommand]func(common.ClientSettingsBase, *Request, net.Conn) {
	r := map[common.RedisCommand]func(common.ClientSettingsBase, *Request, net.Conn){
		common.RedisCommandCommand: CommandCommandHandler,
		common.RedisCommandSet:     CommandSetHandler,
		common.RedisCommandGet:     CommandGetHandler,
	}
	return r
}

var Router = initRouter()
