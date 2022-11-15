package main

import "fmt"

func main() {
	arr := []int{2, 3, 1, 4}
	var ds []int
	target := 6
	FindCombination(arr, ds, 0, len(arr), target)
}

func FindCombination(arr []int, ds []int, ind int, n int, target int) {
	if ind == n {
		if target == 0 {
			for _, v := range ds {
				fmt.Printf("%d ", v)
			}
			fmt.Println()
		}
	} else {
		if arr[ind] <= target {
			ds = append(ds, arr[ind])
			FindCombination(arr, ds, ind, n, target-arr[ind])
			ds = ds[:len(ds)-1]
		}
		FindCombination(arr, ds, ind+1, n, target)
	}

}
