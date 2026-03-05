package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beanstalkd/go-beanstalk"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	// 1. 连接到 Beanstalkd 服务
	// BeanstalkdAddr 在 config.go 中定义
	conn, err := beanstalk.Dial("tcp", BeanstalkdAddr)
	if err != nil {
		log.Fatalf("无法连接到 Beanstalkd: %v", err)
	}
	defer conn.Close()

	// 2. 创建 TubeSet 消费者
	// Watch 命令用于告诉 Beanstalkd 我们要监听哪些 Tube
	// NewTubeSet 可以同时监听多个 Tube
	tubeSet := beanstalk.NewTubeSet(conn, TubeName)

	log.Printf("开始监听 Tube: [%s] ... \n", TubeName)

	for {

		fmt.Println()

		// 3. 获取（预留）任务
		// Reserve 命令会阻塞等待直到有任务可用
		// ReserveTimeout 可以设置超时时间，如果不设置超时，会一直阻塞
		// 当任务被 Reserve 后，它的状态变为 "reserved"，其他消费者无法获取该任务
		id, body, err := tubeSet.Reserve(ReserveTimeout)

		if err != nil {
			// 如果是超时错误，说明暂时没有任务，继续循环
			if connErr, ok := err.(beanstalk.ConnError); ok && connErr.Err == beanstalk.ErrTimeout {
				// log.Println("暂时没有任务...") // 为了保持日志清爽，这里注释掉
				continue
			} else if connErr, ok := err.(beanstalk.ConnError); ok && connErr.Err == beanstalk.ErrDeadline {
				continue
			}

			// 其他错误可能是连接断开等
			log.Printf("获取任务失败: %v", err)
			time.Sleep(1 * time.Second)
			continue
		}

		log.Printf("收到任务 ID: %d, 内容: %s", id, body)

		// 4. 处理任务逻辑
		processTask(id, body)

		// 5. 删除任务
		// 只有删除任务才算真正完成，否则 TTR 超时后任务会重新变成 Ready 状态被其他消费者获取。
		// 如果处理失败，可以选择 Release 释放任务回队列，或者 Bury 掩埋任务。
		// 这里演示处理成功后删除任务。
		err = conn.Delete(id)
		if err != nil {
			log.Printf("删除任务失败 ID: %d, 错误: %v", id, err)
		} else {
			log.Printf("任务处理完成并删除 ID: %d", id)
		}
	}
}

// processTask 模拟处理任务
func processTask(id uint64, body []byte) {
	spew.Dump("正在处理任务 ID:", id, "内容:", string(body))
	fmt.Printf("任务执行开始时间 %s \n", time.Now().Format(time.DateTime))
	// 模拟耗时操作
	time.Sleep(500 * time.Millisecond)
}
