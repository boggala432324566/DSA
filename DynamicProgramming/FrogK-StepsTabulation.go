package main

import "math"

func FrogJumpKTabulation(n int, k int) int {
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
		for j := 1; j <= k; j++ {
			if i >= j {
				stepOne := dp[i-j] + int(math.Abs(float64(energy[i]-energy[i-j])))
				stepTwo := dp[i-j] + int(math.Abs(float64(energy[i]-energy[i-j])))
				dp[i] = int(math.Min(float64(stepOne), float64(stepTwo)))
			}

		}
	}
	return dp[n]
}
