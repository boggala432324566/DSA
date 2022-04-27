package main

import "fmt"

func main() {
	var arr = []int{3, 4, 5, 1, 6}
	SecondLargest(arr)
}
func SecondLargest(arr []int) {
	length := len(arr)
	largest := Largest(arr)
	res := -1
	for i := 0; i < length; i++ {
		if arr[i] != largest {
			if res == -1 {
				res = i
			} else if arr[res] < arr[i] {
				res = i
			}
		}
	}
	fmt.Printf("Second Largest Element:%d", arr[res])
}

func Largest(arr []int) int {
	length := len(arr)
	largest := arr[0]
	for i := 0; i < length; i++ {
		if largest < arr[i] {
			largest = arr[i]
		}
	}
	return largest
}
