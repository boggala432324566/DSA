package main

func (l *LinkedList) RemoveDuplicatesInSortedLinkedList() {
	pointer := l.head
	for pointer != nil && pointer.next != nil {
		for pointer.next != nil && pointer.val == pointer.next.val {
			pointer.next = pointer.next.next
		}
		pointer = pointer.next
	}
}
