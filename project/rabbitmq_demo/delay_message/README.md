# RabbitMQ 延迟队列

## 第一种：传统方法，基于死信交换机（死信队列 + TTL）

对于基于死信交换机的方法，需要创建两个队列：一个用于延迟存储，消息过期后转发到死信交换机，另一个是实际消费的队列。

测试步骤：

1. 启动消费者

```bash
go run consumer_dlx.go
```

2. 启动生产者

```bash
go run producer_dlx.go
```

> 但是这种方式有一个非常严重的弊端：
> 因为队列参数的不可变性，RabbitMQ 的队列参数（如 x-message-ttl）在首次声明后不可更改。若尝试用不同参数重新声明同名队列，会触发 `PRECONDITION_FAILED` 错误。
这也就意味着，一旦队列的配置发生变化，比如 TTL 时间变化，那么就需要重新创建队列，否则无法生效。也就是说，死信队列方案依赖队列级别的 TTL，所有消息共享相同延迟时间，无法实现消息级动态延迟。
> 解决方案可以参考 `producer_dlx1.go` 但是还是不能解决 **先进先出** 的问题。（msg1 预计延迟50s被消费，msg2 预计延迟5s被消费，msg1 先进入队列，理应 msg2 先被消费，但是这里依旧 msg1 先被消费）


## 第二种：基于官方延迟插件 [rabbitmq-delayed-message-exchange](https://github.com/rabbitmq/rabbitmq-delayed-message-exchange)

使用插件的话，则可以直接声明延迟类型的交换机，并发布消息时设置延迟头。

官方 docker 镜像 `rabbitmq:3.11-management` 默认不包含延迟插件，我们需要整一个 `rabbitmq_delayed_message_exchange` 延迟插件。

先尝试使用以下命令来开启插件

```bash
# 先启用插件（需要管理员权限）
rabbitmq-plugins enable rabbitmq_delayed_message_exchange
```

如果报错，说明插件不存在，需要手动安装。

### 方法一：使用预装插件的镜像

```bash
# 停止并删除旧容器
docker stop rabbitmq && docker rm rabbitmq

# 使用带插件的镜像
docker run -d --name rabbitmq \
  -p 5672:5672 -p 15672:15672 \
  pivotalrabbitmq/rabbitmq-delayed-message-exchange
```

### 方法二：手动安装插件

```bash
# 进入容器
docker exec -it rabbitmq bash

# 容器内执行以下命令
apt-get update && apt-get install -y wget
wget https://github.com/rabbitmq/rabbitmq-delayed-message-exchange/releases/download/3.11.1/rabbitmq_delayed_message_exchange-3.11.1.ez
mv rabbitmq_delayed_message_exchange-3.11.1.ez $RABBITMQ_HOME/plugins/

# 退出容器后重启
docker restart rabbitmq
```

### 方法三：通过 Dockerfile 定制镜像

Dockerfile 内容如下：

```Dockerfile
FROM rabbitmq:3.11-management

RUN set -eux; \
    apt-get update; \
    apt-get install -y --no-install-recommends wget; \
    wget https://github.com/rabbitmq/rabbitmq-delayed-message-exchange/releases/download/3.11.1/rabbitmq_delayed_message_exchange-3.11.1.ez -P /plugins; \
    chown rabbitmq:rabbitmq /plugins/rabbitmq_delayed_message_exchange-3.11.1.ez; \
    rabbitmq-plugins enable --offline rabbitmq_delayed_message_exchange
```

构建并运行：

```bash
docker build -t alex-rabbitmq .
docker run -d -p 5672:5672 -p 15672:15672 alex-rabbitmq
```

### 验证安装

```bash
# 查看已启用插件列表
docker exec rabbitmq rabbitmq-plugins list

# 或者直接进入容器之后，执行
rabbitmq-plugins list | grep rabbitmq_delayed_message_exchange

# 应显示：
#[E*] rabbitmq_delayed_message_exchange 3.11.1
# 则表示插件已启用
```

### 测试

1. 启动生产者

```bash
go run producer_delay_plugin.go
```

2. 启动消费者

```bash
go run consumer_delay_plugin.go
```

## 两种方案对比

| 特性       | 死信队列方案           | 插件方案                   |
|------------|------------------------|----------------------------|
| 安装要求   | 无需插件               | 需要安装延迟插件           |
| 延迟精度   | 队列级别TTL            | 消息级别延迟               |
| 灵活性     | 所有消息相同延迟       | 每条消息可设不同延迟       |
| 消息顺序   | 先进先出               | 根据到期时间排序           |
| 性能影响   | 高（定期扫描队列）     | 低（时间轮算法）           |
| 适用场景   | 固定延迟批量消息       | 需要动态调整延迟的场景     |

推荐**优先使用官方插件方案**，原因如下：
- 灵活性：支持消息级延迟，每条消息可独立设置TTL，适应动态业务需求（如不同订单有不同的超时时间）。
- 性能优势：避免死信队列的消息阻塞问题，处理效率更高。
- 功能完整性：原生支持延迟队列特性，减少代码冗余（如无需维护多个队列）。