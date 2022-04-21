package main

import (
	"fmt"
)

func main() {
	SubSets("abc", "", 0)
}

func SubSets(str string, current string, i int) {
	if i == len(str) {
		fmt.Printf("Sub String:%s\n", current)
		return
	}
	SubSets(str, current, i+1)
	SubSets(str, current+string(str[i]), i+1)
}
