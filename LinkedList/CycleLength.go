package main

import "fmt"

func (l *LinkedList) CycleLength() {
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
			temp := slow.next
			cycleLength := 1
			for temp != slow {
				cycleLength++
				temp = temp.next
			}
			fmt.Printf("Cycle Found With Length:%d", cycleLength)
			break
		}
	}
	if !cycle {
		fmt.Println("Cycle Not Found")
	}
}
