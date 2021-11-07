/*
类的多态
*/
package main

import "fmt"

// interface 本质就是一个指针
type AnimalIF interface {
	Sleep()
	GetColor() string // 获取动物的颜色
	GetType() string  // 获取动物的种类
}

// 定义一个具体的猫类
type Cat struct {
	color string // 猫的颜色
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is Sleep")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

// 定义一个具体的狗类
type Dog struct {
	color string // 狗的颜色
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is Sleep")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

// 定义一个方法获取接口的实现
func showAnimal(animal AnimalIF) {
	animal.Sleep()
	fmt.Println("color = ", animal.GetColor())
	fmt.Println("kind = ", animal.GetType())
}

func main() {
	/*
		var animal AnimalIF  // 接口的数据类型，父类指针

		animal = &Cat{"Green"}
		animal.Sleep()  // 调用的就是 Cat 的 Sleep() 方法

		animal = &Dog{"Yellow"}
		animal.Sleep()  // 调用的就是 Dog 的 Sleep() 方法
	*/

	cat := Cat{"Green"}
	dog := Dog{"Yellow"}

	showAnimal(&cat)
	/*
	Cat is Sleep
	color =  Green
	kind =  Cat
	*/

	showAnimal(&dog)
	/*
	Dog is Sleep
	color =  Yellow
	kind =  Dog
	*/

}
