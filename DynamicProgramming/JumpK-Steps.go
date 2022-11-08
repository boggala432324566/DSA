package main

import "math"

func JumpKSteps(n int, k int) int {
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = -1
	}
	return JumpK(n, k, dp)
}

func JumpK(n int, k int, dp []int) int {
	if n <= 1 {
		return n
	}
	if dp[n] != -1 {
		return dp[n]
	}
	minJumps := math.MaxInt
	for i := 1; i <= k; i++ {
		dp[n] = int(math.Min(float64(JumpK(n-i, k, dp)), float64(minJumps)))
	}
	return dp[n]
}
