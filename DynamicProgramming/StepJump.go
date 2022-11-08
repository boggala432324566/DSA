package main

func StepJump(n int) int {
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = -1
	}
	return Jump(n, dp)
}
func Jump(n int, dp []int) int {
	if n <= 1 {
		return n
	}
	if dp[n] != -1 {
		return dp[n]
	}
	dp[n] = Jump(n-1, dp) + Jump(n-2, dp)
	return dp[n]
}
