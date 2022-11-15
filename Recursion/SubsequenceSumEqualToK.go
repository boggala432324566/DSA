package main

import "fmt"

func main() {
	arr := []int{1, 2, 1}
	var ds []int
	SubSequenceSumEqualToK(arr, ds, 0, len(arr), 0, 2)
}

func SubSequenceSumEqualToK(arr []int, ds []int, ind int, n int, sum int, k int) {
	if ind == n {
		if sum == k {
			for _, v := range ds {
				fmt.Printf("%d ", v)
			}
			fmt.Println()
		}
	} else {
		ds = append(ds, arr[ind])
		sum = sum + arr[ind]
		SubSequenceSumEqualToK(arr, ds, ind+1, n, sum, k)
		sum = sum - ds[len(ds)-1]
		ds = ds[:len(ds)-1]
		SubSequenceSumEqualToK(arr, ds, ind+1, n, sum, k)
	}

}
