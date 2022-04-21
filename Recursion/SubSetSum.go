package main

import "fmt"

func main() {
	subSets := SubSetSum([]int{10, 5, 2, 3, 6}, 5, 8)
	fmt.Printf("SubSets Count:%d\n", subSets)
}

func SubSetSum(arr []int, n int, sum int) int {
	if n == 0 {
		if sum == 0 {
			return 1
		}
		return 0
	}
	return SubSetSum(arr, n-1, sum) + SubSetSum(arr, n-1, sum-arr[n-1])
}
