package server

import (
	"ne_cache/common"
	"net"
)

func initValidator() map[common.RedisCommand][]int {
	r := map[common.RedisCommand][]int{
		common.RedisCommandCommand: {1},
		common.RedisCommandSet:     {3, 5, 6},
		common.RedisCommandGet:     {2},
	}
	return r
}

var Validator = initValidator()

func Validate(settings common.ClientSettingsBase, request *Request, conn net.Conn) bool {
	if v, ok := Validator[request.Command]; ok {
		for _, l := range v {
			if len(request.Params) == l {
				return true
			}
		}
		resp := Response{
			Conn: conn,
		}
		resp.WrongNumberOfArguments(request.Command)
		return false
	} else {
		// 404
		resp := Response{
			Conn: conn,
		}
		resp.UnknownCommand(request.Command)
		return false
	}
}
