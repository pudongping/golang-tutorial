package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beanstalkd/go-beanstalk"
)

func main() {
	// 1. 连接到 Beanstalkd 服务
	conn, err := beanstalk.Dial("tcp", BeanstalkdAddr)
	if err != nil {
		log.Fatalf("无法连接到 Beanstalkd 错误原因为 %+v", err)
	}
	defer conn.Close()

	// 2. 创建一个 Tube 生产者
	// Use 命令用于告诉 Beanstalkd 我们要向哪个 Tube 放入任务
	// 如果不指定，默认是 "default"
	tube := beanstalk.NewTube(conn, TubeName)

	log.Printf("开始向 Tube [%s] 发送任务...", TubeName)

	// 3. 发送任务
	// 模拟发送 5 个任务
	for i := 1; i <= 5; i++ {
		payload := []byte(fmt.Sprintf("你好 Beanstalkd ----> %d - Time: %s", i, time.Now().Format(time.DateTime)))

		// Put 参数说明:
		// body: 任务的具体数据（字节数组）
		// pri: 优先级，数字越小优先级越高 (0-2^32-1)，默认 1024。0 为最高优先级。
		// delay: 延迟时间，0 表示立即放入 Ready 状态，让消费者可以立即获取。
		// ttr: Time To Run，任务预留时间（消费者处理该任务允许的最大时间）。
		//      如果消费者在 TTR 时间内没有发送 delete, release 或 touch 命令，
		//      任务会重回 Ready 状态，可能被其他消费者再次获取。
		id, err := tube.Put(payload, 1024, 0, 10*time.Second)
		if err != nil {
			log.Printf("任务发送失败: %v \n", err)
			continue
		}

		log.Printf("成功发送任务 ID: %d, 内容: %s \n", id, payload)
		time.Sleep(1 * time.Second)
	}

	log.Println("所有任务发送完毕")
}
