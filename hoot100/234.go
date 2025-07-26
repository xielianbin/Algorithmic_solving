package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func showList(head *ListNode) {
	p := head
	for p != nil {
		fmt.Printf("%d\t", p.Val)
		p = p.Next
	}
	fmt.Println("\n")
}
func isPalindrome(head *ListNode) bool {
	vals := []int{}
	p := head
	for p != nil {
		vals = append(vals, p.Val)
		p = p.Next
	}
	n := len(vals)
	for i := 0; i < n/2; i++ {
		if vals[i] != vals[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
	head := &ListNode{}
	arr := []int{1, 2}
	//arr := []int{1, 2, 2, 1}
	head.Val = arr[0]
	p := head
	for i := 1; i < len(arr); i++ {
		node := &ListNode{Val: arr[i], Next: nil}
		p.Next = node
		p = p.Next
	}
	showList(head)
	fmt.Println(isPalindrome(head))

}
