package node_manage

import (
    "neko_server_go"
)

func main() {
    o := neko_server_go.Options{}
    neko_server_go.StartAPP(Settings, &Router, &o)

    CheckTimer()
}
