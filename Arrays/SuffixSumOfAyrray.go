package main

import (
	"fmt"
)

func main() {
	var arr = []int{7, 2, 9, 8, 5, 3}
	sufSum := suffixSum(arr)
	for i := len(sufSum) - 1; i >= 0; i-- {
		fmt.Printf("Suffix Sum Of %d:%d\n", len(sufSum)-i-1, sufSum[i])

	}
}

func suffixSum(arr []int) []int {
	var sufSum []int
	sufSum = append(sufSum, arr[len(arr)-1])
	length := len(arr)
	for i := length - 2; i >= 0; i-- {
		sufSum = append(sufSum, sufSum[length-i-2]+arr[i])
	}
	return sufSum
}
