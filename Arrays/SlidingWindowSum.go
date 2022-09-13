package main

import (
	"fmt"
)

func main() {
	var arr = []int{7, 2, 9, 8, 5, 3, 5, 3, 8, 9, 1}
	slidingWindowSize := 3
	slidingWindowSum(arr, slidingWindowSize)
}

func slidingWindowSum(arr []int, slidingWindowSize int) {
	if len(arr) == 0 {
		return
	}
	sum := 0
	for i := 0; i < slidingWindowSize; i++ {
		sum = sum + arr[i]
	}
	fmt.Printf("Sliding window sum:%d\n", sum)
	for i := 1; i <= len(arr)-slidingWindowSize; i++ {
		sum = sum + arr[i+slidingWindowSize-1] - arr[i-1]
		fmt.Printf("Sliding window sum:%d\n", sum)
	}
}
