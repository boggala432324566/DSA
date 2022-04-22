package main

import "fmt"

func main() {
	var arr = [5]int{3, 4, 5, 1, 6}
	Search(arr, 6)
}

func Search(arr [5]int, ele int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == ele {
			fmt.Printf("Element %d Found At:%d\n", ele, i)
			break
		}
	}
}
