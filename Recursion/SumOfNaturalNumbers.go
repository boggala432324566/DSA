package main

import (
	"fmt"
)

func main() {	
	sum := SumOfNaturalNumbers(5)
	fmt.Println("Sum Of Natural Numbers:", sum)
}

func SumOfNaturalNumbers(n int) int {
	if n == 0 {
		return 0
	}
	return n + SumOfNaturalNumbers(n-1)
}
