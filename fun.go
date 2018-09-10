package main

import (
	ff "fmt"
)

type testInt func(int) bool

func isOdd(num int) bool {
	if num%2 == 0 {
		return false
	}
	return true
}
func isEven(integer int) bool {
	return !isOdd(integer)
}
func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		// 通过运行传进来的函数判断返回值
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	slice := []int{1, 2, 3, 5, 6, 7}
	ff.Println("slice= ", slice)
	odd := filter(slice, isOdd)
	ff.Println("Odd elements of slice are: ", odd)
	even := filter(slice, isEven)
	ff.Println("Odd elements of slice are: ", even)
}
