package viper_demo

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

// Viper 会按照下面的优先级：
// 1. 显式调用 Set 方法设置的值
// 2. 命令行参数（flag）
// 3. 环境变量
// 4. 配置文件
// 5. key/value 存储（如 etcd、consul）
// 6. 默认值

func GetConfig4YamlFile() {
	viper.SetConfigFile("./config_demo.yaml") // 指定配置文件路径

	// 指定配置文件名（不带扩展名），它只会去找文件名为 config_demo 的文件，
	// 比如如果有 config_demo.yaml 和 config_demo.json 两个文件时，其实这两个文件都有可能会读到
	// 因此，强烈建议不要在同目录下放置多个同名不同后缀的文件，如果存在时，则直接使用 viper.SetConfigFile() 方法
	// viper.SetConfigName("config_demo")

	// 需要注意的是：这个配置基本上是配合远程配置中心使用的，比如 etcd、consul、zookeeper 等，告诉 viper 当前的数据使用什么格式去解析
	// viper.SetConfigType("yaml")        // 指定配置文件类型（专用于从远程获取配置信息时指定配置）

	// viper.AddConfigPath(".")           // 指定查找配置文件的路径（这里使用相对路径）

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

	// 查看某个配置是否存在
	fmt.Printf("一定存在 mysql.host: %v\n", viper.IsSet("mysql.host"))
	fmt.Printf("一定不存在 mysql.host1: %v\n", viper.IsSet("mysql.host1"))

	// 设置默认值
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
