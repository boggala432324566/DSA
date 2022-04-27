package main

import "fmt"

func main() {
	var arr = []int{1, 2, 3, 10, 5, 6}
	reversedArray := ReverseArray(arr)
	fmt.Printf("Array elments:")
	for _, v := range reversedArray {
		fmt.Printf("%d ", v)
	}
}

func ReverseArray(arr []int) []int {
	length := len(arr)
	low := 0
	high := length - 1
	for low < high {
		arr[low], arr[high] = arr[high], arr[low]
		low++
		high--
	}
	return arr
}
