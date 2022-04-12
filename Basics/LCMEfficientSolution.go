package main

import (
	"fmt"
)

func main() {
	LCMEfficientSolution(3, 4)
}

func LCMEfficientSolution(a, b int) {
	gcd := GCDUsingEuclideanAlgorithmRec(a, b)  
  fmt.Printf("LCM of (%d,%d) is :%d\n", a, b, (a*b)/gcd)  //lcm(a,b)=(a*b)/gcd(a,b)
}

func GCDUsingEuclideanAlgorithmRec(a, b int) int {
	if b == 0 {
		return a
	}
	return GCDUsingEuclideanAlgorithmRec(b, a%b)

}
