package main

import (
	"fmt"
	"reflect"
)

type Human struct {
	name  string
	age   int
	phone string
}

type Element interface{}
type List [] Element

func (h Human) String() string {
	return h.name
}

func main() {
	var a interface{}
	var i int = 5
	s := "Hello world"
	a = i
	fmt.Println(a)
	a = s
	fmt.Println(a)
	fmt.Println(test("sun"))
	fmt.Println(test(123))

	bob := Human{"bob", 39, "000-7777-XXX"}
	fmt.Println("This Human is : ", bob)

	// 判断变量类型
	list := make(List, 3)
	list[2] = Human{"bob", 39, "000-7777-XXX"}
	if value, ok := list[2].(Human); ok {
		fmt.Println(value, ok)
	}
	// 反射获取变量类型和值
	t := reflect.TypeOf(a)
	fmt.Println(t)
	x := float64(3.4)
	v := reflect.ValueOf(x)
	fmt.Println(v.Type())
	fmt.Println(v.Kind() == reflect.Float64)
	fmt.Println(v.Float())
	p := reflect.ValueOf(&x)
	m := p.Elem()
	m.SetFloat(1.2)
	fmt.Println(m.Float(), x)

	n := []int{1, 2, 3, 4, 5}
	hello(n)
	c := hello1(1, 2, 3);
	fmt.Println(c)
}

// 空接口可以接受任何类型参数
func test(in ...interface{}) interface{} {
	return in
}
func hello(arr ...[]int) {
	fmt.Println(arr)
}
func hello1(n ...int) int {
	size := len(n);
	sum := 0
	for size > 0 {
		sum += n[size-1]
		size--
	}
	return sum
}
