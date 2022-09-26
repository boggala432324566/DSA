package main

import "fmt"

func main() {
	pascalTriangle(5)
}

func pascalTriangle(numRows int) {
	var arr = make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		arr[i] = make([]int, i+1)
		arr[i][0] = 1
		arr[i][i] = 1
		for j := 1; j < i; j++ {
			arr[i][j] = arr[i-1][j-1] + arr[i-1][j]
		}
	}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
