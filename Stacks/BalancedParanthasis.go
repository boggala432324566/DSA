package main

import "fmt"

func (s *Stack) balancedParanthasis() {
	input := "{()[]}"
	isBalaced := s.checkBalacedOrNot(input)
	fmt.Println(isBalaced)

}

func (s *Stack) checkBalacedOrNot(input string) bool {
	inputArr := []rune(input)
	for _, c := range inputArr {
		sc := string(c)
		if sc == "{" || sc == "[" || sc == "(" {
			s.push(sc)
		} else {
			if s.isEmpty() {
				return false
			}
			if !s.isEmpty() && sc == ")" || sc == "}" || sc == "]" {
				s.pop()
			}
		}
	}
	if s.isEmpty() {
		return true
	}
	return false
}
