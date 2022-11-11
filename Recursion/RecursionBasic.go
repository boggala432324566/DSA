package main

import "fmt"

var count int = 3

func main() {
	PrintNum(0)
}
func PrintNum(n int) {
	if n == count {
		return
	}
	fmt.Println(n)
	n++
	PrintNum(n)
	fmt.Println(n)
}
