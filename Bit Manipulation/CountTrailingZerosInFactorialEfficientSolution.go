package main

import (
	"fmt"	
)

func main() {	
	CountTrailingZerosEfficientSolution(100)
}

func CountTrailingZerosEfficientSolution(n int) {
	count := 0
	for i := 5; i <= n; i = i * 5 {
		count = count + (n / i)
	}
	fmt.Println("Trailing Zeros:", count)
}
