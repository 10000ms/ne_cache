package server

import (
	"ne_cache/client_server/common"
	"neko_server_go/utils"
	"net"
)

func CommandHandler(settings common.SettingsBase, request *Request, conn net.Conn) {
	utils.LogDebug("insssssadadsdasdsad")
	resp := Response{
		Conn: conn,
	}
	resp.OKStatus()
}
