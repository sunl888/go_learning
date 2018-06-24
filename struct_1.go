package main

import "fmt"

type Skills []int

type Human struct {
	name   string
	age    int
	weight int
}
type Student struct {
	Human
	Skills
	description string
}

func main() {
	mark := Student{Human{"sunlong", 18, 120}, Skills{1, 2, 3, 4}, "hello"}

	zhangsan := new(Student)
	zhangsan.name = "zhangsan"
	zhangsan.age = 99
	zhangsan.description = "world"
	zhangsan.weight = 119
	arr := [5]int{1, 2, 3, 4, 5}
	zhangsan.Skills = arr[2:]
	fmt.Println(mark.Human, zhangsan)
}
