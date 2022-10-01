package main

import "fmt"

type BSTNode struct {
	val   int
	left  *BSTNode
	right *BSTNode
}

type BST struct {
	root *BSTNode
}

func (t *BST) GetNode(value int) *BSTNode {
	node := new(BSTNode)
	node.val = value
	node.left = nil
	node.right = nil
	return node
}

func (t *BST) InsertNode(value int) {
	newNode := t.GetNode(value)
	if t.root == nil {
		t.root = newNode
		return
	}
	prev := &BSTNode{}
	temp := t.root

	for temp != nil {
		if value < temp.val {
			prev = temp
			temp = temp.left
		} else if value > temp.val {
			prev = temp
			temp = temp.right
		}
	}

	if value < prev.val {
		prev.left = newNode
	} else {
		prev.right = newNode
	}

}

func (t *BST) InorderTraversalRecursion(root *BSTNode) {
	if root == nil {
		return
	}
	t.InorderTraversalRecursion(root.left)
	fmt.Printf("%d ", root.val)
	t.InorderTraversalRecursion(root.right)
}

func main() {
	bst := BST{}
	bst.InsertNode(15)
	bst.InsertNode(10)
	bst.InsertNode(5)
	bst.InsertNode(8)
	bst.InsertNode(25)
	bst.InsertNode(20)
	bst.InsertNode(28)
	bst.InorderTraversalRecursion(bst.root)
}
