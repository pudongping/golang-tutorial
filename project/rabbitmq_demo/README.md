# RabbitMQ

- [官网示例教程](https://www.rabbitmq.com/getstarted.html)
- [官方示例代码，涵盖多种客户端示例](https://github.com/rabbitmq/rabbitmq-tutorials)

## 安装

使用 docker 安装：[rabbitmq docker 官方文档地址](https://registry.hub.docker.com/_/rabbitmq/)

```shell
# 安装 rabbitmq 3.11

# 临时启用
docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.11-management

# 长久使用
docker run -it --name alex-rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.11-management

# 访问 web 管理界面
# 默认的账号和密码都为 guest
curl http://127.0.0.1:15672
```