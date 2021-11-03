/*
结构体标签在 json 中的应用
*/
package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"` // encoding/json 标准库需要提供标签
	Year   int      `json:"year"`
	Price  int      `json:"rmb"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"长津湖", 2021, 70, []string{"吴京", "易洋千玺"}}

	// 编码的过程：结构体 ====> json 字符串
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error", err)
		return
	}

	// jsonStr = {"title":"长津湖","year":2021,"rmb":70,"actors":["吴京","易洋千玺"]}
	fmt.Printf("jsonStr = %s\n", jsonStr)

	// 解码的过程：json 字符串 ====> 结构体
	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("json unmarshal error", err)
		return
	}
	// myMovie = {长津湖 2021 70 [吴京 易洋千玺]}
	fmt.Printf("myMovie = %v\n", myMovie)

}
