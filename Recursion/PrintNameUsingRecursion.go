package main

import "fmt"

func main() {
	PrintName(1, 5)
}
func PrintName(i int, n int) {
	if i > n {
		return
	}
	fmt.Println("Name")
	PrintName(i+1, n)
}
