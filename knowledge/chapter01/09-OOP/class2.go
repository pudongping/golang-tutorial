/*
面向对象继承
*/
package main

import "fmt"

// 定义一个 Human 父类
type Human struct {
	name string
	sex  string
}

// 定义父类的 Eat 方法
func (this *Human) Eat() {
	fmt.Println("Human.Eat() ....")
}

// 定义父类的 Walk 方法
func (this *Human) Walk() {
	fmt.Println("Human.Walk() ....")
}

// ================== 定义一个子类 =======================

type SuperMan struct {
	Human // SuperMan 类继承了 Human 类的方法

	// 新添加的属性
	level int
}

// 重定义父类的方法 Eat()
func (this *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat() ....")
}

// 子类的新方法
func (this *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly() ....")
}

func (this *SuperMan) Print()  {
	fmt.Println("name =", this.name)
	fmt.Println("sex =", this.sex)
	fmt.Println("level =", this.level)
}

func main() {

	// 实例化一个类
	h := Human{"Alex", "male"}

	h.Eat()
	h.Walk()

	// 定义一个子类对象
	// 第一种定义子类对象的方式：
	// s := SuperMan{Human{"Mark", "female"}, 88}
	// 第二种定义子类对象的方式：
	var s SuperMan
	s.name = "Mark"
	s.sex = "female"
	s.level = 88

	s.Walk() // 调用父类的方法
	s.Eat()  // 调用子类的方法
	s.Fly()  // 调用子类的方法

	s.Print()

}

/*
Human.Eat() ....
Human.Walk() ....
Human.Walk() ....
SuperMan.Eat() ....
SuperMan.Fly() ....
name = Mark
sex = female
level = 88
*/