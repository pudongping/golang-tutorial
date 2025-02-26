// producer_delay_plugin.go
package main

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	url           = "amqp://guest:guest@localhost:5672/"
	delayExchange = "delayed_exchange" // 延迟交换机
	delayQueue    = "delayed_queue"
)

func failOnError(msg string, err error) {
	if err != nil {
		log.Fatalf("%s ==> %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial(url)
	failOnError("连接失败", err)
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError("打开通道失败", err)
	defer ch.Close()

	// 声明延迟交换机（必须设置x-delayed-type）
	args := amqp.Table{"x-delayed-type": "direct"}
	err = ch.ExchangeDeclare(
		delayExchange,
		"x-delayed-message", // 特殊类型
		true,                // 持久化
		false,
		false,
		false,
		args,
	)
	failOnError("声明交换机失败", err)

	// 声明队列并绑定
	_, err = ch.QueueDeclare(
		delayQueue,
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError("声明队列失败", err)

	err = ch.QueueBind(
		delayQueue,
		"delay", // 路由键
		delayExchange,
		false,
		nil,
	)
	failOnError("绑定失败", err)

	// 发送延迟消息
	headers := amqp.Table{"x-delay": 120000} // 2分钟延迟，单位：毫秒
	body := "插件延迟消息：" + time.Now().Format("15:04:05")
	err = ch.PublishWithContext(
		context.Background(),
		delayExchange,
		"delay",
		false,
		false,
		amqp.Publishing{
			Headers:     headers,
			ContentType: "text/plain",
			Body:        []byte(body),
			Timestamp:   time.Now(), // 消息发送时间
		})
	failOnError("发送失败", err)
	log.Printf(" [x] 已发送消息: %s", body)
}
