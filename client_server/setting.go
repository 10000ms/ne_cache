package main

import (
	"ne_cache/client_server/common"
)

var Settings = common.SettingsBase{
	SettingsServerAddr: ":6380",
	SettingsBufferSize: 256,
}
