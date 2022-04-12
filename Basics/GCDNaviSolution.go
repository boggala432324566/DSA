package main

import (
	"fmt"
	"math"
)

func main() {
	GCD(4, 4)
}

func GCD(a, b int) {
	res := int(math.Min(float64(a), float64(b)))
	for res > 0 {
		if a%res == 0 && b%res == 0 {
			break
		}
		res--
	}
	fmt.Printf("GCD of (%d,%d) is :%d\n", a, b, res)
}
