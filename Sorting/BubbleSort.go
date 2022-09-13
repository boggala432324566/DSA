package main

import "fmt"

func main() {
	var arr = []int{2, 4, 8, 3, 20, 1, -8}
	bubbleSort(arr)

}

func bubbleSort(arr []int) {
	if len(arr) == 0 {
		return
	}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	fmt.Printf("Bubble Sort:")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}
}
