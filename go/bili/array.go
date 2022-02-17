package main

import (
	"fmt"
)

func main() {
	var a  = [3]int{1,2,3}
	// var b = [...]int{1,2,3,4,5,65}
	// var c = [...]int{3:111, 2:22, 99:233}
	var d[2]int = a
	a[1] = 4
	// for i :=0; i < len(a); i++ {
	// 	fmt.Println("输入成绩")
	// 	fmt.Scanln(&a[i])
	// }



	fmt.Println(a, d)
}