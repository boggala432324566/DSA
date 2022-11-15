package main

import (
	"fmt"
)

func main() {
	var ans []int
	input := []int{1, 2, 4}
	fre := make([]int, len(input))
	for i := 0; i < len(input); i++ {
		fre[i] = 0
	}
	Permutations(input, ans, fre)
}

func Permutations(arr []int, ans []int, fre []int) {
	if len(ans) == len(arr) {
		fmt.Println(ans)
		return
	}
	for i := 0; i < len(arr); i++ {
		if fre[i] != 1 {
			fre[i] = 1
			ans = append(ans, arr[i])
			Permutations(arr, ans, fre)
			ans = ans[:len(ans)-1]
			fre[i] = 0
		}
	}

}
