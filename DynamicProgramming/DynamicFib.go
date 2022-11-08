package main

func DynamicFib(n int) int {
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = -1
	}
	return Fib(n, dp)
}
func Fib(n int, dp []int) int {
	if n <= 1 {
		return n
	}
	if dp[n] != -1 {
		return dp[n]
	}
	dp[n] = Fib(n-1, dp) + Fib(n-2, dp)
	return dp[n]
}
