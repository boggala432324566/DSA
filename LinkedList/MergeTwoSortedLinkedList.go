package main

import "fmt"

func (l *LinkedList) MergeLinkedList(head2 *Node) {
	head := l.merge(head2)
	fmt.Printf("Merge List:")
	for head != nil {
		fmt.Printf("%d  ", head.val)
		head = head.next
	}
}

func (l *LinkedList) merge(head2 *Node) *Node {
	l1 := &Node{}
	l2 := &Node{}
	head1 := l.head
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}
	if head1.val <= head2.val {
		l1 = head1
		l2 = head2
	} else {
		l1 = head2
		l2 = head1
	}
	newHead := l1
	temp := &Node{}
	for l1 != nil {
		if l1.val < l2.val {
			temp = l1
			l1 = l1.next
		} else {
			temp.next = l2
			l1, l2 = l2, l1
		}
	}
	temp.next = l2
	return newHead

}
