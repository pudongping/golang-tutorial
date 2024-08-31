package json_demo

import (
	"encoding/json"
	"fmt"
	"log"
)

type Profile struct {
	Website string `json:"website"`
	Email   string `json:"email"`
}

type OmitEmpty1 struct {
	Name     string   `json:"name"`
	Age      int      `json:"age,omitempty"`
	Hobby    []string `json:"hobby,omitempty"`
	*Profile `json:"profile,omitempty"`
}

// 忽略嵌套结构体空值字段
func OmitEmpty1Demo() {
	// 如果字段的值为空，则不输出到 json 字符串中
	data := OmitEmpty1{
		Name: "Alex",
		Age:  18,
	}
	b, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("json marshal failed: %v", err)
	}
	// OmitEmpty1Demo: {"name":"Alex","age":18}
	fmt.Printf("OmitEmpty1Demo: %s\n", string(b))
}
