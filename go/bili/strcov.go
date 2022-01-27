package main 

import (
	"fmt"
	"strconv"
)

func main () {
	var n int = 17
	var ns = strconv.FormatInt(int64(n), 10)
	fmt.Println(ns)
	var f float32 = 0.64
	var fs string = strconv.FormatFloat(float64(f),'f', 4, 32)
	fmt.Println(fs)
	var t bool = false
	var ts string = strconv.FormatBool(t)
	fmt.Println(ts)
}