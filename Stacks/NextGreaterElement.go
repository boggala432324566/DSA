package main

import "fmt"

var stack []int

func main() {
	var arr = []int{2, 4, 8, 3, 20}
	nextGreaterEle(arr)

}

func nextGreaterEle(arr []int) {
	var temp [100]int
	for i, val := range arr {
		for !isEmpty() && arr[peek()] < val {
			temp[peek()] = val
			pop()
		}
		push(i)
	}

	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d-->%d\n", i, temp[i])
	}
}

func push(ele int) {
	stack = append(stack, ele)
}
func pop() int {
	ele := stack[len(stack)-1]
	stack = stack[0 : len(stack)-1]
	return ele
}

func isEmpty() bool {
	return len(stack) == 0
}

func peek() int {
	return stack[len(stack)-1]
}
