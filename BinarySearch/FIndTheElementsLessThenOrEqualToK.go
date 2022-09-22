package main

import (
	"fmt"
)

func main() {
	var arr = []int{3, 6, 7, 7, 7, 8, 9, 10}
	index := lessThenK(arr, 7)
	fmt.Printf("Number of ele:%d\n", index)
}

func lessThenK(arr []int, k int) int {
	if len(arr) == 0 {
		return -1
	}
	length := len(arr)
	l := 0
	r := length - 1
	for l <= r {
		mid := (l + r) / 2

		if k >= arr[mid] {
			l = mid + 1

		} else {
			r = mid - 1
		}

	}
	return l
}
