package main

import (
	"fmt"
)

func main() {
	var arr = []int{1, 0, 0, 4, 0, 6, 7}
	MoveAllZerosToEnd(arr)
}

func MoveAllZerosToEnd(arr []int) {
	if len(arr) == 0 {
		return
	}
	i := 0
	for j := 0; j < len(arr); j++ {
		if arr[j] != 0 {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	fmt.Print("Array elements:")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}
}
