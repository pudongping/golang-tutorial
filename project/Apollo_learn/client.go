package main

import (
	"fmt"
	"time"

	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"
)

// Apollo 官方提供了一个开发环境
// Demo Environment:
// http://81.68.181.139
// User/Password: apollo/admin
func main() {
	c := &config.AppConfig{
		AppID:          "000111",
		Cluster:        "dev",
		IP:             "http://81.68.181.139:8080",
		NamespaceName:  "application",
		IsBackupConfig: true,
		Secret:         "4843921ea2144c8785f88c6f3404df2c",
	}

	client, err := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})
	if err != nil {
		fmt.Println("初始化Apollo配置失败", err)
		return
	}
	fmt.Println("初始化Apollo配置成功")

	checkKey(c.NamespaceName, client)

	fmt.Println("-------------------")

	cache := client.GetConfigCache(c.NamespaceName)
	value, _ := cache.Get("aps.redis.auth")
	fmt.Println(value)

	time.Sleep(5 * time.Second)
}

func checkKey(namespace string, client agollo.Client) {
	cache := client.GetConfigCache(namespace)
	count := 0
	cache.Range(func(key, value interface{}) bool {
		fmt.Println("key : ", key, ", value :", value)
		count++
		return true
	})
	if count < 1 {
		panic("config key can not be null")
	}
}
