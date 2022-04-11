package main

import (
	"fmt"
	"math"
)

func main() {
	PowerSet("abc")
}

func PowerSet(str string) {
	length := len(str)
	counter := math.Pow(2, float64(length))
	for c := 0; c <= int(counter); c++ {
		for i := 0; i < length; i++ {
			if c&(1<<i) != 0 {
				fmt.Printf("%c", str[i])
			}
		}
		fmt.Println()
	}
}
