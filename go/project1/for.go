package main

import (
	"fmt"
	"os"
	// "strings"
)

func main() {
	// var a = 1
	// for i:=0; i< 10; i++ {
	// 	fmt.Println(i)
	// }

	// for {
	// 	if(a> 10) {
	// 		break;
	// 	}
	// 	fmt.Println(a)
	// 	a++
	// }
	// fmt.Println(a)
	for index, item := range os.Args {
		
		var temp = item
		if(index == 0) {
			var len = len(item)
			temp = item[len-3:len]
		}
		fmt.Println(temp, index, "\n")
	}
	// var str = strings.Join(os.Args, "-")
	// fmt.Println(str)
}
