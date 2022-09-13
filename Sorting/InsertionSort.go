package main

import "fmt"

func main() {
	var arr = []int{2, 4, 8, 3, 20, 1, -8}
	insertionSort(arr)

}

func insertionSort(arr []int) {
	if len(arr) == 0 {
		return
	}
	for i := 1; i < len(arr); i++ {
		val := arr[i]
		j := i - 1
		for j >= 0 && val < arr[j] {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = val
	}

	fmt.Printf("Insertion Sort:")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}
}
