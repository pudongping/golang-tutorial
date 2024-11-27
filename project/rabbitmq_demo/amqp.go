package rabbitmq_demo

import (
	"context"
	"log"
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
	url       = "amqp://" + username + ":" + password + "@" + host + ":5672/" + vhost
	queueName = "my_queue" // 定义一个队列名称
)

func failOnError(msg string, err error) {
	if err != nil {
		log.Fatalf("%s ==> %s", msg, err)
	}
}

// 演示一个消息生产者
func ProducerMessage() {
	// 1. 创建 RabbitMQ 连接
	// 该连接抽象了套接字连接，并为我们处理协议版本协商和认证等
	conn, err := amqp.Dial(url)
	failOnError("无法连接到 RabbitMQ", err)
	defer conn.Close()

	// 2. 创建一个通道，
	ch, err := conn.Channel()
	failOnError("无法打开通道", err)
	defer ch.Close()

	// 3. 声明消息要发送到的队列
	q, err := ch.QueueDeclare(
		queueName, // 队列名称
		false,     // 持久化
		false,     // 自动删除
		false,     // 排他性
		false,     // 等待服务器确认
		nil,       // 参数
	)
	failOnError("无法声明队列", err)

	// 4. 将信息发布到声明的队列中
	message := "Hello, RabbitMQ! " + time.Now().Format("2006-01-02 15:04:05")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(
		ctx,
		"",     // 交换机名称
		q.Name, // 队列名称
		false,  // 必需的
		false,  // 立即发布
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	failOnError("无法发送消息", err)

	// 如果需要检查队列，可以使用 `rabbitmqctl list_queues` 命令来查看
	log.Println("消息已发送:", message)
}

// 演示一个消息消费者
func ConsumerMessage() {
	// 创建 RabbitMQ 连接
	conn, err := amqp.Dial(url)
	failOnError("无法连接到 RabbitMQ", err)
	defer conn.Close()

	// 创建通道
	ch, err := conn.Channel()
	failOnError("无法打开通道", err)
	defer ch.Close()

	// 声明队列
	// 请注意：我们也在这里声明队列。因为我们可能在发布者（生产者）之前启动使用者（消费者）
	// 所以，我们希望在尝试使用队列中的消息之前确保队列存在
	q, err := ch.QueueDeclare(
		queueName, // 队列名称
		false,     // 持久化
		false,     // 自动删除
		false,     // 排他性
		false,     // 等待服务器确认
		nil,       // 参数
	)
	failOnError("无法声明队列", err)

	// 获取接收消息的 Delivery 通道
	msgs, err := ch.Consume(
		q.Name, // 队列名称（消费者要监听和生产者一样的队列）
		"",     // 消费者名称
		true,   // 自动应答
		false,  // 排他性
		false,  // 不等待服务器确认
		false,  // 消费者取消通知
		nil,    // 参数
	)
	failOnError("无法接收消息", err)

	// 获取消息队列中的消息
	forever := make(chan struct{})
	go func() {
		for msg := range msgs {
			log.Printf("收到消息: %s \n", string(msg.Body))
			// 手动回复 ack
			// msg.Ack(false)
		}
	}()

	log.Printf(" [Consumer] Waiting for messages. To exit press CTRL+C \n")
	<-forever

}
