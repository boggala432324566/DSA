package main

import (
	"fmt"
)

func (l *LinkedList) FindMiddleElement() {
	slow := l.head
	fast := l.head.next
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
	}
	fmt.Printf("Middle Element:%d", slow.val)
}
