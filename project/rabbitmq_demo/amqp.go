package rabbitmq_demo

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func ProducerMessage() {
	// 创建 RabbitMQ 连接
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("无法连接到 RabbitMQ:", err)
	}
	defer conn.Close()

	// 创建通道
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("无法打开通道:", err)
	}
	defer ch.Close()

	// 声明队列
	queueName := "my_queue"
	_, err = ch.QueueDeclare(
		queueName, // 队列名称
		false,     // 持久化
		false,     // 自动删除
		false,     // 排他性
		false,     // 等待服务器确认
		nil,       // 参数
	)
	if err != nil {
		log.Fatal("无法声明队列:", err)
	}

	// 发送消息
	message := "Hello, RabbitMQ!" + time.Now().Format("2006-01-02 15:04:05")
	err = ch.PublishWithContext(
		context.Background(),
		"",        // 交换机名称
		queueName, // 队列名称
		false,     // 必需的
		false,     // 立即发布
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		log.Fatal("无法发送消息:", err)
	}

	log.Println("消息已发送:", message)
}

func ConsumerMessage() {
	// 创建 RabbitMQ 连接
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("无法连接到 RabbitMQ:", err)
	}
	defer conn.Close()

	// 创建通道
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("无法打开通道:", err)
	}
	defer ch.Close()

	// 声明队列
	queueName := "my_queue"
	_, err = ch.QueueDeclare(
		queueName, // 队列名称
		false,     // 持久化
		false,     // 自动删除
		false,     // 排他性
		false,     // 等待服务器确认
		nil,       // 参数
	)
	if err != nil {
		log.Fatal("无法声明队列:", err)
	}

	// 接收消息
	msgs, err := ch.Consume(
		queueName, // 队列名称
		"",        // 消费者名称
		true,      // 自动应答
		false,     // 排他性
		false,     // 不等待服务器确认
		false,     // 消费者取消通知
		nil,       // 参数
	)
	if err != nil {
		log.Fatal("无法接收消息:", err)
	}

	for msg := range msgs {
		log.Println("收到消息:", string(msg.Body))
	}
}
