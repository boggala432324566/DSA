package main

type Stack struct {
	stack []*TreeNode
}

func (s *Stack) Push(node *TreeNode) {
	s.stack = append(s.stack, node)
}

func (s *Stack) Pop() *TreeNode {
	tNode := s.stack[len(s.stack)-1]
	s.stack = s.stack[0 : len(s.stack)-1]
	return tNode
}

func (s *Stack) Peek() *TreeNode {
	return s.stack[len(s.stack)-1]

}

func (s *Stack) Size() int {
	return len(s.stack)
}

func (s *Stack) isEmpty() bool {
	return len(s.stack) == 0
}
