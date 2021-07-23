package server

import (
	"ne_cache/client_server/common"
	"neko_server_go/utils"
	"net"
)

func ProcessConn(settings common.SettingsBase, c net.Conn) {
	// TODO 连接池管理功能，避免keep alive的连接太多
	defer func() {
		_ = c.Close()
	}()
	remoteAddr := c.RemoteAddr()
	utils.LogInfo(remoteAddr, " connect success")
	r := RequestHandler{
		Conn:    c,
		EndConn: false,
	}
	r.Process(settings)
}

func StartServer(settings common.SettingsBase) {
	l, err := net.Listen("tcp", settings.SettingsServerAddr)
	if err != nil {
		utils.LogError("listen error:", err)
		return
	}

	utils.LogInfo("start client server")

	for {
		c, err := l.Accept()
		if err != nil {
			utils.LogError("accept error:", err)
			break
		}
		// start a new goroutine to handle
		// the new connection.
		go ProcessConn(settings, c)
	}
}
