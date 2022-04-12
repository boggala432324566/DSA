package main

import (
  "fmt"
)

func main() {
	digitSum := getSum(534)
	fmt.Println("Digits Sum:", digitSum)
}

func getSum(n int) int {
	if n == 0 {
		return 0
	}
	return getSum(n/10) + n%10
}
