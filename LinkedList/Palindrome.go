package main

import "fmt"

func (l *LinkedList) palindrome() {
	prev := &Node{}
	slow := l.head
	fast := l.head
	for fast != nil && fast.next != nil {
		prev = slow
		fast = fast.next.next
		slow = slow.next
	}
	temp := &Node{}
	if fast != nil {
		temp = slow.next
	}
	prev.next = nil
	secondHead := l.reverseLinkedList(temp)
	palindrome := l.compare(l.head, secondHead)
	if palindrome {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not Palindrome")
	}
}

func (l *LinkedList) reverseLinkedList(head *Node) *Node {
	prev := &Node{}
	prev = nil
	curr := head
	nxt := head.next

	for nxt != nil {
		curr.next = prev
		prev = curr
		curr = nxt
		nxt = nxt.next
	}
	curr.next = prev
	return curr
}
func (l *LinkedList) compare(head1 *Node, head2 *Node) bool {
	for head1 != nil && head2 != nil {
		if head1.val != head2.val {
			return false
		} else {
			head1 = head1.next
			head2 = head2.next
		}
	}
	return head1 == nil && head2 == nil
}
