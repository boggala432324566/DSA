package main

import (
	"fmt"	
)

func main() {
	GCDUsingEuclideanAlgorithm(30, 20)
}

func GCDUsingEuclideanAlgorithm(a, b int) {
	for a != b {
		if a > b {
			a = a - b
		} else {
			b = b - a
		}
	}
	fmt.Printf("GCD is :%d\n", a)
}
