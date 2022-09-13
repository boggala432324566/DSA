package main

import (
	"fmt"
)

func main() {
	var arr = []int{1, 3, 4, 2}
	n := 5
	missingNumber(arr, n)
}

func missingNumber(arr []int, n int) {
	if len(arr) == 0 {
		return
	}
	expectedSum := n * (n + 1) / 2
	actualSum := 0
	for i := 0; i < len(arr); i++ {
		actualSum = actualSum + arr[i]
	}
	fmt.Printf("Missing num:%d\n", expectedSum-actualSum)

}
