package main

import (
	"fmt"
)

func main() {	
	f := factorial(5)
	fmt.Println("Factorial:", f)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
