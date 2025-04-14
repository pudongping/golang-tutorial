package wire_pkg_demo

import (
	"fmt"
)

type B struct {
	Desc string
}

func NewB(a *A) B {
	return B{Desc: fmt.Sprintf("我的名字是 %s, 我今年 %d 岁", a.Name, a.Age)}
}
