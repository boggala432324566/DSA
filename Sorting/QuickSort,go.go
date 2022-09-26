package main

import "fmt"

func main() {
	var arr = []int{2, 4, 8, 3, 20, 1}
	quickSort(arr, 0, len(arr)-1)
	fmt.Printf("Merge Sort:")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}
}

func quickSort(arr []int, l int, r int) {
	if l < r {
		pi := partition(arr, l, r)
		quickSort(arr, l, pi-1)
		quickSort(arr, pi+1, r)

	}
}

func partition(arr []int, l int, r int) int {
	pivot := r
	i, j := l, l
	for ; j <= r-1; j++ {
		if arr[j] < arr[pivot] {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[pivot] = arr[pivot], arr[i]
	return i
}
