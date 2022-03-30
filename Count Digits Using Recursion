package main

import "fmt"

func main() {	
	fmt.Println("Digit Count Rec:", DigitCountRec(1234))
}
func DigitCountRec(num int) int {
	if num == 0 {
		return 0
	}
	return 1 + DigitCountRec(num/10)
}
