package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pa := &ListNode{}
	pb := &ListNode{}
	pa = headA
	pb = headB
	for pa != pb {
		pa = pa.Next
		pb = pb.Next
		if pa == nil {
			pa = headB
		}
		if pb == nil {
			pb = headA
		}
	}
	return pa
}
func show(head *ListNode) {
	p := &ListNode{}
	p = head
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}
}
func main() {
	a := []int{4, 1}
	b := []int{5, 6, 1}
	com := []int{8, 4, 5}
	headA := &ListNode{a[0], nil}
	headB := &ListNode{b[0], nil}
	conL := &ListNode{com[0], nil}

	q := &ListNode{}

	q = conL
	for i := 1; i < len(com); i++ {
		node := &ListNode{com[i], nil}
		q.Next = node
		q = q.Next
	}
	q = headA
	for i := 1; i < len(a); i++ {
		node := &ListNode{a[i], nil}
		q.Next = node
		q = q.Next
	}
	q = headB
	for i := 1; i < len(b); i++ {
		node := &ListNode{b[i], nil}
		q.Next = node
		q = q.Next
	}

	res := &ListNode{}
	res = getIntersectionNode(headA, headB)
	fmt.Println(res.Val)
	//show(headA)
	//show(headB)
}
