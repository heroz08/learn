package main

import "fmt"

var (
	n = 2
	c = "hello"
)

func main() {
	fmt.Println("hello world")
	age()
}

func age() {
	var a = 21
	fmt.Println("age", a)
	var n1, n2, n3 = 1, '2', 2.8
	sex:="男"
	fmt.Println(n1, n2, n3, n, c, sex)
}

// var n1,n2,n3 = 1, '2',2.8

// func log(){
// 	fmt.Println(v)
// }

// log(n1)
