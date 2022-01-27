package main

import (
	"time"
	"math/rand"
	"fmt"
)

func main () {
	rand.Seed(time.Now().UTC().UnixNano())
	var test = rand.Intn(10)
	fmt.Println(test)
}