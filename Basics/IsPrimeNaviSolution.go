package main

import (
	"fmt"	
)

func main() {	
	num := 5
	isPrime := IsPrime(num)
	fmt.Printf("%d is Prime:%v\n", num, isPrime)
}

func IsPrime(num int) bool {
	if num == 1 {
		return false
	}
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
