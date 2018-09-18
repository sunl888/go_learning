package main

func main() {
	for i := 0; i < 3; i++ {
		// 闭包对外部捕获的变量采用引用方式访问，因此 i 最终为3
		defer func() { println(i) }()
	}

	for j := 0; j < 3; j++ {
		j := j // 定义一个循环体内局部变量 j
		defer func() {
			println(j)
		}()
	}

	for k := 0; k < 3; k++ {
		// 将迭代变量通过闭包函数的参数传入，defer语句会马上对调用参数求值
		defer func(k int) {
			println(k)
		}(k)
	}
}
