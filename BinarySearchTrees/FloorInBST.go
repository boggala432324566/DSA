package main

func (t *BST) Floor(ceilNode *BSTNode) {
	ceil := BSTNode{}
	r := t.root
	for r != nil {
		if r.val == ceilNode.val {
			ceil.val = r.val
			break
		}
		if ceilNode.val < r.val {
			r = r.left
		} else {
			ceil.val = r.val
			r = r.right
		}
	}
}
