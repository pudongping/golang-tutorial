package main

import (
	"fmt"
	. "go-tutorial/knowledge/chapter04/03-oop-with-type-composite/animal"
)

func main()  {
	animal := NewAnimal("中华田园犬")
	pet := NewPet("宠物狗")
	dog := NewDog(&animal, pet)

	fmt.Println(dog.GetName())  // 宠物狗
	fmt.Println(dog.Call())  // 动物的叫声……汪汪汪
	fmt.Println(dog.FavorFood())  // 爱吃的食物……骨头
}