/**
结构体调用函数的实质：类型决定其调用
方法实际上是函数的语法糖，函数允许入参为 nil，方法就允许调用者为 nil
*/
package main

import (
	"fmt"
)

type HeroT struct {
}

func (h *HeroT) HelloHeroT() {
	fmt.Println("我是 HelloHeroT 方法")
}

func (h *HeroT) HelloHeroTT(name string) {
	fmt.Println("我是 HelloHeroTT 方法 name = ", name)
}

func HelloHeroT(h *HeroT) {
	fmt.Println("我是111")
}

func HelloHeroTT(h *HeroT, name string) {
	fmt.Println("我是111 name = ", name)
}

func main() {
	var t *HeroT

	fmt.Printf("t 是什么 %+v 类型是什么 %T \n", t, t)

	t.HelloHeroT()
	HelloHeroT(t)

	t.HelloHeroTT("alex")
	HelloHeroTT(t, "alex")

}

/**
t 是什么 <nil> 类型是什么 *main.HeroT
我是 HelloHeroT 方法
我是111
我是 HelloHeroTT 方法 name =  alex
我是111 name =  alex
*/
