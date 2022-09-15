package main

import "fmt"

func (l *LinkedList) CycleStartNode() {
	slow := l.head
	fast := l.head
	cycle := false
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			cycle = true
			slow = l.head
			for slow != fast {
				slow = slow.next
				fast = fast.next
			}
			fmt.Printf("Cycle Start At:%d", slow.val)
			break
		}
	}
	if !cycle {
		fmt.Println("Cycle Not Found")
	}
}
