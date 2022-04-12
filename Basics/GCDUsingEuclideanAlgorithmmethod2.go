package main

import (
	"fmt"	
)

func main() {
	gcd := GCDUsingEuclideanAlgorithmRec(4, 5)
	fmt.Printf("GCD is :%d\n", gcd)
}


func GCDUsingEuclideanAlgorithmRec(a, b int) int {
	if b == 0 {
		return a
	}
	return GCDUsingEuclideanAlgorithmRec(b, a%b)
}
