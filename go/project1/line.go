package main

import (
	"fmt"
	"os"
	"bufio"
	// "io/ioutil"
	// "strings"
)

// func main () {
// 	var obj = make(map[string]int)
// 	var input = bufio.NewScanner(os.Stdin)
// 	for input.Scan() {
// 		if input.Text() {
// 			break
// 		}
// 		obj[input.Text()]++
// 	}

// 	for index, item := range(obj) {
// 		fmt.Println(item, index)
// 	}
// 	fmt.Println(obj)
// }

func main() {
	var obj = make(map[string]int)
	var files = os.Args[1:]
	if len(files) != 0 {
		for _, file := range(files) {
			f, err := os.Open(file)
			if err != nil {
				continue
			}
			line(f, obj)
			f.Close()
		}
	}
	for index, item := range(obj) {
		fmt.Printf("key:%s,\tnum:%d", index, item)
		fmt.Println(index, "ln")
	}
}

func line(f *os.File, obj map[string]int) {
	var input = bufio.NewScanner(f)
	for input.Scan() {
		if input.Text() == "" {
			break
		}
		obj[input.Text()]++
		if obj[input.Text()] > 1 {
			fmt.Println("some ", input.Text())
		}
	}
}

// func main() {
// 	var obj = make(map[string]int)
// 	var paths = os.Args[1:]
// 	if len(paths) != 0 {
// 		for _, path := range(paths) {
// 			data, err := ioutil.ReadFile(path)
// 			if err != nil {
// 				continue
// 			}
// 			stringData := string(data)
// 			splitData := strings.Split(stringData, "\n")
// 			for _, item := range(splitData) {
// 				obj[item]++
				
// 			}
// 		}
// 	}
// 	for index, item := range(obj) {
// 		fmt.Printf("key: %s, num: %d\n",index, item)
// 	}
// }