package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
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
	url = "amqp://" + username + ":" + password + "@" + host + ":5672/" + vhost
)

func failOnError(msg string, err error) {
	if err != nil {
		log.Fatalf("%s ==> %s", msg, err)
	}
}

// 生成随机字符串
func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		// A-Z
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

// 生成随机整数
func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func bodyFrom(args []string) int {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "30"
	} else {
		s = strings.Join(args[1:], " ")
	}
	n, err := strconv.Atoi(s)
	failOnError("无法转换参数为整数", err)
	return n
}

func fibonacciRPC(n int) (res int, err error) {
	conn, err := amqp.Dial(url)
	failOnError("无法连接到 RabbitMQ", err)
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError("无法打开通道", err)
	defer ch.Close()

	// 声明交换器
	q, err := ch.QueueDeclare(
		"",    // 空字符串作为队列名称，将会得到一个随机生成的名称，类似 amq.gen-6OzD2FA4N-tCo_C4pA1UmQ
		false, // 非持久队列
		false,
		true, // 独占队列
		false,
		nil,
	)
	failOnError("声明队列失败", err)

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

	// 生成唯一的 correlationID
	corrID := randomString(32)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(
		ctx,
		"",          // 交换器名称
		"rpc_queue", // 路由 key
		false,
		false,
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: corrID, // 设置为每个请求的唯一值
			ReplyTo:       q.Name, // 设置为回调队列
			Body:          []byte(strconv.Itoa(n)),
		})
	failOnError("发布消息失败", err)

	for d := range msgs {
		// 验证响应与请求的关系
		if corrID == d.CorrelationId {
			res, err = strconv.Atoi(string(d.Body))
			failOnError("无法转换消息为整数", err)
			break
		}
	}

	return
}

func main() {
	// 生成随机种子
	rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

	n := bodyFrom(os.Args)

	log.Printf(" [👉] 请求 fib(%d) \n", n)
	res, err := fibonacciRPC(n)
	failOnError("RPC 失败", err)

	log.Printf(" [👉] 结果 fib(%d) = %d \n", n, res)
}
