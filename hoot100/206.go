package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 采用递归
func reverseList(head *ListNode) *ListNode {
	var pre, cur *ListNode = nil, head
	for cur != nil {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}
func showList(head *ListNode) {
	p := head
	for p != nil {
		fmt.Printf("%d\t", p.Val)
		p = p.Next
	}
	fmt.Println("\n")
}
func main() {
	head := &ListNode{}
	//arr := []int{1, 2}
	arr := []int{1, 2, 3, 4, 5}
	head.Val = arr[0]
	p := head
	for i := 1; i < len(arr); i++ {
		node := &ListNode{Val: arr[i], Next: nil}
		p.Next = node
		p = p.Next
	}
	showList(head)
	showList(reverseList(head))

}
