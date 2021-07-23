package common

type SettingsBase struct {
	// SettingsServerAddr 服务监听的地址
	SettingsServerAddr string
	// SettingsBufferSize TCP每次读大小
	SettingsBufferSize int
}

type RedisCommand string
