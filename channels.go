package main

func main() {
	a := []int{1, 4, 6, 3, -2, -5, 10, -8}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c //从 c 接收数据
	println(x, y)
}

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total // 发送 total 到 c
}
