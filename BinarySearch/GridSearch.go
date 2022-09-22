package main

import "fmt"

func main() {
	var arr = [4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	index := gridSearch(arr, 15)
	fmt.Printf("Number of ele:%d\n", index)
}

func gridSearch(arr [4][4]int, k int) int {
	if len(arr) == 0 {
		return -1
	}
	row := len(arr)
	col := len(arr[0]) - 1
	rowNum := getRow(arr, 0, row-1, k, col)
	if rowNum == -1 {
		return -1
	}
	l := 0
	r := len(arr[0]) - 1
	for l <= r {
		mid := (l + r) / 2
		if arr[rowNum][mid] == k {
			return k
		}
		if k > arr[rowNum][mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

func getRow(arr [4][4]int, l int, r int, k int, col int) int {
	for l <= r {
		mid := (l + r) / 2
		if arr[mid][0] == k || k <= arr[mid][col] {
			return mid
		}
		if k > arr[mid][0] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}
