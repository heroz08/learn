package main

import (
	"fmt"
	"time"
)

func main() {
	// now
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(int(now.Month()))
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println("---------------------")
	fmt.Println(now.Format("2006-01-02 15:04:05"))
}