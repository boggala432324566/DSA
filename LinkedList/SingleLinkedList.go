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
		fmt.Printf("%d", temp.val)
	}
}

func main() {
	linkedList := &LinkedList{}
	linkedList.AddNode(1)
	linkedList.AddNode(2)
	linkedList.AddNode(3)
	linkedList.AddNode(4)
	linkedList.AddNode(5)
	//linkedList.DeleteGivenNode(linkedList.head.next.next)
	//linkedList.FindMiddleElement()
	//linkedList.DeleteGivenNodeFromEnd(2)
	linkedList.ReverseLinkedList()
	linkedList.PrintLinkedList()
}
