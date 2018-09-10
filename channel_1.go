package main

import "fmt"

func test(c chan int) {
	c <- 10
}

func main() {
	c := make(chan int, 2)
	go test(c)
	a := <-c
	fmt.Println(a)

	// 指定 channel 容量
	d := make(chan int, 2)
	d <- 1
	d <- 2
	fmt.Println(<-d, <-d)
	close(c)
	close(d)
}
