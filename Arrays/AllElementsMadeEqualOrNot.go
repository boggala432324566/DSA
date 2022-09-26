package main

import "fmt"

func main() {
	var arr = []int{2, 3, 1}
	flag := AllElementsMadeEqualOrNot(arr)
	if flag {
		fmt.Printf("We can make equal\n")
	} else {
		fmt.Printf("We can't make equal\n")
	}
}

func AllElementsMadeEqualOrNot(arr []int) bool {
	flag := true
	for i := 1; i < len(arr); i++ {
		if arr[0] != arr[i] && (arr[0]+1) != arr[i] && (arr[0]-1) != arr[i] {
			flag = false
		}
	}
	return flag
}
