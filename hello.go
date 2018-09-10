package main

import (
	"fmt"
)

//func say(s string) {
//	for i := 0; i < 5; i++ {
//		runtime.Gosched()
//		fmt.Println(s)
//	}
//}
func main() {
	//go say("world")
	//say("hello")
	//var a [3]int
	//fmt.Println(a)
	//var b = [...]int{1, 2, 3}
	//fmt.Println(b)
	//var c = [...]int{2: 3, 1: 2}
	//fmt.Println(c)
	//var d = [...]int{1, 2, 4: 5, 6}
	//fmt.Println(d)
	//var aa = 1;
	//fmt.Println(aa)
	//var bb, cc, dd = 1, 2, 3;
	//fmt.Println(bb, cc, dd)
	//bbb, ccc, ddd := 4, 'a', 6
	//fmt.Println(bbb, ccc, ddd)
	//ab := "hello world"
	//fmt.Println(ab)
	//_, h := 1, 2
	//fmt.Println(h)
	//const PI = 3.141592653
	//fmt.Println(PI)
	//var enabled, _ = true, false
	//enabled = false;
	//if enabled {
	//	fmt.Println("hello")
	//} else {
	//	fmt.Println("world")
	//}
	//var t rune;
	//t = 123
	//fmt.Println(t)
	//var abc byte = 'w'
	//fmt.Println(abc)
	//var cccc complex64 = 5 + 0i
	//var dddd complex64 = 5 + 0i
	//fmt.Printf("%v", cccc+dddd)
	//fmt.Printf("\n-------------\n")
	//var x complex128 = complex(1, 2) // 1+2i
	//var y complex128 = complex(3, 4) // 3+4i
	//fmt.Println(x * y)               // "(-5+10i)"
	//fmt.Println(real(x * y))         // "-5"
	//fmt.Println(imag(x * y))         // "10"
	//fmt.Println(x + y)
	//
	//s := "hello"
	//cb := []byte(s)  // 将字符串 s 转换为 []byte 类型
	//cb[0] = 'c'
	//s2 := string(cb)  // 再转换回 string 类型
	//fmt.Printf("%s\n", s2)
	//
	//s1 := "hello"
	//s1 = "c" + s1[1:] // 字符串虽不能更改，但可进行切片操作
	//fmt.Printf("%s\n", s1)

	//fmt.Println(a, b, c, d, e, f, g)

	//	var b = []int{ 1,2,3};
	//var arr [10]int;
	//arr[0] = 1
	//fmt.Println(arr)
	//b := [3]int{1, 2, 3}
	//fmt.Println(b)
	//
	//doubleArr := [2][4]int{[4]int{4, 5}, [4]int{1, 3}}
	//fmt.Println(doubleArr);
	//easyArr := [...][3]int{{1, 2, 3}, {4, 5, 6}}
	//fmt.Println(easyArr)
	//var fslice []int
	slice := []byte{'a', 'b', 'c', 'd'}
	slice = []byte{1, 2, 3}
	fmt.Println(slice)
	// 区别
	array := []int{1, 2}
	array = []int{1}
	fmt.Println(array)
	// 声明一个含有10个元素,元素类型为byte的数组
	var ar = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	// 声明两个含有byte的slice
	var a, b []byte
	// a指向数组的第3个元素开始，并到第五个元素结束，
	a = ar[2:5]
	//现在a含有的元素: ar[2]、ar[3]和ar[4]
	// b是数组ar的另一个slice
	b = ar[3:5]
	// b的元素是：ar[3]和ar[4]
	fmt.Println(a, b)
	b[0] = 1 // slice是对数组的引用
	//fmt.Println(ar)
	a = b[0:7]
	//fmt.Println(cap(b), a, b)
	a = append(b, 2)
	fmt.Println(a, ar)
	slice = ar[2:4:5]
	fmt.Println(slice, cap(slice))
	nums := make(map[string]int)
	// not of thread safe
	nums["a"] = 1
	nums["b"] = 2
	nums["c"] = 3
	fmt.Println(nums, len(nums))
	rating := map[string]float32{
		"C":      5,
		"Go":     10,
		"Python": 1.5,
		"C++":    2,
	}
	csarpRating, ok := rating["C++"]
	if ok {
		fmt.Println("This value is ", csarpRating)
	} else {
		fmt.Println("NO")
	}
	delete(rating, "C")
	// map 也是引用 和 slice 一样
	m := make(map[string]string)
	m["Hello"] = "Bonjour"
	m1 := m
	m1["Hello"] = "Salut"
	fmt.Println(m)

	arr := [10]int{1, 2, 3}
	t1 := make([]int, 1)
	t1 = arr[:3]
	fmt.Println(t1)

	sum := 1
	for sum < 1000 {
		sum += sum // n+n <1000
		fmt.Println(sum)
	}

	for k, v := range m1 {
		fmt.Println("map's key:", k)
		fmt.Println("map's val:", v)
	}
	for k := range m1 {
		fmt.Println("map's key:", k)
	}
	for _, v := range m1 {
		fmt.Println("map's val:", v)
	}

	i := 'a'
	switch i {
	case 1:
		fmt.Println(1)
	case 2, 3, 4:
		fmt.Println(2, 3, 4)
	case 'a':
		fmt.Println('a')
		fallthrough //默认相当于每个case最后带有break,但是可以使用fallthrough强制执行后面的case代码
	default:
		fmt.Println("Error")
	}
	fmt.Println(calc(1, -2))
	fmt.Println(myMax(10.5, -2.4, 3.9999))
	bb := 4
	fmt.Println(inc(&bb), bb)
	//defer fmt.Println("closed 1")
	//defer fmt.Println("closed 2")
	//defer fmt.Println("closed 3")
	//fmt.Println("test")
	// defer 类似栈 先进后出
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
	defer_call()

}

// auto run
func init() {
	fmt.Println("hello world11")
}
func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
}
func calc(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func inc(a *int) int {
	*a += 1
	return *a
}
func myMax(slice ...float32) float32 {
	maxValue := slice[0]
	for _, x := range slice {
		if x > maxValue {
			maxValue = x
		}
	}
	return maxValue
}

const (
	a       = iota //a=0
	b       = "B"
	c       = iota             //c=2
	d, e, f = iota, iota, iota //d=3,e=3,f=3
	g       = iota             //g = 4
)
