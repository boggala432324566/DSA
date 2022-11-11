package main

import "fmt"

var max int = 0

func MaximumSumOfSubSequences(arr []int) {
	path := make([]int, len(arr))
	MaximumSubSequence(arr, 0, path)
	fmt.Println(max)

}
func MaximumSubSequence(arr []int, ind int, path []int) {
	if len(arr) == ind {
		sum := 0
		for _, v := range path {
			sum = sum + v
		}
		if sum > max {
			max = sum
		}

	} else {
		MaximumSubSequence(arr, ind+1, path)
		path = append(path, arr[ind])
		MaximumSubSequence(arr, ind+1, path)
		path = arr[:ind]
	}
}
