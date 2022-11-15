package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	var ds []int
	SubSeq(arr, ds, 0, len(arr))
}

func SubSeq(arr []int, ds []int, i int, n int) {
	if i == n {
		for _, v := range ds {
			fmt.Printf("%d ", v)
		}
		fmt.Println()

	} else {
		ds = append(ds, arr[i])
		SubSeq(arr, ds, i+1, n)
		ds = ds[:len(ds)-1]
		SubSeq(arr, ds, i+1, n)
	}

}
