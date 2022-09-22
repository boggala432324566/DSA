package main

import (
	"fmt"
)

func main() {
	var arr = []int{3, 6, 7, 8, 9, 10}
	index := binarySearch(arr, 10)
	fmt.Printf("Element At Index:%d\n", index)
}

func binarySearch(arr []int, ele int) int {
	if len(arr) == 0 {
		return -1
	}
	length := len(arr)
	l := 0
	r := length - 1
	for l <= r {
		mid := (l + r) / 2
		if arr[mid] == ele {
			return mid
		}
		if ele > arr[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}

	}
	return -1
}
