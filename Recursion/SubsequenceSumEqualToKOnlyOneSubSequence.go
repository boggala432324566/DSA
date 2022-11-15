package main

import "fmt"

func main() {
	arr := []int{1, 2, 1}
	var ds []int
	SubSequenceSumEqualToKOnlyOne(arr, ds, 0, len(arr), 0, 2)
}

func SubSequenceSumEqualToKOnlyOne(arr []int, ds []int, ind int, n int, sum int, k int) bool {
	if ind == n {
		if sum == k {
			for _, v := range ds {
				fmt.Printf("%d ", v)
			}
			fmt.Println()
			return true
		}
		return false
	} else {
		ds = append(ds, arr[ind])
		sum = sum + arr[ind]
		if SubSequenceSumEqualToKOnlyOne(arr, ds, ind+1, n, sum, k) {
			return true
		}
		sum = sum - ds[len(ds)-1]
		ds = ds[:len(ds)-1]
		if SubSequenceSumEqualToKOnlyOne(arr, ds, ind+1, n, sum, k) {
			return true
		}
	}
	return false
}
