package main

import (
	"fmt"	
)

func main() {
	str := "abbbba"
	isPal := isPalindrome(str, 0, len(str)-1)
	fmt.Printf("%v Is Palindrome:%v", str, isPal)
}

func isPalindrome(str string, start, end int) bool {
	if start >= end {
		return true
	}
	return (str[start] == str[end]) && isPalindrome(str, start+1, end-1)
}
