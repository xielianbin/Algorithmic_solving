package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	cur := head
	nxt := head
	var pre *ListNode = nil
	for nxt != nil {
		nxt = nxt.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}
func main() {
	a := []int{1, 2, 3, 4, 5}
	head := &ListNode{Val: a[0], Next: nil}
	p := head
	for i := 1; i < len(a); i++ {
		t := &ListNode{Val: a[i], Next: nil}
		p.Next = t
		p = p.Next
	}
	q := reverseList(head)
	for q != nil {
		fmt.Println(q.Val)
		q = q.Next
	}
}
