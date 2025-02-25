# RPC

示例代码中的 RPC 工作流程如下

![](./rpc工作流程.png)

## 运行

先启动服务器

```bash
go run rpc_server.go
```

要请求斐波那契数列，则运行客户端

```bash
# 请求斐波那契数列的第 20 个数
go run rpc_client.go 20
```