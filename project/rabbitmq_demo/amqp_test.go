package rabbitmq_demo

import (
	"testing"
)

// 最简单的生产者 go test -run TestProducerMessage
func TestProducerMessage(t *testing.T) {
	// 如果需要检查队列，可以使用 `rabbitmqctl list_queues` 命令来查看
	ProducerMessage()
}

// 最简单的消费者 go test -run TestConsumerMessage
func TestConsumerMessage(t *testing.T) {
	ConsumerMessage()
}

func TestCallDelayMessage(t *testing.T) {
	CallDelayMessage()
}
