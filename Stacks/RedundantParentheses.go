package main

import "fmt"

func (s *Stack) redundantParentheses() {
	input := "(((a+b)+c))"
	isRedundant := s.checkRedundantParentheses(input)
	fmt.Println(isRedundant)

}

func (s *Stack) checkRedundantParentheses(input string) bool {
	inputArr := []rune(input)
	op := 0
	for _, c := range inputArr {
		sc := string(c)
		if sc == "{" || sc == "[" || sc == "(" {
			s.push(sc)
		} else if sc == "+" || sc == "-" || sc == "*" || sc == "/" {
			op = op + 1
		} else {
			if s.isEmpty() {
				return false
			}
			if !s.isEmpty() && sc == ")" || sc == "}" || sc == "]" {
				op = op - 1
				s.pop()
			}
		}
	}
	if s.isEmpty() && op == 0 {
		return true
	}
	return false
}
