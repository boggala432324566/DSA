package main

import (
	"fmt"	
)

func main() {
	num := 73	
	isPrime := IsPrimeMoreEfficientSolution(num)
	fmt.Printf("%d is Prime:%v\n", num, isPrime)
}

func IsPrimeMoreEfficientSolution(num int) bool {
	if num == 1 {
		return false
	}
	if num == 2 || num == 3 {
		return true
	}
	if num%2 == 0 || num%3 == 0 {
		return false
	}
	for i := 5; i*i < num; i = i + 6 {
		if num%i == 0 && num%(i+2) == 0 {
			return false
		}
	}
	return true
}
