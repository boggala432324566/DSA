package main

import "fmt"

func PathsToReachEnd(m int, n int) {
	fmt.Println(ReachEnd(m, n))
}

func ReachEnd(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	return ReachEnd(m-1, n) + ReachEnd(m, n-1)

}
