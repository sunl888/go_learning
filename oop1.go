package main

import (
	"math"
	"fmt"
)

// 矩形
type Rectangle struct {
	width, height float64
}

// 園
type Circle struct {
	radius float64
}
// func (r ReceiverType) funcName(parameters) (results)
func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

// 不推薦 這種方法不夠靈活
func area(r Rectangle) float64 {
	return r.width * r.height
}

func main() {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	c1 := Circle{10}
	c2 := Circle{25}
	fmt.Println("矩形面積是 ", area(r1))
	fmt.Println("矩形面積是 ", r1.area())
	fmt.Println("矩形面積是 ", r2.area())
	fmt.Println("園面積是 ", c1.area())
	fmt.Println("園面積是 ", c2.area())
}
