package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("file/test.txt")
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	b := []byte("test")
	file.Write(b)
	defer file.Close()
}
