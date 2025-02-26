package main

import (
	"context"
	"fmt"
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
	queueName = "ttl_queue" // 定义一个队列名称
)

func failOnError(msg string, err error) {
	if err != nil {
		log.Fatalf("%s ==> %s", msg, err)
	}
}

func ProducerMessage() {
	conn, err := amqp.Dial(url)
	failOnError("无法连接到 RabbitMQ", err)
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	ch, err := conn.Channel()
	failOnError("无法打开通道", err)
	defer func(ch *amqp.Channel) {
		err := ch.Close()
		if err != nil {

		}
	}(ch)

	argsQue := make(map[string]interface{})
	// 添加过期时间（10秒消息存活时间）
	argsQue["x-message-ttl"] = int32((10 * time.Second).Milliseconds()) // 单位毫秒

	// 声明消息队列
	queue, err := ch.QueueDeclare(
		queueName, // 队列名称
		true,      // 持久化
		false,     // 自动删除
		false,     // 排他性
		false,     // 等待服务器确认
		argsQue)
	failOnError("声明队列失败", err)

	// 发送消息
	for i := 1; i <= 10; i++ {
		// 模拟发送消息的耗时
		time.Sleep(time.Second)

		// 消息内容
		msg := fmt.Sprintf("第 %d 次发送消息，发送时间：%s", i, time.Now().Format("2006-01-02 15:04:05"))
		err := ch.PublishWithContext(
			context.Background(),
			"",         // 交换器名称，填空字符串表示使用默认交换器
			queue.Name, // 队列名称
			false,      // 必需的
			false,      // 立即发布
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(msg),
			})

		failOnError("发送消息失败", err)
		fmt.Println(msg)
	}

}

func main() {
	ProducerMessage()
}
