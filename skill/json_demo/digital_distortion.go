package json_demo

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
)

type User struct {
	// 有了 string 后，序列化时 user_id 会被转为字符串
	// 反序列化时，json 字符串中的 user_id 会被转为 int64 类型
	UserID int64  `json:"user_id,string"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
}

func DigitalDistortionDemo() {
	// 一般后端都会将后端的数据进行序列化为 json 字符串格式给到前端
	data := User{
		UserID: math.MaxInt64, // 故意设置为最大值
		Name:   "Alex",
		Age:    18,
	}
	// json 序列化
	b, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("json marshal failed: %v", err)
	}
	// 在结构体中不加 string 时，打印结果如下
	// r1: {"user_id":9223372036854775807,"name":"Alex","age":18}
	// 加了 string 后，打印结果如下
	// r1: {"user_id":"9223372036854775807","name":"Alex","age":18}
	fmt.Printf("r1: %s\n", string(b))

	// 反序列化
	// 因为考虑到我们给出去的 user_id 会超出 js 数字类型的上限，为了保证不会类型失真，我们考虑通过字符串来传输 user_id
	// 那当序列化的 user_id 是一个字符串类型，然而我们后端的结构体中是需要一个 int64 类型时
	// 我们只需要在结构中添加一个 string
	s := `{"user_id":"9223372036854775807","name":"Alex","age":18}`
	var user User
	if err := json.Unmarshal([]byte(s), &user); err != nil {
		log.Fatalf("json unmarshal failed: %v", err)
	}
	// r2: {UserID:9223372036854775807 Name:Alex Age:18}
	fmt.Printf("r2: %+v\n", user)
}
