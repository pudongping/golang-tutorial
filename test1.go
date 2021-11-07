package main

import (
	"fmt"
	"strings"
)

func main()  {

	keywords := []string{"坏蛋", "坏人", "发票", "傻子", "傻大个", "傻人"}
	content := "不要发票，你就是一个傻子，只会发呆"

	for _, keyword := range keywords {
		fmt.Println(keyword)
		content = strings.ReplaceAll(content, keyword, "**")
	}

	fmt.Println(content)

}
