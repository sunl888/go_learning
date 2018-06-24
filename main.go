package main

import "fmt"

// func main() {
//fmt.Printf("hello world")
//// 變量
//var hello int = 123
//fmt.Printf("\n%d", hello);
//var a, c, b int = 1, 2, 3
//fmt.Printf("%d%d%d", a, b, c)
//var vname1, vname2, vname3 = a, 1, c
//fmt.Printf("%d%d%d", vname1, vname2, vname3)
//v1, v2, v3 := 1, 2, 3
//fmt.Printf("%d%d%d", v1, v2, v3)
//_, bb := 10, 20
//fmt.Printf("%d", bb);
//const C = 123
//const Pi float32 = 3.1415926
//
//var cc complex64 = 5 + 5i
////output: (5+5i)
//fmt.Printf("Value is: %v", cc)
//
//a1, a2 := "\n123\n", `456`
//fmt.Print(a1, a2)

//var s string = "hello1"
//ccc := []byte(s)
//ccc[0] = 'c'
//s2 := string(ccc)
//fmt.Printf("%s \n", string([]byte(s)))
//var m string = `hello2`
//fmt.Printf("%s \n", s + m)
//s := "hello"
//s = "c" + s[1:] // 切片
//fmt.Printf("%s\n", s)
//
//m := `hello
//		world!`
//fmt.Printf("%s\n", m)
//err := errors.New("email is not found!")
//if err != nil {
//	fmt.Print(err)
//}
//bytes := [5]byte{'h', 'e', 'l', 'l', 'o'}
//fmt.Printf("%s\n", bytes)
//primes := [4]int{2, 3, 4, 5}
//fmt.Printf("%d%d%d%d", primes[0], primes[1], primes[2], primes[3])
// const (
// 	i      = 100
// 	pi     = 3.1415
// 	prefix = "go"
// )

// var (
// 	a int
// 	b string
// 	c bool
// )
// a = 1
// b = `hello`
// c = true
// fmt.Println(a, b, c)
// }
//const (
//	x = iota // x == 0
//	y = iota // y == 1
//	z = iota // z == 2
//	w        // 常量声明省略值时，默认和之前一个值的字面相同。这里隐式地说w = iota，因此w == 3。其实上面y和z可同样不用"= iota"
//)
//
//const v = iota // 每遇到一个const关键字，iota就会重置，此时v == 0
//
//const (
//	h, i, j = iota, iota, iota //h=0,i=0,j=0 iota在同一行值相同
//)
//
//const (
//	a       = iota //a=0
//	b       = "B"
//	c       = iota             //c=2
//	d, e, f = iota, iota, iota //d=3,e=3,f=3
//	g       = iota             //g = 4
//)

/*func main() {
	//A, _, aa := 1, 2, 3
	//fmt.Println(a, b, c, d, e, f, g, h, i, j, x, y, z, w, v, aa, A)
	var arr [10] int
	arr[0] = 2
	fmt.Println(arr)

	a := [3]int{1, 2, 3}
	fmt.Println(a)

	b := [...]string{`123`, "456", "789"}
	fmt.Println(b)

	c := [2][2]int{
		[2]int{11, 12},
		[2]int{21, 22},
	}
	fmt.Println(c)

	d := [2][2] int{
		{11, 12},
		{21, 22},
	}
	fmt.Println(d)

	// slice
	slice := [] byte{'a', 'b', 'c', 'd', 'e'}
	slice = [] byte{1, 2, 3}
	fmt.Println(slice)

	var ar = [5]byte{'a', 'b', 'c', 'd', 'e'}
	var a_slice []byte

	a_slice = ar[1:3]
	fmt.Println(a_slice, ar[:3], ar[:])
	ar = [5]byte{1, 2, 3, 4, 5}
	fmt.Println(a_slice, ar[:3], cap(ar[:]), len(ar[3:]))

	var array [10]int
	slice1 := array[2:4]
	fmt.Println(cap(slice1), len(slice1))

	slice2 := array[2:4:5]
	fmt.Println(cap(slice2), len(slice2))

	// map
	var nums map[string]int
	nums = make(map[string]int)
	nums["a"] = 1
	nums["b"] = 2
	nums["c"] = 3
	fmt.Println(nums, nums["a"], nums[`b`], nums[`c`], len(nums))

	rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}
	rating_val, ok := rating["C#"]
	if ok {
		fmt.Println(rating, rating_val)
	} else {
		fmt.Println("error")
	}
	delete(rating, "C")

	m := make(map[string]string)
	m["Hello"] = "Bonjour"
	fmt.Println(m)
	m1 := m
	m1["Hello"] = "Slaut"
	fmt.Println(m)

	if isOk, isNull := false, true; m["Hello"] == "Slaut" {
		isOk = true
		isNull = false
		fmt.Println(isOk, isNull)
	} else {
		fmt.Println(isOk, isNull)
	}
	myGoto()

	fmt.Println(getSum())

	getSlice()

	a1, a2 := Change(1, 2)
	fmt.Println(a1, a2)

	fmt.Println(max(11, 2))

	fmt.Println(SumAndProduct(1, 2))
	x := 2
	l := add(&x)
	fmt.Println(l,x)
}*/
/*func main() {
	x := 2
	l := add(&x)
	fmt.Println(*l,x)
}

// goto
func myGoto() {
	i := 0
Here:
	fmt.Println(i)
	i++
	if i < 10 {
		goto Here
	}
}

// sum
func getSum() int {
	sum := 0
	for index := 0; index < 10; index ++ {
		sum += index
		if index > 5 {
			break
		}
	}
	return sum
}

func getSlice() {
	//map1 := map[string]int{"a": 1, "b": 2,"c":3}
	var map2 map[string]int
	//nums = make(map[string]int)
	for _, v := range map2 {
		fmt.Println(v)
	}
}

// 兩個數交換位置
func Change(a uint, b uint) (uint, uint) {
	a = a ^ b // 1 ^ 2 => 001 ^ 010 = 011
	b = a ^ b // 3 ^ 2 => 011 ^ 010 = 001
	a = a ^ b // 3 ^ 1 => 011 ^ 001 = 010
	return a, b
}

func max(a, b int) int {
	if a > b {
		b = a
	}
	return b
}

func SumAndProduct(A, B int) (add int, Multiplied int) {
	add = A + B
	Multiplied = A * B
	return
}

func getSun(arg ...int) (sum int){
	sum = 0
	for _, n := range arg {
		sum += n
	}
	return
}
func add (a *int) *int {
	*a += 1
	return a
}*/

//简单的一个函数，实现了参数+1的操作
func add(a *int) int { // 请注意，
	*a = *a+1 // 修改了a的值
	return *a // 返回新值
}

func main() {
	x := 3

	fmt.Println("x = ", x)  // 应该输出 "x = 3"
	x1 := add(&x)  // 调用 add1(&x) 传x的地址

	fmt.Println("x+1 = ", x1) // 应该输出 "x+1 = 4"
	fmt.Println("x = ", x)    // 应该输出 "x = 4"

	for i := 0; i < 5; i++ {
		fmt.Println(i+10)
		defer fmt.Printf("%d ", i)
	}
}
