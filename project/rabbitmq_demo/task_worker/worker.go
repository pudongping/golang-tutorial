package main

import (
	"bytes"
	"log"
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

func Worker() {
	conn, err := amqp.Dial(url)
	failOnError("无法连接到 RabbitMQ", err)
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError("无法打开通道", err)
	defer ch.Close()

	// 声明队列
	// 请注意：我们也在这里声明队列。因为我们可能在发布者（生产者）之前启动使用者（消费者）
	// 所以，我们希望在尝试使用队列中的消息之前确保队列存在
	q, err := ch.QueueDeclare(
		taskQueueName, // 队列名称
		true,          // 持久化
		false,         // 自动删除
		false,         // 排他性
		false,         // 等待服务器确认
		nil,           // 参数
	)
	failOnError("无法声明队列", err)

	// 这里，我们将预取计数设置为 1
	// 这就告诉 RabbitMQ 不要一次向一个 worker 发出多个消息，换句话来说就是
	// 在处理并确认前一条消息之前，不要向 worker 发送新消息
	err = ch.Qos(1, 0, false)
	failOnError("设置 Qos 失败", err)

	msgs, err := ch.Consume(
		q.Name, // 队列名称（消费者要监听和生产者一样的队列）
		"",     // 消费者名称
		false,  // 自动应答（这里标记为，关闭自动消息确认，意思也就是，我们需要进行手动消息确认）
		false,  // 排他性
		false,  // 不等待服务器确认
		false,  // 消费者取消通知
		nil,    // 参数
	)
	failOnError("无法接收消息", err)

	// 开启循环不断地消费消息
	forever := make(chan struct{})
	go func() {
		for d := range msgs {

			// 这里的用意是：每当消息正文中出现一个点 `.` 就伪造一秒钟的延时，用来模拟耗时任务
			dotCount := bytes.Count(d.Body, []byte(".")) // 数一下有多少个 `.`
			t := time.Duration(dotCount)
			log.Printf("🫡 收到消息 dotCount %v 🫱 %s \n", dotCount, d.Body)
			time.Sleep(t * time.Second) // 模拟耗时的任务执行

			// 这将确认一次传递（手动传递消息确认）
			// 消息确认必须在接收消息的同一通道 Channel 上发送。如果使用不同的通道进行消息确认将导致通道级协议异常
			// 如果没有这一行 ack 确认，那么消息是不会从队列里面删除的，可以通过 rabbitmqctl list_queues name messages_ready messages_unacknowledged 命令进行查看
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func main() {
	log.Println("worker start.")
	// 作为消费者
	Worker()
}
