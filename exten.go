package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}
type Student struct {
	Human  //匿名字段
	school string
}
type Employee struct {
	Human   //匿名字段
	company string
}

func (h *Human) SayHai() {
	fmt.Println("Hi, I am", h.name)
}

// 方法重写
func (h *Employee) SayHai() {
	fmt.Println("Hi, my name is", h.name)
}
func main() {
	mark := Student{Human{
		"mark", 25, "15705547511",
	}, "MIT"}
	sam := Employee{Human{
		"sam", 15, "15705547512",
	}, "Golang Inc"}
	// 方法的继承
	mark.SayHai()
	sam.SayHai()
}
