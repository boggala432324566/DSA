package main

import (
	"fmt"
	"math"
)

func main() {
	var arr = []int{7, 2, 9, 8, 5, 3}
	min, max := minMax(arr)
	fmt.Printf("Min:%d\nMax:%d", min, max)
}

func minMax(arr []int) (int, int) {
	if len(arr) == 0 {
		return -1, -1
	}
	min, max := math.MaxInt, math.MinInt
	length := len(arr)
	for i := 0; i < length; i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min, max
}
