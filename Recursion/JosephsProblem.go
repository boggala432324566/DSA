package main

import "fmt"

func main() {
	survivor := Josephs(8, 2)
	fmt.Printf("Survivor:%d\n", survivor)// if numbering starts from 1 then print survivor+1
}

func Josephs(n int, k int) int {
	if n == 1 {
		return 0
	}
	return (Josephs(n-1, k) + k) % n
}
