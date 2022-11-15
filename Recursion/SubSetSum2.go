package main

import (
	"fmt"
	"sort"
)

func main() {
	var ans []int
	SubSetSumII(0, 0, []int{1, 2, 3}, 3, &ans)
	sort.Ints(ans)
	fmt.Printf("SubSets sum:%d\n", ans)
}

func SubSetSumII(ind int, sum int, arr []int, n int, ans *[]int) {
	if ind == n {
		*ans = append(*ans, sum)
		return
	}
	SubSetSumII(ind+1, sum+arr[ind], arr, n, ans)
	SubSetSumII(ind+1, sum, arr, n, ans)
}
