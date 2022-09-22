package main

import (
	"fmt"
)

func main() {
	var arr = []int{3, 6, 7, 7, 7, 8, 9, 10}
	index := firstOccurrence(arr, 7)
	fmt.Printf("Element At Index:%d\n", index)
}

func firstOccurrence(arr []int, ele int) int {
	if len(arr) == 0 {
		return -1
	}
	foundAt := -1
	length := len(arr)
	l := 0
	r := length - 1
	for l <= r {
		mid := (l + r) / 2
		if ele <= arr[mid] {
			foundAt = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return foundAt
}
