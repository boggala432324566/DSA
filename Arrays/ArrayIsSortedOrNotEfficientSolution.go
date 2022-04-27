package main

import "fmt"

func main() {
	var arr = []int{1, 2, 3, 4, 5, 6}
	flag := IsSorted(arr)
	if flag {
		fmt.Printf("Array is sorted\n")
	} else {
		fmt.Printf("Array is not sorted\n")
	}
}

func IsSorted(arr []int) bool {
	length := len(arr)
	for i := 1; i < length; i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}
	return true
}
