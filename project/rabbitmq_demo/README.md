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

# 浏览器访问访问 web 管理界面
# 默认的账号和密码都为 guest
http://127.0.0.1:15672
```

## 使用 Go RabbitMQ 客户端

这里我们已经没有使用 `github.com/streadway/amqp` 库了，因为这个库现在已经没有维护了。这个库也推荐直接使用 `https://github.com/rabbitmq/amqp091-go` 库。

## 一些常用命令

命令 | 含义 |
--- | ---
rabbitmqctl list_queues | 查看队列
rabbitmqctl list_queues name messages_ready messages_unacknowledged | 打印忘记确认的队列信息

## 代码示例

- [一个最简单的生产者和消费者](./simple)
- [工作队列/任务队列](./task_worker)：消息确认、消息持久化、消息公平分发