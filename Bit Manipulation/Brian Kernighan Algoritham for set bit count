package main

import (
	"fmt"	
)

func main() {	
	BrianKernighanAlgorithmForCountSetBits(35)
}

func BrianKernighanAlgorithmForCountSetBits(n int) {
	count := 0
	for n != 0 {
		n = n & (n - 1)
		count++
	}
	fmt.Println("Set Bits Count:", count)
}
