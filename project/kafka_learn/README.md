# Kafka

在 Go 社区中目前有 3 个比较常用的 Kafka 客户端库，分别是：

- [segmentio/kafka-go](https://github.com/segmentio/kafka-go)：更加简单、更易用
- [IBM/sarama](https://github.com/IBM/sarama)：不太易用
- [confluentinc/confluent-kafka-go](https://github.com/confluentinc/confluent-kafka-go)：它是一个基于cgo的 [librdkafka](https://github.com/confluentinc/librdkafka) 包装，在项目中使用它会引入对C库的依赖。

## 本地运行

可以直接参考这里的 [docker-compose.yml](https://github.com/pudongping/polyglot-script-box/tree/master/kafka)

## 使用

kafka-go 提供了两套与 Kafka 交互的 API。

- 低级别（ low-level）：基于与 Kafka 服务器的原始网络连接实现。
- 高级别（high-level）：对于常用读写操作封装了一套更易用的API。

**通常建议直接使用高级别的交互API。**

## 注意

如果觉得英文官方文档不方便理解的话，也可以看这篇翻译的文档 [Go操作Kafka之kafka-go](https://www.liwenzhou.com/posts/Go/kafka-go/)