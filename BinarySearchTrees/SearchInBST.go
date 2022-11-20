package main

import "fmt"

func (t *BST) Search(searchNode *BSTNode) {
	searchFlag := false
	r := t.root
	for r != nil {
		if r.val == searchNode.val {
			fmt.Println("Node Found")
			searchFlag = true
			break
		}
		if searchNode.val < r.val {
			r = r.left
		} else {
			r = r.right
		}
	}
	if !searchFlag {
		fmt.Println("Node Not Found")
	}
}
