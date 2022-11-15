package main

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

type BinaryTree struct {
	root  *TreeNode
	queue Queue
	stack Stack
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
	//binaryTree.LevelOrderTraversal()
	//binaryTree.PreorderTraversalRecursion(binaryTree.root)
	//binaryTree.InorderTraversalRecursion(binaryTree.root)
	//binaryTree.PostorderTraversalRecursion(binaryTree.root)
	//binaryTree.PreorderTraversalIterative(binaryTree.root)
	//binaryTree.InorderTraversalIterative(binaryTree.root)
	//binaryTree.PostorderTraversalIterative(binaryTree.root)
	//binaryTree.VerticalOrderTraversal(binaryTree.root)
	//binaryTree.TreeHeight(binaryTree.root)
	//binaryTree.CheckBalancedBinaryTree(binaryTree.root)
	//binaryTree.TreeDiameter(binaryTree.root)
	//binaryTree2 := &BinaryTree{queue: q}
	//binaryTree2.root = binaryTree2.GetNode(6)
	//binaryTree2.root.left = binaryTree2.GetNode(10)
	//binaryTree2.root.right = binaryTree2.GetNode(15)
	//binaryTree2.root.left.left = binaryTree2.GetNode(20)
	//binaryTree2.root.left.right = binaryTree2.GetNode(21)
	//binaryTree2.root.right.left = binaryTree2.GetNode(34)
	//binaryTree.TreesAreIdentical(binaryTree.root, binaryTree2.root)
	//binaryTree.TreeTopView(binaryTree.root)
	//binaryTree.TreeButtomView(binaryTree.root)
	//binaryTree.TreeRightView(binaryTree.root)
	//binaryTree.TreeLeftView(binaryTree.root)
	binaryTree.BoundaryTraversal(binaryTree.root)

}
