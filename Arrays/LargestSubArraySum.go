package main

import (
	"fmt"
)

func main() {
	var arr = []int{2, 4, -8, 3, -20, 1, 8}
	maxSum := maxSubArraySum(arr)
	fmt.Printf("Max Sub Arry Sum:%d", maxSum)
}

func maxSubArraySum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	sumTillNow := 0
	maxSum := 0
	for i := 1; i < len(arr); i++ {
		sumTillNow = sumTillNow + arr[i]
		if maxSum < sumTillNow {
			maxSum = sumTillNow
		}
		if sumTillNow < 0 {
			sumTillNow = 0
		}
	}
	return maxSum
}
