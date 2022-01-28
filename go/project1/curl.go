package main

import (
	"fmt"
	"net/http"
	// "io/ioutil"
	"io"
	"os"
	"strings"
)

// func main() {
// 	for _, url := range(os.Args[1:]) {
// 		resp, err := http.Get(url)
// 		if err != nil {
// 			fmt.Printf("错误啦:%s", err)
// 			os.Exit(1)
// 		}
// 		body, err := ioutil.ReadAll(resp.Body)
// 		resp.Body.Close()
// 		if err != nil {
// 			fmt.Println(err)
// 			os.Exit(1)
// 		}
// 		fmt.Printf("body:%s\n", body)
// 	}
// }

// func main() {
// 	for _, url := range(os.Args[1:]) {
// 		resp, err := http.Get(url)
// 		if err != nil {
// 			fmt.Printf("错误啦:%s", err)
// 			os.Exit(1)
// 		}
// 		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
// 			fmt.Printf("error:%s", err)
// 			os.Exit(0)
// 		}
// 		resp.Body.Close()
// 	}
// }

func main() {
	for _, url := range(os.Args[1:]) {
		has := strings.HasPrefix(url, "http://")
		if !has {
			url = "http://" + url
		}
		fmt.Println(url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("错误啦:%s", err)
			os.Exit(1)
		}
		fmt.Println(resp.Status)
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Printf("error:%s", err)
			os.Exit(0)
		}
		resp.Body.Close()
	}
}