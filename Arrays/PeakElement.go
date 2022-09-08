package main

import (
	"fmt"
)

func main() {
	var arr = []int{7, 2, 9, 8, 5, 3}
	peak := peakElement(arr)
	fmt.Printf("Peak Element:%d\n", peak)
}

func peakElement(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	peakEle := -1
	length := len(arr)
	for i := 1; i < (length - 1); i++ {
		if arr[i] > arr[i+1] && arr[i] > arr[i-1] {
			peakEle = arr[i]
			break
		}

	}
	return peakEle
}
