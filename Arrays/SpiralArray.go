package main

import "fmt"

func main() {
	var arr = [6][6]int{}
	spiralArray(arr)
}

func spiralArray(arr [6][6]int) {
	count := 1
	top := 0
	right := len(arr[0]) - 1
	left := 0
	buttom := len(arr) - 1
	for top <= buttom && left <= right {
		for i := left; i <= right; i++ {
			arr[top][i] = count
			count++
		}
		top++
		for i := top; i <= buttom; i++ {
			arr[i][right] = count
			count++
		}
		right--
		for i := right; i >= left; i-- {
			arr[buttom][i] = count
			count++
		}
		buttom--
		for i := buttom; i >= top; i-- {
			arr[i][left] = count
			count++
		}
		left++
	}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d", arr[i])
		fmt.Println()
	}
}
