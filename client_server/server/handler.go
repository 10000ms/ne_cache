package server

import (
	"ne_cache/client_server/common"
	"ne_cache/client_server/node"
	"neko_server_go/utils"
	"net"
)

func CommandCommandHandler(settings common.SettingsBase, request *Request, conn net.Conn) {
	utils.LogDebug("CommandCommandHandler")
	resp := Response{
		Conn: conn,
	}
	resp.OKStatus()
}

func CommandGetHandler(settings common.SettingsBase, request *Request, conn net.Conn) {
	utils.LogDebug("CommandGetHandler")
	resp := Response{
		Conn: conn,
	}
	if len(request.Params) == 2 {
		key := string(request.Params[1].Content)
		s := node.NodeManager.GetNode(key)
		cache := make([]byte, 0)
		var err error
		if s != nil {
			cache, err = s.NodeGet(key)
			if err != nil {
				utils.LogError(err)
				resp.InternalError()
			}
			resp.Send(cache)
		} else {
			resp.InternalError()
		}

	} else {
		resp.WrongNumberOfArguments(request.Command)
	}
}

func CommandSetHandler(settings common.SettingsBase, request *Request, conn net.Conn) {
	// TODO, 支持完整的参数
	utils.LogDebug("CommandSetHandler")
	resp := Response{
		Conn: conn,
	}
	if len(request.Params) == 3 {
		key := string(request.Params[1].Content)
		s := node.NodeManager.GetNode(key)
		var err error
		if s != nil {
			err = s.NodeSet(string(request.Params[1].Content), request.Params[1].Content, 0)
			if err != nil {
				utils.LogError(err)
				resp.InternalError()
			}
			resp.OKStatus()
		} else {
			resp.InternalError()
		}

	} else {
		resp.WrongNumberOfArguments(request.Command)
	}
}
