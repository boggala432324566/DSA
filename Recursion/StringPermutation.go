package main

import "fmt"

func main() {
	Permute("abc", 0)
}

func Permute(str string, i int) {
	if i == len(str)-1 {
		fmt.Printf("%s\n", str)
	}
	for j := i; j < len(str); j++ {
		s1 := []rune(str)
		s1[i], s1[j] = s1[j], s1[i]
		Permute(string(s1), i+1)
		s2 := []rune(str)
		s2[j], s2[i] = s2[i], s2[j]
	}
}
