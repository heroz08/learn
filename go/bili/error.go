package main
import (
	"fmt"
	"errors"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	test()
	fmt.Println("end")
}

func test() {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			// panic(errors.New("test"))
		}
	}()
	n1 := 2
	n2 := 0
	err := errors.New("error la")
	// panic(err)
	result := n1/n2
	fmt.Println(result, err)
}