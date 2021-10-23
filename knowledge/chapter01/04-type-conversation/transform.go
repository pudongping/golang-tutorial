package main

import (
	"fmt"
	"strconv"
)

func main() {

	// 整型 <=> 整型
	v1 := -1
	v2 := uint8(v1)
	fmt.Printf("v1 = %d, type of v1 = %T\n", v1, v1) // v1 = -1, type of v1 = int
	fmt.Printf("v2 = %d, type of v2 = %T\n", v2, v1) // v2 = 255, type of v2 = int

	// 整型 <=> 浮点型
	v3 := 99
	v4 := float32(v3)
	v5 := float64(v3)
	fmt.Printf("v3 = %d, type of v3 = %T\n", v3, v3) // v3 = 99, type of v3 = int

	fmt.Printf("v4 = %f, type of v4 = %T\n", v4, v4) // v4 = 99.000000, type of v4 = float32
	fmt.Printf("v5 = %f, type of v5 = %T\n", v5, v5) // v5 = 99.000000, type of v5 = float64
	// 浮点型 <=> 整型
	v6 := 99.99
	v7 := int(v6)
	fmt.Printf("v6 = %f, type of v6 = %T\n", v6, v6) // v6 = 99.990000, type of v6 = float64
	fmt.Printf("v7 = %d, type of v7 = %T\n", v7, v7) // v7 = 99, type of v7 = int

	// 目前 Go 语言不支持将数值类型转化为布尔型

	// 字符串和其他基本类型之间的转化
	// 整型 <=> 字符串
	v8 := 65
	v9 := string(v8)
	fmt.Printf("v8 = %d, type of v8 = %T\n", v8, v8) // v8 = 65, type of v8 = int
	fmt.Printf("v9 = %s, type of v9 = %T\n", v9, v9) // v9 = A, type of v9 = string
	v10 := 30028
	v11 := string(v10)
	fmt.Printf("v10 = %d, type of v10 = %T\n", v10, v10) // v10 = 30028, type of v10 = int
	fmt.Printf("v11 = %s, type of v11 = %T\n", v11, v11) // v11 = 界, type of v11 = string

	// 数组 <=> 字符串
	v12 := []byte{'h', 'e', 'l', 'l', 'o'}
	v13 := string(v12)
	fmt.Printf("v12 = %s, type of v12 = %T\n", v12, v12) // v12 = hello, type of v12 = []uint8
	fmt.Printf("v13 = %s, type of v13 = %T\n", v13, v13) // v13 = hello, type of v13 = string

	// Go 语言默认不支持将字符串类型强制转化为数值类型，即使字符串中包含数字也不行
	v14 := "100"
	v15, _ := strconv.Atoi(v14)                          // 将字符串转化为整型
	fmt.Printf("v14 = %s, type of v14 = %T\n", v14, v14) // v14 = 100, type of v14 = string
	fmt.Printf("v15 = %d, type of v15 = %T\n", v15, v15) // v15 = 100, type of v15 = int

	v16 := 100
	v17 := strconv.Itoa(v16)                             // 将整型转化为字符串
	fmt.Printf("v16 = %d, type of v16 = %T\n", v16, v16) // v16 = 100, type of v16 = int
	fmt.Printf("v17 = %s, type of v17 = %T\n", v17, v17) // v17 = 100, type of v17 = string

	v18 := "false"
	v19, _ := strconv.ParseBool(v18)                     // 将字符串转化为布尔型
	v20 := strconv.FormatBool(v19)                       // 将布尔值转化为字符串
	fmt.Printf("v18 = %s, type of v18 = %T\n", v18, v18) // v18 = false, type of v18 = string
	fmt.Printf("v19 = %t, type of v19 = %T\n", v19, v19) // v19 = false, type of v19 = bool
	fmt.Printf("v20 = %s, type of v20 = %T\n", v20, v20) // v20 = false, type of v20 = string

	v21 := "100"
	v22, _ := strconv.ParseInt(v21, 10, 64)              // 将字符串转化为整型，第二个参数表示进制，第三个参数表示最大位数
	v23 := strconv.FormatInt(v22, 10)                    // 将整型转化为字符串，第二个参数表示进制
	fmt.Printf("v21 = %s, type of v21 = %T\n", v21, v21) // v21 = 100, type of v21 = string
	fmt.Printf("v22 = %d, type of v22 = %T\n", v22, v22) // v22 = 100, type of v22 = int64
	fmt.Printf("v23 = %s, type of v23 = %T\n", v23, v23) // v23 = 100, type of v23 = string

	v24, _ := strconv.ParseUint(v23, 10, 64)             // 将字符串转化为无符号整型，第二个参数表示进制，第三个参数表示最大位数
	v25 := strconv.FormatUint(v24, 10)                   // 将整型转化为字符串，第二个参数表示进制
	fmt.Printf("v24 = %d, type of v24 = %T\n", v24, v24) // v24 = 100, type of v24 = uint64
	fmt.Printf("v25 = %s, type of v25 = %T\n", v25, v25) // v25 = 100, type of v25 = string

	v26 := "99.99"
	v27, _ := strconv.ParseFloat(v26, 64) // 将字符串转化为浮点型，第二个参数表示精度
	v28 := strconv.FormatFloat(v27, 'E', -1, 64)
	fmt.Printf("v26 = %s, type of v26 = %T\n", v26, v26) // v26 = 99.99, type of v26 = string
	fmt.Printf("v27 = %f, type of v27 = %T\n", v27, v27) // v27 = 99.990000, type of v27 = float64
	fmt.Printf("v28 = %s, type of v28 = %T\n", v28, v28) // v28 = 9.999E+01, type of v28 = string

	v29 := strconv.Quote("Hello, 世界")                    // 为字符串加引号
	v30 := strconv.QuoteToASCII("Hello, 世界")             // 将字符串转化为 ASCII 编码
	fmt.Printf("v29 = %s, type of v29 = %T\n", v29, v29) // v29 = "Hello, 世界", type of v29 = string
	fmt.Printf("v30 = %s, type of v30 = %T\n", v30, v30) // v30 = "Hello, \u4e16\u754c", type of v30 = string

}
