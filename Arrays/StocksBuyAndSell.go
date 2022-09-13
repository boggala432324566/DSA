package main

import (
	"fmt"
)

func main() {
	var arr = []int{2, 4, 5, 3, 7, 1, 8}
	maxProfit := stockBuyAndSell(arr)
	fmt.Printf("Max Profit:%d", maxProfit)
}

func stockBuyAndSell(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	maxProfit := 0
	for i := 1; i < len(arr); i++ {
		if arr[i-1] < arr[i] {
			maxProfit += arr[i] - arr[i-1]
		}
	}
	return maxProfit
}
