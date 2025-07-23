package main

import (
	"fmt"
)

func longestConsecutive(nums []int) int {

	if len(nums) < 2 {
		return len(nums)
	}
	res := make(map[int]bool)
	for _, n := range nums {
		res[n] = true
	}
	maxn := 1
	for n := range res {
		if res[n-1] {
			continue
		}
		y := n + 1
		for res[y] {
			y++
		}
		maxn = max(maxn, y-n)
	}

	return maxn
}
func main() {
	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	//nums := []int{1, 0, 1, 2}
	//nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums))
}
