package main

import (
	"fmt"
	"os"
)

func main() {
	// 最终的权限是 0777 - umask 的值
	os.Mkdir("sunlong", 0777)
	os.MkdirAll("test/test1/test2", 0777)
	err := os.Remove("sunlong")
	if err != nil {
		fmt.Println(err)
	}
	os.RemoveAll("sunlong")
}
