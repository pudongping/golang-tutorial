/*
反射解析结构体标签 tag
*/
package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	// 可以定义多个标签，通过空格做分割
	Name string `info:"name" doc:"我的名字"`
	Sex  string `info:"sex"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()

	for i := 0; i < t.NumField(); i++ {
		taginfo := t.Field(i).Tag.Get("info")
		tagdoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info ==", taginfo, " doc ==", tagdoc)
		// info == name  doc == 我的名字
		// info == sex  doc ==
	}

}

func main() {

	var re resume
	findTag(&re)

}
