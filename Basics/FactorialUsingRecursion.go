package main

import (
	"fmt"
)

func main() {	
	num := 4
	f := factRec(num)
	fmt.Printf("Factorial of %d is:%d\n", num, f)
}

func factRec(n int) int {
	if n == 1 {
		return 1
	}
	return n * factRec(n-1)
}
