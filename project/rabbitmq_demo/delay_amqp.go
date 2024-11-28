package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func CallDelayMessage() {
	// RabbitMQ 连接信息
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("无法连接到 RabbitMQ：%s", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("无法打开通道：%s", err)
	}
	defer ch.Close()

	// 声明交换机
	exchangeName := "delay_exchange"
	err = ch.ExchangeDeclare(
		exchangeName, // 交换机名称
		"direct",     // 交换机类型
		true,         // 持久化
		false,        // 自动删除
		false,        // 内部使用
		false,        // 等待服务器确认
		nil,          // 参数
	)
	if err != nil {
		log.Fatalf("无法声明交换机：%s", err)
	}

	// 声明延迟队列
	delayQueueName := "delay_queue"
	_, err = ch.QueueDeclare(
		delayQueueName, // 队列名称
		true,           // 持久化
		false,          // 自动删除
		false,          // 排他性
		false,          // 等待服务器确认
		amqp.Table{
			"x-dead-letter-exchange":    exchangeName,                               // 死信交换机
			"x-dead-letter-routing-key": delayQueueName,                             // 死信队列路由键
			"x-message-ttl":             int32(10 * time.Second / time.Millisecond), // 消息过期时间（毫秒）
			"x-expires":                 int32(30 * time.Second / time.Millisecond), // 队列过期时间（毫秒）
		},
	)
	if err != nil {
		log.Fatalf("无法声明延迟队列：%s", err)
	}

	// 声明死信队列
	deadLetterQueueName := "dead_letter_queue"
	_, err = ch.QueueDeclare(
		deadLetterQueueName, // 队列名称
		true,                // 持久化
		false,               // 自动删除
		false,               // 排他性
		false,               // 等待服务器确认
		nil,                 // 参数
	)
	if err != nil {
		log.Fatalf("无法声明死信队列：%s", err)
	}

	// 绑定延迟队列与交换机
	err = ch.QueueBind(
		delayQueueName, // 队列名称
		delayQueueName, // 路由键
		exchangeName,   // 交换机名称
		false,          // 不等待服务器确认
		nil,            // 参数
	)
	if err != nil {
		log.Fatalf("无法绑定延迟队列与交换机：%s", err)
	}

	// 创建等待组
	var wg sync.WaitGroup

	// 消费者消费延迟消息
	wg.Add(1)
	go ConsumerDelayMessage(ch, delayQueueName, &wg)

	// 生产者发送延迟消息
	go ProducerDelayMessage(ch, exchangeName, delayQueueName)

	// 等待消费者执行完成
	wg.Wait()

	log.Println("程序终止")
}

// 生产者发送延迟消息
func ProducerDelayMessage(ch *amqp.Channel, exchangeName, queueName string) {
	for i := 1; i <= 10; i++ {
		message := fmt.Sprintf("投递时间：%s Delayed message %d", time.Now().Format("2006-01-02 15:04:05"), i)
		err := ch.PublishWithContext(
			context.Background(),
			exchangeName, // 交换机名称
			queueName,    // 路由键
			false,        // 必需的
			false,        // 立即发布
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(message),
			},
		)
		if err != nil {
			log.Printf("无法发送延迟消息：%s", err)
		} else {
			log.Printf("发送延迟消息：%s", message)
		}

		time.Sleep(1 * time.Second)
	}
}

// 消费者消费延迟消息
func ConsumerDelayMessage(ch *amqp.Channel, queueName string, wg *sync.WaitGroup) {
	defer wg.Done()

	msgs, err := ch.Consume(
		queueName, // 队列名称
		"",        // 消费者标识符
		false,     // 自动应答
		false,     // 排他性
		false,     // 不等待服务器确认
		false,     // 消费者取消通知
		nil,       // 参数
	)
	if err != nil {
		log.Fatalf("无法注册消费者：%s", err)
	}

	// 使用通道来传递退出信号
	exit := make(chan struct{})

	// 启动一个协程来检查是否需要退出
	go func() {
		// 等待一段时间后发送退出信号
		time.Sleep(30 * time.Second)
		close(exit)
	}()

	for {
		select {
		case msg, ok := <-msgs:
			if ok {
				log.Printf("接收时间为：%s 收到消息：%s", time.Now().Format("2006-01-02 15:04:05"), msg.Body)
				msg.Ack(false)
			} else {
				// 通道已关闭，退出循环
				return
			}
		case <-exit:
			// 收到退出信号，退出循环
			return
		}
	}
}
