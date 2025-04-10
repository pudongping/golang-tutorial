# 组织代码

当前程序同时对外提供 HTTP API 和 gRPC API

## 测试

### 启动

```bash
go run main.go transport.go endpoint.go service.go
```

### 测试 HTTP API

```bash
curl -XPOST -d'{"a":1,"b":2}' localhost:8080/sum

curl -XPOST -d'{"a":"你好啊","b":"骚年"}' localhost:8080/concat
```

### 测试 gRPC API

```bash
cd ../sample_grpc_srv && go test -v ./...
```