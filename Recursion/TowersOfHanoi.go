package main

import (
	"fmt"
)

func main() {
	TOH(3, 'a', 'b', 'c')
}

func TOH(n int, a rune, b rune, c rune) {
	if n == 1 {
		fmt.Printf("Move 1  From :%c to %c\n", a, c)
		return
	}
	TOH(n-1, a, c, b)
	fmt.Printf("Move %d From :%c to %c\n", n, a, c)
	TOH(n-1, b, a, c)
}
