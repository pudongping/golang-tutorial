package main

import (
	"log"

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

const (
	exchangeName = "logs_sample"
	exchangeType = "fanout" // 扇出
)

func failOnError(msg string, err error) {
	if err != nil {
		log.Fatalf("%s ==> %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial(url)
	failOnError("无法连接到 RabbitMQ", err)
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError("无法打开通道", err)
	defer ch.Close()

	// 声明交换器
	err = ch.ExchangeDeclare(
		exchangeName,
		exchangeType,
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError("声明交换器失败", err)

	q, err := ch.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	failOnError("声明队列失败", err)

	// 绑定是交换器和队列之间的关系。这可以简单地理解为：队列对来自此交换器的消息感兴趣。
	err = ch.QueueBind(
		q.Name, // 队列名称
		"",
		exchangeName,
		false,
		nil,
	)
	failOnError("绑定队列失败", err)

	msgs, err := ch.Consume(
		q.Name, // 队列名称
		"",     // 消费者
		true,   // 设置为 true，表示自动应答
		false,
		false,
		false,
		nil,
	)
	failOnError("注册一个消费者失败", err)

	forever := make(chan struct{})

	go func() {
		for d := range msgs {
			log.Printf("🫡 收到消息 🫱 %s \n", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C \n")
	<-forever

}
