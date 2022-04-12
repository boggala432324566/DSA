package main

import (
	"fmt"
	"math"
)

func main() {	
	LCM(3, 4)
}

func LCM(a, b int) {
	res := int(math.Max(float64(a), float64(b)))
	for {
		if res%a == 0 && res%b == 0 {
			break
		}
		res++
	}
	fmt.Printf("LCM of (%d,%d) is :%d\n", a, b, res)
}
