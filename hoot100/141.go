package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	p1, p2 := head, head
	for p2 != nil && p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
		if p1 == p2 {
			return true
		}
	}
	return false
}

func main() {
	head := &ListNode{}
	//arr := []int{1, 2}
	arr := []int{3, 2, 0, 1}
	head.Val = arr[0]
	p := head
	q := head
	for i := 1; i < len(arr); i++ {
		node := &ListNode{Val: arr[i], Next: nil}
		p.Next = node
		if i == 1 {
			q = node
		}
		if i == 3 {
			p.Next = q
		}
		p = p.Next
	}
	//showList(head)
	fmt.Println(hasCycle(head))

}
