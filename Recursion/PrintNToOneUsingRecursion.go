package main

import (
	"fmt"
)

func main() {
	PrintNToOneUsingRecursion(10)
}

func PrintNToOneUsingRecursion(n int) {
	if n == 0 {
		return
	}
	fmt.Print(n)
	PrintNToOneUsingRecursion(n - 1)

}
