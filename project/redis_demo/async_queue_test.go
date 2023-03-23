package redis_demo

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis"
)

func TestAsyncQueueProducer(t *testing.T) {
	payload := []byte(`{"task": "send_email", "email": "test@example.com", "content": "hello world"}`)

	// 模拟将任务放入队列
	err := NewAsyncQueue().Enqueue(payload)
	if err != nil {
		fmt.Println("错误为：", err)
	} else {
		fmt.Println("任务投递成功")
	}

}

func TestAsyncQueueConsumer(t *testing.T) {

	asyncQueueObj := NewAsyncQueue()

	for {
		val, err := asyncQueueObj.Dequeue()
		if err == redis.Nil {
			fmt.Println("队列已经消费完毕，跳过本次循环")
			continue
		} else if err != nil {
			fmt.Println("出错啦，错误原因：", err)
			break
		}

		// 反序列化任务
		var task map[string]interface{}
		if err := json.Unmarshal(val, &task); err != nil {
			fmt.Println("反序列化失败：", err)
			continue
		}

		fmt.Println("取出的任务信息为：", task)

		// 后面可以执行对应的任务
	}

}

func TestAsyncDelayQueueProducer(t *testing.T) {
	asyncDelayQueueObj := NewAsyncDelayQueue()

	for i := 0; i < 10; i++ {

		payload := map[string]interface{}{
			"task":    "send_email",
			"email":   "test@example.com",
			"content": "hello worlds",
			"times":   i,
			"now":     time.Now(),
		}
		payloadByte, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("有错误：", err)
			continue
		}
		// 加入异步延迟队列
		err = asyncDelayQueueObj.Enqueue(payloadByte, int64(i))
		if err != nil {
			fmt.Println("加入异步延迟队列时，有错误：", err)
			continue
		}

	}
}

func TestAsyncDelayQueueConsumer(t *testing.T) {
	asyncDelayQueueObj := NewAsyncDelayQueue()

	for {
		res, err := asyncDelayQueueObj.RedisClient.ZRangeWithScores(asyncDelayQueueObj.QueueName, 0, 0).Result()
		if err == redis.Nil {
			fmt.Println("队列已经消费完毕，跳过本次循环")
			continue
		} else if err != nil {
			fmt.Println("出错啦，错误原因：", err)
			break
		}

		if len(res) == 0 || res[0].Score > float64(time.Now().Unix()) {
			fmt.Println("取不到数据，或者现在还没有到执行时间")
			continue
		}

		// 取出分数最小的任务
		val, err := asyncDelayQueueObj.RedisClient.ZPopMin(asyncDelayQueueObj.QueueName, 1).Result()
		if err != nil {
			fmt.Println("取出任务失败：", err)
			break
		}
		
		// 反序列化任务
		var task map[string]interface{}
		if err := json.Unmarshal([]byte(val[0].Member.(string)), &task); err != nil {
			fmt.Println("反序列化失败：", err)
			continue
		}

		fmt.Println("取出的任务信息为：", task)

		// 后面可以执行对应的任务

	}

}
