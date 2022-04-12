package main

import (
	"fmt"	
)

func main() {
	fact(4)
}

func fact(n int) {
	f := 1
	for i := 1; i <= n; i++ {
		f = f * i
	}
	fmt.Printf("Factorial of %d is:%d\n", n, f)
}
