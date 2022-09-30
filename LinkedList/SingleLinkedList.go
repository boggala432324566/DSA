package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) GetNode(value int) *Node {
	node := new(Node)
	node.val = value
	node.next = nil
	return node
}
func (l *LinkedList) AddNode(value int) {
	node := l.GetNode(value)
	if l.head == nil {
		l.head = node
	} else {
		temp := l.head
		for ; temp.next != nil; temp = temp.next {

		}
		temp.next = node
	}
}
func (l *LinkedList) PrintLinkedList() {
	temp := l.head
	fmt.Printf("List Elements:")
	for ; temp != nil; temp = temp.next {
		fmt.Printf("%d  ", temp.val)
	}
}

func main() {
	linkedList := &LinkedList{}
	//linkedList.AddNode(6)
	linkedList.AddNode(10)
	linkedList.AddNode(20)
	linkedList.AddNode(30)
	linkedList.AddNode(21)
	linkedList.AddNode(50)
	linkedList.AddNode(70)
	linkedList.AddNode(90)
	linkedList.AddNode(61)
	//linkedList.head.next.next.next.next = linkedList.head.next
	//linkedList.DeleteGivenNode(linkedList.head.next.next)
	//linkedList.FindMiddleElement()
	//linkedList.DeleteGivenNodeFromEnd(2)
	//linkedList.ReverseLinkedList()
	//linkedList.PrintLinkedList()
	//linkedList.DetectCycle()
	//linkedList.CycleLength()
	//linkedList.palindrome()
	//node := linkedList.GetNode(4)
	//node.next = linkedList.GetNode(13)
	//node.next.next = linkedList.GetNode(50)
	//linkedList.MergeLinkedList(nil)
	//linkedList.AlternateLinkedListNodes()
	//linkedList.RemoveDuplicatesInSortedLinkedList()
	linkedList.reverseByGroupOfK(linkedList.head, 8, 3)
	linkedList.PrintLinkedList()
}
