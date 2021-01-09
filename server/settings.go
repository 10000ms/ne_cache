package main

import (
	"errors"
	"neko_server_go"
	"path/filepath"
	"runtime"
)

func getPath() string {
	_, str, _, ok := runtime.Caller(1)
	if !ok {
		panic(errors.New("get path error"))
	}
	r, err := filepath.Abs(filepath.Dir(str))
	if err != nil {
		panic(errors.New("get filepath Abs error"))
	}
	return r
}

var Settings = neko_server_go.Setting{
	"ServiceName": "ne_cache_proxy_server",
	"nodeManageAddr": "127.0.0.1:8090",
	"Host":        "",
	"Port":        "8080",
	"Debug":       true,
	"Path":        getPath(),
}
