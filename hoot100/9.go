package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	xl1 := []int{}
	for x != 0 {
		xl1 = append(xl1, x%10)
		x = x / 10
	}
	lengthXl1 := len(xl1)
	for i := 0; i < lengthXl1/2; i++ {
		if xl1[i] != xl1[lengthXl1-i-1] {
			return false
		}
	}
	return true
}
func main() {
	x := 121
	fmt.Println(isPalindrome(x))
}
