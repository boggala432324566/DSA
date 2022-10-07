package main

type Stack struct {
	stack []string
}

func (s *Stack) push(ele string) {
	s.stack = append(s.stack, ele)
}
func (s *Stack) pop() string {
	ele := s.stack[len(s.stack)-1]
	s.stack = s.stack[0 : len(s.stack)-1]
	return ele
}

func (s *Stack) peek() string {
	return s.stack[len(s.stack)-1]
}

func (s *Stack) size() int {
	return len(s.stack)
}

func (s *Stack) isEmpty() bool {
	return len(s.stack) == 0
}
