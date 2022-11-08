package main

import "math"

func FrogJumpTabulation(n int) int {
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = -1
	}
	energy := make([]int, n+1)
	for i := 0; i <= n; i++ {
		energy[i] = i + 10
	}
	dp[0] = 0
	for i := 1; i <= n; i++ {
		stepOne := dp[i-1] + int(math.Abs(float64(energy[i]-energy[i-1])))
		stepTwo := math.MaxInt
		if i > 1 {
			stepTwo = dp[i-2] + int(math.Abs(float64(energy[i]-energy[i-2])))
		}
		dp[i] = int(math.Min(float64(stepOne), float64(stepTwo)))
	}
	return dp[n]
}
