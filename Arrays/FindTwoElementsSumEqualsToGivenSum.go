package main

import (
	"fmt"
	"sort"
)

type ArrayOrder struct {
	index int
	ele   int
}

func main() {
	var arr = []int{2, 4, 5, 3, 7, 1, 8}
	sum := 10
	twoElementsSumEqualToGivenSum(arr, sum)

}

func twoElementsSumEqualToGivenSum(arr []int, sum int) {
	var arrayEle []ArrayOrder
	if len(arr) == 0 {
		return
	}
	for i, v := range arr {
		aEle := ArrayOrder{
			index: i,
			ele:   v,
		}
		arrayEle = append(arrayEle, aEle)
	}
	sort.Slice(arrayEle, func(i, j int) bool {
		return arrayEle[i].ele < arrayEle[j].ele
	})
	first := 0
	last := len(arrayEle) - 1
	firstEle := ArrayOrder{}
	secondEle := ArrayOrder{}
	for first < last {
		tSum := arrayEle[first].ele + arrayEle[last].ele
		if tSum == sum {
			firstEle = arrayEle[first]
			secondEle = arrayEle[last]
			break
		} else if tSum > sum {
			last--
		} else {
			first++
		}
	}
	fmt.Printf("Indexes i,j:%d,%d", firstEle.index, secondEle.index)
}
