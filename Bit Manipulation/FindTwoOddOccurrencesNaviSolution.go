package main

import (
	"fmt"
)

func main() {	
	FindOddOccurrence()
}

func FindOddOccurrence() {
	var arr = []int32{2, 2, 3, 9, 5, 5, 6, 6, 7, 7, 7, 7}
	for i := 0; i < len(arr); i++ {
		count := 0
		for j := 0; j < len(arr); j++ {
			if arr[i] == arr[j] {
				count++
			}
		}
		if count%2 != 0 {
			fmt.Println("Odd Occurrence:", arr[i])
		}
	}
}
