package main

import "fmt"

func main() {
	var arr = [5]int{3, 4, 5, 1, 6}
	LargestNavi(arr)
}

func LargestNavi(arr [5]int) {
	length := len(arr)
	for i := 0; i < length; i++ {
		flag := true
		j := 0
		for j = 0; j < length; j++ {
			if arr[i] < arr[j] {
				flag = false
				break
			}
		}
		if flag == true && j != length {
			fmt.Printf("Largest Element:%d\n", arr[j])
		} else if j == length {
			fmt.Printf("Largest Element:%d\n", arr[j-1])
		}
	}
}
