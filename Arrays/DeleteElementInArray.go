package main

import "fmt"

func main() {
	var arr = [5]int{3, 4, 5, 1, 6}
	DeleteEle(arr, 5)
}

func DeleteEle(arr [5]int, ele int) {
	length := len(arr)
	i := 0
	for i = 0; i < length; i++ {
		if arr[i] == ele {
			break
		}
	}
	if i == length {
		return
	}
	for j := i; j < length-1; j++ {
		arr[j] = arr[j+1]
	}
	fmt.Printf("Array Elements\n")
	for j := 0; j < length-1; j++ {
		fmt.Printf("%d", arr[j])
	}
}
