package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k < 2 {
		return head
	}
	// 获取链表长度
	p := head
	listLen := 0
	for p != nil {
		listLen++
		p = p.Next
	}
	groups := listLen / k         //进行groups组翻转
	tmpHead := &ListNode{Val: -1} //临时虚拟头节点
	now := tmpHead

	// 2.进行groups组的反转
	// 2.1 定义每组连接节点反转的指针变量
	cur := head //每组反转（操作的）当前节点
	nxt := head //每组反转（操作的）下一个节点

	var pre *ListNode //每组反转（操作的）前序节点
	for i := 0; i < groups; i++ {
		// 2.2 对每一组k个节点进行翻转进行翻转
		for j := 0; j < k; j++ {
			nxt = cur.Next
			cur.Next = pre
			pre = cur
			cur = nxt
		}
		//2.3 翻转之后，进行衔接
		now.Next = pre
		for now.Next != nil {
			now = now.Next
		}
		// 2.4 每一组K个节点反转完成后，使反转（操作的）前序节点指向空(Null)
		pre = nil
	}
	//3. 衔接剩余的节点
	now.Next = cur
	return tmpHead.Next
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
	q := reverseKGroup(head, 2)
	for q != nil {
		fmt.Println(q.Val)
		q = q.Next
	}
}
