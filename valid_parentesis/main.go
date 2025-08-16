package main

import "fmt"

func main() {
	// https://leetcode.com/problems/valid-parentheses
	fmt.Println(isValid("()"))
}

func isValid(s string) bool {
	openingMap := map[string]string{"(": ")", "[": "]", "{": "}"}
	closingMap := map[string]string{")": "(", "]": "[", "}": "{"}
	st := newStack()
	for _, c := range s {
		str := string(c)
		if _, isOpening := openingMap[str]; isOpening {
			st.Push(str)
			continue
		}
		if openingTag, isClosing := closingMap[str]; isClosing {
			if st.Top() != openingTag {
				return false
			}
			st.Pop()
		}
	}
	return st.IsEmpty()
}

type stack struct {
	items []string
}

func newStack() *stack {
	return &stack{items: make([]string, 0)}
}

func (s *stack) Push(item string) {
	s.items = append(s.items, item)
}

func (s *stack) Pop() string {
	var zero string
	total := len(s.items)
	if total == 0 {
		return zero
	}
	if total == 1 {
		item := s.items[0]
		s.items = make([]string, 0)
		return item
	}
	item := s.items[total-1]
	s.items = s.items[:total-1]
	return item
}

func (s *stack) Top() string {
	var zero string
	total := len(s.items)
	if total == 0 {
		return zero
	}
	return s.items[total-1]
}

func (s *stack) IsEmpty() bool {
	return len(s.items) == 0
}
