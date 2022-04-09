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
	for n != 1 {
		if n%2 != 0 {
			return false
		}
		n = n / 2
	}
	return true
}
