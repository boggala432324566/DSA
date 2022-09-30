package main

import (
	"fmt"
	"math"
)

var arr [10]int
var seg [4 * 10]int

func SegmentTreeBuildMax(ind int, low int, high int) {
	if low == high {
		seg[ind] = arr[low]
		return
	}
	mid := (low + high) / 2
	SegmentTreeBuildMax((2*ind)+1, low, mid)
	SegmentTreeBuildMax((2*ind)+2, mid+1, high)
	if seg[(2*ind)+1] > seg[(2*ind)+2] {
		seg[ind] = seg[(2*ind)+1]
	} else {
		seg[ind] = seg[(2*ind)+2]
	}
}
func FindMax(ind int, low int, high int, l int, r int) int {
	if low >= l && high <= r {
		return seg[ind]
	}
	if high < l || r < low {
		return math.MinInt
	}
	mid := (low + high) / 2
	left := FindMax((2*ind)+1, low, mid, l, r)
	right := FindMax((2*ind)+2, mid+1, high, l, r)
	if left > right {
		return left
	}
	return right
}

func SegmentTreeBuildMin(ind int, low int, high int) {
	if low == high {
		seg[ind] = arr[low]
		return
	}
	mid := (low + high) / 2
	SegmentTreeBuildMin((2*ind)+1, low, mid)
	SegmentTreeBuildMin((2*ind)+2, mid+1, high)
	if seg[(2*ind)+1] < seg[(2*ind)+2] {
		seg[ind] = seg[(2*ind)+1]
	} else {
		seg[ind] = seg[(2*ind)+2]
	}
}
func FindMin(ind int, low int, high int, l int, r int) int {
	if low >= l && high <= r {
		return seg[ind]
	}
	if high < l || r < low {
		return math.MaxInt
	}
	mid := (low + high) / 2
	left := FindMin((2*ind)+1, low, mid, l, r)
	right := FindMin((2*ind)+2, mid+1, high, l, r)
	if left < right {
		return left
	}
	return right
}

func main() {
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}
	//SegmentTreeBuildMax(0, 0, len(arr)-1)
	//fmt.Println(FindMax(0, 0, len(arr)-1, 3, 8))
	SegmentTreeBuildMin(0, 0, len(arr)-1)
	fmt.Println(FindMin(0, 0, len(arr)-1, 3, 8))
}
