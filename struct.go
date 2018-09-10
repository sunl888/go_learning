package main

import "fmt"

type person struct {
	name string
	age  int
}

// 比较两个人的年龄
func older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, (p1.age - p2.age) * -1
}

func main() {
	p1 := person{"lili", 18}
	p2 := person{"xiaoming", 100}
	var a, b = older(p1, p2)
	fmt.Println(a, b)
}
