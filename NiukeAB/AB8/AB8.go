package main

import "fmt"

type CirQue struct {
	items []string
}

func (c *CirQue) Push(item string) {
	c.items = append(c.items, item)
}
func (c *CirQue) Pop() string {
	if c.Size() == 0 {
		return "empty"
	}
	item := c.items[0]
	c.items = c.items[1:]
	return item
}
func (c *CirQue) Size() int {
	return len(c.items)
}
func (c *CirQue) Front() string {
	if c.Size() == 0 {
		return "empty"
	}
	return c.items[0]
}
func main() {
	var cq CirQue
	var lgtn int
	var optn int
	fmt.Scan(&lgtn, &optn)
	for i := 0; i < optn; i++ {
		var op string
		fmt.Scan(&op)
		switch op {
		case "push":
			var x string
			fmt.Scan(&x)
			if cq.Size() > lgtn-1 {
				fmt.Println("full")
			} else {
				cq.Push(x)
			}
		case "pop":
			fmt.Println(cq.Pop())
		case "front":
			fmt.Println(cq.Front())
		}
	}
}
