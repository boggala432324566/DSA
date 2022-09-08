package main

import (
	"fmt"
)

func main() {
	var arr = []int{7, 2, 9, 8, 5, 3}
	preSum := prefixSum(arr)
	for i := 0; i < len(preSum); i++ {
		fmt.Printf("Prefix Sum Of %d:%d\n", i, preSum[i])

	}
}

func prefixSum(arr []int) []int {
	var preSum []int
	preSum = append(preSum, arr[0])
	length := len(arr)
	for i := 1; i < length; i++ {
		preSum = append(preSum, preSum[i-1]+arr[i])

	}
	return preSum
}
