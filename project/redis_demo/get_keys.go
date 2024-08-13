package redis_demo

import (
	"fmt"
)

func scanKeysDemo1() {
	var cursor uint64
	rdb, err := RDBClient()
	if err != nil {
		panic(err)
	}

	for {
		var keys []string
		var err error
		// Scan 命令用于迭代数据库中的数据库键。
		keys, cursor, err = rdb.Scan(cursor, "*", 0).Result()
		if err != nil {
			panic(err)
		}

		// 处理 keys
		for _, key := range keys {
			fmt.Printf("key: %s\n", key)
		}

		// 如果 cursor 为 0，说明已经遍历完成，退出循环
		if cursor == 0 {
			break
		}
	}

}

func scanKeysDemo2() {
	rdb, err := RDBClient()
	if err != nil {
		panic(err)
	}

	// 针对这种需要遍历大量 key 的场景，go-redis 提供了一个更简单的方法 Iterator
	iter := rdb.Scan(0, "*", 50).Iterator()
	for iter.Next() {
		fmt.Printf("key: %s\n", iter.Val())
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}

	// 此外，对于 redis 中的 set、hash、zset 等类型，也可以使用 Iterator 进行遍历
	// 例如：
	// iter := rdb.SScan("set_key", 0,	"*", 50).Iterator()
	// iter := rdb.HScan("hash_key", 0,	"*", 50).Iterator()
	// iter := rdb.ZScan("zset_key", 0	"*", 50).Iterator()
}
