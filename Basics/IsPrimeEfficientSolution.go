package main

import (
	"fmt"
)

func main() {	
	num := 73
	isPrime := IsPrimeEfficientSolution(num)
	fmt.Printf("%d is Prime:%v\n", num, isPrime)
}

func IsPrimeEfficientSolution(num int) bool {
	if num == 1 {
		return false
	}
	for i := 2; i*i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
