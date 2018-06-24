package main

import "fmt"

type ages int

type money float32

type months map[string]int

func main() {
	m := months{
		"January":  31,
		"February": 28,
		"December":31,
	}
	fmt.Println(m)
}
