package main

import "fmt"

func SubSequences(arr []int) {
	path := make([]int, len(arr))
	PrintSubSeq(arr, 0, path)

}
func PrintSubSeq(arr []int, ind int, path []int) {
	if len(arr) == ind {
		for _, v := range path {
			fmt.Printf("%d", v)
		}
		fmt.Println()

	} else {
		PrintSubSeq(arr, ind+1, path)
		path = append(path, arr[ind])
		PrintSubSeq(arr, ind+1, path)
		path = arr[:ind]
	}
}
