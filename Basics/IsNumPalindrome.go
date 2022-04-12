package main

import (
	"fmt"
)

func main() {	
	IsPalindrome(339)
}

func IsPalindrome(num int) {
	temp := num
	rev := 0
	for temp != 0 {
		n := temp % 10
		rev = rev*10 + n
		temp = temp / 10
	}
	if num == rev {
		fmt.Println(num, "is Palindrome")
	} else {
		fmt.Println(num, "is Not Palindrome")
	}
}
