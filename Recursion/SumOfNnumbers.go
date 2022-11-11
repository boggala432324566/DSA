package main

import "fmt"

func main() {
	fmt.Println(SumN(5))
}

func SumN(i int) int {
	if i < 1 {
		return 0
	}
	sum := i
	i--
	return sum + SumN(i)
}
