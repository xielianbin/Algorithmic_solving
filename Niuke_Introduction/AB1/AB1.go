package main

import (
	"fmt"
)

type Stack struct {
	items []string
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) < 1
}
func (s *Stack) Size() int {
	return len(s.items)
}
func (s *Stack) Pop() string {
	if s.IsEmpty() {
		a := s.items[s.Size()-1]
		s.items = s.items[:s.Size()-1]
		return a
	} else {
		return "error"
	}
}
func (s *Stack) Push(a string) {
	s.items = append(s.items, a)
}
func main() {
	var s Stack
	var num int
	fmt.Scanf(&num)
	fmt.Printf(num)
	for i := 0; i < num; i++ {
		var a string
		fmt.Scanf(&a)
		fmt.Printf(a)
		if a[0:3] == "pop" {
			fmt.Printf(s.Pop())
		} else if a[0:4] == "push" {
			s.Push(a[len(a)-1:])
		}
	}
}
