package main

import (
	"fmt"	
)

func main() {
	isPowerOfTwo := PowerOfTwo(8)
	fmt.Println("Is Power Of 2:", isPowerOfTwo)
}

func PowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	return n&(n-1) == 0
}
