package main

import "fmt"

func main() {
	var arr = []int{2, 4, 8, 3, 20, 1, -8}
	selectionSort(arr)

}

func selectionSort(arr []int) {
	if len(arr) == 0 {
		return
	}
	for i := 0; i < len(arr); i++ {
		min, index := arr[i], i
		for j := i + 1; j < len(arr); j++ {
			if min > arr[j] {
				min, index = arr[j], j
			}
		}
		arr[i], arr[index] = arr[index], arr[i]
	}

	fmt.Printf("Selection Sort:")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}
}
