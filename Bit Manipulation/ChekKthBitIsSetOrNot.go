package main

import (
	"fmt"	
)

func main() {	
	CheckKthBitIsSetOrNot(5, 3)
}

func CheckKthBitIsSetOrNot(n, k int) {
       // Using left shifit operator
	if n&(1<<(k-1)) != 0 {
		fmt.Println("yes")
	} else {
		fmt.Println("No")
	}
	// Using right shifit operator
	/*
	if (n>>(k-1))&1 == 1 {
		fmt.Println("yes")
	} else {
		fmt.Println("No")
	}*/
}
