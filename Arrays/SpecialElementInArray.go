package main

import "fmt"

func main() {
	var arr = []int{5, 5, 2, 5, 8}
	spCount := SpecialElementInArray(arr)
	fmt.Printf("Special Element Count:%d", spCount)
}
func SpecialElementInArray(arr []int) int {
	specialElementCount := 0
	totalEvenSum, totalOddSum, currentEvenSum, currentOddSum, newEvenSum, newOddSum := 0, 0, 0, 0, 0, 0

	for i := 0; i < len(arr); i++ {
		if i%2 == 0 {
			totalEvenSum = totalEvenSum + arr[i]
		} else {
			totalOddSum = totalOddSum + arr[i]
		}
	}
	for i := 0; i < len(arr); i++ {
		if i%2 == 0 {
			currentEvenSum = currentEvenSum + arr[i]
			newEvenSum = currentEvenSum + (totalOddSum - currentOddSum - arr[i])
			newOddSum = currentOddSum + (totalEvenSum - currentEvenSum)
		} else {
			currentOddSum = currentOddSum + arr[i]
			newEvenSum = currentEvenSum + (totalOddSum - currentOddSum)
			newOddSum = currentOddSum + (totalEvenSum - currentEvenSum - arr[i])
		}
		if newEvenSum == newOddSum {
			specialElementCount++
		}
	}
	return specialElementCount

}
