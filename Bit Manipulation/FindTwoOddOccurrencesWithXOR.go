package main

import (
	"fmt"	
)

func main() {	
	FindTwoOddOccurrenceWithXOR()
}

func FindTwoOddOccurrenceWithXOR() {
	var arr = []int32{2, 2, 3, 9, 5, 5, 6, 6, 7, 7, 7, 7}
	var xr int32 = 0
	var first, second int32 = 0, 0
	for _, val := range arr {
		xr = xr ^ val
	}
	setBit := xr &^ (xr - 1)
	for _, val := range arr {
		if val&setBit == 0 {
			first = first ^ val
		}
		if val&setBit != 0 {
			second = second ^ val
		}
	}

	fmt.Println("Odd Occurrences:", first, second)

}
