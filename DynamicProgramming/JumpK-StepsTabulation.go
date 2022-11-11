package main

func StepJumpKStepsTabulation(n int, k int) int {
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = -1
	}
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= k; j++ {
			if i >= j {
				dp[i] += dp[i-j]
			}

		}
	}

	return dp[n]
}
