package main

import (
	"fmt"
)

func main() {	
	FindOddOccurrenceWithXOR()
}

func FindOddOccurrenceWithXOR() {
	var arr = []int32{2, 2, 3, 5, 5, 6, 6, 7, 7, 7, 7}
	var res int32 = 0
	for _, val := range arr {
		res = res ^ val
	}
	fmt.Println("Odd Occurrence:", res)
}
