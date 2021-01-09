# ne_cache


### 示例运行
```shell
bash ./docker_build.sh

docker-compose up -d
```


### grpc构建指令参考 
```

./protoc --go_out=./ ./node/grpc/health.proto 
./protoc --go-grpc_out=require_unimplemented_servers=false:./ ./node/grpc/health.proto

```
