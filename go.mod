module github.com/10000ms/ne_cache

go 1.14

require (
	github.com/shirou/gopsutil v3.20.11+incompatible // indirect
	google.golang.org/grpc v1.34.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.0.1 // indirect
	ne_cache v0.0.0
	neko_server_go v0.0.0
)

replace (
	ne_cache v0.0.0 => ./
	neko_server_go v0.0.0 => github.com/10000ms/neko_server_go v0.0.3
)
