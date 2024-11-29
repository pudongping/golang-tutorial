package main

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	taskQueueName = "task_queue"
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

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello_rabbitmq"
	} else {
		s = strings.Join(args[1:], " ")
	}

	return s
}

func Task() {
	conn, err := amqp.Dial(url)
	failOnError("无法连接到 RabbitMQ", err)
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError("无法打开通道", err)
	defer ch.Close()

	q, err := ch.QueueDeclare(
		taskQueueName, // 队列名称
		true,          // 持久化（如果不设置为 true，那么当 RabbitMQ 服务器停止运行或者崩溃时，消息就会丢失）
		false,         // 自动删除
		false,         // 排他性
		false,         // 等待服务器确认
		nil,           // 参数
	)
	failOnError("无法声明队列", err)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := bodyFrom(os.Args)
	err = ch.PublishWithContext(
		ctx,
		"",     // 交换器名称
		q.Name, // 队列名称
		false,  // 必需的
		false,  // 立即发布
		amqp.Publishing{
			ContentType:  "text/plain",
			DeliveryMode: amqp.Persistent, // 持久（交付模式：瞬态/持久）=> 将消息标记为持久的（在队列中标记为“持久化”还不行，还一定需要在发送消息的时候标记为“持久”）
			Body:         []byte(body),
		},
	)

	failOnError("无法发送消息", err)
	log.Printf("消息已发送： %s \n", body)
}

func main() {
	// 作为生产者
	Task()
}
