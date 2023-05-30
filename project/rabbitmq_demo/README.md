# RabbitMQ

- [官网使用示例](https://www.rabbitmq.com/getstarted.html)

## 安装

使用 docker 安装

```shell
# 安装 rabbitmq 3.11

# 临时启用
docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.11-management

# 长久使用
docker run -it --name alex-rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.11-management

# 访问 web 管理界面
curl http://127.0.0.1:15672
```