package main

import "fmt"

func main() {
	var arr = []int{3, 4, 5, 1, 6}
	SecondLargest(arr)
}
func SecondLargest(arr []int) {
	length := len(arr)
	largest := 0
	res := -1
	for i := 0; i < length; i++ {
		if arr[i] > arr[largest] {
			res = largest
			largest = i
		} else if arr[i] != arr[largest] {
			if res == -1 || arr[res] < arr[i] {
				res = i
			}
		}
	}
	fmt.Printf("Second Largest Element:%d", arr[res])
}
