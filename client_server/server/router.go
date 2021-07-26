package server

import (
	"ne_cache/client_server/common"
	"net"
)

func router() map[common.RedisCommand]func(common.SettingsBase, *Request, net.Conn) {
	r := map[common.RedisCommand]func(common.SettingsBase, *Request, net.Conn) {
		common.RedisCommandCommand: CommandCommandHandler,
		common.RedisCommandSet: CommandSetHandler,
		common.RedisCommandGet: CommandGetHandler,
	}
	return r
}

var Router = router()
