package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"time"
)

func main() {
	h1 := sha256.New()
	h1.Write([]byte("password"))
	fmt.Printf("% x\n", h1.Sum(nil))

	h2 := sha1.New()
	h2.Write([]byte("password"))
	fmt.Printf("% x\n", h2.Sum(nil))

	h3 := md5.New()
	h3.Write([]byte("password"))
	fmt.Printf("%x \n", h3.Sum(nil))

	// salt
	h4:=md5.New()
	h4.Write([]byte("password"))
	h4.Write([]byte(time.Now().String()))
	fmt.Printf("%x \n", h4.Sum(nil))

	fmt.Println(^3)
}
