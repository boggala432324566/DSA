package main

import (
	"fmt"	
)

func main() {
	CountTrailingZeros(20)
}

func factRec(n int64) int64 {
	if n == 1 {
		return 1
	}
	return n * factRec(n-1)
}

func CountTrailingZeros(n int64) {
	count := 0
	fact := factRec(n)
	for fact%10 == 0 {
		count++
		fact = fact / 10
	}
	fmt.Println("Trailing Zeros:", count)
}
