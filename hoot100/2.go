package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p1 := l1
	p2 := l2
	t := 0

	for p1 != nil {
		prep1 := p1
		tmp := 0
		if p2 == nil {
			tmp = p1.Val + t
		} else {
			tmp = p1.Val + p2.Val + t
		}

		p1.Val = tmp % 10
		t = tmp / 10
		p1 = p1.Next
		if p2 != nil {
			p2 = p2.Next
		}
		if p1 == nil && p2 != nil {
			tp := p1
			p1 = p2
			prep1.Next = p1
			p2 = tp
		}
		if p1 == nil && p2 == nil && t == 1 {
			tp := &ListNode{t, nil}
			prep1.Next = tp
			break
		}
	}
	return l1
}
func show(head *ListNode) {
	p := &ListNode{}
	p = head
	for p != nil {
		fmt.Print(p.Val)
		p = p.Next
	}
	fmt.Println()
}
func main() {
	a := []int{2, 4, 9}
	b := []int{5, 6, 4, 9}
	l1 := &ListNode{a[0], nil}
	l2 := &ListNode{b[0], nil}
	q := &ListNode{}

	q = l1
	for i := 1; i < len(a); i++ {
		node := &ListNode{a[i], nil}
		q.Next = node
		q = q.Next
	}
	q = l2
	for i := 1; i < len(b); i++ {
		node := &ListNode{b[i], nil}
		q.Next = node
		q = q.Next
	}

	show(l1)
	addTwoNumbers(l1, l2)
	show(l1)
}
