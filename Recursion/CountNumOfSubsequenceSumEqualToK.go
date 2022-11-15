package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	fmt.Println(CountNumOfSubsequenceSumEqualToK(arr, 0, len(arr), 0, 2))
}

func CountNumOfSubsequenceSumEqualToK(arr []int, ind int, n int, sum int, k int) int {
	if ind == n {
		if sum == k {
			return 1
		}
		return 0
	}
	sum = sum + arr[ind]
	l := CountNumOfSubsequenceSumEqualToK(arr, ind+1, n, sum, k)
	sum = sum - arr[ind]
	r := CountNumOfSubsequenceSumEqualToK(arr, ind+1, n, sum, k)
	return l + r
}
