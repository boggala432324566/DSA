package main

import "math"

func FrogJump(n int) int {
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = -1
	}
	energy := make([]int, n+1)
	for i := 0; i <= n; i++ {
		energy[i] = i + 10
	}
	return FJump(n, dp, energy)
}
func FJump(n int, dp []int, energy []int) int {
	if n == 0 {
		return 0
	}
	if dp[n] != -1 {
		return dp[n]
	}
	stepOne := FJump(n-1, dp, energy) + int(math.Abs(float64(energy[n]-energy[n-1])))
	stepTwo := math.MaxInt
	if n > 1 {
		stepTwo = FJump(n-2, dp, energy) + int(math.Abs(float64(energy[n]-energy[n-2])))
	}

	dp[n] = int(math.Min(float64(stepOne), float64(stepTwo)))
	return dp[n]
}
