package main

import "fmt"

func main() {
	str := "madame"
	flag := pal(str, 0, len(str)-1)
	fmt.Println(flag)
}

func pal(str string, l int, r int) bool {
	if l > r {
		return true
	}
	if str[l] != str[r] {
		return false
	}
	l++
	r--
	return pal(str, l, r)
}
