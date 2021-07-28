# ne_cache

## 说明

ne_cache总共包含4个部分：缓存节点、节点管理、Web缓存服务、RESP缓存服务。使用`docker_build.sh`可以自动打包成一个镜像，根据指令在不同机器中启动不同部分

### 缓存节点

*启动指令*

```
/ne_cache/bin/node
```

负责缓存的实际储存，过期检查，容量检查。可多个节点部署，必须后于节点管理服务启动。

### 节点管理

*启动指令*

```
/ne_cache/bin/nodemanage
```

负责缓存节点的服务注册和健康监测。单节点部署，必须首先启动。

### Web缓存服务

*启动指令*

```
/ne_cache/bin/server
```

负责整个分布式缓存系统对外的缓存获取和缓存设置，会定期向节点管理服务获取最新的节点信息，并利用一致性哈希算法计算不同缓存节点管理的缓存段。可多个节点部署，必须后于节点管理服务启动。

### RESP缓存服务

*启动指令*

```
/ne_cache/bin/client_server
```

支持 Redis 协议规范 RESP (REdis Serialization Protocol)
的缓存服务，功能上面和Web缓存服务相似，负责整个分布式缓存系统对外的缓存获取和缓存设置，会定期向节点管理服务获取最新的节点信息，并利用一致性哈希算法计算不同缓存节点管理的缓存段。可多个节点部署，必须后于节点管理服务启动。

## API

### 设值

```
curl --location --request POST 'http://{{缓存服务地址}}/v1/cache/set/{{要缓存的key}}?expire={{缓存过期（milliseconds）}}' \
--header 'Content-Type: text/plain' \
--data-raw '{{需要缓存的信息}}'
```

### 取值

请求响应body即为缓存值，没有这个key则为空

```
curl --location --request GET 'http://{{缓存服务地址}}/v1/cache/get/{{要获取的key}}'
```

## 示例运行

```shell
bash ./docker_build.sh

docker-compose up -d
```

## grpc构建指令参考

```

./protoc --go_out=./ ./node/grpc/health.proto 
./protoc --go-grpc_out=require_unimplemented_servers=false:./ ./node/grpc/health.proto

```
