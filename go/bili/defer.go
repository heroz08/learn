package main

import "fmt"

func main() {
	var a = 123
	var b = 233
	defer fmt.Println(a, b)
	a += b
	fmt.Println(a)
}