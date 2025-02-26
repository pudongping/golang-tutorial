// producer_dlx.go
package main

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	url         = "amqp://guest:guest@localhost:5672/"
	delayQueue  = "delay_queue"   // 延迟存储队列
	finalQueue  = "process_queue" // 实际消费队列
	dlxExchange = "dlx_exchange"  // 死信交换机
)

func failOnError(msg string, err error) {
	if err != nil {
		log.Fatalf("%s ==> %s", msg, err)
	}
}

func main() {
	// 1. 建立连接
	conn, err := amqp.Dial(url)
	failOnError("连接失败", err)
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError("打开通道失败", err)
	defer ch.Close()

	// 2. 声明死信交换机
	err = ch.ExchangeDeclare(
		dlxExchange, // 交换机名称
		"direct",    // 类型
		true,        // 持久化
		false,       // 自动删除
		false,       // 内部
		false,       // 等待确认
		nil,
	)
	failOnError("声明交换机失败", err)

	// 3. 声明最终消费队列
	_, err = ch.QueueDeclare(
		finalQueue,
		true,  // 持久化
		false, // 自动删除
		false, // 排他
		false, // 等待
		nil,
	)
	failOnError("声明队列失败", err)

	// 绑定队列到死信交换机
	err = ch.QueueBind(
		finalQueue,  // 队列名称
		"delay_key", // 路由键
		dlxExchange, // 交换机
		false,       // 等待
		nil,
	)
	failOnError("队列绑定失败", err)

	// 4. 声明延迟队列（带死信配置）
	args := amqp.Table{
		"x-dead-letter-exchange":    dlxExchange, // 指定死信交换机
		"x-dead-letter-routing-key": "delay_key", // 指定路由键
		// "x-message-ttl":             20000,       // 20s TTL（单位毫秒） -----> 移除 x-message-ttl 参数
	}
	_, err = ch.QueueDeclare(
		delayQueue,
		true, // 持久化
		false,
		false,
		false,
		args,
	)
	failOnError("声明延迟队列失败", err)

	// 5. 发送延迟消息
	body := "producer_dlx 延迟消息：" + time.Now().Format("15:04:05")
	err = ch.PublishWithContext(
		context.Background(),
		"",         // 使用默认交换机
		delayQueue, // 发送到延迟队列
		false,      // 强制
		false,      // 立即
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
			Expiration:  "3000", // 发送消息时设置每条消息的 TTL（单位：毫秒） -----> 添加 Expiration 字段，动态设置延迟时间 （此处为 5s）
		})
	failOnError("发送失败", err)
	log.Printf(" [x] 已发送消息: %s", body)
}
