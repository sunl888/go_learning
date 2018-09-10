package main

import (
	"fmt"
	"runtime"
	"time"
)

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		runtime.Gosched()
	}
}

// 不要通过共享内存来通信，而要通过通信来共享内存。
func main() {
	go say("bbbb")
	say("aaaa")
	time.Sleep(10000)
}
