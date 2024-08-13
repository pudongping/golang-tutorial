package redis_demo

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		PoolSize: 25, // 连接池大小
	})

	// 测试连接
	_, err := client.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("连接Redis失败，错误原因：%v", err))
	}

	return client
}

// AsyncQueue 异步队列
type AsyncQueue struct {
	RedisClient *redis.Client
	QueueName   string
}

func NewAsyncQueue() *AsyncQueue {
	return &AsyncQueue{
		RedisClient: NewRedisClient(),
		QueueName:   "async_queue_{channel}",
	}
}

func (a *AsyncQueue) Enqueue(jobPayload []byte) error {
	return a.RedisClient.LPush(a.QueueName, jobPayload).Err()
}

func (a *AsyncQueue) Dequeue() ([]byte, error) {
	return a.RedisClient.RPop(a.QueueName).Bytes()
}

// AsyncDelayQueue 异步延迟队列
type AsyncDelayQueue struct {
	RedisClient *redis.Client
	QueueName   string
}

func NewAsyncDelayQueue() *AsyncDelayQueue {
	return &AsyncDelayQueue{
		RedisClient: NewRedisClient(),
		QueueName:   "async_delay_queue_{channel}",
	}
}

// Enqueue 加入异步延迟队列
// jobPayload 任务载荷
// delay 延迟时间（单位，秒）
func (a *AsyncDelayQueue) Enqueue(jobPayload []byte, delay int64) error {
	return a.RedisClient.ZAdd(a.QueueName, redis.Z{
		Score:  float64(time.Now().Unix() + delay),
		Member: jobPayload,
	}).Err()
}
