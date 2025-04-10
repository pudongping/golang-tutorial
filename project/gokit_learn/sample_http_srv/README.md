# 一个基本的 http 示例

## Go kit 关键概念

1. 传输层（Transport layer）

传输域绑定到具体的传输协议，如 HTTP 或 gRPC。在一个微服务可能支持一个或多个传输协议的世界中，这是非常强大的：你可以在单个微服务中支持原有的 HTTP API 和新增的 RPC 服务。

2. 端点层（Endpoint layer）

端点就像控制器上的动作/处理程序; 它是安全性和抗脆弱性逻辑的所在。如果实现两种传输(HTTP 和 gRPC) ，则可能有两种将请求发送到同一端点的方法。

3. 服务层（Service layer）

服务（指Go kit中的service层）是实现所有业务逻辑的地方。服务层通常将多个端点粘合在一起。在 Go kit 中，服务层通常被抽象为接口，这些接口的实现包含业务逻辑。Go kit 服务层应该努力遵守整洁架构或六边形架构。也就是说，业务逻辑不需要了解端点（尤其是传输域）概念：你的服务层不应该关心HTTP 头或 gRPC 错误代码。

**请求在第 1 层进入服务，向下流到第 3 层，响应则相反。**

## 测试

```bash
go run transports.go endpoints.go service.go
```

请求

```bash
curl -XPOST -d'{"a":1,"b":2}' localhost:8080/sum

curl -XPOST -d'{"a":"你好啊","b":"骚年"}' localhost:8080/concat
```
