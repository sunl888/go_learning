package main

import "fmt"

// 類(方法)的繼承
type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human //匿名字段
	school string
}

type Employee struct {
	Human //匿名字段
	company string
}

// 在 human 上面定义了一个 method
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// Employee 的 method 重寫 Human 的 method
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", e.name, e.phone)
}

func main() {
	mark := Student{Human{"Mark", 25, "222-222-YYYY"}, "MIT"}
	sam := Employee{Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}

	mark.SayHi()
	sam.SayHi()
}
