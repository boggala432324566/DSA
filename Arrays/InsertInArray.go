package main

import "fmt"

func main() {
	var arr = [6]int{3, 4, 5, 1, 6}
	insert(arr, 5, 6, 9, 2)
}

func insert(arr [6]int, n int, cap int, ele int, pos int) {
	if n == cap {
		return
	}
	idx := pos - 1
	for i := len(arr) - 2; i >= idx; i-- {
		arr[i+1] = arr[i]
	}
	arr[idx] = ele
	fmt.Printf("Array Elements\n")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d", arr[i])
	}
}
