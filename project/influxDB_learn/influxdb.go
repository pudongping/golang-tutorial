package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

const (
	// 配置信息
	serverURL   = "http://127.0.0.1:8086"                                                                    // InfluxDB 地址
	token       = "c8RSuTvbVbdqBBKdXKcf12_8VJJhZMYjHET4rBi7Qsa94Zj575JBvyHiBsuSLkHHaHT0r44Mbw1lwujR-8e1xA==" // 认证令牌（需替换为实际值）
	org         = "alex_og"                                                                                  // 组织名称（需替换）
	bucket      = "alex_bk"                                                                                  // 存储桶名称（需替换）
	measurement = "sensor_data"                                                                              // 表名（Measurement）
)

func connInflux() influxdb2.Client {
	// 1. 创建 InfluxDB 客户端
	client := influxdb2.NewClient(serverURL, token)

	return client
}

func GenerateRandomDigit() int {
	num, err := rand.Int(rand.Reader, big.NewInt(9))
	if err != nil {
		return 1
	}
	return int(num.Int64()) + 1
}

func GenerateRandomNumber(length int) (int, error) {
	const digits = "0123456789"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			return 0, err
		}
		result[i] = digits[num.Int64()]
	}

	numInt, err := strconv.Atoi(string(result))
	if err != nil {
		return 0, err
	}
	return numInt, nil
}

func writePoints(client influxdb2.Client) {
	writeAPI := client.WriteAPIBlocking(org, bucket)

	t1, err := GenerateRandomNumber(GenerateRandomDigit())
	if err != nil {
		panic(fmt.Sprintf("生成随机数 temperature 失败: %v", err))
	}
	t2, err := GenerateRandomNumber(GenerateRandomDigit())
	if err != nil {
		panic(fmt.Sprintf("生成随机数 humidity 失败: %v", err))
	}

	// ========== C (Create): 写入数据 ==========
	// 创建一个数据点（Point）
	p := influxdb2.NewPoint(
		measurement,
		map[string]string{"device_id": "d1", "location": "room1"},          // Tags（索引字段，用于快速过滤）
		map[string]interface{}{"temperature": float64(t1), "humidity": t2}, // Fields（实际数据）
		time.Now(), // 时间戳
	)

	// 写入数据点
	if err := writeAPI.WritePoint(context.Background(), p); err != nil {
		panic(fmt.Sprintf("写入失败: %v", err))
	}
	fmt.Println("数据写入成功！")
}

func fetchData(client influxdb2.Client) {
	queryAPI := client.QueryAPI(org)

	// ========== R (Read): 查询数据 ==========
	// 构建 Flux 查询语句（查询最近1小时的数据）
	query := fmt.Sprintf(`from(bucket: "%s")
		|> range(start: -1h)
		|> filter(fn: (r) => r._measurement == "%s" and r.device_id == "d1")`, bucket, measurement)

	// 执行查询
	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		panic(fmt.Sprintf("查询失败: %v", err))
	}

	// 遍历查询结果
	fmt.Println("\n查询结果:")
	for result.Next() {
		record := result.Record()
		fmt.Printf(
			"[%s] %s: %v\n",
			record.Time().Format(time.RFC3339),
			record.Field(), // 字段名（如 temperature）
			record.Value(), // 字段值
		)
		fmt.Println()
	}
	if result.Err() != nil {
		panic(fmt.Sprintf("结果解析错误: %v", result.Err()))
	}
}

func main() {

	client := connInflux()
	defer client.Close() // 确保最后关闭客户端

	// 尝试每隔 500ms 写入一次数据，写 100 次
	// for i := 0; i < 100; i++ {
	// 	writePoints(client)
	// 	time.Sleep(500 * time.Millisecond)
	// }

	// 查询数据
	fetchData(client)
}
