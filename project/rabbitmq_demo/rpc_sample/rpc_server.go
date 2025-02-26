package main

import (
	"context"
	"log"
	"strconv"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// 所需参数
const (
	host     = "localhost" // 服务接入地址
	username = "guest"     // 角色控制台对应的角色名称
	password = "guest"     // 角色对应的密钥
	vhost    = ""          // 要使用的Vhost全称
)

const (
	url = "amqp://" + username + ":" + password + "@" + host + ":5672/" + vhost
)

func failOnError(msg string, err error) {
	if err != nil {
		log.Fatalf("%s ==> %s", msg, err)
	}
}

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

// 模拟发送消息
func main() {
	conn, err := amqp.Dial(url)
	failOnError("无法连接到 RabbitMQ", err)
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError("无法打开通道", err)
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"rpc_queue", // 队列名称
		false,       // 是否持久化
		false,       // 是否自动删除
		false,       // 是否排他
		false,       // 是否阻塞
		nil,         // 额外属性
	)
	failOnError("声明队列失败", err)

	// 设置每次投递一个消息
	err = ch.Qos(
		1,     // 消费者未确认消息的最大数量
		0,     // 消费者未确认消息的最大字节数
		false, // 应用于整个通道
	)
	failOnError("设置 Qos 失败", err)

	msgs, err := ch.Consume(
		q.Name, // 队列名称
		"",     // 消费者
		false,  // 是否自动应答
		false,  // 是否排他
		false,  // no-local
		false,  // no-wait
		nil,    // 额外属性
	)
	failOnError("注册一个消费者失败", err)

	forever := make(chan struct{})

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		for d := range msgs {
			n, err := strconv.Atoi(string(d.Body))
			failOnError("无法转换消息为整数", err)

			log.Printf("收到请求: %d", n)
			// 执行RPC业务
			response := fib(n)

			// 将结果发布到执行的回调队列(返回RPC调用结果)
			err = ch.PublishWithContext(
				ctx,
				"",        // 交换器名称
				d.ReplyTo, // 路由 key
				false,     // 强制
				false,     // 立即
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          []byte(strconv.Itoa(response)),
				})
			failOnError("无法发送消息", err)

			// 确认消息
			d.Ack(false)
		}

	}()

	log.Printf("🐇 等待 RPC 请求 🐇")
	<-forever

}
