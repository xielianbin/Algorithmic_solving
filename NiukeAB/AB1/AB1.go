package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	var reader = bufio.NewReader(os.Stdin)
	num, _ := reader.ReadString('\n')
	numInt, _ := strconv.Atoi(num[0:1])
	// fmt.Println(len(num))
	// for i, c := range num {
	// 	fmt.Printf("%d:%d  ", i, int(c))
	// }

	// fmt.Scanf(&num)
	// fmt.Printf(num)
	for i := 0; i < numInt; i++ {
		var a string
		a, _ = reader.ReadString('\n')

		if len(a) > 3 {
			if a[0:3] == "pop" {
				fmt.Println(s.Pop())
			} else if a[0:3] == "top" {
				fmt.Println(s.Top())
			}
		}
		if len(a) > 4 {
			if a[0:4] == "push" {
				// fmt.Println(a[len(a)-3 : len(a)-2])
				s.Push(a[len(a)-3 : len(a)-2])
			}
		}
	}
}
