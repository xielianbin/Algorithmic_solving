package main

import "fmt"

func mySqrt(x int) int {
	left, right := 0, x
	for left+1 < right {
		m := (left + right) / 2
		if m*m <= x {
			left = m
		} else {
			right = m
		}
	}
	return left
}
func main() {
	fmt.Println(mySqrt(4))
}
