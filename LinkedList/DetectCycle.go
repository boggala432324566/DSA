package main

import "fmt"

func (l *LinkedList) DetectCycle() {
	slow := l.head
	fast := l.head
	cycle := false
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			fmt.Println("Cycle Found")
			cycle = true
			break
		}
	}
	if !cycle {
		fmt.Println("Cycle Not Found")
	}
}
