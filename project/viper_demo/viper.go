package viper_demo

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

// Viper 会按照下面的优先级。每个项目的优先级都高于它下面的项目:
// 显示调用 Set 设置值
// 命令行参数（flag）
// 环境变量
// 配置文件
// key/value 存储
// 默认值

func GetConfig4YamlFile() {
	// viper.SetConfigFile("./config_demo.yaml") // 指定配置文件路径
	// 和同时设置配置文件名和配置文件扩展名一样
	viper.SetConfigName("config_demo") // 指定配置文件名（不带扩展名）
	viper.SetConfigType("yaml")        // 指定配置文件类型
	viper.AddConfigPath(".")           // 指定查找配置文件的路径

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Printf("Config file not found: %v\n", err)
			return
		} else {
			// 读取配置文件失败
			fmt.Printf("Failed to read config file: %v\n", err)
			return
		}
	}

	// 建立默认值
	viper.SetDefault("port", 8081)

	// 读取所有的配置信息
	spew.Dump(viper.AllSettings())
	fmt.Println("-------------------")
	fmt.Printf("port: %v\n", viper.Get("port"))
	fmt.Printf("password: %v\n", viper.Get("mysql.password"))

	// 覆盖配置文件中的值
	fmt.Println("-------------------")
	viper.Set("port", 8082)
	fmt.Printf("port: %v\n", viper.Get("port"))
}

func GetConfig4etcd() {
	viper.AddRemoteProvider("etcd", "http://127.0.0.1:4001", "/config/config.yaml")
	viper.SetConfigType("yaml")
	err := viper.ReadRemoteConfig()
	if err != nil {
		panic(err)
	}
}

// GetConfig4Consul 从 consul 获取配置信息
func GetConfig4Consul() {
	// port: 8080
	// env: local
	var err error
	err = viper.AddRemoteProvider("consul", "http://127.0.0.1:8500", "/config/local_config")
	if err != nil {
		panic(err)
	}

	// 设置配置文件的类型
	viper.SetConfigType("yaml")
	// 读取远程配置
	err = viper.ReadRemoteConfig()
	if err != nil {
		if _, ok := err.(viper.RemoteConfigError); ok {
			fmt.Println("远程配置信息没有找到")
			return
		} else {
			panic(err)
		}
	}

	// 读取所有的配置信息
	spew.Dump(viper.AllSettings())
	fmt.Println("-------------------")
	fmt.Printf("port: %v\n", viper.Get("port"))
	fmt.Printf("env: %v\n", viper.Get("env"))
	fmt.Println("-------------------")

	// 这里需要注意，一定要使用 `mapstructure` 这个 tag，否则无法解析
	type cfg struct {
		Port int    `mapstructure:"port"`
		Env  string `mapstructure:"env"`
	}
	var c cfg
	// Unmarshal 将 viper 的配置信息解析到结构体中
	err = viper.Unmarshal(&c)
	if err != nil {
		panic(err)
	}
	spew.Dump(c)

}
