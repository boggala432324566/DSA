package main

import "math"

func FrogKJump(n int, k int) int {
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = -1
	}
	energy := make([]int, n+1)
	for i := 0; i <= n; i++ {
		energy[i] = i + 10
	}
	return FJumpK(n, k, dp, energy)
}
func FJumpK(n int, k int, dp []int, energy []int) int {
	if n == 0 {
		return 0
	}
	if dp[n] != -1 {
		return dp[n]
	}
	minCost := math.MaxInt
	for i := 1; i <= k; i++ {
		if n-i >= 0 {
			steps := FJumpK(n-i, k, dp, energy) + int(math.Abs(float64(energy[n]-energy[n-i])))
			dp[n] = int(math.Min(float64(steps), float64(minCost)))
		}
	}
	return dp[n]
}
