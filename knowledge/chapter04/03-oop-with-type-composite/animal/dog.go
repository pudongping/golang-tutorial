package animal

type Dog struct {
	// 这里可以直接写 *Animal 和 Pet
	// 之所以写成 animal *Animal 是为组合类型设置了别名，方便引用
	animal *Animal // 继承了 Animal
	pet    Pet
}

func NewDog(animal *Animal, pet Pet) Dog {
	return Dog{animal: animal, pet: pet}
}

func (d Dog) FavorFood() string {
	// d.animal 调用父类 Animal 中的 FavorFood 方法
	return d.animal.FavorFood() + "骨头"
}

func (d Dog) Call() string {
	return d.animal.Call() + "汪汪汪"
}

func (d Dog) GetName() string {
	return d.pet.GetName()
}
