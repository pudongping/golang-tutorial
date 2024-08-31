package json_demo

import (
	"encoding/json"
	"fmt"
	"log"
)

type OmitEmpty2 struct {
	UserID   int64  `json:"user_id"`
	Name     string `json:"name"`
	Password string `json:"-"`
}

// 忽略字段
func OmitEmpty2Demo() {
	u1 := OmitEmpty2{
		UserID:   1,
		Name:     "Alex",
		Password: "123456",
	}
	b, err := json.Marshal(u1)
	if err != nil {
		log.Fatalf("json marshal failed: %v", err)
	}

	// OmitEmpty2Demo: {"user_id":1,"name":"Alex"}
	// 注意 Password 字段没有被输出
	fmt.Printf("OmitEmpty2Demo: %s\n", string(b))
}
