package main

import "fmt"
type Quene struct{
	 s1 []string
	 s2 []string
}
func (q *Quene) Push(x string) {
	q.s1=append(q.s1,x)
}
func (q *Quene) Pop() (x string) {
	if len(q.s2)>0{
		x=q.s2[len(q.s2)-1]
		q.s2=q.s2[:len(q.s2)-1]
		return 
	}
	if len(q.s1)<1{
		return "error"
	}
	x=q.s1[0]
	q.s1=q.s1[1:]
	for i:=len(q.s1)-1;i>-1;i--{
		q.s2=append(q.s2,q.s1[i])
	}
	q.s1=q.s1[:0]
	return 
}
func (q *Quene) Front() (x string) {
	if len(q.s2)>0{
		return q.s2[len(q.s2)-1]
	}
	if len(q.s1)<1{
		return "error"
	}
	return q.s1[0]
}
func main() {
	var n int
	fmt.Scan(&n)
	var q Quene
	for i := 0; i < n; i++ {
		var op, x string
		fmt.Scan(&op)
		if op == "push" {
			fmt.Scan(&x)
			q.Push(x)
		} else if op == "pop" {
			x = q.Pop()
			fmt.Println(x)
		} else if op == "front" {
			x = q.Front()
			fmt.Println(x)
		}
	}		
}