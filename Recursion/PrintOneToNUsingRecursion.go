package main

import (
	"fmt"	
)

func main() {	
	PrintOneToNUsingRecursion(10)
}

func PrintOneToNUsingRecursion(n int) {
	if n == 0 {
		return
	}
	PrintOneToNUsingRecursion(n - 1)
	fmt.Print(n)
}
