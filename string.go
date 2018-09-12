package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 字符串s中是否包含substr
	fmt.Println(strings.Contains("seafood", "foo")) //true
	fmt.Println(strings.Contains("seafood", "bar")) //false
	fmt.Println(strings.Contains("seafood", ""))    //true
	fmt.Println(strings.Contains("", ""))           //true
	// 字符串连接
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ","))
	//字符串查找
	fmt.Println(strings.Index("chicken", "ken")) //
	fmt.Println(strings.Index("chicen", "dmr"))  // -1 没有找到
	// 重复字符串n次
	fmt.Println("ba" + strings.Repeat("na", 2))
	// 字符串替换 n表示替换的次数 n小于0表示全部替换
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
	// 字符串切割
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
	// 去除字符串头部和尾部指定的字符串
	fmt.Printf("[%q]", strings.Trim(" !!! Achtung !!! ", "! "))
	//去除字符串的空格符，并按照空格分割返回slice
	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))

	fmt.Println("")

	// Append 系列函数将整数等转换为字符串后，添加到现有的字节数组中。
	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 4567, 10)
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcdefg")
	str = strconv.AppendQuoteRune(str, '单')
	fmt.Println(string(str))
	// Format 系列函数把其他类型的转换为字符串
	a := strconv.FormatBool(false)
	b := strconv.FormatFloat(123.23, 'g', 12, 64)
	c := strconv.FormatInt(1234, 10)
	d := strconv.FormatUint(12345, 10)
	e := strconv.Itoa(1023)
	fmt.Println(a, b, c, d, e)
	// Parse 系列函数把字符串转换为其他类型
	//a, err := strconv.ParseBool("false")
	//checkError(err)
	//b, err := strconv.ParseFloat("123.23", 64)
	//checkError(err)
	//c, err := strconv.ParseInt("1234", 10, 64)
	//checkError(err)
	//d, err := strconv.ParseUint("12345", 10, 64)
	//checkError(err)
	//e, err := strconv.Atoi("1023")
	//checkError(err)
	//fmt.Println(a, b, c, d, e)
}

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
