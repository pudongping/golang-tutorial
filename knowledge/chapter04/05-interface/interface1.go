package main

import "fmt"

type A interface {
	Foo()
}

// Go 语言也支持类似的「接口继承」特性，但是由于不支持 extends 关键字，所以其实现和类的继承一样，是通过组合来完成的。
type B interface {
	A
	Bar()
}

type T struct {
}

// 在 Go 语言中，又与传统的接口继承有些不同，因为接口实现不是强制的，
// 是根据类实现的方法来动态判定的，比如我们上面的 T 类可以只实现 Foo 方法，
// 也可以只实现 Bar 方法，也可以都不实现。如果只实现了 Foo 方法，则 T 实现了接口 A；
// 如果只实现了 Bar 方法，则既没有实现接口 A 也没有实现接口 B，只有两个方法都实现了系统才会判定实现了接口 B
func (t T) Foo() {
	fmt.Println("call Foo function from interface A.")
}

func (t T) Bar() {
	fmt.Println("call Bar function from interface B.")
}
