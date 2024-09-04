package json_demo

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"

	"github.com/vmihailenco/msgpack/v5"
)

func JsonEnDeDemo() {
	d1 := make(map[string]interface{})
	d2 := make(map[string]interface{})

	var (
		age    int     = 18
		name   string  = "Alex"
		height float32 = 1.75
	)

	d1["name"] = name
	d1["age"] = age
	d1["height"] = height

	ret, err := json.Marshal(d1)
	if err != nil {
		fmt.Printf("json.Marshal failed: %v\n", err)
		return
	}
	// json.Marshal: {"age":18,"height":1.75,"name":"Alex"}
	fmt.Printf("json.Marshal: %s\n", string(ret))

	err = json.Unmarshal(ret, &d2)
	if err != nil {
		fmt.Printf("json.Unmarshal failed: %v\n", err)
		return
	}
	// json.Unmarshal: map[age:18 height:1.75 name:Alex]
	fmt.Printf("json.Unmarshal: %v\n", d2)

	// 这里我们可以发现一个问题：Go 语言中的 json 包在序列化 interface{} 类型时，会将数字类型（整型、浮点型等）都序列化为 float64 类型
	for k, v := range d2 {
		// key: age, value: 18, type:float64
		// key: height, value: 1.75, type:float64
		// key: name, value: Alex, type:string
		fmt.Printf("key: %s, value: %v, type:%T \n", k, v, v)
	}
}

// gob 序列化示例
// 标准库 gob 是 golang 提供的“私有”的编解码方式，它的效率会比json，xml等更高，特别适合在 Go 语言程序间传递数据
func GobEnDeDemo() {
	d1 := make(map[string]interface{})
	d2 := make(map[string]interface{})

	var (
		age    int     = 18
		name   string  = "Alex"
		height float32 = 1.75
	)

	d1["name"] = name
	d1["age"] = age
	d1["height"] = height

	// encode
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(d1)
	if err != nil {
		fmt.Printf("gob.Encode failed: %v\n", err)
		return
	}
	b := buf.Bytes()
	// gob.Encode:  [13 127 4 1 2 255 128 0 1 12 1 16 0 0 57 255 128 0 3 4 110 97 109 101 6 115 116 114 105 110 103 12 6 0 4 65 108 101 120 3 97 103 101 3 105 110 116 4 2 0 36 6 104 101 105 103 104 116 7 102 108 111 97 116 51 50 8 4 0 254 252 63]
	fmt.Println("gob.Encode: ", b)

	// decode
	dec := gob.NewDecoder(bytes.NewBuffer(b))
	err = dec.Decode(&d2)
	if err != nil {
		fmt.Printf("gob.Decode failed: %v\n", err)
		return
	}
	// gob.Decode: map[age:18 height:1.75 name:Alex]
	fmt.Printf("gob.Decode: %v\n", d2)

	for k, v := range d2 {
		// key: name, value: Alex, type:string
		// key: age, value: 18, type:int
		// key: height, value: 1.75, type:float32
		fmt.Printf("key: %s, value: %v, type:%T \n", k, v, v)
	}

}

func MsgpackEnDeDemo() {
	// msgpack 是一种高效的二进制序列化格式，它允许你在多种语言(如JSON)之间交换数据。但它更快更小。
	// msgpack 序列化示例
	d1 := make(map[string]interface{})
	d2 := make(map[string]interface{})

	var (
		age    int     = 18
		name   string  = "Alex"
		height float32 = 1.75
	)

	d1["name"] = name
	d1["age"] = age
	d1["height"] = height

	// encode
	b, err := msgpack.Marshal(d1)
	if err != nil {
		fmt.Printf("msgpack.Marshal failed: %v\n", err)
		return
	}
	// msgpack.Marshal:  [131 164 110 97 109 101 164 65 108 101 120 163 97 103 101 18 166 104 101 105 103 104 116 202 63 224 0 0]
	fmt.Println("msgpack.Marshal: ", b)

	// decode
	err = msgpack.Unmarshal(b, &d2)
	if err != nil {
		fmt.Printf("msgpack.Unmarshal failed: %v\n", err)
		return
	}
	// msgpack.Unmarshal: map[age:18 height:1.75 name:Alex]
	fmt.Printf("msgpack.Unmarshal: %v\n", d2)

	for k, v := range d2 {
		// key: age, value: 18, type:int8
		// key: height, value: 1.75, type:float32
		// key: name, value: Alex, type:string
		fmt.Printf("key: %s, value: %v, type:%T \n", k, v, v)
	}

}
