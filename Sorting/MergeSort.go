package main

import "fmt"

func main() {
	var arr = []int{2, 4, 8, 3, 20, 1, -8}
	mergeSort(arr, 0, len(arr)-1)
	fmt.Printf("Merge Sort:")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}
}

func mergeSort(arr []int, l int, r int) {
	if l < r {
		mid := (l + r) / 2
		mergeSort(arr, l, mid)
		mergeSort(arr, mid+1, r)
		merge(arr, l, mid, r)
	}
}

func merge(arr []int, left int, mid int, right int) {
	lengthOfLeft := mid - left + 1
	lengthOfRight := right - mid
	lArr := make([]int, lengthOfLeft)
	rArr := make([]int, lengthOfRight)
	for i := 0; i < lengthOfLeft; i++ {
		lArr[i] = arr[i+left]
	}
	for i := 0; i < lengthOfRight; i++ {
		rArr[i] = arr[i+mid+1]
	}
	arrIndex := left
	i := 0
	j := 0
	for i < lengthOfLeft && j < lengthOfRight {
		if lArr[i] <= rArr[j] {
			arr[arrIndex] = lArr[i]
			i++
		} else {
			arr[arrIndex] = rArr[j]
			j++
		}
		arrIndex++
	}
	for i < lengthOfLeft {
		arr[arrIndex] = lArr[i]
		i++
		arrIndex++
	}
	for j < lengthOfRight {
		arr[arrIndex] = rArr[j]
		j++
		arrIndex++
	}
}
