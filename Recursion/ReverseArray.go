package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	Reverse(arr, 0, len(arr)-1)
	for _, v := range arr {
		fmt.Println(v)
	}
}

func Reverse(arr []int, l, r int) {
	if l > r {
		return
	}
	arr[r], arr[l] = arr[l], arr[r]
	l++
	r--
	Reverse(arr, l, r)
}
