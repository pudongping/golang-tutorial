package main

import (
	"log"
	"os"

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
	exchangeName = "logs_direct"
	exchangeType = "direct" // 扇出
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
		"",    // 空字符串作为队列名称，将会得到一个随机生成的名称，类似 amq.gen-6OzD2FA4N-tCo_C4pA1UmQ
		false, // 非持久队列
		false,
		true, // 独占队列（当前声明队列的连接关闭后即被删除）
		false,
		nil,
	)
	failOnError("声明队列失败", err)

	if len(os.Args) < 2 {
		log.Printf("Usage: %s [info] [warning] [error]", os.Args[0])
		os.Exit(0)
	}

	for _, s := range os.Args[1:] {
		log.Printf("绑定队列 [%s] 到交换器 [%s] 使用路由 key [%s] ", q.Name, exchangeName, s)

		// 绑定是交换器和队列之间的关系。这可以简单地理解为：队列对来自此交换器的消息感兴趣。
		err = ch.QueueBind(
			q.Name, // 队列名称
			s,
			exchangeName,
			false,
			nil,
		)
		failOnError("绑定队列失败", err)

	}

	// 创建消费者并消费指定消息队列中的消息
	msgs, err := ch.Consume(
		q.Name, // 队列名称
		"",     // 消费者
		true,
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
