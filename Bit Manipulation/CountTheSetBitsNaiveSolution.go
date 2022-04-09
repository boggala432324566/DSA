package main

import (
	"fmt"
)

func main() {
	CountSetBits(7)
}

func CountSetBits(n int) {
	count := 0
	for n != 0 {
		if n%2 != 0 {
			count++
		}
		n = n / 2
	}
	fmt.Println("Set Bits Count:", count)
}
