// consumer_delay_plugin.go
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

	msgs, err := ch.Consume(
		"delayed_queue",
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError("注册消费者失败", err)

	log.Printf(" [*] 等待延迟消息...")

	for d := range msgs {
		log.Printf("接收时间: %s → 消息内容: %s (原始发送时间: %s)",
			time.Now().Format("15:04:05"),
			d.Body,
			d.Timestamp.Format("15:04:05"))
		d.Ack(false)
	}
}
