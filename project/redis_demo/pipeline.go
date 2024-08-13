package redis_demo

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func RDBClient() (*redis.Client, error) {
	// 创建一个 Redis 客户端
	// 也可以使用数据源名称（DSN）来创建
	// redis://<user>:<pass>@localhost:6379/<db>
	opt, err := redis.ParseURL("redis://localhost:6379/0")
	if err != nil {
		return nil, err
	}
	client := redis.NewClient(opt)

	// 通过 cient.Ping() 来检查是否成功连接到了 redis 服务器
	_, err = client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}

func pipeline1() {
	rdb, err := RDBClient()
	if err != nil {
		panic(err)
	}

	// 使用 Pipeline 批量执行多个命令
	// 以下代码相当于将
	// INCR pipeline_counter
	// EXPIRE pipeline_counter 10
	// 作为一个事务一次性执行
	pipe := rdb.Pipeline()
	incr := pipe.Incr("pipeline_counter")
	pipe.Expire("pipeline_counter", 10*time.Second)
	cmds, err := pipe.Exec()
	if err != nil {
		panic(err)
	}

	// 获取 incr 命令的执行结果
	fmt.Println("pipeline_counter:", incr.Val())
	for _, cmd := range cmds {
		// 这里可以遍历出 pipe.Exec() 返回的所有命令执行结果
		fmt.Printf("cmd: %#v \n", cmd)
	}
}

func pipeline2() {
	rdb, err := RDBClient()
	if err != nil {
		panic(err)
	}

	var incr *redis.IntCmd

	// 以下代码相当于将
	// INCR pipeline_counter
	// EXPIRE pipeline_counter 10
	// 作为一个事务一次性执行
	cmds, err := rdb.Pipelined(func(pipe redis.Pipeliner) error {
		incr = pipe.Incr("pipeline_counter")
		pipe.Expire("pipeline_counter", 10*time.Second)
		return nil
	})
	if err != nil {
		panic(err)
	}

	// 获取 incr 命令的执行结果
	fmt.Println("pipeline_counter:", incr.Val())

	for _, cmd := range cmds {
		fmt.Printf("cmd: %#v \n", cmd)
	}

}

func pipeline3() {
	rdb, err := RDBClient()
	if err != nil {
		panic(err)
	}

	// 使用 TxPipeline 批量执行多个命令
	// 以下代码就相当于执行了
	// MULTI
	// INCR pipeline_counter
	// EXPIRE pipeline_counter 10
	// EXEC
	pipe := rdb.TxPipeline()
	incr := pipe.Incr("pipeline_counter")
	pipe.Expire("pipeline_counter", 10*time.Second)
	_, err = pipe.Exec()
	if err != nil {
		panic(err)
	}

	// 获取 incr 命令的执行结果
	fmt.Println("pipeline_counter:", incr.Val())
}

func pipeline4() {
	rdb, err := RDBClient()
	if err != nil {
		panic(err)
	}

	var incr *redis.IntCmd

	// 以下代码就相当于执行了
	// MULTI
	// INCR pipeline_counter
	// EXPIRE pipeline_counter 10
	// EXEC
	_, err = rdb.TxPipelined(func(pipe redis.Pipeliner) error {
		incr = pipe.Incr("pipeline_counter")
		pipe.Expire("pipeline_counter", 10*time.Second)
		return nil
	})
	if err != nil {
		panic(err)
	}

	// 获取 incr 命令的执行结果
	fmt.Println("pipeline_counter:", incr.Val())
}

func watchDemo() {
	rdb, err := RDBClient()
	if err != nil {
		panic(err)
	}

	// watch_key 为被监视的 key
	// 以下代码模拟了一个事务，其中包含了一个 watch 操作
	// 如果 watch_key 在事务执行期间被修改，那么事务将不会执行
	// 事务执行失败时，返回 redis.TxFailedErr
	// 以下代码相当于执行
	// WATCH watch_key
	// MULTI
	// GET watch_key
	// SET watch_key 1 EX 60
	// MULTI
	key := "watch_key"
	err = rdb.Watch(func(tx *redis.Tx) error {
		num, err := tx.Get(key).Int()
		if err != nil && !errors.Is(err, redis.Nil) {
			return err
		}

		// 模拟并发情况下的数据变更
		time.Sleep(5 * time.Second)

		_, err = tx.TxPipelined(func(pipe redis.Pipeliner) error {
			// 在事务中对 key 的值进行 +1 操作
			pipe.Set(key, num+1, time.Second*60)
			return nil
		})

		return nil
	}, key)

	if errors.Is(err, redis.TxFailedErr) {
		fmt.Println("事务执行失败")
	}
}
