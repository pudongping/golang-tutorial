package wire_pkg_demo

type A struct {
	Name string
	Age  int
}

func NewA(name string, age int) *A {
	return &A{
		Name: name,
		Age:  age,
	}
}
