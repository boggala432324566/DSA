package main

import (
	"fmt"	
)

func main() {	
	f := fib(5)
	fmt.Println("Nth fib number:", f)
}

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
