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
	if !s.IsEmpty() {
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
func (s *Stack) Top() string {
	if !s.IsEmpty() {
		return s.items[s.Size()-1]
	} else {
		return "error"
	}
}
func main() {
	var s Stack
	var num int
	fmt.Scan(&num)
	// numInt, _ := strconv.Atoi(num[0:1])
	// fmt.Println(len(num))
	// for i, c := range num {
	// 	fmt.Printf("%d:%d  ", i, int(c))
	// }

	// fmt.Scanf(&num)
	// fmt.Printf(num)
	for i := 0; i < num; i++ {
		var a string
		fmt.Scan(&a)
		if a == "pop" {
			fmt.Println(s.Pop())
		} else if a == "top" {
			fmt.Println(s.Top())
		} else if a == "push" {
			var b string
			fmt.Scan(&b)
			// fmt.Println(a[len(a)-3 : len(a)-2])
			s.Push(b)
		}
	}
}
