package main

import (
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

func ConsumerMessage() {

	// 1. 建立与RabbitMQ服务器的连接
	conn, err := amqp.Dial(url)
	failOnError("无法连接到 RabbitMQ", err)
	defer func() {
		// 优雅关闭连接，在main函数结束时执行
		if err := conn.Close(); err != nil {
			log.Printf("关闭连接失败: %v", err)
		}
	}()

	// 2. 创建通信通道
	ch, err := conn.Channel()
	failOnError("无法打开通道", err)
	defer func() {
		// 确保通道关闭
		if err := ch.Close(); err != nil {
			log.Printf("关闭通道失败: %v", err)
		}
	}()

	// 3. 声明队列（幂等操作，如果队列已存在则直接使用）
	// 参数必须与生产者完全一致，否则会报错
	_, err = ch.QueueDeclare(
		queueName, // 队列名称
		true,      // 持久化队列（与生产者设置一致）
		false,     // 自动删除（消费者断开后不删除队列）
		false,     // 排他性（仅限于当前连接）
		false,     // 不等待服务器响应
		amqp.Table{
			"x-message-ttl": int32(10 * time.Second / time.Millisecond), // 必须与生产者相同的TTL设置
		},
	)
	failOnError("声明队列失败", err)

	// 4. 设置QoS（服务质量控制）
	// 每个消费者最多同时处理1条未确认的消息
	err = ch.Qos(
		1,     // 预取计数 （实现公平调度，防止单个消费者占用所有消息）
		0,     // 预取大小（字节），0表示不限制
		false, // 全局设置（true表示应用于整个通道）
	)
	failOnError("设置QoS失败", err)

	// 5. 创建消费者
	msgs, err := ch.Consume(
		queueName, // 监听的队列名称
		"",        // 消费者标签（空字符串由服务器自动生成）
		false,     // 自动确认（关闭以支持手动确认）
		false,     // 排他性（多个消费者可以同时消费）
		false,     // 不等待服务器响应
		false,     // 其他参数
		nil,       // 额外参数
	)
	failOnError("注册消费者失败", err)

	// 6. 启动消息处理循环
	log.Println(" [*] 等待消息。退出请按 CTRL+C")

	// 创建退出信号通道
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			// 处理接收到的消息
			log.Printf("收到消息: %s", d.Body)
			log.Printf("消息属性: %+v", d.Headers)

			// 模拟消息处理耗时
			time.Sleep(500 * time.Millisecond)

			// 手动确认消息（重要！否则消息会留在队列中）
			// 如果Auto-Ack设置为true，则不需要此操作
			if err := d.Ack(false); err != nil {
				log.Printf("消息确认失败: %v", err)
			}
		}
	}()

	// 阻塞主线程，保持消费者运行
	<-forever
}

func main() {
	ConsumerMessage()
}
