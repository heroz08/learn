package main

import (
	"fmt"
	"io"
	"os"
	"net/http"
	"time"
	"io/ioutil"
) 

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range(os.Args[1:]) {
		go fetch(url, ch)
	}
	for range(os.Args[1:]) {
		fmt.Println(<-ch)
	}
	ss := time.Since(start).Seconds()
	fmt.Printf("%.2fs elapsed\n", ss )
}

func fetch(url string, ch chan <- string){
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error")
		return
	}
	body, err := io.Copy(ioutil.Discard,resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Println("错误")
		return
	}

	se := time.Since(start).Seconds()

	ch <- fmt.Sprintf("%.2fs %7d %s", se, body, url)


}