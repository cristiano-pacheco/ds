package main

import (
	"ds/stack"
	"fmt"
)

func main() {
	s := stack.New[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	fmt.Println(s.Peak())
}
