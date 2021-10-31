/*
类的定义、初始化和成员方法
*/
package main

import "fmt"

// 通过结构体来实现类的声明
type Student struct {
	id    uint
	name  string
	male  bool
	score float64
}

// 定义形如 NewXXX 这样的全局函数（首字母大写）作为类的初始化函数
func NewStudent(id uint, name string, score float64) *Student {
	return &Student{
		id:    id,
		name:  name,
		score: score,
	}
}

// 定义成员方法
// 值方法
// 要为 Go 类定义成员方法，需要在 func 和方法名之间声明方法所属的类型（有的地方将其称之为接收者声明）
// 通过在函数签名中增加接收者声明的方式定义了函数所归属的类型，这个时候，函数就不再是普通的函数，而是类的成员方法了
func (s Student) GetName() string {
	return s.name
}

// 指针方法
func (s *Student) SetName(name string) {
	s.name = name
}

// 类似于 php 中调用类的 toString 方法，以字符串格式打印类的实例
// 无需显式调用 String 方法，Go 语言会自动调用该方法来打印
func (s Student) String() string {
	return fmt.Sprintf("{id: %d, name: %s, male: %t, score: %f}", s.id, s.name, s.male, s.score)
}

func main() {
	student := NewStudent(1, "alex", 100)
	fmt.Println("student =", student) // student = {id: 1, name: alex, male: false, score: 100.000000}
	student.SetName("alex pu")
	fmt.Println("student =", student)        // student = {id: 1, name: alex pu, male: false, score: 100.000000}
	fmt.Println("Name =", student.GetName()) // Name = alex pu
}
