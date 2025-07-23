package main

import (
	"fmt"
)

func maxArea(height []int) int {
	n := len(height)
	lp := 0
	rp := n - 1
	maxn := 0
	for lp < rp {
		maxn = max(min(height[lp], height[rp])*(rp-lp), maxn)
		if height[lp] > height[rp] {
			rp--
		} else {
			lp++
		}
	}
	return maxn
}
func main() {
	//nums := []int{0, 1, 0, 3, 12}
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	//nums := []int{100, 4, 200, 1, 3, 2}

	fmt.Println(maxArea(nums))
}
