package main

import "fmt"

func PathsToReachEndTabulation(m int, n int) {
	count := make([][]int, m, n)
	for i := 0; i < m; i++ {
		count[i][0] = 1
	}
	for j := 0; j < n; j++ {
		count[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			count[i][j] = count[i-1][j] + count[i][j-1]
		}
	}
	fmt.Println(count[m-1][n-1])
}
