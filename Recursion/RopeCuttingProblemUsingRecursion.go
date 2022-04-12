package main

import (
	"fmt"	
)

func main() {	
	ropeCutting := RopeCutting(23, 11, 9, 12)
	fmt.Println("Num of Cuttings:", ropeCutting)
}

func RopeCutting(ropeLength, l1, l2, l3 int) int {
	if ropeLength == 0 {
		return 0
	}
	if ropeLength <= 0 {
		return -1
	}
	res := int(math.Max(float64(RopeCutting(ropeLength-l1, l1, l2, l3)), math.Max(float64(RopeCutting(ropeLength-l2, l1, l2, l3)), float64(RopeCutting(ropeLength-l3, l1, l2, l3)))))
	if res == -1 {
		return -1
	}
	return res + 1
}
