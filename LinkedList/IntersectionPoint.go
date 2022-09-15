package main

import "fmt"

func (l *LinkedList) IntersectionPoint(head2 *Node) {
	l1 := 0
	l2 := 0
	t1 := l.head
	t2 := head2
	for t1 != nil {
		l1++
		t1 = t1.next
	}
	for t2 != nil {
		l2++
		t2 = t2.next
	}
	diff := 0
	if l1 > l2 {
		diff = l1 - l2
		for i := 0; i < diff; i++ {
			l.head = l.head.next
		}
	} else {
		diff = l2 - l1
		for i := 0; i < diff; i++ {
			head2 = head2.next
		}
	}
	for l.head != head2 {
		l.head = l.head.next
		head2 = head2.next
	}
	fmt.Printf("Intersection Point:%d", head2.val)
}
