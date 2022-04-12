package main

import (
	"fmt"	
)

func main() {
	BinaryRepresentationUsingRecursion(12)
}

func BinaryRepresentationUsingRecursion(n int) {
	if n == 0 {
		return
	}
	BinaryRepresentationUsingRecursion(n / 2)
	fmt.Print(n % 2)

}
