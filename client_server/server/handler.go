package server

import (
	"ne_cache/common"
	"neko_server_go/utils"
	"net"
	"time"
)

func CommandCommandHandler(settings common.ClientSettingsBase, request *Request, conn net.Conn) {
	utils.LogDebug("CommandCommandHandler")
	resp := Response{
		Conn: conn,
	}
	resp.OKStatus()
}

func CommandGetHandler(settings common.ClientSettingsBase, request *Request, conn net.Conn) {
	utils.LogDebug("CommandGetHandler")
	resp := Response{
		Conn: conn,
	}
	key := string(request.Params[1].Content)
	s := common.NodeManager.GetNode(key)
	cache := make([]byte, 0)
	var err error
	if s != nil {
		cache, err = s.NodeGet(key)
		if err != nil {
			utils.LogError(err)
			resp.InternalError()
		}
		if len(cache) == 0 {
			resp.NoSuchKey()
		} else {
			resp.SendBulkStrings(cache)
		}
	} else {
		resp.InternalError()
	}
}

func CommandSetHandler(settings common.ClientSettingsBase, request *Request, conn net.Conn) {
	utils.LogDebug("CommandSetHandler")
	resp := Response{
		Conn: conn,
	}
	var expire int64
	if len(request.Params) == 3 {
		expire = 0
	} else if len(request.Params) == 5 {
		expirationIdx := 3
		if string(request.Params[expirationIdx].Content) == "EX" {
			expire = int64(common.BytesStringToInt(request.Params[expirationIdx+1].Content)) * int64(time.Second)
		} else if string(request.Params[expirationIdx].Content) == "PX" {
			expire = int64(common.BytesStringToInt(request.Params[expirationIdx+1].Content)) * int64(time.Millisecond)
		} else {
			resp.ParamsError()
			return
		}
	} else {
		// TODO, 考虑是否支持 NX 和 XX
	}
	key := string(request.Params[1].Content)
	s := common.NodeManager.GetNode(key)
	var err error
	if s != nil {
		err = s.NodeSet(string(request.Params[1].Content), request.Params[2].Content, expire)
		if err != nil {
			utils.LogError(err)
			resp.InternalError()
		}
		resp.OKStatus()
	} else {
		resp.InternalError()
	}
}
