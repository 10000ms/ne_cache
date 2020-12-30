.DEFAULT_GOAL := all

all: build-go

build-go: build-node build-node_manage build-server

build-node:
	CGO_ENABLED=0 go build -o ./bin/node ./node/main.go ./node/settings.go

build-node_manage:
	CGO_ENABLED=0 go build -o ./bin/nodemanage ./node_manage/main.go ./node_manage/settings.go ./node_manage/router.go

build-server:
	CGO_ENABLED=0 go build -o ./bin/server ./server/main.go ./server/settings.go ./server/router.go
