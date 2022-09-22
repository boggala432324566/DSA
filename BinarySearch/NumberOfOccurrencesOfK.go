package main

import (
	"fmt"
)

func main() {
	var arr = []int{3, 6, 7, 8, 9, 10, 10, 10, 10}
	l := lessThenKIndex(arr, 10)
	g := greaterThenKIndex(arr, 10)
	fmt.Printf("Number of occurrences:%d\n", l-g-1)
}

func greaterThenKIndex(arr []int, k int) int {
	if len(arr) == 0 {
		return -1
	}
	length := len(arr)
	l := 0
	r := length - 1
	for l <= r {
		mid := (l + r) / 2

		if k <= arr[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}

	}
	return r
}

func lessThenKIndex(arr []int, k int) int {
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
