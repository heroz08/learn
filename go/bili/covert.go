package main

import (
	"fmt"
)

func main() {
	var n1 int= 100
	var n2 float32 = float32(n1)
	fmt.Println(n1, n2)
	var n3 int32 = 888888
	var n4 int64 = int64(n3)
	fmt.Println(n4)

	var n int = 19
	var ns string = fmt.Sprintf("%d", n)
	var b bool = true
	var bs string = fmt.Sprintf("%t", b)
	var f float32 = 0.77
	var fs string = fmt.Sprintf("%f", f)
	var byte byte = 'c'
	var bytes string = fmt.Sprintf("%c", byte)

	fmt.Println(ns, bs, fs, bytes)
}