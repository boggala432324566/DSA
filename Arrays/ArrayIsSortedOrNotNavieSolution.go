package main

import "fmt"

func main() {
	var arr = []int{1, 2, 3, 10, 5, 6}
	flag := IsSorted(arr)
	if flag {
		fmt.Printf("Array is sorted\n")
	} else {
		fmt.Printf("Array is not sorted\n")
	}
}

func IsSorted(arr []int) bool {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if arr[j] < arr[i] {
				return false
			}
		}
	}
	return true
}
