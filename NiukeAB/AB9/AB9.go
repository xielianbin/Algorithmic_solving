package main

import "fmt"

type LinkList struct {
	items []int
}

func (l *LinkList) Insert(x, y int) {
	for i := 0; i < l.Size(); i++ {
		if l.items[i] == x {
			l.items = append(l.items[:i+1], l.items[i:]...)
			l.items[i] = y
			return
		}
	}
	l.items = append(l.items, y)
}
func (l *LinkList) Size() (lgt int) {
	return len(l.items)
}
func (l *LinkList) Delete(x int) {
	for i := 0; i < l.Size(); i++ {
		if l.items[i] == x {
			l.items = append(l.items[:i], l.items[i+1:]...)
			return
		}
	}
}
func (l *LinkList) Print() {
	for i := 0; i < l.Size(); i++ {
		fmt.Print(l.items[i], " ")
	}
}
func main() {
	var ll LinkList
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var op string
		fmt.Scan(&op)
		switch op {
		case "insert":
			var x, y int
			fmt.Scan(&x, &y)
			ll.Insert(x, y)
		case "delete":
			var x int
			fmt.Scan(&x)
			ll.Delete(x)
		}
	}
	if ll.Size() == 0 {
		fmt.Println("NULL")
	} else {
		ll.Print()
	}
}
