package main
 

import "fmt"

func chanage(n1 *int, n2 *int)(b1 int, b2 int) {
	var t int
	t = *n1
	*n1 = *n2
	*n2 = t
	b1 = *n1
	b2 = *n2
	return 
}

func main () {
	var n = 1
	var b = 2
	fmt.Println(n,b)
	nn,bb := chanage(&n, &b);
	fmt.Println(nn, bb)
}