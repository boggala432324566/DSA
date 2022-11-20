package main

func (t *BST) Ceil(ceilNode *BSTNode) {
	ceil := BSTNode{}
	r := t.root
	for r != nil {
		if r.val == ceilNode.val {
			ceil.val = r.val
			break
		}
		if ceilNode.val < r.val {
			ceil.val = r.val
			r = r.left
		} else {
			r = r.right
		}
	}
}
