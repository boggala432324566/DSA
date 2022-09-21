package main

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

type BinaryTree struct {
	root  *TreeNode
	queue Queue
}

func (t *BinaryTree) GetNode(value int) *TreeNode {
	node := new(TreeNode)
	node.val = value
	node.left = nil
	node.right = nil
	return node
}

func main() {
	q := Queue{}
	binaryTree := &BinaryTree{queue: q}
	binaryTree.root = binaryTree.GetNode(6)
	binaryTree.root.left = binaryTree.GetNode(10)
	binaryTree.root.right = binaryTree.GetNode(15)
	binaryTree.root.left.left = binaryTree.GetNode(20)
	binaryTree.root.left.right = binaryTree.GetNode(21)
	binaryTree.root.right.left = binaryTree.GetNode(34)
	binaryTree.LevelOrderTraversal()
}
