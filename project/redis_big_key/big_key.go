package redis_big_key

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

func NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       2,
	})

	return client
}

// GenerateRandomString 生成指定大小的随机字符串
func GenerateRandomString(size int) string {
	if 0 <= size {
		return ""
	}
	rand.Seed(time.Now().UnixNano())
	chars := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	result := make([]byte, size)
	for i := 0; i < size; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

// WriteBigKey 写入大 key
func WriteBigKey() {
	start := time.Now()

	client := NewRedisClient()
	// 使用完毕后，关闭连接
	defer client.Close()
	var err error

	// 写入字符串类型的键，大小为 5M
	largeStringValue := GenerateRandomString(5 * 1024 * 1024)
	if err = client.Set("large_string_key", largeStringValue, 0).Err(); err != nil {
		log.Fatalf("写入字符串类型的键失败，错误信息为：%s", err.Error())
	}

	// 写入哈希类型的键，元素个数大于 200
	hashData := make(map[string]interface{})
	var hashLock sync.RWMutex
	for i := 0; i <= 200; i++ {
		field := fmt.Sprintf("field_%d", i)
		value := fmt.Sprintf("value_%d", i)
		hashLock.Lock()
		hashData[field] = value
		hashLock.Unlock()
	}
	if err = client.HMSet("large_hash_key", hashData).Err(); err != nil {
		log.Fatalf("写入哈希类型的键失败，错误信息为：%s", err.Error())
	}

	// 写入列表类型的键，元素个数大于 200
	listData := make([]interface{}, 0, 200)
	for i := 0; i <= 200; i++ {
		listData = append(listData, fmt.Sprintf("value_%d", i))
	}
	if err = client.LPush("large_list_key", listData...).Err(); err != nil {
		log.Fatalf("写入列表类型的键失败，错误信息为：%s", err.Error())
	}

	// 写入集合类型的键，元素个数大于 200
	setData := make([]interface{}, 0, 200)
	for i := 0; i <= 200; i++ {
		setData = append(setData, fmt.Sprintf("%d", i))
	}
	if err = client.SAdd("large_set_key2", setData...).Err(); err != nil {
		log.Fatalf("写入集合类型的键失败，错误信息为：%s", err.Error())
	}

	// 写入有序集合类型的键，元素个数大于 200
	zsetData := make([]redis.Z, 0, 200)
	for i := 0; i <= 200; i++ {
		zsetData = append(zsetData, redis.Z{
			Score:  float64(i),
			Member: fmt.Sprintf("value_%d", i),
		})
	}
	if err = client.ZAdd("large_zset_key", zsetData...).Err(); err != nil {
		log.Fatalf("写入有序集合类型的键失败，错误信息为：%s", err.Error())
	}

	fmt.Println(fmt.Sprintf("写入大 key 总耗时：%s", time.Since(start).String()))
}

// ScanBigKey 扫描大 key
// maxMemory 单位为 b
func ScanBigKey(maxMemory int64) []string {
	if maxMemory <= 0 {
		return nil
	}
	var cursor uint64
	var keys []string
	client := NewRedisClient()
	defer client.Close()
	start := time.Now()
	maxKeys := make([]string, 0, 1000)

	for {
		var err error
		keys, cursor, err = client.Scan(cursor, "*", 1).Result()
		if err != nil {
			log.Fatalf("扫描大 key 失败，错误信息为：%s", err.Error())
		}

		// 检查每个键的内存占用情况
		for _, key := range keys {
			// memory 单位为 byte
			memory, err := client.MemoryUsage(key).Result()
			if err != nil {
				log.Fatalf("获取键 %s 的内存占用失败，错误信息为：%s", key, err.Error())
			}

			// 如果内存占用超过指定最大内存时，则打印出来
			if memory > maxMemory {
				log.Printf("键 %s 的内存占用为 %f MB", key, float64(memory)/(1024*1024))
				maxKeys = append(maxKeys, key)
			}
		}

		// 如果 cursor 为 0，说明已经遍历完成，退出循环
		if cursor == 0 {
			break
		}
	}

	fmt.Println(fmt.Sprintf("扫描大 key 总耗时：%s", time.Since(start).String()))

	return maxKeys
}

func ClearKeys(keys []string) {
	if 0 == len(keys) {
		return
	}
	start := time.Now()
	client := NewRedisClient()
	defer client.Close()

	pipe := client.Pipeline()
	pipe.Unlink(keys...)
	_, err := pipe.Exec()
	if err != nil {
		log.Fatalf("删除 key 失败，错误信息为：%s", err.Error())
	}

	fmt.Println(fmt.Sprintf("删除 key 总耗时：%s", time.Since(start).String()))
}

func WriteKeysToFile(keys []string) error {
	if 0 == len(keys) {
		return nil
	}
	content := strings.Join(keys, "\n")
	err := ioutil.WriteFile("./bigKey.txt", []byte(content), 0644)
	return err
}
