module github.com/10000ms/ne_cache

go 1.14

require (
	ne_cache v0.0.0
	neko_server_go v0.0.0
	google.golang.org/grpc v1.34.0
	google.golang.org/grpc/examples v0.0.0-20201216234656-d79063fdde28
)

replace (
	neko_server_go v0.0.0 => github.com/10000ms/neko_server_go v0.0.3
	ne_cache v0.0.0 => ./
)
