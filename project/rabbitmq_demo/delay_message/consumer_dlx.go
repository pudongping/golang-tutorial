// consumer_dlx.go
package main

import (
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(msg string, err error) {
	if err != nil {
		log.Fatalf("%s ==> %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError("连接失败", err)
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError("打开通道失败", err)
	defer ch.Close()

	// 直接消费最终队列
	msgs, err := ch.Consume(
		"process_queue", // 队列名称
		"",              // 消费者标识
		false,           // 自动确认
		false,           // 排他
		false,           // 不等待
		false,
		nil,
	)
	failOnError("注册消费者失败", err)

	log.Printf(" [*] 等待延迟消息...")

	for d := range msgs {
		log.Printf("接收时间: %s → 消息内容: %s",
			time.Now().Format("15:04:05"),
			d.Body)
		// 模拟处理耗时
		time.Sleep(500 * time.Millisecond)
		d.Ack(false)
	}
}
