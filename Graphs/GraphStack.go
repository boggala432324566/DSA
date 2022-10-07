package main

type GraphStack struct {
	stack []*Vertex
}

func (s *GraphStack) push(ele *Vertex) {
	s.stack = append(s.stack, ele)
}
func (s *GraphStack) pop() *Vertex {
	ele := s.stack[len(s.stack)-1]
	s.stack = s.stack[0 : len(s.stack)-1]
	return ele
}

func (s *GraphStack) peek() *Vertex {
	return s.stack[len(s.stack)-1]
}

func (s *GraphStack) size() int {
	return len(s.stack)
}

func (s *GraphStack) isEmpty() bool {
	return len(s.stack) == 0
}
