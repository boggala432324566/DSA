package main

import "math"

func LCS(str1 string, str2 string, m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	if str1[m-1] == str2[n-1] {
		return 1 + LCS(str1, str2, m-1, n-1)
	} else {
		return int(math.Max(float64(LCS(str1, str2, m-1, n)), float64(LCS(str1, str2, m, n-1))))
	}

}
